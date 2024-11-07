package json

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementUnmarshal(f *File, s ir.Struct) {
	if s.IsPrimitive {
		implementUnmarshalPrimitive(f, s)
	} else {
		implementUnmarshalElement(f, s)
		implementUnmarshalStruct(f, s)
	}
}

func implementUnmarshalPrimitive(f *File, s ir.Struct) {
	var ty string
	for _, field := range s.Fields {
		if field.Name == "Value" {
			ty = field.PossibleTypes[0].Name
			break
		}
	}

	var unmarshal Code
	if s.Name == "Decimal" {
		unmarshal = Id("value").Op("=").String().Params(Id("b"))
	} else {
		unmarshal = If().Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Params(Id("b"), Op("&").Id("value")).Op(";").Id("err").Op("!=").Nil().Block(
			Return(Id("err")),
		)
	}

	var assign Code
	if s.Name == "Xhtml" {
		assign = Id("*r").Op("=").Id(s.Name).Values(Id("Value").Op(":").Id("value"))
	} else {
		assign = Id("*r").Op("=").Id(s.Name).Values(Id("Value").Op(":").Id("&value"))
	}

	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Params(Error()).Block(
		Var().Id("value").Id(ty),
		unmarshal,
		assign,
		Return().Nil(),
	)
}

func implementUnmarshalElement(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Params(Error()).Block(
		Var().Id("m").Id("json"+s.Name),
		If().Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Params(Id("b"), Op("&").Id("m")).Op(";").Id("err").Op("!=").Nil().Block(
			Return(Id("err")),
		),
		Return(Id("r.unmarshalJSON").Params(Id("m"))),
	)
}

func implementUnmarshalStruct(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("unmarshalJSON").Params(Id("m").Id("json" + s.Name)).Params(Error()).BlockFunc(func(g *Group) {
		for _, sf := range s.Fields {
			if sf.Polymorph {
				for _, t := range sf.PossibleTypes {
					ifCond := Id("m." + sf.Name + t.Name).Op("!=").Nil()
					if t.IsPrimitive {
						ifCond.Op("||").Id("m." + sf.Name + t.Name + "PrimitiveElement").Op("!=").Nil()
					}

					g.If(ifCond).BlockFunc(func(g *Group) {
						g.If(Id("r." + sf.Name).Op("!=").Nil()).Block(
							Return(Qual("fmt", "Errorf").Params(Lit(`multiple values for field "` + sf.Name + `"`))),
						)

						g.Id("v").Op(":=").Id("m." + sf.Name + t.Name)

						if t.IsPrimitive {
							g.If(Id("m." + sf.Name + t.Name + "PrimitiveElement").Op("!=").Nil()).BlockFunc(func(g *Group) {
								g.If(Id("v").Op("==").Nil()).Block(
									Id("v").Op("=").Op("&").Id(t.Name).Values(),
								)
								g.Id("v.Id").Op("=").Id("m." + sf.Name + t.Name + "PrimitiveElement.Id")
								g.Id("v.Extension").Op("=").Id("m." + sf.Name + t.Name + "PrimitiveElement.Extension")
							})
						}

						g.Id("r." + sf.Name).Op("=").Id("v")
					})
				}
			} else {
				t := sf.PossibleTypes[0]

				if t.IsNestedResource {
					if sf.Multiple {
						g.Id("r."+sf.Name).Op("=").Make(Index().Qual("fhir-toolbox/model", "Resource"), Lit(0), Len(Id("m."+sf.Name)))
						g.For(List(Id("_"), Id("v")).Op(":=").Range().Id("m." + sf.Name)).Block(
							Id("r."+sf.Name).Op("=").Append(Id("r."+sf.Name), Id("v.Resource")),
						)
					} else {
						g.If(Id("m." + sf.Name).Op("!=").Nil()).Block(
							Id("r." + sf.Name).Op("=").Id("&m." + sf.Name + ".Resource"),
						)
					}
				} else {
					g.Id("r." + sf.Name).Op("=").Id("m." + sf.Name)
				}

				if t.IsPrimitive {
					if sf.Multiple {
						g.For(List(Id("i"), Id("e")).Op(":=").Range().Id("m."+sf.Name+"PrimitiveElement")).Block(
							If(Len(Id("r."+sf.Name)).Op("<=").Id("i")).Block(
								Id("r."+sf.Name).Op("=").Append(Id("r."+sf.Name), Id(t.Name).Values()),
							),
							If(Id("e").Op("!=").Nil()).Block(
								Id("r."+sf.Name).Index(Id("i")).Id(".Id").Op("=").Id("e.Id"),
								Id("r."+sf.Name).Index(Id("i")).Id(".Extension").Op("=").Id("e.Extension"),
							),
						)
					} else {
						g.If(Id("m." + sf.Name + "PrimitiveElement").Op("!=").Nil()).BlockFunc(func(g *Group) {
							if sf.Optional {
								g.If(Id("r." + sf.Name).Op("==").Nil()).Block(
									Id("r." + sf.Name).Op("=").Op("&").Id(t.Name).Values(),
								)
							}
							g.Id("r." + sf.Name + ".Id").Op("=").Id("m." + sf.Name + "PrimitiveElement.Id")
							if sf.Name != "Div" {
								g.Id("r." + sf.Name + ".Extension").Op("=").Id("m." + sf.Name + "PrimitiveElement.Extension")
							}
						})
					}
				}
			}
		}

		g.Return(Nil())
	})
}
