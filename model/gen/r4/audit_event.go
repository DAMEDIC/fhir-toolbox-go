package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent struct {
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
	// Identifier for a family of the event.  For example, a menu item, program, rule, policy, function code, application name or URL. It identifies the performed function.
	Type Coding
	// Identifier for the category of event.
	Subtype []Coding
	// Indicator for type of action performed during the event that generated the audit.
	Action *Code
	// The period during which the activity occurred.
	Period *Period
	// The time when the event was recorded.
	Recorded Instant
	// Indicates whether the event succeeded or failed.
	Outcome *Code
	// A free text description of the outcome of the event.
	OutcomeDesc *String
	// The purposeOfUse (reason) that was used during the event being recorded.
	PurposeOfEvent []CodeableConcept
	// An actor taking an active role in the event or activity that is logged.
	Agent []AuditEventAgent
	// The system that is reporting the event.
	Source AuditEventSource
	// Specific instances of data or objects that have been accessed.
	Entity []AuditEventEntity
}

func (r AuditEvent) ResourceType() string {
	return "AuditEvent"
}
func (r AuditEvent) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonAuditEvent struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []ContainedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Type                          Coding              `json:"type,omitempty"`
	Subtype                       []Coding            `json:"subtype,omitempty"`
	Action                        *Code               `json:"action,omitempty"`
	ActionPrimitiveElement        *primitiveElement   `json:"_action,omitempty"`
	Period                        *Period             `json:"period,omitempty"`
	Recorded                      Instant             `json:"recorded,omitempty"`
	RecordedPrimitiveElement      *primitiveElement   `json:"_recorded,omitempty"`
	Outcome                       *Code               `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement   `json:"_outcome,omitempty"`
	OutcomeDesc                   *String             `json:"outcomeDesc,omitempty"`
	OutcomeDescPrimitiveElement   *primitiveElement   `json:"_outcomeDesc,omitempty"`
	PurposeOfEvent                []CodeableConcept   `json:"purposeOfEvent,omitempty"`
	Agent                         []AuditEventAgent   `json:"agent,omitempty"`
	Source                        AuditEventSource    `json:"source,omitempty"`
	Entity                        []AuditEventEntity  `json:"entity,omitempty"`
}

func (r AuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEvent) marshalJSON() jsonAuditEvent {
	m := jsonAuditEvent{}
	m.ResourceType = "AuditEvent"
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
	m.Type = r.Type
	m.Subtype = r.Subtype
	if r.Action != nil && r.Action.Value != nil {
		m.Action = r.Action
	}
	if r.Action != nil && (r.Action.Id != nil || r.Action.Extension != nil) {
		m.ActionPrimitiveElement = &primitiveElement{Id: r.Action.Id, Extension: r.Action.Extension}
	}
	m.Period = r.Period
	if r.Recorded.Value != nil {
		m.Recorded = r.Recorded
	}
	if r.Recorded.Id != nil || r.Recorded.Extension != nil {
		m.RecordedPrimitiveElement = &primitiveElement{Id: r.Recorded.Id, Extension: r.Recorded.Extension}
	}
	if r.Outcome != nil && r.Outcome.Value != nil {
		m.Outcome = r.Outcome
	}
	if r.Outcome != nil && (r.Outcome.Id != nil || r.Outcome.Extension != nil) {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	if r.OutcomeDesc != nil && r.OutcomeDesc.Value != nil {
		m.OutcomeDesc = r.OutcomeDesc
	}
	if r.OutcomeDesc != nil && (r.OutcomeDesc.Id != nil || r.OutcomeDesc.Extension != nil) {
		m.OutcomeDescPrimitiveElement = &primitiveElement{Id: r.OutcomeDesc.Id, Extension: r.OutcomeDesc.Extension}
	}
	m.PurposeOfEvent = r.PurposeOfEvent
	m.Agent = r.Agent
	m.Source = r.Source
	m.Entity = r.Entity
	return m
}
func (r *AuditEvent) UnmarshalJSON(b []byte) error {
	var m jsonAuditEvent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEvent) unmarshalJSON(m jsonAuditEvent) error {
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
	r.Type = m.Type
	r.Subtype = m.Subtype
	r.Action = m.Action
	if m.ActionPrimitiveElement != nil {
		if r.Action == nil {
			r.Action = &Code{}
		}
		r.Action.Id = m.ActionPrimitiveElement.Id
		r.Action.Extension = m.ActionPrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Recorded = m.Recorded
	if m.RecordedPrimitiveElement != nil {
		r.Recorded.Id = m.RecordedPrimitiveElement.Id
		r.Recorded.Extension = m.RecordedPrimitiveElement.Extension
	}
	r.Outcome = m.Outcome
	if m.OutcomePrimitiveElement != nil {
		if r.Outcome == nil {
			r.Outcome = &Code{}
		}
		r.Outcome.Id = m.OutcomePrimitiveElement.Id
		r.Outcome.Extension = m.OutcomePrimitiveElement.Extension
	}
	r.OutcomeDesc = m.OutcomeDesc
	if m.OutcomeDescPrimitiveElement != nil {
		if r.OutcomeDesc == nil {
			r.OutcomeDesc = &String{}
		}
		r.OutcomeDesc.Id = m.OutcomeDescPrimitiveElement.Id
		r.OutcomeDesc.Extension = m.OutcomeDescPrimitiveElement.Extension
	}
	r.PurposeOfEvent = m.PurposeOfEvent
	r.Agent = m.Agent
	r.Source = m.Source
	r.Entity = m.Entity
	return nil
}
func (r AuditEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "AuditEvent"
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subtype, xml.StartElement{Name: xml.Name{Local: "subtype"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recorded, xml.StartElement{Name: xml.Name{Local: "recorded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "outcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OutcomeDesc, xml.StartElement{Name: xml.Name{Local: "outcomeDesc"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PurposeOfEvent, xml.StartElement{Name: xml.Name{Local: "purposeOfEvent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Agent, xml.StartElement{Name: xml.Name{Local: "agent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Entity, xml.StartElement{Name: xml.Name{Local: "entity"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEvent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "subtype":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subtype = append(r.Subtype, v)
			case "action":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			case "recorded":
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recorded = v
			case "outcome":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = &v
			case "outcomeDesc":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OutcomeDesc = &v
			case "purposeOfEvent":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PurposeOfEvent = append(r.PurposeOfEvent, v)
			case "agent":
				var v AuditEventAgent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Agent = append(r.Agent, v)
			case "source":
				var v AuditEventSource
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = v
			case "entity":
				var v AuditEventEntity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entity = append(r.Entity, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEvent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An actor taking an active role in the event or activity that is logged.
type AuditEventAgent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specification of the participation type the user plays when performing the event.
	Type *CodeableConcept
	// The security role that the user was acting under, that come from local codes defined by the access control security system (e.g. RBAC, ABAC) used in the local context.
	Role []CodeableConcept
	// Reference to who this agent is that was involved in the event.
	Who *Reference
	// Alternative agent Identifier. For a human, this should be a user identifier text string from authentication system. This identifier would be one known to a common authentication system (e.g. single sign-on), if available.
	AltId *String
	// Human-meaningful name for the agent.
	Name *String
	// Indicator that the user is or is not the requestor, or initiator, for the event being audited.
	Requestor Boolean
	// Where the event occurred.
	Location *Reference
	// The policy or plan that authorized the activity being recorded. Typically, a single activity may have multiple applicable policies, such as patient consent, guarantor funding, etc. The policy would also indicate the security token used.
	Policy []Uri
	// Type of media involved. Used when the event is about exporting/importing onto media.
	Media *Coding
	// Logical network location for application activity, if the activity has a network location.
	Network *AuditEventAgentNetwork
	// The reason (purpose of use), specific to this agent, that was used during the event being recorded.
	PurposeOfUse []CodeableConcept
}
type jsonAuditEventAgent struct {
	Id                        *string                 `json:"id,omitempty"`
	Extension                 []Extension             `json:"extension,omitempty"`
	ModifierExtension         []Extension             `json:"modifierExtension,omitempty"`
	Type                      *CodeableConcept        `json:"type,omitempty"`
	Role                      []CodeableConcept       `json:"role,omitempty"`
	Who                       *Reference              `json:"who,omitempty"`
	AltId                     *String                 `json:"altId,omitempty"`
	AltIdPrimitiveElement     *primitiveElement       `json:"_altId,omitempty"`
	Name                      *String                 `json:"name,omitempty"`
	NamePrimitiveElement      *primitiveElement       `json:"_name,omitempty"`
	Requestor                 Boolean                 `json:"requestor,omitempty"`
	RequestorPrimitiveElement *primitiveElement       `json:"_requestor,omitempty"`
	Location                  *Reference              `json:"location,omitempty"`
	Policy                    []Uri                   `json:"policy,omitempty"`
	PolicyPrimitiveElement    []*primitiveElement     `json:"_policy,omitempty"`
	Media                     *Coding                 `json:"media,omitempty"`
	Network                   *AuditEventAgentNetwork `json:"network,omitempty"`
	PurposeOfUse              []CodeableConcept       `json:"purposeOfUse,omitempty"`
}

func (r AuditEventAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEventAgent) marshalJSON() jsonAuditEventAgent {
	m := jsonAuditEventAgent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Role = r.Role
	m.Who = r.Who
	if r.AltId != nil && r.AltId.Value != nil {
		m.AltId = r.AltId
	}
	if r.AltId != nil && (r.AltId.Id != nil || r.AltId.Extension != nil) {
		m.AltIdPrimitiveElement = &primitiveElement{Id: r.AltId.Id, Extension: r.AltId.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Requestor.Value != nil {
		m.Requestor = r.Requestor
	}
	if r.Requestor.Id != nil || r.Requestor.Extension != nil {
		m.RequestorPrimitiveElement = &primitiveElement{Id: r.Requestor.Id, Extension: r.Requestor.Extension}
	}
	m.Location = r.Location
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
	m.Media = r.Media
	m.Network = r.Network
	m.PurposeOfUse = r.PurposeOfUse
	return m
}
func (r *AuditEventAgent) UnmarshalJSON(b []byte) error {
	var m jsonAuditEventAgent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEventAgent) unmarshalJSON(m jsonAuditEventAgent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Role = m.Role
	r.Who = m.Who
	r.AltId = m.AltId
	if m.AltIdPrimitiveElement != nil {
		if r.AltId == nil {
			r.AltId = &String{}
		}
		r.AltId.Id = m.AltIdPrimitiveElement.Id
		r.AltId.Extension = m.AltIdPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Requestor = m.Requestor
	if m.RequestorPrimitiveElement != nil {
		r.Requestor.Id = m.RequestorPrimitiveElement.Id
		r.Requestor.Extension = m.RequestorPrimitiveElement.Extension
	}
	r.Location = m.Location
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
	r.Media = m.Media
	r.Network = m.Network
	r.PurposeOfUse = m.PurposeOfUse
	return nil
}
func (r AuditEventAgent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.AltId, xml.StartElement{Name: xml.Name{Local: "altId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Requestor, xml.StartElement{Name: xml.Name{Local: "requestor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Policy, xml.StartElement{Name: xml.Name{Local: "policy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Media, xml.StartElement{Name: xml.Name{Local: "media"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Network, xml.StartElement{Name: xml.Name{Local: "network"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PurposeOfUse, xml.StartElement{Name: xml.Name{Local: "purposeOfUse"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEventAgent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Who = &v
			case "altId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AltId = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "requestor":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requestor = v
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "policy":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Policy = append(r.Policy, v)
			case "media":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Media = &v
			case "network":
				var v AuditEventAgentNetwork
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Network = &v
			case "purposeOfUse":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PurposeOfUse = append(r.PurposeOfUse, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEventAgent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Logical network location for application activity, if the activity has a network location.
type AuditEventAgentNetwork struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An identifier for the network access point of the user device for the audit event.
	Address *String
	// An identifier for the type of network access point that originated the audit event.
	Type *Code
}
type jsonAuditEventAgentNetwork struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Address                 *String           `json:"address,omitempty"`
	AddressPrimitiveElement *primitiveElement `json:"_address,omitempty"`
	Type                    *Code             `json:"type,omitempty"`
	TypePrimitiveElement    *primitiveElement `json:"_type,omitempty"`
}

func (r AuditEventAgentNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEventAgentNetwork) marshalJSON() jsonAuditEventAgentNetwork {
	m := jsonAuditEventAgentNetwork{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Address != nil && r.Address.Value != nil {
		m.Address = r.Address
	}
	if r.Address != nil && (r.Address.Id != nil || r.Address.Extension != nil) {
		m.AddressPrimitiveElement = &primitiveElement{Id: r.Address.Id, Extension: r.Address.Extension}
	}
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	return m
}
func (r *AuditEventAgentNetwork) UnmarshalJSON(b []byte) error {
	var m jsonAuditEventAgentNetwork
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEventAgentNetwork) unmarshalJSON(m jsonAuditEventAgentNetwork) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Address = m.Address
	if m.AddressPrimitiveElement != nil {
		if r.Address == nil {
			r.Address = &String{}
		}
		r.Address.Id = m.AddressPrimitiveElement.Id
		r.Address.Extension = m.AddressPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	return nil
}
func (r AuditEventAgentNetwork) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEventAgentNetwork) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "address":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = &v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEventAgentNetwork) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The system that is reporting the event.
type AuditEventSource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Logical source location within the healthcare enterprise network.  For example, a hospital or other provider location within a multi-entity provider group.
	Site *String
	// Identifier of the source where the event was detected.
	Observer Reference
	// Code specifying the type of source where event originated.
	Type []Coding
}
type jsonAuditEventSource struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Site                 *String           `json:"site,omitempty"`
	SitePrimitiveElement *primitiveElement `json:"_site,omitempty"`
	Observer             Reference         `json:"observer,omitempty"`
	Type                 []Coding          `json:"type,omitempty"`
}

func (r AuditEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEventSource) marshalJSON() jsonAuditEventSource {
	m := jsonAuditEventSource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Site != nil && r.Site.Value != nil {
		m.Site = r.Site
	}
	if r.Site != nil && (r.Site.Id != nil || r.Site.Extension != nil) {
		m.SitePrimitiveElement = &primitiveElement{Id: r.Site.Id, Extension: r.Site.Extension}
	}
	m.Observer = r.Observer
	m.Type = r.Type
	return m
}
func (r *AuditEventSource) UnmarshalJSON(b []byte) error {
	var m jsonAuditEventSource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEventSource) unmarshalJSON(m jsonAuditEventSource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Site = m.Site
	if m.SitePrimitiveElement != nil {
		if r.Site == nil {
			r.Site = &String{}
		}
		r.Site.Id = m.SitePrimitiveElement.Id
		r.Site.Extension = m.SitePrimitiveElement.Extension
	}
	r.Observer = m.Observer
	r.Type = m.Type
	return nil
}
func (r AuditEventSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Site, xml.StartElement{Name: xml.Name{Local: "site"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Observer, xml.StartElement{Name: xml.Name{Local: "observer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEventSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "site":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Site = &v
			case "observer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Observer = v
			case "type":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEventSource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Specific instances of data or objects that have been accessed.
type AuditEventEntity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies a specific instance of the entity. The reference should be version specific.
	What *Reference
	// The type of the object that was involved in this audit event.
	Type *Coding
	// Code representing the role the entity played in the event being audited.
	Role *Coding
	// Identifier for the data life-cycle stage for the entity.
	Lifecycle *Coding
	// Security labels for the identified entity.
	SecurityLabel []Coding
	// A name of the entity in the audit event.
	Name *String
	// Text that describes the entity in more detail.
	Description *String
	// The query parameters for a query-type entities.
	Query *Base64Binary
	// Tagged value pairs for conveying additional information about the entity.
	Detail []AuditEventEntityDetail
}
type jsonAuditEventEntity struct {
	Id                          *string                  `json:"id,omitempty"`
	Extension                   []Extension              `json:"extension,omitempty"`
	ModifierExtension           []Extension              `json:"modifierExtension,omitempty"`
	What                        *Reference               `json:"what,omitempty"`
	Type                        *Coding                  `json:"type,omitempty"`
	Role                        *Coding                  `json:"role,omitempty"`
	Lifecycle                   *Coding                  `json:"lifecycle,omitempty"`
	SecurityLabel               []Coding                 `json:"securityLabel,omitempty"`
	Name                        *String                  `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement        `json:"_name,omitempty"`
	Description                 *String                  `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement        `json:"_description,omitempty"`
	Query                       *Base64Binary            `json:"query,omitempty"`
	QueryPrimitiveElement       *primitiveElement        `json:"_query,omitempty"`
	Detail                      []AuditEventEntityDetail `json:"detail,omitempty"`
}

func (r AuditEventEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEventEntity) marshalJSON() jsonAuditEventEntity {
	m := jsonAuditEventEntity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.What = r.What
	m.Type = r.Type
	m.Role = r.Role
	m.Lifecycle = r.Lifecycle
	m.SecurityLabel = r.SecurityLabel
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Query != nil && r.Query.Value != nil {
		m.Query = r.Query
	}
	if r.Query != nil && (r.Query.Id != nil || r.Query.Extension != nil) {
		m.QueryPrimitiveElement = &primitiveElement{Id: r.Query.Id, Extension: r.Query.Extension}
	}
	m.Detail = r.Detail
	return m
}
func (r *AuditEventEntity) UnmarshalJSON(b []byte) error {
	var m jsonAuditEventEntity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEventEntity) unmarshalJSON(m jsonAuditEventEntity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.What = m.What
	r.Type = m.Type
	r.Role = m.Role
	r.Lifecycle = m.Lifecycle
	r.SecurityLabel = m.SecurityLabel
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Query = m.Query
	if m.QueryPrimitiveElement != nil {
		if r.Query == nil {
			r.Query = &Base64Binary{}
		}
		r.Query.Id = m.QueryPrimitiveElement.Id
		r.Query.Extension = m.QueryPrimitiveElement.Extension
	}
	r.Detail = m.Detail
	return nil
}
func (r AuditEventEntity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.What, xml.StartElement{Name: xml.Name{Local: "what"}})
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
	err = e.EncodeElement(r.Lifecycle, xml.StartElement{Name: xml.Name{Local: "lifecycle"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabel, xml.StartElement{Name: xml.Name{Local: "securityLabel"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Query, xml.StartElement{Name: xml.Name{Local: "query"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEventEntity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "what":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.What = &v
			case "type":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "role":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = &v
			case "lifecycle":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Lifecycle = &v
			case "securityLabel":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabel = append(r.SecurityLabel, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "query":
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Query = &v
			case "detail":
				var v AuditEventEntityDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEventEntity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Tagged value pairs for conveying additional information about the entity.
type AuditEventEntityDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of extra detail provided in the value.
	Type String
	// The  value of the extra detail.
	Value isAuditEventEntityDetailValue
}
type isAuditEventEntityDetailValue interface {
	isAuditEventEntityDetailValue()
}

func (r String) isAuditEventEntityDetailValue()       {}
func (r Base64Binary) isAuditEventEntityDetailValue() {}

type jsonAuditEventEntityDetail struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	Type                              String            `json:"type,omitempty"`
	TypePrimitiveElement              *primitiveElement `json:"_type,omitempty"`
	ValueString                       *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement `json:"_valueString,omitempty"`
	ValueBase64Binary                 *Base64Binary     `json:"valueBase64Binary,omitempty"`
	ValueBase64BinaryPrimitiveElement *primitiveElement `json:"_valueBase64Binary,omitempty"`
}

func (r AuditEventEntityDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AuditEventEntityDetail) marshalJSON() jsonAuditEventEntityDetail {
	m := jsonAuditEventEntityDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	switch v := r.Value.(type) {
	case String:
		if v.Value != nil {
			m.ValueString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.ValueString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *AuditEventEntityDetail) UnmarshalJSON(b []byte) error {
	var m jsonAuditEventEntityDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AuditEventEntityDetail) unmarshalJSON(m jsonAuditEventEntityDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueBase64Binary != nil || m.ValueBase64BinaryPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBase64Binary
		if m.ValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.ValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.ValueBase64BinaryPrimitiveElement.Extension
		}
		r.Value = v
	}
	return nil
}
func (r AuditEventEntityDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Value.(type) {
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Base64Binary, *Base64Binary:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBase64Binary"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AuditEventEntityDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "valueString":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueBase64Binary":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AuditEventEntityDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
