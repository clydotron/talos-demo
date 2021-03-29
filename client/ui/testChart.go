package ui

import (
	"bytes"
	"encoding/base64"
	"net/url"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/wcharczuk/go-chart" //exposes "chart"
)

type TestChartX struct {
	app.Compo

	encodedPC string
}

func (c *TestChartX) OnNav(ctx app.Context, url *url.URL) {

	// get the info for node x

	c.renderPieChart()
}

func (c *TestChartX) renderPieChart() {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 1, Label: "!!"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	pie.Render(chart.PNG, buffer)

	c.encodedPC = base64.StdEncoding.EncodeToString(buffer.Bytes())
	c.Update()
}

func (c *TestChartX) Render() app.UI {
	h := ""
	if c.encodedPC == "" {
		h = " hidden"
	}
	enc := "data:image/png;base64, " + c.encodedPC

	return app.Div().
		Class("bg-purple-400").
		Body(
			app.Text("Chart"),
			app.Img().Src("/api/v1/blue").Class("p-2").Alt("blue"),
			app.Img().Src("/api/v1/red").Class("p-2").Alt("red"),
			app.Img().Src(enc).Class("p-2"+h).Alt("orange"),
			app.Img().Src(`data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUA
			AAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO
				9TXL0Y4OHwAAAABJRU5ErkJggg==`).Alt("Red dot").Class("p-2"),
		)
}
