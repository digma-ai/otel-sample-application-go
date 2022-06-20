module github.com/digma-ai/otel-sample-application-go/src/authservice

go 1.18

require (
	github.com/digma-ai/otel-go-instrumentation/echo v0.0.0-00010101000000-000000000000
	github.com/digma-ai/otel-sample-application-go/src/otelconfigure v0.0.1
	github.com/labstack/echo/v4 v4.7.2
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.32.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

replace github.com/digma-ai/otel-go-instrumentation => ../../../otel-go-instrumentation

replace github.com/digma-ai/otel-go-instrumentation/echo => ../../../otel-go-instrumentation/echo

replace github.com/digma-ai/otel-sample-application-go/src/otelconfigure => ../otelconfigure

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/digma-ai/otel-go-instrumentation v0.0.0-20220526013256-80db951cfb9b // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/proto/otlp v0.16.0 // indirect
	golang.org/x/crypto v0.0.0-20220517005047-85d78b3ac167 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	google.golang.org/genproto v0.0.0-20220608133413-ed9918b62aac // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
