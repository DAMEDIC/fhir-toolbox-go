package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceDefinition struct {
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
	// Unique instance identifiers assigned to a device by the software, manufacturers, other organizations or owners. For example: handle ID.
	Identifier []Identifier
	// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier
	// A name of the manufacturer.
	Manufacturer isDeviceDefinitionManufacturer
	// A name given to the device to identify it.
	DeviceName []DeviceDefinitionDeviceName
	// The model number for the device.
	ModelNumber *String
	// What kind of device or device system this is.
	Type *CodeableConcept
	// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
	Specialization []DeviceDefinitionSpecialization
	// The available versions of the device, e.g., software versions.
	Version []String
	// Safety characteristics of the device.
	Safety []CodeableConcept
	// Shelf Life and storage information.
	ShelfLifeStorage []ProductShelfLife
	// Dimensions, color etc.
	PhysicalCharacteristics *ProdCharacteristic
	// Language code for the human-readable text strings produced by the device (all supported).
	LanguageCode []CodeableConcept
	// Device capabilities.
	Capability []DeviceDefinitionCapability
	// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
	Property []DeviceDefinitionProperty
	// An organization that is responsible for the provision and ongoing maintenance of the device.
	Owner *Reference
	// Contact details for an organization or a particular human that is responsible for the device.
	Contact []ContactPoint
	// A network address on which the device may be contacted directly.
	Url *Uri
	// Access to on-line information about the device.
	OnlineInformation *Uri
	// Descriptive information, usage information or implantation information that is not captured in an existing element.
	Note []Annotation
	// The quantity of the device present in the packaging (e.g. the number of devices present in a pack, or the number of devices in the same package of the medicinal product).
	Quantity *Quantity
	// The parent device it can be part of.
	ParentDevice *Reference
	// A substance used to create the material(s) of which the device is made.
	Material []DeviceDefinitionMaterial
}

func (r DeviceDefinition) ResourceType() string {
	return "DeviceDefinition"
}
func (r DeviceDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isDeviceDefinitionManufacturer interface {
	isDeviceDefinitionManufacturer()
}

func (r String) isDeviceDefinitionManufacturer()    {}
func (r Reference) isDeviceDefinitionManufacturer() {}

type jsonDeviceDefinition struct {
	ResourceType                       string                                `json:"resourceType"`
	Id                                 *Id                                   `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement                     `json:"_id,omitempty"`
	Meta                               *Meta                                 `json:"meta,omitempty"`
	ImplicitRules                      *Uri                                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement                     `json:"_implicitRules,omitempty"`
	Language                           *Code                                 `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement                     `json:"_language,omitempty"`
	Text                               *Narrative                            `json:"text,omitempty"`
	Contained                          []ContainedResource                   `json:"contained,omitempty"`
	Extension                          []Extension                           `json:"extension,omitempty"`
	ModifierExtension                  []Extension                           `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier                          `json:"identifier,omitempty"`
	UdiDeviceIdentifier                []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`
	ManufacturerString                 *String                               `json:"manufacturerString,omitempty"`
	ManufacturerStringPrimitiveElement *primitiveElement                     `json:"_manufacturerString,omitempty"`
	ManufacturerReference              *Reference                            `json:"manufacturerReference,omitempty"`
	DeviceName                         []DeviceDefinitionDeviceName          `json:"deviceName,omitempty"`
	ModelNumber                        *String                               `json:"modelNumber,omitempty"`
	ModelNumberPrimitiveElement        *primitiveElement                     `json:"_modelNumber,omitempty"`
	Type                               *CodeableConcept                      `json:"type,omitempty"`
	Specialization                     []DeviceDefinitionSpecialization      `json:"specialization,omitempty"`
	Version                            []String                              `json:"version,omitempty"`
	VersionPrimitiveElement            []*primitiveElement                   `json:"_version,omitempty"`
	Safety                             []CodeableConcept                     `json:"safety,omitempty"`
	ShelfLifeStorage                   []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	PhysicalCharacteristics            *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	LanguageCode                       []CodeableConcept                     `json:"languageCode,omitempty"`
	Capability                         []DeviceDefinitionCapability          `json:"capability,omitempty"`
	Property                           []DeviceDefinitionProperty            `json:"property,omitempty"`
	Owner                              *Reference                            `json:"owner,omitempty"`
	Contact                            []ContactPoint                        `json:"contact,omitempty"`
	Url                                *Uri                                  `json:"url,omitempty"`
	UrlPrimitiveElement                *primitiveElement                     `json:"_url,omitempty"`
	OnlineInformation                  *Uri                                  `json:"onlineInformation,omitempty"`
	OnlineInformationPrimitiveElement  *primitiveElement                     `json:"_onlineInformation,omitempty"`
	Note                               []Annotation                          `json:"note,omitempty"`
	Quantity                           *Quantity                             `json:"quantity,omitempty"`
	ParentDevice                       *Reference                            `json:"parentDevice,omitempty"`
	Material                           []DeviceDefinitionMaterial            `json:"material,omitempty"`
}

func (r DeviceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinition) marshalJSON() jsonDeviceDefinition {
	m := jsonDeviceDefinition{}
	m.ResourceType = "DeviceDefinition"
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
	m.UdiDeviceIdentifier = r.UdiDeviceIdentifier
	switch v := r.Manufacturer.(type) {
	case String:
		if v.Value != nil {
			m.ManufacturerString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ManufacturerStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.ManufacturerString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ManufacturerStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Reference:
		m.ManufacturerReference = &v
	case *Reference:
		m.ManufacturerReference = v
	}
	m.DeviceName = r.DeviceName
	if r.ModelNumber != nil && r.ModelNumber.Value != nil {
		m.ModelNumber = r.ModelNumber
	}
	if r.ModelNumber != nil && (r.ModelNumber.Id != nil || r.ModelNumber.Extension != nil) {
		m.ModelNumberPrimitiveElement = &primitiveElement{Id: r.ModelNumber.Id, Extension: r.ModelNumber.Extension}
	}
	m.Type = r.Type
	m.Specialization = r.Specialization
	anyVersionValue := false
	for _, e := range r.Version {
		if e.Value != nil {
			anyVersionValue = true
			break
		}
	}
	if anyVersionValue {
		m.Version = r.Version
	}
	anyVersionIdOrExtension := false
	for _, e := range r.Version {
		if e.Id != nil || e.Extension != nil {
			anyVersionIdOrExtension = true
			break
		}
	}
	if anyVersionIdOrExtension {
		m.VersionPrimitiveElement = make([]*primitiveElement, 0, len(r.Version))
		for _, e := range r.Version {
			if e.Id != nil || e.Extension != nil {
				m.VersionPrimitiveElement = append(m.VersionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.VersionPrimitiveElement = append(m.VersionPrimitiveElement, nil)
			}
		}
	}
	m.Safety = r.Safety
	m.ShelfLifeStorage = r.ShelfLifeStorage
	m.PhysicalCharacteristics = r.PhysicalCharacteristics
	m.LanguageCode = r.LanguageCode
	m.Capability = r.Capability
	m.Property = r.Property
	m.Owner = r.Owner
	m.Contact = r.Contact
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.OnlineInformation != nil && r.OnlineInformation.Value != nil {
		m.OnlineInformation = r.OnlineInformation
	}
	if r.OnlineInformation != nil && (r.OnlineInformation.Id != nil || r.OnlineInformation.Extension != nil) {
		m.OnlineInformationPrimitiveElement = &primitiveElement{Id: r.OnlineInformation.Id, Extension: r.OnlineInformation.Extension}
	}
	m.Note = r.Note
	m.Quantity = r.Quantity
	m.ParentDevice = r.ParentDevice
	m.Material = r.Material
	return m
}
func (r *DeviceDefinition) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinition) unmarshalJSON(m jsonDeviceDefinition) error {
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
	r.UdiDeviceIdentifier = m.UdiDeviceIdentifier
	if m.ManufacturerString != nil || m.ManufacturerStringPrimitiveElement != nil {
		if r.Manufacturer != nil {
			return fmt.Errorf("multiple values for field \"Manufacturer\"")
		}
		v := m.ManufacturerString
		if m.ManufacturerStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ManufacturerStringPrimitiveElement.Id
			v.Extension = m.ManufacturerStringPrimitiveElement.Extension
		}
		r.Manufacturer = v
	}
	if m.ManufacturerReference != nil {
		if r.Manufacturer != nil {
			return fmt.Errorf("multiple values for field \"Manufacturer\"")
		}
		v := m.ManufacturerReference
		r.Manufacturer = v
	}
	r.DeviceName = m.DeviceName
	r.ModelNumber = m.ModelNumber
	if m.ModelNumberPrimitiveElement != nil {
		if r.ModelNumber == nil {
			r.ModelNumber = &String{}
		}
		r.ModelNumber.Id = m.ModelNumberPrimitiveElement.Id
		r.ModelNumber.Extension = m.ModelNumberPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Specialization = m.Specialization
	r.Version = m.Version
	for i, e := range m.VersionPrimitiveElement {
		if len(r.Version) <= i {
			r.Version = append(r.Version, String{})
		}
		if e != nil {
			r.Version[i].Id = e.Id
			r.Version[i].Extension = e.Extension
		}
	}
	r.Safety = m.Safety
	r.ShelfLifeStorage = m.ShelfLifeStorage
	r.PhysicalCharacteristics = m.PhysicalCharacteristics
	r.LanguageCode = m.LanguageCode
	r.Capability = m.Capability
	r.Property = m.Property
	r.Owner = m.Owner
	r.Contact = m.Contact
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.OnlineInformation = m.OnlineInformation
	if m.OnlineInformationPrimitiveElement != nil {
		if r.OnlineInformation == nil {
			r.OnlineInformation = &Uri{}
		}
		r.OnlineInformation.Id = m.OnlineInformationPrimitiveElement.Id
		r.OnlineInformation.Extension = m.OnlineInformationPrimitiveElement.Extension
	}
	r.Note = m.Note
	r.Quantity = m.Quantity
	r.ParentDevice = m.ParentDevice
	r.Material = m.Material
	return nil
}
func (r DeviceDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Unique device identifier (UDI) assigned to device label or package.  Note that the Device may include multiple udiCarriers as it either may include just the udiCarrier for the jurisdiction it is sold, or for multiple jurisdictions it could have been sold.
type DeviceDefinitionUdiDeviceIdentifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The identifier that is to be associated with every Device that references this DeviceDefintiion for the issuer and jurisdication porvided in the DeviceDefinition.udiDeviceIdentifier.
	DeviceIdentifier String
	// The organization that assigns the identifier algorithm.
	Issuer Uri
	// The jurisdiction to which the deviceIdentifier applies.
	Jurisdiction Uri
}
type jsonDeviceDefinitionUdiDeviceIdentifier struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	DeviceIdentifier                 String            `json:"deviceIdentifier,omitempty"`
	DeviceIdentifierPrimitiveElement *primitiveElement `json:"_deviceIdentifier,omitempty"`
	Issuer                           Uri               `json:"issuer,omitempty"`
	IssuerPrimitiveElement           *primitiveElement `json:"_issuer,omitempty"`
	Jurisdiction                     Uri               `json:"jurisdiction,omitempty"`
	JurisdictionPrimitiveElement     *primitiveElement `json:"_jurisdiction,omitempty"`
}

func (r DeviceDefinitionUdiDeviceIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionUdiDeviceIdentifier) marshalJSON() jsonDeviceDefinitionUdiDeviceIdentifier {
	m := jsonDeviceDefinitionUdiDeviceIdentifier{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.DeviceIdentifier.Value != nil {
		m.DeviceIdentifier = r.DeviceIdentifier
	}
	if r.DeviceIdentifier.Id != nil || r.DeviceIdentifier.Extension != nil {
		m.DeviceIdentifierPrimitiveElement = &primitiveElement{Id: r.DeviceIdentifier.Id, Extension: r.DeviceIdentifier.Extension}
	}
	if r.Issuer.Value != nil {
		m.Issuer = r.Issuer
	}
	if r.Issuer.Id != nil || r.Issuer.Extension != nil {
		m.IssuerPrimitiveElement = &primitiveElement{Id: r.Issuer.Id, Extension: r.Issuer.Extension}
	}
	if r.Jurisdiction.Value != nil {
		m.Jurisdiction = r.Jurisdiction
	}
	if r.Jurisdiction.Id != nil || r.Jurisdiction.Extension != nil {
		m.JurisdictionPrimitiveElement = &primitiveElement{Id: r.Jurisdiction.Id, Extension: r.Jurisdiction.Extension}
	}
	return m
}
func (r *DeviceDefinitionUdiDeviceIdentifier) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionUdiDeviceIdentifier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionUdiDeviceIdentifier) unmarshalJSON(m jsonDeviceDefinitionUdiDeviceIdentifier) error {
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
	return nil
}
func (r DeviceDefinitionUdiDeviceIdentifier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A name given to the device to identify it.
type DeviceDefinitionDeviceName struct {
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
type jsonDeviceDefinitionDeviceName struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Name                 String            `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
}

func (r DeviceDefinitionDeviceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionDeviceName) marshalJSON() jsonDeviceDefinitionDeviceName {
	m := jsonDeviceDefinitionDeviceName{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	return m
}
func (r *DeviceDefinitionDeviceName) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionDeviceName
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionDeviceName) unmarshalJSON(m jsonDeviceDefinitionDeviceName) error {
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
func (r DeviceDefinitionDeviceName) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The capabilities supported on a  device, the standards to which the device conforms for a particular purpose, and used for the communication.
type DeviceDefinitionSpecialization struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The standard that is used to operate and communicate.
	SystemType String
	// The version of the standard that is used to operate and communicate.
	Version *String
}
type jsonDeviceDefinitionSpecialization struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	SystemType                 String            `json:"systemType,omitempty"`
	SystemTypePrimitiveElement *primitiveElement `json:"_systemType,omitempty"`
	Version                    *String           `json:"version,omitempty"`
	VersionPrimitiveElement    *primitiveElement `json:"_version,omitempty"`
}

func (r DeviceDefinitionSpecialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionSpecialization) marshalJSON() jsonDeviceDefinitionSpecialization {
	m := jsonDeviceDefinitionSpecialization{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.SystemType.Value != nil {
		m.SystemType = r.SystemType
	}
	if r.SystemType.Id != nil || r.SystemType.Extension != nil {
		m.SystemTypePrimitiveElement = &primitiveElement{Id: r.SystemType.Id, Extension: r.SystemType.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	return m
}
func (r *DeviceDefinitionSpecialization) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionSpecialization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionSpecialization) unmarshalJSON(m jsonDeviceDefinitionSpecialization) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.SystemType = m.SystemType
	if m.SystemTypePrimitiveElement != nil {
		r.SystemType.Id = m.SystemTypePrimitiveElement.Id
		r.SystemType.Extension = m.SystemTypePrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	return nil
}
func (r DeviceDefinitionSpecialization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Device capabilities.
type DeviceDefinitionCapability struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of capability.
	Type CodeableConcept
	// Description of capability.
	Description []CodeableConcept
}
type jsonDeviceDefinitionCapability struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type,omitempty"`
	Description       []CodeableConcept `json:"description,omitempty"`
}

func (r DeviceDefinitionCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionCapability) marshalJSON() jsonDeviceDefinitionCapability {
	m := jsonDeviceDefinitionCapability{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Description = r.Description
	return m
}
func (r *DeviceDefinitionCapability) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionCapability
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionCapability) unmarshalJSON(m jsonDeviceDefinitionCapability) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Description = m.Description
	return nil
}
func (r DeviceDefinitionCapability) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties.
type DeviceDefinitionProperty struct {
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
type jsonDeviceDefinitionProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type,omitempty"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

func (r DeviceDefinitionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionProperty) marshalJSON() jsonDeviceDefinitionProperty {
	m := jsonDeviceDefinitionProperty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.ValueQuantity = r.ValueQuantity
	m.ValueCode = r.ValueCode
	return m
}
func (r *DeviceDefinitionProperty) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionProperty) unmarshalJSON(m jsonDeviceDefinitionProperty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.ValueQuantity = m.ValueQuantity
	r.ValueCode = m.ValueCode
	return nil
}
func (r DeviceDefinitionProperty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A substance used to create the material(s) of which the device is made.
type DeviceDefinitionMaterial struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The substance.
	Substance CodeableConcept
	// Indicates an alternative material of the device.
	Alternate *Boolean
	// Whether the substance is a known or suspected allergen.
	AllergenicIndicator *Boolean
}
type jsonDeviceDefinitionMaterial struct {
	Id                                  *string           `json:"id,omitempty"`
	Extension                           []Extension       `json:"extension,omitempty"`
	ModifierExtension                   []Extension       `json:"modifierExtension,omitempty"`
	Substance                           CodeableConcept   `json:"substance,omitempty"`
	Alternate                           *Boolean          `json:"alternate,omitempty"`
	AlternatePrimitiveElement           *primitiveElement `json:"_alternate,omitempty"`
	AllergenicIndicator                 *Boolean          `json:"allergenicIndicator,omitempty"`
	AllergenicIndicatorPrimitiveElement *primitiveElement `json:"_allergenicIndicator,omitempty"`
}

func (r DeviceDefinitionMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceDefinitionMaterial) marshalJSON() jsonDeviceDefinitionMaterial {
	m := jsonDeviceDefinitionMaterial{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Substance = r.Substance
	if r.Alternate != nil && r.Alternate.Value != nil {
		m.Alternate = r.Alternate
	}
	if r.Alternate != nil && (r.Alternate.Id != nil || r.Alternate.Extension != nil) {
		m.AlternatePrimitiveElement = &primitiveElement{Id: r.Alternate.Id, Extension: r.Alternate.Extension}
	}
	if r.AllergenicIndicator != nil && r.AllergenicIndicator.Value != nil {
		m.AllergenicIndicator = r.AllergenicIndicator
	}
	if r.AllergenicIndicator != nil && (r.AllergenicIndicator.Id != nil || r.AllergenicIndicator.Extension != nil) {
		m.AllergenicIndicatorPrimitiveElement = &primitiveElement{Id: r.AllergenicIndicator.Id, Extension: r.AllergenicIndicator.Extension}
	}
	return m
}
func (r *DeviceDefinitionMaterial) UnmarshalJSON(b []byte) error {
	var m jsonDeviceDefinitionMaterial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceDefinitionMaterial) unmarshalJSON(m jsonDeviceDefinitionMaterial) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Substance = m.Substance
	r.Alternate = m.Alternate
	if m.AlternatePrimitiveElement != nil {
		if r.Alternate == nil {
			r.Alternate = &Boolean{}
		}
		r.Alternate.Id = m.AlternatePrimitiveElement.Id
		r.Alternate.Extension = m.AlternatePrimitiveElement.Extension
	}
	r.AllergenicIndicator = m.AllergenicIndicator
	if m.AllergenicIndicatorPrimitiveElement != nil {
		if r.AllergenicIndicator == nil {
			r.AllergenicIndicator = &Boolean{}
		}
		r.AllergenicIndicator.Id = m.AllergenicIndicatorPrimitiveElement.Id
		r.AllergenicIndicator.Extension = m.AllergenicIndicatorPrimitiveElement.Extension
	}
	return nil
}
func (r DeviceDefinitionMaterial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
