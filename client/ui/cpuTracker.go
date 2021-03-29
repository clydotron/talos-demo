package ui

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	//exposes "chart"
)

type CpuTracker struct {
	app.Compo

	id  string
	eb  *utils.EventBus
	sub *utils.EventBusSubscriber

	history     []float64
	encChart    string
	FillColor   string
	StrokeColor string
}

func NewCpuTracker(id string, eb *utils.EventBus) *CpuTracker {

	tt := &CpuTracker{id: id,
		eb:  eb,
		sub: utils.NewEventBusSubscriber("PI", eb),
	}
	tt.FillColor = "00dd00"
	tt.StrokeColor = "33ff88"

	return tt
}

func (c *CpuTracker) handleEvent(d utils.DataEvent) {
	if pi, ok := d.Data.(*models.ProcessInfo); ok {
		if pi.ID == c.id {
			c.history = append(c.history, pi.CPU)
			go c.updateChart()
		}
	} else {
		//programming error:
		log.Fatal("failed to convert to ProcessInfo")
	}
}

func (c *CpuTracker) updateChart() {

	// @todo make color configurable
	const kMaxSize = 50

	yvals := []float64{}

	hs := len(c.history)
	si := 0
	if hs > kMaxSize {
		si = hs - kMaxSize
		yvals = c.history[si:]
	} else {
		yvals = c.history
		// now fill in the remaining slots with zeros
		zvals := make([]float64, kMaxSize-hs)
		yvals = append(yvals, zvals...)
		//yvals = append(yvals, 0.0)
	}

	// make the sequence for x:
	// if history size < maxSize, add a double at the size...
	xvals := []float64{}
	for i := 0; i < kMaxSize; i++ {
		xvals = append(xvals, float64(i))
	}

	if hs < kMaxSize {
		xvals[hs] = float64(hs - 1)
	}

	//chart.ValueSequence()
	// do the heavy lifting off the main thread. Update the encoded image once we are ready
	//fmt.Println("updating chart!")
	graph := chart.Chart{
		Height: 150,
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: 100.0,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: drawing.ColorFromHex(c.FillColor),                 // will supercede defaults
					FillColor:   drawing.ColorFromHex(c.StrokeColor).WithAlpha(64), // will supercede defaults
				},
				YValues: chart.Seq{Sequence: chart.ValueSequence(yvals...)}.Values(),
				XValues: chart.Seq{Sequence: chart.ValueSequence(xvals...)}.Values(),
				//XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1).WithEnd(kMaxSize)}.Values(),
				//YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(50).WithMax(150)}.Values(),
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	enc := base64.StdEncoding.EncodeToString(buffer.Bytes())

	app.Dispatch(func() {
		c.encChart = enc
		c.Update()
	})
}

func (c *CpuTracker) OnMount(ctx app.Context) {
	//fmt.Println("CpuTracker.OnMount: ")

	req := &models.ProcessInfoCpuHistoryRequest{
		ID:  c.id,
		Num: 50,
		Callback: func(h []float64) {
			app.Dispatch(func() {
				c.history = append(c.history, h...)
				c.Update()
			})
		},
	}
	c.eb.Publish("pi_req", req)

	c.sub.Start(c.handleEvent)
}

// OnDismount ...
func (c *CpuTracker) OnDismount() {
	//defer fmt.Println("CpuTracker dismounted")
	c.sub.Stop()
}

func (c *CpuTracker) Render() app.UI {
	h := ""
	if c.encChart == "" {
		h = " hidden"
		//return nil
	}

	enc := "data:image/png;base64, " + c.encChart

	return app.Div().
		Class("w-screen p-2"+h).
		Body(
			app.Text(fmt.Sprintln("CPU usage:", c.id)),
			app.Img().Src(enc).Alt("cpu"),
		)

}
