package capabilitiesR4

import (
	"context"
	r4 "fhir-toolbox/model/gen/r4"
)

type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (r4.Account, error)
}
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (r4.ActivityDefinition, error)
}
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (r4.AdverseEvent, error)
}
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (r4.AllergyIntolerance, error)
}
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (r4.Appointment, error)
}
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (r4.AppointmentResponse, error)
}
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (r4.AuditEvent, error)
}
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (r4.Basic, error)
}
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (r4.Binary, error)
}
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4.BiologicallyDerivedProduct, error)
}
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (r4.BodyStructure, error)
}
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (r4.Bundle, error)
}
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (r4.CapabilityStatement, error)
}
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (r4.CarePlan, error)
}
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (r4.CareTeam, error)
}
type CatalogEntryRead interface {
	ReadCatalogEntry(ctx context.Context, id string) (r4.CatalogEntry, error)
}
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (r4.ChargeItem, error)
}
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (r4.ChargeItemDefinition, error)
}
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (r4.Claim, error)
}
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (r4.ClaimResponse, error)
}
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (r4.ClinicalImpression, error)
}
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (r4.CodeSystem, error)
}
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (r4.Communication, error)
}
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (r4.CommunicationRequest, error)
}
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (r4.CompartmentDefinition, error)
}
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (r4.Composition, error)
}
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (r4.ConceptMap, error)
}
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (r4.Condition, error)
}
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (r4.Consent, error)
}
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (r4.Contract, error)
}
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (r4.Coverage, error)
}
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4.CoverageEligibilityRequest, error)
}
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4.CoverageEligibilityResponse, error)
}
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (r4.DetectedIssue, error)
}
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (r4.Device, error)
}
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (r4.DeviceDefinition, error)
}
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (r4.DeviceMetric, error)
}
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (r4.DeviceRequest, error)
}
type DeviceUseStatementRead interface {
	ReadDeviceUseStatement(ctx context.Context, id string) (r4.DeviceUseStatement, error)
}
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (r4.DiagnosticReport, error)
}
type DocumentManifestRead interface {
	ReadDocumentManifest(ctx context.Context, id string) (r4.DocumentManifest, error)
}
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (r4.DocumentReference, error)
}
type EffectEvidenceSynthesisRead interface {
	ReadEffectEvidenceSynthesis(ctx context.Context, id string) (r4.EffectEvidenceSynthesis, error)
}
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (r4.Encounter, error)
}
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (r4.Endpoint, error)
}
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (r4.EnrollmentRequest, error)
}
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (r4.EnrollmentResponse, error)
}
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (r4.EpisodeOfCare, error)
}
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (r4.EventDefinition, error)
}
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (r4.Evidence, error)
}
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (r4.EvidenceVariable, error)
}
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (r4.ExampleScenario, error)
}
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (r4.ExplanationOfBenefit, error)
}
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (r4.FamilyMemberHistory, error)
}
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (r4.Flag, error)
}
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (r4.Goal, error)
}
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (r4.GraphDefinition, error)
}
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (r4.Group, error)
}
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (r4.GuidanceResponse, error)
}
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (r4.HealthcareService, error)
}
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (r4.ImagingStudy, error)
}
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (r4.Immunization, error)
}
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (r4.ImmunizationEvaluation, error)
}
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (r4.ImmunizationRecommendation, error)
}
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (r4.ImplementationGuide, error)
}
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (r4.InsurancePlan, error)
}
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (r4.Invoice, error)
}
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (r4.Library, error)
}
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (r4.Linkage, error)
}
type ListRead interface {
	ReadList(ctx context.Context, id string) (r4.List, error)
}
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (r4.Location, error)
}
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (r4.Measure, error)
}
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (r4.MeasureReport, error)
}
type MediaRead interface {
	ReadMedia(ctx context.Context, id string) (r4.Media, error)
}
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (r4.Medication, error)
}
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (r4.MedicationAdministration, error)
}
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (r4.MedicationDispense, error)
}
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (r4.MedicationKnowledge, error)
}
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (r4.MedicationRequest, error)
}
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (r4.MedicationStatement, error)
}
type MedicinalProductRead interface {
	ReadMedicinalProduct(ctx context.Context, id string) (r4.MedicinalProduct, error)
}
type MedicinalProductAuthorizationRead interface {
	ReadMedicinalProductAuthorization(ctx context.Context, id string) (r4.MedicinalProductAuthorization, error)
}
type MedicinalProductContraindicationRead interface {
	ReadMedicinalProductContraindication(ctx context.Context, id string) (r4.MedicinalProductContraindication, error)
}
type MedicinalProductIndicationRead interface {
	ReadMedicinalProductIndication(ctx context.Context, id string) (r4.MedicinalProductIndication, error)
}
type MedicinalProductIngredientRead interface {
	ReadMedicinalProductIngredient(ctx context.Context, id string) (r4.MedicinalProductIngredient, error)
}
type MedicinalProductInteractionRead interface {
	ReadMedicinalProductInteraction(ctx context.Context, id string) (r4.MedicinalProductInteraction, error)
}
type MedicinalProductManufacturedRead interface {
	ReadMedicinalProductManufactured(ctx context.Context, id string) (r4.MedicinalProductManufactured, error)
}
type MedicinalProductPackagedRead interface {
	ReadMedicinalProductPackaged(ctx context.Context, id string) (r4.MedicinalProductPackaged, error)
}
type MedicinalProductPharmaceuticalRead interface {
	ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (r4.MedicinalProductPharmaceutical, error)
}
type MedicinalProductUndesirableEffectRead interface {
	ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (r4.MedicinalProductUndesirableEffect, error)
}
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (r4.MessageDefinition, error)
}
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (r4.MessageHeader, error)
}
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (r4.MolecularSequence, error)
}
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (r4.NamingSystem, error)
}
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (r4.NutritionOrder, error)
}
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (r4.Observation, error)
}
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (r4.ObservationDefinition, error)
}
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (r4.OperationDefinition, error)
}
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (r4.OperationOutcome, error)
}
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (r4.Organization, error)
}
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (r4.OrganizationAffiliation, error)
}
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (r4.Parameters, error)
}
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (r4.Patient, error)
}
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (r4.PaymentNotice, error)
}
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (r4.PaymentReconciliation, error)
}
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (r4.Person, error)
}
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (r4.PlanDefinition, error)
}
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (r4.Practitioner, error)
}
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (r4.PractitionerRole, error)
}
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (r4.Procedure, error)
}
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (r4.Provenance, error)
}
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (r4.Questionnaire, error)
}
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (r4.QuestionnaireResponse, error)
}
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (r4.RelatedPerson, error)
}
type RequestGroupRead interface {
	ReadRequestGroup(ctx context.Context, id string) (r4.RequestGroup, error)
}
type ResearchDefinitionRead interface {
	ReadResearchDefinition(ctx context.Context, id string) (r4.ResearchDefinition, error)
}
type ResearchElementDefinitionRead interface {
	ReadResearchElementDefinition(ctx context.Context, id string) (r4.ResearchElementDefinition, error)
}
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (r4.ResearchStudy, error)
}
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (r4.ResearchSubject, error)
}
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (r4.RiskAssessment, error)
}
type RiskEvidenceSynthesisRead interface {
	ReadRiskEvidenceSynthesis(ctx context.Context, id string) (r4.RiskEvidenceSynthesis, error)
}
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (r4.Schedule, error)
}
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (r4.SearchParameter, error)
}
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (r4.ServiceRequest, error)
}
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (r4.Slot, error)
}
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (r4.Specimen, error)
}
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (r4.SpecimenDefinition, error)
}
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (r4.StructureDefinition, error)
}
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (r4.StructureMap, error)
}
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (r4.Subscription, error)
}
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (r4.Substance, error)
}
type SubstanceNucleicAcidRead interface {
	ReadSubstanceNucleicAcid(ctx context.Context, id string) (r4.SubstanceNucleicAcid, error)
}
type SubstancePolymerRead interface {
	ReadSubstancePolymer(ctx context.Context, id string) (r4.SubstancePolymer, error)
}
type SubstanceProteinRead interface {
	ReadSubstanceProtein(ctx context.Context, id string) (r4.SubstanceProtein, error)
}
type SubstanceReferenceInformationRead interface {
	ReadSubstanceReferenceInformation(ctx context.Context, id string) (r4.SubstanceReferenceInformation, error)
}
type SubstanceSourceMaterialRead interface {
	ReadSubstanceSourceMaterial(ctx context.Context, id string) (r4.SubstanceSourceMaterial, error)
}
type SubstanceSpecificationRead interface {
	ReadSubstanceSpecification(ctx context.Context, id string) (r4.SubstanceSpecification, error)
}
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (r4.SupplyDelivery, error)
}
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (r4.SupplyRequest, error)
}
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (r4.Task, error)
}
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (r4.TerminologyCapabilities, error)
}
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (r4.TestReport, error)
}
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (r4.TestScript, error)
}
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (r4.ValueSet, error)
}
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (r4.VerificationResult, error)
}
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (r4.VisionPrescription, error)
}
