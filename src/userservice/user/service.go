package domain

import (
	"context"
	"errors"
	"runtime"
	"time"

	"github.com/digma-ai/otel-sample-application-go/src/authenticator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	ExtraLatency time.Duration
)

type Service interface {
	List() ([]User, error)
	Add(ctx context.Context, user User) error
	Get(ctx context.Context, id string) (User, error)
	Init()
}

func NewUserService() Service {
	return &userService{}
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type userService struct {
	users map[string]User
}

var ErrIdInvalid = errors.New("user id too long")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

func (u *userService) Init() {
	u.users = make(map[string]User)
}

func (u *userService) Get(ctx context.Context, id string) (User, error) {
	tracer := otel.GetTracerProvider().Tracer("UserService")
	ctx, span :=
		tracer.Start(ctx, funcName(0))
	defer span.End(trace.WithStackTrace(true))

	// defer func() { //rethrow
	// 	if err := recover(); err != nil {
	// 		e, _ := err.(error)
	// 		panic(e.Error() + "my new panic")
	// 	}
	// }()

	time.Sleep(ExtraLatency)
	value, found := u.users[id]
	authenticator.Authenticate(ctx, found)

	return value, nil
}
func funcName(depth int) string {
	pc, _, _, ok := runtime.Caller(depth + 1)
	if ok {
		fn := runtime.FuncForPC(pc)
		return fn.Name()
	}
	return ""
}

// func (u *userService) Get(ctx context.Context, id string) (User, error) {
// 	time.Sleep(ExtraLatency)
// 	value, found := u.users[id]
// 	authenticator.Authenticate(ctx, found)
// 	return value, nil
// }

func (u *userService) List() ([]User, error) {
	time.Sleep(ExtraLatency)

	v := []User{}
	for _, value := range u.users {
		v = append(v, value)
	}
	return v, nil
}

func (u *userService) Add(ctx context.Context, user User) error {
	tracer := otel.GetTracerProvider().Tracer("UserService")
	ctx, span :=
		tracer.Start(ctx, funcName(0))
	defer span.End(trace.WithStackTrace(true))
	time.Sleep(2 * time.Second)
	if len(user.Id) > 5 {
		panic("invalid user id: " + user.Id)
	}
	if _, ok := u.users[user.Id]; ok {
		return ErrUserAlreadyExists
	}

	u.users[user.Id] = user
	return nil
}
