package utils_test

import (
	"testing"
	"time"

	"github.com/clydotron/talos-demo/client/utils"
)

func TestEventBusSubscriber(t *testing.T) {

	topic := "test"
	badTopic := "bad"

	eb := utils.NewEventBus()
	ebc := utils.NewEventBusSubscriber(topic, eb)

	dataCh := make(chan utils.DataEvent)
	fcn := func(d utils.DataEvent) {
		dataCh <- d
	}

	ebc.Start(fcn)
	defer ebc.Stop()

	eb.Publish(topic, nil)

	timeout := time.After(100 * time.Millisecond)

	select {
	case <-timeout:
		t.Fatal("did not finish in time")
	case d := <-dataCh:
		if d.Topic != topic {
			t.Fatal("wrong topic!")
		}
	}

	eb.Publish(badTopic, nil)
	eb.Publish(badTopic, nil)
	//eb.Publish(topic, nil)

	timeout = time.After(100 * time.Millisecond)

	select {
	case <-timeout:
		// expected (can we shorten the timeout?)
		//t.Fatal("did not finish in time")
	case d := <-dataCh:
		if d.Topic != topic {
			t.Fatal("wrong topic!")
		}
	}
}
