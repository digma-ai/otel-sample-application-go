package domain

import (
	"context"
	"errors"
	"log"
	"runtime"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	ExtraLatency time.Duration
)

type Service interface {
	List(ctx context.Context) ([]User, error)
	Add(ctx context.Context, user User) error
	Get(ctx context.Context, id string) (User, error)
}

func NewUserService(db *pgxpool.Pool) Service {
	return &userService{db: db}
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type userService struct {
	db *pgxpool.Pool
}

var ErrIdInvalid = errors.New("user id too long")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

func (u *userService) Get(ctx context.Context, id string) (User, error) {
	tracer := otel.GetTracerProvider().Tracer("UserService")
	_, span :=
		tracer.Start(ctx, funcName(0))
	defer span.End(trace.WithStackTrace(true))

	// defer func() { //rethrow
	// 	if err := recover(); err != nil {
	// 		e, _ := err.(error)
	// 		panic(e.Error() + "my new panic")
	// 	}
	// }()

	time.Sleep(ExtraLatency)
	rows, err := u.db.Query(ctx, "select * from users where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(users) != 0 {
		return users[0], nil
	} else {
		return User{}, errors.New("user not found")
	}
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
// check}

func (u *userService) List(ctx context.Context) ([]User, error) {
	tracer := otel.GetTracerProvider().Tracer("UserService")
	_, span :=
		tracer.Start(ctx, funcName(0))
	defer span.End(trace.WithStackTrace(true))
	time.Sleep(ExtraLatency)

	rows, err := u.db.Query(ctx, "select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users, nil
}

func (u *userService) Add(ctx context.Context, user User) error {
	tracer := otel.GetTracerProvider().Tracer("UserService")
	_, span :=
		tracer.Start(ctx, funcName(0))
	defer span.End(trace.WithStackTrace(true))
	time.Sleep(2 * time.Second)
	if len(user.Id) > 5 {
		panic("invalid user id: " + user.Id)
	}

	if _, err := u.db.Exec(ctx, `insert into users(id, name) values ($1, $2)`, user.Id, user.Name); err != nil {
		return err
	}

	// if _, ok := u.users[user.Id]; ok {
	// 	return ErrUserAlreadyExists
	// }

	// u.users[user.Id] = user
	return nil
}
