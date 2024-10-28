package xml

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementMarshalContained(f *File) {
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

func ImplementUnmarshalContained(resources []ir.Struct, f *File) {
	f.Func().Params(Id("cr").Op("*").Id("ContainedResource")).Id("UnmarshalXML").Params(
		Id("d").Op("*").Qual("encoding/xml", "Decoder"),
		Id("start").Qual("encoding/xml", "StartElement"),
	).Params(Error()).Block(
		// if name is lower means we are dealing with a contained resource
		If(Qual("unicode", "IsLower").Call(Rune().Call(Id("start.Name.Local").Index(Lit(0))))).Block(
			Err().Op(":=").Id("d.Decode").Call(Op("cr")),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Id("err")),
			),
			For().Block(
				Id("t, err").Op(":=").Id("d.Token()"),
				If(Err().Op("!=").Nil()).Block(
					Return(Id("err")),
				),
				Id("_, ok").Op(":=").Id("t.").Params(Qual("encoding/xml", "EndElement")),
				If(Id("ok")).Block(Break()),
			),
			Return(Nil()),
		),

		If(Id("start.Name.Space").Op("!=").Lit(NamespaceFHIR)).Block(
			Return(Qual("fmt", "Errorf").Params(
				Lit("invalid namespace: \"%s\", expected: \""+NamespaceFHIR+"\""),
				Id("start.Name.Space"),
			)),
		),

		Switch(Id("start.Name.Local")).BlockFunc(func(g *Group) {
			for _, r := range resources {
				g.Case(Lit(r.Name)).Block(
					Var().Id("r").Id(r.Name),
					Id("err").Op(":=").Id("d.DecodeElement").Call(Op("&r"), Id("&start")),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					Op("*").Id("cr").Op("=").Id("ContainedResource").Values(Id("r")),
					Return(Nil()),
				)
			}

			g.Default().Block(
				Return(Qual("fmt", "Errorf").Call(Lit("unknown resource type: %s"), Id("start.Name.Local"))),
			)
		}),
	)
}
