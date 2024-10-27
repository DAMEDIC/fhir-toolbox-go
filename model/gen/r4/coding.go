package r4

import "encoding/json"

// Base StructureDefinition for Coding Type: A reference to a code defined by a terminology system.
//
// References to codes are very common in healthcare models.
type Coding struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The identification of the code system that defines the meaning of the symbol in the code.
	System *Uri
	// The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Version *String
	// A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Code *Code
	// A representation of the meaning of the code in the system, following the rules of the system.
	Display *String
	// Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
	UserSelected *Boolean
}
type jsonCoding struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	System                       *Uri              `json:"system,omitempty"`
	SystemPrimitiveElement       *primitiveElement `json:"_system,omitempty"`
	Version                      *String           `json:"version,omitempty"`
	VersionPrimitiveElement      *primitiveElement `json:"_version,omitempty"`
	Code                         *Code             `json:"code,omitempty"`
	CodePrimitiveElement         *primitiveElement `json:"_code,omitempty"`
	Display                      *String           `json:"display,omitempty"`
	DisplayPrimitiveElement      *primitiveElement `json:"_display,omitempty"`
	UserSelected                 *Boolean          `json:"userSelected,omitempty"`
	UserSelectedPrimitiveElement *primitiveElement `json:"_userSelected,omitempty"`
}

func (r Coding) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Coding) marshalJSON() jsonCoding {
	m := jsonCoding{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.System != nil && r.System.Value != nil {
		m.System = r.System
	}
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Code != nil && r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Display != nil && r.Display.Value != nil {
		m.Display = r.Display
	}
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	if r.UserSelected != nil && r.UserSelected.Value != nil {
		m.UserSelected = r.UserSelected
	}
	if r.UserSelected != nil && (r.UserSelected.Id != nil || r.UserSelected.Extension != nil) {
		m.UserSelectedPrimitiveElement = &primitiveElement{Id: r.UserSelected.Id, Extension: r.UserSelected.Extension}
	}
	return m
}
func (r *Coding) UnmarshalJSON(b []byte) error {
	var m jsonCoding
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Coding) unmarshalJSON(m jsonCoding) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		if r.System == nil {
			r.System = &Uri{}
		}
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		if r.Code == nil {
			r.Code = &Code{}
		}
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		if r.Display == nil {
			r.Display = &String{}
		}
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.UserSelected = m.UserSelected
	if m.UserSelectedPrimitiveElement != nil {
		if r.UserSelected == nil {
			r.UserSelected = &Boolean{}
		}
		r.UserSelected.Id = m.UserSelectedPrimitiveElement.Id
		r.UserSelected.Extension = m.UserSelectedPrimitiveElement.Extension
	}
	return nil
}
func (r Coding) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
