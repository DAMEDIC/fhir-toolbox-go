package r4b

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
)

// A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
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
	// Allows a service to provide a unique, business identifier for the request.
	Identifier []Identifier
	// A canonical URL referencing a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesCanonical []Canonical
	// A URL referencing an externally defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesUri []Uri
	// A plan, proposal or order that is fulfilled in whole or in part by this request.
	BasedOn []Reference
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []Reference
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *Identifier
	// The current state of the request. For request groups, the status reflects the status of all the requests in the group.
	Status Code
	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain.
	Intent Code
	// Indicates how quickly the request should be addressed with respect to other requests.
	Priority *Code
	// A code that identifies what the overall request group is.
	Code *CodeableConcept
	// The subject for which the request group was created.
	Subject *Reference
	// Describes the context of the request group, if any.
	Encounter *Reference
	// Indicates when the request group was created.
	AuthoredOn *DateTime
	// Provides a reference to the author of the request group.
	Author *Reference
	// Describes the reason for the request group in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this request group.
	ReasonReference []Reference
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation
	// The actions, if any, produced by the evaluation of the artifact.
	Action []RequestGroupAction
}

// The actions, if any, produced by the evaluation of the artifact.
type RequestGroupAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A user-visible prefix for the action.
	Prefix *String
	// The title of the action displayed to a user.
	Title *String
	// A short description of the action used to provide a summary to display to the user.
	Description *String
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *String
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *Code
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for a section of a documentation template.
	Code []CodeableConcept
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	Condition []RequestGroupActionCondition
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []RequestGroupActionRelatedAction
	// An optional value describing when the action should be performed.
	Timing isRequestGroupActionTiming
	// The participant that should perform or be responsible for this action.
	Participant []Reference
	// The type of action to perform (create, update, remove).
	Type *CodeableConcept
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *Code
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *Code
	// Defines expectations around whether an action is required.
	RequiredBehavior *Code
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *Code
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *Code
	// The resource that is the target of the action (e.g. CommunicationRequest).
	Resource *Reference
	// Sub actions.
	Action []RequestGroupAction
}
type isRequestGroupActionTiming interface {
	model.Element
	isRequestGroupActionTiming()
}

func (r DateTime) isRequestGroupActionTiming() {}
func (r Age) isRequestGroupActionTiming()      {}
func (r Period) isRequestGroupActionTiming()   {}
func (r Duration) isRequestGroupActionTiming() {}
func (r Range) isRequestGroupActionTiming()    {}
func (r Timing) isRequestGroupActionTiming()   {}

// An expression that describes applicability criteria, or start/stop conditions for the action.
type RequestGroupActionCondition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of condition.
	Kind Code
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	Expression *Expression
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type RequestGroupActionRelatedAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The element id of the action this is related to.
	ActionId Id
	// The relationship of this action to the related action.
	Relationship Code
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	Offset isRequestGroupActionRelatedActionOffset
}
type isRequestGroupActionRelatedActionOffset interface {
	model.Element
	isRequestGroupActionRelatedActionOffset()
}

func (r Duration) isRequestGroupActionRelatedActionOffset() {}
func (r Range) isRequestGroupActionRelatedActionOffset()    {}
func (r RequestGroup) ResourceType() string {
	return "RequestGroup"
}
func (r RequestGroup) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r RequestGroup) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.InstantiatesCanonical {
		s += i.MemSize()
	}
	s += (cap(r.InstantiatesCanonical) - len(r.InstantiatesCanonical)) * int(unsafe.Sizeof(Canonical{}))
	for _, i := range r.InstantiatesUri {
		s += i.MemSize()
	}
	s += (cap(r.InstantiatesUri) - len(r.InstantiatesUri)) * int(unsafe.Sizeof(Uri{}))
	for _, i := range r.BasedOn {
		s += i.MemSize()
	}
	s += (cap(r.BasedOn) - len(r.BasedOn)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Replaces {
		s += i.MemSize()
	}
	s += (cap(r.Replaces) - len(r.Replaces)) * int(unsafe.Sizeof(Reference{}))
	if r.GroupIdentifier != nil {
		s += r.GroupIdentifier.MemSize()
	}
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	s += r.Intent.MemSize() - int(unsafe.Sizeof(r.Intent))
	if r.Priority != nil {
		s += r.Priority.MemSize()
	}
	if r.Code != nil {
		s += r.Code.MemSize()
	}
	if r.Subject != nil {
		s += r.Subject.MemSize()
	}
	if r.Encounter != nil {
		s += r.Encounter.MemSize()
	}
	if r.AuthoredOn != nil {
		s += r.AuthoredOn.MemSize()
	}
	if r.Author != nil {
		s += r.Author.MemSize()
	}
	for _, i := range r.ReasonCode {
		s += i.MemSize()
	}
	s += (cap(r.ReasonCode) - len(r.ReasonCode)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ReasonReference {
		s += i.MemSize()
	}
	s += (cap(r.ReasonReference) - len(r.ReasonReference)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	for _, i := range r.Action {
		s += i.MemSize()
	}
	s += (cap(r.Action) - len(r.Action)) * int(unsafe.Sizeof(RequestGroupAction{}))
	return s
}
func (r RequestGroupAction) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Prefix != nil {
		s += r.Prefix.MemSize()
	}
	if r.Title != nil {
		s += r.Title.MemSize()
	}
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.TextEquivalent != nil {
		s += r.TextEquivalent.MemSize()
	}
	if r.Priority != nil {
		s += r.Priority.MemSize()
	}
	for _, i := range r.Code {
		s += i.MemSize()
	}
	s += (cap(r.Code) - len(r.Code)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.Documentation {
		s += i.MemSize()
	}
	s += (cap(r.Documentation) - len(r.Documentation)) * int(unsafe.Sizeof(RelatedArtifact{}))
	for _, i := range r.Condition {
		s += i.MemSize()
	}
	s += (cap(r.Condition) - len(r.Condition)) * int(unsafe.Sizeof(RequestGroupActionCondition{}))
	for _, i := range r.RelatedAction {
		s += i.MemSize()
	}
	s += (cap(r.RelatedAction) - len(r.RelatedAction)) * int(unsafe.Sizeof(RequestGroupActionRelatedAction{}))
	if r.Timing != nil {
		s += r.Timing.MemSize()
	}
	for _, i := range r.Participant {
		s += i.MemSize()
	}
	s += (cap(r.Participant) - len(r.Participant)) * int(unsafe.Sizeof(Reference{}))
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.GroupingBehavior != nil {
		s += r.GroupingBehavior.MemSize()
	}
	if r.SelectionBehavior != nil {
		s += r.SelectionBehavior.MemSize()
	}
	if r.RequiredBehavior != nil {
		s += r.RequiredBehavior.MemSize()
	}
	if r.PrecheckBehavior != nil {
		s += r.PrecheckBehavior.MemSize()
	}
	if r.CardinalityBehavior != nil {
		s += r.CardinalityBehavior.MemSize()
	}
	if r.Resource != nil {
		s += r.Resource.MemSize()
	}
	for _, i := range r.Action {
		s += i.MemSize()
	}
	s += (cap(r.Action) - len(r.Action)) * int(unsafe.Sizeof(RequestGroupAction{}))
	return s
}
func (r RequestGroupActionCondition) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Kind.MemSize() - int(unsafe.Sizeof(r.Kind))
	if r.Expression != nil {
		s += r.Expression.MemSize()
	}
	return s
}
func (r RequestGroupActionRelatedAction) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.ActionId.MemSize() - int(unsafe.Sizeof(r.ActionId))
	s += r.Relationship.MemSize() - int(unsafe.Sizeof(r.Relationship))
	if r.Offset != nil {
		s += r.Offset.MemSize()
	}
	return s
}
func (r RequestGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RequestGroupAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RequestGroupActionCondition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RequestGroupActionRelatedAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RequestGroup) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"RequestGroup\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.InstantiatesCanonical {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"instantiatesCanonical\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.InstantiatesCanonical)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_instantiatesCanonical\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.InstantiatesCanonical {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.InstantiatesUri {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"instantiatesUri\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.InstantiatesUri)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_instantiatesUri\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.InstantiatesUri {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if len(r.BasedOn) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"basedOn\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BasedOn {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Replaces) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"replaces\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Replaces {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.GroupIdentifier != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"groupIdentifier\":"))
		if err != nil {
			return err
		}
		err = r.GroupIdentifier.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"intent\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Intent)
		if err != nil {
			return err
		}
	}
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		p := primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_intent\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && r.Priority.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"priority\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Priority)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		p := primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_priority\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Code != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"code\":"))
		if err != nil {
			return err
		}
		err = r.Code.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Subject != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subject\":"))
		if err != nil {
			return err
		}
		err = r.Subject.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Encounter != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"encounter\":"))
		if err != nil {
			return err
		}
		err = r.Encounter.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AuthoredOn != nil && r.AuthoredOn.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authoredOn\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AuthoredOn)
		if err != nil {
			return err
		}
	}
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		p := primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_authoredOn\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Author != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"author\":"))
		if err != nil {
			return err
		}
		err = r.Author.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.ReasonCode) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reasonCode\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ReasonCode {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ReasonReference) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reasonReference\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ReasonReference {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Note) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"note\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Note {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Action) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"action\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Action {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupAction) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RequestGroupAction) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Prefix != nil && r.Prefix.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"prefix\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Prefix)
		if err != nil {
			return err
		}
	}
	if r.Prefix != nil && (r.Prefix.Id != nil || r.Prefix.Extension != nil) {
		p := primitiveElement{Id: r.Prefix.Id, Extension: r.Prefix.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_prefix\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Title != nil && r.Title.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"title\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Title)
		if err != nil {
			return err
		}
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		p := primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_title\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.TextEquivalent != nil && r.TextEquivalent.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"textEquivalent\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.TextEquivalent)
		if err != nil {
			return err
		}
	}
	if r.TextEquivalent != nil && (r.TextEquivalent.Id != nil || r.TextEquivalent.Extension != nil) {
		p := primitiveElement{Id: r.TextEquivalent.Id, Extension: r.TextEquivalent.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_textEquivalent\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && r.Priority.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"priority\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Priority)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		p := primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_priority\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Code) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"code\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Code {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Documentation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"documentation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Documentation {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Condition) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"condition\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Condition {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.RelatedAction) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relatedAction\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RelatedAction {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	switch v := r.Timing.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timingDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_timingDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timingDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_timingDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Timing:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Timing:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timingTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	if len(r.Participant) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"participant\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Participant {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Type != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		err = r.Type.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.GroupingBehavior != nil && r.GroupingBehavior.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"groupingBehavior\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.GroupingBehavior)
		if err != nil {
			return err
		}
	}
	if r.GroupingBehavior != nil && (r.GroupingBehavior.Id != nil || r.GroupingBehavior.Extension != nil) {
		p := primitiveElement{Id: r.GroupingBehavior.Id, Extension: r.GroupingBehavior.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_groupingBehavior\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SelectionBehavior != nil && r.SelectionBehavior.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"selectionBehavior\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SelectionBehavior)
		if err != nil {
			return err
		}
	}
	if r.SelectionBehavior != nil && (r.SelectionBehavior.Id != nil || r.SelectionBehavior.Extension != nil) {
		p := primitiveElement{Id: r.SelectionBehavior.Id, Extension: r.SelectionBehavior.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_selectionBehavior\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.RequiredBehavior != nil && r.RequiredBehavior.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"requiredBehavior\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.RequiredBehavior)
		if err != nil {
			return err
		}
	}
	if r.RequiredBehavior != nil && (r.RequiredBehavior.Id != nil || r.RequiredBehavior.Extension != nil) {
		p := primitiveElement{Id: r.RequiredBehavior.Id, Extension: r.RequiredBehavior.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_requiredBehavior\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PrecheckBehavior != nil && r.PrecheckBehavior.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"precheckBehavior\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PrecheckBehavior)
		if err != nil {
			return err
		}
	}
	if r.PrecheckBehavior != nil && (r.PrecheckBehavior.Id != nil || r.PrecheckBehavior.Extension != nil) {
		p := primitiveElement{Id: r.PrecheckBehavior.Id, Extension: r.PrecheckBehavior.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_precheckBehavior\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CardinalityBehavior != nil && r.CardinalityBehavior.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"cardinalityBehavior\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CardinalityBehavior)
		if err != nil {
			return err
		}
	}
	if r.CardinalityBehavior != nil && (r.CardinalityBehavior.Id != nil || r.CardinalityBehavior.Extension != nil) {
		p := primitiveElement{Id: r.CardinalityBehavior.Id, Extension: r.CardinalityBehavior.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_cardinalityBehavior\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Resource != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"resource\":"))
		if err != nil {
			return err
		}
		err = r.Resource.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Action) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"action\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Action {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupActionCondition) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RequestGroupActionCondition) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"kind\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Kind)
		if err != nil {
			return err
		}
	}
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		p := primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_kind\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Expression != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"expression\":"))
		if err != nil {
			return err
		}
		err = r.Expression.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupActionRelatedAction) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r RequestGroupActionRelatedAction) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"actionId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ActionId)
		if err != nil {
			return err
		}
	}
	if r.ActionId.Id != nil || r.ActionId.Extension != nil {
		p := primitiveElement{Id: r.ActionId.Id, Extension: r.ActionId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_actionId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relationship\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Relationship)
		if err != nil {
			return err
		}
	}
	if r.Relationship.Id != nil || r.Relationship.Extension != nil {
		p := primitiveElement{Id: r.Relationship.Id, Extension: r.Relationship.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_relationship\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Offset.(type) {
	case Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offsetDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offsetDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offsetRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offsetRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *RequestGroup) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *RequestGroup) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RequestGroup element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RequestGroup element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "instantiatesCanonical":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for i := 0; d.More(); i++ {
				var v Canonical
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.InstantiatesCanonical) <= i {
					r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
				}
				r.InstantiatesCanonical[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "_instantiatesCanonical":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.InstantiatesCanonical) <= i {
					r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
				}
				r.InstantiatesCanonical[i].Id = v.Id
				r.InstantiatesCanonical[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "instantiatesUri":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for i := 0; d.More(); i++ {
				var v Uri
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.InstantiatesUri) <= i {
					r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
				}
				r.InstantiatesUri[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "_instantiatesUri":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.InstantiatesUri) <= i {
					r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
				}
				r.InstantiatesUri[i].Id = v.Id
				r.InstantiatesUri[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "basedOn":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "replaces":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Replaces = append(r.Replaces, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "groupIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.GroupIdentifier = &v
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "intent":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Intent.Value = v.Value
		case "_intent":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Intent.Id = v.Id
			r.Intent.Extension = v.Extension
		case "priority":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Value = v.Value
		case "_priority":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Id = v.Id
			r.Priority.Extension = v.Extension
		case "code":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Code = &v
		case "subject":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Subject = &v
		case "encounter":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Encounter = &v
		case "authoredOn":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AuthoredOn == nil {
				r.AuthoredOn = &DateTime{}
			}
			r.AuthoredOn.Value = v.Value
		case "_authoredOn":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AuthoredOn == nil {
				r.AuthoredOn = &DateTime{}
			}
			r.AuthoredOn.Id = v.Id
			r.AuthoredOn.Extension = v.Extension
		case "author":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Author = &v
		case "reasonCode":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ReasonCode = append(r.ReasonCode, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "reasonReference":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ReasonReference = append(r.ReasonReference, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v Annotation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		case "action":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroup element", t)
			}
			for d.More() {
				var v RequestGroupAction
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroup element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in RequestGroup", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RequestGroup element", t)
	}
	return nil
}
func (r *RequestGroupAction) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RequestGroupAction element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RequestGroupAction element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "prefix":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Prefix == nil {
				r.Prefix = &String{}
			}
			r.Prefix.Value = v.Value
		case "_prefix":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Prefix == nil {
				r.Prefix = &String{}
			}
			r.Prefix.Id = v.Id
			r.Prefix.Extension = v.Extension
		case "title":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Title == nil {
				r.Title = &String{}
			}
			r.Title.Value = v.Value
		case "_title":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Title == nil {
				r.Title = &String{}
			}
			r.Title.Id = v.Id
			r.Title.Extension = v.Extension
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "textEquivalent":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.TextEquivalent == nil {
				r.TextEquivalent = &String{}
			}
			r.TextEquivalent.Value = v.Value
		case "_textEquivalent":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.TextEquivalent == nil {
				r.TextEquivalent = &String{}
			}
			r.TextEquivalent.Id = v.Id
			r.TextEquivalent.Extension = v.Extension
		case "priority":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Value = v.Value
		case "_priority":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Id = v.Id
			r.Priority.Extension = v.Extension
		case "code":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "documentation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v RelatedArtifact
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Documentation = append(r.Documentation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "condition":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v RequestGroupActionCondition
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Condition = append(r.Condition, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "relatedAction":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v RequestGroupActionRelatedAction
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RelatedAction = append(r.RelatedAction, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "timingDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Timing != nil {
				r.Timing = DateTime{
					Extension: r.Timing.(DateTime).Extension,
					Id:        r.Timing.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Timing = v
			}
		case "_timingDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Timing != nil {
				r.Timing = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Timing.(DateTime).Value,
				}
			} else {
				r.Timing = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "timingAge":
			var v Age
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = v
		case "timingPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = v
		case "timingDuration":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = v
		case "timingRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = v
		case "timingTiming":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = v
		case "participant":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Participant = append(r.Participant, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "groupingBehavior":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.GroupingBehavior == nil {
				r.GroupingBehavior = &Code{}
			}
			r.GroupingBehavior.Value = v.Value
		case "_groupingBehavior":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.GroupingBehavior == nil {
				r.GroupingBehavior = &Code{}
			}
			r.GroupingBehavior.Id = v.Id
			r.GroupingBehavior.Extension = v.Extension
		case "selectionBehavior":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SelectionBehavior == nil {
				r.SelectionBehavior = &Code{}
			}
			r.SelectionBehavior.Value = v.Value
		case "_selectionBehavior":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SelectionBehavior == nil {
				r.SelectionBehavior = &Code{}
			}
			r.SelectionBehavior.Id = v.Id
			r.SelectionBehavior.Extension = v.Extension
		case "requiredBehavior":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.RequiredBehavior == nil {
				r.RequiredBehavior = &Code{}
			}
			r.RequiredBehavior.Value = v.Value
		case "_requiredBehavior":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.RequiredBehavior == nil {
				r.RequiredBehavior = &Code{}
			}
			r.RequiredBehavior.Id = v.Id
			r.RequiredBehavior.Extension = v.Extension
		case "precheckBehavior":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PrecheckBehavior == nil {
				r.PrecheckBehavior = &Code{}
			}
			r.PrecheckBehavior.Value = v.Value
		case "_precheckBehavior":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PrecheckBehavior == nil {
				r.PrecheckBehavior = &Code{}
			}
			r.PrecheckBehavior.Id = v.Id
			r.PrecheckBehavior.Extension = v.Extension
		case "cardinalityBehavior":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CardinalityBehavior == nil {
				r.CardinalityBehavior = &Code{}
			}
			r.CardinalityBehavior.Value = v.Value
		case "_cardinalityBehavior":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CardinalityBehavior == nil {
				r.CardinalityBehavior = &Code{}
			}
			r.CardinalityBehavior.Id = v.Id
			r.CardinalityBehavior.Extension = v.Extension
		case "resource":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Resource = &v
		case "action":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupAction element", t)
			}
			for d.More() {
				var v RequestGroupAction
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupAction element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in RequestGroupAction", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RequestGroupAction element", t)
	}
	return nil
}
func (r *RequestGroupActionCondition) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RequestGroupActionCondition element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RequestGroupActionCondition element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupActionCondition element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupActionCondition element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupActionCondition element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupActionCondition element", t)
			}
		case "kind":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Kind.Value = v.Value
		case "_kind":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Kind.Id = v.Id
			r.Kind.Extension = v.Extension
		case "expression":
			var v Expression
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Expression = &v
		default:
			return fmt.Errorf("invalid field: %s in RequestGroupActionCondition", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RequestGroupActionCondition element", t)
	}
	return nil
}
func (r *RequestGroupActionRelatedAction) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in RequestGroupActionRelatedAction element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in RequestGroupActionRelatedAction element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupActionRelatedAction element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupActionRelatedAction element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in RequestGroupActionRelatedAction element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in RequestGroupActionRelatedAction element", t)
			}
		case "actionId":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.ActionId.Value = v.Value
		case "_actionId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ActionId.Id = v.Id
			r.ActionId.Extension = v.Extension
		case "relationship":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Relationship.Value = v.Value
		case "_relationship":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Relationship.Id = v.Id
			r.Relationship.Extension = v.Extension
		case "offsetDuration":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Offset = v
		case "offsetRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Offset = v
		default:
			return fmt.Errorf("invalid field: %s in RequestGroupActionRelatedAction", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in RequestGroupActionRelatedAction element", t)
	}
	return nil
}
func (r RequestGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "RequestGroup"
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
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
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
	err = e.EncodeElement(r.InstantiatesCanonical, xml.StartElement{Name: xml.Name{Local: "instantiatesCanonical"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InstantiatesUri, xml.StartElement{Name: xml.Name{Local: "instantiatesUri"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Replaces, xml.StartElement{Name: xml.Name{Local: "replaces"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GroupIdentifier, xml.StartElement{Name: xml.Name{Local: "groupIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Intent, xml.StartElement{Name: xml.Name{Local: "intent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthoredOn, xml.StartElement{Name: xml.Name{Local: "authoredOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonCode, xml.StartElement{Name: xml.Name{Local: "reasonCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonReference, xml.StartElement{Name: xml.Name{Local: "reasonReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Prefix, xml.StartElement{Name: xml.Name{Local: "prefix"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TextEquivalent, xml.StartElement{Name: xml.Name{Local: "textEquivalent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Condition, xml.StartElement{Name: xml.Name{Local: "condition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatedAction, xml.StartElement{Name: xml.Name{Local: "relatedAction"}})
	if err != nil {
		return err
	}
	switch v := r.Timing.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingDateTime"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingAge"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingPeriod"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingDuration"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingRange"}})
		if err != nil {
			return err
		}
	case Timing, *Timing:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timingTiming"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Participant, xml.StartElement{Name: xml.Name{Local: "participant"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GroupingBehavior, xml.StartElement{Name: xml.Name{Local: "groupingBehavior"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SelectionBehavior, xml.StartElement{Name: xml.Name{Local: "selectionBehavior"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequiredBehavior, xml.StartElement{Name: xml.Name{Local: "requiredBehavior"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PrecheckBehavior, xml.StartElement{Name: xml.Name{Local: "precheckBehavior"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CardinalityBehavior, xml.StartElement{Name: xml.Name{Local: "cardinalityBehavior"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupActionCondition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Kind, xml.StartElement{Name: xml.Name{Local: "kind"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Expression, xml.StartElement{Name: xml.Name{Local: "expression"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r RequestGroupActionRelatedAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ActionId, xml.StartElement{Name: xml.Name{Local: "actionId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relationship, xml.StartElement{Name: xml.Name{Local: "relationship"}})
	if err != nil {
		return err
	}
	switch v := r.Offset.(type) {
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "offsetDuration"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "offsetRange"}})
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
func (r *RequestGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "instantiatesCanonical":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesCanonical = append(r.InstantiatesCanonical, v)
			case "instantiatesUri":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesUri = append(r.InstantiatesUri, v)
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "replaces":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Replaces = append(r.Replaces, v)
			case "groupIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GroupIdentifier = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "intent":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Intent = v
			case "priority":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "authoredOn":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthoredOn = &v
			case "author":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = &v
			case "reasonCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonCode = append(r.ReasonCode, v)
			case "reasonReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonReference = append(r.ReasonReference, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "action":
				var v RequestGroupAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *RequestGroupAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "prefix":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Prefix = &v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "textEquivalent":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TextEquivalent = &v
			case "priority":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			case "documentation":
				var v RelatedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = append(r.Documentation, v)
			case "condition":
				var v RequestGroupActionCondition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Condition = append(r.Condition, v)
			case "relatedAction":
				var v RequestGroupActionRelatedAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatedAction = append(r.RelatedAction, v)
			case "timingDateTime":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "timingAge":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "timingPeriod":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "timingDuration":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "timingRange":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "timingTiming":
				if r.Timing != nil {
					return fmt.Errorf("multiple values for field \"Timing\"")
				}
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = v
			case "participant":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Participant = append(r.Participant, v)
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "groupingBehavior":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GroupingBehavior = &v
			case "selectionBehavior":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SelectionBehavior = &v
			case "requiredBehavior":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequiredBehavior = &v
			case "precheckBehavior":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PrecheckBehavior = &v
			case "cardinalityBehavior":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CardinalityBehavior = &v
			case "resource":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = &v
			case "action":
				var v RequestGroupAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *RequestGroupActionCondition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "kind":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Kind = v
			case "expression":
				var v Expression
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Expression = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *RequestGroupActionRelatedAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "actionId":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ActionId = v
			case "relationship":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationship = v
			case "offsetDuration":
				if r.Offset != nil {
					return fmt.Errorf("multiple values for field \"Offset\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Offset = v
			case "offsetRange":
				if r.Offset != nil {
					return fmt.Errorf("multiple values for field \"Offset\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Offset = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r RequestGroup) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instantiatesCanonical") {
		for _, v := range r.InstantiatesCanonical {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instantiatesUri") {
		for _, v := range r.InstantiatesUri {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "basedOn") {
		for _, v := range r.BasedOn {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "replaces") {
		for _, v := range r.Replaces {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "groupIdentifier") {
		if r.GroupIdentifier != nil {
			children = append(children, *r.GroupIdentifier)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "intent") {
		children = append(children, r.Intent)
	}
	if len(name) == 0 || slices.Contains(name, "priority") {
		if r.Priority != nil {
			children = append(children, *r.Priority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		if r.Code != nil {
			children = append(children, *r.Code)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subject") {
		if r.Subject != nil {
			children = append(children, *r.Subject)
		}
	}
	if len(name) == 0 || slices.Contains(name, "encounter") {
		if r.Encounter != nil {
			children = append(children, *r.Encounter)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authoredOn") {
		if r.AuthoredOn != nil {
			children = append(children, *r.AuthoredOn)
		}
	}
	if len(name) == 0 || slices.Contains(name, "author") {
		if r.Author != nil {
			children = append(children, *r.Author)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reasonCode") {
		for _, v := range r.ReasonCode {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reasonReference") {
		for _, v := range r.ReasonReference {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "action") {
		for _, v := range r.Action {
			children = append(children, v)
		}
	}
	return children
}
func (r RequestGroup) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RequestGroup to Boolean")
}
func (r RequestGroup) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RequestGroup to String")
}
func (r RequestGroup) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RequestGroup to Integer")
}
func (r RequestGroup) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RequestGroup to Decimal")
}
func (r RequestGroup) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RequestGroup to Date")
}
func (r RequestGroup) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RequestGroup to Time")
}
func (r RequestGroup) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RequestGroup to DateTime")
}
func (r RequestGroup) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RequestGroup to Quantity")
}
func (r RequestGroup) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroup
	switch other := other.(type) {
	case RequestGroup:
		o = other
	case *RequestGroup:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroup) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroup
	switch other := other.(type) {
	case RequestGroup:
		o = other
	case *RequestGroup:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroup) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "InstantiatesCanonical",
			Type: "List<FHIR.Canonical>",
		}, {
			Name: "InstantiatesUri",
			Type: "List<FHIR.Uri>",
		}, {
			Name: "BasedOn",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Replaces",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "GroupIdentifier",
			Type: "FHIR.Identifier",
		}, {
			Name: "Status",
			Type: "FHIR.Code",
		}, {
			Name: "Intent",
			Type: "FHIR.Code",
		}, {
			Name: "Priority",
			Type: "FHIR.Code",
		}, {
			Name: "Code",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Subject",
			Type: "FHIR.Reference",
		}, {
			Name: "Encounter",
			Type: "FHIR.Reference",
		}, {
			Name: "AuthoredOn",
			Type: "FHIR.DateTime",
		}, {
			Name: "Author",
			Type: "FHIR.Reference",
		}, {
			Name: "ReasonCode",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "ReasonReference",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Note",
			Type: "List<FHIR.Annotation>",
		}, {
			Name: "Action",
			Type: "List<FHIR.RequestGroupAction>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "RequestGroup",
			Namespace: "FHIR",
		},
	}
}
func (r RequestGroupAction) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "prefix") {
		if r.Prefix != nil {
			children = append(children, *r.Prefix)
		}
	}
	if len(name) == 0 || slices.Contains(name, "title") {
		if r.Title != nil {
			children = append(children, *r.Title)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "textEquivalent") {
		if r.TextEquivalent != nil {
			children = append(children, *r.TextEquivalent)
		}
	}
	if len(name) == 0 || slices.Contains(name, "priority") {
		if r.Priority != nil {
			children = append(children, *r.Priority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		for _, v := range r.Code {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "documentation") {
		for _, v := range r.Documentation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "condition") {
		for _, v := range r.Condition {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relatedAction") {
		for _, v := range r.RelatedAction {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timing") {
		if r.Timing != nil {
			children = append(children, r.Timing)
		}
	}
	if len(name) == 0 || slices.Contains(name, "participant") {
		for _, v := range r.Participant {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "groupingBehavior") {
		if r.GroupingBehavior != nil {
			children = append(children, *r.GroupingBehavior)
		}
	}
	if len(name) == 0 || slices.Contains(name, "selectionBehavior") {
		if r.SelectionBehavior != nil {
			children = append(children, *r.SelectionBehavior)
		}
	}
	if len(name) == 0 || slices.Contains(name, "requiredBehavior") {
		if r.RequiredBehavior != nil {
			children = append(children, *r.RequiredBehavior)
		}
	}
	if len(name) == 0 || slices.Contains(name, "precheckBehavior") {
		if r.PrecheckBehavior != nil {
			children = append(children, *r.PrecheckBehavior)
		}
	}
	if len(name) == 0 || slices.Contains(name, "cardinalityBehavior") {
		if r.CardinalityBehavior != nil {
			children = append(children, *r.CardinalityBehavior)
		}
	}
	if len(name) == 0 || slices.Contains(name, "resource") {
		if r.Resource != nil {
			children = append(children, *r.Resource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "action") {
		for _, v := range r.Action {
			children = append(children, v)
		}
	}
	return children
}
func (r RequestGroupAction) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RequestGroupAction to Boolean")
}
func (r RequestGroupAction) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RequestGroupAction to String")
}
func (r RequestGroupAction) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RequestGroupAction to Integer")
}
func (r RequestGroupAction) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RequestGroupAction to Decimal")
}
func (r RequestGroupAction) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RequestGroupAction to Date")
}
func (r RequestGroupAction) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RequestGroupAction to Time")
}
func (r RequestGroupAction) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RequestGroupAction to DateTime")
}
func (r RequestGroupAction) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RequestGroupAction to Quantity")
}
func (r RequestGroupAction) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupAction
	switch other := other.(type) {
	case RequestGroupAction:
		o = other
	case *RequestGroupAction:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupAction) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupAction
	switch other := other.(type) {
	case RequestGroupAction:
		o = other
	case *RequestGroupAction:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupAction) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Prefix",
			Type: "FHIR.String",
		}, {
			Name: "Title",
			Type: "FHIR.String",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "TextEquivalent",
			Type: "FHIR.String",
		}, {
			Name: "Priority",
			Type: "FHIR.Code",
		}, {
			Name: "Code",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "Documentation",
			Type: "List<FHIR.RelatedArtifact>",
		}, {
			Name: "Condition",
			Type: "List<FHIR.RequestGroupActionCondition>",
		}, {
			Name: "RelatedAction",
			Type: "List<FHIR.RequestGroupActionRelatedAction>",
		}, {
			Name: "Timing",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Participant",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "GroupingBehavior",
			Type: "FHIR.Code",
		}, {
			Name: "SelectionBehavior",
			Type: "FHIR.Code",
		}, {
			Name: "RequiredBehavior",
			Type: "FHIR.Code",
		}, {
			Name: "PrecheckBehavior",
			Type: "FHIR.Code",
		}, {
			Name: "CardinalityBehavior",
			Type: "FHIR.Code",
		}, {
			Name: "Resource",
			Type: "FHIR.Reference",
		}, {
			Name: "Action",
			Type: "List<FHIR.RequestGroupAction>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "RequestGroupAction",
			Namespace: "FHIR",
		},
	}
}
func (r RequestGroupActionCondition) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "kind") {
		children = append(children, r.Kind)
	}
	if len(name) == 0 || slices.Contains(name, "expression") {
		if r.Expression != nil {
			children = append(children, *r.Expression)
		}
	}
	return children
}
func (r RequestGroupActionCondition) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Boolean")
}
func (r RequestGroupActionCondition) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to String")
}
func (r RequestGroupActionCondition) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Integer")
}
func (r RequestGroupActionCondition) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Decimal")
}
func (r RequestGroupActionCondition) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Date")
}
func (r RequestGroupActionCondition) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Time")
}
func (r RequestGroupActionCondition) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to DateTime")
}
func (r RequestGroupActionCondition) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RequestGroupActionCondition to Quantity")
}
func (r RequestGroupActionCondition) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupActionCondition
	switch other := other.(type) {
	case RequestGroupActionCondition:
		o = other
	case *RequestGroupActionCondition:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupActionCondition) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupActionCondition
	switch other := other.(type) {
	case RequestGroupActionCondition:
		o = other
	case *RequestGroupActionCondition:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupActionCondition) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Kind",
			Type: "FHIR.Code",
		}, {
			Name: "Expression",
			Type: "FHIR.Expression",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "RequestGroupActionCondition",
			Namespace: "FHIR",
		},
	}
}
func (r RequestGroupActionRelatedAction) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "actionId") {
		children = append(children, r.ActionId)
	}
	if len(name) == 0 || slices.Contains(name, "relationship") {
		children = append(children, r.Relationship)
	}
	if len(name) == 0 || slices.Contains(name, "offset") {
		if r.Offset != nil {
			children = append(children, r.Offset)
		}
	}
	return children
}
func (r RequestGroupActionRelatedAction) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Boolean")
}
func (r RequestGroupActionRelatedAction) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to String")
}
func (r RequestGroupActionRelatedAction) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Integer")
}
func (r RequestGroupActionRelatedAction) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Decimal")
}
func (r RequestGroupActionRelatedAction) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Date")
}
func (r RequestGroupActionRelatedAction) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Time")
}
func (r RequestGroupActionRelatedAction) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to DateTime")
}
func (r RequestGroupActionRelatedAction) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert RequestGroupActionRelatedAction to Quantity")
}
func (r RequestGroupActionRelatedAction) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupActionRelatedAction
	switch other := other.(type) {
	case RequestGroupActionRelatedAction:
		o = other
	case *RequestGroupActionRelatedAction:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupActionRelatedAction) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o RequestGroupActionRelatedAction
	switch other := other.(type) {
	case RequestGroupActionRelatedAction:
		o = other
	case *RequestGroupActionRelatedAction:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r RequestGroupActionRelatedAction) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ActionId",
			Type: "FHIR.Id",
		}, {
			Name: "Relationship",
			Type: "FHIR.Code",
		}, {
			Name: "Offset",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "RequestGroupActionRelatedAction",
			Namespace: "FHIR",
		},
	}
}
