package views

import (
	"fmt"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Machines ...
type machines struct {
	app.Compo
	eb  *utils.EventBus
	sub *utils.EventBusSubscriber
	MI  []models.MachineInfo
}

func NewMachinesView(eb *utils.EventBus) *machines {
	mv := &machines{
		eb:  eb,
		sub: utils.NewEventBusSubscriber("MI", eb), //events.MachineInfo
	}
	return mv
}

func (mv *machines) UpdateM(mi *[]models.MachineInfo) {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		mv.MI = *mi
		mv.Update()
	})
}

func (mv *machines) handleEvent(d utils.DataEvent) {
	//switch on the event type?
	app.Dispatch(func() {
		if mi, ok := d.Data.(*models.MachineInfo); ok {
			fmt.Println(mi)
			//either update an existing MI, or add one
			mv.Update()
		}
	})
}

func (c *machines) OnMount(ctx app.Context) {
	fmt.Println("component mounted")

	r := &models.MachineInfoRequest{
		ID: "allMachines",
		Callback: func(mx []models.MachineInfo) {
			app.Dispatch(func() {
				c.MI = mx
				c.Update()
			})
		},
	}
	//request machine info
	c.eb.Publish("mi_req", r)

	c.sub.Start(c.handleEvent)
}
func (c *machines) OnDismount() {
	c.sub.Stop()
}

// Render ...
func (c *machines) Render() app.UI {

	return app.Div().Class("h-screen w-screen").
		Body(
			&ui.NavBar{},
			app.Div().Class("h-screen w-screen bg-gray-100 pt-20 p-8 flex flex-col").
				Body(
					app.Range(c.MI).Slice(func(i int) app.UI {
						return app.Ul().Class("w-48 h-full bg-gray-300").
							Body(
								&ui.Machine{MI: c.MI[i]},
							)
					}),
				),
		)

}
