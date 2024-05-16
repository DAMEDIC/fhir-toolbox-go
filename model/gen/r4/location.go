package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
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
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
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
	Contained                              []containedResource        `json:"contained,omitempty"`
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
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.OperationalStatus = r.OperationalStatus
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Alias = r.Alias
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
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Mode = r.Mode
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
	m.AvailabilityExceptions = r.AvailabilityExceptions
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
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.OperationalStatus = m.OperationalStatus
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	for i, e := range m.AliasPrimitiveElement {
		if len(r.Alias) > i {
			r.Alias[i].Id = e.Id
			r.Alias[i].Extension = e.Extension
		} else {
			r.Alias = append(r.Alias, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
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
		r.AvailabilityExceptions.Id = m.AvailabilityExceptionsPrimitiveElement.Id
		r.AvailabilityExceptions.Extension = m.AvailabilityExceptionsPrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	return nil
}
func (r Location) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.Longitude = r.Longitude
	if r.Longitude.Id != nil || r.Longitude.Extension != nil {
		m.LongitudePrimitiveElement = &primitiveElement{Id: r.Longitude.Id, Extension: r.Longitude.Extension}
	}
	m.Latitude = r.Latitude
	if r.Latitude.Id != nil || r.Latitude.Extension != nil {
		m.LatitudePrimitiveElement = &primitiveElement{Id: r.Latitude.Id, Extension: r.Latitude.Extension}
	}
	m.Altitude = r.Altitude
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
		r.Altitude.Id = m.AltitudePrimitiveElement.Id
		r.Altitude.Extension = m.AltitudePrimitiveElement.Extension
	}
	return nil
}
func (r LocationPosition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.DaysOfWeek = r.DaysOfWeek
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
	m.AllDay = r.AllDay
	if r.AllDay != nil && (r.AllDay.Id != nil || r.AllDay.Extension != nil) {
		m.AllDayPrimitiveElement = &primitiveElement{Id: r.AllDay.Id, Extension: r.AllDay.Extension}
	}
	m.OpeningTime = r.OpeningTime
	if r.OpeningTime != nil && (r.OpeningTime.Id != nil || r.OpeningTime.Extension != nil) {
		m.OpeningTimePrimitiveElement = &primitiveElement{Id: r.OpeningTime.Id, Extension: r.OpeningTime.Extension}
	}
	m.ClosingTime = r.ClosingTime
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
		if len(r.DaysOfWeek) > i {
			r.DaysOfWeek[i].Id = e.Id
			r.DaysOfWeek[i].Extension = e.Extension
		} else {
			r.DaysOfWeek = append(r.DaysOfWeek, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.AllDay = m.AllDay
	if m.AllDayPrimitiveElement != nil {
		r.AllDay.Id = m.AllDayPrimitiveElement.Id
		r.AllDay.Extension = m.AllDayPrimitiveElement.Extension
	}
	r.OpeningTime = m.OpeningTime
	if m.OpeningTimePrimitiveElement != nil {
		r.OpeningTime.Id = m.OpeningTimePrimitiveElement.Id
		r.OpeningTime.Extension = m.OpeningTimePrimitiveElement.Extension
	}
	r.ClosingTime = m.ClosingTime
	if m.ClosingTimePrimitiveElement != nil {
		r.ClosingTime.Id = m.ClosingTimePrimitiveElement.Id
		r.ClosingTime.Extension = m.ClosingTimePrimitiveElement.Extension
	}
	return nil
}
func (r LocationHoursOfOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
