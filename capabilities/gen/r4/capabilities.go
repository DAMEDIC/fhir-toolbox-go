package capabilitiesR4

import (
	capabilities "fhir-toolbox/capabilities"
	search "fhir-toolbox/capabilities/search"
)

func AllCapabilities(api any) capabilities.Capabilities {
	read := []string{}
	search := map[string]search.Capabilities{}
	if _, ok := api.(AccountRead); ok {
		read = append(read, "Account")
	}
	if c, ok := api.(AccountSearch); ok {
		search["Account"] = c.SearchCapabilitiesAccount()
	}
	if _, ok := api.(ActivityDefinitionRead); ok {
		read = append(read, "ActivityDefinition")
	}
	if c, ok := api.(ActivityDefinitionSearch); ok {
		search["ActivityDefinition"] = c.SearchCapabilitiesActivityDefinition()
	}
	if _, ok := api.(AdverseEventRead); ok {
		read = append(read, "AdverseEvent")
	}
	if c, ok := api.(AdverseEventSearch); ok {
		search["AdverseEvent"] = c.SearchCapabilitiesAdverseEvent()
	}
	if _, ok := api.(AllergyIntoleranceRead); ok {
		read = append(read, "AllergyIntolerance")
	}
	if c, ok := api.(AllergyIntoleranceSearch); ok {
		search["AllergyIntolerance"] = c.SearchCapabilitiesAllergyIntolerance()
	}
	if _, ok := api.(AppointmentRead); ok {
		read = append(read, "Appointment")
	}
	if c, ok := api.(AppointmentSearch); ok {
		search["Appointment"] = c.SearchCapabilitiesAppointment()
	}
	if _, ok := api.(AppointmentResponseRead); ok {
		read = append(read, "AppointmentResponse")
	}
	if c, ok := api.(AppointmentResponseSearch); ok {
		search["AppointmentResponse"] = c.SearchCapabilitiesAppointmentResponse()
	}
	if _, ok := api.(AuditEventRead); ok {
		read = append(read, "AuditEvent")
	}
	if c, ok := api.(AuditEventSearch); ok {
		search["AuditEvent"] = c.SearchCapabilitiesAuditEvent()
	}
	if _, ok := api.(BasicRead); ok {
		read = append(read, "Basic")
	}
	if c, ok := api.(BasicSearch); ok {
		search["Basic"] = c.SearchCapabilitiesBasic()
	}
	if _, ok := api.(BinaryRead); ok {
		read = append(read, "Binary")
	}
	if c, ok := api.(BinarySearch); ok {
		search["Binary"] = c.SearchCapabilitiesBinary()
	}
	if _, ok := api.(BiologicallyDerivedProductRead); ok {
		read = append(read, "BiologicallyDerivedProduct")
	}
	if c, ok := api.(BiologicallyDerivedProductSearch); ok {
		search["BiologicallyDerivedProduct"] = c.SearchCapabilitiesBiologicallyDerivedProduct()
	}
	if _, ok := api.(BodyStructureRead); ok {
		read = append(read, "BodyStructure")
	}
	if c, ok := api.(BodyStructureSearch); ok {
		search["BodyStructure"] = c.SearchCapabilitiesBodyStructure()
	}
	if _, ok := api.(BundleRead); ok {
		read = append(read, "Bundle")
	}
	if c, ok := api.(BundleSearch); ok {
		search["Bundle"] = c.SearchCapabilitiesBundle()
	}
	if _, ok := api.(CapabilityStatementRead); ok {
		read = append(read, "CapabilityStatement")
	}
	if c, ok := api.(CapabilityStatementSearch); ok {
		search["CapabilityStatement"] = c.SearchCapabilitiesCapabilityStatement()
	}
	if _, ok := api.(CarePlanRead); ok {
		read = append(read, "CarePlan")
	}
	if c, ok := api.(CarePlanSearch); ok {
		search["CarePlan"] = c.SearchCapabilitiesCarePlan()
	}
	if _, ok := api.(CareTeamRead); ok {
		read = append(read, "CareTeam")
	}
	if c, ok := api.(CareTeamSearch); ok {
		search["CareTeam"] = c.SearchCapabilitiesCareTeam()
	}
	if _, ok := api.(CatalogEntryRead); ok {
		read = append(read, "CatalogEntry")
	}
	if c, ok := api.(CatalogEntrySearch); ok {
		search["CatalogEntry"] = c.SearchCapabilitiesCatalogEntry()
	}
	if _, ok := api.(ChargeItemRead); ok {
		read = append(read, "ChargeItem")
	}
	if c, ok := api.(ChargeItemSearch); ok {
		search["ChargeItem"] = c.SearchCapabilitiesChargeItem()
	}
	if _, ok := api.(ChargeItemDefinitionRead); ok {
		read = append(read, "ChargeItemDefinition")
	}
	if c, ok := api.(ChargeItemDefinitionSearch); ok {
		search["ChargeItemDefinition"] = c.SearchCapabilitiesChargeItemDefinition()
	}
	if _, ok := api.(ClaimRead); ok {
		read = append(read, "Claim")
	}
	if c, ok := api.(ClaimSearch); ok {
		search["Claim"] = c.SearchCapabilitiesClaim()
	}
	if _, ok := api.(ClaimResponseRead); ok {
		read = append(read, "ClaimResponse")
	}
	if c, ok := api.(ClaimResponseSearch); ok {
		search["ClaimResponse"] = c.SearchCapabilitiesClaimResponse()
	}
	if _, ok := api.(ClinicalImpressionRead); ok {
		read = append(read, "ClinicalImpression")
	}
	if c, ok := api.(ClinicalImpressionSearch); ok {
		search["ClinicalImpression"] = c.SearchCapabilitiesClinicalImpression()
	}
	if _, ok := api.(CodeSystemRead); ok {
		read = append(read, "CodeSystem")
	}
	if c, ok := api.(CodeSystemSearch); ok {
		search["CodeSystem"] = c.SearchCapabilitiesCodeSystem()
	}
	if _, ok := api.(CommunicationRead); ok {
		read = append(read, "Communication")
	}
	if c, ok := api.(CommunicationSearch); ok {
		search["Communication"] = c.SearchCapabilitiesCommunication()
	}
	if _, ok := api.(CommunicationRequestRead); ok {
		read = append(read, "CommunicationRequest")
	}
	if c, ok := api.(CommunicationRequestSearch); ok {
		search["CommunicationRequest"] = c.SearchCapabilitiesCommunicationRequest()
	}
	if _, ok := api.(CompartmentDefinitionRead); ok {
		read = append(read, "CompartmentDefinition")
	}
	if c, ok := api.(CompartmentDefinitionSearch); ok {
		search["CompartmentDefinition"] = c.SearchCapabilitiesCompartmentDefinition()
	}
	if _, ok := api.(CompositionRead); ok {
		read = append(read, "Composition")
	}
	if c, ok := api.(CompositionSearch); ok {
		search["Composition"] = c.SearchCapabilitiesComposition()
	}
	if _, ok := api.(ConceptMapRead); ok {
		read = append(read, "ConceptMap")
	}
	if c, ok := api.(ConceptMapSearch); ok {
		search["ConceptMap"] = c.SearchCapabilitiesConceptMap()
	}
	if _, ok := api.(ConditionRead); ok {
		read = append(read, "Condition")
	}
	if c, ok := api.(ConditionSearch); ok {
		search["Condition"] = c.SearchCapabilitiesCondition()
	}
	if _, ok := api.(ConsentRead); ok {
		read = append(read, "Consent")
	}
	if c, ok := api.(ConsentSearch); ok {
		search["Consent"] = c.SearchCapabilitiesConsent()
	}
	if _, ok := api.(ContractRead); ok {
		read = append(read, "Contract")
	}
	if c, ok := api.(ContractSearch); ok {
		search["Contract"] = c.SearchCapabilitiesContract()
	}
	if _, ok := api.(CoverageRead); ok {
		read = append(read, "Coverage")
	}
	if c, ok := api.(CoverageSearch); ok {
		search["Coverage"] = c.SearchCapabilitiesCoverage()
	}
	if _, ok := api.(CoverageEligibilityRequestRead); ok {
		read = append(read, "CoverageEligibilityRequest")
	}
	if c, ok := api.(CoverageEligibilityRequestSearch); ok {
		search["CoverageEligibilityRequest"] = c.SearchCapabilitiesCoverageEligibilityRequest()
	}
	if _, ok := api.(CoverageEligibilityResponseRead); ok {
		read = append(read, "CoverageEligibilityResponse")
	}
	if c, ok := api.(CoverageEligibilityResponseSearch); ok {
		search["CoverageEligibilityResponse"] = c.SearchCapabilitiesCoverageEligibilityResponse()
	}
	if _, ok := api.(DetectedIssueRead); ok {
		read = append(read, "DetectedIssue")
	}
	if c, ok := api.(DetectedIssueSearch); ok {
		search["DetectedIssue"] = c.SearchCapabilitiesDetectedIssue()
	}
	if _, ok := api.(DeviceRead); ok {
		read = append(read, "Device")
	}
	if c, ok := api.(DeviceSearch); ok {
		search["Device"] = c.SearchCapabilitiesDevice()
	}
	if _, ok := api.(DeviceDefinitionRead); ok {
		read = append(read, "DeviceDefinition")
	}
	if c, ok := api.(DeviceDefinitionSearch); ok {
		search["DeviceDefinition"] = c.SearchCapabilitiesDeviceDefinition()
	}
	if _, ok := api.(DeviceMetricRead); ok {
		read = append(read, "DeviceMetric")
	}
	if c, ok := api.(DeviceMetricSearch); ok {
		search["DeviceMetric"] = c.SearchCapabilitiesDeviceMetric()
	}
	if _, ok := api.(DeviceRequestRead); ok {
		read = append(read, "DeviceRequest")
	}
	if c, ok := api.(DeviceRequestSearch); ok {
		search["DeviceRequest"] = c.SearchCapabilitiesDeviceRequest()
	}
	if _, ok := api.(DeviceUseStatementRead); ok {
		read = append(read, "DeviceUseStatement")
	}
	if c, ok := api.(DeviceUseStatementSearch); ok {
		search["DeviceUseStatement"] = c.SearchCapabilitiesDeviceUseStatement()
	}
	if _, ok := api.(DiagnosticReportRead); ok {
		read = append(read, "DiagnosticReport")
	}
	if c, ok := api.(DiagnosticReportSearch); ok {
		search["DiagnosticReport"] = c.SearchCapabilitiesDiagnosticReport()
	}
	if _, ok := api.(DocumentManifestRead); ok {
		read = append(read, "DocumentManifest")
	}
	if c, ok := api.(DocumentManifestSearch); ok {
		search["DocumentManifest"] = c.SearchCapabilitiesDocumentManifest()
	}
	if _, ok := api.(DocumentReferenceRead); ok {
		read = append(read, "DocumentReference")
	}
	if c, ok := api.(DocumentReferenceSearch); ok {
		search["DocumentReference"] = c.SearchCapabilitiesDocumentReference()
	}
	if _, ok := api.(EffectEvidenceSynthesisRead); ok {
		read = append(read, "EffectEvidenceSynthesis")
	}
	if c, ok := api.(EffectEvidenceSynthesisSearch); ok {
		search["EffectEvidenceSynthesis"] = c.SearchCapabilitiesEffectEvidenceSynthesis()
	}
	if _, ok := api.(EncounterRead); ok {
		read = append(read, "Encounter")
	}
	if c, ok := api.(EncounterSearch); ok {
		search["Encounter"] = c.SearchCapabilitiesEncounter()
	}
	if _, ok := api.(EndpointRead); ok {
		read = append(read, "Endpoint")
	}
	if c, ok := api.(EndpointSearch); ok {
		search["Endpoint"] = c.SearchCapabilitiesEndpoint()
	}
	if _, ok := api.(EnrollmentRequestRead); ok {
		read = append(read, "EnrollmentRequest")
	}
	if c, ok := api.(EnrollmentRequestSearch); ok {
		search["EnrollmentRequest"] = c.SearchCapabilitiesEnrollmentRequest()
	}
	if _, ok := api.(EnrollmentResponseRead); ok {
		read = append(read, "EnrollmentResponse")
	}
	if c, ok := api.(EnrollmentResponseSearch); ok {
		search["EnrollmentResponse"] = c.SearchCapabilitiesEnrollmentResponse()
	}
	if _, ok := api.(EpisodeOfCareRead); ok {
		read = append(read, "EpisodeOfCare")
	}
	if c, ok := api.(EpisodeOfCareSearch); ok {
		search["EpisodeOfCare"] = c.SearchCapabilitiesEpisodeOfCare()
	}
	if _, ok := api.(EventDefinitionRead); ok {
		read = append(read, "EventDefinition")
	}
	if c, ok := api.(EventDefinitionSearch); ok {
		search["EventDefinition"] = c.SearchCapabilitiesEventDefinition()
	}
	if _, ok := api.(EvidenceRead); ok {
		read = append(read, "Evidence")
	}
	if c, ok := api.(EvidenceSearch); ok {
		search["Evidence"] = c.SearchCapabilitiesEvidence()
	}
	if _, ok := api.(EvidenceVariableRead); ok {
		read = append(read, "EvidenceVariable")
	}
	if c, ok := api.(EvidenceVariableSearch); ok {
		search["EvidenceVariable"] = c.SearchCapabilitiesEvidenceVariable()
	}
	if _, ok := api.(ExampleScenarioRead); ok {
		read = append(read, "ExampleScenario")
	}
	if c, ok := api.(ExampleScenarioSearch); ok {
		search["ExampleScenario"] = c.SearchCapabilitiesExampleScenario()
	}
	if _, ok := api.(ExplanationOfBenefitRead); ok {
		read = append(read, "ExplanationOfBenefit")
	}
	if c, ok := api.(ExplanationOfBenefitSearch); ok {
		search["ExplanationOfBenefit"] = c.SearchCapabilitiesExplanationOfBenefit()
	}
	if _, ok := api.(FamilyMemberHistoryRead); ok {
		read = append(read, "FamilyMemberHistory")
	}
	if c, ok := api.(FamilyMemberHistorySearch); ok {
		search["FamilyMemberHistory"] = c.SearchCapabilitiesFamilyMemberHistory()
	}
	if _, ok := api.(FlagRead); ok {
		read = append(read, "Flag")
	}
	if c, ok := api.(FlagSearch); ok {
		search["Flag"] = c.SearchCapabilitiesFlag()
	}
	if _, ok := api.(GoalRead); ok {
		read = append(read, "Goal")
	}
	if c, ok := api.(GoalSearch); ok {
		search["Goal"] = c.SearchCapabilitiesGoal()
	}
	if _, ok := api.(GraphDefinitionRead); ok {
		read = append(read, "GraphDefinition")
	}
	if c, ok := api.(GraphDefinitionSearch); ok {
		search["GraphDefinition"] = c.SearchCapabilitiesGraphDefinition()
	}
	if _, ok := api.(GroupRead); ok {
		read = append(read, "Group")
	}
	if c, ok := api.(GroupSearch); ok {
		search["Group"] = c.SearchCapabilitiesGroup()
	}
	if _, ok := api.(GuidanceResponseRead); ok {
		read = append(read, "GuidanceResponse")
	}
	if c, ok := api.(GuidanceResponseSearch); ok {
		search["GuidanceResponse"] = c.SearchCapabilitiesGuidanceResponse()
	}
	if _, ok := api.(HealthcareServiceRead); ok {
		read = append(read, "HealthcareService")
	}
	if c, ok := api.(HealthcareServiceSearch); ok {
		search["HealthcareService"] = c.SearchCapabilitiesHealthcareService()
	}
	if _, ok := api.(ImagingStudyRead); ok {
		read = append(read, "ImagingStudy")
	}
	if c, ok := api.(ImagingStudySearch); ok {
		search["ImagingStudy"] = c.SearchCapabilitiesImagingStudy()
	}
	if _, ok := api.(ImmunizationRead); ok {
		read = append(read, "Immunization")
	}
	if c, ok := api.(ImmunizationSearch); ok {
		search["Immunization"] = c.SearchCapabilitiesImmunization()
	}
	if _, ok := api.(ImmunizationEvaluationRead); ok {
		read = append(read, "ImmunizationEvaluation")
	}
	if c, ok := api.(ImmunizationEvaluationSearch); ok {
		search["ImmunizationEvaluation"] = c.SearchCapabilitiesImmunizationEvaluation()
	}
	if _, ok := api.(ImmunizationRecommendationRead); ok {
		read = append(read, "ImmunizationRecommendation")
	}
	if c, ok := api.(ImmunizationRecommendationSearch); ok {
		search["ImmunizationRecommendation"] = c.SearchCapabilitiesImmunizationRecommendation()
	}
	if _, ok := api.(ImplementationGuideRead); ok {
		read = append(read, "ImplementationGuide")
	}
	if c, ok := api.(ImplementationGuideSearch); ok {
		search["ImplementationGuide"] = c.SearchCapabilitiesImplementationGuide()
	}
	if _, ok := api.(InsurancePlanRead); ok {
		read = append(read, "InsurancePlan")
	}
	if c, ok := api.(InsurancePlanSearch); ok {
		search["InsurancePlan"] = c.SearchCapabilitiesInsurancePlan()
	}
	if _, ok := api.(InvoiceRead); ok {
		read = append(read, "Invoice")
	}
	if c, ok := api.(InvoiceSearch); ok {
		search["Invoice"] = c.SearchCapabilitiesInvoice()
	}
	if _, ok := api.(LibraryRead); ok {
		read = append(read, "Library")
	}
	if c, ok := api.(LibrarySearch); ok {
		search["Library"] = c.SearchCapabilitiesLibrary()
	}
	if _, ok := api.(LinkageRead); ok {
		read = append(read, "Linkage")
	}
	if c, ok := api.(LinkageSearch); ok {
		search["Linkage"] = c.SearchCapabilitiesLinkage()
	}
	if _, ok := api.(ListRead); ok {
		read = append(read, "List")
	}
	if c, ok := api.(ListSearch); ok {
		search["List"] = c.SearchCapabilitiesList()
	}
	if _, ok := api.(LocationRead); ok {
		read = append(read, "Location")
	}
	if c, ok := api.(LocationSearch); ok {
		search["Location"] = c.SearchCapabilitiesLocation()
	}
	if _, ok := api.(MeasureRead); ok {
		read = append(read, "Measure")
	}
	if c, ok := api.(MeasureSearch); ok {
		search["Measure"] = c.SearchCapabilitiesMeasure()
	}
	if _, ok := api.(MeasureReportRead); ok {
		read = append(read, "MeasureReport")
	}
	if c, ok := api.(MeasureReportSearch); ok {
		search["MeasureReport"] = c.SearchCapabilitiesMeasureReport()
	}
	if _, ok := api.(MediaRead); ok {
		read = append(read, "Media")
	}
	if c, ok := api.(MediaSearch); ok {
		search["Media"] = c.SearchCapabilitiesMedia()
	}
	if _, ok := api.(MedicationRead); ok {
		read = append(read, "Medication")
	}
	if c, ok := api.(MedicationSearch); ok {
		search["Medication"] = c.SearchCapabilitiesMedication()
	}
	if _, ok := api.(MedicationAdministrationRead); ok {
		read = append(read, "MedicationAdministration")
	}
	if c, ok := api.(MedicationAdministrationSearch); ok {
		search["MedicationAdministration"] = c.SearchCapabilitiesMedicationAdministration()
	}
	if _, ok := api.(MedicationDispenseRead); ok {
		read = append(read, "MedicationDispense")
	}
	if c, ok := api.(MedicationDispenseSearch); ok {
		search["MedicationDispense"] = c.SearchCapabilitiesMedicationDispense()
	}
	if _, ok := api.(MedicationKnowledgeRead); ok {
		read = append(read, "MedicationKnowledge")
	}
	if c, ok := api.(MedicationKnowledgeSearch); ok {
		search["MedicationKnowledge"] = c.SearchCapabilitiesMedicationKnowledge()
	}
	if _, ok := api.(MedicationRequestRead); ok {
		read = append(read, "MedicationRequest")
	}
	if c, ok := api.(MedicationRequestSearch); ok {
		search["MedicationRequest"] = c.SearchCapabilitiesMedicationRequest()
	}
	if _, ok := api.(MedicationStatementRead); ok {
		read = append(read, "MedicationStatement")
	}
	if c, ok := api.(MedicationStatementSearch); ok {
		search["MedicationStatement"] = c.SearchCapabilitiesMedicationStatement()
	}
	if _, ok := api.(MedicinalProductRead); ok {
		read = append(read, "MedicinalProduct")
	}
	if c, ok := api.(MedicinalProductSearch); ok {
		search["MedicinalProduct"] = c.SearchCapabilitiesMedicinalProduct()
	}
	if _, ok := api.(MedicinalProductAuthorizationRead); ok {
		read = append(read, "MedicinalProductAuthorization")
	}
	if c, ok := api.(MedicinalProductAuthorizationSearch); ok {
		search["MedicinalProductAuthorization"] = c.SearchCapabilitiesMedicinalProductAuthorization()
	}
	if _, ok := api.(MedicinalProductContraindicationRead); ok {
		read = append(read, "MedicinalProductContraindication")
	}
	if c, ok := api.(MedicinalProductContraindicationSearch); ok {
		search["MedicinalProductContraindication"] = c.SearchCapabilitiesMedicinalProductContraindication()
	}
	if _, ok := api.(MedicinalProductIndicationRead); ok {
		read = append(read, "MedicinalProductIndication")
	}
	if c, ok := api.(MedicinalProductIndicationSearch); ok {
		search["MedicinalProductIndication"] = c.SearchCapabilitiesMedicinalProductIndication()
	}
	if _, ok := api.(MedicinalProductIngredientRead); ok {
		read = append(read, "MedicinalProductIngredient")
	}
	if c, ok := api.(MedicinalProductIngredientSearch); ok {
		search["MedicinalProductIngredient"] = c.SearchCapabilitiesMedicinalProductIngredient()
	}
	if _, ok := api.(MedicinalProductInteractionRead); ok {
		read = append(read, "MedicinalProductInteraction")
	}
	if c, ok := api.(MedicinalProductInteractionSearch); ok {
		search["MedicinalProductInteraction"] = c.SearchCapabilitiesMedicinalProductInteraction()
	}
	if _, ok := api.(MedicinalProductManufacturedRead); ok {
		read = append(read, "MedicinalProductManufactured")
	}
	if c, ok := api.(MedicinalProductManufacturedSearch); ok {
		search["MedicinalProductManufactured"] = c.SearchCapabilitiesMedicinalProductManufactured()
	}
	if _, ok := api.(MedicinalProductPackagedRead); ok {
		read = append(read, "MedicinalProductPackaged")
	}
	if c, ok := api.(MedicinalProductPackagedSearch); ok {
		search["MedicinalProductPackaged"] = c.SearchCapabilitiesMedicinalProductPackaged()
	}
	if _, ok := api.(MedicinalProductPharmaceuticalRead); ok {
		read = append(read, "MedicinalProductPharmaceutical")
	}
	if c, ok := api.(MedicinalProductPharmaceuticalSearch); ok {
		search["MedicinalProductPharmaceutical"] = c.SearchCapabilitiesMedicinalProductPharmaceutical()
	}
	if _, ok := api.(MedicinalProductUndesirableEffectRead); ok {
		read = append(read, "MedicinalProductUndesirableEffect")
	}
	if c, ok := api.(MedicinalProductUndesirableEffectSearch); ok {
		search["MedicinalProductUndesirableEffect"] = c.SearchCapabilitiesMedicinalProductUndesirableEffect()
	}
	if _, ok := api.(MessageDefinitionRead); ok {
		read = append(read, "MessageDefinition")
	}
	if c, ok := api.(MessageDefinitionSearch); ok {
		search["MessageDefinition"] = c.SearchCapabilitiesMessageDefinition()
	}
	if _, ok := api.(MessageHeaderRead); ok {
		read = append(read, "MessageHeader")
	}
	if c, ok := api.(MessageHeaderSearch); ok {
		search["MessageHeader"] = c.SearchCapabilitiesMessageHeader()
	}
	if _, ok := api.(MolecularSequenceRead); ok {
		read = append(read, "MolecularSequence")
	}
	if c, ok := api.(MolecularSequenceSearch); ok {
		search["MolecularSequence"] = c.SearchCapabilitiesMolecularSequence()
	}
	if _, ok := api.(NamingSystemRead); ok {
		read = append(read, "NamingSystem")
	}
	if c, ok := api.(NamingSystemSearch); ok {
		search["NamingSystem"] = c.SearchCapabilitiesNamingSystem()
	}
	if _, ok := api.(NutritionOrderRead); ok {
		read = append(read, "NutritionOrder")
	}
	if c, ok := api.(NutritionOrderSearch); ok {
		search["NutritionOrder"] = c.SearchCapabilitiesNutritionOrder()
	}
	if _, ok := api.(ObservationRead); ok {
		read = append(read, "Observation")
	}
	if c, ok := api.(ObservationSearch); ok {
		search["Observation"] = c.SearchCapabilitiesObservation()
	}
	if _, ok := api.(ObservationDefinitionRead); ok {
		read = append(read, "ObservationDefinition")
	}
	if c, ok := api.(ObservationDefinitionSearch); ok {
		search["ObservationDefinition"] = c.SearchCapabilitiesObservationDefinition()
	}
	if _, ok := api.(OperationDefinitionRead); ok {
		read = append(read, "OperationDefinition")
	}
	if c, ok := api.(OperationDefinitionSearch); ok {
		search["OperationDefinition"] = c.SearchCapabilitiesOperationDefinition()
	}
	if _, ok := api.(OperationOutcomeRead); ok {
		read = append(read, "OperationOutcome")
	}
	if c, ok := api.(OperationOutcomeSearch); ok {
		search["OperationOutcome"] = c.SearchCapabilitiesOperationOutcome()
	}
	if _, ok := api.(OrganizationRead); ok {
		read = append(read, "Organization")
	}
	if c, ok := api.(OrganizationSearch); ok {
		search["Organization"] = c.SearchCapabilitiesOrganization()
	}
	if _, ok := api.(OrganizationAffiliationRead); ok {
		read = append(read, "OrganizationAffiliation")
	}
	if c, ok := api.(OrganizationAffiliationSearch); ok {
		search["OrganizationAffiliation"] = c.SearchCapabilitiesOrganizationAffiliation()
	}
	if _, ok := api.(ParametersRead); ok {
		read = append(read, "Parameters")
	}
	if c, ok := api.(ParametersSearch); ok {
		search["Parameters"] = c.SearchCapabilitiesParameters()
	}
	if _, ok := api.(PatientRead); ok {
		read = append(read, "Patient")
	}
	if c, ok := api.(PatientSearch); ok {
		search["Patient"] = c.SearchCapabilitiesPatient()
	}
	if _, ok := api.(PaymentNoticeRead); ok {
		read = append(read, "PaymentNotice")
	}
	if c, ok := api.(PaymentNoticeSearch); ok {
		search["PaymentNotice"] = c.SearchCapabilitiesPaymentNotice()
	}
	if _, ok := api.(PaymentReconciliationRead); ok {
		read = append(read, "PaymentReconciliation")
	}
	if c, ok := api.(PaymentReconciliationSearch); ok {
		search["PaymentReconciliation"] = c.SearchCapabilitiesPaymentReconciliation()
	}
	if _, ok := api.(PersonRead); ok {
		read = append(read, "Person")
	}
	if c, ok := api.(PersonSearch); ok {
		search["Person"] = c.SearchCapabilitiesPerson()
	}
	if _, ok := api.(PlanDefinitionRead); ok {
		read = append(read, "PlanDefinition")
	}
	if c, ok := api.(PlanDefinitionSearch); ok {
		search["PlanDefinition"] = c.SearchCapabilitiesPlanDefinition()
	}
	if _, ok := api.(PractitionerRead); ok {
		read = append(read, "Practitioner")
	}
	if c, ok := api.(PractitionerSearch); ok {
		search["Practitioner"] = c.SearchCapabilitiesPractitioner()
	}
	if _, ok := api.(PractitionerRoleRead); ok {
		read = append(read, "PractitionerRole")
	}
	if c, ok := api.(PractitionerRoleSearch); ok {
		search["PractitionerRole"] = c.SearchCapabilitiesPractitionerRole()
	}
	if _, ok := api.(ProcedureRead); ok {
		read = append(read, "Procedure")
	}
	if c, ok := api.(ProcedureSearch); ok {
		search["Procedure"] = c.SearchCapabilitiesProcedure()
	}
	if _, ok := api.(ProvenanceRead); ok {
		read = append(read, "Provenance")
	}
	if c, ok := api.(ProvenanceSearch); ok {
		search["Provenance"] = c.SearchCapabilitiesProvenance()
	}
	if _, ok := api.(QuestionnaireRead); ok {
		read = append(read, "Questionnaire")
	}
	if c, ok := api.(QuestionnaireSearch); ok {
		search["Questionnaire"] = c.SearchCapabilitiesQuestionnaire()
	}
	if _, ok := api.(QuestionnaireResponseRead); ok {
		read = append(read, "QuestionnaireResponse")
	}
	if c, ok := api.(QuestionnaireResponseSearch); ok {
		search["QuestionnaireResponse"] = c.SearchCapabilitiesQuestionnaireResponse()
	}
	if _, ok := api.(RelatedPersonRead); ok {
		read = append(read, "RelatedPerson")
	}
	if c, ok := api.(RelatedPersonSearch); ok {
		search["RelatedPerson"] = c.SearchCapabilitiesRelatedPerson()
	}
	if _, ok := api.(RequestGroupRead); ok {
		read = append(read, "RequestGroup")
	}
	if c, ok := api.(RequestGroupSearch); ok {
		search["RequestGroup"] = c.SearchCapabilitiesRequestGroup()
	}
	if _, ok := api.(ResearchDefinitionRead); ok {
		read = append(read, "ResearchDefinition")
	}
	if c, ok := api.(ResearchDefinitionSearch); ok {
		search["ResearchDefinition"] = c.SearchCapabilitiesResearchDefinition()
	}
	if _, ok := api.(ResearchElementDefinitionRead); ok {
		read = append(read, "ResearchElementDefinition")
	}
	if c, ok := api.(ResearchElementDefinitionSearch); ok {
		search["ResearchElementDefinition"] = c.SearchCapabilitiesResearchElementDefinition()
	}
	if _, ok := api.(ResearchStudyRead); ok {
		read = append(read, "ResearchStudy")
	}
	if c, ok := api.(ResearchStudySearch); ok {
		search["ResearchStudy"] = c.SearchCapabilitiesResearchStudy()
	}
	if _, ok := api.(ResearchSubjectRead); ok {
		read = append(read, "ResearchSubject")
	}
	if c, ok := api.(ResearchSubjectSearch); ok {
		search["ResearchSubject"] = c.SearchCapabilitiesResearchSubject()
	}
	if _, ok := api.(RiskAssessmentRead); ok {
		read = append(read, "RiskAssessment")
	}
	if c, ok := api.(RiskAssessmentSearch); ok {
		search["RiskAssessment"] = c.SearchCapabilitiesRiskAssessment()
	}
	if _, ok := api.(RiskEvidenceSynthesisRead); ok {
		read = append(read, "RiskEvidenceSynthesis")
	}
	if c, ok := api.(RiskEvidenceSynthesisSearch); ok {
		search["RiskEvidenceSynthesis"] = c.SearchCapabilitiesRiskEvidenceSynthesis()
	}
	if _, ok := api.(ScheduleRead); ok {
		read = append(read, "Schedule")
	}
	if c, ok := api.(ScheduleSearch); ok {
		search["Schedule"] = c.SearchCapabilitiesSchedule()
	}
	if _, ok := api.(SearchParameterRead); ok {
		read = append(read, "SearchParameter")
	}
	if c, ok := api.(SearchParameterSearch); ok {
		search["SearchParameter"] = c.SearchCapabilitiesSearchParameter()
	}
	if _, ok := api.(ServiceRequestRead); ok {
		read = append(read, "ServiceRequest")
	}
	if c, ok := api.(ServiceRequestSearch); ok {
		search["ServiceRequest"] = c.SearchCapabilitiesServiceRequest()
	}
	if _, ok := api.(SlotRead); ok {
		read = append(read, "Slot")
	}
	if c, ok := api.(SlotSearch); ok {
		search["Slot"] = c.SearchCapabilitiesSlot()
	}
	if _, ok := api.(SpecimenRead); ok {
		read = append(read, "Specimen")
	}
	if c, ok := api.(SpecimenSearch); ok {
		search["Specimen"] = c.SearchCapabilitiesSpecimen()
	}
	if _, ok := api.(SpecimenDefinitionRead); ok {
		read = append(read, "SpecimenDefinition")
	}
	if c, ok := api.(SpecimenDefinitionSearch); ok {
		search["SpecimenDefinition"] = c.SearchCapabilitiesSpecimenDefinition()
	}
	if _, ok := api.(StructureDefinitionRead); ok {
		read = append(read, "StructureDefinition")
	}
	if c, ok := api.(StructureDefinitionSearch); ok {
		search["StructureDefinition"] = c.SearchCapabilitiesStructureDefinition()
	}
	if _, ok := api.(StructureMapRead); ok {
		read = append(read, "StructureMap")
	}
	if c, ok := api.(StructureMapSearch); ok {
		search["StructureMap"] = c.SearchCapabilitiesStructureMap()
	}
	if _, ok := api.(SubscriptionRead); ok {
		read = append(read, "Subscription")
	}
	if c, ok := api.(SubscriptionSearch); ok {
		search["Subscription"] = c.SearchCapabilitiesSubscription()
	}
	if _, ok := api.(SubstanceRead); ok {
		read = append(read, "Substance")
	}
	if c, ok := api.(SubstanceSearch); ok {
		search["Substance"] = c.SearchCapabilitiesSubstance()
	}
	if _, ok := api.(SubstanceNucleicAcidRead); ok {
		read = append(read, "SubstanceNucleicAcid")
	}
	if c, ok := api.(SubstanceNucleicAcidSearch); ok {
		search["SubstanceNucleicAcid"] = c.SearchCapabilitiesSubstanceNucleicAcid()
	}
	if _, ok := api.(SubstancePolymerRead); ok {
		read = append(read, "SubstancePolymer")
	}
	if c, ok := api.(SubstancePolymerSearch); ok {
		search["SubstancePolymer"] = c.SearchCapabilitiesSubstancePolymer()
	}
	if _, ok := api.(SubstanceProteinRead); ok {
		read = append(read, "SubstanceProtein")
	}
	if c, ok := api.(SubstanceProteinSearch); ok {
		search["SubstanceProtein"] = c.SearchCapabilitiesSubstanceProtein()
	}
	if _, ok := api.(SubstanceReferenceInformationRead); ok {
		read = append(read, "SubstanceReferenceInformation")
	}
	if c, ok := api.(SubstanceReferenceInformationSearch); ok {
		search["SubstanceReferenceInformation"] = c.SearchCapabilitiesSubstanceReferenceInformation()
	}
	if _, ok := api.(SubstanceSourceMaterialRead); ok {
		read = append(read, "SubstanceSourceMaterial")
	}
	if c, ok := api.(SubstanceSourceMaterialSearch); ok {
		search["SubstanceSourceMaterial"] = c.SearchCapabilitiesSubstanceSourceMaterial()
	}
	if _, ok := api.(SubstanceSpecificationRead); ok {
		read = append(read, "SubstanceSpecification")
	}
	if c, ok := api.(SubstanceSpecificationSearch); ok {
		search["SubstanceSpecification"] = c.SearchCapabilitiesSubstanceSpecification()
	}
	if _, ok := api.(SupplyDeliveryRead); ok {
		read = append(read, "SupplyDelivery")
	}
	if c, ok := api.(SupplyDeliverySearch); ok {
		search["SupplyDelivery"] = c.SearchCapabilitiesSupplyDelivery()
	}
	if _, ok := api.(SupplyRequestRead); ok {
		read = append(read, "SupplyRequest")
	}
	if c, ok := api.(SupplyRequestSearch); ok {
		search["SupplyRequest"] = c.SearchCapabilitiesSupplyRequest()
	}
	if _, ok := api.(TaskRead); ok {
		read = append(read, "Task")
	}
	if c, ok := api.(TaskSearch); ok {
		search["Task"] = c.SearchCapabilitiesTask()
	}
	if _, ok := api.(TerminologyCapabilitiesRead); ok {
		read = append(read, "TerminologyCapabilities")
	}
	if c, ok := api.(TerminologyCapabilitiesSearch); ok {
		search["TerminologyCapabilities"] = c.SearchCapabilitiesTerminologyCapabilities()
	}
	if _, ok := api.(TestReportRead); ok {
		read = append(read, "TestReport")
	}
	if c, ok := api.(TestReportSearch); ok {
		search["TestReport"] = c.SearchCapabilitiesTestReport()
	}
	if _, ok := api.(TestScriptRead); ok {
		read = append(read, "TestScript")
	}
	if c, ok := api.(TestScriptSearch); ok {
		search["TestScript"] = c.SearchCapabilitiesTestScript()
	}
	if _, ok := api.(ValueSetRead); ok {
		read = append(read, "ValueSet")
	}
	if c, ok := api.(ValueSetSearch); ok {
		search["ValueSet"] = c.SearchCapabilitiesValueSet()
	}
	if _, ok := api.(VerificationResultRead); ok {
		read = append(read, "VerificationResult")
	}
	if c, ok := api.(VerificationResultSearch); ok {
		search["VerificationResult"] = c.SearchCapabilitiesVerificationResult()
	}
	if _, ok := api.(VisionPrescriptionRead); ok {
		read = append(read, "VisionPrescription")
	}
	if c, ok := api.(VisionPrescriptionSearch); ok {
		search["VisionPrescription"] = c.SearchCapabilitiesVisionPrescription()
	}
	return capabilities.Capabilities{
		ReadInteractions:   read,
		SearchCapabilities: search,
	}
}
