package views

import (
	"fmt"
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

/*
todos:
build mechanism to get all events up to ___ (some point) on mount
if the user hits clear, remember that point for future mounts
*/
// Events ...
type Events struct {
	app.Compo
	events      []models.EventInfo
	eb          *utils.EventBus
	sub         *utils.EventBusSubscriber
	lastEventTS time.Time
}

// NewEventsView factory function
func NewEventsView(eb *utils.EventBus) *Events {
	e := &Events{
		eb:          eb,
		sub:         utils.NewEventBusSubscriber("event", eb),
		lastEventTS: time.Now(),
	}
	e.sub.Name = "Events"
	return e
}

func (c *Events) handleEvent(d utils.DataEvent) {

	app.Dispatch(func() {
		if ei, ok := d.Data.(*models.EventInfo); ok {
			c.events = append(c.events, *ei)
			c.lastEventTS = ei.TimeStamp
			c.Update()
		}
	})
}

func (c *Events) OnMount(ctx app.Context) {
	fmt.Println("Events onMount >start<")
	defer fmt.Println("Events onMount >end<")

	// request all events since the last one we received
	// @todo - for the first one do we want to go back in time? (or just start at the present)
	ri := &models.EventInfoRequest{
		StartTime: c.lastEventTS,
		Callback: func(events []models.EventInfo) {
			app.Dispatch(func() {
				c.events = append(c.events, events...)
				c.Update()
			})
		},
	}
	c.eb.Publish("event_req", ri)

	c.sub.Start(c.handleEvent)
}

func (c *Events) OnDismount() {
	c.sub.Stop()
}

func (c *Events) clearEvents(ctx app.Context, e app.Event) {
	app.Dispatch(func() {
		c.events = []models.EventInfo{}
		c.Update()
	})
}

// Render ...
func (c *Events) Render() app.UI {

	// @otodo make something in the event info clickable to bring up detailed view
	// make event info component?
	return app.Div().Class("h-screen w-screen").
		Body(
			&ui.NavBar{},
			app.Div().Class("pt-20 flex flex-col px-2").
				Body(
					app.Table().
						Body(
							app.Button().Class("rounded bg-indigo-200 p-1 mb-2").Text("Clear").OnClick(c.clearEvents),
							app.Tr().Class("bg-gray-200").Body(
								//bg color, change alignment
								app.Th().Class("text-left").Text("Date"),
								app.Th().Class("text-left").Text("Event"),
								app.Th().Class("text-left").Text("ID"),
							),
							app.Range(c.events).Slice(func(i int) app.UI {
								e := c.events[i]
								bgcolor := "bg-blue-200"
								if i%2 == 1 {
									bgcolor = "bg-blue-100"
								}
								return app.Tr().Class(bgcolor).
									Body(
										app.Td().Text(e.TimeStamp.Format("02 Jan 2006 at 15:04:05")),
										app.Td().Text(e.Name),
										app.Td().Text(e.ID),
									)

							}),
						),
				),
		)
}
