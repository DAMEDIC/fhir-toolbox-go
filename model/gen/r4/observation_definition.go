package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
//
// In a catalog of health-related services that use or produce observations and measurements, this resource describes the expected characteristics of these observation / measurements.
type ObservationDefinition struct {
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
	// A code that classifies the general type of observation.
	Category []CodeableConcept
	// Describes what will be observed. Sometimes this is called the observation "name".
	Code CodeableConcept
	// A unique identifier assigned to this ObservationDefinition artifact.
	Identifier []Identifier
	// The data types allowed for the value element of the instance observations conforming to this ObservationDefinition.
	PermittedDataType []Code
	// Multiple results allowed for observations conforming to this ObservationDefinition.
	MultipleResultsAllowed *Boolean
	// The method or technique used to perform the observation.
	Method *CodeableConcept
	// The preferred name to be used when reporting the results of observations conforming to this ObservationDefinition.
	PreferredReportName *String
	// Characteristics for quantitative results of this observation.
	QuantitativeDetails *ObservationDefinitionQuantitativeDetails
	// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
	QualifiedInterval []ObservationDefinitionQualifiedInterval
	// The set of valid coded results for the observations  conforming to this ObservationDefinition.
	ValidCodedValueSet *Reference
	// The set of normal coded results for the observations conforming to this ObservationDefinition.
	NormalCodedValueSet *Reference
	// The set of abnormal coded results for the observation conforming to this ObservationDefinition.
	AbnormalCodedValueSet *Reference
	// The set of critical coded results for the observation conforming to this ObservationDefinition.
	CriticalCodedValueSet *Reference
}

func (r ObservationDefinition) ResourceType() string {
	return "ObservationDefinition"
}
func (r ObservationDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonObservationDefinition struct {
	ResourceType                           string                                    `json:"resourceType"`
	Id                                     *Id                                       `json:"id,omitempty"`
	IdPrimitiveElement                     *primitiveElement                         `json:"_id,omitempty"`
	Meta                                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules                          *Uri                                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement          *primitiveElement                         `json:"_implicitRules,omitempty"`
	Language                               *Code                                     `json:"language,omitempty"`
	LanguagePrimitiveElement               *primitiveElement                         `json:"_language,omitempty"`
	Text                                   *Narrative                                `json:"text,omitempty"`
	Contained                              []ContainedResource                       `json:"contained,omitempty"`
	Extension                              []Extension                               `json:"extension,omitempty"`
	ModifierExtension                      []Extension                               `json:"modifierExtension,omitempty"`
	Category                               []CodeableConcept                         `json:"category,omitempty"`
	Code                                   CodeableConcept                           `json:"code,omitempty"`
	Identifier                             []Identifier                              `json:"identifier,omitempty"`
	PermittedDataType                      []Code                                    `json:"permittedDataType,omitempty"`
	PermittedDataTypePrimitiveElement      []*primitiveElement                       `json:"_permittedDataType,omitempty"`
	MultipleResultsAllowed                 *Boolean                                  `json:"multipleResultsAllowed,omitempty"`
	MultipleResultsAllowedPrimitiveElement *primitiveElement                         `json:"_multipleResultsAllowed,omitempty"`
	Method                                 *CodeableConcept                          `json:"method,omitempty"`
	PreferredReportName                    *String                                   `json:"preferredReportName,omitempty"`
	PreferredReportNamePrimitiveElement    *primitiveElement                         `json:"_preferredReportName,omitempty"`
	QuantitativeDetails                    *ObservationDefinitionQuantitativeDetails `json:"quantitativeDetails,omitempty"`
	QualifiedInterval                      []ObservationDefinitionQualifiedInterval  `json:"qualifiedInterval,omitempty"`
	ValidCodedValueSet                     *Reference                                `json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet                    *Reference                                `json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet                  *Reference                                `json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet                  *Reference                                `json:"criticalCodedValueSet,omitempty"`
}

func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ObservationDefinition) marshalJSON() jsonObservationDefinition {
	m := jsonObservationDefinition{}
	m.ResourceType = "ObservationDefinition"
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
	m.Category = r.Category
	m.Code = r.Code
	m.Identifier = r.Identifier
	anyPermittedDataTypeValue := false
	for _, e := range r.PermittedDataType {
		if e.Value != nil {
			anyPermittedDataTypeValue = true
			break
		}
	}
	if anyPermittedDataTypeValue {
		m.PermittedDataType = r.PermittedDataType
	}
	anyPermittedDataTypeIdOrExtension := false
	for _, e := range r.PermittedDataType {
		if e.Id != nil || e.Extension != nil {
			anyPermittedDataTypeIdOrExtension = true
			break
		}
	}
	if anyPermittedDataTypeIdOrExtension {
		m.PermittedDataTypePrimitiveElement = make([]*primitiveElement, 0, len(r.PermittedDataType))
		for _, e := range r.PermittedDataType {
			if e.Id != nil || e.Extension != nil {
				m.PermittedDataTypePrimitiveElement = append(m.PermittedDataTypePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PermittedDataTypePrimitiveElement = append(m.PermittedDataTypePrimitiveElement, nil)
			}
		}
	}
	if r.MultipleResultsAllowed != nil && r.MultipleResultsAllowed.Value != nil {
		m.MultipleResultsAllowed = r.MultipleResultsAllowed
	}
	if r.MultipleResultsAllowed != nil && (r.MultipleResultsAllowed.Id != nil || r.MultipleResultsAllowed.Extension != nil) {
		m.MultipleResultsAllowedPrimitiveElement = &primitiveElement{Id: r.MultipleResultsAllowed.Id, Extension: r.MultipleResultsAllowed.Extension}
	}
	m.Method = r.Method
	if r.PreferredReportName != nil && r.PreferredReportName.Value != nil {
		m.PreferredReportName = r.PreferredReportName
	}
	if r.PreferredReportName != nil && (r.PreferredReportName.Id != nil || r.PreferredReportName.Extension != nil) {
		m.PreferredReportNamePrimitiveElement = &primitiveElement{Id: r.PreferredReportName.Id, Extension: r.PreferredReportName.Extension}
	}
	m.QuantitativeDetails = r.QuantitativeDetails
	m.QualifiedInterval = r.QualifiedInterval
	m.ValidCodedValueSet = r.ValidCodedValueSet
	m.NormalCodedValueSet = r.NormalCodedValueSet
	m.AbnormalCodedValueSet = r.AbnormalCodedValueSet
	m.CriticalCodedValueSet = r.CriticalCodedValueSet
	return m
}
func (r *ObservationDefinition) UnmarshalJSON(b []byte) error {
	var m jsonObservationDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ObservationDefinition) unmarshalJSON(m jsonObservationDefinition) error {
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
	r.Category = m.Category
	r.Code = m.Code
	r.Identifier = m.Identifier
	r.PermittedDataType = m.PermittedDataType
	for i, e := range m.PermittedDataTypePrimitiveElement {
		if len(r.PermittedDataType) <= i {
			r.PermittedDataType = append(r.PermittedDataType, Code{})
		}
		if e != nil {
			r.PermittedDataType[i].Id = e.Id
			r.PermittedDataType[i].Extension = e.Extension
		}
	}
	r.MultipleResultsAllowed = m.MultipleResultsAllowed
	if m.MultipleResultsAllowedPrimitiveElement != nil {
		if r.MultipleResultsAllowed == nil {
			r.MultipleResultsAllowed = &Boolean{}
		}
		r.MultipleResultsAllowed.Id = m.MultipleResultsAllowedPrimitiveElement.Id
		r.MultipleResultsAllowed.Extension = m.MultipleResultsAllowedPrimitiveElement.Extension
	}
	r.Method = m.Method
	r.PreferredReportName = m.PreferredReportName
	if m.PreferredReportNamePrimitiveElement != nil {
		if r.PreferredReportName == nil {
			r.PreferredReportName = &String{}
		}
		r.PreferredReportName.Id = m.PreferredReportNamePrimitiveElement.Id
		r.PreferredReportName.Extension = m.PreferredReportNamePrimitiveElement.Extension
	}
	r.QuantitativeDetails = m.QuantitativeDetails
	r.QualifiedInterval = m.QualifiedInterval
	r.ValidCodedValueSet = m.ValidCodedValueSet
	r.NormalCodedValueSet = m.NormalCodedValueSet
	r.AbnormalCodedValueSet = m.AbnormalCodedValueSet
	r.CriticalCodedValueSet = m.CriticalCodedValueSet
	return nil
}
func (r ObservationDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ObservationDefinition"
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PermittedDataType, xml.StartElement{Name: xml.Name{Local: "permittedDataType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MultipleResultsAllowed, xml.StartElement{Name: xml.Name{Local: "multipleResultsAllowed"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreferredReportName, xml.StartElement{Name: xml.Name{Local: "preferredReportName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QuantitativeDetails, xml.StartElement{Name: xml.Name{Local: "quantitativeDetails"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QualifiedInterval, xml.StartElement{Name: xml.Name{Local: "qualifiedInterval"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidCodedValueSet, xml.StartElement{Name: xml.Name{Local: "validCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NormalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "normalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AbnormalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "abnormalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CriticalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "criticalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ObservationDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "permittedDataType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PermittedDataType = append(r.PermittedDataType, v)
			case "multipleResultsAllowed":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MultipleResultsAllowed = &v
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "preferredReportName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreferredReportName = &v
			case "quantitativeDetails":
				var v ObservationDefinitionQuantitativeDetails
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QuantitativeDetails = &v
			case "qualifiedInterval":
				var v ObservationDefinitionQualifiedInterval
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QualifiedInterval = append(r.QualifiedInterval, v)
			case "validCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidCodedValueSet = &v
			case "normalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NormalCodedValueSet = &v
			case "abnormalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AbnormalCodedValueSet = &v
			case "criticalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CriticalCodedValueSet = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ObservationDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Characteristics for quantitative results of this observation.
type ObservationDefinitionQuantitativeDetails struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Customary unit used to report quantitative results of observations conforming to this ObservationDefinition.
	CustomaryUnit *CodeableConcept
	// SI unit used to report quantitative results of observations conforming to this ObservationDefinition.
	Unit *CodeableConcept
	// Factor for converting value expressed with SI unit to value expressed with customary unit.
	ConversionFactor *Decimal
	// Number of digits after decimal separator when the results of such observations are of type Quantity.
	DecimalPrecision *Integer
}
type jsonObservationDefinitionQuantitativeDetails struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	CustomaryUnit                    *CodeableConcept  `json:"customaryUnit,omitempty"`
	Unit                             *CodeableConcept  `json:"unit,omitempty"`
	ConversionFactor                 *Decimal          `json:"conversionFactor,omitempty"`
	ConversionFactorPrimitiveElement *primitiveElement `json:"_conversionFactor,omitempty"`
	DecimalPrecision                 *Integer          `json:"decimalPrecision,omitempty"`
	DecimalPrecisionPrimitiveElement *primitiveElement `json:"_decimalPrecision,omitempty"`
}

func (r ObservationDefinitionQuantitativeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ObservationDefinitionQuantitativeDetails) marshalJSON() jsonObservationDefinitionQuantitativeDetails {
	m := jsonObservationDefinitionQuantitativeDetails{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.CustomaryUnit = r.CustomaryUnit
	m.Unit = r.Unit
	if r.ConversionFactor != nil && r.ConversionFactor.Value != nil {
		m.ConversionFactor = r.ConversionFactor
	}
	if r.ConversionFactor != nil && (r.ConversionFactor.Id != nil || r.ConversionFactor.Extension != nil) {
		m.ConversionFactorPrimitiveElement = &primitiveElement{Id: r.ConversionFactor.Id, Extension: r.ConversionFactor.Extension}
	}
	if r.DecimalPrecision != nil && r.DecimalPrecision.Value != nil {
		m.DecimalPrecision = r.DecimalPrecision
	}
	if r.DecimalPrecision != nil && (r.DecimalPrecision.Id != nil || r.DecimalPrecision.Extension != nil) {
		m.DecimalPrecisionPrimitiveElement = &primitiveElement{Id: r.DecimalPrecision.Id, Extension: r.DecimalPrecision.Extension}
	}
	return m
}
func (r *ObservationDefinitionQuantitativeDetails) UnmarshalJSON(b []byte) error {
	var m jsonObservationDefinitionQuantitativeDetails
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ObservationDefinitionQuantitativeDetails) unmarshalJSON(m jsonObservationDefinitionQuantitativeDetails) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.CustomaryUnit = m.CustomaryUnit
	r.Unit = m.Unit
	r.ConversionFactor = m.ConversionFactor
	if m.ConversionFactorPrimitiveElement != nil {
		if r.ConversionFactor == nil {
			r.ConversionFactor = &Decimal{}
		}
		r.ConversionFactor.Id = m.ConversionFactorPrimitiveElement.Id
		r.ConversionFactor.Extension = m.ConversionFactorPrimitiveElement.Extension
	}
	r.DecimalPrecision = m.DecimalPrecision
	if m.DecimalPrecisionPrimitiveElement != nil {
		if r.DecimalPrecision == nil {
			r.DecimalPrecision = &Integer{}
		}
		r.DecimalPrecision.Id = m.DecimalPrecisionPrimitiveElement.Id
		r.DecimalPrecision.Extension = m.DecimalPrecisionPrimitiveElement.Extension
	}
	return nil
}
func (r ObservationDefinitionQuantitativeDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.CustomaryUnit, xml.StartElement{Name: xml.Name{Local: "customaryUnit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Unit, xml.StartElement{Name: xml.Name{Local: "unit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConversionFactor, xml.StartElement{Name: xml.Name{Local: "conversionFactor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DecimalPrecision, xml.StartElement{Name: xml.Name{Local: "decimalPrecision"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ObservationDefinitionQuantitativeDetails) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "customaryUnit":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CustomaryUnit = &v
			case "unit":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Unit = &v
			case "conversionFactor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConversionFactor = &v
			case "decimalPrecision":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DecimalPrecision = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ObservationDefinitionQuantitativeDetails) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
type ObservationDefinitionQualifiedInterval struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of interval of values for continuous or ordinal observations conforming to this ObservationDefinition.
	Category *Code
	// The low and high values determining the interval. There may be only one of the two.
	Range *Range
	// Codes to indicate the health context the range applies to. For example, the normal or therapeutic range.
	Context *CodeableConcept
	// Codes to indicate the target population this reference range applies to.
	AppliesTo []CodeableConcept
	// Sex of the population the range applies to.
	Gender *Code
	// The age at which this reference range is applicable. This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *Range
	// The gestational age to which this reference range is applicable, in the context of pregnancy.
	GestationalAge *Range
	// Text based condition for which the reference range is valid.
	Condition *String
}
type jsonObservationDefinitionQualifiedInterval struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Category                  *Code             `json:"category,omitempty"`
	CategoryPrimitiveElement  *primitiveElement `json:"_category,omitempty"`
	Range                     *Range            `json:"range,omitempty"`
	Context                   *CodeableConcept  `json:"context,omitempty"`
	AppliesTo                 []CodeableConcept `json:"appliesTo,omitempty"`
	Gender                    *Code             `json:"gender,omitempty"`
	GenderPrimitiveElement    *primitiveElement `json:"_gender,omitempty"`
	Age                       *Range            `json:"age,omitempty"`
	GestationalAge            *Range            `json:"gestationalAge,omitempty"`
	Condition                 *String           `json:"condition,omitempty"`
	ConditionPrimitiveElement *primitiveElement `json:"_condition,omitempty"`
}

func (r ObservationDefinitionQualifiedInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ObservationDefinitionQualifiedInterval) marshalJSON() jsonObservationDefinitionQualifiedInterval {
	m := jsonObservationDefinitionQualifiedInterval{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Category != nil && r.Category.Value != nil {
		m.Category = r.Category
	}
	if r.Category != nil && (r.Category.Id != nil || r.Category.Extension != nil) {
		m.CategoryPrimitiveElement = &primitiveElement{Id: r.Category.Id, Extension: r.Category.Extension}
	}
	m.Range = r.Range
	m.Context = r.Context
	m.AppliesTo = r.AppliesTo
	if r.Gender != nil && r.Gender.Value != nil {
		m.Gender = r.Gender
	}
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		m.GenderPrimitiveElement = &primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
	}
	m.Age = r.Age
	m.GestationalAge = r.GestationalAge
	if r.Condition != nil && r.Condition.Value != nil {
		m.Condition = r.Condition
	}
	if r.Condition != nil && (r.Condition.Id != nil || r.Condition.Extension != nil) {
		m.ConditionPrimitiveElement = &primitiveElement{Id: r.Condition.Id, Extension: r.Condition.Extension}
	}
	return m
}
func (r *ObservationDefinitionQualifiedInterval) UnmarshalJSON(b []byte) error {
	var m jsonObservationDefinitionQualifiedInterval
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ObservationDefinitionQualifiedInterval) unmarshalJSON(m jsonObservationDefinitionQualifiedInterval) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	if m.CategoryPrimitiveElement != nil {
		if r.Category == nil {
			r.Category = &Code{}
		}
		r.Category.Id = m.CategoryPrimitiveElement.Id
		r.Category.Extension = m.CategoryPrimitiveElement.Extension
	}
	r.Range = m.Range
	r.Context = m.Context
	r.AppliesTo = m.AppliesTo
	r.Gender = m.Gender
	if m.GenderPrimitiveElement != nil {
		if r.Gender == nil {
			r.Gender = &Code{}
		}
		r.Gender.Id = m.GenderPrimitiveElement.Id
		r.Gender.Extension = m.GenderPrimitiveElement.Extension
	}
	r.Age = m.Age
	r.GestationalAge = m.GestationalAge
	r.Condition = m.Condition
	if m.ConditionPrimitiveElement != nil {
		if r.Condition == nil {
			r.Condition = &String{}
		}
		r.Condition.Id = m.ConditionPrimitiveElement.Id
		r.Condition.Extension = m.ConditionPrimitiveElement.Extension
	}
	return nil
}
func (r ObservationDefinitionQualifiedInterval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Range, xml.StartElement{Name: xml.Name{Local: "range"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Context, xml.StartElement{Name: xml.Name{Local: "context"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AppliesTo, xml.StartElement{Name: xml.Name{Local: "appliesTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Gender, xml.StartElement{Name: xml.Name{Local: "gender"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Age, xml.StartElement{Name: xml.Name{Local: "age"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GestationalAge, xml.StartElement{Name: xml.Name{Local: "gestationalAge"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Condition, xml.StartElement{Name: xml.Name{Local: "condition"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ObservationDefinitionQualifiedInterval) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = &v
			case "range":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Range = &v
			case "context":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Context = &v
			case "appliesTo":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AppliesTo = append(r.AppliesTo, v)
			case "gender":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Gender = &v
			case "age":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Age = &v
			case "gestationalAge":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GestationalAge = &v
			case "condition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Condition = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ObservationDefinitionQualifiedInterval) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
