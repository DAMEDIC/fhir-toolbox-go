package concreteR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	search "fhir-toolbox/capabilities/search"
)

func (w InternalWrapper) SearchCapabilitiesAccount() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Account")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAccount(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Account", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesActivityDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ActivityDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchActivityDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ActivityDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesAdverseEvent() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("AdverseEvent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAdverseEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "AdverseEvent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesAllergyIntolerance() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("AllergyIntolerance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAllergyIntolerance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "AllergyIntolerance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesAppointment() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Appointment")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAppointment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Appointment", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesAppointmentResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("AppointmentResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAppointmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "AppointmentResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesAuditEvent() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("AuditEvent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchAuditEvent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "AuditEvent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesBasic() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Basic")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchBasic(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Basic", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesBinary() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Binary")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchBinary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Binary", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesBiologicallyDerivedProduct() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("BiologicallyDerivedProduct")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchBiologicallyDerivedProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "BiologicallyDerivedProduct", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesBodyStructure() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("BodyStructure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchBodyStructure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "BodyStructure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesBundle() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Bundle")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchBundle(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Bundle", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCapabilityStatement() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CapabilityStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCapabilityStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CapabilityStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCarePlan() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CarePlan")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCarePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CarePlan", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCareTeam() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CareTeam")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCareTeam(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CareTeam", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCatalogEntry() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CatalogEntry")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCatalogEntry(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CatalogEntry", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesChargeItem() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ChargeItem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchChargeItem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ChargeItem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesChargeItemDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ChargeItemDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchChargeItemDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ChargeItemDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesClaim() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Claim")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchClaim(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Claim", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesClaimResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ClaimResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchClaimResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ClaimResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesClinicalImpression() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ClinicalImpression")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchClinicalImpression(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ClinicalImpression", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCodeSystem() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CodeSystem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCodeSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CodeSystem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCommunication() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Communication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCommunication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Communication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCommunicationRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CommunicationRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCommunicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CommunicationRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCompartmentDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CompartmentDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCompartmentDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CompartmentDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesComposition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Composition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchComposition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Composition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesConceptMap() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ConceptMap")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchConceptMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ConceptMap", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCondition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Condition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCondition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Condition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesConsent() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Consent")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchConsent(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Consent", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesContract() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Contract")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchContract(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Contract", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCoverage() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Coverage")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCoverage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Coverage", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCoverageEligibilityRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CoverageEligibilityRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCoverageEligibilityRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CoverageEligibilityRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesCoverageEligibilityResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("CoverageEligibilityResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchCoverageEligibilityResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "CoverageEligibilityResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDetectedIssue() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DetectedIssue")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDetectedIssue(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DetectedIssue", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDevice() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Device")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDevice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Device", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDeviceDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DeviceDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDeviceDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DeviceDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDeviceMetric() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DeviceMetric")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDeviceMetric(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DeviceMetric", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDeviceRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DeviceRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDeviceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DeviceRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDeviceUseStatement() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DeviceUseStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDeviceUseStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DeviceUseStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDiagnosticReport() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DiagnosticReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDiagnosticReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DiagnosticReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDocumentManifest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DocumentManifest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDocumentManifest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DocumentManifest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesDocumentReference() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("DocumentReference")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchDocumentReference(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "DocumentReference", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEffectEvidenceSynthesis() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EffectEvidenceSynthesis")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEffectEvidenceSynthesis(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EffectEvidenceSynthesis", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEncounter() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Encounter")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEncounter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Encounter", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEndpoint() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Endpoint")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEndpoint(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Endpoint", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEnrollmentRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EnrollmentRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEnrollmentRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EnrollmentRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEnrollmentResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EnrollmentResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEnrollmentResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EnrollmentResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEpisodeOfCare() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EpisodeOfCare")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEpisodeOfCare(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EpisodeOfCare", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEventDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EventDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEventDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EventDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEvidence() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Evidence")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEvidence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Evidence", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesEvidenceVariable() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("EvidenceVariable")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchEvidenceVariable(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "EvidenceVariable", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesExampleScenario() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ExampleScenario")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchExampleScenario(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ExampleScenario", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesExplanationOfBenefit() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ExplanationOfBenefit")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchExplanationOfBenefit(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ExplanationOfBenefit", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesFamilyMemberHistory() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("FamilyMemberHistory")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchFamilyMemberHistory(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "FamilyMemberHistory", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesFlag() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Flag")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchFlag(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Flag", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesGoal() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Goal")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchGoal(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Goal", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesGraphDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("GraphDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchGraphDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "GraphDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesGroup() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Group")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchGroup(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Group", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesGuidanceResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("GuidanceResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchGuidanceResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "GuidanceResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesHealthcareService() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("HealthcareService")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchHealthcareService(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "HealthcareService", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesImagingStudy() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ImagingStudy")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchImagingStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ImagingStudy", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesImmunization() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Immunization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchImmunization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Immunization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesImmunizationEvaluation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ImmunizationEvaluation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchImmunizationEvaluation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ImmunizationEvaluation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesImmunizationRecommendation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ImmunizationRecommendation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchImmunizationRecommendation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ImmunizationRecommendation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesImplementationGuide() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ImplementationGuide")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchImplementationGuide(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ImplementationGuide", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesInsurancePlan() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("InsurancePlan")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchInsurancePlan(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "InsurancePlan", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesInvoice() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Invoice")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchInvoice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Invoice", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesLibrary() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Library")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchLibrary(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Library", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesLinkage() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Linkage")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchLinkage(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Linkage", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesList() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("List")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchList(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "List", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesLocation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Location")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchLocation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Location", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMeasure() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Measure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMeasure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Measure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMeasureReport() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MeasureReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMeasureReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MeasureReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedia() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Media")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedia(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Media", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedication() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Medication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Medication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicationAdministration() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicationAdministration")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicationAdministration(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicationAdministration", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicationDispense() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicationDispense")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicationDispense(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicationDispense", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicationKnowledge() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicationKnowledge")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicationKnowledge(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicationKnowledge", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicationRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicationRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicationRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicationRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicationStatement() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicationStatement")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicationStatement(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicationStatement", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProduct() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProduct")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProduct(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProduct", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductAuthorization() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductAuthorization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductAuthorization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductAuthorization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductContraindication() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductContraindication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductContraindication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductContraindication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductIndication() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductIndication")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductIndication(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductIndication", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductIngredient() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductIngredient")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductIngredient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductIngredient", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductInteraction() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductInteraction")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductInteraction(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductInteraction", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductManufactured() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductManufactured")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductManufactured(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductManufactured", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductPackaged() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductPackaged")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductPackaged(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductPackaged", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductPharmaceutical() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductPharmaceutical")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductPharmaceutical(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductPharmaceutical", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMedicinalProductUndesirableEffect() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MedicinalProductUndesirableEffect")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMedicinalProductUndesirableEffect(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MedicinalProductUndesirableEffect", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMessageDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MessageDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMessageDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MessageDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMessageHeader() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MessageHeader")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMessageHeader(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MessageHeader", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesMolecularSequence() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("MolecularSequence")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchMolecularSequence(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "MolecularSequence", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesNamingSystem() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("NamingSystem")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchNamingSystem(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "NamingSystem", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesNutritionOrder() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("NutritionOrder")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchNutritionOrder(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "NutritionOrder", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesObservation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Observation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Observation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesObservationDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ObservationDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchObservationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ObservationDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesOperationDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("OperationDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchOperationDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "OperationDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesOperationOutcome() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("OperationOutcome")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchOperationOutcome(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "OperationOutcome", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesOrganization() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Organization")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchOrganization(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Organization", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesOrganizationAffiliation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("OrganizationAffiliation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchOrganizationAffiliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "OrganizationAffiliation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesParameters() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Parameters")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchParameters(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Parameters", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPatient() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Patient")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Patient", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPaymentNotice() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("PaymentNotice")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPaymentNotice(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "PaymentNotice", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPaymentReconciliation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("PaymentReconciliation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPaymentReconciliation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "PaymentReconciliation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPerson() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Person")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Person", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPlanDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("PlanDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPlanDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "PlanDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPractitioner() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Practitioner")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPractitioner(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Practitioner", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesPractitionerRole() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("PractitionerRole")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchPractitionerRole(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "PractitionerRole", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesProcedure() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Procedure")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchProcedure(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Procedure", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesProvenance() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Provenance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchProvenance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Provenance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesQuestionnaire() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Questionnaire")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchQuestionnaire(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Questionnaire", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesQuestionnaireResponse() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("QuestionnaireResponse")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchQuestionnaireResponse(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "QuestionnaireResponse", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesRelatedPerson() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("RelatedPerson")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchRelatedPerson(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "RelatedPerson", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesRequestGroup() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("RequestGroup")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchRequestGroup(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "RequestGroup", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesResearchDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ResearchDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchResearchDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ResearchDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesResearchElementDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ResearchElementDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchResearchElementDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ResearchElementDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesResearchStudy() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ResearchStudy")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchResearchStudy(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ResearchStudy", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesResearchSubject() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ResearchSubject")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchResearchSubject(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ResearchSubject", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesRiskAssessment() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("RiskAssessment")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchRiskAssessment(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "RiskAssessment", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesRiskEvidenceSynthesis() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("RiskEvidenceSynthesis")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchRiskEvidenceSynthesis(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "RiskEvidenceSynthesis", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSchedule() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Schedule")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSchedule(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Schedule", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSearchParameter() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SearchParameter")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSearchParameter(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SearchParameter", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesServiceRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ServiceRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchServiceRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ServiceRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSlot() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Slot")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSlot(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Slot", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSpecimen() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Specimen")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSpecimen(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Specimen", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSpecimenDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SpecimenDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSpecimenDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SpecimenDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesStructureDefinition() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("StructureDefinition")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchStructureDefinition(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "StructureDefinition", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesStructureMap() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("StructureMap")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchStructureMap(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "StructureMap", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubscription() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Subscription")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubscription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Subscription", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstance() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Substance")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstance(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Substance", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstanceNucleicAcid() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstanceNucleicAcid")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstanceNucleicAcid(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstanceNucleicAcid", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstancePolymer() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstancePolymer")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstancePolymer(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstancePolymer", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstanceProtein() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstanceProtein")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstanceProtein(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstanceProtein", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstanceReferenceInformation() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstanceReferenceInformation")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstanceReferenceInformation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstanceReferenceInformation", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstanceSourceMaterial() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstanceSourceMaterial")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstanceSourceMaterial(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstanceSourceMaterial", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSubstanceSpecification() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SubstanceSpecification")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSubstanceSpecification(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SubstanceSpecification", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSupplyDelivery() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SupplyDelivery")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSupplyDelivery(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SupplyDelivery", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesSupplyRequest() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("SupplyRequest")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchSupplyRequest(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "SupplyRequest", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesTask() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("Task")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchTask(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "Task", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesTerminologyCapabilities() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("TerminologyCapabilities")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchTerminologyCapabilities(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "TerminologyCapabilities", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesTestReport() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("TestReport")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchTestReport(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "TestReport", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesTestScript() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("TestScript")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchTestScript(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "TestScript", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesValueSet() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("ValueSet")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchValueSet(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "ValueSet", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesVerificationResult() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("VerificationResult")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchVerificationResult(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "VerificationResult", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
func (w InternalWrapper) SearchCapabilitiesVisionPrescription() search.Capabilities {
	c, err := w.GenericAPI.SearchCapabilities("VisionPrescription")
	if err != nil {
		return search.Capabilities{}
	}
	return c
}
func (w InternalWrapper) SearchVisionPrescription(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	v, err := w.GenericAPI.Search(ctx, "VisionPrescription", options)
	if err != nil {
		return search.Result{}, err
	}
	return v, nil
}
