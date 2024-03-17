package types

import "encoding/json"

// Base StructureDefinition for SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
//
// There is a need for a concise way to handle the data produced by devices that sample a physical state at a high frequency.
type SampledData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *Decimal
	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *Decimal
	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Dimensions PositiveInt
	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin Quantity
	// The length of time between sampling times, measured in milliseconds.
	Period Decimal
	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *Decimal
	// A series of data points which are decimal values separated by a single space (character u20). The special values "E" (error), "L" (below detection limit) and "U" (above detection limit) can also be used in place of a decimal value.
	Data *String
}

func (s SampledData) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
