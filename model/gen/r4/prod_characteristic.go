package r4

import "encoding/json"

// Base StructureDefinition for ProdCharacteristic Type: The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type ProdCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Where applicable, the height can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Height *Quantity
	// Where applicable, the width can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Width *Quantity
	// Where applicable, the depth can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Depth *Quantity
	// Where applicable, the weight can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	Weight *Quantity
	// Where applicable, the nominal volume can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	NominalVolume *Quantity
	// Where applicable, the external diameter can be specified using a numerical value and its unit of measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used.
	ExternalDiameter *Quantity
	// Where applicable, the shape can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Shape *String
	// Where applicable, the color can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Color []String
	// Where applicable, the imprint can be specified as text.
	Imprint []String
	// Where applicable, the image can be provided The format of the image attachment shall be specified by regional implementations.
	Image []Attachment
	// Where applicable, the scoring can be specified An appropriate controlled vocabulary shall be used The term and the term identifier shall be used.
	Scoring *CodeableConcept
}
type jsonProdCharacteristic struct {
	Id                      *string             `json:"id,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	Height                  *Quantity           `json:"height,omitempty"`
	Width                   *Quantity           `json:"width,omitempty"`
	Depth                   *Quantity           `json:"depth,omitempty"`
	Weight                  *Quantity           `json:"weight,omitempty"`
	NominalVolume           *Quantity           `json:"nominalVolume,omitempty"`
	ExternalDiameter        *Quantity           `json:"externalDiameter,omitempty"`
	Shape                   *String             `json:"shape,omitempty"`
	ShapePrimitiveElement   *primitiveElement   `json:"_shape,omitempty"`
	Color                   []String            `json:"color,omitempty"`
	ColorPrimitiveElement   []*primitiveElement `json:"_color,omitempty"`
	Imprint                 []String            `json:"imprint,omitempty"`
	ImprintPrimitiveElement []*primitiveElement `json:"_imprint,omitempty"`
	Image                   []Attachment        `json:"image,omitempty"`
	Scoring                 *CodeableConcept    `json:"scoring,omitempty"`
}

func (r ProdCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ProdCharacteristic) marshalJSON() jsonProdCharacteristic {
	m := jsonProdCharacteristic{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Height = r.Height
	m.Width = r.Width
	m.Depth = r.Depth
	m.Weight = r.Weight
	m.NominalVolume = r.NominalVolume
	m.ExternalDiameter = r.ExternalDiameter
	if r.Shape != nil && r.Shape.Value != nil {
		m.Shape = r.Shape
	}
	if r.Shape != nil && (r.Shape.Id != nil || r.Shape.Extension != nil) {
		m.ShapePrimitiveElement = &primitiveElement{Id: r.Shape.Id, Extension: r.Shape.Extension}
	}
	anyColorValue := false
	for _, e := range r.Color {
		if e.Value != nil {
			anyColorValue = true
			break
		}
	}
	if anyColorValue {
		m.Color = r.Color
	}
	anyColorIdOrExtension := false
	for _, e := range r.Color {
		if e.Id != nil || e.Extension != nil {
			anyColorIdOrExtension = true
			break
		}
	}
	if anyColorIdOrExtension {
		m.ColorPrimitiveElement = make([]*primitiveElement, 0, len(r.Color))
		for _, e := range r.Color {
			if e.Id != nil || e.Extension != nil {
				m.ColorPrimitiveElement = append(m.ColorPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ColorPrimitiveElement = append(m.ColorPrimitiveElement, nil)
			}
		}
	}
	anyImprintValue := false
	for _, e := range r.Imprint {
		if e.Value != nil {
			anyImprintValue = true
			break
		}
	}
	if anyImprintValue {
		m.Imprint = r.Imprint
	}
	anyImprintIdOrExtension := false
	for _, e := range r.Imprint {
		if e.Id != nil || e.Extension != nil {
			anyImprintIdOrExtension = true
			break
		}
	}
	if anyImprintIdOrExtension {
		m.ImprintPrimitiveElement = make([]*primitiveElement, 0, len(r.Imprint))
		for _, e := range r.Imprint {
			if e.Id != nil || e.Extension != nil {
				m.ImprintPrimitiveElement = append(m.ImprintPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ImprintPrimitiveElement = append(m.ImprintPrimitiveElement, nil)
			}
		}
	}
	m.Image = r.Image
	m.Scoring = r.Scoring
	return m
}
func (r *ProdCharacteristic) UnmarshalJSON(b []byte) error {
	var m jsonProdCharacteristic
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ProdCharacteristic) unmarshalJSON(m jsonProdCharacteristic) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Height = m.Height
	r.Width = m.Width
	r.Depth = m.Depth
	r.Weight = m.Weight
	r.NominalVolume = m.NominalVolume
	r.ExternalDiameter = m.ExternalDiameter
	r.Shape = m.Shape
	if m.ShapePrimitiveElement != nil {
		if r.Shape == nil {
			r.Shape = &String{}
		}
		r.Shape.Id = m.ShapePrimitiveElement.Id
		r.Shape.Extension = m.ShapePrimitiveElement.Extension
	}
	r.Color = m.Color
	for i, e := range m.ColorPrimitiveElement {
		if len(r.Color) <= i {
			r.Color = append(r.Color, String{})
		}
		if e != nil {
			r.Color[i].Id = e.Id
			r.Color[i].Extension = e.Extension
		}
	}
	r.Imprint = m.Imprint
	for i, e := range m.ImprintPrimitiveElement {
		if len(r.Imprint) <= i {
			r.Imprint = append(r.Imprint, String{})
		}
		if e != nil {
			r.Imprint[i].Id = e.Id
			r.Imprint[i].Extension = e.Extension
		}
	}
	r.Image = m.Image
	r.Scoring = m.Scoring
	return nil
}
func (r ProdCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
