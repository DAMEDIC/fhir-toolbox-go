package ir

import "github.com/iancoleman/strcase"

func toGoFileCasing(s string) string {
	return strcase.ToSnake(s)
}

func toGoTypeCasing(s string) string {
	return strcase.ToCamel(s)
}

func toGoFieldCasing(s string) string {
	return strcase.ToCamel(s)
}
