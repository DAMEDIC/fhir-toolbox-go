package capabilitiesR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	model "fhir-toolbox/model"
)

type AccountSearch interface {
	SearchCapabilitiesAccount() search.Capabilities
	SearchAccount(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ActivityDefinitionSearch interface {
	SearchCapabilitiesActivityDefinition() search.Capabilities
	SearchActivityDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type AdverseEventSearch interface {
	SearchCapabilitiesAdverseEvent() search.Capabilities
	SearchAdverseEvent(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type AllergyIntoleranceSearch interface {
	SearchCapabilitiesAllergyIntolerance() search.Capabilities
	SearchAllergyIntolerance(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type AppointmentSearch interface {
	SearchCapabilitiesAppointment() search.Capabilities
	SearchAppointment(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type AppointmentResponseSearch interface {
	SearchCapabilitiesAppointmentResponse() search.Capabilities
	SearchAppointmentResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type AuditEventSearch interface {
	SearchCapabilitiesAuditEvent() search.Capabilities
	SearchAuditEvent(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type BasicSearch interface {
	SearchCapabilitiesBasic() search.Capabilities
	SearchBasic(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type BinarySearch interface {
	SearchCapabilitiesBinary() search.Capabilities
	SearchBinary(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type BiologicallyDerivedProductSearch interface {
	SearchCapabilitiesBiologicallyDerivedProduct() search.Capabilities
	SearchBiologicallyDerivedProduct(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type BodyStructureSearch interface {
	SearchCapabilitiesBodyStructure() search.Capabilities
	SearchBodyStructure(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type BundleSearch interface {
	SearchCapabilitiesBundle() search.Capabilities
	SearchBundle(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CapabilityStatementSearch interface {
	SearchCapabilitiesCapabilityStatement() search.Capabilities
	SearchCapabilityStatement(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CarePlanSearch interface {
	SearchCapabilitiesCarePlan() search.Capabilities
	SearchCarePlan(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CareTeamSearch interface {
	SearchCapabilitiesCareTeam() search.Capabilities
	SearchCareTeam(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CatalogEntrySearch interface {
	SearchCapabilitiesCatalogEntry() search.Capabilities
	SearchCatalogEntry(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ChargeItemSearch interface {
	SearchCapabilitiesChargeItem() search.Capabilities
	SearchChargeItem(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ChargeItemDefinitionSearch interface {
	SearchCapabilitiesChargeItemDefinition() search.Capabilities
	SearchChargeItemDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ClaimSearch interface {
	SearchCapabilitiesClaim() search.Capabilities
	SearchClaim(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ClaimResponseSearch interface {
	SearchCapabilitiesClaimResponse() search.Capabilities
	SearchClaimResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ClinicalImpressionSearch interface {
	SearchCapabilitiesClinicalImpression() search.Capabilities
	SearchClinicalImpression(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CodeSystemSearch interface {
	SearchCapabilitiesCodeSystem() search.Capabilities
	SearchCodeSystem(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CommunicationSearch interface {
	SearchCapabilitiesCommunication() search.Capabilities
	SearchCommunication(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CommunicationRequestSearch interface {
	SearchCapabilitiesCommunicationRequest() search.Capabilities
	SearchCommunicationRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CompartmentDefinitionSearch interface {
	SearchCapabilitiesCompartmentDefinition() search.Capabilities
	SearchCompartmentDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CompositionSearch interface {
	SearchCapabilitiesComposition() search.Capabilities
	SearchComposition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ConceptMapSearch interface {
	SearchCapabilitiesConceptMap() search.Capabilities
	SearchConceptMap(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ConditionSearch interface {
	SearchCapabilitiesCondition() search.Capabilities
	SearchCondition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ConsentSearch interface {
	SearchCapabilitiesConsent() search.Capabilities
	SearchConsent(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ContractSearch interface {
	SearchCapabilitiesContract() search.Capabilities
	SearchContract(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CoverageSearch interface {
	SearchCapabilitiesCoverage() search.Capabilities
	SearchCoverage(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CoverageEligibilityRequestSearch interface {
	SearchCapabilitiesCoverageEligibilityRequest() search.Capabilities
	SearchCoverageEligibilityRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type CoverageEligibilityResponseSearch interface {
	SearchCapabilitiesCoverageEligibilityResponse() search.Capabilities
	SearchCoverageEligibilityResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DetectedIssueSearch interface {
	SearchCapabilitiesDetectedIssue() search.Capabilities
	SearchDetectedIssue(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DeviceSearch interface {
	SearchCapabilitiesDevice() search.Capabilities
	SearchDevice(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DeviceDefinitionSearch interface {
	SearchCapabilitiesDeviceDefinition() search.Capabilities
	SearchDeviceDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DeviceMetricSearch interface {
	SearchCapabilitiesDeviceMetric() search.Capabilities
	SearchDeviceMetric(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DeviceRequestSearch interface {
	SearchCapabilitiesDeviceRequest() search.Capabilities
	SearchDeviceRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DeviceUseStatementSearch interface {
	SearchCapabilitiesDeviceUseStatement() search.Capabilities
	SearchDeviceUseStatement(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DiagnosticReportSearch interface {
	SearchCapabilitiesDiagnosticReport() search.Capabilities
	SearchDiagnosticReport(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DocumentManifestSearch interface {
	SearchCapabilitiesDocumentManifest() search.Capabilities
	SearchDocumentManifest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type DocumentReferenceSearch interface {
	SearchCapabilitiesDocumentReference() search.Capabilities
	SearchDocumentReference(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EffectEvidenceSynthesisSearch interface {
	SearchCapabilitiesEffectEvidenceSynthesis() search.Capabilities
	SearchEffectEvidenceSynthesis(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EncounterSearch interface {
	SearchCapabilitiesEncounter() search.Capabilities
	SearchEncounter(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EndpointSearch interface {
	SearchCapabilitiesEndpoint() search.Capabilities
	SearchEndpoint(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EnrollmentRequestSearch interface {
	SearchCapabilitiesEnrollmentRequest() search.Capabilities
	SearchEnrollmentRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EnrollmentResponseSearch interface {
	SearchCapabilitiesEnrollmentResponse() search.Capabilities
	SearchEnrollmentResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EpisodeOfCareSearch interface {
	SearchCapabilitiesEpisodeOfCare() search.Capabilities
	SearchEpisodeOfCare(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EventDefinitionSearch interface {
	SearchCapabilitiesEventDefinition() search.Capabilities
	SearchEventDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EvidenceSearch interface {
	SearchCapabilitiesEvidence() search.Capabilities
	SearchEvidence(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type EvidenceVariableSearch interface {
	SearchCapabilitiesEvidenceVariable() search.Capabilities
	SearchEvidenceVariable(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ExampleScenarioSearch interface {
	SearchCapabilitiesExampleScenario() search.Capabilities
	SearchExampleScenario(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ExplanationOfBenefitSearch interface {
	SearchCapabilitiesExplanationOfBenefit() search.Capabilities
	SearchExplanationOfBenefit(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type FamilyMemberHistorySearch interface {
	SearchCapabilitiesFamilyMemberHistory() search.Capabilities
	SearchFamilyMemberHistory(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type FlagSearch interface {
	SearchCapabilitiesFlag() search.Capabilities
	SearchFlag(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type GoalSearch interface {
	SearchCapabilitiesGoal() search.Capabilities
	SearchGoal(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type GraphDefinitionSearch interface {
	SearchCapabilitiesGraphDefinition() search.Capabilities
	SearchGraphDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type GroupSearch interface {
	SearchCapabilitiesGroup() search.Capabilities
	SearchGroup(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type GuidanceResponseSearch interface {
	SearchCapabilitiesGuidanceResponse() search.Capabilities
	SearchGuidanceResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type HealthcareServiceSearch interface {
	SearchCapabilitiesHealthcareService() search.Capabilities
	SearchHealthcareService(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ImagingStudySearch interface {
	SearchCapabilitiesImagingStudy() search.Capabilities
	SearchImagingStudy(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationSearch interface {
	SearchCapabilitiesImmunization() search.Capabilities
	SearchImmunization(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationEvaluationSearch interface {
	SearchCapabilitiesImmunizationEvaluation() search.Capabilities
	SearchImmunizationEvaluation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ImmunizationRecommendationSearch interface {
	SearchCapabilitiesImmunizationRecommendation() search.Capabilities
	SearchImmunizationRecommendation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ImplementationGuideSearch interface {
	SearchCapabilitiesImplementationGuide() search.Capabilities
	SearchImplementationGuide(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type InsurancePlanSearch interface {
	SearchCapabilitiesInsurancePlan() search.Capabilities
	SearchInsurancePlan(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type InvoiceSearch interface {
	SearchCapabilitiesInvoice() search.Capabilities
	SearchInvoice(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type LibrarySearch interface {
	SearchCapabilitiesLibrary() search.Capabilities
	SearchLibrary(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type LinkageSearch interface {
	SearchCapabilitiesLinkage() search.Capabilities
	SearchLinkage(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ListSearch interface {
	SearchCapabilitiesList() search.Capabilities
	SearchList(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type LocationSearch interface {
	SearchCapabilitiesLocation() search.Capabilities
	SearchLocation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MeasureSearch interface {
	SearchCapabilitiesMeasure() search.Capabilities
	SearchMeasure(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MeasureReportSearch interface {
	SearchCapabilitiesMeasureReport() search.Capabilities
	SearchMeasureReport(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MediaSearch interface {
	SearchCapabilitiesMedia() search.Capabilities
	SearchMedia(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationSearch interface {
	SearchCapabilitiesMedication() search.Capabilities
	SearchMedication(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationAdministrationSearch interface {
	SearchCapabilitiesMedicationAdministration() search.Capabilities
	SearchMedicationAdministration(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationDispenseSearch interface {
	SearchCapabilitiesMedicationDispense() search.Capabilities
	SearchMedicationDispense(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationKnowledgeSearch interface {
	SearchCapabilitiesMedicationKnowledge() search.Capabilities
	SearchMedicationKnowledge(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationRequestSearch interface {
	SearchCapabilitiesMedicationRequest() search.Capabilities
	SearchMedicationRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicationStatementSearch interface {
	SearchCapabilitiesMedicationStatement() search.Capabilities
	SearchMedicationStatement(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductSearch interface {
	SearchCapabilitiesMedicinalProduct() search.Capabilities
	SearchMedicinalProduct(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductAuthorizationSearch interface {
	SearchCapabilitiesMedicinalProductAuthorization() search.Capabilities
	SearchMedicinalProductAuthorization(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductContraindicationSearch interface {
	SearchCapabilitiesMedicinalProductContraindication() search.Capabilities
	SearchMedicinalProductContraindication(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductIndicationSearch interface {
	SearchCapabilitiesMedicinalProductIndication() search.Capabilities
	SearchMedicinalProductIndication(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductIngredientSearch interface {
	SearchCapabilitiesMedicinalProductIngredient() search.Capabilities
	SearchMedicinalProductIngredient(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductInteractionSearch interface {
	SearchCapabilitiesMedicinalProductInteraction() search.Capabilities
	SearchMedicinalProductInteraction(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductManufacturedSearch interface {
	SearchCapabilitiesMedicinalProductManufactured() search.Capabilities
	SearchMedicinalProductManufactured(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductPackagedSearch interface {
	SearchCapabilitiesMedicinalProductPackaged() search.Capabilities
	SearchMedicinalProductPackaged(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductPharmaceuticalSearch interface {
	SearchCapabilitiesMedicinalProductPharmaceutical() search.Capabilities
	SearchMedicinalProductPharmaceutical(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MedicinalProductUndesirableEffectSearch interface {
	SearchCapabilitiesMedicinalProductUndesirableEffect() search.Capabilities
	SearchMedicinalProductUndesirableEffect(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MessageDefinitionSearch interface {
	SearchCapabilitiesMessageDefinition() search.Capabilities
	SearchMessageDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MessageHeaderSearch interface {
	SearchCapabilitiesMessageHeader() search.Capabilities
	SearchMessageHeader(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type MolecularSequenceSearch interface {
	SearchCapabilitiesMolecularSequence() search.Capabilities
	SearchMolecularSequence(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type NamingSystemSearch interface {
	SearchCapabilitiesNamingSystem() search.Capabilities
	SearchNamingSystem(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type NutritionOrderSearch interface {
	SearchCapabilitiesNutritionOrder() search.Capabilities
	SearchNutritionOrder(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ObservationSearch interface {
	SearchCapabilitiesObservation() search.Capabilities
	SearchObservation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ObservationDefinitionSearch interface {
	SearchCapabilitiesObservationDefinition() search.Capabilities
	SearchObservationDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type OperationDefinitionSearch interface {
	SearchCapabilitiesOperationDefinition() search.Capabilities
	SearchOperationDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type OperationOutcomeSearch interface {
	SearchCapabilitiesOperationOutcome() search.Capabilities
	SearchOperationOutcome(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type OrganizationSearch interface {
	SearchCapabilitiesOrganization() search.Capabilities
	SearchOrganization(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type OrganizationAffiliationSearch interface {
	SearchCapabilitiesOrganizationAffiliation() search.Capabilities
	SearchOrganizationAffiliation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ParametersSearch interface {
	SearchCapabilitiesParameters() search.Capabilities
	SearchParameters(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PatientSearch interface {
	SearchCapabilitiesPatient() search.Capabilities
	SearchPatient(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PaymentNoticeSearch interface {
	SearchCapabilitiesPaymentNotice() search.Capabilities
	SearchPaymentNotice(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PaymentReconciliationSearch interface {
	SearchCapabilitiesPaymentReconciliation() search.Capabilities
	SearchPaymentReconciliation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PersonSearch interface {
	SearchCapabilitiesPerson() search.Capabilities
	SearchPerson(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PlanDefinitionSearch interface {
	SearchCapabilitiesPlanDefinition() search.Capabilities
	SearchPlanDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PractitionerSearch interface {
	SearchCapabilitiesPractitioner() search.Capabilities
	SearchPractitioner(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type PractitionerRoleSearch interface {
	SearchCapabilitiesPractitionerRole() search.Capabilities
	SearchPractitionerRole(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ProcedureSearch interface {
	SearchCapabilitiesProcedure() search.Capabilities
	SearchProcedure(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ProvenanceSearch interface {
	SearchCapabilitiesProvenance() search.Capabilities
	SearchProvenance(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type QuestionnaireSearch interface {
	SearchCapabilitiesQuestionnaire() search.Capabilities
	SearchQuestionnaire(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type QuestionnaireResponseSearch interface {
	SearchCapabilitiesQuestionnaireResponse() search.Capabilities
	SearchQuestionnaireResponse(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type RelatedPersonSearch interface {
	SearchCapabilitiesRelatedPerson() search.Capabilities
	SearchRelatedPerson(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type RequestGroupSearch interface {
	SearchCapabilitiesRequestGroup() search.Capabilities
	SearchRequestGroup(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ResearchDefinitionSearch interface {
	SearchCapabilitiesResearchDefinition() search.Capabilities
	SearchResearchDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ResearchElementDefinitionSearch interface {
	SearchCapabilitiesResearchElementDefinition() search.Capabilities
	SearchResearchElementDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ResearchStudySearch interface {
	SearchCapabilitiesResearchStudy() search.Capabilities
	SearchResearchStudy(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ResearchSubjectSearch interface {
	SearchCapabilitiesResearchSubject() search.Capabilities
	SearchResearchSubject(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type RiskAssessmentSearch interface {
	SearchCapabilitiesRiskAssessment() search.Capabilities
	SearchRiskAssessment(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type RiskEvidenceSynthesisSearch interface {
	SearchCapabilitiesRiskEvidenceSynthesis() search.Capabilities
	SearchRiskEvidenceSynthesis(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ScheduleSearch interface {
	SearchCapabilitiesSchedule() search.Capabilities
	SearchSchedule(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SearchParameterSearch interface {
	SearchCapabilitiesSearchParameter() search.Capabilities
	SearchSearchParameter(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ServiceRequestSearch interface {
	SearchCapabilitiesServiceRequest() search.Capabilities
	SearchServiceRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SlotSearch interface {
	SearchCapabilitiesSlot() search.Capabilities
	SearchSlot(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SpecimenSearch interface {
	SearchCapabilitiesSpecimen() search.Capabilities
	SearchSpecimen(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SpecimenDefinitionSearch interface {
	SearchCapabilitiesSpecimenDefinition() search.Capabilities
	SearchSpecimenDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type StructureDefinitionSearch interface {
	SearchCapabilitiesStructureDefinition() search.Capabilities
	SearchStructureDefinition(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type StructureMapSearch interface {
	SearchCapabilitiesStructureMap() search.Capabilities
	SearchStructureMap(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubscriptionSearch interface {
	SearchCapabilitiesSubscription() search.Capabilities
	SearchSubscription(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSearch interface {
	SearchCapabilitiesSubstance() search.Capabilities
	SearchSubstance(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceNucleicAcidSearch interface {
	SearchCapabilitiesSubstanceNucleicAcid() search.Capabilities
	SearchSubstanceNucleicAcid(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstancePolymerSearch interface {
	SearchCapabilitiesSubstancePolymer() search.Capabilities
	SearchSubstancePolymer(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceProteinSearch interface {
	SearchCapabilitiesSubstanceProtein() search.Capabilities
	SearchSubstanceProtein(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceReferenceInformationSearch interface {
	SearchCapabilitiesSubstanceReferenceInformation() search.Capabilities
	SearchSubstanceReferenceInformation(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSourceMaterialSearch interface {
	SearchCapabilitiesSubstanceSourceMaterial() search.Capabilities
	SearchSubstanceSourceMaterial(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SubstanceSpecificationSearch interface {
	SearchCapabilitiesSubstanceSpecification() search.Capabilities
	SearchSubstanceSpecification(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SupplyDeliverySearch interface {
	SearchCapabilitiesSupplyDelivery() search.Capabilities
	SearchSupplyDelivery(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type SupplyRequestSearch interface {
	SearchCapabilitiesSupplyRequest() search.Capabilities
	SearchSupplyRequest(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type TaskSearch interface {
	SearchCapabilitiesTask() search.Capabilities
	SearchTask(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type TerminologyCapabilitiesSearch interface {
	SearchCapabilitiesTerminologyCapabilities() search.Capabilities
	SearchTerminologyCapabilities(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type TestReportSearch interface {
	SearchCapabilitiesTestReport() search.Capabilities
	SearchTestReport(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type TestScriptSearch interface {
	SearchCapabilitiesTestScript() search.Capabilities
	SearchTestScript(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type ValueSetSearch interface {
	SearchCapabilitiesValueSet() search.Capabilities
	SearchValueSet(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type VerificationResultSearch interface {
	SearchCapabilitiesVerificationResult() search.Capabilities
	SearchVerificationResult(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
type VisionPrescriptionSearch interface {
	SearchCapabilitiesVisionPrescription() search.Capabilities
	SearchVisionPrescription(ctx context.Context, options search.Options) ([]model.Resource, capabilities.FHIRError)
}
