package concreteR4

import (
	"context"
	capabilities "fhir-toolbox/capabilities"
	r4 "fhir-toolbox/model/gen/r4"
)

func (w InternalWrapper) ReadAccount(ctx context.Context, id string) (r4.Account, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Account", id)
	if err != nil {
		return r4.Account{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Account)
	if !ok {
		return r4.Account{}, capabilities.InvalidResourceError{ResourceType: "Account"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadActivityDefinition(ctx context.Context, id string) (r4.ActivityDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ActivityDefinition", id)
	if err != nil {
		return r4.ActivityDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ActivityDefinition)
	if !ok {
		return r4.ActivityDefinition{}, capabilities.InvalidResourceError{ResourceType: "ActivityDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadAdverseEvent(ctx context.Context, id string) (r4.AdverseEvent, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "AdverseEvent", id)
	if err != nil {
		return r4.AdverseEvent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AdverseEvent)
	if !ok {
		return r4.AdverseEvent{}, capabilities.InvalidResourceError{ResourceType: "AdverseEvent"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadAllergyIntolerance(ctx context.Context, id string) (r4.AllergyIntolerance, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "AllergyIntolerance", id)
	if err != nil {
		return r4.AllergyIntolerance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AllergyIntolerance)
	if !ok {
		return r4.AllergyIntolerance{}, capabilities.InvalidResourceError{ResourceType: "AllergyIntolerance"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadAppointment(ctx context.Context, id string) (r4.Appointment, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Appointment", id)
	if err != nil {
		return r4.Appointment{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Appointment)
	if !ok {
		return r4.Appointment{}, capabilities.InvalidResourceError{ResourceType: "Appointment"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadAppointmentResponse(ctx context.Context, id string) (r4.AppointmentResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "AppointmentResponse", id)
	if err != nil {
		return r4.AppointmentResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AppointmentResponse)
	if !ok {
		return r4.AppointmentResponse{}, capabilities.InvalidResourceError{ResourceType: "AppointmentResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadAuditEvent(ctx context.Context, id string) (r4.AuditEvent, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "AuditEvent", id)
	if err != nil {
		return r4.AuditEvent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.AuditEvent)
	if !ok {
		return r4.AuditEvent{}, capabilities.InvalidResourceError{ResourceType: "AuditEvent"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadBasic(ctx context.Context, id string) (r4.Basic, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Basic", id)
	if err != nil {
		return r4.Basic{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Basic)
	if !ok {
		return r4.Basic{}, capabilities.InvalidResourceError{ResourceType: "Basic"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadBinary(ctx context.Context, id string) (r4.Binary, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Binary", id)
	if err != nil {
		return r4.Binary{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Binary)
	if !ok {
		return r4.Binary{}, capabilities.InvalidResourceError{ResourceType: "Binary"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadBiologicallyDerivedProduct(ctx context.Context, id string) (r4.BiologicallyDerivedProduct, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "BiologicallyDerivedProduct", id)
	if err != nil {
		return r4.BiologicallyDerivedProduct{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.BiologicallyDerivedProduct)
	if !ok {
		return r4.BiologicallyDerivedProduct{}, capabilities.InvalidResourceError{ResourceType: "BiologicallyDerivedProduct"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadBodyStructure(ctx context.Context, id string) (r4.BodyStructure, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "BodyStructure", id)
	if err != nil {
		return r4.BodyStructure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.BodyStructure)
	if !ok {
		return r4.BodyStructure{}, capabilities.InvalidResourceError{ResourceType: "BodyStructure"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadBundle(ctx context.Context, id string) (r4.Bundle, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Bundle", id)
	if err != nil {
		return r4.Bundle{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Bundle)
	if !ok {
		return r4.Bundle{}, capabilities.InvalidResourceError{ResourceType: "Bundle"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCapabilityStatement(ctx context.Context, id string) (r4.CapabilityStatement, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CapabilityStatement", id)
	if err != nil {
		return r4.CapabilityStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CapabilityStatement)
	if !ok {
		return r4.CapabilityStatement{}, capabilities.InvalidResourceError{ResourceType: "CapabilityStatement"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCarePlan(ctx context.Context, id string) (r4.CarePlan, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CarePlan", id)
	if err != nil {
		return r4.CarePlan{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CarePlan)
	if !ok {
		return r4.CarePlan{}, capabilities.InvalidResourceError{ResourceType: "CarePlan"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCareTeam(ctx context.Context, id string) (r4.CareTeam, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CareTeam", id)
	if err != nil {
		return r4.CareTeam{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CareTeam)
	if !ok {
		return r4.CareTeam{}, capabilities.InvalidResourceError{ResourceType: "CareTeam"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCatalogEntry(ctx context.Context, id string) (r4.CatalogEntry, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CatalogEntry", id)
	if err != nil {
		return r4.CatalogEntry{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CatalogEntry)
	if !ok {
		return r4.CatalogEntry{}, capabilities.InvalidResourceError{ResourceType: "CatalogEntry"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadChargeItem(ctx context.Context, id string) (r4.ChargeItem, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ChargeItem", id)
	if err != nil {
		return r4.ChargeItem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ChargeItem)
	if !ok {
		return r4.ChargeItem{}, capabilities.InvalidResourceError{ResourceType: "ChargeItem"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadChargeItemDefinition(ctx context.Context, id string) (r4.ChargeItemDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ChargeItemDefinition", id)
	if err != nil {
		return r4.ChargeItemDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ChargeItemDefinition)
	if !ok {
		return r4.ChargeItemDefinition{}, capabilities.InvalidResourceError{ResourceType: "ChargeItemDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadClaim(ctx context.Context, id string) (r4.Claim, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Claim", id)
	if err != nil {
		return r4.Claim{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Claim)
	if !ok {
		return r4.Claim{}, capabilities.InvalidResourceError{ResourceType: "Claim"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadClaimResponse(ctx context.Context, id string) (r4.ClaimResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ClaimResponse", id)
	if err != nil {
		return r4.ClaimResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ClaimResponse)
	if !ok {
		return r4.ClaimResponse{}, capabilities.InvalidResourceError{ResourceType: "ClaimResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadClinicalImpression(ctx context.Context, id string) (r4.ClinicalImpression, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ClinicalImpression", id)
	if err != nil {
		return r4.ClinicalImpression{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ClinicalImpression)
	if !ok {
		return r4.ClinicalImpression{}, capabilities.InvalidResourceError{ResourceType: "ClinicalImpression"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCodeSystem(ctx context.Context, id string) (r4.CodeSystem, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CodeSystem", id)
	if err != nil {
		return r4.CodeSystem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CodeSystem)
	if !ok {
		return r4.CodeSystem{}, capabilities.InvalidResourceError{ResourceType: "CodeSystem"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCommunication(ctx context.Context, id string) (r4.Communication, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Communication", id)
	if err != nil {
		return r4.Communication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Communication)
	if !ok {
		return r4.Communication{}, capabilities.InvalidResourceError{ResourceType: "Communication"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCommunicationRequest(ctx context.Context, id string) (r4.CommunicationRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CommunicationRequest", id)
	if err != nil {
		return r4.CommunicationRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CommunicationRequest)
	if !ok {
		return r4.CommunicationRequest{}, capabilities.InvalidResourceError{ResourceType: "CommunicationRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCompartmentDefinition(ctx context.Context, id string) (r4.CompartmentDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CompartmentDefinition", id)
	if err != nil {
		return r4.CompartmentDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CompartmentDefinition)
	if !ok {
		return r4.CompartmentDefinition{}, capabilities.InvalidResourceError{ResourceType: "CompartmentDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadComposition(ctx context.Context, id string) (r4.Composition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Composition", id)
	if err != nil {
		return r4.Composition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Composition)
	if !ok {
		return r4.Composition{}, capabilities.InvalidResourceError{ResourceType: "Composition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadConceptMap(ctx context.Context, id string) (r4.ConceptMap, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ConceptMap", id)
	if err != nil {
		return r4.ConceptMap{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ConceptMap)
	if !ok {
		return r4.ConceptMap{}, capabilities.InvalidResourceError{ResourceType: "ConceptMap"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCondition(ctx context.Context, id string) (r4.Condition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Condition", id)
	if err != nil {
		return r4.Condition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Condition)
	if !ok {
		return r4.Condition{}, capabilities.InvalidResourceError{ResourceType: "Condition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadConsent(ctx context.Context, id string) (r4.Consent, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Consent", id)
	if err != nil {
		return r4.Consent{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Consent)
	if !ok {
		return r4.Consent{}, capabilities.InvalidResourceError{ResourceType: "Consent"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadContract(ctx context.Context, id string) (r4.Contract, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Contract", id)
	if err != nil {
		return r4.Contract{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Contract)
	if !ok {
		return r4.Contract{}, capabilities.InvalidResourceError{ResourceType: "Contract"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCoverage(ctx context.Context, id string) (r4.Coverage, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Coverage", id)
	if err != nil {
		return r4.Coverage{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Coverage)
	if !ok {
		return r4.Coverage{}, capabilities.InvalidResourceError{ResourceType: "Coverage"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCoverageEligibilityRequest(ctx context.Context, id string) (r4.CoverageEligibilityRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CoverageEligibilityRequest", id)
	if err != nil {
		return r4.CoverageEligibilityRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CoverageEligibilityRequest)
	if !ok {
		return r4.CoverageEligibilityRequest{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadCoverageEligibilityResponse(ctx context.Context, id string) (r4.CoverageEligibilityResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "CoverageEligibilityResponse", id)
	if err != nil {
		return r4.CoverageEligibilityResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.CoverageEligibilityResponse)
	if !ok {
		return r4.CoverageEligibilityResponse{}, capabilities.InvalidResourceError{ResourceType: "CoverageEligibilityResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDetectedIssue(ctx context.Context, id string) (r4.DetectedIssue, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DetectedIssue", id)
	if err != nil {
		return r4.DetectedIssue{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DetectedIssue)
	if !ok {
		return r4.DetectedIssue{}, capabilities.InvalidResourceError{ResourceType: "DetectedIssue"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDevice(ctx context.Context, id string) (r4.Device, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Device", id)
	if err != nil {
		return r4.Device{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Device)
	if !ok {
		return r4.Device{}, capabilities.InvalidResourceError{ResourceType: "Device"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDeviceDefinition(ctx context.Context, id string) (r4.DeviceDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DeviceDefinition", id)
	if err != nil {
		return r4.DeviceDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceDefinition)
	if !ok {
		return r4.DeviceDefinition{}, capabilities.InvalidResourceError{ResourceType: "DeviceDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDeviceMetric(ctx context.Context, id string) (r4.DeviceMetric, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DeviceMetric", id)
	if err != nil {
		return r4.DeviceMetric{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceMetric)
	if !ok {
		return r4.DeviceMetric{}, capabilities.InvalidResourceError{ResourceType: "DeviceMetric"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDeviceRequest(ctx context.Context, id string) (r4.DeviceRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DeviceRequest", id)
	if err != nil {
		return r4.DeviceRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceRequest)
	if !ok {
		return r4.DeviceRequest{}, capabilities.InvalidResourceError{ResourceType: "DeviceRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDeviceUseStatement(ctx context.Context, id string) (r4.DeviceUseStatement, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DeviceUseStatement", id)
	if err != nil {
		return r4.DeviceUseStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DeviceUseStatement)
	if !ok {
		return r4.DeviceUseStatement{}, capabilities.InvalidResourceError{ResourceType: "DeviceUseStatement"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDiagnosticReport(ctx context.Context, id string) (r4.DiagnosticReport, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DiagnosticReport", id)
	if err != nil {
		return r4.DiagnosticReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DiagnosticReport)
	if !ok {
		return r4.DiagnosticReport{}, capabilities.InvalidResourceError{ResourceType: "DiagnosticReport"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDocumentManifest(ctx context.Context, id string) (r4.DocumentManifest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DocumentManifest", id)
	if err != nil {
		return r4.DocumentManifest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DocumentManifest)
	if !ok {
		return r4.DocumentManifest{}, capabilities.InvalidResourceError{ResourceType: "DocumentManifest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadDocumentReference(ctx context.Context, id string) (r4.DocumentReference, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "DocumentReference", id)
	if err != nil {
		return r4.DocumentReference{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.DocumentReference)
	if !ok {
		return r4.DocumentReference{}, capabilities.InvalidResourceError{ResourceType: "DocumentReference"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEffectEvidenceSynthesis(ctx context.Context, id string) (r4.EffectEvidenceSynthesis, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EffectEvidenceSynthesis", id)
	if err != nil {
		return r4.EffectEvidenceSynthesis{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EffectEvidenceSynthesis)
	if !ok {
		return r4.EffectEvidenceSynthesis{}, capabilities.InvalidResourceError{ResourceType: "EffectEvidenceSynthesis"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEncounter(ctx context.Context, id string) (r4.Encounter, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Encounter", id)
	if err != nil {
		return r4.Encounter{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Encounter)
	if !ok {
		return r4.Encounter{}, capabilities.InvalidResourceError{ResourceType: "Encounter"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEndpoint(ctx context.Context, id string) (r4.Endpoint, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Endpoint", id)
	if err != nil {
		return r4.Endpoint{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Endpoint)
	if !ok {
		return r4.Endpoint{}, capabilities.InvalidResourceError{ResourceType: "Endpoint"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEnrollmentRequest(ctx context.Context, id string) (r4.EnrollmentRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EnrollmentRequest", id)
	if err != nil {
		return r4.EnrollmentRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EnrollmentRequest)
	if !ok {
		return r4.EnrollmentRequest{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEnrollmentResponse(ctx context.Context, id string) (r4.EnrollmentResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EnrollmentResponse", id)
	if err != nil {
		return r4.EnrollmentResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EnrollmentResponse)
	if !ok {
		return r4.EnrollmentResponse{}, capabilities.InvalidResourceError{ResourceType: "EnrollmentResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEpisodeOfCare(ctx context.Context, id string) (r4.EpisodeOfCare, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EpisodeOfCare", id)
	if err != nil {
		return r4.EpisodeOfCare{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EpisodeOfCare)
	if !ok {
		return r4.EpisodeOfCare{}, capabilities.InvalidResourceError{ResourceType: "EpisodeOfCare"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEventDefinition(ctx context.Context, id string) (r4.EventDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EventDefinition", id)
	if err != nil {
		return r4.EventDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EventDefinition)
	if !ok {
		return r4.EventDefinition{}, capabilities.InvalidResourceError{ResourceType: "EventDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEvidence(ctx context.Context, id string) (r4.Evidence, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Evidence", id)
	if err != nil {
		return r4.Evidence{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Evidence)
	if !ok {
		return r4.Evidence{}, capabilities.InvalidResourceError{ResourceType: "Evidence"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadEvidenceVariable(ctx context.Context, id string) (r4.EvidenceVariable, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "EvidenceVariable", id)
	if err != nil {
		return r4.EvidenceVariable{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.EvidenceVariable)
	if !ok {
		return r4.EvidenceVariable{}, capabilities.InvalidResourceError{ResourceType: "EvidenceVariable"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadExampleScenario(ctx context.Context, id string) (r4.ExampleScenario, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ExampleScenario", id)
	if err != nil {
		return r4.ExampleScenario{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ExampleScenario)
	if !ok {
		return r4.ExampleScenario{}, capabilities.InvalidResourceError{ResourceType: "ExampleScenario"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadExplanationOfBenefit(ctx context.Context, id string) (r4.ExplanationOfBenefit, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ExplanationOfBenefit", id)
	if err != nil {
		return r4.ExplanationOfBenefit{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ExplanationOfBenefit)
	if !ok {
		return r4.ExplanationOfBenefit{}, capabilities.InvalidResourceError{ResourceType: "ExplanationOfBenefit"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadFamilyMemberHistory(ctx context.Context, id string) (r4.FamilyMemberHistory, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "FamilyMemberHistory", id)
	if err != nil {
		return r4.FamilyMemberHistory{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.FamilyMemberHistory)
	if !ok {
		return r4.FamilyMemberHistory{}, capabilities.InvalidResourceError{ResourceType: "FamilyMemberHistory"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadFlag(ctx context.Context, id string) (r4.Flag, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Flag", id)
	if err != nil {
		return r4.Flag{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Flag)
	if !ok {
		return r4.Flag{}, capabilities.InvalidResourceError{ResourceType: "Flag"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadGoal(ctx context.Context, id string) (r4.Goal, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Goal", id)
	if err != nil {
		return r4.Goal{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Goal)
	if !ok {
		return r4.Goal{}, capabilities.InvalidResourceError{ResourceType: "Goal"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadGraphDefinition(ctx context.Context, id string) (r4.GraphDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "GraphDefinition", id)
	if err != nil {
		return r4.GraphDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.GraphDefinition)
	if !ok {
		return r4.GraphDefinition{}, capabilities.InvalidResourceError{ResourceType: "GraphDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadGroup(ctx context.Context, id string) (r4.Group, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Group", id)
	if err != nil {
		return r4.Group{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Group)
	if !ok {
		return r4.Group{}, capabilities.InvalidResourceError{ResourceType: "Group"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadGuidanceResponse(ctx context.Context, id string) (r4.GuidanceResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "GuidanceResponse", id)
	if err != nil {
		return r4.GuidanceResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.GuidanceResponse)
	if !ok {
		return r4.GuidanceResponse{}, capabilities.InvalidResourceError{ResourceType: "GuidanceResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadHealthcareService(ctx context.Context, id string) (r4.HealthcareService, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "HealthcareService", id)
	if err != nil {
		return r4.HealthcareService{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.HealthcareService)
	if !ok {
		return r4.HealthcareService{}, capabilities.InvalidResourceError{ResourceType: "HealthcareService"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadImagingStudy(ctx context.Context, id string) (r4.ImagingStudy, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ImagingStudy", id)
	if err != nil {
		return r4.ImagingStudy{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImagingStudy)
	if !ok {
		return r4.ImagingStudy{}, capabilities.InvalidResourceError{ResourceType: "ImagingStudy"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadImmunization(ctx context.Context, id string) (r4.Immunization, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Immunization", id)
	if err != nil {
		return r4.Immunization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Immunization)
	if !ok {
		return r4.Immunization{}, capabilities.InvalidResourceError{ResourceType: "Immunization"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadImmunizationEvaluation(ctx context.Context, id string) (r4.ImmunizationEvaluation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ImmunizationEvaluation", id)
	if err != nil {
		return r4.ImmunizationEvaluation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImmunizationEvaluation)
	if !ok {
		return r4.ImmunizationEvaluation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationEvaluation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadImmunizationRecommendation(ctx context.Context, id string) (r4.ImmunizationRecommendation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ImmunizationRecommendation", id)
	if err != nil {
		return r4.ImmunizationRecommendation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImmunizationRecommendation)
	if !ok {
		return r4.ImmunizationRecommendation{}, capabilities.InvalidResourceError{ResourceType: "ImmunizationRecommendation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadImplementationGuide(ctx context.Context, id string) (r4.ImplementationGuide, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ImplementationGuide", id)
	if err != nil {
		return r4.ImplementationGuide{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ImplementationGuide)
	if !ok {
		return r4.ImplementationGuide{}, capabilities.InvalidResourceError{ResourceType: "ImplementationGuide"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadInsurancePlan(ctx context.Context, id string) (r4.InsurancePlan, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "InsurancePlan", id)
	if err != nil {
		return r4.InsurancePlan{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.InsurancePlan)
	if !ok {
		return r4.InsurancePlan{}, capabilities.InvalidResourceError{ResourceType: "InsurancePlan"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadInvoice(ctx context.Context, id string) (r4.Invoice, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Invoice", id)
	if err != nil {
		return r4.Invoice{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Invoice)
	if !ok {
		return r4.Invoice{}, capabilities.InvalidResourceError{ResourceType: "Invoice"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadLibrary(ctx context.Context, id string) (r4.Library, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Library", id)
	if err != nil {
		return r4.Library{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Library)
	if !ok {
		return r4.Library{}, capabilities.InvalidResourceError{ResourceType: "Library"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadLinkage(ctx context.Context, id string) (r4.Linkage, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Linkage", id)
	if err != nil {
		return r4.Linkage{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Linkage)
	if !ok {
		return r4.Linkage{}, capabilities.InvalidResourceError{ResourceType: "Linkage"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadList(ctx context.Context, id string) (r4.List, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "List", id)
	if err != nil {
		return r4.List{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.List)
	if !ok {
		return r4.List{}, capabilities.InvalidResourceError{ResourceType: "List"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadLocation(ctx context.Context, id string) (r4.Location, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Location", id)
	if err != nil {
		return r4.Location{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Location)
	if !ok {
		return r4.Location{}, capabilities.InvalidResourceError{ResourceType: "Location"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMeasure(ctx context.Context, id string) (r4.Measure, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Measure", id)
	if err != nil {
		return r4.Measure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Measure)
	if !ok {
		return r4.Measure{}, capabilities.InvalidResourceError{ResourceType: "Measure"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMeasureReport(ctx context.Context, id string) (r4.MeasureReport, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MeasureReport", id)
	if err != nil {
		return r4.MeasureReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MeasureReport)
	if !ok {
		return r4.MeasureReport{}, capabilities.InvalidResourceError{ResourceType: "MeasureReport"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedia(ctx context.Context, id string) (r4.Media, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Media", id)
	if err != nil {
		return r4.Media{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Media)
	if !ok {
		return r4.Media{}, capabilities.InvalidResourceError{ResourceType: "Media"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedication(ctx context.Context, id string) (r4.Medication, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Medication", id)
	if err != nil {
		return r4.Medication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Medication)
	if !ok {
		return r4.Medication{}, capabilities.InvalidResourceError{ResourceType: "Medication"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicationAdministration(ctx context.Context, id string) (r4.MedicationAdministration, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicationAdministration", id)
	if err != nil {
		return r4.MedicationAdministration{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationAdministration)
	if !ok {
		return r4.MedicationAdministration{}, capabilities.InvalidResourceError{ResourceType: "MedicationAdministration"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicationDispense(ctx context.Context, id string) (r4.MedicationDispense, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicationDispense", id)
	if err != nil {
		return r4.MedicationDispense{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationDispense)
	if !ok {
		return r4.MedicationDispense{}, capabilities.InvalidResourceError{ResourceType: "MedicationDispense"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicationKnowledge(ctx context.Context, id string) (r4.MedicationKnowledge, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicationKnowledge", id)
	if err != nil {
		return r4.MedicationKnowledge{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationKnowledge)
	if !ok {
		return r4.MedicationKnowledge{}, capabilities.InvalidResourceError{ResourceType: "MedicationKnowledge"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicationRequest(ctx context.Context, id string) (r4.MedicationRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicationRequest", id)
	if err != nil {
		return r4.MedicationRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationRequest)
	if !ok {
		return r4.MedicationRequest{}, capabilities.InvalidResourceError{ResourceType: "MedicationRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicationStatement(ctx context.Context, id string) (r4.MedicationStatement, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicationStatement", id)
	if err != nil {
		return r4.MedicationStatement{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicationStatement)
	if !ok {
		return r4.MedicationStatement{}, capabilities.InvalidResourceError{ResourceType: "MedicationStatement"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProduct(ctx context.Context, id string) (r4.MedicinalProduct, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProduct", id)
	if err != nil {
		return r4.MedicinalProduct{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProduct)
	if !ok {
		return r4.MedicinalProduct{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProduct"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductAuthorization(ctx context.Context, id string) (r4.MedicinalProductAuthorization, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductAuthorization", id)
	if err != nil {
		return r4.MedicinalProductAuthorization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductAuthorization)
	if !ok {
		return r4.MedicinalProductAuthorization{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductAuthorization"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductContraindication(ctx context.Context, id string) (r4.MedicinalProductContraindication, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductContraindication", id)
	if err != nil {
		return r4.MedicinalProductContraindication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductContraindication)
	if !ok {
		return r4.MedicinalProductContraindication{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductContraindication"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductIndication(ctx context.Context, id string) (r4.MedicinalProductIndication, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductIndication", id)
	if err != nil {
		return r4.MedicinalProductIndication{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductIndication)
	if !ok {
		return r4.MedicinalProductIndication{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductIndication"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductIngredient(ctx context.Context, id string) (r4.MedicinalProductIngredient, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductIngredient", id)
	if err != nil {
		return r4.MedicinalProductIngredient{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductIngredient)
	if !ok {
		return r4.MedicinalProductIngredient{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductIngredient"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductInteraction(ctx context.Context, id string) (r4.MedicinalProductInteraction, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductInteraction", id)
	if err != nil {
		return r4.MedicinalProductInteraction{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductInteraction)
	if !ok {
		return r4.MedicinalProductInteraction{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductInteraction"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductManufactured(ctx context.Context, id string) (r4.MedicinalProductManufactured, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductManufactured", id)
	if err != nil {
		return r4.MedicinalProductManufactured{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductManufactured)
	if !ok {
		return r4.MedicinalProductManufactured{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductManufactured"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductPackaged(ctx context.Context, id string) (r4.MedicinalProductPackaged, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductPackaged", id)
	if err != nil {
		return r4.MedicinalProductPackaged{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductPackaged)
	if !ok {
		return r4.MedicinalProductPackaged{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductPackaged"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductPharmaceutical(ctx context.Context, id string) (r4.MedicinalProductPharmaceutical, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductPharmaceutical", id)
	if err != nil {
		return r4.MedicinalProductPharmaceutical{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductPharmaceutical)
	if !ok {
		return r4.MedicinalProductPharmaceutical{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductPharmaceutical"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMedicinalProductUndesirableEffect(ctx context.Context, id string) (r4.MedicinalProductUndesirableEffect, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MedicinalProductUndesirableEffect", id)
	if err != nil {
		return r4.MedicinalProductUndesirableEffect{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MedicinalProductUndesirableEffect)
	if !ok {
		return r4.MedicinalProductUndesirableEffect{}, capabilities.InvalidResourceError{ResourceType: "MedicinalProductUndesirableEffect"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMessageDefinition(ctx context.Context, id string) (r4.MessageDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MessageDefinition", id)
	if err != nil {
		return r4.MessageDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MessageDefinition)
	if !ok {
		return r4.MessageDefinition{}, capabilities.InvalidResourceError{ResourceType: "MessageDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMessageHeader(ctx context.Context, id string) (r4.MessageHeader, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MessageHeader", id)
	if err != nil {
		return r4.MessageHeader{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MessageHeader)
	if !ok {
		return r4.MessageHeader{}, capabilities.InvalidResourceError{ResourceType: "MessageHeader"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadMolecularSequence(ctx context.Context, id string) (r4.MolecularSequence, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "MolecularSequence", id)
	if err != nil {
		return r4.MolecularSequence{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.MolecularSequence)
	if !ok {
		return r4.MolecularSequence{}, capabilities.InvalidResourceError{ResourceType: "MolecularSequence"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadNamingSystem(ctx context.Context, id string) (r4.NamingSystem, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "NamingSystem", id)
	if err != nil {
		return r4.NamingSystem{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.NamingSystem)
	if !ok {
		return r4.NamingSystem{}, capabilities.InvalidResourceError{ResourceType: "NamingSystem"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadNutritionOrder(ctx context.Context, id string) (r4.NutritionOrder, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "NutritionOrder", id)
	if err != nil {
		return r4.NutritionOrder{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.NutritionOrder)
	if !ok {
		return r4.NutritionOrder{}, capabilities.InvalidResourceError{ResourceType: "NutritionOrder"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Observation", id)
	if err != nil {
		return r4.Observation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Observation)
	if !ok {
		return r4.Observation{}, capabilities.InvalidResourceError{ResourceType: "Observation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadObservationDefinition(ctx context.Context, id string) (r4.ObservationDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ObservationDefinition", id)
	if err != nil {
		return r4.ObservationDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ObservationDefinition)
	if !ok {
		return r4.ObservationDefinition{}, capabilities.InvalidResourceError{ResourceType: "ObservationDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadOperationDefinition(ctx context.Context, id string) (r4.OperationDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "OperationDefinition", id)
	if err != nil {
		return r4.OperationDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OperationDefinition)
	if !ok {
		return r4.OperationDefinition{}, capabilities.InvalidResourceError{ResourceType: "OperationDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadOperationOutcome(ctx context.Context, id string) (r4.OperationOutcome, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "OperationOutcome", id)
	if err != nil {
		return r4.OperationOutcome{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OperationOutcome)
	if !ok {
		return r4.OperationOutcome{}, capabilities.InvalidResourceError{ResourceType: "OperationOutcome"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadOrganization(ctx context.Context, id string) (r4.Organization, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Organization", id)
	if err != nil {
		return r4.Organization{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Organization)
	if !ok {
		return r4.Organization{}, capabilities.InvalidResourceError{ResourceType: "Organization"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadOrganizationAffiliation(ctx context.Context, id string) (r4.OrganizationAffiliation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "OrganizationAffiliation", id)
	if err != nil {
		return r4.OrganizationAffiliation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.OrganizationAffiliation)
	if !ok {
		return r4.OrganizationAffiliation{}, capabilities.InvalidResourceError{ResourceType: "OrganizationAffiliation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadParameters(ctx context.Context, id string) (r4.Parameters, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Parameters", id)
	if err != nil {
		return r4.Parameters{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Parameters)
	if !ok {
		return r4.Parameters{}, capabilities.InvalidResourceError{ResourceType: "Parameters"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Patient", id)
	if err != nil {
		return r4.Patient{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Patient)
	if !ok {
		return r4.Patient{}, capabilities.InvalidResourceError{ResourceType: "Patient"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPaymentNotice(ctx context.Context, id string) (r4.PaymentNotice, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "PaymentNotice", id)
	if err != nil {
		return r4.PaymentNotice{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PaymentNotice)
	if !ok {
		return r4.PaymentNotice{}, capabilities.InvalidResourceError{ResourceType: "PaymentNotice"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPaymentReconciliation(ctx context.Context, id string) (r4.PaymentReconciliation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "PaymentReconciliation", id)
	if err != nil {
		return r4.PaymentReconciliation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PaymentReconciliation)
	if !ok {
		return r4.PaymentReconciliation{}, capabilities.InvalidResourceError{ResourceType: "PaymentReconciliation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPerson(ctx context.Context, id string) (r4.Person, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Person", id)
	if err != nil {
		return r4.Person{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Person)
	if !ok {
		return r4.Person{}, capabilities.InvalidResourceError{ResourceType: "Person"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPlanDefinition(ctx context.Context, id string) (r4.PlanDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "PlanDefinition", id)
	if err != nil {
		return r4.PlanDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PlanDefinition)
	if !ok {
		return r4.PlanDefinition{}, capabilities.InvalidResourceError{ResourceType: "PlanDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPractitioner(ctx context.Context, id string) (r4.Practitioner, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Practitioner", id)
	if err != nil {
		return r4.Practitioner{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Practitioner)
	if !ok {
		return r4.Practitioner{}, capabilities.InvalidResourceError{ResourceType: "Practitioner"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadPractitionerRole(ctx context.Context, id string) (r4.PractitionerRole, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "PractitionerRole", id)
	if err != nil {
		return r4.PractitionerRole{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.PractitionerRole)
	if !ok {
		return r4.PractitionerRole{}, capabilities.InvalidResourceError{ResourceType: "PractitionerRole"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadProcedure(ctx context.Context, id string) (r4.Procedure, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Procedure", id)
	if err != nil {
		return r4.Procedure{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Procedure)
	if !ok {
		return r4.Procedure{}, capabilities.InvalidResourceError{ResourceType: "Procedure"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadProvenance(ctx context.Context, id string) (r4.Provenance, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Provenance", id)
	if err != nil {
		return r4.Provenance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Provenance)
	if !ok {
		return r4.Provenance{}, capabilities.InvalidResourceError{ResourceType: "Provenance"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadQuestionnaire(ctx context.Context, id string) (r4.Questionnaire, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Questionnaire", id)
	if err != nil {
		return r4.Questionnaire{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Questionnaire)
	if !ok {
		return r4.Questionnaire{}, capabilities.InvalidResourceError{ResourceType: "Questionnaire"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadQuestionnaireResponse(ctx context.Context, id string) (r4.QuestionnaireResponse, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "QuestionnaireResponse", id)
	if err != nil {
		return r4.QuestionnaireResponse{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.QuestionnaireResponse)
	if !ok {
		return r4.QuestionnaireResponse{}, capabilities.InvalidResourceError{ResourceType: "QuestionnaireResponse"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadRelatedPerson(ctx context.Context, id string) (r4.RelatedPerson, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "RelatedPerson", id)
	if err != nil {
		return r4.RelatedPerson{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RelatedPerson)
	if !ok {
		return r4.RelatedPerson{}, capabilities.InvalidResourceError{ResourceType: "RelatedPerson"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadRequestGroup(ctx context.Context, id string) (r4.RequestGroup, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "RequestGroup", id)
	if err != nil {
		return r4.RequestGroup{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RequestGroup)
	if !ok {
		return r4.RequestGroup{}, capabilities.InvalidResourceError{ResourceType: "RequestGroup"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadResearchDefinition(ctx context.Context, id string) (r4.ResearchDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ResearchDefinition", id)
	if err != nil {
		return r4.ResearchDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchDefinition)
	if !ok {
		return r4.ResearchDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadResearchElementDefinition(ctx context.Context, id string) (r4.ResearchElementDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ResearchElementDefinition", id)
	if err != nil {
		return r4.ResearchElementDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchElementDefinition)
	if !ok {
		return r4.ResearchElementDefinition{}, capabilities.InvalidResourceError{ResourceType: "ResearchElementDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadResearchStudy(ctx context.Context, id string) (r4.ResearchStudy, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ResearchStudy", id)
	if err != nil {
		return r4.ResearchStudy{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchStudy)
	if !ok {
		return r4.ResearchStudy{}, capabilities.InvalidResourceError{ResourceType: "ResearchStudy"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadResearchSubject(ctx context.Context, id string) (r4.ResearchSubject, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ResearchSubject", id)
	if err != nil {
		return r4.ResearchSubject{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ResearchSubject)
	if !ok {
		return r4.ResearchSubject{}, capabilities.InvalidResourceError{ResourceType: "ResearchSubject"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadRiskAssessment(ctx context.Context, id string) (r4.RiskAssessment, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "RiskAssessment", id)
	if err != nil {
		return r4.RiskAssessment{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RiskAssessment)
	if !ok {
		return r4.RiskAssessment{}, capabilities.InvalidResourceError{ResourceType: "RiskAssessment"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadRiskEvidenceSynthesis(ctx context.Context, id string) (r4.RiskEvidenceSynthesis, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "RiskEvidenceSynthesis", id)
	if err != nil {
		return r4.RiskEvidenceSynthesis{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.RiskEvidenceSynthesis)
	if !ok {
		return r4.RiskEvidenceSynthesis{}, capabilities.InvalidResourceError{ResourceType: "RiskEvidenceSynthesis"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSchedule(ctx context.Context, id string) (r4.Schedule, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Schedule", id)
	if err != nil {
		return r4.Schedule{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Schedule)
	if !ok {
		return r4.Schedule{}, capabilities.InvalidResourceError{ResourceType: "Schedule"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSearchParameter(ctx context.Context, id string) (r4.SearchParameter, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SearchParameter", id)
	if err != nil {
		return r4.SearchParameter{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SearchParameter)
	if !ok {
		return r4.SearchParameter{}, capabilities.InvalidResourceError{ResourceType: "SearchParameter"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadServiceRequest(ctx context.Context, id string) (r4.ServiceRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ServiceRequest", id)
	if err != nil {
		return r4.ServiceRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ServiceRequest)
	if !ok {
		return r4.ServiceRequest{}, capabilities.InvalidResourceError{ResourceType: "ServiceRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSlot(ctx context.Context, id string) (r4.Slot, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Slot", id)
	if err != nil {
		return r4.Slot{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Slot)
	if !ok {
		return r4.Slot{}, capabilities.InvalidResourceError{ResourceType: "Slot"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSpecimen(ctx context.Context, id string) (r4.Specimen, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Specimen", id)
	if err != nil {
		return r4.Specimen{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Specimen)
	if !ok {
		return r4.Specimen{}, capabilities.InvalidResourceError{ResourceType: "Specimen"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSpecimenDefinition(ctx context.Context, id string) (r4.SpecimenDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SpecimenDefinition", id)
	if err != nil {
		return r4.SpecimenDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SpecimenDefinition)
	if !ok {
		return r4.SpecimenDefinition{}, capabilities.InvalidResourceError{ResourceType: "SpecimenDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadStructureDefinition(ctx context.Context, id string) (r4.StructureDefinition, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "StructureDefinition", id)
	if err != nil {
		return r4.StructureDefinition{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.StructureDefinition)
	if !ok {
		return r4.StructureDefinition{}, capabilities.InvalidResourceError{ResourceType: "StructureDefinition"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadStructureMap(ctx context.Context, id string) (r4.StructureMap, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "StructureMap", id)
	if err != nil {
		return r4.StructureMap{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.StructureMap)
	if !ok {
		return r4.StructureMap{}, capabilities.InvalidResourceError{ResourceType: "StructureMap"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubscription(ctx context.Context, id string) (r4.Subscription, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Subscription", id)
	if err != nil {
		return r4.Subscription{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Subscription)
	if !ok {
		return r4.Subscription{}, capabilities.InvalidResourceError{ResourceType: "Subscription"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstance(ctx context.Context, id string) (r4.Substance, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Substance", id)
	if err != nil {
		return r4.Substance{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Substance)
	if !ok {
		return r4.Substance{}, capabilities.InvalidResourceError{ResourceType: "Substance"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstanceNucleicAcid(ctx context.Context, id string) (r4.SubstanceNucleicAcid, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstanceNucleicAcid", id)
	if err != nil {
		return r4.SubstanceNucleicAcid{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceNucleicAcid)
	if !ok {
		return r4.SubstanceNucleicAcid{}, capabilities.InvalidResourceError{ResourceType: "SubstanceNucleicAcid"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstancePolymer(ctx context.Context, id string) (r4.SubstancePolymer, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstancePolymer", id)
	if err != nil {
		return r4.SubstancePolymer{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstancePolymer)
	if !ok {
		return r4.SubstancePolymer{}, capabilities.InvalidResourceError{ResourceType: "SubstancePolymer"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstanceProtein(ctx context.Context, id string) (r4.SubstanceProtein, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstanceProtein", id)
	if err != nil {
		return r4.SubstanceProtein{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceProtein)
	if !ok {
		return r4.SubstanceProtein{}, capabilities.InvalidResourceError{ResourceType: "SubstanceProtein"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstanceReferenceInformation(ctx context.Context, id string) (r4.SubstanceReferenceInformation, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstanceReferenceInformation", id)
	if err != nil {
		return r4.SubstanceReferenceInformation{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceReferenceInformation)
	if !ok {
		return r4.SubstanceReferenceInformation{}, capabilities.InvalidResourceError{ResourceType: "SubstanceReferenceInformation"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstanceSourceMaterial(ctx context.Context, id string) (r4.SubstanceSourceMaterial, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstanceSourceMaterial", id)
	if err != nil {
		return r4.SubstanceSourceMaterial{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceSourceMaterial)
	if !ok {
		return r4.SubstanceSourceMaterial{}, capabilities.InvalidResourceError{ResourceType: "SubstanceSourceMaterial"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSubstanceSpecification(ctx context.Context, id string) (r4.SubstanceSpecification, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SubstanceSpecification", id)
	if err != nil {
		return r4.SubstanceSpecification{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SubstanceSpecification)
	if !ok {
		return r4.SubstanceSpecification{}, capabilities.InvalidResourceError{ResourceType: "SubstanceSpecification"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSupplyDelivery(ctx context.Context, id string) (r4.SupplyDelivery, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SupplyDelivery", id)
	if err != nil {
		return r4.SupplyDelivery{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SupplyDelivery)
	if !ok {
		return r4.SupplyDelivery{}, capabilities.InvalidResourceError{ResourceType: "SupplyDelivery"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadSupplyRequest(ctx context.Context, id string) (r4.SupplyRequest, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "SupplyRequest", id)
	if err != nil {
		return r4.SupplyRequest{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.SupplyRequest)
	if !ok {
		return r4.SupplyRequest{}, capabilities.InvalidResourceError{ResourceType: "SupplyRequest"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadTask(ctx context.Context, id string) (r4.Task, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "Task", id)
	if err != nil {
		return r4.Task{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.Task)
	if !ok {
		return r4.Task{}, capabilities.InvalidResourceError{ResourceType: "Task"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadTerminologyCapabilities(ctx context.Context, id string) (r4.TerminologyCapabilities, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "TerminologyCapabilities", id)
	if err != nil {
		return r4.TerminologyCapabilities{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TerminologyCapabilities)
	if !ok {
		return r4.TerminologyCapabilities{}, capabilities.InvalidResourceError{ResourceType: "TerminologyCapabilities"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadTestReport(ctx context.Context, id string) (r4.TestReport, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "TestReport", id)
	if err != nil {
		return r4.TestReport{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TestReport)
	if !ok {
		return r4.TestReport{}, capabilities.InvalidResourceError{ResourceType: "TestReport"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadTestScript(ctx context.Context, id string) (r4.TestScript, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "TestScript", id)
	if err != nil {
		return r4.TestScript{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.TestScript)
	if !ok {
		return r4.TestScript{}, capabilities.InvalidResourceError{ResourceType: "TestScript"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadValueSet(ctx context.Context, id string) (r4.ValueSet, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "ValueSet", id)
	if err != nil {
		return r4.ValueSet{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.ValueSet)
	if !ok {
		return r4.ValueSet{}, capabilities.InvalidResourceError{ResourceType: "ValueSet"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadVerificationResult(ctx context.Context, id string) (r4.VerificationResult, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "VerificationResult", id)
	if err != nil {
		return r4.VerificationResult{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.VerificationResult)
	if !ok {
		return r4.VerificationResult{}, capabilities.InvalidResourceError{ResourceType: "VerificationResult"}
	}
	return impl, nil
}
func (w InternalWrapper) ReadVisionPrescription(ctx context.Context, id string) (r4.VisionPrescription, capabilities.FHIRError) {
	v, err := w.GenericAPI.Read(ctx, "VisionPrescription", id)
	if err != nil {
		return r4.VisionPrescription{}, err
	}
	contained, ok := v.(r4.ContainedResource)
	if ok {
		v = contained.Resource
	}
	impl, ok := v.(r4.VisionPrescription)
	if !ok {
		return r4.VisionPrescription{}, capabilities.InvalidResourceError{ResourceType: "VisionPrescription"}
	}
	return impl, nil
}
