package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Information about a medication that is used to support knowledge.
type MedicationKnowledge struct {
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
	// A code that specifies this medication, or a textual description if no code is available. Usage note: This could be a standard medication code such as a code from RxNorm, SNOMED CT, IDMP etc. It could also be a national or local formulary code, optionally with translations to other code systems.
	Code *CodeableConcept
	// A code to indicate if the medication is in active use.  The status refers to the validity about the information of the medication and not to its medicinal properties.
	Status *Code
	// Describes the details of the manufacturer of the medication product.  This is not intended to represent the distributor of a medication product.
	Manufacturer *Reference
	// Describes the form of the item.  Powder; tablets; capsule.
	DoseForm *CodeableConcept
	// Specific amount of the drug in the packaged product.  For example, when specifying a product that has the same strength (For example, Insulin glargine 100 unit per mL solution for injection), this attribute provides additional clarification of the package amount (For example, 3 mL, 10mL, etc.).
	Amount *Quantity
	// Additional names for a medication, for example, the name(s) given to a medication in different countries.  For example, acetaminophen and paracetamol or salbutamol and albuterol.
	Synonym []String
	// Associated or related knowledge about a medication.
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge
	// Associated or related medications.  For example, if the medication is a branded product (e.g. Crestor), this is the Therapeutic Moeity (e.g. Rosuvastatin) or if this is a generic medication (e.g. Rosuvastatin), this would link to a branded product (e.g. Crestor).
	AssociatedMedication []Reference
	// Category of the medication or product (e.g. branded product, therapeutic moeity, generic product, innovator product, etc.).
	ProductType []CodeableConcept
	// Associated documentation about the medication.
	Monograph []MedicationKnowledgeMonograph
	// Identifies a particular constituent of interest in the product.
	Ingredient []MedicationKnowledgeIngredient
	// The instructions for preparing the medication.
	PreparationInstruction *Markdown
	// The intended or approved route of administration.
	IntendedRoute []CodeableConcept
	// The price of the medication.
	Cost []MedicationKnowledgeCost
	// The program under which the medication is reviewed.
	MonitoringProgram []MedicationKnowledgeMonitoringProgram
	// Guidelines for the administration of the medication.
	AdministrationGuidelines []MedicationKnowledgeAdministrationGuidelines
	// Categorization of the medication within a formulary or classification system.
	MedicineClassification []MedicationKnowledgeMedicineClassification
	// Information that only applies to packages (not products).
	Packaging *MedicationKnowledgePackaging
	// Specifies descriptive properties of the medicine, such as color, shape, imprints, etc.
	DrugCharacteristic []MedicationKnowledgeDrugCharacteristic
	// Potential clinical issue with or between medication(s) (for example, drug-drug interaction, drug-disease contraindication, drug-allergy interaction, etc.).
	Contraindication []Reference
	// Regulatory information about a medication.
	Regulatory []MedicationKnowledgeRegulatory
	// The time course of drug absorption, distribution, metabolism and excretion of a medication from the body.
	Kinetics []MedicationKnowledgeKinetics
}

func (r MedicationKnowledge) ResourceType() string {
	return "MedicationKnowledge"
}
func (r MedicationKnowledge) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicationKnowledge struct {
	ResourceType                           string                                          `json:"resourceType"`
	Id                                     *Id                                             `json:"id,omitempty"`
	IdPrimitiveElement                     *primitiveElement                               `json:"_id,omitempty"`
	Meta                                   *Meta                                           `json:"meta,omitempty"`
	ImplicitRules                          *Uri                                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement          *primitiveElement                               `json:"_implicitRules,omitempty"`
	Language                               *Code                                           `json:"language,omitempty"`
	LanguagePrimitiveElement               *primitiveElement                               `json:"_language,omitempty"`
	Text                                   *Narrative                                      `json:"text,omitempty"`
	Contained                              []ContainedResource                             `json:"contained,omitempty"`
	Extension                              []Extension                                     `json:"extension,omitempty"`
	ModifierExtension                      []Extension                                     `json:"modifierExtension,omitempty"`
	Code                                   *CodeableConcept                                `json:"code,omitempty"`
	Status                                 *Code                                           `json:"status,omitempty"`
	StatusPrimitiveElement                 *primitiveElement                               `json:"_status,omitempty"`
	Manufacturer                           *Reference                                      `json:"manufacturer,omitempty"`
	DoseForm                               *CodeableConcept                                `json:"doseForm,omitempty"`
	Amount                                 *Quantity                                       `json:"amount,omitempty"`
	Synonym                                []String                                        `json:"synonym,omitempty"`
	SynonymPrimitiveElement                []*primitiveElement                             `json:"_synonym,omitempty"`
	RelatedMedicationKnowledge             []MedicationKnowledgeRelatedMedicationKnowledge `json:"relatedMedicationKnowledge,omitempty"`
	AssociatedMedication                   []Reference                                     `json:"associatedMedication,omitempty"`
	ProductType                            []CodeableConcept                               `json:"productType,omitempty"`
	Monograph                              []MedicationKnowledgeMonograph                  `json:"monograph,omitempty"`
	Ingredient                             []MedicationKnowledgeIngredient                 `json:"ingredient,omitempty"`
	PreparationInstruction                 *Markdown                                       `json:"preparationInstruction,omitempty"`
	PreparationInstructionPrimitiveElement *primitiveElement                               `json:"_preparationInstruction,omitempty"`
	IntendedRoute                          []CodeableConcept                               `json:"intendedRoute,omitempty"`
	Cost                                   []MedicationKnowledgeCost                       `json:"cost,omitempty"`
	MonitoringProgram                      []MedicationKnowledgeMonitoringProgram          `json:"monitoringProgram,omitempty"`
	AdministrationGuidelines               []MedicationKnowledgeAdministrationGuidelines   `json:"administrationGuidelines,omitempty"`
	MedicineClassification                 []MedicationKnowledgeMedicineClassification     `json:"medicineClassification,omitempty"`
	Packaging                              *MedicationKnowledgePackaging                   `json:"packaging,omitempty"`
	DrugCharacteristic                     []MedicationKnowledgeDrugCharacteristic         `json:"drugCharacteristic,omitempty"`
	Contraindication                       []Reference                                     `json:"contraindication,omitempty"`
	Regulatory                             []MedicationKnowledgeRegulatory                 `json:"regulatory,omitempty"`
	Kinetics                               []MedicationKnowledgeKinetics                   `json:"kinetics,omitempty"`
}

func (r MedicationKnowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledge) marshalJSON() jsonMedicationKnowledge {
	m := jsonMedicationKnowledge{}
	m.ResourceType = "MedicationKnowledge"
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
	m.Code = r.Code
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Manufacturer = r.Manufacturer
	m.DoseForm = r.DoseForm
	m.Amount = r.Amount
	anySynonymValue := false
	for _, e := range r.Synonym {
		if e.Value != nil {
			anySynonymValue = true
			break
		}
	}
	if anySynonymValue {
		m.Synonym = r.Synonym
	}
	anySynonymIdOrExtension := false
	for _, e := range r.Synonym {
		if e.Id != nil || e.Extension != nil {
			anySynonymIdOrExtension = true
			break
		}
	}
	if anySynonymIdOrExtension {
		m.SynonymPrimitiveElement = make([]*primitiveElement, 0, len(r.Synonym))
		for _, e := range r.Synonym {
			if e.Id != nil || e.Extension != nil {
				m.SynonymPrimitiveElement = append(m.SynonymPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SynonymPrimitiveElement = append(m.SynonymPrimitiveElement, nil)
			}
		}
	}
	m.RelatedMedicationKnowledge = r.RelatedMedicationKnowledge
	m.AssociatedMedication = r.AssociatedMedication
	m.ProductType = r.ProductType
	m.Monograph = r.Monograph
	m.Ingredient = r.Ingredient
	if r.PreparationInstruction != nil && r.PreparationInstruction.Value != nil {
		m.PreparationInstruction = r.PreparationInstruction
	}
	if r.PreparationInstruction != nil && (r.PreparationInstruction.Id != nil || r.PreparationInstruction.Extension != nil) {
		m.PreparationInstructionPrimitiveElement = &primitiveElement{Id: r.PreparationInstruction.Id, Extension: r.PreparationInstruction.Extension}
	}
	m.IntendedRoute = r.IntendedRoute
	m.Cost = r.Cost
	m.MonitoringProgram = r.MonitoringProgram
	m.AdministrationGuidelines = r.AdministrationGuidelines
	m.MedicineClassification = r.MedicineClassification
	m.Packaging = r.Packaging
	m.DrugCharacteristic = r.DrugCharacteristic
	m.Contraindication = r.Contraindication
	m.Regulatory = r.Regulatory
	m.Kinetics = r.Kinetics
	return m
}
func (r *MedicationKnowledge) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledge
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledge) unmarshalJSON(m jsonMedicationKnowledge) error {
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
	r.DoseForm = m.DoseForm
	r.Amount = m.Amount
	r.Synonym = m.Synonym
	for i, e := range m.SynonymPrimitiveElement {
		if len(r.Synonym) <= i {
			r.Synonym = append(r.Synonym, String{})
		}
		if e != nil {
			r.Synonym[i].Id = e.Id
			r.Synonym[i].Extension = e.Extension
		}
	}
	r.RelatedMedicationKnowledge = m.RelatedMedicationKnowledge
	r.AssociatedMedication = m.AssociatedMedication
	r.ProductType = m.ProductType
	r.Monograph = m.Monograph
	r.Ingredient = m.Ingredient
	r.PreparationInstruction = m.PreparationInstruction
	if m.PreparationInstructionPrimitiveElement != nil {
		if r.PreparationInstruction == nil {
			r.PreparationInstruction = &Markdown{}
		}
		r.PreparationInstruction.Id = m.PreparationInstructionPrimitiveElement.Id
		r.PreparationInstruction.Extension = m.PreparationInstructionPrimitiveElement.Extension
	}
	r.IntendedRoute = m.IntendedRoute
	r.Cost = m.Cost
	r.MonitoringProgram = m.MonitoringProgram
	r.AdministrationGuidelines = m.AdministrationGuidelines
	r.MedicineClassification = m.MedicineClassification
	r.Packaging = m.Packaging
	r.DrugCharacteristic = m.DrugCharacteristic
	r.Contraindication = m.Contraindication
	r.Regulatory = m.Regulatory
	r.Kinetics = m.Kinetics
	return nil
}
func (r MedicationKnowledge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicationKnowledge"
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
	err = e.EncodeElement(r.DoseForm, xml.StartElement{Name: xml.Name{Local: "doseForm"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Synonym, xml.StartElement{Name: xml.Name{Local: "synonym"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatedMedicationKnowledge, xml.StartElement{Name: xml.Name{Local: "relatedMedicationKnowledge"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AssociatedMedication, xml.StartElement{Name: xml.Name{Local: "associatedMedication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductType, xml.StartElement{Name: xml.Name{Local: "productType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Monograph, xml.StartElement{Name: xml.Name{Local: "monograph"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Ingredient, xml.StartElement{Name: xml.Name{Local: "ingredient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreparationInstruction, xml.StartElement{Name: xml.Name{Local: "preparationInstruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntendedRoute, xml.StartElement{Name: xml.Name{Local: "intendedRoute"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Cost, xml.StartElement{Name: xml.Name{Local: "cost"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MonitoringProgram, xml.StartElement{Name: xml.Name{Local: "monitoringProgram"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdministrationGuidelines, xml.StartElement{Name: xml.Name{Local: "administrationGuidelines"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MedicineClassification, xml.StartElement{Name: xml.Name{Local: "medicineClassification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Packaging, xml.StartElement{Name: xml.Name{Local: "packaging"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DrugCharacteristic, xml.StartElement{Name: xml.Name{Local: "drugCharacteristic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contraindication, xml.StartElement{Name: xml.Name{Local: "contraindication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Regulatory, xml.StartElement{Name: xml.Name{Local: "regulatory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Kinetics, xml.StartElement{Name: xml.Name{Local: "kinetics"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "doseForm":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseForm = &v
			case "amount":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			case "synonym":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Synonym = append(r.Synonym, v)
			case "relatedMedicationKnowledge":
				var v MedicationKnowledgeRelatedMedicationKnowledge
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatedMedicationKnowledge = append(r.RelatedMedicationKnowledge, v)
			case "associatedMedication":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AssociatedMedication = append(r.AssociatedMedication, v)
			case "productType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductType = append(r.ProductType, v)
			case "monograph":
				var v MedicationKnowledgeMonograph
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Monograph = append(r.Monograph, v)
			case "ingredient":
				var v MedicationKnowledgeIngredient
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Ingredient = append(r.Ingredient, v)
			case "preparationInstruction":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreparationInstruction = &v
			case "intendedRoute":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntendedRoute = append(r.IntendedRoute, v)
			case "cost":
				var v MedicationKnowledgeCost
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Cost = append(r.Cost, v)
			case "monitoringProgram":
				var v MedicationKnowledgeMonitoringProgram
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MonitoringProgram = append(r.MonitoringProgram, v)
			case "administrationGuidelines":
				var v MedicationKnowledgeAdministrationGuidelines
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdministrationGuidelines = append(r.AdministrationGuidelines, v)
			case "medicineClassification":
				var v MedicationKnowledgeMedicineClassification
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MedicineClassification = append(r.MedicineClassification, v)
			case "packaging":
				var v MedicationKnowledgePackaging
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Packaging = &v
			case "drugCharacteristic":
				var v MedicationKnowledgeDrugCharacteristic
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DrugCharacteristic = append(r.DrugCharacteristic, v)
			case "contraindication":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contraindication = append(r.Contraindication, v)
			case "regulatory":
				var v MedicationKnowledgeRegulatory
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Regulatory = append(r.Regulatory, v)
			case "kinetics":
				var v MedicationKnowledgeKinetics
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Kinetics = append(r.Kinetics, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledge) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Associated or related knowledge about a medication.
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of the associated medication knowledge reference.
	Type CodeableConcept
	// Associated documentation about the associated medication knowledge.
	Reference []Reference
}
type jsonMedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type,omitempty"`
	Reference         []Reference     `json:"reference,omitempty"`
}

func (r MedicationKnowledgeRelatedMedicationKnowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeRelatedMedicationKnowledge) marshalJSON() jsonMedicationKnowledgeRelatedMedicationKnowledge {
	m := jsonMedicationKnowledgeRelatedMedicationKnowledge{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Reference = r.Reference
	return m
}
func (r *MedicationKnowledgeRelatedMedicationKnowledge) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeRelatedMedicationKnowledge
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeRelatedMedicationKnowledge) unmarshalJSON(m jsonMedicationKnowledgeRelatedMedicationKnowledge) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Reference = m.Reference
	return nil
}
func (r MedicationKnowledgeRelatedMedicationKnowledge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeRelatedMedicationKnowledge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = append(r.Reference, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeRelatedMedicationKnowledge) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Associated documentation about the medication.
type MedicationKnowledgeMonograph struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of documentation about the medication. (e.g. professional monograph, patient education monograph).
	Type *CodeableConcept
	// Associated documentation about the medication.
	Source *Reference
}
type jsonMedicationKnowledgeMonograph struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
}

func (r MedicationKnowledgeMonograph) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeMonograph) marshalJSON() jsonMedicationKnowledgeMonograph {
	m := jsonMedicationKnowledgeMonograph{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Source = r.Source
	return m
}
func (r *MedicationKnowledgeMonograph) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeMonograph
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeMonograph) unmarshalJSON(m jsonMedicationKnowledgeMonograph) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Source = m.Source
	return nil
}
func (r MedicationKnowledgeMonograph) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeMonograph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeMonograph) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Identifies a particular constituent of interest in the product.
type MedicationKnowledgeIngredient struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The actual ingredient - either a substance (simple ingredient) or another medication.
	Item isMedicationKnowledgeIngredientItem
	// Indication of whether this ingredient affects the therapeutic action of the drug.
	IsActive *Boolean
	// Specifies how many (or how much) of the items there are in this Medication.  For example, 250 mg per tablet.  This is expressed as a ratio where the numerator is 250mg and the denominator is 1 tablet.
	Strength *Ratio
}
type isMedicationKnowledgeIngredientItem interface {
	isMedicationKnowledgeIngredientItem()
}

func (r CodeableConcept) isMedicationKnowledgeIngredientItem() {}
func (r Reference) isMedicationKnowledgeIngredientItem()       {}

type jsonMedicationKnowledgeIngredient struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	ItemCodeableConcept      *CodeableConcept  `json:"itemCodeableConcept,omitempty"`
	ItemReference            *Reference        `json:"itemReference,omitempty"`
	IsActive                 *Boolean          `json:"isActive,omitempty"`
	IsActivePrimitiveElement *primitiveElement `json:"_isActive,omitempty"`
	Strength                 *Ratio            `json:"strength,omitempty"`
}

func (r MedicationKnowledgeIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeIngredient) marshalJSON() jsonMedicationKnowledgeIngredient {
	m := jsonMedicationKnowledgeIngredient{}
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
func (r *MedicationKnowledgeIngredient) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeIngredient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeIngredient) unmarshalJSON(m jsonMedicationKnowledgeIngredient) error {
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
func (r MedicationKnowledgeIngredient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r *MedicationKnowledgeIngredient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (r MedicationKnowledgeIngredient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The price of the medication.
type MedicationKnowledgeCost struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of the cost information.  For example, manufacturers' cost, patient cost, claim reimbursement cost, actual acquisition cost.
	Type CodeableConcept
	// The source or owner that assigns the price to the medication.
	Source *String
	// The price of the medication.
	Cost Money
}
type jsonMedicationKnowledgeCost struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Type                   CodeableConcept   `json:"type,omitempty"`
	Source                 *String           `json:"source,omitempty"`
	SourcePrimitiveElement *primitiveElement `json:"_source,omitempty"`
	Cost                   Money             `json:"cost,omitempty"`
}

func (r MedicationKnowledgeCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeCost) marshalJSON() jsonMedicationKnowledgeCost {
	m := jsonMedicationKnowledgeCost{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Source != nil && r.Source.Value != nil {
		m.Source = r.Source
	}
	if r.Source != nil && (r.Source.Id != nil || r.Source.Extension != nil) {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	m.Cost = r.Cost
	return m
}
func (r *MedicationKnowledgeCost) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeCost
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeCost) unmarshalJSON(m jsonMedicationKnowledgeCost) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		if r.Source == nil {
			r.Source = &String{}
		}
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.Cost = m.Cost
	return nil
}
func (r MedicationKnowledgeCost) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Cost, xml.StartElement{Name: xml.Name{Local: "cost"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeCost) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "source":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			case "cost":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Cost = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeCost) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The program under which the medication is reviewed.
type MedicationKnowledgeMonitoringProgram struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of program under which the medication is monitored.
	Type *CodeableConcept
	// Name of the reviewing program.
	Name *String
}
type jsonMedicationKnowledgeMonitoringProgram struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
	Name                 *String           `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
}

func (r MedicationKnowledgeMonitoringProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeMonitoringProgram) marshalJSON() jsonMedicationKnowledgeMonitoringProgram {
	m := jsonMedicationKnowledgeMonitoringProgram{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	return m
}
func (r *MedicationKnowledgeMonitoringProgram) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeMonitoringProgram
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeMonitoringProgram) unmarshalJSON(m jsonMedicationKnowledgeMonitoringProgram) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	return nil
}
func (r MedicationKnowledgeMonitoringProgram) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeMonitoringProgram) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeMonitoringProgram) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Guidelines for the administration of the medication.
type MedicationKnowledgeAdministrationGuidelines struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Dosage for the medication for the specific guidelines.
	Dosage []MedicationKnowledgeAdministrationGuidelinesDosage
	// Indication for use that apply to the specific administration guidelines.
	Indication isMedicationKnowledgeAdministrationGuidelinesIndication
	// Characteristics of the patient that are relevant to the administration guidelines (for example, height, weight, gender, etc.).
	PatientCharacteristics []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics
}
type isMedicationKnowledgeAdministrationGuidelinesIndication interface {
	isMedicationKnowledgeAdministrationGuidelinesIndication()
}

func (r CodeableConcept) isMedicationKnowledgeAdministrationGuidelinesIndication() {}
func (r Reference) isMedicationKnowledgeAdministrationGuidelinesIndication()       {}

type jsonMedicationKnowledgeAdministrationGuidelines struct {
	Id                        *string                                                             `json:"id,omitempty"`
	Extension                 []Extension                                                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                         `json:"modifierExtension,omitempty"`
	Dosage                    []MedicationKnowledgeAdministrationGuidelinesDosage                 `json:"dosage,omitempty"`
	IndicationCodeableConcept *CodeableConcept                                                    `json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference                                                          `json:"indicationReference,omitempty"`
	PatientCharacteristics    []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `json:"patientCharacteristics,omitempty"`
}

func (r MedicationKnowledgeAdministrationGuidelines) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeAdministrationGuidelines) marshalJSON() jsonMedicationKnowledgeAdministrationGuidelines {
	m := jsonMedicationKnowledgeAdministrationGuidelines{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Dosage = r.Dosage
	switch v := r.Indication.(type) {
	case CodeableConcept:
		m.IndicationCodeableConcept = &v
	case *CodeableConcept:
		m.IndicationCodeableConcept = v
	case Reference:
		m.IndicationReference = &v
	case *Reference:
		m.IndicationReference = v
	}
	m.PatientCharacteristics = r.PatientCharacteristics
	return m
}
func (r *MedicationKnowledgeAdministrationGuidelines) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeAdministrationGuidelines
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeAdministrationGuidelines) unmarshalJSON(m jsonMedicationKnowledgeAdministrationGuidelines) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Dosage = m.Dosage
	if m.IndicationCodeableConcept != nil {
		if r.Indication != nil {
			return fmt.Errorf("multiple values for field \"Indication\"")
		}
		v := m.IndicationCodeableConcept
		r.Indication = v
	}
	if m.IndicationReference != nil {
		if r.Indication != nil {
			return fmt.Errorf("multiple values for field \"Indication\"")
		}
		v := m.IndicationReference
		r.Indication = v
	}
	r.PatientCharacteristics = m.PatientCharacteristics
	return nil
}
func (r MedicationKnowledgeAdministrationGuidelines) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Dosage, xml.StartElement{Name: xml.Name{Local: "dosage"}})
	if err != nil {
		return err
	}
	switch v := r.Indication.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "indicationCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "indicationReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.PatientCharacteristics, xml.StartElement{Name: xml.Name{Local: "patientCharacteristics"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeAdministrationGuidelines) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "dosage":
				var v MedicationKnowledgeAdministrationGuidelinesDosage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Dosage = append(r.Dosage, v)
			case "indicationCodeableConcept":
				if r.Indication != nil {
					return fmt.Errorf("multiple values for field \"Indication\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Indication = v
			case "indicationReference":
				if r.Indication != nil {
					return fmt.Errorf("multiple values for field \"Indication\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Indication = v
			case "patientCharacteristics":
				var v MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PatientCharacteristics = append(r.PatientCharacteristics, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeAdministrationGuidelines) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Dosage for the medication for the specific guidelines.
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of dosage (for example, prophylaxis, maintenance, therapeutic, etc.).
	Type CodeableConcept
	// Dosage for the medication for the specific guidelines.
	Dosage []Dosage
}
type jsonMedicationKnowledgeAdministrationGuidelinesDosage struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type,omitempty"`
	Dosage            []Dosage        `json:"dosage,omitempty"`
}

func (r MedicationKnowledgeAdministrationGuidelinesDosage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeAdministrationGuidelinesDosage) marshalJSON() jsonMedicationKnowledgeAdministrationGuidelinesDosage {
	m := jsonMedicationKnowledgeAdministrationGuidelinesDosage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Dosage = r.Dosage
	return m
}
func (r *MedicationKnowledgeAdministrationGuidelinesDosage) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeAdministrationGuidelinesDosage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeAdministrationGuidelinesDosage) unmarshalJSON(m jsonMedicationKnowledgeAdministrationGuidelinesDosage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Dosage = m.Dosage
	return nil
}
func (r MedicationKnowledgeAdministrationGuidelinesDosage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Dosage, xml.StartElement{Name: xml.Name{Local: "dosage"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeAdministrationGuidelinesDosage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "dosage":
				var v Dosage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Dosage = append(r.Dosage, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeAdministrationGuidelinesDosage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Characteristics of the patient that are relevant to the administration guidelines (for example, height, weight, gender, etc.).
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specific characteristic that is relevant to the administration guideline (e.g. height, weight, gender).
	Characteristic isMedicationKnowledgeAdministrationGuidelinesPatientCharacteristicsCharacteristic
	// The specific characteristic (e.g. height, weight, gender, etc.).
	Value []String
}
type isMedicationKnowledgeAdministrationGuidelinesPatientCharacteristicsCharacteristic interface {
	isMedicationKnowledgeAdministrationGuidelinesPatientCharacteristicsCharacteristic()
}

func (r CodeableConcept) isMedicationKnowledgeAdministrationGuidelinesPatientCharacteristicsCharacteristic() {
}
func (r Quantity) isMedicationKnowledgeAdministrationGuidelinesPatientCharacteristicsCharacteristic() {
}

type jsonMedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	Id                            *string             `json:"id,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	CharacteristicCodeableConcept *CodeableConcept    `json:"characteristicCodeableConcept,omitempty"`
	CharacteristicQuantity        *Quantity           `json:"characteristicQuantity,omitempty"`
	Value                         []String            `json:"value,omitempty"`
	ValuePrimitiveElement         []*primitiveElement `json:"_value,omitempty"`
}

func (r MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) marshalJSON() jsonMedicationKnowledgeAdministrationGuidelinesPatientCharacteristics {
	m := jsonMedicationKnowledgeAdministrationGuidelinesPatientCharacteristics{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Characteristic.(type) {
	case CodeableConcept:
		m.CharacteristicCodeableConcept = &v
	case *CodeableConcept:
		m.CharacteristicCodeableConcept = v
	case Quantity:
		m.CharacteristicQuantity = &v
	case *Quantity:
		m.CharacteristicQuantity = v
	}
	anyValueValue := false
	for _, e := range r.Value {
		if e.Value != nil {
			anyValueValue = true
			break
		}
	}
	if anyValueValue {
		m.Value = r.Value
	}
	anyValueIdOrExtension := false
	for _, e := range r.Value {
		if e.Id != nil || e.Extension != nil {
			anyValueIdOrExtension = true
			break
		}
	}
	if anyValueIdOrExtension {
		m.ValuePrimitiveElement = make([]*primitiveElement, 0, len(r.Value))
		for _, e := range r.Value {
			if e.Id != nil || e.Extension != nil {
				m.ValuePrimitiveElement = append(m.ValuePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ValuePrimitiveElement = append(m.ValuePrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeAdministrationGuidelinesPatientCharacteristics
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) unmarshalJSON(m jsonMedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.CharacteristicCodeableConcept != nil {
		if r.Characteristic != nil {
			return fmt.Errorf("multiple values for field \"Characteristic\"")
		}
		v := m.CharacteristicCodeableConcept
		r.Characteristic = v
	}
	if m.CharacteristicQuantity != nil {
		if r.Characteristic != nil {
			return fmt.Errorf("multiple values for field \"Characteristic\"")
		}
		v := m.CharacteristicQuantity
		r.Characteristic = v
	}
	r.Value = m.Value
	for i, e := range m.ValuePrimitiveElement {
		if len(r.Value) <= i {
			r.Value = append(r.Value, String{})
		}
		if e != nil {
			r.Value[i].Id = e.Id
			r.Value[i].Extension = e.Extension
		}
	}
	return nil
}
func (r MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Characteristic.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "characteristicCodeableConcept"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "characteristicQuantity"}})
		if err != nil {
			return err
		}
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
func (r *MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "characteristicCodeableConcept":
				if r.Characteristic != nil {
					return fmt.Errorf("multiple values for field \"Characteristic\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Characteristic = v
			case "characteristicQuantity":
				if r.Characteristic != nil {
					return fmt.Errorf("multiple values for field \"Characteristic\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Characteristic = v
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = append(r.Value, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Categorization of the medication within a formulary or classification system.
type MedicationKnowledgeMedicineClassification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of category for the medication (for example, therapeutic classification, therapeutic sub-classification).
	Type CodeableConcept
	// Specific category assigned to the medication (e.g. anti-infective, anti-hypertensive, antibiotic, etc.).
	Classification []CodeableConcept
}
type jsonMedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type,omitempty"`
	Classification    []CodeableConcept `json:"classification,omitempty"`
}

func (r MedicationKnowledgeMedicineClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeMedicineClassification) marshalJSON() jsonMedicationKnowledgeMedicineClassification {
	m := jsonMedicationKnowledgeMedicineClassification{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Classification = r.Classification
	return m
}
func (r *MedicationKnowledgeMedicineClassification) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeMedicineClassification
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeMedicineClassification) unmarshalJSON(m jsonMedicationKnowledgeMedicineClassification) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Classification = m.Classification
	return nil
}
func (r MedicationKnowledgeMedicineClassification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Classification, xml.StartElement{Name: xml.Name{Local: "classification"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeMedicineClassification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "classification":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeMedicineClassification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information that only applies to packages (not products).
type MedicationKnowledgePackaging struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that defines the specific type of packaging that the medication can be found in (e.g. blister sleeve, tube, bottle).
	Type *CodeableConcept
	// The number of product units the package would contain if fully loaded.
	Quantity *Quantity
}
type jsonMedicationKnowledgePackaging struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}

func (r MedicationKnowledgePackaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgePackaging) marshalJSON() jsonMedicationKnowledgePackaging {
	m := jsonMedicationKnowledgePackaging{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Quantity = r.Quantity
	return m
}
func (r *MedicationKnowledgePackaging) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgePackaging
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgePackaging) unmarshalJSON(m jsonMedicationKnowledgePackaging) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Quantity = m.Quantity
	return nil
}
func (r MedicationKnowledgePackaging) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgePackaging) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgePackaging) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Specifies descriptive properties of the medicine, such as color, shape, imprints, etc.
type MedicationKnowledgeDrugCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code specifying which characteristic of the medicine is being described (for example, colour, shape, imprint).
	Type *CodeableConcept
	// Description of the characteristic.
	Value isMedicationKnowledgeDrugCharacteristicValue
}
type isMedicationKnowledgeDrugCharacteristicValue interface {
	isMedicationKnowledgeDrugCharacteristicValue()
}

func (r CodeableConcept) isMedicationKnowledgeDrugCharacteristicValue() {}
func (r String) isMedicationKnowledgeDrugCharacteristicValue()          {}
func (r Quantity) isMedicationKnowledgeDrugCharacteristicValue()        {}
func (r Base64Binary) isMedicationKnowledgeDrugCharacteristicValue()    {}

type jsonMedicationKnowledgeDrugCharacteristic struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	Type                              *CodeableConcept  `json:"type,omitempty"`
	ValueCodeableConcept              *CodeableConcept  `json:"valueCodeableConcept,omitempty"`
	ValueString                       *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement `json:"_valueString,omitempty"`
	ValueQuantity                     *Quantity         `json:"valueQuantity,omitempty"`
	ValueBase64Binary                 *Base64Binary     `json:"valueBase64Binary,omitempty"`
	ValueBase64BinaryPrimitiveElement *primitiveElement `json:"_valueBase64Binary,omitempty"`
}

func (r MedicationKnowledgeDrugCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeDrugCharacteristic) marshalJSON() jsonMedicationKnowledgeDrugCharacteristic {
	m := jsonMedicationKnowledgeDrugCharacteristic{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	switch v := r.Value.(type) {
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
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
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *MedicationKnowledgeDrugCharacteristic) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeDrugCharacteristic
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeDrugCharacteristic) unmarshalJSON(m jsonMedicationKnowledgeDrugCharacteristic) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
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
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueBase64Binary != nil || m.ValueBase64BinaryPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBase64Binary
		if m.ValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.ValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.ValueBase64BinaryPrimitiveElement.Extension
		}
		r.Value = v
	}
	return nil
}
func (r MedicationKnowledgeDrugCharacteristic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Value.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableConcept"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Base64Binary, *Base64Binary:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBase64Binary"}})
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
func (r *MedicationKnowledgeDrugCharacteristic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "valueCodeableConcept":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableConcept
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
			case "valueBase64Binary":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Base64Binary
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
func (r MedicationKnowledgeDrugCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Regulatory information about a medication.
type MedicationKnowledgeRegulatory struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The authority that is specifying the regulations.
	RegulatoryAuthority Reference
	// Specifies if changes are allowed when dispensing a medication from a regulatory perspective.
	Substitution []MedicationKnowledgeRegulatorySubstitution
	// Specifies the schedule of a medication in jurisdiction.
	Schedule []MedicationKnowledgeRegulatorySchedule
	// The maximum number of units of the medication that can be dispensed in a period.
	MaxDispense *MedicationKnowledgeRegulatoryMaxDispense
}
type jsonMedicationKnowledgeRegulatory struct {
	Id                  *string                                     `json:"id,omitempty"`
	Extension           []Extension                                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `json:"regulatoryAuthority,omitempty"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
	Schedule            []MedicationKnowledgeRegulatorySchedule     `json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `json:"maxDispense,omitempty"`
}

func (r MedicationKnowledgeRegulatory) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeRegulatory) marshalJSON() jsonMedicationKnowledgeRegulatory {
	m := jsonMedicationKnowledgeRegulatory{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.RegulatoryAuthority = r.RegulatoryAuthority
	m.Substitution = r.Substitution
	m.Schedule = r.Schedule
	m.MaxDispense = r.MaxDispense
	return m
}
func (r *MedicationKnowledgeRegulatory) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeRegulatory
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeRegulatory) unmarshalJSON(m jsonMedicationKnowledgeRegulatory) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.RegulatoryAuthority = m.RegulatoryAuthority
	r.Substitution = m.Substitution
	r.Schedule = m.Schedule
	r.MaxDispense = m.MaxDispense
	return nil
}
func (r MedicationKnowledgeRegulatory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.RegulatoryAuthority, xml.StartElement{Name: xml.Name{Local: "regulatoryAuthority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Substitution, xml.StartElement{Name: xml.Name{Local: "substitution"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Schedule, xml.StartElement{Name: xml.Name{Local: "schedule"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDispense, xml.StartElement{Name: xml.Name{Local: "maxDispense"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeRegulatory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "regulatoryAuthority":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RegulatoryAuthority = v
			case "substitution":
				var v MedicationKnowledgeRegulatorySubstitution
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substitution = append(r.Substitution, v)
			case "schedule":
				var v MedicationKnowledgeRegulatorySchedule
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Schedule = append(r.Schedule, v)
			case "maxDispense":
				var v MedicationKnowledgeRegulatoryMaxDispense
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDispense = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeRegulatory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Specifies if changes are allowed when dispensing a medication from a regulatory perspective.
type MedicationKnowledgeRegulatorySubstitution struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specifies the type of substitution allowed.
	Type CodeableConcept
	// Specifies if regulation allows for changes in the medication when dispensing.
	Allowed Boolean
}
type jsonMedicationKnowledgeRegulatorySubstitution struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Type                    CodeableConcept   `json:"type,omitempty"`
	Allowed                 Boolean           `json:"allowed,omitempty"`
	AllowedPrimitiveElement *primitiveElement `json:"_allowed,omitempty"`
}

func (r MedicationKnowledgeRegulatorySubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeRegulatorySubstitution) marshalJSON() jsonMedicationKnowledgeRegulatorySubstitution {
	m := jsonMedicationKnowledgeRegulatorySubstitution{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Allowed.Value != nil {
		m.Allowed = r.Allowed
	}
	if r.Allowed.Id != nil || r.Allowed.Extension != nil {
		m.AllowedPrimitiveElement = &primitiveElement{Id: r.Allowed.Id, Extension: r.Allowed.Extension}
	}
	return m
}
func (r *MedicationKnowledgeRegulatorySubstitution) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeRegulatorySubstitution
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeRegulatorySubstitution) unmarshalJSON(m jsonMedicationKnowledgeRegulatorySubstitution) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Allowed = m.Allowed
	if m.AllowedPrimitiveElement != nil {
		r.Allowed.Id = m.AllowedPrimitiveElement.Id
		r.Allowed.Extension = m.AllowedPrimitiveElement.Extension
	}
	return nil
}
func (r MedicationKnowledgeRegulatorySubstitution) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Allowed, xml.StartElement{Name: xml.Name{Local: "allowed"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeRegulatorySubstitution) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "allowed":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Allowed = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeRegulatorySubstitution) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Specifies the schedule of a medication in jurisdiction.
type MedicationKnowledgeRegulatorySchedule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specifies the specific drug schedule.
	Schedule CodeableConcept
}
type jsonMedicationKnowledgeRegulatorySchedule struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Schedule          CodeableConcept `json:"schedule,omitempty"`
}

func (r MedicationKnowledgeRegulatorySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeRegulatorySchedule) marshalJSON() jsonMedicationKnowledgeRegulatorySchedule {
	m := jsonMedicationKnowledgeRegulatorySchedule{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Schedule = r.Schedule
	return m
}
func (r *MedicationKnowledgeRegulatorySchedule) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeRegulatorySchedule
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeRegulatorySchedule) unmarshalJSON(m jsonMedicationKnowledgeRegulatorySchedule) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Schedule = m.Schedule
	return nil
}
func (r MedicationKnowledgeRegulatorySchedule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Schedule, xml.StartElement{Name: xml.Name{Local: "schedule"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeRegulatorySchedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "schedule":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Schedule = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeRegulatorySchedule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The maximum number of units of the medication that can be dispensed in a period.
type MedicationKnowledgeRegulatoryMaxDispense struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The maximum number of units of the medication that can be dispensed.
	Quantity Quantity
	// The period that applies to the maximum number of units.
	Period *Duration
}
type jsonMedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          Quantity    `json:"quantity,omitempty"`
	Period            *Duration   `json:"period,omitempty"`
}

func (r MedicationKnowledgeRegulatoryMaxDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeRegulatoryMaxDispense) marshalJSON() jsonMedicationKnowledgeRegulatoryMaxDispense {
	m := jsonMedicationKnowledgeRegulatoryMaxDispense{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Quantity = r.Quantity
	m.Period = r.Period
	return m
}
func (r *MedicationKnowledgeRegulatoryMaxDispense) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeRegulatoryMaxDispense
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeRegulatoryMaxDispense) unmarshalJSON(m jsonMedicationKnowledgeRegulatoryMaxDispense) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Quantity = m.Quantity
	r.Period = m.Period
	return nil
}
func (r MedicationKnowledgeRegulatoryMaxDispense) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeRegulatoryMaxDispense) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = v
			case "period":
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeRegulatoryMaxDispense) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The time course of drug absorption, distribution, metabolism and excretion of a medication from the body.
type MedicationKnowledgeKinetics struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The drug concentration measured at certain discrete points in time.
	AreaUnderCurve []Quantity
	// The median lethal dose of a drug.
	LethalDose50 []Quantity
	// The time required for any specified property (e.g., the concentration of a substance in the body) to decrease by half.
	HalfLifePeriod *Duration
}
type jsonMedicationKnowledgeKinetics struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	AreaUnderCurve    []Quantity  `json:"areaUnderCurve,omitempty"`
	LethalDose50      []Quantity  `json:"lethalDose50,omitempty"`
	HalfLifePeriod    *Duration   `json:"halfLifePeriod,omitempty"`
}

func (r MedicationKnowledgeKinetics) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationKnowledgeKinetics) marshalJSON() jsonMedicationKnowledgeKinetics {
	m := jsonMedicationKnowledgeKinetics{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.AreaUnderCurve = r.AreaUnderCurve
	m.LethalDose50 = r.LethalDose50
	m.HalfLifePeriod = r.HalfLifePeriod
	return m
}
func (r *MedicationKnowledgeKinetics) UnmarshalJSON(b []byte) error {
	var m jsonMedicationKnowledgeKinetics
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationKnowledgeKinetics) unmarshalJSON(m jsonMedicationKnowledgeKinetics) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.AreaUnderCurve = m.AreaUnderCurve
	r.LethalDose50 = m.LethalDose50
	r.HalfLifePeriod = m.HalfLifePeriod
	return nil
}
func (r MedicationKnowledgeKinetics) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.AreaUnderCurve, xml.StartElement{Name: xml.Name{Local: "areaUnderCurve"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LethalDose50, xml.StartElement{Name: xml.Name{Local: "lethalDose50"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HalfLifePeriod, xml.StartElement{Name: xml.Name{Local: "halfLifePeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicationKnowledgeKinetics) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "areaUnderCurve":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AreaUnderCurve = append(r.AreaUnderCurve, v)
			case "lethalDose50":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LethalDose50 = append(r.LethalDose50, v)
			case "halfLifePeriod":
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HalfLifePeriod = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicationKnowledgeKinetics) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
