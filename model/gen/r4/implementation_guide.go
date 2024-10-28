package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
//
// An implementation guide is able to define default profiles that must apply to any use of a resource, so validation services may need to take one or more implementation guide resources when validating.
type ImplementationGuide struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An absolute URI that is used to identify this implementation guide when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this implementation guide is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the implementation guide is stored on different servers.
	Url Uri
	// The identifier that is used to identify this version of the implementation guide when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the implementation guide author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the implementation guide. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// A short, descriptive, user-friendly title for the implementation guide.
	Title *String
	// The status of this implementation guide. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this implementation guide is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the implementation guide was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the implementation guide changes.
	Date *DateTime
	// The name of the organization or individual that published the implementation guide.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the implementation guide from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate implementation guide instances.
	UseContext []UsageContext
	// A legal or geographic region in which the implementation guide is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the implementation guide and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the implementation guide.
	Copyright *Markdown
	// The NPM package name for this Implementation Guide, used in the NPM package distribution, which is the primary mechanism by which FHIR based tooling manages IG dependencies. This value must be globally unique, and should be assigned with care.
	PackageId Id
	// The license that applies to this Implementation Guide, using an SPDX license code, or 'not-open-source'.
	License *Code
	// The version(s) of the FHIR specification that this ImplementationGuide targets - e.g. describes how to use. The value of this element is the formal version of the specification, without the revision number, e.g. [publication].[major].[minor], which is 4.0.1. for this version.
	FhirVersion []Code
	// Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
	DependsOn []ImplementationGuideDependsOn
	// A set of profiles that all resources covered by this implementation guide must conform to.
	Global []ImplementationGuideGlobal
	// The information needed by an IG publisher tool to publish the whole implementation guide.
	Definition *ImplementationGuideDefinition
	// Information about an assembled implementation guide, created by the publication tooling.
	Manifest *ImplementationGuideManifest
}

func (r ImplementationGuide) ResourceType() string {
	return "ImplementationGuide"
}
func (r ImplementationGuide) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonImplementationGuide struct {
	ResourceType                  string                         `json:"resourceType"`
	Id                            *Id                            `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement              `json:"_id,omitempty"`
	Meta                          *Meta                          `json:"meta,omitempty"`
	ImplicitRules                 *Uri                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement              `json:"_implicitRules,omitempty"`
	Language                      *Code                          `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement              `json:"_language,omitempty"`
	Text                          *Narrative                     `json:"text,omitempty"`
	Contained                     []ContainedResource            `json:"contained,omitempty"`
	Extension                     []Extension                    `json:"extension,omitempty"`
	ModifierExtension             []Extension                    `json:"modifierExtension,omitempty"`
	Url                           Uri                            `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement              `json:"_url,omitempty"`
	Version                       *String                        `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement              `json:"_version,omitempty"`
	Name                          String                         `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement              `json:"_name,omitempty"`
	Title                         *String                        `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement              `json:"_title,omitempty"`
	Status                        Code                           `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement              `json:"_status,omitempty"`
	Experimental                  *Boolean                       `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement              `json:"_experimental,omitempty"`
	Date                          *DateTime                      `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement              `json:"_date,omitempty"`
	Publisher                     *String                        `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement              `json:"_publisher,omitempty"`
	Contact                       []ContactDetail                `json:"contact,omitempty"`
	Description                   *Markdown                      `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement              `json:"_description,omitempty"`
	UseContext                    []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept              `json:"jurisdiction,omitempty"`
	Copyright                     *Markdown                      `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement              `json:"_copyright,omitempty"`
	PackageId                     Id                             `json:"packageId,omitempty"`
	PackageIdPrimitiveElement     *primitiveElement              `json:"_packageId,omitempty"`
	License                       *Code                          `json:"license,omitempty"`
	LicensePrimitiveElement       *primitiveElement              `json:"_license,omitempty"`
	FhirVersion                   []Code                         `json:"fhirVersion,omitempty"`
	FhirVersionPrimitiveElement   []*primitiveElement            `json:"_fhirVersion,omitempty"`
	DependsOn                     []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`
	Global                        []ImplementationGuideGlobal    `json:"global,omitempty"`
	Definition                    *ImplementationGuideDefinition `json:"definition,omitempty"`
	Manifest                      *ImplementationGuideManifest   `json:"manifest,omitempty"`
}

func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuide) marshalJSON() jsonImplementationGuide {
	m := jsonImplementationGuide{}
	m.ResourceType = "ImplementationGuide"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Experimental != nil && r.Experimental.Value != nil {
		m.Experimental = r.Experimental
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		m.Publisher = r.Publisher
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	if r.PackageId.Value != nil {
		m.PackageId = r.PackageId
	}
	if r.PackageId.Id != nil || r.PackageId.Extension != nil {
		m.PackageIdPrimitiveElement = &primitiveElement{Id: r.PackageId.Id, Extension: r.PackageId.Extension}
	}
	if r.License != nil && r.License.Value != nil {
		m.License = r.License
	}
	if r.License != nil && (r.License.Id != nil || r.License.Extension != nil) {
		m.LicensePrimitiveElement = &primitiveElement{Id: r.License.Id, Extension: r.License.Extension}
	}
	anyFhirVersionValue := false
	for _, e := range r.FhirVersion {
		if e.Value != nil {
			anyFhirVersionValue = true
			break
		}
	}
	if anyFhirVersionValue {
		m.FhirVersion = r.FhirVersion
	}
	anyFhirVersionIdOrExtension := false
	for _, e := range r.FhirVersion {
		if e.Id != nil || e.Extension != nil {
			anyFhirVersionIdOrExtension = true
			break
		}
	}
	if anyFhirVersionIdOrExtension {
		m.FhirVersionPrimitiveElement = make([]*primitiveElement, 0, len(r.FhirVersion))
		for _, e := range r.FhirVersion {
			if e.Id != nil || e.Extension != nil {
				m.FhirVersionPrimitiveElement = append(m.FhirVersionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.FhirVersionPrimitiveElement = append(m.FhirVersionPrimitiveElement, nil)
			}
		}
	}
	m.DependsOn = r.DependsOn
	m.Global = r.Global
	m.Definition = r.Definition
	m.Manifest = r.Manifest
	return m
}
func (r *ImplementationGuide) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuide
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuide) unmarshalJSON(m jsonImplementationGuide) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		if r.Experimental == nil {
			r.Experimental = &Boolean{}
		}
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		if r.Publisher == nil {
			r.Publisher = &String{}
		}
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.PackageId = m.PackageId
	if m.PackageIdPrimitiveElement != nil {
		r.PackageId.Id = m.PackageIdPrimitiveElement.Id
		r.PackageId.Extension = m.PackageIdPrimitiveElement.Extension
	}
	r.License = m.License
	if m.LicensePrimitiveElement != nil {
		if r.License == nil {
			r.License = &Code{}
		}
		r.License.Id = m.LicensePrimitiveElement.Id
		r.License.Extension = m.LicensePrimitiveElement.Extension
	}
	r.FhirVersion = m.FhirVersion
	for i, e := range m.FhirVersionPrimitiveElement {
		if len(r.FhirVersion) <= i {
			r.FhirVersion = append(r.FhirVersion, Code{})
		}
		if e != nil {
			r.FhirVersion[i].Id = e.Id
			r.FhirVersion[i].Extension = e.Extension
		}
	}
	r.DependsOn = m.DependsOn
	r.Global = m.Global
	r.Definition = m.Definition
	r.Manifest = m.Manifest
	return nil
}
func (r ImplementationGuide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ImplementationGuide"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Experimental, xml.StartElement{Name: xml.Name{Local: "experimental"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UseContext, xml.StartElement{Name: xml.Name{Local: "useContext"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PackageId, xml.StartElement{Name: xml.Name{Local: "packageId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.License, xml.StartElement{Name: xml.Name{Local: "license"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FhirVersion, xml.StartElement{Name: xml.Name{Local: "fhirVersion"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DependsOn, xml.StartElement{Name: xml.Name{Local: "dependsOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Global, xml.StartElement{Name: xml.Name{Local: "global"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Definition, xml.StartElement{Name: xml.Name{Local: "definition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manifest, xml.StartElement{Name: xml.Name{Local: "manifest"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = v
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "experimental":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Experimental = &v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "publisher":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "contact":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "useContext":
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UseContext = append(r.UseContext, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			case "packageId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PackageId = v
			case "license":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.License = &v
			case "fhirVersion":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FhirVersion = append(r.FhirVersion, v)
			case "dependsOn":
				var v ImplementationGuideDependsOn
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DependsOn = append(r.DependsOn, v)
			case "global":
				var v ImplementationGuideGlobal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Global = append(r.Global, v)
			case "definition":
				var v ImplementationGuideDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Definition = &v
			case "manifest":
				var v ImplementationGuideManifest
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manifest = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuide) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
type ImplementationGuideDependsOn struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A canonical reference to the Implementation guide for the dependency.
	Uri Canonical
	// The NPM package name for the Implementation Guide that this IG depends on.
	PackageId *Id
	// The version of the IG that is depended on, when the correct version is required to understand the IG correctly.
	Version *String
}
type jsonImplementationGuideDependsOn struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Uri                       Canonical         `json:"uri,omitempty"`
	UriPrimitiveElement       *primitiveElement `json:"_uri,omitempty"`
	PackageId                 *Id               `json:"packageId,omitempty"`
	PackageIdPrimitiveElement *primitiveElement `json:"_packageId,omitempty"`
	Version                   *String           `json:"version,omitempty"`
	VersionPrimitiveElement   *primitiveElement `json:"_version,omitempty"`
}

func (r ImplementationGuideDependsOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDependsOn) marshalJSON() jsonImplementationGuideDependsOn {
	m := jsonImplementationGuideDependsOn{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Uri.Value != nil {
		m.Uri = r.Uri
	}
	if r.Uri.Id != nil || r.Uri.Extension != nil {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	if r.PackageId != nil && r.PackageId.Value != nil {
		m.PackageId = r.PackageId
	}
	if r.PackageId != nil && (r.PackageId.Id != nil || r.PackageId.Extension != nil) {
		m.PackageIdPrimitiveElement = &primitiveElement{Id: r.PackageId.Id, Extension: r.PackageId.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	return m
}
func (r *ImplementationGuideDependsOn) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDependsOn
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDependsOn) unmarshalJSON(m jsonImplementationGuideDependsOn) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	r.PackageId = m.PackageId
	if m.PackageIdPrimitiveElement != nil {
		if r.PackageId == nil {
			r.PackageId = &Id{}
		}
		r.PackageId.Id = m.PackageIdPrimitiveElement.Id
		r.PackageId.Extension = m.PackageIdPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideDependsOn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Uri, xml.StartElement{Name: xml.Name{Local: "uri"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PackageId, xml.StartElement{Name: xml.Name{Local: "packageId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDependsOn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "uri":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Uri = v
			case "packageId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PackageId = &v
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDependsOn) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A set of profiles that all resources covered by this implementation guide must conform to.
type ImplementationGuideGlobal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of resource that all instances must conform to.
	Type Code
	// A reference to the profile that all instances must conform to.
	Profile Canonical
}
type jsonImplementationGuideGlobal struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Type                    Code              `json:"type,omitempty"`
	TypePrimitiveElement    *primitiveElement `json:"_type,omitempty"`
	Profile                 Canonical         `json:"profile,omitempty"`
	ProfilePrimitiveElement *primitiveElement `json:"_profile,omitempty"`
}

func (r ImplementationGuideGlobal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideGlobal) marshalJSON() jsonImplementationGuideGlobal {
	m := jsonImplementationGuideGlobal{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Profile.Value != nil {
		m.Profile = r.Profile
	}
	if r.Profile.Id != nil || r.Profile.Extension != nil {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	return m
}
func (r *ImplementationGuideGlobal) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideGlobal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideGlobal) unmarshalJSON(m jsonImplementationGuideGlobal) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideGlobal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideGlobal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "profile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideGlobal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The information needed by an IG publisher tool to publish the whole implementation guide.
type ImplementationGuideDefinition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A logical group of resources. Logical groups can be used when building pages.
	Grouping []ImplementationGuideDefinitionGrouping
	// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
	Resource []ImplementationGuideDefinitionResource
	// A page / section in the implementation guide. The root page is the implementation guide home page.
	Page *ImplementationGuideDefinitionPage
	// Defines how IG is built by tools.
	Parameter []ImplementationGuideDefinitionParameter
	// A template for building resources.
	Template []ImplementationGuideDefinitionTemplate
}
type jsonImplementationGuideDefinition struct {
	Id                *string                                  `json:"id,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `json:"resource,omitempty"`
	Page              *ImplementationGuideDefinitionPage       `json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `json:"template,omitempty"`
}

func (r ImplementationGuideDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinition) marshalJSON() jsonImplementationGuideDefinition {
	m := jsonImplementationGuideDefinition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Grouping = r.Grouping
	m.Resource = r.Resource
	m.Page = r.Page
	m.Parameter = r.Parameter
	m.Template = r.Template
	return m
}
func (r *ImplementationGuideDefinition) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinition) unmarshalJSON(m jsonImplementationGuideDefinition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Grouping = m.Grouping
	r.Resource = m.Resource
	r.Page = m.Page
	r.Parameter = m.Parameter
	r.Template = m.Template
	return nil
}
func (r ImplementationGuideDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Grouping, xml.StartElement{Name: xml.Name{Local: "grouping"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Page, xml.StartElement{Name: xml.Name{Local: "page"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Parameter, xml.StartElement{Name: xml.Name{Local: "parameter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Template, xml.StartElement{Name: xml.Name{Local: "template"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "grouping":
				var v ImplementationGuideDefinitionGrouping
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Grouping = append(r.Grouping, v)
			case "resource":
				var v ImplementationGuideDefinitionResource
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = append(r.Resource, v)
			case "page":
				var v ImplementationGuideDefinitionPage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Page = &v
			case "parameter":
				var v ImplementationGuideDefinitionParameter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Parameter = append(r.Parameter, v)
			case "template":
				var v ImplementationGuideDefinitionTemplate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Template = append(r.Template, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A logical group of resources. Logical groups can be used when building pages.
type ImplementationGuideDefinitionGrouping struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The human-readable title to display for the package of resources when rendering the implementation guide.
	Name String
	// Human readable text describing the package.
	Description *String
}
type jsonImplementationGuideDefinitionGrouping struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Name                        String            `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r ImplementationGuideDefinitionGrouping) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinitionGrouping) marshalJSON() jsonImplementationGuideDefinitionGrouping {
	m := jsonImplementationGuideDefinitionGrouping{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *ImplementationGuideDefinitionGrouping) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinitionGrouping
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinitionGrouping) unmarshalJSON(m jsonImplementationGuideDefinitionGrouping) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideDefinitionGrouping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinitionGrouping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinitionGrouping) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideDefinitionResource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Where this resource is found.
	Reference Reference
	// Indicates the FHIR Version(s) this artifact is intended to apply to. If no versions are specified, the resource is assumed to apply to all the versions stated in ImplementationGuide.fhirVersion.
	FhirVersion []Code
	// A human assigned name for the resource. All resources SHOULD have a name, but the name may be extracted from the resource (e.g. ValueSet.name).
	Name *String
	// A description of the reason that a resource has been included in the implementation guide.
	Description *String
	// If true or a reference, indicates the resource is an example instance.  If a reference is present, indicates that the example is an example of the specified profile.
	Example isImplementationGuideDefinitionResourceExample
	// Reference to the id of the grouping this resource appears in.
	GroupingId *Id
}
type isImplementationGuideDefinitionResourceExample interface {
	isImplementationGuideDefinitionResourceExample()
}

func (r Boolean) isImplementationGuideDefinitionResourceExample()   {}
func (r Canonical) isImplementationGuideDefinitionResourceExample() {}

type jsonImplementationGuideDefinitionResource struct {
	Id                               *string             `json:"id,omitempty"`
	Extension                        []Extension         `json:"extension,omitempty"`
	ModifierExtension                []Extension         `json:"modifierExtension,omitempty"`
	Reference                        Reference           `json:"reference,omitempty"`
	FhirVersion                      []Code              `json:"fhirVersion,omitempty"`
	FhirVersionPrimitiveElement      []*primitiveElement `json:"_fhirVersion,omitempty"`
	Name                             *String             `json:"name,omitempty"`
	NamePrimitiveElement             *primitiveElement   `json:"_name,omitempty"`
	Description                      *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement   `json:"_description,omitempty"`
	ExampleBoolean                   *Boolean            `json:"exampleBoolean,omitempty"`
	ExampleBooleanPrimitiveElement   *primitiveElement   `json:"_exampleBoolean,omitempty"`
	ExampleCanonical                 *Canonical          `json:"exampleCanonical,omitempty"`
	ExampleCanonicalPrimitiveElement *primitiveElement   `json:"_exampleCanonical,omitempty"`
	GroupingId                       *Id                 `json:"groupingId,omitempty"`
	GroupingIdPrimitiveElement       *primitiveElement   `json:"_groupingId,omitempty"`
}

func (r ImplementationGuideDefinitionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinitionResource) marshalJSON() jsonImplementationGuideDefinitionResource {
	m := jsonImplementationGuideDefinitionResource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Reference = r.Reference
	anyFhirVersionValue := false
	for _, e := range r.FhirVersion {
		if e.Value != nil {
			anyFhirVersionValue = true
			break
		}
	}
	if anyFhirVersionValue {
		m.FhirVersion = r.FhirVersion
	}
	anyFhirVersionIdOrExtension := false
	for _, e := range r.FhirVersion {
		if e.Id != nil || e.Extension != nil {
			anyFhirVersionIdOrExtension = true
			break
		}
	}
	if anyFhirVersionIdOrExtension {
		m.FhirVersionPrimitiveElement = make([]*primitiveElement, 0, len(r.FhirVersion))
		for _, e := range r.FhirVersion {
			if e.Id != nil || e.Extension != nil {
				m.FhirVersionPrimitiveElement = append(m.FhirVersionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.FhirVersionPrimitiveElement = append(m.FhirVersionPrimitiveElement, nil)
			}
		}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	switch v := r.Example.(type) {
	case Boolean:
		if v.Value != nil {
			m.ExampleBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.ExampleBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		if v.Value != nil {
			m.ExampleCanonical = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		if v.Value != nil {
			m.ExampleCanonical = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.GroupingId != nil && r.GroupingId.Value != nil {
		m.GroupingId = r.GroupingId
	}
	if r.GroupingId != nil && (r.GroupingId.Id != nil || r.GroupingId.Extension != nil) {
		m.GroupingIdPrimitiveElement = &primitiveElement{Id: r.GroupingId.Id, Extension: r.GroupingId.Extension}
	}
	return m
}
func (r *ImplementationGuideDefinitionResource) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinitionResource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinitionResource) unmarshalJSON(m jsonImplementationGuideDefinitionResource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Reference = m.Reference
	r.FhirVersion = m.FhirVersion
	for i, e := range m.FhirVersionPrimitiveElement {
		if len(r.FhirVersion) <= i {
			r.FhirVersion = append(r.FhirVersion, Code{})
		}
		if e != nil {
			r.FhirVersion[i].Id = e.Id
			r.FhirVersion[i].Extension = e.Extension
		}
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	if m.ExampleBoolean != nil || m.ExampleBooleanPrimitiveElement != nil {
		if r.Example != nil {
			return fmt.Errorf("multiple values for field \"Example\"")
		}
		v := m.ExampleBoolean
		if m.ExampleBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ExampleBooleanPrimitiveElement.Id
			v.Extension = m.ExampleBooleanPrimitiveElement.Extension
		}
		r.Example = v
	}
	if m.ExampleCanonical != nil || m.ExampleCanonicalPrimitiveElement != nil {
		if r.Example != nil {
			return fmt.Errorf("multiple values for field \"Example\"")
		}
		v := m.ExampleCanonical
		if m.ExampleCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ExampleCanonicalPrimitiveElement.Id
			v.Extension = m.ExampleCanonicalPrimitiveElement.Extension
		}
		r.Example = v
	}
	r.GroupingId = m.GroupingId
	if m.GroupingIdPrimitiveElement != nil {
		if r.GroupingId == nil {
			r.GroupingId = &Id{}
		}
		r.GroupingId.Id = m.GroupingIdPrimitiveElement.Id
		r.GroupingId.Extension = m.GroupingIdPrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideDefinitionResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FhirVersion, xml.StartElement{Name: xml.Name{Local: "fhirVersion"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	switch v := r.Example.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "exampleBoolean"}})
		if err != nil {
			return err
		}
	case Canonical, *Canonical:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "exampleCanonical"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.GroupingId, xml.StartElement{Name: xml.Name{Local: "groupingId"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinitionResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = v
			case "fhirVersion":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FhirVersion = append(r.FhirVersion, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "exampleBoolean":
				if r.Example != nil {
					return fmt.Errorf("multiple values for field \"Example\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Example = v
			case "exampleCanonical":
				if r.Example != nil {
					return fmt.Errorf("multiple values for field \"Example\"")
				}
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Example = v
			case "groupingId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GroupingId = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinitionResource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A page / section in the implementation guide. The root page is the implementation guide home page.
type ImplementationGuideDefinitionPage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The source address for the page.
	Name isImplementationGuideDefinitionPageName
	// A short title used to represent this page in navigational structures such as table of contents, bread crumbs, etc.
	Title String
	// A code that indicates how the page is generated.
	Generation Code
	// Nested Pages/Sections under this page.
	Page []ImplementationGuideDefinitionPage
}
type isImplementationGuideDefinitionPageName interface {
	isImplementationGuideDefinitionPageName()
}

func (r Url) isImplementationGuideDefinitionPageName()       {}
func (r Reference) isImplementationGuideDefinitionPageName() {}

type jsonImplementationGuideDefinitionPage struct {
	Id                         *string                             `json:"id,omitempty"`
	Extension                  []Extension                         `json:"extension,omitempty"`
	ModifierExtension          []Extension                         `json:"modifierExtension,omitempty"`
	NameUrl                    *Url                                `json:"nameUrl,omitempty"`
	NameUrlPrimitiveElement    *primitiveElement                   `json:"_nameUrl,omitempty"`
	NameReference              *Reference                          `json:"nameReference,omitempty"`
	Title                      String                              `json:"title,omitempty"`
	TitlePrimitiveElement      *primitiveElement                   `json:"_title,omitempty"`
	Generation                 Code                                `json:"generation,omitempty"`
	GenerationPrimitiveElement *primitiveElement                   `json:"_generation,omitempty"`
	Page                       []ImplementationGuideDefinitionPage `json:"page,omitempty"`
}

func (r ImplementationGuideDefinitionPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinitionPage) marshalJSON() jsonImplementationGuideDefinitionPage {
	m := jsonImplementationGuideDefinitionPage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Name.(type) {
	case Url:
		if v.Value != nil {
			m.NameUrl = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.NameUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Url:
		if v.Value != nil {
			m.NameUrl = v
		}
		if v.Id != nil || v.Extension != nil {
			m.NameUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Reference:
		m.NameReference = &v
	case *Reference:
		m.NameReference = v
	}
	if r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title.Id != nil || r.Title.Extension != nil {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Generation.Value != nil {
		m.Generation = r.Generation
	}
	if r.Generation.Id != nil || r.Generation.Extension != nil {
		m.GenerationPrimitiveElement = &primitiveElement{Id: r.Generation.Id, Extension: r.Generation.Extension}
	}
	m.Page = r.Page
	return m
}
func (r *ImplementationGuideDefinitionPage) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinitionPage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinitionPage) unmarshalJSON(m jsonImplementationGuideDefinitionPage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.NameUrl != nil || m.NameUrlPrimitiveElement != nil {
		if r.Name != nil {
			return fmt.Errorf("multiple values for field \"Name\"")
		}
		v := m.NameUrl
		if m.NameUrlPrimitiveElement != nil {
			if v == nil {
				v = &Url{}
			}
			v.Id = m.NameUrlPrimitiveElement.Id
			v.Extension = m.NameUrlPrimitiveElement.Extension
		}
		r.Name = v
	}
	if m.NameReference != nil {
		if r.Name != nil {
			return fmt.Errorf("multiple values for field \"Name\"")
		}
		v := m.NameReference
		r.Name = v
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Generation = m.Generation
	if m.GenerationPrimitiveElement != nil {
		r.Generation.Id = m.GenerationPrimitiveElement.Id
		r.Generation.Extension = m.GenerationPrimitiveElement.Extension
	}
	r.Page = m.Page
	return nil
}
func (r ImplementationGuideDefinitionPage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	switch v := r.Name.(type) {
	case Url, *Url:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "nameUrl"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "nameReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Generation, xml.StartElement{Name: xml.Name{Local: "generation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Page, xml.StartElement{Name: xml.Name{Local: "page"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinitionPage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "nameUrl":
				if r.Name != nil {
					return fmt.Errorf("multiple values for field \"Name\"")
				}
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "nameReference":
				if r.Name != nil {
					return fmt.Errorf("multiple values for field \"Name\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = v
			case "generation":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Generation = v
			case "page":
				var v ImplementationGuideDefinitionPage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Page = append(r.Page, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinitionPage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Defines how IG is built by tools.
type ImplementationGuideDefinitionParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// apply | path-resource | path-pages | path-tx-cache | expansion-parameter | rule-broken-links | generate-xml | generate-json | generate-turtle | html-template.
	Code Code
	// Value for named type.
	Value String
}
type jsonImplementationGuideDefinitionParameter struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Code                  Code              `json:"code,omitempty"`
	CodePrimitiveElement  *primitiveElement `json:"_code,omitempty"`
	Value                 String            `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r ImplementationGuideDefinitionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinitionParameter) marshalJSON() jsonImplementationGuideDefinitionParameter {
	m := jsonImplementationGuideDefinitionParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *ImplementationGuideDefinitionParameter) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinitionParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinitionParameter) unmarshalJSON(m jsonImplementationGuideDefinitionParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideDefinitionParameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinitionParameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinitionParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A template for building resources.
type ImplementationGuideDefinitionTemplate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of template specified.
	Code Code
	// The source location for the template.
	Source String
	// The scope in which the template applies.
	Scope *String
}
type jsonImplementationGuideDefinitionTemplate struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Code                   Code              `json:"code,omitempty"`
	CodePrimitiveElement   *primitiveElement `json:"_code,omitempty"`
	Source                 String            `json:"source,omitempty"`
	SourcePrimitiveElement *primitiveElement `json:"_source,omitempty"`
	Scope                  *String           `json:"scope,omitempty"`
	ScopePrimitiveElement  *primitiveElement `json:"_scope,omitempty"`
}

func (r ImplementationGuideDefinitionTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideDefinitionTemplate) marshalJSON() jsonImplementationGuideDefinitionTemplate {
	m := jsonImplementationGuideDefinitionTemplate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Source.Value != nil {
		m.Source = r.Source
	}
	if r.Source.Id != nil || r.Source.Extension != nil {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	if r.Scope != nil && r.Scope.Value != nil {
		m.Scope = r.Scope
	}
	if r.Scope != nil && (r.Scope.Id != nil || r.Scope.Extension != nil) {
		m.ScopePrimitiveElement = &primitiveElement{Id: r.Scope.Id, Extension: r.Scope.Extension}
	}
	return m
}
func (r *ImplementationGuideDefinitionTemplate) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideDefinitionTemplate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideDefinitionTemplate) unmarshalJSON(m jsonImplementationGuideDefinitionTemplate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.Scope = m.Scope
	if m.ScopePrimitiveElement != nil {
		if r.Scope == nil {
			r.Scope = &String{}
		}
		r.Scope.Id = m.ScopePrimitiveElement.Id
		r.Scope.Extension = m.ScopePrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideDefinitionTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Scope, xml.StartElement{Name: xml.Name{Local: "scope"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideDefinitionTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "source":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = v
			case "scope":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Scope = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideDefinitionTemplate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information about an assembled implementation guide, created by the publication tooling.
type ImplementationGuideManifest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A pointer to official web page, PDF or other rendering of the implementation guide.
	Rendering *Url
	// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
	Resource []ImplementationGuideManifestResource
	// Information about a page within the IG.
	Page []ImplementationGuideManifestPage
	// Indicates a relative path to an image that exists within the IG.
	Image []String
	// Indicates the relative path of an additional non-page, non-image file that is part of the IG - e.g. zip, jar and similar files that could be the target of a hyperlink in a derived IG.
	Other []String
}
type jsonImplementationGuideManifest struct {
	Id                        *string                               `json:"id,omitempty"`
	Extension                 []Extension                           `json:"extension,omitempty"`
	ModifierExtension         []Extension                           `json:"modifierExtension,omitempty"`
	Rendering                 *Url                                  `json:"rendering,omitempty"`
	RenderingPrimitiveElement *primitiveElement                     `json:"_rendering,omitempty"`
	Resource                  []ImplementationGuideManifestResource `json:"resource,omitempty"`
	Page                      []ImplementationGuideManifestPage     `json:"page,omitempty"`
	Image                     []String                              `json:"image,omitempty"`
	ImagePrimitiveElement     []*primitiveElement                   `json:"_image,omitempty"`
	Other                     []String                              `json:"other,omitempty"`
	OtherPrimitiveElement     []*primitiveElement                   `json:"_other,omitempty"`
}

func (r ImplementationGuideManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideManifest) marshalJSON() jsonImplementationGuideManifest {
	m := jsonImplementationGuideManifest{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Rendering != nil && r.Rendering.Value != nil {
		m.Rendering = r.Rendering
	}
	if r.Rendering != nil && (r.Rendering.Id != nil || r.Rendering.Extension != nil) {
		m.RenderingPrimitiveElement = &primitiveElement{Id: r.Rendering.Id, Extension: r.Rendering.Extension}
	}
	m.Resource = r.Resource
	m.Page = r.Page
	anyImageValue := false
	for _, e := range r.Image {
		if e.Value != nil {
			anyImageValue = true
			break
		}
	}
	if anyImageValue {
		m.Image = r.Image
	}
	anyImageIdOrExtension := false
	for _, e := range r.Image {
		if e.Id != nil || e.Extension != nil {
			anyImageIdOrExtension = true
			break
		}
	}
	if anyImageIdOrExtension {
		m.ImagePrimitiveElement = make([]*primitiveElement, 0, len(r.Image))
		for _, e := range r.Image {
			if e.Id != nil || e.Extension != nil {
				m.ImagePrimitiveElement = append(m.ImagePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ImagePrimitiveElement = append(m.ImagePrimitiveElement, nil)
			}
		}
	}
	anyOtherValue := false
	for _, e := range r.Other {
		if e.Value != nil {
			anyOtherValue = true
			break
		}
	}
	if anyOtherValue {
		m.Other = r.Other
	}
	anyOtherIdOrExtension := false
	for _, e := range r.Other {
		if e.Id != nil || e.Extension != nil {
			anyOtherIdOrExtension = true
			break
		}
	}
	if anyOtherIdOrExtension {
		m.OtherPrimitiveElement = make([]*primitiveElement, 0, len(r.Other))
		for _, e := range r.Other {
			if e.Id != nil || e.Extension != nil {
				m.OtherPrimitiveElement = append(m.OtherPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.OtherPrimitiveElement = append(m.OtherPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ImplementationGuideManifest) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideManifest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideManifest) unmarshalJSON(m jsonImplementationGuideManifest) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Rendering = m.Rendering
	if m.RenderingPrimitiveElement != nil {
		if r.Rendering == nil {
			r.Rendering = &Url{}
		}
		r.Rendering.Id = m.RenderingPrimitiveElement.Id
		r.Rendering.Extension = m.RenderingPrimitiveElement.Extension
	}
	r.Resource = m.Resource
	r.Page = m.Page
	r.Image = m.Image
	for i, e := range m.ImagePrimitiveElement {
		if len(r.Image) <= i {
			r.Image = append(r.Image, String{})
		}
		if e != nil {
			r.Image[i].Id = e.Id
			r.Image[i].Extension = e.Extension
		}
	}
	r.Other = m.Other
	for i, e := range m.OtherPrimitiveElement {
		if len(r.Other) <= i {
			r.Other = append(r.Other, String{})
		}
		if e != nil {
			r.Other[i].Id = e.Id
			r.Other[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ImplementationGuideManifest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Rendering, xml.StartElement{Name: xml.Name{Local: "rendering"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Page, xml.StartElement{Name: xml.Name{Local: "page"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Image, xml.StartElement{Name: xml.Name{Local: "image"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Other, xml.StartElement{Name: xml.Name{Local: "other"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideManifest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "rendering":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rendering = &v
			case "resource":
				var v ImplementationGuideManifestResource
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = append(r.Resource, v)
			case "page":
				var v ImplementationGuideManifestPage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Page = append(r.Page, v)
			case "image":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Image = append(r.Image, v)
			case "other":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Other = append(r.Other, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideManifest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideManifestResource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Where this resource is found.
	Reference Reference
	// If true or a reference, indicates the resource is an example instance.  If a reference is present, indicates that the example is an example of the specified profile.
	Example isImplementationGuideManifestResourceExample
	// The relative path for primary page for this resource within the IG.
	RelativePath *Url
}
type isImplementationGuideManifestResourceExample interface {
	isImplementationGuideManifestResourceExample()
}

func (r Boolean) isImplementationGuideManifestResourceExample()   {}
func (r Canonical) isImplementationGuideManifestResourceExample() {}

type jsonImplementationGuideManifestResource struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	Reference                        Reference         `json:"reference,omitempty"`
	ExampleBoolean                   *Boolean          `json:"exampleBoolean,omitempty"`
	ExampleBooleanPrimitiveElement   *primitiveElement `json:"_exampleBoolean,omitempty"`
	ExampleCanonical                 *Canonical        `json:"exampleCanonical,omitempty"`
	ExampleCanonicalPrimitiveElement *primitiveElement `json:"_exampleCanonical,omitempty"`
	RelativePath                     *Url              `json:"relativePath,omitempty"`
	RelativePathPrimitiveElement     *primitiveElement `json:"_relativePath,omitempty"`
}

func (r ImplementationGuideManifestResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideManifestResource) marshalJSON() jsonImplementationGuideManifestResource {
	m := jsonImplementationGuideManifestResource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Reference = r.Reference
	switch v := r.Example.(type) {
	case Boolean:
		if v.Value != nil {
			m.ExampleBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.ExampleBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		if v.Value != nil {
			m.ExampleCanonical = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		if v.Value != nil {
			m.ExampleCanonical = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ExampleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.RelativePath != nil && r.RelativePath.Value != nil {
		m.RelativePath = r.RelativePath
	}
	if r.RelativePath != nil && (r.RelativePath.Id != nil || r.RelativePath.Extension != nil) {
		m.RelativePathPrimitiveElement = &primitiveElement{Id: r.RelativePath.Id, Extension: r.RelativePath.Extension}
	}
	return m
}
func (r *ImplementationGuideManifestResource) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideManifestResource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideManifestResource) unmarshalJSON(m jsonImplementationGuideManifestResource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Reference = m.Reference
	if m.ExampleBoolean != nil || m.ExampleBooleanPrimitiveElement != nil {
		if r.Example != nil {
			return fmt.Errorf("multiple values for field \"Example\"")
		}
		v := m.ExampleBoolean
		if m.ExampleBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ExampleBooleanPrimitiveElement.Id
			v.Extension = m.ExampleBooleanPrimitiveElement.Extension
		}
		r.Example = v
	}
	if m.ExampleCanonical != nil || m.ExampleCanonicalPrimitiveElement != nil {
		if r.Example != nil {
			return fmt.Errorf("multiple values for field \"Example\"")
		}
		v := m.ExampleCanonical
		if m.ExampleCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ExampleCanonicalPrimitiveElement.Id
			v.Extension = m.ExampleCanonicalPrimitiveElement.Extension
		}
		r.Example = v
	}
	r.RelativePath = m.RelativePath
	if m.RelativePathPrimitiveElement != nil {
		if r.RelativePath == nil {
			r.RelativePath = &Url{}
		}
		r.RelativePath.Id = m.RelativePathPrimitiveElement.Id
		r.RelativePath.Extension = m.RelativePathPrimitiveElement.Extension
	}
	return nil
}
func (r ImplementationGuideManifestResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	switch v := r.Example.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "exampleBoolean"}})
		if err != nil {
			return err
		}
	case Canonical, *Canonical:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "exampleCanonical"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.RelativePath, xml.StartElement{Name: xml.Name{Local: "relativePath"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideManifestResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = v
			case "exampleBoolean":
				if r.Example != nil {
					return fmt.Errorf("multiple values for field \"Example\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Example = v
			case "exampleCanonical":
				if r.Example != nil {
					return fmt.Errorf("multiple values for field \"Example\"")
				}
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Example = v
			case "relativePath":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelativePath = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideManifestResource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information about a page within the IG.
type ImplementationGuideManifestPage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Relative path to the page.
	Name String
	// Label for the page intended for human display.
	Title *String
	// The name of an anchor available on the page.
	Anchor []String
}
type jsonImplementationGuideManifestPage struct {
	Id                     *string             `json:"id,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	ModifierExtension      []Extension         `json:"modifierExtension,omitempty"`
	Name                   String              `json:"name,omitempty"`
	NamePrimitiveElement   *primitiveElement   `json:"_name,omitempty"`
	Title                  *String             `json:"title,omitempty"`
	TitlePrimitiveElement  *primitiveElement   `json:"_title,omitempty"`
	Anchor                 []String            `json:"anchor,omitempty"`
	AnchorPrimitiveElement []*primitiveElement `json:"_anchor,omitempty"`
}

func (r ImplementationGuideManifestPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImplementationGuideManifestPage) marshalJSON() jsonImplementationGuideManifestPage {
	m := jsonImplementationGuideManifestPage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	anyAnchorValue := false
	for _, e := range r.Anchor {
		if e.Value != nil {
			anyAnchorValue = true
			break
		}
	}
	if anyAnchorValue {
		m.Anchor = r.Anchor
	}
	anyAnchorIdOrExtension := false
	for _, e := range r.Anchor {
		if e.Id != nil || e.Extension != nil {
			anyAnchorIdOrExtension = true
			break
		}
	}
	if anyAnchorIdOrExtension {
		m.AnchorPrimitiveElement = make([]*primitiveElement, 0, len(r.Anchor))
		for _, e := range r.Anchor {
			if e.Id != nil || e.Extension != nil {
				m.AnchorPrimitiveElement = append(m.AnchorPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.AnchorPrimitiveElement = append(m.AnchorPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ImplementationGuideManifestPage) UnmarshalJSON(b []byte) error {
	var m jsonImplementationGuideManifestPage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImplementationGuideManifestPage) unmarshalJSON(m jsonImplementationGuideManifestPage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Anchor = m.Anchor
	for i, e := range m.AnchorPrimitiveElement {
		if len(r.Anchor) <= i {
			r.Anchor = append(r.Anchor, String{})
		}
		if e != nil {
			r.Anchor[i].Id = e.Id
			r.Anchor[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ImplementationGuideManifestPage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Anchor, xml.StartElement{Name: xml.Name{Local: "anchor"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImplementationGuideManifestPage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "anchor":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Anchor = append(r.Anchor, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImplementationGuideManifestPage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
