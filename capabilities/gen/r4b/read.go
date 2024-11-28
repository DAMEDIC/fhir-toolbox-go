package capabilitiesR4B

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4b "fhir-toolbox/model/gen/r4b"
)

type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (r4b.Account, capabilities.FHIRError)
}
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (r4b.ActivityDefinition, capabilities.FHIRError)
}
type AdministrableProductDefinitionRead interface {
	ReadAdministrableProductDefinition(ctx context.Context, id string) (r4b.AdministrableProductDefinition, capabilities.FHIRError)
}
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (r4b.AdverseEvent, capabilities.FHIRError)
}
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (r4b.AllergyIntolerance, capabilities.FHIRError)
}
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (r4b.Appointment, capabilities.FHIRError)
}
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (r4b.AppointmentResponse, capabilities.FHIRError)
}
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (r4b.AuditEvent, capabilities.FHIRError)
}
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (r4b.Basic, capabilities.FHIRError)
}
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (r4b.Binary, capabilities.FHIRError)
}
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4b.BiologicallyDerivedProduct, capabilities.FHIRError)
}
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (r4b.BodyStructure, capabilities.FHIRError)
}
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (r4b.Bundle, capabilities.FHIRError)
}
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (r4b.CapabilityStatement, capabilities.FHIRError)
}
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (r4b.CarePlan, capabilities.FHIRError)
}
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (r4b.CareTeam, capabilities.FHIRError)
}
type CatalogEntryRead interface {
	ReadCatalogEntry(ctx context.Context, id string) (r4b.CatalogEntry, capabilities.FHIRError)
}
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (r4b.ChargeItem, capabilities.FHIRError)
}
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (r4b.ChargeItemDefinition, capabilities.FHIRError)
}
type CitationRead interface {
	ReadCitation(ctx context.Context, id string) (r4b.Citation, capabilities.FHIRError)
}
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (r4b.Claim, capabilities.FHIRError)
}
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (r4b.ClaimResponse, capabilities.FHIRError)
}
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (r4b.ClinicalImpression, capabilities.FHIRError)
}
type ClinicalUseDefinitionRead interface {
	ReadClinicalUseDefinition(ctx context.Context, id string) (r4b.ClinicalUseDefinition, capabilities.FHIRError)
}
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (r4b.CodeSystem, capabilities.FHIRError)
}
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (r4b.Communication, capabilities.FHIRError)
}
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (r4b.CommunicationRequest, capabilities.FHIRError)
}
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (r4b.CompartmentDefinition, capabilities.FHIRError)
}
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (r4b.Composition, capabilities.FHIRError)
}
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (r4b.ConceptMap, capabilities.FHIRError)
}
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (r4b.Condition, capabilities.FHIRError)
}
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (r4b.Consent, capabilities.FHIRError)
}
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (r4b.Contract, capabilities.FHIRError)
}
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (r4b.Coverage, capabilities.FHIRError)
}
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4b.CoverageEligibilityRequest, capabilities.FHIRError)
}
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4b.CoverageEligibilityResponse, capabilities.FHIRError)
}
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (r4b.DetectedIssue, capabilities.FHIRError)
}
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (r4b.Device, capabilities.FHIRError)
}
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (r4b.DeviceDefinition, capabilities.FHIRError)
}
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (r4b.DeviceMetric, capabilities.FHIRError)
}
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (r4b.DeviceRequest, capabilities.FHIRError)
}
type DeviceUseStatementRead interface {
	ReadDeviceUseStatement(ctx context.Context, id string) (r4b.DeviceUseStatement, capabilities.FHIRError)
}
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (r4b.DiagnosticReport, capabilities.FHIRError)
}
type DocumentManifestRead interface {
	ReadDocumentManifest(ctx context.Context, id string) (r4b.DocumentManifest, capabilities.FHIRError)
}
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (r4b.DocumentReference, capabilities.FHIRError)
}
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (r4b.Encounter, capabilities.FHIRError)
}
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (r4b.Endpoint, capabilities.FHIRError)
}
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (r4b.EnrollmentRequest, capabilities.FHIRError)
}
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (r4b.EnrollmentResponse, capabilities.FHIRError)
}
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (r4b.EpisodeOfCare, capabilities.FHIRError)
}
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (r4b.EventDefinition, capabilities.FHIRError)
}
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (r4b.Evidence, capabilities.FHIRError)
}
type EvidenceReportRead interface {
	ReadEvidenceReport(ctx context.Context, id string) (r4b.EvidenceReport, capabilities.FHIRError)
}
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (r4b.EvidenceVariable, capabilities.FHIRError)
}
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (r4b.ExampleScenario, capabilities.FHIRError)
}
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (r4b.ExplanationOfBenefit, capabilities.FHIRError)
}
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (r4b.FamilyMemberHistory, capabilities.FHIRError)
}
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (r4b.Flag, capabilities.FHIRError)
}
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (r4b.Goal, capabilities.FHIRError)
}
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (r4b.GraphDefinition, capabilities.FHIRError)
}
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (r4b.Group, capabilities.FHIRError)
}
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (r4b.GuidanceResponse, capabilities.FHIRError)
}
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (r4b.HealthcareService, capabilities.FHIRError)
}
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (r4b.ImagingStudy, capabilities.FHIRError)
}
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (r4b.Immunization, capabilities.FHIRError)
}
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (r4b.ImmunizationEvaluation, capabilities.FHIRError)
}
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (r4b.ImmunizationRecommendation, capabilities.FHIRError)
}
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (r4b.ImplementationGuide, capabilities.FHIRError)
}
type IngredientRead interface {
	ReadIngredient(ctx context.Context, id string) (r4b.Ingredient, capabilities.FHIRError)
}
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (r4b.InsurancePlan, capabilities.FHIRError)
}
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (r4b.Invoice, capabilities.FHIRError)
}
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (r4b.Library, capabilities.FHIRError)
}
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (r4b.Linkage, capabilities.FHIRError)
}
type ListRead interface {
	ReadList(ctx context.Context, id string) (r4b.List, capabilities.FHIRError)
}
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (r4b.Location, capabilities.FHIRError)
}
type ManufacturedItemDefinitionRead interface {
	ReadManufacturedItemDefinition(ctx context.Context, id string) (r4b.ManufacturedItemDefinition, capabilities.FHIRError)
}
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (r4b.Measure, capabilities.FHIRError)
}
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (r4b.MeasureReport, capabilities.FHIRError)
}
type MediaRead interface {
	ReadMedia(ctx context.Context, id string) (r4b.Media, capabilities.FHIRError)
}
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (r4b.Medication, capabilities.FHIRError)
}
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (r4b.MedicationAdministration, capabilities.FHIRError)
}
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (r4b.MedicationDispense, capabilities.FHIRError)
}
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (r4b.MedicationKnowledge, capabilities.FHIRError)
}
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (r4b.MedicationRequest, capabilities.FHIRError)
}
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (r4b.MedicationStatement, capabilities.FHIRError)
}
type MedicinalProductDefinitionRead interface {
	ReadMedicinalProductDefinition(ctx context.Context, id string) (r4b.MedicinalProductDefinition, capabilities.FHIRError)
}
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (r4b.MessageDefinition, capabilities.FHIRError)
}
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (r4b.MessageHeader, capabilities.FHIRError)
}
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (r4b.MolecularSequence, capabilities.FHIRError)
}
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (r4b.NamingSystem, capabilities.FHIRError)
}
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (r4b.NutritionOrder, capabilities.FHIRError)
}
type NutritionProductRead interface {
	ReadNutritionProduct(ctx context.Context, id string) (r4b.NutritionProduct, capabilities.FHIRError)
}
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (r4b.Observation, capabilities.FHIRError)
}
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (r4b.ObservationDefinition, capabilities.FHIRError)
}
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (r4b.OperationDefinition, capabilities.FHIRError)
}
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (r4b.OperationOutcome, capabilities.FHIRError)
}
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (r4b.Organization, capabilities.FHIRError)
}
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (r4b.OrganizationAffiliation, capabilities.FHIRError)
}
type PackagedProductDefinitionRead interface {
	ReadPackagedProductDefinition(ctx context.Context, id string) (r4b.PackagedProductDefinition, capabilities.FHIRError)
}
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (r4b.Parameters, capabilities.FHIRError)
}
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (r4b.Patient, capabilities.FHIRError)
}
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (r4b.PaymentNotice, capabilities.FHIRError)
}
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (r4b.PaymentReconciliation, capabilities.FHIRError)
}
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (r4b.Person, capabilities.FHIRError)
}
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (r4b.PlanDefinition, capabilities.FHIRError)
}
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (r4b.Practitioner, capabilities.FHIRError)
}
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (r4b.PractitionerRole, capabilities.FHIRError)
}
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (r4b.Procedure, capabilities.FHIRError)
}
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (r4b.Provenance, capabilities.FHIRError)
}
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (r4b.Questionnaire, capabilities.FHIRError)
}
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (r4b.QuestionnaireResponse, capabilities.FHIRError)
}
type RegulatedAuthorizationRead interface {
	ReadRegulatedAuthorization(ctx context.Context, id string) (r4b.RegulatedAuthorization, capabilities.FHIRError)
}
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (r4b.RelatedPerson, capabilities.FHIRError)
}
type RequestGroupRead interface {
	ReadRequestGroup(ctx context.Context, id string) (r4b.RequestGroup, capabilities.FHIRError)
}
type ResearchDefinitionRead interface {
	ReadResearchDefinition(ctx context.Context, id string) (r4b.ResearchDefinition, capabilities.FHIRError)
}
type ResearchElementDefinitionRead interface {
	ReadResearchElementDefinition(ctx context.Context, id string) (r4b.ResearchElementDefinition, capabilities.FHIRError)
}
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (r4b.ResearchStudy, capabilities.FHIRError)
}
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (r4b.ResearchSubject, capabilities.FHIRError)
}
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (r4b.RiskAssessment, capabilities.FHIRError)
}
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (r4b.Schedule, capabilities.FHIRError)
}
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (r4b.SearchParameter, capabilities.FHIRError)
}
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (r4b.ServiceRequest, capabilities.FHIRError)
}
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (r4b.Slot, capabilities.FHIRError)
}
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (r4b.Specimen, capabilities.FHIRError)
}
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (r4b.SpecimenDefinition, capabilities.FHIRError)
}
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (r4b.StructureDefinition, capabilities.FHIRError)
}
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (r4b.StructureMap, capabilities.FHIRError)
}
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (r4b.Subscription, capabilities.FHIRError)
}
type SubscriptionStatusRead interface {
	ReadSubscriptionStatus(ctx context.Context, id string) (r4b.SubscriptionStatus, capabilities.FHIRError)
}
type SubscriptionTopicRead interface {
	ReadSubscriptionTopic(ctx context.Context, id string) (r4b.SubscriptionTopic, capabilities.FHIRError)
}
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (r4b.Substance, capabilities.FHIRError)
}
type SubstanceDefinitionRead interface {
	ReadSubstanceDefinition(ctx context.Context, id string) (r4b.SubstanceDefinition, capabilities.FHIRError)
}
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (r4b.SupplyDelivery, capabilities.FHIRError)
}
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (r4b.SupplyRequest, capabilities.FHIRError)
}
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (r4b.Task, capabilities.FHIRError)
}
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (r4b.TerminologyCapabilities, capabilities.FHIRError)
}
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (r4b.TestReport, capabilities.FHIRError)
}
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (r4b.TestScript, capabilities.FHIRError)
}
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (r4b.ValueSet, capabilities.FHIRError)
}
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (r4b.VerificationResult, capabilities.FHIRError)
}
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (r4b.VisionPrescription, capabilities.FHIRError)
}
