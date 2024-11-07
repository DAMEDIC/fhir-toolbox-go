package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Details and position information for a physical place where services are provided and resources and participants may be stored, found, contained, or accommodated.
type Location struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique code or number identifying the location to its users.
	Identifier []Identifier
	// The status property covers the general availability of the resource, not the current value which may be covered by the operationStatus, or by a schedule/slots if they are configured for the location.
	Status *Code
	// The operational status covers operation values most relevant to beds (but can also apply to rooms/units/chairs/etc. such as an isolation unit/dialysis chair). This typically covers concepts such as contamination, housekeeping, and other activities like maintenance.
	OperationalStatus *Coding
	// Name of the location as used by humans. Does not need to be unique.
	Name *String
	// A list of alternate names that the location is known as, or was known as, in the past.
	Alias []String
	// Description of the Location, which helps in finding or referencing the place.
	Description *String
	// Indicates whether a resource instance represents a specific location or a class of locations.
	Mode *Code
	// Indicates the type of function performed at the location.
	Type []CodeableConcept
	// The contact details of communication devices available at the location. This can include phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	Telecom []ContactPoint
	// Physical location.
	Address *Address
	// Physical form of the location, e.g. building, room, vehicle, road.
	PhysicalType *CodeableConcept
	// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
	Position *LocationPosition
	// The organization responsible for the provisioning and upkeep of the location.
	ManagingOrganization *Reference
	// Another Location of which this Location is physically a part of.
	PartOf *Reference
	// What days/times during a week is this location usually open.
	HoursOfOperation []LocationHoursOfOperation
	// A description of when the locations opening ours are different to normal, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as detailed in the opening hours Times.
	AvailabilityExceptions *String
	// Technical endpoints providing access to services operated for the location.
	Endpoint []Reference
}

func (r Location) ResourceType() string {
	return "Location"
}
func (r Location) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonLocation struct {
	ResourceType                           string                     `json:"resourceType"`
	Id                                     *Id                        `json:"id,omitempty"`
	IdPrimitiveElement                     *primitiveElement          `json:"_id,omitempty"`
	Meta                                   *Meta                      `json:"meta,omitempty"`
	ImplicitRules                          *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement          *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                               *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement               *primitiveElement          `json:"_language,omitempty"`
	Text                                   *Narrative                 `json:"text,omitempty"`
	Contained                              []ContainedResource        `json:"contained,omitempty"`
	Extension                              []Extension                `json:"extension,omitempty"`
	ModifierExtension                      []Extension                `json:"modifierExtension,omitempty"`
	Identifier                             []Identifier               `json:"identifier,omitempty"`
	Status                                 *Code                      `json:"status,omitempty"`
	StatusPrimitiveElement                 *primitiveElement          `json:"_status,omitempty"`
	OperationalStatus                      *Coding                    `json:"operationalStatus,omitempty"`
	Name                                   *String                    `json:"name,omitempty"`
	NamePrimitiveElement                   *primitiveElement          `json:"_name,omitempty"`
	Alias                                  []String                   `json:"alias,omitempty"`
	AliasPrimitiveElement                  []*primitiveElement        `json:"_alias,omitempty"`
	Description                            *String                    `json:"description,omitempty"`
	DescriptionPrimitiveElement            *primitiveElement          `json:"_description,omitempty"`
	Mode                                   *Code                      `json:"mode,omitempty"`
	ModePrimitiveElement                   *primitiveElement          `json:"_mode,omitempty"`
	Type                                   []CodeableConcept          `json:"type,omitempty"`
	Telecom                                []ContactPoint             `json:"telecom,omitempty"`
	Address                                *Address                   `json:"address,omitempty"`
	PhysicalType                           *CodeableConcept           `json:"physicalType,omitempty"`
	Position                               *LocationPosition          `json:"position,omitempty"`
	ManagingOrganization                   *Reference                 `json:"managingOrganization,omitempty"`
	PartOf                                 *Reference                 `json:"partOf,omitempty"`
	HoursOfOperation                       []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`
	AvailabilityExceptions                 *String                    `json:"availabilityExceptions,omitempty"`
	AvailabilityExceptionsPrimitiveElement *primitiveElement          `json:"_availabilityExceptions,omitempty"`
	Endpoint                               []Reference                `json:"endpoint,omitempty"`
}

func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Location) marshalJSON() jsonLocation {
	m := jsonLocation{}
	m.ResourceType = "Location"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.OperationalStatus = r.OperationalStatus
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	anyAliasValue := false
	for _, e := range r.Alias {
		if e.Value != nil {
			anyAliasValue = true
			break
		}
	}
	if anyAliasValue {
		m.Alias = r.Alias
	}
	anyAliasIdOrExtension := false
	for _, e := range r.Alias {
		if e.Id != nil || e.Extension != nil {
			anyAliasIdOrExtension = true
			break
		}
	}
	if anyAliasIdOrExtension {
		m.AliasPrimitiveElement = make([]*primitiveElement, 0, len(r.Alias))
		for _, e := range r.Alias {
			if e.Id != nil || e.Extension != nil {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, nil)
			}
		}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Mode != nil && r.Mode.Value != nil {
		m.Mode = r.Mode
	}
	if r.Mode != nil && (r.Mode.Id != nil || r.Mode.Extension != nil) {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Type = r.Type
	m.Telecom = r.Telecom
	m.Address = r.Address
	m.PhysicalType = r.PhysicalType
	m.Position = r.Position
	m.ManagingOrganization = r.ManagingOrganization
	m.PartOf = r.PartOf
	m.HoursOfOperation = r.HoursOfOperation
	if r.AvailabilityExceptions != nil && r.AvailabilityExceptions.Value != nil {
		m.AvailabilityExceptions = r.AvailabilityExceptions
	}
	if r.AvailabilityExceptions != nil && (r.AvailabilityExceptions.Id != nil || r.AvailabilityExceptions.Extension != nil) {
		m.AvailabilityExceptionsPrimitiveElement = &primitiveElement{Id: r.AvailabilityExceptions.Id, Extension: r.AvailabilityExceptions.Extension}
	}
	m.Endpoint = r.Endpoint
	return m
}
func (r *Location) UnmarshalJSON(b []byte) error {
	var m jsonLocation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Location) unmarshalJSON(m jsonLocation) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		if r.Status == nil {
			r.Status = &Code{}
		}
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.OperationalStatus = m.OperationalStatus
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	for i, e := range m.AliasPrimitiveElement {
		if len(r.Alias) <= i {
			r.Alias = append(r.Alias, String{})
		}
		if e != nil {
			r.Alias[i].Id = e.Id
			r.Alias[i].Extension = e.Extension
		}
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		if r.Mode == nil {
			r.Mode = &Code{}
		}
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Telecom = m.Telecom
	r.Address = m.Address
	r.PhysicalType = m.PhysicalType
	r.Position = m.Position
	r.ManagingOrganization = m.ManagingOrganization
	r.PartOf = m.PartOf
	r.HoursOfOperation = m.HoursOfOperation
	r.AvailabilityExceptions = m.AvailabilityExceptions
	if m.AvailabilityExceptionsPrimitiveElement != nil {
		if r.AvailabilityExceptions == nil {
			r.AvailabilityExceptions = &String{}
		}
		r.AvailabilityExceptions.Id = m.AvailabilityExceptionsPrimitiveElement.Id
		r.AvailabilityExceptions.Extension = m.AvailabilityExceptionsPrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	return nil
}
func (r Location) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Location"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OperationalStatus, xml.StartElement{Name: xml.Name{Local: "operationalStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Alias, xml.StartElement{Name: xml.Name{Local: "alias"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PhysicalType, xml.StartElement{Name: xml.Name{Local: "physicalType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Position, xml.StartElement{Name: xml.Name{Local: "position"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ManagingOrganization, xml.StartElement{Name: xml.Name{Local: "managingOrganization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartOf, xml.StartElement{Name: xml.Name{Local: "partOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HoursOfOperation, xml.StartElement{Name: xml.Name{Local: "hoursOfOperation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailabilityExceptions, xml.StartElement{Name: xml.Name{Local: "availabilityExceptions"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Endpoint, xml.StartElement{Name: xml.Name{Local: "endpoint"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Location) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "operationalStatus":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OperationalStatus = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "alias":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Alias = append(r.Alias, v)
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "mode":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Mode = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "address":
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = &v
			case "physicalType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PhysicalType = &v
			case "position":
				var v LocationPosition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Position = &v
			case "managingOrganization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ManagingOrganization = &v
			case "partOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartOf = &v
			case "hoursOfOperation":
				var v LocationHoursOfOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HoursOfOperation = append(r.HoursOfOperation, v)
			case "availabilityExceptions":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailabilityExceptions = &v
			case "endpoint":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Endpoint = append(r.Endpoint, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Location) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
type LocationPosition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Longitude. The value domain and the interpretation are the same as for the text of the longitude element in KML (see notes below).
	Longitude Decimal
	// Latitude. The value domain and the interpretation are the same as for the text of the latitude element in KML (see notes below).
	Latitude Decimal
	// Altitude. The value domain and the interpretation are the same as for the text of the altitude element in KML (see notes below).
	Altitude *Decimal
}
type jsonLocationPosition struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Longitude                 Decimal           `json:"longitude,omitempty"`
	LongitudePrimitiveElement *primitiveElement `json:"_longitude,omitempty"`
	Latitude                  Decimal           `json:"latitude,omitempty"`
	LatitudePrimitiveElement  *primitiveElement `json:"_latitude,omitempty"`
	Altitude                  *Decimal          `json:"altitude,omitempty"`
	AltitudePrimitiveElement  *primitiveElement `json:"_altitude,omitempty"`
}

func (r LocationPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r LocationPosition) marshalJSON() jsonLocationPosition {
	m := jsonLocationPosition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Longitude.Value != nil {
		m.Longitude = r.Longitude
	}
	if r.Longitude.Id != nil || r.Longitude.Extension != nil {
		m.LongitudePrimitiveElement = &primitiveElement{Id: r.Longitude.Id, Extension: r.Longitude.Extension}
	}
	if r.Latitude.Value != nil {
		m.Latitude = r.Latitude
	}
	if r.Latitude.Id != nil || r.Latitude.Extension != nil {
		m.LatitudePrimitiveElement = &primitiveElement{Id: r.Latitude.Id, Extension: r.Latitude.Extension}
	}
	if r.Altitude != nil && r.Altitude.Value != nil {
		m.Altitude = r.Altitude
	}
	if r.Altitude != nil && (r.Altitude.Id != nil || r.Altitude.Extension != nil) {
		m.AltitudePrimitiveElement = &primitiveElement{Id: r.Altitude.Id, Extension: r.Altitude.Extension}
	}
	return m
}
func (r *LocationPosition) UnmarshalJSON(b []byte) error {
	var m jsonLocationPosition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *LocationPosition) unmarshalJSON(m jsonLocationPosition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Longitude = m.Longitude
	if m.LongitudePrimitiveElement != nil {
		r.Longitude.Id = m.LongitudePrimitiveElement.Id
		r.Longitude.Extension = m.LongitudePrimitiveElement.Extension
	}
	r.Latitude = m.Latitude
	if m.LatitudePrimitiveElement != nil {
		r.Latitude.Id = m.LatitudePrimitiveElement.Id
		r.Latitude.Extension = m.LatitudePrimitiveElement.Extension
	}
	r.Altitude = m.Altitude
	if m.AltitudePrimitiveElement != nil {
		if r.Altitude == nil {
			r.Altitude = &Decimal{}
		}
		r.Altitude.Id = m.AltitudePrimitiveElement.Id
		r.Altitude.Extension = m.AltitudePrimitiveElement.Extension
	}
	return nil
}
func (r LocationPosition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Longitude, xml.StartElement{Name: xml.Name{Local: "longitude"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Latitude, xml.StartElement{Name: xml.Name{Local: "latitude"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Altitude, xml.StartElement{Name: xml.Name{Local: "altitude"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *LocationPosition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "longitude":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Longitude = v
			case "latitude":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Latitude = v
			case "altitude":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Altitude = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r LocationPosition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// What days/times during a week is this location usually open.
type LocationHoursOfOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []Code
	// The Location is open all day.
	AllDay *Boolean
	// Time that the Location opens.
	OpeningTime *Time
	// Time that the Location closes.
	ClosingTime *Time
}
type jsonLocationHoursOfOperation struct {
	Id                          *string             `json:"id,omitempty"`
	Extension                   []Extension         `json:"extension,omitempty"`
	ModifierExtension           []Extension         `json:"modifierExtension,omitempty"`
	DaysOfWeek                  []Code              `json:"daysOfWeek,omitempty"`
	DaysOfWeekPrimitiveElement  []*primitiveElement `json:"_daysOfWeek,omitempty"`
	AllDay                      *Boolean            `json:"allDay,omitempty"`
	AllDayPrimitiveElement      *primitiveElement   `json:"_allDay,omitempty"`
	OpeningTime                 *Time               `json:"openingTime,omitempty"`
	OpeningTimePrimitiveElement *primitiveElement   `json:"_openingTime,omitempty"`
	ClosingTime                 *Time               `json:"closingTime,omitempty"`
	ClosingTimePrimitiveElement *primitiveElement   `json:"_closingTime,omitempty"`
}

func (r LocationHoursOfOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r LocationHoursOfOperation) marshalJSON() jsonLocationHoursOfOperation {
	m := jsonLocationHoursOfOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyDaysOfWeekValue := false
	for _, e := range r.DaysOfWeek {
		if e.Value != nil {
			anyDaysOfWeekValue = true
			break
		}
	}
	if anyDaysOfWeekValue {
		m.DaysOfWeek = r.DaysOfWeek
	}
	anyDaysOfWeekIdOrExtension := false
	for _, e := range r.DaysOfWeek {
		if e.Id != nil || e.Extension != nil {
			anyDaysOfWeekIdOrExtension = true
			break
		}
	}
	if anyDaysOfWeekIdOrExtension {
		m.DaysOfWeekPrimitiveElement = make([]*primitiveElement, 0, len(r.DaysOfWeek))
		for _, e := range r.DaysOfWeek {
			if e.Id != nil || e.Extension != nil {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, nil)
			}
		}
	}
	if r.AllDay != nil && r.AllDay.Value != nil {
		m.AllDay = r.AllDay
	}
	if r.AllDay != nil && (r.AllDay.Id != nil || r.AllDay.Extension != nil) {
		m.AllDayPrimitiveElement = &primitiveElement{Id: r.AllDay.Id, Extension: r.AllDay.Extension}
	}
	if r.OpeningTime != nil && r.OpeningTime.Value != nil {
		m.OpeningTime = r.OpeningTime
	}
	if r.OpeningTime != nil && (r.OpeningTime.Id != nil || r.OpeningTime.Extension != nil) {
		m.OpeningTimePrimitiveElement = &primitiveElement{Id: r.OpeningTime.Id, Extension: r.OpeningTime.Extension}
	}
	if r.ClosingTime != nil && r.ClosingTime.Value != nil {
		m.ClosingTime = r.ClosingTime
	}
	if r.ClosingTime != nil && (r.ClosingTime.Id != nil || r.ClosingTime.Extension != nil) {
		m.ClosingTimePrimitiveElement = &primitiveElement{Id: r.ClosingTime.Id, Extension: r.ClosingTime.Extension}
	}
	return m
}
func (r *LocationHoursOfOperation) UnmarshalJSON(b []byte) error {
	var m jsonLocationHoursOfOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *LocationHoursOfOperation) unmarshalJSON(m jsonLocationHoursOfOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DaysOfWeek = m.DaysOfWeek
	for i, e := range m.DaysOfWeekPrimitiveElement {
		if len(r.DaysOfWeek) <= i {
			r.DaysOfWeek = append(r.DaysOfWeek, Code{})
		}
		if e != nil {
			r.DaysOfWeek[i].Id = e.Id
			r.DaysOfWeek[i].Extension = e.Extension
		}
	}
	r.AllDay = m.AllDay
	if m.AllDayPrimitiveElement != nil {
		if r.AllDay == nil {
			r.AllDay = &Boolean{}
		}
		r.AllDay.Id = m.AllDayPrimitiveElement.Id
		r.AllDay.Extension = m.AllDayPrimitiveElement.Extension
	}
	r.OpeningTime = m.OpeningTime
	if m.OpeningTimePrimitiveElement != nil {
		if r.OpeningTime == nil {
			r.OpeningTime = &Time{}
		}
		r.OpeningTime.Id = m.OpeningTimePrimitiveElement.Id
		r.OpeningTime.Extension = m.OpeningTimePrimitiveElement.Extension
	}
	r.ClosingTime = m.ClosingTime
	if m.ClosingTimePrimitiveElement != nil {
		if r.ClosingTime == nil {
			r.ClosingTime = &Time{}
		}
		r.ClosingTime.Id = m.ClosingTimePrimitiveElement.Id
		r.ClosingTime.Extension = m.ClosingTimePrimitiveElement.Extension
	}
	return nil
}
func (r LocationHoursOfOperation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DaysOfWeek, xml.StartElement{Name: xml.Name{Local: "daysOfWeek"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AllDay, xml.StartElement{Name: xml.Name{Local: "allDay"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OpeningTime, xml.StartElement{Name: xml.Name{Local: "openingTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClosingTime, xml.StartElement{Name: xml.Name{Local: "closingTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *LocationHoursOfOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "daysOfWeek":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DaysOfWeek = append(r.DaysOfWeek, v)
			case "allDay":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AllDay = &v
			case "openingTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OpeningTime = &v
			case "closingTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClosingTime = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r LocationHoursOfOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
