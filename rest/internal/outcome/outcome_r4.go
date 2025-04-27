//go:build r4 || !(r4 || r4b || r5)

package outcome

import (
	"errors"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
)

func init() {
	toOOR4 = func(err error) (model.Resource, basic.OperationOutcome, bool) {
		var (
			oo r4.OperationOutcome
			ok bool
		)
		ok = errors.As(err, &oo)
		if !ok {
			return oo, basic.OperationOutcome{}, false
		}

		ooBasic := basic.OperationOutcome{}
		for _, e := range oo.Issue {
			issue := basic.OperationOutcomeIssue{}
			if e.Severity.Value != nil {
				issue.Severity = basic.Code{Value: e.Severity.Value}
			}
			if e.Code.Value != nil {
				issue.Code = basic.Code{Value: e.Code.Value}
			}
			if e.Diagnostics != nil && e.Diagnostics.Value != nil {
				issue.Diagnostics = &basic.String{Value: e.Diagnostics.Value}
			}
			ooBasic.Issue = append(ooBasic.Issue, issue)
		}

		return oo, ooBasic, true
	}
}
