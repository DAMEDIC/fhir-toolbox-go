package dispatchR4

import (
	"context"
	r4 "fhir-toolbox/capabilities/gen/r4"
	dispatch "fhir-toolbox/dispatch"
)

func SearchParams(api any, resourceType string) ([]string, error) {
	switch resourceType {
	case "Account":
		impl, ok := api.(r4.AccountSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Account"}
			}
		}
		return impl.SearchParamsAccount(), nil
	case "ActivityDefinition":
		impl, ok := api.(r4.ActivityDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ActivityDefinition"}
			}
		}
		return impl.SearchParamsActivityDefinition(), nil
	case "AdverseEvent":
		impl, ok := api.(r4.AdverseEventSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "AdverseEvent"}
			}
		}
		return impl.SearchParamsAdverseEvent(), nil
	case "AllergyIntolerance":
		impl, ok := api.(r4.AllergyIntoleranceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "AllergyIntolerance"}
			}
		}
		return impl.SearchParamsAllergyIntolerance(), nil
	case "Appointment":
		impl, ok := api.(r4.AppointmentSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Appointment"}
			}
		}
		return impl.SearchParamsAppointment(), nil
	case "AppointmentResponse":
		impl, ok := api.(r4.AppointmentResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "AppointmentResponse"}
			}
		}
		return impl.SearchParamsAppointmentResponse(), nil
	case "AuditEvent":
		impl, ok := api.(r4.AuditEventSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "AuditEvent"}
			}
		}
		return impl.SearchParamsAuditEvent(), nil
	case "Basic":
		impl, ok := api.(r4.BasicSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Basic"}
			}
		}
		return impl.SearchParamsBasic(), nil
	case "Binary":
		impl, ok := api.(r4.BinarySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Binary"}
			}
		}
		return impl.SearchParamsBinary(), nil
	case "BiologicallyDerivedProduct":
		impl, ok := api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "BiologicallyDerivedProduct"}
			}
		}
		return impl.SearchParamsBiologicallyDerivedProduct(), nil
	case "BodyStructure":
		impl, ok := api.(r4.BodyStructureSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "BodyStructure"}
			}
		}
		return impl.SearchParamsBodyStructure(), nil
	case "Bundle":
		impl, ok := api.(r4.BundleSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Bundle"}
			}
		}
		return impl.SearchParamsBundle(), nil
	case "CapabilityStatement":
		impl, ok := api.(r4.CapabilityStatementSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CapabilityStatement"}
			}
		}
		return impl.SearchParamsCapabilityStatement(), nil
	case "CarePlan":
		impl, ok := api.(r4.CarePlanSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CarePlan"}
			}
		}
		return impl.SearchParamsCarePlan(), nil
	case "CareTeam":
		impl, ok := api.(r4.CareTeamSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CareTeam"}
			}
		}
		return impl.SearchParamsCareTeam(), nil
	case "CatalogEntry":
		impl, ok := api.(r4.CatalogEntrySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CatalogEntry"}
			}
		}
		return impl.SearchParamsCatalogEntry(), nil
	case "ChargeItem":
		impl, ok := api.(r4.ChargeItemSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ChargeItem"}
			}
		}
		return impl.SearchParamsChargeItem(), nil
	case "ChargeItemDefinition":
		impl, ok := api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ChargeItemDefinition"}
			}
		}
		return impl.SearchParamsChargeItemDefinition(), nil
	case "Claim":
		impl, ok := api.(r4.ClaimSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Claim"}
			}
		}
		return impl.SearchParamsClaim(), nil
	case "ClaimResponse":
		impl, ok := api.(r4.ClaimResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ClaimResponse"}
			}
		}
		return impl.SearchParamsClaimResponse(), nil
	case "ClinicalImpression":
		impl, ok := api.(r4.ClinicalImpressionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ClinicalImpression"}
			}
		}
		return impl.SearchParamsClinicalImpression(), nil
	case "CodeSystem":
		impl, ok := api.(r4.CodeSystemSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CodeSystem"}
			}
		}
		return impl.SearchParamsCodeSystem(), nil
	case "Communication":
		impl, ok := api.(r4.CommunicationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Communication"}
			}
		}
		return impl.SearchParamsCommunication(), nil
	case "CommunicationRequest":
		impl, ok := api.(r4.CommunicationRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CommunicationRequest"}
			}
		}
		return impl.SearchParamsCommunicationRequest(), nil
	case "CompartmentDefinition":
		impl, ok := api.(r4.CompartmentDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CompartmentDefinition"}
			}
		}
		return impl.SearchParamsCompartmentDefinition(), nil
	case "Composition":
		impl, ok := api.(r4.CompositionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Composition"}
			}
		}
		return impl.SearchParamsComposition(), nil
	case "ConceptMap":
		impl, ok := api.(r4.ConceptMapSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ConceptMap"}
			}
		}
		return impl.SearchParamsConceptMap(), nil
	case "Condition":
		impl, ok := api.(r4.ConditionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Condition"}
			}
		}
		return impl.SearchParamsCondition(), nil
	case "Consent":
		impl, ok := api.(r4.ConsentSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Consent"}
			}
		}
		return impl.SearchParamsConsent(), nil
	case "Contract":
		impl, ok := api.(r4.ContractSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Contract"}
			}
		}
		return impl.SearchParamsContract(), nil
	case "Coverage":
		impl, ok := api.(r4.CoverageSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Coverage"}
			}
		}
		return impl.SearchParamsCoverage(), nil
	case "CoverageEligibilityRequest":
		impl, ok := api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CoverageEligibilityRequest"}
			}
		}
		return impl.SearchParamsCoverageEligibilityRequest(), nil
	case "CoverageEligibilityResponse":
		impl, ok := api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "CoverageEligibilityResponse"}
			}
		}
		return impl.SearchParamsCoverageEligibilityResponse(), nil
	case "DetectedIssue":
		impl, ok := api.(r4.DetectedIssueSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DetectedIssue"}
			}
		}
		return impl.SearchParamsDetectedIssue(), nil
	case "Device":
		impl, ok := api.(r4.DeviceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Device"}
			}
		}
		return impl.SearchParamsDevice(), nil
	case "DeviceDefinition":
		impl, ok := api.(r4.DeviceDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceDefinition"}
			}
		}
		return impl.SearchParamsDeviceDefinition(), nil
	case "DeviceMetric":
		impl, ok := api.(r4.DeviceMetricSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceMetric"}
			}
		}
		return impl.SearchParamsDeviceMetric(), nil
	case "DeviceRequest":
		impl, ok := api.(r4.DeviceRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceRequest"}
			}
		}
		return impl.SearchParamsDeviceRequest(), nil
	case "DeviceUseStatement":
		impl, ok := api.(r4.DeviceUseStatementSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DeviceUseStatement"}
			}
		}
		return impl.SearchParamsDeviceUseStatement(), nil
	case "DiagnosticReport":
		impl, ok := api.(r4.DiagnosticReportSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DiagnosticReport"}
			}
		}
		return impl.SearchParamsDiagnosticReport(), nil
	case "DocumentManifest":
		impl, ok := api.(r4.DocumentManifestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DocumentManifest"}
			}
		}
		return impl.SearchParamsDocumentManifest(), nil
	case "DocumentReference":
		impl, ok := api.(r4.DocumentReferenceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "DocumentReference"}
			}
		}
		return impl.SearchParamsDocumentReference(), nil
	case "EffectEvidenceSynthesis":
		impl, ok := api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EffectEvidenceSynthesis"}
			}
		}
		return impl.SearchParamsEffectEvidenceSynthesis(), nil
	case "Encounter":
		impl, ok := api.(r4.EncounterSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Encounter"}
			}
		}
		return impl.SearchParamsEncounter(), nil
	case "Endpoint":
		impl, ok := api.(r4.EndpointSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Endpoint"}
			}
		}
		return impl.SearchParamsEndpoint(), nil
	case "EnrollmentRequest":
		impl, ok := api.(r4.EnrollmentRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EnrollmentRequest"}
			}
		}
		return impl.SearchParamsEnrollmentRequest(), nil
	case "EnrollmentResponse":
		impl, ok := api.(r4.EnrollmentResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EnrollmentResponse"}
			}
		}
		return impl.SearchParamsEnrollmentResponse(), nil
	case "EpisodeOfCare":
		impl, ok := api.(r4.EpisodeOfCareSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EpisodeOfCare"}
			}
		}
		return impl.SearchParamsEpisodeOfCare(), nil
	case "EventDefinition":
		impl, ok := api.(r4.EventDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EventDefinition"}
			}
		}
		return impl.SearchParamsEventDefinition(), nil
	case "Evidence":
		impl, ok := api.(r4.EvidenceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Evidence"}
			}
		}
		return impl.SearchParamsEvidence(), nil
	case "EvidenceVariable":
		impl, ok := api.(r4.EvidenceVariableSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "EvidenceVariable"}
			}
		}
		return impl.SearchParamsEvidenceVariable(), nil
	case "ExampleScenario":
		impl, ok := api.(r4.ExampleScenarioSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ExampleScenario"}
			}
		}
		return impl.SearchParamsExampleScenario(), nil
	case "ExplanationOfBenefit":
		impl, ok := api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ExplanationOfBenefit"}
			}
		}
		return impl.SearchParamsExplanationOfBenefit(), nil
	case "FamilyMemberHistory":
		impl, ok := api.(r4.FamilyMemberHistorySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "FamilyMemberHistory"}
			}
		}
		return impl.SearchParamsFamilyMemberHistory(), nil
	case "Flag":
		impl, ok := api.(r4.FlagSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Flag"}
			}
		}
		return impl.SearchParamsFlag(), nil
	case "Goal":
		impl, ok := api.(r4.GoalSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Goal"}
			}
		}
		return impl.SearchParamsGoal(), nil
	case "GraphDefinition":
		impl, ok := api.(r4.GraphDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "GraphDefinition"}
			}
		}
		return impl.SearchParamsGraphDefinition(), nil
	case "Group":
		impl, ok := api.(r4.GroupSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Group"}
			}
		}
		return impl.SearchParamsGroup(), nil
	case "GuidanceResponse":
		impl, ok := api.(r4.GuidanceResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "GuidanceResponse"}
			}
		}
		return impl.SearchParamsGuidanceResponse(), nil
	case "HealthcareService":
		impl, ok := api.(r4.HealthcareServiceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "HealthcareService"}
			}
		}
		return impl.SearchParamsHealthcareService(), nil
	case "ImagingStudy":
		impl, ok := api.(r4.ImagingStudySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImagingStudy"}
			}
		}
		return impl.SearchParamsImagingStudy(), nil
	case "Immunization":
		impl, ok := api.(r4.ImmunizationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Immunization"}
			}
		}
		return impl.SearchParamsImmunization(), nil
	case "ImmunizationEvaluation":
		impl, ok := api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImmunizationEvaluation"}
			}
		}
		return impl.SearchParamsImmunizationEvaluation(), nil
	case "ImmunizationRecommendation":
		impl, ok := api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImmunizationRecommendation"}
			}
		}
		return impl.SearchParamsImmunizationRecommendation(), nil
	case "ImplementationGuide":
		impl, ok := api.(r4.ImplementationGuideSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ImplementationGuide"}
			}
		}
		return impl.SearchParamsImplementationGuide(), nil
	case "InsurancePlan":
		impl, ok := api.(r4.InsurancePlanSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "InsurancePlan"}
			}
		}
		return impl.SearchParamsInsurancePlan(), nil
	case "Invoice":
		impl, ok := api.(r4.InvoiceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Invoice"}
			}
		}
		return impl.SearchParamsInvoice(), nil
	case "Library":
		impl, ok := api.(r4.LibrarySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Library"}
			}
		}
		return impl.SearchParamsLibrary(), nil
	case "Linkage":
		impl, ok := api.(r4.LinkageSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Linkage"}
			}
		}
		return impl.SearchParamsLinkage(), nil
	case "List":
		impl, ok := api.(r4.ListSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "List"}
			}
		}
		return impl.SearchParamsList(), nil
	case "Location":
		impl, ok := api.(r4.LocationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Location"}
			}
		}
		return impl.SearchParamsLocation(), nil
	case "Measure":
		impl, ok := api.(r4.MeasureSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Measure"}
			}
		}
		return impl.SearchParamsMeasure(), nil
	case "MeasureReport":
		impl, ok := api.(r4.MeasureReportSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MeasureReport"}
			}
		}
		return impl.SearchParamsMeasureReport(), nil
	case "Media":
		impl, ok := api.(r4.MediaSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Media"}
			}
		}
		return impl.SearchParamsMedia(), nil
	case "Medication":
		impl, ok := api.(r4.MedicationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Medication"}
			}
		}
		return impl.SearchParamsMedication(), nil
	case "MedicationAdministration":
		impl, ok := api.(r4.MedicationAdministrationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationAdministration"}
			}
		}
		return impl.SearchParamsMedicationAdministration(), nil
	case "MedicationDispense":
		impl, ok := api.(r4.MedicationDispenseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationDispense"}
			}
		}
		return impl.SearchParamsMedicationDispense(), nil
	case "MedicationKnowledge":
		impl, ok := api.(r4.MedicationKnowledgeSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationKnowledge"}
			}
		}
		return impl.SearchParamsMedicationKnowledge(), nil
	case "MedicationRequest":
		impl, ok := api.(r4.MedicationRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationRequest"}
			}
		}
		return impl.SearchParamsMedicationRequest(), nil
	case "MedicationStatement":
		impl, ok := api.(r4.MedicationStatementSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicationStatement"}
			}
		}
		return impl.SearchParamsMedicationStatement(), nil
	case "MedicinalProduct":
		impl, ok := api.(r4.MedicinalProductSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProduct"}
			}
		}
		return impl.SearchParamsMedicinalProduct(), nil
	case "MedicinalProductAuthorization":
		impl, ok := api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductAuthorization"}
			}
		}
		return impl.SearchParamsMedicinalProductAuthorization(), nil
	case "MedicinalProductContraindication":
		impl, ok := api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductContraindication"}
			}
		}
		return impl.SearchParamsMedicinalProductContraindication(), nil
	case "MedicinalProductIndication":
		impl, ok := api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductIndication"}
			}
		}
		return impl.SearchParamsMedicinalProductIndication(), nil
	case "MedicinalProductIngredient":
		impl, ok := api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductIngredient"}
			}
		}
		return impl.SearchParamsMedicinalProductIngredient(), nil
	case "MedicinalProductInteraction":
		impl, ok := api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductInteraction"}
			}
		}
		return impl.SearchParamsMedicinalProductInteraction(), nil
	case "MedicinalProductManufactured":
		impl, ok := api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductManufactured"}
			}
		}
		return impl.SearchParamsMedicinalProductManufactured(), nil
	case "MedicinalProductPackaged":
		impl, ok := api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductPackaged"}
			}
		}
		return impl.SearchParamsMedicinalProductPackaged(), nil
	case "MedicinalProductPharmaceutical":
		impl, ok := api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductPharmaceutical"}
			}
		}
		return impl.SearchParamsMedicinalProductPharmaceutical(), nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MedicinalProductUndesirableEffect"}
			}
		}
		return impl.SearchParamsMedicinalProductUndesirableEffect(), nil
	case "MessageDefinition":
		impl, ok := api.(r4.MessageDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MessageDefinition"}
			}
		}
		return impl.SearchParamsMessageDefinition(), nil
	case "MessageHeader":
		impl, ok := api.(r4.MessageHeaderSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MessageHeader"}
			}
		}
		return impl.SearchParamsMessageHeader(), nil
	case "MolecularSequence":
		impl, ok := api.(r4.MolecularSequenceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "MolecularSequence"}
			}
		}
		return impl.SearchParamsMolecularSequence(), nil
	case "NamingSystem":
		impl, ok := api.(r4.NamingSystemSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "NamingSystem"}
			}
		}
		return impl.SearchParamsNamingSystem(), nil
	case "NutritionOrder":
		impl, ok := api.(r4.NutritionOrderSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "NutritionOrder"}
			}
		}
		return impl.SearchParamsNutritionOrder(), nil
	case "Observation":
		impl, ok := api.(r4.ObservationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Observation"}
			}
		}
		return impl.SearchParamsObservation(), nil
	case "ObservationDefinition":
		impl, ok := api.(r4.ObservationDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ObservationDefinition"}
			}
		}
		return impl.SearchParamsObservationDefinition(), nil
	case "OperationDefinition":
		impl, ok := api.(r4.OperationDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "OperationDefinition"}
			}
		}
		return impl.SearchParamsOperationDefinition(), nil
	case "OperationOutcome":
		impl, ok := api.(r4.OperationOutcomeSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "OperationOutcome"}
			}
		}
		return impl.SearchParamsOperationOutcome(), nil
	case "Organization":
		impl, ok := api.(r4.OrganizationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Organization"}
			}
		}
		return impl.SearchParamsOrganization(), nil
	case "OrganizationAffiliation":
		impl, ok := api.(r4.OrganizationAffiliationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "OrganizationAffiliation"}
			}
		}
		return impl.SearchParamsOrganizationAffiliation(), nil
	case "Parameters":
		impl, ok := api.(r4.ParametersSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Parameters"}
			}
		}
		return impl.SearchParamsParameters(), nil
	case "Patient":
		impl, ok := api.(r4.PatientSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Patient"}
			}
		}
		return impl.SearchParamsPatient(), nil
	case "PaymentNotice":
		impl, ok := api.(r4.PaymentNoticeSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "PaymentNotice"}
			}
		}
		return impl.SearchParamsPaymentNotice(), nil
	case "PaymentReconciliation":
		impl, ok := api.(r4.PaymentReconciliationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "PaymentReconciliation"}
			}
		}
		return impl.SearchParamsPaymentReconciliation(), nil
	case "Person":
		impl, ok := api.(r4.PersonSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Person"}
			}
		}
		return impl.SearchParamsPerson(), nil
	case "PlanDefinition":
		impl, ok := api.(r4.PlanDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "PlanDefinition"}
			}
		}
		return impl.SearchParamsPlanDefinition(), nil
	case "Practitioner":
		impl, ok := api.(r4.PractitionerSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Practitioner"}
			}
		}
		return impl.SearchParamsPractitioner(), nil
	case "PractitionerRole":
		impl, ok := api.(r4.PractitionerRoleSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "PractitionerRole"}
			}
		}
		return impl.SearchParamsPractitionerRole(), nil
	case "Procedure":
		impl, ok := api.(r4.ProcedureSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Procedure"}
			}
		}
		return impl.SearchParamsProcedure(), nil
	case "Provenance":
		impl, ok := api.(r4.ProvenanceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Provenance"}
			}
		}
		return impl.SearchParamsProvenance(), nil
	case "Questionnaire":
		impl, ok := api.(r4.QuestionnaireSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Questionnaire"}
			}
		}
		return impl.SearchParamsQuestionnaire(), nil
	case "QuestionnaireResponse":
		impl, ok := api.(r4.QuestionnaireResponseSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "QuestionnaireResponse"}
			}
		}
		return impl.SearchParamsQuestionnaireResponse(), nil
	case "RelatedPerson":
		impl, ok := api.(r4.RelatedPersonSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "RelatedPerson"}
			}
		}
		return impl.SearchParamsRelatedPerson(), nil
	case "RequestGroup":
		impl, ok := api.(r4.RequestGroupSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "RequestGroup"}
			}
		}
		return impl.SearchParamsRequestGroup(), nil
	case "ResearchDefinition":
		impl, ok := api.(r4.ResearchDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchDefinition"}
			}
		}
		return impl.SearchParamsResearchDefinition(), nil
	case "ResearchElementDefinition":
		impl, ok := api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchElementDefinition"}
			}
		}
		return impl.SearchParamsResearchElementDefinition(), nil
	case "ResearchStudy":
		impl, ok := api.(r4.ResearchStudySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchStudy"}
			}
		}
		return impl.SearchParamsResearchStudy(), nil
	case "ResearchSubject":
		impl, ok := api.(r4.ResearchSubjectSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ResearchSubject"}
			}
		}
		return impl.SearchParamsResearchSubject(), nil
	case "RiskAssessment":
		impl, ok := api.(r4.RiskAssessmentSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "RiskAssessment"}
			}
		}
		return impl.SearchParamsRiskAssessment(), nil
	case "RiskEvidenceSynthesis":
		impl, ok := api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "RiskEvidenceSynthesis"}
			}
		}
		return impl.SearchParamsRiskEvidenceSynthesis(), nil
	case "Schedule":
		impl, ok := api.(r4.ScheduleSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Schedule"}
			}
		}
		return impl.SearchParamsSchedule(), nil
	case "SearchParameter":
		impl, ok := api.(r4.SearchParameterSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SearchParameter"}
			}
		}
		return impl.SearchParamsSearchParameter(), nil
	case "ServiceRequest":
		impl, ok := api.(r4.ServiceRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ServiceRequest"}
			}
		}
		return impl.SearchParamsServiceRequest(), nil
	case "Slot":
		impl, ok := api.(r4.SlotSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Slot"}
			}
		}
		return impl.SearchParamsSlot(), nil
	case "Specimen":
		impl, ok := api.(r4.SpecimenSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Specimen"}
			}
		}
		return impl.SearchParamsSpecimen(), nil
	case "SpecimenDefinition":
		impl, ok := api.(r4.SpecimenDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SpecimenDefinition"}
			}
		}
		return impl.SearchParamsSpecimenDefinition(), nil
	case "StructureDefinition":
		impl, ok := api.(r4.StructureDefinitionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "StructureDefinition"}
			}
		}
		return impl.SearchParamsStructureDefinition(), nil
	case "StructureMap":
		impl, ok := api.(r4.StructureMapSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "StructureMap"}
			}
		}
		return impl.SearchParamsStructureMap(), nil
	case "Subscription":
		impl, ok := api.(r4.SubscriptionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Subscription"}
			}
		}
		return impl.SearchParamsSubscription(), nil
	case "Substance":
		impl, ok := api.(r4.SubstanceSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Substance"}
			}
		}
		return impl.SearchParamsSubstance(), nil
	case "SubstanceNucleicAcid":
		impl, ok := api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceNucleicAcid"}
			}
		}
		return impl.SearchParamsSubstanceNucleicAcid(), nil
	case "SubstancePolymer":
		impl, ok := api.(r4.SubstancePolymerSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstancePolymer"}
			}
		}
		return impl.SearchParamsSubstancePolymer(), nil
	case "SubstanceProtein":
		impl, ok := api.(r4.SubstanceProteinSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceProtein"}
			}
		}
		return impl.SearchParamsSubstanceProtein(), nil
	case "SubstanceReferenceInformation":
		impl, ok := api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceReferenceInformation"}
			}
		}
		return impl.SearchParamsSubstanceReferenceInformation(), nil
	case "SubstanceSourceMaterial":
		impl, ok := api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceSourceMaterial"}
			}
		}
		return impl.SearchParamsSubstanceSourceMaterial(), nil
	case "SubstanceSpecification":
		impl, ok := api.(r4.SubstanceSpecificationSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SubstanceSpecification"}
			}
		}
		return impl.SearchParamsSubstanceSpecification(), nil
	case "SupplyDelivery":
		impl, ok := api.(r4.SupplyDeliverySearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SupplyDelivery"}
			}
		}
		return impl.SearchParamsSupplyDelivery(), nil
	case "SupplyRequest":
		impl, ok := api.(r4.SupplyRequestSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "SupplyRequest"}
			}
		}
		return impl.SearchParamsSupplyRequest(), nil
	case "Task":
		impl, ok := api.(r4.TaskSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "Task"}
			}
		}
		return impl.SearchParamsTask(), nil
	case "TerminologyCapabilities":
		impl, ok := api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "TerminologyCapabilities"}
			}
		}
		return impl.SearchParamsTerminologyCapabilities(), nil
	case "TestReport":
		impl, ok := api.(r4.TestReportSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "TestReport"}
			}
		}
		return impl.SearchParamsTestReport(), nil
	case "TestScript":
		impl, ok := api.(r4.TestScriptSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "TestScript"}
			}
		}
		return impl.SearchParamsTestScript(), nil
	case "ValueSet":
		impl, ok := api.(r4.ValueSetSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "ValueSet"}
			}
		}
		return impl.SearchParamsValueSet(), nil
	case "VerificationResult":
		impl, ok := api.(r4.VerificationResultSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "VerificationResult"}
			}
		}
		return impl.SearchParamsVerificationResult(), nil
	case "VisionPrescription":
		impl, ok := api.(r4.VisionPrescriptionSearch)
		if !ok {
			if !ok {
				return nil, dispatch.NotImplementedError{Interaction: "SearchParams", ResourceType: "VisionPrescription"}
			}
		}
		return impl.SearchParamsVisionPrescription(), nil
	default:
		return nil, dispatch.UnknownResourceError{ResourceType: resourceType}
	}
}
func Search(ctx context.Context, api any, resourceType string, parameters map[string]string) ([]any, error) {
	switch resourceType {
	case "Account":
		impl, ok := api.(r4.AccountSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Account"}
		}
		v, err := impl.SearchAccount(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ActivityDefinition":
		impl, ok := api.(r4.ActivityDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ActivityDefinition"}
		}
		v, err := impl.SearchActivityDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AdverseEvent":
		impl, ok := api.(r4.AdverseEventSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "AdverseEvent"}
		}
		v, err := impl.SearchAdverseEvent(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AllergyIntolerance":
		impl, ok := api.(r4.AllergyIntoleranceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "AllergyIntolerance"}
		}
		v, err := impl.SearchAllergyIntolerance(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Appointment":
		impl, ok := api.(r4.AppointmentSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Appointment"}
		}
		v, err := impl.SearchAppointment(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AppointmentResponse":
		impl, ok := api.(r4.AppointmentResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "AppointmentResponse"}
		}
		v, err := impl.SearchAppointmentResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "AuditEvent":
		impl, ok := api.(r4.AuditEventSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "AuditEvent"}
		}
		v, err := impl.SearchAuditEvent(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Basic":
		impl, ok := api.(r4.BasicSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Basic"}
		}
		v, err := impl.SearchBasic(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Binary":
		impl, ok := api.(r4.BinarySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Binary"}
		}
		v, err := impl.SearchBinary(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "BiologicallyDerivedProduct":
		impl, ok := api.(r4.BiologicallyDerivedProductSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "BiologicallyDerivedProduct"}
		}
		v, err := impl.SearchBiologicallyDerivedProduct(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "BodyStructure":
		impl, ok := api.(r4.BodyStructureSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "BodyStructure"}
		}
		v, err := impl.SearchBodyStructure(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Bundle":
		impl, ok := api.(r4.BundleSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Bundle"}
		}
		v, err := impl.SearchBundle(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CapabilityStatement":
		impl, ok := api.(r4.CapabilityStatementSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CapabilityStatement"}
		}
		v, err := impl.SearchCapabilityStatement(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CarePlan":
		impl, ok := api.(r4.CarePlanSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CarePlan"}
		}
		v, err := impl.SearchCarePlan(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CareTeam":
		impl, ok := api.(r4.CareTeamSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CareTeam"}
		}
		v, err := impl.SearchCareTeam(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CatalogEntry":
		impl, ok := api.(r4.CatalogEntrySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CatalogEntry"}
		}
		v, err := impl.SearchCatalogEntry(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ChargeItem":
		impl, ok := api.(r4.ChargeItemSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ChargeItem"}
		}
		v, err := impl.SearchChargeItem(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ChargeItemDefinition":
		impl, ok := api.(r4.ChargeItemDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ChargeItemDefinition"}
		}
		v, err := impl.SearchChargeItemDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Claim":
		impl, ok := api.(r4.ClaimSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Claim"}
		}
		v, err := impl.SearchClaim(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ClaimResponse":
		impl, ok := api.(r4.ClaimResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ClaimResponse"}
		}
		v, err := impl.SearchClaimResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ClinicalImpression":
		impl, ok := api.(r4.ClinicalImpressionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ClinicalImpression"}
		}
		v, err := impl.SearchClinicalImpression(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CodeSystem":
		impl, ok := api.(r4.CodeSystemSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CodeSystem"}
		}
		v, err := impl.SearchCodeSystem(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Communication":
		impl, ok := api.(r4.CommunicationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Communication"}
		}
		v, err := impl.SearchCommunication(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CommunicationRequest":
		impl, ok := api.(r4.CommunicationRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CommunicationRequest"}
		}
		v, err := impl.SearchCommunicationRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CompartmentDefinition":
		impl, ok := api.(r4.CompartmentDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CompartmentDefinition"}
		}
		v, err := impl.SearchCompartmentDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Composition":
		impl, ok := api.(r4.CompositionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Composition"}
		}
		v, err := impl.SearchComposition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ConceptMap":
		impl, ok := api.(r4.ConceptMapSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ConceptMap"}
		}
		v, err := impl.SearchConceptMap(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Condition":
		impl, ok := api.(r4.ConditionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Condition"}
		}
		v, err := impl.SearchCondition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Consent":
		impl, ok := api.(r4.ConsentSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Consent"}
		}
		v, err := impl.SearchConsent(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Contract":
		impl, ok := api.(r4.ContractSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Contract"}
		}
		v, err := impl.SearchContract(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Coverage":
		impl, ok := api.(r4.CoverageSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Coverage"}
		}
		v, err := impl.SearchCoverage(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CoverageEligibilityRequest":
		impl, ok := api.(r4.CoverageEligibilityRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CoverageEligibilityRequest"}
		}
		v, err := impl.SearchCoverageEligibilityRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "CoverageEligibilityResponse":
		impl, ok := api.(r4.CoverageEligibilityResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "CoverageEligibilityResponse"}
		}
		v, err := impl.SearchCoverageEligibilityResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DetectedIssue":
		impl, ok := api.(r4.DetectedIssueSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DetectedIssue"}
		}
		v, err := impl.SearchDetectedIssue(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Device":
		impl, ok := api.(r4.DeviceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Device"}
		}
		v, err := impl.SearchDevice(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceDefinition":
		impl, ok := api.(r4.DeviceDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DeviceDefinition"}
		}
		v, err := impl.SearchDeviceDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceMetric":
		impl, ok := api.(r4.DeviceMetricSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DeviceMetric"}
		}
		v, err := impl.SearchDeviceMetric(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceRequest":
		impl, ok := api.(r4.DeviceRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DeviceRequest"}
		}
		v, err := impl.SearchDeviceRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DeviceUseStatement":
		impl, ok := api.(r4.DeviceUseStatementSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DeviceUseStatement"}
		}
		v, err := impl.SearchDeviceUseStatement(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DiagnosticReport":
		impl, ok := api.(r4.DiagnosticReportSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DiagnosticReport"}
		}
		v, err := impl.SearchDiagnosticReport(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DocumentManifest":
		impl, ok := api.(r4.DocumentManifestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DocumentManifest"}
		}
		v, err := impl.SearchDocumentManifest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "DocumentReference":
		impl, ok := api.(r4.DocumentReferenceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "DocumentReference"}
		}
		v, err := impl.SearchDocumentReference(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EffectEvidenceSynthesis":
		impl, ok := api.(r4.EffectEvidenceSynthesisSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EffectEvidenceSynthesis"}
		}
		v, err := impl.SearchEffectEvidenceSynthesis(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Encounter":
		impl, ok := api.(r4.EncounterSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Encounter"}
		}
		v, err := impl.SearchEncounter(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Endpoint":
		impl, ok := api.(r4.EndpointSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Endpoint"}
		}
		v, err := impl.SearchEndpoint(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EnrollmentRequest":
		impl, ok := api.(r4.EnrollmentRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EnrollmentRequest"}
		}
		v, err := impl.SearchEnrollmentRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EnrollmentResponse":
		impl, ok := api.(r4.EnrollmentResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EnrollmentResponse"}
		}
		v, err := impl.SearchEnrollmentResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EpisodeOfCare":
		impl, ok := api.(r4.EpisodeOfCareSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EpisodeOfCare"}
		}
		v, err := impl.SearchEpisodeOfCare(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EventDefinition":
		impl, ok := api.(r4.EventDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EventDefinition"}
		}
		v, err := impl.SearchEventDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Evidence":
		impl, ok := api.(r4.EvidenceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Evidence"}
		}
		v, err := impl.SearchEvidence(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "EvidenceVariable":
		impl, ok := api.(r4.EvidenceVariableSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "EvidenceVariable"}
		}
		v, err := impl.SearchEvidenceVariable(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ExampleScenario":
		impl, ok := api.(r4.ExampleScenarioSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ExampleScenario"}
		}
		v, err := impl.SearchExampleScenario(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ExplanationOfBenefit":
		impl, ok := api.(r4.ExplanationOfBenefitSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ExplanationOfBenefit"}
		}
		v, err := impl.SearchExplanationOfBenefit(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "FamilyMemberHistory":
		impl, ok := api.(r4.FamilyMemberHistorySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "FamilyMemberHistory"}
		}
		v, err := impl.SearchFamilyMemberHistory(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Flag":
		impl, ok := api.(r4.FlagSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Flag"}
		}
		v, err := impl.SearchFlag(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Goal":
		impl, ok := api.(r4.GoalSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Goal"}
		}
		v, err := impl.SearchGoal(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "GraphDefinition":
		impl, ok := api.(r4.GraphDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "GraphDefinition"}
		}
		v, err := impl.SearchGraphDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Group":
		impl, ok := api.(r4.GroupSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Group"}
		}
		v, err := impl.SearchGroup(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "GuidanceResponse":
		impl, ok := api.(r4.GuidanceResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "GuidanceResponse"}
		}
		v, err := impl.SearchGuidanceResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "HealthcareService":
		impl, ok := api.(r4.HealthcareServiceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "HealthcareService"}
		}
		v, err := impl.SearchHealthcareService(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImagingStudy":
		impl, ok := api.(r4.ImagingStudySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ImagingStudy"}
		}
		v, err := impl.SearchImagingStudy(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Immunization":
		impl, ok := api.(r4.ImmunizationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Immunization"}
		}
		v, err := impl.SearchImmunization(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImmunizationEvaluation":
		impl, ok := api.(r4.ImmunizationEvaluationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ImmunizationEvaluation"}
		}
		v, err := impl.SearchImmunizationEvaluation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImmunizationRecommendation":
		impl, ok := api.(r4.ImmunizationRecommendationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ImmunizationRecommendation"}
		}
		v, err := impl.SearchImmunizationRecommendation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ImplementationGuide":
		impl, ok := api.(r4.ImplementationGuideSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ImplementationGuide"}
		}
		v, err := impl.SearchImplementationGuide(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "InsurancePlan":
		impl, ok := api.(r4.InsurancePlanSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "InsurancePlan"}
		}
		v, err := impl.SearchInsurancePlan(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Invoice":
		impl, ok := api.(r4.InvoiceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Invoice"}
		}
		v, err := impl.SearchInvoice(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Library":
		impl, ok := api.(r4.LibrarySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Library"}
		}
		v, err := impl.SearchLibrary(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Linkage":
		impl, ok := api.(r4.LinkageSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Linkage"}
		}
		v, err := impl.SearchLinkage(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "List":
		impl, ok := api.(r4.ListSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "List"}
		}
		v, err := impl.SearchList(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Location":
		impl, ok := api.(r4.LocationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Location"}
		}
		v, err := impl.SearchLocation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Measure":
		impl, ok := api.(r4.MeasureSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Measure"}
		}
		v, err := impl.SearchMeasure(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MeasureReport":
		impl, ok := api.(r4.MeasureReportSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MeasureReport"}
		}
		v, err := impl.SearchMeasureReport(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Media":
		impl, ok := api.(r4.MediaSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Media"}
		}
		v, err := impl.SearchMedia(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Medication":
		impl, ok := api.(r4.MedicationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Medication"}
		}
		v, err := impl.SearchMedication(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationAdministration":
		impl, ok := api.(r4.MedicationAdministrationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicationAdministration"}
		}
		v, err := impl.SearchMedicationAdministration(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationDispense":
		impl, ok := api.(r4.MedicationDispenseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicationDispense"}
		}
		v, err := impl.SearchMedicationDispense(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationKnowledge":
		impl, ok := api.(r4.MedicationKnowledgeSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicationKnowledge"}
		}
		v, err := impl.SearchMedicationKnowledge(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationRequest":
		impl, ok := api.(r4.MedicationRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicationRequest"}
		}
		v, err := impl.SearchMedicationRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicationStatement":
		impl, ok := api.(r4.MedicationStatementSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicationStatement"}
		}
		v, err := impl.SearchMedicationStatement(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProduct":
		impl, ok := api.(r4.MedicinalProductSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProduct"}
		}
		v, err := impl.SearchMedicinalProduct(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductAuthorization":
		impl, ok := api.(r4.MedicinalProductAuthorizationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductAuthorization"}
		}
		v, err := impl.SearchMedicinalProductAuthorization(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductContraindication":
		impl, ok := api.(r4.MedicinalProductContraindicationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductContraindication"}
		}
		v, err := impl.SearchMedicinalProductContraindication(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductIndication":
		impl, ok := api.(r4.MedicinalProductIndicationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductIndication"}
		}
		v, err := impl.SearchMedicinalProductIndication(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductIngredient":
		impl, ok := api.(r4.MedicinalProductIngredientSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductIngredient"}
		}
		v, err := impl.SearchMedicinalProductIngredient(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductInteraction":
		impl, ok := api.(r4.MedicinalProductInteractionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductInteraction"}
		}
		v, err := impl.SearchMedicinalProductInteraction(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductManufactured":
		impl, ok := api.(r4.MedicinalProductManufacturedSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductManufactured"}
		}
		v, err := impl.SearchMedicinalProductManufactured(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductPackaged":
		impl, ok := api.(r4.MedicinalProductPackagedSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductPackaged"}
		}
		v, err := impl.SearchMedicinalProductPackaged(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductPharmaceutical":
		impl, ok := api.(r4.MedicinalProductPharmaceuticalSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductPharmaceutical"}
		}
		v, err := impl.SearchMedicinalProductPharmaceutical(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MedicinalProductUndesirableEffect":
		impl, ok := api.(r4.MedicinalProductUndesirableEffectSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MedicinalProductUndesirableEffect"}
		}
		v, err := impl.SearchMedicinalProductUndesirableEffect(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MessageDefinition":
		impl, ok := api.(r4.MessageDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MessageDefinition"}
		}
		v, err := impl.SearchMessageDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MessageHeader":
		impl, ok := api.(r4.MessageHeaderSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MessageHeader"}
		}
		v, err := impl.SearchMessageHeader(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "MolecularSequence":
		impl, ok := api.(r4.MolecularSequenceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "MolecularSequence"}
		}
		v, err := impl.SearchMolecularSequence(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "NamingSystem":
		impl, ok := api.(r4.NamingSystemSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "NamingSystem"}
		}
		v, err := impl.SearchNamingSystem(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "NutritionOrder":
		impl, ok := api.(r4.NutritionOrderSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "NutritionOrder"}
		}
		v, err := impl.SearchNutritionOrder(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Observation":
		impl, ok := api.(r4.ObservationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Observation"}
		}
		v, err := impl.SearchObservation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ObservationDefinition":
		impl, ok := api.(r4.ObservationDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ObservationDefinition"}
		}
		v, err := impl.SearchObservationDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OperationDefinition":
		impl, ok := api.(r4.OperationDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "OperationDefinition"}
		}
		v, err := impl.SearchOperationDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OperationOutcome":
		impl, ok := api.(r4.OperationOutcomeSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "OperationOutcome"}
		}
		v, err := impl.SearchOperationOutcome(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Organization":
		impl, ok := api.(r4.OrganizationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Organization"}
		}
		v, err := impl.SearchOrganization(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "OrganizationAffiliation":
		impl, ok := api.(r4.OrganizationAffiliationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "OrganizationAffiliation"}
		}
		v, err := impl.SearchOrganizationAffiliation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Parameters":
		impl, ok := api.(r4.ParametersSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Parameters"}
		}
		v, err := impl.SearchParameters(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Patient":
		impl, ok := api.(r4.PatientSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Patient"}
		}
		v, err := impl.SearchPatient(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PaymentNotice":
		impl, ok := api.(r4.PaymentNoticeSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "PaymentNotice"}
		}
		v, err := impl.SearchPaymentNotice(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PaymentReconciliation":
		impl, ok := api.(r4.PaymentReconciliationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "PaymentReconciliation"}
		}
		v, err := impl.SearchPaymentReconciliation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Person":
		impl, ok := api.(r4.PersonSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Person"}
		}
		v, err := impl.SearchPerson(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PlanDefinition":
		impl, ok := api.(r4.PlanDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "PlanDefinition"}
		}
		v, err := impl.SearchPlanDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Practitioner":
		impl, ok := api.(r4.PractitionerSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Practitioner"}
		}
		v, err := impl.SearchPractitioner(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "PractitionerRole":
		impl, ok := api.(r4.PractitionerRoleSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "PractitionerRole"}
		}
		v, err := impl.SearchPractitionerRole(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Procedure":
		impl, ok := api.(r4.ProcedureSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Procedure"}
		}
		v, err := impl.SearchProcedure(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Provenance":
		impl, ok := api.(r4.ProvenanceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Provenance"}
		}
		v, err := impl.SearchProvenance(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Questionnaire":
		impl, ok := api.(r4.QuestionnaireSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Questionnaire"}
		}
		v, err := impl.SearchQuestionnaire(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "QuestionnaireResponse":
		impl, ok := api.(r4.QuestionnaireResponseSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "QuestionnaireResponse"}
		}
		v, err := impl.SearchQuestionnaireResponse(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RelatedPerson":
		impl, ok := api.(r4.RelatedPersonSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "RelatedPerson"}
		}
		v, err := impl.SearchRelatedPerson(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RequestGroup":
		impl, ok := api.(r4.RequestGroupSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "RequestGroup"}
		}
		v, err := impl.SearchRequestGroup(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchDefinition":
		impl, ok := api.(r4.ResearchDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ResearchDefinition"}
		}
		v, err := impl.SearchResearchDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchElementDefinition":
		impl, ok := api.(r4.ResearchElementDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ResearchElementDefinition"}
		}
		v, err := impl.SearchResearchElementDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchStudy":
		impl, ok := api.(r4.ResearchStudySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ResearchStudy"}
		}
		v, err := impl.SearchResearchStudy(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ResearchSubject":
		impl, ok := api.(r4.ResearchSubjectSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ResearchSubject"}
		}
		v, err := impl.SearchResearchSubject(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RiskAssessment":
		impl, ok := api.(r4.RiskAssessmentSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "RiskAssessment"}
		}
		v, err := impl.SearchRiskAssessment(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "RiskEvidenceSynthesis":
		impl, ok := api.(r4.RiskEvidenceSynthesisSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "RiskEvidenceSynthesis"}
		}
		v, err := impl.SearchRiskEvidenceSynthesis(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Schedule":
		impl, ok := api.(r4.ScheduleSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Schedule"}
		}
		v, err := impl.SearchSchedule(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SearchParameter":
		impl, ok := api.(r4.SearchParameterSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SearchParameter"}
		}
		v, err := impl.SearchSearchParameter(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ServiceRequest":
		impl, ok := api.(r4.ServiceRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ServiceRequest"}
		}
		v, err := impl.SearchServiceRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Slot":
		impl, ok := api.(r4.SlotSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Slot"}
		}
		v, err := impl.SearchSlot(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Specimen":
		impl, ok := api.(r4.SpecimenSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Specimen"}
		}
		v, err := impl.SearchSpecimen(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SpecimenDefinition":
		impl, ok := api.(r4.SpecimenDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SpecimenDefinition"}
		}
		v, err := impl.SearchSpecimenDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "StructureDefinition":
		impl, ok := api.(r4.StructureDefinitionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "StructureDefinition"}
		}
		v, err := impl.SearchStructureDefinition(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "StructureMap":
		impl, ok := api.(r4.StructureMapSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "StructureMap"}
		}
		v, err := impl.SearchStructureMap(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Subscription":
		impl, ok := api.(r4.SubscriptionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Subscription"}
		}
		v, err := impl.SearchSubscription(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Substance":
		impl, ok := api.(r4.SubstanceSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Substance"}
		}
		v, err := impl.SearchSubstance(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceNucleicAcid":
		impl, ok := api.(r4.SubstanceNucleicAcidSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceNucleicAcid"}
		}
		v, err := impl.SearchSubstanceNucleicAcid(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstancePolymer":
		impl, ok := api.(r4.SubstancePolymerSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstancePolymer"}
		}
		v, err := impl.SearchSubstancePolymer(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceProtein":
		impl, ok := api.(r4.SubstanceProteinSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceProtein"}
		}
		v, err := impl.SearchSubstanceProtein(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceReferenceInformation":
		impl, ok := api.(r4.SubstanceReferenceInformationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceReferenceInformation"}
		}
		v, err := impl.SearchSubstanceReferenceInformation(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceSourceMaterial":
		impl, ok := api.(r4.SubstanceSourceMaterialSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceSourceMaterial"}
		}
		v, err := impl.SearchSubstanceSourceMaterial(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SubstanceSpecification":
		impl, ok := api.(r4.SubstanceSpecificationSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SubstanceSpecification"}
		}
		v, err := impl.SearchSubstanceSpecification(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SupplyDelivery":
		impl, ok := api.(r4.SupplyDeliverySearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SupplyDelivery"}
		}
		v, err := impl.SearchSupplyDelivery(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "SupplyRequest":
		impl, ok := api.(r4.SupplyRequestSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "SupplyRequest"}
		}
		v, err := impl.SearchSupplyRequest(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "Task":
		impl, ok := api.(r4.TaskSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "Task"}
		}
		v, err := impl.SearchTask(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TerminologyCapabilities":
		impl, ok := api.(r4.TerminologyCapabilitiesSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "TerminologyCapabilities"}
		}
		v, err := impl.SearchTerminologyCapabilities(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TestReport":
		impl, ok := api.(r4.TestReportSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "TestReport"}
		}
		v, err := impl.SearchTestReport(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "TestScript":
		impl, ok := api.(r4.TestScriptSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "TestScript"}
		}
		v, err := impl.SearchTestScript(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "ValueSet":
		impl, ok := api.(r4.ValueSetSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "ValueSet"}
		}
		v, err := impl.SearchValueSet(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "VerificationResult":
		impl, ok := api.(r4.VerificationResultSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "VerificationResult"}
		}
		v, err := impl.SearchVerificationResult(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	case "VisionPrescription":
		impl, ok := api.(r4.VisionPrescriptionSearch)
		if !ok {
			return nil, dispatch.NotImplementedError{Interaction: "Search", ResourceType: "VisionPrescription"}
		}
		v, err := impl.SearchVisionPrescription(ctx, parameters)
		if err != nil {
			return nil, err
		}
		r := make([]any, 0, len(v))
		for _, e := range v {
			r = append(r, e)
		}
		return r, nil
	default:
		return nil, dispatch.UnknownResourceError{ResourceType: resourceType}
	}
}
