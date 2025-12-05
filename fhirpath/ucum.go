package fhirpath

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/apd/v3"
	"github.com/iimos/ucum"
	"github.com/iimos/ucum/ucumapd"
)

const defaultUCUMPrecision uint32 = 28

var defaultUCUMConverter = ucumapd.NewConverter(ucum.NewConverter())

type ucumConverterKey struct{}

// WithUCUMConverter installs a custom UCUM converter into the context. The converter
// is used for quantity conversions during FHIRPath evaluation.
func WithUCUMConverter(ctx context.Context, converter ucumapd.Converter) context.Context {
	if converter == nil {
		return ctx
	}
	if ctx == nil {
		return context.WithValue(context.Background(), ucumConverterKey{}, converter)
	}
	return context.WithValue(ctx, ucumConverterKey{}, converter)
}

func ucumConverter(ctx context.Context) ucumapd.Converter {
	if ctx != nil {
		if converter, ok := ctx.Value(ucumConverterKey{}).(ucumapd.Converter); ok && converter != nil {
			return converter
		}
	}
	return defaultUCUMConverter
}

func convertDecimalUnit(ctx context.Context, value *apd.Decimal, fromUnit, toUnit string) (*apd.Decimal, error) {
	from := strings.TrimSpace(fromUnit)
	to := strings.TrimSpace(toUnit)

	if from == "" {
		from = "1"
	}
	if to == "" {
		to = "1"
	}
	if from == to {
		return value, nil
	}
	if from == "1" || to == "1" {
		return nil, fmt.Errorf("cannot convert between dimensionless units %q and %q", from, to)
	}

	converter := ucumConverter(ctx)
	converted, err := converter.ConvDecimal(value, from, to, ucumAPDContext(ctx))
	if err != nil {
		return nil, err
	}
	return converted, nil
}

func ucumAPDContext(ctx context.Context) *apd.Context {
	base := apdContext(ctx)
	if base.Precision > 0 {
		return base
	}
	return base.WithPrecision(defaultUCUMPrecision)
}

func canonicalUCUMUnit(unit string) string {
	trimmed := strings.TrimSpace(unit)
	if strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}") && len(trimmed) > 2 {
		trimmed = trimmed[1 : len(trimmed)-1]
	}
	lower := strings.ToLower(trimmed)

	switch lower {
	case UnitYear, UnitYears:
		return "a"
	case UnitMonth, UnitMonths:
		return "mo"
	case UnitWeek, UnitWeeks:
		return "wk"
	case UnitDay, UnitDays:
		return "d"
	case UnitHour, UnitHours:
		return "h"
	case UnitMinute, UnitMinutes:
		return "min"
	case UnitSecond, UnitSeconds, UnitS:
		return "s"
	case UnitMillisecond, UnitMilliseconds, UnitMs:
		return "ms"
	default:
		return trimmed
	}
}

func displayQuantityUnit(unit String) string {
	s := strings.TrimSpace(string(unit))
	if s == "" {
		return s
	}
	if s == "1" {
		return s
	}
	if isCalendarLiteralUnit(unit) {
		return s
	}
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		return s
	}
	if _, err := ucum.Parse([]byte(s)); err == nil {
		return s
	}
	return fmt.Sprintf("{%s}", s)
}
