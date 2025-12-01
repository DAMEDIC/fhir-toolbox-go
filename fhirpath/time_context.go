package fhirpath

import (
	"context"
	"time"
)

type evaluationInstantKey struct{}

// withEvaluationInstant ensures the context carries the evaluation anchor time.
// Callers should invoke this once per evaluation so that temporal functions can
// behave deterministically per FHIRPath spec (FHIRPath v3.0.0, Utility Functions:
// now()/timeOfDay()/today() must reuse the same instant for an expression).
func withEvaluationInstant(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if _, ok := ctx.Value(evaluationInstantKey{}).(time.Time); ok {
		return ctx
	}
	return context.WithValue(ctx, evaluationInstantKey{}, currentEvaluationInstant())
}

// WithEvaluationTime allows callers to inject a fixed instant (useful for deterministic tests).
func WithEvaluationTime(ctx context.Context, instant time.Time) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, evaluationInstantKey{}, instant.Truncate(time.Millisecond))
}

func evaluationInstant(ctx context.Context) time.Time {
	if ctx == nil {
		return currentEvaluationInstant()
	}
	if instant, ok := ctx.Value(evaluationInstantKey{}).(time.Time); ok {
		return instant
	}
	return currentEvaluationInstant()
}

func currentEvaluationInstant() time.Time {
	return time.Now().Truncate(time.Millisecond)
}
