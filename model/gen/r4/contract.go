package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type Contract struct {
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
	// Unique identifier for this Contract or a derivative that references a Source Contract.
	Identifier []Identifier
	// Canonical identifier for this contract, represented as a URI (globally unique).
	Url *Uri
	// An edition identifier used for business purposes to label business significant variants.
	Version *String
	// The status of the resource instance.
	Status *Code
	// Legal states of the formation of a legal instrument, which is a formally executed written document that can be formally attributed to its author, records and formally expresses a legally enforceable act, process, or contractual duty, obligation, or right, and therefore evidences that act, process, or agreement.
	LegalState *CodeableConcept
	// The URL pointing to a FHIR-defined Contract Definition that is adhered to in whole or part by this Contract.
	InstantiatesCanonical *Reference
	// The URL pointing to an externally maintained definition that is adhered to in whole or in part by this Contract.
	InstantiatesUri *Uri
	// The minimal content derived from the basal information source at a specific stage in its lifecycle.
	ContentDerivative *CodeableConcept
	// When this  Contract was issued.
	Issued *DateTime
	// Relevant time or time-period when this Contract is applicable.
	Applies *Period
	// Event resulting in discontinuation or termination of this Contract instance by one or more parties to the contract.
	ExpirationType *CodeableConcept
	// The target entity impacted by or of interest to parties to the agreement.
	Subject []Reference
	// A formally or informally recognized grouping of people, principals, organizations, or jurisdictions formed for the purpose of achieving some form of collective action such as the promulgation, administration and enforcement of contracts and policies.
	Authority []Reference
	// Recognized governance framework or system operating with a circumscribed scope in accordance with specified principles, policies, processes or procedures for managing rights, actions, or behaviors of parties or principals relative to resources.
	Domain []Reference
	// Sites in which the contract is complied with,  exercised, or in force.
	Site []Reference
	// A natural language name identifying this Contract definition, derivative, or instance in any legal state. Provides additional information about its content. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for this Contract definition, derivative, or instance in any legal state.t giving additional information about its content.
	Title *String
	// An explanatory or alternate user-friendly title for this Contract definition, derivative, or instance in any legal state.t giving additional information about its content.
	Subtitle *String
	// Alternative representation of the title for this Contract definition, derivative, or instance in any legal state., e.g., a domain specific contract number related to legislation.
	Alias []String
	// The individual or organization that authored the Contract definition, derivative, or instance in any legal state.
	Author *Reference
	// A selector of legal concerns for this Contract definition, derivative, or instance in any legal state.
	Scope *CodeableConcept
	// Narrows the range of legal concerns to focus on the achievement of specific contractual objectives.
	Topic isContractTopic
	// A high-level category for the legal instrument, whether constructed as a Contract definition, derivative, or instance in any legal state.  Provides additional information about its content within the context of the Contract's scope to distinguish the kinds of systems that would be interested in the contract.
	Type *CodeableConcept
	// Sub-category for the Contract that distinguishes the kinds of systems that would be interested in the Contract within the context of the Contract's scope.
	SubType []CodeableConcept
	// Precusory content developed with a focus and intent of supporting the formation a Contract instance, which may be associated with and transformable into a Contract.
	ContentDefinition *ContractContentDefinition
	// One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
	Term []ContractTerm
	// Information that may be needed by/relevant to the performer in their execution of this term action.
	SupportingInfo []Reference
	// Links to Provenance records for past versions of this Contract definition, derivative, or instance, which identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the Contract.  The Provence.entity indicates the target that was changed in the update. http://build.fhir.org/provenance-definitions.html#Provenance.entity.
	RelevantHistory []Reference
	// Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
	Signer []ContractSigner
	// The "patient friendly language" versionof the Contract in whole or in parts. "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
	Friendly []ContractFriendly
	// List of Legal expressions or representations of this Contract.
	Legal []ContractLegal
	// List of Computable Policy Rule Language Representations of this Contract.
	Rule []ContractRule
	// Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
	LegallyBinding isContractLegallyBinding
}

func (r Contract) ResourceType() string {
	return "Contract"
}
func (r Contract) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isContractTopic interface {
	isContractTopic()
}

func (r CodeableConcept) isContractTopic() {}
func (r Reference) isContractTopic()       {}

type isContractLegallyBinding interface {
	isContractLegallyBinding()
}

func (r Attachment) isContractLegallyBinding() {}
func (r Reference) isContractLegallyBinding()  {}

type jsonContract struct {
	ResourceType                    string                     `json:"resourceType"`
	Id                              *Id                        `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement          `json:"_id,omitempty"`
	Meta                            *Meta                      `json:"meta,omitempty"`
	ImplicitRules                   *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                        *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement          `json:"_language,omitempty"`
	Text                            *Narrative                 `json:"text,omitempty"`
	Contained                       []ContainedResource        `json:"contained,omitempty"`
	Extension                       []Extension                `json:"extension,omitempty"`
	ModifierExtension               []Extension                `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier               `json:"identifier,omitempty"`
	Url                             *Uri                       `json:"url,omitempty"`
	UrlPrimitiveElement             *primitiveElement          `json:"_url,omitempty"`
	Version                         *String                    `json:"version,omitempty"`
	VersionPrimitiveElement         *primitiveElement          `json:"_version,omitempty"`
	Status                          *Code                      `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement          `json:"_status,omitempty"`
	LegalState                      *CodeableConcept           `json:"legalState,omitempty"`
	InstantiatesCanonical           *Reference                 `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri                 *Uri                       `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement *primitiveElement          `json:"_instantiatesUri,omitempty"`
	ContentDerivative               *CodeableConcept           `json:"contentDerivative,omitempty"`
	Issued                          *DateTime                  `json:"issued,omitempty"`
	IssuedPrimitiveElement          *primitiveElement          `json:"_issued,omitempty"`
	Applies                         *Period                    `json:"applies,omitempty"`
	ExpirationType                  *CodeableConcept           `json:"expirationType,omitempty"`
	Subject                         []Reference                `json:"subject,omitempty"`
	Authority                       []Reference                `json:"authority,omitempty"`
	Domain                          []Reference                `json:"domain,omitempty"`
	Site                            []Reference                `json:"site,omitempty"`
	Name                            *String                    `json:"name,omitempty"`
	NamePrimitiveElement            *primitiveElement          `json:"_name,omitempty"`
	Title                           *String                    `json:"title,omitempty"`
	TitlePrimitiveElement           *primitiveElement          `json:"_title,omitempty"`
	Subtitle                        *String                    `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement        *primitiveElement          `json:"_subtitle,omitempty"`
	Alias                           []String                   `json:"alias,omitempty"`
	AliasPrimitiveElement           []*primitiveElement        `json:"_alias,omitempty"`
	Author                          *Reference                 `json:"author,omitempty"`
	Scope                           *CodeableConcept           `json:"scope,omitempty"`
	TopicCodeableConcept            *CodeableConcept           `json:"topicCodeableConcept,omitempty"`
	TopicReference                  *Reference                 `json:"topicReference,omitempty"`
	Type                            *CodeableConcept           `json:"type,omitempty"`
	SubType                         []CodeableConcept          `json:"subType,omitempty"`
	ContentDefinition               *ContractContentDefinition `json:"contentDefinition,omitempty"`
	Term                            []ContractTerm             `json:"term,omitempty"`
	SupportingInfo                  []Reference                `json:"supportingInfo,omitempty"`
	RelevantHistory                 []Reference                `json:"relevantHistory,omitempty"`
	Signer                          []ContractSigner           `json:"signer,omitempty"`
	Friendly                        []ContractFriendly         `json:"friendly,omitempty"`
	Legal                           []ContractLegal            `json:"legal,omitempty"`
	Rule                            []ContractRule             `json:"rule,omitempty"`
	LegallyBindingAttachment        *Attachment                `json:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference         *Reference                 `json:"legallyBindingReference,omitempty"`
}

func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Contract) marshalJSON() jsonContract {
	m := jsonContract{}
	m.ResourceType = "Contract"
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
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.LegalState = r.LegalState
	m.InstantiatesCanonical = r.InstantiatesCanonical
	if r.InstantiatesUri != nil && r.InstantiatesUri.Value != nil {
		m.InstantiatesUri = r.InstantiatesUri
	}
	if r.InstantiatesUri != nil && (r.InstantiatesUri.Id != nil || r.InstantiatesUri.Extension != nil) {
		m.InstantiatesUriPrimitiveElement = &primitiveElement{Id: r.InstantiatesUri.Id, Extension: r.InstantiatesUri.Extension}
	}
	m.ContentDerivative = r.ContentDerivative
	if r.Issued != nil && r.Issued.Value != nil {
		m.Issued = r.Issued
	}
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Applies = r.Applies
	m.ExpirationType = r.ExpirationType
	m.Subject = r.Subject
	m.Authority = r.Authority
	m.Domain = r.Domain
	m.Site = r.Site
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Subtitle != nil && r.Subtitle.Value != nil {
		m.Subtitle = r.Subtitle
	}
	if r.Subtitle != nil && (r.Subtitle.Id != nil || r.Subtitle.Extension != nil) {
		m.SubtitlePrimitiveElement = &primitiveElement{Id: r.Subtitle.Id, Extension: r.Subtitle.Extension}
	}
	anyAliasValue := false
	for _, e := range r.Alias {
		if e.Value != nil {
			anyAliasValue = true
			break
		}
	}
	if anyAliasValue {
		m.Alias = r.Alias
	}
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
	m.Author = r.Author
	m.Scope = r.Scope
	switch v := r.Topic.(type) {
	case CodeableConcept:
		m.TopicCodeableConcept = &v
	case *CodeableConcept:
		m.TopicCodeableConcept = v
	case Reference:
		m.TopicReference = &v
	case *Reference:
		m.TopicReference = v
	}
	m.Type = r.Type
	m.SubType = r.SubType
	m.ContentDefinition = r.ContentDefinition
	m.Term = r.Term
	m.SupportingInfo = r.SupportingInfo
	m.RelevantHistory = r.RelevantHistory
	m.Signer = r.Signer
	m.Friendly = r.Friendly
	m.Legal = r.Legal
	m.Rule = r.Rule
	switch v := r.LegallyBinding.(type) {
	case Attachment:
		m.LegallyBindingAttachment = &v
	case *Attachment:
		m.LegallyBindingAttachment = v
	case Reference:
		m.LegallyBindingReference = &v
	case *Reference:
		m.LegallyBindingReference = v
	}
	return m
}
func (r *Contract) UnmarshalJSON(b []byte) error {
	var m jsonContract
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Contract) unmarshalJSON(m jsonContract) error {
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
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		if r.Status == nil {
			r.Status = &Code{}
		}
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.LegalState = m.LegalState
	r.InstantiatesCanonical = m.InstantiatesCanonical
	r.InstantiatesUri = m.InstantiatesUri
	if m.InstantiatesUriPrimitiveElement != nil {
		if r.InstantiatesUri == nil {
			r.InstantiatesUri = &Uri{}
		}
		r.InstantiatesUri.Id = m.InstantiatesUriPrimitiveElement.Id
		r.InstantiatesUri.Extension = m.InstantiatesUriPrimitiveElement.Extension
	}
	r.ContentDerivative = m.ContentDerivative
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		if r.Issued == nil {
			r.Issued = &DateTime{}
		}
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Applies = m.Applies
	r.ExpirationType = m.ExpirationType
	r.Subject = m.Subject
	r.Authority = m.Authority
	r.Domain = m.Domain
	r.Site = m.Site
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Subtitle = m.Subtitle
	if m.SubtitlePrimitiveElement != nil {
		if r.Subtitle == nil {
			r.Subtitle = &String{}
		}
		r.Subtitle.Id = m.SubtitlePrimitiveElement.Id
		r.Subtitle.Extension = m.SubtitlePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	for i, e := range m.AliasPrimitiveElement {
		if len(r.Alias) <= i {
			r.Alias = append(r.Alias, String{})
		}
		if e != nil {
			r.Alias[i].Id = e.Id
			r.Alias[i].Extension = e.Extension
		}
	}
	r.Author = m.Author
	r.Scope = m.Scope
	if m.TopicCodeableConcept != nil {
		if r.Topic != nil {
			return fmt.Errorf("multiple values for field \"Topic\"")
		}
		v := m.TopicCodeableConcept
		r.Topic = v
	}
	if m.TopicReference != nil {
		if r.Topic != nil {
			return fmt.Errorf("multiple values for field \"Topic\"")
		}
		v := m.TopicReference
		r.Topic = v
	}
	r.Type = m.Type
	r.SubType = m.SubType
	r.ContentDefinition = m.ContentDefinition
	r.Term = m.Term
	r.SupportingInfo = m.SupportingInfo
	r.RelevantHistory = m.RelevantHistory
	r.Signer = m.Signer
	r.Friendly = m.Friendly
	r.Legal = m.Legal
	r.Rule = m.Rule
	if m.LegallyBindingAttachment != nil {
		if r.LegallyBinding != nil {
			return fmt.Errorf("multiple values for field \"LegallyBinding\"")
		}
		v := m.LegallyBindingAttachment
		r.LegallyBinding = v
	}
	if m.LegallyBindingReference != nil {
		if r.LegallyBinding != nil {
			return fmt.Errorf("multiple values for field \"LegallyBinding\"")
		}
		v := m.LegallyBindingReference
		r.LegallyBinding = v
	}
	return nil
}
func (r Contract) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Contract"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LegalState, xml.StartElement{Name: xml.Name{Local: "legalState"}})
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
	err = e.EncodeElement(r.ContentDerivative, xml.StartElement{Name: xml.Name{Local: "contentDerivative"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Issued, xml.StartElement{Name: xml.Name{Local: "issued"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Applies, xml.StartElement{Name: xml.Name{Local: "applies"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExpirationType, xml.StartElement{Name: xml.Name{Local: "expirationType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Authority, xml.StartElement{Name: xml.Name{Local: "authority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Domain, xml.StartElement{Name: xml.Name{Local: "domain"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Site, xml.StartElement{Name: xml.Name{Local: "site"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subtitle, xml.StartElement{Name: xml.Name{Local: "subtitle"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Alias, xml.StartElement{Name: xml.Name{Local: "alias"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Scope, xml.StartElement{Name: xml.Name{Local: "scope"}})
	if err != nil {
		return err
	}
	switch v := r.Topic.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "topicCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "topicReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubType, xml.StartElement{Name: xml.Name{Local: "subType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContentDefinition, xml.StartElement{Name: xml.Name{Local: "contentDefinition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Term, xml.StartElement{Name: xml.Name{Local: "term"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingInfo, xml.StartElement{Name: xml.Name{Local: "supportingInfo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelevantHistory, xml.StartElement{Name: xml.Name{Local: "relevantHistory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Signer, xml.StartElement{Name: xml.Name{Local: "signer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Friendly, xml.StartElement{Name: xml.Name{Local: "friendly"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Legal, xml.StartElement{Name: xml.Name{Local: "legal"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Rule, xml.StartElement{Name: xml.Name{Local: "rule"}})
	if err != nil {
		return err
	}
	switch v := r.LegallyBinding.(type) {
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "legallyBindingAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "legallyBindingReference"}})
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
func (r *Contract) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "legalState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegalState = &v
			case "instantiatesCanonical":
				var v Reference
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
			case "contentDerivative":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContentDerivative = &v
			case "issued":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Issued = &v
			case "applies":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Applies = &v
			case "expirationType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExpirationType = &v
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = append(r.Subject, v)
			case "authority":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authority = append(r.Authority, v)
			case "domain":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Domain = append(r.Domain, v)
			case "site":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Site = append(r.Site, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "subtitle":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subtitle = &v
			case "alias":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Alias = append(r.Alias, v)
			case "author":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = &v
			case "scope":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Scope = &v
			case "topicCodeableConcept":
				if r.Topic != nil {
					return fmt.Errorf("multiple values for field \"Topic\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = v
			case "topicReference":
				if r.Topic != nil {
					return fmt.Errorf("multiple values for field \"Topic\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "subType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubType = append(r.SubType, v)
			case "contentDefinition":
				var v ContractContentDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContentDefinition = &v
			case "term":
				var v ContractTerm
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Term = append(r.Term, v)
			case "supportingInfo":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingInfo = append(r.SupportingInfo, v)
			case "relevantHistory":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelevantHistory = append(r.RelevantHistory, v)
			case "signer":
				var v ContractSigner
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Signer = append(r.Signer, v)
			case "friendly":
				var v ContractFriendly
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Friendly = append(r.Friendly, v)
			case "legal":
				var v ContractLegal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Legal = append(r.Legal, v)
			case "rule":
				var v ContractRule
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rule = append(r.Rule, v)
			case "legallyBindingAttachment":
				if r.LegallyBinding != nil {
					return fmt.Errorf("multiple values for field \"LegallyBinding\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegallyBinding = v
			case "legallyBindingReference":
				if r.LegallyBinding != nil {
					return fmt.Errorf("multiple values for field \"LegallyBinding\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegallyBinding = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Contract) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Precusory content developed with a focus and intent of supporting the formation a Contract instance, which may be associated with and transformable into a Contract.
type ContractContentDefinition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Precusory content structure and use, i.e., a boilerplate, template, application for a contract such as an insurance policy or benefits under a program, e.g., workers compensation.
	Type CodeableConcept
	// Detailed Precusory content type.
	SubType *CodeableConcept
	// The  individual or organization that published the Contract precursor content.
	Publisher *Reference
	// The date (and optionally time) when the contract was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the contract changes.
	PublicationDate *DateTime
	// amended | appended | cancelled | disputed | entered-in-error | executable | executed | negotiable | offered | policy | rejected | renewed | revoked | resolved | terminated.
	PublicationStatus Code
	// A copyright statement relating to Contract precursor content. Copyright statements are generally legal restrictions on the use and publishing of the Contract precursor content.
	Copyright *Markdown
}
type jsonContractContentDefinition struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	Type                              CodeableConcept   `json:"type,omitempty"`
	SubType                           *CodeableConcept  `json:"subType,omitempty"`
	Publisher                         *Reference        `json:"publisher,omitempty"`
	PublicationDate                   *DateTime         `json:"publicationDate,omitempty"`
	PublicationDatePrimitiveElement   *primitiveElement `json:"_publicationDate,omitempty"`
	PublicationStatus                 Code              `json:"publicationStatus,omitempty"`
	PublicationStatusPrimitiveElement *primitiveElement `json:"_publicationStatus,omitempty"`
	Copyright                         *Markdown         `json:"copyright,omitempty"`
	CopyrightPrimitiveElement         *primitiveElement `json:"_copyright,omitempty"`
}

func (r ContractContentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractContentDefinition) marshalJSON() jsonContractContentDefinition {
	m := jsonContractContentDefinition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.SubType = r.SubType
	m.Publisher = r.Publisher
	if r.PublicationDate != nil && r.PublicationDate.Value != nil {
		m.PublicationDate = r.PublicationDate
	}
	if r.PublicationDate != nil && (r.PublicationDate.Id != nil || r.PublicationDate.Extension != nil) {
		m.PublicationDatePrimitiveElement = &primitiveElement{Id: r.PublicationDate.Id, Extension: r.PublicationDate.Extension}
	}
	if r.PublicationStatus.Value != nil {
		m.PublicationStatus = r.PublicationStatus
	}
	if r.PublicationStatus.Id != nil || r.PublicationStatus.Extension != nil {
		m.PublicationStatusPrimitiveElement = &primitiveElement{Id: r.PublicationStatus.Id, Extension: r.PublicationStatus.Extension}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	return m
}
func (r *ContractContentDefinition) UnmarshalJSON(b []byte) error {
	var m jsonContractContentDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractContentDefinition) unmarshalJSON(m jsonContractContentDefinition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.SubType = m.SubType
	r.Publisher = m.Publisher
	r.PublicationDate = m.PublicationDate
	if m.PublicationDatePrimitiveElement != nil {
		if r.PublicationDate == nil {
			r.PublicationDate = &DateTime{}
		}
		r.PublicationDate.Id = m.PublicationDatePrimitiveElement.Id
		r.PublicationDate.Extension = m.PublicationDatePrimitiveElement.Extension
	}
	r.PublicationStatus = m.PublicationStatus
	if m.PublicationStatusPrimitiveElement != nil {
		r.PublicationStatus.Id = m.PublicationStatusPrimitiveElement.Id
		r.PublicationStatus.Extension = m.PublicationStatusPrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	return nil
}
func (r ContractContentDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.SubType, xml.StartElement{Name: xml.Name{Local: "subType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PublicationDate, xml.StartElement{Name: xml.Name{Local: "publicationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PublicationStatus, xml.StartElement{Name: xml.Name{Local: "publicationStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractContentDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubType = &v
			case "publisher":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "publicationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublicationDate = &v
			case "publicationStatus":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublicationStatus = v
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractContentDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
type ContractTerm struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique identifier for this particular Contract Provision.
	Identifier *Identifier
	// When this Contract Provision was issued.
	Issued *DateTime
	// Relevant time or time-period when this Contract Provision is applicable.
	Applies *Period
	// The entity that the term applies to.
	Topic isContractTermTopic
	// A legal clause or condition contained within a contract that requires one or both parties to perform a particular requirement by some specified time or prevents one or both parties from performing a particular requirement by some specified time.
	Type *CodeableConcept
	// A specialized legal clause or condition based on overarching contract type.
	SubType *CodeableConcept
	// Statement of a provision in a policy or a contract.
	Text *String
	// Security labels that protect the handling of information about the term and its elements, which may be specifically identified..
	SecurityLabel []ContractTermSecurityLabel
	// The matter of concern in the context of this provision of the agrement.
	Offer ContractTermOffer
	// Contract Term Asset List.
	Asset []ContractTermAsset
	// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
	Action []ContractTermAction
	// Nested group of Contract Provisions.
	Group []ContractTerm
}
type isContractTermTopic interface {
	isContractTermTopic()
}

func (r CodeableConcept) isContractTermTopic() {}
func (r Reference) isContractTermTopic()       {}

type jsonContractTerm struct {
	Id                     *string                     `json:"id,omitempty"`
	Extension              []Extension                 `json:"extension,omitempty"`
	ModifierExtension      []Extension                 `json:"modifierExtension,omitempty"`
	Identifier             *Identifier                 `json:"identifier,omitempty"`
	Issued                 *DateTime                   `json:"issued,omitempty"`
	IssuedPrimitiveElement *primitiveElement           `json:"_issued,omitempty"`
	Applies                *Period                     `json:"applies,omitempty"`
	TopicCodeableConcept   *CodeableConcept            `json:"topicCodeableConcept,omitempty"`
	TopicReference         *Reference                  `json:"topicReference,omitempty"`
	Type                   *CodeableConcept            `json:"type,omitempty"`
	SubType                *CodeableConcept            `json:"subType,omitempty"`
	Text                   *String                     `json:"text,omitempty"`
	TextPrimitiveElement   *primitiveElement           `json:"_text,omitempty"`
	SecurityLabel          []ContractTermSecurityLabel `json:"securityLabel,omitempty"`
	Offer                  ContractTermOffer           `json:"offer,omitempty"`
	Asset                  []ContractTermAsset         `json:"asset,omitempty"`
	Action                 []ContractTermAction        `json:"action,omitempty"`
	Group                  []ContractTerm              `json:"group,omitempty"`
}

func (r ContractTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTerm) marshalJSON() jsonContractTerm {
	m := jsonContractTerm{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	if r.Issued != nil && r.Issued.Value != nil {
		m.Issued = r.Issued
	}
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Applies = r.Applies
	switch v := r.Topic.(type) {
	case CodeableConcept:
		m.TopicCodeableConcept = &v
	case *CodeableConcept:
		m.TopicCodeableConcept = v
	case Reference:
		m.TopicReference = &v
	case *Reference:
		m.TopicReference = v
	}
	m.Type = r.Type
	m.SubType = r.SubType
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.SecurityLabel = r.SecurityLabel
	m.Offer = r.Offer
	m.Asset = r.Asset
	m.Action = r.Action
	m.Group = r.Group
	return m
}
func (r *ContractTerm) UnmarshalJSON(b []byte) error {
	var m jsonContractTerm
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTerm) unmarshalJSON(m jsonContractTerm) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		if r.Issued == nil {
			r.Issued = &DateTime{}
		}
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Applies = m.Applies
	if m.TopicCodeableConcept != nil {
		if r.Topic != nil {
			return fmt.Errorf("multiple values for field \"Topic\"")
		}
		v := m.TopicCodeableConcept
		r.Topic = v
	}
	if m.TopicReference != nil {
		if r.Topic != nil {
			return fmt.Errorf("multiple values for field \"Topic\"")
		}
		v := m.TopicReference
		r.Topic = v
	}
	r.Type = m.Type
	r.SubType = m.SubType
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.SecurityLabel = m.SecurityLabel
	r.Offer = m.Offer
	r.Asset = m.Asset
	r.Action = m.Action
	r.Group = m.Group
	return nil
}
func (r ContractTerm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Issued, xml.StartElement{Name: xml.Name{Local: "issued"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Applies, xml.StartElement{Name: xml.Name{Local: "applies"}})
	if err != nil {
		return err
	}
	switch v := r.Topic.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "topicCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "topicReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubType, xml.StartElement{Name: xml.Name{Local: "subType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabel, xml.StartElement{Name: xml.Name{Local: "securityLabel"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Offer, xml.StartElement{Name: xml.Name{Local: "offer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Asset, xml.StartElement{Name: xml.Name{Local: "asset"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Action, xml.StartElement{Name: xml.Name{Local: "action"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Group, xml.StartElement{Name: xml.Name{Local: "group"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTerm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "issued":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Issued = &v
			case "applies":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Applies = &v
			case "topicCodeableConcept":
				if r.Topic != nil {
					return fmt.Errorf("multiple values for field \"Topic\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = v
			case "topicReference":
				if r.Topic != nil {
					return fmt.Errorf("multiple values for field \"Topic\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "subType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubType = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "securityLabel":
				var v ContractTermSecurityLabel
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabel = append(r.SecurityLabel, v)
			case "offer":
				var v ContractTermOffer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Offer = v
			case "asset":
				var v ContractTermAsset
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Asset = append(r.Asset, v)
			case "action":
				var v ContractTermAction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Action = append(r.Action, v)
			case "group":
				var v ContractTerm
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Group = append(r.Group, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTerm) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Security labels that protect the handling of information about the term and its elements, which may be specifically identified..
type ContractTermSecurityLabel struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Number used to link this term or term element to the applicable Security Label.
	Number []UnsignedInt
	// Security label privacy tag that species the level of confidentiality protection required for this term and/or term elements.
	Classification Coding
	// Security label privacy tag that species the applicable privacy and security policies governing this term and/or term elements.
	Category []Coding
	// Security label privacy tag that species the manner in which term and/or term elements are to be protected.
	Control []Coding
}
type jsonContractTermSecurityLabel struct {
	Id                     *string             `json:"id,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	ModifierExtension      []Extension         `json:"modifierExtension,omitempty"`
	Number                 []UnsignedInt       `json:"number,omitempty"`
	NumberPrimitiveElement []*primitiveElement `json:"_number,omitempty"`
	Classification         Coding              `json:"classification,omitempty"`
	Category               []Coding            `json:"category,omitempty"`
	Control                []Coding            `json:"control,omitempty"`
}

func (r ContractTermSecurityLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermSecurityLabel) marshalJSON() jsonContractTermSecurityLabel {
	m := jsonContractTermSecurityLabel{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyNumberValue := false
	for _, e := range r.Number {
		if e.Value != nil {
			anyNumberValue = true
			break
		}
	}
	if anyNumberValue {
		m.Number = r.Number
	}
	anyNumberIdOrExtension := false
	for _, e := range r.Number {
		if e.Id != nil || e.Extension != nil {
			anyNumberIdOrExtension = true
			break
		}
	}
	if anyNumberIdOrExtension {
		m.NumberPrimitiveElement = make([]*primitiveElement, 0, len(r.Number))
		for _, e := range r.Number {
			if e.Id != nil || e.Extension != nil {
				m.NumberPrimitiveElement = append(m.NumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NumberPrimitiveElement = append(m.NumberPrimitiveElement, nil)
			}
		}
	}
	m.Classification = r.Classification
	m.Category = r.Category
	m.Control = r.Control
	return m
}
func (r *ContractTermSecurityLabel) UnmarshalJSON(b []byte) error {
	var m jsonContractTermSecurityLabel
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermSecurityLabel) unmarshalJSON(m jsonContractTermSecurityLabel) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Number = m.Number
	for i, e := range m.NumberPrimitiveElement {
		if len(r.Number) <= i {
			r.Number = append(r.Number, UnsignedInt{})
		}
		if e != nil {
			r.Number[i].Id = e.Id
			r.Number[i].Extension = e.Extension
		}
	}
	r.Classification = m.Classification
	r.Category = m.Category
	r.Control = m.Control
	return nil
}
func (r ContractTermSecurityLabel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Number, xml.StartElement{Name: xml.Name{Local: "number"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Classification, xml.StartElement{Name: xml.Name{Local: "classification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Control, xml.StartElement{Name: xml.Name{Local: "control"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermSecurityLabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "number":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Number = append(r.Number, v)
			case "classification":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classification = v
			case "category":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "control":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Control = append(r.Control, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermSecurityLabel) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The matter of concern in the context of this provision of the agrement.
type ContractTermOffer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique identifier for this particular Contract Provision.
	Identifier []Identifier
	// Offer Recipient.
	Party []ContractTermOfferParty
	// The owner of an asset has the residual control rights over the asset: the right to decide all usages of the asset in any way not inconsistent with a prior contract, custom, or law (Hart, 1995, p. 30).
	Topic *Reference
	// Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions, e.g. life time maximum benefit.
	Type *CodeableConcept
	// Type of choice made by accepting party with respect to an offer made by an offeror/ grantee.
	Decision *CodeableConcept
	// How the decision about a Contract was conveyed.
	DecisionMode []CodeableConcept
	// Response to offer text.
	Answer []ContractTermOfferAnswer
	// Human readable form of this Contract Offer.
	Text *String
	// The id of the clause or question text of the offer in the referenced questionnaire/response.
	LinkId []String
	// Security labels that protects the offer.
	SecurityLabelNumber []UnsignedInt
}
type jsonContractTermOffer struct {
	Id                                  *string                   `json:"id,omitempty"`
	Extension                           []Extension               `json:"extension,omitempty"`
	ModifierExtension                   []Extension               `json:"modifierExtension,omitempty"`
	Identifier                          []Identifier              `json:"identifier,omitempty"`
	Party                               []ContractTermOfferParty  `json:"party,omitempty"`
	Topic                               *Reference                `json:"topic,omitempty"`
	Type                                *CodeableConcept          `json:"type,omitempty"`
	Decision                            *CodeableConcept          `json:"decision,omitempty"`
	DecisionMode                        []CodeableConcept         `json:"decisionMode,omitempty"`
	Answer                              []ContractTermOfferAnswer `json:"answer,omitempty"`
	Text                                *String                   `json:"text,omitempty"`
	TextPrimitiveElement                *primitiveElement         `json:"_text,omitempty"`
	LinkId                              []String                  `json:"linkId,omitempty"`
	LinkIdPrimitiveElement              []*primitiveElement       `json:"_linkId,omitempty"`
	SecurityLabelNumber                 []UnsignedInt             `json:"securityLabelNumber,omitempty"`
	SecurityLabelNumberPrimitiveElement []*primitiveElement       `json:"_securityLabelNumber,omitempty"`
}

func (r ContractTermOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermOffer) marshalJSON() jsonContractTermOffer {
	m := jsonContractTermOffer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Party = r.Party
	m.Topic = r.Topic
	m.Type = r.Type
	m.Decision = r.Decision
	m.DecisionMode = r.DecisionMode
	m.Answer = r.Answer
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	anyLinkIdValue := false
	for _, e := range r.LinkId {
		if e.Value != nil {
			anyLinkIdValue = true
			break
		}
	}
	if anyLinkIdValue {
		m.LinkId = r.LinkId
	}
	anyLinkIdIdOrExtension := false
	for _, e := range r.LinkId {
		if e.Id != nil || e.Extension != nil {
			anyLinkIdIdOrExtension = true
			break
		}
	}
	if anyLinkIdIdOrExtension {
		m.LinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.LinkId))
		for _, e := range r.LinkId {
			if e.Id != nil || e.Extension != nil {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, nil)
			}
		}
	}
	anySecurityLabelNumberValue := false
	for _, e := range r.SecurityLabelNumber {
		if e.Value != nil {
			anySecurityLabelNumberValue = true
			break
		}
	}
	if anySecurityLabelNumberValue {
		m.SecurityLabelNumber = r.SecurityLabelNumber
	}
	anySecurityLabelNumberIdOrExtension := false
	for _, e := range r.SecurityLabelNumber {
		if e.Id != nil || e.Extension != nil {
			anySecurityLabelNumberIdOrExtension = true
			break
		}
	}
	if anySecurityLabelNumberIdOrExtension {
		m.SecurityLabelNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.SecurityLabelNumber))
		for _, e := range r.SecurityLabelNumber {
			if e.Id != nil || e.Extension != nil {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ContractTermOffer) UnmarshalJSON(b []byte) error {
	var m jsonContractTermOffer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermOffer) unmarshalJSON(m jsonContractTermOffer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Party = m.Party
	r.Topic = m.Topic
	r.Type = m.Type
	r.Decision = m.Decision
	r.DecisionMode = m.DecisionMode
	r.Answer = m.Answer
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.LinkId = m.LinkId
	for i, e := range m.LinkIdPrimitiveElement {
		if len(r.LinkId) <= i {
			r.LinkId = append(r.LinkId, String{})
		}
		if e != nil {
			r.LinkId[i].Id = e.Id
			r.LinkId[i].Extension = e.Extension
		}
	}
	r.SecurityLabelNumber = m.SecurityLabelNumber
	for i, e := range m.SecurityLabelNumberPrimitiveElement {
		if len(r.SecurityLabelNumber) <= i {
			r.SecurityLabelNumber = append(r.SecurityLabelNumber, UnsignedInt{})
		}
		if e != nil {
			r.SecurityLabelNumber[i].Id = e.Id
			r.SecurityLabelNumber[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ContractTermOffer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Party, xml.StartElement{Name: xml.Name{Local: "party"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Topic, xml.StartElement{Name: xml.Name{Local: "topic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Decision, xml.StartElement{Name: xml.Name{Local: "decision"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DecisionMode, xml.StartElement{Name: xml.Name{Local: "decisionMode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Answer, xml.StartElement{Name: xml.Name{Local: "answer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LinkId, xml.StartElement{Name: xml.Name{Local: "linkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabelNumber, xml.StartElement{Name: xml.Name{Local: "securityLabelNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermOffer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "party":
				var v ContractTermOfferParty
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Party = append(r.Party, v)
			case "topic":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "decision":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Decision = &v
			case "decisionMode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DecisionMode = append(r.DecisionMode, v)
			case "answer":
				var v ContractTermOfferAnswer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Answer = append(r.Answer, v)
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "linkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LinkId = append(r.LinkId, v)
			case "securityLabelNumber":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabelNumber = append(r.SecurityLabelNumber, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermOffer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Offer Recipient.
type ContractTermOfferParty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Participant in the offer.
	Reference []Reference
	// How the party participates in the offer.
	Role CodeableConcept
}
type jsonContractTermOfferParty struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Reference         []Reference     `json:"reference,omitempty"`
	Role              CodeableConcept `json:"role,omitempty"`
}

func (r ContractTermOfferParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermOfferParty) marshalJSON() jsonContractTermOfferParty {
	m := jsonContractTermOfferParty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Reference = r.Reference
	m.Role = r.Role
	return m
}
func (r *ContractTermOfferParty) UnmarshalJSON(b []byte) error {
	var m jsonContractTermOfferParty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermOfferParty) unmarshalJSON(m jsonContractTermOfferParty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Reference = m.Reference
	r.Role = m.Role
	return nil
}
func (r ContractTermOfferParty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermOfferParty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = append(r.Reference, v)
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermOfferParty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Response to offer text.
type ContractTermOfferAnswer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Response to an offer clause or question text,  which enables selection of values to be agreed to, e.g., the period of participation, the date of occupancy of a rental, warrently duration, or whether biospecimen may be used for further research.
	Value isContractTermOfferAnswerValue
}
type isContractTermOfferAnswerValue interface {
	isContractTermOfferAnswerValue()
}

func (r Boolean) isContractTermOfferAnswerValue()    {}
func (r Decimal) isContractTermOfferAnswerValue()    {}
func (r Integer) isContractTermOfferAnswerValue()    {}
func (r Date) isContractTermOfferAnswerValue()       {}
func (r DateTime) isContractTermOfferAnswerValue()   {}
func (r Time) isContractTermOfferAnswerValue()       {}
func (r String) isContractTermOfferAnswerValue()     {}
func (r Uri) isContractTermOfferAnswerValue()        {}
func (r Attachment) isContractTermOfferAnswerValue() {}
func (r Coding) isContractTermOfferAnswerValue()     {}
func (r Quantity) isContractTermOfferAnswerValue()   {}
func (r Reference) isContractTermOfferAnswerValue()  {}

type jsonContractTermOfferAnswer struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	ValueBoolean                  *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueDecimal                  *Decimal          `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement  *primitiveElement `json:"_valueDecimal,omitempty"`
	ValueInteger                  *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement `json:"_valueInteger,omitempty"`
	ValueDate                     *Date             `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement     *primitiveElement `json:"_valueDate,omitempty"`
	ValueDateTime                 *DateTime         `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement `json:"_valueDateTime,omitempty"`
	ValueTime                     *Time             `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement     *primitiveElement `json:"_valueTime,omitempty"`
	ValueString                   *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement `json:"_valueString,omitempty"`
	ValueUri                      *Uri              `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement      *primitiveElement `json:"_valueUri,omitempty"`
	ValueAttachment               *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding                   *Coding           `json:"valueCoding,omitempty"`
	ValueQuantity                 *Quantity         `json:"valueQuantity,omitempty"`
	ValueReference                *Reference        `json:"valueReference,omitempty"`
}

func (r ContractTermOfferAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermOfferAnswer) marshalJSON() jsonContractTermOfferAnswer {
	m := jsonContractTermOfferAnswer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Value.(type) {
	case Boolean:
		if v.Value != nil {
			m.ValueBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.ValueBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		if v.Value != nil {
			m.ValueDecimal = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		if v.Value != nil {
			m.ValueDecimal = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		if v.Value != nil {
			m.ValueInteger = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		if v.Value != nil {
			m.ValueInteger = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		if v.Value != nil {
			m.ValueDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.ValueDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		if v.Value != nil {
			m.ValueDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.ValueDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		if v.Value != nil {
			m.ValueTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		if v.Value != nil {
			m.ValueTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
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
	case Uri:
		if v.Value != nil {
			m.ValueUri = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		if v.Value != nil {
			m.ValueUri = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	return m
}
func (r *ContractTermOfferAnswer) UnmarshalJSON(b []byte) error {
	var m jsonContractTermOfferAnswer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermOfferAnswer) unmarshalJSON(m jsonContractTermOfferAnswer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDecimal != nil || m.ValueDecimalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDecimal
		if m.ValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ValueDecimalPrimitiveElement.Id
			v.Extension = m.ValueDecimalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDate != nil || m.ValueDatePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDate
		if m.ValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ValueDatePrimitiveElement.Id
			v.Extension = m.ValueDatePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDateTime != nil || m.ValueDateTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDateTime
		if m.ValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ValueDateTimePrimitiveElement.Id
			v.Extension = m.ValueDateTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueTime != nil || m.ValueTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTime
		if m.ValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.ValueTimePrimitiveElement.Id
			v.Extension = m.ValueTimePrimitiveElement.Extension
		}
		r.Value = v
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
	if m.ValueUri != nil || m.ValueUriPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUri
		if m.ValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ValueUriPrimitiveElement.Id
			v.Extension = m.ValueUriPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	return nil
}
func (r ContractTermOfferAnswer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Value.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Decimal, *Decimal:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDecimal"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger"}})
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
	case Time, *Time:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTime"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUri"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case Coding, *Coding:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCoding"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
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
func (r *ContractTermOfferAnswer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermOfferAnswer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Contract Term Asset List.
type ContractTermAsset struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Differentiates the kind of the asset .
	Scope *CodeableConcept
	// Target entity type about which the term may be concerned.
	Type []CodeableConcept
	// Associated entities.
	TypeReference []Reference
	// May be a subtype or part of an offered asset.
	Subtype []CodeableConcept
	// Specifies the applicability of the term to an asset resource instance, and instances it refers to orinstances that refer to it, and/or are owned by the offeree.
	Relationship *Coding
	// Circumstance of the asset.
	Context []ContractTermAssetContext
	// Description of the quality and completeness of the asset that imay be a factor in its valuation.
	Condition *String
	// Type of Asset availability for use or ownership.
	PeriodType []CodeableConcept
	// Asset relevant contractual time period.
	Period []Period
	// Time period of asset use.
	UsePeriod []Period
	// Clause or question text (Prose Object) concerning the asset in a linked form, such as a QuestionnaireResponse used in the formation of the contract.
	Text *String
	// Id [identifier??] of the clause or question text about the asset in the referenced form or QuestionnaireResponse.
	LinkId []String
	// Response to assets.
	Answer []ContractTermOfferAnswer
	// Security labels that protects the asset.
	SecurityLabelNumber []UnsignedInt
	// Contract Valued Item List.
	ValuedItem []ContractTermAssetValuedItem
}
type jsonContractTermAsset struct {
	Id                                  *string                       `json:"id,omitempty"`
	Extension                           []Extension                   `json:"extension,omitempty"`
	ModifierExtension                   []Extension                   `json:"modifierExtension,omitempty"`
	Scope                               *CodeableConcept              `json:"scope,omitempty"`
	Type                                []CodeableConcept             `json:"type,omitempty"`
	TypeReference                       []Reference                   `json:"typeReference,omitempty"`
	Subtype                             []CodeableConcept             `json:"subtype,omitempty"`
	Relationship                        *Coding                       `json:"relationship,omitempty"`
	Context                             []ContractTermAssetContext    `json:"context,omitempty"`
	Condition                           *String                       `json:"condition,omitempty"`
	ConditionPrimitiveElement           *primitiveElement             `json:"_condition,omitempty"`
	PeriodType                          []CodeableConcept             `json:"periodType,omitempty"`
	Period                              []Period                      `json:"period,omitempty"`
	UsePeriod                           []Period                      `json:"usePeriod,omitempty"`
	Text                                *String                       `json:"text,omitempty"`
	TextPrimitiveElement                *primitiveElement             `json:"_text,omitempty"`
	LinkId                              []String                      `json:"linkId,omitempty"`
	LinkIdPrimitiveElement              []*primitiveElement           `json:"_linkId,omitempty"`
	Answer                              []ContractTermOfferAnswer     `json:"answer,omitempty"`
	SecurityLabelNumber                 []UnsignedInt                 `json:"securityLabelNumber,omitempty"`
	SecurityLabelNumberPrimitiveElement []*primitiveElement           `json:"_securityLabelNumber,omitempty"`
	ValuedItem                          []ContractTermAssetValuedItem `json:"valuedItem,omitempty"`
}

func (r ContractTermAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermAsset) marshalJSON() jsonContractTermAsset {
	m := jsonContractTermAsset{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Scope = r.Scope
	m.Type = r.Type
	m.TypeReference = r.TypeReference
	m.Subtype = r.Subtype
	m.Relationship = r.Relationship
	m.Context = r.Context
	if r.Condition != nil && r.Condition.Value != nil {
		m.Condition = r.Condition
	}
	if r.Condition != nil && (r.Condition.Id != nil || r.Condition.Extension != nil) {
		m.ConditionPrimitiveElement = &primitiveElement{Id: r.Condition.Id, Extension: r.Condition.Extension}
	}
	m.PeriodType = r.PeriodType
	m.Period = r.Period
	m.UsePeriod = r.UsePeriod
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	anyLinkIdValue := false
	for _, e := range r.LinkId {
		if e.Value != nil {
			anyLinkIdValue = true
			break
		}
	}
	if anyLinkIdValue {
		m.LinkId = r.LinkId
	}
	anyLinkIdIdOrExtension := false
	for _, e := range r.LinkId {
		if e.Id != nil || e.Extension != nil {
			anyLinkIdIdOrExtension = true
			break
		}
	}
	if anyLinkIdIdOrExtension {
		m.LinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.LinkId))
		for _, e := range r.LinkId {
			if e.Id != nil || e.Extension != nil {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, nil)
			}
		}
	}
	m.Answer = r.Answer
	anySecurityLabelNumberValue := false
	for _, e := range r.SecurityLabelNumber {
		if e.Value != nil {
			anySecurityLabelNumberValue = true
			break
		}
	}
	if anySecurityLabelNumberValue {
		m.SecurityLabelNumber = r.SecurityLabelNumber
	}
	anySecurityLabelNumberIdOrExtension := false
	for _, e := range r.SecurityLabelNumber {
		if e.Id != nil || e.Extension != nil {
			anySecurityLabelNumberIdOrExtension = true
			break
		}
	}
	if anySecurityLabelNumberIdOrExtension {
		m.SecurityLabelNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.SecurityLabelNumber))
		for _, e := range r.SecurityLabelNumber {
			if e.Id != nil || e.Extension != nil {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, nil)
			}
		}
	}
	m.ValuedItem = r.ValuedItem
	return m
}
func (r *ContractTermAsset) UnmarshalJSON(b []byte) error {
	var m jsonContractTermAsset
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermAsset) unmarshalJSON(m jsonContractTermAsset) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Scope = m.Scope
	r.Type = m.Type
	r.TypeReference = m.TypeReference
	r.Subtype = m.Subtype
	r.Relationship = m.Relationship
	r.Context = m.Context
	r.Condition = m.Condition
	if m.ConditionPrimitiveElement != nil {
		if r.Condition == nil {
			r.Condition = &String{}
		}
		r.Condition.Id = m.ConditionPrimitiveElement.Id
		r.Condition.Extension = m.ConditionPrimitiveElement.Extension
	}
	r.PeriodType = m.PeriodType
	r.Period = m.Period
	r.UsePeriod = m.UsePeriod
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.LinkId = m.LinkId
	for i, e := range m.LinkIdPrimitiveElement {
		if len(r.LinkId) <= i {
			r.LinkId = append(r.LinkId, String{})
		}
		if e != nil {
			r.LinkId[i].Id = e.Id
			r.LinkId[i].Extension = e.Extension
		}
	}
	r.Answer = m.Answer
	r.SecurityLabelNumber = m.SecurityLabelNumber
	for i, e := range m.SecurityLabelNumberPrimitiveElement {
		if len(r.SecurityLabelNumber) <= i {
			r.SecurityLabelNumber = append(r.SecurityLabelNumber, UnsignedInt{})
		}
		if e != nil {
			r.SecurityLabelNumber[i].Id = e.Id
			r.SecurityLabelNumber[i].Extension = e.Extension
		}
	}
	r.ValuedItem = m.ValuedItem
	return nil
}
func (r ContractTermAsset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Scope, xml.StartElement{Name: xml.Name{Local: "scope"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TypeReference, xml.StartElement{Name: xml.Name{Local: "typeReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subtype, xml.StartElement{Name: xml.Name{Local: "subtype"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relationship, xml.StartElement{Name: xml.Name{Local: "relationship"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Context, xml.StartElement{Name: xml.Name{Local: "context"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Condition, xml.StartElement{Name: xml.Name{Local: "condition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PeriodType, xml.StartElement{Name: xml.Name{Local: "periodType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UsePeriod, xml.StartElement{Name: xml.Name{Local: "usePeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LinkId, xml.StartElement{Name: xml.Name{Local: "linkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Answer, xml.StartElement{Name: xml.Name{Local: "answer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabelNumber, xml.StartElement{Name: xml.Name{Local: "securityLabelNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValuedItem, xml.StartElement{Name: xml.Name{Local: "valuedItem"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermAsset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "scope":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Scope = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "typeReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TypeReference = append(r.TypeReference, v)
			case "subtype":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subtype = append(r.Subtype, v)
			case "relationship":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationship = &v
			case "context":
				var v ContractTermAssetContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Context = append(r.Context, v)
			case "condition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Condition = &v
			case "periodType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PeriodType = append(r.PeriodType, v)
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = append(r.Period, v)
			case "usePeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UsePeriod = append(r.UsePeriod, v)
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "linkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LinkId = append(r.LinkId, v)
			case "answer":
				var v ContractTermOfferAnswer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Answer = append(r.Answer, v)
			case "securityLabelNumber":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabelNumber = append(r.SecurityLabelNumber, v)
			case "valuedItem":
				var v ContractTermAssetValuedItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValuedItem = append(r.ValuedItem, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermAsset) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Circumstance of the asset.
type ContractTermAssetContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Asset context reference may include the creator, custodian, or owning Person or Organization (e.g., bank, repository),  location held, e.g., building,  jurisdiction.
	Reference *Reference
	// Coded representation of the context generally or of the Referenced entity, such as the asset holder type or location.
	Code []CodeableConcept
	// Context description.
	Text *String
}
type jsonContractTermAssetContext struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Reference            *Reference        `json:"reference,omitempty"`
	Code                 []CodeableConcept `json:"code,omitempty"`
	Text                 *String           `json:"text,omitempty"`
	TextPrimitiveElement *primitiveElement `json:"_text,omitempty"`
}

func (r ContractTermAssetContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermAssetContext) marshalJSON() jsonContractTermAssetContext {
	m := jsonContractTermAssetContext{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Reference = r.Reference
	m.Code = r.Code
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	return m
}
func (r *ContractTermAssetContext) UnmarshalJSON(b []byte) error {
	var m jsonContractTermAssetContext
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermAssetContext) unmarshalJSON(m jsonContractTermAssetContext) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Reference = m.Reference
	r.Code = m.Code
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	return nil
}
func (r ContractTermAssetContext) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermAssetContext) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermAssetContext) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Contract Valued Item List.
type ContractTermAssetValuedItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specific type of Contract Valued Item that may be priced.
	Entity isContractTermAssetValuedItemEntity
	// Identifies a Contract Valued Item instance.
	Identifier *Identifier
	// Indicates the time during which this Contract ValuedItem information is effective.
	EffectiveTime *DateTime
	// Specifies the units by which the Contract Valued Item is measured or counted, and quantifies the countable or measurable Contract Valued Item instances.
	Quantity *Quantity
	// A Contract Valued Item unit valuation measure.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of the Contract Valued Item delivered. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// An amount that expresses the weighting (based on difficulty, cost and/or resource intensiveness) associated with the Contract Valued Item delivered. The concept of Points allows for assignment of point values for a Contract Valued Item, such that a monetary amount can be assigned to each point.
	Points *Decimal
	// Expresses the product of the Contract Valued Item unitQuantity and the unitPriceAmt. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money
	// Terms of valuation.
	Payment *String
	// When payment is due.
	PaymentDate *DateTime
	// Who will make payment.
	Responsible *Reference
	// Who will receive payment.
	Recipient *Reference
	// Id  of the clause or question text related to the context of this valuedItem in the referenced form or QuestionnaireResponse.
	LinkId []String
	// A set of security labels that define which terms are controlled by this condition.
	SecurityLabelNumber []UnsignedInt
}
type isContractTermAssetValuedItemEntity interface {
	isContractTermAssetValuedItemEntity()
}

func (r CodeableConcept) isContractTermAssetValuedItemEntity() {}
func (r Reference) isContractTermAssetValuedItemEntity()       {}

type jsonContractTermAssetValuedItem struct {
	Id                                  *string             `json:"id,omitempty"`
	Extension                           []Extension         `json:"extension,omitempty"`
	ModifierExtension                   []Extension         `json:"modifierExtension,omitempty"`
	EntityCodeableConcept               *CodeableConcept    `json:"entityCodeableConcept,omitempty"`
	EntityReference                     *Reference          `json:"entityReference,omitempty"`
	Identifier                          *Identifier         `json:"identifier,omitempty"`
	EffectiveTime                       *DateTime           `json:"effectiveTime,omitempty"`
	EffectiveTimePrimitiveElement       *primitiveElement   `json:"_effectiveTime,omitempty"`
	Quantity                            *Quantity           `json:"quantity,omitempty"`
	UnitPrice                           *Money              `json:"unitPrice,omitempty"`
	Factor                              *Decimal            `json:"factor,omitempty"`
	FactorPrimitiveElement              *primitiveElement   `json:"_factor,omitempty"`
	Points                              *Decimal            `json:"points,omitempty"`
	PointsPrimitiveElement              *primitiveElement   `json:"_points,omitempty"`
	Net                                 *Money              `json:"net,omitempty"`
	Payment                             *String             `json:"payment,omitempty"`
	PaymentPrimitiveElement             *primitiveElement   `json:"_payment,omitempty"`
	PaymentDate                         *DateTime           `json:"paymentDate,omitempty"`
	PaymentDatePrimitiveElement         *primitiveElement   `json:"_paymentDate,omitempty"`
	Responsible                         *Reference          `json:"responsible,omitempty"`
	Recipient                           *Reference          `json:"recipient,omitempty"`
	LinkId                              []String            `json:"linkId,omitempty"`
	LinkIdPrimitiveElement              []*primitiveElement `json:"_linkId,omitempty"`
	SecurityLabelNumber                 []UnsignedInt       `json:"securityLabelNumber,omitempty"`
	SecurityLabelNumberPrimitiveElement []*primitiveElement `json:"_securityLabelNumber,omitempty"`
}

func (r ContractTermAssetValuedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermAssetValuedItem) marshalJSON() jsonContractTermAssetValuedItem {
	m := jsonContractTermAssetValuedItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Entity.(type) {
	case CodeableConcept:
		m.EntityCodeableConcept = &v
	case *CodeableConcept:
		m.EntityCodeableConcept = v
	case Reference:
		m.EntityReference = &v
	case *Reference:
		m.EntityReference = v
	}
	m.Identifier = r.Identifier
	if r.EffectiveTime != nil && r.EffectiveTime.Value != nil {
		m.EffectiveTime = r.EffectiveTime
	}
	if r.EffectiveTime != nil && (r.EffectiveTime.Id != nil || r.EffectiveTime.Extension != nil) {
		m.EffectiveTimePrimitiveElement = &primitiveElement{Id: r.EffectiveTime.Id, Extension: r.EffectiveTime.Extension}
	}
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	if r.Factor != nil && r.Factor.Value != nil {
		m.Factor = r.Factor
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	if r.Points != nil && r.Points.Value != nil {
		m.Points = r.Points
	}
	if r.Points != nil && (r.Points.Id != nil || r.Points.Extension != nil) {
		m.PointsPrimitiveElement = &primitiveElement{Id: r.Points.Id, Extension: r.Points.Extension}
	}
	m.Net = r.Net
	if r.Payment != nil && r.Payment.Value != nil {
		m.Payment = r.Payment
	}
	if r.Payment != nil && (r.Payment.Id != nil || r.Payment.Extension != nil) {
		m.PaymentPrimitiveElement = &primitiveElement{Id: r.Payment.Id, Extension: r.Payment.Extension}
	}
	if r.PaymentDate != nil && r.PaymentDate.Value != nil {
		m.PaymentDate = r.PaymentDate
	}
	if r.PaymentDate != nil && (r.PaymentDate.Id != nil || r.PaymentDate.Extension != nil) {
		m.PaymentDatePrimitiveElement = &primitiveElement{Id: r.PaymentDate.Id, Extension: r.PaymentDate.Extension}
	}
	m.Responsible = r.Responsible
	m.Recipient = r.Recipient
	anyLinkIdValue := false
	for _, e := range r.LinkId {
		if e.Value != nil {
			anyLinkIdValue = true
			break
		}
	}
	if anyLinkIdValue {
		m.LinkId = r.LinkId
	}
	anyLinkIdIdOrExtension := false
	for _, e := range r.LinkId {
		if e.Id != nil || e.Extension != nil {
			anyLinkIdIdOrExtension = true
			break
		}
	}
	if anyLinkIdIdOrExtension {
		m.LinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.LinkId))
		for _, e := range r.LinkId {
			if e.Id != nil || e.Extension != nil {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, nil)
			}
		}
	}
	anySecurityLabelNumberValue := false
	for _, e := range r.SecurityLabelNumber {
		if e.Value != nil {
			anySecurityLabelNumberValue = true
			break
		}
	}
	if anySecurityLabelNumberValue {
		m.SecurityLabelNumber = r.SecurityLabelNumber
	}
	anySecurityLabelNumberIdOrExtension := false
	for _, e := range r.SecurityLabelNumber {
		if e.Id != nil || e.Extension != nil {
			anySecurityLabelNumberIdOrExtension = true
			break
		}
	}
	if anySecurityLabelNumberIdOrExtension {
		m.SecurityLabelNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.SecurityLabelNumber))
		for _, e := range r.SecurityLabelNumber {
			if e.Id != nil || e.Extension != nil {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ContractTermAssetValuedItem) UnmarshalJSON(b []byte) error {
	var m jsonContractTermAssetValuedItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermAssetValuedItem) unmarshalJSON(m jsonContractTermAssetValuedItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.EntityCodeableConcept != nil {
		if r.Entity != nil {
			return fmt.Errorf("multiple values for field \"Entity\"")
		}
		v := m.EntityCodeableConcept
		r.Entity = v
	}
	if m.EntityReference != nil {
		if r.Entity != nil {
			return fmt.Errorf("multiple values for field \"Entity\"")
		}
		v := m.EntityReference
		r.Entity = v
	}
	r.Identifier = m.Identifier
	r.EffectiveTime = m.EffectiveTime
	if m.EffectiveTimePrimitiveElement != nil {
		if r.EffectiveTime == nil {
			r.EffectiveTime = &DateTime{}
		}
		r.EffectiveTime.Id = m.EffectiveTimePrimitiveElement.Id
		r.EffectiveTime.Extension = m.EffectiveTimePrimitiveElement.Extension
	}
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		if r.Factor == nil {
			r.Factor = &Decimal{}
		}
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Points = m.Points
	if m.PointsPrimitiveElement != nil {
		if r.Points == nil {
			r.Points = &Decimal{}
		}
		r.Points.Id = m.PointsPrimitiveElement.Id
		r.Points.Extension = m.PointsPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.Payment = m.Payment
	if m.PaymentPrimitiveElement != nil {
		if r.Payment == nil {
			r.Payment = &String{}
		}
		r.Payment.Id = m.PaymentPrimitiveElement.Id
		r.Payment.Extension = m.PaymentPrimitiveElement.Extension
	}
	r.PaymentDate = m.PaymentDate
	if m.PaymentDatePrimitiveElement != nil {
		if r.PaymentDate == nil {
			r.PaymentDate = &DateTime{}
		}
		r.PaymentDate.Id = m.PaymentDatePrimitiveElement.Id
		r.PaymentDate.Extension = m.PaymentDatePrimitiveElement.Extension
	}
	r.Responsible = m.Responsible
	r.Recipient = m.Recipient
	r.LinkId = m.LinkId
	for i, e := range m.LinkIdPrimitiveElement {
		if len(r.LinkId) <= i {
			r.LinkId = append(r.LinkId, String{})
		}
		if e != nil {
			r.LinkId[i].Id = e.Id
			r.LinkId[i].Extension = e.Extension
		}
	}
	r.SecurityLabelNumber = m.SecurityLabelNumber
	for i, e := range m.SecurityLabelNumberPrimitiveElement {
		if len(r.SecurityLabelNumber) <= i {
			r.SecurityLabelNumber = append(r.SecurityLabelNumber, UnsignedInt{})
		}
		if e != nil {
			r.SecurityLabelNumber[i].Id = e.Id
			r.SecurityLabelNumber[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ContractTermAssetValuedItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Entity.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "entityCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "entityReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EffectiveTime, xml.StartElement{Name: xml.Name{Local: "effectiveTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Points, xml.StartElement{Name: xml.Name{Local: "points"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Payment, xml.StartElement{Name: xml.Name{Local: "payment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaymentDate, xml.StartElement{Name: xml.Name{Local: "paymentDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Responsible, xml.StartElement{Name: xml.Name{Local: "responsible"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recipient, xml.StartElement{Name: xml.Name{Local: "recipient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LinkId, xml.StartElement{Name: xml.Name{Local: "linkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabelNumber, xml.StartElement{Name: xml.Name{Local: "securityLabelNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermAssetValuedItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "entityCodeableConcept":
				if r.Entity != nil {
					return fmt.Errorf("multiple values for field \"Entity\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entity = v
			case "entityReference":
				if r.Entity != nil {
					return fmt.Errorf("multiple values for field \"Entity\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entity = v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "effectiveTime":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EffectiveTime = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "points":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Points = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "payment":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Payment = &v
			case "paymentDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaymentDate = &v
			case "responsible":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Responsible = &v
			case "recipient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recipient = &v
			case "linkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LinkId = append(r.LinkId, v)
			case "securityLabelNumber":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabelNumber = append(r.SecurityLabelNumber, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermAssetValuedItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
type ContractTermAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// True if the term prohibits the  action.
	DoNotPerform *Boolean
	// Activity or service obligation to be done or not done, performed or not performed, effectuated or not by this Contract term.
	Type CodeableConcept
	// Entity of the action.
	Subject []ContractTermActionSubject
	// Reason or purpose for the action stipulated by this Contract Provision.
	Intent CodeableConcept
	// Id [identifier??] of the clause or question text related to this action in the referenced form or QuestionnaireResponse.
	LinkId []String
	// Current state of the term action.
	Status CodeableConcept
	// Encounter or Episode with primary association to specified term activity.
	Context *Reference
	// Id [identifier??] of the clause or question text related to the requester of this action in the referenced form or QuestionnaireResponse.
	ContextLinkId []String
	// When action happens.
	Occurrence isContractTermActionOccurrence
	// Who or what initiated the action and has responsibility for its activation.
	Requester []Reference
	// Id [identifier??] of the clause or question text related to the requester of this action in the referenced form or QuestionnaireResponse.
	RequesterLinkId []String
	// The type of individual that is desired or required to perform or not perform the action.
	PerformerType []CodeableConcept
	// The type of role or competency of an individual desired or required to perform or not perform the action.
	PerformerRole *CodeableConcept
	// Indicates who or what is being asked to perform (or not perform) the ction.
	Performer *Reference
	// Id [identifier??] of the clause or question text related to the reason type or reference of this  action in the referenced form or QuestionnaireResponse.
	PerformerLinkId []String
	// Rationale for the action to be performed or not performed. Describes why the action is permitted or prohibited.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies permitting or not permitting this action.
	ReasonReference []Reference
	// Describes why the action is to be performed or not performed in textual form.
	Reason []String
	// Id [identifier??] of the clause or question text related to the reason type or reference of this  action in the referenced form or QuestionnaireResponse.
	ReasonLinkId []String
	// Comments made about the term action made by the requester, performer, subject or other participants.
	Note []Annotation
	// Security labels that protects the action.
	SecurityLabelNumber []UnsignedInt
}
type isContractTermActionOccurrence interface {
	isContractTermActionOccurrence()
}

func (r DateTime) isContractTermActionOccurrence() {}
func (r Period) isContractTermActionOccurrence()   {}
func (r Timing) isContractTermActionOccurrence()   {}

type jsonContractTermAction struct {
	Id                                  *string                     `json:"id,omitempty"`
	Extension                           []Extension                 `json:"extension,omitempty"`
	ModifierExtension                   []Extension                 `json:"modifierExtension,omitempty"`
	DoNotPerform                        *Boolean                    `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement        *primitiveElement           `json:"_doNotPerform,omitempty"`
	Type                                CodeableConcept             `json:"type,omitempty"`
	Subject                             []ContractTermActionSubject `json:"subject,omitempty"`
	Intent                              CodeableConcept             `json:"intent,omitempty"`
	LinkId                              []String                    `json:"linkId,omitempty"`
	LinkIdPrimitiveElement              []*primitiveElement         `json:"_linkId,omitempty"`
	Status                              CodeableConcept             `json:"status,omitempty"`
	Context                             *Reference                  `json:"context,omitempty"`
	ContextLinkId                       []String                    `json:"contextLinkId,omitempty"`
	ContextLinkIdPrimitiveElement       []*primitiveElement         `json:"_contextLinkId,omitempty"`
	OccurrenceDateTime                  *DateTime                   `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement  *primitiveElement           `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                    *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                    *Timing                     `json:"occurrenceTiming,omitempty"`
	Requester                           []Reference                 `json:"requester,omitempty"`
	RequesterLinkId                     []String                    `json:"requesterLinkId,omitempty"`
	RequesterLinkIdPrimitiveElement     []*primitiveElement         `json:"_requesterLinkId,omitempty"`
	PerformerType                       []CodeableConcept           `json:"performerType,omitempty"`
	PerformerRole                       *CodeableConcept            `json:"performerRole,omitempty"`
	Performer                           *Reference                  `json:"performer,omitempty"`
	PerformerLinkId                     []String                    `json:"performerLinkId,omitempty"`
	PerformerLinkIdPrimitiveElement     []*primitiveElement         `json:"_performerLinkId,omitempty"`
	ReasonCode                          []CodeableConcept           `json:"reasonCode,omitempty"`
	ReasonReference                     []Reference                 `json:"reasonReference,omitempty"`
	Reason                              []String                    `json:"reason,omitempty"`
	ReasonPrimitiveElement              []*primitiveElement         `json:"_reason,omitempty"`
	ReasonLinkId                        []String                    `json:"reasonLinkId,omitempty"`
	ReasonLinkIdPrimitiveElement        []*primitiveElement         `json:"_reasonLinkId,omitempty"`
	Note                                []Annotation                `json:"note,omitempty"`
	SecurityLabelNumber                 []UnsignedInt               `json:"securityLabelNumber,omitempty"`
	SecurityLabelNumberPrimitiveElement []*primitiveElement         `json:"_securityLabelNumber,omitempty"`
}

func (r ContractTermAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermAction) marshalJSON() jsonContractTermAction {
	m := jsonContractTermAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.DoNotPerform != nil && r.DoNotPerform.Value != nil {
		m.DoNotPerform = r.DoNotPerform
	}
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	m.Type = r.Type
	m.Subject = r.Subject
	m.Intent = r.Intent
	anyLinkIdValue := false
	for _, e := range r.LinkId {
		if e.Value != nil {
			anyLinkIdValue = true
			break
		}
	}
	if anyLinkIdValue {
		m.LinkId = r.LinkId
	}
	anyLinkIdIdOrExtension := false
	for _, e := range r.LinkId {
		if e.Id != nil || e.Extension != nil {
			anyLinkIdIdOrExtension = true
			break
		}
	}
	if anyLinkIdIdOrExtension {
		m.LinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.LinkId))
		for _, e := range r.LinkId {
			if e.Id != nil || e.Extension != nil {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LinkIdPrimitiveElement = append(m.LinkIdPrimitiveElement, nil)
			}
		}
	}
	m.Status = r.Status
	m.Context = r.Context
	anyContextLinkIdValue := false
	for _, e := range r.ContextLinkId {
		if e.Value != nil {
			anyContextLinkIdValue = true
			break
		}
	}
	if anyContextLinkIdValue {
		m.ContextLinkId = r.ContextLinkId
	}
	anyContextLinkIdIdOrExtension := false
	for _, e := range r.ContextLinkId {
		if e.Id != nil || e.Extension != nil {
			anyContextLinkIdIdOrExtension = true
			break
		}
	}
	if anyContextLinkIdIdOrExtension {
		m.ContextLinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.ContextLinkId))
		for _, e := range r.ContextLinkId {
			if e.Id != nil || e.Extension != nil {
				m.ContextLinkIdPrimitiveElement = append(m.ContextLinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ContextLinkIdPrimitiveElement = append(m.ContextLinkIdPrimitiveElement, nil)
			}
		}
	}
	switch v := r.Occurrence.(type) {
	case DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.OccurrencePeriod = &v
	case *Period:
		m.OccurrencePeriod = v
	case Timing:
		m.OccurrenceTiming = &v
	case *Timing:
		m.OccurrenceTiming = v
	}
	m.Requester = r.Requester
	anyRequesterLinkIdValue := false
	for _, e := range r.RequesterLinkId {
		if e.Value != nil {
			anyRequesterLinkIdValue = true
			break
		}
	}
	if anyRequesterLinkIdValue {
		m.RequesterLinkId = r.RequesterLinkId
	}
	anyRequesterLinkIdIdOrExtension := false
	for _, e := range r.RequesterLinkId {
		if e.Id != nil || e.Extension != nil {
			anyRequesterLinkIdIdOrExtension = true
			break
		}
	}
	if anyRequesterLinkIdIdOrExtension {
		m.RequesterLinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.RequesterLinkId))
		for _, e := range r.RequesterLinkId {
			if e.Id != nil || e.Extension != nil {
				m.RequesterLinkIdPrimitiveElement = append(m.RequesterLinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.RequesterLinkIdPrimitiveElement = append(m.RequesterLinkIdPrimitiveElement, nil)
			}
		}
	}
	m.PerformerType = r.PerformerType
	m.PerformerRole = r.PerformerRole
	m.Performer = r.Performer
	anyPerformerLinkIdValue := false
	for _, e := range r.PerformerLinkId {
		if e.Value != nil {
			anyPerformerLinkIdValue = true
			break
		}
	}
	if anyPerformerLinkIdValue {
		m.PerformerLinkId = r.PerformerLinkId
	}
	anyPerformerLinkIdIdOrExtension := false
	for _, e := range r.PerformerLinkId {
		if e.Id != nil || e.Extension != nil {
			anyPerformerLinkIdIdOrExtension = true
			break
		}
	}
	if anyPerformerLinkIdIdOrExtension {
		m.PerformerLinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.PerformerLinkId))
		for _, e := range r.PerformerLinkId {
			if e.Id != nil || e.Extension != nil {
				m.PerformerLinkIdPrimitiveElement = append(m.PerformerLinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PerformerLinkIdPrimitiveElement = append(m.PerformerLinkIdPrimitiveElement, nil)
			}
		}
	}
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	anyReasonValue := false
	for _, e := range r.Reason {
		if e.Value != nil {
			anyReasonValue = true
			break
		}
	}
	if anyReasonValue {
		m.Reason = r.Reason
	}
	anyReasonIdOrExtension := false
	for _, e := range r.Reason {
		if e.Id != nil || e.Extension != nil {
			anyReasonIdOrExtension = true
			break
		}
	}
	if anyReasonIdOrExtension {
		m.ReasonPrimitiveElement = make([]*primitiveElement, 0, len(r.Reason))
		for _, e := range r.Reason {
			if e.Id != nil || e.Extension != nil {
				m.ReasonPrimitiveElement = append(m.ReasonPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ReasonPrimitiveElement = append(m.ReasonPrimitiveElement, nil)
			}
		}
	}
	anyReasonLinkIdValue := false
	for _, e := range r.ReasonLinkId {
		if e.Value != nil {
			anyReasonLinkIdValue = true
			break
		}
	}
	if anyReasonLinkIdValue {
		m.ReasonLinkId = r.ReasonLinkId
	}
	anyReasonLinkIdIdOrExtension := false
	for _, e := range r.ReasonLinkId {
		if e.Id != nil || e.Extension != nil {
			anyReasonLinkIdIdOrExtension = true
			break
		}
	}
	if anyReasonLinkIdIdOrExtension {
		m.ReasonLinkIdPrimitiveElement = make([]*primitiveElement, 0, len(r.ReasonLinkId))
		for _, e := range r.ReasonLinkId {
			if e.Id != nil || e.Extension != nil {
				m.ReasonLinkIdPrimitiveElement = append(m.ReasonLinkIdPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ReasonLinkIdPrimitiveElement = append(m.ReasonLinkIdPrimitiveElement, nil)
			}
		}
	}
	m.Note = r.Note
	anySecurityLabelNumberValue := false
	for _, e := range r.SecurityLabelNumber {
		if e.Value != nil {
			anySecurityLabelNumberValue = true
			break
		}
	}
	if anySecurityLabelNumberValue {
		m.SecurityLabelNumber = r.SecurityLabelNumber
	}
	anySecurityLabelNumberIdOrExtension := false
	for _, e := range r.SecurityLabelNumber {
		if e.Id != nil || e.Extension != nil {
			anySecurityLabelNumberIdOrExtension = true
			break
		}
	}
	if anySecurityLabelNumberIdOrExtension {
		m.SecurityLabelNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.SecurityLabelNumber))
		for _, e := range r.SecurityLabelNumber {
			if e.Id != nil || e.Extension != nil {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SecurityLabelNumberPrimitiveElement = append(m.SecurityLabelNumberPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ContractTermAction) UnmarshalJSON(b []byte) error {
	var m jsonContractTermAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermAction) unmarshalJSON(m jsonContractTermAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		if r.DoNotPerform == nil {
			r.DoNotPerform = &Boolean{}
		}
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Subject = m.Subject
	r.Intent = m.Intent
	r.LinkId = m.LinkId
	for i, e := range m.LinkIdPrimitiveElement {
		if len(r.LinkId) <= i {
			r.LinkId = append(r.LinkId, String{})
		}
		if e != nil {
			r.LinkId[i].Id = e.Id
			r.LinkId[i].Extension = e.Extension
		}
	}
	r.Status = m.Status
	r.Context = m.Context
	r.ContextLinkId = m.ContextLinkId
	for i, e := range m.ContextLinkIdPrimitiveElement {
		if len(r.ContextLinkId) <= i {
			r.ContextLinkId = append(r.ContextLinkId, String{})
		}
		if e != nil {
			r.ContextLinkId[i].Id = e.Id
			r.ContextLinkId[i].Extension = e.Extension
		}
	}
	if m.OccurrenceDateTime != nil || m.OccurrenceDateTimePrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceDateTime
		if m.OccurrenceDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurrenceDateTimePrimitiveElement.Id
			v.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	if m.OccurrencePeriod != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrencePeriod
		r.Occurrence = v
	}
	if m.OccurrenceTiming != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceTiming
		r.Occurrence = v
	}
	r.Requester = m.Requester
	r.RequesterLinkId = m.RequesterLinkId
	for i, e := range m.RequesterLinkIdPrimitiveElement {
		if len(r.RequesterLinkId) <= i {
			r.RequesterLinkId = append(r.RequesterLinkId, String{})
		}
		if e != nil {
			r.RequesterLinkId[i].Id = e.Id
			r.RequesterLinkId[i].Extension = e.Extension
		}
	}
	r.PerformerType = m.PerformerType
	r.PerformerRole = m.PerformerRole
	r.Performer = m.Performer
	r.PerformerLinkId = m.PerformerLinkId
	for i, e := range m.PerformerLinkIdPrimitiveElement {
		if len(r.PerformerLinkId) <= i {
			r.PerformerLinkId = append(r.PerformerLinkId, String{})
		}
		if e != nil {
			r.PerformerLinkId[i].Id = e.Id
			r.PerformerLinkId[i].Extension = e.Extension
		}
	}
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Reason = m.Reason
	for i, e := range m.ReasonPrimitiveElement {
		if len(r.Reason) <= i {
			r.Reason = append(r.Reason, String{})
		}
		if e != nil {
			r.Reason[i].Id = e.Id
			r.Reason[i].Extension = e.Extension
		}
	}
	r.ReasonLinkId = m.ReasonLinkId
	for i, e := range m.ReasonLinkIdPrimitiveElement {
		if len(r.ReasonLinkId) <= i {
			r.ReasonLinkId = append(r.ReasonLinkId, String{})
		}
		if e != nil {
			r.ReasonLinkId[i].Id = e.Id
			r.ReasonLinkId[i].Extension = e.Extension
		}
	}
	r.Note = m.Note
	r.SecurityLabelNumber = m.SecurityLabelNumber
	for i, e := range m.SecurityLabelNumberPrimitiveElement {
		if len(r.SecurityLabelNumber) <= i {
			r.SecurityLabelNumber = append(r.SecurityLabelNumber, UnsignedInt{})
		}
		if e != nil {
			r.SecurityLabelNumber[i].Id = e.Id
			r.SecurityLabelNumber[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ContractTermAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DoNotPerform, xml.StartElement{Name: xml.Name{Local: "doNotPerform"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Intent, xml.StartElement{Name: xml.Name{Local: "intent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LinkId, xml.StartElement{Name: xml.Name{Local: "linkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Context, xml.StartElement{Name: xml.Name{Local: "context"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContextLinkId, xml.StartElement{Name: xml.Name{Local: "contextLinkId"}})
	if err != nil {
		return err
	}
	switch v := r.Occurrence.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurrenceDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurrencePeriod"}})
		if err != nil {
			return err
		}
	case Timing, *Timing:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurrenceTiming"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Requester, xml.StartElement{Name: xml.Name{Local: "requester"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RequesterLinkId, xml.StartElement{Name: xml.Name{Local: "requesterLinkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PerformerType, xml.StartElement{Name: xml.Name{Local: "performerType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PerformerRole, xml.StartElement{Name: xml.Name{Local: "performerRole"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Performer, xml.StartElement{Name: xml.Name{Local: "performer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PerformerLinkId, xml.StartElement{Name: xml.Name{Local: "performerLinkId"}})
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
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonLinkId, xml.StartElement{Name: xml.Name{Local: "reasonLinkId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SecurityLabelNumber, xml.StartElement{Name: xml.Name{Local: "securityLabelNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "doNotPerform":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoNotPerform = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "subject":
				var v ContractTermActionSubject
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = append(r.Subject, v)
			case "intent":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Intent = v
			case "linkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LinkId = append(r.LinkId, v)
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "context":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Context = &v
			case "contextLinkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContextLinkId = append(r.ContextLinkId, v)
			case "occurrenceDateTime":
				if r.Occurrence != nil {
					return fmt.Errorf("multiple values for field \"Occurrence\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurrence = v
			case "occurrencePeriod":
				if r.Occurrence != nil {
					return fmt.Errorf("multiple values for field \"Occurrence\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurrence = v
			case "occurrenceTiming":
				if r.Occurrence != nil {
					return fmt.Errorf("multiple values for field \"Occurrence\"")
				}
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurrence = v
			case "requester":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requester = append(r.Requester, v)
			case "requesterLinkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RequesterLinkId = append(r.RequesterLinkId, v)
			case "performerType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PerformerType = append(r.PerformerType, v)
			case "performerRole":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PerformerRole = &v
			case "performer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Performer = &v
			case "performerLinkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PerformerLinkId = append(r.PerformerLinkId, v)
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
			case "reason":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			case "reasonLinkId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonLinkId = append(r.ReasonLinkId, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "securityLabelNumber":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SecurityLabelNumber = append(r.SecurityLabelNumber, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Entity of the action.
type ContractTermActionSubject struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The entity the action is performed or not performed on or for.
	Reference []Reference
	// Role type of agent assigned roles in this Contract.
	Role *CodeableConcept
}
type jsonContractTermActionSubject struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Reference         []Reference      `json:"reference,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

func (r ContractTermActionSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractTermActionSubject) marshalJSON() jsonContractTermActionSubject {
	m := jsonContractTermActionSubject{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Reference = r.Reference
	m.Role = r.Role
	return m
}
func (r *ContractTermActionSubject) UnmarshalJSON(b []byte) error {
	var m jsonContractTermActionSubject
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractTermActionSubject) unmarshalJSON(m jsonContractTermActionSubject) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Reference = m.Reference
	r.Role = m.Role
	return nil
}
func (r ContractTermActionSubject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ContractTermActionSubject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "reference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = append(r.Reference, v)
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractTermActionSubject) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
type ContractSigner struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Role of this Contract signer, e.g. notary, grantee.
	Type Coding
	// Party which is a signator to this Contract.
	Party Reference
	// Legally binding Contract DSIG signature contents in Base64.
	Signature []Signature
}
type jsonContractSigner struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type,omitempty"`
	Party             Reference   `json:"party,omitempty"`
	Signature         []Signature `json:"signature,omitempty"`
}

func (r ContractSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractSigner) marshalJSON() jsonContractSigner {
	m := jsonContractSigner{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Party = r.Party
	m.Signature = r.Signature
	return m
}
func (r *ContractSigner) UnmarshalJSON(b []byte) error {
	var m jsonContractSigner
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractSigner) unmarshalJSON(m jsonContractSigner) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Party = m.Party
	r.Signature = m.Signature
	return nil
}
func (r ContractSigner) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Party, xml.StartElement{Name: xml.Name{Local: "party"}})
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
func (r *ContractSigner) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "party":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Party = v
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
func (r ContractSigner) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The "patient friendly language" versionof the Contract in whole or in parts. "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
type ContractFriendly struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension and ensure understandability.
	Content isContractFriendlyContent
}
type isContractFriendlyContent interface {
	isContractFriendlyContent()
}

func (r Attachment) isContractFriendlyContent() {}
func (r Reference) isContractFriendlyContent()  {}

type jsonContractFriendly struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `json:"contentReference,omitempty"`
}

func (r ContractFriendly) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractFriendly) marshalJSON() jsonContractFriendly {
	m := jsonContractFriendly{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Content.(type) {
	case Attachment:
		m.ContentAttachment = &v
	case *Attachment:
		m.ContentAttachment = v
	case Reference:
		m.ContentReference = &v
	case *Reference:
		m.ContentReference = v
	}
	return m
}
func (r *ContractFriendly) UnmarshalJSON(b []byte) error {
	var m jsonContractFriendly
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractFriendly) unmarshalJSON(m jsonContractFriendly) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ContentAttachment != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentAttachment
		r.Content = v
	}
	if m.ContentReference != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentReference
		r.Content = v
	}
	return nil
}
func (r ContractFriendly) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Content.(type) {
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentReference"}})
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
func (r *ContractFriendly) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "contentAttachment":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			case "contentReference":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractFriendly) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// List of Legal expressions or representations of this Contract.
type ContractLegal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Contract legal text in human renderable form.
	Content isContractLegalContent
}
type isContractLegalContent interface {
	isContractLegalContent()
}

func (r Attachment) isContractLegalContent() {}
func (r Reference) isContractLegalContent()  {}

type jsonContractLegal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `json:"contentReference,omitempty"`
}

func (r ContractLegal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractLegal) marshalJSON() jsonContractLegal {
	m := jsonContractLegal{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Content.(type) {
	case Attachment:
		m.ContentAttachment = &v
	case *Attachment:
		m.ContentAttachment = v
	case Reference:
		m.ContentReference = &v
	case *Reference:
		m.ContentReference = v
	}
	return m
}
func (r *ContractLegal) UnmarshalJSON(b []byte) error {
	var m jsonContractLegal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractLegal) unmarshalJSON(m jsonContractLegal) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ContentAttachment != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentAttachment
		r.Content = v
	}
	if m.ContentReference != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentReference
		r.Content = v
	}
	return nil
}
func (r ContractLegal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Content.(type) {
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentReference"}})
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
func (r *ContractLegal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "contentAttachment":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			case "contentReference":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractLegal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// List of Computable Policy Rule Language Representations of this Contract.
type ContractRule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal).
	Content isContractRuleContent
}
type isContractRuleContent interface {
	isContractRuleContent()
}

func (r Attachment) isContractRuleContent() {}
func (r Reference) isContractRuleContent()  {}

type jsonContractRule struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `json:"contentReference,omitempty"`
}

func (r ContractRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ContractRule) marshalJSON() jsonContractRule {
	m := jsonContractRule{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Content.(type) {
	case Attachment:
		m.ContentAttachment = &v
	case *Attachment:
		m.ContentAttachment = v
	case Reference:
		m.ContentReference = &v
	case *Reference:
		m.ContentReference = v
	}
	return m
}
func (r *ContractRule) UnmarshalJSON(b []byte) error {
	var m jsonContractRule
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ContractRule) unmarshalJSON(m jsonContractRule) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ContentAttachment != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentAttachment
		r.Content = v
	}
	if m.ContentReference != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentReference
		r.Content = v
	}
	return nil
}
func (r ContractRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Content.(type) {
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentReference"}})
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
func (r *ContractRule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "contentAttachment":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			case "contentReference":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ContractRule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
