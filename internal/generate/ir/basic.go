package ir

func CollectBasic(r4 ResourceOrType, other []ResourceOrType) ResourceOrType {
	var basicStructs []Struct
	collectBasicFields(r4.Structs[0], r4, other, &basicStructs)
	return ResourceOrType{
		Name:        r4.Name,
		FileName:    r4.FileName,
		IsResource:  r4.IsResource,
		IsPrimitive: r4.IsPrimitive,
		Structs:     basicStructs,
	}
}

func collectBasicFields(s Struct, r4 ResourceOrType, other []ResourceOrType, structCollector *[]Struct) {
	// break recursive loops
	for _, rs := range *structCollector {
		if s.Name == rs.Name {
			return
		}
	}

	structIdx := len(*structCollector)
	*structCollector = append(*structCollector, s)
	var structFields []StructField

	for _, f := range s.Fields {
		if allContainField(s, f, other) {
			structFields = append(structFields, f)
			fieldStructs := findFieldStruct(f, r4.Structs)
			for _, fs := range fieldStructs {
				collectBasicFields(fs, r4, other, structCollector)
			}

		}
	}
	(*structCollector)[structIdx].Fields = structFields
}

func allContainField(s Struct, f StructField, typs []ResourceOrType) bool {
out:
	for _, t := range typs {
		for _, ts := range t.Structs {
			if s.Name == ts.Name {
				for _, tf := range ts.Fields {
					if tf.Name == f.Name {
						continue out
					}
				}
			}
		}

		return false
	}
	return true
}

func findFieldStruct(f StructField, structs []Struct) []Struct {
	var foundStructs []Struct
	for _, t := range f.PossibleTypes {
		for _, s := range structs {
			if t.Name == s.Name {
				foundStructs = append(foundStructs, s)
			}
		}
	}
	return foundStructs
}
