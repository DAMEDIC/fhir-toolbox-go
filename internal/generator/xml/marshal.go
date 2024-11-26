// Package json generated code for (un)marshalling FHIR reosurces to and from XML.
package xml

import (
	"fhir-toolbox/internal/generator"
	"fhir-toolbox/internal/generator/ir"
	. "github.com/dave/jennifer/jen"
	"strings"
)

var (
	NamespaceFHIR  = "http://hl7.org/fhir"
	NamespaceXHTML = "http://www.w3.org/1999/xhtml"
)

type MarshalGenerator struct {
	generator.NoOpGenerator
}

func (g MarshalGenerator) GenerateType(f *File, rt ir.ResourceOrType) bool {
	for _, t := range rt.Structs {
		implementMarshal(f, t)
	}

	return true
}

func (g MarshalGenerator) GenerateAdditional(f func(fileName string, pkgName string) *File, release string, rt []ir.ResourceOrType) {
	implementMarshalContained(f("contained", strings.ToLower(release)))
}

func implementMarshal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("MarshalXML").Params(
		Id("e").Op("*").Qual("encoding/xml", "Encoder"),
		Id("start").Qual("encoding/xml", "StartElement"),
	).Params(Error()).BlockFunc(func(g *Group) {

		if s.IsResource {
			g.Id("start.Name.Local").Op("=").Lit(s.Name)
		} else if s.Name == "Xhtml" {
			g.Id("start.Name.Space").Op("=").Lit(NamespaceXHTML)
		}

		collectAttrs(g, s)

		if s.Name == "Xhtml" {
			writeXHTMLElement(g)

			return
		}

		g.Err().Op(":=").Id("e.EncodeToken").Params(Id("start"))
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)

		for _, f := range s.Fields {
			if !s.IsResource && f.Name == "Id" {
				continue
			}
			if s.IsPrimitive && f.Name == "Value" {
				continue
			}
			if s.Name == "Extension" && f.Name == "Url" {
				continue
			}

			marshalField(g, f)
		}

		g.Err().Op("=").Id("e.EncodeToken").Params(Id("start.End()"))
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)

		g.Return(Nil())
	})
}

func collectAttrs(g *Group, s ir.Struct) {
	if !s.IsResource {
		g.If(Id("r.Id").Op("!=").Nil()).Block(
			Id("start.Attr").Op("=").Append(Id("start.Attr"), Qual("encoding/xml", "Attr").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit("id"),
				}),
				Id("Value"): Id("*r.Id"),
			})),
		)
	}
	if s.IsPrimitive && s.Name != "Xhtml" {
		marshalPrimitiveValueAttr(g, s)
	}
	if s.Name == "Extension" {
		g.Id("start.Attr").Op("=").Append(Id("start.Attr"), Qual("encoding/xml", "Attr").Values(Dict{
			Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
				Id("Local"): Lit("url"),
			}),
			Id("Value"): Id("r.Url"),
		}))
	}
}

func marshalPrimitiveValueAttr(g *Group, s ir.Struct) {
	if s.Name == "Boolean" {
		g.If(Id("r.Value").Op("!=").Nil()).Block(
			Id("start.Attr").Op("=").Append(Id("start.Attr"), Qual("encoding/xml", "Attr").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit("value"),
				}),
				Id("Value"): Qual("strconv", "FormatBool").Call(Id("*r.Value")),
			})),
		)
	} else if s.Name == "Integer" || s.Name == "PositiveInt" || s.Name == "UnsignedInt" {
		g.If(Id("r.Value").Op("!=").Nil()).Block(
			Id("start.Attr").Op("=").Append(Id("start.Attr"), Qual("encoding/xml", "Attr").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit("value"),
				}),
				Id("Value"): Qual("strconv", "FormatInt").Call(Int64().Call(Id("*r.Value")), Lit(10)),
			})),
		)
	} else {
		g.If(Id("r.Value").Op("!=").Nil()).Block(
			Id("start.Attr").Op("=").Append(Id("start.Attr"), Qual("encoding/xml", "Attr").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit("value"),
				}),
				Id("Value"): Id("*r.Value"),
			})),
		)
	}
}

func marshalField(g *Group, f ir.StructField) {
	if f.Polymorph {
		g.Switch(Id("v").Op(":=").Id("r").Op(".").Id(f.Name).Op(".").Params(Type())).BlockFunc(func(g *Group) {
			for _, t := range f.PossibleTypes {
				marshalCase(g, f, t)
			}
		})
	} else {
		t := f.PossibleTypes[0]

		if t.IsNestedResource {
			encodeContainedResource(g, f)
		} else {
			encodeElement(g, f)
		}
	}
}

func marshalCase(g *Group, f ir.StructField, t ir.FieldType) {
	g.Case(List(Id(t.Name), Op("*").Id(t.Name))).BlockFunc(func(g *Group) {
		g.Err().Op(":=").Id("e.EncodeElement").Call(
			Id("v"),
			Qual("encoding/xml", "StartElement").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit(f.MarshalName + t.Name),
				}),
			}),
		)
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
	})
}

func encodeContainedResource(g *Group, f ir.StructField) {
	if f.Multiple {
		g.Id("v").Op(":=").Id("make").Params(Id("[]ContainedResource"), Lit(0), Len(Id("r."+f.Name)))
		g.For(Id("_, c").Op(":=").Range().Id("r." + f.Name)).Block(
			Id("v").Op("=").Append(Id("v"), Id("ContainedResource").Values(Id("c"))),
		)
		g.Err().Op("=").Id("e.EncodeElement").Call(
			Id("v"),
			Qual("encoding/xml", "StartElement").Values(Dict{
				Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
					Id("Local"): Lit(f.MarshalName),
				}),
			}),
		)
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
	} else {
		g.If(Id("r."+f.Name).Op("!=").Nil()).Block(
			Id("v").Op(":=").Id("ContainedResource").Values(Id("*r."+f.Name)),
			Err().Op("=").Id("e.EncodeElement").Call(
				Id("v"),
				Qual("encoding/xml", "StartElement").Values(Dict{
					Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
						Id("Local"): Lit(f.MarshalName),
					}),
				}),
			),
			If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			),
		)
	}
}

func encodeElement(g *Group, f ir.StructField) {
	g.Err().Op("=").Id("e.EncodeElement").Call(
		Id("r."+f.Name),
		Qual("encoding/xml", "StartElement").Values(Dict{
			Id("Name"): Qual("encoding/xml", "Name").Values(Dict{
				Id("Local"): Lit(f.MarshalName),
			}),
		}),
	)
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
}

func writeXHTMLElement(g *Group) {
	g.Var().Id("v").Struct(
		Id("V").Op("*").String().Tag(map[string]string{"xml": ",innerxml"}),
	)
	g.Id("v.V").Op("=").Id("&r.Value")
	g.Id("err").Op(":=").Id("e.EncodeElement").Call(Id("v"), Id("start"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("err")),
	)
	g.Return(Nil())
}

func implementMarshalContained(f *File) {
	f.Func().Params(Id("r").Id("ContainedResource")).Id("MarshalXML").Params(
		Id("e").Op("*").Qual("encoding/xml", "Encoder"),
		Id("start").Qual("encoding/xml", "StartElement"),
	).Error().Block(
		If(Id("start.Name.Local").Op("==").Lit("ContainedResource")).Block(
			Id("start.Name.Space").Op("=").Lit(NamespaceFHIR),
			Return(Id("e").Op(".").Id("EncodeElement").Params(Id("r.Resource"), Id("start"))),
		).Else().Block(
			Err().Op(":=").Id("e.EncodeToken").Params(Id("start")),
			If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			),

			Err().Op("=").Id("e.Encode").Params(Id("r.Resource")),
			If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			),

			Err().Op("=").Id("e.EncodeToken").Params(Id("start.End()")),
			If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			),

			Return(Nil()),
		),
	)
}
