package capabilitiesR4

import (
	"context"
	model "fhir-toolbox/model/gen/r4"
)

type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (model.Account, error)
}
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (model.ActivityDefinition, error)
}
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (model.AdverseEvent, error)
}
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (model.AllergyIntolerance, error)
}
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (model.Appointment, error)
}
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (model.AppointmentResponse, error)
}
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (model.AuditEvent, error)
}
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (model.Basic, error)
}
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (model.Binary, error)
}
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (model.BiologicallyDerivedProduct, error)
}
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (model.BodyStructure, error)
}
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (model.Bundle, error)
}
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (model.CapabilityStatement, error)
}
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (model.CarePlan, error)
}
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (model.CareTeam, error)
}
type CatalogEntryRead interface {
	ReadCatalogEntry(ctx context.Context, id string) (model.CatalogEntry, error)
}
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (model.ChargeItem, error)
}
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (model.ChargeItemDefinition, error)
}
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (model.Claim, error)
}
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (model.ClaimResponse, error)
}
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (model.ClinicalImpression, error)
}
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (model.CodeSystem, error)
}
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (model.Communication, error)
}
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (model.CommunicationRequest, error)
}
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (model.CompartmentDefinition, error)
}
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (model.Composition, error)
}
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (model.ConceptMap, error)
}
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (model.Condition, error)
}
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (model.Consent, error)
}
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (model.Contract, error)
}
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (model.Coverage, error)
}
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (model.CoverageEligibilityRequest, error)
}
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (model.CoverageEligibilityResponse, error)
}
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (model.DetectedIssue, error)
}
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (model.Device, error)
}
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (model.DeviceDefinition, error)
}
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (model.DeviceMetric, error)
}
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (model.DeviceRequest, error)
}
type DeviceUseStatementRead interface {
	ReadDeviceUseStatement(ctx context.Context, id string) (model.DeviceUseStatement, error)
}
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (model.DiagnosticReport, error)
}
type DocumentManifestRead interface {
	ReadDocumentManifest(ctx context.Context, id string) (model.DocumentManifest, error)
}
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (model.DocumentReference, error)
}
type EffectEvidenceSynthesisRead interface {
	ReadEffectEvidenceSynthesis(ctx context.Context, id string) (model.EffectEvidenceSynthesis, error)
}
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (model.Encounter, error)
}
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (model.Endpoint, error)
}
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (model.EnrollmentRequest, error)
}
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (model.EnrollmentResponse, error)
}
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (model.EpisodeOfCare, error)
}
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (model.EventDefinition, error)
}
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (model.Evidence, error)
}
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (model.EvidenceVariable, error)
}
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (model.ExampleScenario, error)
}
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (model.ExplanationOfBenefit, error)
}
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (model.FamilyMemberHistory, error)
}
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (model.Flag, error)
}
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (model.Goal, error)
}
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (model.GraphDefinition, error)
}
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (model.Group, error)
}
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (model.GuidanceResponse, error)
}
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (model.HealthcareService, error)
}
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (model.ImagingStudy, error)
}
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (model.Immunization, error)
}
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (model.ImmunizationEvaluation, error)
}
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (model.ImmunizationRecommendation, error)
}
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (model.ImplementationGuide, error)
}
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (model.InsurancePlan, error)
}
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (model.Invoice, error)
}
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (model.Library, error)
}
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (model.Linkage, error)
}
type ListRead interface {
	ReadList(ctx context.Context, id string) (model.List, error)
}
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (model.Location, error)
}
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (model.Measure, error)
}
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (model.MeasureReport, error)
}
type MediaRead interface {
	ReadMedia(ctx context.Context, id string) (model.Media, error)
}
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (model.Medication, error)
}
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (model.MedicationAdministration, error)
}
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (model.MedicationDispense, error)
}
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (model.MedicationKnowledge, error)
}
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (model.MedicationRequest, error)
}
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (model.MedicationStatement, error)
}
type MedicinalProductRead interface {
	ReadMedicinalProduct(ctx context.Context, id string) (model.MedicinalProduct, error)
}
type MedicinalProductAuthorizationRead interface {
	ReadMedicinalProductAuthorization(ctx context.Context, id string) (model.MedicinalProductAuthorization, error)
}
type MedicinalProductContraindicationRead interface {
	ReadMedicinalProductContraindication(ctx context.Context, id string) (model.MedicinalProductContraindication, error)
}
type MedicinalProductIndicationRead interface {
	ReadMedicinalProductIndication(ctx context.Context, id string) (model.MedicinalProductIndication, error)
}
type MedicinalProductIngredientRead interface {
	ReadMedicinalProductIngredient(ctx context.Context, id string) (model.MedicinalProductIngredient, error)
}
type MedicinalProductInteractionRead interface {
	ReadMedicinalProductInteraction(ctx context.Context, id string) (model.MedicinalProductInteraction, error)
}
type MedicinalProductManufacturedRead interface {
	ReadMedicinalProductManufactured(ctx context.Context, id string) (model.MedicinalProductManufactured, error)
}
type MedicinalProductPackagedRead interface {
	ReadMedicinalProductPackaged(ctx context.Context, id string) (model.MedicinalProductPackaged, error)
}
type MedicinalProductPharmaceuticalRead interface {
	ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (model.MedicinalProductPharmaceutical, error)
}
type MedicinalProductUndesirableEffectRead interface {
	ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (model.MedicinalProductUndesirableEffect, error)
}
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (model.MessageDefinition, error)
}
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (model.MessageHeader, error)
}
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (model.MolecularSequence, error)
}
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (model.NamingSystem, error)
}
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (model.NutritionOrder, error)
}
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (model.Observation, error)
}
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (model.ObservationDefinition, error)
}
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (model.OperationDefinition, error)
}
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (model.OperationOutcome, error)
}
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (model.Organization, error)
}
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (model.OrganizationAffiliation, error)
}
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (model.Parameters, error)
}
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (model.Patient, error)
}
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (model.PaymentNotice, error)
}
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (model.PaymentReconciliation, error)
}
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (model.Person, error)
}
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (model.PlanDefinition, error)
}
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (model.Practitioner, error)
}
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (model.PractitionerRole, error)
}
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (model.Procedure, error)
}
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (model.Provenance, error)
}
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (model.Questionnaire, error)
}
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (model.QuestionnaireResponse, error)
}
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (model.RelatedPerson, error)
}
type RequestGroupRead interface {
	ReadRequestGroup(ctx context.Context, id string) (model.RequestGroup, error)
}
type ResearchDefinitionRead interface {
	ReadResearchDefinition(ctx context.Context, id string) (model.ResearchDefinition, error)
}
type ResearchElementDefinitionRead interface {
	ReadResearchElementDefinition(ctx context.Context, id string) (model.ResearchElementDefinition, error)
}
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (model.ResearchStudy, error)
}
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (model.ResearchSubject, error)
}
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (model.RiskAssessment, error)
}
type RiskEvidenceSynthesisRead interface {
	ReadRiskEvidenceSynthesis(ctx context.Context, id string) (model.RiskEvidenceSynthesis, error)
}
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (model.Schedule, error)
}
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (model.SearchParameter, error)
}
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (model.ServiceRequest, error)
}
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (model.Slot, error)
}
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (model.Specimen, error)
}
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (model.SpecimenDefinition, error)
}
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (model.StructureDefinition, error)
}
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (model.StructureMap, error)
}
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (model.Subscription, error)
}
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (model.Substance, error)
}
type SubstanceNucleicAcidRead interface {
	ReadSubstanceNucleicAcid(ctx context.Context, id string) (model.SubstanceNucleicAcid, error)
}
type SubstancePolymerRead interface {
	ReadSubstancePolymer(ctx context.Context, id string) (model.SubstancePolymer, error)
}
type SubstanceProteinRead interface {
	ReadSubstanceProtein(ctx context.Context, id string) (model.SubstanceProtein, error)
}
type SubstanceReferenceInformationRead interface {
	ReadSubstanceReferenceInformation(ctx context.Context, id string) (model.SubstanceReferenceInformation, error)
}
type SubstanceSourceMaterialRead interface {
	ReadSubstanceSourceMaterial(ctx context.Context, id string) (model.SubstanceSourceMaterial, error)
}
type SubstanceSpecificationRead interface {
	ReadSubstanceSpecification(ctx context.Context, id string) (model.SubstanceSpecification, error)
}
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (model.SupplyDelivery, error)
}
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (model.SupplyRequest, error)
}
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (model.Task, error)
}
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (model.TerminologyCapabilities, error)
}
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (model.TestReport, error)
}
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (model.TestScript, error)
}
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (model.ValueSet, error)
}
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (model.VerificationResult, error)
}
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (model.VisionPrescription, error)
}
