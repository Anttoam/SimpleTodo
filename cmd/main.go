package main

import (
	"log"
	"time"

	"github.com/Anttoam/golang-htmx-todos/config"
	"github.com/Anttoam/golang-htmx-todos/internal/controller"
	"github.com/Anttoam/golang-htmx-todos/internal/repository"
	"github.com/Anttoam/golang-htmx-todos/internal/usecase"
	"github.com/Anttoam/golang-htmx-todos/pkg/storage"
	"github.com/Anttoam/golang-htmx-todos/pkg/turso"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/Anttoam/golang-htmx-todos/docs"
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

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	redis := storage.NewRedisClient(cfg)
	store := session.New(session.Config{
		Storage:      redis,
		Expiration:   12 * time.Hour,
		KeyLookup:    "cookie:session_id",
		CookiePath:   "/",
		CookieDomain: "localhost",
		// CookieSecure: true,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
	})
	app.Static("/dist", "./dist")
	controller.NewUserController(app, userUsecase, store)

	controller.NewTodoController(app, todoUsecase, store)

	log.Fatal(app.Listen(":8080"))
}
