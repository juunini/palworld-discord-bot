package web

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
	"github.com/juunini/palworld-discord-bot/src/web/handlers"
)

//go:embed public/*
var public embed.FS

var app *fiber.App

func Listen(port int) error {
	app = fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(public), ".html"),
	})

	app.Patch("/config", handlers.PatchConfig(Shutdown))
	app.Get("/", handlers.Index)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     true,
	}))

	console_decoration.PrintError(i18n.WebServerOpeningMessage + "\n")

	privateIPs, err := utils.PrivateIPs()
	if err != nil {
		return err
	}

	for _, ip := range privateIPs {
		console_decoration.PrintSuccess(fmt.Sprintf("- http://%s:%d", ip, port))
	}

	return app.Listen(fmt.Sprintf(":%d", port))
}
