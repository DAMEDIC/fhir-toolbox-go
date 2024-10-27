package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A physical entity which is the primary unit of operational and/or administrative interest in a study.
type ResearchSubject struct {
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
	// Identifiers assigned to this research subject for a study.
	Identifier []Identifier
	// The current state of the subject.
	Status Code
	// The dates the subject began and ended their participation in the study.
	Period *Period
	// Reference to the study the subject is participating in.
	Study Reference
	// The record of the person or animal who is involved in the study.
	Individual Reference
	// The name of the arm in the study the subject is expected to follow as part of this study.
	AssignedArm *String
	// The name of the arm in the study the subject actually followed as part of this study.
	ActualArm *String
	// A record of the patient's informed agreement to participate in the study.
	Consent *Reference
}

func (r ResearchSubject) ResourceType() string {
	return "ResearchSubject"
}
func (r ResearchSubject) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonResearchSubject struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []ContainedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Period                        *Period             `json:"period,omitempty"`
	Study                         Reference           `json:"study,omitempty"`
	Individual                    Reference           `json:"individual,omitempty"`
	AssignedArm                   *String             `json:"assignedArm,omitempty"`
	AssignedArmPrimitiveElement   *primitiveElement   `json:"_assignedArm,omitempty"`
	ActualArm                     *String             `json:"actualArm,omitempty"`
	ActualArmPrimitiveElement     *primitiveElement   `json:"_actualArm,omitempty"`
	Consent                       *Reference          `json:"consent,omitempty"`
}

func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchSubject) marshalJSON() jsonResearchSubject {
	m := jsonResearchSubject{}
	m.ResourceType = "ResearchSubject"
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
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Period = r.Period
	m.Study = r.Study
	m.Individual = r.Individual
	if r.AssignedArm != nil && r.AssignedArm.Value != nil {
		m.AssignedArm = r.AssignedArm
	}
	if r.AssignedArm != nil && (r.AssignedArm.Id != nil || r.AssignedArm.Extension != nil) {
		m.AssignedArmPrimitiveElement = &primitiveElement{Id: r.AssignedArm.Id, Extension: r.AssignedArm.Extension}
	}
	if r.ActualArm != nil && r.ActualArm.Value != nil {
		m.ActualArm = r.ActualArm
	}
	if r.ActualArm != nil && (r.ActualArm.Id != nil || r.ActualArm.Extension != nil) {
		m.ActualArmPrimitiveElement = &primitiveElement{Id: r.ActualArm.Id, Extension: r.ActualArm.Extension}
	}
	m.Consent = r.Consent
	return m
}
func (r *ResearchSubject) UnmarshalJSON(b []byte) error {
	var m jsonResearchSubject
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchSubject) unmarshalJSON(m jsonResearchSubject) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Study = m.Study
	r.Individual = m.Individual
	r.AssignedArm = m.AssignedArm
	if m.AssignedArmPrimitiveElement != nil {
		if r.AssignedArm == nil {
			r.AssignedArm = &String{}
		}
		r.AssignedArm.Id = m.AssignedArmPrimitiveElement.Id
		r.AssignedArm.Extension = m.AssignedArmPrimitiveElement.Extension
	}
	r.ActualArm = m.ActualArm
	if m.ActualArmPrimitiveElement != nil {
		if r.ActualArm == nil {
			r.ActualArm = &String{}
		}
		r.ActualArm.Id = m.ActualArmPrimitiveElement.Id
		r.ActualArm.Extension = m.ActualArmPrimitiveElement.Extension
	}
	r.Consent = m.Consent
	return nil
}
func (r ResearchSubject) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
