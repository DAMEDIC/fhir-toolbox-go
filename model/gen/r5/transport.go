package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"unsafe"
)

// Record of transport.
type Transport struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifier for the transport event that is used to identify it across multiple disparate systems.
	Identifier []Identifier
	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Transport.
	InstantiatesCanonical *Canonical
	// The URL pointing to an *externally* maintained  protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Transport.
	InstantiatesUri *Uri
	// BasedOn refers to a higher-level authorization that triggered the creation of the transport.  It references a "request" resource such as a ServiceRequest or Transport, which is distinct from the "request" resource the Transport is seeking to fulfill.  This latter resource is referenced by FocusOn.  For example, based on a ServiceRequest (= BasedOn), a transport is created to fulfill a procedureRequest ( = FocusOn ) to transport a specimen to the lab.
	BasedOn []Reference
	// A shared identifier common to multiple independent Request instances that were activated/authorized more or less simultaneously by a single author.  The presence of the same identifier on each request ties those requests together and may have business ramifications in terms of reporting of results, billing, etc.  E.g. a requisition number shared by a set of lab tests ordered together, or a prescription number shared by all meds ordered at one time.
	GroupIdentifier *Identifier
	// A larger event of which this particular event is a component or step.
	PartOf []Reference
	// A code specifying the state of the transport event.
	Status *Code
	// An explanation as to why this transport is held, failed, was refused, etc.
	StatusReason *CodeableConcept
	// Indicates the "level" of actionability associated with the Transport, i.e. i+R[9]Cs this a proposed transport, a planned transport, an actionable transport, etc.
	Intent Code
	// Indicates how quickly the Transport should be addressed with respect to other requests.
	Priority *Code
	// A name or code (or both) briefly describing what the transport involves.
	Code *CodeableConcept
	// A free-text description of what is to be performed.
	Description *String
	// The request being actioned or the resource being manipulated by this transport.
	Focus *Reference
	// The entity who benefits from the performance of the service specified in the transport (e.g., the patient).
	For *Reference
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this transport was created.
	Encounter *Reference
	// Identifies the completion time of the event (the occurrence).
	CompletionTime *DateTime
	// The date and time this transport was created.
	AuthoredOn *DateTime
	// The date and time of last modification to this transport.
	LastModified *DateTime
	// The creator of the transport.
	Requester *Reference
	// The kind of participant that should perform the transport.
	PerformerType []CodeableConcept
	// Individual organization or Device currently responsible for transport execution.
	Owner *Reference
	// Principal physical location where this transport is performed.
	Location *Reference
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Transport.
	Insurance []Reference
	// Free-text information captured about the transport as it progresses.
	Note []Annotation
	// Links to Provenance records for past versions of this Transport that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the transport.
	RelevantHistory []Reference
	// If the Transport.focus is a request resource and the transport is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	Restriction *TransportRestriction
	// Additional information that may be needed in the execution of the transport.
	Input []TransportInput
	// Outputs produced by the Transport.
	Output []TransportOutput
	// The desired or final location for the transport.
	RequestedLocation Reference
	// The current location for the entity to be transported.
	CurrentLocation Reference
	// A resource reference indicating why this transport needs to be performed.
	Reason *CodeableReference
	// The transport event prior to this one.
	History *Reference
}

// If the Transport.focus is a request resource and the transport is seeking fulfillment (i.e. is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
type TransportRestriction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the number of times the requested action should occur.
	Repetitions *PositiveInt
	// Over what time-period is fulfillment sought.
	Period *Period
	// For requests that are targeted to more than one potential recipient/target, to identify who is fulfillment is sought for.
	Recipient []Reference
}

// Additional information that may be needed in the execution of the transport.
type TransportInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code or description indicating how the input is intended to be used as part of the transport execution.
	Type CodeableConcept
	// The value of the input parameter as a basic type.
	Value isTransportInputValue
}
type isTransportInputValue interface {
	model.Element
	isTransportInputValue()
}

func (r Base64Binary) isTransportInputValue()          {}
func (r Boolean) isTransportInputValue()               {}
func (r Canonical) isTransportInputValue()             {}
func (r Code) isTransportInputValue()                  {}
func (r Date) isTransportInputValue()                  {}
func (r DateTime) isTransportInputValue()              {}
func (r Decimal) isTransportInputValue()               {}
func (r Id) isTransportInputValue()                    {}
func (r Instant) isTransportInputValue()               {}
func (r Integer) isTransportInputValue()               {}
func (r Integer64) isTransportInputValue()             {}
func (r Markdown) isTransportInputValue()              {}
func (r Oid) isTransportInputValue()                   {}
func (r PositiveInt) isTransportInputValue()           {}
func (r String) isTransportInputValue()                {}
func (r Time) isTransportInputValue()                  {}
func (r UnsignedInt) isTransportInputValue()           {}
func (r Uri) isTransportInputValue()                   {}
func (r Url) isTransportInputValue()                   {}
func (r Uuid) isTransportInputValue()                  {}
func (r Address) isTransportInputValue()               {}
func (r Age) isTransportInputValue()                   {}
func (r Annotation) isTransportInputValue()            {}
func (r Attachment) isTransportInputValue()            {}
func (r CodeableConcept) isTransportInputValue()       {}
func (r CodeableReference) isTransportInputValue()     {}
func (r Coding) isTransportInputValue()                {}
func (r ContactPoint) isTransportInputValue()          {}
func (r Count) isTransportInputValue()                 {}
func (r Distance) isTransportInputValue()              {}
func (r Duration) isTransportInputValue()              {}
func (r HumanName) isTransportInputValue()             {}
func (r Identifier) isTransportInputValue()            {}
func (r Money) isTransportInputValue()                 {}
func (r Period) isTransportInputValue()                {}
func (r Quantity) isTransportInputValue()              {}
func (r Range) isTransportInputValue()                 {}
func (r Ratio) isTransportInputValue()                 {}
func (r RatioRange) isTransportInputValue()            {}
func (r Reference) isTransportInputValue()             {}
func (r SampledData) isTransportInputValue()           {}
func (r Signature) isTransportInputValue()             {}
func (r Timing) isTransportInputValue()                {}
func (r ContactDetail) isTransportInputValue()         {}
func (r DataRequirement) isTransportInputValue()       {}
func (r Expression) isTransportInputValue()            {}
func (r ParameterDefinition) isTransportInputValue()   {}
func (r RelatedArtifact) isTransportInputValue()       {}
func (r TriggerDefinition) isTransportInputValue()     {}
func (r UsageContext) isTransportInputValue()          {}
func (r Availability) isTransportInputValue()          {}
func (r ExtendedContactDetail) isTransportInputValue() {}
func (r Dosage) isTransportInputValue()                {}
func (r Meta) isTransportInputValue()                  {}

// Outputs produced by the Transport.
type TransportOutput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the Output parameter.
	Type CodeableConcept
	// The value of the Output parameter as a basic type.
	Value isTransportOutputValue
}
type isTransportOutputValue interface {
	model.Element
	isTransportOutputValue()
}

func (r Base64Binary) isTransportOutputValue()          {}
func (r Boolean) isTransportOutputValue()               {}
func (r Canonical) isTransportOutputValue()             {}
func (r Code) isTransportOutputValue()                  {}
func (r Date) isTransportOutputValue()                  {}
func (r DateTime) isTransportOutputValue()              {}
func (r Decimal) isTransportOutputValue()               {}
func (r Id) isTransportOutputValue()                    {}
func (r Instant) isTransportOutputValue()               {}
func (r Integer) isTransportOutputValue()               {}
func (r Integer64) isTransportOutputValue()             {}
func (r Markdown) isTransportOutputValue()              {}
func (r Oid) isTransportOutputValue()                   {}
func (r PositiveInt) isTransportOutputValue()           {}
func (r String) isTransportOutputValue()                {}
func (r Time) isTransportOutputValue()                  {}
func (r UnsignedInt) isTransportOutputValue()           {}
func (r Uri) isTransportOutputValue()                   {}
func (r Url) isTransportOutputValue()                   {}
func (r Uuid) isTransportOutputValue()                  {}
func (r Address) isTransportOutputValue()               {}
func (r Age) isTransportOutputValue()                   {}
func (r Annotation) isTransportOutputValue()            {}
func (r Attachment) isTransportOutputValue()            {}
func (r CodeableConcept) isTransportOutputValue()       {}
func (r CodeableReference) isTransportOutputValue()     {}
func (r Coding) isTransportOutputValue()                {}
func (r ContactPoint) isTransportOutputValue()          {}
func (r Count) isTransportOutputValue()                 {}
func (r Distance) isTransportOutputValue()              {}
func (r Duration) isTransportOutputValue()              {}
func (r HumanName) isTransportOutputValue()             {}
func (r Identifier) isTransportOutputValue()            {}
func (r Money) isTransportOutputValue()                 {}
func (r Period) isTransportOutputValue()                {}
func (r Quantity) isTransportOutputValue()              {}
func (r Range) isTransportOutputValue()                 {}
func (r Ratio) isTransportOutputValue()                 {}
func (r RatioRange) isTransportOutputValue()            {}
func (r Reference) isTransportOutputValue()             {}
func (r SampledData) isTransportOutputValue()           {}
func (r Signature) isTransportOutputValue()             {}
func (r Timing) isTransportOutputValue()                {}
func (r ContactDetail) isTransportOutputValue()         {}
func (r DataRequirement) isTransportOutputValue()       {}
func (r Expression) isTransportOutputValue()            {}
func (r ParameterDefinition) isTransportOutputValue()   {}
func (r RelatedArtifact) isTransportOutputValue()       {}
func (r TriggerDefinition) isTransportOutputValue()     {}
func (r UsageContext) isTransportOutputValue()          {}
func (r Availability) isTransportOutputValue()          {}
func (r ExtendedContactDetail) isTransportOutputValue() {}
func (r Dosage) isTransportOutputValue()                {}
func (r Meta) isTransportOutputValue()                  {}
func (r Transport) ResourceType() string {
	return "Transport"
}
func (r Transport) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Transport) MemSize() int {
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
	if r.InstantiatesCanonical != nil {
		s += r.InstantiatesCanonical.MemSize()
	}
	if r.InstantiatesUri != nil {
		s += r.InstantiatesUri.MemSize()
	}
	for _, i := range r.BasedOn {
		s += i.MemSize()
	}
	s += (cap(r.BasedOn) - len(r.BasedOn)) * int(unsafe.Sizeof(Reference{}))
	if r.GroupIdentifier != nil {
		s += r.GroupIdentifier.MemSize()
	}
	for _, i := range r.PartOf {
		s += i.MemSize()
	}
	s += (cap(r.PartOf) - len(r.PartOf)) * int(unsafe.Sizeof(Reference{}))
	if r.Status != nil {
		s += r.Status.MemSize()
	}
	if r.StatusReason != nil {
		s += r.StatusReason.MemSize()
	}
	s += r.Intent.MemSize() - int(unsafe.Sizeof(r.Intent))
	if r.Priority != nil {
		s += r.Priority.MemSize()
	}
	if r.Code != nil {
		s += r.Code.MemSize()
	}
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Focus != nil {
		s += r.Focus.MemSize()
	}
	if r.For != nil {
		s += r.For.MemSize()
	}
	if r.Encounter != nil {
		s += r.Encounter.MemSize()
	}
	if r.CompletionTime != nil {
		s += r.CompletionTime.MemSize()
	}
	if r.AuthoredOn != nil {
		s += r.AuthoredOn.MemSize()
	}
	if r.LastModified != nil {
		s += r.LastModified.MemSize()
	}
	if r.Requester != nil {
		s += r.Requester.MemSize()
	}
	for _, i := range r.PerformerType {
		s += i.MemSize()
	}
	s += (cap(r.PerformerType) - len(r.PerformerType)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Owner != nil {
		s += r.Owner.MemSize()
	}
	if r.Location != nil {
		s += r.Location.MemSize()
	}
	for _, i := range r.Insurance {
		s += i.MemSize()
	}
	s += (cap(r.Insurance) - len(r.Insurance)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	for _, i := range r.RelevantHistory {
		s += i.MemSize()
	}
	s += (cap(r.RelevantHistory) - len(r.RelevantHistory)) * int(unsafe.Sizeof(Reference{}))
	if r.Restriction != nil {
		s += r.Restriction.MemSize()
	}
	for _, i := range r.Input {
		s += i.MemSize()
	}
	s += (cap(r.Input) - len(r.Input)) * int(unsafe.Sizeof(TransportInput{}))
	for _, i := range r.Output {
		s += i.MemSize()
	}
	s += (cap(r.Output) - len(r.Output)) * int(unsafe.Sizeof(TransportOutput{}))
	s += r.RequestedLocation.MemSize() - int(unsafe.Sizeof(r.RequestedLocation))
	s += r.CurrentLocation.MemSize() - int(unsafe.Sizeof(r.CurrentLocation))
	if r.Reason != nil {
		s += r.Reason.MemSize()
	}
	if r.History != nil {
		s += r.History.MemSize()
	}
	return s
}
func (r TransportRestriction) MemSize() int {
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
	if r.Repetitions != nil {
		s += r.Repetitions.MemSize()
	}
	if r.Period != nil {
		s += r.Period.MemSize()
	}
	for _, i := range r.Recipient {
		s += i.MemSize()
	}
	s += (cap(r.Recipient) - len(r.Recipient)) * int(unsafe.Sizeof(Reference{}))
	return s
}
func (r TransportInput) MemSize() int {
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
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Value != nil {
		s += r.Value.MemSize()
	}
	return s
}
func (r TransportOutput) MemSize() int {
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
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Value != nil {
		s += r.Value.MemSize()
	}
	return s
}
func (r Transport) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Transport) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Transport) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Transport\""))
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
	if r.InstantiatesCanonical != nil && r.InstantiatesCanonical.Value != nil {
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
	if r.InstantiatesCanonical != nil && (r.InstantiatesCanonical.Id != nil || r.InstantiatesCanonical.Extension != nil) {
		p := primitiveElement{Id: r.InstantiatesCanonical.Id, Extension: r.InstantiatesCanonical.Extension}
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
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.InstantiatesUri != nil && r.InstantiatesUri.Value != nil {
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
	if r.InstantiatesUri != nil && (r.InstantiatesUri.Id != nil || r.InstantiatesUri.Extension != nil) {
		p := primitiveElement{Id: r.InstantiatesUri.Id, Extension: r.InstantiatesUri.Extension}
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
		err = p.marshalJSON(w)
		if err != nil {
			return err
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
	if len(r.PartOf) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"partOf\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PartOf {
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
	if r.Status != nil && r.Status.Value != nil {
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
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
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
	if r.StatusReason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"statusReason\":"))
		if err != nil {
			return err
		}
		err = r.StatusReason.marshalJSON(w)
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
	if r.Focus != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"focus\":"))
		if err != nil {
			return err
		}
		err = r.Focus.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.For != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"for\":"))
		if err != nil {
			return err
		}
		err = r.For.marshalJSON(w)
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
	if r.CompletionTime != nil && r.CompletionTime.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"completionTime\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CompletionTime)
		if err != nil {
			return err
		}
	}
	if r.CompletionTime != nil && (r.CompletionTime.Id != nil || r.CompletionTime.Extension != nil) {
		p := primitiveElement{Id: r.CompletionTime.Id, Extension: r.CompletionTime.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_completionTime\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	if r.LastModified != nil && r.LastModified.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastModified\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastModified)
		if err != nil {
			return err
		}
	}
	if r.LastModified != nil && (r.LastModified.Id != nil || r.LastModified.Extension != nil) {
		p := primitiveElement{Id: r.LastModified.Id, Extension: r.LastModified.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastModified\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Requester != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"requester\":"))
		if err != nil {
			return err
		}
		err = r.Requester.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.PerformerType) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"performerType\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PerformerType {
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
	if r.Owner != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"owner\":"))
		if err != nil {
			return err
		}
		err = r.Owner.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Location != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"location\":"))
		if err != nil {
			return err
		}
		err = r.Location.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Insurance) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"insurance\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Insurance {
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
	if len(r.RelevantHistory) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relevantHistory\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RelevantHistory {
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
	if r.Restriction != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"restriction\":"))
		if err != nil {
			return err
		}
		err = r.Restriction.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Input) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"input\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Input {
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
	if len(r.Output) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"output\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Output {
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
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"requestedLocation\":"))
	if err != nil {
		return err
	}
	err = r.RequestedLocation.marshalJSON(w)
	if err != nil {
		return err
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"currentLocation\":"))
	if err != nil {
		return err
	}
	err = r.CurrentLocation.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Reason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reason\":"))
		if err != nil {
			return err
		}
		err = r.Reason.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.History != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"history\":"))
		if err != nil {
			return err
		}
		err = r.History.marshalJSON(w)
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
func (r TransportRestriction) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r TransportRestriction) marshalJSON(w io.Writer) error {
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
	if r.Repetitions != nil && r.Repetitions.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"repetitions\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Repetitions)
		if err != nil {
			return err
		}
	}
	if r.Repetitions != nil && (r.Repetitions.Id != nil || r.Repetitions.Extension != nil) {
		p := primitiveElement{Id: r.Repetitions.Id, Extension: r.Repetitions.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_repetitions\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Period != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"period\":"))
		if err != nil {
			return err
		}
		err = r.Period.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Recipient) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"recipient\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Recipient {
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
func (r TransportInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r TransportInput) marshalJSON(w io.Writer) error {
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
	switch v := r.Value.(type) {
	case Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
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
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAge\":"))
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
		_, err = w.Write([]byte("\"valueAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
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
		_, err = w.Write([]byte("\"valueDuration\":"))
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
		_, err = w.Write([]byte("\"valueDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
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
		_, err = w.Write([]byte("\"valuePeriod\":"))
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
		_, err = w.Write([]byte("\"valuePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
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
		_, err = w.Write([]byte("\"valueTiming\":"))
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
		_, err = w.Write([]byte("\"valueTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
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
func (r TransportOutput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r TransportOutput) marshalJSON(w io.Writer) error {
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
	switch v := r.Value.(type) {
	case Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
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
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAge\":"))
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
		_, err = w.Write([]byte("\"valueAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
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
		_, err = w.Write([]byte("\"valueDuration\":"))
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
		_, err = w.Write([]byte("\"valueDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
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
		_, err = w.Write([]byte("\"valuePeriod\":"))
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
		_, err = w.Write([]byte("\"valuePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
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
		_, err = w.Write([]byte("\"valueTiming\":"))
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
		_, err = w.Write([]byte("\"valueTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
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
func (r *Transport) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *Transport) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Transport element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "instantiatesCanonical":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.InstantiatesCanonical == nil {
				r.InstantiatesCanonical = &Canonical{}
			}
			r.InstantiatesCanonical.Value = v.Value
		case "_instantiatesCanonical":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.InstantiatesCanonical == nil {
				r.InstantiatesCanonical = &Canonical{}
			}
			r.InstantiatesCanonical.Id = v.Id
			r.InstantiatesCanonical.Extension = v.Extension
		case "instantiatesUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.InstantiatesUri == nil {
				r.InstantiatesUri = &Uri{}
			}
			r.InstantiatesUri.Value = v.Value
		case "_instantiatesUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.InstantiatesUri == nil {
				r.InstantiatesUri = &Uri{}
			}
			r.InstantiatesUri.Id = v.Id
			r.InstantiatesUri.Extension = v.Extension
		case "basedOn":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "groupIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.GroupIdentifier = &v
		case "partOf":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PartOf = append(r.PartOf, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Status == nil {
				r.Status = &Code{}
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Status == nil {
				r.Status = &Code{}
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "statusReason":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.StatusReason = &v
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
		case "focus":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Focus = &v
		case "for":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.For = &v
		case "encounter":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Encounter = &v
		case "completionTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CompletionTime == nil {
				r.CompletionTime = &DateTime{}
			}
			r.CompletionTime.Value = v.Value
		case "_completionTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CompletionTime == nil {
				r.CompletionTime = &DateTime{}
			}
			r.CompletionTime.Id = v.Id
			r.CompletionTime.Extension = v.Extension
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
		case "lastModified":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastModified == nil {
				r.LastModified = &DateTime{}
			}
			r.LastModified.Value = v.Value
		case "_lastModified":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastModified == nil {
				r.LastModified = &DateTime{}
			}
			r.LastModified.Id = v.Id
			r.LastModified.Extension = v.Extension
		case "requester":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Requester = &v
		case "performerType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PerformerType = append(r.PerformerType, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "owner":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Owner = &v
		case "location":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = &v
		case "insurance":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "relevantHistory":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RelevantHistory = append(r.RelevantHistory, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "restriction":
			var v TransportRestriction
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Restriction = &v
		case "input":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v TransportInput
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Input = append(r.Input, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "output":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Transport element", t)
			}
			for d.More() {
				var v TransportOutput
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Output = append(r.Output, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Transport element", t)
			}
		case "requestedLocation":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.RequestedLocation = v
		case "currentLocation":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CurrentLocation = v
		case "reason":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Reason = &v
		case "history":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.History = &v
		default:
			return fmt.Errorf("invalid field: %s in Transport", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Transport element", t)
	}
	return nil
}
func (r *TransportRestriction) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in TransportRestriction element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in TransportRestriction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportRestriction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportRestriction element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportRestriction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportRestriction element", t)
			}
		case "repetitions":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Repetitions == nil {
				r.Repetitions = &PositiveInt{}
			}
			r.Repetitions.Value = v.Value
		case "_repetitions":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Repetitions == nil {
				r.Repetitions = &PositiveInt{}
			}
			r.Repetitions.Id = v.Id
			r.Repetitions.Extension = v.Extension
		case "period":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Period = &v
		case "recipient":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportRestriction element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Recipient = append(r.Recipient, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportRestriction element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in TransportRestriction", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in TransportRestriction element", t)
	}
	return nil
}
func (r *TransportInput) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in TransportInput element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in TransportInput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportInput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportInput element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportInput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportInput element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "valueBase64Binary":
			var v Base64Binary
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: r.Value.(Base64Binary).Extension,
					Id:        r.Value.(Base64Binary).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBase64Binary":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Base64Binary).Value,
				}
			} else {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueBoolean":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: r.Value.(Boolean).Extension,
					Id:        r.Value.(Boolean).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBoolean":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Boolean).Value,
				}
			} else {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCanonical":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: r.Value.(Canonical).Extension,
					Id:        r.Value.(Canonical).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCanonical":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Canonical).Value,
				}
			} else {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCode":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: r.Value.(Code).Extension,
					Id:        r.Value.(Code).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCode":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Code).Value,
				}
			} else {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: r.Value.(Date).Extension,
					Id:        r.Value.(Date).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Date).Value,
				}
			} else {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: r.Value.(DateTime).Extension,
					Id:        r.Value.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(DateTime).Value,
				}
			} else {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDecimal":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: r.Value.(Decimal).Extension,
					Id:        r.Value.(Decimal).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDecimal":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Decimal).Value,
				}
			} else {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueId":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: r.Value.(Id).Extension,
					Id:        r.Value.(Id).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Id).Value,
				}
			} else {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInstant":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: r.Value.(Instant).Extension,
					Id:        r.Value.(Instant).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInstant":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Instant).Value,
				}
			} else {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: r.Value.(Integer).Extension,
					Id:        r.Value.(Integer).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer).Value,
				}
			} else {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger64":
			var v Integer64
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: r.Value.(Integer64).Extension,
					Id:        r.Value.(Integer64).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger64":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer64).Value,
				}
			} else {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueMarkdown":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: r.Value.(Markdown).Extension,
					Id:        r.Value.(Markdown).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueMarkdown":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Markdown).Value,
				}
			} else {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueOid":
			var v Oid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: r.Value.(Oid).Extension,
					Id:        r.Value.(Oid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueOid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Oid).Value,
				}
			} else {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valuePositiveInt":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: r.Value.(PositiveInt).Extension,
					Id:        r.Value.(PositiveInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valuePositiveInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(PositiveInt).Value,
				}
			} else {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: r.Value.(String).Extension,
					Id:        r.Value.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(String).Value,
				}
			} else {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueTime":
			var v Time
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: r.Value.(Time).Extension,
					Id:        r.Value.(Time).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Time).Value,
				}
			} else {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUnsignedInt":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: r.Value.(UnsignedInt).Extension,
					Id:        r.Value.(UnsignedInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUnsignedInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(UnsignedInt).Value,
				}
			} else {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: r.Value.(Uri).Extension,
					Id:        r.Value.(Uri).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uri).Value,
				}
			} else {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUrl":
			var v Url
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: r.Value.(Url).Extension,
					Id:        r.Value.(Url).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUrl":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Url).Value,
				}
			} else {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUuid":
			var v Uuid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: r.Value.(Uuid).Extension,
					Id:        r.Value.(Uuid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUuid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uuid).Value,
				}
			} else {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueAddress":
			var v Address
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAge":
			var v Age
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAnnotation":
			var v Annotation
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableReference":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCoding":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactPoint":
			var v ContactPoint
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCount":
			var v Count
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDistance":
			var v Distance
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDuration":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueHumanName":
			var v HumanName
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMoney":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valuePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatio":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatioRange":
			var v RatioRange
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSampledData":
			var v SampledData
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTiming":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactDetail":
			var v ContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDataRequirement":
			var v DataRequirement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExpression":
			var v Expression
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueParameterDefinition":
			var v ParameterDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRelatedArtifact":
			var v RelatedArtifact
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTriggerDefinition":
			var v TriggerDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueUsageContext":
			var v UsageContext
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAvailability":
			var v Availability
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExtendedContactDetail":
			var v ExtendedContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDosage":
			var v Dosage
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMeta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		default:
			return fmt.Errorf("invalid field: %s in TransportInput", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in TransportInput element", t)
	}
	return nil
}
func (r *TransportOutput) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in TransportOutput element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in TransportOutput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportOutput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportOutput element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in TransportOutput element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in TransportOutput element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "valueBase64Binary":
			var v Base64Binary
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: r.Value.(Base64Binary).Extension,
					Id:        r.Value.(Base64Binary).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBase64Binary":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Base64Binary).Value,
				}
			} else {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueBoolean":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: r.Value.(Boolean).Extension,
					Id:        r.Value.(Boolean).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBoolean":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Boolean).Value,
				}
			} else {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCanonical":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: r.Value.(Canonical).Extension,
					Id:        r.Value.(Canonical).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCanonical":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Canonical).Value,
				}
			} else {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCode":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: r.Value.(Code).Extension,
					Id:        r.Value.(Code).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCode":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Code).Value,
				}
			} else {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: r.Value.(Date).Extension,
					Id:        r.Value.(Date).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Date).Value,
				}
			} else {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: r.Value.(DateTime).Extension,
					Id:        r.Value.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(DateTime).Value,
				}
			} else {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDecimal":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: r.Value.(Decimal).Extension,
					Id:        r.Value.(Decimal).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDecimal":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Decimal).Value,
				}
			} else {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueId":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: r.Value.(Id).Extension,
					Id:        r.Value.(Id).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Id).Value,
				}
			} else {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInstant":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: r.Value.(Instant).Extension,
					Id:        r.Value.(Instant).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInstant":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Instant).Value,
				}
			} else {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: r.Value.(Integer).Extension,
					Id:        r.Value.(Integer).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer).Value,
				}
			} else {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger64":
			var v Integer64
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: r.Value.(Integer64).Extension,
					Id:        r.Value.(Integer64).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger64":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer64).Value,
				}
			} else {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueMarkdown":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: r.Value.(Markdown).Extension,
					Id:        r.Value.(Markdown).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueMarkdown":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Markdown).Value,
				}
			} else {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueOid":
			var v Oid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: r.Value.(Oid).Extension,
					Id:        r.Value.(Oid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueOid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Oid).Value,
				}
			} else {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valuePositiveInt":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: r.Value.(PositiveInt).Extension,
					Id:        r.Value.(PositiveInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valuePositiveInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(PositiveInt).Value,
				}
			} else {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: r.Value.(String).Extension,
					Id:        r.Value.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(String).Value,
				}
			} else {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueTime":
			var v Time
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: r.Value.(Time).Extension,
					Id:        r.Value.(Time).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Time).Value,
				}
			} else {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUnsignedInt":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: r.Value.(UnsignedInt).Extension,
					Id:        r.Value.(UnsignedInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUnsignedInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(UnsignedInt).Value,
				}
			} else {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: r.Value.(Uri).Extension,
					Id:        r.Value.(Uri).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uri).Value,
				}
			} else {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUrl":
			var v Url
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: r.Value.(Url).Extension,
					Id:        r.Value.(Url).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUrl":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Url).Value,
				}
			} else {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUuid":
			var v Uuid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: r.Value.(Uuid).Extension,
					Id:        r.Value.(Uuid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUuid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uuid).Value,
				}
			} else {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueAddress":
			var v Address
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAge":
			var v Age
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAnnotation":
			var v Annotation
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableReference":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCoding":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactPoint":
			var v ContactPoint
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCount":
			var v Count
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDistance":
			var v Distance
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDuration":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueHumanName":
			var v HumanName
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMoney":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valuePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatio":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatioRange":
			var v RatioRange
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSampledData":
			var v SampledData
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTiming":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactDetail":
			var v ContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDataRequirement":
			var v DataRequirement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExpression":
			var v Expression
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueParameterDefinition":
			var v ParameterDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRelatedArtifact":
			var v RelatedArtifact
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTriggerDefinition":
			var v TriggerDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueUsageContext":
			var v UsageContext
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAvailability":
			var v Availability
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExtendedContactDetail":
			var v ExtendedContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDosage":
			var v Dosage
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMeta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		default:
			return fmt.Errorf("invalid field: %s in TransportOutput", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in TransportOutput element", t)
	}
	return nil
}
func (r Transport) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Transport"
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
	err = e.EncodeElement(r.GroupIdentifier, xml.StartElement{Name: xml.Name{Local: "groupIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartOf, xml.StartElement{Name: xml.Name{Local: "partOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusReason, xml.StartElement{Name: xml.Name{Local: "statusReason"}})
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Focus, xml.StartElement{Name: xml.Name{Local: "focus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.For, xml.StartElement{Name: xml.Name{Local: "for"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CompletionTime, xml.StartElement{Name: xml.Name{Local: "completionTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthoredOn, xml.StartElement{Name: xml.Name{Local: "authoredOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastModified, xml.StartElement{Name: xml.Name{Local: "lastModified"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Requester, xml.StartElement{Name: xml.Name{Local: "requester"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PerformerType, xml.StartElement{Name: xml.Name{Local: "performerType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Owner, xml.StartElement{Name: xml.Name{Local: "owner"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurance, xml.StartElement{Name: xml.Name{Local: "insurance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelevantHistory, xml.StartElement{Name: xml.Name{Local: "relevantHistory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Restriction, xml.StartElement{Name: xml.Name{Local: "restriction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Input, xml.StartElement{Name: xml.Name{Local: "input"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Output, xml.StartElement{Name: xml.Name{Local: "output"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequestedLocation, xml.StartElement{Name: xml.Name{Local: "requestedLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CurrentLocation, xml.StartElement{Name: xml.Name{Local: "currentLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.History, xml.StartElement{Name: xml.Name{Local: "history"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r TransportRestriction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Repetitions, xml.StartElement{Name: xml.Name{Local: "repetitions"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recipient, xml.StartElement{Name: xml.Name{Local: "recipient"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r TransportInput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	case Base64Binary, *Base64Binary:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBase64Binary"}})
		if err != nil {
			return err
		}
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Canonical, *Canonical:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCanonical"}})
		if err != nil {
			return err
		}
	case Code, *Code:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCode"}})
		if err != nil {
			return err
		}
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDate"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDateTime"}})
		if err != nil {
			return err
		}
	case Decimal, *Decimal:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDecimal"}})
		if err != nil {
			return err
		}
	case Id, *Id:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueId"}})
		if err != nil {
			return err
		}
	case Instant, *Instant:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInstant"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger"}})
		if err != nil {
			return err
		}
	case Integer64, *Integer64:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger64"}})
		if err != nil {
			return err
		}
	case Markdown, *Markdown:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMarkdown"}})
		if err != nil {
			return err
		}
	case Oid, *Oid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueOid"}})
		if err != nil {
			return err
		}
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Time, *Time:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTime"}})
		if err != nil {
			return err
		}
	case UnsignedInt, *UnsignedInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUnsignedInt"}})
		if err != nil {
			return err
		}
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUri"}})
		if err != nil {
			return err
		}
	case Url, *Url:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUrl"}})
		if err != nil {
			return err
		}
	case Uuid, *Uuid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUuid"}})
		if err != nil {
			return err
		}
	case Address, *Address:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAddress"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAge"}})
		if err != nil {
			return err
		}
	case Annotation, *Annotation:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAnnotation"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableConcept"}})
		if err != nil {
			return err
		}
	case CodeableReference, *CodeableReference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableReference"}})
		if err != nil {
			return err
		}
	case Coding, *Coding:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCoding"}})
		if err != nil {
			return err
		}
	case ContactPoint, *ContactPoint:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactPoint"}})
		if err != nil {
			return err
		}
	case Count, *Count:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCount"}})
		if err != nil {
			return err
		}
	case Distance, *Distance:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDistance"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDuration"}})
		if err != nil {
			return err
		}
	case HumanName, *HumanName:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueHumanName"}})
		if err != nil {
			return err
		}
	case Identifier, *Identifier:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueIdentifier"}})
		if err != nil {
			return err
		}
	case Money, *Money:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMoney"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePeriod"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRange"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatio"}})
		if err != nil {
			return err
		}
	case RatioRange, *RatioRange:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatioRange"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
		if err != nil {
			return err
		}
	case SampledData, *SampledData:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSampledData"}})
		if err != nil {
			return err
		}
	case Signature, *Signature:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSignature"}})
		if err != nil {
			return err
		}
	case Timing, *Timing:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTiming"}})
		if err != nil {
			return err
		}
	case ContactDetail, *ContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactDetail"}})
		if err != nil {
			return err
		}
	case DataRequirement, *DataRequirement:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDataRequirement"}})
		if err != nil {
			return err
		}
	case Expression, *Expression:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExpression"}})
		if err != nil {
			return err
		}
	case ParameterDefinition, *ParameterDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueParameterDefinition"}})
		if err != nil {
			return err
		}
	case RelatedArtifact, *RelatedArtifact:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRelatedArtifact"}})
		if err != nil {
			return err
		}
	case TriggerDefinition, *TriggerDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTriggerDefinition"}})
		if err != nil {
			return err
		}
	case UsageContext, *UsageContext:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUsageContext"}})
		if err != nil {
			return err
		}
	case Availability, *Availability:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAvailability"}})
		if err != nil {
			return err
		}
	case ExtendedContactDetail, *ExtendedContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExtendedContactDetail"}})
		if err != nil {
			return err
		}
	case Dosage, *Dosage:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDosage"}})
		if err != nil {
			return err
		}
	case Meta, *Meta:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMeta"}})
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
func (r TransportOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	case Base64Binary, *Base64Binary:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBase64Binary"}})
		if err != nil {
			return err
		}
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Canonical, *Canonical:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCanonical"}})
		if err != nil {
			return err
		}
	case Code, *Code:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCode"}})
		if err != nil {
			return err
		}
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDate"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDateTime"}})
		if err != nil {
			return err
		}
	case Decimal, *Decimal:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDecimal"}})
		if err != nil {
			return err
		}
	case Id, *Id:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueId"}})
		if err != nil {
			return err
		}
	case Instant, *Instant:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInstant"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger"}})
		if err != nil {
			return err
		}
	case Integer64, *Integer64:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger64"}})
		if err != nil {
			return err
		}
	case Markdown, *Markdown:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMarkdown"}})
		if err != nil {
			return err
		}
	case Oid, *Oid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueOid"}})
		if err != nil {
			return err
		}
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Time, *Time:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTime"}})
		if err != nil {
			return err
		}
	case UnsignedInt, *UnsignedInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUnsignedInt"}})
		if err != nil {
			return err
		}
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUri"}})
		if err != nil {
			return err
		}
	case Url, *Url:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUrl"}})
		if err != nil {
			return err
		}
	case Uuid, *Uuid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUuid"}})
		if err != nil {
			return err
		}
	case Address, *Address:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAddress"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAge"}})
		if err != nil {
			return err
		}
	case Annotation, *Annotation:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAnnotation"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableConcept"}})
		if err != nil {
			return err
		}
	case CodeableReference, *CodeableReference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableReference"}})
		if err != nil {
			return err
		}
	case Coding, *Coding:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCoding"}})
		if err != nil {
			return err
		}
	case ContactPoint, *ContactPoint:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactPoint"}})
		if err != nil {
			return err
		}
	case Count, *Count:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCount"}})
		if err != nil {
			return err
		}
	case Distance, *Distance:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDistance"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDuration"}})
		if err != nil {
			return err
		}
	case HumanName, *HumanName:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueHumanName"}})
		if err != nil {
			return err
		}
	case Identifier, *Identifier:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueIdentifier"}})
		if err != nil {
			return err
		}
	case Money, *Money:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMoney"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePeriod"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRange"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatio"}})
		if err != nil {
			return err
		}
	case RatioRange, *RatioRange:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatioRange"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
		if err != nil {
			return err
		}
	case SampledData, *SampledData:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSampledData"}})
		if err != nil {
			return err
		}
	case Signature, *Signature:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSignature"}})
		if err != nil {
			return err
		}
	case Timing, *Timing:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTiming"}})
		if err != nil {
			return err
		}
	case ContactDetail, *ContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactDetail"}})
		if err != nil {
			return err
		}
	case DataRequirement, *DataRequirement:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDataRequirement"}})
		if err != nil {
			return err
		}
	case Expression, *Expression:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExpression"}})
		if err != nil {
			return err
		}
	case ParameterDefinition, *ParameterDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueParameterDefinition"}})
		if err != nil {
			return err
		}
	case RelatedArtifact, *RelatedArtifact:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRelatedArtifact"}})
		if err != nil {
			return err
		}
	case TriggerDefinition, *TriggerDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTriggerDefinition"}})
		if err != nil {
			return err
		}
	case UsageContext, *UsageContext:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUsageContext"}})
		if err != nil {
			return err
		}
	case Availability, *Availability:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAvailability"}})
		if err != nil {
			return err
		}
	case ExtendedContactDetail, *ExtendedContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExtendedContactDetail"}})
		if err != nil {
			return err
		}
	case Dosage, *Dosage:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDosage"}})
		if err != nil {
			return err
		}
	case Meta, *Meta:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMeta"}})
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
func (r *Transport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.InstantiatesCanonical = &v
			case "instantiatesUri":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesUri = &v
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "groupIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GroupIdentifier = &v
			case "partOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartOf = append(r.PartOf, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "statusReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusReason = &v
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "focus":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Focus = &v
			case "for":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.For = &v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "completionTime":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CompletionTime = &v
			case "authoredOn":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthoredOn = &v
			case "lastModified":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastModified = &v
			case "requester":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requester = &v
			case "performerType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PerformerType = append(r.PerformerType, v)
			case "owner":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Owner = &v
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "insurance":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "relevantHistory":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelevantHistory = append(r.RelevantHistory, v)
			case "restriction":
				var v TransportRestriction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Restriction = &v
			case "input":
				var v TransportInput
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Input = append(r.Input, v)
			case "output":
				var v TransportOutput
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Output = append(r.Output, v)
			case "requestedLocation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequestedLocation = v
			case "currentLocation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CurrentLocation = v
			case "reason":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = &v
			case "history":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.History = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *TransportRestriction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "repetitions":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Repetitions = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			case "recipient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recipient = append(r.Recipient, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *TransportInput) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Type = v
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
			case "valueBoolean":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCanonical":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCode":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDate":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDateTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDecimal":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueId":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInstant":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger64":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer64
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMarkdown":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueOid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Oid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePositiveInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
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
			case "valueTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUnsignedInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUri":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUrl":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUuid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uuid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAddress":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAge":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAnnotation":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAttachment":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableConcept":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCoding":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactPoint":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCount":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Count
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDistance":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Distance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDuration":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueHumanName":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueIdentifier":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMoney":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePeriod":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueQuantity":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatio":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatioRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RatioRange
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSampledData":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v SampledData
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSignature":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTiming":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDataRequirement":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DataRequirement
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExpression":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Expression
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueParameterDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ParameterDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRelatedArtifact":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RelatedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTriggerDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v TriggerDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUsageContext":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAvailability":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Availability
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExtendedContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ExtendedContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDosage":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Dosage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMeta":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Meta
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
func (r *TransportOutput) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Type = v
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
			case "valueBoolean":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCanonical":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCode":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDate":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDateTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDecimal":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueId":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInstant":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger64":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer64
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMarkdown":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueOid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Oid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePositiveInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
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
			case "valueTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUnsignedInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUri":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUrl":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUuid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uuid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAddress":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAge":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAnnotation":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAttachment":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableConcept":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCoding":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactPoint":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCount":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Count
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDistance":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Distance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDuration":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueHumanName":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueIdentifier":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMoney":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePeriod":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueQuantity":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatio":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatioRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RatioRange
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSampledData":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v SampledData
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSignature":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTiming":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDataRequirement":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DataRequirement
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExpression":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Expression
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueParameterDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ParameterDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRelatedArtifact":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RelatedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTriggerDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v TriggerDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUsageContext":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAvailability":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Availability
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExtendedContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ExtendedContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDosage":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Dosage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMeta":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Meta
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
