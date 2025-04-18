package fhirpath_test

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/cockroachdb/apd/v3"
	"testing"
	"time"
)

func TestDateArithmetic(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		date     fhirpath.Date
		quantity fhirpath.Quantity
		wantAdd  fhirpath.Date
		wantSub  fhirpath.Date
		wantErr  bool
	}{
		{
			name: "add/subtract one year",
			date: fhirpath.Date{
				Value:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantAdd: fhirpath.Date{
				Value:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			wantSub: fhirpath.Date{
				Value:     time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
		},
		{
			name: "add/subtract one month with month end adjustment",
			date: fhirpath.Date{
				Value:     time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "month",
			},
			wantAdd: fhirpath.Date{
				Value:     time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			wantSub: fhirpath.Date{
				Value:     time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
		},
		{
			name: "add/subtract one week",
			date: fhirpath.Date{
				Value:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "week",
			},
			wantAdd: fhirpath.Date{
				Value:     time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			wantSub: fhirpath.Date{
				Value:     time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
		},
		{
			name: "invalid unit",
			date: fhirpath.Date{
				Value:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "hour",
			},
			wantErr: true,
		},
		{
			name: "empty date",
			date: fhirpath.Date{},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantErr: true,
		},
		{
			name: "decimal year",
			date: fhirpath.Date{
				Value:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(15, -1)}, // 1.5 years
				Unit:  "year",
			},
			wantAdd: fhirpath.Date{
				Value:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
			wantSub: fhirpath.Date{
				Value:     time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DatePrecisionFull,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test addition
			got, err := tt.date.Add(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantAdd {
				t.Errorf("Add() = %v, want %v", got, tt.wantAdd)
			}

			// Test subtraction
			got, err = tt.date.Subtract(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantSub {
				t.Errorf("Subtract() = %v, want %v", got, tt.wantSub)
			}
		})
	}
}

func TestTimeArithmetic(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		time     fhirpath.Time
		quantity fhirpath.Quantity
		wantAdd  fhirpath.Time
		wantSub  fhirpath.Time
		wantErr  bool
	}{
		{
			name: "add/subtract one hour",
			time: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "hour",
			},
			wantAdd: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 13, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			wantSub: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 11, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
		},
		{
			name: "add/subtract one minute",
			time: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "minute",
			},
			wantAdd: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 1, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			wantSub: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 11, 59, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
		},
		{
			name: "add/subtract one second",
			time: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "second",
			},
			wantAdd: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 1, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			wantSub: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 11, 59, 59, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
		},
		{
			name: "invalid unit",
			time: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantErr: true,
		},
		{
			name: "empty time",
			time: fhirpath.Time{},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "hour",
			},
			wantErr: true,
		},
		{
			name: "millisecond precision",
			time: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1500, 0)},
				Unit:  "millisecond",
			},
			wantAdd: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 12, 0, 1, 500000000, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
			wantSub: fhirpath.Time{
				Value:     time.Date(0, 1, 1, 11, 59, 58, 500000000, time.UTC),
				Precision: fhirpath.TimePrecisionFull,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test addition
			got, err := tt.time.Add(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantAdd {
				t.Errorf("Add() = %v, want %v", got, tt.wantAdd)
			}

			// Test subtraction
			got, err = tt.time.Subtract(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantSub {
				t.Errorf("Subtract() = %v, want %v", got, tt.wantSub)
			}
		})
	}
}

func TestDateTimeArithmetic(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		datetime fhirpath.DateTime
		quantity fhirpath.Quantity
		wantAdd  fhirpath.DateTime
		wantSub  fhirpath.DateTime
		wantErr  bool
	}{
		{
			name: "add/subtract one year",
			datetime: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantAdd: fhirpath.DateTime{
				Value:     time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			wantSub: fhirpath.DateTime{
				Value:     time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
		},
		{
			name: "add/subtract one hour",
			datetime: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "hour",
			},
			wantAdd: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 13, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			wantSub: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
		},
		{
			name: "partial date - year only",
			datetime: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionYear,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantAdd: fhirpath.DateTime{
				Value:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionYear,
			},
			wantSub: fhirpath.DateTime{
				Value:     time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionYear,
			},
		},
		{
			name: "month end adjustment",
			datetime: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 31, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "month",
			},
			wantAdd: fhirpath.DateTime{
				Value:     time.Date(2020, 2, 29, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			wantSub: fhirpath.DateTime{
				Value:     time.Date(2019, 12, 31, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
		},
		{
			name:     "empty datetime",
			datetime: fhirpath.DateTime{},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1, 0)},
				Unit:  "year",
			},
			wantErr: true,
		},
		{
			name: "millisecond precision",
			datetime: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			quantity: fhirpath.Quantity{
				Value: fhirpath.Decimal{Value: apd.New(1500, 0)},
				Unit:  "millisecond",
			},
			wantAdd: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 12, 0, 1, 500000000, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
			wantSub: fhirpath.DateTime{
				Value:     time.Date(2020, 1, 1, 11, 59, 58, 500000000, time.UTC),
				Precision: fhirpath.DateTimePrecisionFull,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test addition
			got, err := tt.datetime.Add(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantAdd {
				t.Errorf("Add() = %v, want %v", got, tt.wantAdd)
			}

			// Test subtraction
			got, err = tt.datetime.Subtract(ctx, tt.quantity)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got != tt.wantSub {
				t.Errorf("Subtract() = %v, want %v", got, tt.wantSub)
			}
		})
	}
}
