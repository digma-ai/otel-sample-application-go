package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	digmamux "github.com/digma-ai/otel-go-instrumentation/mux"
	"github.com/digma-ai/otel-sample-application-go/src/otelconfigure"
	domain "github.com/digma-ai/otel-sample-application-go/src/mux-sample/user"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	port    = "8011"
	appName = "user-microservice"
)

func main() {
	// injected latency
	if s := os.Getenv("EXTRA_LATENCY"); s != "" {
		v, err := time.ParseDuration(s)
		if err != nil {
			log.Fatalf("failed to parse EXTRA_LATENCY (%s) as time.Duration: %+v", v, err) //%+v: variant will include the struct’s field names.
		}
		domain.ExtraLatency = v
		log.Printf("extra latency enabled (duration: %v)", v)
	} else {
		domain.ExtraLatency = time.Duration(0)
	}

	shutdown := otelconfigure.InitTracer("user-service", []string{
		"github.com/digma-ai/otel-sample-application-go/src/authenticator",
		"github.com/digma-ai/otel-sample-application-go/src/otelconfigure",
	})
	defer shutdown()

	service := domain.NewUserService()
	service.Init()
	controller := *domain.NewUserController(service)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(otelmux.Middleware(appName))
	router.Use(digmamux.Middleware(router))
	//router.Use(handlers.RecoveryHandler())

	router.HandleFunc("/hc", healthCheck).Methods("GET")
	router.HandleFunc("/users", controller.Add).Methods("POST")
	router.HandleFunc("/users/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/users", controller.All).Methods("GET")

	fmt.Println("listening on :" + port)
	err := http.ListenAndServe(":"+port, router)
	handleErr(err, "failed to listen & serve")
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	tracer := otel.GetTracerProvider().Tracer("Global")
	_, span := tracer.Start(req.Context(), getCurrentFuncName())
	defer span.End(trace.WithStackTrace(true))
	if span != nil {
		panic("dont panic")
	}

	setResponse(w, "ok", http.StatusOK)
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func setResponse(w http.ResponseWriter, output interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(output)
}
func getCurrentFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		return fn.Name()
	}
	return ""
}
