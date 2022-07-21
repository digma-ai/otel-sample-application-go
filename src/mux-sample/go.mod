module github.com/digma-ai/otel-sample-application-go/src/mux-sample

go 1.17

require (
	github.com/digma-ai/otel-go-instrumentation/mux v1.0.1
	github.com/digma-ai/otel-sample-application-go/src/authenticator v0.0.1
	github.com/digma-ai/otel-sample-application-go/src/otelconfigure v0.0.1
	github.com/gorilla/mux v1.8.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.32.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/trace v1.8.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/puddle v1.2.2-0.20220404125616-4e959849469a // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
)

replace github.com/digma-ai/otel-sample-application-go/src/otelconfigure => ../otelconfigure

replace github.com/digma-ai/otel-sample-application-go/src/authenticator => ../authenticator

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/digma-ai/otel-go-instrumentation v1.0.10 // indirect
	github.com/exaring/otelpgx v0.1.0
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/jackc/pgx/v5 v5.0.0-alpha.5
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/proto/otlp v0.16.0 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220608133413-ed9918b62aac // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
