package capabilitiesR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	model "fhir-toolbox/model"
)

type AccountSearch interface {
	SearchCapabilitiesAccount() capabilities.SearchCapabilities
	SearchAccount(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ActivityDefinitionSearch interface {
	SearchCapabilitiesActivityDefinition() capabilities.SearchCapabilities
	SearchActivityDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type AdverseEventSearch interface {
	SearchCapabilitiesAdverseEvent() capabilities.SearchCapabilities
	SearchAdverseEvent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type AllergyIntoleranceSearch interface {
	SearchCapabilitiesAllergyIntolerance() capabilities.SearchCapabilities
	SearchAllergyIntolerance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type AppointmentSearch interface {
	SearchCapabilitiesAppointment() capabilities.SearchCapabilities
	SearchAppointment(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type AppointmentResponseSearch interface {
	SearchCapabilitiesAppointmentResponse() capabilities.SearchCapabilities
	SearchAppointmentResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type AuditEventSearch interface {
	SearchCapabilitiesAuditEvent() capabilities.SearchCapabilities
	SearchAuditEvent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type BasicSearch interface {
	SearchCapabilitiesBasic() capabilities.SearchCapabilities
	SearchBasic(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type BinarySearch interface {
	SearchCapabilitiesBinary() capabilities.SearchCapabilities
	SearchBinary(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type BiologicallyDerivedProductSearch interface {
	SearchCapabilitiesBiologicallyDerivedProduct() capabilities.SearchCapabilities
	SearchBiologicallyDerivedProduct(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type BodyStructureSearch interface {
	SearchCapabilitiesBodyStructure() capabilities.SearchCapabilities
	SearchBodyStructure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type BundleSearch interface {
	SearchCapabilitiesBundle() capabilities.SearchCapabilities
	SearchBundle(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CapabilityStatementSearch interface {
	SearchCapabilitiesCapabilityStatement() capabilities.SearchCapabilities
	SearchCapabilityStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CarePlanSearch interface {
	SearchCapabilitiesCarePlan() capabilities.SearchCapabilities
	SearchCarePlan(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CareTeamSearch interface {
	SearchCapabilitiesCareTeam() capabilities.SearchCapabilities
	SearchCareTeam(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CatalogEntrySearch interface {
	SearchCapabilitiesCatalogEntry() capabilities.SearchCapabilities
	SearchCatalogEntry(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ChargeItemSearch interface {
	SearchCapabilitiesChargeItem() capabilities.SearchCapabilities
	SearchChargeItem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ChargeItemDefinitionSearch interface {
	SearchCapabilitiesChargeItemDefinition() capabilities.SearchCapabilities
	SearchChargeItemDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ClaimSearch interface {
	SearchCapabilitiesClaim() capabilities.SearchCapabilities
	SearchClaim(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ClaimResponseSearch interface {
	SearchCapabilitiesClaimResponse() capabilities.SearchCapabilities
	SearchClaimResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ClinicalImpressionSearch interface {
	SearchCapabilitiesClinicalImpression() capabilities.SearchCapabilities
	SearchClinicalImpression(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CodeSystemSearch interface {
	SearchCapabilitiesCodeSystem() capabilities.SearchCapabilities
	SearchCodeSystem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CommunicationSearch interface {
	SearchCapabilitiesCommunication() capabilities.SearchCapabilities
	SearchCommunication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CommunicationRequestSearch interface {
	SearchCapabilitiesCommunicationRequest() capabilities.SearchCapabilities
	SearchCommunicationRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CompartmentDefinitionSearch interface {
	SearchCapabilitiesCompartmentDefinition() capabilities.SearchCapabilities
	SearchCompartmentDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CompositionSearch interface {
	SearchCapabilitiesComposition() capabilities.SearchCapabilities
	SearchComposition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ConceptMapSearch interface {
	SearchCapabilitiesConceptMap() capabilities.SearchCapabilities
	SearchConceptMap(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ConditionSearch interface {
	SearchCapabilitiesCondition() capabilities.SearchCapabilities
	SearchCondition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ConsentSearch interface {
	SearchCapabilitiesConsent() capabilities.SearchCapabilities
	SearchConsent(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ContractSearch interface {
	SearchCapabilitiesContract() capabilities.SearchCapabilities
	SearchContract(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CoverageSearch interface {
	SearchCapabilitiesCoverage() capabilities.SearchCapabilities
	SearchCoverage(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CoverageEligibilityRequestSearch interface {
	SearchCapabilitiesCoverageEligibilityRequest() capabilities.SearchCapabilities
	SearchCoverageEligibilityRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type CoverageEligibilityResponseSearch interface {
	SearchCapabilitiesCoverageEligibilityResponse() capabilities.SearchCapabilities
	SearchCoverageEligibilityResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DetectedIssueSearch interface {
	SearchCapabilitiesDetectedIssue() capabilities.SearchCapabilities
	SearchDetectedIssue(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DeviceSearch interface {
	SearchCapabilitiesDevice() capabilities.SearchCapabilities
	SearchDevice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DeviceDefinitionSearch interface {
	SearchCapabilitiesDeviceDefinition() capabilities.SearchCapabilities
	SearchDeviceDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DeviceMetricSearch interface {
	SearchCapabilitiesDeviceMetric() capabilities.SearchCapabilities
	SearchDeviceMetric(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DeviceRequestSearch interface {
	SearchCapabilitiesDeviceRequest() capabilities.SearchCapabilities
	SearchDeviceRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DeviceUseStatementSearch interface {
	SearchCapabilitiesDeviceUseStatement() capabilities.SearchCapabilities
	SearchDeviceUseStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DiagnosticReportSearch interface {
	SearchCapabilitiesDiagnosticReport() capabilities.SearchCapabilities
	SearchDiagnosticReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DocumentManifestSearch interface {
	SearchCapabilitiesDocumentManifest() capabilities.SearchCapabilities
	SearchDocumentManifest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type DocumentReferenceSearch interface {
	SearchCapabilitiesDocumentReference() capabilities.SearchCapabilities
	SearchDocumentReference(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EffectEvidenceSynthesisSearch interface {
	SearchCapabilitiesEffectEvidenceSynthesis() capabilities.SearchCapabilities
	SearchEffectEvidenceSynthesis(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EncounterSearch interface {
	SearchCapabilitiesEncounter() capabilities.SearchCapabilities
	SearchEncounter(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EndpointSearch interface {
	SearchCapabilitiesEndpoint() capabilities.SearchCapabilities
	SearchEndpoint(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EnrollmentRequestSearch interface {
	SearchCapabilitiesEnrollmentRequest() capabilities.SearchCapabilities
	SearchEnrollmentRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EnrollmentResponseSearch interface {
	SearchCapabilitiesEnrollmentResponse() capabilities.SearchCapabilities
	SearchEnrollmentResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EpisodeOfCareSearch interface {
	SearchCapabilitiesEpisodeOfCare() capabilities.SearchCapabilities
	SearchEpisodeOfCare(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EventDefinitionSearch interface {
	SearchCapabilitiesEventDefinition() capabilities.SearchCapabilities
	SearchEventDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EvidenceSearch interface {
	SearchCapabilitiesEvidence() capabilities.SearchCapabilities
	SearchEvidence(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type EvidenceVariableSearch interface {
	SearchCapabilitiesEvidenceVariable() capabilities.SearchCapabilities
	SearchEvidenceVariable(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ExampleScenarioSearch interface {
	SearchCapabilitiesExampleScenario() capabilities.SearchCapabilities
	SearchExampleScenario(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ExplanationOfBenefitSearch interface {
	SearchCapabilitiesExplanationOfBenefit() capabilities.SearchCapabilities
	SearchExplanationOfBenefit(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type FamilyMemberHistorySearch interface {
	SearchCapabilitiesFamilyMemberHistory() capabilities.SearchCapabilities
	SearchFamilyMemberHistory(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type FlagSearch interface {
	SearchCapabilitiesFlag() capabilities.SearchCapabilities
	SearchFlag(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type GoalSearch interface {
	SearchCapabilitiesGoal() capabilities.SearchCapabilities
	SearchGoal(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type GraphDefinitionSearch interface {
	SearchCapabilitiesGraphDefinition() capabilities.SearchCapabilities
	SearchGraphDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type GroupSearch interface {
	SearchCapabilitiesGroup() capabilities.SearchCapabilities
	SearchGroup(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type GuidanceResponseSearch interface {
	SearchCapabilitiesGuidanceResponse() capabilities.SearchCapabilities
	SearchGuidanceResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type HealthcareServiceSearch interface {
	SearchCapabilitiesHealthcareService() capabilities.SearchCapabilities
	SearchHealthcareService(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ImagingStudySearch interface {
	SearchCapabilitiesImagingStudy() capabilities.SearchCapabilities
	SearchImagingStudy(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ImmunizationSearch interface {
	SearchCapabilitiesImmunization() capabilities.SearchCapabilities
	SearchImmunization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ImmunizationEvaluationSearch interface {
	SearchCapabilitiesImmunizationEvaluation() capabilities.SearchCapabilities
	SearchImmunizationEvaluation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ImmunizationRecommendationSearch interface {
	SearchCapabilitiesImmunizationRecommendation() capabilities.SearchCapabilities
	SearchImmunizationRecommendation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ImplementationGuideSearch interface {
	SearchCapabilitiesImplementationGuide() capabilities.SearchCapabilities
	SearchImplementationGuide(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type InsurancePlanSearch interface {
	SearchCapabilitiesInsurancePlan() capabilities.SearchCapabilities
	SearchInsurancePlan(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type InvoiceSearch interface {
	SearchCapabilitiesInvoice() capabilities.SearchCapabilities
	SearchInvoice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type LibrarySearch interface {
	SearchCapabilitiesLibrary() capabilities.SearchCapabilities
	SearchLibrary(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type LinkageSearch interface {
	SearchCapabilitiesLinkage() capabilities.SearchCapabilities
	SearchLinkage(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ListSearch interface {
	SearchCapabilitiesList() capabilities.SearchCapabilities
	SearchList(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type LocationSearch interface {
	SearchCapabilitiesLocation() capabilities.SearchCapabilities
	SearchLocation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MeasureSearch interface {
	SearchCapabilitiesMeasure() capabilities.SearchCapabilities
	SearchMeasure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MeasureReportSearch interface {
	SearchCapabilitiesMeasureReport() capabilities.SearchCapabilities
	SearchMeasureReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MediaSearch interface {
	SearchCapabilitiesMedia() capabilities.SearchCapabilities
	SearchMedia(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationSearch interface {
	SearchCapabilitiesMedication() capabilities.SearchCapabilities
	SearchMedication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationAdministrationSearch interface {
	SearchCapabilitiesMedicationAdministration() capabilities.SearchCapabilities
	SearchMedicationAdministration(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationDispenseSearch interface {
	SearchCapabilitiesMedicationDispense() capabilities.SearchCapabilities
	SearchMedicationDispense(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationKnowledgeSearch interface {
	SearchCapabilitiesMedicationKnowledge() capabilities.SearchCapabilities
	SearchMedicationKnowledge(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationRequestSearch interface {
	SearchCapabilitiesMedicationRequest() capabilities.SearchCapabilities
	SearchMedicationRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicationStatementSearch interface {
	SearchCapabilitiesMedicationStatement() capabilities.SearchCapabilities
	SearchMedicationStatement(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductSearch interface {
	SearchCapabilitiesMedicinalProduct() capabilities.SearchCapabilities
	SearchMedicinalProduct(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductAuthorizationSearch interface {
	SearchCapabilitiesMedicinalProductAuthorization() capabilities.SearchCapabilities
	SearchMedicinalProductAuthorization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductContraindicationSearch interface {
	SearchCapabilitiesMedicinalProductContraindication() capabilities.SearchCapabilities
	SearchMedicinalProductContraindication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductIndicationSearch interface {
	SearchCapabilitiesMedicinalProductIndication() capabilities.SearchCapabilities
	SearchMedicinalProductIndication(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductIngredientSearch interface {
	SearchCapabilitiesMedicinalProductIngredient() capabilities.SearchCapabilities
	SearchMedicinalProductIngredient(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductInteractionSearch interface {
	SearchCapabilitiesMedicinalProductInteraction() capabilities.SearchCapabilities
	SearchMedicinalProductInteraction(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductManufacturedSearch interface {
	SearchCapabilitiesMedicinalProductManufactured() capabilities.SearchCapabilities
	SearchMedicinalProductManufactured(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductPackagedSearch interface {
	SearchCapabilitiesMedicinalProductPackaged() capabilities.SearchCapabilities
	SearchMedicinalProductPackaged(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductPharmaceuticalSearch interface {
	SearchCapabilitiesMedicinalProductPharmaceutical() capabilities.SearchCapabilities
	SearchMedicinalProductPharmaceutical(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MedicinalProductUndesirableEffectSearch interface {
	SearchCapabilitiesMedicinalProductUndesirableEffect() capabilities.SearchCapabilities
	SearchMedicinalProductUndesirableEffect(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MessageDefinitionSearch interface {
	SearchCapabilitiesMessageDefinition() capabilities.SearchCapabilities
	SearchMessageDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MessageHeaderSearch interface {
	SearchCapabilitiesMessageHeader() capabilities.SearchCapabilities
	SearchMessageHeader(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type MolecularSequenceSearch interface {
	SearchCapabilitiesMolecularSequence() capabilities.SearchCapabilities
	SearchMolecularSequence(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type NamingSystemSearch interface {
	SearchCapabilitiesNamingSystem() capabilities.SearchCapabilities
	SearchNamingSystem(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type NutritionOrderSearch interface {
	SearchCapabilitiesNutritionOrder() capabilities.SearchCapabilities
	SearchNutritionOrder(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ObservationSearch interface {
	SearchCapabilitiesObservation() capabilities.SearchCapabilities
	SearchObservation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ObservationDefinitionSearch interface {
	SearchCapabilitiesObservationDefinition() capabilities.SearchCapabilities
	SearchObservationDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type OperationDefinitionSearch interface {
	SearchCapabilitiesOperationDefinition() capabilities.SearchCapabilities
	SearchOperationDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type OperationOutcomeSearch interface {
	SearchCapabilitiesOperationOutcome() capabilities.SearchCapabilities
	SearchOperationOutcome(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type OrganizationSearch interface {
	SearchCapabilitiesOrganization() capabilities.SearchCapabilities
	SearchOrganization(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type OrganizationAffiliationSearch interface {
	SearchCapabilitiesOrganizationAffiliation() capabilities.SearchCapabilities
	SearchOrganizationAffiliation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ParametersSearch interface {
	SearchCapabilitiesParameters() capabilities.SearchCapabilities
	SearchParameters(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PatientSearch interface {
	SearchCapabilitiesPatient() capabilities.SearchCapabilities
	SearchPatient(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PaymentNoticeSearch interface {
	SearchCapabilitiesPaymentNotice() capabilities.SearchCapabilities
	SearchPaymentNotice(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PaymentReconciliationSearch interface {
	SearchCapabilitiesPaymentReconciliation() capabilities.SearchCapabilities
	SearchPaymentReconciliation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PersonSearch interface {
	SearchCapabilitiesPerson() capabilities.SearchCapabilities
	SearchPerson(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PlanDefinitionSearch interface {
	SearchCapabilitiesPlanDefinition() capabilities.SearchCapabilities
	SearchPlanDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PractitionerSearch interface {
	SearchCapabilitiesPractitioner() capabilities.SearchCapabilities
	SearchPractitioner(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type PractitionerRoleSearch interface {
	SearchCapabilitiesPractitionerRole() capabilities.SearchCapabilities
	SearchPractitionerRole(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ProcedureSearch interface {
	SearchCapabilitiesProcedure() capabilities.SearchCapabilities
	SearchProcedure(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ProvenanceSearch interface {
	SearchCapabilitiesProvenance() capabilities.SearchCapabilities
	SearchProvenance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type QuestionnaireSearch interface {
	SearchCapabilitiesQuestionnaire() capabilities.SearchCapabilities
	SearchQuestionnaire(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type QuestionnaireResponseSearch interface {
	SearchCapabilitiesQuestionnaireResponse() capabilities.SearchCapabilities
	SearchQuestionnaireResponse(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type RelatedPersonSearch interface {
	SearchCapabilitiesRelatedPerson() capabilities.SearchCapabilities
	SearchRelatedPerson(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type RequestGroupSearch interface {
	SearchCapabilitiesRequestGroup() capabilities.SearchCapabilities
	SearchRequestGroup(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ResearchDefinitionSearch interface {
	SearchCapabilitiesResearchDefinition() capabilities.SearchCapabilities
	SearchResearchDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ResearchElementDefinitionSearch interface {
	SearchCapabilitiesResearchElementDefinition() capabilities.SearchCapabilities
	SearchResearchElementDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ResearchStudySearch interface {
	SearchCapabilitiesResearchStudy() capabilities.SearchCapabilities
	SearchResearchStudy(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ResearchSubjectSearch interface {
	SearchCapabilitiesResearchSubject() capabilities.SearchCapabilities
	SearchResearchSubject(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type RiskAssessmentSearch interface {
	SearchCapabilitiesRiskAssessment() capabilities.SearchCapabilities
	SearchRiskAssessment(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type RiskEvidenceSynthesisSearch interface {
	SearchCapabilitiesRiskEvidenceSynthesis() capabilities.SearchCapabilities
	SearchRiskEvidenceSynthesis(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ScheduleSearch interface {
	SearchCapabilitiesSchedule() capabilities.SearchCapabilities
	SearchSchedule(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SearchParameterSearch interface {
	SearchCapabilitiesSearchParameter() capabilities.SearchCapabilities
	SearchSearchParameter(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ServiceRequestSearch interface {
	SearchCapabilitiesServiceRequest() capabilities.SearchCapabilities
	SearchServiceRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SlotSearch interface {
	SearchCapabilitiesSlot() capabilities.SearchCapabilities
	SearchSlot(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SpecimenSearch interface {
	SearchCapabilitiesSpecimen() capabilities.SearchCapabilities
	SearchSpecimen(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SpecimenDefinitionSearch interface {
	SearchCapabilitiesSpecimenDefinition() capabilities.SearchCapabilities
	SearchSpecimenDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type StructureDefinitionSearch interface {
	SearchCapabilitiesStructureDefinition() capabilities.SearchCapabilities
	SearchStructureDefinition(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type StructureMapSearch interface {
	SearchCapabilitiesStructureMap() capabilities.SearchCapabilities
	SearchStructureMap(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubscriptionSearch interface {
	SearchCapabilitiesSubscription() capabilities.SearchCapabilities
	SearchSubscription(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceSearch interface {
	SearchCapabilitiesSubstance() capabilities.SearchCapabilities
	SearchSubstance(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceNucleicAcidSearch interface {
	SearchCapabilitiesSubstanceNucleicAcid() capabilities.SearchCapabilities
	SearchSubstanceNucleicAcid(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstancePolymerSearch interface {
	SearchCapabilitiesSubstancePolymer() capabilities.SearchCapabilities
	SearchSubstancePolymer(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceProteinSearch interface {
	SearchCapabilitiesSubstanceProtein() capabilities.SearchCapabilities
	SearchSubstanceProtein(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceReferenceInformationSearch interface {
	SearchCapabilitiesSubstanceReferenceInformation() capabilities.SearchCapabilities
	SearchSubstanceReferenceInformation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceSourceMaterialSearch interface {
	SearchCapabilitiesSubstanceSourceMaterial() capabilities.SearchCapabilities
	SearchSubstanceSourceMaterial(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SubstanceSpecificationSearch interface {
	SearchCapabilitiesSubstanceSpecification() capabilities.SearchCapabilities
	SearchSubstanceSpecification(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SupplyDeliverySearch interface {
	SearchCapabilitiesSupplyDelivery() capabilities.SearchCapabilities
	SearchSupplyDelivery(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type SupplyRequestSearch interface {
	SearchCapabilitiesSupplyRequest() capabilities.SearchCapabilities
	SearchSupplyRequest(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type TaskSearch interface {
	SearchCapabilitiesTask() capabilities.SearchCapabilities
	SearchTask(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type TerminologyCapabilitiesSearch interface {
	SearchCapabilitiesTerminologyCapabilities() capabilities.SearchCapabilities
	SearchTerminologyCapabilities(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type TestReportSearch interface {
	SearchCapabilitiesTestReport() capabilities.SearchCapabilities
	SearchTestReport(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type TestScriptSearch interface {
	SearchCapabilitiesTestScript() capabilities.SearchCapabilities
	SearchTestScript(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type ValueSetSearch interface {
	SearchCapabilitiesValueSet() capabilities.SearchCapabilities
	SearchValueSet(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type VerificationResultSearch interface {
	SearchCapabilitiesVerificationResult() capabilities.SearchCapabilities
	SearchVerificationResult(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
type VisionPrescriptionSearch interface {
	SearchCapabilitiesVisionPrescription() capabilities.SearchCapabilities
	SearchVisionPrescription(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, error)
}
