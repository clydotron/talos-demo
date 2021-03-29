package views

import (
	"fmt"
	"net/url"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Machines ...
type Machines struct {
	app.Compo

	MI  []models.MachineInfo
	upx ui.Updater
}

func (mv *Machines) UpdateM(mi *[]models.MachineInfo) {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		mv.MI = *mi
		mv.Update()
	})
}

// OnNav ...
func (c *Machines) OnNav(ctx app.Context, url *url.URL) {
	app.Log("onNav - machines")
	fmt.Println(url)

	// 	//fmt.Println(url)
}

func (c *Machines) OnMount(ctx app.Context) {
	fmt.Println("component mounted")
}

// Render ...
func (c *Machines) Render() app.UI {

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
