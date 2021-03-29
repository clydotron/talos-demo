package models

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/clydotron/talos-demo/client/utils"
)

type ProcessInfoEx struct {
	id         string
	name       string
	cpu        float64 //being lazy
	cpuHistory []float64
}

type ProcessInfoSource struct {
	eb      *utils.EventBus
	sub     *utils.EventBusSubscriber
	ticker  *time.Ticker
	doneCh  chan bool
	eventId int

	m         sync.RWMutex
	processes map[string]*ProcessInfoEx
}

// NewProcessInfoSource ...
func NewProcessInfoSource(eb *utils.EventBus) *ProcessInfoSource {
	ps := &ProcessInfoSource{
		eb:        eb,
		processes: map[string]*ProcessInfoEx{},
		sub:       utils.NewEventBusSubscriber("pi_req", eb),
	}
	//hook some additional things up?
	return ps
}

// Start ...
func (ps *ProcessInfoSource) Start() {
	//start sending events every second

	rand.Seed(time.Now().UnixNano())
	ps.ticker = time.NewTicker(time.Second)
	ps.doneCh = make(chan bool)

	go func() {
		for {
			select {
			case <-ps.doneCh:
				return
			case <-ps.ticker.C:
				ps.SendUpdate()
			}
		}
	}()

	ps.sub.Start(ps.handleEvent)
}

//
func (ps *ProcessInfoSource) handleEvent(d utils.DataEvent) {
	if d.Topic == "pi_req" {
		if req, ok := d.Data.(*ProcessInfoCpuHistoryRequest); ok {
			req.Callback(ps.GetProcessHistory(req.ID, req.Num))
		} else {
			fmt.Println("ProcessInfoSource - bad :(")
		}
	}
}

func (ps *ProcessInfoSource) GetProcessHistory(id string, num int) []float64 {
	if p, ok := ps.processes[id]; ok {

		si := 0
		hs := len(p.cpuHistory)

		if hs > num {
			si = hs - num
		}
		return p.cpuHistory[si:]
	}
	return []float64{}
}

// AddProcess ...
func (ps *ProcessInfoSource) AddProcess(pi ProcessInfo) {
	ps.m.Lock()
	defer ps.m.Unlock()

	// check to see if the process is already in the map
	// (we dont really care that much for this example)
	if _, exists := ps.processes[pi.ID]; exists {
		fmt.Println("Process already exists:", pi)
		return
	}

	pp := &ProcessInfoEx{
		id:   pi.ID,
		name: pi.Name,
	}

	ps.processes[pi.ID] = pp
}

// RemoveProcess ...
func (ps *ProcessInfoSource) RemoveProcess(id string) {

	ps.m.Lock()
	defer ps.m.Unlock()

	_, ok := ps.processes[id]
	if ok {
		delete(ps.processes, id)
	}
}

// Stop ...
func (ps *ProcessInfoSource) Stop() {
	ps.doneCh <- true

	ps.sub.Stop()
}

func (ps *ProcessInfoSource) SendUpdate() {

	ps.m.Lock()
	defer ps.m.Unlock()

	// for each process in the list, use a random number to determine if cpu usage went up/down/unchangged
	// +/- 5
	for k, v := range ps.processes {
		r := (10.0 * rand.Float64()) - 5.0

		c := math.Min(100.0, math.Max(0.0, v.cpu+(r*2.0)))
		v.cpu = c
		v.cpuHistory = append(v.cpuHistory, c)

		// shouldnt have to do this any more
		ps.processes[k] = v

		x := &ProcessInfo{
			ID:   v.id,
			Name: v.name,
			CPU:  c,
		}

		ps.eb.Publish("PI", x)
	}
}
