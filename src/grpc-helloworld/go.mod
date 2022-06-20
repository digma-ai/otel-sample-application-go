module github.com/digma-ai/otel-sample-application-go/grpc-helloworld

go 1.18

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.32.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	golang.org/x/net v0.0.0-20220615171555-694bf12d69de
	google.golang.org/grpc v1.47.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/digma-ai/otel-go-instrumentation v0.0.0-20220526013256-80db951cfb9b // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/proto/otlp v0.16.0 // indirect
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220616135557-88e70c0c3a90 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

require (
	github.com/digma-ai/otel-go-instrumentation/grpc v1.0.0
	github.com/digma-ai/otel-sample-application-go/src/otelconfigure v0.0.1
)

replace github.com/digma-ai/otel-sample-application-go/src/otelconfigure => ../otelconfigure

replace github.com/digma-ai/otel-go-instrumentation => ../../../otel-go-instrumentation

replace github.com/digma-ai/otel-go-instrumentation/grpc => ../../../otel-go-instrumentation/grpc
