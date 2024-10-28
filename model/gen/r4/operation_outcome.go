package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A collection of error, warning, or information messages that result from a system action.
type OperationOutcome struct {
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
	// An error, warning, or information message that results from a system action.
	Issue []OperationOutcomeIssue
}

func (r OperationOutcome) ResourceType() string {
	return "OperationOutcome"
}
func (r OperationOutcome) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonOperationOutcome struct {
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
	Issue                         []OperationOutcomeIssue `json:"issue,omitempty"`
}

func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationOutcome) marshalJSON() jsonOperationOutcome {
	m := jsonOperationOutcome{}
	m.ResourceType = "OperationOutcome"
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
	m.Issue = r.Issue
	return m
}
func (r *OperationOutcome) UnmarshalJSON(b []byte) error {
	var m jsonOperationOutcome
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationOutcome) unmarshalJSON(m jsonOperationOutcome) error {
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
	r.Issue = m.Issue
	return nil
}
func (r OperationOutcome) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "OperationOutcome"
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
	err = e.EncodeElement(r.Issue, xml.StartElement{Name: xml.Name{Local: "issue"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationOutcome) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "issue":
				var v OperationOutcomeIssue
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Issue = append(r.Issue, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationOutcome) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An error, warning, or information message that results from a system action.
type OperationOutcomeIssue struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates whether the issue indicates a variation from successful processing.
	Severity Code
	// Describes the type of the issue. The system that creates an OperationOutcome SHALL choose the most applicable code from the IssueType value set, and may additional provide its own code for the error in the details element.
	Code Code
	// Additional details about the error. This may be a text description of the error or a system code that identifies the error.
	Details *CodeableConcept
	// Additional diagnostic information about the issue.
	Diagnostics *String
	// This element is deprecated because it is XML specific. It is replaced by issue.expression, which is format independent, and simpler to parse.
	//
	// For resource issues, this will be a simple XPath limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.  For HTTP errors, will be "http." + the parameter name.
	Location []String
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	Expression []String
}
type jsonOperationOutcomeIssue struct {
	Id                          *string             `json:"id,omitempty"`
	Extension                   []Extension         `json:"extension,omitempty"`
	ModifierExtension           []Extension         `json:"modifierExtension,omitempty"`
	Severity                    Code                `json:"severity,omitempty"`
	SeverityPrimitiveElement    *primitiveElement   `json:"_severity,omitempty"`
	Code                        Code                `json:"code,omitempty"`
	CodePrimitiveElement        *primitiveElement   `json:"_code,omitempty"`
	Details                     *CodeableConcept    `json:"details,omitempty"`
	Diagnostics                 *String             `json:"diagnostics,omitempty"`
	DiagnosticsPrimitiveElement *primitiveElement   `json:"_diagnostics,omitempty"`
	Location                    []String            `json:"location,omitempty"`
	LocationPrimitiveElement    []*primitiveElement `json:"_location,omitempty"`
	Expression                  []String            `json:"expression,omitempty"`
	ExpressionPrimitiveElement  []*primitiveElement `json:"_expression,omitempty"`
}

func (r OperationOutcomeIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationOutcomeIssue) marshalJSON() jsonOperationOutcomeIssue {
	m := jsonOperationOutcomeIssue{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Severity.Value != nil {
		m.Severity = r.Severity
	}
	if r.Severity.Id != nil || r.Severity.Extension != nil {
		m.SeverityPrimitiveElement = &primitiveElement{Id: r.Severity.Id, Extension: r.Severity.Extension}
	}
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Details = r.Details
	if r.Diagnostics != nil && r.Diagnostics.Value != nil {
		m.Diagnostics = r.Diagnostics
	}
	if r.Diagnostics != nil && (r.Diagnostics.Id != nil || r.Diagnostics.Extension != nil) {
		m.DiagnosticsPrimitiveElement = &primitiveElement{Id: r.Diagnostics.Id, Extension: r.Diagnostics.Extension}
	}
	anyLocationValue := false
	for _, e := range r.Location {
		if e.Value != nil {
			anyLocationValue = true
			break
		}
	}
	if anyLocationValue {
		m.Location = r.Location
	}
	anyLocationIdOrExtension := false
	for _, e := range r.Location {
		if e.Id != nil || e.Extension != nil {
			anyLocationIdOrExtension = true
			break
		}
	}
	if anyLocationIdOrExtension {
		m.LocationPrimitiveElement = make([]*primitiveElement, 0, len(r.Location))
		for _, e := range r.Location {
			if e.Id != nil || e.Extension != nil {
				m.LocationPrimitiveElement = append(m.LocationPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LocationPrimitiveElement = append(m.LocationPrimitiveElement, nil)
			}
		}
	}
	anyExpressionValue := false
	for _, e := range r.Expression {
		if e.Value != nil {
			anyExpressionValue = true
			break
		}
	}
	if anyExpressionValue {
		m.Expression = r.Expression
	}
	anyExpressionIdOrExtension := false
	for _, e := range r.Expression {
		if e.Id != nil || e.Extension != nil {
			anyExpressionIdOrExtension = true
			break
		}
	}
	if anyExpressionIdOrExtension {
		m.ExpressionPrimitiveElement = make([]*primitiveElement, 0, len(r.Expression))
		for _, e := range r.Expression {
			if e.Id != nil || e.Extension != nil {
				m.ExpressionPrimitiveElement = append(m.ExpressionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ExpressionPrimitiveElement = append(m.ExpressionPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *OperationOutcomeIssue) UnmarshalJSON(b []byte) error {
	var m jsonOperationOutcomeIssue
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationOutcomeIssue) unmarshalJSON(m jsonOperationOutcomeIssue) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Severity = m.Severity
	if m.SeverityPrimitiveElement != nil {
		r.Severity.Id = m.SeverityPrimitiveElement.Id
		r.Severity.Extension = m.SeverityPrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Details = m.Details
	r.Diagnostics = m.Diagnostics
	if m.DiagnosticsPrimitiveElement != nil {
		if r.Diagnostics == nil {
			r.Diagnostics = &String{}
		}
		r.Diagnostics.Id = m.DiagnosticsPrimitiveElement.Id
		r.Diagnostics.Extension = m.DiagnosticsPrimitiveElement.Extension
	}
	r.Location = m.Location
	for i, e := range m.LocationPrimitiveElement {
		if len(r.Location) <= i {
			r.Location = append(r.Location, String{})
		}
		if e != nil {
			r.Location[i].Id = e.Id
			r.Location[i].Extension = e.Extension
		}
	}
	r.Expression = m.Expression
	for i, e := range m.ExpressionPrimitiveElement {
		if len(r.Expression) <= i {
			r.Expression = append(r.Expression, String{})
		}
		if e != nil {
			r.Expression[i].Id = e.Id
			r.Expression[i].Extension = e.Extension
		}
	}
	return nil
}
func (r OperationOutcomeIssue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Severity, xml.StartElement{Name: xml.Name{Local: "severity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Details, xml.StartElement{Name: xml.Name{Local: "details"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Diagnostics, xml.StartElement{Name: xml.Name{Local: "diagnostics"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Expression, xml.StartElement{Name: xml.Name{Local: "expression"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationOutcomeIssue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "severity":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Severity = v
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "details":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Details = &v
			case "diagnostics":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Diagnostics = &v
			case "location":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			case "expression":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Expression = append(r.Expression, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationOutcomeIssue) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
