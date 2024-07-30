package genericR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/capabilities/gen/r4"
	search "fhir-toolbox/capabilities/search"
)

func (w wrapper) SearchCapabilities(resourceType string) (search.Capabilities, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.api.(r4.AccountSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Account"}
		}
		return impl.SearchCapabilitiesAccount(), nil
	case "ActivityDefinition":
		impl, ok := w.api.(r4.ActivityDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActivityDefinition"}
		}
		return impl.SearchCapabilitiesActivityDefinition(), nil
	case "AdverseEvent":
		impl, ok := w.api.(r4.AdverseEventSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdverseEvent"}
		}
		return impl.SearchCapabilitiesAdverseEvent(), nil
	case "AllergyIntolerance":
		impl, ok := w.api.(r4.AllergyIntoleranceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AllergyIntolerance"}
		}
		return impl.SearchCapabilitiesAllergyIntolerance(), nil
	case "Appointment":
		impl, ok := w.api.(r4.AppointmentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Appointment"}
		}
		return impl.SearchCapabilitiesAppointment(), nil
	case "AppointmentResponse":
		impl, ok := w.api.(r4.AppointmentResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AppointmentResponse"}
		}
		return impl.SearchCapabilitiesAppointmentResponse(), nil
	case "AuditEvent":
		impl, ok := w.api.(r4.AuditEventSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AuditEvent"}
		}
		return impl.SearchCapabilitiesAuditEvent(), nil
	case "Basic":
		impl, ok := w.api.(r4.BasicSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Basic"}
		}
		return impl.SearchCapabilitiesBasic(), nil
	case "Binary":
		impl, ok := w.api.(r4.BinarySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Binary"}
		}
		return impl.SearchCapabilitiesBinary(), nil
	case "BiologicallyDerivedProduct":
		impl, ok := w.api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.SearchCapabilitiesBiologicallyDerivedProduct(), nil
	case "BodyStructure":
		impl, ok := w.api.(r4.BodyStructureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BodyStructure"}
		}
		return impl.SearchCapabilitiesBodyStructure(), nil
	case "Bundle":
		impl, ok := w.api.(r4.BundleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Bundle"}
		}
		return impl.SearchCapabilitiesBundle(), nil
	case "CapabilityStatement":
		impl, ok := w.api.(r4.CapabilityStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CapabilityStatement"}
		}
		return impl.SearchCapabilitiesCapabilityStatement(), nil
	case "CarePlan":
		impl, ok := w.api.(r4.CarePlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CarePlan"}
		}
		return impl.SearchCapabilitiesCarePlan(), nil
	case "CareTeam":
		impl, ok := w.api.(r4.CareTeamSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CareTeam"}
		}
		return impl.SearchCapabilitiesCareTeam(), nil
	case "CatalogEntry":
		impl, ok := w.api.(r4.CatalogEntrySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CatalogEntry"}
		}
		return impl.SearchCapabilitiesCatalogEntry(), nil
	case "ChargeItem":
		impl, ok := w.api.(r4.ChargeItemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItem"}
		}
		return impl.SearchCapabilitiesChargeItem(), nil
	case "ChargeItemDefinition":
		impl, ok := w.api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItemDefinition"}
		}
		return impl.SearchCapabilitiesChargeItemDefinition(), nil
	case "Claim":
		impl, ok := w.api.(r4.ClaimSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Claim"}
		}
		return impl.SearchCapabilitiesClaim(), nil
	case "ClaimResponse":
		impl, ok := w.api.(r4.ClaimResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClaimResponse"}
		}
		return impl.SearchCapabilitiesClaimResponse(), nil
	case "ClinicalImpression":
		impl, ok := w.api.(r4.ClinicalImpressionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalImpression"}
		}
		return impl.SearchCapabilitiesClinicalImpression(), nil
	case "CodeSystem":
		impl, ok := w.api.(r4.CodeSystemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CodeSystem"}
		}
		return impl.SearchCapabilitiesCodeSystem(), nil
	case "Communication":
		impl, ok := w.api.(r4.CommunicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Communication"}
		}
		return impl.SearchCapabilitiesCommunication(), nil
	case "CommunicationRequest":
		impl, ok := w.api.(r4.CommunicationRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CommunicationRequest"}
		}
		return impl.SearchCapabilitiesCommunicationRequest(), nil
	case "CompartmentDefinition":
		impl, ok := w.api.(r4.CompartmentDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CompartmentDefinition"}
		}
		return impl.SearchCapabilitiesCompartmentDefinition(), nil
	case "Composition":
		impl, ok := w.api.(r4.CompositionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Composition"}
		}
		return impl.SearchCapabilitiesComposition(), nil
	case "ConceptMap":
		impl, ok := w.api.(r4.ConceptMapSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConceptMap"}
		}
		return impl.SearchCapabilitiesConceptMap(), nil
	case "Condition":
		impl, ok := w.api.(r4.ConditionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Condition"}
		}
		return impl.SearchCapabilitiesCondition(), nil
	case "Consent":
		impl, ok := w.api.(r4.ConsentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Consent"}
		}
		return impl.SearchCapabilitiesConsent(), nil
	case "Contract":
		impl, ok := w.api.(r4.ContractSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Contract"}
		}
		return impl.SearchCapabilitiesContract(), nil
	case "Coverage":
		impl, ok := w.api.(r4.CoverageSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Coverage"}
		}
		return impl.SearchCapabilitiesCoverage(), nil
	case "CoverageEligibilityRequest":
		impl, ok := w.api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.SearchCapabilitiesCoverageEligibilityRequest(), nil
	case "CoverageEligibilityResponse":
		impl, ok := w.api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.SearchCapabilitiesCoverageEligibilityResponse(), nil
	case "DetectedIssue":
		impl, ok := w.api.(r4.DetectedIssueSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DetectedIssue"}
		}
		return impl.SearchCapabilitiesDetectedIssue(), nil
	case "Device":
		impl, ok := w.api.(r4.DeviceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Device"}
		}
		return impl.SearchCapabilitiesDevice(), nil
	case "DeviceDefinition":
		impl, ok := w.api.(r4.DeviceDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchCapabilitiesDeviceDefinition(), nil
	case "DeviceMetric":
		impl, ok := w.api.(r4.DeviceMetricSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceMetric"}
		}
		return impl.SearchCapabilitiesDeviceMetric(), nil
	case "DeviceRequest":
		impl, ok := w.api.(r4.DeviceRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceRequest"}
		}
		return impl.SearchCapabilitiesDeviceRequest(), nil
	case "DeviceUseStatement":
		impl, ok := w.api.(r4.DeviceUseStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUseStatement"}
		}
		return impl.SearchCapabilitiesDeviceUseStatement(), nil
	case "DiagnosticReport":
		impl, ok := w.api.(r4.DiagnosticReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchCapabilitiesDiagnosticReport(), nil
	case "DocumentManifest":
		impl, ok := w.api.(r4.DocumentManifestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentManifest"}
		}
		return impl.SearchCapabilitiesDocumentManifest(), nil
	case "DocumentReference":
		impl, ok := w.api.(r4.DocumentReferenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchCapabilitiesDocumentReference(), nil
	case "EffectEvidenceSynthesis":
		impl, ok := w.api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.SearchCapabilitiesEffectEvidenceSynthesis(), nil
	case "Encounter":
		impl, ok := w.api.(r4.EncounterSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchCapabilitiesEncounter(), nil
	case "Endpoint":
		impl, ok := w.api.(r4.EndpointSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Endpoint"}
		}
		return impl.SearchCapabilitiesEndpoint(), nil
	case "EnrollmentRequest":
		impl, ok := w.api.(r4.EnrollmentRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentRequest"}
		}
		return impl.SearchCapabilitiesEnrollmentRequest(), nil
	case "EnrollmentResponse":
		impl, ok := w.api.(r4.EnrollmentResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentResponse"}
		}
		return impl.SearchCapabilitiesEnrollmentResponse(), nil
	case "EpisodeOfCare":
		impl, ok := w.api.(r4.EpisodeOfCareSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EpisodeOfCare"}
		}
		return impl.SearchCapabilitiesEpisodeOfCare(), nil
	case "EventDefinition":
		impl, ok := w.api.(r4.EventDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EventDefinition"}
		}
		return impl.SearchCapabilitiesEventDefinition(), nil
	case "Evidence":
		impl, ok := w.api.(r4.EvidenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Evidence"}
		}
		return impl.SearchCapabilitiesEvidence(), nil
	case "EvidenceVariable":
		impl, ok := w.api.(r4.EvidenceVariableSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceVariable"}
		}
		return impl.SearchCapabilitiesEvidenceVariable(), nil
	case "ExampleScenario":
		impl, ok := w.api.(r4.ExampleScenarioSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExampleScenario"}
		}
		return impl.SearchCapabilitiesExampleScenario(), nil
	case "ExplanationOfBenefit":
		impl, ok := w.api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.SearchCapabilitiesExplanationOfBenefit(), nil
	case "FamilyMemberHistory":
		impl, ok := w.api.(r4.FamilyMemberHistorySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FamilyMemberHistory"}
		}
		return impl.SearchCapabilitiesFamilyMemberHistory(), nil
	case "Flag":
		impl, ok := w.api.(r4.FlagSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Flag"}
		}
		return impl.SearchCapabilitiesFlag(), nil
	case "Goal":
		impl, ok := w.api.(r4.GoalSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Goal"}
		}
		return impl.SearchCapabilitiesGoal(), nil
	case "GraphDefinition":
		impl, ok := w.api.(r4.GraphDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GraphDefinition"}
		}
		return impl.SearchCapabilitiesGraphDefinition(), nil
	case "Group":
		impl, ok := w.api.(r4.GroupSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Group"}
		}
		return impl.SearchCapabilitiesGroup(), nil
	case "GuidanceResponse":
		impl, ok := w.api.(r4.GuidanceResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GuidanceResponse"}
		}
		return impl.SearchCapabilitiesGuidanceResponse(), nil
	case "HealthcareService":
		impl, ok := w.api.(r4.HealthcareServiceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "HealthcareService"}
		}
		return impl.SearchCapabilitiesHealthcareService(), nil
	case "ImagingStudy":
		impl, ok := w.api.(r4.ImagingStudySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingStudy"}
		}
		return impl.SearchCapabilitiesImagingStudy(), nil
	case "Immunization":
		impl, ok := w.api.(r4.ImmunizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Immunization"}
		}
		return impl.SearchCapabilitiesImmunization(), nil
	case "ImmunizationEvaluation":
		impl, ok := w.api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.SearchCapabilitiesImmunizationEvaluation(), nil
	case "ImmunizationRecommendation":
		impl, ok := w.api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.SearchCapabilitiesImmunizationRecommendation(), nil
	case "ImplementationGuide":
		impl, ok := w.api.(r4.ImplementationGuideSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImplementationGuide"}
		}
		return impl.SearchCapabilitiesImplementationGuide(), nil
	case "InsurancePlan":
		impl, ok := w.api.(r4.InsurancePlanSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchCapabilitiesInsurancePlan(), nil
	case "Invoice":
		impl, ok := w.api.(r4.InvoiceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Invoice"}
		}
		return impl.SearchCapabilitiesInvoice(), nil
	case "Library":
		impl, ok := w.api.(r4.LibrarySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Library"}
		}
		return impl.SearchCapabilitiesLibrary(), nil
	case "Linkage":
		impl, ok := w.api.(r4.LinkageSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Linkage"}
		}
		return impl.SearchCapabilitiesLinkage(), nil
	case "List":
		impl, ok := w.api.(r4.ListSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "List"}
		}
		return impl.SearchCapabilitiesList(), nil
	case "Location":
		impl, ok := w.api.(r4.LocationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Location"}
		}
		return impl.SearchCapabilitiesLocation(), nil
	case "Measure":
		impl, ok := w.api.(r4.MeasureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Measure"}
		}
		return impl.SearchCapabilitiesMeasure(), nil
	case "MeasureReport":
		impl, ok := w.api.(r4.MeasureReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MeasureReport"}
		}
		return impl.SearchCapabilitiesMeasureReport(), nil
	case "Media":
		impl, ok := w.api.(r4.MediaSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Media"}
		}
		return impl.SearchCapabilitiesMedia(), nil
	case "Medication":
		impl, ok := w.api.(r4.MedicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Medication"}
		}
		return impl.SearchCapabilitiesMedication(), nil
	case "MedicationAdministration":
		impl, ok := w.api.(r4.MedicationAdministrationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationAdministration"}
		}
		return impl.SearchCapabilitiesMedicationAdministration(), nil
	case "MedicationDispense":
		impl, ok := w.api.(r4.MedicationDispenseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationDispense"}
		}
		return impl.SearchCapabilitiesMedicationDispense(), nil
	case "MedicationKnowledge":
		impl, ok := w.api.(r4.MedicationKnowledgeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationKnowledge"}
		}
		return impl.SearchCapabilitiesMedicationKnowledge(), nil
	case "MedicationRequest":
		impl, ok := w.api.(r4.MedicationRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationRequest"}
		}
		return impl.SearchCapabilitiesMedicationRequest(), nil
	case "MedicationStatement":
		impl, ok := w.api.(r4.MedicationStatementSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationStatement"}
		}
		return impl.SearchCapabilitiesMedicationStatement(), nil
	case "MedicinalProduct":
		impl, ok := w.api.(r4.MedicinalProductSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProduct"}
		}
		return impl.SearchCapabilitiesMedicinalProduct(), nil
	case "MedicinalProductAuthorization":
		impl, ok := w.api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.SearchCapabilitiesMedicinalProductAuthorization(), nil
	case "MedicinalProductContraindication":
		impl, ok := w.api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.SearchCapabilitiesMedicinalProductContraindication(), nil
	case "MedicinalProductIndication":
		impl, ok := w.api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIndication"}
		}
		return impl.SearchCapabilitiesMedicinalProductIndication(), nil
	case "MedicinalProductIngredient":
		impl, ok := w.api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.SearchCapabilitiesMedicinalProductIngredient(), nil
	case "MedicinalProductInteraction":
		impl, ok := w.api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.SearchCapabilitiesMedicinalProductInteraction(), nil
	case "MedicinalProductManufactured":
		impl, ok := w.api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.SearchCapabilitiesMedicinalProductManufactured(), nil
	case "MedicinalProductPackaged":
		impl, ok := w.api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.SearchCapabilitiesMedicinalProductPackaged(), nil
	case "MedicinalProductPharmaceutical":
		impl, ok := w.api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.SearchCapabilitiesMedicinalProductPharmaceutical(), nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.SearchCapabilitiesMedicinalProductUndesirableEffect(), nil
	case "MessageDefinition":
		impl, ok := w.api.(r4.MessageDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageDefinition"}
		}
		return impl.SearchCapabilitiesMessageDefinition(), nil
	case "MessageHeader":
		impl, ok := w.api.(r4.MessageHeaderSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageHeader"}
		}
		return impl.SearchCapabilitiesMessageHeader(), nil
	case "MolecularSequence":
		impl, ok := w.api.(r4.MolecularSequenceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MolecularSequence"}
		}
		return impl.SearchCapabilitiesMolecularSequence(), nil
	case "NamingSystem":
		impl, ok := w.api.(r4.NamingSystemSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NamingSystem"}
		}
		return impl.SearchCapabilitiesNamingSystem(), nil
	case "NutritionOrder":
		impl, ok := w.api.(r4.NutritionOrderSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchCapabilitiesNutritionOrder(), nil
	case "Observation":
		impl, ok := w.api.(r4.ObservationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Observation"}
		}
		return impl.SearchCapabilitiesObservation(), nil
	case "ObservationDefinition":
		impl, ok := w.api.(r4.ObservationDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ObservationDefinition"}
		}
		return impl.SearchCapabilitiesObservationDefinition(), nil
	case "OperationDefinition":
		impl, ok := w.api.(r4.OperationDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationDefinition"}
		}
		return impl.SearchCapabilitiesOperationDefinition(), nil
	case "OperationOutcome":
		impl, ok := w.api.(r4.OperationOutcomeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationOutcome"}
		}
		return impl.SearchCapabilitiesOperationOutcome(), nil
	case "Organization":
		impl, ok := w.api.(r4.OrganizationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Organization"}
		}
		return impl.SearchCapabilitiesOrganization(), nil
	case "OrganizationAffiliation":
		impl, ok := w.api.(r4.OrganizationAffiliationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OrganizationAffiliation"}
		}
		return impl.SearchCapabilitiesOrganizationAffiliation(), nil
	case "Parameters":
		impl, ok := w.api.(r4.ParametersSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Parameters"}
		}
		return impl.SearchCapabilitiesParameters(), nil
	case "Patient":
		impl, ok := w.api.(r4.PatientSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Patient"}
		}
		return impl.SearchCapabilitiesPatient(), nil
	case "PaymentNotice":
		impl, ok := w.api.(r4.PaymentNoticeSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentNotice"}
		}
		return impl.SearchCapabilitiesPaymentNotice(), nil
	case "PaymentReconciliation":
		impl, ok := w.api.(r4.PaymentReconciliationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentReconciliation"}
		}
		return impl.SearchCapabilitiesPaymentReconciliation(), nil
	case "Person":
		impl, ok := w.api.(r4.PersonSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Person"}
		}
		return impl.SearchCapabilitiesPerson(), nil
	case "PlanDefinition":
		impl, ok := w.api.(r4.PlanDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PlanDefinition"}
		}
		return impl.SearchCapabilitiesPlanDefinition(), nil
	case "Practitioner":
		impl, ok := w.api.(r4.PractitionerSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Practitioner"}
		}
		return impl.SearchCapabilitiesPractitioner(), nil
	case "PractitionerRole":
		impl, ok := w.api.(r4.PractitionerRoleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PractitionerRole"}
		}
		return impl.SearchCapabilitiesPractitionerRole(), nil
	case "Procedure":
		impl, ok := w.api.(r4.ProcedureSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Procedure"}
		}
		return impl.SearchCapabilitiesProcedure(), nil
	case "Provenance":
		impl, ok := w.api.(r4.ProvenanceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Provenance"}
		}
		return impl.SearchCapabilitiesProvenance(), nil
	case "Questionnaire":
		impl, ok := w.api.(r4.QuestionnaireSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Questionnaire"}
		}
		return impl.SearchCapabilitiesQuestionnaire(), nil
	case "QuestionnaireResponse":
		impl, ok := w.api.(r4.QuestionnaireResponseSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "QuestionnaireResponse"}
		}
		return impl.SearchCapabilitiesQuestionnaireResponse(), nil
	case "RelatedPerson":
		impl, ok := w.api.(r4.RelatedPersonSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchCapabilitiesRelatedPerson(), nil
	case "RequestGroup":
		impl, ok := w.api.(r4.RequestGroupSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestGroup"}
		}
		return impl.SearchCapabilitiesRequestGroup(), nil
	case "ResearchDefinition":
		impl, ok := w.api.(r4.ResearchDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchDefinition"}
		}
		return impl.SearchCapabilitiesResearchDefinition(), nil
	case "ResearchElementDefinition":
		impl, ok := w.api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchElementDefinition"}
		}
		return impl.SearchCapabilitiesResearchElementDefinition(), nil
	case "ResearchStudy":
		impl, ok := w.api.(r4.ResearchStudySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchStudy"}
		}
		return impl.SearchCapabilitiesResearchStudy(), nil
	case "ResearchSubject":
		impl, ok := w.api.(r4.ResearchSubjectSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchSubject"}
		}
		return impl.SearchCapabilitiesResearchSubject(), nil
	case "RiskAssessment":
		impl, ok := w.api.(r4.RiskAssessmentSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskAssessment"}
		}
		return impl.SearchCapabilitiesRiskAssessment(), nil
	case "RiskEvidenceSynthesis":
		impl, ok := w.api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.SearchCapabilitiesRiskEvidenceSynthesis(), nil
	case "Schedule":
		impl, ok := w.api.(r4.ScheduleSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Schedule"}
		}
		return impl.SearchCapabilitiesSchedule(), nil
	case "SearchParameter":
		impl, ok := w.api.(r4.SearchParameterSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SearchParameter"}
		}
		return impl.SearchCapabilitiesSearchParameter(), nil
	case "ServiceRequest":
		impl, ok := w.api.(r4.ServiceRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ServiceRequest"}
		}
		return impl.SearchCapabilitiesServiceRequest(), nil
	case "Slot":
		impl, ok := w.api.(r4.SlotSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Slot"}
		}
		return impl.SearchCapabilitiesSlot(), nil
	case "Specimen":
		impl, ok := w.api.(r4.SpecimenSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Specimen"}
		}
		return impl.SearchCapabilitiesSpecimen(), nil
	case "SpecimenDefinition":
		impl, ok := w.api.(r4.SpecimenDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SpecimenDefinition"}
		}
		return impl.SearchCapabilitiesSpecimenDefinition(), nil
	case "StructureDefinition":
		impl, ok := w.api.(r4.StructureDefinitionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureDefinition"}
		}
		return impl.SearchCapabilitiesStructureDefinition(), nil
	case "StructureMap":
		impl, ok := w.api.(r4.StructureMapSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureMap"}
		}
		return impl.SearchCapabilitiesStructureMap(), nil
	case "Subscription":
		impl, ok := w.api.(r4.SubscriptionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Subscription"}
		}
		return impl.SearchCapabilitiesSubscription(), nil
	case "Substance":
		impl, ok := w.api.(r4.SubstanceSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchCapabilitiesSubstance(), nil
	case "SubstanceNucleicAcid":
		impl, ok := w.api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.SearchCapabilitiesSubstanceNucleicAcid(), nil
	case "SubstancePolymer":
		impl, ok := w.api.(r4.SubstancePolymerSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstancePolymer"}
		}
		return impl.SearchCapabilitiesSubstancePolymer(), nil
	case "SubstanceProtein":
		impl, ok := w.api.(r4.SubstanceProteinSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceProtein"}
		}
		return impl.SearchCapabilitiesSubstanceProtein(), nil
	case "SubstanceReferenceInformation":
		impl, ok := w.api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.SearchCapabilitiesSubstanceReferenceInformation(), nil
	case "SubstanceSourceMaterial":
		impl, ok := w.api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.SearchCapabilitiesSubstanceSourceMaterial(), nil
	case "SubstanceSpecification":
		impl, ok := w.api.(r4.SubstanceSpecificationSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSpecification"}
		}
		return impl.SearchCapabilitiesSubstanceSpecification(), nil
	case "SupplyDelivery":
		impl, ok := w.api.(r4.SupplyDeliverySearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyDelivery"}
		}
		return impl.SearchCapabilitiesSupplyDelivery(), nil
	case "SupplyRequest":
		impl, ok := w.api.(r4.SupplyRequestSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyRequest"}
		}
		return impl.SearchCapabilitiesSupplyRequest(), nil
	case "Task":
		impl, ok := w.api.(r4.TaskSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Task"}
		}
		return impl.SearchCapabilitiesTask(), nil
	case "TerminologyCapabilities":
		impl, ok := w.api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TerminologyCapabilities"}
		}
		return impl.SearchCapabilitiesTerminologyCapabilities(), nil
	case "TestReport":
		impl, ok := w.api.(r4.TestReportSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestReport"}
		}
		return impl.SearchCapabilitiesTestReport(), nil
	case "TestScript":
		impl, ok := w.api.(r4.TestScriptSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestScript"}
		}
		return impl.SearchCapabilitiesTestScript(), nil
	case "ValueSet":
		impl, ok := w.api.(r4.ValueSetSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ValueSet"}
		}
		return impl.SearchCapabilitiesValueSet(), nil
	case "VerificationResult":
		impl, ok := w.api.(r4.VerificationResultSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VerificationResult"}
		}
		return impl.SearchCapabilitiesVerificationResult(), nil
	case "VisionPrescription":
		impl, ok := w.api.(r4.VisionPrescriptionSearch)
		if !ok {
			return search.Capabilities{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VisionPrescription"}
		}
		return impl.SearchCapabilitiesVisionPrescription(), nil
	default:
		return search.Capabilities{}, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
func (w wrapper) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := w.api.(r4.AccountSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Account"}
		}
		return impl.SearchAccount(ctx, options)
	case "ActivityDefinition":
		impl, ok := w.api.(r4.ActivityDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ActivityDefinition"}
		}
		return impl.SearchActivityDefinition(ctx, options)
	case "AdverseEvent":
		impl, ok := w.api.(r4.AdverseEventSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AdverseEvent"}
		}
		return impl.SearchAdverseEvent(ctx, options)
	case "AllergyIntolerance":
		impl, ok := w.api.(r4.AllergyIntoleranceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AllergyIntolerance"}
		}
		return impl.SearchAllergyIntolerance(ctx, options)
	case "Appointment":
		impl, ok := w.api.(r4.AppointmentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Appointment"}
		}
		return impl.SearchAppointment(ctx, options)
	case "AppointmentResponse":
		impl, ok := w.api.(r4.AppointmentResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AppointmentResponse"}
		}
		return impl.SearchAppointmentResponse(ctx, options)
	case "AuditEvent":
		impl, ok := w.api.(r4.AuditEventSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "AuditEvent"}
		}
		return impl.SearchAuditEvent(ctx, options)
	case "Basic":
		impl, ok := w.api.(r4.BasicSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Basic"}
		}
		return impl.SearchBasic(ctx, options)
	case "Binary":
		impl, ok := w.api.(r4.BinarySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Binary"}
		}
		return impl.SearchBinary(ctx, options)
	case "BiologicallyDerivedProduct":
		impl, ok := w.api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BiologicallyDerivedProduct"}
		}
		return impl.SearchBiologicallyDerivedProduct(ctx, options)
	case "BodyStructure":
		impl, ok := w.api.(r4.BodyStructureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "BodyStructure"}
		}
		return impl.SearchBodyStructure(ctx, options)
	case "Bundle":
		impl, ok := w.api.(r4.BundleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Bundle"}
		}
		return impl.SearchBundle(ctx, options)
	case "CapabilityStatement":
		impl, ok := w.api.(r4.CapabilityStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CapabilityStatement"}
		}
		return impl.SearchCapabilityStatement(ctx, options)
	case "CarePlan":
		impl, ok := w.api.(r4.CarePlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CarePlan"}
		}
		return impl.SearchCarePlan(ctx, options)
	case "CareTeam":
		impl, ok := w.api.(r4.CareTeamSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CareTeam"}
		}
		return impl.SearchCareTeam(ctx, options)
	case "CatalogEntry":
		impl, ok := w.api.(r4.CatalogEntrySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CatalogEntry"}
		}
		return impl.SearchCatalogEntry(ctx, options)
	case "ChargeItem":
		impl, ok := w.api.(r4.ChargeItemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItem"}
		}
		return impl.SearchChargeItem(ctx, options)
	case "ChargeItemDefinition":
		impl, ok := w.api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ChargeItemDefinition"}
		}
		return impl.SearchChargeItemDefinition(ctx, options)
	case "Claim":
		impl, ok := w.api.(r4.ClaimSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Claim"}
		}
		return impl.SearchClaim(ctx, options)
	case "ClaimResponse":
		impl, ok := w.api.(r4.ClaimResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClaimResponse"}
		}
		return impl.SearchClaimResponse(ctx, options)
	case "ClinicalImpression":
		impl, ok := w.api.(r4.ClinicalImpressionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ClinicalImpression"}
		}
		return impl.SearchClinicalImpression(ctx, options)
	case "CodeSystem":
		impl, ok := w.api.(r4.CodeSystemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CodeSystem"}
		}
		return impl.SearchCodeSystem(ctx, options)
	case "Communication":
		impl, ok := w.api.(r4.CommunicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Communication"}
		}
		return impl.SearchCommunication(ctx, options)
	case "CommunicationRequest":
		impl, ok := w.api.(r4.CommunicationRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CommunicationRequest"}
		}
		return impl.SearchCommunicationRequest(ctx, options)
	case "CompartmentDefinition":
		impl, ok := w.api.(r4.CompartmentDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CompartmentDefinition"}
		}
		return impl.SearchCompartmentDefinition(ctx, options)
	case "Composition":
		impl, ok := w.api.(r4.CompositionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Composition"}
		}
		return impl.SearchComposition(ctx, options)
	case "ConceptMap":
		impl, ok := w.api.(r4.ConceptMapSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ConceptMap"}
		}
		return impl.SearchConceptMap(ctx, options)
	case "Condition":
		impl, ok := w.api.(r4.ConditionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Condition"}
		}
		return impl.SearchCondition(ctx, options)
	case "Consent":
		impl, ok := w.api.(r4.ConsentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Consent"}
		}
		return impl.SearchConsent(ctx, options)
	case "Contract":
		impl, ok := w.api.(r4.ContractSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Contract"}
		}
		return impl.SearchContract(ctx, options)
	case "Coverage":
		impl, ok := w.api.(r4.CoverageSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Coverage"}
		}
		return impl.SearchCoverage(ctx, options)
	case "CoverageEligibilityRequest":
		impl, ok := w.api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityRequest"}
		}
		return impl.SearchCoverageEligibilityRequest(ctx, options)
	case "CoverageEligibilityResponse":
		impl, ok := w.api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "CoverageEligibilityResponse"}
		}
		return impl.SearchCoverageEligibilityResponse(ctx, options)
	case "DetectedIssue":
		impl, ok := w.api.(r4.DetectedIssueSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DetectedIssue"}
		}
		return impl.SearchDetectedIssue(ctx, options)
	case "Device":
		impl, ok := w.api.(r4.DeviceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Device"}
		}
		return impl.SearchDevice(ctx, options)
	case "DeviceDefinition":
		impl, ok := w.api.(r4.DeviceDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceDefinition"}
		}
		return impl.SearchDeviceDefinition(ctx, options)
	case "DeviceMetric":
		impl, ok := w.api.(r4.DeviceMetricSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceMetric"}
		}
		return impl.SearchDeviceMetric(ctx, options)
	case "DeviceRequest":
		impl, ok := w.api.(r4.DeviceRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceRequest"}
		}
		return impl.SearchDeviceRequest(ctx, options)
	case "DeviceUseStatement":
		impl, ok := w.api.(r4.DeviceUseStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DeviceUseStatement"}
		}
		return impl.SearchDeviceUseStatement(ctx, options)
	case "DiagnosticReport":
		impl, ok := w.api.(r4.DiagnosticReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DiagnosticReport"}
		}
		return impl.SearchDiagnosticReport(ctx, options)
	case "DocumentManifest":
		impl, ok := w.api.(r4.DocumentManifestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentManifest"}
		}
		return impl.SearchDocumentManifest(ctx, options)
	case "DocumentReference":
		impl, ok := w.api.(r4.DocumentReferenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "DocumentReference"}
		}
		return impl.SearchDocumentReference(ctx, options)
	case "EffectEvidenceSynthesis":
		impl, ok := w.api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EffectEvidenceSynthesis"}
		}
		return impl.SearchEffectEvidenceSynthesis(ctx, options)
	case "Encounter":
		impl, ok := w.api.(r4.EncounterSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Encounter"}
		}
		return impl.SearchEncounter(ctx, options)
	case "Endpoint":
		impl, ok := w.api.(r4.EndpointSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Endpoint"}
		}
		return impl.SearchEndpoint(ctx, options)
	case "EnrollmentRequest":
		impl, ok := w.api.(r4.EnrollmentRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentRequest"}
		}
		return impl.SearchEnrollmentRequest(ctx, options)
	case "EnrollmentResponse":
		impl, ok := w.api.(r4.EnrollmentResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EnrollmentResponse"}
		}
		return impl.SearchEnrollmentResponse(ctx, options)
	case "EpisodeOfCare":
		impl, ok := w.api.(r4.EpisodeOfCareSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EpisodeOfCare"}
		}
		return impl.SearchEpisodeOfCare(ctx, options)
	case "EventDefinition":
		impl, ok := w.api.(r4.EventDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EventDefinition"}
		}
		return impl.SearchEventDefinition(ctx, options)
	case "Evidence":
		impl, ok := w.api.(r4.EvidenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Evidence"}
		}
		return impl.SearchEvidence(ctx, options)
	case "EvidenceVariable":
		impl, ok := w.api.(r4.EvidenceVariableSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "EvidenceVariable"}
		}
		return impl.SearchEvidenceVariable(ctx, options)
	case "ExampleScenario":
		impl, ok := w.api.(r4.ExampleScenarioSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExampleScenario"}
		}
		return impl.SearchExampleScenario(ctx, options)
	case "ExplanationOfBenefit":
		impl, ok := w.api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ExplanationOfBenefit"}
		}
		return impl.SearchExplanationOfBenefit(ctx, options)
	case "FamilyMemberHistory":
		impl, ok := w.api.(r4.FamilyMemberHistorySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "FamilyMemberHistory"}
		}
		return impl.SearchFamilyMemberHistory(ctx, options)
	case "Flag":
		impl, ok := w.api.(r4.FlagSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Flag"}
		}
		return impl.SearchFlag(ctx, options)
	case "Goal":
		impl, ok := w.api.(r4.GoalSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Goal"}
		}
		return impl.SearchGoal(ctx, options)
	case "GraphDefinition":
		impl, ok := w.api.(r4.GraphDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GraphDefinition"}
		}
		return impl.SearchGraphDefinition(ctx, options)
	case "Group":
		impl, ok := w.api.(r4.GroupSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Group"}
		}
		return impl.SearchGroup(ctx, options)
	case "GuidanceResponse":
		impl, ok := w.api.(r4.GuidanceResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "GuidanceResponse"}
		}
		return impl.SearchGuidanceResponse(ctx, options)
	case "HealthcareService":
		impl, ok := w.api.(r4.HealthcareServiceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "HealthcareService"}
		}
		return impl.SearchHealthcareService(ctx, options)
	case "ImagingStudy":
		impl, ok := w.api.(r4.ImagingStudySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImagingStudy"}
		}
		return impl.SearchImagingStudy(ctx, options)
	case "Immunization":
		impl, ok := w.api.(r4.ImmunizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Immunization"}
		}
		return impl.SearchImmunization(ctx, options)
	case "ImmunizationEvaluation":
		impl, ok := w.api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationEvaluation"}
		}
		return impl.SearchImmunizationEvaluation(ctx, options)
	case "ImmunizationRecommendation":
		impl, ok := w.api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImmunizationRecommendation"}
		}
		return impl.SearchImmunizationRecommendation(ctx, options)
	case "ImplementationGuide":
		impl, ok := w.api.(r4.ImplementationGuideSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ImplementationGuide"}
		}
		return impl.SearchImplementationGuide(ctx, options)
	case "InsurancePlan":
		impl, ok := w.api.(r4.InsurancePlanSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "InsurancePlan"}
		}
		return impl.SearchInsurancePlan(ctx, options)
	case "Invoice":
		impl, ok := w.api.(r4.InvoiceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Invoice"}
		}
		return impl.SearchInvoice(ctx, options)
	case "Library":
		impl, ok := w.api.(r4.LibrarySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Library"}
		}
		return impl.SearchLibrary(ctx, options)
	case "Linkage":
		impl, ok := w.api.(r4.LinkageSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Linkage"}
		}
		return impl.SearchLinkage(ctx, options)
	case "List":
		impl, ok := w.api.(r4.ListSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "List"}
		}
		return impl.SearchList(ctx, options)
	case "Location":
		impl, ok := w.api.(r4.LocationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Location"}
		}
		return impl.SearchLocation(ctx, options)
	case "Measure":
		impl, ok := w.api.(r4.MeasureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Measure"}
		}
		return impl.SearchMeasure(ctx, options)
	case "MeasureReport":
		impl, ok := w.api.(r4.MeasureReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MeasureReport"}
		}
		return impl.SearchMeasureReport(ctx, options)
	case "Media":
		impl, ok := w.api.(r4.MediaSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Media"}
		}
		return impl.SearchMedia(ctx, options)
	case "Medication":
		impl, ok := w.api.(r4.MedicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Medication"}
		}
		return impl.SearchMedication(ctx, options)
	case "MedicationAdministration":
		impl, ok := w.api.(r4.MedicationAdministrationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationAdministration"}
		}
		return impl.SearchMedicationAdministration(ctx, options)
	case "MedicationDispense":
		impl, ok := w.api.(r4.MedicationDispenseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationDispense"}
		}
		return impl.SearchMedicationDispense(ctx, options)
	case "MedicationKnowledge":
		impl, ok := w.api.(r4.MedicationKnowledgeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationKnowledge"}
		}
		return impl.SearchMedicationKnowledge(ctx, options)
	case "MedicationRequest":
		impl, ok := w.api.(r4.MedicationRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationRequest"}
		}
		return impl.SearchMedicationRequest(ctx, options)
	case "MedicationStatement":
		impl, ok := w.api.(r4.MedicationStatementSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicationStatement"}
		}
		return impl.SearchMedicationStatement(ctx, options)
	case "MedicinalProduct":
		impl, ok := w.api.(r4.MedicinalProductSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProduct"}
		}
		return impl.SearchMedicinalProduct(ctx, options)
	case "MedicinalProductAuthorization":
		impl, ok := w.api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductAuthorization"}
		}
		return impl.SearchMedicinalProductAuthorization(ctx, options)
	case "MedicinalProductContraindication":
		impl, ok := w.api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductContraindication"}
		}
		return impl.SearchMedicinalProductContraindication(ctx, options)
	case "MedicinalProductIndication":
		impl, ok := w.api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIndication"}
		}
		return impl.SearchMedicinalProductIndication(ctx, options)
	case "MedicinalProductIngredient":
		impl, ok := w.api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductIngredient"}
		}
		return impl.SearchMedicinalProductIngredient(ctx, options)
	case "MedicinalProductInteraction":
		impl, ok := w.api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductInteraction"}
		}
		return impl.SearchMedicinalProductInteraction(ctx, options)
	case "MedicinalProductManufactured":
		impl, ok := w.api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductManufactured"}
		}
		return impl.SearchMedicinalProductManufactured(ctx, options)
	case "MedicinalProductPackaged":
		impl, ok := w.api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPackaged"}
		}
		return impl.SearchMedicinalProductPackaged(ctx, options)
	case "MedicinalProductPharmaceutical":
		impl, ok := w.api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		return impl.SearchMedicinalProductPharmaceutical(ctx, options)
	case "MedicinalProductUndesirableEffect":
		impl, ok := w.api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		return impl.SearchMedicinalProductUndesirableEffect(ctx, options)
	case "MessageDefinition":
		impl, ok := w.api.(r4.MessageDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageDefinition"}
		}
		return impl.SearchMessageDefinition(ctx, options)
	case "MessageHeader":
		impl, ok := w.api.(r4.MessageHeaderSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MessageHeader"}
		}
		return impl.SearchMessageHeader(ctx, options)
	case "MolecularSequence":
		impl, ok := w.api.(r4.MolecularSequenceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "MolecularSequence"}
		}
		return impl.SearchMolecularSequence(ctx, options)
	case "NamingSystem":
		impl, ok := w.api.(r4.NamingSystemSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NamingSystem"}
		}
		return impl.SearchNamingSystem(ctx, options)
	case "NutritionOrder":
		impl, ok := w.api.(r4.NutritionOrderSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "NutritionOrder"}
		}
		return impl.SearchNutritionOrder(ctx, options)
	case "Observation":
		impl, ok := w.api.(r4.ObservationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Observation"}
		}
		return impl.SearchObservation(ctx, options)
	case "ObservationDefinition":
		impl, ok := w.api.(r4.ObservationDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ObservationDefinition"}
		}
		return impl.SearchObservationDefinition(ctx, options)
	case "OperationDefinition":
		impl, ok := w.api.(r4.OperationDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationDefinition"}
		}
		return impl.SearchOperationDefinition(ctx, options)
	case "OperationOutcome":
		impl, ok := w.api.(r4.OperationOutcomeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OperationOutcome"}
		}
		return impl.SearchOperationOutcome(ctx, options)
	case "Organization":
		impl, ok := w.api.(r4.OrganizationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Organization"}
		}
		return impl.SearchOrganization(ctx, options)
	case "OrganizationAffiliation":
		impl, ok := w.api.(r4.OrganizationAffiliationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "OrganizationAffiliation"}
		}
		return impl.SearchOrganizationAffiliation(ctx, options)
	case "Parameters":
		impl, ok := w.api.(r4.ParametersSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Parameters"}
		}
		return impl.SearchParameters(ctx, options)
	case "Patient":
		impl, ok := w.api.(r4.PatientSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Patient"}
		}
		return impl.SearchPatient(ctx, options)
	case "PaymentNotice":
		impl, ok := w.api.(r4.PaymentNoticeSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentNotice"}
		}
		return impl.SearchPaymentNotice(ctx, options)
	case "PaymentReconciliation":
		impl, ok := w.api.(r4.PaymentReconciliationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PaymentReconciliation"}
		}
		return impl.SearchPaymentReconciliation(ctx, options)
	case "Person":
		impl, ok := w.api.(r4.PersonSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Person"}
		}
		return impl.SearchPerson(ctx, options)
	case "PlanDefinition":
		impl, ok := w.api.(r4.PlanDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PlanDefinition"}
		}
		return impl.SearchPlanDefinition(ctx, options)
	case "Practitioner":
		impl, ok := w.api.(r4.PractitionerSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Practitioner"}
		}
		return impl.SearchPractitioner(ctx, options)
	case "PractitionerRole":
		impl, ok := w.api.(r4.PractitionerRoleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "PractitionerRole"}
		}
		return impl.SearchPractitionerRole(ctx, options)
	case "Procedure":
		impl, ok := w.api.(r4.ProcedureSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Procedure"}
		}
		return impl.SearchProcedure(ctx, options)
	case "Provenance":
		impl, ok := w.api.(r4.ProvenanceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Provenance"}
		}
		return impl.SearchProvenance(ctx, options)
	case "Questionnaire":
		impl, ok := w.api.(r4.QuestionnaireSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Questionnaire"}
		}
		return impl.SearchQuestionnaire(ctx, options)
	case "QuestionnaireResponse":
		impl, ok := w.api.(r4.QuestionnaireResponseSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "QuestionnaireResponse"}
		}
		return impl.SearchQuestionnaireResponse(ctx, options)
	case "RelatedPerson":
		impl, ok := w.api.(r4.RelatedPersonSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RelatedPerson"}
		}
		return impl.SearchRelatedPerson(ctx, options)
	case "RequestGroup":
		impl, ok := w.api.(r4.RequestGroupSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RequestGroup"}
		}
		return impl.SearchRequestGroup(ctx, options)
	case "ResearchDefinition":
		impl, ok := w.api.(r4.ResearchDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchDefinition"}
		}
		return impl.SearchResearchDefinition(ctx, options)
	case "ResearchElementDefinition":
		impl, ok := w.api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchElementDefinition"}
		}
		return impl.SearchResearchElementDefinition(ctx, options)
	case "ResearchStudy":
		impl, ok := w.api.(r4.ResearchStudySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchStudy"}
		}
		return impl.SearchResearchStudy(ctx, options)
	case "ResearchSubject":
		impl, ok := w.api.(r4.ResearchSubjectSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ResearchSubject"}
		}
		return impl.SearchResearchSubject(ctx, options)
	case "RiskAssessment":
		impl, ok := w.api.(r4.RiskAssessmentSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskAssessment"}
		}
		return impl.SearchRiskAssessment(ctx, options)
	case "RiskEvidenceSynthesis":
		impl, ok := w.api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "RiskEvidenceSynthesis"}
		}
		return impl.SearchRiskEvidenceSynthesis(ctx, options)
	case "Schedule":
		impl, ok := w.api.(r4.ScheduleSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Schedule"}
		}
		return impl.SearchSchedule(ctx, options)
	case "SearchParameter":
		impl, ok := w.api.(r4.SearchParameterSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SearchParameter"}
		}
		return impl.SearchSearchParameter(ctx, options)
	case "ServiceRequest":
		impl, ok := w.api.(r4.ServiceRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ServiceRequest"}
		}
		return impl.SearchServiceRequest(ctx, options)
	case "Slot":
		impl, ok := w.api.(r4.SlotSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Slot"}
		}
		return impl.SearchSlot(ctx, options)
	case "Specimen":
		impl, ok := w.api.(r4.SpecimenSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Specimen"}
		}
		return impl.SearchSpecimen(ctx, options)
	case "SpecimenDefinition":
		impl, ok := w.api.(r4.SpecimenDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SpecimenDefinition"}
		}
		return impl.SearchSpecimenDefinition(ctx, options)
	case "StructureDefinition":
		impl, ok := w.api.(r4.StructureDefinitionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureDefinition"}
		}
		return impl.SearchStructureDefinition(ctx, options)
	case "StructureMap":
		impl, ok := w.api.(r4.StructureMapSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "StructureMap"}
		}
		return impl.SearchStructureMap(ctx, options)
	case "Subscription":
		impl, ok := w.api.(r4.SubscriptionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Subscription"}
		}
		return impl.SearchSubscription(ctx, options)
	case "Substance":
		impl, ok := w.api.(r4.SubstanceSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Substance"}
		}
		return impl.SearchSubstance(ctx, options)
	case "SubstanceNucleicAcid":
		impl, ok := w.api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceNucleicAcid"}
		}
		return impl.SearchSubstanceNucleicAcid(ctx, options)
	case "SubstancePolymer":
		impl, ok := w.api.(r4.SubstancePolymerSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstancePolymer"}
		}
		return impl.SearchSubstancePolymer(ctx, options)
	case "SubstanceProtein":
		impl, ok := w.api.(r4.SubstanceProteinSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceProtein"}
		}
		return impl.SearchSubstanceProtein(ctx, options)
	case "SubstanceReferenceInformation":
		impl, ok := w.api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceReferenceInformation"}
		}
		return impl.SearchSubstanceReferenceInformation(ctx, options)
	case "SubstanceSourceMaterial":
		impl, ok := w.api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSourceMaterial"}
		}
		return impl.SearchSubstanceSourceMaterial(ctx, options)
	case "SubstanceSpecification":
		impl, ok := w.api.(r4.SubstanceSpecificationSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SubstanceSpecification"}
		}
		return impl.SearchSubstanceSpecification(ctx, options)
	case "SupplyDelivery":
		impl, ok := w.api.(r4.SupplyDeliverySearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyDelivery"}
		}
		return impl.SearchSupplyDelivery(ctx, options)
	case "SupplyRequest":
		impl, ok := w.api.(r4.SupplyRequestSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "SupplyRequest"}
		}
		return impl.SearchSupplyRequest(ctx, options)
	case "Task":
		impl, ok := w.api.(r4.TaskSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "Task"}
		}
		return impl.SearchTask(ctx, options)
	case "TerminologyCapabilities":
		impl, ok := w.api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TerminologyCapabilities"}
		}
		return impl.SearchTerminologyCapabilities(ctx, options)
	case "TestReport":
		impl, ok := w.api.(r4.TestReportSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestReport"}
		}
		return impl.SearchTestReport(ctx, options)
	case "TestScript":
		impl, ok := w.api.(r4.TestScriptSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "TestScript"}
		}
		return impl.SearchTestScript(ctx, options)
	case "ValueSet":
		impl, ok := w.api.(r4.ValueSetSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "ValueSet"}
		}
		return impl.SearchValueSet(ctx, options)
	case "VerificationResult":
		impl, ok := w.api.(r4.VerificationResultSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VerificationResult"}
		}
		return impl.SearchVerificationResult(ctx, options)
	case "VisionPrescription":
		impl, ok := w.api.(r4.VisionPrescriptionSearch)
		if !ok {
			return search.Result{}, capabilities.NotImplementedError{Interaction: "search", ResourceType: "VisionPrescription"}
		}
		return impl.SearchVisionPrescription(ctx, options)
	default:
		return search.Result{}, capabilities.UnknownResourceError{ResourceType: resourceType}
	}
}
