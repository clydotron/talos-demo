package ui

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Updater ...
type UpdateButton struct {
	app.Compo
	updateAvailable bool
}

// OnAppUpdate satisfies the app.Updater interface. It is called when the app is
// updated in background.
func (c *UpdateButton) OnAppUpdate(ctx app.Context) {
	c.updateAvailable = ctx.AppUpdateAvailable // Reports that an app update is available.
	c.Update()                                 // Triggers UI update.
}

func (c *UpdateButton) onUpdateClick(ctx app.Context, e app.Event) {
	// Reloads the page to display the modifications.
	app.Reload()
}

// Render ...
func (c *UpdateButton) Render() app.UI {

	h := " hidden"
	if c.updateAvailable {
		h = ""
	}
	// make conditional in class?
	return app.Div().Class(h).
		Body(
			//Displays an Update button when an update is available.
			app.If(c.updateAvailable,
				app.Button().
					Class("bg-indigo-300 rounded p-2 m-2").
					Text("Update!").
					OnClick(c.onUpdateClick),
			),
		)

}
