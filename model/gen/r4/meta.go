package r4

import "encoding/json"

// Base StructureDefinition for Meta Type: The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The version specific identifier, as it appears in the version portion of the URL. This value changes when the resource is created, updated, or deleted.
	VersionId *Id
	// When the resource last changed - e.g. when the version changed.
	LastUpdated *Instant
	// A uri that identifies the source system of the resource. This provides a minimal amount of [Provenance](provenance.html#) information that can be used to track or differentiate the source of information in the resource. The source may identify another FHIR server, document, message, database, etc.
	Source *Uri
	// A list of profiles (references to [StructureDefinition](structuredefinition.html#) resources) that this resource claims to conform to. The URL is a reference to [StructureDefinition.url](structuredefinition-definitions.html#StructureDefinition.url).
	Profile []Canonical
	// Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Security []Coding
	// Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	Tag []Coding
}
type jsonMeta struct {
	Id                          *string             `json:"id,omitempty"`
	Extension                   []Extension         `json:"extension,omitempty"`
	VersionId                   *Id                 `json:"versionId,omitempty"`
	VersionIdPrimitiveElement   *primitiveElement   `json:"_versionId,omitempty"`
	LastUpdated                 *Instant            `json:"lastUpdated,omitempty"`
	LastUpdatedPrimitiveElement *primitiveElement   `json:"_lastUpdated,omitempty"`
	Source                      *Uri                `json:"source,omitempty"`
	SourcePrimitiveElement      *primitiveElement   `json:"_source,omitempty"`
	Profile                     []Canonical         `json:"profile,omitempty"`
	ProfilePrimitiveElement     []*primitiveElement `json:"_profile,omitempty"`
	Security                    []Coding            `json:"security,omitempty"`
	Tag                         []Coding            `json:"tag,omitempty"`
}

func (r Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Meta) marshalJSON() jsonMeta {
	m := jsonMeta{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.VersionId != nil && r.VersionId.Value != nil {
		m.VersionId = r.VersionId
	}
	if r.VersionId != nil && (r.VersionId.Id != nil || r.VersionId.Extension != nil) {
		m.VersionIdPrimitiveElement = &primitiveElement{Id: r.VersionId.Id, Extension: r.VersionId.Extension}
	}
	if r.LastUpdated != nil && r.LastUpdated.Value != nil {
		m.LastUpdated = r.LastUpdated
	}
	if r.LastUpdated != nil && (r.LastUpdated.Id != nil || r.LastUpdated.Extension != nil) {
		m.LastUpdatedPrimitiveElement = &primitiveElement{Id: r.LastUpdated.Id, Extension: r.LastUpdated.Extension}
	}
	if r.Source != nil && r.Source.Value != nil {
		m.Source = r.Source
	}
	if r.Source != nil && (r.Source.Id != nil || r.Source.Extension != nil) {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	anyProfileValue := false
	for _, e := range r.Profile {
		if e.Value != nil {
			anyProfileValue = true
			break
		}
	}
	if anyProfileValue {
		m.Profile = r.Profile
	}
	anyProfileIdOrExtension := false
	for _, e := range r.Profile {
		if e.Id != nil || e.Extension != nil {
			anyProfileIdOrExtension = true
			break
		}
	}
	if anyProfileIdOrExtension {
		m.ProfilePrimitiveElement = make([]*primitiveElement, 0, len(r.Profile))
		for _, e := range r.Profile {
			if e.Id != nil || e.Extension != nil {
				m.ProfilePrimitiveElement = append(m.ProfilePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ProfilePrimitiveElement = append(m.ProfilePrimitiveElement, nil)
			}
		}
	}
	m.Security = r.Security
	m.Tag = r.Tag
	return m
}
func (r *Meta) UnmarshalJSON(b []byte) error {
	var m jsonMeta
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Meta) unmarshalJSON(m jsonMeta) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.VersionId = m.VersionId
	if m.VersionIdPrimitiveElement != nil {
		if r.VersionId == nil {
			r.VersionId = &Id{}
		}
		r.VersionId.Id = m.VersionIdPrimitiveElement.Id
		r.VersionId.Extension = m.VersionIdPrimitiveElement.Extension
	}
	r.LastUpdated = m.LastUpdated
	if m.LastUpdatedPrimitiveElement != nil {
		if r.LastUpdated == nil {
			r.LastUpdated = &Instant{}
		}
		r.LastUpdated.Id = m.LastUpdatedPrimitiveElement.Id
		r.LastUpdated.Extension = m.LastUpdatedPrimitiveElement.Extension
	}
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		if r.Source == nil {
			r.Source = &Uri{}
		}
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	for i, e := range m.ProfilePrimitiveElement {
		if len(r.Profile) <= i {
			r.Profile = append(r.Profile, Canonical{})
		}
		if e != nil {
			r.Profile[i].Id = e.Id
			r.Profile[i].Extension = e.Extension
		}
	}
	r.Security = m.Security
	r.Tag = m.Tag
	return nil
}
func (r Meta) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
