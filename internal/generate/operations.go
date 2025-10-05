package generate

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

// Operations holds discovered operation codes per release.
type Operations struct {
	// System holds codes available at system-level (/$code)
	System []string
	// ByResource maps resource name -> op code -> level flags
	ByResource map[string]map[string]resOpLevels
}

type resOpLevels struct {
	HasType     bool
	HasInstance bool
}

var (
	trRowRe   = regexp.MustCompile(`<tr><td><a href="[^"]+">[^<]+</a></td><td>([^<]+)</td></tr>`) // captures RHS paths segment
	pathSplit = regexp.MustCompile(`\s*\|\s*`)
)

// FetchOperationsForRelease scrapes the FHIR operations list for the given release (e.g., "R4", "R4B", "R5").
func FetchOperationsForRelease(release string) (Operations, error) {
	url := fmt.Sprintf("https://hl7.org/fhir/%s/operationslist.html", strings.ToUpper(release))
	resp, err := http.Get(url)
	if err != nil {
		return Operations{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Operations{}, fmt.Errorf("unexpected status fetching operationslist: %s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Operations{}, err
	}
	html := string(b)

	// Extract all rows with paths
	matches := trRowRe.FindAllStringSubmatch(html, -1)
	if len(matches) == 0 {
		return Operations{}, fmt.Errorf("no operation rows found for %s", release)
	}

	ops := Operations{ByResource: map[string]map[string]resOpLevels{}}
	systemSet := map[string]struct{}{}

	for _, m := range matches {
		rhs := strings.TrimSpace(m[1])
		parts := pathSplit.Split(rhs, -1)
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			// Determine level and extract code + resource
			// Examples:
			// [base]/$meta
			// [base]/[Resource]/$validate
			// [base]/[Resource]/[id]/$validate
			// [base]/Patient/$everything
			// [base]/Patient/[id]/$everything
			// [base]/[CanonicalResource]/$current-canonical (R5)
			if !strings.HasPrefix(p, "[base]/") {
				continue
			}
			rest := strings.TrimPrefix(p, "[base]/")

			// system-level: begins with $code
			if strings.HasPrefix(rest, "$") {
				code := strings.TrimPrefix(rest, "$")
				addSystem(systemSet, code)
				continue
			}

			segs := strings.Split(rest, "/")
			if len(segs) < 2 {
				continue
			}

			// seg0 is resource indicator
			resSeg := segs[0]
			// skip placeholder filters we can't map precisely by treating as generic to all resources
			var resList []string
			switch resSeg {
			case "[Resource]", "[CanonicalResource]":
				// will be expanded later across known resources
				resList = nil
			default:
				// explicit resource name
				resList = []string{resSeg}
			}

			// code is always prefixed by $ in the last segment
			last := segs[len(segs)-1]
			if !strings.HasPrefix(last, "$") {
				continue
			}
			code := strings.TrimPrefix(last, "$")

			// determine level: presence of [id]
			hasInstance := false
			hasType := false
			for _, s := range segs {
				if s == "[id]" {
					hasInstance = true
				}
			}
			// If instance-only pattern, mark instance; if there is no [id], it's type-level
			if !hasInstance {
				hasType = true
			} else {
				// Some rows include both type and instance alternatives via |; we'll handle both occurrences separately
			}

			if resList == nil {
				// generic placeholder -> mark on special key "*" for later expansion
				addResOp(&ops, "*", code, hasType, hasInstance)
			} else {
				for _, r := range resList {
					addResOp(&ops, r, code, hasType, hasInstance)
				}
			}
		}
	}

	// Expand generic placeholders to all resources at generation time; the caller will filter
	// We keep them stored under "*" and the generator expands per actual resource.
	// Prepare system list sorted
	for k := range systemSet {
		ops.System = append(ops.System, k)
	}
	sort.Strings(ops.System)

	return ops, nil
}

func addSystem(set map[string]struct{}, code string) {
	if _, ok := set[code]; !ok {
		set[code] = struct{}{}
	}
}

func addResOp(ops *Operations, res, code string, hasType, hasInstance bool) {
	if ops.ByResource[res] == nil {
		ops.ByResource[res] = map[string]resOpLevels{}
	}
	v := ops.ByResource[res][code]
	v.HasType = v.HasType || hasType
	v.HasInstance = v.HasInstance || hasInstance
	ops.ByResource[res][code] = v
}
