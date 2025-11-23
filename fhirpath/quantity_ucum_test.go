package fhirpath

import (
	"context"
	"testing"

	"github.com/cockroachdb/apd/v3"
)

func TestConvertQuantityToUnit(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := []struct {
		name     string
		value    Quantity
		target   String
		wantUnit String
		want     string
	}{
		{
			name:     "mass mg to g",
			value:    Quantity{Value: mustDecimal(t, "4000"), Unit: "mg"},
			target:   "g",
			wantUnit: "g",
			want:     "4",
		},
		{
			name:     "mass mg to g precise",
			value:    Quantity{Value: mustDecimal(t, "4040"), Unit: "mg"},
			target:   "g",
			wantUnit: "g",
			want:     "4.04",
		},
		{
			name:     "week to day",
			value:    Quantity{Value: mustDecimal(t, "1"), Unit: "week"},
			target:   "day",
			wantUnit: canonicalQuantityUnit("day"),
			want:     "7",
		},
	}

	for _, tc := range testCases {
		ct := tc
		t.Run(ct.name, func(t *testing.T) {
			t.Parallel()
			converted, err := convertQuantityToUnit(ctx, ct.value, ct.target)
			if err != nil {
				t.Fatalf("convertQuantityToUnit returned error: %v", err)
			}

			if converted.Unit != ct.wantUnit {
				t.Fatalf("expected unit %s, got %s", ct.wantUnit, converted.Unit)
			}

			wantValue := mustAPDDecimal(t, ct.want)
			if converted.Value.Value.Cmp(wantValue) != 0 {
				t.Fatalf("expected value %s, got %s", wantValue, converted.Value.Value)
			}
		})
	}
}

func mustDecimal(t *testing.T, in string) Decimal {
	t.Helper()
	d, _, err := apd.NewFromString(in)
	if err != nil {
		t.Fatalf("failed to parse decimal %q: %v", in, err)
	}
	return Decimal{Value: d}
}

func mustAPDDecimal(t *testing.T, in string) *apd.Decimal {
	t.Helper()
	d, _, err := apd.NewFromString(in)
	if err != nil {
		t.Fatalf("failed to parse decimal %q: %v", in, err)
	}
	return d
}

func TestQuantityEquivalentIgnoresPrecision(t *testing.T) {
	left := Quantity{Value: mustDecimal(t, "4"), Unit: "g"}
	right := Quantity{Value: mustDecimal(t, "4040"), Unit: "mg"}

	if eq, ok := left.Equal(right); ok && eq {
		t.Fatalf("expected equality to be false")
	}
	if !left.Equivalent(right) {
		t.Fatalf("expected quantities to be equivalent")
	}
}

func TestQuantityCalendarEqualityAndEquivalence(t *testing.T) {
	yearCalendar, err := ParseQuantity("1 year")
	if err != nil {
		t.Fatalf("ParseQuantity: %v", err)
	}
	yearDefinite, err := ParseQuantity("1 'a'")
	if err != nil {
		t.Fatalf("ParseQuantity: %v", err)
	}

	if eq, ok := yearCalendar.Equal(yearDefinite); ok {
		t.Fatalf("expected calendar vs definite equality to be non-comparable, got eq=%v", eq)
	}
	if !yearCalendar.Equivalent(yearDefinite) {
		t.Fatalf("expected calendar vs definite equivalence to be true")
	}

	if twoYear, _ := ParseQuantity("2 year"); twoYear.Equivalent(yearDefinite) {
		t.Fatalf("expected differing values to be non-equivalent")
	}
}

func TestQuantityCalendarSecondsEquality(t *testing.T) {
	secCalendar, err := ParseQuantity("1 second")
	if err != nil {
		t.Fatalf("ParseQuantity: %v", err)
	}
	secDefinite, err := ParseQuantity("1 's'")
	if err != nil {
		t.Fatalf("ParseQuantity: %v", err)
	}

	if eq, ok := secCalendar.Equal(secDefinite); !ok || !eq {
		t.Fatalf("expected seconds comparison to be true, eq=%v ok=%v", eq, ok)
	}
	if !secCalendar.Equivalent(secDefinite) {
		t.Fatalf("expected seconds equivalence to be true")
	}
}
