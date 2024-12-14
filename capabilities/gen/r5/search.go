package capabilitiesR5

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	search "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
)

type AccountSearch interface {
	SearchCapabilitiesAccount() search.Capabilities
	SearchAccount(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ActivityDefinitionSearch interface {
	SearchCapabilitiesActivityDefinition() search.Capabilities
	SearchActivityDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ActorDefinitionSearch interface {
	SearchCapabilitiesActorDefinition() search.Capabilities
	SearchActorDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AdministrableProductDefinitionSearch interface {
	SearchCapabilitiesAdministrableProductDefinition() search.Capabilities
	SearchAdministrableProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AdverseEventSearch interface {
	SearchCapabilitiesAdverseEvent() search.Capabilities
	SearchAdverseEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AllergyIntoleranceSearch interface {
	SearchCapabilitiesAllergyIntolerance() search.Capabilities
	SearchAllergyIntolerance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AppointmentSearch interface {
	SearchCapabilitiesAppointment() search.Capabilities
	SearchAppointment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AppointmentResponseSearch interface {
	SearchCapabilitiesAppointmentResponse() search.Capabilities
	SearchAppointmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ArtifactAssessmentSearch interface {
	SearchCapabilitiesArtifactAssessment() search.Capabilities
	SearchArtifactAssessment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type AuditEventSearch interface {
	SearchCapabilitiesAuditEvent() search.Capabilities
	SearchAuditEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BasicSearch interface {
	SearchCapabilitiesBasic() search.Capabilities
	SearchBasic(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BinarySearch interface {
	SearchCapabilitiesBinary() search.Capabilities
	SearchBinary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BiologicallyDerivedProductSearch interface {
	SearchCapabilitiesBiologicallyDerivedProduct() search.Capabilities
	SearchBiologicallyDerivedProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BiologicallyDerivedProductDispenseSearch interface {
	SearchCapabilitiesBiologicallyDerivedProductDispense() search.Capabilities
	SearchBiologicallyDerivedProductDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BodyStructureSearch interface {
	SearchCapabilitiesBodyStructure() search.Capabilities
	SearchBodyStructure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type BundleSearch interface {
	SearchCapabilitiesBundle() search.Capabilities
	SearchBundle(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CapabilityStatementSearch interface {
	SearchCapabilitiesCapabilityStatement() search.Capabilities
	SearchCapabilityStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CarePlanSearch interface {
	SearchCapabilitiesCarePlan() search.Capabilities
	SearchCarePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CareTeamSearch interface {
	SearchCapabilitiesCareTeam() search.Capabilities
	SearchCareTeam(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ChargeItemSearch interface {
	SearchCapabilitiesChargeItem() search.Capabilities
	SearchChargeItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ChargeItemDefinitionSearch interface {
	SearchCapabilitiesChargeItemDefinition() search.Capabilities
	SearchChargeItemDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CitationSearch interface {
	SearchCapabilitiesCitation() search.Capabilities
	SearchCitation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ClaimSearch interface {
	SearchCapabilitiesClaim() search.Capabilities
	SearchClaim(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ClaimResponseSearch interface {
	SearchCapabilitiesClaimResponse() search.Capabilities
	SearchClaimResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ClinicalImpressionSearch interface {
	SearchCapabilitiesClinicalImpression() search.Capabilities
	SearchClinicalImpression(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ClinicalUseDefinitionSearch interface {
	SearchCapabilitiesClinicalUseDefinition() search.Capabilities
	SearchClinicalUseDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CodeSystemSearch interface {
	SearchCapabilitiesCodeSystem() search.Capabilities
	SearchCodeSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CommunicationSearch interface {
	SearchCapabilitiesCommunication() search.Capabilities
	SearchCommunication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CommunicationRequestSearch interface {
	SearchCapabilitiesCommunicationRequest() search.Capabilities
	SearchCommunicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CompartmentDefinitionSearch interface {
	SearchCapabilitiesCompartmentDefinition() search.Capabilities
	SearchCompartmentDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CompositionSearch interface {
	SearchCapabilitiesComposition() search.Capabilities
	SearchComposition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ConceptMapSearch interface {
	SearchCapabilitiesConceptMap() search.Capabilities
	SearchConceptMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ConditionSearch interface {
	SearchCapabilitiesCondition() search.Capabilities
	SearchCondition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ConditionDefinitionSearch interface {
	SearchCapabilitiesConditionDefinition() search.Capabilities
	SearchConditionDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ConsentSearch interface {
	SearchCapabilitiesConsent() search.Capabilities
	SearchConsent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ContractSearch interface {
	SearchCapabilitiesContract() search.Capabilities
	SearchContract(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CoverageSearch interface {
	SearchCapabilitiesCoverage() search.Capabilities
	SearchCoverage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CoverageEligibilityRequestSearch interface {
	SearchCapabilitiesCoverageEligibilityRequest() search.Capabilities
	SearchCoverageEligibilityRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type CoverageEligibilityResponseSearch interface {
	SearchCapabilitiesCoverageEligibilityResponse() search.Capabilities
	SearchCoverageEligibilityResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DetectedIssueSearch interface {
	SearchCapabilitiesDetectedIssue() search.Capabilities
	SearchDetectedIssue(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceSearch interface {
	SearchCapabilitiesDevice() search.Capabilities
	SearchDevice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceAssociationSearch interface {
	SearchCapabilitiesDeviceAssociation() search.Capabilities
	SearchDeviceAssociation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceDefinitionSearch interface {
	SearchCapabilitiesDeviceDefinition() search.Capabilities
	SearchDeviceDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceDispenseSearch interface {
	SearchCapabilitiesDeviceDispense() search.Capabilities
	SearchDeviceDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceMetricSearch interface {
	SearchCapabilitiesDeviceMetric() search.Capabilities
	SearchDeviceMetric(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceRequestSearch interface {
	SearchCapabilitiesDeviceRequest() search.Capabilities
	SearchDeviceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DeviceUsageSearch interface {
	SearchCapabilitiesDeviceUsage() search.Capabilities
	SearchDeviceUsage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DiagnosticReportSearch interface {
	SearchCapabilitiesDiagnosticReport() search.Capabilities
	SearchDiagnosticReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type DocumentReferenceSearch interface {
	SearchCapabilitiesDocumentReference() search.Capabilities
	SearchDocumentReference(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EncounterSearch interface {
	SearchCapabilitiesEncounter() search.Capabilities
	SearchEncounter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EncounterHistorySearch interface {
	SearchCapabilitiesEncounterHistory() search.Capabilities
	SearchEncounterHistory(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EndpointSearch interface {
	SearchCapabilitiesEndpoint() search.Capabilities
	SearchEndpoint(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EnrollmentRequestSearch interface {
	SearchCapabilitiesEnrollmentRequest() search.Capabilities
	SearchEnrollmentRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EnrollmentResponseSearch interface {
	SearchCapabilitiesEnrollmentResponse() search.Capabilities
	SearchEnrollmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EpisodeOfCareSearch interface {
	SearchCapabilitiesEpisodeOfCare() search.Capabilities
	SearchEpisodeOfCare(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EventDefinitionSearch interface {
	SearchCapabilitiesEventDefinition() search.Capabilities
	SearchEventDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EvidenceSearch interface {
	SearchCapabilitiesEvidence() search.Capabilities
	SearchEvidence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EvidenceReportSearch interface {
	SearchCapabilitiesEvidenceReport() search.Capabilities
	SearchEvidenceReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type EvidenceVariableSearch interface {
	SearchCapabilitiesEvidenceVariable() search.Capabilities
	SearchEvidenceVariable(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ExampleScenarioSearch interface {
	SearchCapabilitiesExampleScenario() search.Capabilities
	SearchExampleScenario(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ExplanationOfBenefitSearch interface {
	SearchCapabilitiesExplanationOfBenefit() search.Capabilities
	SearchExplanationOfBenefit(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type FamilyMemberHistorySearch interface {
	SearchCapabilitiesFamilyMemberHistory() search.Capabilities
	SearchFamilyMemberHistory(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type FlagSearch interface {
	SearchCapabilitiesFlag() search.Capabilities
	SearchFlag(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type FormularyItemSearch interface {
	SearchCapabilitiesFormularyItem() search.Capabilities
	SearchFormularyItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type GenomicStudySearch interface {
	SearchCapabilitiesGenomicStudy() search.Capabilities
	SearchGenomicStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type GoalSearch interface {
	SearchCapabilitiesGoal() search.Capabilities
	SearchGoal(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type GraphDefinitionSearch interface {
	SearchCapabilitiesGraphDefinition() search.Capabilities
	SearchGraphDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type GroupSearch interface {
	SearchCapabilitiesGroup() search.Capabilities
	SearchGroup(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type GuidanceResponseSearch interface {
	SearchCapabilitiesGuidanceResponse() search.Capabilities
	SearchGuidanceResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type HealthcareServiceSearch interface {
	SearchCapabilitiesHealthcareService() search.Capabilities
	SearchHealthcareService(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImagingSelectionSearch interface {
	SearchCapabilitiesImagingSelection() search.Capabilities
	SearchImagingSelection(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImagingStudySearch interface {
	SearchCapabilitiesImagingStudy() search.Capabilities
	SearchImagingStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImmunizationSearch interface {
	SearchCapabilitiesImmunization() search.Capabilities
	SearchImmunization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImmunizationEvaluationSearch interface {
	SearchCapabilitiesImmunizationEvaluation() search.Capabilities
	SearchImmunizationEvaluation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImmunizationRecommendationSearch interface {
	SearchCapabilitiesImmunizationRecommendation() search.Capabilities
	SearchImmunizationRecommendation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ImplementationGuideSearch interface {
	SearchCapabilitiesImplementationGuide() search.Capabilities
	SearchImplementationGuide(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type IngredientSearch interface {
	SearchCapabilitiesIngredient() search.Capabilities
	SearchIngredient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type InsurancePlanSearch interface {
	SearchCapabilitiesInsurancePlan() search.Capabilities
	SearchInsurancePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type InventoryItemSearch interface {
	SearchCapabilitiesInventoryItem() search.Capabilities
	SearchInventoryItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type InventoryReportSearch interface {
	SearchCapabilitiesInventoryReport() search.Capabilities
	SearchInventoryReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type InvoiceSearch interface {
	SearchCapabilitiesInvoice() search.Capabilities
	SearchInvoice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type LibrarySearch interface {
	SearchCapabilitiesLibrary() search.Capabilities
	SearchLibrary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type LinkageSearch interface {
	SearchCapabilitiesLinkage() search.Capabilities
	SearchLinkage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ListSearch interface {
	SearchCapabilitiesList() search.Capabilities
	SearchList(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type LocationSearch interface {
	SearchCapabilitiesLocation() search.Capabilities
	SearchLocation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ManufacturedItemDefinitionSearch interface {
	SearchCapabilitiesManufacturedItemDefinition() search.Capabilities
	SearchManufacturedItemDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MeasureSearch interface {
	SearchCapabilitiesMeasure() search.Capabilities
	SearchMeasure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MeasureReportSearch interface {
	SearchCapabilitiesMeasureReport() search.Capabilities
	SearchMeasureReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationSearch interface {
	SearchCapabilitiesMedication() search.Capabilities
	SearchMedication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationAdministrationSearch interface {
	SearchCapabilitiesMedicationAdministration() search.Capabilities
	SearchMedicationAdministration(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationDispenseSearch interface {
	SearchCapabilitiesMedicationDispense() search.Capabilities
	SearchMedicationDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationKnowledgeSearch interface {
	SearchCapabilitiesMedicationKnowledge() search.Capabilities
	SearchMedicationKnowledge(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationRequestSearch interface {
	SearchCapabilitiesMedicationRequest() search.Capabilities
	SearchMedicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicationStatementSearch interface {
	SearchCapabilitiesMedicationStatement() search.Capabilities
	SearchMedicationStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MedicinalProductDefinitionSearch interface {
	SearchCapabilitiesMedicinalProductDefinition() search.Capabilities
	SearchMedicinalProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MessageDefinitionSearch interface {
	SearchCapabilitiesMessageDefinition() search.Capabilities
	SearchMessageDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MessageHeaderSearch interface {
	SearchCapabilitiesMessageHeader() search.Capabilities
	SearchMessageHeader(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type MolecularSequenceSearch interface {
	SearchCapabilitiesMolecularSequence() search.Capabilities
	SearchMolecularSequence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type NamingSystemSearch interface {
	SearchCapabilitiesNamingSystem() search.Capabilities
	SearchNamingSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type NutritionIntakeSearch interface {
	SearchCapabilitiesNutritionIntake() search.Capabilities
	SearchNutritionIntake(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type NutritionOrderSearch interface {
	SearchCapabilitiesNutritionOrder() search.Capabilities
	SearchNutritionOrder(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type NutritionProductSearch interface {
	SearchCapabilitiesNutritionProduct() search.Capabilities
	SearchNutritionProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ObservationSearch interface {
	SearchCapabilitiesObservation() search.Capabilities
	SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ObservationDefinitionSearch interface {
	SearchCapabilitiesObservationDefinition() search.Capabilities
	SearchObservationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type OperationDefinitionSearch interface {
	SearchCapabilitiesOperationDefinition() search.Capabilities
	SearchOperationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type OperationOutcomeSearch interface {
	SearchCapabilitiesOperationOutcome() search.Capabilities
	SearchOperationOutcome(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type OrganizationSearch interface {
	SearchCapabilitiesOrganization() search.Capabilities
	SearchOrganization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type OrganizationAffiliationSearch interface {
	SearchCapabilitiesOrganizationAffiliation() search.Capabilities
	SearchOrganizationAffiliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PackagedProductDefinitionSearch interface {
	SearchCapabilitiesPackagedProductDefinition() search.Capabilities
	SearchPackagedProductDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ParametersSearch interface {
	SearchCapabilitiesParameters() search.Capabilities
	SearchParameters(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PatientSearch interface {
	SearchCapabilitiesPatient() search.Capabilities
	SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PaymentNoticeSearch interface {
	SearchCapabilitiesPaymentNotice() search.Capabilities
	SearchPaymentNotice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PaymentReconciliationSearch interface {
	SearchCapabilitiesPaymentReconciliation() search.Capabilities
	SearchPaymentReconciliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PermissionSearch interface {
	SearchCapabilitiesPermission() search.Capabilities
	SearchPermission(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PersonSearch interface {
	SearchCapabilitiesPerson() search.Capabilities
	SearchPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PlanDefinitionSearch interface {
	SearchCapabilitiesPlanDefinition() search.Capabilities
	SearchPlanDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PractitionerSearch interface {
	SearchCapabilitiesPractitioner() search.Capabilities
	SearchPractitioner(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type PractitionerRoleSearch interface {
	SearchCapabilitiesPractitionerRole() search.Capabilities
	SearchPractitionerRole(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ProcedureSearch interface {
	SearchCapabilitiesProcedure() search.Capabilities
	SearchProcedure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ProvenanceSearch interface {
	SearchCapabilitiesProvenance() search.Capabilities
	SearchProvenance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type QuestionnaireSearch interface {
	SearchCapabilitiesQuestionnaire() search.Capabilities
	SearchQuestionnaire(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type QuestionnaireResponseSearch interface {
	SearchCapabilitiesQuestionnaireResponse() search.Capabilities
	SearchQuestionnaireResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type RegulatedAuthorizationSearch interface {
	SearchCapabilitiesRegulatedAuthorization() search.Capabilities
	SearchRegulatedAuthorization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type RelatedPersonSearch interface {
	SearchCapabilitiesRelatedPerson() search.Capabilities
	SearchRelatedPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type RequestOrchestrationSearch interface {
	SearchCapabilitiesRequestOrchestration() search.Capabilities
	SearchRequestOrchestration(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type RequirementsSearch interface {
	SearchCapabilitiesRequirements() search.Capabilities
	SearchRequirements(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ResearchStudySearch interface {
	SearchCapabilitiesResearchStudy() search.Capabilities
	SearchResearchStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ResearchSubjectSearch interface {
	SearchCapabilitiesResearchSubject() search.Capabilities
	SearchResearchSubject(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type RiskAssessmentSearch interface {
	SearchCapabilitiesRiskAssessment() search.Capabilities
	SearchRiskAssessment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ScheduleSearch interface {
	SearchCapabilitiesSchedule() search.Capabilities
	SearchSchedule(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SearchParameterSearch interface {
	SearchCapabilitiesSearchParameter() search.Capabilities
	SearchSearchParameter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ServiceRequestSearch interface {
	SearchCapabilitiesServiceRequest() search.Capabilities
	SearchServiceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SlotSearch interface {
	SearchCapabilitiesSlot() search.Capabilities
	SearchSlot(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SpecimenSearch interface {
	SearchCapabilitiesSpecimen() search.Capabilities
	SearchSpecimen(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SpecimenDefinitionSearch interface {
	SearchCapabilitiesSpecimenDefinition() search.Capabilities
	SearchSpecimenDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type StructureDefinitionSearch interface {
	SearchCapabilitiesStructureDefinition() search.Capabilities
	SearchStructureDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type StructureMapSearch interface {
	SearchCapabilitiesStructureMap() search.Capabilities
	SearchStructureMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubscriptionSearch interface {
	SearchCapabilitiesSubscription() search.Capabilities
	SearchSubscription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubscriptionStatusSearch interface {
	SearchCapabilitiesSubscriptionStatus() search.Capabilities
	SearchSubscriptionStatus(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubscriptionTopicSearch interface {
	SearchCapabilitiesSubscriptionTopic() search.Capabilities
	SearchSubscriptionTopic(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceSearch interface {
	SearchCapabilitiesSubstance() search.Capabilities
	SearchSubstance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceDefinitionSearch interface {
	SearchCapabilitiesSubstanceDefinition() search.Capabilities
	SearchSubstanceDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceNucleicAcidSearch interface {
	SearchCapabilitiesSubstanceNucleicAcid() search.Capabilities
	SearchSubstanceNucleicAcid(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstancePolymerSearch interface {
	SearchCapabilitiesSubstancePolymer() search.Capabilities
	SearchSubstancePolymer(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceProteinSearch interface {
	SearchCapabilitiesSubstanceProtein() search.Capabilities
	SearchSubstanceProtein(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceReferenceInformationSearch interface {
	SearchCapabilitiesSubstanceReferenceInformation() search.Capabilities
	SearchSubstanceReferenceInformation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SubstanceSourceMaterialSearch interface {
	SearchCapabilitiesSubstanceSourceMaterial() search.Capabilities
	SearchSubstanceSourceMaterial(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SupplyDeliverySearch interface {
	SearchCapabilitiesSupplyDelivery() search.Capabilities
	SearchSupplyDelivery(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type SupplyRequestSearch interface {
	SearchCapabilitiesSupplyRequest() search.Capabilities
	SearchSupplyRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TaskSearch interface {
	SearchCapabilitiesTask() search.Capabilities
	SearchTask(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TerminologyCapabilitiesSearch interface {
	SearchCapabilitiesTerminologyCapabilities() search.Capabilities
	SearchTerminologyCapabilities(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TestPlanSearch interface {
	SearchCapabilitiesTestPlan() search.Capabilities
	SearchTestPlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TestReportSearch interface {
	SearchCapabilitiesTestReport() search.Capabilities
	SearchTestReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TestScriptSearch interface {
	SearchCapabilitiesTestScript() search.Capabilities
	SearchTestScript(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type TransportSearch interface {
	SearchCapabilitiesTransport() search.Capabilities
	SearchTransport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type ValueSetSearch interface {
	SearchCapabilitiesValueSet() search.Capabilities
	SearchValueSet(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type VerificationResultSearch interface {
	SearchCapabilitiesVerificationResult() search.Capabilities
	SearchVerificationResult(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
type VisionPrescriptionSearch interface {
	SearchCapabilitiesVisionPrescription() search.Capabilities
	SearchVisionPrescription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError)
}
