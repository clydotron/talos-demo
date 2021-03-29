package ui

import (
	"github.com/clydotron/talos-demo/client/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// ControlPlane ...
type ControlPlane struct {
	app.Compo

	CPI models.ControlPlaneInfo
}

// Render ...
func (c *ControlPlane) Render() app.UI {
	return app.Div().Class("bg-white shadow p-3 mt-2 rounded").
		Body(
			app.Div().Class("text-left").Body(
				app.H3().Class("mb-2 text-gray-700").Text(c.CPI.Name),
				app.P().Class("text-grey-600 text-sm").Text(c.CPI.Status),
			),
			app.Div().Class("mt-4").Body(
				app.A().Class("no-underline mr-4 text-blue-500 hover:text-blue-400").Href("#").Text("details"),
			),
		)
}
