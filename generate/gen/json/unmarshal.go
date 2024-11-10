package json

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementUnmarshal(f *File, s ir.Struct) {
	if s.IsResource {
		implementUnmarshalExternal(f, s)
	}

	if s.IsPrimitive {
		implementUnmarshalPrimitive(f, s)
	} else {
		implementUnmarshalInternal(f, s)
	}
}

func implementUnmarshalExternal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("UnmarshalJSON").Params(
		Id("b").Index().Byte(),
	).Params(Error()).Block(
		Id("d").Op(":=").Qual("encoding/json", "NewDecoder").Call(
			Qual("bytes", "NewReader").Call(Id("b")),
		),
		Return(Id("r").Dot("unmarshalJSON").Call(Id("d"))),
	)
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
		unmarshal = Id("v").Op("=").String().Params(Id("b"))
	} else {
		unmarshal = If(
			Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Params(Id("b"), Op("&").Id("v")),
			Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
		)
	}

	var assign Code
	if s.Name == "Xhtml" {
		assign = Id("*r").Op("=").Id(s.Name).Values(Id("Value").Op(":").Id("v"))
	} else {
		assign = Id("*r").Op("=").Id(s.Name).Values(Id("Value").Op(":").Id("&v"))
	}

	f.Func().Params(
		Id("r").Op("*").Id(s.Name),
	).Id("UnmarshalJSON").Params(
		Id("b").Index().Byte(),
	).Params(Error()).Block(
		Var().Id("v").Id(ty),
		unmarshal,
		assign,
		Return().Nil(),
	)
}

func implementUnmarshalInternal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Op("*").Id(s.Name)).Id("unmarshalJSON").Params(
		Id("d").Op("*").Qual("encoding/json", "Decoder"),
	).Params(Error()).BlockFunc(func(g *Group) {
		g.List(Id("t"), Err()).Op(":=").Id("d").Dot("Token").Call()
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
		g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('{')).Block(
			returnInvalidTokenError(s.Name, "'{'"),
		)

		unmarshalFields(g, s)

		g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
		g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('}')).Block(
			returnInvalidTokenError(s.Name, "'}'"),
		)
		g.Return(Nil())
	})
}

func unmarshalFields(g *Group, s ir.Struct) {
	g.For(Id("d").Dot("More").Call()).Block(
		List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call(),
		If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		),
		List(Id("f"), Id("ok")).Op(":=").Id("t").Op(".").Call(String()),
		If(Op("!").Id("ok")).Block(
			returnInvalidTokenError(s.Name, "field name"),
		),
		Switch(Id("f")).BlockFunc(func(g *Group) {
			if s.IsResource {
				g.Case(Lit("resourceType")).Block(
					List(Id("_"), Id("err")).Op(":=").Id("d").Dot("Token").Call(),
					If(Err().Op("!=").Nil()).Block(
						Return(Id("err")),
					),
				)
			}

			for _, f := range s.Fields {
				unmarshalField(g, s, f)
			}

			g.Default().Block(
				Return(Qual("fmt", "Errorf").Params(
					Lit("invalid field: %s in "+s.Name+""),
					Id("f"),
				)),
			)
		}),
	)
}

func unmarshalField(g *Group, s ir.Struct, f ir.StructField) {
	if f.Polymorph {
		unmarshalCasePolymorph(g, s, f)
	} else {
		unmarshalCaseNonPolymorph(g, s, f)
	}
}

func unmarshalCasePolymorph(g *Group, s ir.Struct, f ir.StructField) {
	for _, t := range f.PossibleTypes {
		g.Case(Lit(f.MarshalName + t.Name)).BlockFunc(func(g *Group) {
			unmarshalToValue(g, s, f, t)

			if t.IsPrimitive {
				g.If(Id("r." + f.Name).Op("!=").Nil()).Block(
					Id("r." + f.Name).Op("=").Id(t.Name).Values(Dict{
						Id("Id"):        Id("r." + f.Name).Dot("").Call(Id(t.Name)).Dot("Id"),
						Id("Extension"): Id("r." + f.Name).Dot("").Call(Id(t.Name)).Dot("Extension"),
						Id("Value"):     Id("v").Dot("Value"),
					}),
				).Else().Block(
					Id("r." + f.Name).Op("=").Id("v"),
				)
			} else {
				g.Id("r." + f.Name).Op("=").Id("v")
			}
		})

		if t.IsPrimitive {
			g.Case(Lit("_" + f.MarshalName + t.Name)).BlockFunc(func(g *Group) {
				unmarshalPrimitiveElement(g)

				g.If(Id("r." + f.Name).Op("!=").Nil()).Block(
					Id("r." + f.Name).Op("=").Id(t.Name).Values(Dict{
						Id("Id"):        Id("v").Dot("Id"),
						Id("Extension"): Id("v").Dot("Extension"),
						Id("Value"):     Id("r." + f.Name).Dot("").Call(Id(t.Name)).Dot("Value"),
					}),
				).Else().Block(
					Id("r." + f.Name).Op("=").Id(t.Name).Values(Dict{
						Id("Id"):        Id("v").Dot("Id"),
						Id("Extension"): Id("v").Dot("Extension"),
					}),
				)
			})
		}
	}
}

func unmarshalCaseNonPolymorph(g *Group, s ir.Struct, f ir.StructField) {
	t := f.PossibleTypes[0]

	g.Case(Lit(f.MarshalName)).BlockFunc(func(g *Group) {
		if f.Multiple {
			unmarshalMultiple(g, s, f, t)
		} else {
			unmarshalToValue(g, s, f, t)

			if t.IsNestedResource {
				g.Id("r." + f.Name).Op("=").Id("&v").Dot("Resource")
			} else if f.Optional {
				if t.IsPrimitive {
					g.If(Id("r." + f.Name).Op("==").Nil()).Block(
						Id("r." + f.Name).Op("=").Op("&").Id(t.Name).Values(),
					)
					g.Id("r." + f.Name).Dot("Value").Op("=").Id("v.Value")
				} else {
					g.Id("r." + f.Name).Op("=").Id("&v")
				}
			} else {
				if t.IsPrimitive {
					g.Id("r." + f.Name).Dot("Value").Op("=").Id("v.Value")
				} else {
					g.Id("r." + f.Name).Op("=").Id("v")
				}
			}
		}
	})

	if t.IsPrimitive {
		g.Case(Lit("_" + f.MarshalName)).BlockFunc(func(g *Group) {
			if f.Multiple {
				unmarshalPrimitiveElementMultiple(g, s, f, t)
			} else {
				unmarshalPrimitiveElement(g)

				if f.Optional {
					g.If(Id("r." + f.Name).Op("==").Nil()).Block(
						Id("r." + f.Name).Op("=").Op("&").Id(t.Name).Values(),
					)
				}

				g.Id("r." + f.Name).Dot("Id").Op("=").Id("v.Id")
				if !(t.Name == "Xhtml" && f.Name == "Div") {
					g.Id("r." + f.Name).Dot("Extension").Op("=").Id("v.Extension")
				}
			}
		})
	}
}

func unmarshalMultiple(g *Group, s ir.Struct, f ir.StructField, t ir.FieldType) {
	g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
	g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('[')).Block(
		returnInvalidTokenError(s.Name, "'['"),
	)

	if t.IsPrimitive {
		g.For(
			Id("i").Op(":=").Lit(0),
			Id("d").Dot("More").Call(),
			Id("i").Op("++"),
		).BlockFunc(func(g *Group) {
			unmarshalToValue(g, s, f, t)

			g.For(Len(Id("r." + f.Name)).Op("<=").Id("i")).Block(
				Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id(t.Name).Values()),
			)
			g.Id("r." + f.Name).Index(Id("i")).Dot("Value").Op("=").Id("v.Value")
		})
	} else {
		g.For(Id("d").Dot("More").Call()).BlockFunc(func(g *Group) {
			unmarshalToValue(g, s, f, t)

			if t.IsNestedResource {
				g.Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id("v").Dot("Resource"))
			} else {
				g.Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id("v"))
			}
		})
	}

	g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
	g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune(']')).Block(
		returnInvalidTokenError(s.Name, "']'"),
	)
}

func unmarshalToValue(g *Group, s ir.Struct, f ir.StructField, t ir.FieldType) {
	if t.IsNestedResource {
		g.Var().Id("v").Id("ContainedResource")
	} else {
		g.Var().Id("v").Id(t.Name)
	}

	if t.IsPrimitive ||
		(!s.IsResource && f.Name == "Id") ||
		(s.Name == "Extension" && f.Name == "Url") {
		g.Id("err").Op(":=").Id("d").Dot("Decode").Call(Id("&v"))
	} else {
		g.Id("err").Op(":=").Id("v").Dot("unmarshalJSON").Call(Id("d"))
	}

	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("err")),
	)
}

func unmarshalPrimitiveElementMultiple(g *Group, s ir.Struct, f ir.StructField, t ir.FieldType) {
	g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
	g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('[')).Block(
		returnInvalidTokenError(s.Name, "'['"),
	)

	g.For(
		Id("i").Op(":=").Lit(0),
		Id("d").Dot("More").Call(),
		Id("i").Op("++"),
	).BlockFunc(func(g *Group) {
		unmarshalPrimitiveElement(g)

		g.For(Len(Id("r." + f.Name)).Op("<=").Id("i")).Block(
			Id("r."+f.Name).Op("=").Append(Id("r."+f.Name), Id(t.Name).Values()),
		)
		g.Id("r." + f.Name).Index(Id("i")).Dot("Id").Op("=").Id("v.Id")
		g.Id("r." + f.Name).Index(Id("i")).Dot("Extension").Op("=").Id("v.Extension")
	})

	g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
	g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune(']')).Block(
		returnInvalidTokenError(s.Name, "']'"),
	)
}

func unmarshalPrimitiveElement(g *Group) {
	g.Var().Id("v").Id("primitiveElement")

	g.Id("err").Op(":=").Id("v").Dot("unmarshalJSON").Call(Id("d"))

	g.If(Err().Op("!=").Nil()).Block(
		Return(Id("err")),
	)
}

func returnInvalidTokenError(in string, expected string) *Statement {
	return Return(Qual("fmt", "Errorf").Params(
		Lit("invalid token: %v, expected: "+expected+" in "+in+" element"),
		Id("t"),
	))
}
