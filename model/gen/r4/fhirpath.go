package r4

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
		Name: "Url",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Url",
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
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServicePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Coverage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AccountCoverage",
			Namespace: "FHIR",
		},
	}, {
		Name: "Owner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Guarantor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AccountGuarantor",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Account",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Subtitle",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Library",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "Profile",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoNotPerform",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ActivityDefinitionParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Product",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Dosage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Dosage",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SpecimenRequirement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ObservationRequirement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ObservationResultRequirement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Transform",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "DynamicValue",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ActivityDefinitionDynamicValue",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ActivityDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Actuality",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Event",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Detected",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "RecordedDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "ResultingCondition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Seriousness",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Severity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contributor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SuspectEntity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AdverseEventSuspectEntity",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubjectMedicalHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferenceDocument",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Study",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "AdverseEvent",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClinicalStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "VerificationStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Criticality",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Onset",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "RecordedDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Asserter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastOccurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reaction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AllergyIntoleranceReaction",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "AllergyIntolerance",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "CancelationReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceCategory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AppointmentType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
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
		Name: "SupportingInformation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Start",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "End",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "MinutesDuration",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Slot",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AppointmentParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "RequestedPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Appointment",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Appointment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Start",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "End",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParticipantType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Actor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParticipantStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "AppointmentResponse",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subtype",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Action",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
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
		Name: "Recorded",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "OutcomeDesc",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PurposeOfEvent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Agent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AuditEventAgent",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "AuditEventSource",
			Namespace: "FHIR",
		},
	}, {
		Name: "Entity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "AuditEventEntity",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "AuditEvent",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Basic",
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
		Name: "ContentType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "SecurityContext",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Binary",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProductCategory",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProductCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Collection",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "BiologicallyDerivedProductCollection",
			Namespace: "FHIR",
		},
	}, {
		Name: "Processing",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "BiologicallyDerivedProductProcessing",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manipulation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "BiologicallyDerivedProductManipulation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Storage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "BiologicallyDerivedProductStorage",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "BiologicallyDerivedProduct",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Morphology",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "LocationQualifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Image",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "BodyStructure",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Description",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contributor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CareTeam",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Addresses",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Goal",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Activity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CarePlanActivity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CarePlan",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CareTeamParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CareTeam",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Orderable",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferencedItem",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdditionalIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Classification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "ValidityPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidTo",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastUpdated",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdditionalCharacteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdditionalClassification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedEntry",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CatalogEntryRelatedEntry",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CatalogEntry",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "DefinitionUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "DefinitionCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ChargeItemPerformer",
			Namespace: "FHIR",
		},
	}, {
		Name: "PerformingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "RequestingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CostCenter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Bodysite",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FactorOverride",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "PriceOverride",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}, {
		Name: "OverrideReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Enterer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "EnteredDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Service",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Product",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Account",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInformation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ChargeItem",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Title",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "DerivedFromUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Applicability",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ChargeItemDefinitionApplicability",
			Namespace: "FHIR",
		},
	}, {
		Name: "PropertyGroup",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ChargeItemDefinitionPropertyGroup",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ChargeItemDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BillablePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Enterer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FundsReserve",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Related",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimRelated",
			Namespace: "FHIR",
		},
	}, {
		Name: "Prescription",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OriginalPrescription",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payee",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ClaimPayee",
			Namespace: "FHIR",
		},
	}, {
		Name: "Referral",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Facility",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CareTeam",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimCareTeam",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimSupportingInfo",
			Namespace: "FHIR",
		},
	}, {
		Name: "Diagnosis",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimDiagnosis",
			Namespace: "FHIR",
		},
	}, {
		Name: "Procedure",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimProcedure",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimInsurance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Accident",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ClaimAccident",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "Total",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Claim",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requestor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disposition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreAuthRef",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreAuthPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "PayeeType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "AddItem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseAddItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "Adjudication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseItemAdjudication",
			Namespace: "FHIR",
		},
	}, {
		Name: "Total",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseTotal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ClaimResponsePayment",
			Namespace: "FHIR",
		},
	}, {
		Name: "FundsReserve",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FormCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Form",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProcessNote",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseProcessNote",
			Namespace: "FHIR",
		},
	}, {
		Name: "CommunicationRequest",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseInsurance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Error",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClaimResponseError",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ClaimResponse",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Effective",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Assessor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Previous",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Problem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Investigation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClinicalImpressionInvestigation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Protocol",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Summary",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Finding",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ClinicalImpressionFinding",
			Namespace: "FHIR",
		},
	}, {
		Name: "PrognosisCodeableConcept",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PrognosisReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ClinicalImpression",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "CaseSensitive",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValueSet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "HierarchyMeaning",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Compositional",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "VersionNeeded",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Content",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Supplements",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Count",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Filter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeSystemFilter",
			Namespace: "FHIR",
		},
	}, {
		Name: "Property",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeSystemProperty",
			Namespace: "FHIR",
		},
	}, {
		Name: "Concept",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeSystemConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CodeSystem",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "InResponseTo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medium",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "About",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Received",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recipient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payload",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CommunicationPayload",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Communication",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "GroupIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoNotPerform",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medium",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "About",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payload",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CommunicationRequestPayload",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recipient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CommunicationRequest",
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
		Name: "Purpose",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
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
		Name: "Search",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Resource",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CompartmentDefinitionResource",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CompartmentDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Confidentiality",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Attester",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CompositionAttester",
			Namespace: "FHIR",
		},
	}, {
		Name: "Custodian",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatesTo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CompositionRelatesTo",
			Namespace: "FHIR",
		},
	}, {
		Name: "Event",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CompositionEvent",
			Namespace: "FHIR",
		},
	}, {
		Name: "Section",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CompositionSection",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Composition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Group",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ConceptMapGroup",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ConceptMap",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClinicalStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "VerificationStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Severity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Onset",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Abatement",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "RecordedDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Asserter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Stage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ConditionStage",
			Namespace: "FHIR",
		},
	}, {
		Name: "Evidence",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ConditionEvidence",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Condition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Scope",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateTime",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Organization",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Policy",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ConsentPolicy",
			Namespace: "FHIR",
		},
	}, {
		Name: "PolicyRule",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Verification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ConsentVerification",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provision",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ConsentProvision",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Consent",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "LegalState",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "ContentDerivative",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issued",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Applies",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExpirationType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Authority",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Domain",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Site",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Subtitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Scope",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "SubType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ContentDefinition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ContractContentDefinition",
			Namespace: "FHIR",
		},
	}, {
		Name: "Term",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContractTerm",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelevantHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Signer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContractSigner",
			Namespace: "FHIR",
		},
	}, {
		Name: "Friendly",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContractFriendly",
			Namespace: "FHIR",
		},
	}, {
		Name: "Legal",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContractLegal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Rule",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContractRule",
			Namespace: "FHIR",
		},
	}, {
		Name: "LegallyBinding",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Contract",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PolicyHolder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subscriber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubscriberId",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Beneficiary",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Dependent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Relationship",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Payor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Class",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageClass",
			Namespace: "FHIR",
		},
	}, {
		Name: "Order",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Network",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "CostToBeneficiary",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageCostToBeneficiary",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subrogation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contract",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Coverage",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Purpose",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Serviced",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Enterer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Facility",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageEligibilityRequestSupportingInfo",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageEligibilityRequestInsurance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageEligibilityRequestItem",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CoverageEligibilityRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Purpose",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Serviced",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requestor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disposition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageEligibilityResponseInsurance",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreAuthRef",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Form",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Error",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CoverageEligibilityResponseError",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "CoverageEligibilityResponse",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Severity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identified",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Implicated",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Evidence",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DetectedIssueEvidence",
			Namespace: "FHIR",
		},
	}, {
		Name: "Detail",
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
	}, {
		Name: "Mitigation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DetectedIssueMitigation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DetectedIssue",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Definition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "UdiCarrier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceUdiCarrier",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DistinctIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManufactureDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExpirationDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "LotNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "SerialNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "DeviceName",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDeviceName",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModelNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Specialization",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceSpecialization",
			Namespace: "FHIR",
		},
	}, {
		Name: "Version",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceVersion",
			Namespace: "FHIR",
		},
	}, {
		Name: "Property",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceProperty",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Owner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Safety",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Device",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "UdiDeviceIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionUdiDeviceIdentifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "DeviceName",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionDeviceName",
			Namespace: "FHIR",
		},
	}, {
		Name: "ModelNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Specialization",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionSpecialization",
			Namespace: "FHIR",
		},
	}, {
		Name: "Version",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Safety",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ShelfLifeStorage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ProductShelfLife",
			Namespace: "FHIR",
		},
	}, {
		Name: "PhysicalCharacteristics",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ProdCharacteristic",
			Namespace: "FHIR",
		},
	}, {
		Name: "LanguageCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Capability",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionCapability",
			Namespace: "FHIR",
		},
	}, {
		Name: "Property",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionProperty",
			Namespace: "FHIR",
		},
	}, {
		Name: "Owner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
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
		Name: "OnlineInformation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParentDevice",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Material",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceDefinitionMaterial",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DeviceDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Unit",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OperationalStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Color",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "MeasurementPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Timing",
			Namespace: "FHIR",
		},
	}, {
		Name: "Calibration",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceMetricCalibration",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DeviceMetric",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PriorRequest",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "GroupIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parameter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DeviceRequestParameter",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PerformerType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelevantHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DeviceRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DerivedFrom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "RecordedOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DeviceUseStatement",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Effective",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issued",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ResultsInterpreter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specimen",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Result",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImagingStudy",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Media",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DiagnosticReportMedia",
			Namespace: "FHIR",
		},
	}, {
		Name: "Conclusion",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ConclusionCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PresentedForm",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DiagnosticReport",
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
		Name: "MasterIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recipient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Description",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Content",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Related",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DocumentManifestRelated",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DocumentManifest",
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
		Name: "MasterIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "DocStatus",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Date",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Authenticator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Custodian",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatesTo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DocumentReferenceRelatesTo",
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
		Name: "SecurityLabel",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Content",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DocumentReferenceContent",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DocumentReferenceContext",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "DocumentReference",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "SynthesisType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "StudyType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Exposure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExposureAlternative",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SampleSize",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "EffectEvidenceSynthesisSampleSize",
			Namespace: "FHIR",
		},
	}, {
		Name: "ResultsByExposure",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EffectEvidenceSynthesisResultsByExposure",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectEstimate",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EffectEvidenceSynthesisEffectEstimate",
			Namespace: "FHIR",
		},
	}, {
		Name: "Certainty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EffectEvidenceSynthesisCertainty",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EffectEvidenceSynthesis",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "StatusHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EncounterStatusHistory",
			Namespace: "FHIR",
		},
	}, {
		Name: "Class",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClassHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EncounterClassHistory",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "EpisodeOfCare",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EncounterParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Appointment",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Length",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Duration",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Diagnosis",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EncounterDiagnosis",
			Namespace: "FHIR",
		},
	}, {
		Name: "Account",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Hospitalization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "EncounterHospitalization",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EncounterLocation",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceProvider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Encounter",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "ConnectionType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
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
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
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
		Name: "PayloadType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PayloadMimeType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Url",
			Namespace: "FHIR",
		},
	}, {
		Name: "Header",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Endpoint",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Candidate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Coverage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EnrollmentRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disposition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Organization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "RequestProvider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EnrollmentResponse",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "StatusHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EpisodeOfCareStatusHistory",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Diagnosis",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EpisodeOfCareDiagnosis",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "ReferralRequest",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CareManager",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Team",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Account",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EpisodeOfCare",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Subtitle",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Trigger",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TriggerDefinition",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EventDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "ShortTitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subtitle",
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
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExposureBackground",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExposureVariant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Evidence",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "ShortTitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subtitle",
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
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
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
		Name: "Characteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "EvidenceVariableCharacteristic",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "EvidenceVariable",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
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
		Name: "Actor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExampleScenarioActor",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExampleScenarioInstance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Process",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExampleScenarioProcess",
			Namespace: "FHIR",
		},
	}, {
		Name: "Workflow",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ExampleScenario",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BillablePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Enterer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FundsReserveRequested",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FundsReserve",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Related",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitRelated",
			Namespace: "FHIR",
		},
	}, {
		Name: "Prescription",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OriginalPrescription",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payee",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ExplanationOfBenefitPayee",
			Namespace: "FHIR",
		},
	}, {
		Name: "Referral",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Facility",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Claim",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClaimResponse",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disposition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreAuthRef",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreAuthRefPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "CareTeam",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitCareTeam",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitSupportingInfo",
			Namespace: "FHIR",
		},
	}, {
		Name: "Diagnosis",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitDiagnosis",
			Namespace: "FHIR",
		},
	}, {
		Name: "Procedure",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitProcedure",
			Namespace: "FHIR",
		},
	}, {
		Name: "Precedence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitInsurance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Accident",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ExplanationOfBenefitAccident",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "AddItem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitAddItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "Adjudication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitItemAdjudication",
			Namespace: "FHIR",
		},
	}, {
		Name: "Total",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitTotal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ExplanationOfBenefitPayment",
			Namespace: "FHIR",
		},
	}, {
		Name: "FormCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Form",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProcessNote",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitProcessNote",
			Namespace: "FHIR",
		},
	}, {
		Name: "BenefitPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "BenefitBalance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ExplanationOfBenefitBenefitBalance",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ExplanationOfBenefit",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
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
		Name: "DataAbsentReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Relationship",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sex",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Born",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "EstimatedAge",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Deceased",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "FamilyMemberHistoryCondition",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "FamilyMemberHistory",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Flag",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "LifecycleStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "AchievementStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Description",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Start",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "GoalTarget",
			Namespace: "FHIR",
		},
	}, {
		Name: "StatusDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExpressedBy",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Addresses",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "OutcomeCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "OutcomeReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Goal",
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
		Name: "Start",
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
	}, {
		Name: "Link",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "GraphDefinitionLink",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "GraphDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Actual",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingEntity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Characteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "GroupCharacteristic",
			Namespace: "FHIR",
		},
	}, {
		Name: "Member",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "GroupMember",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Group",
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
		Name: "RequestIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Module",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OccurrenceDateTime",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "EvaluationMessage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OutputParameters",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Result",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DataRequirement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirement",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "GuidanceResponse",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProvidedBy",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExtraDetails",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Photo",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "CoverageArea",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceProvisionCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Eligibility",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HealthcareServiceEligibility",
			Namespace: "FHIR",
		},
	}, {
		Name: "Program",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Characteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Communication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferralMethod",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AppointmentRequired",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "AvailableTime",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HealthcareServiceAvailableTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "NotAvailable",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HealthcareServiceNotAvailable",
			Namespace: "FHIR",
		},
	}, {
		Name: "AvailabilityExceptions",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "HealthcareService",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Modality",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Started",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Referrer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Interpreter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "NumberOfSeries",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "NumberOfInstances",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "UnsignedInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProcedureReference",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProcedureCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "Series",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImagingStudySeries",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ImagingStudy",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "VaccineCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorded",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "PrimarySource",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReportOrigin",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "LotNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExpirationDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
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
		Name: "DoseQuantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImmunizationPerformer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "IsSubpotent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "SubpotentReason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Education",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImmunizationEducation",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProgramEligibility",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FundingSource",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reaction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImmunizationReaction",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProtocolApplied",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImmunizationProtocolApplied",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Immunization",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Authority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "TargetDisease",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImmunizationEvent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoseStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoseStatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Series",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoseNumber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "SeriesDoses",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ImmunizationEvaluation",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Authority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recommendation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImmunizationRecommendationRecommendation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ImmunizationRecommendation",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "PackageId",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Id",
			Namespace: "FHIR",
		},
	}, {
		Name: "License",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "FhirVersion",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "DependsOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImplementationGuideDependsOn",
			Namespace: "FHIR",
		},
	}, {
		Name: "Global",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ImplementationGuideGlobal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Definition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ImplementationGuideDefinition",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manifest",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ImplementationGuideManifest",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ImplementationGuide",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Alias",
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
	}, {
		Name: "OwnedBy",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdministeredBy",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CoverageArea",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InsurancePlanContact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Network",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Coverage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InsurancePlanCoverage",
			Namespace: "FHIR",
		},
	}, {
		Name: "Plan",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InsurancePlanPlan",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "InsurancePlan",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "CancelledReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recipient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InvoiceParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issuer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Account",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "LineItem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InvoiceLineItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "TotalPriceComponent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "InvoiceLineItemPriceComponent",
			Namespace: "FHIR",
		},
	}, {
		Name: "TotalNet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}, {
		Name: "TotalGross",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentTerms",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Invoice",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Subtitle",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parameter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ParameterDefinition",
			Namespace: "FHIR",
		},
	}, {
		Name: "DataRequirement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "DataRequirement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Content",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Library",
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
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "LinkageItem",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Linkage",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Mode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
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
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OrderedBy",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Entry",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ListEntry",
			Namespace: "FHIR",
		},
	}, {
		Name: "EmptyReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "List",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "OperationalStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
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
		Name: "Alias",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
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
		Name: "Mode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "PhysicalType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Position",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "LocationPosition",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "HoursOfOperation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "LocationHoursOfOperation",
			Namespace: "FHIR",
		},
	}, {
		Name: "AvailabilityExceptions",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Location",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Subtitle",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Library",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disclaimer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Scoring",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "CompositeScoring",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "RiskAdjustment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "RateAggregation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Rationale",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClinicalRecommendationStatement",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ImprovementNotation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Definition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Guidance",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Group",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MeasureGroup",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupplementalData",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MeasureSupplementalData",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Measure",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Measure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Reporter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "ImprovementNotation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Group",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MeasureReportGroup",
			Namespace: "FHIR",
		},
	}, {
		Name: "EvaluatedResource",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MeasureReport",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Modality",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "View",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issued",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Operator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DeviceName",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Height",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Width",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Frames",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PositiveInt",
			Namespace: "FHIR",
		},
	}, {
		Name: "Duration",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Content",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Media",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Form",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Amount",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Ratio",
			Namespace: "FHIR",
		},
	}, {
		Name: "Ingredient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationIngredient",
			Namespace: "FHIR",
		},
	}, {
		Name: "Batch",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationBatch",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Medication",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instantiates",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medication",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInformation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Effective",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationAdministrationPerformer",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Dosage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationAdministrationDosage",
			Namespace: "FHIR",
		},
	}, {
		Name: "EventHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicationAdministration",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medication",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInformation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationDispensePerformer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthorizingPrescription",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "DaysSupply",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "WhenPrepared",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "WhenHandedOver",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Destination",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Receiver",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "DosageInstruction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Dosage",
			Namespace: "FHIR",
		},
	}, {
		Name: "Substitution",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationDispenseSubstitution",
			Namespace: "FHIR",
		},
	}, {
		Name: "DetectedIssue",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "EventHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicationDispense",
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
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoseForm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Amount",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Synonym",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedMedicationKnowledge",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeRelatedMedicationKnowledge",
			Namespace: "FHIR",
		},
	}, {
		Name: "AssociatedMedication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProductType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Monograph",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeMonograph",
			Namespace: "FHIR",
		},
	}, {
		Name: "Ingredient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeIngredient",
			Namespace: "FHIR",
		},
	}, {
		Name: "PreparationInstruction",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "IntendedRoute",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Cost",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeCost",
			Namespace: "FHIR",
		},
	}, {
		Name: "MonitoringProgram",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeMonitoringProgram",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdministrationGuidelines",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeAdministrationGuidelines",
			Namespace: "FHIR",
		},
	}, {
		Name: "MedicineClassification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeMedicineClassification",
			Namespace: "FHIR",
		},
	}, {
		Name: "Packaging",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationKnowledgePackaging",
			Namespace: "FHIR",
		},
	}, {
		Name: "DrugCharacteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeDrugCharacteristic",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contraindication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Regulatory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeRegulatory",
			Namespace: "FHIR",
		},
	}, {
		Name: "Kinetics",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicationKnowledgeKinetics",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicationKnowledge",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoNotPerform",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reported",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medication",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInformation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PerformerType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "GroupIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "CourseOfTherapyType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "DosageInstruction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Dosage",
			Namespace: "FHIR",
		},
	}, {
		Name: "DispenseRequest",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationRequestDispenseRequest",
			Namespace: "FHIR",
		},
	}, {
		Name: "Substitution",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicationRequestSubstitution",
			Namespace: "FHIR",
		},
	}, {
		Name: "PriorPrescription",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DetectedIssue",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "EventHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicationRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Medication",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Effective",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateAsserted",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "InformationSource",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DerivedFrom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Dosage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Dosage",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicationStatement",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Domain",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Coding",
			Namespace: "FHIR",
		},
	}, {
		Name: "CombinedPharmaceuticalDoseForm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "LegalStatusOfSupply",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdditionalMonitoringIndicator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SpecialMeasures",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaediatricUseIndicator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProductClassification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "MarketingStatus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MarketingStatus",
			Namespace: "FHIR",
		},
	}, {
		Name: "PharmaceuticalProduct",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PackagedMedicinalProduct",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AttachedDocument",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "MasterFile",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ClinicalTrial",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductName",
			Namespace: "FHIR",
		},
	}, {
		Name: "CrossReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManufacturingBusinessOperation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductManufacturingBusinessOperation",
			Namespace: "FHIR",
		},
	}, {
		Name: "SpecialDesignation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductSpecialDesignation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProduct",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Country",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "StatusDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "RestoreDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidityPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "DataExclusivityPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateOfFirstAuthorization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "InternationalBirthDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "LegalBasis",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "JurisdictionalAuthorization",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductAuthorizationJurisdictionalAuthorization",
			Namespace: "FHIR",
		},
	}, {
		Name: "Holder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Regulator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Procedure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicinalProductAuthorizationProcedure",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductAuthorization",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disease",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DiseaseStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comorbidity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "TherapeuticIndication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "OtherTherapy",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductContraindicationOtherTherapy",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Population",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductContraindication",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DiseaseSymptomProcedure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "DiseaseStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comorbidity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "IntendedEffect",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Duration",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "OtherTherapy",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductIndicationOtherTherapy",
			Namespace: "FHIR",
		},
	}, {
		Name: "UndesirableEffect",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Population",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductIndication",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Role",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AllergenicIndicator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SpecifiedSubstance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductIngredientSpecifiedSubstance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Substance",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MedicinalProductIngredientSubstance",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductIngredient",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Interactant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductInteractionInteractant",
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
		Name: "Effect",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Incidence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Management",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductInteraction",
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
		Name: "ManufacturedDoseForm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "UnitOfPresentation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Ingredient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PhysicalCharacteristics",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ProdCharacteristic",
			Namespace: "FHIR",
		},
	}, {
		Name: "OtherCharacteristics",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductManufactured",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "LegalStatusOfSupply",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "MarketingStatus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MarketingStatus",
			Namespace: "FHIR",
		},
	}, {
		Name: "MarketingAuthorization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Manufacturer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BatchIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductPackagedBatchIdentifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "PackageItem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductPackagedPackageItem",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductPackaged",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "AdministrableDoseForm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "UnitOfPresentation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Ingredient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Characteristics",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductPharmaceuticalCharacteristics",
			Namespace: "FHIR",
		},
	}, {
		Name: "RouteOfAdministration",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MedicinalProductPharmaceuticalRouteOfAdministration",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductPharmaceutical",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SymptomConditionEffect",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Classification",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FrequencyOfOccurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Population",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MedicinalProductUndesirableEffect",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "Base",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Event",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Focus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MessageDefinitionFocus",
			Namespace: "FHIR",
		},
	}, {
		Name: "ResponseRequired",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "AllowedResponse",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MessageDefinitionAllowedResponse",
			Namespace: "FHIR",
		},
	}, {
		Name: "Graph",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MessageDefinition",
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
		Name: "Event",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Destination",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MessageHeaderDestination",
			Namespace: "FHIR",
		},
	}, {
		Name: "Sender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Enterer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MessageHeaderSource",
			Namespace: "FHIR",
		},
	}, {
		Name: "Responsible",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Response",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MessageHeaderResponse",
			Namespace: "FHIR",
		},
	}, {
		Name: "Focus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Definition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MessageHeader",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
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
		Name: "CoordinateSystem",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specimen",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferenceSeq",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "MolecularSequenceReferenceSeq",
			Namespace: "FHIR",
		},
	}, {
		Name: "Variant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MolecularSequenceVariant",
			Namespace: "FHIR",
		},
	}, {
		Name: "ObservedSeq",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quality",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MolecularSequenceQuality",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReadCoverage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Repository",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MolecularSequenceRepository",
			Namespace: "FHIR",
		},
	}, {
		Name: "Pointer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "StructureVariant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "MolecularSequenceStructureVariant",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "MolecularSequence",
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
		Name: "Name",
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
		Name: "Kind",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
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
		Name: "Responsible",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "UniqueId",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "NamingSystemUniqueId",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "NamingSystem",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instantiates",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
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
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateTime",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Orderer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AllergyIntolerance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "FoodPreferenceModifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExcludeFoodModifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "OralDiet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "NutritionOrderOralDiet",
			Namespace: "FHIR",
		},
	}, {
		Name: "Supplement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "NutritionOrderSupplement",
			Namespace: "FHIR",
		},
	}, {
		Name: "EnteralFormula",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "NutritionOrderEnteralFormula",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "NutritionOrder",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Focus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Effective",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issued",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Value",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "DataAbsentReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Interpretation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
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
		Name: "Specimen",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Device",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferenceRange",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ObservationReferenceRange",
			Namespace: "FHIR",
		},
	}, {
		Name: "HasMember",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DerivedFrom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Component",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ObservationComponent",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Observation",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "PermittedDataType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "MultipleResultsAllowed",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "PreferredReportName",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "QuantitativeDetails",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ObservationDefinitionQuantitativeDetails",
			Namespace: "FHIR",
		},
	}, {
		Name: "QualifiedInterval",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ObservationDefinitionQualifiedInterval",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidCodedValueSet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "NormalCodedValueSet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AbnormalCodedValueSet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "CriticalCodedValueSet",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ObservationDefinition",
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
		Name: "Kind",
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
		Name: "AffectsState",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "Base",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Resource",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "System",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Instance",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "InputProfile",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "OutputProfile",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parameter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "OperationDefinitionParameter",
			Namespace: "FHIR",
		},
	}, {
		Name: "Overload",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "OperationDefinitionOverload",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "OperationDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Alias",
		Type: fhirpath.TypeSpecifier{
			List:      true,
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
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "OrganizationContact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Organization",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Organization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParticipatingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Network",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "HealthcareService",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "OrganizationAffiliation",
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
		Name: "Parameter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ParametersParameter",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Parameters",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HumanName",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "BirthDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Deceased",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "MaritalStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "MultipleBirth",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Photo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PatientContact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Communication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PatientCommunication",
			Namespace: "FHIR",
		},
	}, {
		Name: "GeneralPractitioner",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Link",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PatientLink",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Patient",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Response",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Provider",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Payee",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recipient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Amount",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "PaymentNotice",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentIssuer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requestor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Disposition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentAmount",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Money",
			Namespace: "FHIR",
		},
	}, {
		Name: "PaymentIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Detail",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PaymentReconciliationDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "FormCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ProcessNote",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PaymentReconciliationProcessNote",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "PaymentReconciliation",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HumanName",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "BirthDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "Photo",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "ManagingOrganization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Link",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PersonLink",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Person",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Subtitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Library",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Goal",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PlanDefinitionGoal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Action",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PlanDefinitionAction",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "PlanDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HumanName",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "BirthDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Photo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
			Namespace: "FHIR",
		},
	}, {
		Name: "Qualification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PractitionerQualification",
			Namespace: "FHIR",
		},
	}, {
		Name: "Communication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Practitioner",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Practitioner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Organization",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "HealthcareService",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "AvailableTime",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PractitionerRoleAvailableTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "NotAvailable",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "PractitionerRoleNotAvailable",
			Namespace: "FHIR",
		},
	}, {
		Name: "AvailabilityExceptions",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endpoint",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "PractitionerRole",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performed",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorder",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Asserter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ProcedurePerformer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Report",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Complication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ComplicationDetail",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "FollowUp",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "FocalDevice",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ProcedureFocalDevice",
			Namespace: "FHIR",
		},
	}, {
		Name: "UsedReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "UsedCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Procedure",
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
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurred",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Recorded",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Policy",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reason",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Activity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Agent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ProvenanceAgent",
			Namespace: "FHIR",
		},
	}, {
		Name: "Entity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ProvenanceEntity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Signature",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Signature",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Provenance",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "DerivedFrom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "SubjectType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
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
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "QuestionnaireItem",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Questionnaire",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Questionnaire",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Authored",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "QuestionnaireResponseItem",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "QuestionnaireResponse",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Relationship",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "HumanName",
			Namespace: "FHIR",
		},
	}, {
		Name: "Telecom",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gender",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "BirthDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "Address",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Address",
			Namespace: "FHIR",
		},
	}, {
		Name: "Photo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Attachment",
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
		Name: "Communication",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedPersonCommunication",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "RelatedPerson",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "GroupIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Action",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RequestGroupAction",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "RequestGroup",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "ShortTitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subtitle",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Library",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Exposure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExposureAlternative",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ResearchDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "ShortTitle",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subtitle",
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
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
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
		Name: "Usage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
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
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Library",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
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
		Name: "VariableType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Characteristic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ResearchElementDefinitionCharacteristic",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ResearchElementDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Protocol",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "PrimaryPurposeType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Phase",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Focus",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "Keyword",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
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
		Name: "Enrollment",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Sponsor",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PrincipalInvestigator",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Site",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonStopped",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Arm",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ResearchStudyArm",
			Namespace: "FHIR",
		},
	}, {
		Name: "Objective",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ResearchStudyObjective",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ResearchStudy",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Period",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Study",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Individual",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "AssignedArm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ActualArm",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Consent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ResearchSubject",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Method",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Basis",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Prediction",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RiskAssessmentPrediction",
			Namespace: "FHIR",
		},
	}, {
		Name: "Mitigation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "RiskAssessment",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "Copyright",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Markdown",
			Namespace: "FHIR",
		},
	}, {
		Name: "ApprovalDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastReviewDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "EffectivePeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Topic",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Author",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Editor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reviewer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "Endorser",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactDetail",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelatedArtifact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RelatedArtifact",
			Namespace: "FHIR",
		},
	}, {
		Name: "SynthesisType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "StudyType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Population",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Exposure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Outcome",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SampleSize",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "RiskEvidenceSynthesisSampleSize",
			Namespace: "FHIR",
		},
	}, {
		Name: "RiskEstimate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "RiskEvidenceSynthesisRiskEstimate",
			Namespace: "FHIR",
		},
	}, {
		Name: "Certainty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "RiskEvidenceSynthesisCertainty",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "RiskEvidenceSynthesis",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "Active",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceCategory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Actor",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PlanningHorizon",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Schedule",
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
		Name: "DerivedFrom",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
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
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Base",
		Type: fhirpath.TypeSpecifier{
			List:      true,
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
		Name: "Expression",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Xpath",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "XpathUsage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "MultipleOr",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "MultipleAnd",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comparator",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Modifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Chain",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Component",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SearchParameterComponent",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SearchParameter",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Replaces",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requisition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "DoNotPerform",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "OrderDetail",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
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
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PerformerType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Performer",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "LocationCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "LocationReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SupportingInfo",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specimen",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "BodySite",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
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
		Name: "RelevantHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ServiceRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceCategory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ServiceType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Specialty",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "AppointmentType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Schedule",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "Start",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "End",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Overbooked",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Slot",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "AccessionIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Type",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subject",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReceivedTime",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parent",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Request",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Collection",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SpecimenCollection",
			Namespace: "FHIR",
		},
	}, {
		Name: "Processing",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SpecimenProcessing",
			Namespace: "FHIR",
		},
	}, {
		Name: "Container",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SpecimenContainer",
			Namespace: "FHIR",
		},
	}, {
		Name: "Condition",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Specimen",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "TypeCollected",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PatientPreparation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "TimeAspect",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Collection",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "TypeTested",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SpecimenDefinitionTypeTested",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SpecimenDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Keyword",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Coding",
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
		Name: "Mapping",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "StructureDefinitionMapping",
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
		Name: "Abstract",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "Context",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "StructureDefinitionContext",
			Namespace: "FHIR",
		},
	}, {
		Name: "ContextInvariant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
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
		Name: "BaseDefinition",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Derivation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Snapshot",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "StructureDefinitionSnapshot",
			Namespace: "FHIR",
		},
	}, {
		Name: "Differential",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "StructureDefinitionDifferential",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "StructureDefinition",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Structure",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "StructureMapStructure",
			Namespace: "FHIR",
		},
	}, {
		Name: "Import",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "Group",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "StructureMapGroup",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "StructureMap",
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
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Contact",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "ContactPoint",
			Namespace: "FHIR",
		},
	}, {
		Name: "End",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Instant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Reason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Criteria",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Error",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Channel",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SubscriptionChannel",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Subscription",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Instance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceInstance",
			Namespace: "FHIR",
		},
	}, {
		Name: "Ingredient",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceIngredient",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Substance",
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
		Name: "SequenceType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "NumberOfSubunits",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "AreaOfHybridisation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "OligoNucleotideType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subunit",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceNucleicAcidSubunit",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstanceNucleicAcid",
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
		Name: "Class",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Geometry",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "CopolymerConnectivity",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Modification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "MonomerSet",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstancePolymerMonomerSet",
			Namespace: "FHIR",
		},
	}, {
		Name: "Repeat",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstancePolymerRepeat",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstancePolymer",
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
		Name: "SequenceType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "NumberOfSubunits",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Integer",
			Namespace: "FHIR",
		},
	}, {
		Name: "DisulfideLinkage",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Subunit",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceProteinSubunit",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstanceProtein",
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
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Gene",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceReferenceInformationGene",
			Namespace: "FHIR",
		},
	}, {
		Name: "GeneElement",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceReferenceInformationGeneElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Classification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceReferenceInformationClassification",
			Namespace: "FHIR",
		},
	}, {
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceReferenceInformationTarget",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstanceReferenceInformation",
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
		Name: "SourceMaterialClass",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SourceMaterialType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "SourceMaterialState",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "OrganismId",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "OrganismName",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParentSubstanceId",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "ParentSubstanceName",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "CountryOfOrigin",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "GeographicalLocation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "DevelopmentStage",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "FractionDescription",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSourceMaterialFractionDescription",
			Namespace: "FHIR",
		},
	}, {
		Name: "Organism",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SubstanceSourceMaterialOrganism",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartDescription",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSourceMaterialPartDescription",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstanceSourceMaterial",
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
		Name: "Domain",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Source",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Comment",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Moiety",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationMoiety",
			Namespace: "FHIR",
		},
	}, {
		Name: "Property",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationProperty",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReferenceInformation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Structure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SubstanceSpecificationStructure",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationCode",
			Namespace: "FHIR",
		},
	}, {
		Name: "Name",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationName",
			Namespace: "FHIR",
		},
	}, {
		Name: "MolecularWeight",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationStructureIsotopeMolecularWeight",
			Namespace: "FHIR",
		},
	}, {
		Name: "Relationship",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SubstanceSpecificationRelationship",
			Namespace: "FHIR",
		},
	}, {
		Name: "NucleicAcid",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Polymer",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Protein",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "SourceMaterial",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SubstanceSpecification",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
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
		Name: "SuppliedItem",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "SupplyDeliverySuppliedItem",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Supplier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Destination",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Receiver",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SupplyDelivery",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Category",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Item",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "Quantity",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Quantity",
			Namespace: "FHIR",
		},
	}, {
		Name: "Parameter",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "SupplyRequestParameter",
			Namespace: "FHIR",
		},
	}, {
		Name: "Occurrence",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "PrimitiveElement",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Supplier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DeliverFrom",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DeliverTo",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "SupplyRequest",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesCanonical",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Canonical",
			Namespace: "FHIR",
		},
	}, {
		Name: "InstantiatesUri",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Uri",
			Namespace: "FHIR",
		},
	}, {
		Name: "BasedOn",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "GroupIdentifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
			Namespace: "FHIR",
		},
	}, {
		Name: "PartOf",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
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
		Name: "StatusReason",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "BusinessStatus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Intent",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Priority",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Code",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "Focus",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "For",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ExecutionPeriod",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Period",
			Namespace: "FHIR",
		},
	}, {
		Name: "AuthoredOn",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastModified",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Requester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "PerformerType",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Owner",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Location",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ReasonReference",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Insurance",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Note",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Annotation",
			Namespace: "FHIR",
		},
	}, {
		Name: "RelevantHistory",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Restriction",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TaskRestriction",
			Namespace: "FHIR",
		},
	}, {
		Name: "Input",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TaskInput",
			Namespace: "FHIR",
		},
	}, {
		Name: "Output",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TaskOutput",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "Task",
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
		Name: "Software",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesSoftware",
			Namespace: "FHIR",
		},
	}, {
		Name: "Implementation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesImplementation",
			Namespace: "FHIR",
		},
	}, {
		Name: "LockedDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
			Namespace: "FHIR",
		},
	}, {
		Name: "CodeSystem",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TerminologyCapabilitiesCodeSystem",
			Namespace: "FHIR",
		},
	}, {
		Name: "Expansion",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesExpansion",
			Namespace: "FHIR",
		},
	}, {
		Name: "CodeSearch",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidateCode",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesValidateCode",
			Namespace: "FHIR",
		},
	}, {
		Name: "Translation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesTranslation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Closure",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TerminologyCapabilitiesClosure",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "TerminologyCapabilities",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Status",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "TestScript",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Result",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Code",
			Namespace: "FHIR",
		},
	}, {
		Name: "Score",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Decimal",
			Namespace: "FHIR",
		},
	}, {
		Name: "Tester",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Issued",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Participant",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestReportParticipant",
			Namespace: "FHIR",
		},
	}, {
		Name: "Setup",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TestReportSetup",
			Namespace: "FHIR",
		},
	}, {
		Name: "Test",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestReportTest",
			Namespace: "FHIR",
		},
	}, {
		Name: "Teardown",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TestReportTeardown",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "TestReport",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Identifier",
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
		Name: "Origin",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestScriptOrigin",
			Namespace: "FHIR",
		},
	}, {
		Name: "Destination",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestScriptDestination",
			Namespace: "FHIR",
		},
	}, {
		Name: "Metadata",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TestScriptMetadata",
			Namespace: "FHIR",
		},
	}, {
		Name: "Fixture",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestScriptFixture",
			Namespace: "FHIR",
		},
	}, {
		Name: "Profile",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Variable",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestScriptVariable",
			Namespace: "FHIR",
		},
	}, {
		Name: "Setup",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TestScriptSetup",
			Namespace: "FHIR",
		},
	}, {
		Name: "Test",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "TestScriptTest",
			Namespace: "FHIR",
		},
	}, {
		Name: "Teardown",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "TestScriptTeardown",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "TestScript",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Immutable",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Boolean",
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
		Name: "Compose",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ValueSetCompose",
			Namespace: "FHIR",
		},
	}, {
		Name: "Expansion",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "ValueSetExpansion",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "ValueSet",
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
		Name: "Target",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "TargetLocation",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "String",
			Namespace: "FHIR",
		},
	}, {
		Name: "Need",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
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
		Name: "StatusDate",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidationType",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "ValidationProcess",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "Frequency",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Timing",
			Namespace: "FHIR",
		},
	}, {
		Name: "LastPerformed",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "NextScheduled",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Date",
			Namespace: "FHIR",
		},
	}, {
		Name: "FailureAction",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "CodeableConcept",
			Namespace: "FHIR",
		},
	}, {
		Name: "PrimarySource",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "VerificationResultPrimarySource",
			Namespace: "FHIR",
		},
	}, {
		Name: "Attestation",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "VerificationResultAttestation",
			Namespace: "FHIR",
		},
	}, {
		Name: "Validator",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "VerificationResultValidator",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "VerificationResult",
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
		Name: "Identifier",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "Identifier",
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
		Name: "Created",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Patient",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "Encounter",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "DateWritten",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "DateTime",
			Namespace: "FHIR",
		},
	}, {
		Name: "Prescriber",
		Type: fhirpath.TypeSpecifier{
			List:      false,
			Name:      "Reference",
			Namespace: "FHIR",
		},
	}, {
		Name: "LensSpecification",
		Type: fhirpath.TypeSpecifier{
			List:      true,
			Name:      "VisionPrescriptionLensSpecification",
			Namespace: "FHIR",
		},
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
		Name:      "VisionPrescription",
		Namespace: "FHIR",
	},
}}
var fhirFunctions = fhirpath.Functions{}
