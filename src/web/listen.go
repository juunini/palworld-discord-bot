package web

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed public/*
var public embed.FS

var app *fiber.App

func Listen(port int) error {
	app = fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     true,
	}))

	return app.Listen(fmt.Sprintf(":%d", port))
}
