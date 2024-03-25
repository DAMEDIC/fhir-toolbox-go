package r4

import (
	"encoding/json"
	"fmt"
)

// Base StructureDefinition for Population Type: A populatioof people with some set of grouping criteria.
type Population struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The age of the specific population.
	Age isPopulationAge
	// The gender of the specific population.
	Gender *CodeableConcept
	// Race of the specific population.
	Race *CodeableConcept
	// The existing physiological conditions of the specific population to which this applies.
	PhysiologicalCondition *CodeableConcept
}
type isPopulationAge interface {
	isPopulationAge()
}

func (r Range) isPopulationAge()           {}
func (r CodeableConcept) isPopulationAge() {}

type jsonPopulation struct {
	Id                     *string          `json:"id,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AgeRange               *Range           `json:"ageRange,omitempty"`
	AgeCodeableConcept     *CodeableConcept `json:"ageCodeableConcept,omitempty"`
	Gender                 *CodeableConcept `json:"gender,omitempty"`
	Race                   *CodeableConcept `json:"race,omitempty"`
	PhysiologicalCondition *CodeableConcept `json:"physiologicalCondition,omitempty"`
}

func (r Population) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Population) marshalJSON() jsonPopulation {
	m := jsonPopulation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Age.(type) {
	case Range:
		m.AgeRange = &v
	case *Range:
		m.AgeRange = v
	case CodeableConcept:
		m.AgeCodeableConcept = &v
	case *CodeableConcept:
		m.AgeCodeableConcept = v
	}
	m.Gender = r.Gender
	m.Race = r.Race
	m.PhysiologicalCondition = r.PhysiologicalCondition
	return m
}
func (r *Population) UnmarshalJSON(b []byte) error {
	var m jsonPopulation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Population) unmarshalJSON(m jsonPopulation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.AgeRange != nil {
		if r.Age != nil {
			return fmt.Errorf("multiple values for field \"Age\"")
		}
		v := m.AgeRange
		r.Age = v
	}
	if m.AgeCodeableConcept != nil {
		if r.Age != nil {
			return fmt.Errorf("multiple values for field \"Age\"")
		}
		v := m.AgeCodeableConcept
		r.Age = v
	}
	r.Gender = m.Gender
	r.Race = m.Race
	r.PhysiologicalCondition = m.PhysiologicalCondition
	return nil
}
func (r Population) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
