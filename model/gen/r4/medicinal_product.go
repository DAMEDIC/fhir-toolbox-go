package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g. regulatory use).
type MedicinalProduct struct {
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
	// Business identifier for this product. Could be an MPID.
	Identifier []Identifier
	// Regulatory type, e.g. Investigational or Authorized.
	Type *CodeableConcept
	// If this medicine applies to human or veterinary uses.
	Domain *Coding
	// The dose form for a single part product, or combined form of a multiple part product.
	CombinedPharmaceuticalDoseForm *CodeableConcept
	// The legal status of supply of the medicinal product as classified by the regulator.
	LegalStatusOfSupply *CodeableConcept
	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons.
	AdditionalMonitoringIndicator *CodeableConcept
	// Whether the Medicinal Product is subject to special measures for regulatory reasons.
	SpecialMeasures []String
	// If authorised for use in children.
	PaediatricUseIndicator *CodeableConcept
	// Allows the product to be classified by various systems.
	ProductClassification []CodeableConcept
	// Marketing status of the medicinal product, in contrast to marketing authorizaton.
	MarketingStatus []MarketingStatus
	// Pharmaceutical aspects of product.
	PharmaceuticalProduct []Reference
	// Package representation for the product.
	PackagedMedicinalProduct []Reference
	// Supporting documentation, typically for regulatory submission.
	AttachedDocument []Reference
	// A master file for to the medicinal product (e.g. Pharmacovigilance System Master File).
	MasterFile []Reference
	// A product specific contact, person (in a role), or an organization.
	Contact []Reference
	// Clinical trials or studies that this product is involved in.
	ClinicalTrial []Reference
	// The product's name, including full name and possibly coded parts.
	Name []MedicinalProductName
	// Reference to another product, e.g. for linking authorised to investigational product.
	CrossReference []Identifier
	// An operation applied to the product, for manufacturing or adminsitrative purpose.
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperation
	// Indicates if the medicinal product has an orphan designation for the treatment of a rare disease.
	SpecialDesignation []MedicinalProductSpecialDesignation
}

func (r MedicinalProduct) ResourceType() string {
	return "MedicinalProduct"
}
func (r MedicinalProduct) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProduct struct {
	ResourceType                    string                                           `json:"resourceType"`
	Id                              *Id                                              `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement                                `json:"_id,omitempty"`
	Meta                            *Meta                                            `json:"meta,omitempty"`
	ImplicitRules                   *Uri                                             `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement                                `json:"_implicitRules,omitempty"`
	Language                        *Code                                            `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement                                `json:"_language,omitempty"`
	Text                            *Narrative                                       `json:"text,omitempty"`
	Contained                       []ContainedResource                              `json:"contained,omitempty"`
	Extension                       []Extension                                      `json:"extension,omitempty"`
	ModifierExtension               []Extension                                      `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier                                     `json:"identifier,omitempty"`
	Type                            *CodeableConcept                                 `json:"type,omitempty"`
	Domain                          *Coding                                          `json:"domain,omitempty"`
	CombinedPharmaceuticalDoseForm  *CodeableConcept                                 `json:"combinedPharmaceuticalDoseForm,omitempty"`
	LegalStatusOfSupply             *CodeableConcept                                 `json:"legalStatusOfSupply,omitempty"`
	AdditionalMonitoringIndicator   *CodeableConcept                                 `json:"additionalMonitoringIndicator,omitempty"`
	SpecialMeasures                 []String                                         `json:"specialMeasures,omitempty"`
	SpecialMeasuresPrimitiveElement []*primitiveElement                              `json:"_specialMeasures,omitempty"`
	PaediatricUseIndicator          *CodeableConcept                                 `json:"paediatricUseIndicator,omitempty"`
	ProductClassification           []CodeableConcept                                `json:"productClassification,omitempty"`
	MarketingStatus                 []MarketingStatus                                `json:"marketingStatus,omitempty"`
	PharmaceuticalProduct           []Reference                                      `json:"pharmaceuticalProduct,omitempty"`
	PackagedMedicinalProduct        []Reference                                      `json:"packagedMedicinalProduct,omitempty"`
	AttachedDocument                []Reference                                      `json:"attachedDocument,omitempty"`
	MasterFile                      []Reference                                      `json:"masterFile,omitempty"`
	Contact                         []Reference                                      `json:"contact,omitempty"`
	ClinicalTrial                   []Reference                                      `json:"clinicalTrial,omitempty"`
	Name                            []MedicinalProductName                           `json:"name,omitempty"`
	CrossReference                  []Identifier                                     `json:"crossReference,omitempty"`
	ManufacturingBusinessOperation  []MedicinalProductManufacturingBusinessOperation `json:"manufacturingBusinessOperation,omitempty"`
	SpecialDesignation              []MedicinalProductSpecialDesignation             `json:"specialDesignation,omitempty"`
}

func (r MedicinalProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProduct) marshalJSON() jsonMedicinalProduct {
	m := jsonMedicinalProduct{}
	m.ResourceType = "MedicinalProduct"
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
	m.Type = r.Type
	m.Domain = r.Domain
	m.CombinedPharmaceuticalDoseForm = r.CombinedPharmaceuticalDoseForm
	m.LegalStatusOfSupply = r.LegalStatusOfSupply
	m.AdditionalMonitoringIndicator = r.AdditionalMonitoringIndicator
	anySpecialMeasuresValue := false
	for _, e := range r.SpecialMeasures {
		if e.Value != nil {
			anySpecialMeasuresValue = true
			break
		}
	}
	if anySpecialMeasuresValue {
		m.SpecialMeasures = r.SpecialMeasures
	}
	anySpecialMeasuresIdOrExtension := false
	for _, e := range r.SpecialMeasures {
		if e.Id != nil || e.Extension != nil {
			anySpecialMeasuresIdOrExtension = true
			break
		}
	}
	if anySpecialMeasuresIdOrExtension {
		m.SpecialMeasuresPrimitiveElement = make([]*primitiveElement, 0, len(r.SpecialMeasures))
		for _, e := range r.SpecialMeasures {
			if e.Id != nil || e.Extension != nil {
				m.SpecialMeasuresPrimitiveElement = append(m.SpecialMeasuresPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SpecialMeasuresPrimitiveElement = append(m.SpecialMeasuresPrimitiveElement, nil)
			}
		}
	}
	m.PaediatricUseIndicator = r.PaediatricUseIndicator
	m.ProductClassification = r.ProductClassification
	m.MarketingStatus = r.MarketingStatus
	m.PharmaceuticalProduct = r.PharmaceuticalProduct
	m.PackagedMedicinalProduct = r.PackagedMedicinalProduct
	m.AttachedDocument = r.AttachedDocument
	m.MasterFile = r.MasterFile
	m.Contact = r.Contact
	m.ClinicalTrial = r.ClinicalTrial
	m.Name = r.Name
	m.CrossReference = r.CrossReference
	m.ManufacturingBusinessOperation = r.ManufacturingBusinessOperation
	m.SpecialDesignation = r.SpecialDesignation
	return m
}
func (r *MedicinalProduct) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProduct
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProduct) unmarshalJSON(m jsonMedicinalProduct) error {
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
	r.Type = m.Type
	r.Domain = m.Domain
	r.CombinedPharmaceuticalDoseForm = m.CombinedPharmaceuticalDoseForm
	r.LegalStatusOfSupply = m.LegalStatusOfSupply
	r.AdditionalMonitoringIndicator = m.AdditionalMonitoringIndicator
	r.SpecialMeasures = m.SpecialMeasures
	for i, e := range m.SpecialMeasuresPrimitiveElement {
		if len(r.SpecialMeasures) <= i {
			r.SpecialMeasures = append(r.SpecialMeasures, String{})
		}
		if e != nil {
			r.SpecialMeasures[i].Id = e.Id
			r.SpecialMeasures[i].Extension = e.Extension
		}
	}
	r.PaediatricUseIndicator = m.PaediatricUseIndicator
	r.ProductClassification = m.ProductClassification
	r.MarketingStatus = m.MarketingStatus
	r.PharmaceuticalProduct = m.PharmaceuticalProduct
	r.PackagedMedicinalProduct = m.PackagedMedicinalProduct
	r.AttachedDocument = m.AttachedDocument
	r.MasterFile = m.MasterFile
	r.Contact = m.Contact
	r.ClinicalTrial = m.ClinicalTrial
	r.Name = m.Name
	r.CrossReference = m.CrossReference
	r.ManufacturingBusinessOperation = m.ManufacturingBusinessOperation
	r.SpecialDesignation = m.SpecialDesignation
	return nil
}
func (r MedicinalProduct) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicinalProduct"
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Domain, xml.StartElement{Name: xml.Name{Local: "domain"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CombinedPharmaceuticalDoseForm, xml.StartElement{Name: xml.Name{Local: "combinedPharmaceuticalDoseForm"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LegalStatusOfSupply, xml.StartElement{Name: xml.Name{Local: "legalStatusOfSupply"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdditionalMonitoringIndicator, xml.StartElement{Name: xml.Name{Local: "additionalMonitoringIndicator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SpecialMeasures, xml.StartElement{Name: xml.Name{Local: "specialMeasures"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaediatricUseIndicator, xml.StartElement{Name: xml.Name{Local: "paediatricUseIndicator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductClassification, xml.StartElement{Name: xml.Name{Local: "productClassification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MarketingStatus, xml.StartElement{Name: xml.Name{Local: "marketingStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PharmaceuticalProduct, xml.StartElement{Name: xml.Name{Local: "pharmaceuticalProduct"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PackagedMedicinalProduct, xml.StartElement{Name: xml.Name{Local: "packagedMedicinalProduct"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AttachedDocument, xml.StartElement{Name: xml.Name{Local: "attachedDocument"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MasterFile, xml.StartElement{Name: xml.Name{Local: "masterFile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClinicalTrial, xml.StartElement{Name: xml.Name{Local: "clinicalTrial"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CrossReference, xml.StartElement{Name: xml.Name{Local: "crossReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ManufacturingBusinessOperation, xml.StartElement{Name: xml.Name{Local: "manufacturingBusinessOperation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SpecialDesignation, xml.StartElement{Name: xml.Name{Local: "specialDesignation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProduct) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "domain":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Domain = &v
			case "combinedPharmaceuticalDoseForm":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CombinedPharmaceuticalDoseForm = &v
			case "legalStatusOfSupply":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegalStatusOfSupply = &v
			case "additionalMonitoringIndicator":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalMonitoringIndicator = &v
			case "specialMeasures":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SpecialMeasures = append(r.SpecialMeasures, v)
			case "paediatricUseIndicator":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaediatricUseIndicator = &v
			case "productClassification":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductClassification = append(r.ProductClassification, v)
			case "marketingStatus":
				var v MarketingStatus
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MarketingStatus = append(r.MarketingStatus, v)
			case "pharmaceuticalProduct":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PharmaceuticalProduct = append(r.PharmaceuticalProduct, v)
			case "packagedMedicinalProduct":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PackagedMedicinalProduct = append(r.PackagedMedicinalProduct, v)
			case "attachedDocument":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AttachedDocument = append(r.AttachedDocument, v)
			case "masterFile":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MasterFile = append(r.MasterFile, v)
			case "contact":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "clinicalTrial":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClinicalTrial = append(r.ClinicalTrial, v)
			case "name":
				var v MedicinalProductName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = append(r.Name, v)
			case "crossReference":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CrossReference = append(r.CrossReference, v)
			case "manufacturingBusinessOperation":
				var v MedicinalProductManufacturingBusinessOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ManufacturingBusinessOperation = append(r.ManufacturingBusinessOperation, v)
			case "specialDesignation":
				var v MedicinalProductSpecialDesignation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SpecialDesignation = append(r.SpecialDesignation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProduct) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The product's name, including full name and possibly coded parts.
type MedicinalProductName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The full product name.
	ProductName String
	// Coding words or phrases of the name.
	NamePart []MedicinalProductNameNamePart
	// Country where the name applies.
	CountryLanguage []MedicinalProductNameCountryLanguage
}
type jsonMedicinalProductName struct {
	Id                          *string                               `json:"id,omitempty"`
	Extension                   []Extension                           `json:"extension,omitempty"`
	ModifierExtension           []Extension                           `json:"modifierExtension,omitempty"`
	ProductName                 String                                `json:"productName,omitempty"`
	ProductNamePrimitiveElement *primitiveElement                     `json:"_productName,omitempty"`
	NamePart                    []MedicinalProductNameNamePart        `json:"namePart,omitempty"`
	CountryLanguage             []MedicinalProductNameCountryLanguage `json:"countryLanguage,omitempty"`
}

func (r MedicinalProductName) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductName) marshalJSON() jsonMedicinalProductName {
	m := jsonMedicinalProductName{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ProductName.Value != nil {
		m.ProductName = r.ProductName
	}
	if r.ProductName.Id != nil || r.ProductName.Extension != nil {
		m.ProductNamePrimitiveElement = &primitiveElement{Id: r.ProductName.Id, Extension: r.ProductName.Extension}
	}
	m.NamePart = r.NamePart
	m.CountryLanguage = r.CountryLanguage
	return m
}
func (r *MedicinalProductName) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductName
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductName) unmarshalJSON(m jsonMedicinalProductName) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ProductName = m.ProductName
	if m.ProductNamePrimitiveElement != nil {
		r.ProductName.Id = m.ProductNamePrimitiveElement.Id
		r.ProductName.Extension = m.ProductNamePrimitiveElement.Extension
	}
	r.NamePart = m.NamePart
	r.CountryLanguage = m.CountryLanguage
	return nil
}
func (r MedicinalProductName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ProductName, xml.StartElement{Name: xml.Name{Local: "productName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NamePart, xml.StartElement{Name: xml.Name{Local: "namePart"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CountryLanguage, xml.StartElement{Name: xml.Name{Local: "countryLanguage"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "productName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductName = v
			case "namePart":
				var v MedicinalProductNameNamePart
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NamePart = append(r.NamePart, v)
			case "countryLanguage":
				var v MedicinalProductNameCountryLanguage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CountryLanguage = append(r.CountryLanguage, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductName) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Coding words or phrases of the name.
type MedicinalProductNameNamePart struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A fragment of a product name.
	Part String
	// Idenifying type for this part of the name (e.g. strength part).
	Type Coding
}
type jsonMedicinalProductNameNamePart struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Part                 String            `json:"part,omitempty"`
	PartPrimitiveElement *primitiveElement `json:"_part,omitempty"`
	Type                 Coding            `json:"type,omitempty"`
}

func (r MedicinalProductNameNamePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductNameNamePart) marshalJSON() jsonMedicinalProductNameNamePart {
	m := jsonMedicinalProductNameNamePart{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Part.Value != nil {
		m.Part = r.Part
	}
	if r.Part.Id != nil || r.Part.Extension != nil {
		m.PartPrimitiveElement = &primitiveElement{Id: r.Part.Id, Extension: r.Part.Extension}
	}
	m.Type = r.Type
	return m
}
func (r *MedicinalProductNameNamePart) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductNameNamePart
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductNameNamePart) unmarshalJSON(m jsonMedicinalProductNameNamePart) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Part = m.Part
	if m.PartPrimitiveElement != nil {
		r.Part.Id = m.PartPrimitiveElement.Id
		r.Part.Extension = m.PartPrimitiveElement.Extension
	}
	r.Type = m.Type
	return nil
}
func (r MedicinalProductNameNamePart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Part, xml.StartElement{Name: xml.Name{Local: "part"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductNameNamePart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "part":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Part = v
			case "type":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductNameNamePart) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Country where the name applies.
type MedicinalProductNameCountryLanguage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Country code for where this name applies.
	Country CodeableConcept
	// Jurisdiction code for where this name applies.
	Jurisdiction *CodeableConcept
	// Language code for this name.
	Language CodeableConcept
}
type jsonMedicinalProductNameCountryLanguage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country,omitempty"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Language          CodeableConcept  `json:"language,omitempty"`
}

func (r MedicinalProductNameCountryLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductNameCountryLanguage) marshalJSON() jsonMedicinalProductNameCountryLanguage {
	m := jsonMedicinalProductNameCountryLanguage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Country = r.Country
	m.Jurisdiction = r.Jurisdiction
	m.Language = r.Language
	return m
}
func (r *MedicinalProductNameCountryLanguage) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductNameCountryLanguage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductNameCountryLanguage) unmarshalJSON(m jsonMedicinalProductNameCountryLanguage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Country = m.Country
	r.Jurisdiction = m.Jurisdiction
	r.Language = m.Language
	return nil
}
func (r MedicinalProductNameCountryLanguage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductNameCountryLanguage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "country":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = v
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = &v
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductNameCountryLanguage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An operation applied to the product, for manufacturing or adminsitrative purpose.
type MedicinalProductManufacturingBusinessOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of manufacturing operation.
	OperationType *CodeableConcept
	// Regulatory authorization reference number.
	AuthorisationReferenceNumber *Identifier
	// Regulatory authorization date.
	EffectiveDate *DateTime
	// To indicate if this proces is commercially confidential.
	ConfidentialityIndicator *CodeableConcept
	// The manufacturer or establishment associated with the process.
	Manufacturer []Reference
	// A regulator which oversees the operation.
	Regulator *Reference
}
type jsonMedicinalProductManufacturingBusinessOperation struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	OperationType                 *CodeableConcept  `json:"operationType,omitempty"`
	AuthorisationReferenceNumber  *Identifier       `json:"authorisationReferenceNumber,omitempty"`
	EffectiveDate                 *DateTime         `json:"effectiveDate,omitempty"`
	EffectiveDatePrimitiveElement *primitiveElement `json:"_effectiveDate,omitempty"`
	ConfidentialityIndicator      *CodeableConcept  `json:"confidentialityIndicator,omitempty"`
	Manufacturer                  []Reference       `json:"manufacturer,omitempty"`
	Regulator                     *Reference        `json:"regulator,omitempty"`
}

func (r MedicinalProductManufacturingBusinessOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductManufacturingBusinessOperation) marshalJSON() jsonMedicinalProductManufacturingBusinessOperation {
	m := jsonMedicinalProductManufacturingBusinessOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.OperationType = r.OperationType
	m.AuthorisationReferenceNumber = r.AuthorisationReferenceNumber
	if r.EffectiveDate != nil && r.EffectiveDate.Value != nil {
		m.EffectiveDate = r.EffectiveDate
	}
	if r.EffectiveDate != nil && (r.EffectiveDate.Id != nil || r.EffectiveDate.Extension != nil) {
		m.EffectiveDatePrimitiveElement = &primitiveElement{Id: r.EffectiveDate.Id, Extension: r.EffectiveDate.Extension}
	}
	m.ConfidentialityIndicator = r.ConfidentialityIndicator
	m.Manufacturer = r.Manufacturer
	m.Regulator = r.Regulator
	return m
}
func (r *MedicinalProductManufacturingBusinessOperation) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductManufacturingBusinessOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductManufacturingBusinessOperation) unmarshalJSON(m jsonMedicinalProductManufacturingBusinessOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.OperationType = m.OperationType
	r.AuthorisationReferenceNumber = m.AuthorisationReferenceNumber
	r.EffectiveDate = m.EffectiveDate
	if m.EffectiveDatePrimitiveElement != nil {
		if r.EffectiveDate == nil {
			r.EffectiveDate = &DateTime{}
		}
		r.EffectiveDate.Id = m.EffectiveDatePrimitiveElement.Id
		r.EffectiveDate.Extension = m.EffectiveDatePrimitiveElement.Extension
	}
	r.ConfidentialityIndicator = m.ConfidentialityIndicator
	r.Manufacturer = m.Manufacturer
	r.Regulator = m.Regulator
	return nil
}
func (r MedicinalProductManufacturingBusinessOperation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.OperationType, xml.StartElement{Name: xml.Name{Local: "operationType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthorisationReferenceNumber, xml.StartElement{Name: xml.Name{Local: "authorisationReferenceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EffectiveDate, xml.StartElement{Name: xml.Name{Local: "effectiveDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConfidentialityIndicator, xml.StartElement{Name: xml.Name{Local: "confidentialityIndicator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manufacturer, xml.StartElement{Name: xml.Name{Local: "manufacturer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Regulator, xml.StartElement{Name: xml.Name{Local: "regulator"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductManufacturingBusinessOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "operationType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OperationType = &v
			case "authorisationReferenceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorisationReferenceNumber = &v
			case "effectiveDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EffectiveDate = &v
			case "confidentialityIndicator":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConfidentialityIndicator = &v
			case "manufacturer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manufacturer = append(r.Manufacturer, v)
			case "regulator":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Regulator = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductManufacturingBusinessOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Indicates if the medicinal product has an orphan designation for the treatment of a rare disease.
type MedicinalProductSpecialDesignation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifier for the designation, or procedure number.
	Identifier []Identifier
	// The type of special designation, e.g. orphan drug, minor use.
	Type *CodeableConcept
	// The intended use of the product, e.g. prevention, treatment.
	IntendedUse *CodeableConcept
	// Condition for which the medicinal use applies.
	Indication isMedicinalProductSpecialDesignationIndication
	// For example granted, pending, expired or withdrawn.
	Status *CodeableConcept
	// Date when the designation was granted.
	Date *DateTime
	// Animal species for which this applies.
	Species *CodeableConcept
}
type isMedicinalProductSpecialDesignationIndication interface {
	isMedicinalProductSpecialDesignationIndication()
}

func (r CodeableConcept) isMedicinalProductSpecialDesignationIndication() {}
func (r Reference) isMedicinalProductSpecialDesignationIndication()       {}

type jsonMedicinalProductSpecialDesignation struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `json:"identifier,omitempty"`
	Type                      *CodeableConcept  `json:"type,omitempty"`
	IntendedUse               *CodeableConcept  `json:"intendedUse,omitempty"`
	IndicationCodeableConcept *CodeableConcept  `json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference        `json:"indicationReference,omitempty"`
	Status                    *CodeableConcept  `json:"status,omitempty"`
	Date                      *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement      *primitiveElement `json:"_date,omitempty"`
	Species                   *CodeableConcept  `json:"species,omitempty"`
}

func (r MedicinalProductSpecialDesignation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductSpecialDesignation) marshalJSON() jsonMedicinalProductSpecialDesignation {
	m := jsonMedicinalProductSpecialDesignation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Type = r.Type
	m.IntendedUse = r.IntendedUse
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
	m.Status = r.Status
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Species = r.Species
	return m
}
func (r *MedicinalProductSpecialDesignation) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductSpecialDesignation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductSpecialDesignation) unmarshalJSON(m jsonMedicinalProductSpecialDesignation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Type = m.Type
	r.IntendedUse = m.IntendedUse
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
	r.Status = m.Status
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Species = m.Species
	return nil
}
func (r MedicinalProductSpecialDesignation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntendedUse, xml.StartElement{Name: xml.Name{Local: "intendedUse"}})
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
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Species, xml.StartElement{Name: xml.Name{Local: "species"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductSpecialDesignation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "intendedUse":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntendedUse = &v
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
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "species":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Species = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductSpecialDesignation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
