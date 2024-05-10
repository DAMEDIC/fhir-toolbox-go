package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
//
// To support documents, and also to capture the EN13606 notion of an attested commit to the patient EHR, and to allow a set of disparate resources at the information/engineering level to be gathered into a clinical statement.
type Composition struct {
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
	// A version-independent identifier for the Composition. This identifier stays constant as the composition is changed over time.
	Identifier *Identifier
	// The workflow/clinical status of this composition. The status is a marker for the clinical standing of the document.
	Status Code
	// Specifies the particular kind of composition (e.g. History and Physical, Discharge Summary, Progress Note). This usually equates to the purpose of making the composition.
	Type CodeableConcept
	// A categorization for the type of the composition - helps for indexing and searching. This may be implied by or derived from the code specified in the Composition Type.
	Category []CodeableConcept
	// Who or what the composition is about. The composition can be about a person, (patient or healthcare practitioner), a device (e.g. a machine) or even a group of subjects (such as a document about a herd of livestock, or a set of patients that share a common exposure).
	Subject *Reference
	// Describes the clinical encounter or type of care this documentation is associated with.
	Encounter *Reference
	// The composition editing time, when the composition was last logically changed by the author.
	Date DateTime
	// Identifies who is responsible for the information in the composition, not necessarily who typed it in.
	Author []Reference
	// Official human-readable label for the composition.
	Title String
	// The code specifying the level of confidentiality of the Composition.
	Confidentiality *Code
	// A participant who has attested to the accuracy of the composition/document.
	Attester []CompositionAttester
	// Identifies the organization or group who is responsible for ongoing maintenance of and access to the composition/document information.
	Custodian *Reference
	// Relationships that this composition has with other compositions or documents that already exist.
	RelatesTo []CompositionRelatesTo
	// The clinical service, such as a colonoscopy or an appendectomy, being documented.
	Event []CompositionEvent
	// The root of the sections that make up the composition.
	Section []CompositionSection
}
type jsonComposition struct {
	ResourceType                    string                 `json:"resourceType"`
	Id                              *Id                    `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement      `json:"_id,omitempty"`
	Meta                            *Meta                  `json:"meta,omitempty"`
	ImplicitRules                   *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                        *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement      `json:"_language,omitempty"`
	Text                            *Narrative             `json:"text,omitempty"`
	Contained                       []containedResource    `json:"contained,omitempty"`
	Extension                       []Extension            `json:"extension,omitempty"`
	ModifierExtension               []Extension            `json:"modifierExtension,omitempty"`
	Identifier                      *Identifier            `json:"identifier,omitempty"`
	Status                          Code                   `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement      `json:"_status,omitempty"`
	Type                            CodeableConcept        `json:"type,omitempty"`
	Category                        []CodeableConcept      `json:"category,omitempty"`
	Subject                         *Reference             `json:"subject,omitempty"`
	Encounter                       *Reference             `json:"encounter,omitempty"`
	Date                            DateTime               `json:"date,omitempty"`
	DatePrimitiveElement            *primitiveElement      `json:"_date,omitempty"`
	Author                          []Reference            `json:"author,omitempty"`
	Title                           String                 `json:"title,omitempty"`
	TitlePrimitiveElement           *primitiveElement      `json:"_title,omitempty"`
	Confidentiality                 *Code                  `json:"confidentiality,omitempty"`
	ConfidentialityPrimitiveElement *primitiveElement      `json:"_confidentiality,omitempty"`
	Attester                        []CompositionAttester  `json:"attester,omitempty"`
	Custodian                       *Reference             `json:"custodian,omitempty"`
	RelatesTo                       []CompositionRelatesTo `json:"relatesTo,omitempty"`
	Event                           []CompositionEvent     `json:"event,omitempty"`
	Section                         []CompositionSection   `json:"section,omitempty"`
}

func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Composition) marshalJSON() jsonComposition {
	m := jsonComposition{}
	m.ResourceType = "Composition"
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	m.Category = r.Category
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.Date = r.Date
	if r.Date.Id != nil || r.Date.Extension != nil {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Author = r.Author
	m.Title = r.Title
	if r.Title.Id != nil || r.Title.Extension != nil {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Confidentiality = r.Confidentiality
	if r.Confidentiality != nil && (r.Confidentiality.Id != nil || r.Confidentiality.Extension != nil) {
		m.ConfidentialityPrimitiveElement = &primitiveElement{Id: r.Confidentiality.Id, Extension: r.Confidentiality.Extension}
	}
	m.Attester = r.Attester
	m.Custodian = r.Custodian
	m.RelatesTo = r.RelatesTo
	m.Event = r.Event
	m.Section = r.Section
	return m
}
func (r *Composition) UnmarshalJSON(b []byte) error {
	var m jsonComposition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Composition) unmarshalJSON(m jsonComposition) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Category = m.Category
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Confidentiality = m.Confidentiality
	if m.ConfidentialityPrimitiveElement != nil {
		r.Confidentiality.Id = m.ConfidentialityPrimitiveElement.Id
		r.Confidentiality.Extension = m.ConfidentialityPrimitiveElement.Extension
	}
	r.Attester = m.Attester
	r.Custodian = m.Custodian
	r.RelatesTo = m.RelatesTo
	r.Event = m.Event
	r.Section = m.Section
	return nil
}
func (r Composition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A participant who has attested to the accuracy of the composition/document.
type CompositionAttester struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of attestation the authenticator offers.
	Mode Code
	// When the composition was attested by the party.
	Time *DateTime
	// Who attested the composition in the specified way.
	Party *Reference
}
type jsonCompositionAttester struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Mode                 Code              `json:"mode,omitempty"`
	ModePrimitiveElement *primitiveElement `json:"_mode,omitempty"`
	Time                 *DateTime         `json:"time,omitempty"`
	TimePrimitiveElement *primitiveElement `json:"_time,omitempty"`
	Party                *Reference        `json:"party,omitempty"`
}

func (r CompositionAttester) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompositionAttester) marshalJSON() jsonCompositionAttester {
	m := jsonCompositionAttester{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Mode = r.Mode
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Time = r.Time
	if r.Time != nil && (r.Time.Id != nil || r.Time.Extension != nil) {
		m.TimePrimitiveElement = &primitiveElement{Id: r.Time.Id, Extension: r.Time.Extension}
	}
	m.Party = r.Party
	return m
}
func (r *CompositionAttester) UnmarshalJSON(b []byte) error {
	var m jsonCompositionAttester
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompositionAttester) unmarshalJSON(m jsonCompositionAttester) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Time = m.Time
	if m.TimePrimitiveElement != nil {
		r.Time.Id = m.TimePrimitiveElement.Id
		r.Time.Extension = m.TimePrimitiveElement.Extension
	}
	r.Party = m.Party
	return nil
}
func (r CompositionAttester) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Relationships that this composition has with other compositions or documents that already exist.
type CompositionRelatesTo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of relationship that this composition has with anther composition or document.
	Code Code
	// The target composition/document of this relationship.
	Target isCompositionRelatesToTarget
}
type isCompositionRelatesToTarget interface {
	isCompositionRelatesToTarget()
}

func (r Identifier) isCompositionRelatesToTarget() {}
func (r Reference) isCompositionRelatesToTarget()  {}

type jsonCompositionRelatesTo struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Code                 Code              `json:"code,omitempty"`
	CodePrimitiveElement *primitiveElement `json:"_code,omitempty"`
	TargetIdentifier     *Identifier       `json:"targetIdentifier,omitempty"`
	TargetReference      *Reference        `json:"targetReference,omitempty"`
}

func (r CompositionRelatesTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompositionRelatesTo) marshalJSON() jsonCompositionRelatesTo {
	m := jsonCompositionRelatesTo{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	switch v := r.Target.(type) {
	case Identifier:
		m.TargetIdentifier = &v
	case *Identifier:
		m.TargetIdentifier = v
	case Reference:
		m.TargetReference = &v
	case *Reference:
		m.TargetReference = v
	}
	return m
}
func (r *CompositionRelatesTo) UnmarshalJSON(b []byte) error {
	var m jsonCompositionRelatesTo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompositionRelatesTo) unmarshalJSON(m jsonCompositionRelatesTo) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	if m.TargetIdentifier != nil {
		if r.Target != nil {
			return fmt.Errorf("multiple values for field \"Target\"")
		}
		v := m.TargetIdentifier
		r.Target = v
	}
	if m.TargetReference != nil {
		if r.Target != nil {
			return fmt.Errorf("multiple values for field \"Target\"")
		}
		v := m.TargetReference
		r.Target = v
	}
	return nil
}
func (r CompositionRelatesTo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The clinical service, such as a colonoscopy or an appendectomy, being documented.
type CompositionEvent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This list of codes represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented. In some cases, the event is inherent in the typeCode, such as a "History and Physical Report" in which the procedure being documented is necessarily a "History and Physical" act.
	Code []CodeableConcept
	// The period of time covered by the documentation. There is no assertion that the documentation is a complete representation for this period, only that it documents events during this time.
	Period *Period
	// The description and/or reference of the event(s) being documented. For example, this could be used to document such a colonoscopy or an appendectomy.
	Detail []Reference
}
type jsonCompositionEvent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

func (r CompositionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompositionEvent) marshalJSON() jsonCompositionEvent {
	m := jsonCompositionEvent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Period = r.Period
	m.Detail = r.Detail
	return m
}
func (r *CompositionEvent) UnmarshalJSON(b []byte) error {
	var m jsonCompositionEvent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompositionEvent) unmarshalJSON(m jsonCompositionEvent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Period = m.Period
	r.Detail = m.Detail
	return nil
}
func (r CompositionEvent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The root of the sections that make up the composition.
type CompositionSection struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The label for this particular section.  This will be part of the rendered content for the document, and is often used to build a table of contents.
	Title *String
	// A code identifying the kind of content contained within the section. This must be consistent with the section title.
	Code *CodeableConcept
	// Identifies who is responsible for the information in this section, not necessarily who typed it in.
	Author []Reference
	// The actual focus of the section when it is not the subject of the composition, but instead represents something or someone associated with the subject such as (for a patient subject) a spouse, parent, fetus, or donor. If not focus is specified, the focus is assumed to be focus of the parent section, or, for a section in the Composition itself, the subject of the composition. Sections with a focus SHALL only include resources where the logical subject (patient, subject, focus, etc.) matches the section focus, or the resources have no logical subject (few resources).
	Focus *Reference
	// A human-readable narrative that contains the attested content of the section, used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative.
	Text *Narrative
	// How the entry list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	Mode *Code
	// Specifies the order applied to the items in the section entries.
	OrderedBy *CodeableConcept
	// A reference to the actual resource from which the narrative in the section is derived.
	Entry []Reference
	// If the section is empty, why the list is empty. An empty section typically has some text explaining the empty reason.
	EmptyReason *CodeableConcept
	// A nested sub-section within this section.
	Section []CompositionSection
}
type jsonCompositionSection struct {
	Id                    *string              `json:"id,omitempty"`
	Extension             []Extension          `json:"extension,omitempty"`
	ModifierExtension     []Extension          `json:"modifierExtension,omitempty"`
	Title                 *String              `json:"title,omitempty"`
	TitlePrimitiveElement *primitiveElement    `json:"_title,omitempty"`
	Code                  *CodeableConcept     `json:"code,omitempty"`
	Author                []Reference          `json:"author,omitempty"`
	Focus                 *Reference           `json:"focus,omitempty"`
	Text                  *Narrative           `json:"text,omitempty"`
	Mode                  *Code                `json:"mode,omitempty"`
	ModePrimitiveElement  *primitiveElement    `json:"_mode,omitempty"`
	OrderedBy             *CodeableConcept     `json:"orderedBy,omitempty"`
	Entry                 []Reference          `json:"entry,omitempty"`
	EmptyReason           *CodeableConcept     `json:"emptyReason,omitempty"`
	Section               []CompositionSection `json:"section,omitempty"`
}

func (r CompositionSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompositionSection) marshalJSON() jsonCompositionSection {
	m := jsonCompositionSection{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Code = r.Code
	m.Author = r.Author
	m.Focus = r.Focus
	m.Text = r.Text
	m.Mode = r.Mode
	if r.Mode != nil && (r.Mode.Id != nil || r.Mode.Extension != nil) {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.OrderedBy = r.OrderedBy
	m.Entry = r.Entry
	m.EmptyReason = r.EmptyReason
	m.Section = r.Section
	return m
}
func (r *CompositionSection) UnmarshalJSON(b []byte) error {
	var m jsonCompositionSection
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompositionSection) unmarshalJSON(m jsonCompositionSection) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Author = m.Author
	r.Focus = m.Focus
	r.Text = m.Text
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.OrderedBy = m.OrderedBy
	r.Entry = m.Entry
	r.EmptyReason = m.EmptyReason
	r.Section = m.Section
	return nil
}
func (r CompositionSection) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
