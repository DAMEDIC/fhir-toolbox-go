package json

import (
	"fhir-toolbox/generate/ir"
	. "github.com/dave/jennifer/jen"
)

func ImplementMarshal(f *File, s ir.Struct) {
	if s.IsPrimitive {
		implementMarshalPrimitive(f, s)
	} else {
		implementMarshalExternal(f, s)
		implementMarshalInternal(f, s)
	}
}

func implementMarshalPrimitive(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).BlockFunc(func(g *Group) {
		if s.Name == "Decimal" {
			g.Return(Index().Byte().Params(Op("*").Id("r").Op(".").Id("Value")), Nil())
		} else {
			g.Var().Id("b").Qual("bytes", "Buffer")
			g.Id("enc").Op(":=").Qual("encoding/json", "NewEncoder").Call(Id("&b"))
			g.Id("enc").Dot("SetEscapeHTML").Call(False())
			g.Err().Op(":=").Id("enc").Dot("Encode").Call(Id("r").Op(".").Id("Value"))
			g.If(Err().Op("!=").Nil()).Block(
				Return(Nil(), Err()),
			)
			g.Return(Id("b").Dot("Bytes").Call(), Nil())
		}
	})
}

func implementMarshalExternal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).Block(
		Var().Id("b").Qual("bytes", "Buffer"),
		Err().Op(":=").Id("r").Dot("marshalJSON").Call(Id("&b")),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), Err()),
		),
		Return(Id("b").Dot("Bytes").Call(), Nil()),
	)
}

func implementMarshalInternal(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("marshalJSON").Params(
		Id("w").Qual("io", "Writer"),
	).Params(Error()).BlockFunc(func(g *Group) {
		g.Var().Err().Error()

		write(g, "{")

		if s.IsResource {
			write(g, `"resourceType":"`+s.Name+`"`)
			g.Id("setComma").Op(":=").True()
		} else {
			g.Id("setComma").Op(":=").False()
		}

		for _, f := range s.Fields {
			if f.Polymorph {
				g.Switch(Id("v").Op(":=").Id("r").Op(".").Id(f.Name).Op(".").Params(Type())).BlockFunc(func(g *Group) {
					for _, t := range f.PossibleTypes {
						implementCasePolymorph(g, f, t, false)
						implementCasePolymorph(g, f, t, true)
					}
				})
			} else {
				t := f.PossibleTypes[0]

				if t.IsNestedResource {
					implementNestedResource(g, f)
				} else if !s.IsResource && f.Name == "Id" {
					g.If(Id("r." + f.Name).Op("!=").Nil()).BlockFunc(func(g *Group) {
						writeKey(g, f.MarshalName)
						writePrimitiveValue(g, "r."+f.Name)
					})
				} else if s.Name == "Extension" && f.Name == "Url" {
					g.BlockFunc(func(g *Group) {
						writeKey(g, f.MarshalName)
						writePrimitiveValue(g, "r."+f.Name)
					})
				} else if t.IsPrimitive {
					implementPrimitive(g, f)
				} else {
					implementElement(g, f)
				}
			}
		}

		write(g, "}")

		g.Return(Nil())
	})
}

func implementNestedResource(g *Group, f ir.StructField) {
	if f.Multiple {
		g.If(Len(Id("r." + f.Name)).Op(">").Lit(0)).BlockFunc(func(g *Group) {
			writeKey(g, f.MarshalName)
			write(g, "[")

			g.Id("setComma").Op("=").False()
			g.For(Id("_, c").Op(":=").Range().Id("r." + f.Name)).BlockFunc(func(g *Group) {
				checkWriteComma(g)
				g.Err().Op("=").Id("ContainedResource").Values(Id("c")).
					Dot("marshalJSON").Call(Id("w"))
				g.If(Err().Op("!=").Nil()).Block(
					Return(Err()),
				)
			})

			write(g, "]")
		})
	} else {
		g.If(Id("r." + f.Name).Op("!=").Nil()).BlockFunc(func(g *Group) {
			writeKey(g, f.MarshalName)
			g.Err().Op("=").Id("ContainedResource").Values(Id("*r." + f.Name)).
				Dot("marshalJSON").Call(Id("w"))
			g.If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			)
		})
	}
}

func implementCasePolymorph(g *Group, f ir.StructField, t ir.FieldType, pointer bool) {
	var c *Statement
	if pointer {
		c = Op("*").Id(t.Name)
	} else {
		c = Id(t.Name)
	}
	g.Case(c).BlockFunc(func(g *Group) {
		if t.IsPrimitive {
			g.If(Id("v.Value").Op("!=").Nil()).BlockFunc(func(g *Group) {
				writeKey(g, f.MarshalName+t.Name)
				writePrimitiveValue(g, "v")
			})

			g.If(Id("v.Id").Op("!=").Nil().Op("||").Id("v.Extension").Op("!=").Nil()).BlockFunc(func(g *Group) {
				g.Id("p").Op(":=").Id("primitiveElement").ValuesFunc(func(g *Group) {
					g.Id("Id").Op(":").Id("v.Id")
					if f.Name != "Div" {
						g.Id("Extension").Op(":").Id("v.Extension")
					}
				})

				writeKey(g, "_"+f.MarshalName+t.Name)
				writeElementValue(g, "p")
			})
		} else {
			writeKey(g, f.MarshalName+t.Name)
			writePrimitiveValue(g, "v")
		}
	})
}

func implementPrimitive(g *Group, f ir.StructField) {
	if f.Multiple {
		g.BlockFunc(func(g *Group) {
			g.Id("anyValue").Op(":=").False()
			g.For(Id("_, e").Op(":=").Range().Id("r." + f.Name)).Block(
				If(Id("e.Value").Op("!=").Nil()).Block(
					Id("anyValue").Op("=").True(),
					Break(),
				),
			)
			g.If(Id("anyValue")).BlockFunc(func(g *Group) {
				writeKey(g, f.MarshalName)
				writePrimitiveValue(g, "r."+f.Name)
			})

			g.Id("anyIdOrExtension").Op(":=").False()
			g.For(Id("_, e").Op(":=").Range().Id("r." + f.Name)).Block(
				If(Id("e.Id").Op("!=").Nil().Op("||").Id("e.Extension").Op("!=").Nil()).Block(
					Id("anyIdOrExtension").Op("=").True(),
					Break(),
				),
			)
			g.If(Id("anyIdOrExtension")).BlockFunc(func(g *Group) {
				writeKey(g, "_"+f.MarshalName)
				write(g, "[")
				g.Id("setComma").Op("=").False()
				g.For(Id("_, e").Op(":=").Range().Id("r." + f.Name)).Block(
					If(Id("e.Id").Op("!=").Nil().Op("||").Id("e.Extension").Op("!=").Nil()).BlockFunc(func(g *Group) {
						checkWriteComma(g)
						g.Id("p").Op(":=").Id("primitiveElement").Values(
							Id("Id").Op(":").Id("e.Id"),
							Id("Extension").Op(":").Id("e.Extension"),
						)
						writeElementValue(g, "p")
					}).Else().BlockFunc(func(g *Group) {
						checkWriteComma(g)
						write(g, "null")
					}),
				)
				write(g, "]")
			})
		})
	} else {
		if f.Name == "Div" {
			g.BlockFunc(func(g *Group) {
				writeKey(g, f.MarshalName)
				writePrimitiveValue(g, "r."+f.Name)
			})
		} else if f.Optional {
			g.If(Id("r." + f.Name).Op("!=").Nil().Op("&&").Id("r." + f.Name + ".Value").Op("!=").Nil()).BlockFunc(func(g *Group) {
				writeKey(g, f.MarshalName)
				writePrimitiveValue(g, "r."+f.Name)
			})
		} else {
			g.BlockFunc(func(g *Group) {
				writeKey(g, f.MarshalName)
				writePrimitiveValue(g, "r."+f.Name)
			})
		}

		g.IfFunc(func(g *Group) {
			if f.Optional {
				g.Id("r." + f.Name).Op("!=").Nil().Op("&&").Params(
					Id("r." + f.Name + ".Id").Op("!=").Nil().Op("||").Id("r." + f.Name + ".Extension").Op("!=").Nil(),
				)
			} else {
				c := g.Id("r." + f.Name + ".Id").Op("!=").Nil()
				if f.Name != "Div" {
					c.Op("||").Id("r." + f.Name + ".Extension").Op("!=").Nil()
				}
			}
		}).BlockFunc(func(g *Group) {
			g.Id("p").Op(":=").Id("primitiveElement").ValuesFunc(func(g *Group) {
				g.Id("Id").Op(":").Id("r." + f.Name + ".Id")
				if f.Name != "Div" {
					g.Id("Extension").Op(":").Id("r." + f.Name + ".Extension")
				}
			})

			writeKey(g, "_"+f.MarshalName)
			writeElementValue(g, "p")
		})
	}
}

func implementElement(g *Group, f ir.StructField) {
	if f.Multiple {
		g.If(Len(Id("r." + f.Name)).Op(">").Lit(0)).BlockFunc(func(g *Group) {
			writeKey(g, f.MarshalName)
			write(g, "[")
			g.Id("setComma").Op("=").False()
			g.For(List(Id("_"), Id("e")).Op(":=").Range().Id("r." + f.Name)).BlockFunc(func(g *Group) {
				checkWriteComma(g)
				g.Err().Op("=").Id("e").Dot("marshalJSON").Call(Id("w"))
				g.If(Err().Op("!=").Nil()).Block(
					Return(Err()),
				)
			})
			write(g, "]")
		})
	} else if f.Optional {
		g.If(Id("r." + f.Name).Op("!=").Nil()).BlockFunc(func(g *Group) {
			writeKey(g, f.MarshalName)
			writeElementValue(g, "r."+f.Name)
		})
	} else {
		writeKey(g, f.MarshalName)
		writeElementValue(g, "r."+f.Name)
	}
}

func writeKey(g *Group, key string) {
	checkWriteComma(g)

	g.List(Id("_"), Err()).Op("=").Id("w").Dot("Write").Call(
		Index().Byte().Call(Lit(`"` + key + `":`)))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
}

func writeElementValue(g *Group, value string) {
	g.Err().Op("=").Id(value).Dot("marshalJSON").Call(Id("w"))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
}

func writePrimitiveValue(g *Group, value string) {
	g.Var().Id("b").Qual("bytes", "Buffer")
	g.Id("enc").Op(":=").Qual("encoding/json", "NewEncoder").Call(Id("&b"))
	g.Id("enc").Dot("SetEscapeHTML").Call(False())
	g.Err().Op(":=").Id("enc").Dot("Encode").Call(Id(value))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
	g.List(Id("_"), Err()).Op("=").Id("w").Dot("Write").Call(Id("b").Dot("Bytes").Call())
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
}

func checkWriteComma(g *Group) {
	g.If(Id("setComma")).BlockFunc(func(g *Group) {
		write(g, ",")
	})
	g.Id("setComma").Op("=").True()
}

func write(g *Group, s string) {
	g.List(Id("_"), Err()).Op("=").Id("w").Dot("Write").Call(
		Index().Byte().Call(Lit(s)))
	g.If(Err().Op("!=").Nil()).Block(
		Return(Err()),
	)
}
