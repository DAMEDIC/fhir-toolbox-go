package dispatchR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/capabilities/gen/r4"
	errors "fhir-toolbox/dispatch/errors"
	model "fhir-toolbox/model"
)

func SearchCapabilities(api any, resourceType string) (capabilities.SearchCapabilities, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := api.(r4.AccountSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Account"}
			}
		}
		return impl.SearchCapabilitiesAccount(), nil
	case "ActivityDefinition":
		impl, ok := api.(r4.ActivityDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ActivityDefinition"}
			}
		}
		return impl.SearchCapabilitiesActivityDefinition(), nil
	case "AdverseEvent":
		impl, ok := api.(r4.AdverseEventSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "AdverseEvent"}
			}
		}
		return impl.SearchCapabilitiesAdverseEvent(), nil
	case "AllergyIntolerance":
		impl, ok := api.(r4.AllergyIntoleranceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "AllergyIntolerance"}
			}
		}
		return impl.SearchCapabilitiesAllergyIntolerance(), nil
	case "Appointment":
		impl, ok := api.(r4.AppointmentSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Appointment"}
			}
		}
		return impl.SearchCapabilitiesAppointment(), nil
	case "AppointmentResponse":
		impl, ok := api.(r4.AppointmentResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "AppointmentResponse"}
			}
		}
		return impl.SearchCapabilitiesAppointmentResponse(), nil
	case "AuditEvent":
		impl, ok := api.(r4.AuditEventSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "AuditEvent"}
			}
		}
		return impl.SearchCapabilitiesAuditEvent(), nil
	case "Basic":
		impl, ok := api.(r4.BasicSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Basic"}
			}
		}
		return impl.SearchCapabilitiesBasic(), nil
	case "Binary":
		impl, ok := api.(r4.BinarySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Binary"}
			}
		}
		return impl.SearchCapabilitiesBinary(), nil
	case "BiologicallyDerivedProduct":
		impl, ok := api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "BiologicallyDerivedProduct"}
			}
		}
		return impl.SearchCapabilitiesBiologicallyDerivedProduct(), nil
	case "BodyStructure":
		impl, ok := api.(r4.BodyStructureSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "BodyStructure"}
			}
		}
		return impl.SearchCapabilitiesBodyStructure(), nil
	case "Bundle":
		impl, ok := api.(r4.BundleSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Bundle"}
			}
		}
		return impl.SearchCapabilitiesBundle(), nil
	case "CapabilityStatement":
		impl, ok := api.(r4.CapabilityStatementSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CapabilityStatement"}
			}
		}
		return impl.SearchCapabilitiesCapabilityStatement(), nil
	case "CarePlan":
		impl, ok := api.(r4.CarePlanSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CarePlan"}
			}
		}
		return impl.SearchCapabilitiesCarePlan(), nil
	case "CareTeam":
		impl, ok := api.(r4.CareTeamSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CareTeam"}
			}
		}
		return impl.SearchCapabilitiesCareTeam(), nil
	case "CatalogEntry":
		impl, ok := api.(r4.CatalogEntrySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CatalogEntry"}
			}
		}
		return impl.SearchCapabilitiesCatalogEntry(), nil
	case "ChargeItem":
		impl, ok := api.(r4.ChargeItemSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ChargeItem"}
			}
		}
		return impl.SearchCapabilitiesChargeItem(), nil
	case "ChargeItemDefinition":
		impl, ok := api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ChargeItemDefinition"}
			}
		}
		return impl.SearchCapabilitiesChargeItemDefinition(), nil
	case "Claim":
		impl, ok := api.(r4.ClaimSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Claim"}
			}
		}
		return impl.SearchCapabilitiesClaim(), nil
	case "ClaimResponse":
		impl, ok := api.(r4.ClaimResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ClaimResponse"}
			}
		}
		return impl.SearchCapabilitiesClaimResponse(), nil
	case "ClinicalImpression":
		impl, ok := api.(r4.ClinicalImpressionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ClinicalImpression"}
			}
		}
		return impl.SearchCapabilitiesClinicalImpression(), nil
	case "CodeSystem":
		impl, ok := api.(r4.CodeSystemSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CodeSystem"}
			}
		}
		return impl.SearchCapabilitiesCodeSystem(), nil
	case "Communication":
		impl, ok := api.(r4.CommunicationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Communication"}
			}
		}
		return impl.SearchCapabilitiesCommunication(), nil
	case "CommunicationRequest":
		impl, ok := api.(r4.CommunicationRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CommunicationRequest"}
			}
		}
		return impl.SearchCapabilitiesCommunicationRequest(), nil
	case "CompartmentDefinition":
		impl, ok := api.(r4.CompartmentDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CompartmentDefinition"}
			}
		}
		return impl.SearchCapabilitiesCompartmentDefinition(), nil
	case "Composition":
		impl, ok := api.(r4.CompositionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Composition"}
			}
		}
		return impl.SearchCapabilitiesComposition(), nil
	case "ConceptMap":
		impl, ok := api.(r4.ConceptMapSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ConceptMap"}
			}
		}
		return impl.SearchCapabilitiesConceptMap(), nil
	case "Condition":
		impl, ok := api.(r4.ConditionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Condition"}
			}
		}
		return impl.SearchCapabilitiesCondition(), nil
	case "Consent":
		impl, ok := api.(r4.ConsentSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Consent"}
			}
		}
		return impl.SearchCapabilitiesConsent(), nil
	case "Contract":
		impl, ok := api.(r4.ContractSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Contract"}
			}
		}
		return impl.SearchCapabilitiesContract(), nil
	case "Coverage":
		impl, ok := api.(r4.CoverageSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Coverage"}
			}
		}
		return impl.SearchCapabilitiesCoverage(), nil
	case "CoverageEligibilityRequest":
		impl, ok := api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CoverageEligibilityRequest"}
			}
		}
		return impl.SearchCapabilitiesCoverageEligibilityRequest(), nil
	case "CoverageEligibilityResponse":
		impl, ok := api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "CoverageEligibilityResponse"}
			}
		}
		return impl.SearchCapabilitiesCoverageEligibilityResponse(), nil
	case "DetectedIssue":
		impl, ok := api.(r4.DetectedIssueSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DetectedIssue"}
			}
		}
		return impl.SearchCapabilitiesDetectedIssue(), nil
	case "Device":
		impl, ok := api.(r4.DeviceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Device"}
			}
		}
		return impl.SearchCapabilitiesDevice(), nil
	case "DeviceDefinition":
		impl, ok := api.(r4.DeviceDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceDefinition"}
			}
		}
		return impl.SearchCapabilitiesDeviceDefinition(), nil
	case "DeviceMetric":
		impl, ok := api.(r4.DeviceMetricSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceMetric"}
			}
		}
		return impl.SearchCapabilitiesDeviceMetric(), nil
	case "DeviceRequest":
		impl, ok := api.(r4.DeviceRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceRequest"}
			}
		}
		return impl.SearchCapabilitiesDeviceRequest(), nil
	case "DeviceUseStatement":
		impl, ok := api.(r4.DeviceUseStatementSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceUseStatement"}
			}
		}
		return impl.SearchCapabilitiesDeviceUseStatement(), nil
	case "DiagnosticReport":
		impl, ok := api.(r4.DiagnosticReportSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DiagnosticReport"}
			}
		}
		return impl.SearchCapabilitiesDiagnosticReport(), nil
	case "DocumentManifest":
		impl, ok := api.(r4.DocumentManifestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DocumentManifest"}
			}
		}
		return impl.SearchCapabilitiesDocumentManifest(), nil
	case "DocumentReference":
		impl, ok := api.(r4.DocumentReferenceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "DocumentReference"}
			}
		}
		return impl.SearchCapabilitiesDocumentReference(), nil
	case "EffectEvidenceSynthesis":
		impl, ok := api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EffectEvidenceSynthesis"}
			}
		}
		return impl.SearchCapabilitiesEffectEvidenceSynthesis(), nil
	case "Encounter":
		impl, ok := api.(r4.EncounterSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Encounter"}
			}
		}
		return impl.SearchCapabilitiesEncounter(), nil
	case "Endpoint":
		impl, ok := api.(r4.EndpointSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Endpoint"}
			}
		}
		return impl.SearchCapabilitiesEndpoint(), nil
	case "EnrollmentRequest":
		impl, ok := api.(r4.EnrollmentRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EnrollmentRequest"}
			}
		}
		return impl.SearchCapabilitiesEnrollmentRequest(), nil
	case "EnrollmentResponse":
		impl, ok := api.(r4.EnrollmentResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EnrollmentResponse"}
			}
		}
		return impl.SearchCapabilitiesEnrollmentResponse(), nil
	case "EpisodeOfCare":
		impl, ok := api.(r4.EpisodeOfCareSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EpisodeOfCare"}
			}
		}
		return impl.SearchCapabilitiesEpisodeOfCare(), nil
	case "EventDefinition":
		impl, ok := api.(r4.EventDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EventDefinition"}
			}
		}
		return impl.SearchCapabilitiesEventDefinition(), nil
	case "Evidence":
		impl, ok := api.(r4.EvidenceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Evidence"}
			}
		}
		return impl.SearchCapabilitiesEvidence(), nil
	case "EvidenceVariable":
		impl, ok := api.(r4.EvidenceVariableSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "EvidenceVariable"}
			}
		}
		return impl.SearchCapabilitiesEvidenceVariable(), nil
	case "ExampleScenario":
		impl, ok := api.(r4.ExampleScenarioSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ExampleScenario"}
			}
		}
		return impl.SearchCapabilitiesExampleScenario(), nil
	case "ExplanationOfBenefit":
		impl, ok := api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ExplanationOfBenefit"}
			}
		}
		return impl.SearchCapabilitiesExplanationOfBenefit(), nil
	case "FamilyMemberHistory":
		impl, ok := api.(r4.FamilyMemberHistorySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "FamilyMemberHistory"}
			}
		}
		return impl.SearchCapabilitiesFamilyMemberHistory(), nil
	case "Flag":
		impl, ok := api.(r4.FlagSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Flag"}
			}
		}
		return impl.SearchCapabilitiesFlag(), nil
	case "Goal":
		impl, ok := api.(r4.GoalSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Goal"}
			}
		}
		return impl.SearchCapabilitiesGoal(), nil
	case "GraphDefinition":
		impl, ok := api.(r4.GraphDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "GraphDefinition"}
			}
		}
		return impl.SearchCapabilitiesGraphDefinition(), nil
	case "Group":
		impl, ok := api.(r4.GroupSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Group"}
			}
		}
		return impl.SearchCapabilitiesGroup(), nil
	case "GuidanceResponse":
		impl, ok := api.(r4.GuidanceResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "GuidanceResponse"}
			}
		}
		return impl.SearchCapabilitiesGuidanceResponse(), nil
	case "HealthcareService":
		impl, ok := api.(r4.HealthcareServiceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "HealthcareService"}
			}
		}
		return impl.SearchCapabilitiesHealthcareService(), nil
	case "ImagingStudy":
		impl, ok := api.(r4.ImagingStudySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImagingStudy"}
			}
		}
		return impl.SearchCapabilitiesImagingStudy(), nil
	case "Immunization":
		impl, ok := api.(r4.ImmunizationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Immunization"}
			}
		}
		return impl.SearchCapabilitiesImmunization(), nil
	case "ImmunizationEvaluation":
		impl, ok := api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImmunizationEvaluation"}
			}
		}
		return impl.SearchCapabilitiesImmunizationEvaluation(), nil
	case "ImmunizationRecommendation":
		impl, ok := api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImmunizationRecommendation"}
			}
		}
		return impl.SearchCapabilitiesImmunizationRecommendation(), nil
	case "ImplementationGuide":
		impl, ok := api.(r4.ImplementationGuideSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImplementationGuide"}
			}
		}
		return impl.SearchCapabilitiesImplementationGuide(), nil
	case "InsurancePlan":
		impl, ok := api.(r4.InsurancePlanSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "InsurancePlan"}
			}
		}
		return impl.SearchCapabilitiesInsurancePlan(), nil
	case "Invoice":
		impl, ok := api.(r4.InvoiceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Invoice"}
			}
		}
		return impl.SearchCapabilitiesInvoice(), nil
	case "Library":
		impl, ok := api.(r4.LibrarySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Library"}
			}
		}
		return impl.SearchCapabilitiesLibrary(), nil
	case "Linkage":
		impl, ok := api.(r4.LinkageSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Linkage"}
			}
		}
		return impl.SearchCapabilitiesLinkage(), nil
	case "List":
		impl, ok := api.(r4.ListSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "List"}
			}
		}
		return impl.SearchCapabilitiesList(), nil
	case "Location":
		impl, ok := api.(r4.LocationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Location"}
			}
		}
		return impl.SearchCapabilitiesLocation(), nil
	case "Measure":
		impl, ok := api.(r4.MeasureSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Measure"}
			}
		}
		return impl.SearchCapabilitiesMeasure(), nil
	case "MeasureReport":
		impl, ok := api.(r4.MeasureReportSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MeasureReport"}
			}
		}
		return impl.SearchCapabilitiesMeasureReport(), nil
	case "Media":
		impl, ok := api.(r4.MediaSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Media"}
			}
		}
		return impl.SearchCapabilitiesMedia(), nil
	case "Medication":
		impl, ok := api.(r4.MedicationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Medication"}
			}
		}
		return impl.SearchCapabilitiesMedication(), nil
	case "MedicationAdministration":
		impl, ok := api.(r4.MedicationAdministrationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationAdministration"}
			}
		}
		return impl.SearchCapabilitiesMedicationAdministration(), nil
	case "MedicationDispense":
		impl, ok := api.(r4.MedicationDispenseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationDispense"}
			}
		}
		return impl.SearchCapabilitiesMedicationDispense(), nil
	case "MedicationKnowledge":
		impl, ok := api.(r4.MedicationKnowledgeSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationKnowledge"}
			}
		}
		return impl.SearchCapabilitiesMedicationKnowledge(), nil
	case "MedicationRequest":
		impl, ok := api.(r4.MedicationRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationRequest"}
			}
		}
		return impl.SearchCapabilitiesMedicationRequest(), nil
	case "MedicationStatement":
		impl, ok := api.(r4.MedicationStatementSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationStatement"}
			}
		}
		return impl.SearchCapabilitiesMedicationStatement(), nil
	case "MedicinalProduct":
		impl, ok := api.(r4.MedicinalProductSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProduct"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProduct(), nil
	case "MedicinalProductAuthorization":
		impl, ok := api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductAuthorization"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductAuthorization(), nil
	case "MedicinalProductContraindication":
		impl, ok := api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductContraindication"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductContraindication(), nil
	case "MedicinalProductIndication":
		impl, ok := api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductIndication"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductIndication(), nil
	case "MedicinalProductIngredient":
		impl, ok := api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductIngredient"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductIngredient(), nil
	case "MedicinalProductInteraction":
		impl, ok := api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductInteraction"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductInteraction(), nil
	case "MedicinalProductManufactured":
		impl, ok := api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductManufactured"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductManufactured(), nil
	case "MedicinalProductPackaged":
		impl, ok := api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductPackaged"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductPackaged(), nil
	case "MedicinalProductPharmaceutical":
		impl, ok := api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductPharmaceutical"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductPharmaceutical(), nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductUndesirableEffect"}
			}
		}
		return impl.SearchCapabilitiesMedicinalProductUndesirableEffect(), nil
	case "MessageDefinition":
		impl, ok := api.(r4.MessageDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MessageDefinition"}
			}
		}
		return impl.SearchCapabilitiesMessageDefinition(), nil
	case "MessageHeader":
		impl, ok := api.(r4.MessageHeaderSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MessageHeader"}
			}
		}
		return impl.SearchCapabilitiesMessageHeader(), nil
	case "MolecularSequence":
		impl, ok := api.(r4.MolecularSequenceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "MolecularSequence"}
			}
		}
		return impl.SearchCapabilitiesMolecularSequence(), nil
	case "NamingSystem":
		impl, ok := api.(r4.NamingSystemSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "NamingSystem"}
			}
		}
		return impl.SearchCapabilitiesNamingSystem(), nil
	case "NutritionOrder":
		impl, ok := api.(r4.NutritionOrderSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "NutritionOrder"}
			}
		}
		return impl.SearchCapabilitiesNutritionOrder(), nil
	case "Observation":
		impl, ok := api.(r4.ObservationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Observation"}
			}
		}
		return impl.SearchCapabilitiesObservation(), nil
	case "ObservationDefinition":
		impl, ok := api.(r4.ObservationDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ObservationDefinition"}
			}
		}
		return impl.SearchCapabilitiesObservationDefinition(), nil
	case "OperationDefinition":
		impl, ok := api.(r4.OperationDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "OperationDefinition"}
			}
		}
		return impl.SearchCapabilitiesOperationDefinition(), nil
	case "OperationOutcome":
		impl, ok := api.(r4.OperationOutcomeSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "OperationOutcome"}
			}
		}
		return impl.SearchCapabilitiesOperationOutcome(), nil
	case "Organization":
		impl, ok := api.(r4.OrganizationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Organization"}
			}
		}
		return impl.SearchCapabilitiesOrganization(), nil
	case "OrganizationAffiliation":
		impl, ok := api.(r4.OrganizationAffiliationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "OrganizationAffiliation"}
			}
		}
		return impl.SearchCapabilitiesOrganizationAffiliation(), nil
	case "Parameters":
		impl, ok := api.(r4.ParametersSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Parameters"}
			}
		}
		return impl.SearchCapabilitiesParameters(), nil
	case "Patient":
		impl, ok := api.(r4.PatientSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Patient"}
			}
		}
		return impl.SearchCapabilitiesPatient(), nil
	case "PaymentNotice":
		impl, ok := api.(r4.PaymentNoticeSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "PaymentNotice"}
			}
		}
		return impl.SearchCapabilitiesPaymentNotice(), nil
	case "PaymentReconciliation":
		impl, ok := api.(r4.PaymentReconciliationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "PaymentReconciliation"}
			}
		}
		return impl.SearchCapabilitiesPaymentReconciliation(), nil
	case "Person":
		impl, ok := api.(r4.PersonSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Person"}
			}
		}
		return impl.SearchCapabilitiesPerson(), nil
	case "PlanDefinition":
		impl, ok := api.(r4.PlanDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "PlanDefinition"}
			}
		}
		return impl.SearchCapabilitiesPlanDefinition(), nil
	case "Practitioner":
		impl, ok := api.(r4.PractitionerSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Practitioner"}
			}
		}
		return impl.SearchCapabilitiesPractitioner(), nil
	case "PractitionerRole":
		impl, ok := api.(r4.PractitionerRoleSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "PractitionerRole"}
			}
		}
		return impl.SearchCapabilitiesPractitionerRole(), nil
	case "Procedure":
		impl, ok := api.(r4.ProcedureSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Procedure"}
			}
		}
		return impl.SearchCapabilitiesProcedure(), nil
	case "Provenance":
		impl, ok := api.(r4.ProvenanceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Provenance"}
			}
		}
		return impl.SearchCapabilitiesProvenance(), nil
	case "Questionnaire":
		impl, ok := api.(r4.QuestionnaireSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Questionnaire"}
			}
		}
		return impl.SearchCapabilitiesQuestionnaire(), nil
	case "QuestionnaireResponse":
		impl, ok := api.(r4.QuestionnaireResponseSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "QuestionnaireResponse"}
			}
		}
		return impl.SearchCapabilitiesQuestionnaireResponse(), nil
	case "RelatedPerson":
		impl, ok := api.(r4.RelatedPersonSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "RelatedPerson"}
			}
		}
		return impl.SearchCapabilitiesRelatedPerson(), nil
	case "RequestGroup":
		impl, ok := api.(r4.RequestGroupSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "RequestGroup"}
			}
		}
		return impl.SearchCapabilitiesRequestGroup(), nil
	case "ResearchDefinition":
		impl, ok := api.(r4.ResearchDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchDefinition"}
			}
		}
		return impl.SearchCapabilitiesResearchDefinition(), nil
	case "ResearchElementDefinition":
		impl, ok := api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchElementDefinition"}
			}
		}
		return impl.SearchCapabilitiesResearchElementDefinition(), nil
	case "ResearchStudy":
		impl, ok := api.(r4.ResearchStudySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchStudy"}
			}
		}
		return impl.SearchCapabilitiesResearchStudy(), nil
	case "ResearchSubject":
		impl, ok := api.(r4.ResearchSubjectSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchSubject"}
			}
		}
		return impl.SearchCapabilitiesResearchSubject(), nil
	case "RiskAssessment":
		impl, ok := api.(r4.RiskAssessmentSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "RiskAssessment"}
			}
		}
		return impl.SearchCapabilitiesRiskAssessment(), nil
	case "RiskEvidenceSynthesis":
		impl, ok := api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "RiskEvidenceSynthesis"}
			}
		}
		return impl.SearchCapabilitiesRiskEvidenceSynthesis(), nil
	case "Schedule":
		impl, ok := api.(r4.ScheduleSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Schedule"}
			}
		}
		return impl.SearchCapabilitiesSchedule(), nil
	case "SearchParameter":
		impl, ok := api.(r4.SearchParameterSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SearchParameter"}
			}
		}
		return impl.SearchCapabilitiesSearchParameter(), nil
	case "ServiceRequest":
		impl, ok := api.(r4.ServiceRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ServiceRequest"}
			}
		}
		return impl.SearchCapabilitiesServiceRequest(), nil
	case "Slot":
		impl, ok := api.(r4.SlotSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Slot"}
			}
		}
		return impl.SearchCapabilitiesSlot(), nil
	case "Specimen":
		impl, ok := api.(r4.SpecimenSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Specimen"}
			}
		}
		return impl.SearchCapabilitiesSpecimen(), nil
	case "SpecimenDefinition":
		impl, ok := api.(r4.SpecimenDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SpecimenDefinition"}
			}
		}
		return impl.SearchCapabilitiesSpecimenDefinition(), nil
	case "StructureDefinition":
		impl, ok := api.(r4.StructureDefinitionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "StructureDefinition"}
			}
		}
		return impl.SearchCapabilitiesStructureDefinition(), nil
	case "StructureMap":
		impl, ok := api.(r4.StructureMapSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "StructureMap"}
			}
		}
		return impl.SearchCapabilitiesStructureMap(), nil
	case "Subscription":
		impl, ok := api.(r4.SubscriptionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Subscription"}
			}
		}
		return impl.SearchCapabilitiesSubscription(), nil
	case "Substance":
		impl, ok := api.(r4.SubstanceSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Substance"}
			}
		}
		return impl.SearchCapabilitiesSubstance(), nil
	case "SubstanceNucleicAcid":
		impl, ok := api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceNucleicAcid"}
			}
		}
		return impl.SearchCapabilitiesSubstanceNucleicAcid(), nil
	case "SubstancePolymer":
		impl, ok := api.(r4.SubstancePolymerSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstancePolymer"}
			}
		}
		return impl.SearchCapabilitiesSubstancePolymer(), nil
	case "SubstanceProtein":
		impl, ok := api.(r4.SubstanceProteinSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceProtein"}
			}
		}
		return impl.SearchCapabilitiesSubstanceProtein(), nil
	case "SubstanceReferenceInformation":
		impl, ok := api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceReferenceInformation"}
			}
		}
		return impl.SearchCapabilitiesSubstanceReferenceInformation(), nil
	case "SubstanceSourceMaterial":
		impl, ok := api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceSourceMaterial"}
			}
		}
		return impl.SearchCapabilitiesSubstanceSourceMaterial(), nil
	case "SubstanceSpecification":
		impl, ok := api.(r4.SubstanceSpecificationSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceSpecification"}
			}
		}
		return impl.SearchCapabilitiesSubstanceSpecification(), nil
	case "SupplyDelivery":
		impl, ok := api.(r4.SupplyDeliverySearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SupplyDelivery"}
			}
		}
		return impl.SearchCapabilitiesSupplyDelivery(), nil
	case "SupplyRequest":
		impl, ok := api.(r4.SupplyRequestSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "SupplyRequest"}
			}
		}
		return impl.SearchCapabilitiesSupplyRequest(), nil
	case "Task":
		impl, ok := api.(r4.TaskSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "Task"}
			}
		}
		return impl.SearchCapabilitiesTask(), nil
	case "TerminologyCapabilities":
		impl, ok := api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "TerminologyCapabilities"}
			}
		}
		return impl.SearchCapabilitiesTerminologyCapabilities(), nil
	case "TestReport":
		impl, ok := api.(r4.TestReportSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "TestReport"}
			}
		}
		return impl.SearchCapabilitiesTestReport(), nil
	case "TestScript":
		impl, ok := api.(r4.TestScriptSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "TestScript"}
			}
		}
		return impl.SearchCapabilitiesTestScript(), nil
	case "ValueSet":
		impl, ok := api.(r4.ValueSetSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "ValueSet"}
			}
		}
		return impl.SearchCapabilitiesValueSet(), nil
	case "VerificationResult":
		impl, ok := api.(r4.VerificationResultSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "VerificationResult"}
			}
		}
		return impl.SearchCapabilitiesVerificationResult(), nil
	case "VisionPrescription":
		impl, ok := api.(r4.VisionPrescriptionSearch)
		if !ok {
			if !ok {
				return capabilities.SearchCapabilities{}, errors.NotImplementedError{Interaction: "SearchParams", ResourceType: "VisionPrescription"}
			}
		}
		return impl.SearchCapabilitiesVisionPrescription(), nil
	default:
		return capabilities.SearchCapabilities{}, errors.UnknownResourceError{ResourceType: resourceType}
	}
}
func Search(ctx context.Context, api any, resourceType string, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError) {
	switch resourceType {
	case "Account":
		impl, ok := api.(r4.AccountSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Account"}
		}
		v, err := impl.SearchAccount(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ActivityDefinition":
		impl, ok := api.(r4.ActivityDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ActivityDefinition"}
		}
		v, err := impl.SearchActivityDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AdverseEvent":
		impl, ok := api.(r4.AdverseEventSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "AdverseEvent"}
		}
		v, err := impl.SearchAdverseEvent(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AllergyIntolerance":
		impl, ok := api.(r4.AllergyIntoleranceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "AllergyIntolerance"}
		}
		v, err := impl.SearchAllergyIntolerance(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Appointment":
		impl, ok := api.(r4.AppointmentSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Appointment"}
		}
		v, err := impl.SearchAppointment(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AppointmentResponse":
		impl, ok := api.(r4.AppointmentResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "AppointmentResponse"}
		}
		v, err := impl.SearchAppointmentResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AuditEvent":
		impl, ok := api.(r4.AuditEventSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "AuditEvent"}
		}
		v, err := impl.SearchAuditEvent(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Basic":
		impl, ok := api.(r4.BasicSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Basic"}
		}
		v, err := impl.SearchBasic(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Binary":
		impl, ok := api.(r4.BinarySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Binary"}
		}
		v, err := impl.SearchBinary(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "BiologicallyDerivedProduct":
		impl, ok := api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "BiologicallyDerivedProduct"}
		}
		v, err := impl.SearchBiologicallyDerivedProduct(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "BodyStructure":
		impl, ok := api.(r4.BodyStructureSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "BodyStructure"}
		}
		v, err := impl.SearchBodyStructure(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Bundle":
		impl, ok := api.(r4.BundleSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Bundle"}
		}
		v, err := impl.SearchBundle(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CapabilityStatement":
		impl, ok := api.(r4.CapabilityStatementSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CapabilityStatement"}
		}
		v, err := impl.SearchCapabilityStatement(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CarePlan":
		impl, ok := api.(r4.CarePlanSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CarePlan"}
		}
		v, err := impl.SearchCarePlan(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CareTeam":
		impl, ok := api.(r4.CareTeamSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CareTeam"}
		}
		v, err := impl.SearchCareTeam(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CatalogEntry":
		impl, ok := api.(r4.CatalogEntrySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CatalogEntry"}
		}
		v, err := impl.SearchCatalogEntry(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ChargeItem":
		impl, ok := api.(r4.ChargeItemSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ChargeItem"}
		}
		v, err := impl.SearchChargeItem(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ChargeItemDefinition":
		impl, ok := api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ChargeItemDefinition"}
		}
		v, err := impl.SearchChargeItemDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Claim":
		impl, ok := api.(r4.ClaimSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Claim"}
		}
		v, err := impl.SearchClaim(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ClaimResponse":
		impl, ok := api.(r4.ClaimResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ClaimResponse"}
		}
		v, err := impl.SearchClaimResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ClinicalImpression":
		impl, ok := api.(r4.ClinicalImpressionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ClinicalImpression"}
		}
		v, err := impl.SearchClinicalImpression(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CodeSystem":
		impl, ok := api.(r4.CodeSystemSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CodeSystem"}
		}
		v, err := impl.SearchCodeSystem(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Communication":
		impl, ok := api.(r4.CommunicationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Communication"}
		}
		v, err := impl.SearchCommunication(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CommunicationRequest":
		impl, ok := api.(r4.CommunicationRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CommunicationRequest"}
		}
		v, err := impl.SearchCommunicationRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CompartmentDefinition":
		impl, ok := api.(r4.CompartmentDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CompartmentDefinition"}
		}
		v, err := impl.SearchCompartmentDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Composition":
		impl, ok := api.(r4.CompositionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Composition"}
		}
		v, err := impl.SearchComposition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ConceptMap":
		impl, ok := api.(r4.ConceptMapSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ConceptMap"}
		}
		v, err := impl.SearchConceptMap(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Condition":
		impl, ok := api.(r4.ConditionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Condition"}
		}
		v, err := impl.SearchCondition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Consent":
		impl, ok := api.(r4.ConsentSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Consent"}
		}
		v, err := impl.SearchConsent(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Contract":
		impl, ok := api.(r4.ContractSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Contract"}
		}
		v, err := impl.SearchContract(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Coverage":
		impl, ok := api.(r4.CoverageSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Coverage"}
		}
		v, err := impl.SearchCoverage(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CoverageEligibilityRequest":
		impl, ok := api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CoverageEligibilityRequest"}
		}
		v, err := impl.SearchCoverageEligibilityRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CoverageEligibilityResponse":
		impl, ok := api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "CoverageEligibilityResponse"}
		}
		v, err := impl.SearchCoverageEligibilityResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DetectedIssue":
		impl, ok := api.(r4.DetectedIssueSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DetectedIssue"}
		}
		v, err := impl.SearchDetectedIssue(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Device":
		impl, ok := api.(r4.DeviceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Device"}
		}
		v, err := impl.SearchDevice(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceDefinition":
		impl, ok := api.(r4.DeviceDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DeviceDefinition"}
		}
		v, err := impl.SearchDeviceDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceMetric":
		impl, ok := api.(r4.DeviceMetricSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DeviceMetric"}
		}
		v, err := impl.SearchDeviceMetric(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceRequest":
		impl, ok := api.(r4.DeviceRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DeviceRequest"}
		}
		v, err := impl.SearchDeviceRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceUseStatement":
		impl, ok := api.(r4.DeviceUseStatementSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DeviceUseStatement"}
		}
		v, err := impl.SearchDeviceUseStatement(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DiagnosticReport":
		impl, ok := api.(r4.DiagnosticReportSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DiagnosticReport"}
		}
		v, err := impl.SearchDiagnosticReport(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DocumentManifest":
		impl, ok := api.(r4.DocumentManifestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DocumentManifest"}
		}
		v, err := impl.SearchDocumentManifest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DocumentReference":
		impl, ok := api.(r4.DocumentReferenceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "DocumentReference"}
		}
		v, err := impl.SearchDocumentReference(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EffectEvidenceSynthesis":
		impl, ok := api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EffectEvidenceSynthesis"}
		}
		v, err := impl.SearchEffectEvidenceSynthesis(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Encounter":
		impl, ok := api.(r4.EncounterSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Encounter"}
		}
		v, err := impl.SearchEncounter(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Endpoint":
		impl, ok := api.(r4.EndpointSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Endpoint"}
		}
		v, err := impl.SearchEndpoint(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EnrollmentRequest":
		impl, ok := api.(r4.EnrollmentRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EnrollmentRequest"}
		}
		v, err := impl.SearchEnrollmentRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EnrollmentResponse":
		impl, ok := api.(r4.EnrollmentResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EnrollmentResponse"}
		}
		v, err := impl.SearchEnrollmentResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EpisodeOfCare":
		impl, ok := api.(r4.EpisodeOfCareSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EpisodeOfCare"}
		}
		v, err := impl.SearchEpisodeOfCare(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EventDefinition":
		impl, ok := api.(r4.EventDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EventDefinition"}
		}
		v, err := impl.SearchEventDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Evidence":
		impl, ok := api.(r4.EvidenceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Evidence"}
		}
		v, err := impl.SearchEvidence(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EvidenceVariable":
		impl, ok := api.(r4.EvidenceVariableSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "EvidenceVariable"}
		}
		v, err := impl.SearchEvidenceVariable(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ExampleScenario":
		impl, ok := api.(r4.ExampleScenarioSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ExampleScenario"}
		}
		v, err := impl.SearchExampleScenario(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ExplanationOfBenefit":
		impl, ok := api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ExplanationOfBenefit"}
		}
		v, err := impl.SearchExplanationOfBenefit(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "FamilyMemberHistory":
		impl, ok := api.(r4.FamilyMemberHistorySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "FamilyMemberHistory"}
		}
		v, err := impl.SearchFamilyMemberHistory(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Flag":
		impl, ok := api.(r4.FlagSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Flag"}
		}
		v, err := impl.SearchFlag(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Goal":
		impl, ok := api.(r4.GoalSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Goal"}
		}
		v, err := impl.SearchGoal(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "GraphDefinition":
		impl, ok := api.(r4.GraphDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "GraphDefinition"}
		}
		v, err := impl.SearchGraphDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Group":
		impl, ok := api.(r4.GroupSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Group"}
		}
		v, err := impl.SearchGroup(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "GuidanceResponse":
		impl, ok := api.(r4.GuidanceResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "GuidanceResponse"}
		}
		v, err := impl.SearchGuidanceResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "HealthcareService":
		impl, ok := api.(r4.HealthcareServiceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "HealthcareService"}
		}
		v, err := impl.SearchHealthcareService(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImagingStudy":
		impl, ok := api.(r4.ImagingStudySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ImagingStudy"}
		}
		v, err := impl.SearchImagingStudy(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Immunization":
		impl, ok := api.(r4.ImmunizationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Immunization"}
		}
		v, err := impl.SearchImmunization(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImmunizationEvaluation":
		impl, ok := api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ImmunizationEvaluation"}
		}
		v, err := impl.SearchImmunizationEvaluation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImmunizationRecommendation":
		impl, ok := api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ImmunizationRecommendation"}
		}
		v, err := impl.SearchImmunizationRecommendation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImplementationGuide":
		impl, ok := api.(r4.ImplementationGuideSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ImplementationGuide"}
		}
		v, err := impl.SearchImplementationGuide(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "InsurancePlan":
		impl, ok := api.(r4.InsurancePlanSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "InsurancePlan"}
		}
		v, err := impl.SearchInsurancePlan(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Invoice":
		impl, ok := api.(r4.InvoiceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Invoice"}
		}
		v, err := impl.SearchInvoice(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Library":
		impl, ok := api.(r4.LibrarySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Library"}
		}
		v, err := impl.SearchLibrary(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Linkage":
		impl, ok := api.(r4.LinkageSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Linkage"}
		}
		v, err := impl.SearchLinkage(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "List":
		impl, ok := api.(r4.ListSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "List"}
		}
		v, err := impl.SearchList(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Location":
		impl, ok := api.(r4.LocationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Location"}
		}
		v, err := impl.SearchLocation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Measure":
		impl, ok := api.(r4.MeasureSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Measure"}
		}
		v, err := impl.SearchMeasure(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MeasureReport":
		impl, ok := api.(r4.MeasureReportSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MeasureReport"}
		}
		v, err := impl.SearchMeasureReport(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Media":
		impl, ok := api.(r4.MediaSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Media"}
		}
		v, err := impl.SearchMedia(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Medication":
		impl, ok := api.(r4.MedicationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Medication"}
		}
		v, err := impl.SearchMedication(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationAdministration":
		impl, ok := api.(r4.MedicationAdministrationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicationAdministration"}
		}
		v, err := impl.SearchMedicationAdministration(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationDispense":
		impl, ok := api.(r4.MedicationDispenseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicationDispense"}
		}
		v, err := impl.SearchMedicationDispense(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationKnowledge":
		impl, ok := api.(r4.MedicationKnowledgeSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicationKnowledge"}
		}
		v, err := impl.SearchMedicationKnowledge(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationRequest":
		impl, ok := api.(r4.MedicationRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicationRequest"}
		}
		v, err := impl.SearchMedicationRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationStatement":
		impl, ok := api.(r4.MedicationStatementSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicationStatement"}
		}
		v, err := impl.SearchMedicationStatement(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProduct":
		impl, ok := api.(r4.MedicinalProductSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProduct"}
		}
		v, err := impl.SearchMedicinalProduct(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductAuthorization":
		impl, ok := api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductAuthorization"}
		}
		v, err := impl.SearchMedicinalProductAuthorization(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductContraindication":
		impl, ok := api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductContraindication"}
		}
		v, err := impl.SearchMedicinalProductContraindication(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductIndication":
		impl, ok := api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductIndication"}
		}
		v, err := impl.SearchMedicinalProductIndication(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductIngredient":
		impl, ok := api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductIngredient"}
		}
		v, err := impl.SearchMedicinalProductIngredient(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductInteraction":
		impl, ok := api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductInteraction"}
		}
		v, err := impl.SearchMedicinalProductInteraction(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductManufactured":
		impl, ok := api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductManufactured"}
		}
		v, err := impl.SearchMedicinalProductManufactured(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductPackaged":
		impl, ok := api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductPackaged"}
		}
		v, err := impl.SearchMedicinalProductPackaged(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductPharmaceutical":
		impl, ok := api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		v, err := impl.SearchMedicinalProductPharmaceutical(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		v, err := impl.SearchMedicinalProductUndesirableEffect(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MessageDefinition":
		impl, ok := api.(r4.MessageDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MessageDefinition"}
		}
		v, err := impl.SearchMessageDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MessageHeader":
		impl, ok := api.(r4.MessageHeaderSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MessageHeader"}
		}
		v, err := impl.SearchMessageHeader(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MolecularSequence":
		impl, ok := api.(r4.MolecularSequenceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "MolecularSequence"}
		}
		v, err := impl.SearchMolecularSequence(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "NamingSystem":
		impl, ok := api.(r4.NamingSystemSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "NamingSystem"}
		}
		v, err := impl.SearchNamingSystem(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "NutritionOrder":
		impl, ok := api.(r4.NutritionOrderSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "NutritionOrder"}
		}
		v, err := impl.SearchNutritionOrder(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Observation":
		impl, ok := api.(r4.ObservationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Observation"}
		}
		v, err := impl.SearchObservation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ObservationDefinition":
		impl, ok := api.(r4.ObservationDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ObservationDefinition"}
		}
		v, err := impl.SearchObservationDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OperationDefinition":
		impl, ok := api.(r4.OperationDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "OperationDefinition"}
		}
		v, err := impl.SearchOperationDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OperationOutcome":
		impl, ok := api.(r4.OperationOutcomeSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "OperationOutcome"}
		}
		v, err := impl.SearchOperationOutcome(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Organization":
		impl, ok := api.(r4.OrganizationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Organization"}
		}
		v, err := impl.SearchOrganization(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OrganizationAffiliation":
		impl, ok := api.(r4.OrganizationAffiliationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "OrganizationAffiliation"}
		}
		v, err := impl.SearchOrganizationAffiliation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Parameters":
		impl, ok := api.(r4.ParametersSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Parameters"}
		}
		v, err := impl.SearchParameters(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Patient":
		impl, ok := api.(r4.PatientSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Patient"}
		}
		v, err := impl.SearchPatient(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PaymentNotice":
		impl, ok := api.(r4.PaymentNoticeSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "PaymentNotice"}
		}
		v, err := impl.SearchPaymentNotice(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PaymentReconciliation":
		impl, ok := api.(r4.PaymentReconciliationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "PaymentReconciliation"}
		}
		v, err := impl.SearchPaymentReconciliation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Person":
		impl, ok := api.(r4.PersonSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Person"}
		}
		v, err := impl.SearchPerson(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PlanDefinition":
		impl, ok := api.(r4.PlanDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "PlanDefinition"}
		}
		v, err := impl.SearchPlanDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Practitioner":
		impl, ok := api.(r4.PractitionerSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Practitioner"}
		}
		v, err := impl.SearchPractitioner(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PractitionerRole":
		impl, ok := api.(r4.PractitionerRoleSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "PractitionerRole"}
		}
		v, err := impl.SearchPractitionerRole(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Procedure":
		impl, ok := api.(r4.ProcedureSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Procedure"}
		}
		v, err := impl.SearchProcedure(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Provenance":
		impl, ok := api.(r4.ProvenanceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Provenance"}
		}
		v, err := impl.SearchProvenance(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Questionnaire":
		impl, ok := api.(r4.QuestionnaireSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Questionnaire"}
		}
		v, err := impl.SearchQuestionnaire(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "QuestionnaireResponse":
		impl, ok := api.(r4.QuestionnaireResponseSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "QuestionnaireResponse"}
		}
		v, err := impl.SearchQuestionnaireResponse(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RelatedPerson":
		impl, ok := api.(r4.RelatedPersonSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "RelatedPerson"}
		}
		v, err := impl.SearchRelatedPerson(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RequestGroup":
		impl, ok := api.(r4.RequestGroupSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "RequestGroup"}
		}
		v, err := impl.SearchRequestGroup(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchDefinition":
		impl, ok := api.(r4.ResearchDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ResearchDefinition"}
		}
		v, err := impl.SearchResearchDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchElementDefinition":
		impl, ok := api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ResearchElementDefinition"}
		}
		v, err := impl.SearchResearchElementDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchStudy":
		impl, ok := api.(r4.ResearchStudySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ResearchStudy"}
		}
		v, err := impl.SearchResearchStudy(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchSubject":
		impl, ok := api.(r4.ResearchSubjectSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ResearchSubject"}
		}
		v, err := impl.SearchResearchSubject(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RiskAssessment":
		impl, ok := api.(r4.RiskAssessmentSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "RiskAssessment"}
		}
		v, err := impl.SearchRiskAssessment(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RiskEvidenceSynthesis":
		impl, ok := api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "RiskEvidenceSynthesis"}
		}
		v, err := impl.SearchRiskEvidenceSynthesis(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Schedule":
		impl, ok := api.(r4.ScheduleSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Schedule"}
		}
		v, err := impl.SearchSchedule(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SearchParameter":
		impl, ok := api.(r4.SearchParameterSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SearchParameter"}
		}
		v, err := impl.SearchSearchParameter(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ServiceRequest":
		impl, ok := api.(r4.ServiceRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ServiceRequest"}
		}
		v, err := impl.SearchServiceRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Slot":
		impl, ok := api.(r4.SlotSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Slot"}
		}
		v, err := impl.SearchSlot(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Specimen":
		impl, ok := api.(r4.SpecimenSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Specimen"}
		}
		v, err := impl.SearchSpecimen(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SpecimenDefinition":
		impl, ok := api.(r4.SpecimenDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SpecimenDefinition"}
		}
		v, err := impl.SearchSpecimenDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "StructureDefinition":
		impl, ok := api.(r4.StructureDefinitionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "StructureDefinition"}
		}
		v, err := impl.SearchStructureDefinition(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "StructureMap":
		impl, ok := api.(r4.StructureMapSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "StructureMap"}
		}
		v, err := impl.SearchStructureMap(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Subscription":
		impl, ok := api.(r4.SubscriptionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Subscription"}
		}
		v, err := impl.SearchSubscription(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Substance":
		impl, ok := api.(r4.SubstanceSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Substance"}
		}
		v, err := impl.SearchSubstance(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceNucleicAcid":
		impl, ok := api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceNucleicAcid"}
		}
		v, err := impl.SearchSubstanceNucleicAcid(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstancePolymer":
		impl, ok := api.(r4.SubstancePolymerSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstancePolymer"}
		}
		v, err := impl.SearchSubstancePolymer(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceProtein":
		impl, ok := api.(r4.SubstanceProteinSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceProtein"}
		}
		v, err := impl.SearchSubstanceProtein(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceReferenceInformation":
		impl, ok := api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceReferenceInformation"}
		}
		v, err := impl.SearchSubstanceReferenceInformation(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceSourceMaterial":
		impl, ok := api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceSourceMaterial"}
		}
		v, err := impl.SearchSubstanceSourceMaterial(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceSpecification":
		impl, ok := api.(r4.SubstanceSpecificationSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceSpecification"}
		}
		v, err := impl.SearchSubstanceSpecification(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SupplyDelivery":
		impl, ok := api.(r4.SupplyDeliverySearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SupplyDelivery"}
		}
		v, err := impl.SearchSupplyDelivery(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SupplyRequest":
		impl, ok := api.(r4.SupplyRequestSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "SupplyRequest"}
		}
		v, err := impl.SearchSupplyRequest(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Task":
		impl, ok := api.(r4.TaskSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "Task"}
		}
		v, err := impl.SearchTask(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TerminologyCapabilities":
		impl, ok := api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "TerminologyCapabilities"}
		}
		v, err := impl.SearchTerminologyCapabilities(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TestReport":
		impl, ok := api.(r4.TestReportSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "TestReport"}
		}
		v, err := impl.SearchTestReport(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TestScript":
		impl, ok := api.(r4.TestScriptSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "TestScript"}
		}
		v, err := impl.SearchTestScript(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ValueSet":
		impl, ok := api.(r4.ValueSetSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "ValueSet"}
		}
		v, err := impl.SearchValueSet(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "VerificationResult":
		impl, ok := api.(r4.VerificationResultSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "VerificationResult"}
		}
		v, err := impl.SearchVerificationResult(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "VisionPrescription":
		impl, ok := api.(r4.VisionPrescriptionSearch)
		if !ok {
			return nil, errors.NotImplementedError{Interaction: "Search", ResourceType: "VisionPrescription"}
		}
		v, err := impl.SearchVisionPrescription(ctx, options)
		if err != nil {
			return nil, err
		}
		r := make([]model.Resource, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	default:
		return nil, errors.UnknownResourceError{ResourceType: resourceType}
	}
}
