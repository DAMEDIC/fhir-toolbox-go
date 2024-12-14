package capabilitiesR5

import (
	"context"
	capabilities "github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	search "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
)

type Generic struct {
	Concrete any
}

func (w Generic) AllCapabilities() capabilities.Capabilities {
	return AllCapabilities(w.Concrete)
}
func (w Generic) Read(ctx context.Context, resourceType string, id string) (model.Resource, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.Concrete.(AccountRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Account"}
		}
		return impl.ReadAccount(ctx, id)
	case "ActivityDefinition":
		impl, ok := w.Concrete.(ActivityDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ActivityDefinition"}
		}
		return impl.ReadActivityDefinition(ctx, id)
	case "ActorDefinition":
		impl, ok := w.Concrete.(ActorDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ActorDefinition"}
		}
		return impl.ReadActorDefinition(ctx, id)
	case "AdministrableProductDefinition":
		impl, ok := w.Concrete.(AdministrableProductDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AdministrableProductDefinition"}
		}
		return impl.ReadAdministrableProductDefinition(ctx, id)
	case "AdverseEvent":
		impl, ok := w.Concrete.(AdverseEventRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AdverseEvent"}
		}
		return impl.ReadAdverseEvent(ctx, id)
	case "AllergyIntolerance":
		impl, ok := w.Concrete.(AllergyIntoleranceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AllergyIntolerance"}
		}
		return impl.ReadAllergyIntolerance(ctx, id)
	case "Appointment":
		impl, ok := w.Concrete.(AppointmentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Appointment"}
		}
		return impl.ReadAppointment(ctx, id)
	case "AppointmentResponse":
		impl, ok := w.Concrete.(AppointmentResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AppointmentResponse"}
		}
		return impl.ReadAppointmentResponse(ctx, id)
	case "ArtifactAssessment":
		impl, ok := w.Concrete.(ArtifactAssessmentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ArtifactAssessment"}
		}
		return impl.ReadArtifactAssessment(ctx, id)
	case "AuditEvent":
		impl, ok := w.Concrete.(AuditEventRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AuditEvent"}
		}
		return impl.ReadAuditEvent(ctx, id)
	case "Basic":
		impl, ok := w.Concrete.(BasicRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Basic"}
		}
		return impl.ReadBasic(ctx, id)
	case "Binary":
		impl, ok := w.Concrete.(BinaryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Binary"}
		}
		return impl.ReadBinary(ctx, id)
	case "BiologicallyDerivedProduct":
		impl, ok := w.Concrete.(BiologicallyDerivedProductRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.ReadBiologicallyDerivedProduct(ctx, id)
	case "BiologicallyDerivedProductDispense":
		impl, ok := w.Concrete.(BiologicallyDerivedProductDispenseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "BiologicallyDerivedProductDispense"}
		}
		return impl.ReadBiologicallyDerivedProductDispense(ctx, id)
	case "BodyStructure":
		impl, ok := w.Concrete.(BodyStructureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "BodyStructure"}
		}
		return impl.ReadBodyStructure(ctx, id)
	case "Bundle":
		impl, ok := w.Concrete.(BundleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Bundle"}
		}
		return impl.ReadBundle(ctx, id)
	case "CapabilityStatement":
		impl, ok := w.Concrete.(CapabilityStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CapabilityStatement"}
		}
		return impl.ReadCapabilityStatement(ctx, id)
	case "CarePlan":
		impl, ok := w.Concrete.(CarePlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CarePlan"}
		}
		return impl.ReadCarePlan(ctx, id)
	case "CareTeam":
		impl, ok := w.Concrete.(CareTeamRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CareTeam"}
		}
		return impl.ReadCareTeam(ctx, id)
	case "ChargeItem":
		impl, ok := w.Concrete.(ChargeItemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ChargeItem"}
		}
		return impl.ReadChargeItem(ctx, id)
	case "ChargeItemDefinition":
		impl, ok := w.Concrete.(ChargeItemDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ChargeItemDefinition"}
		}
		return impl.ReadChargeItemDefinition(ctx, id)
	case "Citation":
		impl, ok := w.Concrete.(CitationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Citation"}
		}
		return impl.ReadCitation(ctx, id)
	case "Claim":
		impl, ok := w.Concrete.(ClaimRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Claim"}
		}
		return impl.ReadClaim(ctx, id)
	case "ClaimResponse":
		impl, ok := w.Concrete.(ClaimResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ClaimResponse"}
		}
		return impl.ReadClaimResponse(ctx, id)
	case "ClinicalImpression":
		impl, ok := w.Concrete.(ClinicalImpressionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ClinicalImpression"}
		}
		return impl.ReadClinicalImpression(ctx, id)
	case "ClinicalUseDefinition":
		impl, ok := w.Concrete.(ClinicalUseDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ClinicalUseDefinition"}
		}
		return impl.ReadClinicalUseDefinition(ctx, id)
	case "CodeSystem":
		impl, ok := w.Concrete.(CodeSystemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CodeSystem"}
		}
		return impl.ReadCodeSystem(ctx, id)
	case "Communication":
		impl, ok := w.Concrete.(CommunicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Communication"}
		}
		return impl.ReadCommunication(ctx, id)
	case "CommunicationRequest":
		impl, ok := w.Concrete.(CommunicationRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CommunicationRequest"}
		}
		return impl.ReadCommunicationRequest(ctx, id)
	case "CompartmentDefinition":
		impl, ok := w.Concrete.(CompartmentDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CompartmentDefinition"}
		}
		return impl.ReadCompartmentDefinition(ctx, id)
	case "Composition":
		impl, ok := w.Concrete.(CompositionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Composition"}
		}
		return impl.ReadComposition(ctx, id)
	case "ConceptMap":
		impl, ok := w.Concrete.(ConceptMapRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ConceptMap"}
		}
		return impl.ReadConceptMap(ctx, id)
	case "Condition":
		impl, ok := w.Concrete.(ConditionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Condition"}
		}
		return impl.ReadCondition(ctx, id)
	case "ConditionDefinition":
		impl, ok := w.Concrete.(ConditionDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ConditionDefinition"}
		}
		return impl.ReadConditionDefinition(ctx, id)
	case "Consent":
		impl, ok := w.Concrete.(ConsentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Consent"}
		}
		return impl.ReadConsent(ctx, id)
	case "Contract":
		impl, ok := w.Concrete.(ContractRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Contract"}
		}
		return impl.ReadContract(ctx, id)
	case "Coverage":
		impl, ok := w.Concrete.(CoverageRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Coverage"}
		}
		return impl.ReadCoverage(ctx, id)
	case "CoverageEligibilityRequest":
		impl, ok := w.Concrete.(CoverageEligibilityRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.ReadCoverageEligibilityRequest(ctx, id)
	case "CoverageEligibilityResponse":
		impl, ok := w.Concrete.(CoverageEligibilityResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.ReadCoverageEligibilityResponse(ctx, id)
	case "DetectedIssue":
		impl, ok := w.Concrete.(DetectedIssueRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DetectedIssue"}
		}
		return impl.ReadDetectedIssue(ctx, id)
	case "Device":
		impl, ok := w.Concrete.(DeviceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Device"}
		}
		return impl.ReadDevice(ctx, id)
	case "DeviceAssociation":
		impl, ok := w.Concrete.(DeviceAssociationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceAssociation"}
		}
		return impl.ReadDeviceAssociation(ctx, id)
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceDefinition"}
		}
		return impl.ReadDeviceDefinition(ctx, id)
	case "DeviceDispense":
		impl, ok := w.Concrete.(DeviceDispenseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceDispense"}
		}
		return impl.ReadDeviceDispense(ctx, id)
	case "DeviceMetric":
		impl, ok := w.Concrete.(DeviceMetricRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceMetric"}
		}
		return impl.ReadDeviceMetric(ctx, id)
	case "DeviceRequest":
		impl, ok := w.Concrete.(DeviceRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceRequest"}
		}
		return impl.ReadDeviceRequest(ctx, id)
	case "DeviceUsage":
		impl, ok := w.Concrete.(DeviceUsageRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceUsage"}
		}
		return impl.ReadDeviceUsage(ctx, id)
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DiagnosticReport"}
		}
		return impl.ReadDiagnosticReport(ctx, id)
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DocumentReference"}
		}
		return impl.ReadDocumentReference(ctx, id)
	case "Encounter":
		impl, ok := w.Concrete.(EncounterRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Encounter"}
		}
		return impl.ReadEncounter(ctx, id)
	case "EncounterHistory":
		impl, ok := w.Concrete.(EncounterHistoryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EncounterHistory"}
		}
		return impl.ReadEncounterHistory(ctx, id)
	case "Endpoint":
		impl, ok := w.Concrete.(EndpointRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Endpoint"}
		}
		return impl.ReadEndpoint(ctx, id)
	case "EnrollmentRequest":
		impl, ok := w.Concrete.(EnrollmentRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EnrollmentRequest"}
		}
		return impl.ReadEnrollmentRequest(ctx, id)
	case "EnrollmentResponse":
		impl, ok := w.Concrete.(EnrollmentResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EnrollmentResponse"}
		}
		return impl.ReadEnrollmentResponse(ctx, id)
	case "EpisodeOfCare":
		impl, ok := w.Concrete.(EpisodeOfCareRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EpisodeOfCare"}
		}
		return impl.ReadEpisodeOfCare(ctx, id)
	case "EventDefinition":
		impl, ok := w.Concrete.(EventDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EventDefinition"}
		}
		return impl.ReadEventDefinition(ctx, id)
	case "Evidence":
		impl, ok := w.Concrete.(EvidenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Evidence"}
		}
		return impl.ReadEvidence(ctx, id)
	case "EvidenceReport":
		impl, ok := w.Concrete.(EvidenceReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EvidenceReport"}
		}
		return impl.ReadEvidenceReport(ctx, id)
	case "EvidenceVariable":
		impl, ok := w.Concrete.(EvidenceVariableRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EvidenceVariable"}
		}
		return impl.ReadEvidenceVariable(ctx, id)
	case "ExampleScenario":
		impl, ok := w.Concrete.(ExampleScenarioRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ExampleScenario"}
		}
		return impl.ReadExampleScenario(ctx, id)
	case "ExplanationOfBenefit":
		impl, ok := w.Concrete.(ExplanationOfBenefitRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.ReadExplanationOfBenefit(ctx, id)
	case "FamilyMemberHistory":
		impl, ok := w.Concrete.(FamilyMemberHistoryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "FamilyMemberHistory"}
		}
		return impl.ReadFamilyMemberHistory(ctx, id)
	case "Flag":
		impl, ok := w.Concrete.(FlagRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Flag"}
		}
		return impl.ReadFlag(ctx, id)
	case "FormularyItem":
		impl, ok := w.Concrete.(FormularyItemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "FormularyItem"}
		}
		return impl.ReadFormularyItem(ctx, id)
	case "GenomicStudy":
		impl, ok := w.Concrete.(GenomicStudyRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "GenomicStudy"}
		}
		return impl.ReadGenomicStudy(ctx, id)
	case "Goal":
		impl, ok := w.Concrete.(GoalRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Goal"}
		}
		return impl.ReadGoal(ctx, id)
	case "GraphDefinition":
		impl, ok := w.Concrete.(GraphDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "GraphDefinition"}
		}
		return impl.ReadGraphDefinition(ctx, id)
	case "Group":
		impl, ok := w.Concrete.(GroupRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Group"}
		}
		return impl.ReadGroup(ctx, id)
	case "GuidanceResponse":
		impl, ok := w.Concrete.(GuidanceResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "GuidanceResponse"}
		}
		return impl.ReadGuidanceResponse(ctx, id)
	case "HealthcareService":
		impl, ok := w.Concrete.(HealthcareServiceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "HealthcareService"}
		}
		return impl.ReadHealthcareService(ctx, id)
	case "ImagingSelection":
		impl, ok := w.Concrete.(ImagingSelectionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImagingSelection"}
		}
		return impl.ReadImagingSelection(ctx, id)
	case "ImagingStudy":
		impl, ok := w.Concrete.(ImagingStudyRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImagingStudy"}
		}
		return impl.ReadImagingStudy(ctx, id)
	case "Immunization":
		impl, ok := w.Concrete.(ImmunizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Immunization"}
		}
		return impl.ReadImmunization(ctx, id)
	case "ImmunizationEvaluation":
		impl, ok := w.Concrete.(ImmunizationEvaluationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.ReadImmunizationEvaluation(ctx, id)
	case "ImmunizationRecommendation":
		impl, ok := w.Concrete.(ImmunizationRecommendationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.ReadImmunizationRecommendation(ctx, id)
	case "ImplementationGuide":
		impl, ok := w.Concrete.(ImplementationGuideRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImplementationGuide"}
		}
		return impl.ReadImplementationGuide(ctx, id)
	case "Ingredient":
		impl, ok := w.Concrete.(IngredientRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Ingredient"}
		}
		return impl.ReadIngredient(ctx, id)
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "InsurancePlan"}
		}
		return impl.ReadInsurancePlan(ctx, id)
	case "InventoryItem":
		impl, ok := w.Concrete.(InventoryItemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "InventoryItem"}
		}
		return impl.ReadInventoryItem(ctx, id)
	case "InventoryReport":
		impl, ok := w.Concrete.(InventoryReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "InventoryReport"}
		}
		return impl.ReadInventoryReport(ctx, id)
	case "Invoice":
		impl, ok := w.Concrete.(InvoiceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Invoice"}
		}
		return impl.ReadInvoice(ctx, id)
	case "Library":
		impl, ok := w.Concrete.(LibraryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Library"}
		}
		return impl.ReadLibrary(ctx, id)
	case "Linkage":
		impl, ok := w.Concrete.(LinkageRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Linkage"}
		}
		return impl.ReadLinkage(ctx, id)
	case "List":
		impl, ok := w.Concrete.(ListRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "List"}
		}
		return impl.ReadList(ctx, id)
	case "Location":
		impl, ok := w.Concrete.(LocationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Location"}
		}
		return impl.ReadLocation(ctx, id)
	case "ManufacturedItemDefinition":
		impl, ok := w.Concrete.(ManufacturedItemDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ManufacturedItemDefinition"}
		}
		return impl.ReadManufacturedItemDefinition(ctx, id)
	case "Measure":
		impl, ok := w.Concrete.(MeasureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Measure"}
		}
		return impl.ReadMeasure(ctx, id)
	case "MeasureReport":
		impl, ok := w.Concrete.(MeasureReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MeasureReport"}
		}
		return impl.ReadMeasureReport(ctx, id)
	case "Medication":
		impl, ok := w.Concrete.(MedicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Medication"}
		}
		return impl.ReadMedication(ctx, id)
	case "MedicationAdministration":
		impl, ok := w.Concrete.(MedicationAdministrationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationAdministration"}
		}
		return impl.ReadMedicationAdministration(ctx, id)
	case "MedicationDispense":
		impl, ok := w.Concrete.(MedicationDispenseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationDispense"}
		}
		return impl.ReadMedicationDispense(ctx, id)
	case "MedicationKnowledge":
		impl, ok := w.Concrete.(MedicationKnowledgeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationKnowledge"}
		}
		return impl.ReadMedicationKnowledge(ctx, id)
	case "MedicationRequest":
		impl, ok := w.Concrete.(MedicationRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationRequest"}
		}
		return impl.ReadMedicationRequest(ctx, id)
	case "MedicationStatement":
		impl, ok := w.Concrete.(MedicationStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationStatement"}
		}
		return impl.ReadMedicationStatement(ctx, id)
	case "MedicinalProductDefinition":
		impl, ok := w.Concrete.(MedicinalProductDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductDefinition"}
		}
		return impl.ReadMedicinalProductDefinition(ctx, id)
	case "MessageDefinition":
		impl, ok := w.Concrete.(MessageDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MessageDefinition"}
		}
		return impl.ReadMessageDefinition(ctx, id)
	case "MessageHeader":
		impl, ok := w.Concrete.(MessageHeaderRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MessageHeader"}
		}
		return impl.ReadMessageHeader(ctx, id)
	case "MolecularSequence":
		impl, ok := w.Concrete.(MolecularSequenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MolecularSequence"}
		}
		return impl.ReadMolecularSequence(ctx, id)
	case "NamingSystem":
		impl, ok := w.Concrete.(NamingSystemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NamingSystem"}
		}
		return impl.ReadNamingSystem(ctx, id)
	case "NutritionIntake":
		impl, ok := w.Concrete.(NutritionIntakeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NutritionIntake"}
		}
		return impl.ReadNutritionIntake(ctx, id)
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NutritionOrder"}
		}
		return impl.ReadNutritionOrder(ctx, id)
	case "NutritionProduct":
		impl, ok := w.Concrete.(NutritionProductRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NutritionProduct"}
		}
		return impl.ReadNutritionProduct(ctx, id)
	case "Observation":
		impl, ok := w.Concrete.(ObservationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Observation"}
		}
		return impl.ReadObservation(ctx, id)
	case "ObservationDefinition":
		impl, ok := w.Concrete.(ObservationDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ObservationDefinition"}
		}
		return impl.ReadObservationDefinition(ctx, id)
	case "OperationDefinition":
		impl, ok := w.Concrete.(OperationDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OperationDefinition"}
		}
		return impl.ReadOperationDefinition(ctx, id)
	case "OperationOutcome":
		impl, ok := w.Concrete.(OperationOutcomeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OperationOutcome"}
		}
		return impl.ReadOperationOutcome(ctx, id)
	case "Organization":
		impl, ok := w.Concrete.(OrganizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Organization"}
		}
		return impl.ReadOrganization(ctx, id)
	case "OrganizationAffiliation":
		impl, ok := w.Concrete.(OrganizationAffiliationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OrganizationAffiliation"}
		}
		return impl.ReadOrganizationAffiliation(ctx, id)
	case "PackagedProductDefinition":
		impl, ok := w.Concrete.(PackagedProductDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PackagedProductDefinition"}
		}
		return impl.ReadPackagedProductDefinition(ctx, id)
	case "Parameters":
		impl, ok := w.Concrete.(ParametersRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Parameters"}
		}
		return impl.ReadParameters(ctx, id)
	case "Patient":
		impl, ok := w.Concrete.(PatientRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Patient"}
		}
		return impl.ReadPatient(ctx, id)
	case "PaymentNotice":
		impl, ok := w.Concrete.(PaymentNoticeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PaymentNotice"}
		}
		return impl.ReadPaymentNotice(ctx, id)
	case "PaymentReconciliation":
		impl, ok := w.Concrete.(PaymentReconciliationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PaymentReconciliation"}
		}
		return impl.ReadPaymentReconciliation(ctx, id)
	case "Permission":
		impl, ok := w.Concrete.(PermissionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Permission"}
		}
		return impl.ReadPermission(ctx, id)
	case "Person":
		impl, ok := w.Concrete.(PersonRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Person"}
		}
		return impl.ReadPerson(ctx, id)
	case "PlanDefinition":
		impl, ok := w.Concrete.(PlanDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PlanDefinition"}
		}
		return impl.ReadPlanDefinition(ctx, id)
	case "Practitioner":
		impl, ok := w.Concrete.(PractitionerRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Practitioner"}
		}
		return impl.ReadPractitioner(ctx, id)
	case "PractitionerRole":
		impl, ok := w.Concrete.(PractitionerRoleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PractitionerRole"}
		}
		return impl.ReadPractitionerRole(ctx, id)
	case "Procedure":
		impl, ok := w.Concrete.(ProcedureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Procedure"}
		}
		return impl.ReadProcedure(ctx, id)
	case "Provenance":
		impl, ok := w.Concrete.(ProvenanceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Provenance"}
		}
		return impl.ReadProvenance(ctx, id)
	case "Questionnaire":
		impl, ok := w.Concrete.(QuestionnaireRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Questionnaire"}
		}
		return impl.ReadQuestionnaire(ctx, id)
	case "QuestionnaireResponse":
		impl, ok := w.Concrete.(QuestionnaireResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "QuestionnaireResponse"}
		}
		return impl.ReadQuestionnaireResponse(ctx, id)
	case "RegulatedAuthorization":
		impl, ok := w.Concrete.(RegulatedAuthorizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RegulatedAuthorization"}
		}
		return impl.ReadRegulatedAuthorization(ctx, id)
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RelatedPerson"}
		}
		return impl.ReadRelatedPerson(ctx, id)
	case "RequestOrchestration":
		impl, ok := w.Concrete.(RequestOrchestrationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RequestOrchestration"}
		}
		return impl.ReadRequestOrchestration(ctx, id)
	case "Requirements":
		impl, ok := w.Concrete.(RequirementsRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Requirements"}
		}
		return impl.ReadRequirements(ctx, id)
	case "ResearchStudy":
		impl, ok := w.Concrete.(ResearchStudyRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchStudy"}
		}
		return impl.ReadResearchStudy(ctx, id)
	case "ResearchSubject":
		impl, ok := w.Concrete.(ResearchSubjectRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchSubject"}
		}
		return impl.ReadResearchSubject(ctx, id)
	case "RiskAssessment":
		impl, ok := w.Concrete.(RiskAssessmentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RiskAssessment"}
		}
		return impl.ReadRiskAssessment(ctx, id)
	case "Schedule":
		impl, ok := w.Concrete.(ScheduleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Schedule"}
		}
		return impl.ReadSchedule(ctx, id)
	case "SearchParameter":
		impl, ok := w.Concrete.(SearchParameterRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SearchParameter"}
		}
		return impl.ReadSearchParameter(ctx, id)
	case "ServiceRequest":
		impl, ok := w.Concrete.(ServiceRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ServiceRequest"}
		}
		return impl.ReadServiceRequest(ctx, id)
	case "Slot":
		impl, ok := w.Concrete.(SlotRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Slot"}
		}
		return impl.ReadSlot(ctx, id)
	case "Specimen":
		impl, ok := w.Concrete.(SpecimenRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Specimen"}
		}
		return impl.ReadSpecimen(ctx, id)
	case "SpecimenDefinition":
		impl, ok := w.Concrete.(SpecimenDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SpecimenDefinition"}
		}
		return impl.ReadSpecimenDefinition(ctx, id)
	case "StructureDefinition":
		impl, ok := w.Concrete.(StructureDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "StructureDefinition"}
		}
		return impl.ReadStructureDefinition(ctx, id)
	case "StructureMap":
		impl, ok := w.Concrete.(StructureMapRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "StructureMap"}
		}
		return impl.ReadStructureMap(ctx, id)
	case "Subscription":
		impl, ok := w.Concrete.(SubscriptionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Subscription"}
		}
		return impl.ReadSubscription(ctx, id)
	case "SubscriptionStatus":
		impl, ok := w.Concrete.(SubscriptionStatusRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubscriptionStatus"}
		}
		return impl.ReadSubscriptionStatus(ctx, id)
	case "SubscriptionTopic":
		impl, ok := w.Concrete.(SubscriptionTopicRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubscriptionTopic"}
		}
		return impl.ReadSubscriptionTopic(ctx, id)
	case "Substance":
		impl, ok := w.Concrete.(SubstanceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Substance"}
		}
		return impl.ReadSubstance(ctx, id)
	case "SubstanceDefinition":
		impl, ok := w.Concrete.(SubstanceDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceDefinition"}
		}
		return impl.ReadSubstanceDefinition(ctx, id)
	case "SubstanceNucleicAcid":
		impl, ok := w.Concrete.(SubstanceNucleicAcidRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.ReadSubstanceNucleicAcid(ctx, id)
	case "SubstancePolymer":
		impl, ok := w.Concrete.(SubstancePolymerRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstancePolymer"}
		}
		return impl.ReadSubstancePolymer(ctx, id)
	case "SubstanceProtein":
		impl, ok := w.Concrete.(SubstanceProteinRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceProtein"}
		}
		return impl.ReadSubstanceProtein(ctx, id)
	case "SubstanceReferenceInformation":
		impl, ok := w.Concrete.(SubstanceReferenceInformationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.ReadSubstanceReferenceInformation(ctx, id)
	case "SubstanceSourceMaterial":
		impl, ok := w.Concrete.(SubstanceSourceMaterialRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.ReadSubstanceSourceMaterial(ctx, id)
	case "SupplyDelivery":
		impl, ok := w.Concrete.(SupplyDeliveryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SupplyDelivery"}
		}
		return impl.ReadSupplyDelivery(ctx, id)
	case "SupplyRequest":
		impl, ok := w.Concrete.(SupplyRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SupplyRequest"}
		}
		return impl.ReadSupplyRequest(ctx, id)
	case "Task":
		impl, ok := w.Concrete.(TaskRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Task"}
		}
		return impl.ReadTask(ctx, id)
	case "TerminologyCapabilities":
		impl, ok := w.Concrete.(TerminologyCapabilitiesRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TerminologyCapabilities"}
		}
		return impl.ReadTerminologyCapabilities(ctx, id)
	case "TestPlan":
		impl, ok := w.Concrete.(TestPlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TestPlan"}
		}
		return impl.ReadTestPlan(ctx, id)
	case "TestReport":
		impl, ok := w.Concrete.(TestReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TestReport"}
		}
		return impl.ReadTestReport(ctx, id)
	case "TestScript":
		impl, ok := w.Concrete.(TestScriptRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TestScript"}
		}
		return impl.ReadTestScript(ctx, id)
	case "Transport":
		impl, ok := w.Concrete.(TransportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Transport"}
		}
		return impl.ReadTransport(ctx, id)
	case "ValueSet":
		impl, ok := w.Concrete.(ValueSetRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ValueSet"}
		}
		return impl.ReadValueSet(ctx, id)
	case "VerificationResult":
		impl, ok := w.Concrete.(VerificationResultRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "VerificationResult"}
		}
		return impl.ReadVerificationResult(ctx, id)
	case "VisionPrescription":
		impl, ok := w.Concrete.(VisionPrescriptionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "VisionPrescription"}
		}
		return impl.ReadVisionPrescription(ctx, id)
	default:
		return nil, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
func (w Generic) SearchCapabilities(resourceType string) (search.Capabilities, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.Concrete.(AccountSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Account"}
		}
		return impl.SearchCapabilitiesAccount(), nil
	case "ActivityDefinition":
		impl, ok := w.Concrete.(ActivityDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActivityDefinition"}
		}
		return impl.SearchCapabilitiesActivityDefinition(), nil
	case "ActorDefinition":
		impl, ok := w.Concrete.(ActorDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActorDefinition"}
		}
		return impl.SearchCapabilitiesActorDefinition(), nil
	case "AdministrableProductDefinition":
		impl, ok := w.Concrete.(AdministrableProductDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdministrableProductDefinition"}
		}
		return impl.SearchCapabilitiesAdministrableProductDefinition(), nil
	case "AdverseEvent":
		impl, ok := w.Concrete.(AdverseEventSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdverseEvent"}
		}
		return impl.SearchCapabilitiesAdverseEvent(), nil
	case "AllergyIntolerance":
		impl, ok := w.Concrete.(AllergyIntoleranceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AllergyIntolerance"}
		}
		return impl.SearchCapabilitiesAllergyIntolerance(), nil
	case "Appointment":
		impl, ok := w.Concrete.(AppointmentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Appointment"}
		}
		return impl.SearchCapabilitiesAppointment(), nil
	case "AppointmentResponse":
		impl, ok := w.Concrete.(AppointmentResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AppointmentResponse"}
		}
		return impl.SearchCapabilitiesAppointmentResponse(), nil
	case "ArtifactAssessment":
		impl, ok := w.Concrete.(ArtifactAssessmentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ArtifactAssessment"}
		}
		return impl.SearchCapabilitiesArtifactAssessment(), nil
	case "AuditEvent":
		impl, ok := w.Concrete.(AuditEventSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AuditEvent"}
		}
		return impl.SearchCapabilitiesAuditEvent(), nil
	case "Basic":
		impl, ok := w.Concrete.(BasicSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Basic"}
		}
		return impl.SearchCapabilitiesBasic(), nil
	case "Binary":
		impl, ok := w.Concrete.(BinarySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Binary"}
		}
		return impl.SearchCapabilitiesBinary(), nil
	case "BiologicallyDerivedProduct":
		impl, ok := w.Concrete.(BiologicallyDerivedProductSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.SearchCapabilitiesBiologicallyDerivedProduct(), nil
	case "BiologicallyDerivedProductDispense":
		impl, ok := w.Concrete.(BiologicallyDerivedProductDispenseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProductDispense"}
		}
		return impl.SearchCapabilitiesBiologicallyDerivedProductDispense(), nil
	case "BodyStructure":
		impl, ok := w.Concrete.(BodyStructureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BodyStructure"}
		}
		return impl.SearchCapabilitiesBodyStructure(), nil
	case "Bundle":
		impl, ok := w.Concrete.(BundleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Bundle"}
		}
		return impl.SearchCapabilitiesBundle(), nil
	case "CapabilityStatement":
		impl, ok := w.Concrete.(CapabilityStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CapabilityStatement"}
		}
		return impl.SearchCapabilitiesCapabilityStatement(), nil
	case "CarePlan":
		impl, ok := w.Concrete.(CarePlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CarePlan"}
		}
		return impl.SearchCapabilitiesCarePlan(), nil
	case "CareTeam":
		impl, ok := w.Concrete.(CareTeamSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CareTeam"}
		}
		return impl.SearchCapabilitiesCareTeam(), nil
	case "ChargeItem":
		impl, ok := w.Concrete.(ChargeItemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItem"}
		}
		return impl.SearchCapabilitiesChargeItem(), nil
	case "ChargeItemDefinition":
		impl, ok := w.Concrete.(ChargeItemDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItemDefinition"}
		}
		return impl.SearchCapabilitiesChargeItemDefinition(), nil
	case "Citation":
		impl, ok := w.Concrete.(CitationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Citation"}
		}
		return impl.SearchCapabilitiesCitation(), nil
	case "Claim":
		impl, ok := w.Concrete.(ClaimSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Claim"}
		}
		return impl.SearchCapabilitiesClaim(), nil
	case "ClaimResponse":
		impl, ok := w.Concrete.(ClaimResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClaimResponse"}
		}
		return impl.SearchCapabilitiesClaimResponse(), nil
	case "ClinicalImpression":
		impl, ok := w.Concrete.(ClinicalImpressionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalImpression"}
		}
		return impl.SearchCapabilitiesClinicalImpression(), nil
	case "ClinicalUseDefinition":
		impl, ok := w.Concrete.(ClinicalUseDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalUseDefinition"}
		}
		return impl.SearchCapabilitiesClinicalUseDefinition(), nil
	case "CodeSystem":
		impl, ok := w.Concrete.(CodeSystemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CodeSystem"}
		}
		return impl.SearchCapabilitiesCodeSystem(), nil
	case "Communication":
		impl, ok := w.Concrete.(CommunicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Communication"}
		}
		return impl.SearchCapabilitiesCommunication(), nil
	case "CommunicationRequest":
		impl, ok := w.Concrete.(CommunicationRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CommunicationRequest"}
		}
		return impl.SearchCapabilitiesCommunicationRequest(), nil
	case "CompartmentDefinition":
		impl, ok := w.Concrete.(CompartmentDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CompartmentDefinition"}
		}
		return impl.SearchCapabilitiesCompartmentDefinition(), nil
	case "Composition":
		impl, ok := w.Concrete.(CompositionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Composition"}
		}
		return impl.SearchCapabilitiesComposition(), nil
	case "ConceptMap":
		impl, ok := w.Concrete.(ConceptMapSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConceptMap"}
		}
		return impl.SearchCapabilitiesConceptMap(), nil
	case "Condition":
		impl, ok := w.Concrete.(ConditionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Condition"}
		}
		return impl.SearchCapabilitiesCondition(), nil
	case "ConditionDefinition":
		impl, ok := w.Concrete.(ConditionDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConditionDefinition"}
		}
		return impl.SearchCapabilitiesConditionDefinition(), nil
	case "Consent":
		impl, ok := w.Concrete.(ConsentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Consent"}
		}
		return impl.SearchCapabilitiesConsent(), nil
	case "Contract":
		impl, ok := w.Concrete.(ContractSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Contract"}
		}
		return impl.SearchCapabilitiesContract(), nil
	case "Coverage":
		impl, ok := w.Concrete.(CoverageSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Coverage"}
		}
		return impl.SearchCapabilitiesCoverage(), nil
	case "CoverageEligibilityRequest":
		impl, ok := w.Concrete.(CoverageEligibilityRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.SearchCapabilitiesCoverageEligibilityRequest(), nil
	case "CoverageEligibilityResponse":
		impl, ok := w.Concrete.(CoverageEligibilityResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.SearchCapabilitiesCoverageEligibilityResponse(), nil
	case "DetectedIssue":
		impl, ok := w.Concrete.(DetectedIssueSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DetectedIssue"}
		}
		return impl.SearchCapabilitiesDetectedIssue(), nil
	case "Device":
		impl, ok := w.Concrete.(DeviceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Device"}
		}
		return impl.SearchCapabilitiesDevice(), nil
	case "DeviceAssociation":
		impl, ok := w.Concrete.(DeviceAssociationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceAssociation"}
		}
		return impl.SearchCapabilitiesDeviceAssociation(), nil
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchCapabilitiesDeviceDefinition(), nil
	case "DeviceDispense":
		impl, ok := w.Concrete.(DeviceDispenseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDispense"}
		}
		return impl.SearchCapabilitiesDeviceDispense(), nil
	case "DeviceMetric":
		impl, ok := w.Concrete.(DeviceMetricSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceMetric"}
		}
		return impl.SearchCapabilitiesDeviceMetric(), nil
	case "DeviceRequest":
		impl, ok := w.Concrete.(DeviceRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceRequest"}
		}
		return impl.SearchCapabilitiesDeviceRequest(), nil
	case "DeviceUsage":
		impl, ok := w.Concrete.(DeviceUsageSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUsage"}
		}
		return impl.SearchCapabilitiesDeviceUsage(), nil
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchCapabilitiesDiagnosticReport(), nil
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchCapabilitiesDocumentReference(), nil
	case "Encounter":
		impl, ok := w.Concrete.(EncounterSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchCapabilitiesEncounter(), nil
	case "EncounterHistory":
		impl, ok := w.Concrete.(EncounterHistorySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EncounterHistory"}
		}
		return impl.SearchCapabilitiesEncounterHistory(), nil
	case "Endpoint":
		impl, ok := w.Concrete.(EndpointSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Endpoint"}
		}
		return impl.SearchCapabilitiesEndpoint(), nil
	case "EnrollmentRequest":
		impl, ok := w.Concrete.(EnrollmentRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentRequest"}
		}
		return impl.SearchCapabilitiesEnrollmentRequest(), nil
	case "EnrollmentResponse":
		impl, ok := w.Concrete.(EnrollmentResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentResponse"}
		}
		return impl.SearchCapabilitiesEnrollmentResponse(), nil
	case "EpisodeOfCare":
		impl, ok := w.Concrete.(EpisodeOfCareSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EpisodeOfCare"}
		}
		return impl.SearchCapabilitiesEpisodeOfCare(), nil
	case "EventDefinition":
		impl, ok := w.Concrete.(EventDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EventDefinition"}
		}
		return impl.SearchCapabilitiesEventDefinition(), nil
	case "Evidence":
		impl, ok := w.Concrete.(EvidenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Evidence"}
		}
		return impl.SearchCapabilitiesEvidence(), nil
	case "EvidenceReport":
		impl, ok := w.Concrete.(EvidenceReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceReport"}
		}
		return impl.SearchCapabilitiesEvidenceReport(), nil
	case "EvidenceVariable":
		impl, ok := w.Concrete.(EvidenceVariableSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceVariable"}
		}
		return impl.SearchCapabilitiesEvidenceVariable(), nil
	case "ExampleScenario":
		impl, ok := w.Concrete.(ExampleScenarioSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExampleScenario"}
		}
		return impl.SearchCapabilitiesExampleScenario(), nil
	case "ExplanationOfBenefit":
		impl, ok := w.Concrete.(ExplanationOfBenefitSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.SearchCapabilitiesExplanationOfBenefit(), nil
	case "FamilyMemberHistory":
		impl, ok := w.Concrete.(FamilyMemberHistorySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FamilyMemberHistory"}
		}
		return impl.SearchCapabilitiesFamilyMemberHistory(), nil
	case "Flag":
		impl, ok := w.Concrete.(FlagSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Flag"}
		}
		return impl.SearchCapabilitiesFlag(), nil
	case "FormularyItem":
		impl, ok := w.Concrete.(FormularyItemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FormularyItem"}
		}
		return impl.SearchCapabilitiesFormularyItem(), nil
	case "GenomicStudy":
		impl, ok := w.Concrete.(GenomicStudySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GenomicStudy"}
		}
		return impl.SearchCapabilitiesGenomicStudy(), nil
	case "Goal":
		impl, ok := w.Concrete.(GoalSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Goal"}
		}
		return impl.SearchCapabilitiesGoal(), nil
	case "GraphDefinition":
		impl, ok := w.Concrete.(GraphDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GraphDefinition"}
		}
		return impl.SearchCapabilitiesGraphDefinition(), nil
	case "Group":
		impl, ok := w.Concrete.(GroupSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Group"}
		}
		return impl.SearchCapabilitiesGroup(), nil
	case "GuidanceResponse":
		impl, ok := w.Concrete.(GuidanceResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GuidanceResponse"}
		}
		return impl.SearchCapabilitiesGuidanceResponse(), nil
	case "HealthcareService":
		impl, ok := w.Concrete.(HealthcareServiceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "HealthcareService"}
		}
		return impl.SearchCapabilitiesHealthcareService(), nil
	case "ImagingSelection":
		impl, ok := w.Concrete.(ImagingSelectionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingSelection"}
		}
		return impl.SearchCapabilitiesImagingSelection(), nil
	case "ImagingStudy":
		impl, ok := w.Concrete.(ImagingStudySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingStudy"}
		}
		return impl.SearchCapabilitiesImagingStudy(), nil
	case "Immunization":
		impl, ok := w.Concrete.(ImmunizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Immunization"}
		}
		return impl.SearchCapabilitiesImmunization(), nil
	case "ImmunizationEvaluation":
		impl, ok := w.Concrete.(ImmunizationEvaluationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.SearchCapabilitiesImmunizationEvaluation(), nil
	case "ImmunizationRecommendation":
		impl, ok := w.Concrete.(ImmunizationRecommendationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.SearchCapabilitiesImmunizationRecommendation(), nil
	case "ImplementationGuide":
		impl, ok := w.Concrete.(ImplementationGuideSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImplementationGuide"}
		}
		return impl.SearchCapabilitiesImplementationGuide(), nil
	case "Ingredient":
		impl, ok := w.Concrete.(IngredientSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Ingredient"}
		}
		return impl.SearchCapabilitiesIngredient(), nil
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchCapabilitiesInsurancePlan(), nil
	case "InventoryItem":
		impl, ok := w.Concrete.(InventoryItemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InventoryItem"}
		}
		return impl.SearchCapabilitiesInventoryItem(), nil
	case "InventoryReport":
		impl, ok := w.Concrete.(InventoryReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InventoryReport"}
		}
		return impl.SearchCapabilitiesInventoryReport(), nil
	case "Invoice":
		impl, ok := w.Concrete.(InvoiceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Invoice"}
		}
		return impl.SearchCapabilitiesInvoice(), nil
	case "Library":
		impl, ok := w.Concrete.(LibrarySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Library"}
		}
		return impl.SearchCapabilitiesLibrary(), nil
	case "Linkage":
		impl, ok := w.Concrete.(LinkageSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Linkage"}
		}
		return impl.SearchCapabilitiesLinkage(), nil
	case "List":
		impl, ok := w.Concrete.(ListSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "List"}
		}
		return impl.SearchCapabilitiesList(), nil
	case "Location":
		impl, ok := w.Concrete.(LocationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Location"}
		}
		return impl.SearchCapabilitiesLocation(), nil
	case "ManufacturedItemDefinition":
		impl, ok := w.Concrete.(ManufacturedItemDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ManufacturedItemDefinition"}
		}
		return impl.SearchCapabilitiesManufacturedItemDefinition(), nil
	case "Measure":
		impl, ok := w.Concrete.(MeasureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Measure"}
		}
		return impl.SearchCapabilitiesMeasure(), nil
	case "MeasureReport":
		impl, ok := w.Concrete.(MeasureReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MeasureReport"}
		}
		return impl.SearchCapabilitiesMeasureReport(), nil
	case "Medication":
		impl, ok := w.Concrete.(MedicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Medication"}
		}
		return impl.SearchCapabilitiesMedication(), nil
	case "MedicationAdministration":
		impl, ok := w.Concrete.(MedicationAdministrationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationAdministration"}
		}
		return impl.SearchCapabilitiesMedicationAdministration(), nil
	case "MedicationDispense":
		impl, ok := w.Concrete.(MedicationDispenseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationDispense"}
		}
		return impl.SearchCapabilitiesMedicationDispense(), nil
	case "MedicationKnowledge":
		impl, ok := w.Concrete.(MedicationKnowledgeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationKnowledge"}
		}
		return impl.SearchCapabilitiesMedicationKnowledge(), nil
	case "MedicationRequest":
		impl, ok := w.Concrete.(MedicationRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationRequest"}
		}
		return impl.SearchCapabilitiesMedicationRequest(), nil
	case "MedicationStatement":
		impl, ok := w.Concrete.(MedicationStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationStatement"}
		}
		return impl.SearchCapabilitiesMedicationStatement(), nil
	case "MedicinalProductDefinition":
		impl, ok := w.Concrete.(MedicinalProductDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductDefinition"}
		}
		return impl.SearchCapabilitiesMedicinalProductDefinition(), nil
	case "MessageDefinition":
		impl, ok := w.Concrete.(MessageDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageDefinition"}
		}
		return impl.SearchCapabilitiesMessageDefinition(), nil
	case "MessageHeader":
		impl, ok := w.Concrete.(MessageHeaderSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageHeader"}
		}
		return impl.SearchCapabilitiesMessageHeader(), nil
	case "MolecularSequence":
		impl, ok := w.Concrete.(MolecularSequenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MolecularSequence"}
		}
		return impl.SearchCapabilitiesMolecularSequence(), nil
	case "NamingSystem":
		impl, ok := w.Concrete.(NamingSystemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NamingSystem"}
		}
		return impl.SearchCapabilitiesNamingSystem(), nil
	case "NutritionIntake":
		impl, ok := w.Concrete.(NutritionIntakeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionIntake"}
		}
		return impl.SearchCapabilitiesNutritionIntake(), nil
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchCapabilitiesNutritionOrder(), nil
	case "NutritionProduct":
		impl, ok := w.Concrete.(NutritionProductSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionProduct"}
		}
		return impl.SearchCapabilitiesNutritionProduct(), nil
	case "Observation":
		impl, ok := w.Concrete.(ObservationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Observation"}
		}
		return impl.SearchCapabilitiesObservation(), nil
	case "ObservationDefinition":
		impl, ok := w.Concrete.(ObservationDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ObservationDefinition"}
		}
		return impl.SearchCapabilitiesObservationDefinition(), nil
	case "OperationDefinition":
		impl, ok := w.Concrete.(OperationDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationDefinition"}
		}
		return impl.SearchCapabilitiesOperationDefinition(), nil
	case "OperationOutcome":
		impl, ok := w.Concrete.(OperationOutcomeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationOutcome"}
		}
		return impl.SearchCapabilitiesOperationOutcome(), nil
	case "Organization":
		impl, ok := w.Concrete.(OrganizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Organization"}
		}
		return impl.SearchCapabilitiesOrganization(), nil
	case "OrganizationAffiliation":
		impl, ok := w.Concrete.(OrganizationAffiliationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OrganizationAffiliation"}
		}
		return impl.SearchCapabilitiesOrganizationAffiliation(), nil
	case "PackagedProductDefinition":
		impl, ok := w.Concrete.(PackagedProductDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PackagedProductDefinition"}
		}
		return impl.SearchCapabilitiesPackagedProductDefinition(), nil
	case "Parameters":
		impl, ok := w.Concrete.(ParametersSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Parameters"}
		}
		return impl.SearchCapabilitiesParameters(), nil
	case "Patient":
		impl, ok := w.Concrete.(PatientSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Patient"}
		}
		return impl.SearchCapabilitiesPatient(), nil
	case "PaymentNotice":
		impl, ok := w.Concrete.(PaymentNoticeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentNotice"}
		}
		return impl.SearchCapabilitiesPaymentNotice(), nil
	case "PaymentReconciliation":
		impl, ok := w.Concrete.(PaymentReconciliationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentReconciliation"}
		}
		return impl.SearchCapabilitiesPaymentReconciliation(), nil
	case "Permission":
		impl, ok := w.Concrete.(PermissionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Permission"}
		}
		return impl.SearchCapabilitiesPermission(), nil
	case "Person":
		impl, ok := w.Concrete.(PersonSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Person"}
		}
		return impl.SearchCapabilitiesPerson(), nil
	case "PlanDefinition":
		impl, ok := w.Concrete.(PlanDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PlanDefinition"}
		}
		return impl.SearchCapabilitiesPlanDefinition(), nil
	case "Practitioner":
		impl, ok := w.Concrete.(PractitionerSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Practitioner"}
		}
		return impl.SearchCapabilitiesPractitioner(), nil
	case "PractitionerRole":
		impl, ok := w.Concrete.(PractitionerRoleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PractitionerRole"}
		}
		return impl.SearchCapabilitiesPractitionerRole(), nil
	case "Procedure":
		impl, ok := w.Concrete.(ProcedureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Procedure"}
		}
		return impl.SearchCapabilitiesProcedure(), nil
	case "Provenance":
		impl, ok := w.Concrete.(ProvenanceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Provenance"}
		}
		return impl.SearchCapabilitiesProvenance(), nil
	case "Questionnaire":
		impl, ok := w.Concrete.(QuestionnaireSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Questionnaire"}
		}
		return impl.SearchCapabilitiesQuestionnaire(), nil
	case "QuestionnaireResponse":
		impl, ok := w.Concrete.(QuestionnaireResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "QuestionnaireResponse"}
		}
		return impl.SearchCapabilitiesQuestionnaireResponse(), nil
	case "RegulatedAuthorization":
		impl, ok := w.Concrete.(RegulatedAuthorizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RegulatedAuthorization"}
		}
		return impl.SearchCapabilitiesRegulatedAuthorization(), nil
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchCapabilitiesRelatedPerson(), nil
	case "RequestOrchestration":
		impl, ok := w.Concrete.(RequestOrchestrationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestOrchestration"}
		}
		return impl.SearchCapabilitiesRequestOrchestration(), nil
	case "Requirements":
		impl, ok := w.Concrete.(RequirementsSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Requirements"}
		}
		return impl.SearchCapabilitiesRequirements(), nil
	case "ResearchStudy":
		impl, ok := w.Concrete.(ResearchStudySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchStudy"}
		}
		return impl.SearchCapabilitiesResearchStudy(), nil
	case "ResearchSubject":
		impl, ok := w.Concrete.(ResearchSubjectSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchSubject"}
		}
		return impl.SearchCapabilitiesResearchSubject(), nil
	case "RiskAssessment":
		impl, ok := w.Concrete.(RiskAssessmentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskAssessment"}
		}
		return impl.SearchCapabilitiesRiskAssessment(), nil
	case "Schedule":
		impl, ok := w.Concrete.(ScheduleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Schedule"}
		}
		return impl.SearchCapabilitiesSchedule(), nil
	case "SearchParameter":
		impl, ok := w.Concrete.(SearchParameterSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SearchParameter"}
		}
		return impl.SearchCapabilitiesSearchParameter(), nil
	case "ServiceRequest":
		impl, ok := w.Concrete.(ServiceRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ServiceRequest"}
		}
		return impl.SearchCapabilitiesServiceRequest(), nil
	case "Slot":
		impl, ok := w.Concrete.(SlotSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Slot"}
		}
		return impl.SearchCapabilitiesSlot(), nil
	case "Specimen":
		impl, ok := w.Concrete.(SpecimenSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Specimen"}
		}
		return impl.SearchCapabilitiesSpecimen(), nil
	case "SpecimenDefinition":
		impl, ok := w.Concrete.(SpecimenDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SpecimenDefinition"}
		}
		return impl.SearchCapabilitiesSpecimenDefinition(), nil
	case "StructureDefinition":
		impl, ok := w.Concrete.(StructureDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureDefinition"}
		}
		return impl.SearchCapabilitiesStructureDefinition(), nil
	case "StructureMap":
		impl, ok := w.Concrete.(StructureMapSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureMap"}
		}
		return impl.SearchCapabilitiesStructureMap(), nil
	case "Subscription":
		impl, ok := w.Concrete.(SubscriptionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Subscription"}
		}
		return impl.SearchCapabilitiesSubscription(), nil
	case "SubscriptionStatus":
		impl, ok := w.Concrete.(SubscriptionStatusSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubscriptionStatus"}
		}
		return impl.SearchCapabilitiesSubscriptionStatus(), nil
	case "SubscriptionTopic":
		impl, ok := w.Concrete.(SubscriptionTopicSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubscriptionTopic"}
		}
		return impl.SearchCapabilitiesSubscriptionTopic(), nil
	case "Substance":
		impl, ok := w.Concrete.(SubstanceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchCapabilitiesSubstance(), nil
	case "SubstanceDefinition":
		impl, ok := w.Concrete.(SubstanceDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceDefinition"}
		}
		return impl.SearchCapabilitiesSubstanceDefinition(), nil
	case "SubstanceNucleicAcid":
		impl, ok := w.Concrete.(SubstanceNucleicAcidSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.SearchCapabilitiesSubstanceNucleicAcid(), nil
	case "SubstancePolymer":
		impl, ok := w.Concrete.(SubstancePolymerSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstancePolymer"}
		}
		return impl.SearchCapabilitiesSubstancePolymer(), nil
	case "SubstanceProtein":
		impl, ok := w.Concrete.(SubstanceProteinSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceProtein"}
		}
		return impl.SearchCapabilitiesSubstanceProtein(), nil
	case "SubstanceReferenceInformation":
		impl, ok := w.Concrete.(SubstanceReferenceInformationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.SearchCapabilitiesSubstanceReferenceInformation(), nil
	case "SubstanceSourceMaterial":
		impl, ok := w.Concrete.(SubstanceSourceMaterialSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.SearchCapabilitiesSubstanceSourceMaterial(), nil
	case "SupplyDelivery":
		impl, ok := w.Concrete.(SupplyDeliverySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyDelivery"}
		}
		return impl.SearchCapabilitiesSupplyDelivery(), nil
	case "SupplyRequest":
		impl, ok := w.Concrete.(SupplyRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyRequest"}
		}
		return impl.SearchCapabilitiesSupplyRequest(), nil
	case "Task":
		impl, ok := w.Concrete.(TaskSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Task"}
		}
		return impl.SearchCapabilitiesTask(), nil
	case "TerminologyCapabilities":
		impl, ok := w.Concrete.(TerminologyCapabilitiesSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TerminologyCapabilities"}
		}
		return impl.SearchCapabilitiesTerminologyCapabilities(), nil
	case "TestPlan":
		impl, ok := w.Concrete.(TestPlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestPlan"}
		}
		return impl.SearchCapabilitiesTestPlan(), nil
	case "TestReport":
		impl, ok := w.Concrete.(TestReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestReport"}
		}
		return impl.SearchCapabilitiesTestReport(), nil
	case "TestScript":
		impl, ok := w.Concrete.(TestScriptSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestScript"}
		}
		return impl.SearchCapabilitiesTestScript(), nil
	case "Transport":
		impl, ok := w.Concrete.(TransportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Transport"}
		}
		return impl.SearchCapabilitiesTransport(), nil
	case "ValueSet":
		impl, ok := w.Concrete.(ValueSetSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ValueSet"}
		}
		return impl.SearchCapabilitiesValueSet(), nil
	case "VerificationResult":
		impl, ok := w.Concrete.(VerificationResultSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VerificationResult"}
		}
		return impl.SearchCapabilitiesVerificationResult(), nil
	case "VisionPrescription":
		impl, ok := w.Concrete.(VisionPrescriptionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VisionPrescription"}
		}
		return impl.SearchCapabilitiesVisionPrescription(), nil
	default:
		return search.Capabilities{}, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
func (w Generic) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.Concrete.(AccountSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Account"}
		}
		return impl.SearchAccount(ctx, options)
	case "ActivityDefinition":
		impl, ok := w.Concrete.(ActivityDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActivityDefinition"}
		}
		return impl.SearchActivityDefinition(ctx, options)
	case "ActorDefinition":
		impl, ok := w.Concrete.(ActorDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActorDefinition"}
		}
		return impl.SearchActorDefinition(ctx, options)
	case "AdministrableProductDefinition":
		impl, ok := w.Concrete.(AdministrableProductDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdministrableProductDefinition"}
		}
		return impl.SearchAdministrableProductDefinition(ctx, options)
	case "AdverseEvent":
		impl, ok := w.Concrete.(AdverseEventSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdverseEvent"}
		}
		return impl.SearchAdverseEvent(ctx, options)
	case "AllergyIntolerance":
		impl, ok := w.Concrete.(AllergyIntoleranceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AllergyIntolerance"}
		}
		return impl.SearchAllergyIntolerance(ctx, options)
	case "Appointment":
		impl, ok := w.Concrete.(AppointmentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Appointment"}
		}
		return impl.SearchAppointment(ctx, options)
	case "AppointmentResponse":
		impl, ok := w.Concrete.(AppointmentResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AppointmentResponse"}
		}
		return impl.SearchAppointmentResponse(ctx, options)
	case "ArtifactAssessment":
		impl, ok := w.Concrete.(ArtifactAssessmentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ArtifactAssessment"}
		}
		return impl.SearchArtifactAssessment(ctx, options)
	case "AuditEvent":
		impl, ok := w.Concrete.(AuditEventSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AuditEvent"}
		}
		return impl.SearchAuditEvent(ctx, options)
	case "Basic":
		impl, ok := w.Concrete.(BasicSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Basic"}
		}
		return impl.SearchBasic(ctx, options)
	case "Binary":
		impl, ok := w.Concrete.(BinarySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Binary"}
		}
		return impl.SearchBinary(ctx, options)
	case "BiologicallyDerivedProduct":
		impl, ok := w.Concrete.(BiologicallyDerivedProductSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.SearchBiologicallyDerivedProduct(ctx, options)
	case "BiologicallyDerivedProductDispense":
		impl, ok := w.Concrete.(BiologicallyDerivedProductDispenseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProductDispense"}
		}
		return impl.SearchBiologicallyDerivedProductDispense(ctx, options)
	case "BodyStructure":
		impl, ok := w.Concrete.(BodyStructureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BodyStructure"}
		}
		return impl.SearchBodyStructure(ctx, options)
	case "Bundle":
		impl, ok := w.Concrete.(BundleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Bundle"}
		}
		return impl.SearchBundle(ctx, options)
	case "CapabilityStatement":
		impl, ok := w.Concrete.(CapabilityStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CapabilityStatement"}
		}
		return impl.SearchCapabilityStatement(ctx, options)
	case "CarePlan":
		impl, ok := w.Concrete.(CarePlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CarePlan"}
		}
		return impl.SearchCarePlan(ctx, options)
	case "CareTeam":
		impl, ok := w.Concrete.(CareTeamSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CareTeam"}
		}
		return impl.SearchCareTeam(ctx, options)
	case "ChargeItem":
		impl, ok := w.Concrete.(ChargeItemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItem"}
		}
		return impl.SearchChargeItem(ctx, options)
	case "ChargeItemDefinition":
		impl, ok := w.Concrete.(ChargeItemDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItemDefinition"}
		}
		return impl.SearchChargeItemDefinition(ctx, options)
	case "Citation":
		impl, ok := w.Concrete.(CitationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Citation"}
		}
		return impl.SearchCitation(ctx, options)
	case "Claim":
		impl, ok := w.Concrete.(ClaimSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Claim"}
		}
		return impl.SearchClaim(ctx, options)
	case "ClaimResponse":
		impl, ok := w.Concrete.(ClaimResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClaimResponse"}
		}
		return impl.SearchClaimResponse(ctx, options)
	case "ClinicalImpression":
		impl, ok := w.Concrete.(ClinicalImpressionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalImpression"}
		}
		return impl.SearchClinicalImpression(ctx, options)
	case "ClinicalUseDefinition":
		impl, ok := w.Concrete.(ClinicalUseDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalUseDefinition"}
		}
		return impl.SearchClinicalUseDefinition(ctx, options)
	case "CodeSystem":
		impl, ok := w.Concrete.(CodeSystemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CodeSystem"}
		}
		return impl.SearchCodeSystem(ctx, options)
	case "Communication":
		impl, ok := w.Concrete.(CommunicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Communication"}
		}
		return impl.SearchCommunication(ctx, options)
	case "CommunicationRequest":
		impl, ok := w.Concrete.(CommunicationRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CommunicationRequest"}
		}
		return impl.SearchCommunicationRequest(ctx, options)
	case "CompartmentDefinition":
		impl, ok := w.Concrete.(CompartmentDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CompartmentDefinition"}
		}
		return impl.SearchCompartmentDefinition(ctx, options)
	case "Composition":
		impl, ok := w.Concrete.(CompositionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Composition"}
		}
		return impl.SearchComposition(ctx, options)
	case "ConceptMap":
		impl, ok := w.Concrete.(ConceptMapSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConceptMap"}
		}
		return impl.SearchConceptMap(ctx, options)
	case "Condition":
		impl, ok := w.Concrete.(ConditionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Condition"}
		}
		return impl.SearchCondition(ctx, options)
	case "ConditionDefinition":
		impl, ok := w.Concrete.(ConditionDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConditionDefinition"}
		}
		return impl.SearchConditionDefinition(ctx, options)
	case "Consent":
		impl, ok := w.Concrete.(ConsentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Consent"}
		}
		return impl.SearchConsent(ctx, options)
	case "Contract":
		impl, ok := w.Concrete.(ContractSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Contract"}
		}
		return impl.SearchContract(ctx, options)
	case "Coverage":
		impl, ok := w.Concrete.(CoverageSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Coverage"}
		}
		return impl.SearchCoverage(ctx, options)
	case "CoverageEligibilityRequest":
		impl, ok := w.Concrete.(CoverageEligibilityRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.SearchCoverageEligibilityRequest(ctx, options)
	case "CoverageEligibilityResponse":
		impl, ok := w.Concrete.(CoverageEligibilityResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.SearchCoverageEligibilityResponse(ctx, options)
	case "DetectedIssue":
		impl, ok := w.Concrete.(DetectedIssueSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DetectedIssue"}
		}
		return impl.SearchDetectedIssue(ctx, options)
	case "Device":
		impl, ok := w.Concrete.(DeviceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Device"}
		}
		return impl.SearchDevice(ctx, options)
	case "DeviceAssociation":
		impl, ok := w.Concrete.(DeviceAssociationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceAssociation"}
		}
		return impl.SearchDeviceAssociation(ctx, options)
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchDeviceDefinition(ctx, options)
	case "DeviceDispense":
		impl, ok := w.Concrete.(DeviceDispenseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDispense"}
		}
		return impl.SearchDeviceDispense(ctx, options)
	case "DeviceMetric":
		impl, ok := w.Concrete.(DeviceMetricSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceMetric"}
		}
		return impl.SearchDeviceMetric(ctx, options)
	case "DeviceRequest":
		impl, ok := w.Concrete.(DeviceRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceRequest"}
		}
		return impl.SearchDeviceRequest(ctx, options)
	case "DeviceUsage":
		impl, ok := w.Concrete.(DeviceUsageSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUsage"}
		}
		return impl.SearchDeviceUsage(ctx, options)
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchDiagnosticReport(ctx, options)
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchDocumentReference(ctx, options)
	case "Encounter":
		impl, ok := w.Concrete.(EncounterSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchEncounter(ctx, options)
	case "EncounterHistory":
		impl, ok := w.Concrete.(EncounterHistorySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EncounterHistory"}
		}
		return impl.SearchEncounterHistory(ctx, options)
	case "Endpoint":
		impl, ok := w.Concrete.(EndpointSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Endpoint"}
		}
		return impl.SearchEndpoint(ctx, options)
	case "EnrollmentRequest":
		impl, ok := w.Concrete.(EnrollmentRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentRequest"}
		}
		return impl.SearchEnrollmentRequest(ctx, options)
	case "EnrollmentResponse":
		impl, ok := w.Concrete.(EnrollmentResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentResponse"}
		}
		return impl.SearchEnrollmentResponse(ctx, options)
	case "EpisodeOfCare":
		impl, ok := w.Concrete.(EpisodeOfCareSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EpisodeOfCare"}
		}
		return impl.SearchEpisodeOfCare(ctx, options)
	case "EventDefinition":
		impl, ok := w.Concrete.(EventDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EventDefinition"}
		}
		return impl.SearchEventDefinition(ctx, options)
	case "Evidence":
		impl, ok := w.Concrete.(EvidenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Evidence"}
		}
		return impl.SearchEvidence(ctx, options)
	case "EvidenceReport":
		impl, ok := w.Concrete.(EvidenceReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceReport"}
		}
		return impl.SearchEvidenceReport(ctx, options)
	case "EvidenceVariable":
		impl, ok := w.Concrete.(EvidenceVariableSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceVariable"}
		}
		return impl.SearchEvidenceVariable(ctx, options)
	case "ExampleScenario":
		impl, ok := w.Concrete.(ExampleScenarioSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExampleScenario"}
		}
		return impl.SearchExampleScenario(ctx, options)
	case "ExplanationOfBenefit":
		impl, ok := w.Concrete.(ExplanationOfBenefitSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.SearchExplanationOfBenefit(ctx, options)
	case "FamilyMemberHistory":
		impl, ok := w.Concrete.(FamilyMemberHistorySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FamilyMemberHistory"}
		}
		return impl.SearchFamilyMemberHistory(ctx, options)
	case "Flag":
		impl, ok := w.Concrete.(FlagSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Flag"}
		}
		return impl.SearchFlag(ctx, options)
	case "FormularyItem":
		impl, ok := w.Concrete.(FormularyItemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FormularyItem"}
		}
		return impl.SearchFormularyItem(ctx, options)
	case "GenomicStudy":
		impl, ok := w.Concrete.(GenomicStudySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GenomicStudy"}
		}
		return impl.SearchGenomicStudy(ctx, options)
	case "Goal":
		impl, ok := w.Concrete.(GoalSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Goal"}
		}
		return impl.SearchGoal(ctx, options)
	case "GraphDefinition":
		impl, ok := w.Concrete.(GraphDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GraphDefinition"}
		}
		return impl.SearchGraphDefinition(ctx, options)
	case "Group":
		impl, ok := w.Concrete.(GroupSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Group"}
		}
		return impl.SearchGroup(ctx, options)
	case "GuidanceResponse":
		impl, ok := w.Concrete.(GuidanceResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GuidanceResponse"}
		}
		return impl.SearchGuidanceResponse(ctx, options)
	case "HealthcareService":
		impl, ok := w.Concrete.(HealthcareServiceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "HealthcareService"}
		}
		return impl.SearchHealthcareService(ctx, options)
	case "ImagingSelection":
		impl, ok := w.Concrete.(ImagingSelectionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingSelection"}
		}
		return impl.SearchImagingSelection(ctx, options)
	case "ImagingStudy":
		impl, ok := w.Concrete.(ImagingStudySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingStudy"}
		}
		return impl.SearchImagingStudy(ctx, options)
	case "Immunization":
		impl, ok := w.Concrete.(ImmunizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Immunization"}
		}
		return impl.SearchImmunization(ctx, options)
	case "ImmunizationEvaluation":
		impl, ok := w.Concrete.(ImmunizationEvaluationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.SearchImmunizationEvaluation(ctx, options)
	case "ImmunizationRecommendation":
		impl, ok := w.Concrete.(ImmunizationRecommendationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.SearchImmunizationRecommendation(ctx, options)
	case "ImplementationGuide":
		impl, ok := w.Concrete.(ImplementationGuideSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImplementationGuide"}
		}
		return impl.SearchImplementationGuide(ctx, options)
	case "Ingredient":
		impl, ok := w.Concrete.(IngredientSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Ingredient"}
		}
		return impl.SearchIngredient(ctx, options)
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchInsurancePlan(ctx, options)
	case "InventoryItem":
		impl, ok := w.Concrete.(InventoryItemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InventoryItem"}
		}
		return impl.SearchInventoryItem(ctx, options)
	case "InventoryReport":
		impl, ok := w.Concrete.(InventoryReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InventoryReport"}
		}
		return impl.SearchInventoryReport(ctx, options)
	case "Invoice":
		impl, ok := w.Concrete.(InvoiceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Invoice"}
		}
		return impl.SearchInvoice(ctx, options)
	case "Library":
		impl, ok := w.Concrete.(LibrarySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Library"}
		}
		return impl.SearchLibrary(ctx, options)
	case "Linkage":
		impl, ok := w.Concrete.(LinkageSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Linkage"}
		}
		return impl.SearchLinkage(ctx, options)
	case "List":
		impl, ok := w.Concrete.(ListSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "List"}
		}
		return impl.SearchList(ctx, options)
	case "Location":
		impl, ok := w.Concrete.(LocationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Location"}
		}
		return impl.SearchLocation(ctx, options)
	case "ManufacturedItemDefinition":
		impl, ok := w.Concrete.(ManufacturedItemDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ManufacturedItemDefinition"}
		}
		return impl.SearchManufacturedItemDefinition(ctx, options)
	case "Measure":
		impl, ok := w.Concrete.(MeasureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Measure"}
		}
		return impl.SearchMeasure(ctx, options)
	case "MeasureReport":
		impl, ok := w.Concrete.(MeasureReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MeasureReport"}
		}
		return impl.SearchMeasureReport(ctx, options)
	case "Medication":
		impl, ok := w.Concrete.(MedicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Medication"}
		}
		return impl.SearchMedication(ctx, options)
	case "MedicationAdministration":
		impl, ok := w.Concrete.(MedicationAdministrationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationAdministration"}
		}
		return impl.SearchMedicationAdministration(ctx, options)
	case "MedicationDispense":
		impl, ok := w.Concrete.(MedicationDispenseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationDispense"}
		}
		return impl.SearchMedicationDispense(ctx, options)
	case "MedicationKnowledge":
		impl, ok := w.Concrete.(MedicationKnowledgeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationKnowledge"}
		}
		return impl.SearchMedicationKnowledge(ctx, options)
	case "MedicationRequest":
		impl, ok := w.Concrete.(MedicationRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationRequest"}
		}
		return impl.SearchMedicationRequest(ctx, options)
	case "MedicationStatement":
		impl, ok := w.Concrete.(MedicationStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationStatement"}
		}
		return impl.SearchMedicationStatement(ctx, options)
	case "MedicinalProductDefinition":
		impl, ok := w.Concrete.(MedicinalProductDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductDefinition"}
		}
		return impl.SearchMedicinalProductDefinition(ctx, options)
	case "MessageDefinition":
		impl, ok := w.Concrete.(MessageDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageDefinition"}
		}
		return impl.SearchMessageDefinition(ctx, options)
	case "MessageHeader":
		impl, ok := w.Concrete.(MessageHeaderSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageHeader"}
		}
		return impl.SearchMessageHeader(ctx, options)
	case "MolecularSequence":
		impl, ok := w.Concrete.(MolecularSequenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MolecularSequence"}
		}
		return impl.SearchMolecularSequence(ctx, options)
	case "NamingSystem":
		impl, ok := w.Concrete.(NamingSystemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NamingSystem"}
		}
		return impl.SearchNamingSystem(ctx, options)
	case "NutritionIntake":
		impl, ok := w.Concrete.(NutritionIntakeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionIntake"}
		}
		return impl.SearchNutritionIntake(ctx, options)
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchNutritionOrder(ctx, options)
	case "NutritionProduct":
		impl, ok := w.Concrete.(NutritionProductSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionProduct"}
		}
		return impl.SearchNutritionProduct(ctx, options)
	case "Observation":
		impl, ok := w.Concrete.(ObservationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Observation"}
		}
		return impl.SearchObservation(ctx, options)
	case "ObservationDefinition":
		impl, ok := w.Concrete.(ObservationDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ObservationDefinition"}
		}
		return impl.SearchObservationDefinition(ctx, options)
	case "OperationDefinition":
		impl, ok := w.Concrete.(OperationDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationDefinition"}
		}
		return impl.SearchOperationDefinition(ctx, options)
	case "OperationOutcome":
		impl, ok := w.Concrete.(OperationOutcomeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationOutcome"}
		}
		return impl.SearchOperationOutcome(ctx, options)
	case "Organization":
		impl, ok := w.Concrete.(OrganizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Organization"}
		}
		return impl.SearchOrganization(ctx, options)
	case "OrganizationAffiliation":
		impl, ok := w.Concrete.(OrganizationAffiliationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OrganizationAffiliation"}
		}
		return impl.SearchOrganizationAffiliation(ctx, options)
	case "PackagedProductDefinition":
		impl, ok := w.Concrete.(PackagedProductDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PackagedProductDefinition"}
		}
		return impl.SearchPackagedProductDefinition(ctx, options)
	case "Parameters":
		impl, ok := w.Concrete.(ParametersSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Parameters"}
		}
		return impl.SearchParameters(ctx, options)
	case "Patient":
		impl, ok := w.Concrete.(PatientSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Patient"}
		}
		return impl.SearchPatient(ctx, options)
	case "PaymentNotice":
		impl, ok := w.Concrete.(PaymentNoticeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentNotice"}
		}
		return impl.SearchPaymentNotice(ctx, options)
	case "PaymentReconciliation":
		impl, ok := w.Concrete.(PaymentReconciliationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentReconciliation"}
		}
		return impl.SearchPaymentReconciliation(ctx, options)
	case "Permission":
		impl, ok := w.Concrete.(PermissionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Permission"}
		}
		return impl.SearchPermission(ctx, options)
	case "Person":
		impl, ok := w.Concrete.(PersonSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Person"}
		}
		return impl.SearchPerson(ctx, options)
	case "PlanDefinition":
		impl, ok := w.Concrete.(PlanDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PlanDefinition"}
		}
		return impl.SearchPlanDefinition(ctx, options)
	case "Practitioner":
		impl, ok := w.Concrete.(PractitionerSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Practitioner"}
		}
		return impl.SearchPractitioner(ctx, options)
	case "PractitionerRole":
		impl, ok := w.Concrete.(PractitionerRoleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PractitionerRole"}
		}
		return impl.SearchPractitionerRole(ctx, options)
	case "Procedure":
		impl, ok := w.Concrete.(ProcedureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Procedure"}
		}
		return impl.SearchProcedure(ctx, options)
	case "Provenance":
		impl, ok := w.Concrete.(ProvenanceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Provenance"}
		}
		return impl.SearchProvenance(ctx, options)
	case "Questionnaire":
		impl, ok := w.Concrete.(QuestionnaireSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Questionnaire"}
		}
		return impl.SearchQuestionnaire(ctx, options)
	case "QuestionnaireResponse":
		impl, ok := w.Concrete.(QuestionnaireResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "QuestionnaireResponse"}
		}
		return impl.SearchQuestionnaireResponse(ctx, options)
	case "RegulatedAuthorization":
		impl, ok := w.Concrete.(RegulatedAuthorizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RegulatedAuthorization"}
		}
		return impl.SearchRegulatedAuthorization(ctx, options)
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchRelatedPerson(ctx, options)
	case "RequestOrchestration":
		impl, ok := w.Concrete.(RequestOrchestrationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestOrchestration"}
		}
		return impl.SearchRequestOrchestration(ctx, options)
	case "Requirements":
		impl, ok := w.Concrete.(RequirementsSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Requirements"}
		}
		return impl.SearchRequirements(ctx, options)
	case "ResearchStudy":
		impl, ok := w.Concrete.(ResearchStudySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchStudy"}
		}
		return impl.SearchResearchStudy(ctx, options)
	case "ResearchSubject":
		impl, ok := w.Concrete.(ResearchSubjectSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchSubject"}
		}
		return impl.SearchResearchSubject(ctx, options)
	case "RiskAssessment":
		impl, ok := w.Concrete.(RiskAssessmentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskAssessment"}
		}
		return impl.SearchRiskAssessment(ctx, options)
	case "Schedule":
		impl, ok := w.Concrete.(ScheduleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Schedule"}
		}
		return impl.SearchSchedule(ctx, options)
	case "SearchParameter":
		impl, ok := w.Concrete.(SearchParameterSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SearchParameter"}
		}
		return impl.SearchSearchParameter(ctx, options)
	case "ServiceRequest":
		impl, ok := w.Concrete.(ServiceRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ServiceRequest"}
		}
		return impl.SearchServiceRequest(ctx, options)
	case "Slot":
		impl, ok := w.Concrete.(SlotSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Slot"}
		}
		return impl.SearchSlot(ctx, options)
	case "Specimen":
		impl, ok := w.Concrete.(SpecimenSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Specimen"}
		}
		return impl.SearchSpecimen(ctx, options)
	case "SpecimenDefinition":
		impl, ok := w.Concrete.(SpecimenDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SpecimenDefinition"}
		}
		return impl.SearchSpecimenDefinition(ctx, options)
	case "StructureDefinition":
		impl, ok := w.Concrete.(StructureDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureDefinition"}
		}
		return impl.SearchStructureDefinition(ctx, options)
	case "StructureMap":
		impl, ok := w.Concrete.(StructureMapSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureMap"}
		}
		return impl.SearchStructureMap(ctx, options)
	case "Subscription":
		impl, ok := w.Concrete.(SubscriptionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Subscription"}
		}
		return impl.SearchSubscription(ctx, options)
	case "SubscriptionStatus":
		impl, ok := w.Concrete.(SubscriptionStatusSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubscriptionStatus"}
		}
		return impl.SearchSubscriptionStatus(ctx, options)
	case "SubscriptionTopic":
		impl, ok := w.Concrete.(SubscriptionTopicSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubscriptionTopic"}
		}
		return impl.SearchSubscriptionTopic(ctx, options)
	case "Substance":
		impl, ok := w.Concrete.(SubstanceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchSubstance(ctx, options)
	case "SubstanceDefinition":
		impl, ok := w.Concrete.(SubstanceDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceDefinition"}
		}
		return impl.SearchSubstanceDefinition(ctx, options)
	case "SubstanceNucleicAcid":
		impl, ok := w.Concrete.(SubstanceNucleicAcidSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.SearchSubstanceNucleicAcid(ctx, options)
	case "SubstancePolymer":
		impl, ok := w.Concrete.(SubstancePolymerSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstancePolymer"}
		}
		return impl.SearchSubstancePolymer(ctx, options)
	case "SubstanceProtein":
		impl, ok := w.Concrete.(SubstanceProteinSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceProtein"}
		}
		return impl.SearchSubstanceProtein(ctx, options)
	case "SubstanceReferenceInformation":
		impl, ok := w.Concrete.(SubstanceReferenceInformationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.SearchSubstanceReferenceInformation(ctx, options)
	case "SubstanceSourceMaterial":
		impl, ok := w.Concrete.(SubstanceSourceMaterialSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.SearchSubstanceSourceMaterial(ctx, options)
	case "SupplyDelivery":
		impl, ok := w.Concrete.(SupplyDeliverySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyDelivery"}
		}
		return impl.SearchSupplyDelivery(ctx, options)
	case "SupplyRequest":
		impl, ok := w.Concrete.(SupplyRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyRequest"}
		}
		return impl.SearchSupplyRequest(ctx, options)
	case "Task":
		impl, ok := w.Concrete.(TaskSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Task"}
		}
		return impl.SearchTask(ctx, options)
	case "TerminologyCapabilities":
		impl, ok := w.Concrete.(TerminologyCapabilitiesSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TerminologyCapabilities"}
		}
		return impl.SearchTerminologyCapabilities(ctx, options)
	case "TestPlan":
		impl, ok := w.Concrete.(TestPlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestPlan"}
		}
		return impl.SearchTestPlan(ctx, options)
	case "TestReport":
		impl, ok := w.Concrete.(TestReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestReport"}
		}
		return impl.SearchTestReport(ctx, options)
	case "TestScript":
		impl, ok := w.Concrete.(TestScriptSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestScript"}
		}
		return impl.SearchTestScript(ctx, options)
	case "Transport":
		impl, ok := w.Concrete.(TransportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Transport"}
		}
		return impl.SearchTransport(ctx, options)
	case "ValueSet":
		impl, ok := w.Concrete.(ValueSetSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ValueSet"}
		}
		return impl.SearchValueSet(ctx, options)
	case "VerificationResult":
		impl, ok := w.Concrete.(VerificationResultSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VerificationResult"}
		}
		return impl.SearchVerificationResult(ctx, options)
	case "VisionPrescription":
		impl, ok := w.Concrete.(VisionPrescriptionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VisionPrescription"}
		}
		return impl.SearchVisionPrescription(ctx, options)
	default:
		return search.Result{}, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
