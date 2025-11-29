package overflow

import (
	"math"
	"testing"
)

type binCase[T Signed] struct {
	name   string
	a, b   T
	want   T
	wantOK bool
}

func runBinCases[T Signed](t *testing.T, cases []binCase[T], op func(T, T) (T, bool)) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := op(tc.a, tc.b)
			if ok != tc.wantOK || got != tc.want {
				t.Fatalf("got (%v, %v), want (%v, %v)", got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	int32Cases := []binCase[int32]{
		{"simple", 1, 1, 2, true},
		{"overflow", math.MaxInt32, 1, math.MinInt32, false},
	}
	int64Cases := []binCase[int64]{
		{"simple", 5, -3, 2, true},
		{"overflow", math.MaxInt64, 1, math.MinInt64, false},
	}
	runBinCases(t, int32Cases, Add[int32])
	runBinCases(t, int64Cases, Add[int64])
}

func TestSub(t *testing.T) {
	int32Cases := []binCase[int32]{
		{"simple", 5, 3, 2, true},
		{"overflow", math.MinInt32, 1, math.MaxInt32, false},
	}
	int64Cases := []binCase[int64]{
		{"simple", 3, -5, 8, true},
		{"overflow", math.MinInt64, 1, math.MaxInt64, false},
	}
	runBinCases(t, int32Cases, Sub[int32])
	runBinCases(t, int64Cases, Sub[int64])
}

func TestMul(t *testing.T) {
	int32Cases := []binCase[int32]{
		{"zero", 0, 5, 0, true},
		{"simple", 6, -7, -42, true},
		{"overflow", math.MaxInt32, 2, -2, false},
	}
	int64Cases := []binCase[int64]{
		{"zero", 0, 123, 0, true},
		{"simple", -9, -9, 81, true},
		{"overflow", math.MaxInt64, 2, -2, false},
	}
	runBinCases(t, int32Cases, Mul[int32])
	runBinCases(t, int64Cases, Mul[int64])
}

func TestDiv(t *testing.T) {
	int32Cases := []binCase[int32]{
		{"simple", 10, 2, 5, true},
		{"zero-divisor", 5, 0, 0, false},
		{"min-overflow", math.MinInt32, -1, math.MinInt32, false},
	}
	int64Cases := []binCase[int64]{
		{"simple", -10, -2, 5, true},
		{"zero-divisor", 0, 0, 0, false},
		{"min-overflow", math.MinInt64, -1, math.MinInt64, false},
	}
	runBinCases(t, int32Cases, Div[int32])
	runBinCases(t, int64Cases, Div[int64])
}

func TestMod(t *testing.T) {
	int32Cases := []binCase[int32]{
		{"simple", 10, 3, 1, true},
		{"zero-divisor", 5, 0, 0, false},
		{"min-overflow", math.MinInt32, -1, 0, false},
	}
	int64Cases := []binCase[int64]{
		{"simple", -10, 3, -1, true},
		{"zero-divisor", 0, 0, 0, false},
		{"min-overflow", math.MinInt64, -1, 0, false},
	}
	runBinCases(t, int32Cases, Mod[int32])
	runBinCases(t, int64Cases, Mod[int64])
}
