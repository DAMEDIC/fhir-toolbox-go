package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
	"io"
	"unicode"
)

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
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ContainedResource) marshalJSON(w io.Writer) error {
	switch t := r.Resource.(type) {
	case Account:
		return t.marshalJSON(w)
	case *Account:
		return t.marshalJSON(w)
	case ActivityDefinition:
		return t.marshalJSON(w)
	case *ActivityDefinition:
		return t.marshalJSON(w)
	case ActorDefinition:
		return t.marshalJSON(w)
	case *ActorDefinition:
		return t.marshalJSON(w)
	case AdministrableProductDefinition:
		return t.marshalJSON(w)
	case *AdministrableProductDefinition:
		return t.marshalJSON(w)
	case AdverseEvent:
		return t.marshalJSON(w)
	case *AdverseEvent:
		return t.marshalJSON(w)
	case AllergyIntolerance:
		return t.marshalJSON(w)
	case *AllergyIntolerance:
		return t.marshalJSON(w)
	case Appointment:
		return t.marshalJSON(w)
	case *Appointment:
		return t.marshalJSON(w)
	case AppointmentResponse:
		return t.marshalJSON(w)
	case *AppointmentResponse:
		return t.marshalJSON(w)
	case ArtifactAssessment:
		return t.marshalJSON(w)
	case *ArtifactAssessment:
		return t.marshalJSON(w)
	case AuditEvent:
		return t.marshalJSON(w)
	case *AuditEvent:
		return t.marshalJSON(w)
	case Basic:
		return t.marshalJSON(w)
	case *Basic:
		return t.marshalJSON(w)
	case Binary:
		return t.marshalJSON(w)
	case *Binary:
		return t.marshalJSON(w)
	case BiologicallyDerivedProduct:
		return t.marshalJSON(w)
	case *BiologicallyDerivedProduct:
		return t.marshalJSON(w)
	case BiologicallyDerivedProductDispense:
		return t.marshalJSON(w)
	case *BiologicallyDerivedProductDispense:
		return t.marshalJSON(w)
	case BodyStructure:
		return t.marshalJSON(w)
	case *BodyStructure:
		return t.marshalJSON(w)
	case Bundle:
		return t.marshalJSON(w)
	case *Bundle:
		return t.marshalJSON(w)
	case CapabilityStatement:
		return t.marshalJSON(w)
	case *CapabilityStatement:
		return t.marshalJSON(w)
	case CarePlan:
		return t.marshalJSON(w)
	case *CarePlan:
		return t.marshalJSON(w)
	case CareTeam:
		return t.marshalJSON(w)
	case *CareTeam:
		return t.marshalJSON(w)
	case ChargeItem:
		return t.marshalJSON(w)
	case *ChargeItem:
		return t.marshalJSON(w)
	case ChargeItemDefinition:
		return t.marshalJSON(w)
	case *ChargeItemDefinition:
		return t.marshalJSON(w)
	case Citation:
		return t.marshalJSON(w)
	case *Citation:
		return t.marshalJSON(w)
	case Claim:
		return t.marshalJSON(w)
	case *Claim:
		return t.marshalJSON(w)
	case ClaimResponse:
		return t.marshalJSON(w)
	case *ClaimResponse:
		return t.marshalJSON(w)
	case ClinicalImpression:
		return t.marshalJSON(w)
	case *ClinicalImpression:
		return t.marshalJSON(w)
	case ClinicalUseDefinition:
		return t.marshalJSON(w)
	case *ClinicalUseDefinition:
		return t.marshalJSON(w)
	case CodeSystem:
		return t.marshalJSON(w)
	case *CodeSystem:
		return t.marshalJSON(w)
	case Communication:
		return t.marshalJSON(w)
	case *Communication:
		return t.marshalJSON(w)
	case CommunicationRequest:
		return t.marshalJSON(w)
	case *CommunicationRequest:
		return t.marshalJSON(w)
	case CompartmentDefinition:
		return t.marshalJSON(w)
	case *CompartmentDefinition:
		return t.marshalJSON(w)
	case Composition:
		return t.marshalJSON(w)
	case *Composition:
		return t.marshalJSON(w)
	case ConceptMap:
		return t.marshalJSON(w)
	case *ConceptMap:
		return t.marshalJSON(w)
	case Condition:
		return t.marshalJSON(w)
	case *Condition:
		return t.marshalJSON(w)
	case ConditionDefinition:
		return t.marshalJSON(w)
	case *ConditionDefinition:
		return t.marshalJSON(w)
	case Consent:
		return t.marshalJSON(w)
	case *Consent:
		return t.marshalJSON(w)
	case Contract:
		return t.marshalJSON(w)
	case *Contract:
		return t.marshalJSON(w)
	case Coverage:
		return t.marshalJSON(w)
	case *Coverage:
		return t.marshalJSON(w)
	case CoverageEligibilityRequest:
		return t.marshalJSON(w)
	case *CoverageEligibilityRequest:
		return t.marshalJSON(w)
	case CoverageEligibilityResponse:
		return t.marshalJSON(w)
	case *CoverageEligibilityResponse:
		return t.marshalJSON(w)
	case DetectedIssue:
		return t.marshalJSON(w)
	case *DetectedIssue:
		return t.marshalJSON(w)
	case Device:
		return t.marshalJSON(w)
	case *Device:
		return t.marshalJSON(w)
	case DeviceAssociation:
		return t.marshalJSON(w)
	case *DeviceAssociation:
		return t.marshalJSON(w)
	case DeviceDefinition:
		return t.marshalJSON(w)
	case *DeviceDefinition:
		return t.marshalJSON(w)
	case DeviceDispense:
		return t.marshalJSON(w)
	case *DeviceDispense:
		return t.marshalJSON(w)
	case DeviceMetric:
		return t.marshalJSON(w)
	case *DeviceMetric:
		return t.marshalJSON(w)
	case DeviceRequest:
		return t.marshalJSON(w)
	case *DeviceRequest:
		return t.marshalJSON(w)
	case DeviceUsage:
		return t.marshalJSON(w)
	case *DeviceUsage:
		return t.marshalJSON(w)
	case DiagnosticReport:
		return t.marshalJSON(w)
	case *DiagnosticReport:
		return t.marshalJSON(w)
	case DocumentReference:
		return t.marshalJSON(w)
	case *DocumentReference:
		return t.marshalJSON(w)
	case Encounter:
		return t.marshalJSON(w)
	case *Encounter:
		return t.marshalJSON(w)
	case EncounterHistory:
		return t.marshalJSON(w)
	case *EncounterHistory:
		return t.marshalJSON(w)
	case Endpoint:
		return t.marshalJSON(w)
	case *Endpoint:
		return t.marshalJSON(w)
	case EnrollmentRequest:
		return t.marshalJSON(w)
	case *EnrollmentRequest:
		return t.marshalJSON(w)
	case EnrollmentResponse:
		return t.marshalJSON(w)
	case *EnrollmentResponse:
		return t.marshalJSON(w)
	case EpisodeOfCare:
		return t.marshalJSON(w)
	case *EpisodeOfCare:
		return t.marshalJSON(w)
	case EventDefinition:
		return t.marshalJSON(w)
	case *EventDefinition:
		return t.marshalJSON(w)
	case Evidence:
		return t.marshalJSON(w)
	case *Evidence:
		return t.marshalJSON(w)
	case EvidenceReport:
		return t.marshalJSON(w)
	case *EvidenceReport:
		return t.marshalJSON(w)
	case EvidenceVariable:
		return t.marshalJSON(w)
	case *EvidenceVariable:
		return t.marshalJSON(w)
	case ExampleScenario:
		return t.marshalJSON(w)
	case *ExampleScenario:
		return t.marshalJSON(w)
	case ExplanationOfBenefit:
		return t.marshalJSON(w)
	case *ExplanationOfBenefit:
		return t.marshalJSON(w)
	case FamilyMemberHistory:
		return t.marshalJSON(w)
	case *FamilyMemberHistory:
		return t.marshalJSON(w)
	case Flag:
		return t.marshalJSON(w)
	case *Flag:
		return t.marshalJSON(w)
	case FormularyItem:
		return t.marshalJSON(w)
	case *FormularyItem:
		return t.marshalJSON(w)
	case GenomicStudy:
		return t.marshalJSON(w)
	case *GenomicStudy:
		return t.marshalJSON(w)
	case Goal:
		return t.marshalJSON(w)
	case *Goal:
		return t.marshalJSON(w)
	case GraphDefinition:
		return t.marshalJSON(w)
	case *GraphDefinition:
		return t.marshalJSON(w)
	case Group:
		return t.marshalJSON(w)
	case *Group:
		return t.marshalJSON(w)
	case GuidanceResponse:
		return t.marshalJSON(w)
	case *GuidanceResponse:
		return t.marshalJSON(w)
	case HealthcareService:
		return t.marshalJSON(w)
	case *HealthcareService:
		return t.marshalJSON(w)
	case ImagingSelection:
		return t.marshalJSON(w)
	case *ImagingSelection:
		return t.marshalJSON(w)
	case ImagingStudy:
		return t.marshalJSON(w)
	case *ImagingStudy:
		return t.marshalJSON(w)
	case Immunization:
		return t.marshalJSON(w)
	case *Immunization:
		return t.marshalJSON(w)
	case ImmunizationEvaluation:
		return t.marshalJSON(w)
	case *ImmunizationEvaluation:
		return t.marshalJSON(w)
	case ImmunizationRecommendation:
		return t.marshalJSON(w)
	case *ImmunizationRecommendation:
		return t.marshalJSON(w)
	case ImplementationGuide:
		return t.marshalJSON(w)
	case *ImplementationGuide:
		return t.marshalJSON(w)
	case Ingredient:
		return t.marshalJSON(w)
	case *Ingredient:
		return t.marshalJSON(w)
	case InsurancePlan:
		return t.marshalJSON(w)
	case *InsurancePlan:
		return t.marshalJSON(w)
	case InventoryItem:
		return t.marshalJSON(w)
	case *InventoryItem:
		return t.marshalJSON(w)
	case InventoryReport:
		return t.marshalJSON(w)
	case *InventoryReport:
		return t.marshalJSON(w)
	case Invoice:
		return t.marshalJSON(w)
	case *Invoice:
		return t.marshalJSON(w)
	case Library:
		return t.marshalJSON(w)
	case *Library:
		return t.marshalJSON(w)
	case Linkage:
		return t.marshalJSON(w)
	case *Linkage:
		return t.marshalJSON(w)
	case List:
		return t.marshalJSON(w)
	case *List:
		return t.marshalJSON(w)
	case Location:
		return t.marshalJSON(w)
	case *Location:
		return t.marshalJSON(w)
	case ManufacturedItemDefinition:
		return t.marshalJSON(w)
	case *ManufacturedItemDefinition:
		return t.marshalJSON(w)
	case Measure:
		return t.marshalJSON(w)
	case *Measure:
		return t.marshalJSON(w)
	case MeasureReport:
		return t.marshalJSON(w)
	case *MeasureReport:
		return t.marshalJSON(w)
	case Medication:
		return t.marshalJSON(w)
	case *Medication:
		return t.marshalJSON(w)
	case MedicationAdministration:
		return t.marshalJSON(w)
	case *MedicationAdministration:
		return t.marshalJSON(w)
	case MedicationDispense:
		return t.marshalJSON(w)
	case *MedicationDispense:
		return t.marshalJSON(w)
	case MedicationKnowledge:
		return t.marshalJSON(w)
	case *MedicationKnowledge:
		return t.marshalJSON(w)
	case MedicationRequest:
		return t.marshalJSON(w)
	case *MedicationRequest:
		return t.marshalJSON(w)
	case MedicationStatement:
		return t.marshalJSON(w)
	case *MedicationStatement:
		return t.marshalJSON(w)
	case MedicinalProductDefinition:
		return t.marshalJSON(w)
	case *MedicinalProductDefinition:
		return t.marshalJSON(w)
	case MessageDefinition:
		return t.marshalJSON(w)
	case *MessageDefinition:
		return t.marshalJSON(w)
	case MessageHeader:
		return t.marshalJSON(w)
	case *MessageHeader:
		return t.marshalJSON(w)
	case MolecularSequence:
		return t.marshalJSON(w)
	case *MolecularSequence:
		return t.marshalJSON(w)
	case NamingSystem:
		return t.marshalJSON(w)
	case *NamingSystem:
		return t.marshalJSON(w)
	case NutritionIntake:
		return t.marshalJSON(w)
	case *NutritionIntake:
		return t.marshalJSON(w)
	case NutritionOrder:
		return t.marshalJSON(w)
	case *NutritionOrder:
		return t.marshalJSON(w)
	case NutritionProduct:
		return t.marshalJSON(w)
	case *NutritionProduct:
		return t.marshalJSON(w)
	case Observation:
		return t.marshalJSON(w)
	case *Observation:
		return t.marshalJSON(w)
	case ObservationDefinition:
		return t.marshalJSON(w)
	case *ObservationDefinition:
		return t.marshalJSON(w)
	case OperationDefinition:
		return t.marshalJSON(w)
	case *OperationDefinition:
		return t.marshalJSON(w)
	case OperationOutcome:
		return t.marshalJSON(w)
	case *OperationOutcome:
		return t.marshalJSON(w)
	case Organization:
		return t.marshalJSON(w)
	case *Organization:
		return t.marshalJSON(w)
	case OrganizationAffiliation:
		return t.marshalJSON(w)
	case *OrganizationAffiliation:
		return t.marshalJSON(w)
	case PackagedProductDefinition:
		return t.marshalJSON(w)
	case *PackagedProductDefinition:
		return t.marshalJSON(w)
	case Parameters:
		return t.marshalJSON(w)
	case *Parameters:
		return t.marshalJSON(w)
	case Patient:
		return t.marshalJSON(w)
	case *Patient:
		return t.marshalJSON(w)
	case PaymentNotice:
		return t.marshalJSON(w)
	case *PaymentNotice:
		return t.marshalJSON(w)
	case PaymentReconciliation:
		return t.marshalJSON(w)
	case *PaymentReconciliation:
		return t.marshalJSON(w)
	case Permission:
		return t.marshalJSON(w)
	case *Permission:
		return t.marshalJSON(w)
	case Person:
		return t.marshalJSON(w)
	case *Person:
		return t.marshalJSON(w)
	case PlanDefinition:
		return t.marshalJSON(w)
	case *PlanDefinition:
		return t.marshalJSON(w)
	case Practitioner:
		return t.marshalJSON(w)
	case *Practitioner:
		return t.marshalJSON(w)
	case PractitionerRole:
		return t.marshalJSON(w)
	case *PractitionerRole:
		return t.marshalJSON(w)
	case Procedure:
		return t.marshalJSON(w)
	case *Procedure:
		return t.marshalJSON(w)
	case Provenance:
		return t.marshalJSON(w)
	case *Provenance:
		return t.marshalJSON(w)
	case Questionnaire:
		return t.marshalJSON(w)
	case *Questionnaire:
		return t.marshalJSON(w)
	case QuestionnaireResponse:
		return t.marshalJSON(w)
	case *QuestionnaireResponse:
		return t.marshalJSON(w)
	case RegulatedAuthorization:
		return t.marshalJSON(w)
	case *RegulatedAuthorization:
		return t.marshalJSON(w)
	case RelatedPerson:
		return t.marshalJSON(w)
	case *RelatedPerson:
		return t.marshalJSON(w)
	case RequestOrchestration:
		return t.marshalJSON(w)
	case *RequestOrchestration:
		return t.marshalJSON(w)
	case Requirements:
		return t.marshalJSON(w)
	case *Requirements:
		return t.marshalJSON(w)
	case ResearchStudy:
		return t.marshalJSON(w)
	case *ResearchStudy:
		return t.marshalJSON(w)
	case ResearchSubject:
		return t.marshalJSON(w)
	case *ResearchSubject:
		return t.marshalJSON(w)
	case RiskAssessment:
		return t.marshalJSON(w)
	case *RiskAssessment:
		return t.marshalJSON(w)
	case Schedule:
		return t.marshalJSON(w)
	case *Schedule:
		return t.marshalJSON(w)
	case SearchParameter:
		return t.marshalJSON(w)
	case *SearchParameter:
		return t.marshalJSON(w)
	case ServiceRequest:
		return t.marshalJSON(w)
	case *ServiceRequest:
		return t.marshalJSON(w)
	case Slot:
		return t.marshalJSON(w)
	case *Slot:
		return t.marshalJSON(w)
	case Specimen:
		return t.marshalJSON(w)
	case *Specimen:
		return t.marshalJSON(w)
	case SpecimenDefinition:
		return t.marshalJSON(w)
	case *SpecimenDefinition:
		return t.marshalJSON(w)
	case StructureDefinition:
		return t.marshalJSON(w)
	case *StructureDefinition:
		return t.marshalJSON(w)
	case StructureMap:
		return t.marshalJSON(w)
	case *StructureMap:
		return t.marshalJSON(w)
	case Subscription:
		return t.marshalJSON(w)
	case *Subscription:
		return t.marshalJSON(w)
	case SubscriptionStatus:
		return t.marshalJSON(w)
	case *SubscriptionStatus:
		return t.marshalJSON(w)
	case SubscriptionTopic:
		return t.marshalJSON(w)
	case *SubscriptionTopic:
		return t.marshalJSON(w)
	case Substance:
		return t.marshalJSON(w)
	case *Substance:
		return t.marshalJSON(w)
	case SubstanceDefinition:
		return t.marshalJSON(w)
	case *SubstanceDefinition:
		return t.marshalJSON(w)
	case SubstanceNucleicAcid:
		return t.marshalJSON(w)
	case *SubstanceNucleicAcid:
		return t.marshalJSON(w)
	case SubstancePolymer:
		return t.marshalJSON(w)
	case *SubstancePolymer:
		return t.marshalJSON(w)
	case SubstanceProtein:
		return t.marshalJSON(w)
	case *SubstanceProtein:
		return t.marshalJSON(w)
	case SubstanceReferenceInformation:
		return t.marshalJSON(w)
	case *SubstanceReferenceInformation:
		return t.marshalJSON(w)
	case SubstanceSourceMaterial:
		return t.marshalJSON(w)
	case *SubstanceSourceMaterial:
		return t.marshalJSON(w)
	case SupplyDelivery:
		return t.marshalJSON(w)
	case *SupplyDelivery:
		return t.marshalJSON(w)
	case SupplyRequest:
		return t.marshalJSON(w)
	case *SupplyRequest:
		return t.marshalJSON(w)
	case Task:
		return t.marshalJSON(w)
	case *Task:
		return t.marshalJSON(w)
	case TerminologyCapabilities:
		return t.marshalJSON(w)
	case *TerminologyCapabilities:
		return t.marshalJSON(w)
	case TestPlan:
		return t.marshalJSON(w)
	case *TestPlan:
		return t.marshalJSON(w)
	case TestReport:
		return t.marshalJSON(w)
	case *TestReport:
		return t.marshalJSON(w)
	case TestScript:
		return t.marshalJSON(w)
	case *TestScript:
		return t.marshalJSON(w)
	case Transport:
		return t.marshalJSON(w)
	case *Transport:
		return t.marshalJSON(w)
	case ValueSet:
		return t.marshalJSON(w)
	case *ValueSet:
		return t.marshalJSON(w)
	case VerificationResult:
		return t.marshalJSON(w)
	case *VerificationResult:
		return t.marshalJSON(w)
	case VisionPrescription:
		return t.marshalJSON(w)
	case *VisionPrescription:
		return t.marshalJSON(w)
	default:
		return fmt.Errorf("unknown resource: %v", t)
	}
}
func (r *ContainedResource) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (cr *ContainedResource) unmarshalJSON(d *json.Decoder) error {
	var rawValue json.RawMessage
	err := d.Decode(&rawValue)
	if err != nil {
		return err
	}
	var t struct {
		ResourceType string `json:"resourceType"`
	}
	err = json.Unmarshal(rawValue, &t)
	if err != nil {
		return err
	}
	d = json.NewDecoder(bytes.NewReader(rawValue))
	switch t.ResourceType {
	case "Account":
		var r Account
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ActivityDefinition":
		var r ActivityDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ActorDefinition":
		var r ActorDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AdministrableProductDefinition":
		var r AdministrableProductDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AdverseEvent":
		var r AdverseEvent
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AllergyIntolerance":
		var r AllergyIntolerance
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Appointment":
		var r Appointment
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AppointmentResponse":
		var r AppointmentResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ArtifactAssessment":
		var r ArtifactAssessment
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AuditEvent":
		var r AuditEvent
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Basic":
		var r Basic
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Binary":
		var r Binary
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BiologicallyDerivedProduct":
		var r BiologicallyDerivedProduct
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BiologicallyDerivedProductDispense":
		var r BiologicallyDerivedProductDispense
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "BodyStructure":
		var r BodyStructure
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Bundle":
		var r Bundle
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CapabilityStatement":
		var r CapabilityStatement
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CarePlan":
		var r CarePlan
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CareTeam":
		var r CareTeam
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItem":
		var r ChargeItem
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ChargeItemDefinition":
		var r ChargeItemDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Citation":
		var r Citation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Claim":
		var r Claim
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClaimResponse":
		var r ClaimResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClinicalImpression":
		var r ClinicalImpression
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ClinicalUseDefinition":
		var r ClinicalUseDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CodeSystem":
		var r CodeSystem
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Communication":
		var r Communication
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CommunicationRequest":
		var r CommunicationRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CompartmentDefinition":
		var r CompartmentDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Composition":
		var r Composition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ConceptMap":
		var r ConceptMap
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Condition":
		var r Condition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ConditionDefinition":
		var r ConditionDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Consent":
		var r Consent
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Contract":
		var r Contract
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Coverage":
		var r Coverage
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityRequest":
		var r CoverageEligibilityRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "CoverageEligibilityResponse":
		var r CoverageEligibilityResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DetectedIssue":
		var r DetectedIssue
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Device":
		var r Device
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceAssociation":
		var r DeviceAssociation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceDefinition":
		var r DeviceDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceDispense":
		var r DeviceDispense
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceMetric":
		var r DeviceMetric
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceRequest":
		var r DeviceRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DeviceUsage":
		var r DeviceUsage
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DiagnosticReport":
		var r DiagnosticReport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "DocumentReference":
		var r DocumentReference
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Encounter":
		var r Encounter
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EncounterHistory":
		var r EncounterHistory
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Endpoint":
		var r Endpoint
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentRequest":
		var r EnrollmentRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EnrollmentResponse":
		var r EnrollmentResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EpisodeOfCare":
		var r EpisodeOfCare
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EventDefinition":
		var r EventDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Evidence":
		var r Evidence
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EvidenceReport":
		var r EvidenceReport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "EvidenceVariable":
		var r EvidenceVariable
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExampleScenario":
		var r ExampleScenario
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ExplanationOfBenefit":
		var r ExplanationOfBenefit
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "FamilyMemberHistory":
		var r FamilyMemberHistory
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Flag":
		var r Flag
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "FormularyItem":
		var r FormularyItem
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GenomicStudy":
		var r GenomicStudy
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Goal":
		var r Goal
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GraphDefinition":
		var r GraphDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Group":
		var r Group
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GuidanceResponse":
		var r GuidanceResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "HealthcareService":
		var r HealthcareService
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImagingSelection":
		var r ImagingSelection
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImagingStudy":
		var r ImagingStudy
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Immunization":
		var r Immunization
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationEvaluation":
		var r ImmunizationEvaluation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImmunizationRecommendation":
		var r ImmunizationRecommendation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ImplementationGuide":
		var r ImplementationGuide
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Ingredient":
		var r Ingredient
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InsurancePlan":
		var r InsurancePlan
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InventoryItem":
		var r InventoryItem
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InventoryReport":
		var r InventoryReport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Invoice":
		var r Invoice
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Library":
		var r Library
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Linkage":
		var r Linkage
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "List":
		var r List
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Location":
		var r Location
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ManufacturedItemDefinition":
		var r ManufacturedItemDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Measure":
		var r Measure
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MeasureReport":
		var r MeasureReport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Medication":
		var r Medication
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationAdministration":
		var r MedicationAdministration
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationDispense":
		var r MedicationDispense
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationKnowledge":
		var r MedicationKnowledge
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationRequest":
		var r MedicationRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicationStatement":
		var r MedicationStatement
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MedicinalProductDefinition":
		var r MedicinalProductDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageDefinition":
		var r MessageDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MessageHeader":
		var r MessageHeader
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "MolecularSequence":
		var r MolecularSequence
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NamingSystem":
		var r NamingSystem
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NutritionIntake":
		var r NutritionIntake
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NutritionOrder":
		var r NutritionOrder
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "NutritionProduct":
		var r NutritionProduct
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Observation":
		var r Observation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ObservationDefinition":
		var r ObservationDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationDefinition":
		var r OperationDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OperationOutcome":
		var r OperationOutcome
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Organization":
		var r Organization
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "OrganizationAffiliation":
		var r OrganizationAffiliation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PackagedProductDefinition":
		var r PackagedProductDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Parameters":
		var r Parameters
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Patient":
		var r Patient
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentNotice":
		var r PaymentNotice
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PaymentReconciliation":
		var r PaymentReconciliation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Permission":
		var r Permission
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Person":
		var r Person
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PlanDefinition":
		var r PlanDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Practitioner":
		var r Practitioner
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "PractitionerRole":
		var r PractitionerRole
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Procedure":
		var r Procedure
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Provenance":
		var r Provenance
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Questionnaire":
		var r Questionnaire
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "QuestionnaireResponse":
		var r QuestionnaireResponse
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RegulatedAuthorization":
		var r RegulatedAuthorization
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RelatedPerson":
		var r RelatedPerson
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RequestOrchestration":
		var r RequestOrchestration
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Requirements":
		var r Requirements
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchStudy":
		var r ResearchStudy
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ResearchSubject":
		var r ResearchSubject
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "RiskAssessment":
		var r RiskAssessment
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Schedule":
		var r Schedule
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SearchParameter":
		var r SearchParameter
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ServiceRequest":
		var r ServiceRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Slot":
		var r Slot
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Specimen":
		var r Specimen
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SpecimenDefinition":
		var r SpecimenDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureDefinition":
		var r StructureDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "StructureMap":
		var r StructureMap
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Subscription":
		var r Subscription
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubscriptionStatus":
		var r SubscriptionStatus
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubscriptionTopic":
		var r SubscriptionTopic
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Substance":
		var r Substance
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceDefinition":
		var r SubstanceDefinition
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceNucleicAcid":
		var r SubstanceNucleicAcid
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstancePolymer":
		var r SubstancePolymer
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceProtein":
		var r SubstanceProtein
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceReferenceInformation":
		var r SubstanceReferenceInformation
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubstanceSourceMaterial":
		var r SubstanceSourceMaterial
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyDelivery":
		var r SupplyDelivery
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SupplyRequest":
		var r SupplyRequest
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Task":
		var r Task
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TerminologyCapabilities":
		var r TerminologyCapabilities
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestPlan":
		var r TestPlan
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestReport":
		var r TestReport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "TestScript":
		var r TestScript
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Transport":
		var r Transport
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "ValueSet":
		var r ValueSet
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VerificationResult":
		var r VerificationResult
		err := r.unmarshalJSON(d)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "VisionPrescription":
		var r VisionPrescription
		err := r.unmarshalJSON(d)
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
	if len(start.Name.Local) > 0 && unicode.IsUpper(rune(start.Name.Local[0])) {
		return e.Encode(r.Resource)
	} else {
		err := e.EncodeToken(start)
		if err != nil {
			return err
		}
		err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
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
	case "ActorDefinition":
		var r ActorDefinition
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "AdministrableProductDefinition":
		var r AdministrableProductDefinition
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
	case "ArtifactAssessment":
		var r ArtifactAssessment
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
	case "BiologicallyDerivedProductDispense":
		var r BiologicallyDerivedProductDispense
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
	case "Citation":
		var r Citation
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
	case "ClinicalUseDefinition":
		var r ClinicalUseDefinition
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
	case "ConditionDefinition":
		var r ConditionDefinition
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
	case "DeviceAssociation":
		var r DeviceAssociation
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
	case "DeviceDispense":
		var r DeviceDispense
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
	case "DeviceUsage":
		var r DeviceUsage
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
	case "DocumentReference":
		var r DocumentReference
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
	case "EncounterHistory":
		var r EncounterHistory
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
	case "EvidenceReport":
		var r EvidenceReport
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
	case "FormularyItem":
		var r FormularyItem
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "GenomicStudy":
		var r GenomicStudy
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
	case "ImagingSelection":
		var r ImagingSelection
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
	case "Ingredient":
		var r Ingredient
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
	case "InventoryItem":
		var r InventoryItem
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "InventoryReport":
		var r InventoryReport
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
	case "ManufacturedItemDefinition":
		var r ManufacturedItemDefinition
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
	case "MedicinalProductDefinition":
		var r MedicinalProductDefinition
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
	case "NutritionIntake":
		var r NutritionIntake
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
	case "NutritionProduct":
		var r NutritionProduct
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
	case "PackagedProductDefinition":
		var r PackagedProductDefinition
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
	case "Permission":
		var r Permission
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
	case "RegulatedAuthorization":
		var r RegulatedAuthorization
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
	case "RequestOrchestration":
		var r RequestOrchestration
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "Requirements":
		var r Requirements
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
	case "SubscriptionStatus":
		var r SubscriptionStatus
		err := d.DecodeElement(&r, &start)
		if err != nil {
			return err
		}
		*cr = ContainedResource{r}
		return nil
	case "SubscriptionTopic":
		var r SubscriptionTopic
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
	case "SubstanceDefinition":
		var r SubstanceDefinition
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
	case "TestPlan":
		var r TestPlan
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
	case "Transport":
		var r Transport
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
