package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/digma-ai/otel-sample-application-go/src/otelconfigure"
	domain "github.com/digma-ai/otel-sample-application-go/src/userservice/user"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
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

	shutdown := otelconfigure.InitTracer("userservice", []string{
		"github.com/digma-ai/otel-sample-application-go/src/authenticator",
		"github.com/digma-ai/otel-sample-application-go/src/otelconfigure",
	})
	defer shutdown()

	tracer := otel.Tracer(appName)

	service := domain.NewUserService()
	service.Init()
	controller := domain.NewUserController(service, tracer)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(otelmux.Middleware(appName))
	router.HandleFunc("/users", controller.Add).Methods("POST")
	router.HandleFunc("/users/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/users", controller.All).Methods("GET")

	fmt.Println("listening on :" + port)
	err := http.ListenAndServe(":"+port, router)
	handleErr(err, "failed to listen & serve")
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
