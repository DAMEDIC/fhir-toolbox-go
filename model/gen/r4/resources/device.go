package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
type DeviceUdiCarrier struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The device identifier (DI) is a mandatory, fixed portion of a UDI that identifies the labeler and the specific version or model of a device.
	DeviceIdentifier *types.String
	// The identity of the authoritative source for UDI generation within a  jurisdiction.  All UDIs are globally unique within a single namespace with the appropriate repository uri as the system.  For example,  UDIs of devices managed in the U.S. by the FDA, the value is  http://hl7.org/fhir/NamingSystem/fda-udi.
	Jurisdiction *types.Uri
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Organization that is charged with issuing UDIs for devices.  For example, the US FDA issuers include :
	// 1) GS1:
	// http://hl7.org/fhir/NamingSystem/gs1-di,
	// 2) HIBCC:
	// http://hl7.org/fhir/NamingSystem/hibcc-dI,
	// 3) ICCBBA for blood containers:
	// http://hl7.org/fhir/NamingSystem/iccbba-blood-di,
	// 4) ICCBA for other devices:
	// http://hl7.org/fhir/NamingSystem/iccbba-other-di.
	Issuer *types.Uri
	// The full UDI carrier of the Automatic Identification and Data Capture (AIDC) technology representation of the barcode string as printed on the packaging of the device - e.g., a barcode or RFID.   Because of limitations on character sets in XML and the need to round-trip JSON data through XML, AIDC Formats *SHALL* be base64 encoded.
	CarrierAidc *types.Base64Binary
	// The full UDI carrier as the human readable form (HRF) representation of the barcode string as printed on the packaging of the device.
	CarrierHrf *types.String
	// A coded entry to indicate how the data was entered.
	EntryType *types.Code
}

func (s DeviceUdiCarrier) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Code that specifies the property DeviceDefinitionPropetyCode (Extensible).
	Type types.CodeableConcept
	// Property value as a quantity.
	ValueQuantity []types.Quantity
	// Property value as a code, e.g., NTP4 (synced to NTP).
	ValueCode []types.CodeableConcept
}

func (s DeviceProperty) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
type DeviceSpecialization struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The standard that is used to operate and communicate.
	SystemType types.CodeableConcept
	// The version of the standard that is used to operate and communicate.
	Version *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s DeviceSpecialization) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actual design of the device or software version running on the device.
type DeviceVersion struct {
	// The version text.
	Value types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of the device version.
	Type *types.CodeableConcept
	// A single component of the device version.
	Component *types.Identifier
}

func (s DeviceVersion) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The name of the device.
	Name types.String
	// The type of deviceName.
	// UDILabelName | UserFriendlyName | PatientReportedName | ManufactureDeviceName | ModelName.
	Type types.Code
}

func (s DeviceDeviceName) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.
//
// Allows institutions to track their devices.
type Device struct {
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The model number for the device.
	ModelNumber *types.String
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The distinct identification string as required by regulation for a human cell, tissue, or cellular and tissue-based product.
	DistinctIdentifier *types.String
	// Patient information, If the device is affixed to a person.
	Patient *types.Reference
	// Provides additional safety characteristics about a medical device.  For example devices containing latex.
	Safety []types.CodeableConcept
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Unique instance identifiers assigned to a device by manufacturers other organizations or owners.
	Identifier []types.Identifier
	// Status of the Device availability.
	Status *types.Code
	// A name of the manufacturer.
	Manufacturer *types.String
	// The serial number assigned by the organization when the device was manufactured.
	SerialNumber *types.String
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The date and time beyond which this device is no longer valid or should not be used (if applicable).
	ExpirationDate *types.DateTime
	// The place where the device can be found.
	Location *types.Reference
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
	UdiCarrier []DeviceUdiCarrier
	// The date and time when the device was manufactured.
	ManufactureDate *types.DateTime
	// An organization that is responsible for the provision and ongoing maintenance of the device.
	Owner *types.Reference
	// The reference to the definition for the device.
	Definition *types.Reference
	// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
	Property []DeviceProperty
	// The parent device.
	Parent *types.Reference
	// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
	Specialization []DeviceSpecialization
	// Reason for the dtatus of the Device availability.
	StatusReason []types.CodeableConcept
	// The part number of the device.
	PartNumber *types.String
	// The actual design of the device or software version running on the device.
	Version []DeviceVersion
	// A network address on which the device may be contacted directly.
	Url *types.Uri
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// This represents the manufacturer's name of the device as provided by the device, from a UDI label, or by a person describing the Device.  This typically would be used when a person provides the name(s) or when the device represents one of the names available from DeviceDefinition.
	DeviceName []DeviceDeviceName
	// The kind or type of device.
	Type *types.CodeableConcept
	// Contact details for an organization or a particular human that is responsible for the device.
	Contact []types.ContactPoint
	// Descriptive information, usage information or implantation information that is not captured in an existing element.
	Note []types.Annotation
	// Lot number assigned by the manufacturer.
	LotNumber *types.String
}

func (s Device) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
