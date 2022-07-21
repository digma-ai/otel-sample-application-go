module github.com/digma-ai/otel-sample-application-go/src/authenticator

go 1.17

require (
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect; indirectus
)

replace github.com/digma-ai/otel-sample-application-go/src/otelconfigure => ../otelconfigure
