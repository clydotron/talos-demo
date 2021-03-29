package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/client/ui"
	"github.com/clydotron/talos-demo/client/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Clusters struct {
	app.Compo

	sub  *utils.EventBusSubscriber
	eb   *utils.EventBus
	view string

	ci *models.ClusterInfo
}

// NewClustersView
func NewClustersView(eb *utils.EventBus) *Clusters {
	e := &Clusters{
		eb:   eb,
		sub:  utils.NewEventBusSubscriber("cluster", eb),
		view: "planes",
	}
	return e
}

func (c *Clusters) handleEvent(d utils.DataEvent) {
	// @todo process any updates to the cluster info here
}

func (c *Clusters) OnNav(ctx app.Context, u *url.URL) {

	// extract the view
	m, err := url.ParseQuery(u.RawQuery)
	if err == nil {
		if val, ok := m["view"]; ok {
			c.view = val[0]
		}

	} else {
		c.view = "planes"
	}

	go c.updateClusterInfo()
}

// updateClusterInfo send a request to the API to get information on the cluster
func (c *Clusters) updateClusterInfo() {

	resp, err := http.Get("http://localhost:8000/api/v1/cluster")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	ci := &models.ClusterInfo{}
	err = json.NewDecoder(resp.Body).Decode(ci)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// update the cluster info on the main ui thread
	app.Dispatch(func() {
		c.ci = ci
		c.Update()
	})
}

func (c *Clusters) OnMount(ctx app.Context) {
	// need a way to get relevant information up to this point:
	// onNav takes care of that
	c.sub.Start(c.handleEvent)
}

func (c *Clusters) OnDismount() {
	defer fmt.Println("Clusters dismounted")
	c.sub.Stop()
}

// here is the question: is it better to have

func (c *Clusters) renderDetailView() app.UI {

	// no cluster info, nothing to render
	if c.ci == nil {
		return nil
	}

	// if planes, no grid?

	// make the number of columns
	return app.Div().Class("bg-gray-100 w-full h-full p-4 grid grid-cols-3 gap-2").
		Body(
			app.Ul().
				Body(
					app.If(c.view == "planes",
						app.Range(c.ci.ControlPlanes).Map(func(k string) app.UI {
							return app.Li().
								Body(
									&ui.ControlPlane{CPI: c.ci.ControlPlanes[k]},
								)
						}),
					).Else(app.Range(c.ci.WorkerNodes).Map(func(k string) app.UI {
						return app.Li().
							Body(
								&ui.WorkerNode{WNI: c.ci.WorkerNodes[k]},
							)
					}),
					),
				),
		)
}

//border-l-4 border-blue-dark
func (c *Clusters) getBorder(v string) string {
	if c.view == v {
		return "border-l-4 border-blue-500"
	}
	return "border-l-4 border-transparent"
}

func (c *Clusters) Render() app.UI {
	pborder := c.getBorder("planes")
	nborder := c.getBorder("nodes")

	return app.Div().Class("h-screen w-screen bg-gray-darker").Body(
		&ui.NavBar{},
		app.Div().Class("pt-16 h-full").ID("main").
			Body(
				// sidebar
				app.Div().Class("bg-gray-600 relative h-full").
					Body(
						app.Div().Class("xl:py-2").
							Body(
								app.Div().Class("group relative sidebar-item with-children text-white").Body(
									app.A().Href("/clusters?view=planes").Class("block xl:flex items-center text-center xl:text-left shadow-light xl:shadow-none py-6 xl:py-2 xl:px-4 hover:bg-gray-700 "+pborder).
										Body(
											app.Raw(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" class="h-6 w-6 text-grey-600 fill-current mx-auto xl:mr-2"><path d="M12 22a10 10 0 1 1 0-20 10 10 0 0 1 0 20zM5.68 7.1A7.96 7.96 0 0 0 4.06 11H5a1 1 0 0 1 0 2h-.94a7.95 7.95 0 0 0 1.32 3.5A9.96 9.96 0 0 1 11 14.05V9a1 1 0 0 1 2 0v5.05a9.96 9.96 0 0 1 5.62 2.45 7.95 7.95 0 0 0 1.32-3.5H19a1 1 0 0 1 0-2h.94a7.96 7.96 0 0 0-1.62-3.9l-.66.66a1 1 0 1 1-1.42-1.42l.67-.66A7.96 7.96 0 0 0 13 4.06V5a1 1 0 0 1-2 0v-.94c-1.46.18-2.8.76-3.9 1.62l.66.66a1 1 0 0 1-1.42 1.42l-.66-.67zM6.71 18a7.97 7.97 0 0 0 10.58 0 7.97 7.97 0 0 0-10.58 0z" class="heroicon-ui"></path></svg>`),
											app.Div().Class("text-white text-xs").Text("Control Planes"),
										),
								),
								app.Div().Class("group relative sidebar-item with-children text-white").Body(
									app.A().Href("/clusters?view=nodes").Class("block xl:flex xl:items-center text-center xl:text-left shadow-light xl:shadow-none py-6 xl:py-2 xl:px-4 hover:bg-gray-700 "+nborder).
										Body(
											app.Raw(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" class="h-6 w-6 text-grey-700 fill-current mx-auto xl:mr-2"><path d="M15 19a3 3 0 0 1-6 0H4a1 1 0 0 1 0-2h1v-6a7 7 0 0 1 4.02-6.34 3 3 0 0 1 5.96 0A7 7 0 0 1 19 11v6h1a1 1 0 0 1 0 2h-5zm-4 0a1 1 0 0 0 2 0h-2zm0-12.9A5 5 0 0 0 7 11v6h10v-6a5 5 0 0 0-4-4.9V5a1 1 0 0 0-2 0v1.1z" class="heroicon-ui"></path></svg>`),
											app.Div().Class("text-white text-xs").Text("Worker Nodes"),
										),
								),
							),
					),
				//content
				c.renderDetailView(),
			),
	)
}
