package r4

import "encoding/json"

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
func (r MarketingStatus) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
