package models_test

import (
	"testing"
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/utils"
)

func TestRequestEvents(t *testing.T) {

	eb := utils.NewEventBus()
	es := models.NewEventSource(eb)

	st := time.Now()
	ts := [5]time.Time{}

	for i := 0; i < 5; i++ {
		ts[i] = st.Add(time.Duration(i) * 20 * time.Millisecond)
		es.Events = append(es.Events, models.EventInfo{TimeStamp: ts[i]})
	}

	//request all
	events := es.GetEventsAfter(ts[0])
	if len(events) != 4 {
		t.Fatal("didnt get enough events")
	}
	// request none
	events = es.GetEventsAfter(ts[4])
	if len(events) > 0 {
		t.Fatal("was expecting zero events")
	}

}
