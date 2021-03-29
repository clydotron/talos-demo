package ui

import (
	"github.com/clydotron/talos-demo/client/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Task ...
type Task struct {
	app.Compo

	TI models.TaskInfo
}

// UpdateData ...
func (t *Task) UpdateData(ti *models.TaskInfo) {

	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		t.TI = *ti
		t.Update()
	})
	// 	//app.Log("updating!")
}

// Render ...
// box with border:
// width = that of container - height (?)
func (c *Task) Render() app.UI {

	//app.Log("render")

	return app.Div().
		Class("bg-green-100 p-4 width-full flex flex-col items-center border border-green-700 hover:bg-green-200").
		Body(
			app.A().Href(`/task/1`).
				Body(
					app.H4().Text(c.TI.Name),
				),
			app.H4().Text(c.TI.Tag),
			app.H4().Text(c.TI.State),
			app.P().Text(c.TI.ContainerID),
			app.P().Text(c.TI.Updated.UTC().Format("02 Jan 2006 at 15:04")),
		)
}
