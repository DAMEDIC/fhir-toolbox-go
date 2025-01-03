package capabilitiesR4

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	r4 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
)

type AccountRead interface {
	ReadAccount(ctx context.Context, id string) (r4.Account, capabilities.FHIRError)
}
type ActivityDefinitionRead interface {
	ReadActivityDefinition(ctx context.Context, id string) (r4.ActivityDefinition, capabilities.FHIRError)
}
type AdverseEventRead interface {
	ReadAdverseEvent(ctx context.Context, id string) (r4.AdverseEvent, capabilities.FHIRError)
}
type AllergyIntoleranceRead interface {
	ReadAllergyIntolerance(ctx context.Context, id string) (r4.AllergyIntolerance, capabilities.FHIRError)
}
type AppointmentRead interface {
	ReadAppointment(ctx context.Context, id string) (r4.Appointment, capabilities.FHIRError)
}
type AppointmentResponseRead interface {
	ReadAppointmentResponse(ctx context.Context, id string) (r4.AppointmentResponse, capabilities.FHIRError)
}
type AuditEventRead interface {
	ReadAuditEvent(ctx context.Context, id string) (r4.AuditEvent, capabilities.FHIRError)
}
type BasicRead interface {
	ReadBasic(ctx context.Context, id string) (r4.Basic, capabilities.FHIRError)
}
type BinaryRead interface {
	ReadBinary(ctx context.Context, id string) (r4.Binary, capabilities.FHIRError)
}
type BiologicallyDerivedProductRead interface {
	ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4.BiologicallyDerivedProduct, capabilities.FHIRError)
}
type BodyStructureRead interface {
	ReadBodyStructure(ctx context.Context, id string) (r4.BodyStructure, capabilities.FHIRError)
}
type BundleRead interface {
	ReadBundle(ctx context.Context, id string) (r4.Bundle, capabilities.FHIRError)
}
type CapabilityStatementRead interface {
	ReadCapabilityStatement(ctx context.Context, id string) (r4.CapabilityStatement, capabilities.FHIRError)
}
type CarePlanRead interface {
	ReadCarePlan(ctx context.Context, id string) (r4.CarePlan, capabilities.FHIRError)
}
type CareTeamRead interface {
	ReadCareTeam(ctx context.Context, id string) (r4.CareTeam, capabilities.FHIRError)
}
type CatalogEntryRead interface {
	ReadCatalogEntry(ctx context.Context, id string) (r4.CatalogEntry, capabilities.FHIRError)
}
type ChargeItemRead interface {
	ReadChargeItem(ctx context.Context, id string) (r4.ChargeItem, capabilities.FHIRError)
}
type ChargeItemDefinitionRead interface {
	ReadChargeItemDefinition(ctx context.Context, id string) (r4.ChargeItemDefinition, capabilities.FHIRError)
}
type ClaimRead interface {
	ReadClaim(ctx context.Context, id string) (r4.Claim, capabilities.FHIRError)
}
type ClaimResponseRead interface {
	ReadClaimResponse(ctx context.Context, id string) (r4.ClaimResponse, capabilities.FHIRError)
}
type ClinicalImpressionRead interface {
	ReadClinicalImpression(ctx context.Context, id string) (r4.ClinicalImpression, capabilities.FHIRError)
}
type CodeSystemRead interface {
	ReadCodeSystem(ctx context.Context, id string) (r4.CodeSystem, capabilities.FHIRError)
}
type CommunicationRead interface {
	ReadCommunication(ctx context.Context, id string) (r4.Communication, capabilities.FHIRError)
}
type CommunicationRequestRead interface {
	ReadCommunicationRequest(ctx context.Context, id string) (r4.CommunicationRequest, capabilities.FHIRError)
}
type CompartmentDefinitionRead interface {
	ReadCompartmentDefinition(ctx context.Context, id string) (r4.CompartmentDefinition, capabilities.FHIRError)
}
type CompositionRead interface {
	ReadComposition(ctx context.Context, id string) (r4.Composition, capabilities.FHIRError)
}
type ConceptMapRead interface {
	ReadConceptMap(ctx context.Context, id string) (r4.ConceptMap, capabilities.FHIRError)
}
type ConditionRead interface {
	ReadCondition(ctx context.Context, id string) (r4.Condition, capabilities.FHIRError)
}
type ConsentRead interface {
	ReadConsent(ctx context.Context, id string) (r4.Consent, capabilities.FHIRError)
}
type ContractRead interface {
	ReadContract(ctx context.Context, id string) (r4.Contract, capabilities.FHIRError)
}
type CoverageRead interface {
	ReadCoverage(ctx context.Context, id string) (r4.Coverage, capabilities.FHIRError)
}
type CoverageEligibilityRequestRead interface {
	ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4.CoverageEligibilityRequest, capabilities.FHIRError)
}
type CoverageEligibilityResponseRead interface {
	ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4.CoverageEligibilityResponse, capabilities.FHIRError)
}
type DetectedIssueRead interface {
	ReadDetectedIssue(ctx context.Context, id string) (r4.DetectedIssue, capabilities.FHIRError)
}
type DeviceRead interface {
	ReadDevice(ctx context.Context, id string) (r4.Device, capabilities.FHIRError)
}
type DeviceDefinitionRead interface {
	ReadDeviceDefinition(ctx context.Context, id string) (r4.DeviceDefinition, capabilities.FHIRError)
}
type DeviceMetricRead interface {
	ReadDeviceMetric(ctx context.Context, id string) (r4.DeviceMetric, capabilities.FHIRError)
}
type DeviceRequestRead interface {
	ReadDeviceRequest(ctx context.Context, id string) (r4.DeviceRequest, capabilities.FHIRError)
}
type DeviceUseStatementRead interface {
	ReadDeviceUseStatement(ctx context.Context, id string) (r4.DeviceUseStatement, capabilities.FHIRError)
}
type DiagnosticReportRead interface {
	ReadDiagnosticReport(ctx context.Context, id string) (r4.DiagnosticReport, capabilities.FHIRError)
}
type DocumentManifestRead interface {
	ReadDocumentManifest(ctx context.Context, id string) (r4.DocumentManifest, capabilities.FHIRError)
}
type DocumentReferenceRead interface {
	ReadDocumentReference(ctx context.Context, id string) (r4.DocumentReference, capabilities.FHIRError)
}
type EffectEvidenceSynthesisRead interface {
	ReadEffectEvidenceSynthesis(ctx context.Context, id string) (r4.EffectEvidenceSynthesis, capabilities.FHIRError)
}
type EncounterRead interface {
	ReadEncounter(ctx context.Context, id string) (r4.Encounter, capabilities.FHIRError)
}
type EndpointRead interface {
	ReadEndpoint(ctx context.Context, id string) (r4.Endpoint, capabilities.FHIRError)
}
type EnrollmentRequestRead interface {
	ReadEnrollmentRequest(ctx context.Context, id string) (r4.EnrollmentRequest, capabilities.FHIRError)
}
type EnrollmentResponseRead interface {
	ReadEnrollmentResponse(ctx context.Context, id string) (r4.EnrollmentResponse, capabilities.FHIRError)
}
type EpisodeOfCareRead interface {
	ReadEpisodeOfCare(ctx context.Context, id string) (r4.EpisodeOfCare, capabilities.FHIRError)
}
type EventDefinitionRead interface {
	ReadEventDefinition(ctx context.Context, id string) (r4.EventDefinition, capabilities.FHIRError)
}
type EvidenceRead interface {
	ReadEvidence(ctx context.Context, id string) (r4.Evidence, capabilities.FHIRError)
}
type EvidenceVariableRead interface {
	ReadEvidenceVariable(ctx context.Context, id string) (r4.EvidenceVariable, capabilities.FHIRError)
}
type ExampleScenarioRead interface {
	ReadExampleScenario(ctx context.Context, id string) (r4.ExampleScenario, capabilities.FHIRError)
}
type ExplanationOfBenefitRead interface {
	ReadExplanationOfBenefit(ctx context.Context, id string) (r4.ExplanationOfBenefit, capabilities.FHIRError)
}
type FamilyMemberHistoryRead interface {
	ReadFamilyMemberHistory(ctx context.Context, id string) (r4.FamilyMemberHistory, capabilities.FHIRError)
}
type FlagRead interface {
	ReadFlag(ctx context.Context, id string) (r4.Flag, capabilities.FHIRError)
}
type GoalRead interface {
	ReadGoal(ctx context.Context, id string) (r4.Goal, capabilities.FHIRError)
}
type GraphDefinitionRead interface {
	ReadGraphDefinition(ctx context.Context, id string) (r4.GraphDefinition, capabilities.FHIRError)
}
type GroupRead interface {
	ReadGroup(ctx context.Context, id string) (r4.Group, capabilities.FHIRError)
}
type GuidanceResponseRead interface {
	ReadGuidanceResponse(ctx context.Context, id string) (r4.GuidanceResponse, capabilities.FHIRError)
}
type HealthcareServiceRead interface {
	ReadHealthcareService(ctx context.Context, id string) (r4.HealthcareService, capabilities.FHIRError)
}
type ImagingStudyRead interface {
	ReadImagingStudy(ctx context.Context, id string) (r4.ImagingStudy, capabilities.FHIRError)
}
type ImmunizationRead interface {
	ReadImmunization(ctx context.Context, id string) (r4.Immunization, capabilities.FHIRError)
}
type ImmunizationEvaluationRead interface {
	ReadImmunizationEvaluation(ctx context.Context, id string) (r4.ImmunizationEvaluation, capabilities.FHIRError)
}
type ImmunizationRecommendationRead interface {
	ReadImmunizationRecommendation(ctx context.Context, id string) (r4.ImmunizationRecommendation, capabilities.FHIRError)
}
type ImplementationGuideRead interface {
	ReadImplementationGuide(ctx context.Context, id string) (r4.ImplementationGuide, capabilities.FHIRError)
}
type InsurancePlanRead interface {
	ReadInsurancePlan(ctx context.Context, id string) (r4.InsurancePlan, capabilities.FHIRError)
}
type InvoiceRead interface {
	ReadInvoice(ctx context.Context, id string) (r4.Invoice, capabilities.FHIRError)
}
type LibraryRead interface {
	ReadLibrary(ctx context.Context, id string) (r4.Library, capabilities.FHIRError)
}
type LinkageRead interface {
	ReadLinkage(ctx context.Context, id string) (r4.Linkage, capabilities.FHIRError)
}
type ListRead interface {
	ReadList(ctx context.Context, id string) (r4.List, capabilities.FHIRError)
}
type LocationRead interface {
	ReadLocation(ctx context.Context, id string) (r4.Location, capabilities.FHIRError)
}
type MeasureRead interface {
	ReadMeasure(ctx context.Context, id string) (r4.Measure, capabilities.FHIRError)
}
type MeasureReportRead interface {
	ReadMeasureReport(ctx context.Context, id string) (r4.MeasureReport, capabilities.FHIRError)
}
type MediaRead interface {
	ReadMedia(ctx context.Context, id string) (r4.Media, capabilities.FHIRError)
}
type MedicationRead interface {
	ReadMedication(ctx context.Context, id string) (r4.Medication, capabilities.FHIRError)
}
type MedicationAdministrationRead interface {
	ReadMedicationAdministration(ctx context.Context, id string) (r4.MedicationAdministration, capabilities.FHIRError)
}
type MedicationDispenseRead interface {
	ReadMedicationDispense(ctx context.Context, id string) (r4.MedicationDispense, capabilities.FHIRError)
}
type MedicationKnowledgeRead interface {
	ReadMedicationKnowledge(ctx context.Context, id string) (r4.MedicationKnowledge, capabilities.FHIRError)
}
type MedicationRequestRead interface {
	ReadMedicationRequest(ctx context.Context, id string) (r4.MedicationRequest, capabilities.FHIRError)
}
type MedicationStatementRead interface {
	ReadMedicationStatement(ctx context.Context, id string) (r4.MedicationStatement, capabilities.FHIRError)
}
type MedicinalProductRead interface {
	ReadMedicinalProduct(ctx context.Context, id string) (r4.MedicinalProduct, capabilities.FHIRError)
}
type MedicinalProductAuthorizationRead interface {
	ReadMedicinalProductAuthorization(ctx context.Context, id string) (r4.MedicinalProductAuthorization, capabilities.FHIRError)
}
type MedicinalProductContraindicationRead interface {
	ReadMedicinalProductContraindication(ctx context.Context, id string) (r4.MedicinalProductContraindication, capabilities.FHIRError)
}
type MedicinalProductIndicationRead interface {
	ReadMedicinalProductIndication(ctx context.Context, id string) (r4.MedicinalProductIndication, capabilities.FHIRError)
}
type MedicinalProductIngredientRead interface {
	ReadMedicinalProductIngredient(ctx context.Context, id string) (r4.MedicinalProductIngredient, capabilities.FHIRError)
}
type MedicinalProductInteractionRead interface {
	ReadMedicinalProductInteraction(ctx context.Context, id string) (r4.MedicinalProductInteraction, capabilities.FHIRError)
}
type MedicinalProductManufacturedRead interface {
	ReadMedicinalProductManufactured(ctx context.Context, id string) (r4.MedicinalProductManufactured, capabilities.FHIRError)
}
type MedicinalProductPackagedRead interface {
	ReadMedicinalProductPackaged(ctx context.Context, id string) (r4.MedicinalProductPackaged, capabilities.FHIRError)
}
type MedicinalProductPharmaceuticalRead interface {
	ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (r4.MedicinalProductPharmaceutical, capabilities.FHIRError)
}
type MedicinalProductUndesirableEffectRead interface {
	ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (r4.MedicinalProductUndesirableEffect, capabilities.FHIRError)
}
type MessageDefinitionRead interface {
	ReadMessageDefinition(ctx context.Context, id string) (r4.MessageDefinition, capabilities.FHIRError)
}
type MessageHeaderRead interface {
	ReadMessageHeader(ctx context.Context, id string) (r4.MessageHeader, capabilities.FHIRError)
}
type MolecularSequenceRead interface {
	ReadMolecularSequence(ctx context.Context, id string) (r4.MolecularSequence, capabilities.FHIRError)
}
type NamingSystemRead interface {
	ReadNamingSystem(ctx context.Context, id string) (r4.NamingSystem, capabilities.FHIRError)
}
type NutritionOrderRead interface {
	ReadNutritionOrder(ctx context.Context, id string) (r4.NutritionOrder, capabilities.FHIRError)
}
type ObservationRead interface {
	ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError)
}
type ObservationDefinitionRead interface {
	ReadObservationDefinition(ctx context.Context, id string) (r4.ObservationDefinition, capabilities.FHIRError)
}
type OperationDefinitionRead interface {
	ReadOperationDefinition(ctx context.Context, id string) (r4.OperationDefinition, capabilities.FHIRError)
}
type OperationOutcomeRead interface {
	ReadOperationOutcome(ctx context.Context, id string) (r4.OperationOutcome, capabilities.FHIRError)
}
type OrganizationRead interface {
	ReadOrganization(ctx context.Context, id string) (r4.Organization, capabilities.FHIRError)
}
type OrganizationAffiliationRead interface {
	ReadOrganizationAffiliation(ctx context.Context, id string) (r4.OrganizationAffiliation, capabilities.FHIRError)
}
type ParametersRead interface {
	ReadParameters(ctx context.Context, id string) (r4.Parameters, capabilities.FHIRError)
}
type PatientRead interface {
	ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError)
}
type PaymentNoticeRead interface {
	ReadPaymentNotice(ctx context.Context, id string) (r4.PaymentNotice, capabilities.FHIRError)
}
type PaymentReconciliationRead interface {
	ReadPaymentReconciliation(ctx context.Context, id string) (r4.PaymentReconciliation, capabilities.FHIRError)
}
type PersonRead interface {
	ReadPerson(ctx context.Context, id string) (r4.Person, capabilities.FHIRError)
}
type PlanDefinitionRead interface {
	ReadPlanDefinition(ctx context.Context, id string) (r4.PlanDefinition, capabilities.FHIRError)
}
type PractitionerRead interface {
	ReadPractitioner(ctx context.Context, id string) (r4.Practitioner, capabilities.FHIRError)
}
type PractitionerRoleRead interface {
	ReadPractitionerRole(ctx context.Context, id string) (r4.PractitionerRole, capabilities.FHIRError)
}
type ProcedureRead interface {
	ReadProcedure(ctx context.Context, id string) (r4.Procedure, capabilities.FHIRError)
}
type ProvenanceRead interface {
	ReadProvenance(ctx context.Context, id string) (r4.Provenance, capabilities.FHIRError)
}
type QuestionnaireRead interface {
	ReadQuestionnaire(ctx context.Context, id string) (r4.Questionnaire, capabilities.FHIRError)
}
type QuestionnaireResponseRead interface {
	ReadQuestionnaireResponse(ctx context.Context, id string) (r4.QuestionnaireResponse, capabilities.FHIRError)
}
type RelatedPersonRead interface {
	ReadRelatedPerson(ctx context.Context, id string) (r4.RelatedPerson, capabilities.FHIRError)
}
type RequestGroupRead interface {
	ReadRequestGroup(ctx context.Context, id string) (r4.RequestGroup, capabilities.FHIRError)
}
type ResearchDefinitionRead interface {
	ReadResearchDefinition(ctx context.Context, id string) (r4.ResearchDefinition, capabilities.FHIRError)
}
type ResearchElementDefinitionRead interface {
	ReadResearchElementDefinition(ctx context.Context, id string) (r4.ResearchElementDefinition, capabilities.FHIRError)
}
type ResearchStudyRead interface {
	ReadResearchStudy(ctx context.Context, id string) (r4.ResearchStudy, capabilities.FHIRError)
}
type ResearchSubjectRead interface {
	ReadResearchSubject(ctx context.Context, id string) (r4.ResearchSubject, capabilities.FHIRError)
}
type RiskAssessmentRead interface {
	ReadRiskAssessment(ctx context.Context, id string) (r4.RiskAssessment, capabilities.FHIRError)
}
type RiskEvidenceSynthesisRead interface {
	ReadRiskEvidenceSynthesis(ctx context.Context, id string) (r4.RiskEvidenceSynthesis, capabilities.FHIRError)
}
type ScheduleRead interface {
	ReadSchedule(ctx context.Context, id string) (r4.Schedule, capabilities.FHIRError)
}
type SearchParameterRead interface {
	ReadSearchParameter(ctx context.Context, id string) (r4.SearchParameter, capabilities.FHIRError)
}
type ServiceRequestRead interface {
	ReadServiceRequest(ctx context.Context, id string) (r4.ServiceRequest, capabilities.FHIRError)
}
type SlotRead interface {
	ReadSlot(ctx context.Context, id string) (r4.Slot, capabilities.FHIRError)
}
type SpecimenRead interface {
	ReadSpecimen(ctx context.Context, id string) (r4.Specimen, capabilities.FHIRError)
}
type SpecimenDefinitionRead interface {
	ReadSpecimenDefinition(ctx context.Context, id string) (r4.SpecimenDefinition, capabilities.FHIRError)
}
type StructureDefinitionRead interface {
	ReadStructureDefinition(ctx context.Context, id string) (r4.StructureDefinition, capabilities.FHIRError)
}
type StructureMapRead interface {
	ReadStructureMap(ctx context.Context, id string) (r4.StructureMap, capabilities.FHIRError)
}
type SubscriptionRead interface {
	ReadSubscription(ctx context.Context, id string) (r4.Subscription, capabilities.FHIRError)
}
type SubstanceRead interface {
	ReadSubstance(ctx context.Context, id string) (r4.Substance, capabilities.FHIRError)
}
type SubstanceNucleicAcidRead interface {
	ReadSubstanceNucleicAcid(ctx context.Context, id string) (r4.SubstanceNucleicAcid, capabilities.FHIRError)
}
type SubstancePolymerRead interface {
	ReadSubstancePolymer(ctx context.Context, id string) (r4.SubstancePolymer, capabilities.FHIRError)
}
type SubstanceProteinRead interface {
	ReadSubstanceProtein(ctx context.Context, id string) (r4.SubstanceProtein, capabilities.FHIRError)
}
type SubstanceReferenceInformationRead interface {
	ReadSubstanceReferenceInformation(ctx context.Context, id string) (r4.SubstanceReferenceInformation, capabilities.FHIRError)
}
type SubstanceSourceMaterialRead interface {
	ReadSubstanceSourceMaterial(ctx context.Context, id string) (r4.SubstanceSourceMaterial, capabilities.FHIRError)
}
type SubstanceSpecificationRead interface {
	ReadSubstanceSpecification(ctx context.Context, id string) (r4.SubstanceSpecification, capabilities.FHIRError)
}
type SupplyDeliveryRead interface {
	ReadSupplyDelivery(ctx context.Context, id string) (r4.SupplyDelivery, capabilities.FHIRError)
}
type SupplyRequestRead interface {
	ReadSupplyRequest(ctx context.Context, id string) (r4.SupplyRequest, capabilities.FHIRError)
}
type TaskRead interface {
	ReadTask(ctx context.Context, id string) (r4.Task, capabilities.FHIRError)
}
type TerminologyCapabilitiesRead interface {
	ReadTerminologyCapabilities(ctx context.Context, id string) (r4.TerminologyCapabilities, capabilities.FHIRError)
}
type TestReportRead interface {
	ReadTestReport(ctx context.Context, id string) (r4.TestReport, capabilities.FHIRError)
}
type TestScriptRead interface {
	ReadTestScript(ctx context.Context, id string) (r4.TestScript, capabilities.FHIRError)
}
type ValueSetRead interface {
	ReadValueSet(ctx context.Context, id string) (r4.ValueSet, capabilities.FHIRError)
}
type VerificationResultRead interface {
	ReadVerificationResult(ctx context.Context, id string) (r4.VerificationResult, capabilities.FHIRError)
}
type VisionPrescriptionRead interface {
	ReadVisionPrescription(ctx context.Context, id string) (r4.VisionPrescription, capabilities.FHIRError)
}
