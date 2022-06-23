# otel-sample-application-go
This repository contains examples for using digma with go:
- [mux-sample](https://github.com/digma-ai/otel-sample-application-go/tree/main/src/mux-sample) for [gorilla/mux](https://github.com/gorilla/mux)
- [grpc-sample](https://github.com/digma-ai/otel-sample-application-go/tree/main/src/grpc-sample) for [grpc/grpc-go](https://github.com/grpc/grpc-go)
- [echo-sample](https://github.com/digma-ai/otel-sample-application-go/tree/main/src/echo-sample) for [labstack/echo](https://github.com/labstack/echo)

Open the it in VSCode via the `workspace.code-workspace` file to try it out.

---
### Branches
#### `main`
Contains only the above samples and uses the latest released version of [otel-go-instrumentation](https://github.com/digma-ai/otel-go-instrumentation).
#### `dev`
Contains the above samples and point to a local [otel-go-instrumentation](https://github.com/digma-ai/otel-go-instrumentation) module (need to be cloneed and placed beside this repo), and replaces all the [otel-go-instrumentation](https://github.com/digma-ai/otel-go-instrumentation) imports with the local one.
