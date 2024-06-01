package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	cfgPath := config.GetConfigPath("local")
	cfgFile, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalln(err)
	}

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

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatalln(err)
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   43200, // 12 hours
		HttpOnly: true,
		// Secure:   true,
	})

	e.Static("/dist", "./dist")
	controller.NewUserController(e, userUsecase, store)
	controller.NewTodoController(e, todoUsecase, store)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/user/login")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
