package ir

import (
	"fhir-toolbox/codegen/model"
	"fmt"
	"slices"
	"strings"
)

var primitives = []string{
	"base64Binary",
	"boolean",
	"canonical",
	"code",
	"date",
	"dateTime",
	"decimal",
	"id",
	"instant",
	"integer",
	"integer64",
	"markdown",
	"oid",
	"positiveInt",
	"string",
	"time",
	"unsignedInt",
	"uri",
	"url",
	"uuid",
	"xhtml",
}

func Parse(bundle *model.Bundle) []SourceFile {
	var sourceFiles []SourceFile

	for _, s := range flattenBundle(bundle) {
		if s.Kind == "logical" {
			continue
		}
		if s.Abstract {
			continue
		}
		if s.Name == "Element" {
			continue
		}

		sourceFile := SourceFile{
			Name: toGoFileCasing(s.Name),
		}

		parseStruct(
			&sourceFile.Structs,
			s.Name,
			s.Kind == "resource",
			s.Snapshot.Element,
			s.Type,
			fmt.Sprintf("%s\n\n%s", s.Description, s.Purpose),
		)

		sourceFiles = append(sourceFiles, sourceFile)
	}

	return sourceFiles
}

func flattenBundle(bundle *model.Bundle) []*model.StructureDefinition {
	var definitions []*model.StructureDefinition

	for _, e := range bundle.Entry {
		if sd, ok := e.Resource.(*model.StructureDefinition); ok {
			definitions = append(definitions, sd)
		}
	}

	return definitions
}

func parseStruct(
	into *[]Struct,
	name string,
	isResource bool,
	elementDefinitions []model.ElementDefinition,
	elementPathStripPrefix string,
	docComment string,
) {
	structName := toGoTypeCasing(name)

	groupedDefinitions := groupElementDefinitionsByPrefix(elementDefinitions, elementPathStripPrefix)

	var fields []StructField
	for fieldName, definitions := range groupedDefinitions {
		if definitions[0].Max == "0" {
			continue
		}

		typeName := structName + toGoTypeCasing(fieldName)

		if len(definitions) > 1 {
			parseStruct(
				into,
				typeName,
				false,
				definitions,
				definitions[0].Path,
				definitions[0].Definition,
			)
		}

		fields = append(fields, parseField(
			structName,
			isResource,
			definitions[0],
			elementPathStripPrefix,
		))
	}

	*into = append(*into, Struct{
		Name:        structName,
		IsPrimitive: slices.Contains(primitives, name),
		Fields:      fields,
		DocComment:  docComment,
	})
}

func groupElementDefinitionsByPrefix(elementDefinitions []model.ElementDefinition, stripPrefix string) map[string][]model.ElementDefinition {
	grouped := make(map[string][]model.ElementDefinition)

	for _, d := range elementDefinitions {
		if d.Path == stripPrefix {
			continue
		}

		fieldName := strings.SplitN(d.Path[len(stripPrefix)+1:], ".", 2)[0]

		if _, exists := grouped[fieldName]; exists {
			grouped[fieldName] = append(grouped[fieldName], d)
		} else {
			grouped[fieldName] = []model.ElementDefinition{d}
		}
	}

	return grouped
}

func parseField(
	structName string,
	isResource bool,
	elementDefinition model.ElementDefinition,
	elementPathStripPrefix string,
) StructField {
	fieldName := elementDefinition.Path[len(elementPathStripPrefix)+1:]
	fieldName, polymorph := strings.CutSuffix(fieldName, "[x]")

	var fieldType FieldType
	if polymorph {
		fieldType.Name = "Element"
		fieldType.Package = ".."
	} else if isResource && fieldName == "id" {
		fieldType.Name = "Id"
		fieldType.Package = "types"
	} else if elementDefinition.Type != nil {
		code := (elementDefinition.Type)[0].Code

		switch code {
		case "BackboneElement", "Element":
			fieldType.Name = structName + toGoTypeCasing(fieldName)
		default:
			fieldType = matchFieldType(elementDefinition.Path, code)
		}
	} else {
		// content reference
		// strip "#"
		fieldType.Name = toGoTypeCasing(elementDefinition.ContentReference[1:])
	}

	return StructField{
		Name:        toGoFieldCasing(fieldName),
		MarshalName: fieldName,
		Type:        fieldType,
		Polymorph:   polymorph,
		Multiple:    elementDefinition.Max == "*",
		Optional:    elementDefinition.Min == 0,
		DocComment:  elementDefinition.Definition,
	}
}

func matchFieldType(path string, code string) FieldType {
	var fieldType FieldType

	// type like http://hl7.org/fhirpath/System.String
	switch t := code[strings.LastIndex(code, "/")+1:]; t {
	case "System.Boolean":
		fieldType.Name = "bool"
	case "System.Integer":
		switch path {
		case "integer64.value":
			fieldType.Name = "int64"
		default:
			fieldType.Name = "int32"
		}
	case "System.String":
		switch path {
		case "unsignedInt.value", "positiveInt.value":
			fieldType.Name = "uint32"
		default:
			fieldType.Name = "string"
		}
	case "System.Decimal":
		fieldType.Name = "string"
	case "System.Date", "System.DateTime", "System.Time":
		fieldType.Name = "string"
	case "Resource":
		fieldType.Name = "Resource"
		fieldType.Package = ".."
	default:
		fieldType.Name = toGoTypeCasing(t)
		fieldType.Package = "types"
	}

	return fieldType
}
