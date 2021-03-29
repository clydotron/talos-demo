package utils_test

import (
	"testing"
	"time"

	"github.com/clydotron/talos-demo/client/utils"
)

func TestEventBusSubscribe(t *testing.T) {

	//doneCh := make(chan bool)
	timeout := time.After(1 * time.Second)

	eb := utils.NewEventBus()

	testCh := make(chan utils.DataEvent)
	eb.Subscribe("test", testCh)
	eb.Publish("test", nil)

	select {
	case <-timeout:
		t.Fatal("did not finish in time")
	case <-testCh:
	}
}
func TestEventBusWrongTopic(t *testing.T) {

	timeout := time.After(200 * time.Millisecond)
	eb := utils.NewEventBus()

	testCh := make(chan utils.DataEvent)
	eb.Subscribe("test", testCh)
	eb.Publish("not_test", nil)

	select {
	case <-testCh:
		t.Fatal("Unexpected message on stream")
	case <-timeout:

	}
}

func TestEventBusUnsubscribe(t *testing.T) {
	timeout := time.After(200 * time.Millisecond)
	eb := utils.NewEventBus()

	testCh := make(chan utils.DataEvent)
	eb.Subscribe("test", testCh)
	eb.Unsubscribe("test", testCh)
	eb.Publish("test", nil)

	select {
	case <-testCh:
		t.Fatal("Unexpected message on stream")
	case <-timeout:
	}
}
