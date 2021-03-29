package ui

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// TaskDetail ...
type TaskDetail struct {
	app.Compo

	TI   models.TaskInfo
	repo *models.TaskInfoRepo
}

// NewTaskDetail create a new instance of TaskDetail
func NewTaskDetail(repo *models.TaskInfoRepo) *TaskDetail {
	var td TaskDetail
	td.repo = repo
	return &td
}

// Updated - Observer interface:
func (c *TaskDetail) Updated(i interface{}) {
	app.Dispatch(func() {
		ti, ok := i.(*models.TaskInfo)
		if ok {
			c.TI = *ti
			c.Update()
		}
	})
}

// GetID - Observer interface
func (c *TaskDetail) GetID() string {
	return c.TI.ContainerID
}

// OnNav ...
func (c *TaskDetail) OnNav(ctx app.Context, url *url.URL) {

	fmt.Println("url:", url, "coming from:", app.Window().Get("document").Get("referrer").String())

	history := app.Window().Get("document").Get("history")
	fmt.Println("history:", history)

	//@todo is there a better way to do this? could i encode ?if=x
	ep := url.EscapedPath()
	s := strings.Split(ep, "/")
	id := s[len(s)-1]

	ti, err := c.repo.GetTaskInfo(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// register as an observer for the TaskInfo
	ti.AddObserver(c)

	// update the TaskInfo (on the UI thread) - delay registering until here?
	// could potentially have stale data?
	app.Dispatch(func() {
		c.TI = *ti
		c.Update()
	})
}

// func (c *TaskDetail) OnMount(ctx app.Context) {
// 	fmt.Println("TaskDetail mounted")
// }

func (c *TaskDetail) OnDismount() {
	fmt.Println("TaskDetail - OnDismount")
	c.TI.RemoveObserver(c)
}

func (c *TaskDetail) EventHandlerX(ctx app.Context, e app.Event) {

	// collapse this into onwe
	doc := app.Window().Get("document")
	ref := doc.Get("referrer")
	x := app.Window().Get("document").Get("referrer").String()
	fmt.Println("Xx: ", x)
	app.Navigate(ref.String())
}

// Render ...
// box with border:
// width = that of container - height (?)
func (c *TaskDetail) Render() app.UI {

	return app.Div().
		Class("bg-gray-100").
		Body(
			&NavBar{},
			app.Div().Class("pt-20 px-2").Body(
				app.H1().Text("Name: "+c.TI.Name),
				app.H3().Text("Tag: "+c.TI.Tag),
				app.H3().Text("State: "+c.TI.State),
				app.H3().Text("ContainerID: "+c.TI.ContainerID),
				app.Br(),
				app.A().Href(app.Window().Get("document").Get("referrer").String()).OnClick(c.EventHandlerX).Text("Back"),
			),
		)
}

//
