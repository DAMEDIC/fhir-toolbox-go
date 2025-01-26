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
	ctx = fhirpath.WithFunctions(ctx, fhirFunctions)
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
		Type: "List<FHIR.Extension>",
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
		Type: "List<FHIR.Extension>",
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
		Type: "FHIR.id",
	}, {
		Name: "meta",
		Type: "FHIR.Meta",
	}, {
		Name: "implicitRules",
		Type: "FHIR.uri",
	}, {
		Name: "language",
		Type: "FHIR.code",
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
		Type: "FHIR.id",
	}, {
		Name: "meta",
		Type: "FHIR.Meta",
	}, {
		Name: "implicitRules",
		Type: "FHIR.uri",
	}, {
		Name: "language",
		Type: "FHIR.code",
	}, {
		Name: "text",
		Type: "FHIR.Narrative",
	}, {
		Name: "contained",
		Type: "List<FHIR.Resource>",
	}, {
		Name: "extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "modifierExtension",
		Type: "List<FHIR.Extension>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Base64Binary",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.bool",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Boolean",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Canonical",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Code",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Date",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "DateTime",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Decimal",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Id",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Instant",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.int32",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Integer",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Markdown",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Oid",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.uint32",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "PositiveInt",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "String",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Time",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.uint32",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "UnsignedInt",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Uri",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Url",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Uuid",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Value",
		Type: "FHIR.string",
	}},
	SimpleTypeInfo: fhirpath.SimpleTypeInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "PrimitiveType",
			Namespace: "FHIR",
		},
		Name:      "Xhtml",
		Namespace: "FHIR",
	},
}, fhirpath.ClassInfo{
	Element: []fhirpath.ClassInfoElement{{
		Name: "Id",
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Use",
		Type: "FHIR.Code",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Text",
		Type: "FHIR.String",
	}, {
		Name: "Line",
		Type: "List<FHIR.String>",
	}, {
		Name: "City",
		Type: "FHIR.String",
	}, {
		Name: "District",
		Type: "FHIR.String",
	}, {
		Name: "State",
		Type: "FHIR.String",
	}, {
		Name: "PostalCode",
		Type: "FHIR.String",
	}, {
		Name: "Country",
		Type: "FHIR.String",
	}, {
		Name: "Period",
		Type: "FHIR.Period",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Author",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Time",
		Type: "FHIR.DateTime",
	}, {
		Name: "Text",
		Type: "FHIR.Markdown",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ContentType",
		Type: "FHIR.Code",
	}, {
		Name: "Language",
		Type: "FHIR.Code",
	}, {
		Name: "Data",
		Type: "FHIR.Base64Binary",
	}, {
		Name: "Url",
		Type: "FHIR.Url",
	}, {
		Name: "Size",
		Type: "FHIR.UnsignedInt",
	}, {
		Name: "Hash",
		Type: "FHIR.Base64Binary",
	}, {
		Name: "Title",
		Type: "FHIR.String",
	}, {
		Name: "Creation",
		Type: "FHIR.DateTime",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Coding",
		Type: "List<FHIR.Coding>",
	}, {
		Name: "Text",
		Type: "FHIR.String",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Version",
		Type: "FHIR.String",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
	}, {
		Name: "Display",
		Type: "FHIR.String",
	}, {
		Name: "UserSelected",
		Type: "FHIR.Boolean",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Name",
		Type: "FHIR.String",
	}, {
		Name: "Telecom",
		Type: "List<FHIR.ContactPoint>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "System",
		Type: "FHIR.Code",
	}, {
		Name: "Value",
		Type: "FHIR.String",
	}, {
		Name: "Use",
		Type: "FHIR.Code",
	}, {
		Name: "Rank",
		Type: "FHIR.PositiveInt",
	}, {
		Name: "Period",
		Type: "FHIR.Period",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Name",
		Type: "FHIR.String",
	}, {
		Name: "Contact",
		Type: "List<FHIR.ContactDetail>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Profile",
		Type: "List<FHIR.Canonical>",
	}, {
		Name: "Subject",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "MustSupport",
		Type: "List<FHIR.String>",
	}, {
		Name: "CodeFilter",
		Type: "List<FHIR.DataRequirementCodeFilter>",
	}, {
		Name: "DateFilter",
		Type: "List<FHIR.DataRequirementDateFilter>",
	}, {
		Name: "Limit",
		Type: "FHIR.PositiveInt",
	}, {
		Name: "Sort",
		Type: "List<FHIR.DataRequirementSort>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Sequence",
		Type: "FHIR.Integer",
	}, {
		Name: "Text",
		Type: "FHIR.String",
	}, {
		Name: "AdditionalInstruction",
		Type: "List<FHIR.CodeableConcept>",
	}, {
		Name: "PatientInstruction",
		Type: "FHIR.String",
	}, {
		Name: "Timing",
		Type: "FHIR.Timing",
	}, {
		Name: "AsNeeded",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Site",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Route",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Method",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "DoseAndRate",
		Type: "List<FHIR.DosageDoseAndRate>",
	}, {
		Name: "MaxDosePerPeriod",
		Type: "FHIR.Ratio",
	}, {
		Name: "MaxDosePerAdministration",
		Type: "FHIR.Quantity",
	}, {
		Name: "MaxDosePerLifetime",
		Type: "FHIR.Quantity",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Path",
		Type: "FHIR.String",
	}, {
		Name: "Representation",
		Type: "List<FHIR.Code>",
	}, {
		Name: "SliceName",
		Type: "FHIR.String",
	}, {
		Name: "SliceIsConstraining",
		Type: "FHIR.Boolean",
	}, {
		Name: "Label",
		Type: "FHIR.String",
	}, {
		Name: "Code",
		Type: "List<FHIR.Coding>",
	}, {
		Name: "Slicing",
		Type: "FHIR.ElementDefinitionSlicing",
	}, {
		Name: "Short",
		Type: "FHIR.String",
	}, {
		Name: "Definition",
		Type: "FHIR.Markdown",
	}, {
		Name: "Comment",
		Type: "FHIR.Markdown",
	}, {
		Name: "Requirements",
		Type: "FHIR.Markdown",
	}, {
		Name: "Alias",
		Type: "List<FHIR.String>",
	}, {
		Name: "Min",
		Type: "FHIR.UnsignedInt",
	}, {
		Name: "Max",
		Type: "FHIR.String",
	}, {
		Name: "Base",
		Type: "FHIR.ElementDefinitionBase",
	}, {
		Name: "ContentReference",
		Type: "FHIR.Uri",
	}, {
		Name: "Type",
		Type: "List<FHIR.ElementDefinitionType>",
	}, {
		Name: "DefaultValue",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "MeaningWhenMissing",
		Type: "FHIR.Markdown",
	}, {
		Name: "OrderMeaning",
		Type: "FHIR.String",
	}, {
		Name: "Fixed",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Pattern",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Example",
		Type: "List<FHIR.ElementDefinitionExample>",
	}, {
		Name: "MinValue",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "MaxValue",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "MaxLength",
		Type: "FHIR.Integer",
	}, {
		Name: "Condition",
		Type: "List<FHIR.Id>",
	}, {
		Name: "Constraint",
		Type: "List<FHIR.ElementDefinitionConstraint>",
	}, {
		Name: "MustSupport",
		Type: "FHIR.Boolean",
	}, {
		Name: "IsModifier",
		Type: "FHIR.Boolean",
	}, {
		Name: "IsModifierReason",
		Type: "FHIR.String",
	}, {
		Name: "IsSummary",
		Type: "FHIR.Boolean",
	}, {
		Name: "Binding",
		Type: "FHIR.ElementDefinitionBinding",
	}, {
		Name: "Mapping",
		Type: "List<FHIR.ElementDefinitionMapping>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Description",
		Type: "FHIR.String",
	}, {
		Name: "Name",
		Type: "FHIR.Id",
	}, {
		Name: "Language",
		Type: "FHIR.Code",
	}, {
		Name: "Expression",
		Type: "FHIR.String",
	}, {
		Name: "Reference",
		Type: "FHIR.Uri",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Url",
		Type: "FHIR.string",
	}, {
		Name: "Value",
		Type: "FHIR.PrimitiveElement",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Use",
		Type: "FHIR.Code",
	}, {
		Name: "Text",
		Type: "FHIR.String",
	}, {
		Name: "Family",
		Type: "FHIR.String",
	}, {
		Name: "Given",
		Type: "List<FHIR.String>",
	}, {
		Name: "Prefix",
		Type: "List<FHIR.String>",
	}, {
		Name: "Suffix",
		Type: "List<FHIR.String>",
	}, {
		Name: "Period",
		Type: "FHIR.Period",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Use",
		Type: "FHIR.Code",
	}, {
		Name: "Type",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Value",
		Type: "FHIR.String",
	}, {
		Name: "Period",
		Type: "FHIR.Period",
	}, {
		Name: "Assigner",
		Type: "FHIR.Reference",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Country",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Jurisdiction",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Status",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "DateRange",
		Type: "FHIR.Period",
	}, {
		Name: "RestoreDate",
		Type: "FHIR.DateTime",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "VersionId",
		Type: "FHIR.Id",
	}, {
		Name: "LastUpdated",
		Type: "FHIR.Instant",
	}, {
		Name: "Source",
		Type: "FHIR.Uri",
	}, {
		Name: "Profile",
		Type: "List<FHIR.Canonical>",
	}, {
		Name: "Security",
		Type: "List<FHIR.Coding>",
	}, {
		Name: "Tag",
		Type: "List<FHIR.Coding>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Currency",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Status",
		Type: "FHIR.Code",
	}, {
		Name: "Div",
		Type: "FHIR.Xhtml",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Name",
		Type: "FHIR.Code",
	}, {
		Name: "Use",
		Type: "FHIR.Code",
	}, {
		Name: "Min",
		Type: "FHIR.Integer",
	}, {
		Name: "Max",
		Type: "FHIR.String",
	}, {
		Name: "Documentation",
		Type: "FHIR.String",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Profile",
		Type: "FHIR.Canonical",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Start",
		Type: "FHIR.DateTime",
	}, {
		Name: "End",
		Type: "FHIR.DateTime",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Age",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Gender",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Race",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "PhysiologicalCondition",
		Type: "FHIR.CodeableConcept",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Height",
		Type: "FHIR.Quantity",
	}, {
		Name: "Width",
		Type: "FHIR.Quantity",
	}, {
		Name: "Depth",
		Type: "FHIR.Quantity",
	}, {
		Name: "Weight",
		Type: "FHIR.Quantity",
	}, {
		Name: "NominalVolume",
		Type: "FHIR.Quantity",
	}, {
		Name: "ExternalDiameter",
		Type: "FHIR.Quantity",
	}, {
		Name: "Shape",
		Type: "FHIR.String",
	}, {
		Name: "Color",
		Type: "List<FHIR.String>",
	}, {
		Name: "Imprint",
		Type: "List<FHIR.String>",
	}, {
		Name: "Image",
		Type: "List<FHIR.Attachment>",
	}, {
		Name: "Scoring",
		Type: "FHIR.CodeableConcept",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "Period",
		Type: "FHIR.Quantity",
	}, {
		Name: "SpecialPrecautionsForStorage",
		Type: "List<FHIR.CodeableConcept>",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Low",
		Type: "FHIR.Quantity",
	}, {
		Name: "High",
		Type: "FHIR.Quantity",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Numerator",
		Type: "FHIR.Quantity",
	}, {
		Name: "Denominator",
		Type: "FHIR.Quantity",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Reference",
		Type: "FHIR.String",
	}, {
		Name: "Type",
		Type: "FHIR.Uri",
	}, {
		Name: "Identifier",
		Type: "FHIR.Identifier",
	}, {
		Name: "Display",
		Type: "FHIR.String",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Label",
		Type: "FHIR.String",
	}, {
		Name: "Display",
		Type: "FHIR.String",
	}, {
		Name: "Citation",
		Type: "FHIR.Markdown",
	}, {
		Name: "Document",
		Type: "FHIR.Attachment",
	}, {
		Name: "Resource",
		Type: "FHIR.Canonical",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Origin",
		Type: "FHIR.Quantity",
	}, {
		Name: "Factor",
		Type: "FHIR.Decimal",
	}, {
		Name: "LowerLimit",
		Type: "FHIR.Decimal",
	}, {
		Name: "UpperLimit",
		Type: "FHIR.Decimal",
	}, {
		Name: "Dimensions",
		Type: "FHIR.PositiveInt",
	}, {
		Name: "Data",
		Type: "FHIR.String",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "List<FHIR.Coding>",
	}, {
		Name: "When",
		Type: "FHIR.Instant",
	}, {
		Name: "Who",
		Type: "FHIR.Reference",
	}, {
		Name: "OnBehalfOf",
		Type: "FHIR.Reference",
	}, {
		Name: "TargetFormat",
		Type: "FHIR.Code",
	}, {
		Name: "SigFormat",
		Type: "FHIR.Code",
	}, {
		Name: "Data",
		Type: "FHIR.Base64Binary",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Amount",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "AmountType",
		Type: "FHIR.CodeableConcept",
	}, {
		Name: "AmountText",
		Type: "FHIR.String",
	}, {
		Name: "ReferenceRange",
		Type: "FHIR.SubstanceAmountReferenceRange",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Event",
		Type: "List<FHIR.DateTime>",
	}, {
		Name: "Repeat",
		Type: "FHIR.TimingRepeat",
	}, {
		Name: "Code",
		Type: "FHIR.CodeableConcept",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Name",
		Type: "FHIR.String",
	}, {
		Name: "Timing",
		Type: "FHIR.PrimitiveElement",
	}, {
		Name: "Data",
		Type: "List<FHIR.DataRequirement>",
	}, {
		Name: "Condition",
		Type: "FHIR.Expression",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Code",
		Type: "FHIR.Coding",
	}, {
		Name: "Value",
		Type: "FHIR.PrimitiveElement",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Comparator",
		Type: "FHIR.Code",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.string",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Value",
		Type: "FHIR.Decimal",
	}, {
		Name: "Unit",
		Type: "FHIR.String",
	}, {
		Name: "System",
		Type: "FHIR.Uri",
	}, {
		Name: "Code",
		Type: "FHIR.Code",
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
		Type: "FHIR.Id",
	}, {
		Name: "Meta",
		Type: "FHIR.Meta",
	}, {
		Name: "ImplicitRules",
		Type: "FHIR.Uri",
	}, {
		Name: "Language",
		Type: "FHIR.Code",
	}, {
		Name: "Identifier",
		Type: "FHIR.Identifier",
	}, {
		Name: "Type",
		Type: "FHIR.Code",
	}, {
		Name: "Timestamp",
		Type: "FHIR.Instant",
	}, {
		Name: "Total",
		Type: "FHIR.UnsignedInt",
	}, {
		Name: "Link",
		Type: "List<FHIR.BundleLink>",
	}, {
		Name: "Entry",
		Type: "List<FHIR.BundleEntry>",
	}, {
		Name: "Signature",
		Type: "FHIR.Signature",
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
		Type: "FHIR.Id",
	}, {
		Name: "Meta",
		Type: "FHIR.Meta",
	}, {
		Name: "ImplicitRules",
		Type: "FHIR.Uri",
	}, {
		Name: "Language",
		Type: "FHIR.Code",
	}, {
		Name: "Text",
		Type: "FHIR.Narrative",
	}, {
		Name: "Contained",
		Type: "List<FHIR.>",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Url",
		Type: "FHIR.Uri",
	}, {
		Name: "Version",
		Type: "FHIR.String",
	}, {
		Name: "Name",
		Type: "FHIR.String",
	}, {
		Name: "Title",
		Type: "FHIR.String",
	}, {
		Name: "Status",
		Type: "FHIR.Code",
	}, {
		Name: "Experimental",
		Type: "FHIR.Boolean",
	}, {
		Name: "Date",
		Type: "FHIR.DateTime",
	}, {
		Name: "Publisher",
		Type: "FHIR.String",
	}, {
		Name: "Contact",
		Type: "List<FHIR.ContactDetail>",
	}, {
		Name: "Description",
		Type: "FHIR.Markdown",
	}, {
		Name: "UseContext",
		Type: "List<FHIR.UsageContext>",
	}, {
		Name: "Jurisdiction",
		Type: "List<FHIR.CodeableConcept>",
	}, {
		Name: "Purpose",
		Type: "FHIR.Markdown",
	}, {
		Name: "Copyright",
		Type: "FHIR.Markdown",
	}, {
		Name: "Kind",
		Type: "FHIR.Code",
	}, {
		Name: "Instantiates",
		Type: "List<FHIR.Canonical>",
	}, {
		Name: "Imports",
		Type: "List<FHIR.Canonical>",
	}, {
		Name: "Software",
		Type: "FHIR.CapabilityStatementSoftware",
	}, {
		Name: "Implementation",
		Type: "FHIR.CapabilityStatementImplementation",
	}, {
		Name: "FhirVersion",
		Type: "FHIR.Code",
	}, {
		Name: "Format",
		Type: "List<FHIR.Code>",
	}, {
		Name: "PatchFormat",
		Type: "List<FHIR.Code>",
	}, {
		Name: "ImplementationGuide",
		Type: "List<FHIR.Canonical>",
	}, {
		Name: "Rest",
		Type: "List<FHIR.CapabilityStatementRest>",
	}, {
		Name: "Messaging",
		Type: "List<FHIR.CapabilityStatementMessaging>",
	}, {
		Name: "Document",
		Type: "List<FHIR.CapabilityStatementDocument>",
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
		Type: "FHIR.Id",
	}, {
		Name: "Meta",
		Type: "FHIR.Meta",
	}, {
		Name: "ImplicitRules",
		Type: "FHIR.Uri",
	}, {
		Name: "Language",
		Type: "FHIR.Code",
	}, {
		Name: "Text",
		Type: "FHIR.Narrative",
	}, {
		Name: "Contained",
		Type: "List<FHIR.>",
	}, {
		Name: "Extension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "ModifierExtension",
		Type: "List<FHIR.Extension>",
	}, {
		Name: "Issue",
		Type: "List<FHIR.OperationOutcomeIssue>",
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
