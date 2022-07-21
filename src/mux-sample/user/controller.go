package domain

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type UserController struct {
	service Service
	tracer  trace.Tracer
}

func NewUserController(service Service) *UserController {
	return &UserController{
		service: service,
		tracer:  otel.Tracer("UserController"),
	}
}

func (controller *UserController) Get(w http.ResponseWriter, req *http.Request) {
	ctx, span := controller.tracer.Start(req.Context(), getCurrentFuncName())
	defer span.End(trace.WithStackTrace(true))

	userId := mux.Vars(req)["id"]
	user, _ := controller.service.Get(ctx, userId)
	setResponse(w, user, http.StatusOK)
}

func (controller *UserController) Add(w http.ResponseWriter, req *http.Request) {
	ctx, span := controller.tracer.Start(req.Context(), getCurrentFuncName())
	defer span.End(trace.WithStackTrace(true))
	var user User
	var m sync.Mutex

	m.Lock()
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		span.SetStatus(otelcodes.Error, "request decode error")
		span.RecordError(err)
		setResponse(w, err, http.StatusBadRequest)
		return
	}

	error := controller.service.Add(ctx, user)
	if error != nil {
		span.RecordError(error)
		setResponse(w, error.Error(), http.StatusBadRequest)
		return
	}
	time.Sleep(8 * time.Second)

	m.Unlock()
	setResponse(w, user, http.StatusCreated)
}

func (controller *UserController) All(w http.ResponseWriter, req *http.Request) {
	_, span := controller.tracer.Start(req.Context(), "controller::All")
	defer span.End(trace.WithStackTrace(true))
	users, _ := controller.service.List(req.Context())
	setResponse(w, users, http.StatusOK)
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
