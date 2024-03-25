package r4

import (
	"encoding/json"
	"fmt"
)

// Todo.
type SubstanceReferenceInformation struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Comment *String
	// Todo.
	Gene []SubstanceReferenceInformationGene
	// Todo.
	GeneElement []SubstanceReferenceInformationGeneElement
	// Todo.
	Classification []SubstanceReferenceInformationClassification
	// Todo.
	Target []SubstanceReferenceInformationTarget
}
type jsonSubstanceReferenceInformation struct {
	ResourceType                  string                                        `json:"resourceType"`
	Id                            *Id                                           `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                             `json:"_id,omitempty"`
	Meta                          *Meta                                         `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                          `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                             `json:"_implicitRules,omitempty"`
	Language                      *Code                                         `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                             `json:"_language,omitempty"`
	Text                          *Narrative                                    `json:"text,omitempty"`
	Contained                     []containedResource                           `json:"contained,omitempty"`
	Extension                     []Extension                                   `json:"extension,omitempty"`
	ModifierExtension             []Extension                                   `json:"modifierExtension,omitempty"`
	Comment                       *String                                       `json:"comment,omitempty"`
	CommentPrimitiveElement       *primitiveElement                             `json:"_comment,omitempty"`
	Gene                          []SubstanceReferenceInformationGene           `json:"gene,omitempty"`
	GeneElement                   []SubstanceReferenceInformationGeneElement    `json:"geneElement,omitempty"`
	Classification                []SubstanceReferenceInformationClassification `json:"classification,omitempty"`
	Target                        []SubstanceReferenceInformationTarget         `json:"target,omitempty"`
}

func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceReferenceInformation) marshalJSON() jsonSubstanceReferenceInformation {
	m := jsonSubstanceReferenceInformation{}
	m.ResourceType = "SubstanceReferenceInformation"
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
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.Gene = r.Gene
	m.GeneElement = r.GeneElement
	m.Classification = r.Classification
	m.Target = r.Target
	return m
}
func (r *SubstanceReferenceInformation) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceReferenceInformation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceReferenceInformation) unmarshalJSON(m jsonSubstanceReferenceInformation) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Gene = m.Gene
	r.GeneElement = m.GeneElement
	r.Classification = m.Classification
	r.Target = m.Target
	return nil
}
func (r SubstanceReferenceInformation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstanceReferenceInformationGene struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	GeneSequenceOrigin *CodeableConcept
	// Todo.
	Gene *CodeableConcept
	// Todo.
	Source []Reference
}
type jsonSubstanceReferenceInformationGene struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `json:"gene,omitempty"`
	Source             []Reference      `json:"source,omitempty"`
}

func (r SubstanceReferenceInformationGene) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceReferenceInformationGene) marshalJSON() jsonSubstanceReferenceInformationGene {
	m := jsonSubstanceReferenceInformationGene{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.GeneSequenceOrigin = r.GeneSequenceOrigin
	m.Gene = r.Gene
	m.Source = r.Source
	return m
}
func (r *SubstanceReferenceInformationGene) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceReferenceInformationGene
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceReferenceInformationGene) unmarshalJSON(m jsonSubstanceReferenceInformationGene) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.GeneSequenceOrigin = m.GeneSequenceOrigin
	r.Gene = m.Gene
	r.Source = m.Source
	return nil
}
func (r SubstanceReferenceInformationGene) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstanceReferenceInformationGeneElement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Type *CodeableConcept
	// Todo.
	Element *Identifier
	// Todo.
	Source []Reference
}
type jsonSubstanceReferenceInformationGeneElement struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Element           *Identifier      `json:"element,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}

func (r SubstanceReferenceInformationGeneElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceReferenceInformationGeneElement) marshalJSON() jsonSubstanceReferenceInformationGeneElement {
	m := jsonSubstanceReferenceInformationGeneElement{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Element = r.Element
	m.Source = r.Source
	return m
}
func (r *SubstanceReferenceInformationGeneElement) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceReferenceInformationGeneElement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceReferenceInformationGeneElement) unmarshalJSON(m jsonSubstanceReferenceInformationGeneElement) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Element = m.Element
	r.Source = m.Source
	return nil
}
func (r SubstanceReferenceInformationGeneElement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstanceReferenceInformationClassification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Domain *CodeableConcept
	// Todo.
	Classification *CodeableConcept
	// Todo.
	Subtype []CodeableConcept
	// Todo.
	Source []Reference
}
type jsonSubstanceReferenceInformationClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Domain            *CodeableConcept  `json:"domain,omitempty"`
	Classification    *CodeableConcept  `json:"classification,omitempty"`
	Subtype           []CodeableConcept `json:"subtype,omitempty"`
	Source            []Reference       `json:"source,omitempty"`
}

func (r SubstanceReferenceInformationClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceReferenceInformationClassification) marshalJSON() jsonSubstanceReferenceInformationClassification {
	m := jsonSubstanceReferenceInformationClassification{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Domain = r.Domain
	m.Classification = r.Classification
	m.Subtype = r.Subtype
	m.Source = r.Source
	return m
}
func (r *SubstanceReferenceInformationClassification) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceReferenceInformationClassification
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceReferenceInformationClassification) unmarshalJSON(m jsonSubstanceReferenceInformationClassification) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Domain = m.Domain
	r.Classification = m.Classification
	r.Subtype = m.Subtype
	r.Source = m.Source
	return nil
}
func (r SubstanceReferenceInformationClassification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstanceReferenceInformationTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Target *Identifier
	// Todo.
	Type *CodeableConcept
	// Todo.
	Interaction *CodeableConcept
	// Todo.
	Organism *CodeableConcept
	// Todo.
	OrganismType *CodeableConcept
	// Todo.
	Amount isSubstanceReferenceInformationTargetAmount
	// Todo.
	AmountType *CodeableConcept
	// Todo.
	Source []Reference
}
type isSubstanceReferenceInformationTargetAmount interface {
	isSubstanceReferenceInformationTargetAmount()
}

func (r Quantity) isSubstanceReferenceInformationTargetAmount() {}
func (r Range) isSubstanceReferenceInformationTargetAmount()    {}
func (r String) isSubstanceReferenceInformationTargetAmount()   {}

type jsonSubstanceReferenceInformationTarget struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Target                       *Identifier       `json:"target,omitempty"`
	Type                         *CodeableConcept  `json:"type,omitempty"`
	Interaction                  *CodeableConcept  `json:"interaction,omitempty"`
	Organism                     *CodeableConcept  `json:"organism,omitempty"`
	OrganismType                 *CodeableConcept  `json:"organismType,omitempty"`
	AmountQuantity               *Quantity         `json:"amountQuantity,omitempty"`
	AmountRange                  *Range            `json:"amountRange,omitempty"`
	AmountString                 *String           `json:"amountString,omitempty"`
	AmountStringPrimitiveElement *primitiveElement `json:"_amountString,omitempty"`
	AmountType                   *CodeableConcept  `json:"amountType,omitempty"`
	Source                       []Reference       `json:"source,omitempty"`
}

func (r SubstanceReferenceInformationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceReferenceInformationTarget) marshalJSON() jsonSubstanceReferenceInformationTarget {
	m := jsonSubstanceReferenceInformationTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Target = r.Target
	m.Type = r.Type
	m.Interaction = r.Interaction
	m.Organism = r.Organism
	m.OrganismType = r.OrganismType
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case Range:
		m.AmountRange = &v
	case *Range:
		m.AmountRange = v
	case String:
		m.AmountString = &v
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AmountString = v
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.AmountType = r.AmountType
	m.Source = r.Source
	return m
}
func (r *SubstanceReferenceInformationTarget) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceReferenceInformationTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceReferenceInformationTarget) unmarshalJSON(m jsonSubstanceReferenceInformationTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Target = m.Target
	r.Type = m.Type
	r.Interaction = m.Interaction
	r.Organism = m.Organism
	r.OrganismType = m.OrganismType
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountRange != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountRange
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	r.AmountType = m.AmountType
	r.Source = m.Source
	return nil
}
func (r SubstanceReferenceInformationTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
