package main

import "github.com/omnipunk/pwa/v9/app"

const (
	installApp = "/app/install"
	updateApp  = "/app/update"
)

func handleAppInstall(ctx app.Context, a app.Action) {
	ctx.ShowAppInstallPrompt()
}

func handleAppUpdate(ctx app.Context, a app.Action) {
	ctx.Reload()
}
