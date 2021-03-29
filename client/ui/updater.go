package ui

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Updater ...
type Updater struct {
	app.Compo
	UpdateAvailable bool
}

// OnAppUpdate satisfies the app.Updater interface. It is called when the app is
// updated in background.
func (c *Updater) OnAppUpdate(ctx app.Context) {
	c.UpdateAvailable = ctx.AppUpdateAvailable // Reports that an app update is available.
	c.Update()                                 // Triggers UI update.
}

func (c *Updater) onUpdateClick(ctx app.Context, e app.Event) {
	// Reloads the page to display the modifications.
	app.Reload()
}

// Render ...
func (c *Updater) Render() app.UI {

	h := " hidden"
	if c.UpdateAvailable {
		h = ""
	}
	// make conditional in class?
	return app.Div().
		Class("bg-blue-100" + h).
		Body(
			//Displays an Update button when an update is available.
			app.If(c.UpdateAvailable,
				app.Button().
					Class("bg-indigo-300 rounded p-3").
					Text("Update!").
					OnClick(c.onUpdateClick),
			),
		)

}
