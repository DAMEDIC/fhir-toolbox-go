package genericR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/capabilities/gen/r4"
	model "fhir-toolbox/model"
)

func (w InternalWrapper) Read(ctx context.Context, resourceType string, id string) (model.Resource, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.API.(r4.AccountRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Account"}
		}
		return impl.ReadAccount(ctx, id)
	case "ActivityDefinition":
		impl, ok := w.API.(r4.ActivityDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ActivityDefinition"}
		}
		return impl.ReadActivityDefinition(ctx, id)
	case "AdverseEvent":
		impl, ok := w.API.(r4.AdverseEventRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AdverseEvent"}
		}
		return impl.ReadAdverseEvent(ctx, id)
	case "AllergyIntolerance":
		impl, ok := w.API.(r4.AllergyIntoleranceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AllergyIntolerance"}
		}
		return impl.ReadAllergyIntolerance(ctx, id)
	case "Appointment":
		impl, ok := w.API.(r4.AppointmentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Appointment"}
		}
		return impl.ReadAppointment(ctx, id)
	case "AppointmentResponse":
		impl, ok := w.API.(r4.AppointmentResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AppointmentResponse"}
		}
		return impl.ReadAppointmentResponse(ctx, id)
	case "AuditEvent":
		impl, ok := w.API.(r4.AuditEventRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "AuditEvent"}
		}
		return impl.ReadAuditEvent(ctx, id)
	case "Basic":
		impl, ok := w.API.(r4.BasicRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Basic"}
		}
		return impl.ReadBasic(ctx, id)
	case "Binary":
		impl, ok := w.API.(r4.BinaryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Binary"}
		}
		return impl.ReadBinary(ctx, id)
	case "BiologicallyDerivedProduct":
		impl, ok := w.API.(r4.BiologicallyDerivedProductRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.ReadBiologicallyDerivedProduct(ctx, id)
	case "BodyStructure":
		impl, ok := w.API.(r4.BodyStructureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "BodyStructure"}
		}
		return impl.ReadBodyStructure(ctx, id)
	case "Bundle":
		impl, ok := w.API.(r4.BundleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Bundle"}
		}
		return impl.ReadBundle(ctx, id)
	case "CapabilityStatement":
		impl, ok := w.API.(r4.CapabilityStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CapabilityStatement"}
		}
		return impl.ReadCapabilityStatement(ctx, id)
	case "CarePlan":
		impl, ok := w.API.(r4.CarePlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CarePlan"}
		}
		return impl.ReadCarePlan(ctx, id)
	case "CareTeam":
		impl, ok := w.API.(r4.CareTeamRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CareTeam"}
		}
		return impl.ReadCareTeam(ctx, id)
	case "CatalogEntry":
		impl, ok := w.API.(r4.CatalogEntryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CatalogEntry"}
		}
		return impl.ReadCatalogEntry(ctx, id)
	case "ChargeItem":
		impl, ok := w.API.(r4.ChargeItemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ChargeItem"}
		}
		return impl.ReadChargeItem(ctx, id)
	case "ChargeItemDefinition":
		impl, ok := w.API.(r4.ChargeItemDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ChargeItemDefinition"}
		}
		return impl.ReadChargeItemDefinition(ctx, id)
	case "Claim":
		impl, ok := w.API.(r4.ClaimRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Claim"}
		}
		return impl.ReadClaim(ctx, id)
	case "ClaimResponse":
		impl, ok := w.API.(r4.ClaimResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ClaimResponse"}
		}
		return impl.ReadClaimResponse(ctx, id)
	case "ClinicalImpression":
		impl, ok := w.API.(r4.ClinicalImpressionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ClinicalImpression"}
		}
		return impl.ReadClinicalImpression(ctx, id)
	case "CodeSystem":
		impl, ok := w.API.(r4.CodeSystemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CodeSystem"}
		}
		return impl.ReadCodeSystem(ctx, id)
	case "Communication":
		impl, ok := w.API.(r4.CommunicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Communication"}
		}
		return impl.ReadCommunication(ctx, id)
	case "CommunicationRequest":
		impl, ok := w.API.(r4.CommunicationRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CommunicationRequest"}
		}
		return impl.ReadCommunicationRequest(ctx, id)
	case "CompartmentDefinition":
		impl, ok := w.API.(r4.CompartmentDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CompartmentDefinition"}
		}
		return impl.ReadCompartmentDefinition(ctx, id)
	case "Composition":
		impl, ok := w.API.(r4.CompositionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Composition"}
		}
		return impl.ReadComposition(ctx, id)
	case "ConceptMap":
		impl, ok := w.API.(r4.ConceptMapRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ConceptMap"}
		}
		return impl.ReadConceptMap(ctx, id)
	case "Condition":
		impl, ok := w.API.(r4.ConditionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Condition"}
		}
		return impl.ReadCondition(ctx, id)
	case "Consent":
		impl, ok := w.API.(r4.ConsentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Consent"}
		}
		return impl.ReadConsent(ctx, id)
	case "Contract":
		impl, ok := w.API.(r4.ContractRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Contract"}
		}
		return impl.ReadContract(ctx, id)
	case "Coverage":
		impl, ok := w.API.(r4.CoverageRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Coverage"}
		}
		return impl.ReadCoverage(ctx, id)
	case "CoverageEligibilityRequest":
		impl, ok := w.API.(r4.CoverageEligibilityRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.ReadCoverageEligibilityRequest(ctx, id)
	case "CoverageEligibilityResponse":
		impl, ok := w.API.(r4.CoverageEligibilityResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.ReadCoverageEligibilityResponse(ctx, id)
	case "DetectedIssue":
		impl, ok := w.API.(r4.DetectedIssueRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DetectedIssue"}
		}
		return impl.ReadDetectedIssue(ctx, id)
	case "Device":
		impl, ok := w.API.(r4.DeviceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Device"}
		}
		return impl.ReadDevice(ctx, id)
	case "DeviceDefinition":
		impl, ok := w.API.(r4.DeviceDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceDefinition"}
		}
		return impl.ReadDeviceDefinition(ctx, id)
	case "DeviceMetric":
		impl, ok := w.API.(r4.DeviceMetricRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceMetric"}
		}
		return impl.ReadDeviceMetric(ctx, id)
	case "DeviceRequest":
		impl, ok := w.API.(r4.DeviceRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceRequest"}
		}
		return impl.ReadDeviceRequest(ctx, id)
	case "DeviceUseStatement":
		impl, ok := w.API.(r4.DeviceUseStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceUseStatement"}
		}
		return impl.ReadDeviceUseStatement(ctx, id)
	case "DiagnosticReport":
		impl, ok := w.API.(r4.DiagnosticReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DiagnosticReport"}
		}
		return impl.ReadDiagnosticReport(ctx, id)
	case "DocumentManifest":
		impl, ok := w.API.(r4.DocumentManifestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DocumentManifest"}
		}
		return impl.ReadDocumentManifest(ctx, id)
	case "DocumentReference":
		impl, ok := w.API.(r4.DocumentReferenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DocumentReference"}
		}
		return impl.ReadDocumentReference(ctx, id)
	case "EffectEvidenceSynthesis":
		impl, ok := w.API.(r4.EffectEvidenceSynthesisRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.ReadEffectEvidenceSynthesis(ctx, id)
	case "Encounter":
		impl, ok := w.API.(r4.EncounterRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Encounter"}
		}
		return impl.ReadEncounter(ctx, id)
	case "Endpoint":
		impl, ok := w.API.(r4.EndpointRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Endpoint"}
		}
		return impl.ReadEndpoint(ctx, id)
	case "EnrollmentRequest":
		impl, ok := w.API.(r4.EnrollmentRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EnrollmentRequest"}
		}
		return impl.ReadEnrollmentRequest(ctx, id)
	case "EnrollmentResponse":
		impl, ok := w.API.(r4.EnrollmentResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EnrollmentResponse"}
		}
		return impl.ReadEnrollmentResponse(ctx, id)
	case "EpisodeOfCare":
		impl, ok := w.API.(r4.EpisodeOfCareRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EpisodeOfCare"}
		}
		return impl.ReadEpisodeOfCare(ctx, id)
	case "EventDefinition":
		impl, ok := w.API.(r4.EventDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EventDefinition"}
		}
		return impl.ReadEventDefinition(ctx, id)
	case "Evidence":
		impl, ok := w.API.(r4.EvidenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Evidence"}
		}
		return impl.ReadEvidence(ctx, id)
	case "EvidenceVariable":
		impl, ok := w.API.(r4.EvidenceVariableRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EvidenceVariable"}
		}
		return impl.ReadEvidenceVariable(ctx, id)
	case "ExampleScenario":
		impl, ok := w.API.(r4.ExampleScenarioRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ExampleScenario"}
		}
		return impl.ReadExampleScenario(ctx, id)
	case "ExplanationOfBenefit":
		impl, ok := w.API.(r4.ExplanationOfBenefitRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.ReadExplanationOfBenefit(ctx, id)
	case "FamilyMemberHistory":
		impl, ok := w.API.(r4.FamilyMemberHistoryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "FamilyMemberHistory"}
		}
		return impl.ReadFamilyMemberHistory(ctx, id)
	case "Flag":
		impl, ok := w.API.(r4.FlagRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Flag"}
		}
		return impl.ReadFlag(ctx, id)
	case "Goal":
		impl, ok := w.API.(r4.GoalRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Goal"}
		}
		return impl.ReadGoal(ctx, id)
	case "GraphDefinition":
		impl, ok := w.API.(r4.GraphDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "GraphDefinition"}
		}
		return impl.ReadGraphDefinition(ctx, id)
	case "Group":
		impl, ok := w.API.(r4.GroupRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Group"}
		}
		return impl.ReadGroup(ctx, id)
	case "GuidanceResponse":
		impl, ok := w.API.(r4.GuidanceResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "GuidanceResponse"}
		}
		return impl.ReadGuidanceResponse(ctx, id)
	case "HealthcareService":
		impl, ok := w.API.(r4.HealthcareServiceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "HealthcareService"}
		}
		return impl.ReadHealthcareService(ctx, id)
	case "ImagingStudy":
		impl, ok := w.API.(r4.ImagingStudyRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImagingStudy"}
		}
		return impl.ReadImagingStudy(ctx, id)
	case "Immunization":
		impl, ok := w.API.(r4.ImmunizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Immunization"}
		}
		return impl.ReadImmunization(ctx, id)
	case "ImmunizationEvaluation":
		impl, ok := w.API.(r4.ImmunizationEvaluationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.ReadImmunizationEvaluation(ctx, id)
	case "ImmunizationRecommendation":
		impl, ok := w.API.(r4.ImmunizationRecommendationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.ReadImmunizationRecommendation(ctx, id)
	case "ImplementationGuide":
		impl, ok := w.API.(r4.ImplementationGuideRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ImplementationGuide"}
		}
		return impl.ReadImplementationGuide(ctx, id)
	case "InsurancePlan":
		impl, ok := w.API.(r4.InsurancePlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "InsurancePlan"}
		}
		return impl.ReadInsurancePlan(ctx, id)
	case "Invoice":
		impl, ok := w.API.(r4.InvoiceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Invoice"}
		}
		return impl.ReadInvoice(ctx, id)
	case "Library":
		impl, ok := w.API.(r4.LibraryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Library"}
		}
		return impl.ReadLibrary(ctx, id)
	case "Linkage":
		impl, ok := w.API.(r4.LinkageRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Linkage"}
		}
		return impl.ReadLinkage(ctx, id)
	case "List":
		impl, ok := w.API.(r4.ListRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "List"}
		}
		return impl.ReadList(ctx, id)
	case "Location":
		impl, ok := w.API.(r4.LocationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Location"}
		}
		return impl.ReadLocation(ctx, id)
	case "Measure":
		impl, ok := w.API.(r4.MeasureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Measure"}
		}
		return impl.ReadMeasure(ctx, id)
	case "MeasureReport":
		impl, ok := w.API.(r4.MeasureReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MeasureReport"}
		}
		return impl.ReadMeasureReport(ctx, id)
	case "Media":
		impl, ok := w.API.(r4.MediaRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Media"}
		}
		return impl.ReadMedia(ctx, id)
	case "Medication":
		impl, ok := w.API.(r4.MedicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Medication"}
		}
		return impl.ReadMedication(ctx, id)
	case "MedicationAdministration":
		impl, ok := w.API.(r4.MedicationAdministrationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationAdministration"}
		}
		return impl.ReadMedicationAdministration(ctx, id)
	case "MedicationDispense":
		impl, ok := w.API.(r4.MedicationDispenseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationDispense"}
		}
		return impl.ReadMedicationDispense(ctx, id)
	case "MedicationKnowledge":
		impl, ok := w.API.(r4.MedicationKnowledgeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationKnowledge"}
		}
		return impl.ReadMedicationKnowledge(ctx, id)
	case "MedicationRequest":
		impl, ok := w.API.(r4.MedicationRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationRequest"}
		}
		return impl.ReadMedicationRequest(ctx, id)
	case "MedicationStatement":
		impl, ok := w.API.(r4.MedicationStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicationStatement"}
		}
		return impl.ReadMedicationStatement(ctx, id)
	case "MedicinalProduct":
		impl, ok := w.API.(r4.MedicinalProductRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProduct"}
		}
		return impl.ReadMedicinalProduct(ctx, id)
	case "MedicinalProductAuthorization":
		impl, ok := w.API.(r4.MedicinalProductAuthorizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.ReadMedicinalProductAuthorization(ctx, id)
	case "MedicinalProductContraindication":
		impl, ok := w.API.(r4.MedicinalProductContraindicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.ReadMedicinalProductContraindication(ctx, id)
	case "MedicinalProductIndication":
		impl, ok := w.API.(r4.MedicinalProductIndicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductIndication"}
		}
		return impl.ReadMedicinalProductIndication(ctx, id)
	case "MedicinalProductIngredient":
		impl, ok := w.API.(r4.MedicinalProductIngredientRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.ReadMedicinalProductIngredient(ctx, id)
	case "MedicinalProductInteraction":
		impl, ok := w.API.(r4.MedicinalProductInteractionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.ReadMedicinalProductInteraction(ctx, id)
	case "MedicinalProductManufactured":
		impl, ok := w.API.(r4.MedicinalProductManufacturedRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.ReadMedicinalProductManufactured(ctx, id)
	case "MedicinalProductPackaged":
		impl, ok := w.API.(r4.MedicinalProductPackagedRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.ReadMedicinalProductPackaged(ctx, id)
	case "MedicinalProductPharmaceutical":
		impl, ok := w.API.(r4.MedicinalProductPharmaceuticalRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.ReadMedicinalProductPharmaceutical(ctx, id)
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.API.(r4.MedicinalProductUndesirableEffectRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.ReadMedicinalProductUndesirableEffect(ctx, id)
	case "MessageDefinition":
		impl, ok := w.API.(r4.MessageDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MessageDefinition"}
		}
		return impl.ReadMessageDefinition(ctx, id)
	case "MessageHeader":
		impl, ok := w.API.(r4.MessageHeaderRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MessageHeader"}
		}
		return impl.ReadMessageHeader(ctx, id)
	case "MolecularSequence":
		impl, ok := w.API.(r4.MolecularSequenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MolecularSequence"}
		}
		return impl.ReadMolecularSequence(ctx, id)
	case "NamingSystem":
		impl, ok := w.API.(r4.NamingSystemRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NamingSystem"}
		}
		return impl.ReadNamingSystem(ctx, id)
	case "NutritionOrder":
		impl, ok := w.API.(r4.NutritionOrderRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NutritionOrder"}
		}
		return impl.ReadNutritionOrder(ctx, id)
	case "Observation":
		impl, ok := w.API.(r4.ObservationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Observation"}
		}
		return impl.ReadObservation(ctx, id)
	case "ObservationDefinition":
		impl, ok := w.API.(r4.ObservationDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ObservationDefinition"}
		}
		return impl.ReadObservationDefinition(ctx, id)
	case "OperationDefinition":
		impl, ok := w.API.(r4.OperationDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OperationDefinition"}
		}
		return impl.ReadOperationDefinition(ctx, id)
	case "OperationOutcome":
		impl, ok := w.API.(r4.OperationOutcomeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OperationOutcome"}
		}
		return impl.ReadOperationOutcome(ctx, id)
	case "Organization":
		impl, ok := w.API.(r4.OrganizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Organization"}
		}
		return impl.ReadOrganization(ctx, id)
	case "OrganizationAffiliation":
		impl, ok := w.API.(r4.OrganizationAffiliationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "OrganizationAffiliation"}
		}
		return impl.ReadOrganizationAffiliation(ctx, id)
	case "Parameters":
		impl, ok := w.API.(r4.ParametersRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Parameters"}
		}
		return impl.ReadParameters(ctx, id)
	case "Patient":
		impl, ok := w.API.(r4.PatientRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Patient"}
		}
		return impl.ReadPatient(ctx, id)
	case "PaymentNotice":
		impl, ok := w.API.(r4.PaymentNoticeRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PaymentNotice"}
		}
		return impl.ReadPaymentNotice(ctx, id)
	case "PaymentReconciliation":
		impl, ok := w.API.(r4.PaymentReconciliationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PaymentReconciliation"}
		}
		return impl.ReadPaymentReconciliation(ctx, id)
	case "Person":
		impl, ok := w.API.(r4.PersonRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Person"}
		}
		return impl.ReadPerson(ctx, id)
	case "PlanDefinition":
		impl, ok := w.API.(r4.PlanDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PlanDefinition"}
		}
		return impl.ReadPlanDefinition(ctx, id)
	case "Practitioner":
		impl, ok := w.API.(r4.PractitionerRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Practitioner"}
		}
		return impl.ReadPractitioner(ctx, id)
	case "PractitionerRole":
		impl, ok := w.API.(r4.PractitionerRoleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "PractitionerRole"}
		}
		return impl.ReadPractitionerRole(ctx, id)
	case "Procedure":
		impl, ok := w.API.(r4.ProcedureRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Procedure"}
		}
		return impl.ReadProcedure(ctx, id)
	case "Provenance":
		impl, ok := w.API.(r4.ProvenanceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Provenance"}
		}
		return impl.ReadProvenance(ctx, id)
	case "Questionnaire":
		impl, ok := w.API.(r4.QuestionnaireRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Questionnaire"}
		}
		return impl.ReadQuestionnaire(ctx, id)
	case "QuestionnaireResponse":
		impl, ok := w.API.(r4.QuestionnaireResponseRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "QuestionnaireResponse"}
		}
		return impl.ReadQuestionnaireResponse(ctx, id)
	case "RelatedPerson":
		impl, ok := w.API.(r4.RelatedPersonRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RelatedPerson"}
		}
		return impl.ReadRelatedPerson(ctx, id)
	case "RequestGroup":
		impl, ok := w.API.(r4.RequestGroupRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RequestGroup"}
		}
		return impl.ReadRequestGroup(ctx, id)
	case "ResearchDefinition":
		impl, ok := w.API.(r4.ResearchDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchDefinition"}
		}
		return impl.ReadResearchDefinition(ctx, id)
	case "ResearchElementDefinition":
		impl, ok := w.API.(r4.ResearchElementDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchElementDefinition"}
		}
		return impl.ReadResearchElementDefinition(ctx, id)
	case "ResearchStudy":
		impl, ok := w.API.(r4.ResearchStudyRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchStudy"}
		}
		return impl.ReadResearchStudy(ctx, id)
	case "ResearchSubject":
		impl, ok := w.API.(r4.ResearchSubjectRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchSubject"}
		}
		return impl.ReadResearchSubject(ctx, id)
	case "RiskAssessment":
		impl, ok := w.API.(r4.RiskAssessmentRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RiskAssessment"}
		}
		return impl.ReadRiskAssessment(ctx, id)
	case "RiskEvidenceSynthesis":
		impl, ok := w.API.(r4.RiskEvidenceSynthesisRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.ReadRiskEvidenceSynthesis(ctx, id)
	case "Schedule":
		impl, ok := w.API.(r4.ScheduleRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Schedule"}
		}
		return impl.ReadSchedule(ctx, id)
	case "SearchParameter":
		impl, ok := w.API.(r4.SearchParameterRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SearchParameter"}
		}
		return impl.ReadSearchParameter(ctx, id)
	case "ServiceRequest":
		impl, ok := w.API.(r4.ServiceRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ServiceRequest"}
		}
		return impl.ReadServiceRequest(ctx, id)
	case "Slot":
		impl, ok := w.API.(r4.SlotRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Slot"}
		}
		return impl.ReadSlot(ctx, id)
	case "Specimen":
		impl, ok := w.API.(r4.SpecimenRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Specimen"}
		}
		return impl.ReadSpecimen(ctx, id)
	case "SpecimenDefinition":
		impl, ok := w.API.(r4.SpecimenDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SpecimenDefinition"}
		}
		return impl.ReadSpecimenDefinition(ctx, id)
	case "StructureDefinition":
		impl, ok := w.API.(r4.StructureDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "StructureDefinition"}
		}
		return impl.ReadStructureDefinition(ctx, id)
	case "StructureMap":
		impl, ok := w.API.(r4.StructureMapRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "StructureMap"}
		}
		return impl.ReadStructureMap(ctx, id)
	case "Subscription":
		impl, ok := w.API.(r4.SubscriptionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Subscription"}
		}
		return impl.ReadSubscription(ctx, id)
	case "Substance":
		impl, ok := w.API.(r4.SubstanceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Substance"}
		}
		return impl.ReadSubstance(ctx, id)
	case "SubstanceNucleicAcid":
		impl, ok := w.API.(r4.SubstanceNucleicAcidRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.ReadSubstanceNucleicAcid(ctx, id)
	case "SubstancePolymer":
		impl, ok := w.API.(r4.SubstancePolymerRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstancePolymer"}
		}
		return impl.ReadSubstancePolymer(ctx, id)
	case "SubstanceProtein":
		impl, ok := w.API.(r4.SubstanceProteinRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceProtein"}
		}
		return impl.ReadSubstanceProtein(ctx, id)
	case "SubstanceReferenceInformation":
		impl, ok := w.API.(r4.SubstanceReferenceInformationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.ReadSubstanceReferenceInformation(ctx, id)
	case "SubstanceSourceMaterial":
		impl, ok := w.API.(r4.SubstanceSourceMaterialRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.ReadSubstanceSourceMaterial(ctx, id)
	case "SubstanceSpecification":
		impl, ok := w.API.(r4.SubstanceSpecificationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceSpecification"}
		}
		return impl.ReadSubstanceSpecification(ctx, id)
	case "SupplyDelivery":
		impl, ok := w.API.(r4.SupplyDeliveryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SupplyDelivery"}
		}
		return impl.ReadSupplyDelivery(ctx, id)
	case "SupplyRequest":
		impl, ok := w.API.(r4.SupplyRequestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SupplyRequest"}
		}
		return impl.ReadSupplyRequest(ctx, id)
	case "Task":
		impl, ok := w.API.(r4.TaskRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Task"}
		}
		return impl.ReadTask(ctx, id)
	case "TerminologyCapabilities":
		impl, ok := w.API.(r4.TerminologyCapabilitiesRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TerminologyCapabilities"}
		}
		return impl.ReadTerminologyCapabilities(ctx, id)
	case "TestReport":
		impl, ok := w.API.(r4.TestReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TestReport"}
		}
		return impl.ReadTestReport(ctx, id)
	case "TestScript":
		impl, ok := w.API.(r4.TestScriptRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "TestScript"}
		}
		return impl.ReadTestScript(ctx, id)
	case "ValueSet":
		impl, ok := w.API.(r4.ValueSetRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ValueSet"}
		}
		return impl.ReadValueSet(ctx, id)
	case "VerificationResult":
		impl, ok := w.API.(r4.VerificationResultRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "VerificationResult"}
		}
		return impl.ReadVerificationResult(ctx, id)
	case "VisionPrescription":
		impl, ok := w.API.(r4.VisionPrescriptionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "VisionPrescription"}
		}
		return impl.ReadVisionPrescription(ctx, id)
	default:
		return nil, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
