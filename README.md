## Talos-Demo
Demonstration of using Golang and WebAssembly to create a responsive web application to potentially visualize kubernetes clusters. Uses a combination of REST and gRPC to access external data, as well as internal fake data generators for this demo.


### How to run
CD into talos-demo/client
```make run```

*optional*
To start gRPC server:
from the root (talos-demo) directory:
```go run server/server.go```

## Views
### Cluster view:
Demonstrates a simple subject/detail view. Vertical bar on the left side contains the subjects, and when selected, the area to the right will show more information for the selected subject. In this case, the user can select between control planes and worker nodes. The current view mode (planes, nodes) is encoded in the url.
Content is currently static.

### Machines View:
Early experiment with css and dynamic content. Features a blinking red/green status indicator (ooo fancy) Updates are generated withing the ui.machine itself.

### Events View:
This view displays a list of 'events' that are updated dynamically (once per second) from an external source. Purpose is to demonstrate a responsive UI that updates as the data being displayed is updated. Uses the event bus to both receive updates, and to request historical data when the view is mounted.

### Charts View:
Uses a 3rd party library (go-chart) to generate graphs showing process cpu usage over time. Data updates are received via EventBus. Historical data is requested from the model (via the EventBus) when the components are mounted.


## Tech:

#### Go-app
For building progressive web apps (PWA) with the Go programming language (Golang) and WebAssembly (Wasm).

[Go-app dev](https://go-app.dev)

[go-app github](https://github.com/maxence-charriere/go-app)


#### Go-chart
Charting package used by the Charts view.

[go-chart](https://github.com/wcharczuk/go-chart)


#### EventBus
Publish-subscribe pattern used to distribute data and updates from the models to the views (and provides a way for the views to request data as well.)


#### things to explore:
WebSockets

[go-echarts](https://github.com/go-echarts/go-echarts)