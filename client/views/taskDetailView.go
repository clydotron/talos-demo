package views

import (
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Machines ...
type TaskDetailView struct {
	app.Compo
}

func (c *TaskDetailView) Render() app.UI {

	return app.Div().Class("h-screen w-screen").
		Body(
			&ui.NavBar{},
		)
}
