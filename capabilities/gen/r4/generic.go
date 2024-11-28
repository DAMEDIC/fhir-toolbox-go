package capabilitiesR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	search "fhir-toolbox/capabilities/search"
	model "fhir-toolbox/model"
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
	case "CatalogEntry":
		impl, ok := w.Concrete.(CatalogEntryRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "CatalogEntry"}
		}
		return impl.ReadCatalogEntry(ctx, id)
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
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceDefinition"}
		}
		return impl.ReadDeviceDefinition(ctx, id)
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
	case "DeviceUseStatement":
		impl, ok := w.Concrete.(DeviceUseStatementRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DeviceUseStatement"}
		}
		return impl.ReadDeviceUseStatement(ctx, id)
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DiagnosticReport"}
		}
		return impl.ReadDiagnosticReport(ctx, id)
	case "DocumentManifest":
		impl, ok := w.Concrete.(DocumentManifestRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DocumentManifest"}
		}
		return impl.ReadDocumentManifest(ctx, id)
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "DocumentReference"}
		}
		return impl.ReadDocumentReference(ctx, id)
	case "EffectEvidenceSynthesis":
		impl, ok := w.Concrete.(EffectEvidenceSynthesisRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.ReadEffectEvidenceSynthesis(ctx, id)
	case "Encounter":
		impl, ok := w.Concrete.(EncounterRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Encounter"}
		}
		return impl.ReadEncounter(ctx, id)
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
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "InsurancePlan"}
		}
		return impl.ReadInsurancePlan(ctx, id)
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
	case "Media":
		impl, ok := w.Concrete.(MediaRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Media"}
		}
		return impl.ReadMedia(ctx, id)
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
	case "MedicinalProduct":
		impl, ok := w.Concrete.(MedicinalProductRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProduct"}
		}
		return impl.ReadMedicinalProduct(ctx, id)
	case "MedicinalProductAuthorization":
		impl, ok := w.Concrete.(MedicinalProductAuthorizationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.ReadMedicinalProductAuthorization(ctx, id)
	case "MedicinalProductContraindication":
		impl, ok := w.Concrete.(MedicinalProductContraindicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.ReadMedicinalProductContraindication(ctx, id)
	case "MedicinalProductIndication":
		impl, ok := w.Concrete.(MedicinalProductIndicationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductIndication"}
		}
		return impl.ReadMedicinalProductIndication(ctx, id)
	case "MedicinalProductIngredient":
		impl, ok := w.Concrete.(MedicinalProductIngredientRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.ReadMedicinalProductIngredient(ctx, id)
	case "MedicinalProductInteraction":
		impl, ok := w.Concrete.(MedicinalProductInteractionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.ReadMedicinalProductInteraction(ctx, id)
	case "MedicinalProductManufactured":
		impl, ok := w.Concrete.(MedicinalProductManufacturedRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.ReadMedicinalProductManufactured(ctx, id)
	case "MedicinalProductPackaged":
		impl, ok := w.Concrete.(MedicinalProductPackagedRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.ReadMedicinalProductPackaged(ctx, id)
	case "MedicinalProductPharmaceutical":
		impl, ok := w.Concrete.(MedicinalProductPharmaceuticalRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.ReadMedicinalProductPharmaceutical(ctx, id)
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.Concrete.(MedicinalProductUndesirableEffectRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.ReadMedicinalProductUndesirableEffect(ctx, id)
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
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "NutritionOrder"}
		}
		return impl.ReadNutritionOrder(ctx, id)
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
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RelatedPerson"}
		}
		return impl.ReadRelatedPerson(ctx, id)
	case "RequestGroup":
		impl, ok := w.Concrete.(RequestGroupRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RequestGroup"}
		}
		return impl.ReadRequestGroup(ctx, id)
	case "ResearchDefinition":
		impl, ok := w.Concrete.(ResearchDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchDefinition"}
		}
		return impl.ReadResearchDefinition(ctx, id)
	case "ResearchElementDefinition":
		impl, ok := w.Concrete.(ResearchElementDefinitionRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "ResearchElementDefinition"}
		}
		return impl.ReadResearchElementDefinition(ctx, id)
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
	case "RiskEvidenceSynthesis":
		impl, ok := w.Concrete.(RiskEvidenceSynthesisRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.ReadRiskEvidenceSynthesis(ctx, id)
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
	case "Substance":
		impl, ok := w.Concrete.(SubstanceRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "Substance"}
		}
		return impl.ReadSubstance(ctx, id)
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
	case "SubstanceSpecification":
		impl, ok := w.Concrete.(SubstanceSpecificationRead)
		if !ok {
			return nil, capabilities.NotImplementedError{Interaction: "read", ResourceType: "SubstanceSpecification"}
		}
		return impl.ReadSubstanceSpecification(ctx, id)
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
	case "CatalogEntry":
		impl, ok := w.Concrete.(CatalogEntrySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CatalogEntry"}
		}
		return impl.SearchCapabilitiesCatalogEntry(), nil
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
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchCapabilitiesDeviceDefinition(), nil
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
	case "DeviceUseStatement":
		impl, ok := w.Concrete.(DeviceUseStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUseStatement"}
		}
		return impl.SearchCapabilitiesDeviceUseStatement(), nil
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchCapabilitiesDiagnosticReport(), nil
	case "DocumentManifest":
		impl, ok := w.Concrete.(DocumentManifestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentManifest"}
		}
		return impl.SearchCapabilitiesDocumentManifest(), nil
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchCapabilitiesDocumentReference(), nil
	case "EffectEvidenceSynthesis":
		impl, ok := w.Concrete.(EffectEvidenceSynthesisSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.SearchCapabilitiesEffectEvidenceSynthesis(), nil
	case "Encounter":
		impl, ok := w.Concrete.(EncounterSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchCapabilitiesEncounter(), nil
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
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchCapabilitiesInsurancePlan(), nil
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
	case "Media":
		impl, ok := w.Concrete.(MediaSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Media"}
		}
		return impl.SearchCapabilitiesMedia(), nil
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
	case "MedicinalProduct":
		impl, ok := w.Concrete.(MedicinalProductSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProduct"}
		}
		return impl.SearchCapabilitiesMedicinalProduct(), nil
	case "MedicinalProductAuthorization":
		impl, ok := w.Concrete.(MedicinalProductAuthorizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.SearchCapabilitiesMedicinalProductAuthorization(), nil
	case "MedicinalProductContraindication":
		impl, ok := w.Concrete.(MedicinalProductContraindicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.SearchCapabilitiesMedicinalProductContraindication(), nil
	case "MedicinalProductIndication":
		impl, ok := w.Concrete.(MedicinalProductIndicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIndication"}
		}
		return impl.SearchCapabilitiesMedicinalProductIndication(), nil
	case "MedicinalProductIngredient":
		impl, ok := w.Concrete.(MedicinalProductIngredientSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.SearchCapabilitiesMedicinalProductIngredient(), nil
	case "MedicinalProductInteraction":
		impl, ok := w.Concrete.(MedicinalProductInteractionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.SearchCapabilitiesMedicinalProductInteraction(), nil
	case "MedicinalProductManufactured":
		impl, ok := w.Concrete.(MedicinalProductManufacturedSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.SearchCapabilitiesMedicinalProductManufactured(), nil
	case "MedicinalProductPackaged":
		impl, ok := w.Concrete.(MedicinalProductPackagedSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.SearchCapabilitiesMedicinalProductPackaged(), nil
	case "MedicinalProductPharmaceutical":
		impl, ok := w.Concrete.(MedicinalProductPharmaceuticalSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.SearchCapabilitiesMedicinalProductPharmaceutical(), nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.Concrete.(MedicinalProductUndesirableEffectSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.SearchCapabilitiesMedicinalProductUndesirableEffect(), nil
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
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchCapabilitiesNutritionOrder(), nil
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
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchCapabilitiesRelatedPerson(), nil
	case "RequestGroup":
		impl, ok := w.Concrete.(RequestGroupSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestGroup"}
		}
		return impl.SearchCapabilitiesRequestGroup(), nil
	case "ResearchDefinition":
		impl, ok := w.Concrete.(ResearchDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchDefinition"}
		}
		return impl.SearchCapabilitiesResearchDefinition(), nil
	case "ResearchElementDefinition":
		impl, ok := w.Concrete.(ResearchElementDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchElementDefinition"}
		}
		return impl.SearchCapabilitiesResearchElementDefinition(), nil
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
	case "RiskEvidenceSynthesis":
		impl, ok := w.Concrete.(RiskEvidenceSynthesisSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.SearchCapabilitiesRiskEvidenceSynthesis(), nil
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
	case "Substance":
		impl, ok := w.Concrete.(SubstanceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchCapabilitiesSubstance(), nil
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
	case "SubstanceSpecification":
		impl, ok := w.Concrete.(SubstanceSpecificationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSpecification"}
		}
		return impl.SearchCapabilitiesSubstanceSpecification(), nil
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
	case "CatalogEntry":
		impl, ok := w.Concrete.(CatalogEntrySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CatalogEntry"}
		}
		return impl.SearchCatalogEntry(ctx, options)
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
	case "DeviceDefinition":
		impl, ok := w.Concrete.(DeviceDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchDeviceDefinition(ctx, options)
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
	case "DeviceUseStatement":
		impl, ok := w.Concrete.(DeviceUseStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUseStatement"}
		}
		return impl.SearchDeviceUseStatement(ctx, options)
	case "DiagnosticReport":
		impl, ok := w.Concrete.(DiagnosticReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchDiagnosticReport(ctx, options)
	case "DocumentManifest":
		impl, ok := w.Concrete.(DocumentManifestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentManifest"}
		}
		return impl.SearchDocumentManifest(ctx, options)
	case "DocumentReference":
		impl, ok := w.Concrete.(DocumentReferenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchDocumentReference(ctx, options)
	case "EffectEvidenceSynthesis":
		impl, ok := w.Concrete.(EffectEvidenceSynthesisSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.SearchEffectEvidenceSynthesis(ctx, options)
	case "Encounter":
		impl, ok := w.Concrete.(EncounterSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchEncounter(ctx, options)
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
	case "InsurancePlan":
		impl, ok := w.Concrete.(InsurancePlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchInsurancePlan(ctx, options)
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
	case "Media":
		impl, ok := w.Concrete.(MediaSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Media"}
		}
		return impl.SearchMedia(ctx, options)
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
	case "MedicinalProduct":
		impl, ok := w.Concrete.(MedicinalProductSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProduct"}
		}
		return impl.SearchMedicinalProduct(ctx, options)
	case "MedicinalProductAuthorization":
		impl, ok := w.Concrete.(MedicinalProductAuthorizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.SearchMedicinalProductAuthorization(ctx, options)
	case "MedicinalProductContraindication":
		impl, ok := w.Concrete.(MedicinalProductContraindicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.SearchMedicinalProductContraindication(ctx, options)
	case "MedicinalProductIndication":
		impl, ok := w.Concrete.(MedicinalProductIndicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIndication"}
		}
		return impl.SearchMedicinalProductIndication(ctx, options)
	case "MedicinalProductIngredient":
		impl, ok := w.Concrete.(MedicinalProductIngredientSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.SearchMedicinalProductIngredient(ctx, options)
	case "MedicinalProductInteraction":
		impl, ok := w.Concrete.(MedicinalProductInteractionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.SearchMedicinalProductInteraction(ctx, options)
	case "MedicinalProductManufactured":
		impl, ok := w.Concrete.(MedicinalProductManufacturedSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.SearchMedicinalProductManufactured(ctx, options)
	case "MedicinalProductPackaged":
		impl, ok := w.Concrete.(MedicinalProductPackagedSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.SearchMedicinalProductPackaged(ctx, options)
	case "MedicinalProductPharmaceutical":
		impl, ok := w.Concrete.(MedicinalProductPharmaceuticalSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.SearchMedicinalProductPharmaceutical(ctx, options)
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.Concrete.(MedicinalProductUndesirableEffectSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.SearchMedicinalProductUndesirableEffect(ctx, options)
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
	case "NutritionOrder":
		impl, ok := w.Concrete.(NutritionOrderSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchNutritionOrder(ctx, options)
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
	case "RelatedPerson":
		impl, ok := w.Concrete.(RelatedPersonSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchRelatedPerson(ctx, options)
	case "RequestGroup":
		impl, ok := w.Concrete.(RequestGroupSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestGroup"}
		}
		return impl.SearchRequestGroup(ctx, options)
	case "ResearchDefinition":
		impl, ok := w.Concrete.(ResearchDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchDefinition"}
		}
		return impl.SearchResearchDefinition(ctx, options)
	case "ResearchElementDefinition":
		impl, ok := w.Concrete.(ResearchElementDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchElementDefinition"}
		}
		return impl.SearchResearchElementDefinition(ctx, options)
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
	case "RiskEvidenceSynthesis":
		impl, ok := w.Concrete.(RiskEvidenceSynthesisSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.SearchRiskEvidenceSynthesis(ctx, options)
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
	case "Substance":
		impl, ok := w.Concrete.(SubstanceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchSubstance(ctx, options)
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
	case "SubstanceSpecification":
		impl, ok := w.Concrete.(SubstanceSpecificationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSpecification"}
		}
		return impl.SearchSubstanceSpecification(ctx, options)
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
