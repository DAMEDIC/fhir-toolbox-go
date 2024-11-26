// Package model contains a simplified representation of the FHIR model,
// necessary for initially bootstrapping the generator.
//
// Future versions may switch to using the generated structures for an earlier version of this module.
package model

import (
	"encoding/json"
)

type Bundle struct {
	Entry []BundleEntry `json:"entry"`
}

type BundleEntry struct {
	Resource Resource `json:"resource"`
}

type Resource interface{}

func (e *BundleEntry) UnmarshalJSON(data []byte) error {
	var t struct {
		R struct {
			T string `json:"resourceType"`
		} `json:"resource"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	switch t.R.T {
	case "StructureDefinition":
		var r struct {
			R StructureDefinition `json:"resource"`
		}
		err := json.Unmarshal(data, &r)
		if err != nil {
			return err
		}
		e.Resource = &r.R
	}
	return nil
}

type StructureDefinition struct {
	URL         string                      `json:"url"`
	Description string                      `json:"description"`
	Purpose     string                      `json:"purpose"`
	Kind        string                      `json:"kind"`
	Abstract    bool                        `json:"abstract"`
	Type        string                      `json:"type"`
	Name        string                      `json:"name"`
	Snapshot    StructureDefinitionSnapshot `json:"snapshot"`
}

type StructureDefinitionSnapshot struct {
	Element []ElementDefinition `json:"element"`
}

type ElementDefinition struct {
	ID               string                  `json:"id"`
	Short            string                  `json:"short"`
	Definition       string                  `json:"definition"`
	Comment          string                  `json:"comment"`
	Path             string                  `json:"path"`
	Min              uint32                  `json:"min"`
	Max              string                  `json:"max"`
	Type             []ElementDefinitionType `json:"type"`
	ContentReference string                  `json:"contentReference"`
}

type ElementDefinitionType struct {
	Code string `json:"code"`
}
