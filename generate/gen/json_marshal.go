package gen

import (
	"fhir-toolbox/generate/ir"
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func generateMarshalJSONStruct(f *File, s ir.Struct) {
	f.Type().Id("json" + s.Name).StructFunc(func(g *Group) {
		if s.IsResource {
			g.Id("ResourceType").Id("string").Tag(map[string]string{"json": "resourceType"})
		}

		for _, sf := range s.Fields {
			if sf.Polymorph {
				for _, t := range sf.PossibleTypes {
					g.Id(sf.Name + t.Name).Op("*").Id(t.Name).
						Tag(map[string]string{"json": sf.MarshalName + t.Name + ",omitempty"})

					if t.IsPrimitive {
						g.Id(sf.Name + t.Name + "PrimitiveElement").Id("*primitiveElement").Tag(map[string]string{"json": "_" + sf.MarshalName + t.Name + ",omitempty"})
					}
				}
			} else {
				stmt := g.Id(sf.Name)

				if sf.Multiple {
					stmt.Index()
				} else if sf.Optional && !sf.Polymorph {
					stmt.Op("*")
				}

				t := sf.PossibleTypes[0]

				if t.IsNestedResource {
					stmt.Id("ContainedResource")
				} else {
					stmt.Id(t.Name)
				}

				stmt.Tag(map[string]string{
					"json": fmt.Sprintf("%s,omitempty", sf.MarshalName),
				})

				if t.IsPrimitive {
					stmt := g.Id(sf.Name + "PrimitiveElement")
					if sf.Multiple {
						stmt.Index()
					}
					stmt.Id("*primitiveElement").Tag(map[string]string{"json": "_" + sf.MarshalName + ",omitempty"})
				}
			}
		}
	})
}

func implementMarshalJSON(f *File, s ir.Struct) {
	if s.IsPrimitive {
		implementMarshalJSONPrimitive(f, s)
	} else {
		implementMarshalJSONElement(f, s)
		implementMarshalJSONStruct(f, s)
	}
}

func implementMarshalJSONPrimitive(f *File, s ir.Struct) {
	var ret Code
	if s.Name == "Decimal" {
		ret = Return(Index().Byte().Params(Op("*").Id("r").Op(".").Id("Value")), Nil())
	} else {
		ret = Return(Qual("encoding/json", "Marshal").Params(Id("r").Op(".").Id("Value")))
	}

	f.Func().Params(Id("r").Id(s.Name)).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).Block(ret)
}

func implementMarshalJSONElement(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).BlockFunc(func(g *Group) {
		g.Return(Qual("encoding/json", "Marshal").Params(Id("r.marshalJSON").Params()))
	})
}

func implementMarshalJSONStruct(f *File, s ir.Struct) {
	f.Func().Params(Id("r").Id(s.Name)).Id("marshalJSON").Params().Params(Id("json" + s.Name)).BlockFunc(func(g *Group) {
		g.Id("m").Op(":=").Id("json" + s.Name).Values()

		if s.IsResource {
			g.Id("m.ResourceType").Op("=").Lit(s.Name)
		}

		for _, sf := range s.Fields {
			if sf.Polymorph {
				g.Switch(Id("v").Op(":=").Id("r").Op(".").Id(sf.Name).Op(".").Params(Type())).BlockFunc(func(g *Group) {
					for _, t := range sf.PossibleTypes {
						implementMarshalCase(g, sf, t, false)
						implementMarshalCase(g, sf, t, true)
					}
				})
			} else {
				t := sf.PossibleTypes[0]

				if t.IsNestedResource {
					if sf.Multiple {
						g.Id("m."+sf.Name).Op("=").Id("make").Params(Id("[]ContainedResource"), Lit(0), Len(Id("r."+sf.Name)))
						g.For(Id("_, c").Op(":=").Range().Id("r." + sf.Name)).Block(
							Id("m."+sf.Name).Op("=").Append(Id("m."+sf.Name), Id("ContainedResource").Values(Id("c"))),
						)
					} else {
						g.If(Id("r." + sf.Name).Op("!=").Nil()).Block(
							Id("m." + sf.Name).Op("=").Id("&ContainedResource").Values(Id("*r." + sf.Name)),
						)
					}
				} else {
					g.Id("m." + sf.Name).Op("=").Id("r." + sf.Name)
				}

				if t.IsPrimitive {
					if sf.Multiple {
						g.Id("any" + sf.Name + "IdOrExtension").Op(":=").False()
						g.For(Id("_, e").Op(":=").Range().Id("r." + sf.Name)).Block(
							If(Id("e.Id").Op("!=").Nil().Op("||").Id("e.Extension").Op("!=").Nil()).Block(
								Id("any"+sf.Name+"IdOrExtension").Op("=").True(),
								Break(),
							),
						)
						g.If(Id("any"+sf.Name+"IdOrExtension")).Block(
							Id("m."+sf.Name+"PrimitiveElement").Op("=").Id("make").Params(Id("[]*primitiveElement"), Lit(0), Len(Id("r."+sf.Name))),
							For(Id("_, e").Op(":=").Range().Id("r."+sf.Name)).Block(
								If(Id("e.Id").Op("!=").Nil().Op("||").Id("e.Extension").Op("!=").Nil()).Block(
									Id("m."+sf.Name+"PrimitiveElement").Op("=").Append(Id("m."+sf.Name+"PrimitiveElement"), Id("&primitiveElement").Values(
										Id("Id").Op(":").Id("e.Id"),
										Id("Extension").Op(":").Id("e.Extension"),
									)),
								).Else().Block(
									Id("m."+sf.Name+"PrimitiveElement").Op("=").Append(Id("m."+sf.Name+"PrimitiveElement"), Nil()),
								),
							),
						)
					} else {
						g.IfFunc(func(g *Group) {
							if sf.Optional {
								g.Id("r." + sf.Name).Op("!=").Nil().Op("&&").Params(
									Id("r." + sf.Name + ".Id").Op("!=").Nil().Op("||").Id("r." + sf.Name + ".Extension").Op("!=").Nil(),
								)
							} else {
								c := g.Id("r." + sf.Name + ".Id").Op("!=").Nil()
								if sf.Name != "Div" {
									c.Op("||").Id("r." + sf.Name + ".Extension").Op("!=").Nil()
								}
							}
						}).Block(
							Id("m." + sf.Name + "PrimitiveElement").Op("=").Id("&primitiveElement").ValuesFunc(func(g *Group) {
								g.Id("Id").Op(":").Id("r." + sf.Name + ".Id")
								if sf.Name != "Div" {
									g.Id("Extension").Op(":").Id("r." + sf.Name + ".Extension")
								}
							}),
						)
					}
				}
			}
		}

		g.Return(Id("m"))
	})
}

func implementMarshalCase(g *Group, sf ir.StructField, t ir.FieldType, pointer bool) {
	var c *Statement
	if pointer {
		c = Op("*").Id(t.Name)
	} else {
		c = Id(t.Name)
	}
	g.Case(c).BlockFunc(func(g *Group) {
		a := g.Id("m." + sf.Name + t.Name).Op("=")

		if pointer {
			a.Id("v")
		} else {
			a.Id("&v")
		}

		if t.IsPrimitive {
			g.If(Id("v.Id").Op("!=").Nil().Op("||").Id("v.Extension").Op("!=").Nil()).Block(
				Id("m."+sf.Name+t.Name+"PrimitiveElement").Op("=").Id("&primitiveElement").Values(
					Id("Id").Op(":").Id("v.Id"),
					Id("Extension").Op(":").Id("v.Extension"),
				),
			)
		}
	})
}
