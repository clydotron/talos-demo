package views_test

import (
	"fmt"
	"testing"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/clydotron/talos-demo/client/views"
)

func TestHandleEvent(t *testing.T) {

	eb := utils.NewEventBus()
	v := views.NewProcessesView(eb)

	//send a PI event, confirm that
	pi := models.ProcessInfo{ID: "p1", Name: "p1"}
	eb.Publish("PI", &pi)

	fmt.Println(v)

	//@todo add actual test
}
