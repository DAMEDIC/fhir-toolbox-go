package capabilitiesR4B

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	search "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	r4b "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
)

type ConcreteAPI interface {
	AccountRead
	AccountSearch
	ActivityDefinitionRead
	ActivityDefinitionSearch
	AdministrableProductDefinitionRead
	AdministrableProductDefinitionSearch
	AdverseEventRead
	AdverseEventSearch
	AllergyIntoleranceRead
	AllergyIntoleranceSearch
	AppointmentRead
	AppointmentSearch
	AppointmentResponseRead
	AppointmentResponseSearch
	AuditEventRead
	AuditEventSearch
	BasicRead
	BasicSearch
	BinaryRead
	BinarySearch
	BiologicallyDerivedProductRead
	BiologicallyDerivedProductSearch
	BodyStructureRead
	BodyStructureSearch
	BundleRead
	BundleSearch
	CapabilityStatementRead
	CapabilityStatementSearch
	CarePlanRead
	CarePlanSearch
	CareTeamRead
	CareTeamSearch
	CatalogEntryRead
	CatalogEntrySearch
	ChargeItemRead
	ChargeItemSearch
	ChargeItemDefinitionRead
	ChargeItemDefinitionSearch
	CitationRead
	CitationSearch
	ClaimRead
	ClaimSearch
	ClaimResponseRead
	ClaimResponseSearch
	ClinicalImpressionRead
	ClinicalImpressionSearch
	ClinicalUseDefinitionRead
	ClinicalUseDefinitionSearch
	CodeSystemRead
	CodeSystemSearch
	CommunicationRead
	CommunicationSearch
	CommunicationRequestRead
	CommunicationRequestSearch
	CompartmentDefinitionRead
	CompartmentDefinitionSearch
	CompositionRead
	CompositionSearch
	ConceptMapRead
	ConceptMapSearch
	ConditionRead
	ConditionSearch
	ConsentRead
	ConsentSearch
	ContractRead
	ContractSearch
	CoverageRead
	CoverageSearch
	CoverageEligibilityRequestRead
	CoverageEligibilityRequestSearch
	CoverageEligibilityResponseRead
	CoverageEligibilityResponseSearch
	DetectedIssueRead
	DetectedIssueSearch
	DeviceRead
	DeviceSearch
	DeviceDefinitionRead
	DeviceDefinitionSearch
	DeviceMetricRead
	DeviceMetricSearch
	DeviceRequestRead
	DeviceRequestSearch
	DeviceUseStatementRead
	DeviceUseStatementSearch
	DiagnosticReportRead
	DiagnosticReportSearch
	DocumentManifestRead
	DocumentManifestSearch
	DocumentReferenceRead
	DocumentReferenceSearch
	EncounterRead
	EncounterSearch
	EndpointRead
	EndpointSearch
	EnrollmentRequestRead
	EnrollmentRequestSearch
	EnrollmentResponseRead
	EnrollmentResponseSearch
	EpisodeOfCareRead
	EpisodeOfCareSearch
	EventDefinitionRead
	EventDefinitionSearch
	EvidenceRead
	EvidenceSearch
	EvidenceReportRead
	EvidenceReportSearch
	EvidenceVariableRead
	EvidenceVariableSearch
	ExampleScenarioRead
	ExampleScenarioSearch
	ExplanationOfBenefitRead
	ExplanationOfBenefitSearch
	FamilyMemberHistoryRead
	FamilyMemberHistorySearch
	FlagRead
	FlagSearch
	GoalRead
	GoalSearch
	GraphDefinitionRead
	GraphDefinitionSearch
	GroupRead
	GroupSearch
	GuidanceResponseRead
	GuidanceResponseSearch
	HealthcareServiceRead
	HealthcareServiceSearch
	ImagingStudyRead
	ImagingStudySearch
	ImmunizationRead
	ImmunizationSearch
	ImmunizationEvaluationRead
	ImmunizationEvaluationSearch
	ImmunizationRecommendationRead
	ImmunizationRecommendationSearch
	ImplementationGuideRead
	ImplementationGuideSearch
	IngredientRead
	IngredientSearch
	InsurancePlanRead
	InsurancePlanSearch
	InvoiceRead
	InvoiceSearch
	LibraryRead
	LibrarySearch
	LinkageRead
	LinkageSearch
	ListRead
	ListSearch
	LocationRead
	LocationSearch
	ManufacturedItemDefinitionRead
	ManufacturedItemDefinitionSearch
	MeasureRead
	MeasureSearch
	MeasureReportRead
	MeasureReportSearch
	MediaRead
	MediaSearch
	MedicationRead
	MedicationSearch
	MedicationAdministrationRead
	MedicationAdministrationSearch
	MedicationDispenseRead
	MedicationDispenseSearch
	MedicationKnowledgeRead
	MedicationKnowledgeSearch
	MedicationRequestRead
	MedicationRequestSearch
	MedicationStatementRead
	MedicationStatementSearch
	MedicinalProductDefinitionRead
	MedicinalProductDefinitionSearch
	MessageDefinitionRead
	MessageDefinitionSearch
	MessageHeaderRead
	MessageHeaderSearch
	MolecularSequenceRead
	MolecularSequenceSearch
	NamingSystemRead
	NamingSystemSearch
	NutritionOrderRead
	NutritionOrderSearch
	NutritionProductRead
	NutritionProductSearch
	ObservationRead
	ObservationSearch
	ObservationDefinitionRead
	ObservationDefinitionSearch
	OperationDefinitionRead
	OperationDefinitionSearch
	OperationOutcomeRead
	OperationOutcomeSearch
	OrganizationRead
	OrganizationSearch
	OrganizationAffiliationRead
	OrganizationAffiliationSearch
	PackagedProductDefinitionRead
	PackagedProductDefinitionSearch
	ParametersRead
	ParametersSearch
	PatientRead
	PatientSearch
	PaymentNoticeRead
	PaymentNoticeSearch
	PaymentReconciliationRead
	PaymentReconciliationSearch
	PersonRead
	PersonSearch
	PlanDefinitionRead
	PlanDefinitionSearch
	PractitionerRead
	PractitionerSearch
	PractitionerRoleRead
	PractitionerRoleSearch
	ProcedureRead
	ProcedureSearch
	ProvenanceRead
	ProvenanceSearch
	QuestionnaireRead
	QuestionnaireSearch
	QuestionnaireResponseRead
	QuestionnaireResponseSearch
	RegulatedAuthorizationRead
	RegulatedAuthorizationSearch
	RelatedPersonRead
	RelatedPersonSearch
	RequestGroupRead
	RequestGroupSearch
	ResearchDefinitionRead
	ResearchDefinitionSearch
	ResearchElementDefinitionRead
	ResearchElementDefinitionSearch
	ResearchStudyRead
	ResearchStudySearch
	ResearchSubjectRead
	ResearchSubjectSearch
	RiskAssessmentRead
	RiskAssessmentSearch
	ScheduleRead
	ScheduleSearch
	SearchParameterRead
	SearchParameterSearch
	ServiceRequestRead
	ServiceRequestSearch
	SlotRead
	SlotSearch
	SpecimenRead
	SpecimenSearch
	SpecimenDefinitionRead
	SpecimenDefinitionSearch
	StructureDefinitionRead
	StructureDefinitionSearch
	StructureMapRead
	StructureMapSearch
	SubscriptionRead
	SubscriptionSearch
	SubscriptionStatusRead
	SubscriptionStatusSearch
	SubscriptionTopicRead
	SubscriptionTopicSearch
	SubstanceRead
	SubstanceSearch
	SubstanceDefinitionRead
	SubstanceDefinitionSearch
	SupplyDeliveryRead
	SupplyDeliverySearch
	SupplyRequestRead
	SupplyRequestSearch
	TaskRead
	TaskSearch
	TerminologyCapabilitiesRead
	TerminologyCapabilitiesSearch
	TestReportRead
	TestReportSearch
	TestScriptRead
	TestScriptSearch
	ValueSetRead
	ValueSetSearch
	VerificationResultRead
	VerificationResultSearch
	VisionPrescriptionRead
	VisionPrescriptionSearch
}
type Concrete struct {
	Generic capabilities.GenericAPI
}

func (w Concrete) ReadAccount(ctx context.Context, id string) (r4b.Account, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Account", id)
	if err != nil {
		return r4b.Account{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Account)
	if !ok {
		return r4b.Account{}, capabilities.InvalidResourceError{ResourceType: "Account"}
	}
	return impl, nil
}
func (w Concrete) ReadActivityDefinition(ctx context.Context, id string) (r4b.ActivityDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ActivityDefinition", id)
	if err != nil {
		return r4b.ActivityDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ActivityDefinition)
	if !ok {
		return r4b.ActivityDefinition{}, capabilities.InvalidResourceError{ResourceType: "ActivityDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadAdministrableProductDefinition(ctx context.Context, id string) (r4b.AdministrableProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AdministrableProductDefinition", id)
	if err != nil {
		return r4b.AdministrableProductDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.AdministrableProductDefinition)
	if !ok {
		return r4b.AdministrableProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "AdministrableProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadAdverseEvent(ctx context.Context, id string) (r4b.AdverseEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AdverseEvent", id)
	if err != nil {
		return r4b.AdverseEvent{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.AdverseEvent)
	if !ok {
		return r4b.AdverseEvent{}, capabilities.InvalidResourceError{ResourceType: "AdverseEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadAllergyIntolerance(ctx context.Context, id string) (r4b.AllergyIntolerance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AllergyIntolerance", id)
	if err != nil {
		return r4b.AllergyIntolerance{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.AllergyIntolerance)
	if !ok {
		return r4b.AllergyIntolerance{}, capabilities.InvalidResourceError{ResourceType: "AllergyIntolerance"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointment(ctx context.Context, id string) (r4b.Appointment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Appointment", id)
	if err != nil {
		return r4b.Appointment{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Appointment)
	if !ok {
		return r4b.Appointment{}, capabilities.InvalidResourceError{ResourceType: "Appointment"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointmentResponse(ctx context.Context, id string) (r4b.AppointmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AppointmentResponse", id)
	if err != nil {
		return r4b.AppointmentResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.AppointmentResponse)
	if !ok {
		return r4b.AppointmentResponse{}, capabilities.InvalidResourceError{ResourceType: "AppointmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadAuditEvent(ctx context.Context, id string) (r4b.AuditEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AuditEvent", id)
	if err != nil {
		return r4b.AuditEvent{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.AuditEvent)
	if !ok {
		return r4b.AuditEvent{}, capabilities.InvalidResourceError{ResourceType: "AuditEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadBasic(ctx context.Context, id string) (r4b.Basic, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Basic", id)
	if err != nil {
		return r4b.Basic{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Basic)
	if !ok {
		return r4b.Basic{}, capabilities.InvalidResourceError{ResourceType: "Basic"}
	}
	return impl, nil
}
func (w Concrete) ReadBinary(ctx context.Context, id string) (r4b.Binary, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Binary", id)
	if err != nil {
		return r4b.Binary{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Binary)
	if !ok {
		return r4b.Binary{}, capabilities.InvalidResourceError{ResourceType: "Binary"}
	}
	return impl, nil
}
func (w Concrete) ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4b.BiologicallyDerivedProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BiologicallyDerivedProduct", id)
	if err != nil {
		return r4b.BiologicallyDerivedProduct{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.BiologicallyDerivedProduct)
	if !ok {
		return r4b.BiologicallyDerivedProduct{}, capabilities.InvalidResourceError{ResourceType: "BiologicallyDerivedProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadBodyStructure(ctx context.Context, id string) (r4b.BodyStructure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BodyStructure", id)
	if err != nil {
		return r4b.BodyStructure{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.BodyStructure)
	if !ok {
		return r4b.BodyStructure{}, capabilities.InvalidResourceError{ResourceType: "BodyStructure"}
	}
	return impl, nil
}
func (w Concrete) ReadBundle(ctx context.Context, id string) (r4b.Bundle, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Bundle", id)
	if err != nil {
		return r4b.Bundle{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Bundle)
	if !ok {
		return r4b.Bundle{}, capabilities.InvalidResourceError{ResourceType: "Bundle"}
	}
	return impl, nil
}
func (w Concrete) ReadCapabilityStatement(ctx context.Context, id string) (r4b.CapabilityStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CapabilityStatement", id)
	if err != nil {
		return r4b.CapabilityStatement{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CapabilityStatement)
	if !ok {
		return r4b.CapabilityStatement{}, capabilities.InvalidResourceError{ResourceType: "CapabilityStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadCarePlan(ctx context.Context, id string) (r4b.CarePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CarePlan", id)
	if err != nil {
		return r4b.CarePlan{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CarePlan)
	if !ok {
		return r4b.CarePlan{}, capabilities.InvalidResourceError{ResourceType: "CarePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadCareTeam(ctx context.Context, id string) (r4b.CareTeam, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CareTeam", id)
	if err != nil {
		return r4b.CareTeam{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CareTeam)
	if !ok {
		return r4b.CareTeam{}, capabilities.InvalidResourceError{ResourceType: "CareTeam"}
	}
	return impl, nil
}
func (w Concrete) ReadCatalogEntry(ctx context.Context, id string) (r4b.CatalogEntry, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CatalogEntry", id)
	if err != nil {
		return r4b.CatalogEntry{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CatalogEntry)
	if !ok {
		return r4b.CatalogEntry{}, capabilities.InvalidResourceError{ResourceType: "CatalogEntry"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItem(ctx context.Context, id string) (r4b.ChargeItem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItem", id)
	if err != nil {
		return r4b.ChargeItem{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ChargeItem)
	if !ok {
		return r4b.ChargeItem{}, capabilities.InvalidResourceError{ResourceType: "ChargeItem"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItemDefinition(ctx context.Context, id string) (r4b.ChargeItemDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItemDefinition", id)
	if err != nil {
		return r4b.ChargeItemDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ChargeItemDefinition)
	if !ok {
		return r4b.ChargeItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ChargeItemDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadCitation(ctx context.Context, id string) (r4b.Citation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Citation", id)
	if err != nil {
		return r4b.Citation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Citation)
	if !ok {
		return r4b.Citation{}, capabilities.InvalidResourceError{ResourceType: "Citation"}
	}
	return impl, nil
}
func (w Concrete) ReadClaim(ctx context.Context, id string) (r4b.Claim, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Claim", id)
	if err != nil {
		return r4b.Claim{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Claim)
	if !ok {
		return r4b.Claim{}, capabilities.InvalidResourceError{ResourceType: "Claim"}
	}
	return impl, nil
}
func (w Concrete) ReadClaimResponse(ctx context.Context, id string) (r4b.ClaimResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClaimResponse", id)
	if err != nil {
		return r4b.ClaimResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ClaimResponse)
	if !ok {
		return r4b.ClaimResponse{}, capabilities.InvalidResourceError{ResourceType: "ClaimResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadClinicalImpression(ctx context.Context, id string) (r4b.ClinicalImpression, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClinicalImpression", id)
	if err != nil {
		return r4b.ClinicalImpression{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ClinicalImpression)
	if !ok {
		return r4b.ClinicalImpression{}, capabilities.InvalidResourceError{ResourceType: "ClinicalImpression"}
	}
	return impl, nil
}
func (w Concrete) ReadClinicalUseDefinition(ctx context.Context, id string) (r4b.ClinicalUseDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClinicalUseDefinition", id)
	if err != nil {
		return r4b.ClinicalUseDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ClinicalUseDefinition)
	if !ok {
		return r4b.ClinicalUseDefinition{}, capabilities.InvalidResourceError{ResourceType: "ClinicalUseDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadCodeSystem(ctx context.Context, id string) (r4b.CodeSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CodeSystem", id)
	if err != nil {
		return r4b.CodeSystem{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CodeSystem)
	if !ok {
		return r4b.CodeSystem{}, capabilities.InvalidResourceError{ResourceType: "CodeSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunication(ctx context.Context, id string) (r4b.Communication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Communication", id)
	if err != nil {
		return r4b.Communication{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Communication)
	if !ok {
		return r4b.Communication{}, capabilities.InvalidResourceError{ResourceType: "Communication"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunicationRequest(ctx context.Context, id string) (r4b.CommunicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CommunicationRequest", id)
	if err != nil {
		return r4b.CommunicationRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CommunicationRequest)
	if !ok {
		return r4b.CommunicationRequest{}, capabilities.InvalidResourceError{ResourceType: "CommunicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCompartmentDefinition(ctx context.Context, id string) (r4b.CompartmentDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CompartmentDefinition", id)
	if err != nil {
		return r4b.CompartmentDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CompartmentDefinition)
	if !ok {
		return r4b.CompartmentDefinition{}, capabilities.InvalidResourceError{ResourceType: "CompartmentDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadComposition(ctx context.Context, id string) (r4b.Composition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Composition", id)
	if err != nil {
		return r4b.Composition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Composition)
	if !ok {
		return r4b.Composition{}, capabilities.InvalidResourceError{ResourceType: "Composition"}
	}
	return impl, nil
}
func (w Concrete) ReadConceptMap(ctx context.Context, id string) (r4b.ConceptMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ConceptMap", id)
	if err != nil {
		return r4b.ConceptMap{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ConceptMap)
	if !ok {
		return r4b.ConceptMap{}, capabilities.InvalidResourceError{ResourceType: "ConceptMap"}
	}
	return impl, nil
}
func (w Concrete) ReadCondition(ctx context.Context, id string) (r4b.Condition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Condition", id)
	if err != nil {
		return r4b.Condition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Condition)
	if !ok {
		return r4b.Condition{}, capabilities.InvalidResourceError{ResourceType: "Condition"}
	}
	return impl, nil
}
func (w Concrete) ReadConsent(ctx context.Context, id string) (r4b.Consent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Consent", id)
	if err != nil {
		return r4b.Consent{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Consent)
	if !ok {
		return r4b.Consent{}, capabilities.InvalidResourceError{ResourceType: "Consent"}
	}
	return impl, nil
}
func (w Concrete) ReadContract(ctx context.Context, id string) (r4b.Contract, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Contract", id)
	if err != nil {
		return r4b.Contract{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Contract)
	if !ok {
		return r4b.Contract{}, capabilities.InvalidResourceError{ResourceType: "Contract"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverage(ctx context.Context, id string) (r4b.Coverage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Coverage", id)
	if err != nil {
		return r4b.Coverage{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Coverage)
	if !ok {
		return r4b.Coverage{}, capabilities.InvalidResourceError{ResourceType: "Coverage"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4b.CoverageEligibilityRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityRequest", id)
	if err != nil {
		return r4b.CoverageEligibilityRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CoverageEligibilityRequest)
	if !ok {
		return r4b.CoverageEligibilityRequest{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4b.CoverageEligibilityResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityResponse", id)
	if err != nil {
		return r4b.CoverageEligibilityResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.CoverageEligibilityResponse)
	if !ok {
		return r4b.CoverageEligibilityResponse{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadDetectedIssue(ctx context.Context, id string) (r4b.DetectedIssue, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DetectedIssue", id)
	if err != nil {
		return r4b.DetectedIssue{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DetectedIssue)
	if !ok {
		return r4b.DetectedIssue{}, capabilities.InvalidResourceError{ResourceType: "DetectedIssue"}
	}
	return impl, nil
}
func (w Concrete) ReadDevice(ctx context.Context, id string) (r4b.Device, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Device", id)
	if err != nil {
		return r4b.Device{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Device)
	if !ok {
		return r4b.Device{}, capabilities.InvalidResourceError{ResourceType: "Device"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceDefinition(ctx context.Context, id string) (r4b.DeviceDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceDefinition", id)
	if err != nil {
		return r4b.DeviceDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DeviceDefinition)
	if !ok {
		return r4b.DeviceDefinition{}, capabilities.InvalidResourceError{ResourceType: "DeviceDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceMetric(ctx context.Context, id string) (r4b.DeviceMetric, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceMetric", id)
	if err != nil {
		return r4b.DeviceMetric{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DeviceMetric)
	if !ok {
		return r4b.DeviceMetric{}, capabilities.InvalidResourceError{ResourceType: "DeviceMetric"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceRequest(ctx context.Context, id string) (r4b.DeviceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceRequest", id)
	if err != nil {
		return r4b.DeviceRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DeviceRequest)
	if !ok {
		return r4b.DeviceRequest{}, capabilities.InvalidResourceError{ResourceType: "DeviceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceUseStatement(ctx context.Context, id string) (r4b.DeviceUseStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceUseStatement", id)
	if err != nil {
		return r4b.DeviceUseStatement{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DeviceUseStatement)
	if !ok {
		return r4b.DeviceUseStatement{}, capabilities.InvalidResourceError{ResourceType: "DeviceUseStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadDiagnosticReport(ctx context.Context, id string) (r4b.DiagnosticReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DiagnosticReport", id)
	if err != nil {
		return r4b.DiagnosticReport{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DiagnosticReport)
	if !ok {
		return r4b.DiagnosticReport{}, capabilities.InvalidResourceError{ResourceType: "DiagnosticReport"}
	}
	return impl, nil
}
func (w Concrete) ReadDocumentManifest(ctx context.Context, id string) (r4b.DocumentManifest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DocumentManifest", id)
	if err != nil {
		return r4b.DocumentManifest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DocumentManifest)
	if !ok {
		return r4b.DocumentManifest{}, capabilities.InvalidResourceError{ResourceType: "DocumentManifest"}
	}
	return impl, nil
}
func (w Concrete) ReadDocumentReference(ctx context.Context, id string) (r4b.DocumentReference, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DocumentReference", id)
	if err != nil {
		return r4b.DocumentReference{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.DocumentReference)
	if !ok {
		return r4b.DocumentReference{}, capabilities.InvalidResourceError{ResourceType: "DocumentReference"}
	}
	return impl, nil
}
func (w Concrete) ReadEncounter(ctx context.Context, id string) (r4b.Encounter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Encounter", id)
	if err != nil {
		return r4b.Encounter{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Encounter)
	if !ok {
		return r4b.Encounter{}, capabilities.InvalidResourceError{ResourceType: "Encounter"}
	}
	return impl, nil
}
func (w Concrete) ReadEndpoint(ctx context.Context, id string) (r4b.Endpoint, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Endpoint", id)
	if err != nil {
		return r4b.Endpoint{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Endpoint)
	if !ok {
		return r4b.Endpoint{}, capabilities.InvalidResourceError{ResourceType: "Endpoint"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentRequest(ctx context.Context, id string) (r4b.EnrollmentRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentRequest", id)
	if err != nil {
		return r4b.EnrollmentRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EnrollmentRequest)
	if !ok {
		return r4b.EnrollmentRequest{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentResponse(ctx context.Context, id string) (r4b.EnrollmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentResponse", id)
	if err != nil {
		return r4b.EnrollmentResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EnrollmentResponse)
	if !ok {
		return r4b.EnrollmentResponse{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadEpisodeOfCare(ctx context.Context, id string) (r4b.EpisodeOfCare, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EpisodeOfCare", id)
	if err != nil {
		return r4b.EpisodeOfCare{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EpisodeOfCare)
	if !ok {
		return r4b.EpisodeOfCare{}, capabilities.InvalidResourceError{ResourceType: "EpisodeOfCare"}
	}
	return impl, nil
}
func (w Concrete) ReadEventDefinition(ctx context.Context, id string) (r4b.EventDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EventDefinition", id)
	if err != nil {
		return r4b.EventDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EventDefinition)
	if !ok {
		return r4b.EventDefinition{}, capabilities.InvalidResourceError{ResourceType: "EventDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidence(ctx context.Context, id string) (r4b.Evidence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Evidence", id)
	if err != nil {
		return r4b.Evidence{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Evidence)
	if !ok {
		return r4b.Evidence{}, capabilities.InvalidResourceError{ResourceType: "Evidence"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidenceReport(ctx context.Context, id string) (r4b.EvidenceReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EvidenceReport", id)
	if err != nil {
		return r4b.EvidenceReport{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EvidenceReport)
	if !ok {
		return r4b.EvidenceReport{}, capabilities.InvalidResourceError{ResourceType: "EvidenceReport"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidenceVariable(ctx context.Context, id string) (r4b.EvidenceVariable, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EvidenceVariable", id)
	if err != nil {
		return r4b.EvidenceVariable{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.EvidenceVariable)
	if !ok {
		return r4b.EvidenceVariable{}, capabilities.InvalidResourceError{ResourceType: "EvidenceVariable"}
	}
	return impl, nil
}
func (w Concrete) ReadExampleScenario(ctx context.Context, id string) (r4b.ExampleScenario, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExampleScenario", id)
	if err != nil {
		return r4b.ExampleScenario{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ExampleScenario)
	if !ok {
		return r4b.ExampleScenario{}, capabilities.InvalidResourceError{ResourceType: "ExampleScenario"}
	}
	return impl, nil
}
func (w Concrete) ReadExplanationOfBenefit(ctx context.Context, id string) (r4b.ExplanationOfBenefit, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExplanationOfBenefit", id)
	if err != nil {
		return r4b.ExplanationOfBenefit{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ExplanationOfBenefit)
	if !ok {
		return r4b.ExplanationOfBenefit{}, capabilities.InvalidResourceError{ResourceType: "ExplanationOfBenefit"}
	}
	return impl, nil
}
func (w Concrete) ReadFamilyMemberHistory(ctx context.Context, id string) (r4b.FamilyMemberHistory, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "FamilyMemberHistory", id)
	if err != nil {
		return r4b.FamilyMemberHistory{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.FamilyMemberHistory)
	if !ok {
		return r4b.FamilyMemberHistory{}, capabilities.InvalidResourceError{ResourceType: "FamilyMemberHistory"}
	}
	return impl, nil
}
func (w Concrete) ReadFlag(ctx context.Context, id string) (r4b.Flag, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Flag", id)
	if err != nil {
		return r4b.Flag{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Flag)
	if !ok {
		return r4b.Flag{}, capabilities.InvalidResourceError{ResourceType: "Flag"}
	}
	return impl, nil
}
func (w Concrete) ReadGoal(ctx context.Context, id string) (r4b.Goal, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Goal", id)
	if err != nil {
		return r4b.Goal{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Goal)
	if !ok {
		return r4b.Goal{}, capabilities.InvalidResourceError{ResourceType: "Goal"}
	}
	return impl, nil
}
func (w Concrete) ReadGraphDefinition(ctx context.Context, id string) (r4b.GraphDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GraphDefinition", id)
	if err != nil {
		return r4b.GraphDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.GraphDefinition)
	if !ok {
		return r4b.GraphDefinition{}, capabilities.InvalidResourceError{ResourceType: "GraphDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadGroup(ctx context.Context, id string) (r4b.Group, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Group", id)
	if err != nil {
		return r4b.Group{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Group)
	if !ok {
		return r4b.Group{}, capabilities.InvalidResourceError{ResourceType: "Group"}
	}
	return impl, nil
}
func (w Concrete) ReadGuidanceResponse(ctx context.Context, id string) (r4b.GuidanceResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GuidanceResponse", id)
	if err != nil {
		return r4b.GuidanceResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.GuidanceResponse)
	if !ok {
		return r4b.GuidanceResponse{}, capabilities.InvalidResourceError{ResourceType: "GuidanceResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadHealthcareService(ctx context.Context, id string) (r4b.HealthcareService, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "HealthcareService", id)
	if err != nil {
		return r4b.HealthcareService{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.HealthcareService)
	if !ok {
		return r4b.HealthcareService{}, capabilities.InvalidResourceError{ResourceType: "HealthcareService"}
	}
	return impl, nil
}
func (w Concrete) ReadImagingStudy(ctx context.Context, id string) (r4b.ImagingStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImagingStudy", id)
	if err != nil {
		return r4b.ImagingStudy{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ImagingStudy)
	if !ok {
		return r4b.ImagingStudy{}, capabilities.InvalidResourceError{ResourceType: "ImagingStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunization(ctx context.Context, id string) (r4b.Immunization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Immunization", id)
	if err != nil {
		return r4b.Immunization{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Immunization)
	if !ok {
		return r4b.Immunization{}, capabilities.InvalidResourceError{ResourceType: "Immunization"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationEvaluation(ctx context.Context, id string) (r4b.ImmunizationEvaluation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationEvaluation", id)
	if err != nil {
		return r4b.ImmunizationEvaluation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ImmunizationEvaluation)
	if !ok {
		return r4b.ImmunizationEvaluation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationEvaluation"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationRecommendation(ctx context.Context, id string) (r4b.ImmunizationRecommendation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationRecommendation", id)
	if err != nil {
		return r4b.ImmunizationRecommendation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ImmunizationRecommendation)
	if !ok {
		return r4b.ImmunizationRecommendation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationRecommendation"}
	}
	return impl, nil
}
func (w Concrete) ReadImplementationGuide(ctx context.Context, id string) (r4b.ImplementationGuide, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImplementationGuide", id)
	if err != nil {
		return r4b.ImplementationGuide{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ImplementationGuide)
	if !ok {
		return r4b.ImplementationGuide{}, capabilities.InvalidResourceError{ResourceType: "ImplementationGuide"}
	}
	return impl, nil
}
func (w Concrete) ReadIngredient(ctx context.Context, id string) (r4b.Ingredient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Ingredient", id)
	if err != nil {
		return r4b.Ingredient{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Ingredient)
	if !ok {
		return r4b.Ingredient{}, capabilities.InvalidResourceError{ResourceType: "Ingredient"}
	}
	return impl, nil
}
func (w Concrete) ReadInsurancePlan(ctx context.Context, id string) (r4b.InsurancePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "InsurancePlan", id)
	if err != nil {
		return r4b.InsurancePlan{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.InsurancePlan)
	if !ok {
		return r4b.InsurancePlan{}, capabilities.InvalidResourceError{ResourceType: "InsurancePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadInvoice(ctx context.Context, id string) (r4b.Invoice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Invoice", id)
	if err != nil {
		return r4b.Invoice{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Invoice)
	if !ok {
		return r4b.Invoice{}, capabilities.InvalidResourceError{ResourceType: "Invoice"}
	}
	return impl, nil
}
func (w Concrete) ReadLibrary(ctx context.Context, id string) (r4b.Library, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Library", id)
	if err != nil {
		return r4b.Library{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Library)
	if !ok {
		return r4b.Library{}, capabilities.InvalidResourceError{ResourceType: "Library"}
	}
	return impl, nil
}
func (w Concrete) ReadLinkage(ctx context.Context, id string) (r4b.Linkage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Linkage", id)
	if err != nil {
		return r4b.Linkage{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Linkage)
	if !ok {
		return r4b.Linkage{}, capabilities.InvalidResourceError{ResourceType: "Linkage"}
	}
	return impl, nil
}
func (w Concrete) ReadList(ctx context.Context, id string) (r4b.List, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "List", id)
	if err != nil {
		return r4b.List{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.List)
	if !ok {
		return r4b.List{}, capabilities.InvalidResourceError{ResourceType: "List"}
	}
	return impl, nil
}
func (w Concrete) ReadLocation(ctx context.Context, id string) (r4b.Location, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Location", id)
	if err != nil {
		return r4b.Location{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Location)
	if !ok {
		return r4b.Location{}, capabilities.InvalidResourceError{ResourceType: "Location"}
	}
	return impl, nil
}
func (w Concrete) ReadManufacturedItemDefinition(ctx context.Context, id string) (r4b.ManufacturedItemDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ManufacturedItemDefinition", id)
	if err != nil {
		return r4b.ManufacturedItemDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ManufacturedItemDefinition)
	if !ok {
		return r4b.ManufacturedItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ManufacturedItemDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasure(ctx context.Context, id string) (r4b.Measure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Measure", id)
	if err != nil {
		return r4b.Measure{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Measure)
	if !ok {
		return r4b.Measure{}, capabilities.InvalidResourceError{ResourceType: "Measure"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasureReport(ctx context.Context, id string) (r4b.MeasureReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MeasureReport", id)
	if err != nil {
		return r4b.MeasureReport{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MeasureReport)
	if !ok {
		return r4b.MeasureReport{}, capabilities.InvalidResourceError{ResourceType: "MeasureReport"}
	}
	return impl, nil
}
func (w Concrete) ReadMedia(ctx context.Context, id string) (r4b.Media, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Media", id)
	if err != nil {
		return r4b.Media{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Media)
	if !ok {
		return r4b.Media{}, capabilities.InvalidResourceError{ResourceType: "Media"}
	}
	return impl, nil
}
func (w Concrete) ReadMedication(ctx context.Context, id string) (r4b.Medication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Medication", id)
	if err != nil {
		return r4b.Medication{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Medication)
	if !ok {
		return r4b.Medication{}, capabilities.InvalidResourceError{ResourceType: "Medication"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationAdministration(ctx context.Context, id string) (r4b.MedicationAdministration, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationAdministration", id)
	if err != nil {
		return r4b.MedicationAdministration{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicationAdministration)
	if !ok {
		return r4b.MedicationAdministration{}, capabilities.InvalidResourceError{ResourceType: "MedicationAdministration"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationDispense(ctx context.Context, id string) (r4b.MedicationDispense, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationDispense", id)
	if err != nil {
		return r4b.MedicationDispense{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicationDispense)
	if !ok {
		return r4b.MedicationDispense{}, capabilities.InvalidResourceError{ResourceType: "MedicationDispense"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationKnowledge(ctx context.Context, id string) (r4b.MedicationKnowledge, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationKnowledge", id)
	if err != nil {
		return r4b.MedicationKnowledge{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicationKnowledge)
	if !ok {
		return r4b.MedicationKnowledge{}, capabilities.InvalidResourceError{ResourceType: "MedicationKnowledge"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationRequest(ctx context.Context, id string) (r4b.MedicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationRequest", id)
	if err != nil {
		return r4b.MedicationRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicationRequest)
	if !ok {
		return r4b.MedicationRequest{}, capabilities.InvalidResourceError{ResourceType: "MedicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationStatement(ctx context.Context, id string) (r4b.MedicationStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationStatement", id)
	if err != nil {
		return r4b.MedicationStatement{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicationStatement)
	if !ok {
		return r4b.MedicationStatement{}, capabilities.InvalidResourceError{ResourceType: "MedicationStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductDefinition(ctx context.Context, id string) (r4b.MedicinalProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductDefinition", id)
	if err != nil {
		return r4b.MedicinalProductDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MedicinalProductDefinition)
	if !ok {
		return r4b.MedicinalProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageDefinition(ctx context.Context, id string) (r4b.MessageDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageDefinition", id)
	if err != nil {
		return r4b.MessageDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MessageDefinition)
	if !ok {
		return r4b.MessageDefinition{}, capabilities.InvalidResourceError{ResourceType: "MessageDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageHeader(ctx context.Context, id string) (r4b.MessageHeader, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageHeader", id)
	if err != nil {
		return r4b.MessageHeader{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MessageHeader)
	if !ok {
		return r4b.MessageHeader{}, capabilities.InvalidResourceError{ResourceType: "MessageHeader"}
	}
	return impl, nil
}
func (w Concrete) ReadMolecularSequence(ctx context.Context, id string) (r4b.MolecularSequence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MolecularSequence", id)
	if err != nil {
		return r4b.MolecularSequence{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.MolecularSequence)
	if !ok {
		return r4b.MolecularSequence{}, capabilities.InvalidResourceError{ResourceType: "MolecularSequence"}
	}
	return impl, nil
}
func (w Concrete) ReadNamingSystem(ctx context.Context, id string) (r4b.NamingSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NamingSystem", id)
	if err != nil {
		return r4b.NamingSystem{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.NamingSystem)
	if !ok {
		return r4b.NamingSystem{}, capabilities.InvalidResourceError{ResourceType: "NamingSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionOrder(ctx context.Context, id string) (r4b.NutritionOrder, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionOrder", id)
	if err != nil {
		return r4b.NutritionOrder{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.NutritionOrder)
	if !ok {
		return r4b.NutritionOrder{}, capabilities.InvalidResourceError{ResourceType: "NutritionOrder"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionProduct(ctx context.Context, id string) (r4b.NutritionProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionProduct", id)
	if err != nil {
		return r4b.NutritionProduct{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.NutritionProduct)
	if !ok {
		return r4b.NutritionProduct{}, capabilities.InvalidResourceError{ResourceType: "NutritionProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadObservation(ctx context.Context, id string) (r4b.Observation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Observation", id)
	if err != nil {
		return r4b.Observation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Observation)
	if !ok {
		return r4b.Observation{}, capabilities.InvalidResourceError{ResourceType: "Observation"}
	}
	return impl, nil
}
func (w Concrete) ReadObservationDefinition(ctx context.Context, id string) (r4b.ObservationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ObservationDefinition", id)
	if err != nil {
		return r4b.ObservationDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ObservationDefinition)
	if !ok {
		return r4b.ObservationDefinition{}, capabilities.InvalidResourceError{ResourceType: "ObservationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationDefinition(ctx context.Context, id string) (r4b.OperationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationDefinition", id)
	if err != nil {
		return r4b.OperationDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.OperationDefinition)
	if !ok {
		return r4b.OperationDefinition{}, capabilities.InvalidResourceError{ResourceType: "OperationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationOutcome(ctx context.Context, id string) (r4b.OperationOutcome, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationOutcome", id)
	if err != nil {
		return r4b.OperationOutcome{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.OperationOutcome)
	if !ok {
		return r4b.OperationOutcome{}, capabilities.InvalidResourceError{ResourceType: "OperationOutcome"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganization(ctx context.Context, id string) (r4b.Organization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Organization", id)
	if err != nil {
		return r4b.Organization{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Organization)
	if !ok {
		return r4b.Organization{}, capabilities.InvalidResourceError{ResourceType: "Organization"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganizationAffiliation(ctx context.Context, id string) (r4b.OrganizationAffiliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OrganizationAffiliation", id)
	if err != nil {
		return r4b.OrganizationAffiliation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.OrganizationAffiliation)
	if !ok {
		return r4b.OrganizationAffiliation{}, capabilities.InvalidResourceError{ResourceType: "OrganizationAffiliation"}
	}
	return impl, nil
}
func (w Concrete) ReadPackagedProductDefinition(ctx context.Context, id string) (r4b.PackagedProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PackagedProductDefinition", id)
	if err != nil {
		return r4b.PackagedProductDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.PackagedProductDefinition)
	if !ok {
		return r4b.PackagedProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "PackagedProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadParameters(ctx context.Context, id string) (r4b.Parameters, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Parameters", id)
	if err != nil {
		return r4b.Parameters{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Parameters)
	if !ok {
		return r4b.Parameters{}, capabilities.InvalidResourceError{ResourceType: "Parameters"}
	}
	return impl, nil
}
func (w Concrete) ReadPatient(ctx context.Context, id string) (r4b.Patient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Patient", id)
	if err != nil {
		return r4b.Patient{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Patient)
	if !ok {
		return r4b.Patient{}, capabilities.InvalidResourceError{ResourceType: "Patient"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentNotice(ctx context.Context, id string) (r4b.PaymentNotice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentNotice", id)
	if err != nil {
		return r4b.PaymentNotice{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.PaymentNotice)
	if !ok {
		return r4b.PaymentNotice{}, capabilities.InvalidResourceError{ResourceType: "PaymentNotice"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentReconciliation(ctx context.Context, id string) (r4b.PaymentReconciliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentReconciliation", id)
	if err != nil {
		return r4b.PaymentReconciliation{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.PaymentReconciliation)
	if !ok {
		return r4b.PaymentReconciliation{}, capabilities.InvalidResourceError{ResourceType: "PaymentReconciliation"}
	}
	return impl, nil
}
func (w Concrete) ReadPerson(ctx context.Context, id string) (r4b.Person, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Person", id)
	if err != nil {
		return r4b.Person{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Person)
	if !ok {
		return r4b.Person{}, capabilities.InvalidResourceError{ResourceType: "Person"}
	}
	return impl, nil
}
func (w Concrete) ReadPlanDefinition(ctx context.Context, id string) (r4b.PlanDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PlanDefinition", id)
	if err != nil {
		return r4b.PlanDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.PlanDefinition)
	if !ok {
		return r4b.PlanDefinition{}, capabilities.InvalidResourceError{ResourceType: "PlanDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitioner(ctx context.Context, id string) (r4b.Practitioner, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Practitioner", id)
	if err != nil {
		return r4b.Practitioner{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Practitioner)
	if !ok {
		return r4b.Practitioner{}, capabilities.InvalidResourceError{ResourceType: "Practitioner"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitionerRole(ctx context.Context, id string) (r4b.PractitionerRole, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PractitionerRole", id)
	if err != nil {
		return r4b.PractitionerRole{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.PractitionerRole)
	if !ok {
		return r4b.PractitionerRole{}, capabilities.InvalidResourceError{ResourceType: "PractitionerRole"}
	}
	return impl, nil
}
func (w Concrete) ReadProcedure(ctx context.Context, id string) (r4b.Procedure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Procedure", id)
	if err != nil {
		return r4b.Procedure{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Procedure)
	if !ok {
		return r4b.Procedure{}, capabilities.InvalidResourceError{ResourceType: "Procedure"}
	}
	return impl, nil
}
func (w Concrete) ReadProvenance(ctx context.Context, id string) (r4b.Provenance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Provenance", id)
	if err != nil {
		return r4b.Provenance{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Provenance)
	if !ok {
		return r4b.Provenance{}, capabilities.InvalidResourceError{ResourceType: "Provenance"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaire(ctx context.Context, id string) (r4b.Questionnaire, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Questionnaire", id)
	if err != nil {
		return r4b.Questionnaire{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Questionnaire)
	if !ok {
		return r4b.Questionnaire{}, capabilities.InvalidResourceError{ResourceType: "Questionnaire"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaireResponse(ctx context.Context, id string) (r4b.QuestionnaireResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "QuestionnaireResponse", id)
	if err != nil {
		return r4b.QuestionnaireResponse{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.QuestionnaireResponse)
	if !ok {
		return r4b.QuestionnaireResponse{}, capabilities.InvalidResourceError{ResourceType: "QuestionnaireResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadRegulatedAuthorization(ctx context.Context, id string) (r4b.RegulatedAuthorization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RegulatedAuthorization", id)
	if err != nil {
		return r4b.RegulatedAuthorization{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.RegulatedAuthorization)
	if !ok {
		return r4b.RegulatedAuthorization{}, capabilities.InvalidResourceError{ResourceType: "RegulatedAuthorization"}
	}
	return impl, nil
}
func (w Concrete) ReadRelatedPerson(ctx context.Context, id string) (r4b.RelatedPerson, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RelatedPerson", id)
	if err != nil {
		return r4b.RelatedPerson{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.RelatedPerson)
	if !ok {
		return r4b.RelatedPerson{}, capabilities.InvalidResourceError{ResourceType: "RelatedPerson"}
	}
	return impl, nil
}
func (w Concrete) ReadRequestGroup(ctx context.Context, id string) (r4b.RequestGroup, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RequestGroup", id)
	if err != nil {
		return r4b.RequestGroup{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.RequestGroup)
	if !ok {
		return r4b.RequestGroup{}, capabilities.InvalidResourceError{ResourceType: "RequestGroup"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchDefinition(ctx context.Context, id string) (r4b.ResearchDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchDefinition", id)
	if err != nil {
		return r4b.ResearchDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ResearchDefinition)
	if !ok {
		return r4b.ResearchDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchElementDefinition(ctx context.Context, id string) (r4b.ResearchElementDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchElementDefinition", id)
	if err != nil {
		return r4b.ResearchElementDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ResearchElementDefinition)
	if !ok {
		return r4b.ResearchElementDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchElementDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchStudy(ctx context.Context, id string) (r4b.ResearchStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchStudy", id)
	if err != nil {
		return r4b.ResearchStudy{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ResearchStudy)
	if !ok {
		return r4b.ResearchStudy{}, capabilities.InvalidResourceError{ResourceType: "ResearchStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchSubject(ctx context.Context, id string) (r4b.ResearchSubject, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchSubject", id)
	if err != nil {
		return r4b.ResearchSubject{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ResearchSubject)
	if !ok {
		return r4b.ResearchSubject{}, capabilities.InvalidResourceError{ResourceType: "ResearchSubject"}
	}
	return impl, nil
}
func (w Concrete) ReadRiskAssessment(ctx context.Context, id string) (r4b.RiskAssessment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RiskAssessment", id)
	if err != nil {
		return r4b.RiskAssessment{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.RiskAssessment)
	if !ok {
		return r4b.RiskAssessment{}, capabilities.InvalidResourceError{ResourceType: "RiskAssessment"}
	}
	return impl, nil
}
func (w Concrete) ReadSchedule(ctx context.Context, id string) (r4b.Schedule, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Schedule", id)
	if err != nil {
		return r4b.Schedule{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Schedule)
	if !ok {
		return r4b.Schedule{}, capabilities.InvalidResourceError{ResourceType: "Schedule"}
	}
	return impl, nil
}
func (w Concrete) ReadSearchParameter(ctx context.Context, id string) (r4b.SearchParameter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SearchParameter", id)
	if err != nil {
		return r4b.SearchParameter{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SearchParameter)
	if !ok {
		return r4b.SearchParameter{}, capabilities.InvalidResourceError{ResourceType: "SearchParameter"}
	}
	return impl, nil
}
func (w Concrete) ReadServiceRequest(ctx context.Context, id string) (r4b.ServiceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ServiceRequest", id)
	if err != nil {
		return r4b.ServiceRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ServiceRequest)
	if !ok {
		return r4b.ServiceRequest{}, capabilities.InvalidResourceError{ResourceType: "ServiceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadSlot(ctx context.Context, id string) (r4b.Slot, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Slot", id)
	if err != nil {
		return r4b.Slot{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Slot)
	if !ok {
		return r4b.Slot{}, capabilities.InvalidResourceError{ResourceType: "Slot"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimen(ctx context.Context, id string) (r4b.Specimen, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Specimen", id)
	if err != nil {
		return r4b.Specimen{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Specimen)
	if !ok {
		return r4b.Specimen{}, capabilities.InvalidResourceError{ResourceType: "Specimen"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimenDefinition(ctx context.Context, id string) (r4b.SpecimenDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SpecimenDefinition", id)
	if err != nil {
		return r4b.SpecimenDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SpecimenDefinition)
	if !ok {
		return r4b.SpecimenDefinition{}, capabilities.InvalidResourceError{ResourceType: "SpecimenDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureDefinition(ctx context.Context, id string) (r4b.StructureDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureDefinition", id)
	if err != nil {
		return r4b.StructureDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.StructureDefinition)
	if !ok {
		return r4b.StructureDefinition{}, capabilities.InvalidResourceError{ResourceType: "StructureDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureMap(ctx context.Context, id string) (r4b.StructureMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureMap", id)
	if err != nil {
		return r4b.StructureMap{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.StructureMap)
	if !ok {
		return r4b.StructureMap{}, capabilities.InvalidResourceError{ResourceType: "StructureMap"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscription(ctx context.Context, id string) (r4b.Subscription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Subscription", id)
	if err != nil {
		return r4b.Subscription{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Subscription)
	if !ok {
		return r4b.Subscription{}, capabilities.InvalidResourceError{ResourceType: "Subscription"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscriptionStatus(ctx context.Context, id string) (r4b.SubscriptionStatus, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubscriptionStatus", id)
	if err != nil {
		return r4b.SubscriptionStatus{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SubscriptionStatus)
	if !ok {
		return r4b.SubscriptionStatus{}, capabilities.InvalidResourceError{ResourceType: "SubscriptionStatus"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscriptionTopic(ctx context.Context, id string) (r4b.SubscriptionTopic, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubscriptionTopic", id)
	if err != nil {
		return r4b.SubscriptionTopic{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SubscriptionTopic)
	if !ok {
		return r4b.SubscriptionTopic{}, capabilities.InvalidResourceError{ResourceType: "SubscriptionTopic"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstance(ctx context.Context, id string) (r4b.Substance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Substance", id)
	if err != nil {
		return r4b.Substance{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Substance)
	if !ok {
		return r4b.Substance{}, capabilities.InvalidResourceError{ResourceType: "Substance"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceDefinition(ctx context.Context, id string) (r4b.SubstanceDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceDefinition", id)
	if err != nil {
		return r4b.SubstanceDefinition{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SubstanceDefinition)
	if !ok {
		return r4b.SubstanceDefinition{}, capabilities.InvalidResourceError{ResourceType: "SubstanceDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyDelivery(ctx context.Context, id string) (r4b.SupplyDelivery, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyDelivery", id)
	if err != nil {
		return r4b.SupplyDelivery{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SupplyDelivery)
	if !ok {
		return r4b.SupplyDelivery{}, capabilities.InvalidResourceError{ResourceType: "SupplyDelivery"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyRequest(ctx context.Context, id string) (r4b.SupplyRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyRequest", id)
	if err != nil {
		return r4b.SupplyRequest{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.SupplyRequest)
	if !ok {
		return r4b.SupplyRequest{}, capabilities.InvalidResourceError{ResourceType: "SupplyRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadTask(ctx context.Context, id string) (r4b.Task, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Task", id)
	if err != nil {
		return r4b.Task{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.Task)
	if !ok {
		return r4b.Task{}, capabilities.InvalidResourceError{ResourceType: "Task"}
	}
	return impl, nil
}
func (w Concrete) ReadTerminologyCapabilities(ctx context.Context, id string) (r4b.TerminologyCapabilities, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TerminologyCapabilities", id)
	if err != nil {
		return r4b.TerminologyCapabilities{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.TerminologyCapabilities)
	if !ok {
		return r4b.TerminologyCapabilities{}, capabilities.InvalidResourceError{ResourceType: "TerminologyCapabilities"}
	}
	return impl, nil
}
func (w Concrete) ReadTestReport(ctx context.Context, id string) (r4b.TestReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestReport", id)
	if err != nil {
		return r4b.TestReport{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.TestReport)
	if !ok {
		return r4b.TestReport{}, capabilities.InvalidResourceError{ResourceType: "TestReport"}
	}
	return impl, nil
}
func (w Concrete) ReadTestScript(ctx context.Context, id string) (r4b.TestScript, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestScript", id)
	if err != nil {
		return r4b.TestScript{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.TestScript)
	if !ok {
		return r4b.TestScript{}, capabilities.InvalidResourceError{ResourceType: "TestScript"}
	}
	return impl, nil
}
func (w Concrete) ReadValueSet(ctx context.Context, id string) (r4b.ValueSet, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ValueSet", id)
	if err != nil {
		return r4b.ValueSet{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.ValueSet)
	if !ok {
		return r4b.ValueSet{}, capabilities.InvalidResourceError{ResourceType: "ValueSet"}
	}
	return impl, nil
}
func (w Concrete) ReadVerificationResult(ctx context.Context, id string) (r4b.VerificationResult, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VerificationResult", id)
	if err != nil {
		return r4b.VerificationResult{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.VerificationResult)
	if !ok {
		return r4b.VerificationResult{}, capabilities.InvalidResourceError{ResourceType: "VerificationResult"}
	}
	return impl, nil
}
func (w Concrete) ReadVisionPrescription(ctx context.Context, id string) (r4b.VisionPrescription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VisionPrescription", id)
	if err != nil {
		return r4b.VisionPrescription{}, err
	}
	contained, ok := v.(r4b.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4b.VisionPrescription)
	if !ok {
		return r4b.VisionPrescription{}, capabilities.InvalidResourceError{ResourceType: "VisionPrescription"}
	}
	return impl, nil
}
func (w Concrete) SearchCapabilitiesAccount() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Account")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAccount(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Account", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesActivityDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ActivityDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchActivityDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ActivityDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAdministrableProductDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("AdministrableProductDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAdministrableProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "AdministrableProductDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAdverseEvent() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("AdverseEvent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAdverseEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "AdverseEvent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAllergyIntolerance() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("AllergyIntolerance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAllergyIntolerance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "AllergyIntolerance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAppointment() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Appointment")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAppointment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Appointment", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAppointmentResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("AppointmentResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAppointmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "AppointmentResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesAuditEvent() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("AuditEvent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchAuditEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "AuditEvent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesBasic() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Basic")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBasic(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Basic", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesBinary() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Binary")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBinary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Binary", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesBiologicallyDerivedProduct() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("BiologicallyDerivedProduct")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBiologicallyDerivedProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "BiologicallyDerivedProduct", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesBodyStructure() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("BodyStructure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBodyStructure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "BodyStructure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesBundle() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Bundle")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBundle(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Bundle", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCapabilityStatement() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CapabilityStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCapabilityStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CapabilityStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCarePlan() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CarePlan")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCarePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CarePlan", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCareTeam() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CareTeam")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCareTeam(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CareTeam", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCatalogEntry() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CatalogEntry")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCatalogEntry(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CatalogEntry", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesChargeItem() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ChargeItem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchChargeItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ChargeItem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesChargeItemDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ChargeItemDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchChargeItemDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ChargeItemDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCitation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Citation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCitation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Citation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesClaim() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Claim")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchClaim(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Claim", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesClaimResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ClaimResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchClaimResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ClaimResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesClinicalImpression() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ClinicalImpression")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchClinicalImpression(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ClinicalImpression", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesClinicalUseDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ClinicalUseDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchClinicalUseDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ClinicalUseDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCodeSystem() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CodeSystem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCodeSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CodeSystem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCommunication() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Communication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCommunication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Communication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCommunicationRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CommunicationRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCommunicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CommunicationRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCompartmentDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CompartmentDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCompartmentDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CompartmentDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesComposition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Composition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchComposition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Composition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesConceptMap() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ConceptMap")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchConceptMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ConceptMap", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCondition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Condition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCondition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Condition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesConsent() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Consent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchConsent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Consent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesContract() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Contract")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchContract(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Contract", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCoverage() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Coverage")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCoverage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Coverage", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCoverageEligibilityRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CoverageEligibilityRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCoverageEligibilityRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CoverageEligibilityRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesCoverageEligibilityResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("CoverageEligibilityResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchCoverageEligibilityResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "CoverageEligibilityResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDetectedIssue() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DetectedIssue")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDetectedIssue(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DetectedIssue", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDevice() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Device")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDevice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Device", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDeviceDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDeviceMetric() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceMetric")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceMetric(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceMetric", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDeviceRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDeviceUseStatement() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceUseStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceUseStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceUseStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDiagnosticReport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DiagnosticReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDiagnosticReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DiagnosticReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDocumentManifest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DocumentManifest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDocumentManifest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DocumentManifest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesDocumentReference() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DocumentReference")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDocumentReference(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DocumentReference", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEncounter() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Encounter")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEncounter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Encounter", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEndpoint() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Endpoint")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEndpoint(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Endpoint", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEnrollmentRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EnrollmentRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEnrollmentRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EnrollmentRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEnrollmentResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EnrollmentResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEnrollmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EnrollmentResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEpisodeOfCare() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EpisodeOfCare")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEpisodeOfCare(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EpisodeOfCare", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEventDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EventDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEventDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EventDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEvidence() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Evidence")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEvidence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Evidence", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEvidenceReport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EvidenceReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEvidenceReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EvidenceReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesEvidenceVariable() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EvidenceVariable")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEvidenceVariable(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EvidenceVariable", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesExampleScenario() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ExampleScenario")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchExampleScenario(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ExampleScenario", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesExplanationOfBenefit() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ExplanationOfBenefit")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchExplanationOfBenefit(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ExplanationOfBenefit", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesFamilyMemberHistory() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("FamilyMemberHistory")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchFamilyMemberHistory(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "FamilyMemberHistory", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesFlag() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Flag")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchFlag(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Flag", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesGoal() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Goal")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchGoal(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Goal", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesGraphDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("GraphDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchGraphDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "GraphDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesGroup() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Group")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchGroup(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Group", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesGuidanceResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("GuidanceResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchGuidanceResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "GuidanceResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesHealthcareService() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("HealthcareService")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchHealthcareService(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "HealthcareService", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesImagingStudy() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ImagingStudy")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImagingStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ImagingStudy", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesImmunization() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Immunization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImmunization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Immunization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesImmunizationEvaluation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ImmunizationEvaluation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImmunizationEvaluation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ImmunizationEvaluation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesImmunizationRecommendation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ImmunizationRecommendation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImmunizationRecommendation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ImmunizationRecommendation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesImplementationGuide() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ImplementationGuide")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImplementationGuide(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ImplementationGuide", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesIngredient() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Ingredient")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchIngredient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Ingredient", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesInsurancePlan() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("InsurancePlan")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchInsurancePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "InsurancePlan", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesInvoice() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Invoice")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchInvoice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Invoice", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesLibrary() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Library")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchLibrary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Library", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesLinkage() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Linkage")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchLinkage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Linkage", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesList() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("List")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchList(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "List", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesLocation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Location")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchLocation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Location", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesManufacturedItemDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ManufacturedItemDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchManufacturedItemDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ManufacturedItemDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMeasure() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Measure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMeasure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Measure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMeasureReport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MeasureReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMeasureReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MeasureReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedia() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Media")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedia(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Media", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedication() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Medication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Medication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicationAdministration() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicationAdministration")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicationAdministration(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicationAdministration", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicationDispense() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicationDispense")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicationDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicationDispense", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicationKnowledge() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicationKnowledge")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicationKnowledge(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicationKnowledge", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicationRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicationRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicationRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicationStatement() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicationStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicationStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicationStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMessageDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MessageDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMessageDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MessageDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMessageHeader() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MessageHeader")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMessageHeader(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MessageHeader", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMolecularSequence() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MolecularSequence")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMolecularSequence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MolecularSequence", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesNamingSystem() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("NamingSystem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchNamingSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "NamingSystem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesNutritionOrder() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("NutritionOrder")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchNutritionOrder(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "NutritionOrder", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesNutritionProduct() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("NutritionProduct")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchNutritionProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "NutritionProduct", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesObservation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Observation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Observation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesObservationDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ObservationDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchObservationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ObservationDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesOperationDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("OperationDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchOperationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "OperationDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesOperationOutcome() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("OperationOutcome")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchOperationOutcome(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "OperationOutcome", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesOrganization() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Organization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchOrganization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Organization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesOrganizationAffiliation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("OrganizationAffiliation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchOrganizationAffiliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "OrganizationAffiliation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPackagedProductDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("PackagedProductDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPackagedProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "PackagedProductDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesParameters() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Parameters")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchParameters(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Parameters", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPatient() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Patient")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Patient", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPaymentNotice() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("PaymentNotice")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPaymentNotice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "PaymentNotice", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPaymentReconciliation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("PaymentReconciliation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPaymentReconciliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "PaymentReconciliation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPerson() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Person")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Person", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPlanDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("PlanDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPlanDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "PlanDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPractitioner() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Practitioner")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPractitioner(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Practitioner", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesPractitionerRole() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("PractitionerRole")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPractitionerRole(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "PractitionerRole", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesProcedure() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Procedure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchProcedure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Procedure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesProvenance() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Provenance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchProvenance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Provenance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesQuestionnaire() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Questionnaire")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchQuestionnaire(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Questionnaire", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesQuestionnaireResponse() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("QuestionnaireResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchQuestionnaireResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "QuestionnaireResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesRegulatedAuthorization() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RegulatedAuthorization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRegulatedAuthorization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RegulatedAuthorization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesRelatedPerson() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RelatedPerson")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRelatedPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RelatedPerson", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesRequestGroup() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RequestGroup")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRequestGroup(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RequestGroup", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesResearchDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ResearchDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchResearchDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ResearchDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesResearchElementDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ResearchElementDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchResearchElementDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ResearchElementDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesResearchStudy() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ResearchStudy")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchResearchStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ResearchStudy", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesResearchSubject() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ResearchSubject")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchResearchSubject(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ResearchSubject", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesRiskAssessment() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RiskAssessment")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRiskAssessment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RiskAssessment", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSchedule() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Schedule")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSchedule(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Schedule", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSearchParameter() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SearchParameter")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSearchParameter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SearchParameter", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesServiceRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ServiceRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchServiceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ServiceRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSlot() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Slot")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSlot(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Slot", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSpecimen() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Specimen")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSpecimen(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Specimen", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSpecimenDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SpecimenDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSpecimenDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SpecimenDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesStructureDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("StructureDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchStructureDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "StructureDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesStructureMap() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("StructureMap")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchStructureMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "StructureMap", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubscription() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Subscription")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubscription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Subscription", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubscriptionStatus() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubscriptionStatus")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubscriptionStatus(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubscriptionStatus", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubscriptionTopic() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubscriptionTopic")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubscriptionTopic(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubscriptionTopic", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstance() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Substance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Substance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstanceDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSupplyDelivery() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SupplyDelivery")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSupplyDelivery(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SupplyDelivery", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSupplyRequest() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SupplyRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSupplyRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SupplyRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesTask() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Task")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTask(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Task", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesTerminologyCapabilities() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("TerminologyCapabilities")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTerminologyCapabilities(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "TerminologyCapabilities", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesTestReport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("TestReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTestReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "TestReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesTestScript() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("TestScript")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTestScript(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "TestScript", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesValueSet() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ValueSet")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchValueSet(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ValueSet", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesVerificationResult() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("VerificationResult")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchVerificationResult(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "VerificationResult", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesVisionPrescription() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("VisionPrescription")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchVisionPrescription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "VisionPrescription", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
