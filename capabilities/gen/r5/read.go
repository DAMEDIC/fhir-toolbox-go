package capabilitiesR5

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	r5 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
)

type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (r5.Account, capabilities.FHIRError)
}
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (r5.ActivityDefinition, capabilities.FHIRError)
}
type ActorDefinitionRead interface {
	ReadActorDefinition(ctx context.Context, id string) (r5.ActorDefinition, capabilities.FHIRError)
}
type AdministrableProductDefinitionRead interface {
	ReadAdministrableProductDefinition(ctx context.Context, id string) (r5.AdministrableProductDefinition, capabilities.FHIRError)
}
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (r5.AdverseEvent, capabilities.FHIRError)
}
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (r5.AllergyIntolerance, capabilities.FHIRError)
}
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (r5.Appointment, capabilities.FHIRError)
}
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (r5.AppointmentResponse, capabilities.FHIRError)
}
type ArtifactAssessmentRead interface {
	ReadArtifactAssessment(ctx context.Context, id string) (r5.ArtifactAssessment, capabilities.FHIRError)
}
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (r5.AuditEvent, capabilities.FHIRError)
}
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (r5.Basic, capabilities.FHIRError)
}
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (r5.Binary, capabilities.FHIRError)
}
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r5.BiologicallyDerivedProduct, capabilities.FHIRError)
}
type BiologicallyDerivedProductDispenseRead interface {
	ReadBiologicallyDerivedProductDispense(ctx context.Context, id string) (r5.BiologicallyDerivedProductDispense, capabilities.FHIRError)
}
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (r5.BodyStructure, capabilities.FHIRError)
}
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (r5.Bundle, capabilities.FHIRError)
}
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (r5.CapabilityStatement, capabilities.FHIRError)
}
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (r5.CarePlan, capabilities.FHIRError)
}
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (r5.CareTeam, capabilities.FHIRError)
}
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (r5.ChargeItem, capabilities.FHIRError)
}
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (r5.ChargeItemDefinition, capabilities.FHIRError)
}
type CitationRead interface {
	ReadCitation(ctx context.Context, id string) (r5.Citation, capabilities.FHIRError)
}
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (r5.Claim, capabilities.FHIRError)
}
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (r5.ClaimResponse, capabilities.FHIRError)
}
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (r5.ClinicalImpression, capabilities.FHIRError)
}
type ClinicalUseDefinitionRead interface {
	ReadClinicalUseDefinition(ctx context.Context, id string) (r5.ClinicalUseDefinition, capabilities.FHIRError)
}
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (r5.CodeSystem, capabilities.FHIRError)
}
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (r5.Communication, capabilities.FHIRError)
}
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (r5.CommunicationRequest, capabilities.FHIRError)
}
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (r5.CompartmentDefinition, capabilities.FHIRError)
}
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (r5.Composition, capabilities.FHIRError)
}
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (r5.ConceptMap, capabilities.FHIRError)
}
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (r5.Condition, capabilities.FHIRError)
}
type ConditionDefinitionRead interface {
	ReadConditionDefinition(ctx context.Context, id string) (r5.ConditionDefinition, capabilities.FHIRError)
}
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (r5.Consent, capabilities.FHIRError)
}
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (r5.Contract, capabilities.FHIRError)
}
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (r5.Coverage, capabilities.FHIRError)
}
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (r5.CoverageEligibilityRequest, capabilities.FHIRError)
}
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (r5.CoverageEligibilityResponse, capabilities.FHIRError)
}
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (r5.DetectedIssue, capabilities.FHIRError)
}
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (r5.Device, capabilities.FHIRError)
}
type DeviceAssociationRead interface {
	ReadDeviceAssociation(ctx context.Context, id string) (r5.DeviceAssociation, capabilities.FHIRError)
}
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (r5.DeviceDefinition, capabilities.FHIRError)
}
type DeviceDispenseRead interface {
	ReadDeviceDispense(ctx context.Context, id string) (r5.DeviceDispense, capabilities.FHIRError)
}
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (r5.DeviceMetric, capabilities.FHIRError)
}
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (r5.DeviceRequest, capabilities.FHIRError)
}
type DeviceUsageRead interface {
	ReadDeviceUsage(ctx context.Context, id string) (r5.DeviceUsage, capabilities.FHIRError)
}
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (r5.DiagnosticReport, capabilities.FHIRError)
}
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (r5.DocumentReference, capabilities.FHIRError)
}
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (r5.Encounter, capabilities.FHIRError)
}
type EncounterHistoryRead interface {
	ReadEncounterHistory(ctx context.Context, id string) (r5.EncounterHistory, capabilities.FHIRError)
}
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (r5.Endpoint, capabilities.FHIRError)
}
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (r5.EnrollmentRequest, capabilities.FHIRError)
}
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (r5.EnrollmentResponse, capabilities.FHIRError)
}
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (r5.EpisodeOfCare, capabilities.FHIRError)
}
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (r5.EventDefinition, capabilities.FHIRError)
}
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (r5.Evidence, capabilities.FHIRError)
}
type EvidenceReportRead interface {
	ReadEvidenceReport(ctx context.Context, id string) (r5.EvidenceReport, capabilities.FHIRError)
}
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (r5.EvidenceVariable, capabilities.FHIRError)
}
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (r5.ExampleScenario, capabilities.FHIRError)
}
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (r5.ExplanationOfBenefit, capabilities.FHIRError)
}
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (r5.FamilyMemberHistory, capabilities.FHIRError)
}
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (r5.Flag, capabilities.FHIRError)
}
type FormularyItemRead interface {
	ReadFormularyItem(ctx context.Context, id string) (r5.FormularyItem, capabilities.FHIRError)
}
type GenomicStudyRead interface {
	ReadGenomicStudy(ctx context.Context, id string) (r5.GenomicStudy, capabilities.FHIRError)
}
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (r5.Goal, capabilities.FHIRError)
}
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (r5.GraphDefinition, capabilities.FHIRError)
}
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (r5.Group, capabilities.FHIRError)
}
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (r5.GuidanceResponse, capabilities.FHIRError)
}
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (r5.HealthcareService, capabilities.FHIRError)
}
type ImagingSelectionRead interface {
	ReadImagingSelection(ctx context.Context, id string) (r5.ImagingSelection, capabilities.FHIRError)
}
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (r5.ImagingStudy, capabilities.FHIRError)
}
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (r5.Immunization, capabilities.FHIRError)
}
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (r5.ImmunizationEvaluation, capabilities.FHIRError)
}
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (r5.ImmunizationRecommendation, capabilities.FHIRError)
}
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (r5.ImplementationGuide, capabilities.FHIRError)
}
type IngredientRead interface {
	ReadIngredient(ctx context.Context, id string) (r5.Ingredient, capabilities.FHIRError)
}
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (r5.InsurancePlan, capabilities.FHIRError)
}
type InventoryItemRead interface {
	ReadInventoryItem(ctx context.Context, id string) (r5.InventoryItem, capabilities.FHIRError)
}
type InventoryReportRead interface {
	ReadInventoryReport(ctx context.Context, id string) (r5.InventoryReport, capabilities.FHIRError)
}
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (r5.Invoice, capabilities.FHIRError)
}
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (r5.Library, capabilities.FHIRError)
}
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (r5.Linkage, capabilities.FHIRError)
}
type ListRead interface {
	ReadList(ctx context.Context, id string) (r5.List, capabilities.FHIRError)
}
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (r5.Location, capabilities.FHIRError)
}
type ManufacturedItemDefinitionRead interface {
	ReadManufacturedItemDefinition(ctx context.Context, id string) (r5.ManufacturedItemDefinition, capabilities.FHIRError)
}
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (r5.Measure, capabilities.FHIRError)
}
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (r5.MeasureReport, capabilities.FHIRError)
}
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (r5.Medication, capabilities.FHIRError)
}
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (r5.MedicationAdministration, capabilities.FHIRError)
}
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (r5.MedicationDispense, capabilities.FHIRError)
}
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (r5.MedicationKnowledge, capabilities.FHIRError)
}
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (r5.MedicationRequest, capabilities.FHIRError)
}
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (r5.MedicationStatement, capabilities.FHIRError)
}
type MedicinalProductDefinitionRead interface {
	ReadMedicinalProductDefinition(ctx context.Context, id string) (r5.MedicinalProductDefinition, capabilities.FHIRError)
}
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (r5.MessageDefinition, capabilities.FHIRError)
}
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (r5.MessageHeader, capabilities.FHIRError)
}
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (r5.MolecularSequence, capabilities.FHIRError)
}
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (r5.NamingSystem, capabilities.FHIRError)
}
type NutritionIntakeRead interface {
	ReadNutritionIntake(ctx context.Context, id string) (r5.NutritionIntake, capabilities.FHIRError)
}
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (r5.NutritionOrder, capabilities.FHIRError)
}
type NutritionProductRead interface {
	ReadNutritionProduct(ctx context.Context, id string) (r5.NutritionProduct, capabilities.FHIRError)
}
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (r5.Observation, capabilities.FHIRError)
}
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (r5.ObservationDefinition, capabilities.FHIRError)
}
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (r5.OperationDefinition, capabilities.FHIRError)
}
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (r5.OperationOutcome, capabilities.FHIRError)
}
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (r5.Organization, capabilities.FHIRError)
}
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (r5.OrganizationAffiliation, capabilities.FHIRError)
}
type PackagedProductDefinitionRead interface {
	ReadPackagedProductDefinition(ctx context.Context, id string) (r5.PackagedProductDefinition, capabilities.FHIRError)
}
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (r5.Parameters, capabilities.FHIRError)
}
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (r5.Patient, capabilities.FHIRError)
}
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (r5.PaymentNotice, capabilities.FHIRError)
}
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (r5.PaymentReconciliation, capabilities.FHIRError)
}
type PermissionRead interface {
	ReadPermission(ctx context.Context, id string) (r5.Permission, capabilities.FHIRError)
}
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (r5.Person, capabilities.FHIRError)
}
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (r5.PlanDefinition, capabilities.FHIRError)
}
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (r5.Practitioner, capabilities.FHIRError)
}
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (r5.PractitionerRole, capabilities.FHIRError)
}
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (r5.Procedure, capabilities.FHIRError)
}
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (r5.Provenance, capabilities.FHIRError)
}
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (r5.Questionnaire, capabilities.FHIRError)
}
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (r5.QuestionnaireResponse, capabilities.FHIRError)
}
type RegulatedAuthorizationRead interface {
	ReadRegulatedAuthorization(ctx context.Context, id string) (r5.RegulatedAuthorization, capabilities.FHIRError)
}
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (r5.RelatedPerson, capabilities.FHIRError)
}
type RequestOrchestrationRead interface {
	ReadRequestOrchestration(ctx context.Context, id string) (r5.RequestOrchestration, capabilities.FHIRError)
}
type RequirementsRead interface {
	ReadRequirements(ctx context.Context, id string) (r5.Requirements, capabilities.FHIRError)
}
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (r5.ResearchStudy, capabilities.FHIRError)
}
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (r5.ResearchSubject, capabilities.FHIRError)
}
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (r5.RiskAssessment, capabilities.FHIRError)
}
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (r5.Schedule, capabilities.FHIRError)
}
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (r5.SearchParameter, capabilities.FHIRError)
}
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (r5.ServiceRequest, capabilities.FHIRError)
}
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (r5.Slot, capabilities.FHIRError)
}
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (r5.Specimen, capabilities.FHIRError)
}
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (r5.SpecimenDefinition, capabilities.FHIRError)
}
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (r5.StructureDefinition, capabilities.FHIRError)
}
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (r5.StructureMap, capabilities.FHIRError)
}
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (r5.Subscription, capabilities.FHIRError)
}
type SubscriptionStatusRead interface {
	ReadSubscriptionStatus(ctx context.Context, id string) (r5.SubscriptionStatus, capabilities.FHIRError)
}
type SubscriptionTopicRead interface {
	ReadSubscriptionTopic(ctx context.Context, id string) (r5.SubscriptionTopic, capabilities.FHIRError)
}
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (r5.Substance, capabilities.FHIRError)
}
type SubstanceDefinitionRead interface {
	ReadSubstanceDefinition(ctx context.Context, id string) (r5.SubstanceDefinition, capabilities.FHIRError)
}
type SubstanceNucleicAcidRead interface {
	ReadSubstanceNucleicAcid(ctx context.Context, id string) (r5.SubstanceNucleicAcid, capabilities.FHIRError)
}
type SubstancePolymerRead interface {
	ReadSubstancePolymer(ctx context.Context, id string) (r5.SubstancePolymer, capabilities.FHIRError)
}
type SubstanceProteinRead interface {
	ReadSubstanceProtein(ctx context.Context, id string) (r5.SubstanceProtein, capabilities.FHIRError)
}
type SubstanceReferenceInformationRead interface {
	ReadSubstanceReferenceInformation(ctx context.Context, id string) (r5.SubstanceReferenceInformation, capabilities.FHIRError)
}
type SubstanceSourceMaterialRead interface {
	ReadSubstanceSourceMaterial(ctx context.Context, id string) (r5.SubstanceSourceMaterial, capabilities.FHIRError)
}
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (r5.SupplyDelivery, capabilities.FHIRError)
}
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (r5.SupplyRequest, capabilities.FHIRError)
}
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (r5.Task, capabilities.FHIRError)
}
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (r5.TerminologyCapabilities, capabilities.FHIRError)
}
type TestPlanRead interface {
	ReadTestPlan(ctx context.Context, id string) (r5.TestPlan, capabilities.FHIRError)
}
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (r5.TestReport, capabilities.FHIRError)
}
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (r5.TestScript, capabilities.FHIRError)
}
type TransportRead interface {
	ReadTransport(ctx context.Context, id string) (r5.Transport, capabilities.FHIRError)
}
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (r5.ValueSet, capabilities.FHIRError)
}
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (r5.VerificationResult, capabilities.FHIRError)
}
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (r5.VisionPrescription, capabilities.FHIRError)
}
