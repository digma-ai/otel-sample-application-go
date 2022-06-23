package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
)

type AuthController struct {
	service Service
	tracer  trace.Tracer
}

func NewAuthController(service Service, tracer trace.Tracer) AuthController {
	return AuthController{
		service: service,
		tracer:  tracer,
	}
}

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Auth!")
}

func (controller *AuthController) Authenticate(c echo.Context) error {
	return c.String(http.StatusOK, "Authenticated!")
}

func (controller AuthController) Error(c echo.Context) error {
	panic("ERROR!!!!")
}
