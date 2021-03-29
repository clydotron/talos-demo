package models

import (
	"fmt"
	"time"

	"github.com/clydotron/talos-demo/client/utils"
)

type EventSource struct {
	Events []EventInfo

	eb  *utils.EventBus
	sub *utils.EventBusSubscriber

	ticker  *time.Ticker
	doneCh  chan bool
	eventId int
}

// NewEventSource
func NewEventSource(eb *utils.EventBus) *EventSource {
	es := &EventSource{
		eb:  eb,
		sub: utils.NewEventBusSubscriber("event_req", eb),
	}
	//hook some additional things up?
	return es
}

// handleEvent ...
func (es *EventSource) handleEvent(d utils.DataEvent) {

	if d.Topic == "event_req" {
		//figure out what events to send
		if req, ok := d.Data.(*EventInfoRequest); ok {
			req.Callback(es.GetEventsAfter(req.StartTime))
		} else {
			fmt.Println("bad :(")
		}
	}
}

// GetEventsAfter get all events after the specified start time
func (es *EventSource) GetEventsAfter(st time.Time) []EventInfo {
	for i, ei := range es.Events {
		if ei.TimeStamp.After(st) {
			return es.Events[i:]
		}
	}
	return []EventInfo{}
}

// Start ...
func (es *EventSource) Start() {

	es.sub.Start(es.handleEvent)
	//start sending events

	// start a ticker:
	es.ticker = time.NewTicker(2000 * time.Millisecond)
	es.doneCh = make(chan bool)

	go func() {
		for {
			select {
			case <-es.doneCh:
				return
			case <-es.ticker.C:
				es.sendEvent()
			}
		}
	}()
}

// Stop ...
func (es *EventSource) Stop() {
	es.sub.Stop()
	es.doneCh <- true
}

func (es *EventSource) sendEvent() {

	data := &EventInfo{
		Name:      fmt.Sprintln("Event", es.eventId),
		ID:        fmt.Sprintln(es.eventId),
		TimeStamp: time.Now(),
	}
	es.eb.Publish("event", data)
	es.Events = append(es.Events, *data)
	es.eventId++
}
