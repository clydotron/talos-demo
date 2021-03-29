package views

import (
	"fmt"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Processes struct {
	app.Compo

	eb  *utils.EventBus
	sub *utils.EventBusSubscriber
	pui map[string]ui.CpuTracker
}

// we should be looking for "process" events (add/remove)
// but i am using this for now:

func NewProcessesView(eb *utils.EventBus) *Processes {

	fmt.Println("NewProcessesView --")
	e := &Processes{
		eb:  eb,
		pui: make(map[string]ui.CpuTracker),
		sub: utils.NewEventBusSubscriber("PI", eb),
	}
	return e
}

func (c *Processes) handleEvent(d utils.DataEvent) {

	app.Dispatch(func() {
		if pi, ok := d.Data.(*models.ProcessInfo); ok {

			if _, exists := c.pui[pi.ID]; !exists {
				fmt.Println("Processes.handleEvent: add:", pi)
				c.pui[pi.ID] = *ui.NewCpuTracker(pi.ID, c.eb)
				//update the fill color?
				c.Update()
			}
		}
	})
}

func (c *Processes) OnMount(ctx app.Context) {
	//fmt.Println("Processes onMount >start< processes:", len(c.pui))
	defer fmt.Println("Processes onMount >end<")

	c.sub.Start(c.handleEvent)
}

// OnDismount ...
func (c *Processes) OnDismount() {
	defer fmt.Println("Processes dismounted!")
	c.sub.Stop()
}

// Render ...
func (c *Processes) Render() app.UI {

	// cycle thru the
	return app.Div().Class("h-screen w-screen").
		Body(
			&ui.NavBar{},
			app.Div().Class("pt-16").Body(

				app.Range(c.pui).Map(func(k string) app.UI {
					v := c.pui[k]
					return app.Div().Body(&v)
				}),
			),
		)

}
