package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource is primarily used for the identification and definition of a medication for the purposes of prescribing, dispensing, and administering a medication as well as for making statements about medication use.
type Medication struct {
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
	// Business identifier for this medication.
	Identifier []Identifier
	// A code (or set of codes) that specify this medication, or a textual description if no code is available. Usage note: This could be a standard medication code such as a code from RxNorm, SNOMED CT, IDMP etc. It could also be a national or local formulary code, optionally with translations to other code systems.
	Code *CodeableConcept
	// A code to indicate if the medication is in active use.
	Status *Code
	// Describes the details of the manufacturer of the medication product.  This is not intended to represent the distributor of a medication product.
	Manufacturer *Reference
	// Describes the form of the item.  Powder; tablets; capsule.
	Form *CodeableConcept
	// Specific amount of the drug in the packaged product.  For example, when specifying a product that has the same strength (For example, Insulin glargine 100 unit per mL solution for injection), this attribute provides additional clarification of the package amount (For example, 3 mL, 10mL, etc.).
	Amount *Ratio
	// Identifies a particular constituent of interest in the product.
	Ingredient []MedicationIngredient
	// Information that only applies to packages (not products).
	Batch *MedicationBatch
}

func (r Medication) ResourceType() string {
	return "Medication"
}
func (r Medication) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedication struct {
	ResourceType                  string                 `json:"resourceType"`
	Id                            *Id                    `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement      `json:"_id,omitempty"`
	Meta                          *Meta                  `json:"meta,omitempty"`
	ImplicitRules                 *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                      *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement      `json:"_language,omitempty"`
	Text                          *Narrative             `json:"text,omitempty"`
	Contained                     []ContainedResource    `json:"contained,omitempty"`
	Extension                     []Extension            `json:"extension,omitempty"`
	ModifierExtension             []Extension            `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier           `json:"identifier,omitempty"`
	Code                          *CodeableConcept       `json:"code,omitempty"`
	Status                        *Code                  `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement      `json:"_status,omitempty"`
	Manufacturer                  *Reference             `json:"manufacturer,omitempty"`
	Form                          *CodeableConcept       `json:"form,omitempty"`
	Amount                        *Ratio                 `json:"amount,omitempty"`
	Ingredient                    []MedicationIngredient `json:"ingredient,omitempty"`
	Batch                         *MedicationBatch       `json:"batch,omitempty"`
}

func (r Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Medication) marshalJSON() jsonMedication {
	m := jsonMedication{}
	m.ResourceType = "Medication"
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
	m.Code = r.Code
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Manufacturer = r.Manufacturer
	m.Form = r.Form
	m.Amount = r.Amount
	m.Ingredient = r.Ingredient
	m.Batch = r.Batch
	return m
}
func (r *Medication) UnmarshalJSON(b []byte) error {
	var m jsonMedication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Medication) unmarshalJSON(m jsonMedication) error {
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
	r.Code = m.Code
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		if r.Status == nil {
			r.Status = &Code{}
		}
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Manufacturer = m.Manufacturer
	r.Form = m.Form
	r.Amount = m.Amount
	r.Ingredient = m.Ingredient
	r.Batch = m.Batch
	return nil
}
func (r Medication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Medication"
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
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manufacturer, xml.StartElement{Name: xml.Name{Local: "manufacturer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Form, xml.StartElement{Name: xml.Name{Local: "form"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Ingredient, xml.StartElement{Name: xml.Name{Local: "ingredient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Batch, xml.StartElement{Name: xml.Name{Local: "batch"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Medication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "manufacturer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manufacturer = &v
			case "form":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Form = &v
			case "amount":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			case "ingredient":
				var v MedicationIngredient
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Ingredient = append(r.Ingredient, v)
			case "batch":
				var v MedicationBatch
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Batch = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Medication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Identifies a particular constituent of interest in the product.
type MedicationIngredient struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The actual ingredient - either a substance (simple ingredient) or another medication of a medication.
	Item isMedicationIngredientItem
	// Indication of whether this ingredient affects the therapeutic action of the drug.
	IsActive *Boolean
	// Specifies how many (or how much) of the items there are in this Medication.  For example, 250 mg per tablet.  This is expressed as a ratio where the numerator is 250mg and the denominator is 1 tablet.
	Strength *Ratio
}
type isMedicationIngredientItem interface {
	isMedicationIngredientItem()
}

func (r CodeableConcept) isMedicationIngredientItem() {}
func (r Reference) isMedicationIngredientItem()       {}

type jsonMedicationIngredient struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	ItemCodeableConcept      *CodeableConcept  `json:"itemCodeableConcept,omitempty"`
	ItemReference            *Reference        `json:"itemReference,omitempty"`
	IsActive                 *Boolean          `json:"isActive,omitempty"`
	IsActivePrimitiveElement *primitiveElement `json:"_isActive,omitempty"`
	Strength                 *Ratio            `json:"strength,omitempty"`
}

func (r MedicationIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationIngredient) marshalJSON() jsonMedicationIngredient {
	m := jsonMedicationIngredient{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Item.(type) {
	case CodeableConcept:
		m.ItemCodeableConcept = &v
	case *CodeableConcept:
		m.ItemCodeableConcept = v
	case Reference:
		m.ItemReference = &v
	case *Reference:
		m.ItemReference = v
	}
	if r.IsActive != nil && r.IsActive.Value != nil {
		m.IsActive = r.IsActive
	}
	if r.IsActive != nil && (r.IsActive.Id != nil || r.IsActive.Extension != nil) {
		m.IsActivePrimitiveElement = &primitiveElement{Id: r.IsActive.Id, Extension: r.IsActive.Extension}
	}
	m.Strength = r.Strength
	return m
}
func (r *MedicationIngredient) UnmarshalJSON(b []byte) error {
	var m jsonMedicationIngredient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationIngredient) unmarshalJSON(m jsonMedicationIngredient) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ItemCodeableConcept != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemCodeableConcept
		r.Item = v
	}
	if m.ItemReference != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemReference
		r.Item = v
	}
	r.IsActive = m.IsActive
	if m.IsActivePrimitiveElement != nil {
		if r.IsActive == nil {
			r.IsActive = &Boolean{}
		}
		r.IsActive.Id = m.IsActivePrimitiveElement.Id
		r.IsActive.Extension = m.IsActivePrimitiveElement.Extension
	}
	r.Strength = m.Strength
	return nil
}
func (r MedicationIngredient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Item.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "itemCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "itemReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.IsActive, xml.StartElement{Name: xml.Name{Local: "isActive"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Strength, xml.StartElement{Name: xml.Name{Local: "strength"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationIngredient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemCodeableConcept":
				if r.Item != nil {
					return fmt.Errorf("multiple values for field \"Item\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = v
			case "itemReference":
				if r.Item != nil {
					return fmt.Errorf("multiple values for field \"Item\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = v
			case "isActive":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IsActive = &v
			case "strength":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strength = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationIngredient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information that only applies to packages (not products).
type MedicationBatch struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The assigned lot number of a batch of the specified product.
	LotNumber *String
	// When this specific batch of product will expire.
	ExpirationDate *DateTime
}
type jsonMedicationBatch struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	LotNumber                      *String           `json:"lotNumber,omitempty"`
	LotNumberPrimitiveElement      *primitiveElement `json:"_lotNumber,omitempty"`
	ExpirationDate                 *DateTime         `json:"expirationDate,omitempty"`
	ExpirationDatePrimitiveElement *primitiveElement `json:"_expirationDate,omitempty"`
}

func (r MedicationBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationBatch) marshalJSON() jsonMedicationBatch {
	m := jsonMedicationBatch{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.LotNumber != nil && r.LotNumber.Value != nil {
		m.LotNumber = r.LotNumber
	}
	if r.LotNumber != nil && (r.LotNumber.Id != nil || r.LotNumber.Extension != nil) {
		m.LotNumberPrimitiveElement = &primitiveElement{Id: r.LotNumber.Id, Extension: r.LotNumber.Extension}
	}
	if r.ExpirationDate != nil && r.ExpirationDate.Value != nil {
		m.ExpirationDate = r.ExpirationDate
	}
	if r.ExpirationDate != nil && (r.ExpirationDate.Id != nil || r.ExpirationDate.Extension != nil) {
		m.ExpirationDatePrimitiveElement = &primitiveElement{Id: r.ExpirationDate.Id, Extension: r.ExpirationDate.Extension}
	}
	return m
}
func (r *MedicationBatch) UnmarshalJSON(b []byte) error {
	var m jsonMedicationBatch
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationBatch) unmarshalJSON(m jsonMedicationBatch) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.LotNumber = m.LotNumber
	if m.LotNumberPrimitiveElement != nil {
		if r.LotNumber == nil {
			r.LotNumber = &String{}
		}
		r.LotNumber.Id = m.LotNumberPrimitiveElement.Id
		r.LotNumber.Extension = m.LotNumberPrimitiveElement.Extension
	}
	r.ExpirationDate = m.ExpirationDate
	if m.ExpirationDatePrimitiveElement != nil {
		if r.ExpirationDate == nil {
			r.ExpirationDate = &DateTime{}
		}
		r.ExpirationDate.Id = m.ExpirationDatePrimitiveElement.Id
		r.ExpirationDate.Extension = m.ExpirationDatePrimitiveElement.Extension
	}
	return nil
}
func (r MedicationBatch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.LotNumber, xml.StartElement{Name: xml.Name{Local: "lotNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExpirationDate, xml.StartElement{Name: xml.Name{Local: "expirationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationBatch) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "lotNumber":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LotNumber = &v
			case "expirationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExpirationDate = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationBatch) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
