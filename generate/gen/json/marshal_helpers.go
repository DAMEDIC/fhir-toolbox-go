package json

import (
	"fhir-toolbox/generate/ir"

	. "github.com/dave/jennifer/jen"
)

func ImplementMarshalContained(f *File) {
	f.Func().Params(Id("r").Id("ContainedResource")).Id("MarshalJSON").Params().Params(Index().Byte(), Error()).BlockFunc(func(g *Group) {
		g.Return(Qual("encoding/json", "Marshal").Params(Id("r.Resource")))
	})
}

func ImplementUnmarshalContained(resources []ir.Struct, f *File) {
	f.Func().Params(Id("cr").Op("*").Id("ContainedResource")).Id("UnmarshalJSON").Params(Id("b").Index().Byte()).Params(Error()).Block(
		Var().Id("t").Struct(
			Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
		),
		If(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("t")).Op("!=").Nil()).Block(
			Return(Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&").Id("t"))),
		),
		Switch(Id("t.ResourceType")).BlockFunc(func(g *Group) {
			for _, r := range resources {
				g.Case(Lit(r.Name)).Block(
					Var().Id("r").Id(r.Name),
					Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(Id("b"), Op("&r")),
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
