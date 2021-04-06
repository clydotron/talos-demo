// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"fmt"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/clydotron/talos-demo/client/views"
	"github.com/clydotron/talos-demo/client/ws"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {

	eventBus := utils.NewEventBus()
	eventSource := models.NewEventSource(eventBus)
	eventSource.Start()
	defer eventSource.Stop()

	piSource := models.NewProcessInfoSource(eventBus)
	piSource.AddProcess(models.ProcessInfo{ID: "P1", Name: "P1", CPU: 10})
	piSource.AddProcess(models.ProcessInfo{ID: "P2", Name: "P2", CPU: 10})
	piSource.AddProcess(models.ProcessInfo{ID: "P3", Name: "P3", CPU: 30})
	piSource.Start()
	defer piSource.Stop()

	taskRepo := models.NewTaskInfoRepo()

	miSource := models.NewMachineInfoSource(eventBus)
	miSource.InitWithFakeData()
	miSource.Start()
	defer miSource.Stop()

	tcx := ws.TestClientX{}
	tcx.Open("ws://localhost:8000/echo")
	defer tcx.Close()

	app.Route("/", &ui.UpdateButton{})
	app.Route("/clusters", views.NewClustersView(eventBus))
	app.Route("/machines", views.NewMachinesView(eventBus))
	app.Route("/events", views.NewEventsView(eventBus))
	app.Route("/charts", views.NewProcessesView(eventBus))
	app.RouteWithRegexp("^/task.*", ui.NewTaskDetail(taskRepo))

	app.Run() // Launches the PWA.

	fmt.Println("DONE")
}
