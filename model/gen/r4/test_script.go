package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScript struct {
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
	// An absolute URI that is used to identify this test script when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this test script is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the test script is stored on different servers.
	Url Uri
	// A formal identifier that is used to identify this test script when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier
	// The identifier that is used to identify this version of the test script when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the test script author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the test script. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// A short, descriptive, user-friendly title for the test script.
	Title *String
	// The status of this test script. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this test script is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the test script was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the test script changes.
	Date *DateTime
	// The name of the organization or individual that published the test script.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the test script from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate test script instances.
	UseContext []UsageContext
	// A legal or geographic region in which the test script is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this test script is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the test script and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the test script.
	Copyright *Markdown
	// An abstract server used in operations within this test script in the origin element.
	Origin []TestScriptOrigin
	// An abstract server used in operations within this test script in the destination element.
	Destination []TestScriptDestination
	// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
	Metadata *TestScriptMetadata
	// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
	Fixture []TestScriptFixture
	// Reference to the profile to be used for validation.
	Profile []Reference
	// Variable is set based either on element value in response body or on header field value in the response headers.
	Variable []TestScriptVariable
	// A series of required setup operations before tests are executed.
	Setup *TestScriptSetup
	// A test in this script.
	Test []TestScriptTest
	// A series of operations required to clean up after all the tests are executed (successfully or otherwise).
	Teardown *TestScriptTeardown
}

func (r TestScript) ResourceType() string {
	return "TestScript"
}
func (r TestScript) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonTestScript struct {
	ResourceType                  string                  `json:"resourceType"`
	Id                            *Id                     `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement       `json:"_id,omitempty"`
	Meta                          *Meta                   `json:"meta,omitempty"`
	ImplicitRules                 *Uri                    `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement       `json:"_implicitRules,omitempty"`
	Language                      *Code                   `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement       `json:"_language,omitempty"`
	Text                          *Narrative              `json:"text,omitempty"`
	Contained                     []ContainedResource     `json:"contained,omitempty"`
	Extension                     []Extension             `json:"extension,omitempty"`
	ModifierExtension             []Extension             `json:"modifierExtension,omitempty"`
	Url                           Uri                     `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement       `json:"_url,omitempty"`
	Identifier                    *Identifier             `json:"identifier,omitempty"`
	Version                       *String                 `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement       `json:"_version,omitempty"`
	Name                          String                  `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement       `json:"_name,omitempty"`
	Title                         *String                 `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement       `json:"_title,omitempty"`
	Status                        Code                    `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement       `json:"_status,omitempty"`
	Experimental                  *Boolean                `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement       `json:"_experimental,omitempty"`
	Date                          *DateTime               `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement       `json:"_date,omitempty"`
	Publisher                     *String                 `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement       `json:"_publisher,omitempty"`
	Contact                       []ContactDetail         `json:"contact,omitempty"`
	Description                   *Markdown               `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement       `json:"_description,omitempty"`
	UseContext                    []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown               `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement       `json:"_purpose,omitempty"`
	Copyright                     *Markdown               `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement       `json:"_copyright,omitempty"`
	Origin                        []TestScriptOrigin      `json:"origin,omitempty"`
	Destination                   []TestScriptDestination `json:"destination,omitempty"`
	Metadata                      *TestScriptMetadata     `json:"metadata,omitempty"`
	Fixture                       []TestScriptFixture     `json:"fixture,omitempty"`
	Profile                       []Reference             `json:"profile,omitempty"`
	Variable                      []TestScriptVariable    `json:"variable,omitempty"`
	Setup                         *TestScriptSetup        `json:"setup,omitempty"`
	Test                          []TestScriptTest        `json:"test,omitempty"`
	Teardown                      *TestScriptTeardown     `json:"teardown,omitempty"`
}

func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScript) marshalJSON() jsonTestScript {
	m := jsonTestScript{}
	m.ResourceType = "TestScript"
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
	m.Identifier = r.Identifier
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
	if r.Purpose != nil && r.Purpose.Value != nil {
		m.Purpose = r.Purpose
	}
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.Origin = r.Origin
	m.Destination = r.Destination
	m.Metadata = r.Metadata
	m.Fixture = r.Fixture
	m.Profile = r.Profile
	m.Variable = r.Variable
	m.Setup = r.Setup
	m.Test = r.Test
	m.Teardown = r.Teardown
	return m
}
func (r *TestScript) UnmarshalJSON(b []byte) error {
	var m jsonTestScript
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScript) unmarshalJSON(m jsonTestScript) error {
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
	r.Identifier = m.Identifier
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
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		if r.Purpose == nil {
			r.Purpose = &Markdown{}
		}
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.Origin = m.Origin
	r.Destination = m.Destination
	r.Metadata = m.Metadata
	r.Fixture = m.Fixture
	r.Profile = m.Profile
	r.Variable = m.Variable
	r.Setup = m.Setup
	r.Test = m.Test
	r.Teardown = m.Teardown
	return nil
}
func (r TestScript) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "TestScript"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
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
	err = e.EncodeElement(r.Purpose, xml.StartElement{Name: xml.Name{Local: "purpose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Origin, xml.StartElement{Name: xml.Name{Local: "origin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Destination, xml.StartElement{Name: xml.Name{Local: "destination"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Metadata, xml.StartElement{Name: xml.Name{Local: "metadata"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Fixture, xml.StartElement{Name: xml.Name{Local: "fixture"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Variable, xml.StartElement{Name: xml.Name{Local: "variable"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Setup, xml.StartElement{Name: xml.Name{Local: "setup"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Test, xml.StartElement{Name: xml.Name{Local: "test"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Teardown, xml.StartElement{Name: xml.Name{Local: "teardown"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScript) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
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
			case "purpose":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Purpose = &v
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			case "origin":
				var v TestScriptOrigin
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Origin = append(r.Origin, v)
			case "destination":
				var v TestScriptDestination
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Destination = append(r.Destination, v)
			case "metadata":
				var v TestScriptMetadata
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Metadata = &v
			case "fixture":
				var v TestScriptFixture
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Fixture = append(r.Fixture, v)
			case "profile":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = append(r.Profile, v)
			case "variable":
				var v TestScriptVariable
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Variable = append(r.Variable, v)
			case "setup":
				var v TestScriptSetup
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Setup = &v
			case "test":
				var v TestScriptTest
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Test = append(r.Test, v)
			case "teardown":
				var v TestScriptTeardown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Teardown = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScript) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An abstract server used in operations within this test script in the origin element.
type TestScriptOrigin struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Abstract name given to an origin server in this test script.  The name is provided as a number starting at 1.
	Index Integer
	// The type of origin profile the test system supports.
	Profile Coding
}
type jsonTestScriptOrigin struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Index                 Integer           `json:"index,omitempty"`
	IndexPrimitiveElement *primitiveElement `json:"_index,omitempty"`
	Profile               Coding            `json:"profile,omitempty"`
}

func (r TestScriptOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptOrigin) marshalJSON() jsonTestScriptOrigin {
	m := jsonTestScriptOrigin{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Index.Value != nil {
		m.Index = r.Index
	}
	if r.Index.Id != nil || r.Index.Extension != nil {
		m.IndexPrimitiveElement = &primitiveElement{Id: r.Index.Id, Extension: r.Index.Extension}
	}
	m.Profile = r.Profile
	return m
}
func (r *TestScriptOrigin) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptOrigin
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptOrigin) unmarshalJSON(m jsonTestScriptOrigin) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Index = m.Index
	if m.IndexPrimitiveElement != nil {
		r.Index.Id = m.IndexPrimitiveElement.Id
		r.Index.Extension = m.IndexPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	return nil
}
func (r TestScriptOrigin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Index, xml.StartElement{Name: xml.Name{Local: "index"}})
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
func (r *TestScriptOrigin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "index":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Index = v
			case "profile":
				var v Coding
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
func (r TestScriptOrigin) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An abstract server used in operations within this test script in the destination element.
type TestScriptDestination struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Abstract name given to a destination server in this test script.  The name is provided as a number starting at 1.
	Index Integer
	// The type of destination profile the test system supports.
	Profile Coding
}
type jsonTestScriptDestination struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Index                 Integer           `json:"index,omitempty"`
	IndexPrimitiveElement *primitiveElement `json:"_index,omitempty"`
	Profile               Coding            `json:"profile,omitempty"`
}

func (r TestScriptDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptDestination) marshalJSON() jsonTestScriptDestination {
	m := jsonTestScriptDestination{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Index.Value != nil {
		m.Index = r.Index
	}
	if r.Index.Id != nil || r.Index.Extension != nil {
		m.IndexPrimitiveElement = &primitiveElement{Id: r.Index.Id, Extension: r.Index.Extension}
	}
	m.Profile = r.Profile
	return m
}
func (r *TestScriptDestination) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptDestination
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptDestination) unmarshalJSON(m jsonTestScriptDestination) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Index = m.Index
	if m.IndexPrimitiveElement != nil {
		r.Index.Id = m.IndexPrimitiveElement.Id
		r.Index.Extension = m.IndexPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	return nil
}
func (r TestScriptDestination) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Index, xml.StartElement{Name: xml.Name{Local: "index"}})
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
func (r *TestScriptDestination) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "index":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Index = v
			case "profile":
				var v Coding
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
func (r TestScriptDestination) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
type TestScriptMetadata struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A link to the FHIR specification that this test is covering.
	Link []TestScriptMetadataLink
	// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
	Capability []TestScriptMetadataCapability
}
type jsonTestScriptMetadata struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `json:"capability,omitempty"`
}

func (r TestScriptMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptMetadata) marshalJSON() jsonTestScriptMetadata {
	m := jsonTestScriptMetadata{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Link = r.Link
	m.Capability = r.Capability
	return m
}
func (r *TestScriptMetadata) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptMetadata
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptMetadata) unmarshalJSON(m jsonTestScriptMetadata) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Link = m.Link
	r.Capability = m.Capability
	return nil
}
func (r TestScriptMetadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Capability, xml.StartElement{Name: xml.Name{Local: "capability"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptMetadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "link":
				var v TestScriptMetadataLink
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Link = append(r.Link, v)
			case "capability":
				var v TestScriptMetadataCapability
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Capability = append(r.Capability, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptMetadata) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A link to the FHIR specification that this test is covering.
type TestScriptMetadataLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// URL to a particular requirement or feature within the FHIR specification.
	Url Uri
	// Short description of the link.
	Description *String
}
type jsonTestScriptMetadataLink struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Url                         Uri               `json:"url,omitempty"`
	UrlPrimitiveElement         *primitiveElement `json:"_url,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r TestScriptMetadataLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptMetadataLink) marshalJSON() jsonTestScriptMetadataLink {
	m := jsonTestScriptMetadataLink{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *TestScriptMetadataLink) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptMetadataLink
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptMetadataLink) unmarshalJSON(m jsonTestScriptMetadataLink) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
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
func (r TestScriptMetadataLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
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
func (r *TestScriptMetadataLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = v
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
func (r TestScriptMetadataLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
type TestScriptMetadataCapability struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether or not the test execution will require the given capabilities of the server in order for this test script to execute.
	Required Boolean
	// Whether or not the test execution will validate the given capabilities of the server in order for this test script to execute.
	Validated Boolean
	// Description of the capabilities that this test script is requiring the server to support.
	Description *String
	// Which origin server these requirements apply to.
	Origin []Integer
	// Which server these requirements apply to.
	Destination *Integer
	// Links to the FHIR specification that describes this interaction and the resources involved in more detail.
	Link []Uri
	// Minimum capabilities required of server for test script to execute successfully.   If server does not meet at a minimum the referenced capability statement, then all tests in this script are skipped.
	Capabilities Canonical
}
type jsonTestScriptMetadataCapability struct {
	Id                           *string             `json:"id,omitempty"`
	Extension                    []Extension         `json:"extension,omitempty"`
	ModifierExtension            []Extension         `json:"modifierExtension,omitempty"`
	Required                     Boolean             `json:"required,omitempty"`
	RequiredPrimitiveElement     *primitiveElement   `json:"_required,omitempty"`
	Validated                    Boolean             `json:"validated,omitempty"`
	ValidatedPrimitiveElement    *primitiveElement   `json:"_validated,omitempty"`
	Description                  *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement   `json:"_description,omitempty"`
	Origin                       []Integer           `json:"origin,omitempty"`
	OriginPrimitiveElement       []*primitiveElement `json:"_origin,omitempty"`
	Destination                  *Integer            `json:"destination,omitempty"`
	DestinationPrimitiveElement  *primitiveElement   `json:"_destination,omitempty"`
	Link                         []Uri               `json:"link,omitempty"`
	LinkPrimitiveElement         []*primitiveElement `json:"_link,omitempty"`
	Capabilities                 Canonical           `json:"capabilities,omitempty"`
	CapabilitiesPrimitiveElement *primitiveElement   `json:"_capabilities,omitempty"`
}

func (r TestScriptMetadataCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptMetadataCapability) marshalJSON() jsonTestScriptMetadataCapability {
	m := jsonTestScriptMetadataCapability{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Required.Value != nil {
		m.Required = r.Required
	}
	if r.Required.Id != nil || r.Required.Extension != nil {
		m.RequiredPrimitiveElement = &primitiveElement{Id: r.Required.Id, Extension: r.Required.Extension}
	}
	if r.Validated.Value != nil {
		m.Validated = r.Validated
	}
	if r.Validated.Id != nil || r.Validated.Extension != nil {
		m.ValidatedPrimitiveElement = &primitiveElement{Id: r.Validated.Id, Extension: r.Validated.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	anyOriginValue := false
	for _, e := range r.Origin {
		if e.Value != nil {
			anyOriginValue = true
			break
		}
	}
	if anyOriginValue {
		m.Origin = r.Origin
	}
	anyOriginIdOrExtension := false
	for _, e := range r.Origin {
		if e.Id != nil || e.Extension != nil {
			anyOriginIdOrExtension = true
			break
		}
	}
	if anyOriginIdOrExtension {
		m.OriginPrimitiveElement = make([]*primitiveElement, 0, len(r.Origin))
		for _, e := range r.Origin {
			if e.Id != nil || e.Extension != nil {
				m.OriginPrimitiveElement = append(m.OriginPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.OriginPrimitiveElement = append(m.OriginPrimitiveElement, nil)
			}
		}
	}
	if r.Destination != nil && r.Destination.Value != nil {
		m.Destination = r.Destination
	}
	if r.Destination != nil && (r.Destination.Id != nil || r.Destination.Extension != nil) {
		m.DestinationPrimitiveElement = &primitiveElement{Id: r.Destination.Id, Extension: r.Destination.Extension}
	}
	anyLinkValue := false
	for _, e := range r.Link {
		if e.Value != nil {
			anyLinkValue = true
			break
		}
	}
	if anyLinkValue {
		m.Link = r.Link
	}
	anyLinkIdOrExtension := false
	for _, e := range r.Link {
		if e.Id != nil || e.Extension != nil {
			anyLinkIdOrExtension = true
			break
		}
	}
	if anyLinkIdOrExtension {
		m.LinkPrimitiveElement = make([]*primitiveElement, 0, len(r.Link))
		for _, e := range r.Link {
			if e.Id != nil || e.Extension != nil {
				m.LinkPrimitiveElement = append(m.LinkPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinkPrimitiveElement = append(m.LinkPrimitiveElement, nil)
			}
		}
	}
	if r.Capabilities.Value != nil {
		m.Capabilities = r.Capabilities
	}
	if r.Capabilities.Id != nil || r.Capabilities.Extension != nil {
		m.CapabilitiesPrimitiveElement = &primitiveElement{Id: r.Capabilities.Id, Extension: r.Capabilities.Extension}
	}
	return m
}
func (r *TestScriptMetadataCapability) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptMetadataCapability
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptMetadataCapability) unmarshalJSON(m jsonTestScriptMetadataCapability) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Required = m.Required
	if m.RequiredPrimitiveElement != nil {
		r.Required.Id = m.RequiredPrimitiveElement.Id
		r.Required.Extension = m.RequiredPrimitiveElement.Extension
	}
	r.Validated = m.Validated
	if m.ValidatedPrimitiveElement != nil {
		r.Validated.Id = m.ValidatedPrimitiveElement.Id
		r.Validated.Extension = m.ValidatedPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Origin = m.Origin
	for i, e := range m.OriginPrimitiveElement {
		if len(r.Origin) <= i {
			r.Origin = append(r.Origin, Integer{})
		}
		if e != nil {
			r.Origin[i].Id = e.Id
			r.Origin[i].Extension = e.Extension
		}
	}
	r.Destination = m.Destination
	if m.DestinationPrimitiveElement != nil {
		if r.Destination == nil {
			r.Destination = &Integer{}
		}
		r.Destination.Id = m.DestinationPrimitiveElement.Id
		r.Destination.Extension = m.DestinationPrimitiveElement.Extension
	}
	r.Link = m.Link
	for i, e := range m.LinkPrimitiveElement {
		if len(r.Link) <= i {
			r.Link = append(r.Link, Uri{})
		}
		if e != nil {
			r.Link[i].Id = e.Id
			r.Link[i].Extension = e.Extension
		}
	}
	r.Capabilities = m.Capabilities
	if m.CapabilitiesPrimitiveElement != nil {
		r.Capabilities.Id = m.CapabilitiesPrimitiveElement.Id
		r.Capabilities.Extension = m.CapabilitiesPrimitiveElement.Extension
	}
	return nil
}
func (r TestScriptMetadataCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Required, xml.StartElement{Name: xml.Name{Local: "required"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Validated, xml.StartElement{Name: xml.Name{Local: "validated"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Origin, xml.StartElement{Name: xml.Name{Local: "origin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Destination, xml.StartElement{Name: xml.Name{Local: "destination"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Capabilities, xml.StartElement{Name: xml.Name{Local: "capabilities"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptMetadataCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "required":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Required = v
			case "validated":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Validated = v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "origin":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Origin = append(r.Origin, v)
			case "destination":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Destination = &v
			case "link":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Link = append(r.Link, v)
			case "capabilities":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Capabilities = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptMetadataCapability) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
type TestScriptFixture struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether or not to implicitly create the fixture during setup. If true, the fixture is automatically created on each server being tested during setup, therefore no create operation is required for this fixture in the TestScript.setup section.
	Autocreate Boolean
	// Whether or not to implicitly delete the fixture during teardown. If true, the fixture is automatically deleted on each server being tested during teardown, therefore no delete operation is required for this fixture in the TestScript.teardown section.
	Autodelete Boolean
	// Reference to the resource (containing the contents of the resource needed for operations).
	Resource *Reference
}
type jsonTestScriptFixture struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Autocreate                 Boolean           `json:"autocreate,omitempty"`
	AutocreatePrimitiveElement *primitiveElement `json:"_autocreate,omitempty"`
	Autodelete                 Boolean           `json:"autodelete,omitempty"`
	AutodeletePrimitiveElement *primitiveElement `json:"_autodelete,omitempty"`
	Resource                   *Reference        `json:"resource,omitempty"`
}

func (r TestScriptFixture) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptFixture) marshalJSON() jsonTestScriptFixture {
	m := jsonTestScriptFixture{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Autocreate.Value != nil {
		m.Autocreate = r.Autocreate
	}
	if r.Autocreate.Id != nil || r.Autocreate.Extension != nil {
		m.AutocreatePrimitiveElement = &primitiveElement{Id: r.Autocreate.Id, Extension: r.Autocreate.Extension}
	}
	if r.Autodelete.Value != nil {
		m.Autodelete = r.Autodelete
	}
	if r.Autodelete.Id != nil || r.Autodelete.Extension != nil {
		m.AutodeletePrimitiveElement = &primitiveElement{Id: r.Autodelete.Id, Extension: r.Autodelete.Extension}
	}
	m.Resource = r.Resource
	return m
}
func (r *TestScriptFixture) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptFixture
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptFixture) unmarshalJSON(m jsonTestScriptFixture) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Autocreate = m.Autocreate
	if m.AutocreatePrimitiveElement != nil {
		r.Autocreate.Id = m.AutocreatePrimitiveElement.Id
		r.Autocreate.Extension = m.AutocreatePrimitiveElement.Extension
	}
	r.Autodelete = m.Autodelete
	if m.AutodeletePrimitiveElement != nil {
		r.Autodelete.Id = m.AutodeletePrimitiveElement.Id
		r.Autodelete.Extension = m.AutodeletePrimitiveElement.Extension
	}
	r.Resource = m.Resource
	return nil
}
func (r TestScriptFixture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Autocreate, xml.StartElement{Name: xml.Name{Local: "autocreate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Autodelete, xml.StartElement{Name: xml.Name{Local: "autodelete"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptFixture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "autocreate":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Autocreate = v
			case "autodelete":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Autodelete = v
			case "resource":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptFixture) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Variable is set based either on element value in response body or on header field value in the response headers.
type TestScriptVariable struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Descriptive name for this variable.
	Name String
	// A default, hard-coded, or user-defined value for this variable.
	DefaultValue *String
	// A free text natural language description of the variable and its purpose.
	Description *String
	// The FHIRPath expression to evaluate against the fixture body. When variables are defined, only one of either expression, headerField or path must be specified.
	Expression *String
	// Will be used to grab the HTTP header field value from the headers that sourceId is pointing to.
	HeaderField *String
	// Displayable text string with hint help information to the user when entering a default value.
	Hint *String
	// XPath or JSONPath to evaluate against the fixture body.  When variables are defined, only one of either expression, headerField or path must be specified.
	Path *String
	// Fixture to evaluate the XPath/JSONPath expression or the headerField  against within this variable.
	SourceId *Id
}
type jsonTestScriptVariable struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Name                         String            `json:"name,omitempty"`
	NamePrimitiveElement         *primitiveElement `json:"_name,omitempty"`
	DefaultValue                 *String           `json:"defaultValue,omitempty"`
	DefaultValuePrimitiveElement *primitiveElement `json:"_defaultValue,omitempty"`
	Description                  *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement `json:"_description,omitempty"`
	Expression                   *String           `json:"expression,omitempty"`
	ExpressionPrimitiveElement   *primitiveElement `json:"_expression,omitempty"`
	HeaderField                  *String           `json:"headerField,omitempty"`
	HeaderFieldPrimitiveElement  *primitiveElement `json:"_headerField,omitempty"`
	Hint                         *String           `json:"hint,omitempty"`
	HintPrimitiveElement         *primitiveElement `json:"_hint,omitempty"`
	Path                         *String           `json:"path,omitempty"`
	PathPrimitiveElement         *primitiveElement `json:"_path,omitempty"`
	SourceId                     *Id               `json:"sourceId,omitempty"`
	SourceIdPrimitiveElement     *primitiveElement `json:"_sourceId,omitempty"`
}

func (r TestScriptVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptVariable) marshalJSON() jsonTestScriptVariable {
	m := jsonTestScriptVariable{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.DefaultValue != nil && r.DefaultValue.Value != nil {
		m.DefaultValue = r.DefaultValue
	}
	if r.DefaultValue != nil && (r.DefaultValue.Id != nil || r.DefaultValue.Extension != nil) {
		m.DefaultValuePrimitiveElement = &primitiveElement{Id: r.DefaultValue.Id, Extension: r.DefaultValue.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Expression != nil && r.Expression.Value != nil {
		m.Expression = r.Expression
	}
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	if r.HeaderField != nil && r.HeaderField.Value != nil {
		m.HeaderField = r.HeaderField
	}
	if r.HeaderField != nil && (r.HeaderField.Id != nil || r.HeaderField.Extension != nil) {
		m.HeaderFieldPrimitiveElement = &primitiveElement{Id: r.HeaderField.Id, Extension: r.HeaderField.Extension}
	}
	if r.Hint != nil && r.Hint.Value != nil {
		m.Hint = r.Hint
	}
	if r.Hint != nil && (r.Hint.Id != nil || r.Hint.Extension != nil) {
		m.HintPrimitiveElement = &primitiveElement{Id: r.Hint.Id, Extension: r.Hint.Extension}
	}
	if r.Path != nil && r.Path.Value != nil {
		m.Path = r.Path
	}
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	if r.SourceId != nil && r.SourceId.Value != nil {
		m.SourceId = r.SourceId
	}
	if r.SourceId != nil && (r.SourceId.Id != nil || r.SourceId.Extension != nil) {
		m.SourceIdPrimitiveElement = &primitiveElement{Id: r.SourceId.Id, Extension: r.SourceId.Extension}
	}
	return m
}
func (r *TestScriptVariable) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptVariable
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptVariable) unmarshalJSON(m jsonTestScriptVariable) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.DefaultValue = m.DefaultValue
	if m.DefaultValuePrimitiveElement != nil {
		if r.DefaultValue == nil {
			r.DefaultValue = &String{}
		}
		r.DefaultValue.Id = m.DefaultValuePrimitiveElement.Id
		r.DefaultValue.Extension = m.DefaultValuePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		if r.Expression == nil {
			r.Expression = &String{}
		}
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	r.HeaderField = m.HeaderField
	if m.HeaderFieldPrimitiveElement != nil {
		if r.HeaderField == nil {
			r.HeaderField = &String{}
		}
		r.HeaderField.Id = m.HeaderFieldPrimitiveElement.Id
		r.HeaderField.Extension = m.HeaderFieldPrimitiveElement.Extension
	}
	r.Hint = m.Hint
	if m.HintPrimitiveElement != nil {
		if r.Hint == nil {
			r.Hint = &String{}
		}
		r.Hint.Id = m.HintPrimitiveElement.Id
		r.Hint.Extension = m.HintPrimitiveElement.Extension
	}
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		if r.Path == nil {
			r.Path = &String{}
		}
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.SourceId = m.SourceId
	if m.SourceIdPrimitiveElement != nil {
		if r.SourceId == nil {
			r.SourceId = &Id{}
		}
		r.SourceId.Id = m.SourceIdPrimitiveElement.Id
		r.SourceId.Extension = m.SourceIdPrimitiveElement.Extension
	}
	return nil
}
func (r TestScriptVariable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DefaultValue, xml.StartElement{Name: xml.Name{Local: "defaultValue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Expression, xml.StartElement{Name: xml.Name{Local: "expression"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HeaderField, xml.StartElement{Name: xml.Name{Local: "headerField"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Hint, xml.StartElement{Name: xml.Name{Local: "hint"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Path, xml.StartElement{Name: xml.Name{Local: "path"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceId, xml.StartElement{Name: xml.Name{Local: "sourceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptVariable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "defaultValue":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DefaultValue = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "expression":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Expression = &v
			case "headerField":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HeaderField = &v
			case "hint":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Hint = &v
			case "path":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Path = &v
			case "sourceId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceId = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptVariable) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A series of required setup operations before tests are executed.
type TestScriptSetup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Action would contain either an operation or an assertion.
	Action []TestScriptSetupAction
}
type jsonTestScriptSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `json:"action,omitempty"`
}

func (r TestScriptSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptSetup) marshalJSON() jsonTestScriptSetup {
	m := jsonTestScriptSetup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	return m
}
func (r *TestScriptSetup) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptSetup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptSetup) unmarshalJSON(m jsonTestScriptSetup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	return nil
}
func (r TestScriptSetup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptSetup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "action":
				var v TestScriptSetupAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptSetup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Action would contain either an operation or an assertion.
type TestScriptSetupAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The operation to perform.
	Operation *TestScriptSetupActionOperation
	// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
	Assert *TestScriptSetupActionAssert
}
type jsonTestScriptSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

func (r TestScriptSetupAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptSetupAction) marshalJSON() jsonTestScriptSetupAction {
	m := jsonTestScriptSetupAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	m.Assert = r.Assert
	return m
}
func (r *TestScriptSetupAction) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptSetupAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptSetupAction) unmarshalJSON(m jsonTestScriptSetupAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	r.Assert = m.Assert
	return nil
}
func (r TestScriptSetupAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Operation, xml.StartElement{Name: xml.Name{Local: "operation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Assert, xml.StartElement{Name: xml.Name{Local: "assert"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptSetupAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "operation":
				var v TestScriptSetupActionOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = &v
			case "assert":
				var v TestScriptSetupActionAssert
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Assert = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptSetupAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The operation to perform.
type TestScriptSetupActionOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Server interaction or operation type.
	Type *Coding
	// The type of the resource.  See http://build.fhir.org/resourcelist.html.
	Resource *Code
	// The label would be used for tracking/logging purposes by test engines.
	Label *String
	// The description would be used by test engines for tracking and reporting purposes.
	Description *String
	// The mime-type to use for RESTful operation in the 'Accept' header.
	Accept *Code
	// The mime-type to use for RESTful operation in the 'Content-Type' header.
	ContentType *Code
	// The server where the request message is destined for.  Must be one of the server numbers listed in TestScript.destination section.
	Destination *Integer
	// Whether or not to implicitly send the request url in encoded format. The default is true to match the standard RESTful client behavior. Set to false when communicating with a server that does not support encoded url paths.
	EncodeRequestUrl Boolean
	// The HTTP method the test engine MUST use for this operation regardless of any other operation details.
	Method *Code
	// The server where the request message originates from.  Must be one of the server numbers listed in TestScript.origin section.
	Origin *Integer
	// Path plus parameters after [type].  Used to set parts of the request URL explicitly.
	Params *String
	// Header elements would be used to set HTTP headers.
	RequestHeader []TestScriptSetupActionOperationRequestHeader
	// The fixture id (maybe new) to map to the request.
	RequestId *Id
	// The fixture id (maybe new) to map to the response.
	ResponseId *Id
	// The id of the fixture used as the body of a PUT or POST request.
	SourceId *Id
	// Id of fixture used for extracting the [id],  [type], and [vid] for GET requests.
	TargetId *Id
	// Complete request URL.
	Url *String
}
type jsonTestScriptSetupActionOperation struct {
	Id                               *string                                       `json:"id,omitempty"`
	Extension                        []Extension                                   `json:"extension,omitempty"`
	ModifierExtension                []Extension                                   `json:"modifierExtension,omitempty"`
	Type                             *Coding                                       `json:"type,omitempty"`
	Resource                         *Code                                         `json:"resource,omitempty"`
	ResourcePrimitiveElement         *primitiveElement                             `json:"_resource,omitempty"`
	Label                            *String                                       `json:"label,omitempty"`
	LabelPrimitiveElement            *primitiveElement                             `json:"_label,omitempty"`
	Description                      *String                                       `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement                             `json:"_description,omitempty"`
	Accept                           *Code                                         `json:"accept,omitempty"`
	AcceptPrimitiveElement           *primitiveElement                             `json:"_accept,omitempty"`
	ContentType                      *Code                                         `json:"contentType,omitempty"`
	ContentTypePrimitiveElement      *primitiveElement                             `json:"_contentType,omitempty"`
	Destination                      *Integer                                      `json:"destination,omitempty"`
	DestinationPrimitiveElement      *primitiveElement                             `json:"_destination,omitempty"`
	EncodeRequestUrl                 Boolean                                       `json:"encodeRequestUrl,omitempty"`
	EncodeRequestUrlPrimitiveElement *primitiveElement                             `json:"_encodeRequestUrl,omitempty"`
	Method                           *Code                                         `json:"method,omitempty"`
	MethodPrimitiveElement           *primitiveElement                             `json:"_method,omitempty"`
	Origin                           *Integer                                      `json:"origin,omitempty"`
	OriginPrimitiveElement           *primitiveElement                             `json:"_origin,omitempty"`
	Params                           *String                                       `json:"params,omitempty"`
	ParamsPrimitiveElement           *primitiveElement                             `json:"_params,omitempty"`
	RequestHeader                    []TestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`
	RequestId                        *Id                                           `json:"requestId,omitempty"`
	RequestIdPrimitiveElement        *primitiveElement                             `json:"_requestId,omitempty"`
	ResponseId                       *Id                                           `json:"responseId,omitempty"`
	ResponseIdPrimitiveElement       *primitiveElement                             `json:"_responseId,omitempty"`
	SourceId                         *Id                                           `json:"sourceId,omitempty"`
	SourceIdPrimitiveElement         *primitiveElement                             `json:"_sourceId,omitempty"`
	TargetId                         *Id                                           `json:"targetId,omitempty"`
	TargetIdPrimitiveElement         *primitiveElement                             `json:"_targetId,omitempty"`
	Url                              *String                                       `json:"url,omitempty"`
	UrlPrimitiveElement              *primitiveElement                             `json:"_url,omitempty"`
}

func (r TestScriptSetupActionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptSetupActionOperation) marshalJSON() jsonTestScriptSetupActionOperation {
	m := jsonTestScriptSetupActionOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Resource != nil && r.Resource.Value != nil {
		m.Resource = r.Resource
	}
	if r.Resource != nil && (r.Resource.Id != nil || r.Resource.Extension != nil) {
		m.ResourcePrimitiveElement = &primitiveElement{Id: r.Resource.Id, Extension: r.Resource.Extension}
	}
	if r.Label != nil && r.Label.Value != nil {
		m.Label = r.Label
	}
	if r.Label != nil && (r.Label.Id != nil || r.Label.Extension != nil) {
		m.LabelPrimitiveElement = &primitiveElement{Id: r.Label.Id, Extension: r.Label.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Accept != nil && r.Accept.Value != nil {
		m.Accept = r.Accept
	}
	if r.Accept != nil && (r.Accept.Id != nil || r.Accept.Extension != nil) {
		m.AcceptPrimitiveElement = &primitiveElement{Id: r.Accept.Id, Extension: r.Accept.Extension}
	}
	if r.ContentType != nil && r.ContentType.Value != nil {
		m.ContentType = r.ContentType
	}
	if r.ContentType != nil && (r.ContentType.Id != nil || r.ContentType.Extension != nil) {
		m.ContentTypePrimitiveElement = &primitiveElement{Id: r.ContentType.Id, Extension: r.ContentType.Extension}
	}
	if r.Destination != nil && r.Destination.Value != nil {
		m.Destination = r.Destination
	}
	if r.Destination != nil && (r.Destination.Id != nil || r.Destination.Extension != nil) {
		m.DestinationPrimitiveElement = &primitiveElement{Id: r.Destination.Id, Extension: r.Destination.Extension}
	}
	if r.EncodeRequestUrl.Value != nil {
		m.EncodeRequestUrl = r.EncodeRequestUrl
	}
	if r.EncodeRequestUrl.Id != nil || r.EncodeRequestUrl.Extension != nil {
		m.EncodeRequestUrlPrimitiveElement = &primitiveElement{Id: r.EncodeRequestUrl.Id, Extension: r.EncodeRequestUrl.Extension}
	}
	if r.Method != nil && r.Method.Value != nil {
		m.Method = r.Method
	}
	if r.Method != nil && (r.Method.Id != nil || r.Method.Extension != nil) {
		m.MethodPrimitiveElement = &primitiveElement{Id: r.Method.Id, Extension: r.Method.Extension}
	}
	if r.Origin != nil && r.Origin.Value != nil {
		m.Origin = r.Origin
	}
	if r.Origin != nil && (r.Origin.Id != nil || r.Origin.Extension != nil) {
		m.OriginPrimitiveElement = &primitiveElement{Id: r.Origin.Id, Extension: r.Origin.Extension}
	}
	if r.Params != nil && r.Params.Value != nil {
		m.Params = r.Params
	}
	if r.Params != nil && (r.Params.Id != nil || r.Params.Extension != nil) {
		m.ParamsPrimitiveElement = &primitiveElement{Id: r.Params.Id, Extension: r.Params.Extension}
	}
	m.RequestHeader = r.RequestHeader
	if r.RequestId != nil && r.RequestId.Value != nil {
		m.RequestId = r.RequestId
	}
	if r.RequestId != nil && (r.RequestId.Id != nil || r.RequestId.Extension != nil) {
		m.RequestIdPrimitiveElement = &primitiveElement{Id: r.RequestId.Id, Extension: r.RequestId.Extension}
	}
	if r.ResponseId != nil && r.ResponseId.Value != nil {
		m.ResponseId = r.ResponseId
	}
	if r.ResponseId != nil && (r.ResponseId.Id != nil || r.ResponseId.Extension != nil) {
		m.ResponseIdPrimitiveElement = &primitiveElement{Id: r.ResponseId.Id, Extension: r.ResponseId.Extension}
	}
	if r.SourceId != nil && r.SourceId.Value != nil {
		m.SourceId = r.SourceId
	}
	if r.SourceId != nil && (r.SourceId.Id != nil || r.SourceId.Extension != nil) {
		m.SourceIdPrimitiveElement = &primitiveElement{Id: r.SourceId.Id, Extension: r.SourceId.Extension}
	}
	if r.TargetId != nil && r.TargetId.Value != nil {
		m.TargetId = r.TargetId
	}
	if r.TargetId != nil && (r.TargetId.Id != nil || r.TargetId.Extension != nil) {
		m.TargetIdPrimitiveElement = &primitiveElement{Id: r.TargetId.Id, Extension: r.TargetId.Extension}
	}
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	return m
}
func (r *TestScriptSetupActionOperation) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptSetupActionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptSetupActionOperation) unmarshalJSON(m jsonTestScriptSetupActionOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Resource = m.Resource
	if m.ResourcePrimitiveElement != nil {
		if r.Resource == nil {
			r.Resource = &Code{}
		}
		r.Resource.Id = m.ResourcePrimitiveElement.Id
		r.Resource.Extension = m.ResourcePrimitiveElement.Extension
	}
	r.Label = m.Label
	if m.LabelPrimitiveElement != nil {
		if r.Label == nil {
			r.Label = &String{}
		}
		r.Label.Id = m.LabelPrimitiveElement.Id
		r.Label.Extension = m.LabelPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Accept = m.Accept
	if m.AcceptPrimitiveElement != nil {
		if r.Accept == nil {
			r.Accept = &Code{}
		}
		r.Accept.Id = m.AcceptPrimitiveElement.Id
		r.Accept.Extension = m.AcceptPrimitiveElement.Extension
	}
	r.ContentType = m.ContentType
	if m.ContentTypePrimitiveElement != nil {
		if r.ContentType == nil {
			r.ContentType = &Code{}
		}
		r.ContentType.Id = m.ContentTypePrimitiveElement.Id
		r.ContentType.Extension = m.ContentTypePrimitiveElement.Extension
	}
	r.Destination = m.Destination
	if m.DestinationPrimitiveElement != nil {
		if r.Destination == nil {
			r.Destination = &Integer{}
		}
		r.Destination.Id = m.DestinationPrimitiveElement.Id
		r.Destination.Extension = m.DestinationPrimitiveElement.Extension
	}
	r.EncodeRequestUrl = m.EncodeRequestUrl
	if m.EncodeRequestUrlPrimitiveElement != nil {
		r.EncodeRequestUrl.Id = m.EncodeRequestUrlPrimitiveElement.Id
		r.EncodeRequestUrl.Extension = m.EncodeRequestUrlPrimitiveElement.Extension
	}
	r.Method = m.Method
	if m.MethodPrimitiveElement != nil {
		if r.Method == nil {
			r.Method = &Code{}
		}
		r.Method.Id = m.MethodPrimitiveElement.Id
		r.Method.Extension = m.MethodPrimitiveElement.Extension
	}
	r.Origin = m.Origin
	if m.OriginPrimitiveElement != nil {
		if r.Origin == nil {
			r.Origin = &Integer{}
		}
		r.Origin.Id = m.OriginPrimitiveElement.Id
		r.Origin.Extension = m.OriginPrimitiveElement.Extension
	}
	r.Params = m.Params
	if m.ParamsPrimitiveElement != nil {
		if r.Params == nil {
			r.Params = &String{}
		}
		r.Params.Id = m.ParamsPrimitiveElement.Id
		r.Params.Extension = m.ParamsPrimitiveElement.Extension
	}
	r.RequestHeader = m.RequestHeader
	r.RequestId = m.RequestId
	if m.RequestIdPrimitiveElement != nil {
		if r.RequestId == nil {
			r.RequestId = &Id{}
		}
		r.RequestId.Id = m.RequestIdPrimitiveElement.Id
		r.RequestId.Extension = m.RequestIdPrimitiveElement.Extension
	}
	r.ResponseId = m.ResponseId
	if m.ResponseIdPrimitiveElement != nil {
		if r.ResponseId == nil {
			r.ResponseId = &Id{}
		}
		r.ResponseId.Id = m.ResponseIdPrimitiveElement.Id
		r.ResponseId.Extension = m.ResponseIdPrimitiveElement.Extension
	}
	r.SourceId = m.SourceId
	if m.SourceIdPrimitiveElement != nil {
		if r.SourceId == nil {
			r.SourceId = &Id{}
		}
		r.SourceId.Id = m.SourceIdPrimitiveElement.Id
		r.SourceId.Extension = m.SourceIdPrimitiveElement.Extension
	}
	r.TargetId = m.TargetId
	if m.TargetIdPrimitiveElement != nil {
		if r.TargetId == nil {
			r.TargetId = &Id{}
		}
		r.TargetId.Id = m.TargetIdPrimitiveElement.Id
		r.TargetId.Extension = m.TargetIdPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &String{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	return nil
}
func (r TestScriptSetupActionOperation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Label, xml.StartElement{Name: xml.Name{Local: "label"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Accept, xml.StartElement{Name: xml.Name{Local: "accept"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContentType, xml.StartElement{Name: xml.Name{Local: "contentType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Destination, xml.StartElement{Name: xml.Name{Local: "destination"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EncodeRequestUrl, xml.StartElement{Name: xml.Name{Local: "encodeRequestUrl"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Origin, xml.StartElement{Name: xml.Name{Local: "origin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Params, xml.StartElement{Name: xml.Name{Local: "params"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestHeader, xml.StartElement{Name: xml.Name{Local: "requestHeader"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestId, xml.StartElement{Name: xml.Name{Local: "requestId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ResponseId, xml.StartElement{Name: xml.Name{Local: "responseId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceId, xml.StartElement{Name: xml.Name{Local: "sourceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetId, xml.StartElement{Name: xml.Name{Local: "targetId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptSetupActionOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "resource":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = &v
			case "label":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Label = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "accept":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Accept = &v
			case "contentType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContentType = &v
			case "destination":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Destination = &v
			case "encodeRequestUrl":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EncodeRequestUrl = v
			case "method":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "origin":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Origin = &v
			case "params":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Params = &v
			case "requestHeader":
				var v TestScriptSetupActionOperationRequestHeader
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestHeader = append(r.RequestHeader, v)
			case "requestId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestId = &v
			case "responseId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResponseId = &v
			case "sourceId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceId = &v
			case "targetId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetId = &v
			case "url":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptSetupActionOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Header elements would be used to set HTTP headers.
type TestScriptSetupActionOperationRequestHeader struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The HTTP header field e.g. "Accept".
	Field String
	// The value of the header e.g. "application/fhir+xml".
	Value String
}
type jsonTestScriptSetupActionOperationRequestHeader struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Field                 String            `json:"field,omitempty"`
	FieldPrimitiveElement *primitiveElement `json:"_field,omitempty"`
	Value                 String            `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r TestScriptSetupActionOperationRequestHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptSetupActionOperationRequestHeader) marshalJSON() jsonTestScriptSetupActionOperationRequestHeader {
	m := jsonTestScriptSetupActionOperationRequestHeader{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Field.Value != nil {
		m.Field = r.Field
	}
	if r.Field.Id != nil || r.Field.Extension != nil {
		m.FieldPrimitiveElement = &primitiveElement{Id: r.Field.Id, Extension: r.Field.Extension}
	}
	if r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *TestScriptSetupActionOperationRequestHeader) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptSetupActionOperationRequestHeader
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptSetupActionOperationRequestHeader) unmarshalJSON(m jsonTestScriptSetupActionOperationRequestHeader) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Field = m.Field
	if m.FieldPrimitiveElement != nil {
		r.Field.Id = m.FieldPrimitiveElement.Id
		r.Field.Extension = m.FieldPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r TestScriptSetupActionOperationRequestHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Field, xml.StartElement{Name: xml.Name{Local: "field"}})
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
func (r *TestScriptSetupActionOperationRequestHeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "field":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Field = v
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
func (r TestScriptSetupActionOperationRequestHeader) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
type TestScriptSetupActionAssert struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The label would be used for tracking/logging purposes by test engines.
	Label *String
	// The description would be used by test engines for tracking and reporting purposes.
	Description *String
	// The direction to use for the assertion.
	Direction *Code
	// Id of the source fixture used as the contents to be evaluated by either the "source/expression" or "sourceId/path" definition.
	CompareToSourceId *String
	// The FHIRPath expression to evaluate against the source fixture. When compareToSourceId is defined, either compareToSourceExpression or compareToSourcePath must be defined, but not both.
	CompareToSourceExpression *String
	// XPath or JSONPath expression to evaluate against the source fixture. When compareToSourceId is defined, either compareToSourceExpression or compareToSourcePath must be defined, but not both.
	CompareToSourcePath *String
	// The mime-type contents to compare against the request or response message 'Content-Type' header.
	ContentType *Code
	// The FHIRPath expression to be evaluated against the request or response message contents - HTTP headers and payload.
	Expression *String
	// The HTTP header field name e.g. 'Location'.
	HeaderField *String
	// The ID of a fixture.  Asserts that the response contains at a minimum the fixture specified by minimumId.
	MinimumId *String
	// Whether or not the test execution performs validation on the bundle navigation links.
	NavigationLinks *Boolean
	// The operator type defines the conditional behavior of the assert. If not defined, the default is equals.
	Operator *Code
	// The XPath or JSONPath expression to be evaluated against the fixture representing the response received from server.
	Path *String
	// The request method or HTTP operation code to compare against that used by the client system under test.
	RequestMethod *Code
	// The value to use in a comparison against the request URL path string.
	RequestUrl *String
	// The type of the resource.  See http://build.fhir.org/resourcelist.html.
	Resource *Code
	// okay | created | noContent | notModified | bad | forbidden | notFound | methodNotAllowed | conflict | gone | preconditionFailed | unprocessable.
	Response *Code
	// The value of the HTTP response code to be tested.
	ResponseCode *String
	// Fixture to evaluate the XPath/JSONPath expression or the headerField  against.
	SourceId *Id
	// The ID of the Profile to validate against.
	ValidateProfileId *Id
	// The value to compare to.
	Value *String
	// Whether or not the test execution will produce a warning only on error for this assert.
	WarningOnly Boolean
}
type jsonTestScriptSetupActionAssert struct {
	Id                                        *string           `json:"id,omitempty"`
	Extension                                 []Extension       `json:"extension,omitempty"`
	ModifierExtension                         []Extension       `json:"modifierExtension,omitempty"`
	Label                                     *String           `json:"label,omitempty"`
	LabelPrimitiveElement                     *primitiveElement `json:"_label,omitempty"`
	Description                               *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement               *primitiveElement `json:"_description,omitempty"`
	Direction                                 *Code             `json:"direction,omitempty"`
	DirectionPrimitiveElement                 *primitiveElement `json:"_direction,omitempty"`
	CompareToSourceId                         *String           `json:"compareToSourceId,omitempty"`
	CompareToSourceIdPrimitiveElement         *primitiveElement `json:"_compareToSourceId,omitempty"`
	CompareToSourceExpression                 *String           `json:"compareToSourceExpression,omitempty"`
	CompareToSourceExpressionPrimitiveElement *primitiveElement `json:"_compareToSourceExpression,omitempty"`
	CompareToSourcePath                       *String           `json:"compareToSourcePath,omitempty"`
	CompareToSourcePathPrimitiveElement       *primitiveElement `json:"_compareToSourcePath,omitempty"`
	ContentType                               *Code             `json:"contentType,omitempty"`
	ContentTypePrimitiveElement               *primitiveElement `json:"_contentType,omitempty"`
	Expression                                *String           `json:"expression,omitempty"`
	ExpressionPrimitiveElement                *primitiveElement `json:"_expression,omitempty"`
	HeaderField                               *String           `json:"headerField,omitempty"`
	HeaderFieldPrimitiveElement               *primitiveElement `json:"_headerField,omitempty"`
	MinimumId                                 *String           `json:"minimumId,omitempty"`
	MinimumIdPrimitiveElement                 *primitiveElement `json:"_minimumId,omitempty"`
	NavigationLinks                           *Boolean          `json:"navigationLinks,omitempty"`
	NavigationLinksPrimitiveElement           *primitiveElement `json:"_navigationLinks,omitempty"`
	Operator                                  *Code             `json:"operator,omitempty"`
	OperatorPrimitiveElement                  *primitiveElement `json:"_operator,omitempty"`
	Path                                      *String           `json:"path,omitempty"`
	PathPrimitiveElement                      *primitiveElement `json:"_path,omitempty"`
	RequestMethod                             *Code             `json:"requestMethod,omitempty"`
	RequestMethodPrimitiveElement             *primitiveElement `json:"_requestMethod,omitempty"`
	RequestUrl                                *String           `json:"requestURL,omitempty"`
	RequestUrlPrimitiveElement                *primitiveElement `json:"_requestURL,omitempty"`
	Resource                                  *Code             `json:"resource,omitempty"`
	ResourcePrimitiveElement                  *primitiveElement `json:"_resource,omitempty"`
	Response                                  *Code             `json:"response,omitempty"`
	ResponsePrimitiveElement                  *primitiveElement `json:"_response,omitempty"`
	ResponseCode                              *String           `json:"responseCode,omitempty"`
	ResponseCodePrimitiveElement              *primitiveElement `json:"_responseCode,omitempty"`
	SourceId                                  *Id               `json:"sourceId,omitempty"`
	SourceIdPrimitiveElement                  *primitiveElement `json:"_sourceId,omitempty"`
	ValidateProfileId                         *Id               `json:"validateProfileId,omitempty"`
	ValidateProfileIdPrimitiveElement         *primitiveElement `json:"_validateProfileId,omitempty"`
	Value                                     *String           `json:"value,omitempty"`
	ValuePrimitiveElement                     *primitiveElement `json:"_value,omitempty"`
	WarningOnly                               Boolean           `json:"warningOnly,omitempty"`
	WarningOnlyPrimitiveElement               *primitiveElement `json:"_warningOnly,omitempty"`
}

func (r TestScriptSetupActionAssert) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptSetupActionAssert) marshalJSON() jsonTestScriptSetupActionAssert {
	m := jsonTestScriptSetupActionAssert{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Label != nil && r.Label.Value != nil {
		m.Label = r.Label
	}
	if r.Label != nil && (r.Label.Id != nil || r.Label.Extension != nil) {
		m.LabelPrimitiveElement = &primitiveElement{Id: r.Label.Id, Extension: r.Label.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Direction != nil && r.Direction.Value != nil {
		m.Direction = r.Direction
	}
	if r.Direction != nil && (r.Direction.Id != nil || r.Direction.Extension != nil) {
		m.DirectionPrimitiveElement = &primitiveElement{Id: r.Direction.Id, Extension: r.Direction.Extension}
	}
	if r.CompareToSourceId != nil && r.CompareToSourceId.Value != nil {
		m.CompareToSourceId = r.CompareToSourceId
	}
	if r.CompareToSourceId != nil && (r.CompareToSourceId.Id != nil || r.CompareToSourceId.Extension != nil) {
		m.CompareToSourceIdPrimitiveElement = &primitiveElement{Id: r.CompareToSourceId.Id, Extension: r.CompareToSourceId.Extension}
	}
	if r.CompareToSourceExpression != nil && r.CompareToSourceExpression.Value != nil {
		m.CompareToSourceExpression = r.CompareToSourceExpression
	}
	if r.CompareToSourceExpression != nil && (r.CompareToSourceExpression.Id != nil || r.CompareToSourceExpression.Extension != nil) {
		m.CompareToSourceExpressionPrimitiveElement = &primitiveElement{Id: r.CompareToSourceExpression.Id, Extension: r.CompareToSourceExpression.Extension}
	}
	if r.CompareToSourcePath != nil && r.CompareToSourcePath.Value != nil {
		m.CompareToSourcePath = r.CompareToSourcePath
	}
	if r.CompareToSourcePath != nil && (r.CompareToSourcePath.Id != nil || r.CompareToSourcePath.Extension != nil) {
		m.CompareToSourcePathPrimitiveElement = &primitiveElement{Id: r.CompareToSourcePath.Id, Extension: r.CompareToSourcePath.Extension}
	}
	if r.ContentType != nil && r.ContentType.Value != nil {
		m.ContentType = r.ContentType
	}
	if r.ContentType != nil && (r.ContentType.Id != nil || r.ContentType.Extension != nil) {
		m.ContentTypePrimitiveElement = &primitiveElement{Id: r.ContentType.Id, Extension: r.ContentType.Extension}
	}
	if r.Expression != nil && r.Expression.Value != nil {
		m.Expression = r.Expression
	}
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	if r.HeaderField != nil && r.HeaderField.Value != nil {
		m.HeaderField = r.HeaderField
	}
	if r.HeaderField != nil && (r.HeaderField.Id != nil || r.HeaderField.Extension != nil) {
		m.HeaderFieldPrimitiveElement = &primitiveElement{Id: r.HeaderField.Id, Extension: r.HeaderField.Extension}
	}
	if r.MinimumId != nil && r.MinimumId.Value != nil {
		m.MinimumId = r.MinimumId
	}
	if r.MinimumId != nil && (r.MinimumId.Id != nil || r.MinimumId.Extension != nil) {
		m.MinimumIdPrimitiveElement = &primitiveElement{Id: r.MinimumId.Id, Extension: r.MinimumId.Extension}
	}
	if r.NavigationLinks != nil && r.NavigationLinks.Value != nil {
		m.NavigationLinks = r.NavigationLinks
	}
	if r.NavigationLinks != nil && (r.NavigationLinks.Id != nil || r.NavigationLinks.Extension != nil) {
		m.NavigationLinksPrimitiveElement = &primitiveElement{Id: r.NavigationLinks.Id, Extension: r.NavigationLinks.Extension}
	}
	if r.Operator != nil && r.Operator.Value != nil {
		m.Operator = r.Operator
	}
	if r.Operator != nil && (r.Operator.Id != nil || r.Operator.Extension != nil) {
		m.OperatorPrimitiveElement = &primitiveElement{Id: r.Operator.Id, Extension: r.Operator.Extension}
	}
	if r.Path != nil && r.Path.Value != nil {
		m.Path = r.Path
	}
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	if r.RequestMethod != nil && r.RequestMethod.Value != nil {
		m.RequestMethod = r.RequestMethod
	}
	if r.RequestMethod != nil && (r.RequestMethod.Id != nil || r.RequestMethod.Extension != nil) {
		m.RequestMethodPrimitiveElement = &primitiveElement{Id: r.RequestMethod.Id, Extension: r.RequestMethod.Extension}
	}
	if r.RequestUrl != nil && r.RequestUrl.Value != nil {
		m.RequestUrl = r.RequestUrl
	}
	if r.RequestUrl != nil && (r.RequestUrl.Id != nil || r.RequestUrl.Extension != nil) {
		m.RequestUrlPrimitiveElement = &primitiveElement{Id: r.RequestUrl.Id, Extension: r.RequestUrl.Extension}
	}
	if r.Resource != nil && r.Resource.Value != nil {
		m.Resource = r.Resource
	}
	if r.Resource != nil && (r.Resource.Id != nil || r.Resource.Extension != nil) {
		m.ResourcePrimitiveElement = &primitiveElement{Id: r.Resource.Id, Extension: r.Resource.Extension}
	}
	if r.Response != nil && r.Response.Value != nil {
		m.Response = r.Response
	}
	if r.Response != nil && (r.Response.Id != nil || r.Response.Extension != nil) {
		m.ResponsePrimitiveElement = &primitiveElement{Id: r.Response.Id, Extension: r.Response.Extension}
	}
	if r.ResponseCode != nil && r.ResponseCode.Value != nil {
		m.ResponseCode = r.ResponseCode
	}
	if r.ResponseCode != nil && (r.ResponseCode.Id != nil || r.ResponseCode.Extension != nil) {
		m.ResponseCodePrimitiveElement = &primitiveElement{Id: r.ResponseCode.Id, Extension: r.ResponseCode.Extension}
	}
	if r.SourceId != nil && r.SourceId.Value != nil {
		m.SourceId = r.SourceId
	}
	if r.SourceId != nil && (r.SourceId.Id != nil || r.SourceId.Extension != nil) {
		m.SourceIdPrimitiveElement = &primitiveElement{Id: r.SourceId.Id, Extension: r.SourceId.Extension}
	}
	if r.ValidateProfileId != nil && r.ValidateProfileId.Value != nil {
		m.ValidateProfileId = r.ValidateProfileId
	}
	if r.ValidateProfileId != nil && (r.ValidateProfileId.Id != nil || r.ValidateProfileId.Extension != nil) {
		m.ValidateProfileIdPrimitiveElement = &primitiveElement{Id: r.ValidateProfileId.Id, Extension: r.ValidateProfileId.Extension}
	}
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	if r.WarningOnly.Value != nil {
		m.WarningOnly = r.WarningOnly
	}
	if r.WarningOnly.Id != nil || r.WarningOnly.Extension != nil {
		m.WarningOnlyPrimitiveElement = &primitiveElement{Id: r.WarningOnly.Id, Extension: r.WarningOnly.Extension}
	}
	return m
}
func (r *TestScriptSetupActionAssert) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptSetupActionAssert
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptSetupActionAssert) unmarshalJSON(m jsonTestScriptSetupActionAssert) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Label = m.Label
	if m.LabelPrimitiveElement != nil {
		if r.Label == nil {
			r.Label = &String{}
		}
		r.Label.Id = m.LabelPrimitiveElement.Id
		r.Label.Extension = m.LabelPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Direction = m.Direction
	if m.DirectionPrimitiveElement != nil {
		if r.Direction == nil {
			r.Direction = &Code{}
		}
		r.Direction.Id = m.DirectionPrimitiveElement.Id
		r.Direction.Extension = m.DirectionPrimitiveElement.Extension
	}
	r.CompareToSourceId = m.CompareToSourceId
	if m.CompareToSourceIdPrimitiveElement != nil {
		if r.CompareToSourceId == nil {
			r.CompareToSourceId = &String{}
		}
		r.CompareToSourceId.Id = m.CompareToSourceIdPrimitiveElement.Id
		r.CompareToSourceId.Extension = m.CompareToSourceIdPrimitiveElement.Extension
	}
	r.CompareToSourceExpression = m.CompareToSourceExpression
	if m.CompareToSourceExpressionPrimitiveElement != nil {
		if r.CompareToSourceExpression == nil {
			r.CompareToSourceExpression = &String{}
		}
		r.CompareToSourceExpression.Id = m.CompareToSourceExpressionPrimitiveElement.Id
		r.CompareToSourceExpression.Extension = m.CompareToSourceExpressionPrimitiveElement.Extension
	}
	r.CompareToSourcePath = m.CompareToSourcePath
	if m.CompareToSourcePathPrimitiveElement != nil {
		if r.CompareToSourcePath == nil {
			r.CompareToSourcePath = &String{}
		}
		r.CompareToSourcePath.Id = m.CompareToSourcePathPrimitiveElement.Id
		r.CompareToSourcePath.Extension = m.CompareToSourcePathPrimitiveElement.Extension
	}
	r.ContentType = m.ContentType
	if m.ContentTypePrimitiveElement != nil {
		if r.ContentType == nil {
			r.ContentType = &Code{}
		}
		r.ContentType.Id = m.ContentTypePrimitiveElement.Id
		r.ContentType.Extension = m.ContentTypePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		if r.Expression == nil {
			r.Expression = &String{}
		}
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	r.HeaderField = m.HeaderField
	if m.HeaderFieldPrimitiveElement != nil {
		if r.HeaderField == nil {
			r.HeaderField = &String{}
		}
		r.HeaderField.Id = m.HeaderFieldPrimitiveElement.Id
		r.HeaderField.Extension = m.HeaderFieldPrimitiveElement.Extension
	}
	r.MinimumId = m.MinimumId
	if m.MinimumIdPrimitiveElement != nil {
		if r.MinimumId == nil {
			r.MinimumId = &String{}
		}
		r.MinimumId.Id = m.MinimumIdPrimitiveElement.Id
		r.MinimumId.Extension = m.MinimumIdPrimitiveElement.Extension
	}
	r.NavigationLinks = m.NavigationLinks
	if m.NavigationLinksPrimitiveElement != nil {
		if r.NavigationLinks == nil {
			r.NavigationLinks = &Boolean{}
		}
		r.NavigationLinks.Id = m.NavigationLinksPrimitiveElement.Id
		r.NavigationLinks.Extension = m.NavigationLinksPrimitiveElement.Extension
	}
	r.Operator = m.Operator
	if m.OperatorPrimitiveElement != nil {
		if r.Operator == nil {
			r.Operator = &Code{}
		}
		r.Operator.Id = m.OperatorPrimitiveElement.Id
		r.Operator.Extension = m.OperatorPrimitiveElement.Extension
	}
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		if r.Path == nil {
			r.Path = &String{}
		}
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.RequestMethod = m.RequestMethod
	if m.RequestMethodPrimitiveElement != nil {
		if r.RequestMethod == nil {
			r.RequestMethod = &Code{}
		}
		r.RequestMethod.Id = m.RequestMethodPrimitiveElement.Id
		r.RequestMethod.Extension = m.RequestMethodPrimitiveElement.Extension
	}
	r.RequestUrl = m.RequestUrl
	if m.RequestUrlPrimitiveElement != nil {
		if r.RequestUrl == nil {
			r.RequestUrl = &String{}
		}
		r.RequestUrl.Id = m.RequestUrlPrimitiveElement.Id
		r.RequestUrl.Extension = m.RequestUrlPrimitiveElement.Extension
	}
	r.Resource = m.Resource
	if m.ResourcePrimitiveElement != nil {
		if r.Resource == nil {
			r.Resource = &Code{}
		}
		r.Resource.Id = m.ResourcePrimitiveElement.Id
		r.Resource.Extension = m.ResourcePrimitiveElement.Extension
	}
	r.Response = m.Response
	if m.ResponsePrimitiveElement != nil {
		if r.Response == nil {
			r.Response = &Code{}
		}
		r.Response.Id = m.ResponsePrimitiveElement.Id
		r.Response.Extension = m.ResponsePrimitiveElement.Extension
	}
	r.ResponseCode = m.ResponseCode
	if m.ResponseCodePrimitiveElement != nil {
		if r.ResponseCode == nil {
			r.ResponseCode = &String{}
		}
		r.ResponseCode.Id = m.ResponseCodePrimitiveElement.Id
		r.ResponseCode.Extension = m.ResponseCodePrimitiveElement.Extension
	}
	r.SourceId = m.SourceId
	if m.SourceIdPrimitiveElement != nil {
		if r.SourceId == nil {
			r.SourceId = &Id{}
		}
		r.SourceId.Id = m.SourceIdPrimitiveElement.Id
		r.SourceId.Extension = m.SourceIdPrimitiveElement.Extension
	}
	r.ValidateProfileId = m.ValidateProfileId
	if m.ValidateProfileIdPrimitiveElement != nil {
		if r.ValidateProfileId == nil {
			r.ValidateProfileId = &Id{}
		}
		r.ValidateProfileId.Id = m.ValidateProfileIdPrimitiveElement.Id
		r.ValidateProfileId.Extension = m.ValidateProfileIdPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		if r.Value == nil {
			r.Value = &String{}
		}
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.WarningOnly = m.WarningOnly
	if m.WarningOnlyPrimitiveElement != nil {
		r.WarningOnly.Id = m.WarningOnlyPrimitiveElement.Id
		r.WarningOnly.Extension = m.WarningOnlyPrimitiveElement.Extension
	}
	return nil
}
func (r TestScriptSetupActionAssert) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Label, xml.StartElement{Name: xml.Name{Local: "label"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Direction, xml.StartElement{Name: xml.Name{Local: "direction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CompareToSourceId, xml.StartElement{Name: xml.Name{Local: "compareToSourceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CompareToSourceExpression, xml.StartElement{Name: xml.Name{Local: "compareToSourceExpression"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CompareToSourcePath, xml.StartElement{Name: xml.Name{Local: "compareToSourcePath"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContentType, xml.StartElement{Name: xml.Name{Local: "contentType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Expression, xml.StartElement{Name: xml.Name{Local: "expression"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HeaderField, xml.StartElement{Name: xml.Name{Local: "headerField"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MinimumId, xml.StartElement{Name: xml.Name{Local: "minimumId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NavigationLinks, xml.StartElement{Name: xml.Name{Local: "navigationLinks"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Operator, xml.StartElement{Name: xml.Name{Local: "operator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Path, xml.StartElement{Name: xml.Name{Local: "path"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestMethod, xml.StartElement{Name: xml.Name{Local: "requestMethod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestUrl, xml.StartElement{Name: xml.Name{Local: "requestURL"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Response, xml.StartElement{Name: xml.Name{Local: "response"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ResponseCode, xml.StartElement{Name: xml.Name{Local: "responseCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceId, xml.StartElement{Name: xml.Name{Local: "sourceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidateProfileId, xml.StartElement{Name: xml.Name{Local: "validateProfileId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WarningOnly, xml.StartElement{Name: xml.Name{Local: "warningOnly"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptSetupActionAssert) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "label":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Label = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "direction":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Direction = &v
			case "compareToSourceId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CompareToSourceId = &v
			case "compareToSourceExpression":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CompareToSourceExpression = &v
			case "compareToSourcePath":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CompareToSourcePath = &v
			case "contentType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContentType = &v
			case "expression":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Expression = &v
			case "headerField":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HeaderField = &v
			case "minimumId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MinimumId = &v
			case "navigationLinks":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NavigationLinks = &v
			case "operator":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operator = &v
			case "path":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Path = &v
			case "requestMethod":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestMethod = &v
			case "requestURL":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestUrl = &v
			case "resource":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = &v
			case "response":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Response = &v
			case "responseCode":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResponseCode = &v
			case "sourceId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceId = &v
			case "validateProfileId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidateProfileId = &v
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			case "warningOnly":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WarningOnly = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptSetupActionAssert) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A test in this script.
type TestScriptTest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of this test used for tracking/logging purposes by test engines.
	Name *String
	// A short description of the test used by test engines for tracking and reporting purposes.
	Description *String
	// Action would contain either an operation or an assertion.
	Action []TestScriptTestAction
}
type jsonTestScriptTest struct {
	Id                          *string                `json:"id,omitempty"`
	Extension                   []Extension            `json:"extension,omitempty"`
	ModifierExtension           []Extension            `json:"modifierExtension,omitempty"`
	Name                        *String                `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement      `json:"_name,omitempty"`
	Description                 *String                `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement      `json:"_description,omitempty"`
	Action                      []TestScriptTestAction `json:"action,omitempty"`
}

func (r TestScriptTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptTest) marshalJSON() jsonTestScriptTest {
	m := jsonTestScriptTest{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
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
	m.Action = r.Action
	return m
}
func (r *TestScriptTest) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptTest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptTest) unmarshalJSON(m jsonTestScriptTest) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
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
	r.Action = m.Action
	return nil
}
func (r TestScriptTest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptTest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Name = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "action":
				var v TestScriptTestAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptTest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Action would contain either an operation or an assertion.
type TestScriptTestAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An operation would involve a REST request to a server.
	Operation *TestScriptSetupActionOperation
	// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
	Assert *TestScriptSetupActionAssert
}
type jsonTestScriptTestAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

func (r TestScriptTestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptTestAction) marshalJSON() jsonTestScriptTestAction {
	m := jsonTestScriptTestAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	m.Assert = r.Assert
	return m
}
func (r *TestScriptTestAction) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptTestAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptTestAction) unmarshalJSON(m jsonTestScriptTestAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	r.Assert = m.Assert
	return nil
}
func (r TestScriptTestAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Operation, xml.StartElement{Name: xml.Name{Local: "operation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Assert, xml.StartElement{Name: xml.Name{Local: "assert"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptTestAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "operation":
				var v TestScriptSetupActionOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = &v
			case "assert":
				var v TestScriptSetupActionAssert
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Assert = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptTestAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A series of operations required to clean up after all the tests are executed (successfully or otherwise).
type TestScriptTeardown struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The teardown action will only contain an operation.
	Action []TestScriptTeardownAction
}
type jsonTestScriptTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `json:"action,omitempty"`
}

func (r TestScriptTeardown) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptTeardown) marshalJSON() jsonTestScriptTeardown {
	m := jsonTestScriptTeardown{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	return m
}
func (r *TestScriptTeardown) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptTeardown
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptTeardown) unmarshalJSON(m jsonTestScriptTeardown) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	return nil
}
func (r TestScriptTeardown) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptTeardown) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "action":
				var v TestScriptTeardownAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptTeardown) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The teardown action will only contain an operation.
type TestScriptTeardownAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An operation would involve a REST request to a server.
	Operation TestScriptSetupActionOperation
}
type jsonTestScriptTeardownAction struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Operation         TestScriptSetupActionOperation `json:"operation,omitempty"`
}

func (r TestScriptTeardownAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestScriptTeardownAction) marshalJSON() jsonTestScriptTeardownAction {
	m := jsonTestScriptTeardownAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	return m
}
func (r *TestScriptTeardownAction) UnmarshalJSON(b []byte) error {
	var m jsonTestScriptTeardownAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestScriptTeardownAction) unmarshalJSON(m jsonTestScriptTeardownAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	return nil
}
func (r TestScriptTeardownAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Operation, xml.StartElement{Name: xml.Name{Local: "operation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestScriptTeardownAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "operation":
				var v TestScriptSetupActionOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestScriptTeardownAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
