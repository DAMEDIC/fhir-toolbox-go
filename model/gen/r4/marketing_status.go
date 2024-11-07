package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for MarketingStatus Type: The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type MarketingStatus struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The country in which the marketing authorisation has been granted shall be specified It should be specified using the ISO 3166 ‑ 1 alpha-2 code elements.
	Country CodeableConcept
	// Where a Medicines Regulatory Agency has granted a marketing authorisation for which specific provisions within a jurisdiction apply, the jurisdiction can be specified using an appropriate controlled terminology The controlled term and the controlled term identifier shall be specified.
	Jurisdiction *CodeableConcept
	// This attribute provides information on the status of the marketing of the medicinal product See ISO/TS 20443 for more information and examples.
	Status CodeableConcept
	// The date when the Medicinal Product is placed on the market by the Marketing Authorisation Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	DateRange Period
	// The date when the Medicinal Product is placed on the market by the Marketing Authorisation Holder (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction shall be provided A complete date consisting of day, month and year shall be specified using the ISO 8601 date format NOTE “Placed on the market” refers to the release of the Medicinal Product into the distribution chain.
	RestoreDate *DateTime
}
type jsonMarketingStatus struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Country                     CodeableConcept   `json:"country,omitempty"`
	Jurisdiction                *CodeableConcept  `json:"jurisdiction,omitempty"`
	Status                      CodeableConcept   `json:"status,omitempty"`
	DateRange                   Period            `json:"dateRange,omitempty"`
	RestoreDate                 *DateTime         `json:"restoreDate,omitempty"`
	RestoreDatePrimitiveElement *primitiveElement `json:"_restoreDate,omitempty"`
}

func (r MarketingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MarketingStatus) marshalJSON() jsonMarketingStatus {
	m := jsonMarketingStatus{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Country = r.Country
	m.Jurisdiction = r.Jurisdiction
	m.Status = r.Status
	m.DateRange = r.DateRange
	if r.RestoreDate != nil && r.RestoreDate.Value != nil {
		m.RestoreDate = r.RestoreDate
	}
	if r.RestoreDate != nil && (r.RestoreDate.Id != nil || r.RestoreDate.Extension != nil) {
		m.RestoreDatePrimitiveElement = &primitiveElement{Id: r.RestoreDate.Id, Extension: r.RestoreDate.Extension}
	}
	return m
}
func (r *MarketingStatus) UnmarshalJSON(b []byte) error {
	var m jsonMarketingStatus
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MarketingStatus) unmarshalJSON(m jsonMarketingStatus) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Country = m.Country
	r.Jurisdiction = m.Jurisdiction
	r.Status = m.Status
	r.DateRange = m.DateRange
	r.RestoreDate = m.RestoreDate
	if m.RestoreDatePrimitiveElement != nil {
		if r.RestoreDate == nil {
			r.RestoreDate = &DateTime{}
		}
		r.RestoreDate.Id = m.RestoreDatePrimitiveElement.Id
		r.RestoreDate.Extension = m.RestoreDatePrimitiveElement.Extension
	}
	return nil
}
func (r MarketingStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateRange, xml.StartElement{Name: xml.Name{Local: "dateRange"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RestoreDate, xml.StartElement{Name: xml.Name{Local: "restoreDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MarketingStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "dateRange":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateRange = v
			case "restoreDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RestoreDate = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MarketingStatus) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
