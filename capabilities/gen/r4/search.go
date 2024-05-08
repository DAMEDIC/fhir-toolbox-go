package r4

import (
	"context"
	model "fhir-toolbox/model/gen/r4"
)

type AccountSearch interface {
	SearchParamsAccount() []string
	SearchAccount(ctx context.Context, parameters map[string]string) ([]model.Account, error)
}
type ActivityDefinitionSearch interface {
	SearchParamsActivityDefinition() []string
	SearchActivityDefinition(ctx context.Context, parameters map[string]string) ([]model.ActivityDefinition, error)
}
type AdverseEventSearch interface {
	SearchParamsAdverseEvent() []string
	SearchAdverseEvent(ctx context.Context, parameters map[string]string) ([]model.AdverseEvent, error)
}
type AllergyIntoleranceSearch interface {
	SearchParamsAllergyIntolerance() []string
	SearchAllergyIntolerance(ctx context.Context, parameters map[string]string) ([]model.AllergyIntolerance, error)
}
type AppointmentSearch interface {
	SearchParamsAppointment() []string
	SearchAppointment(ctx context.Context, parameters map[string]string) ([]model.Appointment, error)
}
type AppointmentResponseSearch interface {
	SearchParamsAppointmentResponse() []string
	SearchAppointmentResponse(ctx context.Context, parameters map[string]string) ([]model.AppointmentResponse, error)
}
type AuditEventSearch interface {
	SearchParamsAuditEvent() []string
	SearchAuditEvent(ctx context.Context, parameters map[string]string) ([]model.AuditEvent, error)
}
type BasicSearch interface {
	SearchParamsBasic() []string
	SearchBasic(ctx context.Context, parameters map[string]string) ([]model.Basic, error)
}
type BinarySearch interface {
	SearchParamsBinary() []string
	SearchBinary(ctx context.Context, parameters map[string]string) ([]model.Binary, error)
}
type BiologicallyDerivedProductSearch interface {
	SearchParamsBiologicallyDerivedProduct() []string
	SearchBiologicallyDerivedProduct(ctx context.Context, parameters map[string]string) ([]model.BiologicallyDerivedProduct, error)
}
type BodyStructureSearch interface {
	SearchParamsBodyStructure() []string
	SearchBodyStructure(ctx context.Context, parameters map[string]string) ([]model.BodyStructure, error)
}
type BundleSearch interface {
	SearchParamsBundle() []string
	SearchBundle(ctx context.Context, parameters map[string]string) ([]model.Bundle, error)
}
type CapabilityStatementSearch interface {
	SearchParamsCapabilityStatement() []string
	SearchCapabilityStatement(ctx context.Context, parameters map[string]string) ([]model.CapabilityStatement, error)
}
type CarePlanSearch interface {
	SearchParamsCarePlan() []string
	SearchCarePlan(ctx context.Context, parameters map[string]string) ([]model.CarePlan, error)
}
type CareTeamSearch interface {
	SearchParamsCareTeam() []string
	SearchCareTeam(ctx context.Context, parameters map[string]string) ([]model.CareTeam, error)
}
type CatalogEntrySearch interface {
	SearchParamsCatalogEntry() []string
	SearchCatalogEntry(ctx context.Context, parameters map[string]string) ([]model.CatalogEntry, error)
}
type ChargeItemSearch interface {
	SearchParamsChargeItem() []string
	SearchChargeItem(ctx context.Context, parameters map[string]string) ([]model.ChargeItem, error)
}
type ChargeItemDefinitionSearch interface {
	SearchParamsChargeItemDefinition() []string
	SearchChargeItemDefinition(ctx context.Context, parameters map[string]string) ([]model.ChargeItemDefinition, error)
}
type ClaimSearch interface {
	SearchParamsClaim() []string
	SearchClaim(ctx context.Context, parameters map[string]string) ([]model.Claim, error)
}
type ClaimResponseSearch interface {
	SearchParamsClaimResponse() []string
	SearchClaimResponse(ctx context.Context, parameters map[string]string) ([]model.ClaimResponse, error)
}
type ClinicalImpressionSearch interface {
	SearchParamsClinicalImpression() []string
	SearchClinicalImpression(ctx context.Context, parameters map[string]string) ([]model.ClinicalImpression, error)
}
type CodeSystemSearch interface {
	SearchParamsCodeSystem() []string
	SearchCodeSystem(ctx context.Context, parameters map[string]string) ([]model.CodeSystem, error)
}
type CommunicationSearch interface {
	SearchParamsCommunication() []string
	SearchCommunication(ctx context.Context, parameters map[string]string) ([]model.Communication, error)
}
type CommunicationRequestSearch interface {
	SearchParamsCommunicationRequest() []string
	SearchCommunicationRequest(ctx context.Context, parameters map[string]string) ([]model.CommunicationRequest, error)
}
type CompartmentDefinitionSearch interface {
	SearchParamsCompartmentDefinition() []string
	SearchCompartmentDefinition(ctx context.Context, parameters map[string]string) ([]model.CompartmentDefinition, error)
}
type CompositionSearch interface {
	SearchParamsComposition() []string
	SearchComposition(ctx context.Context, parameters map[string]string) ([]model.Composition, error)
}
type ConceptMapSearch interface {
	SearchParamsConceptMap() []string
	SearchConceptMap(ctx context.Context, parameters map[string]string) ([]model.ConceptMap, error)
}
type ConditionSearch interface {
	SearchParamsCondition() []string
	SearchCondition(ctx context.Context, parameters map[string]string) ([]model.Condition, error)
}
type ConsentSearch interface {
	SearchParamsConsent() []string
	SearchConsent(ctx context.Context, parameters map[string]string) ([]model.Consent, error)
}
type ContractSearch interface {
	SearchParamsContract() []string
	SearchContract(ctx context.Context, parameters map[string]string) ([]model.Contract, error)
}
type CoverageSearch interface {
	SearchParamsCoverage() []string
	SearchCoverage(ctx context.Context, parameters map[string]string) ([]model.Coverage, error)
}
type CoverageEligibilityRequestSearch interface {
	SearchParamsCoverageEligibilityRequest() []string
	SearchCoverageEligibilityRequest(ctx context.Context, parameters map[string]string) ([]model.CoverageEligibilityRequest, error)
}
type CoverageEligibilityResponseSearch interface {
	SearchParamsCoverageEligibilityResponse() []string
	SearchCoverageEligibilityResponse(ctx context.Context, parameters map[string]string) ([]model.CoverageEligibilityResponse, error)
}
type DetectedIssueSearch interface {
	SearchParamsDetectedIssue() []string
	SearchDetectedIssue(ctx context.Context, parameters map[string]string) ([]model.DetectedIssue, error)
}
type DeviceSearch interface {
	SearchParamsDevice() []string
	SearchDevice(ctx context.Context, parameters map[string]string) ([]model.Device, error)
}
type DeviceDefinitionSearch interface {
	SearchParamsDeviceDefinition() []string
	SearchDeviceDefinition(ctx context.Context, parameters map[string]string) ([]model.DeviceDefinition, error)
}
type DeviceMetricSearch interface {
	SearchParamsDeviceMetric() []string
	SearchDeviceMetric(ctx context.Context, parameters map[string]string) ([]model.DeviceMetric, error)
}
type DeviceRequestSearch interface {
	SearchParamsDeviceRequest() []string
	SearchDeviceRequest(ctx context.Context, parameters map[string]string) ([]model.DeviceRequest, error)
}
type DeviceUseStatementSearch interface {
	SearchParamsDeviceUseStatement() []string
	SearchDeviceUseStatement(ctx context.Context, parameters map[string]string) ([]model.DeviceUseStatement, error)
}
type DiagnosticReportSearch interface {
	SearchParamsDiagnosticReport() []string
	SearchDiagnosticReport(ctx context.Context, parameters map[string]string) ([]model.DiagnosticReport, error)
}
type DocumentManifestSearch interface {
	SearchParamsDocumentManifest() []string
	SearchDocumentManifest(ctx context.Context, parameters map[string]string) ([]model.DocumentManifest, error)
}
type DocumentReferenceSearch interface {
	SearchParamsDocumentReference() []string
	SearchDocumentReference(ctx context.Context, parameters map[string]string) ([]model.DocumentReference, error)
}
type EffectEvidenceSynthesisSearch interface {
	SearchParamsEffectEvidenceSynthesis() []string
	SearchEffectEvidenceSynthesis(ctx context.Context, parameters map[string]string) ([]model.EffectEvidenceSynthesis, error)
}
type EncounterSearch interface {
	SearchParamsEncounter() []string
	SearchEncounter(ctx context.Context, parameters map[string]string) ([]model.Encounter, error)
}
type EndpointSearch interface {
	SearchParamsEndpoint() []string
	SearchEndpoint(ctx context.Context, parameters map[string]string) ([]model.Endpoint, error)
}
type EnrollmentRequestSearch interface {
	SearchParamsEnrollmentRequest() []string
	SearchEnrollmentRequest(ctx context.Context, parameters map[string]string) ([]model.EnrollmentRequest, error)
}
type EnrollmentResponseSearch interface {
	SearchParamsEnrollmentResponse() []string
	SearchEnrollmentResponse(ctx context.Context, parameters map[string]string) ([]model.EnrollmentResponse, error)
}
type EpisodeOfCareSearch interface {
	SearchParamsEpisodeOfCare() []string
	SearchEpisodeOfCare(ctx context.Context, parameters map[string]string) ([]model.EpisodeOfCare, error)
}
type EventDefinitionSearch interface {
	SearchParamsEventDefinition() []string
	SearchEventDefinition(ctx context.Context, parameters map[string]string) ([]model.EventDefinition, error)
}
type EvidenceSearch interface {
	SearchParamsEvidence() []string
	SearchEvidence(ctx context.Context, parameters map[string]string) ([]model.Evidence, error)
}
type EvidenceVariableSearch interface {
	SearchParamsEvidenceVariable() []string
	SearchEvidenceVariable(ctx context.Context, parameters map[string]string) ([]model.EvidenceVariable, error)
}
type ExampleScenarioSearch interface {
	SearchParamsExampleScenario() []string
	SearchExampleScenario(ctx context.Context, parameters map[string]string) ([]model.ExampleScenario, error)
}
type ExplanationOfBenefitSearch interface {
	SearchParamsExplanationOfBenefit() []string
	SearchExplanationOfBenefit(ctx context.Context, parameters map[string]string) ([]model.ExplanationOfBenefit, error)
}
type FamilyMemberHistorySearch interface {
	SearchParamsFamilyMemberHistory() []string
	SearchFamilyMemberHistory(ctx context.Context, parameters map[string]string) ([]model.FamilyMemberHistory, error)
}
type FlagSearch interface {
	SearchParamsFlag() []string
	SearchFlag(ctx context.Context, parameters map[string]string) ([]model.Flag, error)
}
type GoalSearch interface {
	SearchParamsGoal() []string
	SearchGoal(ctx context.Context, parameters map[string]string) ([]model.Goal, error)
}
type GraphDefinitionSearch interface {
	SearchParamsGraphDefinition() []string
	SearchGraphDefinition(ctx context.Context, parameters map[string]string) ([]model.GraphDefinition, error)
}
type GroupSearch interface {
	SearchParamsGroup() []string
	SearchGroup(ctx context.Context, parameters map[string]string) ([]model.Group, error)
}
type GuidanceResponseSearch interface {
	SearchParamsGuidanceResponse() []string
	SearchGuidanceResponse(ctx context.Context, parameters map[string]string) ([]model.GuidanceResponse, error)
}
type HealthcareServiceSearch interface {
	SearchParamsHealthcareService() []string
	SearchHealthcareService(ctx context.Context, parameters map[string]string) ([]model.HealthcareService, error)
}
type ImagingStudySearch interface {
	SearchParamsImagingStudy() []string
	SearchImagingStudy(ctx context.Context, parameters map[string]string) ([]model.ImagingStudy, error)
}
type ImmunizationSearch interface {
	SearchParamsImmunization() []string
	SearchImmunization(ctx context.Context, parameters map[string]string) ([]model.Immunization, error)
}
type ImmunizationEvaluationSearch interface {
	SearchParamsImmunizationEvaluation() []string
	SearchImmunizationEvaluation(ctx context.Context, parameters map[string]string) ([]model.ImmunizationEvaluation, error)
}
type ImmunizationRecommendationSearch interface {
	SearchParamsImmunizationRecommendation() []string
	SearchImmunizationRecommendation(ctx context.Context, parameters map[string]string) ([]model.ImmunizationRecommendation, error)
}
type ImplementationGuideSearch interface {
	SearchParamsImplementationGuide() []string
	SearchImplementationGuide(ctx context.Context, parameters map[string]string) ([]model.ImplementationGuide, error)
}
type InsurancePlanSearch interface {
	SearchParamsInsurancePlan() []string
	SearchInsurancePlan(ctx context.Context, parameters map[string]string) ([]model.InsurancePlan, error)
}
type InvoiceSearch interface {
	SearchParamsInvoice() []string
	SearchInvoice(ctx context.Context, parameters map[string]string) ([]model.Invoice, error)
}
type LibrarySearch interface {
	SearchParamsLibrary() []string
	SearchLibrary(ctx context.Context, parameters map[string]string) ([]model.Library, error)
}
type LinkageSearch interface {
	SearchParamsLinkage() []string
	SearchLinkage(ctx context.Context, parameters map[string]string) ([]model.Linkage, error)
}
type ListSearch interface {
	SearchParamsList() []string
	SearchList(ctx context.Context, parameters map[string]string) ([]model.List, error)
}
type LocationSearch interface {
	SearchParamsLocation() []string
	SearchLocation(ctx context.Context, parameters map[string]string) ([]model.Location, error)
}
type MeasureSearch interface {
	SearchParamsMeasure() []string
	SearchMeasure(ctx context.Context, parameters map[string]string) ([]model.Measure, error)
}
type MeasureReportSearch interface {
	SearchParamsMeasureReport() []string
	SearchMeasureReport(ctx context.Context, parameters map[string]string) ([]model.MeasureReport, error)
}
type MediaSearch interface {
	SearchParamsMedia() []string
	SearchMedia(ctx context.Context, parameters map[string]string) ([]model.Media, error)
}
type MedicationSearch interface {
	SearchParamsMedication() []string
	SearchMedication(ctx context.Context, parameters map[string]string) ([]model.Medication, error)
}
type MedicationAdministrationSearch interface {
	SearchParamsMedicationAdministration() []string
	SearchMedicationAdministration(ctx context.Context, parameters map[string]string) ([]model.MedicationAdministration, error)
}
type MedicationDispenseSearch interface {
	SearchParamsMedicationDispense() []string
	SearchMedicationDispense(ctx context.Context, parameters map[string]string) ([]model.MedicationDispense, error)
}
type MedicationKnowledgeSearch interface {
	SearchParamsMedicationKnowledge() []string
	SearchMedicationKnowledge(ctx context.Context, parameters map[string]string) ([]model.MedicationKnowledge, error)
}
type MedicationRequestSearch interface {
	SearchParamsMedicationRequest() []string
	SearchMedicationRequest(ctx context.Context, parameters map[string]string) ([]model.MedicationRequest, error)
}
type MedicationStatementSearch interface {
	SearchParamsMedicationStatement() []string
	SearchMedicationStatement(ctx context.Context, parameters map[string]string) ([]model.MedicationStatement, error)
}
type MedicinalProductSearch interface {
	SearchParamsMedicinalProduct() []string
	SearchMedicinalProduct(ctx context.Context, parameters map[string]string) ([]model.MedicinalProduct, error)
}
type MedicinalProductAuthorizationSearch interface {
	SearchParamsMedicinalProductAuthorization() []string
	SearchMedicinalProductAuthorization(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductAuthorization, error)
}
type MedicinalProductContraindicationSearch interface {
	SearchParamsMedicinalProductContraindication() []string
	SearchMedicinalProductContraindication(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductContraindication, error)
}
type MedicinalProductIndicationSearch interface {
	SearchParamsMedicinalProductIndication() []string
	SearchMedicinalProductIndication(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductIndication, error)
}
type MedicinalProductIngredientSearch interface {
	SearchParamsMedicinalProductIngredient() []string
	SearchMedicinalProductIngredient(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductIngredient, error)
}
type MedicinalProductInteractionSearch interface {
	SearchParamsMedicinalProductInteraction() []string
	SearchMedicinalProductInteraction(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductInteraction, error)
}
type MedicinalProductManufacturedSearch interface {
	SearchParamsMedicinalProductManufactured() []string
	SearchMedicinalProductManufactured(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductManufactured, error)
}
type MedicinalProductPackagedSearch interface {
	SearchParamsMedicinalProductPackaged() []string
	SearchMedicinalProductPackaged(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductPackaged, error)
}
type MedicinalProductPharmaceuticalSearch interface {
	SearchParamsMedicinalProductPharmaceutical() []string
	SearchMedicinalProductPharmaceutical(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductPharmaceutical, error)
}
type MedicinalProductUndesirableEffectSearch interface {
	SearchParamsMedicinalProductUndesirableEffect() []string
	SearchMedicinalProductUndesirableEffect(ctx context.Context, parameters map[string]string) ([]model.MedicinalProductUndesirableEffect, error)
}
type MessageDefinitionSearch interface {
	SearchParamsMessageDefinition() []string
	SearchMessageDefinition(ctx context.Context, parameters map[string]string) ([]model.MessageDefinition, error)
}
type MessageHeaderSearch interface {
	SearchParamsMessageHeader() []string
	SearchMessageHeader(ctx context.Context, parameters map[string]string) ([]model.MessageHeader, error)
}
type MolecularSequenceSearch interface {
	SearchParamsMolecularSequence() []string
	SearchMolecularSequence(ctx context.Context, parameters map[string]string) ([]model.MolecularSequence, error)
}
type NamingSystemSearch interface {
	SearchParamsNamingSystem() []string
	SearchNamingSystem(ctx context.Context, parameters map[string]string) ([]model.NamingSystem, error)
}
type NutritionOrderSearch interface {
	SearchParamsNutritionOrder() []string
	SearchNutritionOrder(ctx context.Context, parameters map[string]string) ([]model.NutritionOrder, error)
}
type ObservationSearch interface {
	SearchParamsObservation() []string
	SearchObservation(ctx context.Context, parameters map[string]string) ([]model.Observation, error)
}
type ObservationDefinitionSearch interface {
	SearchParamsObservationDefinition() []string
	SearchObservationDefinition(ctx context.Context, parameters map[string]string) ([]model.ObservationDefinition, error)
}
type OperationDefinitionSearch interface {
	SearchParamsOperationDefinition() []string
	SearchOperationDefinition(ctx context.Context, parameters map[string]string) ([]model.OperationDefinition, error)
}
type OperationOutcomeSearch interface {
	SearchParamsOperationOutcome() []string
	SearchOperationOutcome(ctx context.Context, parameters map[string]string) ([]model.OperationOutcome, error)
}
type OrganizationSearch interface {
	SearchParamsOrganization() []string
	SearchOrganization(ctx context.Context, parameters map[string]string) ([]model.Organization, error)
}
type OrganizationAffiliationSearch interface {
	SearchParamsOrganizationAffiliation() []string
	SearchOrganizationAffiliation(ctx context.Context, parameters map[string]string) ([]model.OrganizationAffiliation, error)
}
type ParametersSearch interface {
	SearchParamsParameters() []string
	SearchParameters(ctx context.Context, parameters map[string]string) ([]model.Parameters, error)
}
type PatientSearch interface {
	SearchParamsPatient() []string
	SearchPatient(ctx context.Context, parameters map[string]string) ([]model.Patient, error)
}
type PaymentNoticeSearch interface {
	SearchParamsPaymentNotice() []string
	SearchPaymentNotice(ctx context.Context, parameters map[string]string) ([]model.PaymentNotice, error)
}
type PaymentReconciliationSearch interface {
	SearchParamsPaymentReconciliation() []string
	SearchPaymentReconciliation(ctx context.Context, parameters map[string]string) ([]model.PaymentReconciliation, error)
}
type PersonSearch interface {
	SearchParamsPerson() []string
	SearchPerson(ctx context.Context, parameters map[string]string) ([]model.Person, error)
}
type PlanDefinitionSearch interface {
	SearchParamsPlanDefinition() []string
	SearchPlanDefinition(ctx context.Context, parameters map[string]string) ([]model.PlanDefinition, error)
}
type PractitionerSearch interface {
	SearchParamsPractitioner() []string
	SearchPractitioner(ctx context.Context, parameters map[string]string) ([]model.Practitioner, error)
}
type PractitionerRoleSearch interface {
	SearchParamsPractitionerRole() []string
	SearchPractitionerRole(ctx context.Context, parameters map[string]string) ([]model.PractitionerRole, error)
}
type ProcedureSearch interface {
	SearchParamsProcedure() []string
	SearchProcedure(ctx context.Context, parameters map[string]string) ([]model.Procedure, error)
}
type ProvenanceSearch interface {
	SearchParamsProvenance() []string
	SearchProvenance(ctx context.Context, parameters map[string]string) ([]model.Provenance, error)
}
type QuestionnaireSearch interface {
	SearchParamsQuestionnaire() []string
	SearchQuestionnaire(ctx context.Context, parameters map[string]string) ([]model.Questionnaire, error)
}
type QuestionnaireResponseSearch interface {
	SearchParamsQuestionnaireResponse() []string
	SearchQuestionnaireResponse(ctx context.Context, parameters map[string]string) ([]model.QuestionnaireResponse, error)
}
type RelatedPersonSearch interface {
	SearchParamsRelatedPerson() []string
	SearchRelatedPerson(ctx context.Context, parameters map[string]string) ([]model.RelatedPerson, error)
}
type RequestGroupSearch interface {
	SearchParamsRequestGroup() []string
	SearchRequestGroup(ctx context.Context, parameters map[string]string) ([]model.RequestGroup, error)
}
type ResearchDefinitionSearch interface {
	SearchParamsResearchDefinition() []string
	SearchResearchDefinition(ctx context.Context, parameters map[string]string) ([]model.ResearchDefinition, error)
}
type ResearchElementDefinitionSearch interface {
	SearchParamsResearchElementDefinition() []string
	SearchResearchElementDefinition(ctx context.Context, parameters map[string]string) ([]model.ResearchElementDefinition, error)
}
type ResearchStudySearch interface {
	SearchParamsResearchStudy() []string
	SearchResearchStudy(ctx context.Context, parameters map[string]string) ([]model.ResearchStudy, error)
}
type ResearchSubjectSearch interface {
	SearchParamsResearchSubject() []string
	SearchResearchSubject(ctx context.Context, parameters map[string]string) ([]model.ResearchSubject, error)
}
type RiskAssessmentSearch interface {
	SearchParamsRiskAssessment() []string
	SearchRiskAssessment(ctx context.Context, parameters map[string]string) ([]model.RiskAssessment, error)
}
type RiskEvidenceSynthesisSearch interface {
	SearchParamsRiskEvidenceSynthesis() []string
	SearchRiskEvidenceSynthesis(ctx context.Context, parameters map[string]string) ([]model.RiskEvidenceSynthesis, error)
}
type ScheduleSearch interface {
	SearchParamsSchedule() []string
	SearchSchedule(ctx context.Context, parameters map[string]string) ([]model.Schedule, error)
}
type SearchParameterSearch interface {
	SearchParamsSearchParameter() []string
	SearchSearchParameter(ctx context.Context, parameters map[string]string) ([]model.SearchParameter, error)
}
type ServiceRequestSearch interface {
	SearchParamsServiceRequest() []string
	SearchServiceRequest(ctx context.Context, parameters map[string]string) ([]model.ServiceRequest, error)
}
type SlotSearch interface {
	SearchParamsSlot() []string
	SearchSlot(ctx context.Context, parameters map[string]string) ([]model.Slot, error)
}
type SpecimenSearch interface {
	SearchParamsSpecimen() []string
	SearchSpecimen(ctx context.Context, parameters map[string]string) ([]model.Specimen, error)
}
type SpecimenDefinitionSearch interface {
	SearchParamsSpecimenDefinition() []string
	SearchSpecimenDefinition(ctx context.Context, parameters map[string]string) ([]model.SpecimenDefinition, error)
}
type StructureDefinitionSearch interface {
	SearchParamsStructureDefinition() []string
	SearchStructureDefinition(ctx context.Context, parameters map[string]string) ([]model.StructureDefinition, error)
}
type StructureMapSearch interface {
	SearchParamsStructureMap() []string
	SearchStructureMap(ctx context.Context, parameters map[string]string) ([]model.StructureMap, error)
}
type SubscriptionSearch interface {
	SearchParamsSubscription() []string
	SearchSubscription(ctx context.Context, parameters map[string]string) ([]model.Subscription, error)
}
type SubstanceSearch interface {
	SearchParamsSubstance() []string
	SearchSubstance(ctx context.Context, parameters map[string]string) ([]model.Substance, error)
}
type SubstanceNucleicAcidSearch interface {
	SearchParamsSubstanceNucleicAcid() []string
	SearchSubstanceNucleicAcid(ctx context.Context, parameters map[string]string) ([]model.SubstanceNucleicAcid, error)
}
type SubstancePolymerSearch interface {
	SearchParamsSubstancePolymer() []string
	SearchSubstancePolymer(ctx context.Context, parameters map[string]string) ([]model.SubstancePolymer, error)
}
type SubstanceProteinSearch interface {
	SearchParamsSubstanceProtein() []string
	SearchSubstanceProtein(ctx context.Context, parameters map[string]string) ([]model.SubstanceProtein, error)
}
type SubstanceReferenceInformationSearch interface {
	SearchParamsSubstanceReferenceInformation() []string
	SearchSubstanceReferenceInformation(ctx context.Context, parameters map[string]string) ([]model.SubstanceReferenceInformation, error)
}
type SubstanceSourceMaterialSearch interface {
	SearchParamsSubstanceSourceMaterial() []string
	SearchSubstanceSourceMaterial(ctx context.Context, parameters map[string]string) ([]model.SubstanceSourceMaterial, error)
}
type SubstanceSpecificationSearch interface {
	SearchParamsSubstanceSpecification() []string
	SearchSubstanceSpecification(ctx context.Context, parameters map[string]string) ([]model.SubstanceSpecification, error)
}
type SupplyDeliverySearch interface {
	SearchParamsSupplyDelivery() []string
	SearchSupplyDelivery(ctx context.Context, parameters map[string]string) ([]model.SupplyDelivery, error)
}
type SupplyRequestSearch interface {
	SearchParamsSupplyRequest() []string
	SearchSupplyRequest(ctx context.Context, parameters map[string]string) ([]model.SupplyRequest, error)
}
type TaskSearch interface {
	SearchParamsTask() []string
	SearchTask(ctx context.Context, parameters map[string]string) ([]model.Task, error)
}
type TerminologyCapabilitiesSearch interface {
	SearchParamsTerminologyCapabilities() []string
	SearchTerminologyCapabilities(ctx context.Context, parameters map[string]string) ([]model.TerminologyCapabilities, error)
}
type TestReportSearch interface {
	SearchParamsTestReport() []string
	SearchTestReport(ctx context.Context, parameters map[string]string) ([]model.TestReport, error)
}
type TestScriptSearch interface {
	SearchParamsTestScript() []string
	SearchTestScript(ctx context.Context, parameters map[string]string) ([]model.TestScript, error)
}
type ValueSetSearch interface {
	SearchParamsValueSet() []string
	SearchValueSet(ctx context.Context, parameters map[string]string) ([]model.ValueSet, error)
}
type VerificationResultSearch interface {
	SearchParamsVerificationResult() []string
	SearchVerificationResult(ctx context.Context, parameters map[string]string) ([]model.VerificationResult, error)
}
type VisionPrescriptionSearch interface {
	SearchParamsVisionPrescription() []string
	SearchVisionPrescription(ctx context.Context, parameters map[string]string) ([]model.VisionPrescription, error)
}
