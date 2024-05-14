package dispatchR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/capabilities/gen/r4"
	errors "fhir-toolbox/dispatch/errors"
	model "fhir-toolbox/model"
)

func Read(ctx context.Context, api any, resourceType string, id string) (model.Resource, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := api.(r4.AccountRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Account"}
		}
		return impl.ReadAccount(ctx, id)
	case "ActivityDefinition":
		impl, ok := api.(r4.ActivityDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ActivityDefinition"}
		}
		return impl.ReadActivityDefinition(ctx, id)
	case "AdverseEvent":
		impl, ok := api.(r4.AdverseEventRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "AdverseEvent"}
		}
		return impl.ReadAdverseEvent(ctx, id)
	case "AllergyIntolerance":
		impl, ok := api.(r4.AllergyIntoleranceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "AllergyIntolerance"}
		}
		return impl.ReadAllergyIntolerance(ctx, id)
	case "Appointment":
		impl, ok := api.(r4.AppointmentRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Appointment"}
		}
		return impl.ReadAppointment(ctx, id)
	case "AppointmentResponse":
		impl, ok := api.(r4.AppointmentResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "AppointmentResponse"}
		}
		return impl.ReadAppointmentResponse(ctx, id)
	case "AuditEvent":
		impl, ok := api.(r4.AuditEventRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "AuditEvent"}
		}
		return impl.ReadAuditEvent(ctx, id)
	case "Basic":
		impl, ok := api.(r4.BasicRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Basic"}
		}
		return impl.ReadBasic(ctx, id)
	case "Binary":
		impl, ok := api.(r4.BinaryRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Binary"}
		}
		return impl.ReadBinary(ctx, id)
	case "BiologicallyDerivedProduct":
		impl, ok := api.(r4.BiologicallyDerivedProductRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.ReadBiologicallyDerivedProduct(ctx, id)
	case "BodyStructure":
		impl, ok := api.(r4.BodyStructureRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "BodyStructure"}
		}
		return impl.ReadBodyStructure(ctx, id)
	case "Bundle":
		impl, ok := api.(r4.BundleRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Bundle"}
		}
		return impl.ReadBundle(ctx, id)
	case "CapabilityStatement":
		impl, ok := api.(r4.CapabilityStatementRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CapabilityStatement"}
		}
		return impl.ReadCapabilityStatement(ctx, id)
	case "CarePlan":
		impl, ok := api.(r4.CarePlanRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CarePlan"}
		}
		return impl.ReadCarePlan(ctx, id)
	case "CareTeam":
		impl, ok := api.(r4.CareTeamRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CareTeam"}
		}
		return impl.ReadCareTeam(ctx, id)
	case "CatalogEntry":
		impl, ok := api.(r4.CatalogEntryRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CatalogEntry"}
		}
		return impl.ReadCatalogEntry(ctx, id)
	case "ChargeItem":
		impl, ok := api.(r4.ChargeItemRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ChargeItem"}
		}
		return impl.ReadChargeItem(ctx, id)
	case "ChargeItemDefinition":
		impl, ok := api.(r4.ChargeItemDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ChargeItemDefinition"}
		}
		return impl.ReadChargeItemDefinition(ctx, id)
	case "Claim":
		impl, ok := api.(r4.ClaimRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Claim"}
		}
		return impl.ReadClaim(ctx, id)
	case "ClaimResponse":
		impl, ok := api.(r4.ClaimResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ClaimResponse"}
		}
		return impl.ReadClaimResponse(ctx, id)
	case "ClinicalImpression":
		impl, ok := api.(r4.ClinicalImpressionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ClinicalImpression"}
		}
		return impl.ReadClinicalImpression(ctx, id)
	case "CodeSystem":
		impl, ok := api.(r4.CodeSystemRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CodeSystem"}
		}
		return impl.ReadCodeSystem(ctx, id)
	case "Communication":
		impl, ok := api.(r4.CommunicationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Communication"}
		}
		return impl.ReadCommunication(ctx, id)
	case "CommunicationRequest":
		impl, ok := api.(r4.CommunicationRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CommunicationRequest"}
		}
		return impl.ReadCommunicationRequest(ctx, id)
	case "CompartmentDefinition":
		impl, ok := api.(r4.CompartmentDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CompartmentDefinition"}
		}
		return impl.ReadCompartmentDefinition(ctx, id)
	case "Composition":
		impl, ok := api.(r4.CompositionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Composition"}
		}
		return impl.ReadComposition(ctx, id)
	case "ConceptMap":
		impl, ok := api.(r4.ConceptMapRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ConceptMap"}
		}
		return impl.ReadConceptMap(ctx, id)
	case "Condition":
		impl, ok := api.(r4.ConditionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Condition"}
		}
		return impl.ReadCondition(ctx, id)
	case "Consent":
		impl, ok := api.(r4.ConsentRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Consent"}
		}
		return impl.ReadConsent(ctx, id)
	case "Contract":
		impl, ok := api.(r4.ContractRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Contract"}
		}
		return impl.ReadContract(ctx, id)
	case "Coverage":
		impl, ok := api.(r4.CoverageRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Coverage"}
		}
		return impl.ReadCoverage(ctx, id)
	case "CoverageEligibilityRequest":
		impl, ok := api.(r4.CoverageEligibilityRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.ReadCoverageEligibilityRequest(ctx, id)
	case "CoverageEligibilityResponse":
		impl, ok := api.(r4.CoverageEligibilityResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.ReadCoverageEligibilityResponse(ctx, id)
	case "DetectedIssue":
		impl, ok := api.(r4.DetectedIssueRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DetectedIssue"}
		}
		return impl.ReadDetectedIssue(ctx, id)
	case "Device":
		impl, ok := api.(r4.DeviceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Device"}
		}
		return impl.ReadDevice(ctx, id)
	case "DeviceDefinition":
		impl, ok := api.(r4.DeviceDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DeviceDefinition"}
		}
		return impl.ReadDeviceDefinition(ctx, id)
	case "DeviceMetric":
		impl, ok := api.(r4.DeviceMetricRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DeviceMetric"}
		}
		return impl.ReadDeviceMetric(ctx, id)
	case "DeviceRequest":
		impl, ok := api.(r4.DeviceRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DeviceRequest"}
		}
		return impl.ReadDeviceRequest(ctx, id)
	case "DeviceUseStatement":
		impl, ok := api.(r4.DeviceUseStatementRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DeviceUseStatement"}
		}
		return impl.ReadDeviceUseStatement(ctx, id)
	case "DiagnosticReport":
		impl, ok := api.(r4.DiagnosticReportRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DiagnosticReport"}
		}
		return impl.ReadDiagnosticReport(ctx, id)
	case "DocumentManifest":
		impl, ok := api.(r4.DocumentManifestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DocumentManifest"}
		}
		return impl.ReadDocumentManifest(ctx, id)
	case "DocumentReference":
		impl, ok := api.(r4.DocumentReferenceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "DocumentReference"}
		}
		return impl.ReadDocumentReference(ctx, id)
	case "EffectEvidenceSynthesis":
		impl, ok := api.(r4.EffectEvidenceSynthesisRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.ReadEffectEvidenceSynthesis(ctx, id)
	case "Encounter":
		impl, ok := api.(r4.EncounterRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Encounter"}
		}
		return impl.ReadEncounter(ctx, id)
	case "Endpoint":
		impl, ok := api.(r4.EndpointRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Endpoint"}
		}
		return impl.ReadEndpoint(ctx, id)
	case "EnrollmentRequest":
		impl, ok := api.(r4.EnrollmentRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EnrollmentRequest"}
		}
		return impl.ReadEnrollmentRequest(ctx, id)
	case "EnrollmentResponse":
		impl, ok := api.(r4.EnrollmentResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EnrollmentResponse"}
		}
		return impl.ReadEnrollmentResponse(ctx, id)
	case "EpisodeOfCare":
		impl, ok := api.(r4.EpisodeOfCareRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EpisodeOfCare"}
		}
		return impl.ReadEpisodeOfCare(ctx, id)
	case "EventDefinition":
		impl, ok := api.(r4.EventDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EventDefinition"}
		}
		return impl.ReadEventDefinition(ctx, id)
	case "Evidence":
		impl, ok := api.(r4.EvidenceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Evidence"}
		}
		return impl.ReadEvidence(ctx, id)
	case "EvidenceVariable":
		impl, ok := api.(r4.EvidenceVariableRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "EvidenceVariable"}
		}
		return impl.ReadEvidenceVariable(ctx, id)
	case "ExampleScenario":
		impl, ok := api.(r4.ExampleScenarioRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ExampleScenario"}
		}
		return impl.ReadExampleScenario(ctx, id)
	case "ExplanationOfBenefit":
		impl, ok := api.(r4.ExplanationOfBenefitRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.ReadExplanationOfBenefit(ctx, id)
	case "FamilyMemberHistory":
		impl, ok := api.(r4.FamilyMemberHistoryRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "FamilyMemberHistory"}
		}
		return impl.ReadFamilyMemberHistory(ctx, id)
	case "Flag":
		impl, ok := api.(r4.FlagRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Flag"}
		}
		return impl.ReadFlag(ctx, id)
	case "Goal":
		impl, ok := api.(r4.GoalRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Goal"}
		}
		return impl.ReadGoal(ctx, id)
	case "GraphDefinition":
		impl, ok := api.(r4.GraphDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "GraphDefinition"}
		}
		return impl.ReadGraphDefinition(ctx, id)
	case "Group":
		impl, ok := api.(r4.GroupRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Group"}
		}
		return impl.ReadGroup(ctx, id)
	case "GuidanceResponse":
		impl, ok := api.(r4.GuidanceResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "GuidanceResponse"}
		}
		return impl.ReadGuidanceResponse(ctx, id)
	case "HealthcareService":
		impl, ok := api.(r4.HealthcareServiceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "HealthcareService"}
		}
		return impl.ReadHealthcareService(ctx, id)
	case "ImagingStudy":
		impl, ok := api.(r4.ImagingStudyRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ImagingStudy"}
		}
		return impl.ReadImagingStudy(ctx, id)
	case "Immunization":
		impl, ok := api.(r4.ImmunizationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Immunization"}
		}
		return impl.ReadImmunization(ctx, id)
	case "ImmunizationEvaluation":
		impl, ok := api.(r4.ImmunizationEvaluationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.ReadImmunizationEvaluation(ctx, id)
	case "ImmunizationRecommendation":
		impl, ok := api.(r4.ImmunizationRecommendationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.ReadImmunizationRecommendation(ctx, id)
	case "ImplementationGuide":
		impl, ok := api.(r4.ImplementationGuideRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ImplementationGuide"}
		}
		return impl.ReadImplementationGuide(ctx, id)
	case "InsurancePlan":
		impl, ok := api.(r4.InsurancePlanRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "InsurancePlan"}
		}
		return impl.ReadInsurancePlan(ctx, id)
	case "Invoice":
		impl, ok := api.(r4.InvoiceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Invoice"}
		}
		return impl.ReadInvoice(ctx, id)
	case "Library":
		impl, ok := api.(r4.LibraryRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Library"}
		}
		return impl.ReadLibrary(ctx, id)
	case "Linkage":
		impl, ok := api.(r4.LinkageRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Linkage"}
		}
		return impl.ReadLinkage(ctx, id)
	case "List":
		impl, ok := api.(r4.ListRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "List"}
		}
		return impl.ReadList(ctx, id)
	case "Location":
		impl, ok := api.(r4.LocationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Location"}
		}
		return impl.ReadLocation(ctx, id)
	case "Measure":
		impl, ok := api.(r4.MeasureRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Measure"}
		}
		return impl.ReadMeasure(ctx, id)
	case "MeasureReport":
		impl, ok := api.(r4.MeasureReportRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MeasureReport"}
		}
		return impl.ReadMeasureReport(ctx, id)
	case "Media":
		impl, ok := api.(r4.MediaRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Media"}
		}
		return impl.ReadMedia(ctx, id)
	case "Medication":
		impl, ok := api.(r4.MedicationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Medication"}
		}
		return impl.ReadMedication(ctx, id)
	case "MedicationAdministration":
		impl, ok := api.(r4.MedicationAdministrationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicationAdministration"}
		}
		return impl.ReadMedicationAdministration(ctx, id)
	case "MedicationDispense":
		impl, ok := api.(r4.MedicationDispenseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicationDispense"}
		}
		return impl.ReadMedicationDispense(ctx, id)
	case "MedicationKnowledge":
		impl, ok := api.(r4.MedicationKnowledgeRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicationKnowledge"}
		}
		return impl.ReadMedicationKnowledge(ctx, id)
	case "MedicationRequest":
		impl, ok := api.(r4.MedicationRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicationRequest"}
		}
		return impl.ReadMedicationRequest(ctx, id)
	case "MedicationStatement":
		impl, ok := api.(r4.MedicationStatementRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicationStatement"}
		}
		return impl.ReadMedicationStatement(ctx, id)
	case "MedicinalProduct":
		impl, ok := api.(r4.MedicinalProductRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProduct"}
		}
		return impl.ReadMedicinalProduct(ctx, id)
	case "MedicinalProductAuthorization":
		impl, ok := api.(r4.MedicinalProductAuthorizationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.ReadMedicinalProductAuthorization(ctx, id)
	case "MedicinalProductContraindication":
		impl, ok := api.(r4.MedicinalProductContraindicationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.ReadMedicinalProductContraindication(ctx, id)
	case "MedicinalProductIndication":
		impl, ok := api.(r4.MedicinalProductIndicationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductIndication"}
		}
		return impl.ReadMedicinalProductIndication(ctx, id)
	case "MedicinalProductIngredient":
		impl, ok := api.(r4.MedicinalProductIngredientRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.ReadMedicinalProductIngredient(ctx, id)
	case "MedicinalProductInteraction":
		impl, ok := api.(r4.MedicinalProductInteractionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.ReadMedicinalProductInteraction(ctx, id)
	case "MedicinalProductManufactured":
		impl, ok := api.(r4.MedicinalProductManufacturedRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.ReadMedicinalProductManufactured(ctx, id)
	case "MedicinalProductPackaged":
		impl, ok := api.(r4.MedicinalProductPackagedRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.ReadMedicinalProductPackaged(ctx, id)
	case "MedicinalProductPharmaceutical":
		impl, ok := api.(r4.MedicinalProductPharmaceuticalRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.ReadMedicinalProductPharmaceutical(ctx, id)
	case "MedicinalProductUndesirableEffect":
		impl, ok := api.(r4.MedicinalProductUndesirableEffectRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.ReadMedicinalProductUndesirableEffect(ctx, id)
	case "MessageDefinition":
		impl, ok := api.(r4.MessageDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MessageDefinition"}
		}
		return impl.ReadMessageDefinition(ctx, id)
	case "MessageHeader":
		impl, ok := api.(r4.MessageHeaderRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MessageHeader"}
		}
		return impl.ReadMessageHeader(ctx, id)
	case "MolecularSequence":
		impl, ok := api.(r4.MolecularSequenceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "MolecularSequence"}
		}
		return impl.ReadMolecularSequence(ctx, id)
	case "NamingSystem":
		impl, ok := api.(r4.NamingSystemRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "NamingSystem"}
		}
		return impl.ReadNamingSystem(ctx, id)
	case "NutritionOrder":
		impl, ok := api.(r4.NutritionOrderRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "NutritionOrder"}
		}
		return impl.ReadNutritionOrder(ctx, id)
	case "Observation":
		impl, ok := api.(r4.ObservationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Observation"}
		}
		return impl.ReadObservation(ctx, id)
	case "ObservationDefinition":
		impl, ok := api.(r4.ObservationDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ObservationDefinition"}
		}
		return impl.ReadObservationDefinition(ctx, id)
	case "OperationDefinition":
		impl, ok := api.(r4.OperationDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "OperationDefinition"}
		}
		return impl.ReadOperationDefinition(ctx, id)
	case "OperationOutcome":
		impl, ok := api.(r4.OperationOutcomeRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "OperationOutcome"}
		}
		return impl.ReadOperationOutcome(ctx, id)
	case "Organization":
		impl, ok := api.(r4.OrganizationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Organization"}
		}
		return impl.ReadOrganization(ctx, id)
	case "OrganizationAffiliation":
		impl, ok := api.(r4.OrganizationAffiliationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "OrganizationAffiliation"}
		}
		return impl.ReadOrganizationAffiliation(ctx, id)
	case "Parameters":
		impl, ok := api.(r4.ParametersRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Parameters"}
		}
		return impl.ReadParameters(ctx, id)
	case "Patient":
		impl, ok := api.(r4.PatientRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Patient"}
		}
		return impl.ReadPatient(ctx, id)
	case "PaymentNotice":
		impl, ok := api.(r4.PaymentNoticeRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "PaymentNotice"}
		}
		return impl.ReadPaymentNotice(ctx, id)
	case "PaymentReconciliation":
		impl, ok := api.(r4.PaymentReconciliationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "PaymentReconciliation"}
		}
		return impl.ReadPaymentReconciliation(ctx, id)
	case "Person":
		impl, ok := api.(r4.PersonRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Person"}
		}
		return impl.ReadPerson(ctx, id)
	case "PlanDefinition":
		impl, ok := api.(r4.PlanDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "PlanDefinition"}
		}
		return impl.ReadPlanDefinition(ctx, id)
	case "Practitioner":
		impl, ok := api.(r4.PractitionerRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Practitioner"}
		}
		return impl.ReadPractitioner(ctx, id)
	case "PractitionerRole":
		impl, ok := api.(r4.PractitionerRoleRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "PractitionerRole"}
		}
		return impl.ReadPractitionerRole(ctx, id)
	case "Procedure":
		impl, ok := api.(r4.ProcedureRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Procedure"}
		}
		return impl.ReadProcedure(ctx, id)
	case "Provenance":
		impl, ok := api.(r4.ProvenanceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Provenance"}
		}
		return impl.ReadProvenance(ctx, id)
	case "Questionnaire":
		impl, ok := api.(r4.QuestionnaireRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Questionnaire"}
		}
		return impl.ReadQuestionnaire(ctx, id)
	case "QuestionnaireResponse":
		impl, ok := api.(r4.QuestionnaireResponseRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "QuestionnaireResponse"}
		}
		return impl.ReadQuestionnaireResponse(ctx, id)
	case "RelatedPerson":
		impl, ok := api.(r4.RelatedPersonRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "RelatedPerson"}
		}
		return impl.ReadRelatedPerson(ctx, id)
	case "RequestGroup":
		impl, ok := api.(r4.RequestGroupRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "RequestGroup"}
		}
		return impl.ReadRequestGroup(ctx, id)
	case "ResearchDefinition":
		impl, ok := api.(r4.ResearchDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ResearchDefinition"}
		}
		return impl.ReadResearchDefinition(ctx, id)
	case "ResearchElementDefinition":
		impl, ok := api.(r4.ResearchElementDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ResearchElementDefinition"}
		}
		return impl.ReadResearchElementDefinition(ctx, id)
	case "ResearchStudy":
		impl, ok := api.(r4.ResearchStudyRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ResearchStudy"}
		}
		return impl.ReadResearchStudy(ctx, id)
	case "ResearchSubject":
		impl, ok := api.(r4.ResearchSubjectRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ResearchSubject"}
		}
		return impl.ReadResearchSubject(ctx, id)
	case "RiskAssessment":
		impl, ok := api.(r4.RiskAssessmentRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "RiskAssessment"}
		}
		return impl.ReadRiskAssessment(ctx, id)
	case "RiskEvidenceSynthesis":
		impl, ok := api.(r4.RiskEvidenceSynthesisRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.ReadRiskEvidenceSynthesis(ctx, id)
	case "Schedule":
		impl, ok := api.(r4.ScheduleRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Schedule"}
		}
		return impl.ReadSchedule(ctx, id)
	case "SearchParameter":
		impl, ok := api.(r4.SearchParameterRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SearchParameter"}
		}
		return impl.ReadSearchParameter(ctx, id)
	case "ServiceRequest":
		impl, ok := api.(r4.ServiceRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ServiceRequest"}
		}
		return impl.ReadServiceRequest(ctx, id)
	case "Slot":
		impl, ok := api.(r4.SlotRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Slot"}
		}
		return impl.ReadSlot(ctx, id)
	case "Specimen":
		impl, ok := api.(r4.SpecimenRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Specimen"}
		}
		return impl.ReadSpecimen(ctx, id)
	case "SpecimenDefinition":
		impl, ok := api.(r4.SpecimenDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SpecimenDefinition"}
		}
		return impl.ReadSpecimenDefinition(ctx, id)
	case "StructureDefinition":
		impl, ok := api.(r4.StructureDefinitionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "StructureDefinition"}
		}
		return impl.ReadStructureDefinition(ctx, id)
	case "StructureMap":
		impl, ok := api.(r4.StructureMapRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "StructureMap"}
		}
		return impl.ReadStructureMap(ctx, id)
	case "Subscription":
		impl, ok := api.(r4.SubscriptionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Subscription"}
		}
		return impl.ReadSubscription(ctx, id)
	case "Substance":
		impl, ok := api.(r4.SubstanceRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Substance"}
		}
		return impl.ReadSubstance(ctx, id)
	case "SubstanceNucleicAcid":
		impl, ok := api.(r4.SubstanceNucleicAcidRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.ReadSubstanceNucleicAcid(ctx, id)
	case "SubstancePolymer":
		impl, ok := api.(r4.SubstancePolymerRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstancePolymer"}
		}
		return impl.ReadSubstancePolymer(ctx, id)
	case "SubstanceProtein":
		impl, ok := api.(r4.SubstanceProteinRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstanceProtein"}
		}
		return impl.ReadSubstanceProtein(ctx, id)
	case "SubstanceReferenceInformation":
		impl, ok := api.(r4.SubstanceReferenceInformationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.ReadSubstanceReferenceInformation(ctx, id)
	case "SubstanceSourceMaterial":
		impl, ok := api.(r4.SubstanceSourceMaterialRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.ReadSubstanceSourceMaterial(ctx, id)
	case "SubstanceSpecification":
		impl, ok := api.(r4.SubstanceSpecificationRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SubstanceSpecification"}
		}
		return impl.ReadSubstanceSpecification(ctx, id)
	case "SupplyDelivery":
		impl, ok := api.(r4.SupplyDeliveryRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SupplyDelivery"}
		}
		return impl.ReadSupplyDelivery(ctx, id)
	case "SupplyRequest":
		impl, ok := api.(r4.SupplyRequestRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "SupplyRequest"}
		}
		return impl.ReadSupplyRequest(ctx, id)
	case "Task":
		impl, ok := api.(r4.TaskRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "Task"}
		}
		return impl.ReadTask(ctx, id)
	case "TerminologyCapabilities":
		impl, ok := api.(r4.TerminologyCapabilitiesRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "TerminologyCapabilities"}
		}
		return impl.ReadTerminologyCapabilities(ctx, id)
	case "TestReport":
		impl, ok := api.(r4.TestReportRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "TestReport"}
		}
		return impl.ReadTestReport(ctx, id)
	case "TestScript":
		impl, ok := api.(r4.TestScriptRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "TestScript"}
		}
		return impl.ReadTestScript(ctx, id)
	case "ValueSet":
		impl, ok := api.(r4.ValueSetRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "ValueSet"}
		}
		return impl.ReadValueSet(ctx, id)
	case "VerificationResult":
		impl, ok := api.(r4.VerificationResultRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "VerificationResult"}
		}
		return impl.ReadVerificationResult(ctx, id)
	case "VisionPrescription":
		impl, ok := api.(r4.VisionPrescriptionRead)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Read", ResourceType: "VisionPrescription"}
		}
		return impl.ReadVisionPrescription(ctx, id)
	default:
		return nil, errors.UnknownResourceError{ResourceType: resourceType}
	}
}
