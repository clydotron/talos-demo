package models_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/utils"
)

func TestPiSouceAddProcess(t *testing.T) {

	eb := utils.NewEventBus()
	piSource := models.NewProcessInfoSource(eb)

	pi := models.ProcessInfo{ID: "p1", Name: "p1"}
	piSource.AddProcess(pi)

	// and?
}

func TestPiSouceSendUpdate(t *testing.T) {

	eb := utils.NewEventBus()
	piSource := models.NewProcessInfoSource(eb)
	piSource.AddProcess(models.ProcessInfo{ID: "1"})

	dataCh := make(chan utils.DataEvent)
	eb.Subscribe("PI", dataCh)

	timeout := time.After(200 * time.Millisecond)

	piSource.SendUpdate()

	select {
	case <-timeout:
		t.Fatal("did not finish in time")
	case d := <-dataCh:
		if _, ok := d.Data.(*models.ProcessInfo); !ok {
			t.Fatal("failed to convert")
		} else {
			t.Log("success")
			fmt.Println(d)
		}
	}

}

// add test for requesting historical data
