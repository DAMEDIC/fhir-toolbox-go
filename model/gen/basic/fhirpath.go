package basic

import (
	"context"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
)

func Context() context.Context {
	return WithContext(context.Background())
}
func WithContext(ctx context.Context) context.Context {
	ctx = fhirpath.WithNamespace(ctx, "FHIR")
	ctx = fhirpath.WithTypes(ctx, allFHIRPathTypes)
	ctx = fhirpath.WithFunctions(ctx, fhirpath.FHIRFunctions)
	return ctx
}

var allFHIRPathTypes = []fhirpath.TypeInfo{fhirpath.ClassInfo{SimpleTypeInfo: fhirpath.SimpleTypeInfo{
	BaseType: fhirpath.TypeSpecifier{
		Name:      "Any",
		Namespace: "System",
	},
	Name:      "Base",
	Namespace: "FHIR",
}}, fhirpath.ClassInfo{SimpleTypeInfo: fhirpath.SimpleTypeInfo{
	BaseType: fhirpath.TypeSpecifier{
		Name:      "Base",
		Namespace: "FHIR",
	},
	Name:      "Element",
	Namespace: "FHIR",
}}, fhirpath.ClassInfo{SimpleTypeInfo: fhirpath.SimpleTypeInfo{
	BaseType: fhirpath.TypeSpecifier{
		Name:      "Element",
		Namespace: "FHIR",
	},
	Name:      "DataType",
	Namespace: "FHIR",
}}, fhirpath.ClassInfo{SimpleTypeInfo: fhirpath.SimpleTypeInfo{
	BaseType: fhirpath.TypeSpecifier{
		Name:      "DataType",
		Namespace: "FHIR",
	},
	Name:      "PrimitiveType",
	Namespace: "FHIR",
}}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "modifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "Element",
			Namespace: "FHIR",
		},
		Name:      "BackboneElement",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "modifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "BackboneType",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "id",
		Type: fhirpath.TypeSpecifier{
			Name:      "id",
			Namespace: "FHIR",
		},
	}, {
		Name: "meta",
		Type: fhirpath.TypeSpecifier{
			Name:      "Meta",
			Namespace: "FHIR",
		},
	}, {
		Name: "implicitRules",
		Type: fhirpath.TypeSpecifier{
			Name:      "uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "language",
		Type: fhirpath.TypeSpecifier{
			Name:      "code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "Base",
			Namespace: "FHIR",
		},
		Name:      "Resource",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "id",
		Type: fhirpath.TypeSpecifier{
			Name:      "id",
			Namespace: "FHIR",
		},
	}, {
		Name: "meta",
		Type: fhirpath.TypeSpecifier{
			Name:      "Meta",
			Namespace: "FHIR",
		},
	}, {
		Name: "implicitRules",
		Type: fhirpath.TypeSpecifier{
			Name:      "uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "language",
		Type: fhirpath.TypeSpecifier{
			Name:      "code",
			Namespace: "FHIR",
		},
	}, {
		Name: "text",
		Type: fhirpath.TypeSpecifier{
			Name:      "Narrative",
			Namespace: "FHIR",
		},
	}, {
		Name: "contained",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Resource",
			Namespace: "FHIR",
		},
	}, {
		Name: "extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "modifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "Resource",
			Namespace: "FHIR",
		},
		Name:      "DomainResource",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "base64Binary",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "boolean",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "canonical",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "code",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "date",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "dateTime",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "decimal",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "id",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "instant",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "integer",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "markdown",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "oid",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "positiveInt",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "string",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "time",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "unsignedInt",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "uri",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "url",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "uuid",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "xhtml",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Use",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Line",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "City",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "District",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "State",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PostalCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Country",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Address",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Age",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Time",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Annotation",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ContentType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Language",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Data",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Base64Binary",
			Namespace: "FHIR",
		},
	}, {
		Name: "Url",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Url",
			Namespace: "FHIR",
		},
	}, {
		Name: "Size",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Hash",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Base64Binary",
			Namespace: "FHIR",
		},
	}, {
		Name: "Title",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Creation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Attachment",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Coding",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "CodeableConcept",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Version",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Display",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "UserSelected",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Coding",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ContactDetail",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Use",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Rank",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ContactPoint",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Contributor",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Count",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Profile",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "MustSupport",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "CodeFilter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirementCodeFilter",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateFilter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirementDateFilter",
			Namespace: "FHIR",
		},
	}, {
		Name: "Limit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sort",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirementSort",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "DataRequirement",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Distance",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sequence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdditionalInstruction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PatientInstruction",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Timing",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Timing",
			Namespace: "FHIR",
		},
	}, {
		Name: "AsNeeded",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Site",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Route",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Method",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoseAndRate",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DosageDoseAndRate",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaxDosePerPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Ratio",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaxDosePerAdministration",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaxDosePerLifetime",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Dosage",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Duration",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Path",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Representation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "SliceName",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "SliceIsConstraining",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Label",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Slicing",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ElementDefinitionSlicing",
			Namespace: "FHIR",
		},
	}, {
		Name: "Short",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Definition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requirements",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Alias",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Min",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Max",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Base",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ElementDefinitionBase",
			Namespace: "FHIR",
		},
	}, {
		Name: "ContentReference",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ElementDefinitionType",
			Namespace: "FHIR",
		},
	}, {
		Name: "DefaultValue",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "MeaningWhenMissing",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "OrderMeaning",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Fixed",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Pattern",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Example",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ElementDefinitionExample",
			Namespace: "FHIR",
		},
	}, {
		Name: "MinValue",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaxValue",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaxLength",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "Constraint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ElementDefinitionConstraint",
			Namespace: "FHIR",
		},
	}, {
		Name: "MustSupport",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "IsModifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "IsModifierReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "IsSummary",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Binding",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ElementDefinitionBinding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Mapping",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ElementDefinitionMapping",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ElementDefinition",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Description",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "Language",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Expression",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reference",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Expression",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Url",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Extension",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Use",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Family",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Given",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Prefix",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Suffix",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "HumanName",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Use",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Assigner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Identifier",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Country",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Jurisdiction",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateRange",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "RestoreDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "MarketingStatus",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "VersionId",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastUpdated",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Profile",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Security",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Tag",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Meta",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Currency",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Money",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Div",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Xhtml",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Narrative",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Use",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Min",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Max",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Documentation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Profile",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ParameterDefinition",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Start",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "End",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Period",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Age",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Race",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PhysiologicalCondition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Population",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Height",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Width",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Depth",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Weight",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "NominalVolume",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExternalDiameter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Shape",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Color",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Imprint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Image",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Scoring",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ProdCharacteristic",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "SpecialPrecautionsForStorage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "ProductShelfLife",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Quantity",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Low",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "High",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Range",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Numerator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Denominator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Ratio",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reference",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Display",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Reference",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Label",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Display",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Citation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Document",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Resource",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "RelatedArtifact",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Origin",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Factor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "LowerLimit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "UpperLimit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Dimensions",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Data",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "SampledData",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "When",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Who",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OnBehalfOf",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "TargetFormat",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "SigFormat",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Data",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Base64Binary",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Signature",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Amount",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "AmountType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AmountText",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferenceRange",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SubstanceAmountReferenceRange",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "SubstanceAmount",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Event",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Repeat",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TimingRepeat",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "Timing",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Timing",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Data",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Expression",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "TriggerDefinition",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "UsageContext",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "MoneyQuantity",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "string",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
		Name:      "SimpleQuantity",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "Meta",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Meta",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImplicitRules",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Language",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Timestamp",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Total",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Link",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "BundleLink",
			Namespace: "FHIR",
		},
	}, {
		Name: "Entry",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "BundleEntry",
			Namespace: "FHIR",
		},
	}, {
		Name: "Signature",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Signature",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Bundle",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "Meta",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Meta",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImplicitRules",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Language",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Narrative",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contained",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Url",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Version",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Title",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Experimental",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Date",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Publisher",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Description",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "UseContext",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "UsageContext",
			Namespace: "FHIR",
		},
	}, {
		Name: "Jurisdiction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Purpose",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Kind",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instantiates",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Imports",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Software",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CapabilityStatementSoftware",
			Namespace: "FHIR",
		},
	}, {
		Name: "Implementation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CapabilityStatementImplementation",
			Namespace: "FHIR",
		},
	}, {
		Name: "FhirVersion",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Format",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "PatchFormat",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImplementationGuide",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Rest",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CapabilityStatementRest",
			Namespace: "FHIR",
		},
	}, {
		Name: "Messaging",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CapabilityStatementMessaging",
			Namespace: "FHIR",
		},
	}, {
		Name: "Document",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CapabilityStatementDocument",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CapabilityStatement",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "Meta",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Meta",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImplicitRules",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Language",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Text",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Narrative",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contained",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "",
			Namespace: "FHIR",
		},
	}, {
		Name: "Extension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModifierExtension",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Extension",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issue",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "OperationOutcomeIssue",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "OperationOutcome",
		Namespace: "FHIR",
	},
}}
var fhirFunctions = fhirpath.Functions{}
