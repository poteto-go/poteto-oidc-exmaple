package main

import (
	"github.com/poteto-go/poteto-hono-oidc/app"
	"github.com/poteto-go/poteto-hono-oidc/app/src/config"
)

func main() {
	config.InitAppConfig()

	app := app.NewApp()

	app.Run(config.AppConfig.AppPort)
}
