package models

import (
	"time"

	"github.com/clydotron/talos-demo/client/utils"
)

type MachineInfoSource struct {
	eb  *utils.EventBus
	sub *utils.EventBusSubscriber
	mi  map[string]MachineInfo
}

func NewMachineInfoSource(eb *utils.EventBus) *MachineInfoSource {

	ms := &MachineInfoSource{
		eb:  eb,
		sub: utils.NewEventBusSubscriber("mi_req", eb), //events.MachineInfoRequest (MIR)
		mi:  make(map[string]MachineInfo),
	}
	return ms
}

func (ms *MachineInfoSource) InitWithFakeData() {

	mi := []MachineInfo{
		{
			ID:     "Mx1",
			Name:   "Machine 1",
			Role:   "Manager",
			Status: "running",
			Tasks: []TaskInfo{
				{
					Name:        "Redis",
					Tag:         "3.2.1",
					ContainerID: "badc0ffee",
					State:       "running",
					Updated:     time.Now(),
				},
			},
		},
	}
	for _, m := range mi {
		ms.mi[m.ID] = m
	}
}

//
func (ms *MachineInfoSource) handleEvent(d utils.DataEvent) {
	if d.Topic == "mi_req" {
		if req, ok := d.Data.(*MachineInfoRequest); ok {
			if req.ID == "allMachines" {
				mx := make([]MachineInfo, 0, len(ms.mi))
				for _, m := range ms.mi {
					mx = append(mx, m)
				}
				req.Callback(mx)
			} else {

				//@todo
			}
		}
	}
}

func (mis *MachineInfoSource) Start() {
	mis.sub.Start(mis.handleEvent)
}

func (mis *MachineInfoSource) Stop() {

	mis.sub.Stop()
}
