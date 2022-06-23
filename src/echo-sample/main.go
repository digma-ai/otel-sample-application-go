package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	digmaecho "github.com/digma-ai/otel-go-instrumentation/echo"
	"github.com/digma-ai/otel-sample-application-go/src/echo-sample/auth"
	"github.com/digma-ai/otel-sample-application-go/src/otelconfigure"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	//"example.com/authentication/infrastructure/opentelemetry"
)

var (
	port    = "8011"
	appName = "auth-service"
)

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello!")
}

func main() {
	// injected latency
	if s := os.Getenv("EXTRA_LATENCY"); s != "" {
		v, err := time.ParseDuration(s)
		if err != nil {
			log.Fatalf("failed to parse EXTRA_LATENCY (%s) as time.Duration: %+v", v, err) //%+v: variant will include the struct’s field names.
		}
		auth.ExtraLatency = v
		log.Printf("extra latency enabled (duration: %v)", v)
	} else {
		auth.ExtraLatency = time.Duration(0)
	}

	shutdown := otelconfigure.InitTracer("authentication-service", []string{
		"github.com/digma-ai/otel-sample-application-go/src/otelconfigure",
	})
	defer shutdown()

	tracer := otel.Tracer(appName)

	service := auth.NewAuthService()
	service.Init()
	controller := auth.NewAuthController(service, tracer)

	r := echo.New()
	r.Use(otelecho.Middleware(appName))
	r.Use(digmaecho.Middleware())
	r.GET("/", index)
	//r.GET("/auth", auth.Index)
	r.GET("/auth", controller.Authenticate)
	r.GET("/error", controller.Error)

	fmt.Println("listening on :" + port)
	handleErr(r.Start(":"+port), "failed to listen & serve")
}

func handleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
