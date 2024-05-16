package r4

import (
	"encoding/json"
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
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
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
	Contained                        []containedResource `json:"contained,omitempty"`
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
	m.Target = r.Target
	switch v := r.Occurred.(type) {
	case Period:
		m.OccurredPeriod = &v
	case *Period:
		m.OccurredPeriod = v
	case DateTime:
		m.OccurredDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.OccurredDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.OccurredDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.OccurredDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Recorded = r.Recorded
	if r.Recorded.Id != nil || r.Recorded.Extension != nil {
		m.RecordedPrimitiveElement = &primitiveElement{Id: r.Recorded.Id, Extension: r.Recorded.Extension}
	}
	m.Policy = r.Policy
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
		if len(r.Policy) > i {
			r.Policy[i].Id = e.Id
			r.Policy[i].Extension = e.Extension
		} else {
			r.Policy = append(r.Policy, Uri{Id: e.Id, Extension: e.Extension})
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
func (r Provenance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
func (r ProvenanceAgent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.Role = r.Role
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
func (r ProvenanceEntity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
