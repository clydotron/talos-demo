package ui

import (
	"fmt"
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Machine ...
type Machine struct {
	app.Compo

	MI      models.MachineInfo
	ticker  *time.Ticker
	doneCh  chan bool
	visited bool
}

func (c *Machine) updateStatus() {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.

		if c.MI.Status == "running" {
			c.MI.Status = "stopped"
		} else {
			c.MI.Status = "running"
		}
		c.Update()
	})
}

func (c *Machine) addTask() {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.

		t := &models.TaskInfo{
			Name:  "postgres",
			State: "error",
		}
		c.MI.Tasks = append(c.MI.Tasks, *t)
		c.Update()
	})
}

func (c *Machine) removeTask() {
	app.Dispatch(func() { // Ensures response field is updated on UI goroutine.
		//fmt.Println("remove task")

		if len(c.MI.Tasks) > 0 {
			c.MI.Tasks = c.MI.Tasks[:len(c.MI.Tasks)-1]
		}
		c.Update()
	})
}

func (c *Machine) OnMount(ctx app.Context) {
	fmt.Println("machine mounted")

	if !c.visited {
		// some experiments:
		timer2 := time.NewTimer(2 * time.Second)
		go func() {
			<-timer2.C
			c.updateStatus()
		}()

		timer3 := time.NewTimer(4 * time.Second)
		go func() {
			<-timer3.C
			c.addTask()
		}()

		timer4 := time.NewTimer(8 * time.Second)
		go func() {
			<-timer4.C
			c.removeTask()
		}()

		// start a ticker:
		c.ticker = time.NewTicker(2000 * time.Millisecond)
		c.doneCh = make(chan bool)
		go func() {
			for {
				select {
				case <-c.doneCh:
					return
				case <-c.ticker.C:
					c.updateStatus()
				}
			}
		}()
	}
	c.visited = true
}

func (c *Machine) OnDismount() {
	defer fmt.Println("machine dismounted")

	c.ticker.Stop()
	c.doneCh <- true

}

// Render ...
func (c *Machine) Render() app.UI {
	statusColor := "bg-red-500"
	if c.MI.Status == "running" {
		statusColor = "bg-green-500"
	}

	return app.Div().
		Class("h-full w-full border border-gray-300 relative").
		Body(
			app.Div().Class("flex flex-col items-center space-y-1 h-32 p-4").
				Body(
					app.Div().Class("flex items-center space-x-2").
						Body(
							app.Div().Class("h-3 w-3 rounded-full "+statusColor),
							app.H2().Text(c.MI.Name),
						),
					app.H1().Text(c.MI.Role),
				),
			app.Div().Class("w-full space-y-2 absolute bottom-0").
				Body(
					app.Range(c.MI.Tasks).Slice(func(i int) app.UI {
						return &Task{TI: c.MI.Tasks[i]}
					},
					),
				),
		)
}

// <div class="flex flex-col items-center space-y-1 border h-32 p-4">
// <div class="flex items-center space-x-2">
//   <div class="bg-red-400 h-3 w-3 rounded-full"></div>
//   <h2>Machine 1</h2>
// </div>
// <h3>Manager</h3>
// <h3>15 Gigs free</h3>
// </div>

// <ul class="w-48 h-full bg-blue-700 relative min-x-42">
// <div class="w-full space-y-2 border border-yellow-300 absolute bottom-0">
//   <div class="h-24 bg-purple-400 ">item 1</div>
//   <div class="h-24 bg-purple-400 pt-4">item 2</div>
// </div>
