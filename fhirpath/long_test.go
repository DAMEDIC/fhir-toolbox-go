package fhirpath

import (
	"context"
	"testing"
)

func TestEvaluateLongLiteral(t *testing.T) {
	expr := MustParse("1L")
	result, err := Evaluate(context.Background(), testElement{}, expr)
	if err != nil {
		t.Fatalf("Evaluate returned error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected single result, got %d", len(result))
	}
	val, ok := result[0].(Long)
	if !ok {
		t.Fatalf("expected Long, got %T", result[0])
	}
	if val != 1 {
		t.Fatalf("expected 1L, got %v", val)
	}
}
