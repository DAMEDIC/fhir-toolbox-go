package overflow

import (
	"math"
	"testing"
)

func TestAdd32(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int32
		want   int32
		wantOk bool
	}{
		{"zero_add", 0, 0, 0, true},
		{"positive_add", 5, 3, 8, true},
		{"negative_add", -5, -3, -8, true},
		{"mixed_add", -5, 8, 3, true},
		{"max_overflow", math.MaxInt32, 1, math.MinInt32, false},
		{"min_overflow", math.MinInt32, -1, math.MaxInt32, false},
		{"near_max", math.MaxInt32 - 1, 1, math.MaxInt32, true},
		{"near_min", math.MinInt32 + 1, -1, math.MinInt32, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Add32(tt.a, tt.b)
			if got != tt.want || ok != tt.wantOk {
				t.Errorf("Add32(%d, %d) = (%d, %v), want (%d, %v)",
					tt.a, tt.b, got, ok, tt.want, tt.wantOk)
			}
		})
	}
}

func TestSub32(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int32
		want   int32
		wantOk bool
	}{
		{"zero_sub", 0, 0, 0, true},
		{"positive_sub", 5, 3, 2, true},
		{"negative_sub", -5, -3, -2, true},
		{"mixed_sub", -5, 8, -13, true},
		{"max_overflow", math.MaxInt32, math.MinInt32, -1, false},
		{"min_overflow", math.MinInt32, 1, math.MaxInt32, false},
		{"near_max", math.MaxInt32 - 1, -1, math.MaxInt32, true},
		{"near_min", math.MinInt32 + 1, 1, math.MinInt32, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Sub32(tt.a, tt.b)
			if got != tt.want || ok != tt.wantOk {
				t.Errorf("Sub32(%d, %d) = (%d, %v), want (%d, %v)",
					tt.a, tt.b, got, ok, tt.want, tt.wantOk)
			}
		})
	}
}

func TestMul32(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int32
		want   int32
		wantOk bool
	}{
		{"zero_mul", 0, 5, 0, true},
		{"positive_mul", 5, 3, 15, true},
		{"negative_mul", -5, -3, 15, true},
		{"mixed_mul", -5, 3, -15, true},
		{"max_overflow", math.MaxInt32, 2, -2, false},
		{"min_overflow", math.MinInt32, 2, 0, false},
		{"near_max", math.MaxInt32 / 2, 2, math.MaxInt32 - 1, true},
		{"near_min", math.MinInt32 / 2, 2, math.MinInt32, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Mul32(tt.a, tt.b)
			if got != tt.want || ok != tt.wantOk {
				t.Errorf("Mul32(%d, %d) = (%d, %v), want (%d, %v)",
					tt.a, tt.b, got, ok, tt.want, tt.wantOk)
			}
		})
	}
}

func TestDiv32(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int32
		want   int32
		wantOk bool
	}{
		{"simple_div", 6, 2, 3, true},
		{"negative_div", -6, 2, -3, true},
		{"zero_dividend", 0, 5, 0, true},
		{"zero_divisor", 5, 0, 0, false},
		{"min_by_minus_one", math.MinInt32, -1, math.MinInt32, false},
		{"max_div", math.MaxInt32, 1, math.MaxInt32, true},
		{"min_div", math.MinInt32, 2, math.MinInt32 / 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Div32(tt.a, tt.b)
			if got != tt.want || ok != tt.wantOk {
				t.Errorf("Div32(%d, %d) = (%d, %v), want (%d, %v)",
					tt.a, tt.b, got, ok, tt.want, tt.wantOk)
			}
		})
	}
}

func TestMod32(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int32
		want   int32
		wantOk bool
	}{
		{"simple_mod", 7, 3, 1, true},
		{"negative_dividend", -7, 3, -1, true},
		{"negative_divisor", 7, -3, 1, true},
		{"both_negative", -7, -3, -1, true},
		{"zero_dividend", 0, 5, 0, true},
		{"zero_divisor", 5, 0, 0, false},
		{"min_by_minus_one", math.MinInt32, -1, 0, false},
		{"max_mod", math.MaxInt32, 2, 1, true},
		{"min_mod", math.MinInt32, 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Mod32(tt.a, tt.b)
			if got != tt.want || ok != tt.wantOk {
				t.Errorf("Mod32(%d, %d) = (%d, %v), want (%d, %v)",
					tt.a, tt.b, got, ok, tt.want, tt.wantOk)
			}
		})
	}
}
