package json

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementMarshalContained(resources []ir.Struct, f *File) {
	implementMarshalContainedExternal(f)
	implementMarshalContainedInternal(resources, f)
}

func implementMarshalContainedExternal(f *File) {
	f.Func().Params(Id("r").Id("ContainedResource")).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).Block(
		Var().Id("b").Qual("bytes", "Buffer"),
		Err().Op(":=").Id("r").Dot("marshalJSON").Call(Id("&b")),
		If(Err().Op("!=").Nil()).Block(
			Return(Nil(), Err()),
		),
		Return(Id("b").Dot("Bytes").Call(), Nil()),
	)
}

func implementMarshalContainedInternal(resources []ir.Struct, f *File) {
	f.Func().Params(Id("r").Id("ContainedResource")).Id("marshalJSON").Params(
		Id("w").Qual("io", "Writer"),
	).Params(Error()).Block(
		Switch(Id("t").Op(":=").Id("r").Dot("Resource").Dot("").Call(Type())).BlockFunc(func(g *Group) {
			for _, r := range resources {
				g.Case(Id(r.Name)).Block(
					Return(Id("t").Dot("marshalJSON").Call(Id("w"))),
				)
			}

			g.Default().Block(
				Return(Qual("fmt", "Errorf").Call(Lit("unknown resource: %v"), Id("t"))),
			)
		}),
	)
}

func ImplementUnmarshalContained(resources []ir.Struct, f *File) {
	implementUnmarshalContainedExternal(f)
	implementUnmarshalContainedInternal(resources, f)
}

func implementUnmarshalContainedExternal(f *File) {
	f.Func().Params(Id("r").Op("*").Id("ContainedResource")).Id("UnmarshalJSON").Params(
		Id("b").Index().Byte(),
	).Params(Error()).Block(
		Id("d").Op(":=").Qual("encoding/json", "NewDecoder").Call(
			Qual("bytes", "NewReader").Call(Id("b")),
		),
		Return(Id("r").Dot("unmarshalJSON").Call(Id("d"))),
	)
}

func implementUnmarshalContainedInternal(resources []ir.Struct, f *File) {
	f.Func().Params(Id("cr").Op("*").Id("ContainedResource")).Id("unmarshalJSON").Params(
		Id("d").Op("*").Qual("encoding/json", "Decoder"),
	).Params(Error()).Block(
		Var().Id("rawValue").Qual("encoding/json", "RawMessage"),
		Id("err").Op(":=").Id("d").Dot("Decode").Call(Id("&rawValue")),
		If(Err().Op("!=").Nil()).Block(
			Return(Id("err")),
		),

		Var().Id("t").Struct(
			Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
		),
		Err().Op("=").Qual("encoding/json", "Unmarshal").Call(Id("rawValue"), Op("&").Id("t")),
		If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		),

		Id("d").Op("=").Qual("encoding/json", "NewDecoder").
			Call(Qual("bytes", "NewReader").Call(Id("rawValue"))),

		Switch(Id("t.ResourceType")).BlockFunc(func(g *Group) {
			for _, r := range resources {
				g.Case(Lit(r.Name)).Block(
					Var().Id("r").Id(r.Name),
					Id("err").Op(":=").Id("r").Dot("unmarshalJSON").Call(Id("d")),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					Op("*").Id("cr").Op("=").Id("ContainedResource").Values(Id("r")),
					Return(Nil()),
				)
			}

			g.Default().Block(
				Return(Qual("fmt", "Errorf").Call(Lit("unknown resource type: %s"), Id("t.ResourceType"))),
			)
		}),
	)
}

func ImplementMarshalPrimitiveElement(f *File) {
	f.Func().Params(Id("r").Id("primitiveElement")).Id("marshalJSON").Params(
		Id("w").Qual("io", "Writer"),
	).Params(Error()).BlockFunc(func(g *Group) {
		g.Var().Err().Error()
		write(g, "{")
		g.Id("setComma").Op(":=").False()
		g.If(Id("r.Id").Op("!=").Nil()).BlockFunc(func(g *Group) {
			writeKey(g, "id")
			writePrimitiveValue(g, "r.Id")
		})
		g.If(Len(Id("r.Extension")).Op(">").Lit(0)).BlockFunc(func(g *Group) {
			writeKey(g, "extension")
			write(g, "[")
			g.Id("setComma").Op("=").False()
			g.For(List(Id("_"), Id("e")).Op(":=").Range().Id("r.Extension")).BlockFunc(func(g *Group) {
				checkWriteComma(g)
				g.Err().Op("=").Id("e").Dot("marshalJSON").Call(Id("w"))
				g.If(Err().Op("!=").Nil()).Block(
					Return(Err()),
				)
			})
			write(g, "]")
		})
		write(g, "}")
		g.Return(Nil())
	})
}

func ImplementUnmarshalPrimitiveElement(f *File) {
	f.Func().Params(Id("r").Op("*").Id("primitiveElement")).Id("unmarshalJSON").Params(
		Id("d").Op("*").Qual("encoding/json", "Decoder"),
	).Params(Error()).BlockFunc(func(g *Group) {
		g.List(Id("t"), Err()).Op(":=").Id("d").Dot("Token").Call()
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
		g.If(Id("t").Op("==").Nil()).Block(
			Return(Nil()),
		).Else().If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('{')).Block(
			returnInvalidTokenError("primitive element", "'{'"),
		)

		g.For(Id("d").Dot("More").Call()).Block(
			List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call(),
			If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			),
			List(Id("f"), Id("ok")).Op(":=").Id("t").Op(".").Call(String()),
			If(Op("!").Id("ok")).Block(
				returnInvalidTokenError("primitive element", "field name"),
			),
			Switch(Id("f")).BlockFunc(func(g *Group) {
				g.Case(Lit("id")).Block(
					Var().Id("v").String(),
					Id("err").Op(":=").Id("d").Dot("Decode").Call(Id("&v")),
					If(Err().Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					Id("r.Id").Op("=").Id("&v"),
				)
				g.Case(Lit("extension")).Block(
					List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call(),
					If(Err().Op("!=").Nil()).Block(
						Return(Err()),
					),
					If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('[')).Block(
						returnInvalidTokenError("primitive element", "'['"),
					),
					For(Id("d").Dot("More").Call()).Block(
						Var().Id("v").Id("Extension"),
						Id("err").Op(":=").Id("v").Dot("unmarshalJSON").Call(Id("d")),
						If(Err().Op("!=").Nil()).Block(
							Return(Id("err")),
						),
						Id("r.Extension").Op("=").Append(Id("r.Extension"), Id("v")),
					),
					List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call(),
					If(Err().Op("!=").Nil()).Block(
						Return(Err()),
					),
					If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune(']')).Block(
						returnInvalidTokenError("primitive element", "']'"),
					),
				)

				g.Default().Block(
					Return(Qual("fmt", "Errorf").Params(
						Lit("invalid field: %v in primitive element, expected \"id\" or \"extension\" (at index %v)"),
						Id("t"),
						Id("d").Dot("InputOffset").Call().Op("-").Lit(1),
					)),
				)
			}),
		)

		g.List(Id("t"), Err()).Op("=").Id("d").Dot("Token").Call()
		g.If(Err().Op("!=").Nil()).Block(
			Return(Err()),
		)
		g.If(Id("t").Op("!=").Qual("encoding/json", "Delim")).Params(LitRune('}')).Block(
			returnInvalidTokenError("primitive element", "'}'"),
		)
		g.Return(Nil())
	})
}
