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

// The Citation Resource enables reference to any knowledge artifact for purposes of identification and attribution. The Citation Resource supports existing reference structures and developing publication practices such as versioning, expressing complex contributorship roles, and referencing computable resources.
type Citation struct {
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
	// An absolute URI that is used to identify this citation when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this summary is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the summary is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this citation when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the citation when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the citation author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the citation. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the citation.
	Title *String
	// The status of this summary. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this citation is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the citation was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the citation changes.
	Date *DateTime
	// The name of the organization or individual that published the citation.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the citation from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate citation instances.
	UseContext []UsageContext
	// A legal or geographic region in which the citation is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this citation is needed and why it has been designed as it has.
	Purpose *Markdown
	// Use and/or publishing restrictions for the Citation, not for the cited artifact.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the citation content was or is planned to be in active use.
	EffectivePeriod *Period
	// Who authored the Citation.
	Author []ContactDetail
	// Who edited the Citation.
	Editor []ContactDetail
	// Who reviewed the Citation.
	Reviewer []ContactDetail
	// Who endorsed the Citation.
	Endorser []ContactDetail
	// A human-readable display of the citation.
	Summary []CitationSummary
	// The assignment to an organizing scheme.
	Classification []CitationClassification
	// Used for general notes and annotations not coded elsewhere.
	Note []Annotation
	// The status of the citation.
	CurrentState []CodeableConcept
	// An effective date or period for a status of the citation.
	StatusDate []CitationStatusDate
	// Artifact related to the Citation Resource.
	RelatesTo []CitationRelatesTo
	// The article or artifact being described.
	CitedArtifact *CitationCitedArtifact
}

// A human-readable display of the citation.
type CitationSummary struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Format for display of the citation.
	Style *CodeableConcept
	// The human-readable display of the citation.
	Text Markdown
}

// The assignment to an organizing scheme.
type CitationClassification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of classifier (e.g. publication type, keyword).
	Type *CodeableConcept
	// The specific classification value.
	Classifier []CodeableConcept
}

// An effective date or period for a status of the citation.
type CitationStatusDate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Classification of the status.
	Activity CodeableConcept
	// Either occurred or expected.
	Actual *Boolean
	// When the status started and/or ended.
	Period Period
}

// Artifact related to the Citation Resource.
type CitationRelatesTo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// How the Citation resource relates to the target artifact.
	RelationshipType CodeableConcept
	// The clasification of the related artifact.
	TargetClassifier []CodeableConcept
	// The article or artifact that the Citation Resource is related to.
	Target isCitationRelatesToTarget
}
type isCitationRelatesToTarget interface {
	model.Element
	isCitationRelatesToTarget()
}

func (r Uri) isCitationRelatesToTarget()        {}
func (r Identifier) isCitationRelatesToTarget() {}
func (r Reference) isCitationRelatesToTarget()  {}
func (r Attachment) isCitationRelatesToTarget() {}

// The article or artifact being described.
type CitationCitedArtifact struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A formal identifier that is used to identify this citation when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// A formal identifier that is used to identify things closely related to this citation.
	RelatedIdentifier []Identifier
	// When the cited artifact was accessed.
	DateAccessed *DateTime
	// The defined version of the cited artifact.
	Version *CitationCitedArtifactVersion
	// The status of the cited artifact.
	CurrentState []CodeableConcept
	// An effective date or period for a status of the cited artifact.
	StatusDate []CitationCitedArtifactStatusDate
	// The title details of the article or artifact.
	Title []CitationCitedArtifactTitle
	// Summary of the article or artifact.
	Abstract []CitationCitedArtifactAbstract
	// The component of the article or artifact.
	Part *CitationCitedArtifactPart
	// The artifact related to the cited artifact.
	RelatesTo []CitationCitedArtifactRelatesTo
	// If multiple, used to represent alternative forms of the article that are not separate citations.
	PublicationForm []CitationCitedArtifactPublicationForm
	// Used for any URL for the article or artifact cited.
	WebLocation []CitationCitedArtifactWebLocation
	// The assignment to an organizing scheme.
	Classification []CitationCitedArtifactClassification
	// This element is used to list authors and other contributors, their contact information, specific contributions, and summary statements.
	Contributorship *CitationCitedArtifactContributorship
	// Any additional information or content for the article or artifact.
	Note []Annotation
}

// The defined version of the cited artifact.
type CitationCitedArtifactVersion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The version number or other version identifier.
	Value String
	// Citation for the main version of the cited artifact.
	BaseCitation *Reference
}

// An effective date or period for a status of the cited artifact.
type CitationCitedArtifactStatusDate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Classification of the status.
	Activity CodeableConcept
	// Either occurred or expected.
	Actual *Boolean
	// When the status started and/or ended.
	Period Period
}

// The title details of the article or artifact.
type CitationCitedArtifactTitle struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Used to express the reason or specific aspect for the title.
	Type []CodeableConcept
	// Used to express the specific language.
	Language *CodeableConcept
	// The title of the article or artifact.
	Text Markdown
}

// Summary of the article or artifact.
type CitationCitedArtifactAbstract struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Used to express the reason or specific aspect for the abstract.
	Type *CodeableConcept
	// Used to express the specific language.
	Language *CodeableConcept
	// Abstract content.
	Text Markdown
	// Copyright notice for the abstract.
	Copyright *Markdown
}

// The component of the article or artifact.
type CitationCitedArtifactPart struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of component.
	Type *CodeableConcept
	// The specification of the component.
	Value *String
	// The citation for the full article or artifact.
	BaseCitation *Reference
}

// The artifact related to the cited artifact.
type CitationCitedArtifactRelatesTo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// How the cited artifact relates to the target artifact.
	RelationshipType CodeableConcept
	// The clasification of the related artifact.
	TargetClassifier []CodeableConcept
	// The article or artifact that the cited artifact is related to.
	Target isCitationCitedArtifactRelatesToTarget
}
type isCitationCitedArtifactRelatesToTarget interface {
	model.Element
	isCitationCitedArtifactRelatesToTarget()
}

func (r Uri) isCitationCitedArtifactRelatesToTarget()        {}
func (r Identifier) isCitationCitedArtifactRelatesToTarget() {}
func (r Reference) isCitationCitedArtifactRelatesToTarget()  {}
func (r Attachment) isCitationCitedArtifactRelatesToTarget() {}

// If multiple, used to represent alternative forms of the article that are not separate citations.
type CitationCitedArtifactPublicationForm struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The collection the cited article or artifact is published in.
	PublishedIn *CitationCitedArtifactPublicationFormPublishedIn
	// The specific issue in which the cited article resides.
	PeriodicRelease *CitationCitedArtifactPublicationFormPeriodicRelease
	// The date the article was added to the database, or the date the article was released (which may differ from the journal issue publication date).
	ArticleDate *DateTime
	// The date the article was last revised or updated in the database.
	LastRevisionDate *DateTime
	// Language in which this form of the article is published.
	Language []CodeableConcept
	// Entry number or identifier for inclusion in a database.
	AccessionNumber *String
	// Used for full display of pagination.
	PageString *String
	// Used for isolated representation of first page.
	FirstPage *String
	// Used for isolated representation of last page.
	LastPage *String
	// Actual or approximate number of pages or screens.
	PageCount *String
	// Copyright notice for the full article or artifact.
	Copyright *Markdown
}

// The collection the cited article or artifact is published in.
type CitationCitedArtifactPublicationFormPublishedIn struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Kind of container (e.g. Periodical, database, or book).
	Type *CodeableConcept
	// Journal identifiers include ISSN, ISO Abbreviation and NLMuniqueID; Book identifiers include ISBN.
	Identifier []Identifier
	// Name of the database or title of the book or journal.
	Title *String
	// Name of the publisher.
	Publisher *Reference
	// Geographic location of the publisher.
	PublisherLocation *String
}

// The specific issue in which the cited article resides.
type CitationCitedArtifactPublicationFormPeriodicRelease struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the form of the medium cited. Common codes are "Internet" or "Print".
	CitedMedium *CodeableConcept
	// Volume number of journal in which the article is published.
	Volume *String
	// Issue, part or supplement of journal in which the article is published.
	Issue *String
	// Defining the date on which the issue of the journal was published.
	DateOfPublication *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication
}

// Defining the date on which the issue of the journal was published.
type CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date on which the issue of the journal was published.
	Date *Date
	// Year on which the issue of the journal was published.
	Year *String
	// Month on which the issue of the journal was published.
	Month *String
	// Day on which the issue of the journal was published.
	Day *String
	// Spring, Summer, Fall/Autumn, Winter.
	Season *String
	// Text representation of the date of which the issue of the journal was published.
	Text *String
}

// Used for any URL for the article or artifact cited.
type CitationCitedArtifactWebLocation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code the reason for different URLs, e.g. abstract and full-text.
	Type *CodeableConcept
	// The specific URL.
	Url *Uri
}

// The assignment to an organizing scheme.
type CitationCitedArtifactClassification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of classifier (e.g. publication type, keyword).
	Type *CodeableConcept
	// The specific classification value.
	Classifier []CodeableConcept
	// Provenance and copyright of classification.
	WhoClassified *CitationCitedArtifactClassificationWhoClassified
}

// Provenance and copyright of classification.
type CitationCitedArtifactClassificationWhoClassified struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Person who created the classification.
	Person *Reference
	// Organization who created the classification.
	Organization *Reference
	// The publisher of the classification, not the publisher of the article or artifact being cited.
	Publisher *Reference
	// Rights management statement for the classification.
	ClassifierCopyright *String
	// Acceptable to re-use the classification.
	FreeToShare *Boolean
}

// This element is used to list authors and other contributors, their contact information, specific contributions, and summary statements.
type CitationCitedArtifactContributorship struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates if the list includes all authors and/or contributors.
	Complete *Boolean
	// An individual entity named in the author list or contributor list.
	Entry []CitationCitedArtifactContributorshipEntry
	// Used to record a display of the author/contributor list without separate coding for each list member.
	Summary []CitationCitedArtifactContributorshipSummary
}

// An individual entity named in the author list or contributor list.
type CitationCitedArtifactContributorshipEntry struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A name associated with the individual.
	Name *HumanName
	// Initials for forename.
	Initials *String
	// Used for collective or corporate name as an author.
	CollectiveName *String
	// Unique person identifier.
	Identifier []Identifier
	// Organization affiliated with the entity.
	AffiliationInfo []CitationCitedArtifactContributorshipEntryAffiliationInfo
	// Physical mailing address for the author or contributor.
	Address []Address
	// Email or telephone contact methods for the author or contributor.
	Telecom []ContactPoint
	// This element identifies the specific nature of an individual’s contribution with respect to the cited work.
	ContributionType []CodeableConcept
	// The role of the contributor (e.g. author, editor, reviewer).
	Role *CodeableConcept
	// Contributions with accounting for time or number.
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance
	// Indication of which contributor is the corresponding contributor for the role.
	CorrespondingContact *Boolean
	// Used to code order of authors.
	ListOrder *PositiveInt
}

// Organization affiliated with the entity.
type CitationCitedArtifactContributorshipEntryAffiliationInfo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Display for the organization.
	Affiliation *String
	// Role within the organization, such as professional title.
	Role *String
	// Identifier for the organization.
	Identifier []Identifier
}

// Contributions with accounting for time or number.
type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The specific contribution.
	Type CodeableConcept
	// The time that the contribution was made.
	Time *DateTime
}

// Used to record a display of the author/contributor list without separate coding for each list member.
type CitationCitedArtifactContributorshipSummary struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Used most commonly to express an author list or a contributorship statement.
	Type *CodeableConcept
	// The format for the display string.
	Style *CodeableConcept
	// Used to code the producer or rule for creating the display string.
	Source *CodeableConcept
	// The display string for the author list, contributor list, or contributorship statement.
	Value Markdown
}

func (r Citation) ResourceType() string {
	return "Citation"
}
func (r Citation) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Citation) MemSize() int {
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
	if r.Url != nil {
		s += r.Url.MemSize()
	}
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	if r.Version != nil {
		s += r.Version.MemSize()
	}
	if r.Name != nil {
		s += r.Name.MemSize()
	}
	if r.Title != nil {
		s += r.Title.MemSize()
	}
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	if r.Experimental != nil {
		s += r.Experimental.MemSize()
	}
	if r.Date != nil {
		s += r.Date.MemSize()
	}
	if r.Publisher != nil {
		s += r.Publisher.MemSize()
	}
	for _, i := range r.Contact {
		s += i.MemSize()
	}
	s += (cap(r.Contact) - len(r.Contact)) * int(unsafe.Sizeof(ContactDetail{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	for _, i := range r.UseContext {
		s += i.MemSize()
	}
	s += (cap(r.UseContext) - len(r.UseContext)) * int(unsafe.Sizeof(UsageContext{}))
	for _, i := range r.Jurisdiction {
		s += i.MemSize()
	}
	s += (cap(r.Jurisdiction) - len(r.Jurisdiction)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Purpose != nil {
		s += r.Purpose.MemSize()
	}
	if r.Copyright != nil {
		s += r.Copyright.MemSize()
	}
	if r.ApprovalDate != nil {
		s += r.ApprovalDate.MemSize()
	}
	if r.LastReviewDate != nil {
		s += r.LastReviewDate.MemSize()
	}
	if r.EffectivePeriod != nil {
		s += r.EffectivePeriod.MemSize()
	}
	for _, i := range r.Author {
		s += i.MemSize()
	}
	s += (cap(r.Author) - len(r.Author)) * int(unsafe.Sizeof(ContactDetail{}))
	for _, i := range r.Editor {
		s += i.MemSize()
	}
	s += (cap(r.Editor) - len(r.Editor)) * int(unsafe.Sizeof(ContactDetail{}))
	for _, i := range r.Reviewer {
		s += i.MemSize()
	}
	s += (cap(r.Reviewer) - len(r.Reviewer)) * int(unsafe.Sizeof(ContactDetail{}))
	for _, i := range r.Endorser {
		s += i.MemSize()
	}
	s += (cap(r.Endorser) - len(r.Endorser)) * int(unsafe.Sizeof(ContactDetail{}))
	for _, i := range r.Summary {
		s += i.MemSize()
	}
	s += (cap(r.Summary) - len(r.Summary)) * int(unsafe.Sizeof(CitationSummary{}))
	for _, i := range r.Classification {
		s += i.MemSize()
	}
	s += (cap(r.Classification) - len(r.Classification)) * int(unsafe.Sizeof(CitationClassification{}))
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	for _, i := range r.CurrentState {
		s += i.MemSize()
	}
	s += (cap(r.CurrentState) - len(r.CurrentState)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.StatusDate {
		s += i.MemSize()
	}
	s += (cap(r.StatusDate) - len(r.StatusDate)) * int(unsafe.Sizeof(CitationStatusDate{}))
	for _, i := range r.RelatesTo {
		s += i.MemSize()
	}
	s += (cap(r.RelatesTo) - len(r.RelatesTo)) * int(unsafe.Sizeof(CitationRelatesTo{}))
	if r.CitedArtifact != nil {
		s += r.CitedArtifact.MemSize()
	}
	return s
}
func (r CitationSummary) MemSize() int {
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
	if r.Style != nil {
		s += r.Style.MemSize()
	}
	s += r.Text.MemSize() - int(unsafe.Sizeof(r.Text))
	return s
}
func (r CitationClassification) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	for _, i := range r.Classifier {
		s += i.MemSize()
	}
	s += (cap(r.Classifier) - len(r.Classifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	return s
}
func (r CitationStatusDate) MemSize() int {
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
	s += r.Activity.MemSize() - int(unsafe.Sizeof(r.Activity))
	if r.Actual != nil {
		s += r.Actual.MemSize()
	}
	s += r.Period.MemSize() - int(unsafe.Sizeof(r.Period))
	return s
}
func (r CitationRelatesTo) MemSize() int {
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
	s += r.RelationshipType.MemSize() - int(unsafe.Sizeof(r.RelationshipType))
	for _, i := range r.TargetClassifier {
		s += i.MemSize()
	}
	s += (cap(r.TargetClassifier) - len(r.TargetClassifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Target != nil {
		s += r.Target.MemSize()
	}
	return s
}
func (r CitationCitedArtifact) MemSize() int {
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
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.RelatedIdentifier {
		s += i.MemSize()
	}
	s += (cap(r.RelatedIdentifier) - len(r.RelatedIdentifier)) * int(unsafe.Sizeof(Identifier{}))
	if r.DateAccessed != nil {
		s += r.DateAccessed.MemSize()
	}
	if r.Version != nil {
		s += r.Version.MemSize()
	}
	for _, i := range r.CurrentState {
		s += i.MemSize()
	}
	s += (cap(r.CurrentState) - len(r.CurrentState)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.StatusDate {
		s += i.MemSize()
	}
	s += (cap(r.StatusDate) - len(r.StatusDate)) * int(unsafe.Sizeof(CitationCitedArtifactStatusDate{}))
	for _, i := range r.Title {
		s += i.MemSize()
	}
	s += (cap(r.Title) - len(r.Title)) * int(unsafe.Sizeof(CitationCitedArtifactTitle{}))
	for _, i := range r.Abstract {
		s += i.MemSize()
	}
	s += (cap(r.Abstract) - len(r.Abstract)) * int(unsafe.Sizeof(CitationCitedArtifactAbstract{}))
	if r.Part != nil {
		s += r.Part.MemSize()
	}
	for _, i := range r.RelatesTo {
		s += i.MemSize()
	}
	s += (cap(r.RelatesTo) - len(r.RelatesTo)) * int(unsafe.Sizeof(CitationCitedArtifactRelatesTo{}))
	for _, i := range r.PublicationForm {
		s += i.MemSize()
	}
	s += (cap(r.PublicationForm) - len(r.PublicationForm)) * int(unsafe.Sizeof(CitationCitedArtifactPublicationForm{}))
	for _, i := range r.WebLocation {
		s += i.MemSize()
	}
	s += (cap(r.WebLocation) - len(r.WebLocation)) * int(unsafe.Sizeof(CitationCitedArtifactWebLocation{}))
	for _, i := range r.Classification {
		s += i.MemSize()
	}
	s += (cap(r.Classification) - len(r.Classification)) * int(unsafe.Sizeof(CitationCitedArtifactClassification{}))
	if r.Contributorship != nil {
		s += r.Contributorship.MemSize()
	}
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	return s
}
func (r CitationCitedArtifactVersion) MemSize() int {
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
	s += r.Value.MemSize() - int(unsafe.Sizeof(r.Value))
	if r.BaseCitation != nil {
		s += r.BaseCitation.MemSize()
	}
	return s
}
func (r CitationCitedArtifactStatusDate) MemSize() int {
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
	s += r.Activity.MemSize() - int(unsafe.Sizeof(r.Activity))
	if r.Actual != nil {
		s += r.Actual.MemSize()
	}
	s += r.Period.MemSize() - int(unsafe.Sizeof(r.Period))
	return s
}
func (r CitationCitedArtifactTitle) MemSize() int {
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
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	s += r.Text.MemSize() - int(unsafe.Sizeof(r.Text))
	return s
}
func (r CitationCitedArtifactAbstract) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	s += r.Text.MemSize() - int(unsafe.Sizeof(r.Text))
	if r.Copyright != nil {
		s += r.Copyright.MemSize()
	}
	return s
}
func (r CitationCitedArtifactPart) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.Value != nil {
		s += r.Value.MemSize()
	}
	if r.BaseCitation != nil {
		s += r.BaseCitation.MemSize()
	}
	return s
}
func (r CitationCitedArtifactRelatesTo) MemSize() int {
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
	s += r.RelationshipType.MemSize() - int(unsafe.Sizeof(r.RelationshipType))
	for _, i := range r.TargetClassifier {
		s += i.MemSize()
	}
	s += (cap(r.TargetClassifier) - len(r.TargetClassifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Target != nil {
		s += r.Target.MemSize()
	}
	return s
}
func (r CitationCitedArtifactPublicationForm) MemSize() int {
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
	if r.PublishedIn != nil {
		s += r.PublishedIn.MemSize()
	}
	if r.PeriodicRelease != nil {
		s += r.PeriodicRelease.MemSize()
	}
	if r.ArticleDate != nil {
		s += r.ArticleDate.MemSize()
	}
	if r.LastRevisionDate != nil {
		s += r.LastRevisionDate.MemSize()
	}
	for _, i := range r.Language {
		s += i.MemSize()
	}
	s += (cap(r.Language) - len(r.Language)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.AccessionNumber != nil {
		s += r.AccessionNumber.MemSize()
	}
	if r.PageString != nil {
		s += r.PageString.MemSize()
	}
	if r.FirstPage != nil {
		s += r.FirstPage.MemSize()
	}
	if r.LastPage != nil {
		s += r.LastPage.MemSize()
	}
	if r.PageCount != nil {
		s += r.PageCount.MemSize()
	}
	if r.Copyright != nil {
		s += r.Copyright.MemSize()
	}
	return s
}
func (r CitationCitedArtifactPublicationFormPublishedIn) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	if r.Title != nil {
		s += r.Title.MemSize()
	}
	if r.Publisher != nil {
		s += r.Publisher.MemSize()
	}
	if r.PublisherLocation != nil {
		s += r.PublisherLocation.MemSize()
	}
	return s
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) MemSize() int {
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
	if r.CitedMedium != nil {
		s += r.CitedMedium.MemSize()
	}
	if r.Volume != nil {
		s += r.Volume.MemSize()
	}
	if r.Issue != nil {
		s += r.Issue.MemSize()
	}
	if r.DateOfPublication != nil {
		s += r.DateOfPublication.MemSize()
	}
	return s
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) MemSize() int {
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
	if r.Date != nil {
		s += r.Date.MemSize()
	}
	if r.Year != nil {
		s += r.Year.MemSize()
	}
	if r.Month != nil {
		s += r.Month.MemSize()
	}
	if r.Day != nil {
		s += r.Day.MemSize()
	}
	if r.Season != nil {
		s += r.Season.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	return s
}
func (r CitationCitedArtifactWebLocation) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.Url != nil {
		s += r.Url.MemSize()
	}
	return s
}
func (r CitationCitedArtifactClassification) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	for _, i := range r.Classifier {
		s += i.MemSize()
	}
	s += (cap(r.Classifier) - len(r.Classifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.WhoClassified != nil {
		s += r.WhoClassified.MemSize()
	}
	return s
}
func (r CitationCitedArtifactClassificationWhoClassified) MemSize() int {
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
	if r.Person != nil {
		s += r.Person.MemSize()
	}
	if r.Organization != nil {
		s += r.Organization.MemSize()
	}
	if r.Publisher != nil {
		s += r.Publisher.MemSize()
	}
	if r.ClassifierCopyright != nil {
		s += r.ClassifierCopyright.MemSize()
	}
	if r.FreeToShare != nil {
		s += r.FreeToShare.MemSize()
	}
	return s
}
func (r CitationCitedArtifactContributorship) MemSize() int {
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
	if r.Complete != nil {
		s += r.Complete.MemSize()
	}
	for _, i := range r.Entry {
		s += i.MemSize()
	}
	s += (cap(r.Entry) - len(r.Entry)) * int(unsafe.Sizeof(CitationCitedArtifactContributorshipEntry{}))
	for _, i := range r.Summary {
		s += i.MemSize()
	}
	s += (cap(r.Summary) - len(r.Summary)) * int(unsafe.Sizeof(CitationCitedArtifactContributorshipSummary{}))
	return s
}
func (r CitationCitedArtifactContributorshipEntry) MemSize() int {
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
	if r.Name != nil {
		s += r.Name.MemSize()
	}
	if r.Initials != nil {
		s += r.Initials.MemSize()
	}
	if r.CollectiveName != nil {
		s += r.CollectiveName.MemSize()
	}
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.AffiliationInfo {
		s += i.MemSize()
	}
	s += (cap(r.AffiliationInfo) - len(r.AffiliationInfo)) * int(unsafe.Sizeof(CitationCitedArtifactContributorshipEntryAffiliationInfo{}))
	for _, i := range r.Address {
		s += i.MemSize()
	}
	s += (cap(r.Address) - len(r.Address)) * int(unsafe.Sizeof(Address{}))
	for _, i := range r.Telecom {
		s += i.MemSize()
	}
	s += (cap(r.Telecom) - len(r.Telecom)) * int(unsafe.Sizeof(ContactPoint{}))
	for _, i := range r.ContributionType {
		s += i.MemSize()
	}
	s += (cap(r.ContributionType) - len(r.ContributionType)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Role != nil {
		s += r.Role.MemSize()
	}
	for _, i := range r.ContributionInstance {
		s += i.MemSize()
	}
	s += (cap(r.ContributionInstance) - len(r.ContributionInstance)) * int(unsafe.Sizeof(CitationCitedArtifactContributorshipEntryContributionInstance{}))
	if r.CorrespondingContact != nil {
		s += r.CorrespondingContact.MemSize()
	}
	if r.ListOrder != nil {
		s += r.ListOrder.MemSize()
	}
	return s
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) MemSize() int {
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
	if r.Affiliation != nil {
		s += r.Affiliation.MemSize()
	}
	if r.Role != nil {
		s += r.Role.MemSize()
	}
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	return s
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) MemSize() int {
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
	if r.Time != nil {
		s += r.Time.MemSize()
	}
	return s
}
func (r CitationCitedArtifactContributorshipSummary) MemSize() int {
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
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.Style != nil {
		s += r.Style.MemSize()
	}
	if r.Source != nil {
		s += r.Source.MemSize()
	}
	s += r.Value.MemSize() - int(unsafe.Sizeof(r.Value))
	return s
}
func (r Citation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationSummary) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationClassification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationStatusDate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationRelatesTo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactVersion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactStatusDate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactTitle) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactAbstract) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactPart) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactRelatesTo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactPublicationForm) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactPublicationFormPublishedIn) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactWebLocation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactClassification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactClassificationWhoClassified) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactContributorship) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactContributorshipEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CitationCitedArtifactContributorshipSummary) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Citation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Citation) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Citation\""))
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
	if r.Url != nil && r.Url.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	if r.Version != nil && r.Version.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"version\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Version)
		if err != nil {
			return err
		}
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		p := primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_version\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && r.Name.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Name)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		p := primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_name\":"))
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
	if r.Experimental != nil && r.Experimental.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"experimental\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Experimental)
		if err != nil {
			return err
		}
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		p := primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_experimental\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && r.Date.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"date\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Date)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		p := primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_date\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publisher\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Publisher)
		if err != nil {
			return err
		}
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		p := primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_publisher\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contact) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contact\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Contact {
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
	if len(r.UseContext) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"useContext\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.UseContext {
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
	if len(r.Jurisdiction) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"jurisdiction\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Jurisdiction {
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
	if r.Purpose != nil && r.Purpose.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"purpose\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Purpose)
		if err != nil {
			return err
		}
	}
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		p := primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_purpose\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"copyright\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Copyright)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		p := primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_copyright\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ApprovalDate != nil && r.ApprovalDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"approvalDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ApprovalDate)
		if err != nil {
			return err
		}
	}
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		p := primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_approvalDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastReviewDate != nil && r.LastReviewDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastReviewDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastReviewDate)
		if err != nil {
			return err
		}
	}
	if r.LastReviewDate != nil && (r.LastReviewDate.Id != nil || r.LastReviewDate.Extension != nil) {
		p := primitiveElement{Id: r.LastReviewDate.Id, Extension: r.LastReviewDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastReviewDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.EffectivePeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"effectivePeriod\":"))
		if err != nil {
			return err
		}
		err = r.EffectivePeriod.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Author) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Author {
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
	if len(r.Editor) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"editor\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Editor {
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
	if len(r.Reviewer) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewer\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Reviewer {
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
	if len(r.Endorser) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"endorser\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Endorser {
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
	if len(r.Summary) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"summary\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Summary {
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
	if len(r.Classification) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"classification\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Classification {
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
	if len(r.CurrentState) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"currentState\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.CurrentState {
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
	if len(r.StatusDate) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"statusDate\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.StatusDate {
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
	if len(r.RelatesTo) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relatesTo\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RelatesTo {
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
	if r.CitedArtifact != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"citedArtifact\":"))
		if err != nil {
			return err
		}
		err = r.CitedArtifact.marshalJSON(w)
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
func (r CitationSummary) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationSummary) marshalJSON(w io.Writer) error {
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
	if r.Style != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"style\":"))
		if err != nil {
			return err
		}
		err = r.Style.marshalJSON(w)
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
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationClassification) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationClassification) marshalJSON(w io.Writer) error {
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
	if len(r.Classifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"classifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Classifier {
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
func (r CitationStatusDate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationStatusDate) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"activity\":"))
	if err != nil {
		return err
	}
	err = r.Activity.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Actual != nil && r.Actual.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"actual\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Actual)
		if err != nil {
			return err
		}
	}
	if r.Actual != nil && (r.Actual.Id != nil || r.Actual.Extension != nil) {
		p := primitiveElement{Id: r.Actual.Id, Extension: r.Actual.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_actual\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	_, err = w.Write([]byte("\"period\":"))
	if err != nil {
		return err
	}
	err = r.Period.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CitationRelatesTo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationRelatesTo) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"relationshipType\":"))
	if err != nil {
		return err
	}
	err = r.RelationshipType.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.TargetClassifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetClassifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TargetClassifier {
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
	switch v := r.Target.(type) {
	case Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"targetUri\":"))
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
			_, err = w.Write([]byte("\"_targetUri\":"))
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
			_, err = w.Write([]byte("\"targetUri\":"))
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
			_, err = w.Write([]byte("\"_targetUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetIdentifier\":"))
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
		_, err = w.Write([]byte("\"targetIdentifier\":"))
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
		_, err = w.Write([]byte("\"targetReference\":"))
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
		_, err = w.Write([]byte("\"targetReference\":"))
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
		_, err = w.Write([]byte("\"targetAttachment\":"))
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
		_, err = w.Write([]byte("\"targetAttachment\":"))
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
func (r CitationCitedArtifact) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifact) marshalJSON(w io.Writer) error {
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
	if len(r.RelatedIdentifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relatedIdentifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RelatedIdentifier {
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
	if r.DateAccessed != nil && r.DateAccessed.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dateAccessed\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DateAccessed)
		if err != nil {
			return err
		}
	}
	if r.DateAccessed != nil && (r.DateAccessed.Id != nil || r.DateAccessed.Extension != nil) {
		p := primitiveElement{Id: r.DateAccessed.Id, Extension: r.DateAccessed.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_dateAccessed\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Version != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"version\":"))
		if err != nil {
			return err
		}
		err = r.Version.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.CurrentState) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"currentState\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.CurrentState {
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
	if len(r.StatusDate) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"statusDate\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.StatusDate {
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
	if len(r.Title) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Title {
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
	if len(r.Abstract) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"abstract\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Abstract {
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
	if r.Part != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"part\":"))
		if err != nil {
			return err
		}
		err = r.Part.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.RelatesTo) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relatesTo\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.RelatesTo {
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
	if len(r.PublicationForm) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publicationForm\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PublicationForm {
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
	if len(r.WebLocation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"webLocation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.WebLocation {
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
	if len(r.Classification) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"classification\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Classification {
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
	if r.Contributorship != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contributorship\":"))
		if err != nil {
			return err
		}
		err = r.Contributorship.marshalJSON(w)
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactVersion) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactVersion) marshalJSON(w io.Writer) error {
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
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Value)
		if err != nil {
			return err
		}
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		p := primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_value\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.BaseCitation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"baseCitation\":"))
		if err != nil {
			return err
		}
		err = r.BaseCitation.marshalJSON(w)
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
func (r CitationCitedArtifactStatusDate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactStatusDate) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"activity\":"))
	if err != nil {
		return err
	}
	err = r.Activity.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Actual != nil && r.Actual.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"actual\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Actual)
		if err != nil {
			return err
		}
	}
	if r.Actual != nil && (r.Actual.Id != nil || r.Actual.Extension != nil) {
		p := primitiveElement{Id: r.Actual.Id, Extension: r.Actual.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_actual\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	_, err = w.Write([]byte("\"period\":"))
	if err != nil {
		return err
	}
	err = r.Period.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactTitle) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactTitle) marshalJSON(w io.Writer) error {
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
	if len(r.Type) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
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
	if r.Language != nil {
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
		err = r.Language.marshalJSON(w)
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
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactAbstract) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactAbstract) marshalJSON(w io.Writer) error {
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
	if r.Language != nil {
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
		err = r.Language.marshalJSON(w)
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
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"copyright\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Copyright)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		p := primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_copyright\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactPart) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactPart) marshalJSON(w io.Writer) error {
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
	if r.Value != nil && r.Value.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Value)
		if err != nil {
			return err
		}
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		p := primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_value\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.BaseCitation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"baseCitation\":"))
		if err != nil {
			return err
		}
		err = r.BaseCitation.marshalJSON(w)
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
func (r CitationCitedArtifactRelatesTo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactRelatesTo) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"relationshipType\":"))
	if err != nil {
		return err
	}
	err = r.RelationshipType.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.TargetClassifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetClassifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TargetClassifier {
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
	switch v := r.Target.(type) {
	case Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"targetUri\":"))
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
			_, err = w.Write([]byte("\"_targetUri\":"))
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
			_, err = w.Write([]byte("\"targetUri\":"))
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
			_, err = w.Write([]byte("\"_targetUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetIdentifier\":"))
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
		_, err = w.Write([]byte("\"targetIdentifier\":"))
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
		_, err = w.Write([]byte("\"targetReference\":"))
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
		_, err = w.Write([]byte("\"targetReference\":"))
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
		_, err = w.Write([]byte("\"targetAttachment\":"))
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
		_, err = w.Write([]byte("\"targetAttachment\":"))
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
func (r CitationCitedArtifactPublicationForm) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactPublicationForm) marshalJSON(w io.Writer) error {
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
	if r.PublishedIn != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publishedIn\":"))
		if err != nil {
			return err
		}
		err = r.PublishedIn.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PeriodicRelease != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"periodicRelease\":"))
		if err != nil {
			return err
		}
		err = r.PeriodicRelease.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ArticleDate != nil && r.ArticleDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"articleDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ArticleDate)
		if err != nil {
			return err
		}
	}
	if r.ArticleDate != nil && (r.ArticleDate.Id != nil || r.ArticleDate.Extension != nil) {
		p := primitiveElement{Id: r.ArticleDate.Id, Extension: r.ArticleDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_articleDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastRevisionDate != nil && r.LastRevisionDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastRevisionDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastRevisionDate)
		if err != nil {
			return err
		}
	}
	if r.LastRevisionDate != nil && (r.LastRevisionDate.Id != nil || r.LastRevisionDate.Extension != nil) {
		p := primitiveElement{Id: r.LastRevisionDate.Id, Extension: r.LastRevisionDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastRevisionDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Language) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Language {
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
	if r.AccessionNumber != nil && r.AccessionNumber.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"accessionNumber\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AccessionNumber)
		if err != nil {
			return err
		}
	}
	if r.AccessionNumber != nil && (r.AccessionNumber.Id != nil || r.AccessionNumber.Extension != nil) {
		p := primitiveElement{Id: r.AccessionNumber.Id, Extension: r.AccessionNumber.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_accessionNumber\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PageString != nil && r.PageString.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"pageString\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PageString)
		if err != nil {
			return err
		}
	}
	if r.PageString != nil && (r.PageString.Id != nil || r.PageString.Extension != nil) {
		p := primitiveElement{Id: r.PageString.Id, Extension: r.PageString.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_pageString\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FirstPage != nil && r.FirstPage.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"firstPage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.FirstPage)
		if err != nil {
			return err
		}
	}
	if r.FirstPage != nil && (r.FirstPage.Id != nil || r.FirstPage.Extension != nil) {
		p := primitiveElement{Id: r.FirstPage.Id, Extension: r.FirstPage.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_firstPage\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastPage != nil && r.LastPage.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastPage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastPage)
		if err != nil {
			return err
		}
	}
	if r.LastPage != nil && (r.LastPage.Id != nil || r.LastPage.Extension != nil) {
		p := primitiveElement{Id: r.LastPage.Id, Extension: r.LastPage.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastPage\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PageCount != nil && r.PageCount.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"pageCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PageCount)
		if err != nil {
			return err
		}
	}
	if r.PageCount != nil && (r.PageCount.Id != nil || r.PageCount.Extension != nil) {
		p := primitiveElement{Id: r.PageCount.Id, Extension: r.PageCount.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_pageCount\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"copyright\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Copyright)
		if err != nil {
			return err
		}
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		p := primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_copyright\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactPublicationFormPublishedIn) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactPublicationFormPublishedIn) marshalJSON(w io.Writer) error {
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
	if r.Publisher != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publisher\":"))
		if err != nil {
			return err
		}
		err = r.Publisher.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PublisherLocation != nil && r.PublisherLocation.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publisherLocation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PublisherLocation)
		if err != nil {
			return err
		}
	}
	if r.PublisherLocation != nil && (r.PublisherLocation.Id != nil || r.PublisherLocation.Extension != nil) {
		p := primitiveElement{Id: r.PublisherLocation.Id, Extension: r.PublisherLocation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_publisherLocation\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactPublicationFormPeriodicRelease) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) marshalJSON(w io.Writer) error {
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
	if r.CitedMedium != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"citedMedium\":"))
		if err != nil {
			return err
		}
		err = r.CitedMedium.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Volume != nil && r.Volume.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"volume\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Volume)
		if err != nil {
			return err
		}
	}
	if r.Volume != nil && (r.Volume.Id != nil || r.Volume.Extension != nil) {
		p := primitiveElement{Id: r.Volume.Id, Extension: r.Volume.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_volume\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Issue != nil && r.Issue.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"issue\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Issue)
		if err != nil {
			return err
		}
	}
	if r.Issue != nil && (r.Issue.Id != nil || r.Issue.Extension != nil) {
		p := primitiveElement{Id: r.Issue.Id, Extension: r.Issue.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_issue\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DateOfPublication != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dateOfPublication\":"))
		if err != nil {
			return err
		}
		err = r.DateOfPublication.marshalJSON(w)
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
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) marshalJSON(w io.Writer) error {
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
	if r.Date != nil && r.Date.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"date\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Date)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		p := primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_date\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Year != nil && r.Year.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"year\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Year)
		if err != nil {
			return err
		}
	}
	if r.Year != nil && (r.Year.Id != nil || r.Year.Extension != nil) {
		p := primitiveElement{Id: r.Year.Id, Extension: r.Year.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_year\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Month != nil && r.Month.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"month\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Month)
		if err != nil {
			return err
		}
	}
	if r.Month != nil && (r.Month.Id != nil || r.Month.Extension != nil) {
		p := primitiveElement{Id: r.Month.Id, Extension: r.Month.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_month\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Day != nil && r.Day.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"day\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Day)
		if err != nil {
			return err
		}
	}
	if r.Day != nil && (r.Day.Id != nil || r.Day.Extension != nil) {
		p := primitiveElement{Id: r.Day.Id, Extension: r.Day.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_day\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Season != nil && r.Season.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"season\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Season)
		if err != nil {
			return err
		}
	}
	if r.Season != nil && (r.Season.Id != nil || r.Season.Extension != nil) {
		p := primitiveElement{Id: r.Season.Id, Extension: r.Season.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_season\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil && r.Text.Value != nil {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactWebLocation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactWebLocation) marshalJSON(w io.Writer) error {
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
	if r.Url != nil && r.Url.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactClassification) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactClassification) marshalJSON(w io.Writer) error {
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
	if len(r.Classifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"classifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Classifier {
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
	if r.WhoClassified != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"whoClassified\":"))
		if err != nil {
			return err
		}
		err = r.WhoClassified.marshalJSON(w)
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
func (r CitationCitedArtifactClassificationWhoClassified) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactClassificationWhoClassified) marshalJSON(w io.Writer) error {
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
	if r.Person != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"person\":"))
		if err != nil {
			return err
		}
		err = r.Person.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Organization != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organization\":"))
		if err != nil {
			return err
		}
		err = r.Organization.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Publisher != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"publisher\":"))
		if err != nil {
			return err
		}
		err = r.Publisher.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ClassifierCopyright != nil && r.ClassifierCopyright.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"classifierCopyright\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ClassifierCopyright)
		if err != nil {
			return err
		}
	}
	if r.ClassifierCopyright != nil && (r.ClassifierCopyright.Id != nil || r.ClassifierCopyright.Extension != nil) {
		p := primitiveElement{Id: r.ClassifierCopyright.Id, Extension: r.ClassifierCopyright.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_classifierCopyright\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FreeToShare != nil && r.FreeToShare.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"freeToShare\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.FreeToShare)
		if err != nil {
			return err
		}
	}
	if r.FreeToShare != nil && (r.FreeToShare.Id != nil || r.FreeToShare.Extension != nil) {
		p := primitiveElement{Id: r.FreeToShare.Id, Extension: r.FreeToShare.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_freeToShare\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactContributorship) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactContributorship) marshalJSON(w io.Writer) error {
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
	if r.Complete != nil && r.Complete.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"complete\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Complete)
		if err != nil {
			return err
		}
	}
	if r.Complete != nil && (r.Complete.Id != nil || r.Complete.Extension != nil) {
		p := primitiveElement{Id: r.Complete.Id, Extension: r.Complete.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_complete\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Entry) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"entry\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Entry {
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
	if len(r.Summary) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"summary\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Summary {
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
func (r CitationCitedArtifactContributorshipEntry) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactContributorshipEntry) marshalJSON(w io.Writer) error {
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
	if r.Name != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		err = r.Name.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Initials != nil && r.Initials.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"initials\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Initials)
		if err != nil {
			return err
		}
	}
	if r.Initials != nil && (r.Initials.Id != nil || r.Initials.Extension != nil) {
		p := primitiveElement{Id: r.Initials.Id, Extension: r.Initials.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_initials\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CollectiveName != nil && r.CollectiveName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"collectiveName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CollectiveName)
		if err != nil {
			return err
		}
	}
	if r.CollectiveName != nil && (r.CollectiveName.Id != nil || r.CollectiveName.Extension != nil) {
		p := primitiveElement{Id: r.CollectiveName.Id, Extension: r.CollectiveName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_collectiveName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	if len(r.AffiliationInfo) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"affiliationInfo\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AffiliationInfo {
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
	if len(r.Address) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"address\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Address {
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
	if len(r.Telecom) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"telecom\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Telecom {
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
	if len(r.ContributionType) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contributionType\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ContributionType {
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
	if r.Role != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"role\":"))
		if err != nil {
			return err
		}
		err = r.Role.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.ContributionInstance) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contributionInstance\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ContributionInstance {
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
	if r.CorrespondingContact != nil && r.CorrespondingContact.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"correspondingContact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CorrespondingContact)
		if err != nil {
			return err
		}
	}
	if r.CorrespondingContact != nil && (r.CorrespondingContact.Id != nil || r.CorrespondingContact.Extension != nil) {
		p := primitiveElement{Id: r.CorrespondingContact.Id, Extension: r.CorrespondingContact.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_correspondingContact\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ListOrder != nil && r.ListOrder.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"listOrder\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ListOrder)
		if err != nil {
			return err
		}
	}
	if r.ListOrder != nil && (r.ListOrder.Id != nil || r.ListOrder.Extension != nil) {
		p := primitiveElement{Id: r.ListOrder.Id, Extension: r.ListOrder.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_listOrder\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) marshalJSON(w io.Writer) error {
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
	if r.Affiliation != nil && r.Affiliation.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"affiliation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Affiliation)
		if err != nil {
			return err
		}
	}
	if r.Affiliation != nil && (r.Affiliation.Id != nil || r.Affiliation.Extension != nil) {
		p := primitiveElement{Id: r.Affiliation.Id, Extension: r.Affiliation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_affiliation\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Role != nil && r.Role.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"role\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Role)
		if err != nil {
			return err
		}
	}
	if r.Role != nil && (r.Role.Id != nil || r.Role.Extension != nil) {
		p := primitiveElement{Id: r.Role.Id, Extension: r.Role.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_role\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) marshalJSON(w io.Writer) error {
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
	if r.Time != nil && r.Time.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"time\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Time)
		if err != nil {
			return err
		}
	}
	if r.Time != nil && (r.Time.Id != nil || r.Time.Extension != nil) {
		p := primitiveElement{Id: r.Time.Id, Extension: r.Time.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_time\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r CitationCitedArtifactContributorshipSummary) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CitationCitedArtifactContributorshipSummary) marshalJSON(w io.Writer) error {
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
	if r.Style != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"style\":"))
		if err != nil {
			return err
		}
		err = r.Style.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Source != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"source\":"))
		if err != nil {
			return err
		}
		err = r.Source.marshalJSON(w)
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
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Value)
		if err != nil {
			return err
		}
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		p := primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_value\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r *Citation) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *Citation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Citation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "url":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Value = v.Value
		case "_url":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Id = v.Id
			r.Url.Extension = v.Extension
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "version":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Version == nil {
				r.Version = &String{}
			}
			r.Version.Value = v.Value
		case "_version":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Version == nil {
				r.Version = &String{}
			}
			r.Version.Id = v.Id
			r.Version.Extension = v.Extension
		case "name":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Value = v.Value
		case "_name":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Id = v.Id
			r.Name.Extension = v.Extension
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
		case "experimental":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Experimental == nil {
				r.Experimental = &Boolean{}
			}
			r.Experimental.Value = v.Value
		case "_experimental":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Experimental == nil {
				r.Experimental = &Boolean{}
			}
			r.Experimental.Id = v.Id
			r.Experimental.Extension = v.Extension
		case "date":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &DateTime{}
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &DateTime{}
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "publisher":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Publisher == nil {
				r.Publisher = &String{}
			}
			r.Publisher.Value = v.Value
		case "_publisher":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Publisher == nil {
				r.Publisher = &String{}
			}
			r.Publisher.Id = v.Id
			r.Publisher.Extension = v.Extension
		case "contact":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v ContactDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "description":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &Markdown{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &Markdown{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "useContext":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v UsageContext
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.UseContext = append(r.UseContext, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "jurisdiction":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "purpose":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Purpose == nil {
				r.Purpose = &Markdown{}
			}
			r.Purpose.Value = v.Value
		case "_purpose":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Purpose == nil {
				r.Purpose = &Markdown{}
			}
			r.Purpose.Id = v.Id
			r.Purpose.Extension = v.Extension
		case "copyright":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Value = v.Value
		case "_copyright":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Id = v.Id
			r.Copyright.Extension = v.Extension
		case "approvalDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ApprovalDate == nil {
				r.ApprovalDate = &Date{}
			}
			r.ApprovalDate.Value = v.Value
		case "_approvalDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ApprovalDate == nil {
				r.ApprovalDate = &Date{}
			}
			r.ApprovalDate.Id = v.Id
			r.ApprovalDate.Extension = v.Extension
		case "lastReviewDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastReviewDate == nil {
				r.LastReviewDate = &Date{}
			}
			r.LastReviewDate.Value = v.Value
		case "_lastReviewDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastReviewDate == nil {
				r.LastReviewDate = &Date{}
			}
			r.LastReviewDate.Id = v.Id
			r.LastReviewDate.Extension = v.Extension
		case "effectivePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.EffectivePeriod = &v
		case "author":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v ContactDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Author = append(r.Author, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "editor":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v ContactDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Editor = append(r.Editor, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "reviewer":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v ContactDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Reviewer = append(r.Reviewer, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "endorser":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v ContactDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Endorser = append(r.Endorser, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "summary":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CitationSummary
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Summary = append(r.Summary, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "classification":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CitationClassification
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "currentState":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.CurrentState = append(r.CurrentState, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "statusDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CitationStatusDate
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.StatusDate = append(r.StatusDate, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "relatesTo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Citation element", t)
			}
			for d.More() {
				var v CitationRelatesTo
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RelatesTo = append(r.RelatesTo, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Citation element", t)
			}
		case "citedArtifact":
			var v CitationCitedArtifact
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CitedArtifact = &v
		default:
			return fmt.Errorf("invalid field: %s in Citation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Citation element", t)
	}
	return nil
}
func (r *CitationSummary) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationSummary element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationSummary element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationSummary element", t)
			}
		case "style":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Style = &v
		case "text":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationSummary", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationSummary element", t)
	}
	return nil
}
func (r *CitationClassification) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationClassification element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationClassification element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationClassification element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "classifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationClassification element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Classifier = append(r.Classifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationClassification element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CitationClassification", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationClassification element", t)
	}
	return nil
}
func (r *CitationStatusDate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationStatusDate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationStatusDate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationStatusDate element", t)
			}
		case "activity":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Activity = v
		case "actual":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Actual == nil {
				r.Actual = &Boolean{}
			}
			r.Actual.Value = v.Value
		case "_actual":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Actual == nil {
				r.Actual = &Boolean{}
			}
			r.Actual.Id = v.Id
			r.Actual.Extension = v.Extension
		case "period":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Period = v
		default:
			return fmt.Errorf("invalid field: %s in CitationStatusDate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationStatusDate element", t)
	}
	return nil
}
func (r *CitationRelatesTo) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationRelatesTo element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationRelatesTo element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationRelatesTo element", t)
			}
		case "relationshipType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.RelationshipType = v
		case "targetClassifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationRelatesTo element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TargetClassifier = append(r.TargetClassifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationRelatesTo element", t)
			}
		case "targetUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Target != nil {
				r.Target = Uri{
					Extension: r.Target.(Uri).Extension,
					Id:        r.Target.(Uri).Id,
					Value:     v.Value,
				}
			} else {
				r.Target = v
			}
		case "_targetUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Target != nil {
				r.Target = Uri{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Target.(Uri).Value,
				}
			} else {
				r.Target = Uri{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "targetIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		case "targetReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		case "targetAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		default:
			return fmt.Errorf("invalid field: %s in CitationRelatesTo", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationRelatesTo element", t)
	}
	return nil
}
func (r *CitationCitedArtifact) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifact element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "relatedIdentifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RelatedIdentifier = append(r.RelatedIdentifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "dateAccessed":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DateAccessed == nil {
				r.DateAccessed = &DateTime{}
			}
			r.DateAccessed.Value = v.Value
		case "_dateAccessed":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DateAccessed == nil {
				r.DateAccessed = &DateTime{}
			}
			r.DateAccessed.Id = v.Id
			r.DateAccessed.Extension = v.Extension
		case "version":
			var v CitationCitedArtifactVersion
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Version = &v
		case "currentState":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.CurrentState = append(r.CurrentState, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "statusDate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactStatusDate
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.StatusDate = append(r.StatusDate, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "title":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactTitle
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Title = append(r.Title, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "abstract":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactAbstract
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Abstract = append(r.Abstract, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "part":
			var v CitationCitedArtifactPart
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Part = &v
		case "relatesTo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactRelatesTo
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.RelatesTo = append(r.RelatesTo, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "publicationForm":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactPublicationForm
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PublicationForm = append(r.PublicationForm, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "webLocation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactWebLocation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.WebLocation = append(r.WebLocation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "classification":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
			}
			for d.More() {
				var v CitationCitedArtifactClassification
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		case "contributorship":
			var v CitationCitedArtifactContributorship
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Contributorship = &v
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifact element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifact element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifact", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifact element", t)
	}
	return nil
}
func (r *CitationCitedArtifactVersion) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactVersion element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactVersion element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactVersion element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactVersion element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactVersion element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactVersion element", t)
			}
		case "value":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Value.Value = v.Value
		case "_value":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value.Id = v.Id
			r.Value.Extension = v.Extension
		case "baseCitation":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.BaseCitation = &v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactVersion", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactVersion element", t)
	}
	return nil
}
func (r *CitationCitedArtifactStatusDate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactStatusDate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactStatusDate element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactStatusDate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactStatusDate element", t)
			}
		case "activity":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Activity = v
		case "actual":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Actual == nil {
				r.Actual = &Boolean{}
			}
			r.Actual.Value = v.Value
		case "_actual":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Actual == nil {
				r.Actual = &Boolean{}
			}
			r.Actual.Id = v.Id
			r.Actual.Extension = v.Extension
		case "period":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Period = v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactStatusDate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactStatusDate element", t)
	}
	return nil
}
func (r *CitationCitedArtifactTitle) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactTitle element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactTitle element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactTitle element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactTitle element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactTitle element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactTitle element", t)
			}
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactTitle element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactTitle element", t)
			}
		case "language":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Language = &v
		case "text":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactTitle", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactTitle element", t)
	}
	return nil
}
func (r *CitationCitedArtifactAbstract) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactAbstract element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactAbstract element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactAbstract element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactAbstract element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactAbstract element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactAbstract element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "language":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Language = &v
		case "text":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		case "copyright":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Value = v.Value
		case "_copyright":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Id = v.Id
			r.Copyright.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactAbstract", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactAbstract element", t)
	}
	return nil
}
func (r *CitationCitedArtifactPart) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactPart element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactPart element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPart element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPart element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPart element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPart element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "value":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value == nil {
				r.Value = &String{}
			}
			r.Value.Value = v.Value
		case "_value":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value == nil {
				r.Value = &String{}
			}
			r.Value.Id = v.Id
			r.Value.Extension = v.Extension
		case "baseCitation":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.BaseCitation = &v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactPart", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactPart element", t)
	}
	return nil
}
func (r *CitationCitedArtifactRelatesTo) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactRelatesTo element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactRelatesTo element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactRelatesTo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactRelatesTo element", t)
			}
		case "relationshipType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.RelationshipType = v
		case "targetClassifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactRelatesTo element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TargetClassifier = append(r.TargetClassifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactRelatesTo element", t)
			}
		case "targetUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Target != nil {
				r.Target = Uri{
					Extension: r.Target.(Uri).Extension,
					Id:        r.Target.(Uri).Id,
					Value:     v.Value,
				}
			} else {
				r.Target = v
			}
		case "_targetUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Target != nil {
				r.Target = Uri{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Target.(Uri).Value,
				}
			} else {
				r.Target = Uri{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "targetIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		case "targetReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		case "targetAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Target = v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactRelatesTo", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactRelatesTo element", t)
	}
	return nil
}
func (r *CitationCitedArtifactPublicationForm) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactPublicationForm element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactPublicationForm element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationForm element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationForm element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationForm element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationForm element", t)
			}
		case "publishedIn":
			var v CitationCitedArtifactPublicationFormPublishedIn
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PublishedIn = &v
		case "periodicRelease":
			var v CitationCitedArtifactPublicationFormPeriodicRelease
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PeriodicRelease = &v
		case "articleDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ArticleDate == nil {
				r.ArticleDate = &DateTime{}
			}
			r.ArticleDate.Value = v.Value
		case "_articleDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ArticleDate == nil {
				r.ArticleDate = &DateTime{}
			}
			r.ArticleDate.Id = v.Id
			r.ArticleDate.Extension = v.Extension
		case "lastRevisionDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastRevisionDate == nil {
				r.LastRevisionDate = &DateTime{}
			}
			r.LastRevisionDate.Value = v.Value
		case "_lastRevisionDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastRevisionDate == nil {
				r.LastRevisionDate = &DateTime{}
			}
			r.LastRevisionDate.Id = v.Id
			r.LastRevisionDate.Extension = v.Extension
		case "language":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationForm element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Language = append(r.Language, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationForm element", t)
			}
		case "accessionNumber":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AccessionNumber == nil {
				r.AccessionNumber = &String{}
			}
			r.AccessionNumber.Value = v.Value
		case "_accessionNumber":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AccessionNumber == nil {
				r.AccessionNumber = &String{}
			}
			r.AccessionNumber.Id = v.Id
			r.AccessionNumber.Extension = v.Extension
		case "pageString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PageString == nil {
				r.PageString = &String{}
			}
			r.PageString.Value = v.Value
		case "_pageString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PageString == nil {
				r.PageString = &String{}
			}
			r.PageString.Id = v.Id
			r.PageString.Extension = v.Extension
		case "firstPage":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.FirstPage == nil {
				r.FirstPage = &String{}
			}
			r.FirstPage.Value = v.Value
		case "_firstPage":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.FirstPage == nil {
				r.FirstPage = &String{}
			}
			r.FirstPage.Id = v.Id
			r.FirstPage.Extension = v.Extension
		case "lastPage":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastPage == nil {
				r.LastPage = &String{}
			}
			r.LastPage.Value = v.Value
		case "_lastPage":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastPage == nil {
				r.LastPage = &String{}
			}
			r.LastPage.Id = v.Id
			r.LastPage.Extension = v.Extension
		case "pageCount":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PageCount == nil {
				r.PageCount = &String{}
			}
			r.PageCount.Value = v.Value
		case "_pageCount":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PageCount == nil {
				r.PageCount = &String{}
			}
			r.PageCount.Id = v.Id
			r.PageCount.Extension = v.Extension
		case "copyright":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Value = v.Value
		case "_copyright":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Copyright == nil {
				r.Copyright = &Markdown{}
			}
			r.Copyright.Id = v.Id
			r.Copyright.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactPublicationForm", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactPublicationForm element", t)
	}
	return nil
}
func (r *CitationCitedArtifactPublicationFormPublishedIn) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactPublicationFormPublishedIn element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactPublicationFormPublishedIn element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPublishedIn element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPublishedIn element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPublishedIn element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPublishedIn element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPublishedIn element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPublishedIn element", t)
			}
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
		case "publisher":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Publisher = &v
		case "publisherLocation":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PublisherLocation == nil {
				r.PublisherLocation = &String{}
			}
			r.PublisherLocation.Value = v.Value
		case "_publisherLocation":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PublisherLocation == nil {
				r.PublisherLocation = &String{}
			}
			r.PublisherLocation.Id = v.Id
			r.PublisherLocation.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactPublicationFormPublishedIn", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactPublicationFormPublishedIn element", t)
	}
	return nil
}
func (r *CitationCitedArtifactPublicationFormPeriodicRelease) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
			}
		case "citedMedium":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CitedMedium = &v
		case "volume":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Volume == nil {
				r.Volume = &String{}
			}
			r.Volume.Value = v.Value
		case "_volume":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Volume == nil {
				r.Volume = &String{}
			}
			r.Volume.Id = v.Id
			r.Volume.Extension = v.Extension
		case "issue":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Issue == nil {
				r.Issue = &String{}
			}
			r.Issue.Value = v.Value
		case "_issue":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Issue == nil {
				r.Issue = &String{}
			}
			r.Issue.Id = v.Id
			r.Issue.Extension = v.Extension
		case "dateOfPublication":
			var v CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DateOfPublication = &v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactPublicationFormPeriodicRelease", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactPublicationFormPeriodicRelease element", t)
	}
	return nil
}
func (r *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
			}
		case "date":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "year":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Year == nil {
				r.Year = &String{}
			}
			r.Year.Value = v.Value
		case "_year":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Year == nil {
				r.Year = &String{}
			}
			r.Year.Id = v.Id
			r.Year.Extension = v.Extension
		case "month":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Month == nil {
				r.Month = &String{}
			}
			r.Month.Value = v.Value
		case "_month":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Month == nil {
				r.Month = &String{}
			}
			r.Month.Id = v.Id
			r.Month.Extension = v.Extension
		case "day":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Day == nil {
				r.Day = &String{}
			}
			r.Day.Value = v.Value
		case "_day":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Day == nil {
				r.Day = &String{}
			}
			r.Day.Id = v.Id
			r.Day.Extension = v.Extension
		case "season":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Season == nil {
				r.Season = &String{}
			}
			r.Season.Value = v.Value
		case "_season":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Season == nil {
				r.Season = &String{}
			}
			r.Season.Id = v.Id
			r.Season.Extension = v.Extension
		case "text":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Text == nil {
				r.Text = &String{}
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Text == nil {
				r.Text = &String{}
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication element", t)
	}
	return nil
}
func (r *CitationCitedArtifactWebLocation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactWebLocation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactWebLocation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactWebLocation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactWebLocation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactWebLocation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactWebLocation element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "url":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Value = v.Value
		case "_url":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Id = v.Id
			r.Url.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactWebLocation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactWebLocation element", t)
	}
	return nil
}
func (r *CitationCitedArtifactClassification) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactClassification element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactClassification element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactClassification element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactClassification element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "classifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactClassification element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Classifier = append(r.Classifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactClassification element", t)
			}
		case "whoClassified":
			var v CitationCitedArtifactClassificationWhoClassified
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.WhoClassified = &v
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactClassification", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactClassification element", t)
	}
	return nil
}
func (r *CitationCitedArtifactClassificationWhoClassified) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactClassificationWhoClassified element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactClassificationWhoClassified element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactClassificationWhoClassified element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactClassificationWhoClassified element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactClassificationWhoClassified element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactClassificationWhoClassified element", t)
			}
		case "person":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Person = &v
		case "organization":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Organization = &v
		case "publisher":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Publisher = &v
		case "classifierCopyright":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ClassifierCopyright == nil {
				r.ClassifierCopyright = &String{}
			}
			r.ClassifierCopyright.Value = v.Value
		case "_classifierCopyright":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ClassifierCopyright == nil {
				r.ClassifierCopyright = &String{}
			}
			r.ClassifierCopyright.Id = v.Id
			r.ClassifierCopyright.Extension = v.Extension
		case "freeToShare":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.FreeToShare == nil {
				r.FreeToShare = &Boolean{}
			}
			r.FreeToShare.Value = v.Value
		case "_freeToShare":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.FreeToShare == nil {
				r.FreeToShare = &Boolean{}
			}
			r.FreeToShare.Id = v.Id
			r.FreeToShare.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactClassificationWhoClassified", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactClassificationWhoClassified element", t)
	}
	return nil
}
func (r *CitationCitedArtifactContributorship) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactContributorship element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactContributorship element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorship element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorship element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorship element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorship element", t)
			}
		case "complete":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Complete == nil {
				r.Complete = &Boolean{}
			}
			r.Complete.Value = v.Value
		case "_complete":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Complete == nil {
				r.Complete = &Boolean{}
			}
			r.Complete.Id = v.Id
			r.Complete.Extension = v.Extension
		case "entry":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorship element", t)
			}
			for d.More() {
				var v CitationCitedArtifactContributorshipEntry
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Entry = append(r.Entry, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorship element", t)
			}
		case "summary":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorship element", t)
			}
			for d.More() {
				var v CitationCitedArtifactContributorshipSummary
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Summary = append(r.Summary, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorship element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactContributorship", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactContributorship element", t)
	}
	return nil
}
func (r *CitationCitedArtifactContributorshipEntry) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactContributorshipEntry element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactContributorshipEntry element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "name":
			var v HumanName
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Name = &v
		case "initials":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Initials == nil {
				r.Initials = &String{}
			}
			r.Initials.Value = v.Value
		case "_initials":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Initials == nil {
				r.Initials = &String{}
			}
			r.Initials.Id = v.Id
			r.Initials.Extension = v.Extension
		case "collectiveName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CollectiveName == nil {
				r.CollectiveName = &String{}
			}
			r.CollectiveName.Value = v.Value
		case "_collectiveName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CollectiveName == nil {
				r.CollectiveName = &String{}
			}
			r.CollectiveName.Id = v.Id
			r.CollectiveName.Extension = v.Extension
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "affiliationInfo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
			}
			for d.More() {
				var v CitationCitedArtifactContributorshipEntryAffiliationInfo
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AffiliationInfo = append(r.AffiliationInfo, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "address":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
			}
			for d.More() {
				var v Address
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Address = append(r.Address, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "telecom":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
			}
			for d.More() {
				var v ContactPoint
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "contributionType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ContributionType = append(r.ContributionType, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "role":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Role = &v
		case "contributionInstance":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntry element", t)
			}
			for d.More() {
				var v CitationCitedArtifactContributorshipEntryContributionInstance
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ContributionInstance = append(r.ContributionInstance, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntry element", t)
			}
		case "correspondingContact":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CorrespondingContact == nil {
				r.CorrespondingContact = &Boolean{}
			}
			r.CorrespondingContact.Value = v.Value
		case "_correspondingContact":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CorrespondingContact == nil {
				r.CorrespondingContact = &Boolean{}
			}
			r.CorrespondingContact.Id = v.Id
			r.CorrespondingContact.Extension = v.Extension
		case "listOrder":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ListOrder == nil {
				r.ListOrder = &PositiveInt{}
			}
			r.ListOrder.Value = v.Value
		case "_listOrder":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ListOrder == nil {
				r.ListOrder = &PositiveInt{}
			}
			r.ListOrder.Id = v.Id
			r.ListOrder.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactContributorshipEntry", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactContributorshipEntry element", t)
	}
	return nil
}
func (r *CitationCitedArtifactContributorshipEntryAffiliationInfo) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
			}
		case "affiliation":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Affiliation == nil {
				r.Affiliation = &String{}
			}
			r.Affiliation.Value = v.Value
		case "_affiliation":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Affiliation == nil {
				r.Affiliation = &String{}
			}
			r.Affiliation.Id = v.Id
			r.Affiliation.Extension = v.Extension
		case "role":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Role == nil {
				r.Role = &String{}
			}
			r.Role.Value = v.Value
		case "_role":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Role == nil {
				r.Role = &String{}
			}
			r.Role.Id = v.Id
			r.Role.Extension = v.Extension
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactContributorshipEntryAffiliationInfo", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactContributorshipEntryAffiliationInfo element", t)
	}
	return nil
}
func (r *CitationCitedArtifactContributorshipEntryContributionInstance) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "time":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Time == nil {
				r.Time = &DateTime{}
			}
			r.Time.Value = v.Value
		case "_time":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Time == nil {
				r.Time = &DateTime{}
			}
			r.Time.Id = v.Id
			r.Time.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactContributorshipEntryContributionInstance", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactContributorshipEntryContributionInstance element", t)
	}
	return nil
}
func (r *CitationCitedArtifactContributorshipSummary) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CitationCitedArtifactContributorshipSummary element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CitationCitedArtifactContributorshipSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipSummary element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CitationCitedArtifactContributorshipSummary element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in CitationCitedArtifactContributorshipSummary element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "style":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Style = &v
		case "source":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Source = &v
		case "value":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Value.Value = v.Value
		case "_value":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value.Id = v.Id
			r.Value.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CitationCitedArtifactContributorshipSummary", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CitationCitedArtifactContributorshipSummary element", t)
	}
	return nil
}
func (r Citation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Citation"
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
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
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
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Experimental, xml.StartElement{Name: xml.Name{Local: "experimental"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UseContext, xml.StartElement{Name: xml.Name{Local: "useContext"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Purpose, xml.StartElement{Name: xml.Name{Local: "purpose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ApprovalDate, xml.StartElement{Name: xml.Name{Local: "approvalDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastReviewDate, xml.StartElement{Name: xml.Name{Local: "lastReviewDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EffectivePeriod, xml.StartElement{Name: xml.Name{Local: "effectivePeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Editor, xml.StartElement{Name: xml.Name{Local: "editor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reviewer, xml.StartElement{Name: xml.Name{Local: "reviewer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Endorser, xml.StartElement{Name: xml.Name{Local: "endorser"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Summary, xml.StartElement{Name: xml.Name{Local: "summary"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Classification, xml.StartElement{Name: xml.Name{Local: "classification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CurrentState, xml.StartElement{Name: xml.Name{Local: "currentState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusDate, xml.StartElement{Name: xml.Name{Local: "statusDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatesTo, xml.StartElement{Name: xml.Name{Local: "relatesTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CitedArtifact, xml.StartElement{Name: xml.Name{Local: "citedArtifact"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationSummary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Style, xml.StartElement{Name: xml.Name{Local: "style"}})
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
func (r CitationClassification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Classifier, xml.StartElement{Name: xml.Name{Local: "classifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationStatusDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Activity, xml.StartElement{Name: xml.Name{Local: "activity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Actual, xml.StartElement{Name: xml.Name{Local: "actual"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationRelatesTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.RelationshipType, xml.StartElement{Name: xml.Name{Local: "relationshipType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetClassifier, xml.StartElement{Name: xml.Name{Local: "targetClassifier"}})
	if err != nil {
		return err
	}
	switch v := r.Target.(type) {
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetUri"}})
		if err != nil {
			return err
		}
	case Identifier, *Identifier:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetIdentifier"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetReference"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetAttachment"}})
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
func (r CitationCitedArtifact) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.RelatedIdentifier, xml.StartElement{Name: xml.Name{Local: "relatedIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateAccessed, xml.StartElement{Name: xml.Name{Local: "dateAccessed"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CurrentState, xml.StartElement{Name: xml.Name{Local: "currentState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusDate, xml.StartElement{Name: xml.Name{Local: "statusDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Abstract, xml.StartElement{Name: xml.Name{Local: "abstract"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Part, xml.StartElement{Name: xml.Name{Local: "part"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatesTo, xml.StartElement{Name: xml.Name{Local: "relatesTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PublicationForm, xml.StartElement{Name: xml.Name{Local: "publicationForm"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WebLocation, xml.StartElement{Name: xml.Name{Local: "webLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Classification, xml.StartElement{Name: xml.Name{Local: "classification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contributorship, xml.StartElement{Name: xml.Name{Local: "contributorship"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactVersion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BaseCitation, xml.StartElement{Name: xml.Name{Local: "baseCitation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactStatusDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Activity, xml.StartElement{Name: xml.Name{Local: "activity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Actual, xml.StartElement{Name: xml.Name{Local: "actual"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactTitle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
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
func (r CitationCitedArtifactAbstract) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
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
func (r CitationCitedArtifactPart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BaseCitation, xml.StartElement{Name: xml.Name{Local: "baseCitation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactRelatesTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.RelationshipType, xml.StartElement{Name: xml.Name{Local: "relationshipType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetClassifier, xml.StartElement{Name: xml.Name{Local: "targetClassifier"}})
	if err != nil {
		return err
	}
	switch v := r.Target.(type) {
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetUri"}})
		if err != nil {
			return err
		}
	case Identifier, *Identifier:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetIdentifier"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetReference"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "targetAttachment"}})
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
func (r CitationCitedArtifactPublicationForm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.PublishedIn, xml.StartElement{Name: xml.Name{Local: "publishedIn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PeriodicRelease, xml.StartElement{Name: xml.Name{Local: "periodicRelease"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ArticleDate, xml.StartElement{Name: xml.Name{Local: "articleDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastRevisionDate, xml.StartElement{Name: xml.Name{Local: "lastRevisionDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AccessionNumber, xml.StartElement{Name: xml.Name{Local: "accessionNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PageString, xml.StartElement{Name: xml.Name{Local: "pageString"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FirstPage, xml.StartElement{Name: xml.Name{Local: "firstPage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastPage, xml.StartElement{Name: xml.Name{Local: "lastPage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PageCount, xml.StartElement{Name: xml.Name{Local: "pageCount"}})
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
func (r CitationCitedArtifactPublicationFormPublishedIn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PublisherLocation, xml.StartElement{Name: xml.Name{Local: "publisherLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.CitedMedium, xml.StartElement{Name: xml.Name{Local: "citedMedium"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Volume, xml.StartElement{Name: xml.Name{Local: "volume"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Issue, xml.StartElement{Name: xml.Name{Local: "issue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateOfPublication, xml.StartElement{Name: xml.Name{Local: "dateOfPublication"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Year, xml.StartElement{Name: xml.Name{Local: "year"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Month, xml.StartElement{Name: xml.Name{Local: "month"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Day, xml.StartElement{Name: xml.Name{Local: "day"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Season, xml.StartElement{Name: xml.Name{Local: "season"}})
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
func (r CitationCitedArtifactWebLocation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactClassification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Classifier, xml.StartElement{Name: xml.Name{Local: "classifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WhoClassified, xml.StartElement{Name: xml.Name{Local: "whoClassified"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactClassificationWhoClassified) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Person, xml.StartElement{Name: xml.Name{Local: "person"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Organization, xml.StartElement{Name: xml.Name{Local: "organization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClassifierCopyright, xml.StartElement{Name: xml.Name{Local: "classifierCopyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FreeToShare, xml.StartElement{Name: xml.Name{Local: "freeToShare"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Complete, xml.StartElement{Name: xml.Name{Local: "complete"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Entry, xml.StartElement{Name: xml.Name{Local: "entry"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Summary, xml.StartElement{Name: xml.Name{Local: "summary"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorshipEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Initials, xml.StartElement{Name: xml.Name{Local: "initials"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CollectiveName, xml.StartElement{Name: xml.Name{Local: "collectiveName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AffiliationInfo, xml.StartElement{Name: xml.Name{Local: "affiliationInfo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContributionType, xml.StartElement{Name: xml.Name{Local: "contributionType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContributionInstance, xml.StartElement{Name: xml.Name{Local: "contributionInstance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CorrespondingContact, xml.StartElement{Name: xml.Name{Local: "correspondingContact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ListOrder, xml.StartElement{Name: xml.Name{Local: "listOrder"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Affiliation, xml.StartElement{Name: xml.Name{Local: "affiliation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Time, xml.StartElement{Name: xml.Name{Local: "time"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CitationCitedArtifactContributorshipSummary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Style, xml.StartElement{Name: xml.Name{Local: "style"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Citation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
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
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "experimental":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Experimental = &v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "publisher":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "contact":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "useContext":
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UseContext = append(r.UseContext, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "purpose":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Purpose = &v
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			case "approvalDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ApprovalDate = &v
			case "lastReviewDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastReviewDate = &v
			case "effectivePeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EffectivePeriod = &v
			case "author":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = append(r.Author, v)
			case "editor":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Editor = append(r.Editor, v)
			case "reviewer":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reviewer = append(r.Reviewer, v)
			case "endorser":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Endorser = append(r.Endorser, v)
			case "summary":
				var v CitationSummary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Summary = append(r.Summary, v)
			case "classification":
				var v CitationClassification
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "currentState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CurrentState = append(r.CurrentState, v)
			case "statusDate":
				var v CitationStatusDate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = append(r.StatusDate, v)
			case "relatesTo":
				var v CitationRelatesTo
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatesTo = append(r.RelatesTo, v)
			case "citedArtifact":
				var v CitationCitedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CitedArtifact = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationSummary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "style":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Style = &v
			case "text":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationClassification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "classifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classifier = append(r.Classifier, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationStatusDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "activity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Activity = v
			case "actual":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Actual = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationRelatesTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relationshipType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelationshipType = v
			case "targetClassifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetClassifier = append(r.TargetClassifier, v)
			case "targetUri":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetIdentifier":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetReference":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetAttachment":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifact) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relatedIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatedIdentifier = append(r.RelatedIdentifier, v)
			case "dateAccessed":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateAccessed = &v
			case "version":
				var v CitationCitedArtifactVersion
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "currentState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CurrentState = append(r.CurrentState, v)
			case "statusDate":
				var v CitationCitedArtifactStatusDate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = append(r.StatusDate, v)
			case "title":
				var v CitationCitedArtifactTitle
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = append(r.Title, v)
			case "abstract":
				var v CitationCitedArtifactAbstract
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abstract = append(r.Abstract, v)
			case "part":
				var v CitationCitedArtifactPart
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Part = &v
			case "relatesTo":
				var v CitationCitedArtifactRelatesTo
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatesTo = append(r.RelatesTo, v)
			case "publicationForm":
				var v CitationCitedArtifactPublicationForm
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublicationForm = append(r.PublicationForm, v)
			case "webLocation":
				var v CitationCitedArtifactWebLocation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WebLocation = append(r.WebLocation, v)
			case "classification":
				var v CitationCitedArtifactClassification
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			case "contributorship":
				var v CitationCitedArtifactContributorship
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contributorship = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactVersion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "baseCitation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BaseCitation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactStatusDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "activity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Activity = v
			case "actual":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Actual = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactTitle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Type = append(r.Type, v)
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactAbstract) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = v
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
func (r *CitationCitedArtifactPart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "value":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			case "baseCitation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BaseCitation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactRelatesTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relationshipType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelationshipType = v
			case "targetClassifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetClassifier = append(r.TargetClassifier, v)
			case "targetUri":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetIdentifier":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetReference":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			case "targetAttachment":
				if r.Target != nil {
					return fmt.Errorf("multiple values for field \"Target\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactPublicationForm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "publishedIn":
				var v CitationCitedArtifactPublicationFormPublishedIn
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublishedIn = &v
			case "periodicRelease":
				var v CitationCitedArtifactPublicationFormPeriodicRelease
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PeriodicRelease = &v
			case "articleDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ArticleDate = &v
			case "lastRevisionDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastRevisionDate = &v
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = append(r.Language, v)
			case "accessionNumber":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AccessionNumber = &v
			case "pageString":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PageString = &v
			case "firstPage":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FirstPage = &v
			case "lastPage":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastPage = &v
			case "pageCount":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PageCount = &v
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
func (r *CitationCitedArtifactPublicationFormPublishedIn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "publisher":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "publisherLocation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublisherLocation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactPublicationFormPeriodicRelease) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "citedMedium":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CitedMedium = &v
			case "volume":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Volume = &v
			case "issue":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Issue = &v
			case "dateOfPublication":
				var v CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateOfPublication = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "date":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "year":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Year = &v
			case "month":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Month = &v
			case "day":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Day = &v
			case "season":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Season = &v
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
func (r *CitationCitedArtifactWebLocation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactClassification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "classifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classifier = append(r.Classifier, v)
			case "whoClassified":
				var v CitationCitedArtifactClassificationWhoClassified
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WhoClassified = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactClassificationWhoClassified) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "person":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Person = &v
			case "organization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Organization = &v
			case "publisher":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "classifierCopyright":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClassifierCopyright = &v
			case "freeToShare":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FreeToShare = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactContributorship) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "complete":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Complete = &v
			case "entry":
				var v CitationCitedArtifactContributorshipEntry
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Entry = append(r.Entry, v)
			case "summary":
				var v CitationCitedArtifactContributorshipSummary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Summary = append(r.Summary, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactContributorshipEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "name":
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "initials":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Initials = &v
			case "collectiveName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CollectiveName = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "affiliationInfo":
				var v CitationCitedArtifactContributorshipEntryAffiliationInfo
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AffiliationInfo = append(r.AffiliationInfo, v)
			case "address":
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = append(r.Address, v)
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "contributionType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContributionType = append(r.ContributionType, v)
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = &v
			case "contributionInstance":
				var v CitationCitedArtifactContributorshipEntryContributionInstance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContributionInstance = append(r.ContributionInstance, v)
			case "correspondingContact":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CorrespondingContact = &v
			case "listOrder":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ListOrder = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactContributorshipEntryAffiliationInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "affiliation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Affiliation = &v
			case "role":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactContributorshipEntryContributionInstance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "time":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CitationCitedArtifactContributorshipSummary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "style":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Style = &v
			case "source":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			case "value":
				var v Markdown
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
func (r Citation) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "url") {
		if r.Url != nil {
			children = append(children, *r.Url)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "version") {
		if r.Version != nil {
			children = append(children, *r.Version)
		}
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		if r.Name != nil {
			children = append(children, *r.Name)
		}
	}
	if len(name) == 0 || slices.Contains(name, "title") {
		if r.Title != nil {
			children = append(children, *r.Title)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "experimental") {
		if r.Experimental != nil {
			children = append(children, *r.Experimental)
		}
	}
	if len(name) == 0 || slices.Contains(name, "date") {
		if r.Date != nil {
			children = append(children, *r.Date)
		}
	}
	if len(name) == 0 || slices.Contains(name, "publisher") {
		if r.Publisher != nil {
			children = append(children, *r.Publisher)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contact") {
		for _, v := range r.Contact {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "useContext") {
		for _, v := range r.UseContext {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "jurisdiction") {
		for _, v := range r.Jurisdiction {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "purpose") {
		if r.Purpose != nil {
			children = append(children, *r.Purpose)
		}
	}
	if len(name) == 0 || slices.Contains(name, "copyright") {
		if r.Copyright != nil {
			children = append(children, *r.Copyright)
		}
	}
	if len(name) == 0 || slices.Contains(name, "approvalDate") {
		if r.ApprovalDate != nil {
			children = append(children, *r.ApprovalDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lastReviewDate") {
		if r.LastReviewDate != nil {
			children = append(children, *r.LastReviewDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "effectivePeriod") {
		if r.EffectivePeriod != nil {
			children = append(children, *r.EffectivePeriod)
		}
	}
	if len(name) == 0 || slices.Contains(name, "author") {
		for _, v := range r.Author {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "editor") {
		for _, v := range r.Editor {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewer") {
		for _, v := range r.Reviewer {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "endorser") {
		for _, v := range r.Endorser {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "summary") {
		for _, v := range r.Summary {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "classification") {
		for _, v := range r.Classification {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "currentState") {
		for _, v := range r.CurrentState {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "statusDate") {
		for _, v := range r.StatusDate {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relatesTo") {
		for _, v := range r.RelatesTo {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "citedArtifact") {
		if r.CitedArtifact != nil {
			children = append(children, *r.CitedArtifact)
		}
	}
	return children
}
func (r Citation) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Citation to Boolean")
}
func (r Citation) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Citation to String")
}
func (r Citation) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Citation to Integer")
}
func (r Citation) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Citation to Decimal")
}
func (r Citation) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Citation to Date")
}
func (r Citation) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Citation to Time")
}
func (r Citation) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Citation to DateTime")
}
func (r Citation) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Citation to Quantity")
}
func (r Citation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Citation
	switch other := other.(type) {
	case Citation:
		o = &other
	case *Citation:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r Citation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r Citation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Version",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Name",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Title",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Experimental",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Date",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Publisher",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contact",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "Description",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "UseContext",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "UsageContext",
				Namespace: "FHIR",
			},
		}, {
			Name: "Jurisdiction",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Purpose",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "Copyright",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "ApprovalDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastReviewDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "EffectivePeriod",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}, {
			Name: "Author",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "Editor",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "Reviewer",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "Endorser",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactDetail",
				Namespace: "FHIR",
			},
		}, {
			Name: "Summary",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationSummary",
				Namespace: "FHIR",
			},
		}, {
			Name: "Classification",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationClassification",
				Namespace: "FHIR",
			},
		}, {
			Name: "Note",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Annotation",
				Namespace: "FHIR",
			},
		}, {
			Name: "CurrentState",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "StatusDate",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationStatusDate",
				Namespace: "FHIR",
			},
		}, {
			Name: "RelatesTo",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationRelatesTo",
				Namespace: "FHIR",
			},
		}, {
			Name: "CitedArtifact",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifact",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "Citation",
			Namespace: "FHIR",
		},
	}
}
func (r CitationSummary) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "style") {
		if r.Style != nil {
			children = append(children, *r.Style)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		children = append(children, r.Text)
	}
	return children
}
func (r CitationSummary) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationSummary to Boolean")
}
func (r CitationSummary) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationSummary to String")
}
func (r CitationSummary) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationSummary to Integer")
}
func (r CitationSummary) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationSummary to Decimal")
}
func (r CitationSummary) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationSummary to Date")
}
func (r CitationSummary) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationSummary to Time")
}
func (r CitationSummary) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationSummary to DateTime")
}
func (r CitationSummary) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationSummary to Quantity")
}
func (r CitationSummary) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationSummary
	switch other := other.(type) {
	case CitationSummary:
		o = &other
	case *CitationSummary:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationSummary) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationSummary) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Style",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationSummary",
			Namespace: "FHIR",
		},
	}
}
func (r CitationClassification) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "classifier") {
		for _, v := range r.Classifier {
			children = append(children, v)
		}
	}
	return children
}
func (r CitationClassification) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationClassification to Boolean")
}
func (r CitationClassification) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationClassification to String")
}
func (r CitationClassification) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationClassification to Integer")
}
func (r CitationClassification) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationClassification to Decimal")
}
func (r CitationClassification) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationClassification to Date")
}
func (r CitationClassification) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationClassification to Time")
}
func (r CitationClassification) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationClassification to DateTime")
}
func (r CitationClassification) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationClassification to Quantity")
}
func (r CitationClassification) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationClassification
	switch other := other.(type) {
	case CitationClassification:
		o = &other
	case *CitationClassification:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationClassification) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationClassification) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Classifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationClassification",
			Namespace: "FHIR",
		},
	}
}
func (r CitationStatusDate) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "activity") {
		children = append(children, r.Activity)
	}
	if len(name) == 0 || slices.Contains(name, "actual") {
		if r.Actual != nil {
			children = append(children, *r.Actual)
		}
	}
	if len(name) == 0 || slices.Contains(name, "period") {
		children = append(children, r.Period)
	}
	return children
}
func (r CitationStatusDate) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationStatusDate to Boolean")
}
func (r CitationStatusDate) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationStatusDate to String")
}
func (r CitationStatusDate) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationStatusDate to Integer")
}
func (r CitationStatusDate) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationStatusDate to Decimal")
}
func (r CitationStatusDate) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationStatusDate to Date")
}
func (r CitationStatusDate) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationStatusDate to Time")
}
func (r CitationStatusDate) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationStatusDate to DateTime")
}
func (r CitationStatusDate) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationStatusDate to Quantity")
}
func (r CitationStatusDate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationStatusDate
	switch other := other.(type) {
	case CitationStatusDate:
		o = &other
	case *CitationStatusDate:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationStatusDate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationStatusDate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Activity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Actual",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Period",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationStatusDate",
			Namespace: "FHIR",
		},
	}
}
func (r CitationRelatesTo) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "relationshipType") {
		children = append(children, r.RelationshipType)
	}
	if len(name) == 0 || slices.Contains(name, "targetClassifier") {
		for _, v := range r.TargetClassifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "target") {
		children = append(children, r.Target)
	}
	return children
}
func (r CitationRelatesTo) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationRelatesTo to Boolean")
}
func (r CitationRelatesTo) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationRelatesTo to String")
}
func (r CitationRelatesTo) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationRelatesTo to Integer")
}
func (r CitationRelatesTo) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationRelatesTo to Decimal")
}
func (r CitationRelatesTo) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationRelatesTo to Date")
}
func (r CitationRelatesTo) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationRelatesTo to Time")
}
func (r CitationRelatesTo) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationRelatesTo to DateTime")
}
func (r CitationRelatesTo) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationRelatesTo to Quantity")
}
func (r CitationRelatesTo) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationRelatesTo
	switch other := other.(type) {
	case CitationRelatesTo:
		o = &other
	case *CitationRelatesTo:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationRelatesTo) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationRelatesTo) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "RelationshipType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "TargetClassifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Target",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationRelatesTo",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifact) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relatedIdentifier") {
		for _, v := range r.RelatedIdentifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dateAccessed") {
		if r.DateAccessed != nil {
			children = append(children, *r.DateAccessed)
		}
	}
	if len(name) == 0 || slices.Contains(name, "version") {
		if r.Version != nil {
			children = append(children, *r.Version)
		}
	}
	if len(name) == 0 || slices.Contains(name, "currentState") {
		for _, v := range r.CurrentState {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "statusDate") {
		for _, v := range r.StatusDate {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "title") {
		for _, v := range r.Title {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "abstract") {
		for _, v := range r.Abstract {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "part") {
		if r.Part != nil {
			children = append(children, *r.Part)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relatesTo") {
		for _, v := range r.RelatesTo {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "publicationForm") {
		for _, v := range r.PublicationForm {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "webLocation") {
		for _, v := range r.WebLocation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "classification") {
		for _, v := range r.Classification {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contributorship") {
		if r.Contributorship != nil {
			children = append(children, *r.Contributorship)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	return children
}
func (r CitationCitedArtifact) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifact to Boolean")
}
func (r CitationCitedArtifact) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifact to String")
}
func (r CitationCitedArtifact) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifact to Integer")
}
func (r CitationCitedArtifact) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifact to Decimal")
}
func (r CitationCitedArtifact) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifact to Date")
}
func (r CitationCitedArtifact) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifact to Time")
}
func (r CitationCitedArtifact) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifact to DateTime")
}
func (r CitationCitedArtifact) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifact to Quantity")
}
func (r CitationCitedArtifact) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifact
	switch other := other.(type) {
	case CitationCitedArtifact:
		o = &other
	case *CitationCitedArtifact:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifact) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifact) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "RelatedIdentifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "DateAccessed",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Version",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactVersion",
				Namespace: "FHIR",
			},
		}, {
			Name: "CurrentState",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "StatusDate",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactStatusDate",
				Namespace: "FHIR",
			},
		}, {
			Name: "Title",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactTitle",
				Namespace: "FHIR",
			},
		}, {
			Name: "Abstract",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactAbstract",
				Namespace: "FHIR",
			},
		}, {
			Name: "Part",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactPart",
				Namespace: "FHIR",
			},
		}, {
			Name: "RelatesTo",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactRelatesTo",
				Namespace: "FHIR",
			},
		}, {
			Name: "PublicationForm",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactPublicationForm",
				Namespace: "FHIR",
			},
		}, {
			Name: "WebLocation",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactWebLocation",
				Namespace: "FHIR",
			},
		}, {
			Name: "Classification",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactClassification",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contributorship",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactContributorship",
				Namespace: "FHIR",
			},
		}, {
			Name: "Note",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Annotation",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifact",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactVersion) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "value") {
		children = append(children, r.Value)
	}
	if len(name) == 0 || slices.Contains(name, "baseCitation") {
		if r.BaseCitation != nil {
			children = append(children, *r.BaseCitation)
		}
	}
	return children
}
func (r CitationCitedArtifactVersion) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactVersion to Boolean")
}
func (r CitationCitedArtifactVersion) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactVersion to String")
}
func (r CitationCitedArtifactVersion) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactVersion to Integer")
}
func (r CitationCitedArtifactVersion) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactVersion to Decimal")
}
func (r CitationCitedArtifactVersion) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactVersion to Date")
}
func (r CitationCitedArtifactVersion) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactVersion to Time")
}
func (r CitationCitedArtifactVersion) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactVersion to DateTime")
}
func (r CitationCitedArtifactVersion) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactVersion to Quantity")
}
func (r CitationCitedArtifactVersion) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactVersion
	switch other := other.(type) {
	case CitationCitedArtifactVersion:
		o = &other
	case *CitationCitedArtifactVersion:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactVersion) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactVersion) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Value",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "BaseCitation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactVersion",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactStatusDate) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "activity") {
		children = append(children, r.Activity)
	}
	if len(name) == 0 || slices.Contains(name, "actual") {
		if r.Actual != nil {
			children = append(children, *r.Actual)
		}
	}
	if len(name) == 0 || slices.Contains(name, "period") {
		children = append(children, r.Period)
	}
	return children
}
func (r CitationCitedArtifactStatusDate) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactStatusDate to Boolean")
}
func (r CitationCitedArtifactStatusDate) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactStatusDate to String")
}
func (r CitationCitedArtifactStatusDate) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactStatusDate to Integer")
}
func (r CitationCitedArtifactStatusDate) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactStatusDate to Decimal")
}
func (r CitationCitedArtifactStatusDate) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactStatusDate to Date")
}
func (r CitationCitedArtifactStatusDate) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactStatusDate to Time")
}
func (r CitationCitedArtifactStatusDate) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactStatusDate to DateTime")
}
func (r CitationCitedArtifactStatusDate) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactStatusDate to Quantity")
}
func (r CitationCitedArtifactStatusDate) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactStatusDate
	switch other := other.(type) {
	case CitationCitedArtifactStatusDate:
		o = &other
	case *CitationCitedArtifactStatusDate:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactStatusDate) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactStatusDate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Activity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Actual",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Period",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactStatusDate",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactTitle) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		for _, v := range r.Type {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		children = append(children, r.Text)
	}
	return children
}
func (r CitationCitedArtifactTitle) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactTitle to Boolean")
}
func (r CitationCitedArtifactTitle) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactTitle to String")
}
func (r CitationCitedArtifactTitle) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactTitle to Integer")
}
func (r CitationCitedArtifactTitle) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactTitle to Decimal")
}
func (r CitationCitedArtifactTitle) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactTitle to Date")
}
func (r CitationCitedArtifactTitle) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactTitle to Time")
}
func (r CitationCitedArtifactTitle) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactTitle to DateTime")
}
func (r CitationCitedArtifactTitle) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactTitle to Quantity")
}
func (r CitationCitedArtifactTitle) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactTitle
	switch other := other.(type) {
	case CitationCitedArtifactTitle:
		o = &other
	case *CitationCitedArtifactTitle:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactTitle) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactTitle) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactTitle",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactAbstract) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		children = append(children, r.Text)
	}
	if len(name) == 0 || slices.Contains(name, "copyright") {
		if r.Copyright != nil {
			children = append(children, *r.Copyright)
		}
	}
	return children
}
func (r CitationCitedArtifactAbstract) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactAbstract to Boolean")
}
func (r CitationCitedArtifactAbstract) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactAbstract to String")
}
func (r CitationCitedArtifactAbstract) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactAbstract to Integer")
}
func (r CitationCitedArtifactAbstract) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactAbstract to Decimal")
}
func (r CitationCitedArtifactAbstract) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactAbstract to Date")
}
func (r CitationCitedArtifactAbstract) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactAbstract to Time")
}
func (r CitationCitedArtifactAbstract) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactAbstract to DateTime")
}
func (r CitationCitedArtifactAbstract) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactAbstract to Quantity")
}
func (r CitationCitedArtifactAbstract) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactAbstract
	switch other := other.(type) {
	case CitationCitedArtifactAbstract:
		o = &other
	case *CitationCitedArtifactAbstract:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactAbstract) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactAbstract) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "Copyright",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactAbstract",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactPart) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "value") {
		if r.Value != nil {
			children = append(children, *r.Value)
		}
	}
	if len(name) == 0 || slices.Contains(name, "baseCitation") {
		if r.BaseCitation != nil {
			children = append(children, *r.BaseCitation)
		}
	}
	return children
}
func (r CitationCitedArtifactPart) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactPart to Boolean")
}
func (r CitationCitedArtifactPart) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactPart to String")
}
func (r CitationCitedArtifactPart) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactPart to Integer")
}
func (r CitationCitedArtifactPart) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactPart to Decimal")
}
func (r CitationCitedArtifactPart) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactPart to Date")
}
func (r CitationCitedArtifactPart) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactPart to Time")
}
func (r CitationCitedArtifactPart) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactPart to DateTime")
}
func (r CitationCitedArtifactPart) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactPart to Quantity")
}
func (r CitationCitedArtifactPart) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactPart
	switch other := other.(type) {
	case CitationCitedArtifactPart:
		o = &other
	case *CitationCitedArtifactPart:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactPart) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactPart) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Value",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "BaseCitation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactPart",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactRelatesTo) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "relationshipType") {
		children = append(children, r.RelationshipType)
	}
	if len(name) == 0 || slices.Contains(name, "targetClassifier") {
		for _, v := range r.TargetClassifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "target") {
		children = append(children, r.Target)
	}
	return children
}
func (r CitationCitedArtifactRelatesTo) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Boolean")
}
func (r CitationCitedArtifactRelatesTo) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactRelatesTo to String")
}
func (r CitationCitedArtifactRelatesTo) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Integer")
}
func (r CitationCitedArtifactRelatesTo) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Decimal")
}
func (r CitationCitedArtifactRelatesTo) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Date")
}
func (r CitationCitedArtifactRelatesTo) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Time")
}
func (r CitationCitedArtifactRelatesTo) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactRelatesTo to DateTime")
}
func (r CitationCitedArtifactRelatesTo) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactRelatesTo to Quantity")
}
func (r CitationCitedArtifactRelatesTo) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactRelatesTo
	switch other := other.(type) {
	case CitationCitedArtifactRelatesTo:
		o = &other
	case *CitationCitedArtifactRelatesTo:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactRelatesTo) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactRelatesTo) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "RelationshipType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "TargetClassifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Target",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactRelatesTo",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactPublicationForm) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "publishedIn") {
		if r.PublishedIn != nil {
			children = append(children, *r.PublishedIn)
		}
	}
	if len(name) == 0 || slices.Contains(name, "periodicRelease") {
		if r.PeriodicRelease != nil {
			children = append(children, *r.PeriodicRelease)
		}
	}
	if len(name) == 0 || slices.Contains(name, "articleDate") {
		if r.ArticleDate != nil {
			children = append(children, *r.ArticleDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lastRevisionDate") {
		if r.LastRevisionDate != nil {
			children = append(children, *r.LastRevisionDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		for _, v := range r.Language {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "accessionNumber") {
		if r.AccessionNumber != nil {
			children = append(children, *r.AccessionNumber)
		}
	}
	if len(name) == 0 || slices.Contains(name, "pageString") {
		if r.PageString != nil {
			children = append(children, *r.PageString)
		}
	}
	if len(name) == 0 || slices.Contains(name, "firstPage") {
		if r.FirstPage != nil {
			children = append(children, *r.FirstPage)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lastPage") {
		if r.LastPage != nil {
			children = append(children, *r.LastPage)
		}
	}
	if len(name) == 0 || slices.Contains(name, "pageCount") {
		if r.PageCount != nil {
			children = append(children, *r.PageCount)
		}
	}
	if len(name) == 0 || slices.Contains(name, "copyright") {
		if r.Copyright != nil {
			children = append(children, *r.Copyright)
		}
	}
	return children
}
func (r CitationCitedArtifactPublicationForm) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Boolean")
}
func (r CitationCitedArtifactPublicationForm) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactPublicationForm to String")
}
func (r CitationCitedArtifactPublicationForm) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Integer")
}
func (r CitationCitedArtifactPublicationForm) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Decimal")
}
func (r CitationCitedArtifactPublicationForm) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Date")
}
func (r CitationCitedArtifactPublicationForm) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Time")
}
func (r CitationCitedArtifactPublicationForm) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactPublicationForm to DateTime")
}
func (r CitationCitedArtifactPublicationForm) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactPublicationForm to Quantity")
}
func (r CitationCitedArtifactPublicationForm) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactPublicationForm
	switch other := other.(type) {
	case CitationCitedArtifactPublicationForm:
		o = &other
	case *CitationCitedArtifactPublicationForm:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactPublicationForm) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactPublicationForm) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "PublishedIn",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactPublicationFormPublishedIn",
				Namespace: "FHIR",
			},
		}, {
			Name: "PeriodicRelease",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactPublicationFormPeriodicRelease",
				Namespace: "FHIR",
			},
		}, {
			Name: "ArticleDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastRevisionDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "AccessionNumber",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "PageString",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "FirstPage",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastPage",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "PageCount",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Copyright",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactPublicationForm",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactPublicationFormPublishedIn) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "title") {
		if r.Title != nil {
			children = append(children, *r.Title)
		}
	}
	if len(name) == 0 || slices.Contains(name, "publisher") {
		if r.Publisher != nil {
			children = append(children, *r.Publisher)
		}
	}
	if len(name) == 0 || slices.Contains(name, "publisherLocation") {
		if r.PublisherLocation != nil {
			children = append(children, *r.PublisherLocation)
		}
	}
	return children
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Boolean")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to String")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Integer")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Decimal")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Date")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Time")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to DateTime")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPublishedIn to Quantity")
}
func (r CitationCitedArtifactPublicationFormPublishedIn) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactPublicationFormPublishedIn
	switch other := other.(type) {
	case CitationCitedArtifactPublicationFormPublishedIn:
		o = &other
	case *CitationCitedArtifactPublicationFormPublishedIn:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactPublicationFormPublishedIn) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactPublicationFormPublishedIn) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Title",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Publisher",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "PublisherLocation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactPublicationFormPublishedIn",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "citedMedium") {
		if r.CitedMedium != nil {
			children = append(children, *r.CitedMedium)
		}
	}
	if len(name) == 0 || slices.Contains(name, "volume") {
		if r.Volume != nil {
			children = append(children, *r.Volume)
		}
	}
	if len(name) == 0 || slices.Contains(name, "issue") {
		if r.Issue != nil {
			children = append(children, *r.Issue)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dateOfPublication") {
		if r.DateOfPublication != nil {
			children = append(children, *r.DateOfPublication)
		}
	}
	return children
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Boolean")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to String")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Integer")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Decimal")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Date")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Time")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to DateTime")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicRelease to Quantity")
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactPublicationFormPeriodicRelease
	switch other := other.(type) {
	case CitationCitedArtifactPublicationFormPeriodicRelease:
		o = &other
	case *CitationCitedArtifactPublicationFormPeriodicRelease:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactPublicationFormPeriodicRelease) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "CitedMedium",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Volume",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Issue",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "DateOfPublication",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactPublicationFormPeriodicRelease",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "date") {
		if r.Date != nil {
			children = append(children, *r.Date)
		}
	}
	if len(name) == 0 || slices.Contains(name, "year") {
		if r.Year != nil {
			children = append(children, *r.Year)
		}
	}
	if len(name) == 0 || slices.Contains(name, "month") {
		if r.Month != nil {
			children = append(children, *r.Month)
		}
	}
	if len(name) == 0 || slices.Contains(name, "day") {
		if r.Day != nil {
			children = append(children, *r.Day)
		}
	}
	if len(name) == 0 || slices.Contains(name, "season") {
		if r.Season != nil {
			children = append(children, *r.Season)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	return children
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Boolean")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to String")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Integer")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Decimal")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Date")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Time")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to DateTime")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication to Quantity")
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication
	switch other := other.(type) {
	case CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication:
		o = &other
	case *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Date",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "Year",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Month",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Day",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Season",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactWebLocation) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "url") {
		if r.Url != nil {
			children = append(children, *r.Url)
		}
	}
	return children
}
func (r CitationCitedArtifactWebLocation) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactWebLocation to Boolean")
}
func (r CitationCitedArtifactWebLocation) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactWebLocation to String")
}
func (r CitationCitedArtifactWebLocation) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactWebLocation to Integer")
}
func (r CitationCitedArtifactWebLocation) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactWebLocation to Decimal")
}
func (r CitationCitedArtifactWebLocation) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactWebLocation to Date")
}
func (r CitationCitedArtifactWebLocation) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactWebLocation to Time")
}
func (r CitationCitedArtifactWebLocation) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactWebLocation to DateTime")
}
func (r CitationCitedArtifactWebLocation) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactWebLocation to Quantity")
}
func (r CitationCitedArtifactWebLocation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactWebLocation
	switch other := other.(type) {
	case CitationCitedArtifactWebLocation:
		o = &other
	case *CitationCitedArtifactWebLocation:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactWebLocation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactWebLocation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactWebLocation",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactClassification) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "classifier") {
		for _, v := range r.Classifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "whoClassified") {
		if r.WhoClassified != nil {
			children = append(children, *r.WhoClassified)
		}
	}
	return children
}
func (r CitationCitedArtifactClassification) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactClassification to Boolean")
}
func (r CitationCitedArtifactClassification) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactClassification to String")
}
func (r CitationCitedArtifactClassification) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactClassification to Integer")
}
func (r CitationCitedArtifactClassification) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactClassification to Decimal")
}
func (r CitationCitedArtifactClassification) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactClassification to Date")
}
func (r CitationCitedArtifactClassification) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactClassification to Time")
}
func (r CitationCitedArtifactClassification) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactClassification to DateTime")
}
func (r CitationCitedArtifactClassification) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactClassification to Quantity")
}
func (r CitationCitedArtifactClassification) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactClassification
	switch other := other.(type) {
	case CitationCitedArtifactClassification:
		o = &other
	case *CitationCitedArtifactClassification:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactClassification) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactClassification) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Classifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "WhoClassified",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CitationCitedArtifactClassificationWhoClassified",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactClassification",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactClassificationWhoClassified) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "person") {
		if r.Person != nil {
			children = append(children, *r.Person)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organization") {
		if r.Organization != nil {
			children = append(children, *r.Organization)
		}
	}
	if len(name) == 0 || slices.Contains(name, "publisher") {
		if r.Publisher != nil {
			children = append(children, *r.Publisher)
		}
	}
	if len(name) == 0 || slices.Contains(name, "classifierCopyright") {
		if r.ClassifierCopyright != nil {
			children = append(children, *r.ClassifierCopyright)
		}
	}
	if len(name) == 0 || slices.Contains(name, "freeToShare") {
		if r.FreeToShare != nil {
			children = append(children, *r.FreeToShare)
		}
	}
	return children
}
func (r CitationCitedArtifactClassificationWhoClassified) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Boolean")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to String")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Integer")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Decimal")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Date")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Time")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to DateTime")
}
func (r CitationCitedArtifactClassificationWhoClassified) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactClassificationWhoClassified to Quantity")
}
func (r CitationCitedArtifactClassificationWhoClassified) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactClassificationWhoClassified
	switch other := other.(type) {
	case CitationCitedArtifactClassificationWhoClassified:
		o = &other
	case *CitationCitedArtifactClassificationWhoClassified:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactClassificationWhoClassified) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactClassificationWhoClassified) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Person",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Organization",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Publisher",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "ClassifierCopyright",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "FreeToShare",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactClassificationWhoClassified",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactContributorship) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "complete") {
		if r.Complete != nil {
			children = append(children, *r.Complete)
		}
	}
	if len(name) == 0 || slices.Contains(name, "entry") {
		for _, v := range r.Entry {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "summary") {
		for _, v := range r.Summary {
			children = append(children, v)
		}
	}
	return children
}
func (r CitationCitedArtifactContributorship) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactContributorship to Boolean")
}
func (r CitationCitedArtifactContributorship) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactContributorship to String")
}
func (r CitationCitedArtifactContributorship) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactContributorship to Integer")
}
func (r CitationCitedArtifactContributorship) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactContributorship to Decimal")
}
func (r CitationCitedArtifactContributorship) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactContributorship to Date")
}
func (r CitationCitedArtifactContributorship) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactContributorship to Time")
}
func (r CitationCitedArtifactContributorship) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactContributorship to DateTime")
}
func (r CitationCitedArtifactContributorship) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactContributorship to Quantity")
}
func (r CitationCitedArtifactContributorship) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactContributorship
	switch other := other.(type) {
	case CitationCitedArtifactContributorship:
		o = &other
	case *CitationCitedArtifactContributorship:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactContributorship) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactContributorship) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Complete",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Entry",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactContributorshipEntry",
				Namespace: "FHIR",
			},
		}, {
			Name: "Summary",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactContributorshipSummary",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactContributorship",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactContributorshipEntry) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "name") {
		if r.Name != nil {
			children = append(children, *r.Name)
		}
	}
	if len(name) == 0 || slices.Contains(name, "initials") {
		if r.Initials != nil {
			children = append(children, *r.Initials)
		}
	}
	if len(name) == 0 || slices.Contains(name, "collectiveName") {
		if r.CollectiveName != nil {
			children = append(children, *r.CollectiveName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "affiliationInfo") {
		for _, v := range r.AffiliationInfo {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "address") {
		for _, v := range r.Address {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "telecom") {
		for _, v := range r.Telecom {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contributionType") {
		for _, v := range r.ContributionType {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "role") {
		if r.Role != nil {
			children = append(children, *r.Role)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contributionInstance") {
		for _, v := range r.ContributionInstance {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "correspondingContact") {
		if r.CorrespondingContact != nil {
			children = append(children, *r.CorrespondingContact)
		}
	}
	if len(name) == 0 || slices.Contains(name, "listOrder") {
		if r.ListOrder != nil {
			children = append(children, *r.ListOrder)
		}
	}
	return children
}
func (r CitationCitedArtifactContributorshipEntry) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Boolean")
}
func (r CitationCitedArtifactContributorshipEntry) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to String")
}
func (r CitationCitedArtifactContributorshipEntry) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Integer")
}
func (r CitationCitedArtifactContributorshipEntry) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Decimal")
}
func (r CitationCitedArtifactContributorshipEntry) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Date")
}
func (r CitationCitedArtifactContributorshipEntry) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Time")
}
func (r CitationCitedArtifactContributorshipEntry) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to DateTime")
}
func (r CitationCitedArtifactContributorshipEntry) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntry to Quantity")
}
func (r CitationCitedArtifactContributorshipEntry) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactContributorshipEntry
	switch other := other.(type) {
	case CitationCitedArtifactContributorshipEntry:
		o = &other
	case *CitationCitedArtifactContributorshipEntry:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactContributorshipEntry) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactContributorshipEntry) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Name",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "HumanName",
				Namespace: "FHIR",
			},
		}, {
			Name: "Initials",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "CollectiveName",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "AffiliationInfo",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactContributorshipEntryAffiliationInfo",
				Namespace: "FHIR",
			},
		}, {
			Name: "Address",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Address",
				Namespace: "FHIR",
			},
		}, {
			Name: "Telecom",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ContactPoint",
				Namespace: "FHIR",
			},
		}, {
			Name: "ContributionType",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Role",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ContributionInstance",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CitationCitedArtifactContributorshipEntryContributionInstance",
				Namespace: "FHIR",
			},
		}, {
			Name: "CorrespondingContact",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "ListOrder",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactContributorshipEntry",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "affiliation") {
		if r.Affiliation != nil {
			children = append(children, *r.Affiliation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "role") {
		if r.Role != nil {
			children = append(children, *r.Role)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	return children
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Boolean")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to String")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Integer")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Decimal")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Date")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Time")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to DateTime")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryAffiliationInfo to Quantity")
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactContributorshipEntryAffiliationInfo
	switch other := other.(type) {
	case CitationCitedArtifactContributorshipEntryAffiliationInfo:
		o = &other
	case *CitationCitedArtifactContributorshipEntryAffiliationInfo:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactContributorshipEntryAffiliationInfo) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Affiliation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Role",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactContributorshipEntryAffiliationInfo",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "time") {
		if r.Time != nil {
			children = append(children, *r.Time)
		}
	}
	return children
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Boolean")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to String")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Integer")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Decimal")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Date")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Time")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to DateTime")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactContributorshipEntryContributionInstance to Quantity")
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactContributorshipEntryContributionInstance
	switch other := other.(type) {
	case CitationCitedArtifactContributorshipEntryContributionInstance:
		o = &other
	case *CitationCitedArtifactContributorshipEntryContributionInstance:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactContributorshipEntryContributionInstance) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Time",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactContributorshipEntryContributionInstance",
			Namespace: "FHIR",
		},
	}
}
func (r CitationCitedArtifactContributorshipSummary) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "style") {
		if r.Style != nil {
			children = append(children, *r.Style)
		}
	}
	if len(name) == 0 || slices.Contains(name, "source") {
		if r.Source != nil {
			children = append(children, *r.Source)
		}
	}
	if len(name) == 0 || slices.Contains(name, "value") {
		children = append(children, r.Value)
	}
	return children
}
func (r CitationCitedArtifactContributorshipSummary) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Boolean")
}
func (r CitationCitedArtifactContributorshipSummary) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to String")
}
func (r CitationCitedArtifactContributorshipSummary) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Integer")
}
func (r CitationCitedArtifactContributorshipSummary) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Decimal")
}
func (r CitationCitedArtifactContributorshipSummary) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Date")
}
func (r CitationCitedArtifactContributorshipSummary) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Time")
}
func (r CitationCitedArtifactContributorshipSummary) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to DateTime")
}
func (r CitationCitedArtifactContributorshipSummary) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert CitationCitedArtifactContributorshipSummary to Quantity")
}
func (r CitationCitedArtifactContributorshipSummary) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *CitationCitedArtifactContributorshipSummary
	switch other := other.(type) {
	case CitationCitedArtifactContributorshipSummary:
		o = &other
	case *CitationCitedArtifactContributorshipSummary:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r CitationCitedArtifactContributorshipSummary) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r CitationCitedArtifactContributorshipSummary) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Style",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Source",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Value",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CitationCitedArtifactContributorshipSummary",
			Namespace: "FHIR",
		},
	}
}
