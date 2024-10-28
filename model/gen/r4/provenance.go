package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance struct {
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
	// The Reference(s) that were generated or updated by  the activity described in this resource. A provenance can point to more than one target if multiple resources were created/updated by the same activity.
	Target []Reference
	// The period during which the activity occurred.
	Occurred isProvenanceOccurred
	// The instant of time at which the activity was recorded.
	Recorded Instant
	// Policy or plan the activity was defined by. Typically, a single activity may have multiple applicable policy documents, such as patient consent, guarantor funding, etc.
	Policy []Uri
	// Where the activity occurred, if relevant.
	Location *Reference
	// The reason that the activity was taking place.
	Reason []CodeableConcept
	// An activity is something that occurs over a period of time and acts upon or with entities; it may include consuming, processing, transforming, modifying, relocating, using, or generating entities.
	Activity *CodeableConcept
	// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
	Agent []ProvenanceAgent
	// An entity used in this activity.
	Entity []ProvenanceEntity
	// A digital signature on the target Reference(s). The signer should match a Provenance.agent. The purpose of the signature is indicated.
	Signature []Signature
}

func (r Provenance) ResourceType() string {
	return "Provenance"
}
func (r Provenance) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isProvenanceOccurred interface {
	isProvenanceOccurred()
}

func (r Period) isProvenanceOccurred()   {}
func (r DateTime) isProvenanceOccurred() {}

type jsonProvenance struct {
	ResourceType                     string              `json:"resourceType"`
	Id                               *Id                 `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement   `json:"_id,omitempty"`
	Meta                             *Meta               `json:"meta,omitempty"`
	ImplicitRules                    *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                         *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement   `json:"_language,omitempty"`
	Text                             *Narrative          `json:"text,omitempty"`
	Contained                        []ContainedResource `json:"contained,omitempty"`
	Extension                        []Extension         `json:"extension,omitempty"`
	ModifierExtension                []Extension         `json:"modifierExtension,omitempty"`
	Target                           []Reference         `json:"target,omitempty"`
	OccurredPeriod                   *Period             `json:"occurredPeriod,omitempty"`
	OccurredDateTime                 *DateTime           `json:"occurredDateTime,omitempty"`
	OccurredDateTimePrimitiveElement *primitiveElement   `json:"_occurredDateTime,omitempty"`
	Recorded                         Instant             `json:"recorded,omitempty"`
	RecordedPrimitiveElement         *primitiveElement   `json:"_recorded,omitempty"`
	Policy                           []Uri               `json:"policy,omitempty"`
	PolicyPrimitiveElement           []*primitiveElement `json:"_policy,omitempty"`
	Location                         *Reference          `json:"location,omitempty"`
	Reason                           []CodeableConcept   `json:"reason,omitempty"`
	Activity                         *CodeableConcept    `json:"activity,omitempty"`
	Agent                            []ProvenanceAgent   `json:"agent,omitempty"`
	Entity                           []ProvenanceEntity  `json:"entity,omitempty"`
	Signature                        []Signature         `json:"signature,omitempty"`
}

func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Provenance) marshalJSON() jsonProvenance {
	m := jsonProvenance{}
	m.ResourceType = "Provenance"
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
	m.Target = r.Target
	switch v := r.Occurred.(type) {
	case Period:
		m.OccurredPeriod = &v
	case *Period:
		m.OccurredPeriod = v
	case DateTime:
		if v.Value != nil {
			m.OccurredDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurredDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.OccurredDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurredDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.Recorded.Value != nil {
		m.Recorded = r.Recorded
	}
	if r.Recorded.Id != nil || r.Recorded.Extension != nil {
		m.RecordedPrimitiveElement = &primitiveElement{Id: r.Recorded.Id, Extension: r.Recorded.Extension}
	}
	anyPolicyValue := false
	for _, e := range r.Policy {
		if e.Value != nil {
			anyPolicyValue = true
			break
		}
	}
	if anyPolicyValue {
		m.Policy = r.Policy
	}
	anyPolicyIdOrExtension := false
	for _, e := range r.Policy {
		if e.Id != nil || e.Extension != nil {
			anyPolicyIdOrExtension = true
			break
		}
	}
	if anyPolicyIdOrExtension {
		m.PolicyPrimitiveElement = make([]*primitiveElement, 0, len(r.Policy))
		for _, e := range r.Policy {
			if e.Id != nil || e.Extension != nil {
				m.PolicyPrimitiveElement = append(m.PolicyPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PolicyPrimitiveElement = append(m.PolicyPrimitiveElement, nil)
			}
		}
	}
	m.Location = r.Location
	m.Reason = r.Reason
	m.Activity = r.Activity
	m.Agent = r.Agent
	m.Entity = r.Entity
	m.Signature = r.Signature
	return m
}
func (r *Provenance) UnmarshalJSON(b []byte) error {
	var m jsonProvenance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Provenance) unmarshalJSON(m jsonProvenance) error {
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
	r.Target = m.Target
	if m.OccurredPeriod != nil {
		if r.Occurred != nil {
			return fmt.Errorf("multiple values for field \"Occurred\"")
		}
		v := m.OccurredPeriod
		r.Occurred = v
	}
	if m.OccurredDateTime != nil || m.OccurredDateTimePrimitiveElement != nil {
		if r.Occurred != nil {
			return fmt.Errorf("multiple values for field \"Occurred\"")
		}
		v := m.OccurredDateTime
		if m.OccurredDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurredDateTimePrimitiveElement.Id
			v.Extension = m.OccurredDateTimePrimitiveElement.Extension
		}
		r.Occurred = v
	}
	r.Recorded = m.Recorded
	if m.RecordedPrimitiveElement != nil {
		r.Recorded.Id = m.RecordedPrimitiveElement.Id
		r.Recorded.Extension = m.RecordedPrimitiveElement.Extension
	}
	r.Policy = m.Policy
	for i, e := range m.PolicyPrimitiveElement {
		if len(r.Policy) <= i {
			r.Policy = append(r.Policy, Uri{})
		}
		if e != nil {
			r.Policy[i].Id = e.Id
			r.Policy[i].Extension = e.Extension
		}
	}
	r.Location = m.Location
	r.Reason = m.Reason
	r.Activity = m.Activity
	r.Agent = m.Agent
	r.Entity = m.Entity
	r.Signature = m.Signature
	return nil
}
func (r Provenance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Provenance"
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
	err = e.EncodeElement(r.Target, xml.StartElement{Name: xml.Name{Local: "target"}})
	if err != nil {
		return err
	}
	switch v := r.Occurred.(type) {
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurredPeriod"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurredDateTime"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Recorded, xml.StartElement{Name: xml.Name{Local: "recorded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Policy, xml.StartElement{Name: xml.Name{Local: "policy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Activity, xml.StartElement{Name: xml.Name{Local: "activity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Agent, xml.StartElement{Name: xml.Name{Local: "agent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Entity, xml.StartElement{Name: xml.Name{Local: "entity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Signature, xml.StartElement{Name: xml.Name{Local: "signature"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Provenance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "target":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = append(r.Target, v)
			case "occurredPeriod":
				if r.Occurred != nil {
					return fmt.Errorf("multiple values for field \"Occurred\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurred = v
			case "occurredDateTime":
				if r.Occurred != nil {
					return fmt.Errorf("multiple values for field \"Occurred\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurred = v
			case "recorded":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recorded = v
			case "policy":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Policy = append(r.Policy, v)
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "reason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			case "activity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Activity = &v
			case "agent":
				var v ProvenanceAgent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Agent = append(r.Agent, v)
			case "entity":
				var v ProvenanceEntity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entity = append(r.Entity, v)
			case "signature":
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Signature = append(r.Signature, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Provenance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
type ProvenanceAgent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The participation the agent had with respect to the activity.
	Type *CodeableConcept
	// The function of the agent with respect to the activity. The security role enabling the agent with respect to the activity.
	Role []CodeableConcept
	// The individual, device or organization that participated in the event.
	Who Reference
	// The individual, device, or organization for whom the change was made.
	OnBehalfOf *Reference
}
type jsonProvenanceAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who,omitempty"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

func (r ProvenanceAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ProvenanceAgent) marshalJSON() jsonProvenanceAgent {
	m := jsonProvenanceAgent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Role = r.Role
	m.Who = r.Who
	m.OnBehalfOf = r.OnBehalfOf
	return m
}
func (r *ProvenanceAgent) UnmarshalJSON(b []byte) error {
	var m jsonProvenanceAgent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ProvenanceAgent) unmarshalJSON(m jsonProvenanceAgent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Role = m.Role
	r.Who = m.Who
	r.OnBehalfOf = m.OnBehalfOf
	return nil
}
func (r ProvenanceAgent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Who, xml.StartElement{Name: xml.Name{Local: "who"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OnBehalfOf, xml.StartElement{Name: xml.Name{Local: "onBehalfOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ProvenanceAgent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = append(r.Role, v)
			case "who":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Who = v
			case "onBehalfOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OnBehalfOf = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ProvenanceAgent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An entity used in this activity.
type ProvenanceEntity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// How the entity was used during the activity.
	Role Code
	// Identity of the  Entity used. May be a logical or physical uri and maybe absolute or relative.
	What Reference
	// The entity is attributed to an agent to express the agent's responsibility for that entity, possibly along with other agents. This description can be understood as shorthand for saying that the agent was responsible for the activity which generated the entity.
	Agent []ProvenanceAgent
}
type jsonProvenanceEntity struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Role                 Code              `json:"role,omitempty"`
	RolePrimitiveElement *primitiveElement `json:"_role,omitempty"`
	What                 Reference         `json:"what,omitempty"`
	Agent                []ProvenanceAgent `json:"agent,omitempty"`
}

func (r ProvenanceEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ProvenanceEntity) marshalJSON() jsonProvenanceEntity {
	m := jsonProvenanceEntity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Role.Value != nil {
		m.Role = r.Role
	}
	if r.Role.Id != nil || r.Role.Extension != nil {
		m.RolePrimitiveElement = &primitiveElement{Id: r.Role.Id, Extension: r.Role.Extension}
	}
	m.What = r.What
	m.Agent = r.Agent
	return m
}
func (r *ProvenanceEntity) UnmarshalJSON(b []byte) error {
	var m jsonProvenanceEntity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ProvenanceEntity) unmarshalJSON(m jsonProvenanceEntity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Role = m.Role
	if m.RolePrimitiveElement != nil {
		r.Role.Id = m.RolePrimitiveElement.Id
		r.Role.Extension = m.RolePrimitiveElement.Extension
	}
	r.What = m.What
	r.Agent = m.Agent
	return nil
}
func (r ProvenanceEntity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.What, xml.StartElement{Name: xml.Name{Local: "what"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Agent, xml.StartElement{Name: xml.Name{Local: "agent"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ProvenanceEntity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "role":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = v
			case "what":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.What = v
			case "agent":
				var v ProvenanceAgent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Agent = append(r.Agent, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ProvenanceEntity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
