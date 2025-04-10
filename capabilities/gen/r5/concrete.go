package capabilitiesR5

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	search "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	r5 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
)

type ConcreteAPI interface {
	AccountRead
	AccountSearch
	ActivityDefinitionRead
	ActivityDefinitionSearch
	ActorDefinitionRead
	ActorDefinitionSearch
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
	ArtifactAssessmentRead
	ArtifactAssessmentSearch
	AuditEventRead
	AuditEventSearch
	BasicRead
	BasicSearch
	BinaryRead
	BinarySearch
	BiologicallyDerivedProductRead
	BiologicallyDerivedProductSearch
	BiologicallyDerivedProductDispenseRead
	BiologicallyDerivedProductDispenseSearch
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
	ConditionDefinitionRead
	ConditionDefinitionSearch
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
	DeviceAssociationRead
	DeviceAssociationSearch
	DeviceDefinitionRead
	DeviceDefinitionSearch
	DeviceDispenseRead
	DeviceDispenseSearch
	DeviceMetricRead
	DeviceMetricSearch
	DeviceRequestRead
	DeviceRequestSearch
	DeviceUsageRead
	DeviceUsageSearch
	DiagnosticReportRead
	DiagnosticReportSearch
	DocumentReferenceRead
	DocumentReferenceSearch
	EncounterRead
	EncounterSearch
	EncounterHistoryRead
	EncounterHistorySearch
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
	FormularyItemRead
	FormularyItemSearch
	GenomicStudyRead
	GenomicStudySearch
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
	ImagingSelectionRead
	ImagingSelectionSearch
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
	InventoryItemRead
	InventoryItemSearch
	InventoryReportRead
	InventoryReportSearch
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
	NutritionIntakeRead
	NutritionIntakeSearch
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
	PermissionRead
	PermissionSearch
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
	RequestOrchestrationRead
	RequestOrchestrationSearch
	RequirementsRead
	RequirementsSearch
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
	SubstanceNucleicAcidRead
	SubstanceNucleicAcidSearch
	SubstancePolymerRead
	SubstancePolymerSearch
	SubstanceProteinRead
	SubstanceProteinSearch
	SubstanceReferenceInformationRead
	SubstanceReferenceInformationSearch
	SubstanceSourceMaterialRead
	SubstanceSourceMaterialSearch
	SupplyDeliveryRead
	SupplyDeliverySearch
	SupplyRequestRead
	SupplyRequestSearch
	TaskRead
	TaskSearch
	TerminologyCapabilitiesRead
	TerminologyCapabilitiesSearch
	TestPlanRead
	TestPlanSearch
	TestReportRead
	TestReportSearch
	TestScriptRead
	TestScriptSearch
	TransportRead
	TransportSearch
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

func (w Concrete) ReadAccount(ctx context.Context, id string) (r5.Account, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Account", id)
	if err != nil {
		return r5.Account{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Account)
	if !ok {
		return r5.Account{}, capabilities.InvalidResourceError{ResourceType: "Account"}
	}
	return impl, nil
}
func (w Concrete) ReadActivityDefinition(ctx context.Context, id string) (r5.ActivityDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ActivityDefinition", id)
	if err != nil {
		return r5.ActivityDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ActivityDefinition)
	if !ok {
		return r5.ActivityDefinition{}, capabilities.InvalidResourceError{ResourceType: "ActivityDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadActorDefinition(ctx context.Context, id string) (r5.ActorDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ActorDefinition", id)
	if err != nil {
		return r5.ActorDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ActorDefinition)
	if !ok {
		return r5.ActorDefinition{}, capabilities.InvalidResourceError{ResourceType: "ActorDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadAdministrableProductDefinition(ctx context.Context, id string) (r5.AdministrableProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AdministrableProductDefinition", id)
	if err != nil {
		return r5.AdministrableProductDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.AdministrableProductDefinition)
	if !ok {
		return r5.AdministrableProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "AdministrableProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadAdverseEvent(ctx context.Context, id string) (r5.AdverseEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AdverseEvent", id)
	if err != nil {
		return r5.AdverseEvent{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.AdverseEvent)
	if !ok {
		return r5.AdverseEvent{}, capabilities.InvalidResourceError{ResourceType: "AdverseEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadAllergyIntolerance(ctx context.Context, id string) (r5.AllergyIntolerance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AllergyIntolerance", id)
	if err != nil {
		return r5.AllergyIntolerance{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.AllergyIntolerance)
	if !ok {
		return r5.AllergyIntolerance{}, capabilities.InvalidResourceError{ResourceType: "AllergyIntolerance"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointment(ctx context.Context, id string) (r5.Appointment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Appointment", id)
	if err != nil {
		return r5.Appointment{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Appointment)
	if !ok {
		return r5.Appointment{}, capabilities.InvalidResourceError{ResourceType: "Appointment"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointmentResponse(ctx context.Context, id string) (r5.AppointmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AppointmentResponse", id)
	if err != nil {
		return r5.AppointmentResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.AppointmentResponse)
	if !ok {
		return r5.AppointmentResponse{}, capabilities.InvalidResourceError{ResourceType: "AppointmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadArtifactAssessment(ctx context.Context, id string) (r5.ArtifactAssessment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ArtifactAssessment", id)
	if err != nil {
		return r5.ArtifactAssessment{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ArtifactAssessment)
	if !ok {
		return r5.ArtifactAssessment{}, capabilities.InvalidResourceError{ResourceType: "ArtifactAssessment"}
	}
	return impl, nil
}
func (w Concrete) ReadAuditEvent(ctx context.Context, id string) (r5.AuditEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AuditEvent", id)
	if err != nil {
		return r5.AuditEvent{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.AuditEvent)
	if !ok {
		return r5.AuditEvent{}, capabilities.InvalidResourceError{ResourceType: "AuditEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadBasic(ctx context.Context, id string) (r5.Basic, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Basic", id)
	if err != nil {
		return r5.Basic{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Basic)
	if !ok {
		return r5.Basic{}, capabilities.InvalidResourceError{ResourceType: "Basic"}
	}
	return impl, nil
}
func (w Concrete) ReadBinary(ctx context.Context, id string) (r5.Binary, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Binary", id)
	if err != nil {
		return r5.Binary{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Binary)
	if !ok {
		return r5.Binary{}, capabilities.InvalidResourceError{ResourceType: "Binary"}
	}
	return impl, nil
}
func (w Concrete) ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r5.BiologicallyDerivedProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BiologicallyDerivedProduct", id)
	if err != nil {
		return r5.BiologicallyDerivedProduct{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.BiologicallyDerivedProduct)
	if !ok {
		return r5.BiologicallyDerivedProduct{}, capabilities.InvalidResourceError{ResourceType: "BiologicallyDerivedProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadBiologicallyDerivedProductDispense(ctx context.Context, id string) (r5.BiologicallyDerivedProductDispense, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BiologicallyDerivedProductDispense", id)
	if err != nil {
		return r5.BiologicallyDerivedProductDispense{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.BiologicallyDerivedProductDispense)
	if !ok {
		return r5.BiologicallyDerivedProductDispense{}, capabilities.InvalidResourceError{ResourceType: "BiologicallyDerivedProductDispense"}
	}
	return impl, nil
}
func (w Concrete) ReadBodyStructure(ctx context.Context, id string) (r5.BodyStructure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BodyStructure", id)
	if err != nil {
		return r5.BodyStructure{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.BodyStructure)
	if !ok {
		return r5.BodyStructure{}, capabilities.InvalidResourceError{ResourceType: "BodyStructure"}
	}
	return impl, nil
}
func (w Concrete) ReadBundle(ctx context.Context, id string) (r5.Bundle, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Bundle", id)
	if err != nil {
		return r5.Bundle{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Bundle)
	if !ok {
		return r5.Bundle{}, capabilities.InvalidResourceError{ResourceType: "Bundle"}
	}
	return impl, nil
}
func (w Concrete) ReadCapabilityStatement(ctx context.Context, id string) (r5.CapabilityStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CapabilityStatement", id)
	if err != nil {
		return r5.CapabilityStatement{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CapabilityStatement)
	if !ok {
		return r5.CapabilityStatement{}, capabilities.InvalidResourceError{ResourceType: "CapabilityStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadCarePlan(ctx context.Context, id string) (r5.CarePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CarePlan", id)
	if err != nil {
		return r5.CarePlan{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CarePlan)
	if !ok {
		return r5.CarePlan{}, capabilities.InvalidResourceError{ResourceType: "CarePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadCareTeam(ctx context.Context, id string) (r5.CareTeam, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CareTeam", id)
	if err != nil {
		return r5.CareTeam{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CareTeam)
	if !ok {
		return r5.CareTeam{}, capabilities.InvalidResourceError{ResourceType: "CareTeam"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItem(ctx context.Context, id string) (r5.ChargeItem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItem", id)
	if err != nil {
		return r5.ChargeItem{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ChargeItem)
	if !ok {
		return r5.ChargeItem{}, capabilities.InvalidResourceError{ResourceType: "ChargeItem"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItemDefinition(ctx context.Context, id string) (r5.ChargeItemDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItemDefinition", id)
	if err != nil {
		return r5.ChargeItemDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ChargeItemDefinition)
	if !ok {
		return r5.ChargeItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ChargeItemDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadCitation(ctx context.Context, id string) (r5.Citation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Citation", id)
	if err != nil {
		return r5.Citation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Citation)
	if !ok {
		return r5.Citation{}, capabilities.InvalidResourceError{ResourceType: "Citation"}
	}
	return impl, nil
}
func (w Concrete) ReadClaim(ctx context.Context, id string) (r5.Claim, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Claim", id)
	if err != nil {
		return r5.Claim{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Claim)
	if !ok {
		return r5.Claim{}, capabilities.InvalidResourceError{ResourceType: "Claim"}
	}
	return impl, nil
}
func (w Concrete) ReadClaimResponse(ctx context.Context, id string) (r5.ClaimResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClaimResponse", id)
	if err != nil {
		return r5.ClaimResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ClaimResponse)
	if !ok {
		return r5.ClaimResponse{}, capabilities.InvalidResourceError{ResourceType: "ClaimResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadClinicalImpression(ctx context.Context, id string) (r5.ClinicalImpression, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClinicalImpression", id)
	if err != nil {
		return r5.ClinicalImpression{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ClinicalImpression)
	if !ok {
		return r5.ClinicalImpression{}, capabilities.InvalidResourceError{ResourceType: "ClinicalImpression"}
	}
	return impl, nil
}
func (w Concrete) ReadClinicalUseDefinition(ctx context.Context, id string) (r5.ClinicalUseDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClinicalUseDefinition", id)
	if err != nil {
		return r5.ClinicalUseDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ClinicalUseDefinition)
	if !ok {
		return r5.ClinicalUseDefinition{}, capabilities.InvalidResourceError{ResourceType: "ClinicalUseDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadCodeSystem(ctx context.Context, id string) (r5.CodeSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CodeSystem", id)
	if err != nil {
		return r5.CodeSystem{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CodeSystem)
	if !ok {
		return r5.CodeSystem{}, capabilities.InvalidResourceError{ResourceType: "CodeSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunication(ctx context.Context, id string) (r5.Communication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Communication", id)
	if err != nil {
		return r5.Communication{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Communication)
	if !ok {
		return r5.Communication{}, capabilities.InvalidResourceError{ResourceType: "Communication"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunicationRequest(ctx context.Context, id string) (r5.CommunicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CommunicationRequest", id)
	if err != nil {
		return r5.CommunicationRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CommunicationRequest)
	if !ok {
		return r5.CommunicationRequest{}, capabilities.InvalidResourceError{ResourceType: "CommunicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCompartmentDefinition(ctx context.Context, id string) (r5.CompartmentDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CompartmentDefinition", id)
	if err != nil {
		return r5.CompartmentDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CompartmentDefinition)
	if !ok {
		return r5.CompartmentDefinition{}, capabilities.InvalidResourceError{ResourceType: "CompartmentDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadComposition(ctx context.Context, id string) (r5.Composition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Composition", id)
	if err != nil {
		return r5.Composition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Composition)
	if !ok {
		return r5.Composition{}, capabilities.InvalidResourceError{ResourceType: "Composition"}
	}
	return impl, nil
}
func (w Concrete) ReadConceptMap(ctx context.Context, id string) (r5.ConceptMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ConceptMap", id)
	if err != nil {
		return r5.ConceptMap{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ConceptMap)
	if !ok {
		return r5.ConceptMap{}, capabilities.InvalidResourceError{ResourceType: "ConceptMap"}
	}
	return impl, nil
}
func (w Concrete) ReadCondition(ctx context.Context, id string) (r5.Condition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Condition", id)
	if err != nil {
		return r5.Condition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Condition)
	if !ok {
		return r5.Condition{}, capabilities.InvalidResourceError{ResourceType: "Condition"}
	}
	return impl, nil
}
func (w Concrete) ReadConditionDefinition(ctx context.Context, id string) (r5.ConditionDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ConditionDefinition", id)
	if err != nil {
		return r5.ConditionDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ConditionDefinition)
	if !ok {
		return r5.ConditionDefinition{}, capabilities.InvalidResourceError{ResourceType: "ConditionDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadConsent(ctx context.Context, id string) (r5.Consent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Consent", id)
	if err != nil {
		return r5.Consent{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Consent)
	if !ok {
		return r5.Consent{}, capabilities.InvalidResourceError{ResourceType: "Consent"}
	}
	return impl, nil
}
func (w Concrete) ReadContract(ctx context.Context, id string) (r5.Contract, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Contract", id)
	if err != nil {
		return r5.Contract{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Contract)
	if !ok {
		return r5.Contract{}, capabilities.InvalidResourceError{ResourceType: "Contract"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverage(ctx context.Context, id string) (r5.Coverage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Coverage", id)
	if err != nil {
		return r5.Coverage{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Coverage)
	if !ok {
		return r5.Coverage{}, capabilities.InvalidResourceError{ResourceType: "Coverage"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityRequest(ctx context.Context, id string) (r5.CoverageEligibilityRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityRequest", id)
	if err != nil {
		return r5.CoverageEligibilityRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CoverageEligibilityRequest)
	if !ok {
		return r5.CoverageEligibilityRequest{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityResponse(ctx context.Context, id string) (r5.CoverageEligibilityResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityResponse", id)
	if err != nil {
		return r5.CoverageEligibilityResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.CoverageEligibilityResponse)
	if !ok {
		return r5.CoverageEligibilityResponse{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadDetectedIssue(ctx context.Context, id string) (r5.DetectedIssue, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DetectedIssue", id)
	if err != nil {
		return r5.DetectedIssue{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DetectedIssue)
	if !ok {
		return r5.DetectedIssue{}, capabilities.InvalidResourceError{ResourceType: "DetectedIssue"}
	}
	return impl, nil
}
func (w Concrete) ReadDevice(ctx context.Context, id string) (r5.Device, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Device", id)
	if err != nil {
		return r5.Device{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Device)
	if !ok {
		return r5.Device{}, capabilities.InvalidResourceError{ResourceType: "Device"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceAssociation(ctx context.Context, id string) (r5.DeviceAssociation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceAssociation", id)
	if err != nil {
		return r5.DeviceAssociation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceAssociation)
	if !ok {
		return r5.DeviceAssociation{}, capabilities.InvalidResourceError{ResourceType: "DeviceAssociation"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceDefinition(ctx context.Context, id string) (r5.DeviceDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceDefinition", id)
	if err != nil {
		return r5.DeviceDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceDefinition)
	if !ok {
		return r5.DeviceDefinition{}, capabilities.InvalidResourceError{ResourceType: "DeviceDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceDispense(ctx context.Context, id string) (r5.DeviceDispense, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceDispense", id)
	if err != nil {
		return r5.DeviceDispense{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceDispense)
	if !ok {
		return r5.DeviceDispense{}, capabilities.InvalidResourceError{ResourceType: "DeviceDispense"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceMetric(ctx context.Context, id string) (r5.DeviceMetric, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceMetric", id)
	if err != nil {
		return r5.DeviceMetric{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceMetric)
	if !ok {
		return r5.DeviceMetric{}, capabilities.InvalidResourceError{ResourceType: "DeviceMetric"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceRequest(ctx context.Context, id string) (r5.DeviceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceRequest", id)
	if err != nil {
		return r5.DeviceRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceRequest)
	if !ok {
		return r5.DeviceRequest{}, capabilities.InvalidResourceError{ResourceType: "DeviceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceUsage(ctx context.Context, id string) (r5.DeviceUsage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceUsage", id)
	if err != nil {
		return r5.DeviceUsage{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DeviceUsage)
	if !ok {
		return r5.DeviceUsage{}, capabilities.InvalidResourceError{ResourceType: "DeviceUsage"}
	}
	return impl, nil
}
func (w Concrete) ReadDiagnosticReport(ctx context.Context, id string) (r5.DiagnosticReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DiagnosticReport", id)
	if err != nil {
		return r5.DiagnosticReport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DiagnosticReport)
	if !ok {
		return r5.DiagnosticReport{}, capabilities.InvalidResourceError{ResourceType: "DiagnosticReport"}
	}
	return impl, nil
}
func (w Concrete) ReadDocumentReference(ctx context.Context, id string) (r5.DocumentReference, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DocumentReference", id)
	if err != nil {
		return r5.DocumentReference{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.DocumentReference)
	if !ok {
		return r5.DocumentReference{}, capabilities.InvalidResourceError{ResourceType: "DocumentReference"}
	}
	return impl, nil
}
func (w Concrete) ReadEncounter(ctx context.Context, id string) (r5.Encounter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Encounter", id)
	if err != nil {
		return r5.Encounter{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Encounter)
	if !ok {
		return r5.Encounter{}, capabilities.InvalidResourceError{ResourceType: "Encounter"}
	}
	return impl, nil
}
func (w Concrete) ReadEncounterHistory(ctx context.Context, id string) (r5.EncounterHistory, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EncounterHistory", id)
	if err != nil {
		return r5.EncounterHistory{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EncounterHistory)
	if !ok {
		return r5.EncounterHistory{}, capabilities.InvalidResourceError{ResourceType: "EncounterHistory"}
	}
	return impl, nil
}
func (w Concrete) ReadEndpoint(ctx context.Context, id string) (r5.Endpoint, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Endpoint", id)
	if err != nil {
		return r5.Endpoint{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Endpoint)
	if !ok {
		return r5.Endpoint{}, capabilities.InvalidResourceError{ResourceType: "Endpoint"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentRequest(ctx context.Context, id string) (r5.EnrollmentRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentRequest", id)
	if err != nil {
		return r5.EnrollmentRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EnrollmentRequest)
	if !ok {
		return r5.EnrollmentRequest{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentResponse(ctx context.Context, id string) (r5.EnrollmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentResponse", id)
	if err != nil {
		return r5.EnrollmentResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EnrollmentResponse)
	if !ok {
		return r5.EnrollmentResponse{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadEpisodeOfCare(ctx context.Context, id string) (r5.EpisodeOfCare, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EpisodeOfCare", id)
	if err != nil {
		return r5.EpisodeOfCare{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EpisodeOfCare)
	if !ok {
		return r5.EpisodeOfCare{}, capabilities.InvalidResourceError{ResourceType: "EpisodeOfCare"}
	}
	return impl, nil
}
func (w Concrete) ReadEventDefinition(ctx context.Context, id string) (r5.EventDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EventDefinition", id)
	if err != nil {
		return r5.EventDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EventDefinition)
	if !ok {
		return r5.EventDefinition{}, capabilities.InvalidResourceError{ResourceType: "EventDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidence(ctx context.Context, id string) (r5.Evidence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Evidence", id)
	if err != nil {
		return r5.Evidence{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Evidence)
	if !ok {
		return r5.Evidence{}, capabilities.InvalidResourceError{ResourceType: "Evidence"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidenceReport(ctx context.Context, id string) (r5.EvidenceReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EvidenceReport", id)
	if err != nil {
		return r5.EvidenceReport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EvidenceReport)
	if !ok {
		return r5.EvidenceReport{}, capabilities.InvalidResourceError{ResourceType: "EvidenceReport"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidenceVariable(ctx context.Context, id string) (r5.EvidenceVariable, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EvidenceVariable", id)
	if err != nil {
		return r5.EvidenceVariable{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.EvidenceVariable)
	if !ok {
		return r5.EvidenceVariable{}, capabilities.InvalidResourceError{ResourceType: "EvidenceVariable"}
	}
	return impl, nil
}
func (w Concrete) ReadExampleScenario(ctx context.Context, id string) (r5.ExampleScenario, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExampleScenario", id)
	if err != nil {
		return r5.ExampleScenario{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ExampleScenario)
	if !ok {
		return r5.ExampleScenario{}, capabilities.InvalidResourceError{ResourceType: "ExampleScenario"}
	}
	return impl, nil
}
func (w Concrete) ReadExplanationOfBenefit(ctx context.Context, id string) (r5.ExplanationOfBenefit, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExplanationOfBenefit", id)
	if err != nil {
		return r5.ExplanationOfBenefit{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ExplanationOfBenefit)
	if !ok {
		return r5.ExplanationOfBenefit{}, capabilities.InvalidResourceError{ResourceType: "ExplanationOfBenefit"}
	}
	return impl, nil
}
func (w Concrete) ReadFamilyMemberHistory(ctx context.Context, id string) (r5.FamilyMemberHistory, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "FamilyMemberHistory", id)
	if err != nil {
		return r5.FamilyMemberHistory{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.FamilyMemberHistory)
	if !ok {
		return r5.FamilyMemberHistory{}, capabilities.InvalidResourceError{ResourceType: "FamilyMemberHistory"}
	}
	return impl, nil
}
func (w Concrete) ReadFlag(ctx context.Context, id string) (r5.Flag, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Flag", id)
	if err != nil {
		return r5.Flag{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Flag)
	if !ok {
		return r5.Flag{}, capabilities.InvalidResourceError{ResourceType: "Flag"}
	}
	return impl, nil
}
func (w Concrete) ReadFormularyItem(ctx context.Context, id string) (r5.FormularyItem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "FormularyItem", id)
	if err != nil {
		return r5.FormularyItem{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.FormularyItem)
	if !ok {
		return r5.FormularyItem{}, capabilities.InvalidResourceError{ResourceType: "FormularyItem"}
	}
	return impl, nil
}
func (w Concrete) ReadGenomicStudy(ctx context.Context, id string) (r5.GenomicStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GenomicStudy", id)
	if err != nil {
		return r5.GenomicStudy{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.GenomicStudy)
	if !ok {
		return r5.GenomicStudy{}, capabilities.InvalidResourceError{ResourceType: "GenomicStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadGoal(ctx context.Context, id string) (r5.Goal, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Goal", id)
	if err != nil {
		return r5.Goal{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Goal)
	if !ok {
		return r5.Goal{}, capabilities.InvalidResourceError{ResourceType: "Goal"}
	}
	return impl, nil
}
func (w Concrete) ReadGraphDefinition(ctx context.Context, id string) (r5.GraphDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GraphDefinition", id)
	if err != nil {
		return r5.GraphDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.GraphDefinition)
	if !ok {
		return r5.GraphDefinition{}, capabilities.InvalidResourceError{ResourceType: "GraphDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadGroup(ctx context.Context, id string) (r5.Group, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Group", id)
	if err != nil {
		return r5.Group{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Group)
	if !ok {
		return r5.Group{}, capabilities.InvalidResourceError{ResourceType: "Group"}
	}
	return impl, nil
}
func (w Concrete) ReadGuidanceResponse(ctx context.Context, id string) (r5.GuidanceResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GuidanceResponse", id)
	if err != nil {
		return r5.GuidanceResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.GuidanceResponse)
	if !ok {
		return r5.GuidanceResponse{}, capabilities.InvalidResourceError{ResourceType: "GuidanceResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadHealthcareService(ctx context.Context, id string) (r5.HealthcareService, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "HealthcareService", id)
	if err != nil {
		return r5.HealthcareService{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.HealthcareService)
	if !ok {
		return r5.HealthcareService{}, capabilities.InvalidResourceError{ResourceType: "HealthcareService"}
	}
	return impl, nil
}
func (w Concrete) ReadImagingSelection(ctx context.Context, id string) (r5.ImagingSelection, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImagingSelection", id)
	if err != nil {
		return r5.ImagingSelection{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ImagingSelection)
	if !ok {
		return r5.ImagingSelection{}, capabilities.InvalidResourceError{ResourceType: "ImagingSelection"}
	}
	return impl, nil
}
func (w Concrete) ReadImagingStudy(ctx context.Context, id string) (r5.ImagingStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImagingStudy", id)
	if err != nil {
		return r5.ImagingStudy{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ImagingStudy)
	if !ok {
		return r5.ImagingStudy{}, capabilities.InvalidResourceError{ResourceType: "ImagingStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunization(ctx context.Context, id string) (r5.Immunization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Immunization", id)
	if err != nil {
		return r5.Immunization{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Immunization)
	if !ok {
		return r5.Immunization{}, capabilities.InvalidResourceError{ResourceType: "Immunization"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationEvaluation(ctx context.Context, id string) (r5.ImmunizationEvaluation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationEvaluation", id)
	if err != nil {
		return r5.ImmunizationEvaluation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ImmunizationEvaluation)
	if !ok {
		return r5.ImmunizationEvaluation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationEvaluation"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationRecommendation(ctx context.Context, id string) (r5.ImmunizationRecommendation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationRecommendation", id)
	if err != nil {
		return r5.ImmunizationRecommendation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ImmunizationRecommendation)
	if !ok {
		return r5.ImmunizationRecommendation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationRecommendation"}
	}
	return impl, nil
}
func (w Concrete) ReadImplementationGuide(ctx context.Context, id string) (r5.ImplementationGuide, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImplementationGuide", id)
	if err != nil {
		return r5.ImplementationGuide{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ImplementationGuide)
	if !ok {
		return r5.ImplementationGuide{}, capabilities.InvalidResourceError{ResourceType: "ImplementationGuide"}
	}
	return impl, nil
}
func (w Concrete) ReadIngredient(ctx context.Context, id string) (r5.Ingredient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Ingredient", id)
	if err != nil {
		return r5.Ingredient{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Ingredient)
	if !ok {
		return r5.Ingredient{}, capabilities.InvalidResourceError{ResourceType: "Ingredient"}
	}
	return impl, nil
}
func (w Concrete) ReadInsurancePlan(ctx context.Context, id string) (r5.InsurancePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "InsurancePlan", id)
	if err != nil {
		return r5.InsurancePlan{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.InsurancePlan)
	if !ok {
		return r5.InsurancePlan{}, capabilities.InvalidResourceError{ResourceType: "InsurancePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadInventoryItem(ctx context.Context, id string) (r5.InventoryItem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "InventoryItem", id)
	if err != nil {
		return r5.InventoryItem{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.InventoryItem)
	if !ok {
		return r5.InventoryItem{}, capabilities.InvalidResourceError{ResourceType: "InventoryItem"}
	}
	return impl, nil
}
func (w Concrete) ReadInventoryReport(ctx context.Context, id string) (r5.InventoryReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "InventoryReport", id)
	if err != nil {
		return r5.InventoryReport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.InventoryReport)
	if !ok {
		return r5.InventoryReport{}, capabilities.InvalidResourceError{ResourceType: "InventoryReport"}
	}
	return impl, nil
}
func (w Concrete) ReadInvoice(ctx context.Context, id string) (r5.Invoice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Invoice", id)
	if err != nil {
		return r5.Invoice{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Invoice)
	if !ok {
		return r5.Invoice{}, capabilities.InvalidResourceError{ResourceType: "Invoice"}
	}
	return impl, nil
}
func (w Concrete) ReadLibrary(ctx context.Context, id string) (r5.Library, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Library", id)
	if err != nil {
		return r5.Library{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Library)
	if !ok {
		return r5.Library{}, capabilities.InvalidResourceError{ResourceType: "Library"}
	}
	return impl, nil
}
func (w Concrete) ReadLinkage(ctx context.Context, id string) (r5.Linkage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Linkage", id)
	if err != nil {
		return r5.Linkage{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Linkage)
	if !ok {
		return r5.Linkage{}, capabilities.InvalidResourceError{ResourceType: "Linkage"}
	}
	return impl, nil
}
func (w Concrete) ReadList(ctx context.Context, id string) (r5.List, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "List", id)
	if err != nil {
		return r5.List{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.List)
	if !ok {
		return r5.List{}, capabilities.InvalidResourceError{ResourceType: "List"}
	}
	return impl, nil
}
func (w Concrete) ReadLocation(ctx context.Context, id string) (r5.Location, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Location", id)
	if err != nil {
		return r5.Location{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Location)
	if !ok {
		return r5.Location{}, capabilities.InvalidResourceError{ResourceType: "Location"}
	}
	return impl, nil
}
func (w Concrete) ReadManufacturedItemDefinition(ctx context.Context, id string) (r5.ManufacturedItemDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ManufacturedItemDefinition", id)
	if err != nil {
		return r5.ManufacturedItemDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ManufacturedItemDefinition)
	if !ok {
		return r5.ManufacturedItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ManufacturedItemDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasure(ctx context.Context, id string) (r5.Measure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Measure", id)
	if err != nil {
		return r5.Measure{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Measure)
	if !ok {
		return r5.Measure{}, capabilities.InvalidResourceError{ResourceType: "Measure"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasureReport(ctx context.Context, id string) (r5.MeasureReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MeasureReport", id)
	if err != nil {
		return r5.MeasureReport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MeasureReport)
	if !ok {
		return r5.MeasureReport{}, capabilities.InvalidResourceError{ResourceType: "MeasureReport"}
	}
	return impl, nil
}
func (w Concrete) ReadMedication(ctx context.Context, id string) (r5.Medication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Medication", id)
	if err != nil {
		return r5.Medication{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Medication)
	if !ok {
		return r5.Medication{}, capabilities.InvalidResourceError{ResourceType: "Medication"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationAdministration(ctx context.Context, id string) (r5.MedicationAdministration, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationAdministration", id)
	if err != nil {
		return r5.MedicationAdministration{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicationAdministration)
	if !ok {
		return r5.MedicationAdministration{}, capabilities.InvalidResourceError{ResourceType: "MedicationAdministration"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationDispense(ctx context.Context, id string) (r5.MedicationDispense, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationDispense", id)
	if err != nil {
		return r5.MedicationDispense{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicationDispense)
	if !ok {
		return r5.MedicationDispense{}, capabilities.InvalidResourceError{ResourceType: "MedicationDispense"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationKnowledge(ctx context.Context, id string) (r5.MedicationKnowledge, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationKnowledge", id)
	if err != nil {
		return r5.MedicationKnowledge{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicationKnowledge)
	if !ok {
		return r5.MedicationKnowledge{}, capabilities.InvalidResourceError{ResourceType: "MedicationKnowledge"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationRequest(ctx context.Context, id string) (r5.MedicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationRequest", id)
	if err != nil {
		return r5.MedicationRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicationRequest)
	if !ok {
		return r5.MedicationRequest{}, capabilities.InvalidResourceError{ResourceType: "MedicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationStatement(ctx context.Context, id string) (r5.MedicationStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationStatement", id)
	if err != nil {
		return r5.MedicationStatement{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicationStatement)
	if !ok {
		return r5.MedicationStatement{}, capabilities.InvalidResourceError{ResourceType: "MedicationStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductDefinition(ctx context.Context, id string) (r5.MedicinalProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductDefinition", id)
	if err != nil {
		return r5.MedicinalProductDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MedicinalProductDefinition)
	if !ok {
		return r5.MedicinalProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageDefinition(ctx context.Context, id string) (r5.MessageDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageDefinition", id)
	if err != nil {
		return r5.MessageDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MessageDefinition)
	if !ok {
		return r5.MessageDefinition{}, capabilities.InvalidResourceError{ResourceType: "MessageDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageHeader(ctx context.Context, id string) (r5.MessageHeader, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageHeader", id)
	if err != nil {
		return r5.MessageHeader{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MessageHeader)
	if !ok {
		return r5.MessageHeader{}, capabilities.InvalidResourceError{ResourceType: "MessageHeader"}
	}
	return impl, nil
}
func (w Concrete) ReadMolecularSequence(ctx context.Context, id string) (r5.MolecularSequence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MolecularSequence", id)
	if err != nil {
		return r5.MolecularSequence{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.MolecularSequence)
	if !ok {
		return r5.MolecularSequence{}, capabilities.InvalidResourceError{ResourceType: "MolecularSequence"}
	}
	return impl, nil
}
func (w Concrete) ReadNamingSystem(ctx context.Context, id string) (r5.NamingSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NamingSystem", id)
	if err != nil {
		return r5.NamingSystem{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.NamingSystem)
	if !ok {
		return r5.NamingSystem{}, capabilities.InvalidResourceError{ResourceType: "NamingSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionIntake(ctx context.Context, id string) (r5.NutritionIntake, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionIntake", id)
	if err != nil {
		return r5.NutritionIntake{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.NutritionIntake)
	if !ok {
		return r5.NutritionIntake{}, capabilities.InvalidResourceError{ResourceType: "NutritionIntake"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionOrder(ctx context.Context, id string) (r5.NutritionOrder, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionOrder", id)
	if err != nil {
		return r5.NutritionOrder{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.NutritionOrder)
	if !ok {
		return r5.NutritionOrder{}, capabilities.InvalidResourceError{ResourceType: "NutritionOrder"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionProduct(ctx context.Context, id string) (r5.NutritionProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionProduct", id)
	if err != nil {
		return r5.NutritionProduct{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.NutritionProduct)
	if !ok {
		return r5.NutritionProduct{}, capabilities.InvalidResourceError{ResourceType: "NutritionProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadObservation(ctx context.Context, id string) (r5.Observation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Observation", id)
	if err != nil {
		return r5.Observation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Observation)
	if !ok {
		return r5.Observation{}, capabilities.InvalidResourceError{ResourceType: "Observation"}
	}
	return impl, nil
}
func (w Concrete) ReadObservationDefinition(ctx context.Context, id string) (r5.ObservationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ObservationDefinition", id)
	if err != nil {
		return r5.ObservationDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ObservationDefinition)
	if !ok {
		return r5.ObservationDefinition{}, capabilities.InvalidResourceError{ResourceType: "ObservationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationDefinition(ctx context.Context, id string) (r5.OperationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationDefinition", id)
	if err != nil {
		return r5.OperationDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.OperationDefinition)
	if !ok {
		return r5.OperationDefinition{}, capabilities.InvalidResourceError{ResourceType: "OperationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationOutcome(ctx context.Context, id string) (r5.OperationOutcome, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationOutcome", id)
	if err != nil {
		return r5.OperationOutcome{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.OperationOutcome)
	if !ok {
		return r5.OperationOutcome{}, capabilities.InvalidResourceError{ResourceType: "OperationOutcome"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganization(ctx context.Context, id string) (r5.Organization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Organization", id)
	if err != nil {
		return r5.Organization{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Organization)
	if !ok {
		return r5.Organization{}, capabilities.InvalidResourceError{ResourceType: "Organization"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganizationAffiliation(ctx context.Context, id string) (r5.OrganizationAffiliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OrganizationAffiliation", id)
	if err != nil {
		return r5.OrganizationAffiliation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.OrganizationAffiliation)
	if !ok {
		return r5.OrganizationAffiliation{}, capabilities.InvalidResourceError{ResourceType: "OrganizationAffiliation"}
	}
	return impl, nil
}
func (w Concrete) ReadPackagedProductDefinition(ctx context.Context, id string) (r5.PackagedProductDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PackagedProductDefinition", id)
	if err != nil {
		return r5.PackagedProductDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.PackagedProductDefinition)
	if !ok {
		return r5.PackagedProductDefinition{}, capabilities.InvalidResourceError{ResourceType: "PackagedProductDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadParameters(ctx context.Context, id string) (r5.Parameters, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Parameters", id)
	if err != nil {
		return r5.Parameters{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Parameters)
	if !ok {
		return r5.Parameters{}, capabilities.InvalidResourceError{ResourceType: "Parameters"}
	}
	return impl, nil
}
func (w Concrete) ReadPatient(ctx context.Context, id string) (r5.Patient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Patient", id)
	if err != nil {
		return r5.Patient{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Patient)
	if !ok {
		return r5.Patient{}, capabilities.InvalidResourceError{ResourceType: "Patient"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentNotice(ctx context.Context, id string) (r5.PaymentNotice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentNotice", id)
	if err != nil {
		return r5.PaymentNotice{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.PaymentNotice)
	if !ok {
		return r5.PaymentNotice{}, capabilities.InvalidResourceError{ResourceType: "PaymentNotice"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentReconciliation(ctx context.Context, id string) (r5.PaymentReconciliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentReconciliation", id)
	if err != nil {
		return r5.PaymentReconciliation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.PaymentReconciliation)
	if !ok {
		return r5.PaymentReconciliation{}, capabilities.InvalidResourceError{ResourceType: "PaymentReconciliation"}
	}
	return impl, nil
}
func (w Concrete) ReadPermission(ctx context.Context, id string) (r5.Permission, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Permission", id)
	if err != nil {
		return r5.Permission{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Permission)
	if !ok {
		return r5.Permission{}, capabilities.InvalidResourceError{ResourceType: "Permission"}
	}
	return impl, nil
}
func (w Concrete) ReadPerson(ctx context.Context, id string) (r5.Person, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Person", id)
	if err != nil {
		return r5.Person{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Person)
	if !ok {
		return r5.Person{}, capabilities.InvalidResourceError{ResourceType: "Person"}
	}
	return impl, nil
}
func (w Concrete) ReadPlanDefinition(ctx context.Context, id string) (r5.PlanDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PlanDefinition", id)
	if err != nil {
		return r5.PlanDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.PlanDefinition)
	if !ok {
		return r5.PlanDefinition{}, capabilities.InvalidResourceError{ResourceType: "PlanDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitioner(ctx context.Context, id string) (r5.Practitioner, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Practitioner", id)
	if err != nil {
		return r5.Practitioner{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Practitioner)
	if !ok {
		return r5.Practitioner{}, capabilities.InvalidResourceError{ResourceType: "Practitioner"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitionerRole(ctx context.Context, id string) (r5.PractitionerRole, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PractitionerRole", id)
	if err != nil {
		return r5.PractitionerRole{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.PractitionerRole)
	if !ok {
		return r5.PractitionerRole{}, capabilities.InvalidResourceError{ResourceType: "PractitionerRole"}
	}
	return impl, nil
}
func (w Concrete) ReadProcedure(ctx context.Context, id string) (r5.Procedure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Procedure", id)
	if err != nil {
		return r5.Procedure{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Procedure)
	if !ok {
		return r5.Procedure{}, capabilities.InvalidResourceError{ResourceType: "Procedure"}
	}
	return impl, nil
}
func (w Concrete) ReadProvenance(ctx context.Context, id string) (r5.Provenance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Provenance", id)
	if err != nil {
		return r5.Provenance{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Provenance)
	if !ok {
		return r5.Provenance{}, capabilities.InvalidResourceError{ResourceType: "Provenance"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaire(ctx context.Context, id string) (r5.Questionnaire, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Questionnaire", id)
	if err != nil {
		return r5.Questionnaire{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Questionnaire)
	if !ok {
		return r5.Questionnaire{}, capabilities.InvalidResourceError{ResourceType: "Questionnaire"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaireResponse(ctx context.Context, id string) (r5.QuestionnaireResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "QuestionnaireResponse", id)
	if err != nil {
		return r5.QuestionnaireResponse{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.QuestionnaireResponse)
	if !ok {
		return r5.QuestionnaireResponse{}, capabilities.InvalidResourceError{ResourceType: "QuestionnaireResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadRegulatedAuthorization(ctx context.Context, id string) (r5.RegulatedAuthorization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RegulatedAuthorization", id)
	if err != nil {
		return r5.RegulatedAuthorization{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.RegulatedAuthorization)
	if !ok {
		return r5.RegulatedAuthorization{}, capabilities.InvalidResourceError{ResourceType: "RegulatedAuthorization"}
	}
	return impl, nil
}
func (w Concrete) ReadRelatedPerson(ctx context.Context, id string) (r5.RelatedPerson, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RelatedPerson", id)
	if err != nil {
		return r5.RelatedPerson{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.RelatedPerson)
	if !ok {
		return r5.RelatedPerson{}, capabilities.InvalidResourceError{ResourceType: "RelatedPerson"}
	}
	return impl, nil
}
func (w Concrete) ReadRequestOrchestration(ctx context.Context, id string) (r5.RequestOrchestration, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RequestOrchestration", id)
	if err != nil {
		return r5.RequestOrchestration{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.RequestOrchestration)
	if !ok {
		return r5.RequestOrchestration{}, capabilities.InvalidResourceError{ResourceType: "RequestOrchestration"}
	}
	return impl, nil
}
func (w Concrete) ReadRequirements(ctx context.Context, id string) (r5.Requirements, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Requirements", id)
	if err != nil {
		return r5.Requirements{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Requirements)
	if !ok {
		return r5.Requirements{}, capabilities.InvalidResourceError{ResourceType: "Requirements"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchStudy(ctx context.Context, id string) (r5.ResearchStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchStudy", id)
	if err != nil {
		return r5.ResearchStudy{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ResearchStudy)
	if !ok {
		return r5.ResearchStudy{}, capabilities.InvalidResourceError{ResourceType: "ResearchStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchSubject(ctx context.Context, id string) (r5.ResearchSubject, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchSubject", id)
	if err != nil {
		return r5.ResearchSubject{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ResearchSubject)
	if !ok {
		return r5.ResearchSubject{}, capabilities.InvalidResourceError{ResourceType: "ResearchSubject"}
	}
	return impl, nil
}
func (w Concrete) ReadRiskAssessment(ctx context.Context, id string) (r5.RiskAssessment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RiskAssessment", id)
	if err != nil {
		return r5.RiskAssessment{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.RiskAssessment)
	if !ok {
		return r5.RiskAssessment{}, capabilities.InvalidResourceError{ResourceType: "RiskAssessment"}
	}
	return impl, nil
}
func (w Concrete) ReadSchedule(ctx context.Context, id string) (r5.Schedule, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Schedule", id)
	if err != nil {
		return r5.Schedule{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Schedule)
	if !ok {
		return r5.Schedule{}, capabilities.InvalidResourceError{ResourceType: "Schedule"}
	}
	return impl, nil
}
func (w Concrete) ReadSearchParameter(ctx context.Context, id string) (r5.SearchParameter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SearchParameter", id)
	if err != nil {
		return r5.SearchParameter{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SearchParameter)
	if !ok {
		return r5.SearchParameter{}, capabilities.InvalidResourceError{ResourceType: "SearchParameter"}
	}
	return impl, nil
}
func (w Concrete) ReadServiceRequest(ctx context.Context, id string) (r5.ServiceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ServiceRequest", id)
	if err != nil {
		return r5.ServiceRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ServiceRequest)
	if !ok {
		return r5.ServiceRequest{}, capabilities.InvalidResourceError{ResourceType: "ServiceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadSlot(ctx context.Context, id string) (r5.Slot, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Slot", id)
	if err != nil {
		return r5.Slot{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Slot)
	if !ok {
		return r5.Slot{}, capabilities.InvalidResourceError{ResourceType: "Slot"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimen(ctx context.Context, id string) (r5.Specimen, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Specimen", id)
	if err != nil {
		return r5.Specimen{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Specimen)
	if !ok {
		return r5.Specimen{}, capabilities.InvalidResourceError{ResourceType: "Specimen"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimenDefinition(ctx context.Context, id string) (r5.SpecimenDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SpecimenDefinition", id)
	if err != nil {
		return r5.SpecimenDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SpecimenDefinition)
	if !ok {
		return r5.SpecimenDefinition{}, capabilities.InvalidResourceError{ResourceType: "SpecimenDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureDefinition(ctx context.Context, id string) (r5.StructureDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureDefinition", id)
	if err != nil {
		return r5.StructureDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.StructureDefinition)
	if !ok {
		return r5.StructureDefinition{}, capabilities.InvalidResourceError{ResourceType: "StructureDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureMap(ctx context.Context, id string) (r5.StructureMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureMap", id)
	if err != nil {
		return r5.StructureMap{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.StructureMap)
	if !ok {
		return r5.StructureMap{}, capabilities.InvalidResourceError{ResourceType: "StructureMap"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscription(ctx context.Context, id string) (r5.Subscription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Subscription", id)
	if err != nil {
		return r5.Subscription{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Subscription)
	if !ok {
		return r5.Subscription{}, capabilities.InvalidResourceError{ResourceType: "Subscription"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscriptionStatus(ctx context.Context, id string) (r5.SubscriptionStatus, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubscriptionStatus", id)
	if err != nil {
		return r5.SubscriptionStatus{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubscriptionStatus)
	if !ok {
		return r5.SubscriptionStatus{}, capabilities.InvalidResourceError{ResourceType: "SubscriptionStatus"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscriptionTopic(ctx context.Context, id string) (r5.SubscriptionTopic, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubscriptionTopic", id)
	if err != nil {
		return r5.SubscriptionTopic{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubscriptionTopic)
	if !ok {
		return r5.SubscriptionTopic{}, capabilities.InvalidResourceError{ResourceType: "SubscriptionTopic"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstance(ctx context.Context, id string) (r5.Substance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Substance", id)
	if err != nil {
		return r5.Substance{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Substance)
	if !ok {
		return r5.Substance{}, capabilities.InvalidResourceError{ResourceType: "Substance"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceDefinition(ctx context.Context, id string) (r5.SubstanceDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceDefinition", id)
	if err != nil {
		return r5.SubstanceDefinition{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstanceDefinition)
	if !ok {
		return r5.SubstanceDefinition{}, capabilities.InvalidResourceError{ResourceType: "SubstanceDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceNucleicAcid(ctx context.Context, id string) (r5.SubstanceNucleicAcid, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceNucleicAcid", id)
	if err != nil {
		return r5.SubstanceNucleicAcid{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstanceNucleicAcid)
	if !ok {
		return r5.SubstanceNucleicAcid{}, capabilities.InvalidResourceError{ResourceType: "SubstanceNucleicAcid"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstancePolymer(ctx context.Context, id string) (r5.SubstancePolymer, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstancePolymer", id)
	if err != nil {
		return r5.SubstancePolymer{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstancePolymer)
	if !ok {
		return r5.SubstancePolymer{}, capabilities.InvalidResourceError{ResourceType: "SubstancePolymer"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceProtein(ctx context.Context, id string) (r5.SubstanceProtein, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceProtein", id)
	if err != nil {
		return r5.SubstanceProtein{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstanceProtein)
	if !ok {
		return r5.SubstanceProtein{}, capabilities.InvalidResourceError{ResourceType: "SubstanceProtein"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceReferenceInformation(ctx context.Context, id string) (r5.SubstanceReferenceInformation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceReferenceInformation", id)
	if err != nil {
		return r5.SubstanceReferenceInformation{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstanceReferenceInformation)
	if !ok {
		return r5.SubstanceReferenceInformation{}, capabilities.InvalidResourceError{ResourceType: "SubstanceReferenceInformation"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceSourceMaterial(ctx context.Context, id string) (r5.SubstanceSourceMaterial, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceSourceMaterial", id)
	if err != nil {
		return r5.SubstanceSourceMaterial{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SubstanceSourceMaterial)
	if !ok {
		return r5.SubstanceSourceMaterial{}, capabilities.InvalidResourceError{ResourceType: "SubstanceSourceMaterial"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyDelivery(ctx context.Context, id string) (r5.SupplyDelivery, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyDelivery", id)
	if err != nil {
		return r5.SupplyDelivery{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SupplyDelivery)
	if !ok {
		return r5.SupplyDelivery{}, capabilities.InvalidResourceError{ResourceType: "SupplyDelivery"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyRequest(ctx context.Context, id string) (r5.SupplyRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyRequest", id)
	if err != nil {
		return r5.SupplyRequest{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.SupplyRequest)
	if !ok {
		return r5.SupplyRequest{}, capabilities.InvalidResourceError{ResourceType: "SupplyRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadTask(ctx context.Context, id string) (r5.Task, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Task", id)
	if err != nil {
		return r5.Task{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Task)
	if !ok {
		return r5.Task{}, capabilities.InvalidResourceError{ResourceType: "Task"}
	}
	return impl, nil
}
func (w Concrete) ReadTerminologyCapabilities(ctx context.Context, id string) (r5.TerminologyCapabilities, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TerminologyCapabilities", id)
	if err != nil {
		return r5.TerminologyCapabilities{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.TerminologyCapabilities)
	if !ok {
		return r5.TerminologyCapabilities{}, capabilities.InvalidResourceError{ResourceType: "TerminologyCapabilities"}
	}
	return impl, nil
}
func (w Concrete) ReadTestPlan(ctx context.Context, id string) (r5.TestPlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestPlan", id)
	if err != nil {
		return r5.TestPlan{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.TestPlan)
	if !ok {
		return r5.TestPlan{}, capabilities.InvalidResourceError{ResourceType: "TestPlan"}
	}
	return impl, nil
}
func (w Concrete) ReadTestReport(ctx context.Context, id string) (r5.TestReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestReport", id)
	if err != nil {
		return r5.TestReport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.TestReport)
	if !ok {
		return r5.TestReport{}, capabilities.InvalidResourceError{ResourceType: "TestReport"}
	}
	return impl, nil
}
func (w Concrete) ReadTestScript(ctx context.Context, id string) (r5.TestScript, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestScript", id)
	if err != nil {
		return r5.TestScript{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.TestScript)
	if !ok {
		return r5.TestScript{}, capabilities.InvalidResourceError{ResourceType: "TestScript"}
	}
	return impl, nil
}
func (w Concrete) ReadTransport(ctx context.Context, id string) (r5.Transport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Transport", id)
	if err != nil {
		return r5.Transport{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.Transport)
	if !ok {
		return r5.Transport{}, capabilities.InvalidResourceError{ResourceType: "Transport"}
	}
	return impl, nil
}
func (w Concrete) ReadValueSet(ctx context.Context, id string) (r5.ValueSet, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ValueSet", id)
	if err != nil {
		return r5.ValueSet{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.ValueSet)
	if !ok {
		return r5.ValueSet{}, capabilities.InvalidResourceError{ResourceType: "ValueSet"}
	}
	return impl, nil
}
func (w Concrete) ReadVerificationResult(ctx context.Context, id string) (r5.VerificationResult, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VerificationResult", id)
	if err != nil {
		return r5.VerificationResult{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.VerificationResult)
	if !ok {
		return r5.VerificationResult{}, capabilities.InvalidResourceError{ResourceType: "VerificationResult"}
	}
	return impl, nil
}
func (w Concrete) ReadVisionPrescription(ctx context.Context, id string) (r5.VisionPrescription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VisionPrescription", id)
	if err != nil {
		return r5.VisionPrescription{}, err
	}
	contained, ok := v.(r5.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r5.VisionPrescription)
	if !ok {
		return r5.VisionPrescription{}, capabilities.InvalidResourceError{ResourceType: "VisionPrescription"}
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
func (w Concrete) SearchCapabilitiesActorDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ActorDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchActorDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ActorDefinition", options)
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
func (w Concrete) SearchCapabilitiesArtifactAssessment() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ArtifactAssessment")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchArtifactAssessment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ArtifactAssessment", options)
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
func (w Concrete) SearchCapabilitiesBiologicallyDerivedProductDispense() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("BiologicallyDerivedProductDispense")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchBiologicallyDerivedProductDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "BiologicallyDerivedProductDispense", options)
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
func (w Concrete) SearchCapabilitiesConditionDefinition() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ConditionDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchConditionDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ConditionDefinition", options)
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
func (w Concrete) SearchCapabilitiesDeviceAssociation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceAssociation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceAssociation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceAssociation", options)
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
func (w Concrete) SearchCapabilitiesDeviceDispense() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceDispense")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceDispense", options)
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
func (w Concrete) SearchCapabilitiesDeviceUsage() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("DeviceUsage")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchDeviceUsage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "DeviceUsage", options)
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
func (w Concrete) SearchCapabilitiesEncounterHistory() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EncounterHistory")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEncounterHistory(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EncounterHistory", options)
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
func (w Concrete) SearchCapabilitiesFormularyItem() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("FormularyItem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchFormularyItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "FormularyItem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesGenomicStudy() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("GenomicStudy")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchGenomicStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "GenomicStudy", options)
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
func (w Concrete) SearchCapabilitiesImagingSelection() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("ImagingSelection")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchImagingSelection(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "ImagingSelection", options)
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
func (w Concrete) SearchCapabilitiesInventoryItem() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("InventoryItem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchInventoryItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "InventoryItem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesInventoryReport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("InventoryReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchInventoryReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "InventoryReport", options)
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
func (w Concrete) SearchCapabilitiesNutritionIntake() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("NutritionIntake")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchNutritionIntake(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "NutritionIntake", options)
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
func (w Concrete) SearchCapabilitiesPermission() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Permission")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchPermission(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Permission", options)
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
func (w Concrete) SearchCapabilitiesRequestOrchestration() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RequestOrchestration")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRequestOrchestration(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RequestOrchestration", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesRequirements() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Requirements")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRequirements(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Requirements", options)
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
func (w Concrete) SearchCapabilitiesSubstanceNucleicAcid() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceNucleicAcid")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceNucleicAcid(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceNucleicAcid", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstancePolymer() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstancePolymer")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstancePolymer(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstancePolymer", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstanceProtein() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceProtein")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceProtein(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceProtein", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstanceReferenceInformation() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceReferenceInformation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceReferenceInformation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceReferenceInformation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesSubstanceSourceMaterial() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceSourceMaterial")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceSourceMaterial(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceSourceMaterial", options)
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
func (w Concrete) SearchCapabilitiesTestPlan() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("TestPlan")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTestPlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "TestPlan", options)
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
func (w Concrete) SearchCapabilitiesTransport() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("Transport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchTransport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "Transport", options)
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
