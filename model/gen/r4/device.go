package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
//
// Allows institutions to track their devices.
type Device struct {
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
	// Unique instance identifiers assigned to a device by manufacturers other organizations or owners.
	Identifier []Identifier
	// The reference to the definition for the device.
	Definition *Reference
	// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
	UdiCarrier []DeviceUdiCarrier
	// Status of the Device availability.
	Status *Code
	// Reason for the dtatus of the Device availability.
	StatusReason []CodeableConcept
	// The distinct identification string as required by regulation for a human cell, tissue, or cellular and tissue-based product.
	DistinctIdentifier *String
	// A name of the manufacturer.
	Manufacturer *String
	// The date and time when the device was manufactured.
	ManufactureDate *DateTime
	// The date and time beyond which this device is no longer valid or should not be used (if applicable).
	ExpirationDate *DateTime
	// Lot number assigned by the manufacturer.
	LotNumber *String
	// The serial number assigned by the organization when the device was manufactured.
	SerialNumber *String
	// This represents the manufacturer's name of the device as provided by the device, from a UDI label, or by a person describing the Device.  This typically would be used when a person provides the name(s) or when the device represents one of the names available from DeviceDefinition.
	DeviceName []DeviceDeviceName
	// The model number for the device.
	ModelNumber *String
	// The part number of the device.
	PartNumber *String
	// The kind or type of device.
	Type *CodeableConcept
	// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
	Specialization []DeviceSpecialization
	// The actual design of the device or software version running on the device.
	Version []DeviceVersion
	// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
	Property []DeviceProperty
	// Patient information, If the device is affixed to a person.
	Patient *Reference
	// An organization that is responsible for the provision and ongoing maintenance of the device.
	Owner *Reference
	// Contact details for an organization or a particular human that is responsible for the device.
	Contact []ContactPoint
	// The place where the device can be found.
	Location *Reference
	// A network address on which the device may be contacted directly.
	Url *Uri
	// Descriptive information, usage information or implantation information that is not captured in an existing element.
	Note []Annotation
	// Provides additional safety characteristics about a medical device.  For example devices containing latex.
	Safety []CodeableConcept
	// The parent device.
	Parent *Reference
}

func (r Device) ResourceType() string {
	return "Device"
}

type jsonDevice struct {
	ResourceType                       string                 `json:"resourceType"`
	Id                                 *Id                    `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement      `json:"_id,omitempty"`
	Meta                               *Meta                  `json:"meta,omitempty"`
	ImplicitRules                      *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                           *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement      `json:"_language,omitempty"`
	Text                               *Narrative             `json:"text,omitempty"`
	Contained                          []containedResource    `json:"contained,omitempty"`
	Extension                          []Extension            `json:"extension,omitempty"`
	ModifierExtension                  []Extension            `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier           `json:"identifier,omitempty"`
	Definition                         *Reference             `json:"definition,omitempty"`
	UdiCarrier                         []DeviceUdiCarrier     `json:"udiCarrier,omitempty"`
	Status                             *Code                  `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement      `json:"_status,omitempty"`
	StatusReason                       []CodeableConcept      `json:"statusReason,omitempty"`
	DistinctIdentifier                 *String                `json:"distinctIdentifier,omitempty"`
	DistinctIdentifierPrimitiveElement *primitiveElement      `json:"_distinctIdentifier,omitempty"`
	Manufacturer                       *String                `json:"manufacturer,omitempty"`
	ManufacturerPrimitiveElement       *primitiveElement      `json:"_manufacturer,omitempty"`
	ManufactureDate                    *DateTime              `json:"manufactureDate,omitempty"`
	ManufactureDatePrimitiveElement    *primitiveElement      `json:"_manufactureDate,omitempty"`
	ExpirationDate                     *DateTime              `json:"expirationDate,omitempty"`
	ExpirationDatePrimitiveElement     *primitiveElement      `json:"_expirationDate,omitempty"`
	LotNumber                          *String                `json:"lotNumber,omitempty"`
	LotNumberPrimitiveElement          *primitiveElement      `json:"_lotNumber,omitempty"`
	SerialNumber                       *String                `json:"serialNumber,omitempty"`
	SerialNumberPrimitiveElement       *primitiveElement      `json:"_serialNumber,omitempty"`
	DeviceName                         []DeviceDeviceName     `json:"deviceName,omitempty"`
	ModelNumber                        *String                `json:"modelNumber,omitempty"`
	ModelNumberPrimitiveElement        *primitiveElement      `json:"_modelNumber,omitempty"`
	PartNumber                         *String                `json:"partNumber,omitempty"`
	PartNumberPrimitiveElement         *primitiveElement      `json:"_partNumber,omitempty"`
	Type                               *CodeableConcept       `json:"type,omitempty"`
	Specialization                     []DeviceSpecialization `json:"specialization,omitempty"`
	Version                            []DeviceVersion        `json:"version,omitempty"`
	Property                           []DeviceProperty       `json:"property,omitempty"`
	Patient                            *Reference             `json:"patient,omitempty"`
	Owner                              *Reference             `json:"owner,omitempty"`
	Contact                            []ContactPoint         `json:"contact,omitempty"`
	Location                           *Reference             `json:"location,omitempty"`
	Url                                *Uri                   `json:"url,omitempty"`
	UrlPrimitiveElement                *primitiveElement      `json:"_url,omitempty"`
	Note                               []Annotation           `json:"note,omitempty"`
	Safety                             []CodeableConcept      `json:"safety,omitempty"`
	Parent                             *Reference             `json:"parent,omitempty"`
}

func (r Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Device) marshalJSON() jsonDevice {
	m := jsonDevice{}
	m.ResourceType = "Device"
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
	m.Definition = r.Definition
	m.UdiCarrier = r.UdiCarrier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.DistinctIdentifier = r.DistinctIdentifier
	if r.DistinctIdentifier != nil && (r.DistinctIdentifier.Id != nil || r.DistinctIdentifier.Extension != nil) {
		m.DistinctIdentifierPrimitiveElement = &primitiveElement{Id: r.DistinctIdentifier.Id, Extension: r.DistinctIdentifier.Extension}
	}
	m.Manufacturer = r.Manufacturer
	if r.Manufacturer != nil && (r.Manufacturer.Id != nil || r.Manufacturer.Extension != nil) {
		m.ManufacturerPrimitiveElement = &primitiveElement{Id: r.Manufacturer.Id, Extension: r.Manufacturer.Extension}
	}
	m.ManufactureDate = r.ManufactureDate
	if r.ManufactureDate != nil && (r.ManufactureDate.Id != nil || r.ManufactureDate.Extension != nil) {
		m.ManufactureDatePrimitiveElement = &primitiveElement{Id: r.ManufactureDate.Id, Extension: r.ManufactureDate.Extension}
	}
	m.ExpirationDate = r.ExpirationDate
	if r.ExpirationDate != nil && (r.ExpirationDate.Id != nil || r.ExpirationDate.Extension != nil) {
		m.ExpirationDatePrimitiveElement = &primitiveElement{Id: r.ExpirationDate.Id, Extension: r.ExpirationDate.Extension}
	}
	m.LotNumber = r.LotNumber
	if r.LotNumber != nil && (r.LotNumber.Id != nil || r.LotNumber.Extension != nil) {
		m.LotNumberPrimitiveElement = &primitiveElement{Id: r.LotNumber.Id, Extension: r.LotNumber.Extension}
	}
	m.SerialNumber = r.SerialNumber
	if r.SerialNumber != nil && (r.SerialNumber.Id != nil || r.SerialNumber.Extension != nil) {
		m.SerialNumberPrimitiveElement = &primitiveElement{Id: r.SerialNumber.Id, Extension: r.SerialNumber.Extension}
	}
	m.DeviceName = r.DeviceName
	m.ModelNumber = r.ModelNumber
	if r.ModelNumber != nil && (r.ModelNumber.Id != nil || r.ModelNumber.Extension != nil) {
		m.ModelNumberPrimitiveElement = &primitiveElement{Id: r.ModelNumber.Id, Extension: r.ModelNumber.Extension}
	}
	m.PartNumber = r.PartNumber
	if r.PartNumber != nil && (r.PartNumber.Id != nil || r.PartNumber.Extension != nil) {
		m.PartNumberPrimitiveElement = &primitiveElement{Id: r.PartNumber.Id, Extension: r.PartNumber.Extension}
	}
	m.Type = r.Type
	m.Specialization = r.Specialization
	m.Version = r.Version
	m.Property = r.Property
	m.Patient = r.Patient
	m.Owner = r.Owner
	m.Contact = r.Contact
	m.Location = r.Location
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Note = r.Note
	m.Safety = r.Safety
	m.Parent = r.Parent
	return m
}
func (r *Device) UnmarshalJSON(b []byte) error {
	var m jsonDevice
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Device) unmarshalJSON(m jsonDevice) error {
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
	r.Definition = m.Definition
	r.UdiCarrier = m.UdiCarrier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.DistinctIdentifier = m.DistinctIdentifier
	if m.DistinctIdentifierPrimitiveElement != nil {
		r.DistinctIdentifier.Id = m.DistinctIdentifierPrimitiveElement.Id
		r.DistinctIdentifier.Extension = m.DistinctIdentifierPrimitiveElement.Extension
	}
	r.Manufacturer = m.Manufacturer
	if m.ManufacturerPrimitiveElement != nil {
		r.Manufacturer.Id = m.ManufacturerPrimitiveElement.Id
		r.Manufacturer.Extension = m.ManufacturerPrimitiveElement.Extension
	}
	r.ManufactureDate = m.ManufactureDate
	if m.ManufactureDatePrimitiveElement != nil {
		r.ManufactureDate.Id = m.ManufactureDatePrimitiveElement.Id
		r.ManufactureDate.Extension = m.ManufactureDatePrimitiveElement.Extension
	}
	r.ExpirationDate = m.ExpirationDate
	if m.ExpirationDatePrimitiveElement != nil {
		r.ExpirationDate.Id = m.ExpirationDatePrimitiveElement.Id
		r.ExpirationDate.Extension = m.ExpirationDatePrimitiveElement.Extension
	}
	r.LotNumber = m.LotNumber
	if m.LotNumberPrimitiveElement != nil {
		r.LotNumber.Id = m.LotNumberPrimitiveElement.Id
		r.LotNumber.Extension = m.LotNumberPrimitiveElement.Extension
	}
	r.SerialNumber = m.SerialNumber
	if m.SerialNumberPrimitiveElement != nil {
		r.SerialNumber.Id = m.SerialNumberPrimitiveElement.Id
		r.SerialNumber.Extension = m.SerialNumberPrimitiveElement.Extension
	}
	r.DeviceName = m.DeviceName
	r.ModelNumber = m.ModelNumber
	if m.ModelNumberPrimitiveElement != nil {
		r.ModelNumber.Id = m.ModelNumberPrimitiveElement.Id
		r.ModelNumber.Extension = m.ModelNumberPrimitiveElement.Extension
	}
	r.PartNumber = m.PartNumber
	if m.PartNumberPrimitiveElement != nil {
		r.PartNumber.Id = m.PartNumberPrimitiveElement.Id
		r.PartNumber.Extension = m.PartNumberPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Specialization = m.Specialization
	r.Version = m.Version
	r.Property = m.Property
	r.Patient = m.Patient
	r.Owner = m.Owner
	r.Contact = m.Contact
	r.Location = m.Location
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Note = m.Note
	r.Safety = m.Safety
	r.Parent = m.Parent
	return nil
}
func (r Device) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
type DeviceUdiCarrier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The device identifier (DI) is a mandatory, fixed portion of a UDI that identifies the labeler and the specific version or model of a device.
	DeviceIdentifier *String
	// Organization that is charged with issuing UDIs for devices.  For example, the US FDA issuers include :
	// 1) GS1:
	// http://hl7.org/fhir/NamingSystem/gs1-di,
	// 2) HIBCC:
	// http://hl7.org/fhir/NamingSystem/hibcc-dI,
	// 3) ICCBBA for blood containers:
	// http://hl7.org/fhir/NamingSystem/iccbba-blood-di,
	// 4) ICCBA for other devices:
	// http://hl7.org/fhir/NamingSystem/iccbba-other-di.
	Issuer *Uri
	// The identity of the authoritative source for UDI generation within a  jurisdiction.  All UDIs are globally unique within a single namespace with the appropriate repository uri as the system.  For example,  UDIs of devices managed in the U.S. by the FDA, the value is  http://hl7.org/fhir/NamingSystem/fda-udi.
	Jurisdiction *Uri
	// The full UDI carrier of the Automatic Identification and Data Capture (AIDC) technology representation of the barcode string as printed on the packaging of the device - e.g., a barcode or RFID.   Because of limitations on character sets in XML and the need to round-trip JSON data through XML, AIDC Formats *SHALL* be base64 encoded.
	CarrierAidc *Base64Binary
	// The full UDI carrier as the human readable form (HRF) representation of the barcode string as printed on the packaging of the device.
	CarrierHrf *String
	// A coded entry to indicate how the data was entered.
	EntryType *Code
}
type jsonDeviceUdiCarrier struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	DeviceIdentifier                 *String           `json:"deviceIdentifier,omitempty"`
	DeviceIdentifierPrimitiveElement *primitiveElement `json:"_deviceIdentifier,omitempty"`
	Issuer                           *Uri              `json:"issuer,omitempty"`
	IssuerPrimitiveElement           *primitiveElement `json:"_issuer,omitempty"`
	Jurisdiction                     *Uri              `json:"jurisdiction,omitempty"`
	JurisdictionPrimitiveElement     *primitiveElement `json:"_jurisdiction,omitempty"`
	CarrierAidc                      *Base64Binary     `json:"carrierAIDC,omitempty"`
	CarrierAidcPrimitiveElement      *primitiveElement `json:"_carrierAIDC,omitempty"`
	CarrierHrf                       *String           `json:"carrierHRF,omitempty"`
	CarrierHrfPrimitiveElement       *primitiveElement `json:"_carrierHRF,omitempty"`
	EntryType                        *Code             `json:"entryType,omitempty"`
	EntryTypePrimitiveElement        *primitiveElement `json:"_entryType,omitempty"`
}

func (r DeviceUdiCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceUdiCarrier) marshalJSON() jsonDeviceUdiCarrier {
	m := jsonDeviceUdiCarrier{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.DeviceIdentifier = r.DeviceIdentifier
	if r.DeviceIdentifier != nil && (r.DeviceIdentifier.Id != nil || r.DeviceIdentifier.Extension != nil) {
		m.DeviceIdentifierPrimitiveElement = &primitiveElement{Id: r.DeviceIdentifier.Id, Extension: r.DeviceIdentifier.Extension}
	}
	m.Issuer = r.Issuer
	if r.Issuer != nil && (r.Issuer.Id != nil || r.Issuer.Extension != nil) {
		m.IssuerPrimitiveElement = &primitiveElement{Id: r.Issuer.Id, Extension: r.Issuer.Extension}
	}
	m.Jurisdiction = r.Jurisdiction
	if r.Jurisdiction != nil && (r.Jurisdiction.Id != nil || r.Jurisdiction.Extension != nil) {
		m.JurisdictionPrimitiveElement = &primitiveElement{Id: r.Jurisdiction.Id, Extension: r.Jurisdiction.Extension}
	}
	m.CarrierAidc = r.CarrierAidc
	if r.CarrierAidc != nil && (r.CarrierAidc.Id != nil || r.CarrierAidc.Extension != nil) {
		m.CarrierAidcPrimitiveElement = &primitiveElement{Id: r.CarrierAidc.Id, Extension: r.CarrierAidc.Extension}
	}
	m.CarrierHrf = r.CarrierHrf
	if r.CarrierHrf != nil && (r.CarrierHrf.Id != nil || r.CarrierHrf.Extension != nil) {
		m.CarrierHrfPrimitiveElement = &primitiveElement{Id: r.CarrierHrf.Id, Extension: r.CarrierHrf.Extension}
	}
	m.EntryType = r.EntryType
	if r.EntryType != nil && (r.EntryType.Id != nil || r.EntryType.Extension != nil) {
		m.EntryTypePrimitiveElement = &primitiveElement{Id: r.EntryType.Id, Extension: r.EntryType.Extension}
	}
	return m
}
func (r *DeviceUdiCarrier) UnmarshalJSON(b []byte) error {
	var m jsonDeviceUdiCarrier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceUdiCarrier) unmarshalJSON(m jsonDeviceUdiCarrier) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DeviceIdentifier = m.DeviceIdentifier
	if m.DeviceIdentifierPrimitiveElement != nil {
		r.DeviceIdentifier.Id = m.DeviceIdentifierPrimitiveElement.Id
		r.DeviceIdentifier.Extension = m.DeviceIdentifierPrimitiveElement.Extension
	}
	r.Issuer = m.Issuer
	if m.IssuerPrimitiveElement != nil {
		r.Issuer.Id = m.IssuerPrimitiveElement.Id
		r.Issuer.Extension = m.IssuerPrimitiveElement.Extension
	}
	r.Jurisdiction = m.Jurisdiction
	if m.JurisdictionPrimitiveElement != nil {
		r.Jurisdiction.Id = m.JurisdictionPrimitiveElement.Id
		r.Jurisdiction.Extension = m.JurisdictionPrimitiveElement.Extension
	}
	r.CarrierAidc = m.CarrierAidc
	if m.CarrierAidcPrimitiveElement != nil {
		r.CarrierAidc.Id = m.CarrierAidcPrimitiveElement.Id
		r.CarrierAidc.Extension = m.CarrierAidcPrimitiveElement.Extension
	}
	r.CarrierHrf = m.CarrierHrf
	if m.CarrierHrfPrimitiveElement != nil {
		r.CarrierHrf.Id = m.CarrierHrfPrimitiveElement.Id
		r.CarrierHrf.Extension = m.CarrierHrfPrimitiveElement.Extension
	}
	r.EntryType = m.EntryType
	if m.EntryTypePrimitiveElement != nil {
		r.EntryType.Id = m.EntryTypePrimitiveElement.Id
		r.EntryType.Extension = m.EntryTypePrimitiveElement.Extension
	}
	return nil
}
func (r DeviceUdiCarrier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This represents the manufacturer's name of the device as provided by the device, from a UDI label, or by a person describing the Device.  This typically would be used when a person provides the name(s) or when the device represents one of the names available from DeviceDefinition.
type DeviceDeviceName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the device.
	Name String
	// The type of deviceName.
	// UDILabelName | UserFriendlyName | PatientReportedName | ManufactureDeviceName | ModelName.
	Type Code
}
type jsonDeviceDeviceName struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Name                 String            `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
}

func (r DeviceDeviceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDeviceName) marshalJSON() jsonDeviceDeviceName {
	m := jsonDeviceDeviceName{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	return m
}
func (r *DeviceDeviceName) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDeviceName
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDeviceName) unmarshalJSON(m jsonDeviceDeviceName) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	return nil
}
func (r DeviceDeviceName) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
type DeviceSpecialization struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The standard that is used to operate and communicate.
	SystemType CodeableConcept
	// The version of the standard that is used to operate and communicate.
	Version *String
}
type jsonDeviceSpecialization struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	SystemType              CodeableConcept   `json:"systemType,omitempty"`
	Version                 *String           `json:"version,omitempty"`
	VersionPrimitiveElement *primitiveElement `json:"_version,omitempty"`
}

func (r DeviceSpecialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceSpecialization) marshalJSON() jsonDeviceSpecialization {
	m := jsonDeviceSpecialization{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.SystemType = r.SystemType
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	return m
}
func (r *DeviceSpecialization) UnmarshalJSON(b []byte) error {
	var m jsonDeviceSpecialization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceSpecialization) unmarshalJSON(m jsonDeviceSpecialization) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.SystemType = m.SystemType
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	return nil
}
func (r DeviceSpecialization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actual design of the device or software version running on the device.
type DeviceVersion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of the device version.
	Type *CodeableConcept
	// A single component of the device version.
	Component *Identifier
	// The version text.
	Value String
}
type jsonDeviceVersion struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Type                  *CodeableConcept  `json:"type,omitempty"`
	Component             *Identifier       `json:"component,omitempty"`
	Value                 String            `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r DeviceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceVersion) marshalJSON() jsonDeviceVersion {
	m := jsonDeviceVersion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Component = r.Component
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *DeviceVersion) UnmarshalJSON(b []byte) error {
	var m jsonDeviceVersion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceVersion) unmarshalJSON(m jsonDeviceVersion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Component = m.Component
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r DeviceVersion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
type DeviceProperty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code that specifies the property DeviceDefinitionPropetyCode (Extensible).
	Type CodeableConcept
	// Property value as a quantity.
	ValueQuantity []Quantity
	// Property value as a code, e.g., NTP4 (synced to NTP).
	ValueCode []CodeableConcept
}
type jsonDeviceProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type,omitempty"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

func (r DeviceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceProperty) marshalJSON() jsonDeviceProperty {
	m := jsonDeviceProperty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.ValueQuantity = r.ValueQuantity
	m.ValueCode = r.ValueCode
	return m
}
func (r *DeviceProperty) UnmarshalJSON(b []byte) error {
	var m jsonDeviceProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceProperty) unmarshalJSON(m jsonDeviceProperty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.ValueQuantity = m.ValueQuantity
	r.ValueCode = m.ValueCode
	return nil
}
func (r DeviceProperty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
