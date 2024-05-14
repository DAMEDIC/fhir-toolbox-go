package capabilitiesR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	model "fhir-toolbox/model"
)

type AccountSearch interface {
	SearchCapabilitiesAccount() capabilities.SearchCapabilities
	SearchAccount(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ActivityDefinitionSearch interface {
	SearchCapabilitiesActivityDefinition() capabilities.SearchCapabilities
	SearchActivityDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type AdverseEventSearch interface {
	SearchCapabilitiesAdverseEvent() capabilities.SearchCapabilities
	SearchAdverseEvent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type AllergyIntoleranceSearch interface {
	SearchCapabilitiesAllergyIntolerance() capabilities.SearchCapabilities
	SearchAllergyIntolerance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type AppointmentSearch interface {
	SearchCapabilitiesAppointment() capabilities.SearchCapabilities
	SearchAppointment(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type AppointmentResponseSearch interface {
	SearchCapabilitiesAppointmentResponse() capabilities.SearchCapabilities
	SearchAppointmentResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type AuditEventSearch interface {
	SearchCapabilitiesAuditEvent() capabilities.SearchCapabilities
	SearchAuditEvent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type BasicSearch interface {
	SearchCapabilitiesBasic() capabilities.SearchCapabilities
	SearchBasic(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type BinarySearch interface {
	SearchCapabilitiesBinary() capabilities.SearchCapabilities
	SearchBinary(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type BiologicallyDerivedProductSearch interface {
	SearchCapabilitiesBiologicallyDerivedProduct() capabilities.SearchCapabilities
	SearchBiologicallyDerivedProduct(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type BodyStructureSearch interface {
	SearchCapabilitiesBodyStructure() capabilities.SearchCapabilities
	SearchBodyStructure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type BundleSearch interface {
	SearchCapabilitiesBundle() capabilities.SearchCapabilities
	SearchBundle(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CapabilityStatementSearch interface {
	SearchCapabilitiesCapabilityStatement() capabilities.SearchCapabilities
	SearchCapabilityStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CarePlanSearch interface {
	SearchCapabilitiesCarePlan() capabilities.SearchCapabilities
	SearchCarePlan(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CareTeamSearch interface {
	SearchCapabilitiesCareTeam() capabilities.SearchCapabilities
	SearchCareTeam(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CatalogEntrySearch interface {
	SearchCapabilitiesCatalogEntry() capabilities.SearchCapabilities
	SearchCatalogEntry(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ChargeItemSearch interface {
	SearchCapabilitiesChargeItem() capabilities.SearchCapabilities
	SearchChargeItem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ChargeItemDefinitionSearch interface {
	SearchCapabilitiesChargeItemDefinition() capabilities.SearchCapabilities
	SearchChargeItemDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ClaimSearch interface {
	SearchCapabilitiesClaim() capabilities.SearchCapabilities
	SearchClaim(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ClaimResponseSearch interface {
	SearchCapabilitiesClaimResponse() capabilities.SearchCapabilities
	SearchClaimResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ClinicalImpressionSearch interface {
	SearchCapabilitiesClinicalImpression() capabilities.SearchCapabilities
	SearchClinicalImpression(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CodeSystemSearch interface {
	SearchCapabilitiesCodeSystem() capabilities.SearchCapabilities
	SearchCodeSystem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CommunicationSearch interface {
	SearchCapabilitiesCommunication() capabilities.SearchCapabilities
	SearchCommunication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CommunicationRequestSearch interface {
	SearchCapabilitiesCommunicationRequest() capabilities.SearchCapabilities
	SearchCommunicationRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CompartmentDefinitionSearch interface {
	SearchCapabilitiesCompartmentDefinition() capabilities.SearchCapabilities
	SearchCompartmentDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CompositionSearch interface {
	SearchCapabilitiesComposition() capabilities.SearchCapabilities
	SearchComposition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ConceptMapSearch interface {
	SearchCapabilitiesConceptMap() capabilities.SearchCapabilities
	SearchConceptMap(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ConditionSearch interface {
	SearchCapabilitiesCondition() capabilities.SearchCapabilities
	SearchCondition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ConsentSearch interface {
	SearchCapabilitiesConsent() capabilities.SearchCapabilities
	SearchConsent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ContractSearch interface {
	SearchCapabilitiesContract() capabilities.SearchCapabilities
	SearchContract(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CoverageSearch interface {
	SearchCapabilitiesCoverage() capabilities.SearchCapabilities
	SearchCoverage(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CoverageEligibilityRequestSearch interface {
	SearchCapabilitiesCoverageEligibilityRequest() capabilities.SearchCapabilities
	SearchCoverageEligibilityRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type CoverageEligibilityResponseSearch interface {
	SearchCapabilitiesCoverageEligibilityResponse() capabilities.SearchCapabilities
	SearchCoverageEligibilityResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DetectedIssueSearch interface {
	SearchCapabilitiesDetectedIssue() capabilities.SearchCapabilities
	SearchDetectedIssue(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DeviceSearch interface {
	SearchCapabilitiesDevice() capabilities.SearchCapabilities
	SearchDevice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DeviceDefinitionSearch interface {
	SearchCapabilitiesDeviceDefinition() capabilities.SearchCapabilities
	SearchDeviceDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DeviceMetricSearch interface {
	SearchCapabilitiesDeviceMetric() capabilities.SearchCapabilities
	SearchDeviceMetric(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DeviceRequestSearch interface {
	SearchCapabilitiesDeviceRequest() capabilities.SearchCapabilities
	SearchDeviceRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DeviceUseStatementSearch interface {
	SearchCapabilitiesDeviceUseStatement() capabilities.SearchCapabilities
	SearchDeviceUseStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DiagnosticReportSearch interface {
	SearchCapabilitiesDiagnosticReport() capabilities.SearchCapabilities
	SearchDiagnosticReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DocumentManifestSearch interface {
	SearchCapabilitiesDocumentManifest() capabilities.SearchCapabilities
	SearchDocumentManifest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type DocumentReferenceSearch interface {
	SearchCapabilitiesDocumentReference() capabilities.SearchCapabilities
	SearchDocumentReference(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EffectEvidenceSynthesisSearch interface {
	SearchCapabilitiesEffectEvidenceSynthesis() capabilities.SearchCapabilities
	SearchEffectEvidenceSynthesis(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EncounterSearch interface {
	SearchCapabilitiesEncounter() capabilities.SearchCapabilities
	SearchEncounter(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EndpointSearch interface {
	SearchCapabilitiesEndpoint() capabilities.SearchCapabilities
	SearchEndpoint(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EnrollmentRequestSearch interface {
	SearchCapabilitiesEnrollmentRequest() capabilities.SearchCapabilities
	SearchEnrollmentRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EnrollmentResponseSearch interface {
	SearchCapabilitiesEnrollmentResponse() capabilities.SearchCapabilities
	SearchEnrollmentResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EpisodeOfCareSearch interface {
	SearchCapabilitiesEpisodeOfCare() capabilities.SearchCapabilities
	SearchEpisodeOfCare(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EventDefinitionSearch interface {
	SearchCapabilitiesEventDefinition() capabilities.SearchCapabilities
	SearchEventDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EvidenceSearch interface {
	SearchCapabilitiesEvidence() capabilities.SearchCapabilities
	SearchEvidence(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type EvidenceVariableSearch interface {
	SearchCapabilitiesEvidenceVariable() capabilities.SearchCapabilities
	SearchEvidenceVariable(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ExampleScenarioSearch interface {
	SearchCapabilitiesExampleScenario() capabilities.SearchCapabilities
	SearchExampleScenario(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ExplanationOfBenefitSearch interface {
	SearchCapabilitiesExplanationOfBenefit() capabilities.SearchCapabilities
	SearchExplanationOfBenefit(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type FamilyMemberHistorySearch interface {
	SearchCapabilitiesFamilyMemberHistory() capabilities.SearchCapabilities
	SearchFamilyMemberHistory(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type FlagSearch interface {
	SearchCapabilitiesFlag() capabilities.SearchCapabilities
	SearchFlag(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type GoalSearch interface {
	SearchCapabilitiesGoal() capabilities.SearchCapabilities
	SearchGoal(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type GraphDefinitionSearch interface {
	SearchCapabilitiesGraphDefinition() capabilities.SearchCapabilities
	SearchGraphDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type GroupSearch interface {
	SearchCapabilitiesGroup() capabilities.SearchCapabilities
	SearchGroup(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type GuidanceResponseSearch interface {
	SearchCapabilitiesGuidanceResponse() capabilities.SearchCapabilities
	SearchGuidanceResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type HealthcareServiceSearch interface {
	SearchCapabilitiesHealthcareService() capabilities.SearchCapabilities
	SearchHealthcareService(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ImagingStudySearch interface {
	SearchCapabilitiesImagingStudy() capabilities.SearchCapabilities
	SearchImagingStudy(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationSearch interface {
	SearchCapabilitiesImmunization() capabilities.SearchCapabilities
	SearchImmunization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationEvaluationSearch interface {
	SearchCapabilitiesImmunizationEvaluation() capabilities.SearchCapabilities
	SearchImmunizationEvaluation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationRecommendationSearch interface {
	SearchCapabilitiesImmunizationRecommendation() capabilities.SearchCapabilities
	SearchImmunizationRecommendation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ImplementationGuideSearch interface {
	SearchCapabilitiesImplementationGuide() capabilities.SearchCapabilities
	SearchImplementationGuide(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type InsurancePlanSearch interface {
	SearchCapabilitiesInsurancePlan() capabilities.SearchCapabilities
	SearchInsurancePlan(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type InvoiceSearch interface {
	SearchCapabilitiesInvoice() capabilities.SearchCapabilities
	SearchInvoice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type LibrarySearch interface {
	SearchCapabilitiesLibrary() capabilities.SearchCapabilities
	SearchLibrary(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type LinkageSearch interface {
	SearchCapabilitiesLinkage() capabilities.SearchCapabilities
	SearchLinkage(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ListSearch interface {
	SearchCapabilitiesList() capabilities.SearchCapabilities
	SearchList(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type LocationSearch interface {
	SearchCapabilitiesLocation() capabilities.SearchCapabilities
	SearchLocation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MeasureSearch interface {
	SearchCapabilitiesMeasure() capabilities.SearchCapabilities
	SearchMeasure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MeasureReportSearch interface {
	SearchCapabilitiesMeasureReport() capabilities.SearchCapabilities
	SearchMeasureReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MediaSearch interface {
	SearchCapabilitiesMedia() capabilities.SearchCapabilities
	SearchMedia(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationSearch interface {
	SearchCapabilitiesMedication() capabilities.SearchCapabilities
	SearchMedication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationAdministrationSearch interface {
	SearchCapabilitiesMedicationAdministration() capabilities.SearchCapabilities
	SearchMedicationAdministration(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationDispenseSearch interface {
	SearchCapabilitiesMedicationDispense() capabilities.SearchCapabilities
	SearchMedicationDispense(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationKnowledgeSearch interface {
	SearchCapabilitiesMedicationKnowledge() capabilities.SearchCapabilities
	SearchMedicationKnowledge(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationRequestSearch interface {
	SearchCapabilitiesMedicationRequest() capabilities.SearchCapabilities
	SearchMedicationRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicationStatementSearch interface {
	SearchCapabilitiesMedicationStatement() capabilities.SearchCapabilities
	SearchMedicationStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductSearch interface {
	SearchCapabilitiesMedicinalProduct() capabilities.SearchCapabilities
	SearchMedicinalProduct(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductAuthorizationSearch interface {
	SearchCapabilitiesMedicinalProductAuthorization() capabilities.SearchCapabilities
	SearchMedicinalProductAuthorization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductContraindicationSearch interface {
	SearchCapabilitiesMedicinalProductContraindication() capabilities.SearchCapabilities
	SearchMedicinalProductContraindication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductIndicationSearch interface {
	SearchCapabilitiesMedicinalProductIndication() capabilities.SearchCapabilities
	SearchMedicinalProductIndication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductIngredientSearch interface {
	SearchCapabilitiesMedicinalProductIngredient() capabilities.SearchCapabilities
	SearchMedicinalProductIngredient(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductInteractionSearch interface {
	SearchCapabilitiesMedicinalProductInteraction() capabilities.SearchCapabilities
	SearchMedicinalProductInteraction(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductManufacturedSearch interface {
	SearchCapabilitiesMedicinalProductManufactured() capabilities.SearchCapabilities
	SearchMedicinalProductManufactured(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductPackagedSearch interface {
	SearchCapabilitiesMedicinalProductPackaged() capabilities.SearchCapabilities
	SearchMedicinalProductPackaged(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductPharmaceuticalSearch interface {
	SearchCapabilitiesMedicinalProductPharmaceutical() capabilities.SearchCapabilities
	SearchMedicinalProductPharmaceutical(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductUndesirableEffectSearch interface {
	SearchCapabilitiesMedicinalProductUndesirableEffect() capabilities.SearchCapabilities
	SearchMedicinalProductUndesirableEffect(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MessageDefinitionSearch interface {
	SearchCapabilitiesMessageDefinition() capabilities.SearchCapabilities
	SearchMessageDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MessageHeaderSearch interface {
	SearchCapabilitiesMessageHeader() capabilities.SearchCapabilities
	SearchMessageHeader(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type MolecularSequenceSearch interface {
	SearchCapabilitiesMolecularSequence() capabilities.SearchCapabilities
	SearchMolecularSequence(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type NamingSystemSearch interface {
	SearchCapabilitiesNamingSystem() capabilities.SearchCapabilities
	SearchNamingSystem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type NutritionOrderSearch interface {
	SearchCapabilitiesNutritionOrder() capabilities.SearchCapabilities
	SearchNutritionOrder(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ObservationSearch interface {
	SearchCapabilitiesObservation() capabilities.SearchCapabilities
	SearchObservation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ObservationDefinitionSearch interface {
	SearchCapabilitiesObservationDefinition() capabilities.SearchCapabilities
	SearchObservationDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type OperationDefinitionSearch interface {
	SearchCapabilitiesOperationDefinition() capabilities.SearchCapabilities
	SearchOperationDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type OperationOutcomeSearch interface {
	SearchCapabilitiesOperationOutcome() capabilities.SearchCapabilities
	SearchOperationOutcome(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type OrganizationSearch interface {
	SearchCapabilitiesOrganization() capabilities.SearchCapabilities
	SearchOrganization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type OrganizationAffiliationSearch interface {
	SearchCapabilitiesOrganizationAffiliation() capabilities.SearchCapabilities
	SearchOrganizationAffiliation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ParametersSearch interface {
	SearchCapabilitiesParameters() capabilities.SearchCapabilities
	SearchParameters(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PatientSearch interface {
	SearchCapabilitiesPatient() capabilities.SearchCapabilities
	SearchPatient(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PaymentNoticeSearch interface {
	SearchCapabilitiesPaymentNotice() capabilities.SearchCapabilities
	SearchPaymentNotice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PaymentReconciliationSearch interface {
	SearchCapabilitiesPaymentReconciliation() capabilities.SearchCapabilities
	SearchPaymentReconciliation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PersonSearch interface {
	SearchCapabilitiesPerson() capabilities.SearchCapabilities
	SearchPerson(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PlanDefinitionSearch interface {
	SearchCapabilitiesPlanDefinition() capabilities.SearchCapabilities
	SearchPlanDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PractitionerSearch interface {
	SearchCapabilitiesPractitioner() capabilities.SearchCapabilities
	SearchPractitioner(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type PractitionerRoleSearch interface {
	SearchCapabilitiesPractitionerRole() capabilities.SearchCapabilities
	SearchPractitionerRole(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ProcedureSearch interface {
	SearchCapabilitiesProcedure() capabilities.SearchCapabilities
	SearchProcedure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ProvenanceSearch interface {
	SearchCapabilitiesProvenance() capabilities.SearchCapabilities
	SearchProvenance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type QuestionnaireSearch interface {
	SearchCapabilitiesQuestionnaire() capabilities.SearchCapabilities
	SearchQuestionnaire(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type QuestionnaireResponseSearch interface {
	SearchCapabilitiesQuestionnaireResponse() capabilities.SearchCapabilities
	SearchQuestionnaireResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type RelatedPersonSearch interface {
	SearchCapabilitiesRelatedPerson() capabilities.SearchCapabilities
	SearchRelatedPerson(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type RequestGroupSearch interface {
	SearchCapabilitiesRequestGroup() capabilities.SearchCapabilities
	SearchRequestGroup(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ResearchDefinitionSearch interface {
	SearchCapabilitiesResearchDefinition() capabilities.SearchCapabilities
	SearchResearchDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ResearchElementDefinitionSearch interface {
	SearchCapabilitiesResearchElementDefinition() capabilities.SearchCapabilities
	SearchResearchElementDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ResearchStudySearch interface {
	SearchCapabilitiesResearchStudy() capabilities.SearchCapabilities
	SearchResearchStudy(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ResearchSubjectSearch interface {
	SearchCapabilitiesResearchSubject() capabilities.SearchCapabilities
	SearchResearchSubject(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type RiskAssessmentSearch interface {
	SearchCapabilitiesRiskAssessment() capabilities.SearchCapabilities
	SearchRiskAssessment(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type RiskEvidenceSynthesisSearch interface {
	SearchCapabilitiesRiskEvidenceSynthesis() capabilities.SearchCapabilities
	SearchRiskEvidenceSynthesis(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ScheduleSearch interface {
	SearchCapabilitiesSchedule() capabilities.SearchCapabilities
	SearchSchedule(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SearchParameterSearch interface {
	SearchCapabilitiesSearchParameter() capabilities.SearchCapabilities
	SearchSearchParameter(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ServiceRequestSearch interface {
	SearchCapabilitiesServiceRequest() capabilities.SearchCapabilities
	SearchServiceRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SlotSearch interface {
	SearchCapabilitiesSlot() capabilities.SearchCapabilities
	SearchSlot(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SpecimenSearch interface {
	SearchCapabilitiesSpecimen() capabilities.SearchCapabilities
	SearchSpecimen(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SpecimenDefinitionSearch interface {
	SearchCapabilitiesSpecimenDefinition() capabilities.SearchCapabilities
	SearchSpecimenDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type StructureDefinitionSearch interface {
	SearchCapabilitiesStructureDefinition() capabilities.SearchCapabilities
	SearchStructureDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type StructureMapSearch interface {
	SearchCapabilitiesStructureMap() capabilities.SearchCapabilities
	SearchStructureMap(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubscriptionSearch interface {
	SearchCapabilitiesSubscription() capabilities.SearchCapabilities
	SearchSubscription(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSearch interface {
	SearchCapabilitiesSubstance() capabilities.SearchCapabilities
	SearchSubstance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceNucleicAcidSearch interface {
	SearchCapabilitiesSubstanceNucleicAcid() capabilities.SearchCapabilities
	SearchSubstanceNucleicAcid(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstancePolymerSearch interface {
	SearchCapabilitiesSubstancePolymer() capabilities.SearchCapabilities
	SearchSubstancePolymer(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceProteinSearch interface {
	SearchCapabilitiesSubstanceProtein() capabilities.SearchCapabilities
	SearchSubstanceProtein(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceReferenceInformationSearch interface {
	SearchCapabilitiesSubstanceReferenceInformation() capabilities.SearchCapabilities
	SearchSubstanceReferenceInformation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSourceMaterialSearch interface {
	SearchCapabilitiesSubstanceSourceMaterial() capabilities.SearchCapabilities
	SearchSubstanceSourceMaterial(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSpecificationSearch interface {
	SearchCapabilitiesSubstanceSpecification() capabilities.SearchCapabilities
	SearchSubstanceSpecification(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SupplyDeliverySearch interface {
	SearchCapabilitiesSupplyDelivery() capabilities.SearchCapabilities
	SearchSupplyDelivery(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type SupplyRequestSearch interface {
	SearchCapabilitiesSupplyRequest() capabilities.SearchCapabilities
	SearchSupplyRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type TaskSearch interface {
	SearchCapabilitiesTask() capabilities.SearchCapabilities
	SearchTask(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type TerminologyCapabilitiesSearch interface {
	SearchCapabilitiesTerminologyCapabilities() capabilities.SearchCapabilities
	SearchTerminologyCapabilities(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type TestReportSearch interface {
	SearchCapabilitiesTestReport() capabilities.SearchCapabilities
	SearchTestReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type TestScriptSearch interface {
	SearchCapabilitiesTestScript() capabilities.SearchCapabilities
	SearchTestScript(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type ValueSetSearch interface {
	SearchCapabilitiesValueSet() capabilities.SearchCapabilities
	SearchValueSet(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type VerificationResultSearch interface {
	SearchCapabilitiesVerificationResult() capabilities.SearchCapabilities
	SearchVerificationResult(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
type VisionPrescriptionSearch interface {
	SearchCapabilitiesVisionPrescription() capabilities.SearchCapabilities
	SearchVisionPrescription(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError)
}
