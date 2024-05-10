package capabilitiesR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/model/gen/r4"
)

type AccountSearch interface {
	SearchCapabilitiesAccount() capabilities.SearchCapabilities
	SearchAccount(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Account, error)
}
type ActivityDefinitionSearch interface {
	SearchCapabilitiesActivityDefinition() capabilities.SearchCapabilities
	SearchActivityDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ActivityDefinition, error)
}
type AdverseEventSearch interface {
	SearchCapabilitiesAdverseEvent() capabilities.SearchCapabilities
	SearchAdverseEvent(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.AdverseEvent, error)
}
type AllergyIntoleranceSearch interface {
	SearchCapabilitiesAllergyIntolerance() capabilities.SearchCapabilities
	SearchAllergyIntolerance(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.AllergyIntolerance, error)
}
type AppointmentSearch interface {
	SearchCapabilitiesAppointment() capabilities.SearchCapabilities
	SearchAppointment(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Appointment, error)
}
type AppointmentResponseSearch interface {
	SearchCapabilitiesAppointmentResponse() capabilities.SearchCapabilities
	SearchAppointmentResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.AppointmentResponse, error)
}
type AuditEventSearch interface {
	SearchCapabilitiesAuditEvent() capabilities.SearchCapabilities
	SearchAuditEvent(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.AuditEvent, error)
}
type BasicSearch interface {
	SearchCapabilitiesBasic() capabilities.SearchCapabilities
	SearchBasic(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Basic, error)
}
type BinarySearch interface {
	SearchCapabilitiesBinary() capabilities.SearchCapabilities
	SearchBinary(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Binary, error)
}
type BiologicallyDerivedProductSearch interface {
	SearchCapabilitiesBiologicallyDerivedProduct() capabilities.SearchCapabilities
	SearchBiologicallyDerivedProduct(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.BiologicallyDerivedProduct, error)
}
type BodyStructureSearch interface {
	SearchCapabilitiesBodyStructure() capabilities.SearchCapabilities
	SearchBodyStructure(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.BodyStructure, error)
}
type BundleSearch interface {
	SearchCapabilitiesBundle() capabilities.SearchCapabilities
	SearchBundle(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Bundle, error)
}
type CapabilityStatementSearch interface {
	SearchCapabilitiesCapabilityStatement() capabilities.SearchCapabilities
	SearchCapabilityStatement(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CapabilityStatement, error)
}
type CarePlanSearch interface {
	SearchCapabilitiesCarePlan() capabilities.SearchCapabilities
	SearchCarePlan(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CarePlan, error)
}
type CareTeamSearch interface {
	SearchCapabilitiesCareTeam() capabilities.SearchCapabilities
	SearchCareTeam(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CareTeam, error)
}
type CatalogEntrySearch interface {
	SearchCapabilitiesCatalogEntry() capabilities.SearchCapabilities
	SearchCatalogEntry(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CatalogEntry, error)
}
type ChargeItemSearch interface {
	SearchCapabilitiesChargeItem() capabilities.SearchCapabilities
	SearchChargeItem(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ChargeItem, error)
}
type ChargeItemDefinitionSearch interface {
	SearchCapabilitiesChargeItemDefinition() capabilities.SearchCapabilities
	SearchChargeItemDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ChargeItemDefinition, error)
}
type ClaimSearch interface {
	SearchCapabilitiesClaim() capabilities.SearchCapabilities
	SearchClaim(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Claim, error)
}
type ClaimResponseSearch interface {
	SearchCapabilitiesClaimResponse() capabilities.SearchCapabilities
	SearchClaimResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ClaimResponse, error)
}
type ClinicalImpressionSearch interface {
	SearchCapabilitiesClinicalImpression() capabilities.SearchCapabilities
	SearchClinicalImpression(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ClinicalImpression, error)
}
type CodeSystemSearch interface {
	SearchCapabilitiesCodeSystem() capabilities.SearchCapabilities
	SearchCodeSystem(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CodeSystem, error)
}
type CommunicationSearch interface {
	SearchCapabilitiesCommunication() capabilities.SearchCapabilities
	SearchCommunication(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Communication, error)
}
type CommunicationRequestSearch interface {
	SearchCapabilitiesCommunicationRequest() capabilities.SearchCapabilities
	SearchCommunicationRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CommunicationRequest, error)
}
type CompartmentDefinitionSearch interface {
	SearchCapabilitiesCompartmentDefinition() capabilities.SearchCapabilities
	SearchCompartmentDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CompartmentDefinition, error)
}
type CompositionSearch interface {
	SearchCapabilitiesComposition() capabilities.SearchCapabilities
	SearchComposition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Composition, error)
}
type ConceptMapSearch interface {
	SearchCapabilitiesConceptMap() capabilities.SearchCapabilities
	SearchConceptMap(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ConceptMap, error)
}
type ConditionSearch interface {
	SearchCapabilitiesCondition() capabilities.SearchCapabilities
	SearchCondition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Condition, error)
}
type ConsentSearch interface {
	SearchCapabilitiesConsent() capabilities.SearchCapabilities
	SearchConsent(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Consent, error)
}
type ContractSearch interface {
	SearchCapabilitiesContract() capabilities.SearchCapabilities
	SearchContract(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Contract, error)
}
type CoverageSearch interface {
	SearchCapabilitiesCoverage() capabilities.SearchCapabilities
	SearchCoverage(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Coverage, error)
}
type CoverageEligibilityRequestSearch interface {
	SearchCapabilitiesCoverageEligibilityRequest() capabilities.SearchCapabilities
	SearchCoverageEligibilityRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CoverageEligibilityRequest, error)
}
type CoverageEligibilityResponseSearch interface {
	SearchCapabilitiesCoverageEligibilityResponse() capabilities.SearchCapabilities
	SearchCoverageEligibilityResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.CoverageEligibilityResponse, error)
}
type DetectedIssueSearch interface {
	SearchCapabilitiesDetectedIssue() capabilities.SearchCapabilities
	SearchDetectedIssue(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DetectedIssue, error)
}
type DeviceSearch interface {
	SearchCapabilitiesDevice() capabilities.SearchCapabilities
	SearchDevice(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Device, error)
}
type DeviceDefinitionSearch interface {
	SearchCapabilitiesDeviceDefinition() capabilities.SearchCapabilities
	SearchDeviceDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DeviceDefinition, error)
}
type DeviceMetricSearch interface {
	SearchCapabilitiesDeviceMetric() capabilities.SearchCapabilities
	SearchDeviceMetric(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DeviceMetric, error)
}
type DeviceRequestSearch interface {
	SearchCapabilitiesDeviceRequest() capabilities.SearchCapabilities
	SearchDeviceRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DeviceRequest, error)
}
type DeviceUseStatementSearch interface {
	SearchCapabilitiesDeviceUseStatement() capabilities.SearchCapabilities
	SearchDeviceUseStatement(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DeviceUseStatement, error)
}
type DiagnosticReportSearch interface {
	SearchCapabilitiesDiagnosticReport() capabilities.SearchCapabilities
	SearchDiagnosticReport(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DiagnosticReport, error)
}
type DocumentManifestSearch interface {
	SearchCapabilitiesDocumentManifest() capabilities.SearchCapabilities
	SearchDocumentManifest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DocumentManifest, error)
}
type DocumentReferenceSearch interface {
	SearchCapabilitiesDocumentReference() capabilities.SearchCapabilities
	SearchDocumentReference(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.DocumentReference, error)
}
type EffectEvidenceSynthesisSearch interface {
	SearchCapabilitiesEffectEvidenceSynthesis() capabilities.SearchCapabilities
	SearchEffectEvidenceSynthesis(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EffectEvidenceSynthesis, error)
}
type EncounterSearch interface {
	SearchCapabilitiesEncounter() capabilities.SearchCapabilities
	SearchEncounter(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Encounter, error)
}
type EndpointSearch interface {
	SearchCapabilitiesEndpoint() capabilities.SearchCapabilities
	SearchEndpoint(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Endpoint, error)
}
type EnrollmentRequestSearch interface {
	SearchCapabilitiesEnrollmentRequest() capabilities.SearchCapabilities
	SearchEnrollmentRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EnrollmentRequest, error)
}
type EnrollmentResponseSearch interface {
	SearchCapabilitiesEnrollmentResponse() capabilities.SearchCapabilities
	SearchEnrollmentResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EnrollmentResponse, error)
}
type EpisodeOfCareSearch interface {
	SearchCapabilitiesEpisodeOfCare() capabilities.SearchCapabilities
	SearchEpisodeOfCare(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EpisodeOfCare, error)
}
type EventDefinitionSearch interface {
	SearchCapabilitiesEventDefinition() capabilities.SearchCapabilities
	SearchEventDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EventDefinition, error)
}
type EvidenceSearch interface {
	SearchCapabilitiesEvidence() capabilities.SearchCapabilities
	SearchEvidence(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Evidence, error)
}
type EvidenceVariableSearch interface {
	SearchCapabilitiesEvidenceVariable() capabilities.SearchCapabilities
	SearchEvidenceVariable(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.EvidenceVariable, error)
}
type ExampleScenarioSearch interface {
	SearchCapabilitiesExampleScenario() capabilities.SearchCapabilities
	SearchExampleScenario(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ExampleScenario, error)
}
type ExplanationOfBenefitSearch interface {
	SearchCapabilitiesExplanationOfBenefit() capabilities.SearchCapabilities
	SearchExplanationOfBenefit(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ExplanationOfBenefit, error)
}
type FamilyMemberHistorySearch interface {
	SearchCapabilitiesFamilyMemberHistory() capabilities.SearchCapabilities
	SearchFamilyMemberHistory(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.FamilyMemberHistory, error)
}
type FlagSearch interface {
	SearchCapabilitiesFlag() capabilities.SearchCapabilities
	SearchFlag(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Flag, error)
}
type GoalSearch interface {
	SearchCapabilitiesGoal() capabilities.SearchCapabilities
	SearchGoal(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Goal, error)
}
type GraphDefinitionSearch interface {
	SearchCapabilitiesGraphDefinition() capabilities.SearchCapabilities
	SearchGraphDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.GraphDefinition, error)
}
type GroupSearch interface {
	SearchCapabilitiesGroup() capabilities.SearchCapabilities
	SearchGroup(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Group, error)
}
type GuidanceResponseSearch interface {
	SearchCapabilitiesGuidanceResponse() capabilities.SearchCapabilities
	SearchGuidanceResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.GuidanceResponse, error)
}
type HealthcareServiceSearch interface {
	SearchCapabilitiesHealthcareService() capabilities.SearchCapabilities
	SearchHealthcareService(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.HealthcareService, error)
}
type ImagingStudySearch interface {
	SearchCapabilitiesImagingStudy() capabilities.SearchCapabilities
	SearchImagingStudy(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ImagingStudy, error)
}
type ImmunizationSearch interface {
	SearchCapabilitiesImmunization() capabilities.SearchCapabilities
	SearchImmunization(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Immunization, error)
}
type ImmunizationEvaluationSearch interface {
	SearchCapabilitiesImmunizationEvaluation() capabilities.SearchCapabilities
	SearchImmunizationEvaluation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ImmunizationEvaluation, error)
}
type ImmunizationRecommendationSearch interface {
	SearchCapabilitiesImmunizationRecommendation() capabilities.SearchCapabilities
	SearchImmunizationRecommendation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ImmunizationRecommendation, error)
}
type ImplementationGuideSearch interface {
	SearchCapabilitiesImplementationGuide() capabilities.SearchCapabilities
	SearchImplementationGuide(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ImplementationGuide, error)
}
type InsurancePlanSearch interface {
	SearchCapabilitiesInsurancePlan() capabilities.SearchCapabilities
	SearchInsurancePlan(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.InsurancePlan, error)
}
type InvoiceSearch interface {
	SearchCapabilitiesInvoice() capabilities.SearchCapabilities
	SearchInvoice(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Invoice, error)
}
type LibrarySearch interface {
	SearchCapabilitiesLibrary() capabilities.SearchCapabilities
	SearchLibrary(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Library, error)
}
type LinkageSearch interface {
	SearchCapabilitiesLinkage() capabilities.SearchCapabilities
	SearchLinkage(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Linkage, error)
}
type ListSearch interface {
	SearchCapabilitiesList() capabilities.SearchCapabilities
	SearchList(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.List, error)
}
type LocationSearch interface {
	SearchCapabilitiesLocation() capabilities.SearchCapabilities
	SearchLocation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Location, error)
}
type MeasureSearch interface {
	SearchCapabilitiesMeasure() capabilities.SearchCapabilities
	SearchMeasure(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Measure, error)
}
type MeasureReportSearch interface {
	SearchCapabilitiesMeasureReport() capabilities.SearchCapabilities
	SearchMeasureReport(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MeasureReport, error)
}
type MediaSearch interface {
	SearchCapabilitiesMedia() capabilities.SearchCapabilities
	SearchMedia(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Media, error)
}
type MedicationSearch interface {
	SearchCapabilitiesMedication() capabilities.SearchCapabilities
	SearchMedication(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Medication, error)
}
type MedicationAdministrationSearch interface {
	SearchCapabilitiesMedicationAdministration() capabilities.SearchCapabilities
	SearchMedicationAdministration(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicationAdministration, error)
}
type MedicationDispenseSearch interface {
	SearchCapabilitiesMedicationDispense() capabilities.SearchCapabilities
	SearchMedicationDispense(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicationDispense, error)
}
type MedicationKnowledgeSearch interface {
	SearchCapabilitiesMedicationKnowledge() capabilities.SearchCapabilities
	SearchMedicationKnowledge(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicationKnowledge, error)
}
type MedicationRequestSearch interface {
	SearchCapabilitiesMedicationRequest() capabilities.SearchCapabilities
	SearchMedicationRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicationRequest, error)
}
type MedicationStatementSearch interface {
	SearchCapabilitiesMedicationStatement() capabilities.SearchCapabilities
	SearchMedicationStatement(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicationStatement, error)
}
type MedicinalProductSearch interface {
	SearchCapabilitiesMedicinalProduct() capabilities.SearchCapabilities
	SearchMedicinalProduct(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProduct, error)
}
type MedicinalProductAuthorizationSearch interface {
	SearchCapabilitiesMedicinalProductAuthorization() capabilities.SearchCapabilities
	SearchMedicinalProductAuthorization(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductAuthorization, error)
}
type MedicinalProductContraindicationSearch interface {
	SearchCapabilitiesMedicinalProductContraindication() capabilities.SearchCapabilities
	SearchMedicinalProductContraindication(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductContraindication, error)
}
type MedicinalProductIndicationSearch interface {
	SearchCapabilitiesMedicinalProductIndication() capabilities.SearchCapabilities
	SearchMedicinalProductIndication(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductIndication, error)
}
type MedicinalProductIngredientSearch interface {
	SearchCapabilitiesMedicinalProductIngredient() capabilities.SearchCapabilities
	SearchMedicinalProductIngredient(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductIngredient, error)
}
type MedicinalProductInteractionSearch interface {
	SearchCapabilitiesMedicinalProductInteraction() capabilities.SearchCapabilities
	SearchMedicinalProductInteraction(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductInteraction, error)
}
type MedicinalProductManufacturedSearch interface {
	SearchCapabilitiesMedicinalProductManufactured() capabilities.SearchCapabilities
	SearchMedicinalProductManufactured(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductManufactured, error)
}
type MedicinalProductPackagedSearch interface {
	SearchCapabilitiesMedicinalProductPackaged() capabilities.SearchCapabilities
	SearchMedicinalProductPackaged(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductPackaged, error)
}
type MedicinalProductPharmaceuticalSearch interface {
	SearchCapabilitiesMedicinalProductPharmaceutical() capabilities.SearchCapabilities
	SearchMedicinalProductPharmaceutical(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductPharmaceutical, error)
}
type MedicinalProductUndesirableEffectSearch interface {
	SearchCapabilitiesMedicinalProductUndesirableEffect() capabilities.SearchCapabilities
	SearchMedicinalProductUndesirableEffect(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MedicinalProductUndesirableEffect, error)
}
type MessageDefinitionSearch interface {
	SearchCapabilitiesMessageDefinition() capabilities.SearchCapabilities
	SearchMessageDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MessageDefinition, error)
}
type MessageHeaderSearch interface {
	SearchCapabilitiesMessageHeader() capabilities.SearchCapabilities
	SearchMessageHeader(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MessageHeader, error)
}
type MolecularSequenceSearch interface {
	SearchCapabilitiesMolecularSequence() capabilities.SearchCapabilities
	SearchMolecularSequence(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.MolecularSequence, error)
}
type NamingSystemSearch interface {
	SearchCapabilitiesNamingSystem() capabilities.SearchCapabilities
	SearchNamingSystem(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.NamingSystem, error)
}
type NutritionOrderSearch interface {
	SearchCapabilitiesNutritionOrder() capabilities.SearchCapabilities
	SearchNutritionOrder(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.NutritionOrder, error)
}
type ObservationSearch interface {
	SearchCapabilitiesObservation() capabilities.SearchCapabilities
	SearchObservation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Observation, error)
}
type ObservationDefinitionSearch interface {
	SearchCapabilitiesObservationDefinition() capabilities.SearchCapabilities
	SearchObservationDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ObservationDefinition, error)
}
type OperationDefinitionSearch interface {
	SearchCapabilitiesOperationDefinition() capabilities.SearchCapabilities
	SearchOperationDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.OperationDefinition, error)
}
type OperationOutcomeSearch interface {
	SearchCapabilitiesOperationOutcome() capabilities.SearchCapabilities
	SearchOperationOutcome(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.OperationOutcome, error)
}
type OrganizationSearch interface {
	SearchCapabilitiesOrganization() capabilities.SearchCapabilities
	SearchOrganization(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Organization, error)
}
type OrganizationAffiliationSearch interface {
	SearchCapabilitiesOrganizationAffiliation() capabilities.SearchCapabilities
	SearchOrganizationAffiliation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.OrganizationAffiliation, error)
}
type ParametersSearch interface {
	SearchCapabilitiesParameters() capabilities.SearchCapabilities
	SearchParameters(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Parameters, error)
}
type PatientSearch interface {
	SearchCapabilitiesPatient() capabilities.SearchCapabilities
	SearchPatient(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Patient, error)
}
type PaymentNoticeSearch interface {
	SearchCapabilitiesPaymentNotice() capabilities.SearchCapabilities
	SearchPaymentNotice(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.PaymentNotice, error)
}
type PaymentReconciliationSearch interface {
	SearchCapabilitiesPaymentReconciliation() capabilities.SearchCapabilities
	SearchPaymentReconciliation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.PaymentReconciliation, error)
}
type PersonSearch interface {
	SearchCapabilitiesPerson() capabilities.SearchCapabilities
	SearchPerson(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Person, error)
}
type PlanDefinitionSearch interface {
	SearchCapabilitiesPlanDefinition() capabilities.SearchCapabilities
	SearchPlanDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.PlanDefinition, error)
}
type PractitionerSearch interface {
	SearchCapabilitiesPractitioner() capabilities.SearchCapabilities
	SearchPractitioner(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Practitioner, error)
}
type PractitionerRoleSearch interface {
	SearchCapabilitiesPractitionerRole() capabilities.SearchCapabilities
	SearchPractitionerRole(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.PractitionerRole, error)
}
type ProcedureSearch interface {
	SearchCapabilitiesProcedure() capabilities.SearchCapabilities
	SearchProcedure(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Procedure, error)
}
type ProvenanceSearch interface {
	SearchCapabilitiesProvenance() capabilities.SearchCapabilities
	SearchProvenance(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Provenance, error)
}
type QuestionnaireSearch interface {
	SearchCapabilitiesQuestionnaire() capabilities.SearchCapabilities
	SearchQuestionnaire(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Questionnaire, error)
}
type QuestionnaireResponseSearch interface {
	SearchCapabilitiesQuestionnaireResponse() capabilities.SearchCapabilities
	SearchQuestionnaireResponse(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.QuestionnaireResponse, error)
}
type RelatedPersonSearch interface {
	SearchCapabilitiesRelatedPerson() capabilities.SearchCapabilities
	SearchRelatedPerson(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.RelatedPerson, error)
}
type RequestGroupSearch interface {
	SearchCapabilitiesRequestGroup() capabilities.SearchCapabilities
	SearchRequestGroup(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.RequestGroup, error)
}
type ResearchDefinitionSearch interface {
	SearchCapabilitiesResearchDefinition() capabilities.SearchCapabilities
	SearchResearchDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ResearchDefinition, error)
}
type ResearchElementDefinitionSearch interface {
	SearchCapabilitiesResearchElementDefinition() capabilities.SearchCapabilities
	SearchResearchElementDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ResearchElementDefinition, error)
}
type ResearchStudySearch interface {
	SearchCapabilitiesResearchStudy() capabilities.SearchCapabilities
	SearchResearchStudy(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ResearchStudy, error)
}
type ResearchSubjectSearch interface {
	SearchCapabilitiesResearchSubject() capabilities.SearchCapabilities
	SearchResearchSubject(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ResearchSubject, error)
}
type RiskAssessmentSearch interface {
	SearchCapabilitiesRiskAssessment() capabilities.SearchCapabilities
	SearchRiskAssessment(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.RiskAssessment, error)
}
type RiskEvidenceSynthesisSearch interface {
	SearchCapabilitiesRiskEvidenceSynthesis() capabilities.SearchCapabilities
	SearchRiskEvidenceSynthesis(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.RiskEvidenceSynthesis, error)
}
type ScheduleSearch interface {
	SearchCapabilitiesSchedule() capabilities.SearchCapabilities
	SearchSchedule(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Schedule, error)
}
type SearchParameterSearch interface {
	SearchCapabilitiesSearchParameter() capabilities.SearchCapabilities
	SearchSearchParameter(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SearchParameter, error)
}
type ServiceRequestSearch interface {
	SearchCapabilitiesServiceRequest() capabilities.SearchCapabilities
	SearchServiceRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ServiceRequest, error)
}
type SlotSearch interface {
	SearchCapabilitiesSlot() capabilities.SearchCapabilities
	SearchSlot(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Slot, error)
}
type SpecimenSearch interface {
	SearchCapabilitiesSpecimen() capabilities.SearchCapabilities
	SearchSpecimen(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Specimen, error)
}
type SpecimenDefinitionSearch interface {
	SearchCapabilitiesSpecimenDefinition() capabilities.SearchCapabilities
	SearchSpecimenDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SpecimenDefinition, error)
}
type StructureDefinitionSearch interface {
	SearchCapabilitiesStructureDefinition() capabilities.SearchCapabilities
	SearchStructureDefinition(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.StructureDefinition, error)
}
type StructureMapSearch interface {
	SearchCapabilitiesStructureMap() capabilities.SearchCapabilities
	SearchStructureMap(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.StructureMap, error)
}
type SubscriptionSearch interface {
	SearchCapabilitiesSubscription() capabilities.SearchCapabilities
	SearchSubscription(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Subscription, error)
}
type SubstanceSearch interface {
	SearchCapabilitiesSubstance() capabilities.SearchCapabilities
	SearchSubstance(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Substance, error)
}
type SubstanceNucleicAcidSearch interface {
	SearchCapabilitiesSubstanceNucleicAcid() capabilities.SearchCapabilities
	SearchSubstanceNucleicAcid(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstanceNucleicAcid, error)
}
type SubstancePolymerSearch interface {
	SearchCapabilitiesSubstancePolymer() capabilities.SearchCapabilities
	SearchSubstancePolymer(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstancePolymer, error)
}
type SubstanceProteinSearch interface {
	SearchCapabilitiesSubstanceProtein() capabilities.SearchCapabilities
	SearchSubstanceProtein(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstanceProtein, error)
}
type SubstanceReferenceInformationSearch interface {
	SearchCapabilitiesSubstanceReferenceInformation() capabilities.SearchCapabilities
	SearchSubstanceReferenceInformation(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstanceReferenceInformation, error)
}
type SubstanceSourceMaterialSearch interface {
	SearchCapabilitiesSubstanceSourceMaterial() capabilities.SearchCapabilities
	SearchSubstanceSourceMaterial(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstanceSourceMaterial, error)
}
type SubstanceSpecificationSearch interface {
	SearchCapabilitiesSubstanceSpecification() capabilities.SearchCapabilities
	SearchSubstanceSpecification(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SubstanceSpecification, error)
}
type SupplyDeliverySearch interface {
	SearchCapabilitiesSupplyDelivery() capabilities.SearchCapabilities
	SearchSupplyDelivery(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SupplyDelivery, error)
}
type SupplyRequestSearch interface {
	SearchCapabilitiesSupplyRequest() capabilities.SearchCapabilities
	SearchSupplyRequest(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.SupplyRequest, error)
}
type TaskSearch interface {
	SearchCapabilitiesTask() capabilities.SearchCapabilities
	SearchTask(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.Task, error)
}
type TerminologyCapabilitiesSearch interface {
	SearchCapabilitiesTerminologyCapabilities() capabilities.SearchCapabilities
	SearchTerminologyCapabilities(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.TerminologyCapabilities, error)
}
type TestReportSearch interface {
	SearchCapabilitiesTestReport() capabilities.SearchCapabilities
	SearchTestReport(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.TestReport, error)
}
type TestScriptSearch interface {
	SearchCapabilitiesTestScript() capabilities.SearchCapabilities
	SearchTestScript(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.TestScript, error)
}
type ValueSetSearch interface {
	SearchCapabilitiesValueSet() capabilities.SearchCapabilities
	SearchValueSet(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.ValueSet, error)
}
type VerificationResultSearch interface {
	SearchCapabilitiesVerificationResult() capabilities.SearchCapabilities
	SearchVerificationResult(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.VerificationResult, error)
}
type VisionPrescriptionSearch interface {
	SearchCapabilitiesVisionPrescription() capabilities.SearchCapabilities
	SearchVisionPrescription(ctx context.Context, parameters capabilities.SearchParameters) ([]r4.VisionPrescription, error)
}
