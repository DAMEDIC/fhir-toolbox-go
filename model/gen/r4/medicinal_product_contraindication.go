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
	model.Element
	isMedicinalProductContraindicationOtherTherapyMedication()
}

func (r CodeableConcept) isMedicinalProductContraindicationOtherTherapyMedication() {}
func (r Reference) isMedicinalProductContraindicationOtherTherapyMedication()       {}
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
func (r MedicinalProductContraindication) MemSize() int {
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
	for _, i := range r.Subject {
		s += i.MemSize()
	}
	s += (cap(r.Subject) - len(r.Subject)) * int(unsafe.Sizeof(Reference{}))
	if r.Disease != nil {
		s += r.Disease.MemSize()
	}
	if r.DiseaseStatus != nil {
		s += r.DiseaseStatus.MemSize()
	}
	for _, i := range r.Comorbidity {
		s += i.MemSize()
	}
	s += (cap(r.Comorbidity) - len(r.Comorbidity)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.TherapeuticIndication {
		s += i.MemSize()
	}
	s += (cap(r.TherapeuticIndication) - len(r.TherapeuticIndication)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.OtherTherapy {
		s += i.MemSize()
	}
	s += (cap(r.OtherTherapy) - len(r.OtherTherapy)) * int(unsafe.Sizeof(MedicinalProductContraindicationOtherTherapy{}))
	for _, i := range r.Population {
		s += i.MemSize()
	}
	s += (cap(r.Population) - len(r.Population)) * int(unsafe.Sizeof(Population{}))
	return s
}
func (r MedicinalProductContraindicationOtherTherapy) MemSize() int {
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
	s += r.TherapyRelationshipType.MemSize() - int(unsafe.Sizeof(r.TherapyRelationshipType))
	if r.Medication != nil {
		s += r.Medication.MemSize()
	}
	return s
}
func (r MedicinalProductContraindication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MedicinalProductContraindicationOtherTherapy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MedicinalProductContraindication) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"MedicinalProductContraindication\""))
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
	if len(r.Subject) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subject\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Subject {
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
	if r.Disease != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"disease\":"))
		if err != nil {
			return err
		}
		err = r.Disease.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DiseaseStatus != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"diseaseStatus\":"))
		if err != nil {
			return err
		}
		err = r.DiseaseStatus.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Comorbidity) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"comorbidity\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Comorbidity {
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
	if len(r.TherapeuticIndication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"therapeuticIndication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TherapeuticIndication {
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
	if len(r.OtherTherapy) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"otherTherapy\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.OtherTherapy {
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
	if len(r.Population) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"population\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Population {
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
func (r MedicinalProductContraindicationOtherTherapy) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MedicinalProductContraindicationOtherTherapy) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"therapyRelationshipType\":"))
	if err != nil {
		return err
	}
	err = r.TherapyRelationshipType.marshalJSON(w)
	if err != nil {
		return err
	}
	switch v := r.Medication.(type) {
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"medicationCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"medicationCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"medicationReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"medicationReference\":"))
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductContraindication) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *MedicinalProductContraindication) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MedicinalProductContraindication element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MedicinalProductContraindication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "subject":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Subject = append(r.Subject, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "disease":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Disease = &v
		case "diseaseStatus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DiseaseStatus = &v
		case "comorbidity":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Comorbidity = append(r.Comorbidity, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "therapeuticIndication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TherapeuticIndication = append(r.TherapeuticIndication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "otherTherapy":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
			}
			for d.More() {
				var v MedicinalProductContraindicationOtherTherapy
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.OtherTherapy = append(r.OtherTherapy, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		case "population":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindication element", t)
			}
			for d.More() {
				var v Population
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Population = append(r.Population, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindication element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in MedicinalProductContraindication", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MedicinalProductContraindication element", t)
	}
	return nil
}
func (r *MedicinalProductContraindicationOtherTherapy) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MedicinalProductContraindicationOtherTherapy element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MedicinalProductContraindicationOtherTherapy element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindicationOtherTherapy element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindicationOtherTherapy element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MedicinalProductContraindicationOtherTherapy element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MedicinalProductContraindicationOtherTherapy element", t)
			}
		case "therapyRelationshipType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.TherapyRelationshipType = v
		case "medicationCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Medication = v
		case "medicationReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Medication = v
		default:
			return fmt.Errorf("invalid field: %s in MedicinalProductContraindicationOtherTherapy", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MedicinalProductContraindicationOtherTherapy element", t)
	}
	return nil
}
func (r MedicinalProductContraindication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "MedicinalProductContraindication"
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
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Disease, xml.StartElement{Name: xml.Name{Local: "disease"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DiseaseStatus, xml.StartElement{Name: xml.Name{Local: "diseaseStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comorbidity, xml.StartElement{Name: xml.Name{Local: "comorbidity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TherapeuticIndication, xml.StartElement{Name: xml.Name{Local: "therapeuticIndication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OtherTherapy, xml.StartElement{Name: xml.Name{Local: "otherTherapy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Population, xml.StartElement{Name: xml.Name{Local: "population"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MedicinalProductContraindicationOtherTherapy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.TherapyRelationshipType, xml.StartElement{Name: xml.Name{Local: "therapyRelationshipType"}})
	if err != nil {
		return err
	}
	switch v := r.Medication.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "medicationCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "medicationReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductContraindication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = append(r.Subject, v)
			case "disease":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Disease = &v
			case "diseaseStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DiseaseStatus = &v
			case "comorbidity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comorbidity = append(r.Comorbidity, v)
			case "therapeuticIndication":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TherapeuticIndication = append(r.TherapeuticIndication, v)
			case "otherTherapy":
				var v MedicinalProductContraindicationOtherTherapy
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OtherTherapy = append(r.OtherTherapy, v)
			case "population":
				var v Population
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Population = append(r.Population, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MedicinalProductContraindicationOtherTherapy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "therapyRelationshipType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TherapyRelationshipType = v
			case "medicationCodeableConcept":
				if r.Medication != nil {
					return fmt.Errorf("multiple values for field \"Medication\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Medication = v
			case "medicationReference":
				if r.Medication != nil {
					return fmt.Errorf("multiple values for field \"Medication\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Medication = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductContraindication) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "subject") {
		for _, v := range r.Subject {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "disease") {
		if r.Disease != nil {
			children = append(children, *r.Disease)
		}
	}
	if len(name) == 0 || slices.Contains(name, "diseaseStatus") {
		if r.DiseaseStatus != nil {
			children = append(children, *r.DiseaseStatus)
		}
	}
	if len(name) == 0 || slices.Contains(name, "comorbidity") {
		for _, v := range r.Comorbidity {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "therapeuticIndication") {
		for _, v := range r.TherapeuticIndication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "otherTherapy") {
		for _, v := range r.OtherTherapy {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "population") {
		for _, v := range r.Population {
			children = append(children, v)
		}
	}
	return children
}
func (r MedicinalProductContraindication) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Boolean")
}
func (r MedicinalProductContraindication) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to String")
}
func (r MedicinalProductContraindication) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Integer")
}
func (r MedicinalProductContraindication) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Decimal")
}
func (r MedicinalProductContraindication) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Date")
}
func (r MedicinalProductContraindication) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Time")
}
func (r MedicinalProductContraindication) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to DateTime")
}
func (r MedicinalProductContraindication) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MedicinalProductContraindication to Quantity")
}
func (r MedicinalProductContraindication) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *MedicinalProductContraindication
	switch other := other.(type) {
	case MedicinalProductContraindication:
		o = &other
	case *MedicinalProductContraindication:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MedicinalProductContraindication) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *MedicinalProductContraindication
	switch other := other.(type) {
	case MedicinalProductContraindication:
		o = &other
	case *MedicinalProductContraindication:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MedicinalProductContraindication) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Subject",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Disease",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "DiseaseStatus",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Comorbidity",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "TherapeuticIndication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "OtherTherapy",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MedicinalProductContraindicationOtherTherapy",
				Namespace: "FHIR",
			},
		}, {
			Name: "Population",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Population",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "MedicinalProductContraindication",
			Namespace: "FHIR",
		},
	}
}
func (r MedicinalProductContraindicationOtherTherapy) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "therapyRelationshipType") {
		children = append(children, r.TherapyRelationshipType)
	}
	if len(name) == 0 || slices.Contains(name, "medication") {
		children = append(children, r.Medication)
	}
	return children
}
func (r MedicinalProductContraindicationOtherTherapy) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Boolean")
}
func (r MedicinalProductContraindicationOtherTherapy) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to String")
}
func (r MedicinalProductContraindicationOtherTherapy) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Integer")
}
func (r MedicinalProductContraindicationOtherTherapy) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Decimal")
}
func (r MedicinalProductContraindicationOtherTherapy) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Date")
}
func (r MedicinalProductContraindicationOtherTherapy) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Time")
}
func (r MedicinalProductContraindicationOtherTherapy) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to DateTime")
}
func (r MedicinalProductContraindicationOtherTherapy) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MedicinalProductContraindicationOtherTherapy to Quantity")
}
func (r MedicinalProductContraindicationOtherTherapy) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *MedicinalProductContraindicationOtherTherapy
	switch other := other.(type) {
	case MedicinalProductContraindicationOtherTherapy:
		o = &other
	case *MedicinalProductContraindicationOtherTherapy:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MedicinalProductContraindicationOtherTherapy) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *MedicinalProductContraindicationOtherTherapy
	switch other := other.(type) {
	case MedicinalProductContraindicationOtherTherapy:
		o = &other
	case *MedicinalProductContraindicationOtherTherapy:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MedicinalProductContraindicationOtherTherapy) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "TherapyRelationshipType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Medication",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MedicinalProductContraindicationOtherTherapy",
			Namespace: "FHIR",
		},
	}
}
