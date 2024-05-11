package main

import (
	"log"

	"github.com/Anttoam/golang-htmx-todos/config"
	"github.com/Anttoam/golang-htmx-todos/internal/controller"
	"github.com/Anttoam/golang-htmx-todos/internal/repository"
	"github.com/Anttoam/golang-htmx-todos/internal/usecase"
	"github.com/Anttoam/golang-htmx-todos/pkg/turso"

	"github.com/gofiber/fiber/v2"
)

type Film struct {
	Title    string
	Director string
}

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

	app := fiber.New()
	controller.NewUserController(app, userUsecase)

	log.Fatal(app.Listen(":8080"))
}
