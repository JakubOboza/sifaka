package server

import (
	"embed"
	"errors"
	"fmt"
	"net/http"

	"github.com/JakubOboza/sifaka/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

var (
	ErrPortOutOfBands = errors.New("server port should be between 1024 and 65535")
)

//go:embed views/*
var embedDirViews embed.FS

type App struct {
	port    int
	engine  *fiber.App
	storage storage.Service
}

func New(port int, storageService storage.Service) *App {
	return &App{port: port, storage: storageService}
}

func (app *App) Setup() {
	engine := html.NewFileSystem(http.FS(embedDirViews), ".html")

	app.engine = fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup
	app.engine.Use(logger.New())
	// Routes
	app.engine.Get("/", app.indexPage())
}

func (app *App) Start() error {
	if app.port < 1024 || app.port > 65535 {
		return ErrPortOutOfBands
	}
	return app.engine.Listen(fmt.Sprintf(":%d", app.port))
}
