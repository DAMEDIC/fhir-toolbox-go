package xml

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementUnmarshal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("UnmarshalXML").Params(
		Id("d").Op("*").Qual("encoding/xml", "Decoder"),
		Id("start").Qual("encoding/xml", "StartElement"),
	).Params(Error()).BlockFunc(func(g *Group) {
		//g.If(Id("r").Op("==").Nil()).Block(
		//	Id("*r").Op("=").Id(s.Name).Values(),
		//)

		validateStructNamespace(g, s)
		unmarshalAttrs(g, s)

		if s.Name == "Xhtml" {
			unmarshalXhtml(g)

			return
		}

		unmarshalFields(g, s)
	})
}

func expectedNamespace(s ir.Struct) string {
	if s.Name == "Xhtml" {
		return NamespaceXHTML
	} else {
		return NamespaceFHIR
	}
}

func validateStructNamespace(g *Group, s ir.Struct) {
	expectedNamespace := expectedNamespace(s)

	g.If(Id("start.Name.Space").Op("!=").Lit(expectedNamespace)).Block(
		Return(Qual("fmt", "Errorf").Params(
			Lit("invalid namespace: \"%s\", expected: \""+expectedNamespace+"\""),
			Id("start.Name.Space"),
		)),
	)
}

func validateAttrNamespace(g *Group) *Statement {
	return g.If(Id("a.Name.Space").Op("!=").Lit("")).Block(
		Return(Qual("fmt", "Errorf").Params(
			Lit("invalid attribute namespace: \"%s\", expected default namespace"),
			Id("start.Name.Space"),
		)),
	)
}

func unmarshalAttrs(g *Group, s ir.Struct) {
	g.For(List(Id("_"), Id("a")).Op(":=").Range().Id("start.Attr")).BlockFunc(func(g *Group) {
		validateAttrNamespace(g)

		g.Switch(Id("a.Name.Local")).BlockFunc(func(g *Group) {
			g.Case(Lit("xmlns")).Block(
				// namespace was validated earlier
				Continue(),
			)
			if !s.IsResource {
				g.Case(Lit("id")).Block(
					Id("r.Id").Op("=").Id("&a.Value"),
				)
			}
			if s.IsPrimitive && s.Name != "Xhtml" {
				unmarshalPrimitiveValueAttr(g, s)
			}
			if s.Name == "Extension" {
				g.Case(Lit("url")).Block(
					Id("r.Url").Op("=").Id("a.Value"),
				)
			}
			g.Default().Block(Return(Qual("fmt", "Errorf").Params(
				Lit("invalid attribute: \"%s\""),
				Id("a.Name.Local"),
			)))
		})
	})
}

func unmarshalPrimitiveValueAttr(g *Group, s ir.Struct) {
	g.Case(Lit("value")).BlockFunc(func(g *Group) {
		if s.Name == "Boolean" {
			g.Var().Id("v").Bool()
			g.If(Id("a.Value").Op("==").Lit("true")).Block(
				Id("v").Op("=").True(),
			).Else().If(Id("a.Value").Op("==").Lit("false")).Block(
				Id("v").Op("=").False(),
			).Else().Block(
				Return(Qual("fmt", "Errorf").Params(
					Lit("can not parse \"%s\" as bool"),
					Id("a.Value"),
				)),
			)
			g.Id("r.Value").Op("=").Id("&v")
		} else if s.Name == "Integer" {
			g.List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(Id("a.Value"), Lit(10), Lit(0))
			g.If(Err().Op("!=").Nil()).Block(
				Return(Id("err")),
			)
			g.Id("v").Op(":=").Id("int32").Call(Id("i"))
			g.Id("r.Value").Op("=").Id("&v")
		} else if s.Name == "PositiveInt" || s.Name == "UnsignedInt" {
			g.List(Id("i"), Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(Id("a.Value"), Lit(10), Lit(0))
			g.If(Err().Op("!=").Nil()).Block(
				Return(Id("err")),
			)
			g.Id("v").Op(":=").Id("uint32").Call(Id("i"))
			g.Id("r.Value").Op("=").Id("&v")
		} else {
			g.Id("r.Value").Op("=").Id("&a.Value")
		}
	})
}

func unmarshalFields(g *Group, s ir.Struct) {
	g.For().Block(
		List(Id("token"), Id("err")).Op(":=").Id("d.Token").Params(),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("err")),
		),
		Switch(Id("t").Op(":=").Id("token.(type)")).Block(
			Case(Qual("encoding/xml", "StartElement")).BlockFunc(func(g *Group) {
				g.Switch(Id("t.Name.Local")).BlockFunc(func(g *Group) {
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

						unmarshalField(g, f)
					}
				})
			}),
			Case(Qual("encoding/xml", "EndElement")).Block(Return(Nil())),
		),
	)
}

func unmarshalField(g *Group, f ir.StructField) {
	if f.Polymorph {
		for _, t := range f.PossibleTypes {
			g.Case(Lit(f.MarshalName+t.Name)).Block(
				If(Id("r."+f.Name).Op("!=").Nil()).Block(
					Return(Qual("fmt", "Errorf").Params(Lit(`multiple values for field "`+f.Name+`"`))),
				),
				Var().Id("v").Id(t.Name),
				Id("err").Op(":=").Id("d.DecodeElement").Call(Op("&").Id("v"), Id("&t")),
				If(Err().Op("!=").Nil()).Block(
					Return(Id("err")),
				),
				Id("r."+f.Name).Op("=").Id("v"),
			)
		}
	} else {
		t := f.PossibleTypes[0]

		g.Case(Lit(f.MarshalName)).BlockFunc(func(g *Group) {
			if t.IsNestedResource {
				g.Var().Id("c").Id("ContainedResource")
				g.Id("err").Op(":=").Id("d.DecodeElement").Call(Op("&").Id("c"), Id("&t"))
				g.If(Err().Op("!=").Nil()).Block(
					Return(Id("err")),
				)
				if f.Multiple {
					g.Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id("c.Resource"))
				} else {
					g.Id("r." + f.Name).Op("=").Id("&c.Resource")
				}
			} else {
				g.Var().Id("v").Id(t.Name)

				g.Id("err").Op(":=").Id("d.DecodeElement").Call(Id("&v"), Id("&t"))
				g.If(Err().Op("!=").Nil()).Block(
					Return(Id("err")),
				)

				if f.Multiple {
					g.Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id("v"))
				} else if f.Optional {
					g.Id("r." + f.Name).Op("=").Id("&v")
				} else {
					g.Id("r." + f.Name).Op("=").Id("v")
				}
			}
		})
	}
}

func unmarshalXhtml(g *Group) {
	g.Var().Id("v").Struct(
		Id("V").String().Tag(map[string]string{"xml": ",innerxml"}),
	)
	g.Id("err").Op(":=").Id("d.DecodeElement").Call(Id("&v"), Id("&start"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("err")),
	)
	g.Id("r.Value").Op("=").Id("v.V")
	g.Return(Nil())
}
