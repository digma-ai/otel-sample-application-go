package authenticator

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Authenticate(ctx context.Context, success bool) bool {
	tracer := otel.GetTracerProvider().Tracer("authenticator")
	_, span := tracer.Start(ctx, "Authenticate")
	defer span.End(trace.WithStackTrace(true))
	if !success {
		panic(errors.New("dont panic"))
	}
	return true
}
