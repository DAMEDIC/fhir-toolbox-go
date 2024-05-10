package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

type primitiveElement struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
}
type containedResource struct {
	resource model.Resource
}

func (r containedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.resource)
}
func (cr *containedResource) UnmarshalJSON(b []byte) error {
	var t struct {
		ResourceType string `json:"resourceType"`
	}
	if json.Unmarshal(b, &t) != nil {
		return json.Unmarshal(b, &t)
	}
	switch t.ResourceType {
	case "Account":
		var r Account
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ActivityDefinition":
		var r ActivityDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "AdverseEvent":
		var r AdverseEvent
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "AllergyIntolerance":
		var r AllergyIntolerance
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Appointment":
		var r Appointment
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "AppointmentResponse":
		var r AppointmentResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "AuditEvent":
		var r AuditEvent
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Basic":
		var r Basic
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Binary":
		var r Binary
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "BiologicallyDerivedProduct":
		var r BiologicallyDerivedProduct
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "BodyStructure":
		var r BodyStructure
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Bundle":
		var r Bundle
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CapabilityStatement":
		var r CapabilityStatement
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CarePlan":
		var r CarePlan
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CareTeam":
		var r CareTeam
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CatalogEntry":
		var r CatalogEntry
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ChargeItem":
		var r ChargeItem
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ChargeItemDefinition":
		var r ChargeItemDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Claim":
		var r Claim
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ClaimResponse":
		var r ClaimResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ClinicalImpression":
		var r ClinicalImpression
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CodeSystem":
		var r CodeSystem
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Communication":
		var r Communication
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CommunicationRequest":
		var r CommunicationRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CompartmentDefinition":
		var r CompartmentDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Composition":
		var r Composition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ConceptMap":
		var r ConceptMap
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Condition":
		var r Condition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Consent":
		var r Consent
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Contract":
		var r Contract
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Coverage":
		var r Coverage
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CoverageEligibilityRequest":
		var r CoverageEligibilityRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "CoverageEligibilityResponse":
		var r CoverageEligibilityResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DetectedIssue":
		var r DetectedIssue
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Device":
		var r Device
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DeviceDefinition":
		var r DeviceDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DeviceMetric":
		var r DeviceMetric
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DeviceRequest":
		var r DeviceRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DeviceUseStatement":
		var r DeviceUseStatement
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DiagnosticReport":
		var r DiagnosticReport
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DocumentManifest":
		var r DocumentManifest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "DocumentReference":
		var r DocumentReference
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EffectEvidenceSynthesis":
		var r EffectEvidenceSynthesis
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Encounter":
		var r Encounter
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Endpoint":
		var r Endpoint
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EnrollmentRequest":
		var r EnrollmentRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EnrollmentResponse":
		var r EnrollmentResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EpisodeOfCare":
		var r EpisodeOfCare
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EventDefinition":
		var r EventDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Evidence":
		var r Evidence
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "EvidenceVariable":
		var r EvidenceVariable
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ExampleScenario":
		var r ExampleScenario
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ExplanationOfBenefit":
		var r ExplanationOfBenefit
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "FamilyMemberHistory":
		var r FamilyMemberHistory
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Flag":
		var r Flag
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Goal":
		var r Goal
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "GraphDefinition":
		var r GraphDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Group":
		var r Group
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "GuidanceResponse":
		var r GuidanceResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "HealthcareService":
		var r HealthcareService
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ImagingStudy":
		var r ImagingStudy
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Immunization":
		var r Immunization
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ImmunizationEvaluation":
		var r ImmunizationEvaluation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ImmunizationRecommendation":
		var r ImmunizationRecommendation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ImplementationGuide":
		var r ImplementationGuide
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "InsurancePlan":
		var r InsurancePlan
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Invoice":
		var r Invoice
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Library":
		var r Library
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Linkage":
		var r Linkage
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "List":
		var r List
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Location":
		var r Location
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Measure":
		var r Measure
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MeasureReport":
		var r MeasureReport
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Media":
		var r Media
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Medication":
		var r Medication
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicationAdministration":
		var r MedicationAdministration
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicationDispense":
		var r MedicationDispense
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicationKnowledge":
		var r MedicationKnowledge
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicationRequest":
		var r MedicationRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicationStatement":
		var r MedicationStatement
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProduct":
		var r MedicinalProduct
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductAuthorization":
		var r MedicinalProductAuthorization
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductContraindication":
		var r MedicinalProductContraindication
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductIndication":
		var r MedicinalProductIndication
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductIngredient":
		var r MedicinalProductIngredient
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductInteraction":
		var r MedicinalProductInteraction
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductManufactured":
		var r MedicinalProductManufactured
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductPackaged":
		var r MedicinalProductPackaged
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductPharmaceutical":
		var r MedicinalProductPharmaceutical
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MedicinalProductUndesirableEffect":
		var r MedicinalProductUndesirableEffect
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MessageDefinition":
		var r MessageDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MessageHeader":
		var r MessageHeader
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "MolecularSequence":
		var r MolecularSequence
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "NamingSystem":
		var r NamingSystem
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "NutritionOrder":
		var r NutritionOrder
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Observation":
		var r Observation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ObservationDefinition":
		var r ObservationDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "OperationDefinition":
		var r OperationDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "OperationOutcome":
		var r OperationOutcome
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Organization":
		var r Organization
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "OrganizationAffiliation":
		var r OrganizationAffiliation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Parameters":
		var r Parameters
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Patient":
		var r Patient
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "PaymentNotice":
		var r PaymentNotice
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "PaymentReconciliation":
		var r PaymentReconciliation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Person":
		var r Person
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "PlanDefinition":
		var r PlanDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Practitioner":
		var r Practitioner
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "PractitionerRole":
		var r PractitionerRole
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Procedure":
		var r Procedure
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Provenance":
		var r Provenance
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Questionnaire":
		var r Questionnaire
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "QuestionnaireResponse":
		var r QuestionnaireResponse
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "RelatedPerson":
		var r RelatedPerson
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "RequestGroup":
		var r RequestGroup
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ResearchDefinition":
		var r ResearchDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ResearchElementDefinition":
		var r ResearchElementDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ResearchStudy":
		var r ResearchStudy
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ResearchSubject":
		var r ResearchSubject
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "RiskAssessment":
		var r RiskAssessment
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "RiskEvidenceSynthesis":
		var r RiskEvidenceSynthesis
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Schedule":
		var r Schedule
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SearchParameter":
		var r SearchParameter
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ServiceRequest":
		var r ServiceRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Slot":
		var r Slot
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Specimen":
		var r Specimen
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SpecimenDefinition":
		var r SpecimenDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "StructureDefinition":
		var r StructureDefinition
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "StructureMap":
		var r StructureMap
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Subscription":
		var r Subscription
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Substance":
		var r Substance
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstanceNucleicAcid":
		var r SubstanceNucleicAcid
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstancePolymer":
		var r SubstancePolymer
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstanceProtein":
		var r SubstanceProtein
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstanceReferenceInformation":
		var r SubstanceReferenceInformation
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstanceSourceMaterial":
		var r SubstanceSourceMaterial
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SubstanceSpecification":
		var r SubstanceSpecification
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SupplyDelivery":
		var r SupplyDelivery
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "SupplyRequest":
		var r SupplyRequest
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "Task":
		var r Task
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "TerminologyCapabilities":
		var r TerminologyCapabilities
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "TestReport":
		var r TestReport
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "TestScript":
		var r TestScript
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "ValueSet":
		var r ValueSet
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "VerificationResult":
		var r VerificationResult
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	case "VisionPrescription":
		var r VisionPrescription
		if json.Unmarshal(b, &r) != nil {
			return json.Unmarshal(b, &r)
		}
		*cr = containedResource{resource: r}
	default:
		return fmt.Errorf("unknown resource type: %s", t.ResourceType)
	}
	return nil
}
