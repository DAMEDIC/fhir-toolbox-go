package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The clinical particulars - indications, contraindications etc. of a medicinal product, including for regulatory purposes.
type MedicinalProductContraindication struct {
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
	// The medication for which this is an indication.
	Subject []Reference
	// The disease, symptom or procedure for the contraindication.
	Disease *CodeableConcept
	// The status of the disease or symptom for the contraindication.
	DiseaseStatus *CodeableConcept
	// A comorbidity (concurrent condition) or coinfection.
	Comorbidity []CodeableConcept
	// Information about the use of the medicinal product in relation to other therapies as part of the indication.
	TherapeuticIndication []Reference
	// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
	OtherTherapy []MedicinalProductContraindicationOtherTherapy
	// The population group to which this applies.
	Population []Population
}

func (r MedicinalProductContraindication) ResourceType() string {
	return "MedicinalProductContraindication"
}
func (r MedicinalProductContraindication) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProductContraindication struct {
	ResourceType                  string                                         `json:"resourceType"`
	Id                            *Id                                            `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                              `json:"_id,omitempty"`
	Meta                          *Meta                                          `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                              `json:"_implicitRules,omitempty"`
	Language                      *Code                                          `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                              `json:"_language,omitempty"`
	Text                          *Narrative                                     `json:"text,omitempty"`
	Contained                     []ContainedResource                            `json:"contained,omitempty"`
	Extension                     []Extension                                    `json:"extension,omitempty"`
	ModifierExtension             []Extension                                    `json:"modifierExtension,omitempty"`
	Subject                       []Reference                                    `json:"subject,omitempty"`
	Disease                       *CodeableConcept                               `json:"disease,omitempty"`
	DiseaseStatus                 *CodeableConcept                               `json:"diseaseStatus,omitempty"`
	Comorbidity                   []CodeableConcept                              `json:"comorbidity,omitempty"`
	TherapeuticIndication         []Reference                                    `json:"therapeuticIndication,omitempty"`
	OtherTherapy                  []MedicinalProductContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
	Population                    []Population                                   `json:"population,omitempty"`
}

func (r MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductContraindication) marshalJSON() jsonMedicinalProductContraindication {
	m := jsonMedicinalProductContraindication{}
	m.ResourceType = "MedicinalProductContraindication"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Subject = r.Subject
	m.Disease = r.Disease
	m.DiseaseStatus = r.DiseaseStatus
	m.Comorbidity = r.Comorbidity
	m.TherapeuticIndication = r.TherapeuticIndication
	m.OtherTherapy = r.OtherTherapy
	m.Population = r.Population
	return m
}
func (r *MedicinalProductContraindication) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductContraindication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductContraindication) unmarshalJSON(m jsonMedicinalProductContraindication) error {
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
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Subject = m.Subject
	r.Disease = m.Disease
	r.DiseaseStatus = m.DiseaseStatus
	r.Comorbidity = m.Comorbidity
	r.TherapeuticIndication = m.TherapeuticIndication
	r.OtherTherapy = m.OtherTherapy
	r.Population = m.Population
	return nil
}
func (r MedicinalProductContraindication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
type MedicinalProductContraindicationOtherTherapy struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of relationship between the medicinal product indication or contraindication and another therapy.
	TherapyRelationshipType CodeableConcept
	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication.
	Medication isMedicinalProductContraindicationOtherTherapyMedication
}
type isMedicinalProductContraindicationOtherTherapyMedication interface {
	isMedicinalProductContraindicationOtherTherapyMedication()
}

func (r CodeableConcept) isMedicinalProductContraindicationOtherTherapyMedication() {}
func (r Reference) isMedicinalProductContraindicationOtherTherapyMedication()       {}

type jsonMedicinalProductContraindicationOtherTherapy struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept  `json:"therapyRelationshipType,omitempty"`
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference       `json:"medicationReference,omitempty"`
}

func (r MedicinalProductContraindicationOtherTherapy) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductContraindicationOtherTherapy) marshalJSON() jsonMedicinalProductContraindicationOtherTherapy {
	m := jsonMedicinalProductContraindicationOtherTherapy{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.TherapyRelationshipType = r.TherapyRelationshipType
	switch v := r.Medication.(type) {
	case CodeableConcept:
		m.MedicationCodeableConcept = &v
	case *CodeableConcept:
		m.MedicationCodeableConcept = v
	case Reference:
		m.MedicationReference = &v
	case *Reference:
		m.MedicationReference = v
	}
	return m
}
func (r *MedicinalProductContraindicationOtherTherapy) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductContraindicationOtherTherapy
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductContraindicationOtherTherapy) unmarshalJSON(m jsonMedicinalProductContraindicationOtherTherapy) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.TherapyRelationshipType = m.TherapyRelationshipType
	if m.MedicationCodeableConcept != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationCodeableConcept
		r.Medication = v
	}
	if m.MedicationReference != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationReference
		r.Medication = v
	}
	return nil
}
func (r MedicinalProductContraindicationOtherTherapy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
