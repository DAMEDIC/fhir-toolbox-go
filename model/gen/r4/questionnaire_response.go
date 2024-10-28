package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
//
// To support structured, hierarchical reporting of data gathered using digital forms and other questionnaires.
type QuestionnaireResponse struct {
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
	// A business identifier assigned to a particular completed (or partially completed) questionnaire.
	Identifier *Identifier
	// The order, proposal or plan that is fulfilled in whole or in part by this QuestionnaireResponse.  For example, a ServiceRequest seeking an intake assessment or a decision support recommendation to assess for post-partum depression.
	BasedOn []Reference
	// A procedure or observation that this questionnaire was performed as part of the execution of.  For example, the surgery a checklist was executed as part of.
	PartOf []Reference
	// The Questionnaire that defines and organizes the questions for which answers are being provided.
	Questionnaire *Canonical
	// The position of the questionnaire response within its overall lifecycle.
	Status Code
	// The subject of the questionnaire response.  This could be a patient, organization, practitioner, device, etc.  This is who/what the answers apply to, but is not necessarily the source of information.
	Subject *Reference
	// The Encounter during which this questionnaire response was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// The date and/or time that this set of answers were last changed.
	Authored *DateTime
	// Person who received the answers to the questions in the QuestionnaireResponse and recorded them in the system.
	Author *Reference
	// The person who answered the questions about the subject.
	Source *Reference
	// A group or question item from the original questionnaire for which answers are provided.
	Item []QuestionnaireResponseItem
}

func (r QuestionnaireResponse) ResourceType() string {
	return "QuestionnaireResponse"
}
func (r QuestionnaireResponse) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonQuestionnaireResponse struct {
	ResourceType                  string                      `json:"resourceType"`
	Id                            *Id                         `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement           `json:"_id,omitempty"`
	Meta                          *Meta                       `json:"meta,omitempty"`
	ImplicitRules                 *Uri                        `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement           `json:"_implicitRules,omitempty"`
	Language                      *Code                       `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement           `json:"_language,omitempty"`
	Text                          *Narrative                  `json:"text,omitempty"`
	Contained                     []ContainedResource         `json:"contained,omitempty"`
	Extension                     []Extension                 `json:"extension,omitempty"`
	ModifierExtension             []Extension                 `json:"modifierExtension,omitempty"`
	Identifier                    *Identifier                 `json:"identifier,omitempty"`
	BasedOn                       []Reference                 `json:"basedOn,omitempty"`
	PartOf                        []Reference                 `json:"partOf,omitempty"`
	Questionnaire                 *Canonical                  `json:"questionnaire,omitempty"`
	QuestionnairePrimitiveElement *primitiveElement           `json:"_questionnaire,omitempty"`
	Status                        Code                        `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement           `json:"_status,omitempty"`
	Subject                       *Reference                  `json:"subject,omitempty"`
	Encounter                     *Reference                  `json:"encounter,omitempty"`
	Authored                      *DateTime                   `json:"authored,omitempty"`
	AuthoredPrimitiveElement      *primitiveElement           `json:"_authored,omitempty"`
	Author                        *Reference                  `json:"author,omitempty"`
	Source                        *Reference                  `json:"source,omitempty"`
	Item                          []QuestionnaireResponseItem `json:"item,omitempty"`
}

func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireResponse) marshalJSON() jsonQuestionnaireResponse {
	m := jsonQuestionnaireResponse{}
	m.ResourceType = "QuestionnaireResponse"
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
	m.BasedOn = r.BasedOn
	m.PartOf = r.PartOf
	if r.Questionnaire != nil && r.Questionnaire.Value != nil {
		m.Questionnaire = r.Questionnaire
	}
	if r.Questionnaire != nil && (r.Questionnaire.Id != nil || r.Questionnaire.Extension != nil) {
		m.QuestionnairePrimitiveElement = &primitiveElement{Id: r.Questionnaire.Id, Extension: r.Questionnaire.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	if r.Authored != nil && r.Authored.Value != nil {
		m.Authored = r.Authored
	}
	if r.Authored != nil && (r.Authored.Id != nil || r.Authored.Extension != nil) {
		m.AuthoredPrimitiveElement = &primitiveElement{Id: r.Authored.Id, Extension: r.Authored.Extension}
	}
	m.Author = r.Author
	m.Source = r.Source
	m.Item = r.Item
	return m
}
func (r *QuestionnaireResponse) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireResponse) unmarshalJSON(m jsonQuestionnaireResponse) error {
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
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.Questionnaire = m.Questionnaire
	if m.QuestionnairePrimitiveElement != nil {
		if r.Questionnaire == nil {
			r.Questionnaire = &Canonical{}
		}
		r.Questionnaire.Id = m.QuestionnairePrimitiveElement.Id
		r.Questionnaire.Extension = m.QuestionnairePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Authored = m.Authored
	if m.AuthoredPrimitiveElement != nil {
		if r.Authored == nil {
			r.Authored = &DateTime{}
		}
		r.Authored.Id = m.AuthoredPrimitiveElement.Id
		r.Authored.Extension = m.AuthoredPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Source = m.Source
	r.Item = m.Item
	return nil
}
func (r QuestionnaireResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "QuestionnaireResponse"
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
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartOf, xml.StartElement{Name: xml.Name{Local: "partOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Questionnaire, xml.StartElement{Name: xml.Name{Local: "questionnaire"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Authored, xml.StartElement{Name: xml.Name{Local: "authored"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *QuestionnaireResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "partOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartOf = append(r.PartOf, v)
			case "questionnaire":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Questionnaire = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "authored":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authored = &v
			case "author":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			case "item":
				var v QuestionnaireResponseItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r QuestionnaireResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A group or question item from the original questionnaire for which answers are provided.
type QuestionnaireResponseItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The item from the Questionnaire that corresponds to this item in the QuestionnaireResponse resource.
	LinkId String
	// A reference to an [ElementDefinition](elementdefinition.html) that provides the details for the item.
	Definition *Uri
	// Text that is displayed above the contents of the group or as the text of the question being answered.
	Text *String
	// The respondent's answer(s) to the question.
	Answer []QuestionnaireResponseItemAnswer
	// Questions or sub-groups nested beneath a question or group.
	Item []QuestionnaireResponseItem
}
type jsonQuestionnaireResponseItem struct {
	Id                         *string                           `json:"id,omitempty"`
	Extension                  []Extension                       `json:"extension,omitempty"`
	ModifierExtension          []Extension                       `json:"modifierExtension,omitempty"`
	LinkId                     String                            `json:"linkId,omitempty"`
	LinkIdPrimitiveElement     *primitiveElement                 `json:"_linkId,omitempty"`
	Definition                 *Uri                              `json:"definition,omitempty"`
	DefinitionPrimitiveElement *primitiveElement                 `json:"_definition,omitempty"`
	Text                       *String                           `json:"text,omitempty"`
	TextPrimitiveElement       *primitiveElement                 `json:"_text,omitempty"`
	Answer                     []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`
	Item                       []QuestionnaireResponseItem       `json:"item,omitempty"`
}

func (r QuestionnaireResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireResponseItem) marshalJSON() jsonQuestionnaireResponseItem {
	m := jsonQuestionnaireResponseItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.LinkId.Value != nil {
		m.LinkId = r.LinkId
	}
	if r.LinkId.Id != nil || r.LinkId.Extension != nil {
		m.LinkIdPrimitiveElement = &primitiveElement{Id: r.LinkId.Id, Extension: r.LinkId.Extension}
	}
	if r.Definition != nil && r.Definition.Value != nil {
		m.Definition = r.Definition
	}
	if r.Definition != nil && (r.Definition.Id != nil || r.Definition.Extension != nil) {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.Answer = r.Answer
	m.Item = r.Item
	return m
}
func (r *QuestionnaireResponseItem) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireResponseItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireResponseItem) unmarshalJSON(m jsonQuestionnaireResponseItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.LinkId = m.LinkId
	if m.LinkIdPrimitiveElement != nil {
		r.LinkId.Id = m.LinkIdPrimitiveElement.Id
		r.LinkId.Extension = m.LinkIdPrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		if r.Definition == nil {
			r.Definition = &Uri{}
		}
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Answer = m.Answer
	r.Item = m.Item
	return nil
}
func (r QuestionnaireResponseItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.LinkId, xml.StartElement{Name: xml.Name{Local: "linkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Definition, xml.StartElement{Name: xml.Name{Local: "definition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Answer, xml.StartElement{Name: xml.Name{Local: "answer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *QuestionnaireResponseItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "linkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LinkId = v
			case "definition":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Definition = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "answer":
				var v QuestionnaireResponseItemAnswer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Answer = append(r.Answer, v)
			case "item":
				var v QuestionnaireResponseItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r QuestionnaireResponseItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The respondent's answer(s) to the question.
type QuestionnaireResponseItemAnswer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The answer (or one of the answers) provided by the respondent to the question.
	Value isQuestionnaireResponseItemAnswerValue
	// Nested groups and/or questions found within this particular answer.
	Item []QuestionnaireResponseItem
}
type isQuestionnaireResponseItemAnswerValue interface {
	isQuestionnaireResponseItemAnswerValue()
}

func (r Boolean) isQuestionnaireResponseItemAnswerValue()    {}
func (r Decimal) isQuestionnaireResponseItemAnswerValue()    {}
func (r Integer) isQuestionnaireResponseItemAnswerValue()    {}
func (r Date) isQuestionnaireResponseItemAnswerValue()       {}
func (r DateTime) isQuestionnaireResponseItemAnswerValue()   {}
func (r Time) isQuestionnaireResponseItemAnswerValue()       {}
func (r String) isQuestionnaireResponseItemAnswerValue()     {}
func (r Uri) isQuestionnaireResponseItemAnswerValue()        {}
func (r Attachment) isQuestionnaireResponseItemAnswerValue() {}
func (r Coding) isQuestionnaireResponseItemAnswerValue()     {}
func (r Quantity) isQuestionnaireResponseItemAnswerValue()   {}
func (r Reference) isQuestionnaireResponseItemAnswerValue()  {}

type jsonQuestionnaireResponseItemAnswer struct {
	Id                            *string                     `json:"id,omitempty"`
	Extension                     []Extension                 `json:"extension,omitempty"`
	ModifierExtension             []Extension                 `json:"modifierExtension,omitempty"`
	ValueBoolean                  *Boolean                    `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement           `json:"_valueBoolean,omitempty"`
	ValueDecimal                  *Decimal                    `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement  *primitiveElement           `json:"_valueDecimal,omitempty"`
	ValueInteger                  *Integer                    `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement           `json:"_valueInteger,omitempty"`
	ValueDate                     *Date                       `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement     *primitiveElement           `json:"_valueDate,omitempty"`
	ValueDateTime                 *DateTime                   `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement           `json:"_valueDateTime,omitempty"`
	ValueTime                     *Time                       `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement     *primitiveElement           `json:"_valueTime,omitempty"`
	ValueString                   *String                     `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement           `json:"_valueString,omitempty"`
	ValueUri                      *Uri                        `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement      *primitiveElement           `json:"_valueUri,omitempty"`
	ValueAttachment               *Attachment                 `json:"valueAttachment,omitempty"`
	ValueCoding                   *Coding                     `json:"valueCoding,omitempty"`
	ValueQuantity                 *Quantity                   `json:"valueQuantity,omitempty"`
	ValueReference                *Reference                  `json:"valueReference,omitempty"`
	Item                          []QuestionnaireResponseItem `json:"item,omitempty"`
}

func (r QuestionnaireResponseItemAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r QuestionnaireResponseItemAnswer) marshalJSON() jsonQuestionnaireResponseItemAnswer {
	m := jsonQuestionnaireResponseItemAnswer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Value.(type) {
	case Boolean:
		if v.Value != nil {
			m.ValueBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.ValueBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		if v.Value != nil {
			m.ValueDecimal = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		if v.Value != nil {
			m.ValueDecimal = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		if v.Value != nil {
			m.ValueInteger = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		if v.Value != nil {
			m.ValueInteger = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		if v.Value != nil {
			m.ValueDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.ValueDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		if v.Value != nil {
			m.ValueDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.ValueDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		if v.Value != nil {
			m.ValueTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		if v.Value != nil {
			m.ValueTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		if v.Value != nil {
			m.ValueString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.ValueString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uri:
		if v.Value != nil {
			m.ValueUri = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		if v.Value != nil {
			m.ValueUri = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	m.Item = r.Item
	return m
}
func (r *QuestionnaireResponseItemAnswer) UnmarshalJSON(b []byte) error {
	var m jsonQuestionnaireResponseItemAnswer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *QuestionnaireResponseItemAnswer) unmarshalJSON(m jsonQuestionnaireResponseItemAnswer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDecimal != nil || m.ValueDecimalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDecimal
		if m.ValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ValueDecimalPrimitiveElement.Id
			v.Extension = m.ValueDecimalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDate != nil || m.ValueDatePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDate
		if m.ValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ValueDatePrimitiveElement.Id
			v.Extension = m.ValueDatePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDateTime != nil || m.ValueDateTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDateTime
		if m.ValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ValueDateTimePrimitiveElement.Id
			v.Extension = m.ValueDateTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueTime != nil || m.ValueTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTime
		if m.ValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.ValueTimePrimitiveElement.Id
			v.Extension = m.ValueTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUri != nil || m.ValueUriPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUri
		if m.ValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ValueUriPrimitiveElement.Id
			v.Extension = m.ValueUriPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	r.Item = m.Item
	return nil
}
func (r QuestionnaireResponseItemAnswer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Value.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Decimal, *Decimal:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDecimal"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger"}})
		if err != nil {
			return err
		}
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDate"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDateTime"}})
		if err != nil {
			return err
		}
	case Time, *Time:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTime"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUri"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case Coding, *Coding:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCoding"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *QuestionnaireResponseItemAnswer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "valueBoolean":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDecimal":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDate":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDateTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueString":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUri":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAttachment":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCoding":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueQuantity":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "item":
				var v QuestionnaireResponseItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r QuestionnaireResponseItemAnswer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
