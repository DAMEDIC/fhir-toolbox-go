// DO NOT EDIT!
// Code generated by "github.com/DAMEDIC/fhir-toolbox-go/internal/generate";
// use `make generate` to regenerate.

package capabilitiesR4B

import (
	"context"
	r4b "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
)

// AccountRead needs to be implemented to support the read interaction.
type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (r4b.Account, error)
}

// ActivityDefinitionRead needs to be implemented to support the read interaction.
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (r4b.ActivityDefinition, error)
}

// AdministrableProductDefinitionRead needs to be implemented to support the read interaction.
type AdministrableProductDefinitionRead interface {
	ReadAdministrableProductDefinition(ctx context.Context, id string) (r4b.AdministrableProductDefinition, error)
}

// AdverseEventRead needs to be implemented to support the read interaction.
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (r4b.AdverseEvent, error)
}

// AllergyIntoleranceRead needs to be implemented to support the read interaction.
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (r4b.AllergyIntolerance, error)
}

// AppointmentRead needs to be implemented to support the read interaction.
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (r4b.Appointment, error)
}

// AppointmentResponseRead needs to be implemented to support the read interaction.
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (r4b.AppointmentResponse, error)
}

// AuditEventRead needs to be implemented to support the read interaction.
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (r4b.AuditEvent, error)
}

// BasicRead needs to be implemented to support the read interaction.
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (r4b.Basic, error)
}

// BinaryRead needs to be implemented to support the read interaction.
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (r4b.Binary, error)
}

// BiologicallyDerivedProductRead needs to be implemented to support the read interaction.
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4b.BiologicallyDerivedProduct, error)
}

// BodyStructureRead needs to be implemented to support the read interaction.
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (r4b.BodyStructure, error)
}

// BundleRead needs to be implemented to support the read interaction.
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (r4b.Bundle, error)
}

// CapabilityStatementRead needs to be implemented to support the read interaction.
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (r4b.CapabilityStatement, error)
}

// CarePlanRead needs to be implemented to support the read interaction.
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (r4b.CarePlan, error)
}

// CareTeamRead needs to be implemented to support the read interaction.
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (r4b.CareTeam, error)
}

// CatalogEntryRead needs to be implemented to support the read interaction.
type CatalogEntryRead interface {
	ReadCatalogEntry(ctx context.Context, id string) (r4b.CatalogEntry, error)
}

// ChargeItemRead needs to be implemented to support the read interaction.
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (r4b.ChargeItem, error)
}

// ChargeItemDefinitionRead needs to be implemented to support the read interaction.
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (r4b.ChargeItemDefinition, error)
}

// CitationRead needs to be implemented to support the read interaction.
type CitationRead interface {
	ReadCitation(ctx context.Context, id string) (r4b.Citation, error)
}

// ClaimRead needs to be implemented to support the read interaction.
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (r4b.Claim, error)
}

// ClaimResponseRead needs to be implemented to support the read interaction.
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (r4b.ClaimResponse, error)
}

// ClinicalImpressionRead needs to be implemented to support the read interaction.
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (r4b.ClinicalImpression, error)
}

// ClinicalUseDefinitionRead needs to be implemented to support the read interaction.
type ClinicalUseDefinitionRead interface {
	ReadClinicalUseDefinition(ctx context.Context, id string) (r4b.ClinicalUseDefinition, error)
}

// CodeSystemRead needs to be implemented to support the read interaction.
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (r4b.CodeSystem, error)
}

// CommunicationRead needs to be implemented to support the read interaction.
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (r4b.Communication, error)
}

// CommunicationRequestRead needs to be implemented to support the read interaction.
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (r4b.CommunicationRequest, error)
}

// CompartmentDefinitionRead needs to be implemented to support the read interaction.
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (r4b.CompartmentDefinition, error)
}

// CompositionRead needs to be implemented to support the read interaction.
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (r4b.Composition, error)
}

// ConceptMapRead needs to be implemented to support the read interaction.
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (r4b.ConceptMap, error)
}

// ConditionRead needs to be implemented to support the read interaction.
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (r4b.Condition, error)
}

// ConsentRead needs to be implemented to support the read interaction.
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (r4b.Consent, error)
}

// ContractRead needs to be implemented to support the read interaction.
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (r4b.Contract, error)
}

// CoverageRead needs to be implemented to support the read interaction.
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (r4b.Coverage, error)
}

// CoverageEligibilityRequestRead needs to be implemented to support the read interaction.
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4b.CoverageEligibilityRequest, error)
}

// CoverageEligibilityResponseRead needs to be implemented to support the read interaction.
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4b.CoverageEligibilityResponse, error)
}

// DetectedIssueRead needs to be implemented to support the read interaction.
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (r4b.DetectedIssue, error)
}

// DeviceRead needs to be implemented to support the read interaction.
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (r4b.Device, error)
}

// DeviceDefinitionRead needs to be implemented to support the read interaction.
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (r4b.DeviceDefinition, error)
}

// DeviceMetricRead needs to be implemented to support the read interaction.
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (r4b.DeviceMetric, error)
}

// DeviceRequestRead needs to be implemented to support the read interaction.
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (r4b.DeviceRequest, error)
}

// DeviceUseStatementRead needs to be implemented to support the read interaction.
type DeviceUseStatementRead interface {
	ReadDeviceUseStatement(ctx context.Context, id string) (r4b.DeviceUseStatement, error)
}

// DiagnosticReportRead needs to be implemented to support the read interaction.
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (r4b.DiagnosticReport, error)
}

// DocumentManifestRead needs to be implemented to support the read interaction.
type DocumentManifestRead interface {
	ReadDocumentManifest(ctx context.Context, id string) (r4b.DocumentManifest, error)
}

// DocumentReferenceRead needs to be implemented to support the read interaction.
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (r4b.DocumentReference, error)
}

// EncounterRead needs to be implemented to support the read interaction.
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (r4b.Encounter, error)
}

// EndpointRead needs to be implemented to support the read interaction.
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (r4b.Endpoint, error)
}

// EnrollmentRequestRead needs to be implemented to support the read interaction.
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (r4b.EnrollmentRequest, error)
}

// EnrollmentResponseRead needs to be implemented to support the read interaction.
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (r4b.EnrollmentResponse, error)
}

// EpisodeOfCareRead needs to be implemented to support the read interaction.
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (r4b.EpisodeOfCare, error)
}

// EventDefinitionRead needs to be implemented to support the read interaction.
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (r4b.EventDefinition, error)
}

// EvidenceRead needs to be implemented to support the read interaction.
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (r4b.Evidence, error)
}

// EvidenceReportRead needs to be implemented to support the read interaction.
type EvidenceReportRead interface {
	ReadEvidenceReport(ctx context.Context, id string) (r4b.EvidenceReport, error)
}

// EvidenceVariableRead needs to be implemented to support the read interaction.
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (r4b.EvidenceVariable, error)
}

// ExampleScenarioRead needs to be implemented to support the read interaction.
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (r4b.ExampleScenario, error)
}

// ExplanationOfBenefitRead needs to be implemented to support the read interaction.
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (r4b.ExplanationOfBenefit, error)
}

// FamilyMemberHistoryRead needs to be implemented to support the read interaction.
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (r4b.FamilyMemberHistory, error)
}

// FlagRead needs to be implemented to support the read interaction.
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (r4b.Flag, error)
}

// GoalRead needs to be implemented to support the read interaction.
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (r4b.Goal, error)
}

// GraphDefinitionRead needs to be implemented to support the read interaction.
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (r4b.GraphDefinition, error)
}

// GroupRead needs to be implemented to support the read interaction.
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (r4b.Group, error)
}

// GuidanceResponseRead needs to be implemented to support the read interaction.
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (r4b.GuidanceResponse, error)
}

// HealthcareServiceRead needs to be implemented to support the read interaction.
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (r4b.HealthcareService, error)
}

// ImagingStudyRead needs to be implemented to support the read interaction.
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (r4b.ImagingStudy, error)
}

// ImmunizationRead needs to be implemented to support the read interaction.
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (r4b.Immunization, error)
}

// ImmunizationEvaluationRead needs to be implemented to support the read interaction.
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (r4b.ImmunizationEvaluation, error)
}

// ImmunizationRecommendationRead needs to be implemented to support the read interaction.
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (r4b.ImmunizationRecommendation, error)
}

// ImplementationGuideRead needs to be implemented to support the read interaction.
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (r4b.ImplementationGuide, error)
}

// IngredientRead needs to be implemented to support the read interaction.
type IngredientRead interface {
	ReadIngredient(ctx context.Context, id string) (r4b.Ingredient, error)
}

// InsurancePlanRead needs to be implemented to support the read interaction.
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (r4b.InsurancePlan, error)
}

// InvoiceRead needs to be implemented to support the read interaction.
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (r4b.Invoice, error)
}

// LibraryRead needs to be implemented to support the read interaction.
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (r4b.Library, error)
}

// LinkageRead needs to be implemented to support the read interaction.
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (r4b.Linkage, error)
}

// ListRead needs to be implemented to support the read interaction.
type ListRead interface {
	ReadList(ctx context.Context, id string) (r4b.List, error)
}

// LocationRead needs to be implemented to support the read interaction.
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (r4b.Location, error)
}

// ManufacturedItemDefinitionRead needs to be implemented to support the read interaction.
type ManufacturedItemDefinitionRead interface {
	ReadManufacturedItemDefinition(ctx context.Context, id string) (r4b.ManufacturedItemDefinition, error)
}

// MeasureRead needs to be implemented to support the read interaction.
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (r4b.Measure, error)
}

// MeasureReportRead needs to be implemented to support the read interaction.
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (r4b.MeasureReport, error)
}

// MediaRead needs to be implemented to support the read interaction.
type MediaRead interface {
	ReadMedia(ctx context.Context, id string) (r4b.Media, error)
}

// MedicationRead needs to be implemented to support the read interaction.
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (r4b.Medication, error)
}

// MedicationAdministrationRead needs to be implemented to support the read interaction.
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (r4b.MedicationAdministration, error)
}

// MedicationDispenseRead needs to be implemented to support the read interaction.
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (r4b.MedicationDispense, error)
}

// MedicationKnowledgeRead needs to be implemented to support the read interaction.
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (r4b.MedicationKnowledge, error)
}

// MedicationRequestRead needs to be implemented to support the read interaction.
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (r4b.MedicationRequest, error)
}

// MedicationStatementRead needs to be implemented to support the read interaction.
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (r4b.MedicationStatement, error)
}

// MedicinalProductDefinitionRead needs to be implemented to support the read interaction.
type MedicinalProductDefinitionRead interface {
	ReadMedicinalProductDefinition(ctx context.Context, id string) (r4b.MedicinalProductDefinition, error)
}

// MessageDefinitionRead needs to be implemented to support the read interaction.
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (r4b.MessageDefinition, error)
}

// MessageHeaderRead needs to be implemented to support the read interaction.
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (r4b.MessageHeader, error)
}

// MolecularSequenceRead needs to be implemented to support the read interaction.
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (r4b.MolecularSequence, error)
}

// NamingSystemRead needs to be implemented to support the read interaction.
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (r4b.NamingSystem, error)
}

// NutritionOrderRead needs to be implemented to support the read interaction.
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (r4b.NutritionOrder, error)
}

// NutritionProductRead needs to be implemented to support the read interaction.
type NutritionProductRead interface {
	ReadNutritionProduct(ctx context.Context, id string) (r4b.NutritionProduct, error)
}

// ObservationRead needs to be implemented to support the read interaction.
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (r4b.Observation, error)
}

// ObservationDefinitionRead needs to be implemented to support the read interaction.
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (r4b.ObservationDefinition, error)
}

// OperationDefinitionRead needs to be implemented to support the read interaction.
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (r4b.OperationDefinition, error)
}

// OperationOutcomeRead needs to be implemented to support the read interaction.
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (r4b.OperationOutcome, error)
}

// OrganizationRead needs to be implemented to support the read interaction.
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (r4b.Organization, error)
}

// OrganizationAffiliationRead needs to be implemented to support the read interaction.
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (r4b.OrganizationAffiliation, error)
}

// PackagedProductDefinitionRead needs to be implemented to support the read interaction.
type PackagedProductDefinitionRead interface {
	ReadPackagedProductDefinition(ctx context.Context, id string) (r4b.PackagedProductDefinition, error)
}

// ParametersRead needs to be implemented to support the read interaction.
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (r4b.Parameters, error)
}

// PatientRead needs to be implemented to support the read interaction.
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (r4b.Patient, error)
}

// PaymentNoticeRead needs to be implemented to support the read interaction.
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (r4b.PaymentNotice, error)
}

// PaymentReconciliationRead needs to be implemented to support the read interaction.
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (r4b.PaymentReconciliation, error)
}

// PersonRead needs to be implemented to support the read interaction.
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (r4b.Person, error)
}

// PlanDefinitionRead needs to be implemented to support the read interaction.
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (r4b.PlanDefinition, error)
}

// PractitionerRead needs to be implemented to support the read interaction.
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (r4b.Practitioner, error)
}

// PractitionerRoleRead needs to be implemented to support the read interaction.
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (r4b.PractitionerRole, error)
}

// ProcedureRead needs to be implemented to support the read interaction.
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (r4b.Procedure, error)
}

// ProvenanceRead needs to be implemented to support the read interaction.
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (r4b.Provenance, error)
}

// QuestionnaireRead needs to be implemented to support the read interaction.
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (r4b.Questionnaire, error)
}

// QuestionnaireResponseRead needs to be implemented to support the read interaction.
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (r4b.QuestionnaireResponse, error)
}

// RegulatedAuthorizationRead needs to be implemented to support the read interaction.
type RegulatedAuthorizationRead interface {
	ReadRegulatedAuthorization(ctx context.Context, id string) (r4b.RegulatedAuthorization, error)
}

// RelatedPersonRead needs to be implemented to support the read interaction.
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (r4b.RelatedPerson, error)
}

// RequestGroupRead needs to be implemented to support the read interaction.
type RequestGroupRead interface {
	ReadRequestGroup(ctx context.Context, id string) (r4b.RequestGroup, error)
}

// ResearchDefinitionRead needs to be implemented to support the read interaction.
type ResearchDefinitionRead interface {
	ReadResearchDefinition(ctx context.Context, id string) (r4b.ResearchDefinition, error)
}

// ResearchElementDefinitionRead needs to be implemented to support the read interaction.
type ResearchElementDefinitionRead interface {
	ReadResearchElementDefinition(ctx context.Context, id string) (r4b.ResearchElementDefinition, error)
}

// ResearchStudyRead needs to be implemented to support the read interaction.
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (r4b.ResearchStudy, error)
}

// ResearchSubjectRead needs to be implemented to support the read interaction.
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (r4b.ResearchSubject, error)
}

// RiskAssessmentRead needs to be implemented to support the read interaction.
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (r4b.RiskAssessment, error)
}

// ScheduleRead needs to be implemented to support the read interaction.
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (r4b.Schedule, error)
}

// SearchParameterRead needs to be implemented to support the read interaction.
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (r4b.SearchParameter, error)
}

// ServiceRequestRead needs to be implemented to support the read interaction.
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (r4b.ServiceRequest, error)
}

// SlotRead needs to be implemented to support the read interaction.
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (r4b.Slot, error)
}

// SpecimenRead needs to be implemented to support the read interaction.
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (r4b.Specimen, error)
}

// SpecimenDefinitionRead needs to be implemented to support the read interaction.
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (r4b.SpecimenDefinition, error)
}

// StructureDefinitionRead needs to be implemented to support the read interaction.
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (r4b.StructureDefinition, error)
}

// StructureMapRead needs to be implemented to support the read interaction.
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (r4b.StructureMap, error)
}

// SubscriptionRead needs to be implemented to support the read interaction.
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (r4b.Subscription, error)
}

// SubscriptionStatusRead needs to be implemented to support the read interaction.
type SubscriptionStatusRead interface {
	ReadSubscriptionStatus(ctx context.Context, id string) (r4b.SubscriptionStatus, error)
}

// SubscriptionTopicRead needs to be implemented to support the read interaction.
type SubscriptionTopicRead interface {
	ReadSubscriptionTopic(ctx context.Context, id string) (r4b.SubscriptionTopic, error)
}

// SubstanceRead needs to be implemented to support the read interaction.
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (r4b.Substance, error)
}

// SubstanceDefinitionRead needs to be implemented to support the read interaction.
type SubstanceDefinitionRead interface {
	ReadSubstanceDefinition(ctx context.Context, id string) (r4b.SubstanceDefinition, error)
}

// SupplyDeliveryRead needs to be implemented to support the read interaction.
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (r4b.SupplyDelivery, error)
}

// SupplyRequestRead needs to be implemented to support the read interaction.
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (r4b.SupplyRequest, error)
}

// TaskRead needs to be implemented to support the read interaction.
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (r4b.Task, error)
}

// TerminologyCapabilitiesRead needs to be implemented to support the read interaction.
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (r4b.TerminologyCapabilities, error)
}

// TestReportRead needs to be implemented to support the read interaction.
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (r4b.TestReport, error)
}

// TestScriptRead needs to be implemented to support the read interaction.
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (r4b.TestScript, error)
}

// ValueSetRead needs to be implemented to support the read interaction.
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (r4b.ValueSet, error)
}

// VerificationResultRead needs to be implemented to support the read interaction.
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (r4b.VerificationResult, error)
}

// VisionPrescriptionRead needs to be implemented to support the read interaction.
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (r4b.VisionPrescription, error)
}
