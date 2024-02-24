package web

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/juunini/palworld-discord-bot/src/config"
)

//go:embed public/*
var public embed.FS

var app *fiber.App

func Listen(port int) error {
	app = fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(public), ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("public/index", fiber.Map{
			"LANGUAGE":                    config.LANGUAGE,
			"DISCORD_BOT_ENABLED":         config.DISCORD_BOT_ENABLED,
			"DISCORD_BOT_TOKEN":           config.DISCORD_BOT_TOKEN,
			"DISCORD_ADMIN_USERNAMES":     strings.Join(config.DISCORD_ADMIN_USERNAMES, ", "),
			"DISCORD_COMMAND_IGNORE_CASE": config.DISCORD_COMMAND_IGNORE_CASE,
			"DISCORD_COMMAND_PREFIX":      config.DISCORD_COMMAND_PREFIX,
		})
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     true,
	}))

	return app.Listen(fmt.Sprintf(":%d", port))
}
