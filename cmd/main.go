package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Anttoam/SimpleTodo/config"
	"github.com/Anttoam/SimpleTodo/internal/controller"
	"github.com/Anttoam/SimpleTodo/internal/repository"
	"github.com/Anttoam/SimpleTodo/internal/usecase"
	"github.com/Anttoam/SimpleTodo/pkg/turso"
	"github.com/Anttoam/SimpleTodo/pkg/validation"
	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/Anttoam/SimpleTodo/docs"
)

// @title			Simple Todo API
// @version		1.0
// @description	example todo api
func main() {
	cfg := config.NewConfig()
	db, err := turso.NewLibsqlDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	uri := cfg.Redis.Url
	opts, err := redis.ParseURL(uri)
	if err != nil {
		log.Fatalln(err)
	}

	Isbool, err := strconv.ParseBool(cfg.Redis.Secure)
	if err != nil {
		log.Fatalln(err)
	}

	if strings.HasPrefix(uri, "rediss") {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: Isbool,
		}
	}

	rdb := redis.NewClient(opts)

	store, err := redisstore.NewRedisStore(context.Background(), rdb)
	if err != nil {
		log.Fatalln(err)
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   cfg.Http.Domain,
		MaxAge:   43200, // 12 hours
		HttpOnly: true,
		Secure:   true,
	})

	e.Static("/dist", "./dist")
	controller.NewUserController(e, userUsecase, store)
	controller.NewTodoController(e, todoUsecase, store)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/user/login")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Http.Port)))
}
