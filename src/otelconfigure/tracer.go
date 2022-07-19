package otelconfigure

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/digma-ai/otel-go-instrumentation/detector"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func InitTracer(serviceName string, otherImportPaths []string) func() {
	return InitTracerWithModule(serviceName, "", "", otherImportPaths)
}

func InitTracerWithModule(serviceName string, moduleImportPath string, modulePath string, otherImportPaths []string) func() {
	otlpAddress, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otlpAddress = "localhost:4317"
	}

	ctx := context.Background()

	res, err := resource.New(ctx,
		//resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(serviceName),
			semconv.TelemetrySDKLanguageGo,
		),
		/*
			Resources can also be detected automatically through resource.Detector implementations.
			These Detectors may discover information about the currently running process, the operating system it is running on, the cloud provider hosting that operating system instance, or any number of other resource attributes.
		*/
		resource.WithDetectors(
			&detector.DigmaDetector{
				DeploymentEnvironment:  os.Getenv("DEPLOYMENT_ENV"),
				CommitId:               "",
				ModuleImportPath:       moduleImportPath,
				ModulePath:             modulePath,
				OtherModulesImportPath: otherImportPaths,
			},
		))

	handleErr(err, "failed to create resource")

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otlpAddress),
		otlptracegrpc.WithReconnectionPeriod(2*time.Second),
		//otlptracegrpc.WithDialOption(grpc.WithBlock()
	)

	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	traceExporter, err := otlptrace.New(cancelCtx, traceClient)
	handleErr(err, "failed to create trace exporter")

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter))

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return func() {
		log.Println("TracerProvider: shutting down...")
		// Shutdown will flush any remaining spans and shut down the exporter.
		handleErr(tracerProvider.Shutdown(context.Background()), "failed to shutdown TracerProvider")
	}
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
