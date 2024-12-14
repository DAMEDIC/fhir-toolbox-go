package capabilitiesR4

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	search "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	r4 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
)

type ConcreteAPI interface {
	AccountRead
	AccountSearch
	ActivityDefinitionRead
	ActivityDefinitionSearch
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
	ClaimRead
	ClaimSearch
	ClaimResponseRead
	ClaimResponseSearch
	ClinicalImpressionRead
	ClinicalImpressionSearch
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
	EffectEvidenceSynthesisRead
	EffectEvidenceSynthesisSearch
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
	MedicinalProductRead
	MedicinalProductSearch
	MedicinalProductAuthorizationRead
	MedicinalProductAuthorizationSearch
	MedicinalProductContraindicationRead
	MedicinalProductContraindicationSearch
	MedicinalProductIndicationRead
	MedicinalProductIndicationSearch
	MedicinalProductIngredientRead
	MedicinalProductIngredientSearch
	MedicinalProductInteractionRead
	MedicinalProductInteractionSearch
	MedicinalProductManufacturedRead
	MedicinalProductManufacturedSearch
	MedicinalProductPackagedRead
	MedicinalProductPackagedSearch
	MedicinalProductPharmaceuticalRead
	MedicinalProductPharmaceuticalSearch
	MedicinalProductUndesirableEffectRead
	MedicinalProductUndesirableEffectSearch
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
	RiskEvidenceSynthesisRead
	RiskEvidenceSynthesisSearch
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
	SubstanceRead
	SubstanceSearch
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
	SubstanceSpecificationRead
	SubstanceSpecificationSearch
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

func (w Concrete) ReadAccount(ctx context.Context, id string) (r4.Account, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Account", id)
	if err != nil {
		return r4.Account{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Account)
	if !ok {
		return r4.Account{}, capabilities.InvalidResourceError{ResourceType: "Account"}
	}
	return impl, nil
}
func (w Concrete) ReadActivityDefinition(ctx context.Context, id string) (r4.ActivityDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ActivityDefinition", id)
	if err != nil {
		return r4.ActivityDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ActivityDefinition)
	if !ok {
		return r4.ActivityDefinition{}, capabilities.InvalidResourceError{ResourceType: "ActivityDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadAdverseEvent(ctx context.Context, id string) (r4.AdverseEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AdverseEvent", id)
	if err != nil {
		return r4.AdverseEvent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AdverseEvent)
	if !ok {
		return r4.AdverseEvent{}, capabilities.InvalidResourceError{ResourceType: "AdverseEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadAllergyIntolerance(ctx context.Context, id string) (r4.AllergyIntolerance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AllergyIntolerance", id)
	if err != nil {
		return r4.AllergyIntolerance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AllergyIntolerance)
	if !ok {
		return r4.AllergyIntolerance{}, capabilities.InvalidResourceError{ResourceType: "AllergyIntolerance"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointment(ctx context.Context, id string) (r4.Appointment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Appointment", id)
	if err != nil {
		return r4.Appointment{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Appointment)
	if !ok {
		return r4.Appointment{}, capabilities.InvalidResourceError{ResourceType: "Appointment"}
	}
	return impl, nil
}
func (w Concrete) ReadAppointmentResponse(ctx context.Context, id string) (r4.AppointmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AppointmentResponse", id)
	if err != nil {
		return r4.AppointmentResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AppointmentResponse)
	if !ok {
		return r4.AppointmentResponse{}, capabilities.InvalidResourceError{ResourceType: "AppointmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadAuditEvent(ctx context.Context, id string) (r4.AuditEvent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "AuditEvent", id)
	if err != nil {
		return r4.AuditEvent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AuditEvent)
	if !ok {
		return r4.AuditEvent{}, capabilities.InvalidResourceError{ResourceType: "AuditEvent"}
	}
	return impl, nil
}
func (w Concrete) ReadBasic(ctx context.Context, id string) (r4.Basic, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Basic", id)
	if err != nil {
		return r4.Basic{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Basic)
	if !ok {
		return r4.Basic{}, capabilities.InvalidResourceError{ResourceType: "Basic"}
	}
	return impl, nil
}
func (w Concrete) ReadBinary(ctx context.Context, id string) (r4.Binary, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Binary", id)
	if err != nil {
		return r4.Binary{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Binary)
	if !ok {
		return r4.Binary{}, capabilities.InvalidResourceError{ResourceType: "Binary"}
	}
	return impl, nil
}
func (w Concrete) ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4.BiologicallyDerivedProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BiologicallyDerivedProduct", id)
	if err != nil {
		return r4.BiologicallyDerivedProduct{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.BiologicallyDerivedProduct)
	if !ok {
		return r4.BiologicallyDerivedProduct{}, capabilities.InvalidResourceError{ResourceType: "BiologicallyDerivedProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadBodyStructure(ctx context.Context, id string) (r4.BodyStructure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "BodyStructure", id)
	if err != nil {
		return r4.BodyStructure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.BodyStructure)
	if !ok {
		return r4.BodyStructure{}, capabilities.InvalidResourceError{ResourceType: "BodyStructure"}
	}
	return impl, nil
}
func (w Concrete) ReadBundle(ctx context.Context, id string) (r4.Bundle, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Bundle", id)
	if err != nil {
		return r4.Bundle{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Bundle)
	if !ok {
		return r4.Bundle{}, capabilities.InvalidResourceError{ResourceType: "Bundle"}
	}
	return impl, nil
}
func (w Concrete) ReadCapabilityStatement(ctx context.Context, id string) (r4.CapabilityStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CapabilityStatement", id)
	if err != nil {
		return r4.CapabilityStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CapabilityStatement)
	if !ok {
		return r4.CapabilityStatement{}, capabilities.InvalidResourceError{ResourceType: "CapabilityStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadCarePlan(ctx context.Context, id string) (r4.CarePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CarePlan", id)
	if err != nil {
		return r4.CarePlan{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CarePlan)
	if !ok {
		return r4.CarePlan{}, capabilities.InvalidResourceError{ResourceType: "CarePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadCareTeam(ctx context.Context, id string) (r4.CareTeam, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CareTeam", id)
	if err != nil {
		return r4.CareTeam{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CareTeam)
	if !ok {
		return r4.CareTeam{}, capabilities.InvalidResourceError{ResourceType: "CareTeam"}
	}
	return impl, nil
}
func (w Concrete) ReadCatalogEntry(ctx context.Context, id string) (r4.CatalogEntry, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CatalogEntry", id)
	if err != nil {
		return r4.CatalogEntry{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CatalogEntry)
	if !ok {
		return r4.CatalogEntry{}, capabilities.InvalidResourceError{ResourceType: "CatalogEntry"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItem(ctx context.Context, id string) (r4.ChargeItem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItem", id)
	if err != nil {
		return r4.ChargeItem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ChargeItem)
	if !ok {
		return r4.ChargeItem{}, capabilities.InvalidResourceError{ResourceType: "ChargeItem"}
	}
	return impl, nil
}
func (w Concrete) ReadChargeItemDefinition(ctx context.Context, id string) (r4.ChargeItemDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ChargeItemDefinition", id)
	if err != nil {
		return r4.ChargeItemDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ChargeItemDefinition)
	if !ok {
		return r4.ChargeItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ChargeItemDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadClaim(ctx context.Context, id string) (r4.Claim, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Claim", id)
	if err != nil {
		return r4.Claim{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Claim)
	if !ok {
		return r4.Claim{}, capabilities.InvalidResourceError{ResourceType: "Claim"}
	}
	return impl, nil
}
func (w Concrete) ReadClaimResponse(ctx context.Context, id string) (r4.ClaimResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClaimResponse", id)
	if err != nil {
		return r4.ClaimResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ClaimResponse)
	if !ok {
		return r4.ClaimResponse{}, capabilities.InvalidResourceError{ResourceType: "ClaimResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadClinicalImpression(ctx context.Context, id string) (r4.ClinicalImpression, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ClinicalImpression", id)
	if err != nil {
		return r4.ClinicalImpression{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ClinicalImpression)
	if !ok {
		return r4.ClinicalImpression{}, capabilities.InvalidResourceError{ResourceType: "ClinicalImpression"}
	}
	return impl, nil
}
func (w Concrete) ReadCodeSystem(ctx context.Context, id string) (r4.CodeSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CodeSystem", id)
	if err != nil {
		return r4.CodeSystem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CodeSystem)
	if !ok {
		return r4.CodeSystem{}, capabilities.InvalidResourceError{ResourceType: "CodeSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunication(ctx context.Context, id string) (r4.Communication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Communication", id)
	if err != nil {
		return r4.Communication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Communication)
	if !ok {
		return r4.Communication{}, capabilities.InvalidResourceError{ResourceType: "Communication"}
	}
	return impl, nil
}
func (w Concrete) ReadCommunicationRequest(ctx context.Context, id string) (r4.CommunicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CommunicationRequest", id)
	if err != nil {
		return r4.CommunicationRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CommunicationRequest)
	if !ok {
		return r4.CommunicationRequest{}, capabilities.InvalidResourceError{ResourceType: "CommunicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCompartmentDefinition(ctx context.Context, id string) (r4.CompartmentDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CompartmentDefinition", id)
	if err != nil {
		return r4.CompartmentDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CompartmentDefinition)
	if !ok {
		return r4.CompartmentDefinition{}, capabilities.InvalidResourceError{ResourceType: "CompartmentDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadComposition(ctx context.Context, id string) (r4.Composition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Composition", id)
	if err != nil {
		return r4.Composition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Composition)
	if !ok {
		return r4.Composition{}, capabilities.InvalidResourceError{ResourceType: "Composition"}
	}
	return impl, nil
}
func (w Concrete) ReadConceptMap(ctx context.Context, id string) (r4.ConceptMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ConceptMap", id)
	if err != nil {
		return r4.ConceptMap{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ConceptMap)
	if !ok {
		return r4.ConceptMap{}, capabilities.InvalidResourceError{ResourceType: "ConceptMap"}
	}
	return impl, nil
}
func (w Concrete) ReadCondition(ctx context.Context, id string) (r4.Condition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Condition", id)
	if err != nil {
		return r4.Condition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Condition)
	if !ok {
		return r4.Condition{}, capabilities.InvalidResourceError{ResourceType: "Condition"}
	}
	return impl, nil
}
func (w Concrete) ReadConsent(ctx context.Context, id string) (r4.Consent, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Consent", id)
	if err != nil {
		return r4.Consent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Consent)
	if !ok {
		return r4.Consent{}, capabilities.InvalidResourceError{ResourceType: "Consent"}
	}
	return impl, nil
}
func (w Concrete) ReadContract(ctx context.Context, id string) (r4.Contract, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Contract", id)
	if err != nil {
		return r4.Contract{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Contract)
	if !ok {
		return r4.Contract{}, capabilities.InvalidResourceError{ResourceType: "Contract"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverage(ctx context.Context, id string) (r4.Coverage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Coverage", id)
	if err != nil {
		return r4.Coverage{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Coverage)
	if !ok {
		return r4.Coverage{}, capabilities.InvalidResourceError{ResourceType: "Coverage"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4.CoverageEligibilityRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityRequest", id)
	if err != nil {
		return r4.CoverageEligibilityRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CoverageEligibilityRequest)
	if !ok {
		return r4.CoverageEligibilityRequest{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4.CoverageEligibilityResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "CoverageEligibilityResponse", id)
	if err != nil {
		return r4.CoverageEligibilityResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CoverageEligibilityResponse)
	if !ok {
		return r4.CoverageEligibilityResponse{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadDetectedIssue(ctx context.Context, id string) (r4.DetectedIssue, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DetectedIssue", id)
	if err != nil {
		return r4.DetectedIssue{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DetectedIssue)
	if !ok {
		return r4.DetectedIssue{}, capabilities.InvalidResourceError{ResourceType: "DetectedIssue"}
	}
	return impl, nil
}
func (w Concrete) ReadDevice(ctx context.Context, id string) (r4.Device, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Device", id)
	if err != nil {
		return r4.Device{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Device)
	if !ok {
		return r4.Device{}, capabilities.InvalidResourceError{ResourceType: "Device"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceDefinition(ctx context.Context, id string) (r4.DeviceDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceDefinition", id)
	if err != nil {
		return r4.DeviceDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceDefinition)
	if !ok {
		return r4.DeviceDefinition{}, capabilities.InvalidResourceError{ResourceType: "DeviceDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceMetric(ctx context.Context, id string) (r4.DeviceMetric, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceMetric", id)
	if err != nil {
		return r4.DeviceMetric{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceMetric)
	if !ok {
		return r4.DeviceMetric{}, capabilities.InvalidResourceError{ResourceType: "DeviceMetric"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceRequest(ctx context.Context, id string) (r4.DeviceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceRequest", id)
	if err != nil {
		return r4.DeviceRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceRequest)
	if !ok {
		return r4.DeviceRequest{}, capabilities.InvalidResourceError{ResourceType: "DeviceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadDeviceUseStatement(ctx context.Context, id string) (r4.DeviceUseStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DeviceUseStatement", id)
	if err != nil {
		return r4.DeviceUseStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceUseStatement)
	if !ok {
		return r4.DeviceUseStatement{}, capabilities.InvalidResourceError{ResourceType: "DeviceUseStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadDiagnosticReport(ctx context.Context, id string) (r4.DiagnosticReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DiagnosticReport", id)
	if err != nil {
		return r4.DiagnosticReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DiagnosticReport)
	if !ok {
		return r4.DiagnosticReport{}, capabilities.InvalidResourceError{ResourceType: "DiagnosticReport"}
	}
	return impl, nil
}
func (w Concrete) ReadDocumentManifest(ctx context.Context, id string) (r4.DocumentManifest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DocumentManifest", id)
	if err != nil {
		return r4.DocumentManifest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DocumentManifest)
	if !ok {
		return r4.DocumentManifest{}, capabilities.InvalidResourceError{ResourceType: "DocumentManifest"}
	}
	return impl, nil
}
func (w Concrete) ReadDocumentReference(ctx context.Context, id string) (r4.DocumentReference, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "DocumentReference", id)
	if err != nil {
		return r4.DocumentReference{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DocumentReference)
	if !ok {
		return r4.DocumentReference{}, capabilities.InvalidResourceError{ResourceType: "DocumentReference"}
	}
	return impl, nil
}
func (w Concrete) ReadEffectEvidenceSynthesis(ctx context.Context, id string) (r4.EffectEvidenceSynthesis, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EffectEvidenceSynthesis", id)
	if err != nil {
		return r4.EffectEvidenceSynthesis{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EffectEvidenceSynthesis)
	if !ok {
		return r4.EffectEvidenceSynthesis{}, capabilities.InvalidResourceError{ResourceType: "EffectEvidenceSynthesis"}
	}
	return impl, nil
}
func (w Concrete) ReadEncounter(ctx context.Context, id string) (r4.Encounter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Encounter", id)
	if err != nil {
		return r4.Encounter{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Encounter)
	if !ok {
		return r4.Encounter{}, capabilities.InvalidResourceError{ResourceType: "Encounter"}
	}
	return impl, nil
}
func (w Concrete) ReadEndpoint(ctx context.Context, id string) (r4.Endpoint, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Endpoint", id)
	if err != nil {
		return r4.Endpoint{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Endpoint)
	if !ok {
		return r4.Endpoint{}, capabilities.InvalidResourceError{ResourceType: "Endpoint"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentRequest(ctx context.Context, id string) (r4.EnrollmentRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentRequest", id)
	if err != nil {
		return r4.EnrollmentRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EnrollmentRequest)
	if !ok {
		return r4.EnrollmentRequest{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadEnrollmentResponse(ctx context.Context, id string) (r4.EnrollmentResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EnrollmentResponse", id)
	if err != nil {
		return r4.EnrollmentResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EnrollmentResponse)
	if !ok {
		return r4.EnrollmentResponse{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadEpisodeOfCare(ctx context.Context, id string) (r4.EpisodeOfCare, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EpisodeOfCare", id)
	if err != nil {
		return r4.EpisodeOfCare{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EpisodeOfCare)
	if !ok {
		return r4.EpisodeOfCare{}, capabilities.InvalidResourceError{ResourceType: "EpisodeOfCare"}
	}
	return impl, nil
}
func (w Concrete) ReadEventDefinition(ctx context.Context, id string) (r4.EventDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EventDefinition", id)
	if err != nil {
		return r4.EventDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EventDefinition)
	if !ok {
		return r4.EventDefinition{}, capabilities.InvalidResourceError{ResourceType: "EventDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidence(ctx context.Context, id string) (r4.Evidence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Evidence", id)
	if err != nil {
		return r4.Evidence{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Evidence)
	if !ok {
		return r4.Evidence{}, capabilities.InvalidResourceError{ResourceType: "Evidence"}
	}
	return impl, nil
}
func (w Concrete) ReadEvidenceVariable(ctx context.Context, id string) (r4.EvidenceVariable, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "EvidenceVariable", id)
	if err != nil {
		return r4.EvidenceVariable{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EvidenceVariable)
	if !ok {
		return r4.EvidenceVariable{}, capabilities.InvalidResourceError{ResourceType: "EvidenceVariable"}
	}
	return impl, nil
}
func (w Concrete) ReadExampleScenario(ctx context.Context, id string) (r4.ExampleScenario, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExampleScenario", id)
	if err != nil {
		return r4.ExampleScenario{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ExampleScenario)
	if !ok {
		return r4.ExampleScenario{}, capabilities.InvalidResourceError{ResourceType: "ExampleScenario"}
	}
	return impl, nil
}
func (w Concrete) ReadExplanationOfBenefit(ctx context.Context, id string) (r4.ExplanationOfBenefit, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ExplanationOfBenefit", id)
	if err != nil {
		return r4.ExplanationOfBenefit{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ExplanationOfBenefit)
	if !ok {
		return r4.ExplanationOfBenefit{}, capabilities.InvalidResourceError{ResourceType: "ExplanationOfBenefit"}
	}
	return impl, nil
}
func (w Concrete) ReadFamilyMemberHistory(ctx context.Context, id string) (r4.FamilyMemberHistory, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "FamilyMemberHistory", id)
	if err != nil {
		return r4.FamilyMemberHistory{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.FamilyMemberHistory)
	if !ok {
		return r4.FamilyMemberHistory{}, capabilities.InvalidResourceError{ResourceType: "FamilyMemberHistory"}
	}
	return impl, nil
}
func (w Concrete) ReadFlag(ctx context.Context, id string) (r4.Flag, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Flag", id)
	if err != nil {
		return r4.Flag{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Flag)
	if !ok {
		return r4.Flag{}, capabilities.InvalidResourceError{ResourceType: "Flag"}
	}
	return impl, nil
}
func (w Concrete) ReadGoal(ctx context.Context, id string) (r4.Goal, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Goal", id)
	if err != nil {
		return r4.Goal{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Goal)
	if !ok {
		return r4.Goal{}, capabilities.InvalidResourceError{ResourceType: "Goal"}
	}
	return impl, nil
}
func (w Concrete) ReadGraphDefinition(ctx context.Context, id string) (r4.GraphDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GraphDefinition", id)
	if err != nil {
		return r4.GraphDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.GraphDefinition)
	if !ok {
		return r4.GraphDefinition{}, capabilities.InvalidResourceError{ResourceType: "GraphDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadGroup(ctx context.Context, id string) (r4.Group, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Group", id)
	if err != nil {
		return r4.Group{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Group)
	if !ok {
		return r4.Group{}, capabilities.InvalidResourceError{ResourceType: "Group"}
	}
	return impl, nil
}
func (w Concrete) ReadGuidanceResponse(ctx context.Context, id string) (r4.GuidanceResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "GuidanceResponse", id)
	if err != nil {
		return r4.GuidanceResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.GuidanceResponse)
	if !ok {
		return r4.GuidanceResponse{}, capabilities.InvalidResourceError{ResourceType: "GuidanceResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadHealthcareService(ctx context.Context, id string) (r4.HealthcareService, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "HealthcareService", id)
	if err != nil {
		return r4.HealthcareService{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.HealthcareService)
	if !ok {
		return r4.HealthcareService{}, capabilities.InvalidResourceError{ResourceType: "HealthcareService"}
	}
	return impl, nil
}
func (w Concrete) ReadImagingStudy(ctx context.Context, id string) (r4.ImagingStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImagingStudy", id)
	if err != nil {
		return r4.ImagingStudy{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImagingStudy)
	if !ok {
		return r4.ImagingStudy{}, capabilities.InvalidResourceError{ResourceType: "ImagingStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunization(ctx context.Context, id string) (r4.Immunization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Immunization", id)
	if err != nil {
		return r4.Immunization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Immunization)
	if !ok {
		return r4.Immunization{}, capabilities.InvalidResourceError{ResourceType: "Immunization"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationEvaluation(ctx context.Context, id string) (r4.ImmunizationEvaluation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationEvaluation", id)
	if err != nil {
		return r4.ImmunizationEvaluation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImmunizationEvaluation)
	if !ok {
		return r4.ImmunizationEvaluation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationEvaluation"}
	}
	return impl, nil
}
func (w Concrete) ReadImmunizationRecommendation(ctx context.Context, id string) (r4.ImmunizationRecommendation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImmunizationRecommendation", id)
	if err != nil {
		return r4.ImmunizationRecommendation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImmunizationRecommendation)
	if !ok {
		return r4.ImmunizationRecommendation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationRecommendation"}
	}
	return impl, nil
}
func (w Concrete) ReadImplementationGuide(ctx context.Context, id string) (r4.ImplementationGuide, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ImplementationGuide", id)
	if err != nil {
		return r4.ImplementationGuide{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImplementationGuide)
	if !ok {
		return r4.ImplementationGuide{}, capabilities.InvalidResourceError{ResourceType: "ImplementationGuide"}
	}
	return impl, nil
}
func (w Concrete) ReadInsurancePlan(ctx context.Context, id string) (r4.InsurancePlan, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "InsurancePlan", id)
	if err != nil {
		return r4.InsurancePlan{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.InsurancePlan)
	if !ok {
		return r4.InsurancePlan{}, capabilities.InvalidResourceError{ResourceType: "InsurancePlan"}
	}
	return impl, nil
}
func (w Concrete) ReadInvoice(ctx context.Context, id string) (r4.Invoice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Invoice", id)
	if err != nil {
		return r4.Invoice{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Invoice)
	if !ok {
		return r4.Invoice{}, capabilities.InvalidResourceError{ResourceType: "Invoice"}
	}
	return impl, nil
}
func (w Concrete) ReadLibrary(ctx context.Context, id string) (r4.Library, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Library", id)
	if err != nil {
		return r4.Library{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Library)
	if !ok {
		return r4.Library{}, capabilities.InvalidResourceError{ResourceType: "Library"}
	}
	return impl, nil
}
func (w Concrete) ReadLinkage(ctx context.Context, id string) (r4.Linkage, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Linkage", id)
	if err != nil {
		return r4.Linkage{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Linkage)
	if !ok {
		return r4.Linkage{}, capabilities.InvalidResourceError{ResourceType: "Linkage"}
	}
	return impl, nil
}
func (w Concrete) ReadList(ctx context.Context, id string) (r4.List, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "List", id)
	if err != nil {
		return r4.List{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.List)
	if !ok {
		return r4.List{}, capabilities.InvalidResourceError{ResourceType: "List"}
	}
	return impl, nil
}
func (w Concrete) ReadLocation(ctx context.Context, id string) (r4.Location, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Location", id)
	if err != nil {
		return r4.Location{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Location)
	if !ok {
		return r4.Location{}, capabilities.InvalidResourceError{ResourceType: "Location"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasure(ctx context.Context, id string) (r4.Measure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Measure", id)
	if err != nil {
		return r4.Measure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Measure)
	if !ok {
		return r4.Measure{}, capabilities.InvalidResourceError{ResourceType: "Measure"}
	}
	return impl, nil
}
func (w Concrete) ReadMeasureReport(ctx context.Context, id string) (r4.MeasureReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MeasureReport", id)
	if err != nil {
		return r4.MeasureReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MeasureReport)
	if !ok {
		return r4.MeasureReport{}, capabilities.InvalidResourceError{ResourceType: "MeasureReport"}
	}
	return impl, nil
}
func (w Concrete) ReadMedia(ctx context.Context, id string) (r4.Media, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Media", id)
	if err != nil {
		return r4.Media{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Media)
	if !ok {
		return r4.Media{}, capabilities.InvalidResourceError{ResourceType: "Media"}
	}
	return impl, nil
}
func (w Concrete) ReadMedication(ctx context.Context, id string) (r4.Medication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Medication", id)
	if err != nil {
		return r4.Medication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Medication)
	if !ok {
		return r4.Medication{}, capabilities.InvalidResourceError{ResourceType: "Medication"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationAdministration(ctx context.Context, id string) (r4.MedicationAdministration, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationAdministration", id)
	if err != nil {
		return r4.MedicationAdministration{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationAdministration)
	if !ok {
		return r4.MedicationAdministration{}, capabilities.InvalidResourceError{ResourceType: "MedicationAdministration"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationDispense(ctx context.Context, id string) (r4.MedicationDispense, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationDispense", id)
	if err != nil {
		return r4.MedicationDispense{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationDispense)
	if !ok {
		return r4.MedicationDispense{}, capabilities.InvalidResourceError{ResourceType: "MedicationDispense"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationKnowledge(ctx context.Context, id string) (r4.MedicationKnowledge, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationKnowledge", id)
	if err != nil {
		return r4.MedicationKnowledge{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationKnowledge)
	if !ok {
		return r4.MedicationKnowledge{}, capabilities.InvalidResourceError{ResourceType: "MedicationKnowledge"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationRequest(ctx context.Context, id string) (r4.MedicationRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationRequest", id)
	if err != nil {
		return r4.MedicationRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationRequest)
	if !ok {
		return r4.MedicationRequest{}, capabilities.InvalidResourceError{ResourceType: "MedicationRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicationStatement(ctx context.Context, id string) (r4.MedicationStatement, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicationStatement", id)
	if err != nil {
		return r4.MedicationStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationStatement)
	if !ok {
		return r4.MedicationStatement{}, capabilities.InvalidResourceError{ResourceType: "MedicationStatement"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProduct(ctx context.Context, id string) (r4.MedicinalProduct, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProduct", id)
	if err != nil {
		return r4.MedicinalProduct{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProduct)
	if !ok {
		return r4.MedicinalProduct{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProduct"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductAuthorization(ctx context.Context, id string) (r4.MedicinalProductAuthorization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductAuthorization", id)
	if err != nil {
		return r4.MedicinalProductAuthorization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductAuthorization)
	if !ok {
		return r4.MedicinalProductAuthorization{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductAuthorization"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductContraindication(ctx context.Context, id string) (r4.MedicinalProductContraindication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductContraindication", id)
	if err != nil {
		return r4.MedicinalProductContraindication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductContraindication)
	if !ok {
		return r4.MedicinalProductContraindication{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductContraindication"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductIndication(ctx context.Context, id string) (r4.MedicinalProductIndication, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductIndication", id)
	if err != nil {
		return r4.MedicinalProductIndication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductIndication)
	if !ok {
		return r4.MedicinalProductIndication{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductIndication"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductIngredient(ctx context.Context, id string) (r4.MedicinalProductIngredient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductIngredient", id)
	if err != nil {
		return r4.MedicinalProductIngredient{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductIngredient)
	if !ok {
		return r4.MedicinalProductIngredient{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductIngredient"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductInteraction(ctx context.Context, id string) (r4.MedicinalProductInteraction, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductInteraction", id)
	if err != nil {
		return r4.MedicinalProductInteraction{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductInteraction)
	if !ok {
		return r4.MedicinalProductInteraction{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductInteraction"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductManufactured(ctx context.Context, id string) (r4.MedicinalProductManufactured, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductManufactured", id)
	if err != nil {
		return r4.MedicinalProductManufactured{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductManufactured)
	if !ok {
		return r4.MedicinalProductManufactured{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductManufactured"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductPackaged(ctx context.Context, id string) (r4.MedicinalProductPackaged, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductPackaged", id)
	if err != nil {
		return r4.MedicinalProductPackaged{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductPackaged)
	if !ok {
		return r4.MedicinalProductPackaged{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductPackaged"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (r4.MedicinalProductPharmaceutical, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductPharmaceutical", id)
	if err != nil {
		return r4.MedicinalProductPharmaceutical{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductPharmaceutical)
	if !ok {
		return r4.MedicinalProductPharmaceutical{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductPharmaceutical"}
	}
	return impl, nil
}
func (w Concrete) ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (r4.MedicinalProductUndesirableEffect, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MedicinalProductUndesirableEffect", id)
	if err != nil {
		return r4.MedicinalProductUndesirableEffect{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductUndesirableEffect)
	if !ok {
		return r4.MedicinalProductUndesirableEffect{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductUndesirableEffect"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageDefinition(ctx context.Context, id string) (r4.MessageDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageDefinition", id)
	if err != nil {
		return r4.MessageDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MessageDefinition)
	if !ok {
		return r4.MessageDefinition{}, capabilities.InvalidResourceError{ResourceType: "MessageDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadMessageHeader(ctx context.Context, id string) (r4.MessageHeader, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MessageHeader", id)
	if err != nil {
		return r4.MessageHeader{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MessageHeader)
	if !ok {
		return r4.MessageHeader{}, capabilities.InvalidResourceError{ResourceType: "MessageHeader"}
	}
	return impl, nil
}
func (w Concrete) ReadMolecularSequence(ctx context.Context, id string) (r4.MolecularSequence, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "MolecularSequence", id)
	if err != nil {
		return r4.MolecularSequence{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MolecularSequence)
	if !ok {
		return r4.MolecularSequence{}, capabilities.InvalidResourceError{ResourceType: "MolecularSequence"}
	}
	return impl, nil
}
func (w Concrete) ReadNamingSystem(ctx context.Context, id string) (r4.NamingSystem, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NamingSystem", id)
	if err != nil {
		return r4.NamingSystem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.NamingSystem)
	if !ok {
		return r4.NamingSystem{}, capabilities.InvalidResourceError{ResourceType: "NamingSystem"}
	}
	return impl, nil
}
func (w Concrete) ReadNutritionOrder(ctx context.Context, id string) (r4.NutritionOrder, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "NutritionOrder", id)
	if err != nil {
		return r4.NutritionOrder{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.NutritionOrder)
	if !ok {
		return r4.NutritionOrder{}, capabilities.InvalidResourceError{ResourceType: "NutritionOrder"}
	}
	return impl, nil
}
func (w Concrete) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Observation", id)
	if err != nil {
		return r4.Observation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Observation)
	if !ok {
		return r4.Observation{}, capabilities.InvalidResourceError{ResourceType: "Observation"}
	}
	return impl, nil
}
func (w Concrete) ReadObservationDefinition(ctx context.Context, id string) (r4.ObservationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ObservationDefinition", id)
	if err != nil {
		return r4.ObservationDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ObservationDefinition)
	if !ok {
		return r4.ObservationDefinition{}, capabilities.InvalidResourceError{ResourceType: "ObservationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationDefinition(ctx context.Context, id string) (r4.OperationDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationDefinition", id)
	if err != nil {
		return r4.OperationDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OperationDefinition)
	if !ok {
		return r4.OperationDefinition{}, capabilities.InvalidResourceError{ResourceType: "OperationDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadOperationOutcome(ctx context.Context, id string) (r4.OperationOutcome, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OperationOutcome", id)
	if err != nil {
		return r4.OperationOutcome{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OperationOutcome)
	if !ok {
		return r4.OperationOutcome{}, capabilities.InvalidResourceError{ResourceType: "OperationOutcome"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganization(ctx context.Context, id string) (r4.Organization, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Organization", id)
	if err != nil {
		return r4.Organization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Organization)
	if !ok {
		return r4.Organization{}, capabilities.InvalidResourceError{ResourceType: "Organization"}
	}
	return impl, nil
}
func (w Concrete) ReadOrganizationAffiliation(ctx context.Context, id string) (r4.OrganizationAffiliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "OrganizationAffiliation", id)
	if err != nil {
		return r4.OrganizationAffiliation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OrganizationAffiliation)
	if !ok {
		return r4.OrganizationAffiliation{}, capabilities.InvalidResourceError{ResourceType: "OrganizationAffiliation"}
	}
	return impl, nil
}
func (w Concrete) ReadParameters(ctx context.Context, id string) (r4.Parameters, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Parameters", id)
	if err != nil {
		return r4.Parameters{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Parameters)
	if !ok {
		return r4.Parameters{}, capabilities.InvalidResourceError{ResourceType: "Parameters"}
	}
	return impl, nil
}
func (w Concrete) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Patient", id)
	if err != nil {
		return r4.Patient{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Patient)
	if !ok {
		return r4.Patient{}, capabilities.InvalidResourceError{ResourceType: "Patient"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentNotice(ctx context.Context, id string) (r4.PaymentNotice, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentNotice", id)
	if err != nil {
		return r4.PaymentNotice{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PaymentNotice)
	if !ok {
		return r4.PaymentNotice{}, capabilities.InvalidResourceError{ResourceType: "PaymentNotice"}
	}
	return impl, nil
}
func (w Concrete) ReadPaymentReconciliation(ctx context.Context, id string) (r4.PaymentReconciliation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PaymentReconciliation", id)
	if err != nil {
		return r4.PaymentReconciliation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PaymentReconciliation)
	if !ok {
		return r4.PaymentReconciliation{}, capabilities.InvalidResourceError{ResourceType: "PaymentReconciliation"}
	}
	return impl, nil
}
func (w Concrete) ReadPerson(ctx context.Context, id string) (r4.Person, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Person", id)
	if err != nil {
		return r4.Person{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Person)
	if !ok {
		return r4.Person{}, capabilities.InvalidResourceError{ResourceType: "Person"}
	}
	return impl, nil
}
func (w Concrete) ReadPlanDefinition(ctx context.Context, id string) (r4.PlanDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PlanDefinition", id)
	if err != nil {
		return r4.PlanDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PlanDefinition)
	if !ok {
		return r4.PlanDefinition{}, capabilities.InvalidResourceError{ResourceType: "PlanDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitioner(ctx context.Context, id string) (r4.Practitioner, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Practitioner", id)
	if err != nil {
		return r4.Practitioner{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Practitioner)
	if !ok {
		return r4.Practitioner{}, capabilities.InvalidResourceError{ResourceType: "Practitioner"}
	}
	return impl, nil
}
func (w Concrete) ReadPractitionerRole(ctx context.Context, id string) (r4.PractitionerRole, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "PractitionerRole", id)
	if err != nil {
		return r4.PractitionerRole{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PractitionerRole)
	if !ok {
		return r4.PractitionerRole{}, capabilities.InvalidResourceError{ResourceType: "PractitionerRole"}
	}
	return impl, nil
}
func (w Concrete) ReadProcedure(ctx context.Context, id string) (r4.Procedure, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Procedure", id)
	if err != nil {
		return r4.Procedure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Procedure)
	if !ok {
		return r4.Procedure{}, capabilities.InvalidResourceError{ResourceType: "Procedure"}
	}
	return impl, nil
}
func (w Concrete) ReadProvenance(ctx context.Context, id string) (r4.Provenance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Provenance", id)
	if err != nil {
		return r4.Provenance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Provenance)
	if !ok {
		return r4.Provenance{}, capabilities.InvalidResourceError{ResourceType: "Provenance"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaire(ctx context.Context, id string) (r4.Questionnaire, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Questionnaire", id)
	if err != nil {
		return r4.Questionnaire{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Questionnaire)
	if !ok {
		return r4.Questionnaire{}, capabilities.InvalidResourceError{ResourceType: "Questionnaire"}
	}
	return impl, nil
}
func (w Concrete) ReadQuestionnaireResponse(ctx context.Context, id string) (r4.QuestionnaireResponse, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "QuestionnaireResponse", id)
	if err != nil {
		return r4.QuestionnaireResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.QuestionnaireResponse)
	if !ok {
		return r4.QuestionnaireResponse{}, capabilities.InvalidResourceError{ResourceType: "QuestionnaireResponse"}
	}
	return impl, nil
}
func (w Concrete) ReadRelatedPerson(ctx context.Context, id string) (r4.RelatedPerson, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RelatedPerson", id)
	if err != nil {
		return r4.RelatedPerson{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RelatedPerson)
	if !ok {
		return r4.RelatedPerson{}, capabilities.InvalidResourceError{ResourceType: "RelatedPerson"}
	}
	return impl, nil
}
func (w Concrete) ReadRequestGroup(ctx context.Context, id string) (r4.RequestGroup, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RequestGroup", id)
	if err != nil {
		return r4.RequestGroup{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RequestGroup)
	if !ok {
		return r4.RequestGroup{}, capabilities.InvalidResourceError{ResourceType: "RequestGroup"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchDefinition(ctx context.Context, id string) (r4.ResearchDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchDefinition", id)
	if err != nil {
		return r4.ResearchDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchDefinition)
	if !ok {
		return r4.ResearchDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchElementDefinition(ctx context.Context, id string) (r4.ResearchElementDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchElementDefinition", id)
	if err != nil {
		return r4.ResearchElementDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchElementDefinition)
	if !ok {
		return r4.ResearchElementDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchElementDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchStudy(ctx context.Context, id string) (r4.ResearchStudy, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchStudy", id)
	if err != nil {
		return r4.ResearchStudy{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchStudy)
	if !ok {
		return r4.ResearchStudy{}, capabilities.InvalidResourceError{ResourceType: "ResearchStudy"}
	}
	return impl, nil
}
func (w Concrete) ReadResearchSubject(ctx context.Context, id string) (r4.ResearchSubject, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ResearchSubject", id)
	if err != nil {
		return r4.ResearchSubject{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchSubject)
	if !ok {
		return r4.ResearchSubject{}, capabilities.InvalidResourceError{ResourceType: "ResearchSubject"}
	}
	return impl, nil
}
func (w Concrete) ReadRiskAssessment(ctx context.Context, id string) (r4.RiskAssessment, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RiskAssessment", id)
	if err != nil {
		return r4.RiskAssessment{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RiskAssessment)
	if !ok {
		return r4.RiskAssessment{}, capabilities.InvalidResourceError{ResourceType: "RiskAssessment"}
	}
	return impl, nil
}
func (w Concrete) ReadRiskEvidenceSynthesis(ctx context.Context, id string) (r4.RiskEvidenceSynthesis, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "RiskEvidenceSynthesis", id)
	if err != nil {
		return r4.RiskEvidenceSynthesis{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RiskEvidenceSynthesis)
	if !ok {
		return r4.RiskEvidenceSynthesis{}, capabilities.InvalidResourceError{ResourceType: "RiskEvidenceSynthesis"}
	}
	return impl, nil
}
func (w Concrete) ReadSchedule(ctx context.Context, id string) (r4.Schedule, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Schedule", id)
	if err != nil {
		return r4.Schedule{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Schedule)
	if !ok {
		return r4.Schedule{}, capabilities.InvalidResourceError{ResourceType: "Schedule"}
	}
	return impl, nil
}
func (w Concrete) ReadSearchParameter(ctx context.Context, id string) (r4.SearchParameter, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SearchParameter", id)
	if err != nil {
		return r4.SearchParameter{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SearchParameter)
	if !ok {
		return r4.SearchParameter{}, capabilities.InvalidResourceError{ResourceType: "SearchParameter"}
	}
	return impl, nil
}
func (w Concrete) ReadServiceRequest(ctx context.Context, id string) (r4.ServiceRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ServiceRequest", id)
	if err != nil {
		return r4.ServiceRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ServiceRequest)
	if !ok {
		return r4.ServiceRequest{}, capabilities.InvalidResourceError{ResourceType: "ServiceRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadSlot(ctx context.Context, id string) (r4.Slot, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Slot", id)
	if err != nil {
		return r4.Slot{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Slot)
	if !ok {
		return r4.Slot{}, capabilities.InvalidResourceError{ResourceType: "Slot"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimen(ctx context.Context, id string) (r4.Specimen, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Specimen", id)
	if err != nil {
		return r4.Specimen{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Specimen)
	if !ok {
		return r4.Specimen{}, capabilities.InvalidResourceError{ResourceType: "Specimen"}
	}
	return impl, nil
}
func (w Concrete) ReadSpecimenDefinition(ctx context.Context, id string) (r4.SpecimenDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SpecimenDefinition", id)
	if err != nil {
		return r4.SpecimenDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SpecimenDefinition)
	if !ok {
		return r4.SpecimenDefinition{}, capabilities.InvalidResourceError{ResourceType: "SpecimenDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureDefinition(ctx context.Context, id string) (r4.StructureDefinition, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureDefinition", id)
	if err != nil {
		return r4.StructureDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.StructureDefinition)
	if !ok {
		return r4.StructureDefinition{}, capabilities.InvalidResourceError{ResourceType: "StructureDefinition"}
	}
	return impl, nil
}
func (w Concrete) ReadStructureMap(ctx context.Context, id string) (r4.StructureMap, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "StructureMap", id)
	if err != nil {
		return r4.StructureMap{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.StructureMap)
	if !ok {
		return r4.StructureMap{}, capabilities.InvalidResourceError{ResourceType: "StructureMap"}
	}
	return impl, nil
}
func (w Concrete) ReadSubscription(ctx context.Context, id string) (r4.Subscription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Subscription", id)
	if err != nil {
		return r4.Subscription{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Subscription)
	if !ok {
		return r4.Subscription{}, capabilities.InvalidResourceError{ResourceType: "Subscription"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstance(ctx context.Context, id string) (r4.Substance, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Substance", id)
	if err != nil {
		return r4.Substance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Substance)
	if !ok {
		return r4.Substance{}, capabilities.InvalidResourceError{ResourceType: "Substance"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceNucleicAcid(ctx context.Context, id string) (r4.SubstanceNucleicAcid, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceNucleicAcid", id)
	if err != nil {
		return r4.SubstanceNucleicAcid{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceNucleicAcid)
	if !ok {
		return r4.SubstanceNucleicAcid{}, capabilities.InvalidResourceError{ResourceType: "SubstanceNucleicAcid"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstancePolymer(ctx context.Context, id string) (r4.SubstancePolymer, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstancePolymer", id)
	if err != nil {
		return r4.SubstancePolymer{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstancePolymer)
	if !ok {
		return r4.SubstancePolymer{}, capabilities.InvalidResourceError{ResourceType: "SubstancePolymer"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceProtein(ctx context.Context, id string) (r4.SubstanceProtein, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceProtein", id)
	if err != nil {
		return r4.SubstanceProtein{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceProtein)
	if !ok {
		return r4.SubstanceProtein{}, capabilities.InvalidResourceError{ResourceType: "SubstanceProtein"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceReferenceInformation(ctx context.Context, id string) (r4.SubstanceReferenceInformation, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceReferenceInformation", id)
	if err != nil {
		return r4.SubstanceReferenceInformation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceReferenceInformation)
	if !ok {
		return r4.SubstanceReferenceInformation{}, capabilities.InvalidResourceError{ResourceType: "SubstanceReferenceInformation"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceSourceMaterial(ctx context.Context, id string) (r4.SubstanceSourceMaterial, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceSourceMaterial", id)
	if err != nil {
		return r4.SubstanceSourceMaterial{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceSourceMaterial)
	if !ok {
		return r4.SubstanceSourceMaterial{}, capabilities.InvalidResourceError{ResourceType: "SubstanceSourceMaterial"}
	}
	return impl, nil
}
func (w Concrete) ReadSubstanceSpecification(ctx context.Context, id string) (r4.SubstanceSpecification, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SubstanceSpecification", id)
	if err != nil {
		return r4.SubstanceSpecification{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceSpecification)
	if !ok {
		return r4.SubstanceSpecification{}, capabilities.InvalidResourceError{ResourceType: "SubstanceSpecification"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyDelivery(ctx context.Context, id string) (r4.SupplyDelivery, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyDelivery", id)
	if err != nil {
		return r4.SupplyDelivery{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SupplyDelivery)
	if !ok {
		return r4.SupplyDelivery{}, capabilities.InvalidResourceError{ResourceType: "SupplyDelivery"}
	}
	return impl, nil
}
func (w Concrete) ReadSupplyRequest(ctx context.Context, id string) (r4.SupplyRequest, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "SupplyRequest", id)
	if err != nil {
		return r4.SupplyRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SupplyRequest)
	if !ok {
		return r4.SupplyRequest{}, capabilities.InvalidResourceError{ResourceType: "SupplyRequest"}
	}
	return impl, nil
}
func (w Concrete) ReadTask(ctx context.Context, id string) (r4.Task, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "Task", id)
	if err != nil {
		return r4.Task{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Task)
	if !ok {
		return r4.Task{}, capabilities.InvalidResourceError{ResourceType: "Task"}
	}
	return impl, nil
}
func (w Concrete) ReadTerminologyCapabilities(ctx context.Context, id string) (r4.TerminologyCapabilities, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TerminologyCapabilities", id)
	if err != nil {
		return r4.TerminologyCapabilities{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TerminologyCapabilities)
	if !ok {
		return r4.TerminologyCapabilities{}, capabilities.InvalidResourceError{ResourceType: "TerminologyCapabilities"}
	}
	return impl, nil
}
func (w Concrete) ReadTestReport(ctx context.Context, id string) (r4.TestReport, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestReport", id)
	if err != nil {
		return r4.TestReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TestReport)
	if !ok {
		return r4.TestReport{}, capabilities.InvalidResourceError{ResourceType: "TestReport"}
	}
	return impl, nil
}
func (w Concrete) ReadTestScript(ctx context.Context, id string) (r4.TestScript, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "TestScript", id)
	if err != nil {
		return r4.TestScript{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TestScript)
	if !ok {
		return r4.TestScript{}, capabilities.InvalidResourceError{ResourceType: "TestScript"}
	}
	return impl, nil
}
func (w Concrete) ReadValueSet(ctx context.Context, id string) (r4.ValueSet, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "ValueSet", id)
	if err != nil {
		return r4.ValueSet{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ValueSet)
	if !ok {
		return r4.ValueSet{}, capabilities.InvalidResourceError{ResourceType: "ValueSet"}
	}
	return impl, nil
}
func (w Concrete) ReadVerificationResult(ctx context.Context, id string) (r4.VerificationResult, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VerificationResult", id)
	if err != nil {
		return r4.VerificationResult{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.VerificationResult)
	if !ok {
		return r4.VerificationResult{}, capabilities.InvalidResourceError{ResourceType: "VerificationResult"}
	}
	return impl, nil
}
func (w Concrete) ReadVisionPrescription(ctx context.Context, id string) (r4.VisionPrescription, capabilities.FHIRError) {
	v, err := w.Generic.Read(ctx, "VisionPrescription", id)
	if err != nil {
		return r4.VisionPrescription{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.VisionPrescription)
	if !ok {
		return r4.VisionPrescription{}, capabilities.InvalidResourceError{ResourceType: "VisionPrescription"}
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
func (w Concrete) SearchCapabilitiesEffectEvidenceSynthesis() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("EffectEvidenceSynthesis")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchEffectEvidenceSynthesis(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "EffectEvidenceSynthesis", options)
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
func (w Concrete) SearchCapabilitiesMedicinalProduct() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProduct")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProduct", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductAuthorization() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductAuthorization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductAuthorization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductAuthorization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductContraindication() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductContraindication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductContraindication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductContraindication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductIndication() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductIndication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductIndication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductIndication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductIngredient() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductIngredient")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductIngredient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductIngredient", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductInteraction() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductInteraction")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductInteraction(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductInteraction", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductManufactured() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductManufactured")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductManufactured(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductManufactured", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductPackaged() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductPackaged")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductPackaged(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductPackaged", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductPharmaceutical() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductPharmaceutical")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductPharmaceutical(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductPharmaceutical", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w Concrete) SearchCapabilitiesMedicinalProductUndesirableEffect() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("MedicinalProductUndesirableEffect")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchMedicinalProductUndesirableEffect(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "MedicinalProductUndesirableEffect", options)
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
func (w Concrete) SearchCapabilitiesRiskEvidenceSynthesis() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("RiskEvidenceSynthesis")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchRiskEvidenceSynthesis(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "RiskEvidenceSynthesis", options)
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
func (w Concrete) SearchCapabilitiesSubstanceSpecification() search.Capabilities {
	c, err := w.Generic.SearchCapabilities("SubstanceSpecification")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w Concrete) SearchSubstanceSpecification(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.Generic.Search(ctx, "SubstanceSpecification", options)
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
