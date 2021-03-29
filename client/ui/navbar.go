package ui

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// NavBar ...
type NavBar struct {
	app.Compo
	upx UpdateButton
}

// Render ...
func (c *NavBar) Render() app.UI {
	ss := "py-5 px-3 text-gray-700 hover:text-gray-600"

	return app.Nav().Class("bg-indigo-200 fixed w-screen").
		Body(
			app.Div().Class("px-6 mx-auto flex justify-between").
				Body(
					app.Div().Class("flex space-x-2").
						Body(
							app.A().Class(ss).Href("/clusters").Text("Clusters"),
							app.A().Class(ss).Href("/machines").Text("Machines"),
							app.A().Class(ss).Href("/events").Text("Events"),
							app.A().Class(ss).Href("/charts").Text("Charts"),
						),
					app.Div().Class("flex items-center space-x-1").
						Body(
							&c.upx,
						),
				),
		)
}
