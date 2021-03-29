// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/clydotron/talos-demo/client/views"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type DataStore struct {
	MI []models.MachineInfo
	MV *views.Machines
}

func (d *DataStore) update(mi []models.MachineInfo) {
	d.MI = mi
	d.MV.UpdateM(&mi)
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	//app.Log("setting up routes")

	// client := client.NewClusterClient()
	// err := client.Connect("0.0.0.0:50051")
	// if err != nil {
	// 	log.Fatalln("### Client failed to connect:", err)
	// }
	// defer client.Close()

	eventBus := utils.NewEventBus()
	eventSource := models.NewEventSource(eventBus)
	eventSource.Start()

	piSource := models.NewProcessInfoSource(eventBus)
	piSource.AddProcess(models.ProcessInfo{ID: "P1", Name: "P1", CPU: 10})
	piSource.AddProcess(models.ProcessInfo{ID: "P2", Name: "P2", CPU: 10})
	piSource.AddProcess(models.ProcessInfo{ID: "P3", Name: "P3", CPU: 30})
	piSource.Start()

	taskRepo := models.NewTaskInfoRepo()

	mi := []models.MachineInfo{
		models.MachineInfo{
			Name:   "Machine 1",
			Role:   "Manager",
			Status: "running",
			Tasks: []models.TaskInfo{
				models.TaskInfo{
					Name:        "Redis",
					Tag:         "3.2.1",
					ContainerID: "badc0ffee",
					State:       "running",
					Updated:     time.Now(),
				},
			},
		},
	}

	DS := &DataStore{
		MV: &views.Machines{MI: mi},
		MI: mi,
	}

	//setup the timer:

	app.Route("/", &ui.Updater{}) // hello component is associated with URL path "/".
	app.Route("/clusters", views.NewClustersView(eventBus))
	app.Route("/machines", DS.MV)
	app.Route("/events", views.NewEventsView(eventBus))
	app.Route("/charts", views.NewProcessesView(eventBus))
	app.RouteWithRegexp("^/node.*", &ui.Node{})
	app.RouteWithRegexp("^/task.*", ui.NewTaskDetail(taskRepo))

	// app.Route("/machine", &views.Machines{MI: []models.MachineInfo{
	// 	models.MachineInfo{
	// 		Name: "Machine 1",
	// 		Role: "Manager",
	// 		Tasks: []models.TaskInfo{
	// 			models.TaskInfo{
	// 				Name:        "Redis",
	// 				Tag:         "3.2.1",
	// 				ContainerID: "bunch of hex",
	// 				State:       "running",
	// 				Updated:     time.Now(),
	// 			},
	// 			models.TaskInfo{
	// 				Name:        "Mongo",
	// 				Tag:         "4.2.1",
	// 				ContainerID: "bunch of hex",
	// 				State:       "stopped",
	// 				Updated:     time.Now(),
	// 			},
	// 		},
	// 	},
	// }})
	app.Run() // Launches the PWA.
}

// return app.Div().OnClick(c.onClick)
// }
