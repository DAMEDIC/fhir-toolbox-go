package r4

import "encoding/json"

// Base StructureDefinition for Expression Type: A expression that is evaluated in a specified context and returns a value. The context of use of the expression must specify the context in which the expression is evaluated, and how the result of the expression is used.
type Expression struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A brief, natural language description of the condition that effectively communicates the intended semantics.
	Description *String
	// A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
	Name *Id
	// The media type of the language for the expression.
	Language Code
	// An expression in the specified language that returns a value.
	Expression *String
	// A URI that defines where the expression is found.
	Reference *Uri
}
type jsonExpression struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Name                        *Id               `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	Language                    Code              `json:"language,omitempty"`
	LanguagePrimitiveElement    *primitiveElement `json:"_language,omitempty"`
	Expression                  *String           `json:"expression,omitempty"`
	ExpressionPrimitiveElement  *primitiveElement `json:"_expression,omitempty"`
	Reference                   *Uri              `json:"reference,omitempty"`
	ReferencePrimitiveElement   *primitiveElement `json:"_reference,omitempty"`
}

func (r Expression) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Expression) marshalJSON() jsonExpression {
	m := jsonExpression{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Language = r.Language
	if r.Language.Id != nil || r.Language.Extension != nil {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Expression = r.Expression
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	m.Reference = r.Reference
	if r.Reference != nil && (r.Reference.Id != nil || r.Reference.Extension != nil) {
		m.ReferencePrimitiveElement = &primitiveElement{Id: r.Reference.Id, Extension: r.Reference.Extension}
	}
	return m
}
func (r *Expression) UnmarshalJSON(b []byte) error {
	var m jsonExpression
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Expression) unmarshalJSON(m jsonExpression) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	r.Reference = m.Reference
	if m.ReferencePrimitiveElement != nil {
		r.Reference.Id = m.ReferencePrimitiveElement.Id
		r.Reference.Extension = m.ReferencePrimitiveElement.Extension
	}
	return nil
}
func (r Expression) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
