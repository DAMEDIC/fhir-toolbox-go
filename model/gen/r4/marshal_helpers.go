package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
	"unicode"
)

type primitiveElement struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
}
type ContainedResource struct {
	model.Resource
}

func (r ContainedResource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ContainedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Resource)
}
func (cr *ContainedResource) UnmarshalJSON(b []byte) error {
	var t struct {
		ResourceType string `json:"resourceType"`
	}
	if json.Unmarshal(b, &t) != nil {
		return json.Unmarshal(b, &t)
	}
	switch t.ResourceType {
	case "Account":
		var r Account
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ActivityDefinition":
		var r ActivityDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AdverseEvent":
		var r AdverseEvent
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AllergyIntolerance":
		var r AllergyIntolerance
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Appointment":
		var r Appointment
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AppointmentResponse":
		var r AppointmentResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AuditEvent":
		var r AuditEvent
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Basic":
		var r Basic
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Binary":
		var r Binary
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BiologicallyDerivedProduct":
		var r BiologicallyDerivedProduct
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BodyStructure":
		var r BodyStructure
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Bundle":
		var r Bundle
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CapabilityStatement":
		var r CapabilityStatement
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CarePlan":
		var r CarePlan
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CareTeam":
		var r CareTeam
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CatalogEntry":
		var r CatalogEntry
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItem":
		var r ChargeItem
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItemDefinition":
		var r ChargeItemDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Claim":
		var r Claim
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClaimResponse":
		var r ClaimResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClinicalImpression":
		var r ClinicalImpression
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CodeSystem":
		var r CodeSystem
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Communication":
		var r Communication
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CommunicationRequest":
		var r CommunicationRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CompartmentDefinition":
		var r CompartmentDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Composition":
		var r Composition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ConceptMap":
		var r ConceptMap
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Condition":
		var r Condition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Consent":
		var r Consent
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Contract":
		var r Contract
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Coverage":
		var r Coverage
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityRequest":
		var r CoverageEligibilityRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityResponse":
		var r CoverageEligibilityResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DetectedIssue":
		var r DetectedIssue
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Device":
		var r Device
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceDefinition":
		var r DeviceDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceMetric":
		var r DeviceMetric
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceRequest":
		var r DeviceRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceUseStatement":
		var r DeviceUseStatement
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DiagnosticReport":
		var r DiagnosticReport
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DocumentManifest":
		var r DocumentManifest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DocumentReference":
		var r DocumentReference
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EffectEvidenceSynthesis":
		var r EffectEvidenceSynthesis
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Encounter":
		var r Encounter
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Endpoint":
		var r Endpoint
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentRequest":
		var r EnrollmentRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentResponse":
		var r EnrollmentResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EpisodeOfCare":
		var r EpisodeOfCare
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EventDefinition":
		var r EventDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Evidence":
		var r Evidence
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EvidenceVariable":
		var r EvidenceVariable
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExampleScenario":
		var r ExampleScenario
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExplanationOfBenefit":
		var r ExplanationOfBenefit
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "FamilyMemberHistory":
		var r FamilyMemberHistory
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Flag":
		var r Flag
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Goal":
		var r Goal
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GraphDefinition":
		var r GraphDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Group":
		var r Group
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GuidanceResponse":
		var r GuidanceResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "HealthcareService":
		var r HealthcareService
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImagingStudy":
		var r ImagingStudy
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Immunization":
		var r Immunization
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationEvaluation":
		var r ImmunizationEvaluation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationRecommendation":
		var r ImmunizationRecommendation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImplementationGuide":
		var r ImplementationGuide
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InsurancePlan":
		var r InsurancePlan
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Invoice":
		var r Invoice
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Library":
		var r Library
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Linkage":
		var r Linkage
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "List":
		var r List
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Location":
		var r Location
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Measure":
		var r Measure
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MeasureReport":
		var r MeasureReport
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Media":
		var r Media
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Medication":
		var r Medication
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationAdministration":
		var r MedicationAdministration
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationDispense":
		var r MedicationDispense
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationKnowledge":
		var r MedicationKnowledge
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationRequest":
		var r MedicationRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationStatement":
		var r MedicationStatement
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProduct":
		var r MedicinalProduct
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductAuthorization":
		var r MedicinalProductAuthorization
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductContraindication":
		var r MedicinalProductContraindication
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductIndication":
		var r MedicinalProductIndication
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductIngredient":
		var r MedicinalProductIngredient
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductInteraction":
		var r MedicinalProductInteraction
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductManufactured":
		var r MedicinalProductManufactured
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductPackaged":
		var r MedicinalProductPackaged
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductPharmaceutical":
		var r MedicinalProductPharmaceutical
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductUndesirableEffect":
		var r MedicinalProductUndesirableEffect
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageDefinition":
		var r MessageDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageHeader":
		var r MessageHeader
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MolecularSequence":
		var r MolecularSequence
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NamingSystem":
		var r NamingSystem
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NutritionOrder":
		var r NutritionOrder
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Observation":
		var r Observation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ObservationDefinition":
		var r ObservationDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationDefinition":
		var r OperationDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationOutcome":
		var r OperationOutcome
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Organization":
		var r Organization
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OrganizationAffiliation":
		var r OrganizationAffiliation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Parameters":
		var r Parameters
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Patient":
		var r Patient
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentNotice":
		var r PaymentNotice
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentReconciliation":
		var r PaymentReconciliation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Person":
		var r Person
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PlanDefinition":
		var r PlanDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Practitioner":
		var r Practitioner
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PractitionerRole":
		var r PractitionerRole
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Procedure":
		var r Procedure
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Provenance":
		var r Provenance
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Questionnaire":
		var r Questionnaire
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "QuestionnaireResponse":
		var r QuestionnaireResponse
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RelatedPerson":
		var r RelatedPerson
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RequestGroup":
		var r RequestGroup
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchDefinition":
		var r ResearchDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchElementDefinition":
		var r ResearchElementDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchStudy":
		var r ResearchStudy
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchSubject":
		var r ResearchSubject
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RiskAssessment":
		var r RiskAssessment
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RiskEvidenceSynthesis":
		var r RiskEvidenceSynthesis
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Schedule":
		var r Schedule
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SearchParameter":
		var r SearchParameter
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ServiceRequest":
		var r ServiceRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Slot":
		var r Slot
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Specimen":
		var r Specimen
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SpecimenDefinition":
		var r SpecimenDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureDefinition":
		var r StructureDefinition
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureMap":
		var r StructureMap
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Subscription":
		var r Subscription
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Substance":
		var r Substance
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceNucleicAcid":
		var r SubstanceNucleicAcid
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstancePolymer":
		var r SubstancePolymer
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceProtein":
		var r SubstanceProtein
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceReferenceInformation":
		var r SubstanceReferenceInformation
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceSourceMaterial":
		var r SubstanceSourceMaterial
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceSpecification":
		var r SubstanceSpecification
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyDelivery":
		var r SupplyDelivery
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyRequest":
		var r SupplyRequest
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Task":
		var r Task
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TerminologyCapabilities":
		var r TerminologyCapabilities
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestReport":
		var r TestReport
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestScript":
		var r TestScript
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ValueSet":
		var r ValueSet
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VerificationResult":
		var r VerificationResult
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VisionPrescription":
		var r VisionPrescription
		err := json.Unmarshal(b, &r)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	default:
		return fmt.Errorf("unknown resource type: %s", t.ResourceType)
	}
}
func (r ContainedResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "ContainedResource" {
		start.Name.Space = "http://hl7.org/fhir"
		return e.EncodeElement(r.Resource, start)
	} else {
		err := e.EncodeToken(start)
		if err != nil {
			return err
		}
		err = e.Encode(r.Resource)
		if err != nil {
			return err
		}
		err = e.EncodeToken(start.End())
		if err != nil {
			return err
		}
		return nil
	}
}
func (cr *ContainedResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if unicode.IsLower(rune(start.Name.Local[0])) {
		err := d.Decode(cr)
		if err != nil {
			return err
		}
		for {
			t, err := d.Token()
			if err != nil {
				return err
			}
			_, ok := t.(xml.EndElement)
			if ok {
				break
			}
		}
		return nil
	}
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	switch start.Name.Local {
	case "Account":
		var r Account
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ActivityDefinition":
		var r ActivityDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AdverseEvent":
		var r AdverseEvent
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AllergyIntolerance":
		var r AllergyIntolerance
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Appointment":
		var r Appointment
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AppointmentResponse":
		var r AppointmentResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AuditEvent":
		var r AuditEvent
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Basic":
		var r Basic
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Binary":
		var r Binary
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BiologicallyDerivedProduct":
		var r BiologicallyDerivedProduct
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BodyStructure":
		var r BodyStructure
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Bundle":
		var r Bundle
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CapabilityStatement":
		var r CapabilityStatement
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CarePlan":
		var r CarePlan
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CareTeam":
		var r CareTeam
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CatalogEntry":
		var r CatalogEntry
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItem":
		var r ChargeItem
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItemDefinition":
		var r ChargeItemDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Claim":
		var r Claim
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClaimResponse":
		var r ClaimResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClinicalImpression":
		var r ClinicalImpression
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CodeSystem":
		var r CodeSystem
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Communication":
		var r Communication
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CommunicationRequest":
		var r CommunicationRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CompartmentDefinition":
		var r CompartmentDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Composition":
		var r Composition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ConceptMap":
		var r ConceptMap
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Condition":
		var r Condition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Consent":
		var r Consent
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Contract":
		var r Contract
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Coverage":
		var r Coverage
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityRequest":
		var r CoverageEligibilityRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityResponse":
		var r CoverageEligibilityResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DetectedIssue":
		var r DetectedIssue
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Device":
		var r Device
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceDefinition":
		var r DeviceDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceMetric":
		var r DeviceMetric
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceRequest":
		var r DeviceRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceUseStatement":
		var r DeviceUseStatement
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DiagnosticReport":
		var r DiagnosticReport
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DocumentManifest":
		var r DocumentManifest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DocumentReference":
		var r DocumentReference
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EffectEvidenceSynthesis":
		var r EffectEvidenceSynthesis
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Encounter":
		var r Encounter
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Endpoint":
		var r Endpoint
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentRequest":
		var r EnrollmentRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentResponse":
		var r EnrollmentResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EpisodeOfCare":
		var r EpisodeOfCare
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EventDefinition":
		var r EventDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Evidence":
		var r Evidence
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EvidenceVariable":
		var r EvidenceVariable
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExampleScenario":
		var r ExampleScenario
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExplanationOfBenefit":
		var r ExplanationOfBenefit
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "FamilyMemberHistory":
		var r FamilyMemberHistory
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Flag":
		var r Flag
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Goal":
		var r Goal
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GraphDefinition":
		var r GraphDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Group":
		var r Group
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GuidanceResponse":
		var r GuidanceResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "HealthcareService":
		var r HealthcareService
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImagingStudy":
		var r ImagingStudy
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Immunization":
		var r Immunization
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationEvaluation":
		var r ImmunizationEvaluation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationRecommendation":
		var r ImmunizationRecommendation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImplementationGuide":
		var r ImplementationGuide
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InsurancePlan":
		var r InsurancePlan
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Invoice":
		var r Invoice
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Library":
		var r Library
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Linkage":
		var r Linkage
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "List":
		var r List
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Location":
		var r Location
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Measure":
		var r Measure
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MeasureReport":
		var r MeasureReport
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Media":
		var r Media
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Medication":
		var r Medication
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationAdministration":
		var r MedicationAdministration
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationDispense":
		var r MedicationDispense
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationKnowledge":
		var r MedicationKnowledge
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationRequest":
		var r MedicationRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationStatement":
		var r MedicationStatement
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProduct":
		var r MedicinalProduct
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductAuthorization":
		var r MedicinalProductAuthorization
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductContraindication":
		var r MedicinalProductContraindication
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductIndication":
		var r MedicinalProductIndication
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductIngredient":
		var r MedicinalProductIngredient
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductInteraction":
		var r MedicinalProductInteraction
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductManufactured":
		var r MedicinalProductManufactured
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductPackaged":
		var r MedicinalProductPackaged
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductPharmaceutical":
		var r MedicinalProductPharmaceutical
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductUndesirableEffect":
		var r MedicinalProductUndesirableEffect
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageDefinition":
		var r MessageDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageHeader":
		var r MessageHeader
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MolecularSequence":
		var r MolecularSequence
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NamingSystem":
		var r NamingSystem
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NutritionOrder":
		var r NutritionOrder
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Observation":
		var r Observation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ObservationDefinition":
		var r ObservationDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationDefinition":
		var r OperationDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationOutcome":
		var r OperationOutcome
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Organization":
		var r Organization
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OrganizationAffiliation":
		var r OrganizationAffiliation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Parameters":
		var r Parameters
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Patient":
		var r Patient
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentNotice":
		var r PaymentNotice
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentReconciliation":
		var r PaymentReconciliation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Person":
		var r Person
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PlanDefinition":
		var r PlanDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Practitioner":
		var r Practitioner
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PractitionerRole":
		var r PractitionerRole
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Procedure":
		var r Procedure
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Provenance":
		var r Provenance
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Questionnaire":
		var r Questionnaire
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "QuestionnaireResponse":
		var r QuestionnaireResponse
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RelatedPerson":
		var r RelatedPerson
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RequestGroup":
		var r RequestGroup
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchDefinition":
		var r ResearchDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchElementDefinition":
		var r ResearchElementDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchStudy":
		var r ResearchStudy
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchSubject":
		var r ResearchSubject
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RiskAssessment":
		var r RiskAssessment
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RiskEvidenceSynthesis":
		var r RiskEvidenceSynthesis
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Schedule":
		var r Schedule
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SearchParameter":
		var r SearchParameter
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ServiceRequest":
		var r ServiceRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Slot":
		var r Slot
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Specimen":
		var r Specimen
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SpecimenDefinition":
		var r SpecimenDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureDefinition":
		var r StructureDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureMap":
		var r StructureMap
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Subscription":
		var r Subscription
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Substance":
		var r Substance
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceNucleicAcid":
		var r SubstanceNucleicAcid
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstancePolymer":
		var r SubstancePolymer
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceProtein":
		var r SubstanceProtein
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceReferenceInformation":
		var r SubstanceReferenceInformation
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceSourceMaterial":
		var r SubstanceSourceMaterial
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceSpecification":
		var r SubstanceSpecification
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyDelivery":
		var r SupplyDelivery
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyRequest":
		var r SupplyRequest
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Task":
		var r Task
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TerminologyCapabilities":
		var r TerminologyCapabilities
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestReport":
		var r TestReport
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestScript":
		var r TestScript
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ValueSet":
		var r ValueSet
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VerificationResult":
		var r VerificationResult
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VisionPrescription":
		var r VisionPrescription
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	default:
		return fmt.Errorf("unknown resource type: %s", start.Name.Local)
	}
}
