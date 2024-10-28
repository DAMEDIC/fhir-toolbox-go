package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A summary of information based on the results of executing a TestScript.
type TestReport struct {
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
	// Identifier for the TestScript assigned for external purposes outside the context of FHIR.
	Identifier *Identifier
	// A free text natural language name identifying the executed TestScript.
	Name *String
	// The current state of this test report.
	Status Code
	// Ideally this is an absolute URL that is used to identify the version-specific TestScript that was executed, matching the `TestScript.url`.
	TestScript Reference
	// The overall result from the execution of the TestScript.
	Result Code
	// The final score (percentage of tests passed) resulting from the execution of the TestScript.
	Score *Decimal
	// Name of the tester producing this report (Organization or individual).
	Tester *String
	// When the TestScript was executed and this TestReport was generated.
	Issued *DateTime
	// A participant in the test execution, either the execution engine, a client, or a server.
	Participant []TestReportParticipant
	// The results of the series of required setup operations before the tests were executed.
	Setup *TestReportSetup
	// A test executed from the test script.
	Test []TestReportTest
	// The results of the series of operations required to clean up after all the tests were executed (successfully or otherwise).
	Teardown *TestReportTeardown
}

func (r TestReport) ResourceType() string {
	return "TestReport"
}
func (r TestReport) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonTestReport struct {
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
	Identifier                    *Identifier             `json:"identifier,omitempty"`
	Name                          *String                 `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement       `json:"_name,omitempty"`
	Status                        Code                    `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement       `json:"_status,omitempty"`
	TestScript                    Reference               `json:"testScript,omitempty"`
	Result                        Code                    `json:"result,omitempty"`
	ResultPrimitiveElement        *primitiveElement       `json:"_result,omitempty"`
	Score                         *Decimal                `json:"score,omitempty"`
	ScorePrimitiveElement         *primitiveElement       `json:"_score,omitempty"`
	Tester                        *String                 `json:"tester,omitempty"`
	TesterPrimitiveElement        *primitiveElement       `json:"_tester,omitempty"`
	Issued                        *DateTime               `json:"issued,omitempty"`
	IssuedPrimitiveElement        *primitiveElement       `json:"_issued,omitempty"`
	Participant                   []TestReportParticipant `json:"participant,omitempty"`
	Setup                         *TestReportSetup        `json:"setup,omitempty"`
	Test                          []TestReportTest        `json:"test,omitempty"`
	Teardown                      *TestReportTeardown     `json:"teardown,omitempty"`
}

func (r TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReport) marshalJSON() jsonTestReport {
	m := jsonTestReport{}
	m.ResourceType = "TestReport"
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
	m.Identifier = r.Identifier
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.TestScript = r.TestScript
	if r.Result.Value != nil {
		m.Result = r.Result
	}
	if r.Result.Id != nil || r.Result.Extension != nil {
		m.ResultPrimitiveElement = &primitiveElement{Id: r.Result.Id, Extension: r.Result.Extension}
	}
	if r.Score != nil && r.Score.Value != nil {
		m.Score = r.Score
	}
	if r.Score != nil && (r.Score.Id != nil || r.Score.Extension != nil) {
		m.ScorePrimitiveElement = &primitiveElement{Id: r.Score.Id, Extension: r.Score.Extension}
	}
	if r.Tester != nil && r.Tester.Value != nil {
		m.Tester = r.Tester
	}
	if r.Tester != nil && (r.Tester.Id != nil || r.Tester.Extension != nil) {
		m.TesterPrimitiveElement = &primitiveElement{Id: r.Tester.Id, Extension: r.Tester.Extension}
	}
	if r.Issued != nil && r.Issued.Value != nil {
		m.Issued = r.Issued
	}
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Participant = r.Participant
	m.Setup = r.Setup
	m.Test = r.Test
	m.Teardown = r.Teardown
	return m
}
func (r *TestReport) UnmarshalJSON(b []byte) error {
	var m jsonTestReport
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReport) unmarshalJSON(m jsonTestReport) error {
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
	r.Identifier = m.Identifier
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.TestScript = m.TestScript
	r.Result = m.Result
	if m.ResultPrimitiveElement != nil {
		r.Result.Id = m.ResultPrimitiveElement.Id
		r.Result.Extension = m.ResultPrimitiveElement.Extension
	}
	r.Score = m.Score
	if m.ScorePrimitiveElement != nil {
		if r.Score == nil {
			r.Score = &Decimal{}
		}
		r.Score.Id = m.ScorePrimitiveElement.Id
		r.Score.Extension = m.ScorePrimitiveElement.Extension
	}
	r.Tester = m.Tester
	if m.TesterPrimitiveElement != nil {
		if r.Tester == nil {
			r.Tester = &String{}
		}
		r.Tester.Id = m.TesterPrimitiveElement.Id
		r.Tester.Extension = m.TesterPrimitiveElement.Extension
	}
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		if r.Issued == nil {
			r.Issued = &DateTime{}
		}
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Participant = m.Participant
	r.Setup = m.Setup
	r.Test = m.Test
	r.Teardown = m.Teardown
	return nil
}
func (r TestReport) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "TestReport"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TestScript, xml.StartElement{Name: xml.Name{Local: "testScript"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Result, xml.StartElement{Name: xml.Name{Local: "result"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Score, xml.StartElement{Name: xml.Name{Local: "score"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tester, xml.StartElement{Name: xml.Name{Local: "tester"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Issued, xml.StartElement{Name: xml.Name{Local: "issued"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Participant, xml.StartElement{Name: xml.Name{Local: "participant"}})
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
func (r *TestReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "testScript":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TestScript = v
			case "result":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Result = v
			case "score":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Score = &v
			case "tester":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tester = &v
			case "issued":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Issued = &v
			case "participant":
				var v TestReportParticipant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Participant = append(r.Participant, v)
			case "setup":
				var v TestReportSetup
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Setup = &v
			case "test":
				var v TestReportTest
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Test = append(r.Test, v)
			case "teardown":
				var v TestReportTeardown
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
func (r TestReport) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A participant in the test execution, either the execution engine, a client, or a server.
type TestReportParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of participant.
	Type Code
	// The uri of the participant. An absolute URL is preferred.
	Uri Uri
	// The display name of the participant.
	Display *String
}
type jsonTestReportParticipant struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Type                    Code              `json:"type,omitempty"`
	TypePrimitiveElement    *primitiveElement `json:"_type,omitempty"`
	Uri                     Uri               `json:"uri,omitempty"`
	UriPrimitiveElement     *primitiveElement `json:"_uri,omitempty"`
	Display                 *String           `json:"display,omitempty"`
	DisplayPrimitiveElement *primitiveElement `json:"_display,omitempty"`
}

func (r TestReportParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportParticipant) marshalJSON() jsonTestReportParticipant {
	m := jsonTestReportParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Uri.Value != nil {
		m.Uri = r.Uri
	}
	if r.Uri.Id != nil || r.Uri.Extension != nil {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	if r.Display != nil && r.Display.Value != nil {
		m.Display = r.Display
	}
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	return m
}
func (r *TestReportParticipant) UnmarshalJSON(b []byte) error {
	var m jsonTestReportParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportParticipant) unmarshalJSON(m jsonTestReportParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		if r.Display == nil {
			r.Display = &String{}
		}
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	return nil
}
func (r TestReportParticipant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Uri, xml.StartElement{Name: xml.Name{Local: "uri"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Display, xml.StartElement{Name: xml.Name{Local: "display"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestReportParticipant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "uri":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Uri = v
			case "display":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Display = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestReportParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The results of the series of required setup operations before the tests were executed.
type TestReportSetup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Action would contain either an operation or an assertion.
	Action []TestReportSetupAction
}
type jsonTestReportSetup struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `json:"action,omitempty"`
}

func (r TestReportSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportSetup) marshalJSON() jsonTestReportSetup {
	m := jsonTestReportSetup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	return m
}
func (r *TestReportSetup) UnmarshalJSON(b []byte) error {
	var m jsonTestReportSetup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportSetup) unmarshalJSON(m jsonTestReportSetup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	return nil
}
func (r TestReportSetup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportSetup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportSetupAction
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
func (r TestReportSetup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Action would contain either an operation or an assertion.
type TestReportSetupAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The operation performed.
	Operation *TestReportSetupActionOperation
	// The results of the assertion performed on the previous operations.
	Assert *TestReportSetupActionAssert
}
type jsonTestReportSetupAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

func (r TestReportSetupAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportSetupAction) marshalJSON() jsonTestReportSetupAction {
	m := jsonTestReportSetupAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	m.Assert = r.Assert
	return m
}
func (r *TestReportSetupAction) UnmarshalJSON(b []byte) error {
	var m jsonTestReportSetupAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportSetupAction) unmarshalJSON(m jsonTestReportSetupAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	r.Assert = m.Assert
	return nil
}
func (r TestReportSetupAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportSetupAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportSetupActionOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = &v
			case "assert":
				var v TestReportSetupActionAssert
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
func (r TestReportSetupAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The operation performed.
type TestReportSetupActionOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The result of this operation.
	Result Code
	// An explanatory message associated with the result.
	Message *Markdown
	// A link to further details on the result.
	Detail *Uri
}
type jsonTestReportSetupActionOperation struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Result                  Code              `json:"result,omitempty"`
	ResultPrimitiveElement  *primitiveElement `json:"_result,omitempty"`
	Message                 *Markdown         `json:"message,omitempty"`
	MessagePrimitiveElement *primitiveElement `json:"_message,omitempty"`
	Detail                  *Uri              `json:"detail,omitempty"`
	DetailPrimitiveElement  *primitiveElement `json:"_detail,omitempty"`
}

func (r TestReportSetupActionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportSetupActionOperation) marshalJSON() jsonTestReportSetupActionOperation {
	m := jsonTestReportSetupActionOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Result.Value != nil {
		m.Result = r.Result
	}
	if r.Result.Id != nil || r.Result.Extension != nil {
		m.ResultPrimitiveElement = &primitiveElement{Id: r.Result.Id, Extension: r.Result.Extension}
	}
	if r.Message != nil && r.Message.Value != nil {
		m.Message = r.Message
	}
	if r.Message != nil && (r.Message.Id != nil || r.Message.Extension != nil) {
		m.MessagePrimitiveElement = &primitiveElement{Id: r.Message.Id, Extension: r.Message.Extension}
	}
	if r.Detail != nil && r.Detail.Value != nil {
		m.Detail = r.Detail
	}
	if r.Detail != nil && (r.Detail.Id != nil || r.Detail.Extension != nil) {
		m.DetailPrimitiveElement = &primitiveElement{Id: r.Detail.Id, Extension: r.Detail.Extension}
	}
	return m
}
func (r *TestReportSetupActionOperation) UnmarshalJSON(b []byte) error {
	var m jsonTestReportSetupActionOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportSetupActionOperation) unmarshalJSON(m jsonTestReportSetupActionOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Result = m.Result
	if m.ResultPrimitiveElement != nil {
		r.Result.Id = m.ResultPrimitiveElement.Id
		r.Result.Extension = m.ResultPrimitiveElement.Extension
	}
	r.Message = m.Message
	if m.MessagePrimitiveElement != nil {
		if r.Message == nil {
			r.Message = &Markdown{}
		}
		r.Message.Id = m.MessagePrimitiveElement.Id
		r.Message.Extension = m.MessagePrimitiveElement.Extension
	}
	r.Detail = m.Detail
	if m.DetailPrimitiveElement != nil {
		if r.Detail == nil {
			r.Detail = &Uri{}
		}
		r.Detail.Id = m.DetailPrimitiveElement.Id
		r.Detail.Extension = m.DetailPrimitiveElement.Extension
	}
	return nil
}
func (r TestReportSetupActionOperation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Result, xml.StartElement{Name: xml.Name{Local: "result"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Message, xml.StartElement{Name: xml.Name{Local: "message"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestReportSetupActionOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "result":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Result = v
			case "message":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Message = &v
			case "detail":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestReportSetupActionOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The results of the assertion performed on the previous operations.
type TestReportSetupActionAssert struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The result of this assertion.
	Result Code
	// An explanatory message associated with the result.
	Message *Markdown
	// A link to further details on the result.
	Detail *String
}
type jsonTestReportSetupActionAssert struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Result                  Code              `json:"result,omitempty"`
	ResultPrimitiveElement  *primitiveElement `json:"_result,omitempty"`
	Message                 *Markdown         `json:"message,omitempty"`
	MessagePrimitiveElement *primitiveElement `json:"_message,omitempty"`
	Detail                  *String           `json:"detail,omitempty"`
	DetailPrimitiveElement  *primitiveElement `json:"_detail,omitempty"`
}

func (r TestReportSetupActionAssert) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportSetupActionAssert) marshalJSON() jsonTestReportSetupActionAssert {
	m := jsonTestReportSetupActionAssert{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Result.Value != nil {
		m.Result = r.Result
	}
	if r.Result.Id != nil || r.Result.Extension != nil {
		m.ResultPrimitiveElement = &primitiveElement{Id: r.Result.Id, Extension: r.Result.Extension}
	}
	if r.Message != nil && r.Message.Value != nil {
		m.Message = r.Message
	}
	if r.Message != nil && (r.Message.Id != nil || r.Message.Extension != nil) {
		m.MessagePrimitiveElement = &primitiveElement{Id: r.Message.Id, Extension: r.Message.Extension}
	}
	if r.Detail != nil && r.Detail.Value != nil {
		m.Detail = r.Detail
	}
	if r.Detail != nil && (r.Detail.Id != nil || r.Detail.Extension != nil) {
		m.DetailPrimitiveElement = &primitiveElement{Id: r.Detail.Id, Extension: r.Detail.Extension}
	}
	return m
}
func (r *TestReportSetupActionAssert) UnmarshalJSON(b []byte) error {
	var m jsonTestReportSetupActionAssert
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportSetupActionAssert) unmarshalJSON(m jsonTestReportSetupActionAssert) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Result = m.Result
	if m.ResultPrimitiveElement != nil {
		r.Result.Id = m.ResultPrimitiveElement.Id
		r.Result.Extension = m.ResultPrimitiveElement.Extension
	}
	r.Message = m.Message
	if m.MessagePrimitiveElement != nil {
		if r.Message == nil {
			r.Message = &Markdown{}
		}
		r.Message.Id = m.MessagePrimitiveElement.Id
		r.Message.Extension = m.MessagePrimitiveElement.Extension
	}
	r.Detail = m.Detail
	if m.DetailPrimitiveElement != nil {
		if r.Detail == nil {
			r.Detail = &String{}
		}
		r.Detail.Id = m.DetailPrimitiveElement.Id
		r.Detail.Extension = m.DetailPrimitiveElement.Extension
	}
	return nil
}
func (r TestReportSetupActionAssert) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Result, xml.StartElement{Name: xml.Name{Local: "result"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Message, xml.StartElement{Name: xml.Name{Local: "message"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TestReportSetupActionAssert) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "result":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Result = v
			case "message":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Message = &v
			case "detail":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TestReportSetupActionAssert) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A test executed from the test script.
type TestReportTest struct {
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
	Action []TestReportTestAction
}
type jsonTestReportTest struct {
	Id                          *string                `json:"id,omitempty"`
	Extension                   []Extension            `json:"extension,omitempty"`
	ModifierExtension           []Extension            `json:"modifierExtension,omitempty"`
	Name                        *String                `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement      `json:"_name,omitempty"`
	Description                 *String                `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement      `json:"_description,omitempty"`
	Action                      []TestReportTestAction `json:"action,omitempty"`
}

func (r TestReportTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportTest) marshalJSON() jsonTestReportTest {
	m := jsonTestReportTest{}
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
func (r *TestReportTest) UnmarshalJSON(b []byte) error {
	var m jsonTestReportTest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportTest) unmarshalJSON(m jsonTestReportTest) error {
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
func (r TestReportTest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportTest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportTestAction
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
func (r TestReportTest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Action would contain either an operation or an assertion.
type TestReportTestAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An operation would involve a REST request to a server.
	Operation *TestReportSetupActionOperation
	// The results of the assertion performed on the previous operations.
	Assert *TestReportSetupActionAssert
}
type jsonTestReportTestAction struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

func (r TestReportTestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportTestAction) marshalJSON() jsonTestReportTestAction {
	m := jsonTestReportTestAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	m.Assert = r.Assert
	return m
}
func (r *TestReportTestAction) UnmarshalJSON(b []byte) error {
	var m jsonTestReportTestAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportTestAction) unmarshalJSON(m jsonTestReportTestAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	r.Assert = m.Assert
	return nil
}
func (r TestReportTestAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportTestAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportSetupActionOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = &v
			case "assert":
				var v TestReportSetupActionAssert
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
func (r TestReportTestAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The results of the series of operations required to clean up after all the tests were executed (successfully or otherwise).
type TestReportTeardown struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The teardown action will only contain an operation.
	Action []TestReportTeardownAction
}
type jsonTestReportTeardown struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `json:"action,omitempty"`
}

func (r TestReportTeardown) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportTeardown) marshalJSON() jsonTestReportTeardown {
	m := jsonTestReportTeardown{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	return m
}
func (r *TestReportTeardown) UnmarshalJSON(b []byte) error {
	var m jsonTestReportTeardown
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportTeardown) unmarshalJSON(m jsonTestReportTeardown) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	return nil
}
func (r TestReportTeardown) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportTeardown) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportTeardownAction
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
func (r TestReportTeardown) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The teardown action will only contain an operation.
type TestReportTeardownAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An operation would involve a REST request to a server.
	Operation TestReportSetupActionOperation
}
type jsonTestReportTeardownAction struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Operation         TestReportSetupActionOperation `json:"operation,omitempty"`
}

func (r TestReportTeardownAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TestReportTeardownAction) marshalJSON() jsonTestReportTeardownAction {
	m := jsonTestReportTeardownAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Operation = r.Operation
	return m
}
func (r *TestReportTeardownAction) UnmarshalJSON(b []byte) error {
	var m jsonTestReportTeardownAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TestReportTeardownAction) unmarshalJSON(m jsonTestReportTeardownAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Operation = m.Operation
	return nil
}
func (r TestReportTeardownAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *TestReportTeardownAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v TestReportSetupActionOperation
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
func (r TestReportTeardownAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
