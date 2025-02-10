package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
)

// A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendation struct {
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
	// A unique identifier assigned to this particular recommendation record.
	Identifier []Identifier
	// The patient the recommendation(s) are for.
	Patient Reference
	// The date the immunization recommendation(s) were created.
	Date DateTime
	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *Reference
	// Vaccine administration recommendations.
	Recommendation []ImmunizationRecommendationRecommendation
}

// Vaccine administration recommendations.
type ImmunizationRecommendationRecommendation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Vaccine(s) or vaccine group that pertain to the recommendation.
	VaccineCode []CodeableConcept
	// The targeted disease for the recommendation.
	TargetDisease *CodeableConcept
	// Vaccine(s) which should not be used to fulfill the recommendation.
	ContraindicatedVaccineCode []CodeableConcept
	// Indicates the patient status with respect to the path to immunity for the target disease.
	ForecastStatus CodeableConcept
	// The reason for the assigned forecast status.
	ForecastReason []CodeableConcept
	// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion
	// Contains the description about the protocol under which the vaccine was administered.
	Description *String
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *String
	// Nominal position of the recommended dose in a series (e.g. dose 2 is the next recommended dose).
	DoseNumber isImmunizationRecommendationRecommendationDoseNumber
	// The recommended number of doses to achieve immunity.
	SeriesDoses isImmunizationRecommendationRecommendationSeriesDoses
	// Immunization event history and/or evaluation that supports the status and recommendation.
	SupportingImmunization []Reference
	// Patient Information that supports the status and recommendation.  This includes patient observations, adverse reactions and allergy/intolerance information.
	SupportingPatientInformation []Reference
}
type isImmunizationRecommendationRecommendationDoseNumber interface {
	model.Element
	isImmunizationRecommendationRecommendationDoseNumber()
}

func (r PositiveInt) isImmunizationRecommendationRecommendationDoseNumber() {}
func (r String) isImmunizationRecommendationRecommendationDoseNumber()      {}

type isImmunizationRecommendationRecommendationSeriesDoses interface {
	model.Element
	isImmunizationRecommendationRecommendationSeriesDoses()
}

func (r PositiveInt) isImmunizationRecommendationRecommendationSeriesDoses() {}
func (r String) isImmunizationRecommendationRecommendationSeriesDoses()      {}

// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
type ImmunizationRecommendationRecommendationDateCriterion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date classification of recommendation.  For example, earliest date to give, latest date to give, etc.
	Code CodeableConcept
	// The date whose meaning is specified by dateCriterion.code.
	Value DateTime
}

func (r ImmunizationRecommendation) ResourceType() string {
	return "ImmunizationRecommendation"
}
func (r ImmunizationRecommendation) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r ImmunizationRecommendation) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	s += r.Patient.MemSize() - int(unsafe.Sizeof(r.Patient))
	s += r.Date.MemSize() - int(unsafe.Sizeof(r.Date))
	if r.Authority != nil {
		s += r.Authority.MemSize()
	}
	for _, i := range r.Recommendation {
		s += i.MemSize()
	}
	s += (cap(r.Recommendation) - len(r.Recommendation)) * int(unsafe.Sizeof(ImmunizationRecommendationRecommendation{}))
	return s
}
func (r ImmunizationRecommendationRecommendation) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.VaccineCode {
		s += i.MemSize()
	}
	s += (cap(r.VaccineCode) - len(r.VaccineCode)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.TargetDisease != nil {
		s += r.TargetDisease.MemSize()
	}
	for _, i := range r.ContraindicatedVaccineCode {
		s += i.MemSize()
	}
	s += (cap(r.ContraindicatedVaccineCode) - len(r.ContraindicatedVaccineCode)) * int(unsafe.Sizeof(CodeableConcept{}))
	s += r.ForecastStatus.MemSize() - int(unsafe.Sizeof(r.ForecastStatus))
	for _, i := range r.ForecastReason {
		s += i.MemSize()
	}
	s += (cap(r.ForecastReason) - len(r.ForecastReason)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.DateCriterion {
		s += i.MemSize()
	}
	s += (cap(r.DateCriterion) - len(r.DateCriterion)) * int(unsafe.Sizeof(ImmunizationRecommendationRecommendationDateCriterion{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Series != nil {
		s += r.Series.MemSize()
	}
	if r.DoseNumber != nil {
		s += r.DoseNumber.MemSize()
	}
	if r.SeriesDoses != nil {
		s += r.SeriesDoses.MemSize()
	}
	for _, i := range r.SupportingImmunization {
		s += i.MemSize()
	}
	s += (cap(r.SupportingImmunization) - len(r.SupportingImmunization)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.SupportingPatientInformation {
		s += i.MemSize()
	}
	s += (cap(r.SupportingPatientInformation) - len(r.SupportingPatientInformation)) * int(unsafe.Sizeof(Reference{}))
	return s
}
func (r ImmunizationRecommendationRecommendationDateCriterion) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Code.MemSize() - int(unsafe.Sizeof(r.Code))
	s += r.Value.MemSize() - int(unsafe.Sizeof(r.Value))
	return s
}
func (r ImmunizationRecommendation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationRecommendationRecommendation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationRecommendationRecommendationDateCriterion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationRecommendation) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"ImmunizationRecommendation\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"patient\":"))
	if err != nil {
		return err
	}
	err = r.Patient.marshalJSON(w)
	if err != nil {
		return err
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"date\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Date)
		if err != nil {
			return err
		}
	}
	if r.Date.Id != nil || r.Date.Extension != nil {
		p := primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_date\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Authority != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authority\":"))
		if err != nil {
			return err
		}
		err = r.Authority.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Recommendation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"recommendation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Recommendation {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationRecommendationRecommendation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationRecommendationRecommendation) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.VaccineCode) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"vaccineCode\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.VaccineCode {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.TargetDisease != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetDisease\":"))
		if err != nil {
			return err
		}
		err = r.TargetDisease.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.ContraindicatedVaccineCode) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contraindicatedVaccineCode\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ContraindicatedVaccineCode {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"forecastStatus\":"))
	if err != nil {
		return err
	}
	err = r.ForecastStatus.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.ForecastReason) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"forecastReason\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ForecastReason {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.DateCriterion) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dateCriterion\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.DateCriterion {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Series != nil && r.Series.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"series\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Series)
		if err != nil {
			return err
		}
	}
	if r.Series != nil && (r.Series.Id != nil || r.Series.Extension != nil) {
		p := primitiveElement{Id: r.Series.Id, Extension: r.Series.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_series\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.DoseNumber.(type) {
	case PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"doseNumberPositiveInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_doseNumberPositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"doseNumberPositiveInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_doseNumberPositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"doseNumberString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_doseNumberString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"doseNumberString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_doseNumberString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	}
	switch v := r.SeriesDoses.(type) {
	case PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"seriesDosesPositiveInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_seriesDosesPositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"seriesDosesPositiveInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_seriesDosesPositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"seriesDosesString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_seriesDosesString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"seriesDosesString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_seriesDosesString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	}
	if len(r.SupportingImmunization) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supportingImmunization\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SupportingImmunization {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.SupportingPatientInformation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supportingPatientInformation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SupportingPatientInformation {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationRecommendationRecommendationDateCriterion) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationRecommendationRecommendationDateCriterion) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"code\":"))
	if err != nil {
		return err
	}
	err = r.Code.marshalJSON(w)
	if err != nil {
		return err
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Value)
		if err != nil {
			return err
		}
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		p := primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_value\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *ImmunizationRecommendation) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *ImmunizationRecommendation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationRecommendation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationRecommendation element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendation element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendation element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendation element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendation element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendation element", t)
			}
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = v
		case "date":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "authority":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Authority = &v
		case "recommendation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendation element", t)
			}
			for d.More() {
				var v ImmunizationRecommendationRecommendation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Recommendation = append(r.Recommendation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendation element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationRecommendation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationRecommendation element", t)
	}
	return nil
}
func (r *ImmunizationRecommendationRecommendation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationRecommendationRecommendation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationRecommendationRecommendation element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "vaccineCode":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.VaccineCode = append(r.VaccineCode, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "targetDisease":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.TargetDisease = &v
		case "contraindicatedVaccineCode":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ContraindicatedVaccineCode = append(r.ContraindicatedVaccineCode, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "forecastStatus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ForecastStatus = v
		case "forecastReason":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ForecastReason = append(r.ForecastReason, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "dateCriterion":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v ImmunizationRecommendationRecommendationDateCriterion
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.DateCriterion = append(r.DateCriterion, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "series":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Series == nil {
				r.Series = &String{}
			}
			r.Series.Value = v.Value
		case "_series":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Series == nil {
				r.Series = &String{}
			}
			r.Series.Id = v.Id
			r.Series.Extension = v.Extension
		case "doseNumberPositiveInt":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DoseNumber != nil {
				r.DoseNumber = PositiveInt{
					Extension: r.DoseNumber.(PositiveInt).Extension,
					Id:        r.DoseNumber.(PositiveInt).Id,
					Value:     v.Value,
				}
			} else {
				r.DoseNumber = v
			}
		case "_doseNumberPositiveInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DoseNumber != nil {
				r.DoseNumber = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.DoseNumber.(PositiveInt).Value,
				}
			} else {
				r.DoseNumber = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "doseNumberString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DoseNumber != nil {
				r.DoseNumber = String{
					Extension: r.DoseNumber.(String).Extension,
					Id:        r.DoseNumber.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.DoseNumber = v
			}
		case "_doseNumberString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DoseNumber != nil {
				r.DoseNumber = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.DoseNumber.(String).Value,
				}
			} else {
				r.DoseNumber = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "seriesDosesPositiveInt":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SeriesDoses != nil {
				r.SeriesDoses = PositiveInt{
					Extension: r.SeriesDoses.(PositiveInt).Extension,
					Id:        r.SeriesDoses.(PositiveInt).Id,
					Value:     v.Value,
				}
			} else {
				r.SeriesDoses = v
			}
		case "_seriesDosesPositiveInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SeriesDoses != nil {
				r.SeriesDoses = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.SeriesDoses.(PositiveInt).Value,
				}
			} else {
				r.SeriesDoses = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "seriesDosesString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SeriesDoses != nil {
				r.SeriesDoses = String{
					Extension: r.SeriesDoses.(String).Extension,
					Id:        r.SeriesDoses.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.SeriesDoses = v
			}
		case "_seriesDosesString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SeriesDoses != nil {
				r.SeriesDoses = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.SeriesDoses.(String).Value,
				}
			} else {
				r.SeriesDoses = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "supportingImmunization":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SupportingImmunization = append(r.SupportingImmunization, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		case "supportingPatientInformation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendation element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SupportingPatientInformation = append(r.SupportingPatientInformation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendation element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationRecommendationRecommendation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationRecommendationRecommendation element", t)
	}
	return nil
}
func (r *ImmunizationRecommendationRecommendationDateCriterion) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationRecommendationRecommendationDateCriterion element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationRecommendationRecommendationDateCriterion element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendationDateCriterion element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendationDateCriterion element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationRecommendationRecommendationDateCriterion element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationRecommendationRecommendationDateCriterion element", t)
			}
		case "code":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Code = v
		case "value":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Value.Value = v.Value
		case "_value":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value.Id = v.Id
			r.Value.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationRecommendationRecommendationDateCriterion", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationRecommendationRecommendationDateCriterion element", t)
	}
	return nil
}
func (r ImmunizationRecommendation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "ImmunizationRecommendation"
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
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
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
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Authority, xml.StartElement{Name: xml.Name{Local: "authority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recommendation, xml.StartElement{Name: xml.Name{Local: "recommendation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationRecommendationRecommendation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.VaccineCode, xml.StartElement{Name: xml.Name{Local: "vaccineCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetDisease, xml.StartElement{Name: xml.Name{Local: "targetDisease"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContraindicatedVaccineCode, xml.StartElement{Name: xml.Name{Local: "contraindicatedVaccineCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ForecastStatus, xml.StartElement{Name: xml.Name{Local: "forecastStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ForecastReason, xml.StartElement{Name: xml.Name{Local: "forecastReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateCriterion, xml.StartElement{Name: xml.Name{Local: "dateCriterion"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Series, xml.StartElement{Name: xml.Name{Local: "series"}})
	if err != nil {
		return err
	}
	switch v := r.DoseNumber.(type) {
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseNumberPositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseNumberString"}})
		if err != nil {
			return err
		}
	}
	switch v := r.SeriesDoses.(type) {
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "seriesDosesPositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "seriesDosesString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.SupportingImmunization, xml.StartElement{Name: xml.Name{Local: "supportingImmunization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingPatientInformation, xml.StartElement{Name: xml.Name{Local: "supportingPatientInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationRecommendationRecommendationDateCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *ImmunizationRecommendation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Identifier = append(r.Identifier, v)
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = v
			case "authority":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authority = &v
			case "recommendation":
				var v ImmunizationRecommendationRecommendation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recommendation = append(r.Recommendation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationRecommendationRecommendation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "vaccineCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VaccineCode = append(r.VaccineCode, v)
			case "targetDisease":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetDisease = &v
			case "contraindicatedVaccineCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContraindicatedVaccineCode = append(r.ContraindicatedVaccineCode, v)
			case "forecastStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ForecastStatus = v
			case "forecastReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ForecastReason = append(r.ForecastReason, v)
			case "dateCriterion":
				var v ImmunizationRecommendationRecommendationDateCriterion
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateCriterion = append(r.DateCriterion, v)
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "series":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Series = &v
			case "doseNumberPositiveInt":
				if r.DoseNumber != nil {
					return fmt.Errorf("multiple values for field \"DoseNumber\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseNumber = v
			case "doseNumberString":
				if r.DoseNumber != nil {
					return fmt.Errorf("multiple values for field \"DoseNumber\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseNumber = v
			case "seriesDosesPositiveInt":
				if r.SeriesDoses != nil {
					return fmt.Errorf("multiple values for field \"SeriesDoses\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SeriesDoses = v
			case "seriesDosesString":
				if r.SeriesDoses != nil {
					return fmt.Errorf("multiple values for field \"SeriesDoses\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SeriesDoses = v
			case "supportingImmunization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingImmunization = append(r.SupportingImmunization, v)
			case "supportingPatientInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingPatientInformation = append(r.SupportingPatientInformation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationRecommendationRecommendationDateCriterion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "value":
				var v DateTime
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
func (r ImmunizationRecommendation) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		children = append(children, r.Patient)
	}
	if len(name) == 0 || slices.Contains(name, "date") {
		children = append(children, r.Date)
	}
	if len(name) == 0 || slices.Contains(name, "authority") {
		if r.Authority != nil {
			children = append(children, *r.Authority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "recommendation") {
		for _, v := range r.Recommendation {
			children = append(children, v)
		}
	}
	return children
}
func (r ImmunizationRecommendation) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Boolean")
}
func (r ImmunizationRecommendation) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to String")
}
func (r ImmunizationRecommendation) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Integer")
}
func (r ImmunizationRecommendation) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Decimal")
}
func (r ImmunizationRecommendation) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Date")
}
func (r ImmunizationRecommendation) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Time")
}
func (r ImmunizationRecommendation) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to DateTime")
}
func (r ImmunizationRecommendation) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationRecommendation to Quantity")
}
func (r ImmunizationRecommendation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendation
	switch other := other.(type) {
	case ImmunizationRecommendation:
		o = other
	case *ImmunizationRecommendation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendation
	switch other := other.(type) {
	case ImmunizationRecommendation:
		o = other
	case *ImmunizationRecommendation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "Patient",
			Type: "FHIR.Reference",
		}, {
			Name: "Date",
			Type: "FHIR.DateTime",
		}, {
			Name: "Authority",
			Type: "FHIR.Reference",
		}, {
			Name: "Recommendation",
			Type: "List<FHIR.ImmunizationRecommendationRecommendation>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationRecommendation",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationRecommendationRecommendation) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "vaccineCode") {
		for _, v := range r.VaccineCode {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "targetDisease") {
		if r.TargetDisease != nil {
			children = append(children, *r.TargetDisease)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contraindicatedVaccineCode") {
		for _, v := range r.ContraindicatedVaccineCode {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "forecastStatus") {
		children = append(children, r.ForecastStatus)
	}
	if len(name) == 0 || slices.Contains(name, "forecastReason") {
		for _, v := range r.ForecastReason {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dateCriterion") {
		for _, v := range r.DateCriterion {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "series") {
		if r.Series != nil {
			children = append(children, *r.Series)
		}
	}
	if len(name) == 0 || slices.Contains(name, "doseNumber") {
		if r.DoseNumber != nil {
			children = append(children, r.DoseNumber)
		}
	}
	if len(name) == 0 || slices.Contains(name, "seriesDoses") {
		if r.SeriesDoses != nil {
			children = append(children, r.SeriesDoses)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supportingImmunization") {
		for _, v := range r.SupportingImmunization {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supportingPatientInformation") {
		for _, v := range r.SupportingPatientInformation {
			children = append(children, v)
		}
	}
	return children
}
func (r ImmunizationRecommendationRecommendation) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Boolean")
}
func (r ImmunizationRecommendationRecommendation) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to String")
}
func (r ImmunizationRecommendationRecommendation) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Integer")
}
func (r ImmunizationRecommendationRecommendation) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Decimal")
}
func (r ImmunizationRecommendationRecommendation) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Date")
}
func (r ImmunizationRecommendationRecommendation) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Time")
}
func (r ImmunizationRecommendationRecommendation) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to DateTime")
}
func (r ImmunizationRecommendationRecommendation) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendation to Quantity")
}
func (r ImmunizationRecommendationRecommendation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendationRecommendation
	switch other := other.(type) {
	case ImmunizationRecommendationRecommendation:
		o = other
	case *ImmunizationRecommendationRecommendation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendationRecommendation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendationRecommendation
	switch other := other.(type) {
	case ImmunizationRecommendationRecommendation:
		o = other
	case *ImmunizationRecommendationRecommendation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendationRecommendation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "VaccineCode",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "TargetDisease",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "ContraindicatedVaccineCode",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "ForecastStatus",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "ForecastReason",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "DateCriterion",
			Type: "List<FHIR.ImmunizationRecommendationRecommendationDateCriterion>",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "Series",
			Type: "FHIR.String",
		}, {
			Name: "DoseNumber",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "SeriesDoses",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "SupportingImmunization",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "SupportingPatientInformation",
			Type: "List<FHIR.Reference>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationRecommendationRecommendation",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationRecommendationRecommendationDateCriterion) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		children = append(children, r.Code)
	}
	if len(name) == 0 || slices.Contains(name, "value") {
		children = append(children, r.Value)
	}
	return children
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Boolean")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to String")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Integer")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Decimal")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Date")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Time")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to DateTime")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationRecommendationRecommendationDateCriterion to Quantity")
}
func (r ImmunizationRecommendationRecommendationDateCriterion) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendationRecommendationDateCriterion
	switch other := other.(type) {
	case ImmunizationRecommendationRecommendationDateCriterion:
		o = other
	case *ImmunizationRecommendationRecommendationDateCriterion:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendationRecommendationDateCriterion) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o ImmunizationRecommendationRecommendationDateCriterion
	switch other := other.(type) {
	case ImmunizationRecommendationRecommendationDateCriterion:
		o = other
	case *ImmunizationRecommendationRecommendationDateCriterion:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r ImmunizationRecommendationRecommendationDateCriterion) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Code",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Value",
			Type: "FHIR.DateTime",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationRecommendationRecommendationDateCriterion",
			Namespace: "FHIR",
		},
	}
}
