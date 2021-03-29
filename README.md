

Go / wasm app
For building progressive web apps (PWA) with the Go programming language (Golang) and WebAssembly (Wasm).

[]

[optional]
To start gRPC server:
from the root (talos-demo) directory:
go run server/server.go

## Tech:
### EventBus - publish-subscribe pattern - 
used to distribute data and updates from the models to the views (and provides a way for the views to request data as well)

### Cluster view:
Demonstrates a simple subject/detail view. Vertical bar on the left side contains the subjects, and when selected, the area to the right will show more information for the selected subject. In this case, the user can select between Control plans and worker nodes. The current view mode (planes, nodes) is encoded in the url
Content is currently static.

### Machines View:
Early experiment with css and dynamic content. Features a blinking red/green status indicator (ooo fancy) Updates are generated withing the ui.machine itself

### Events View:
This view displays a list of 'events' that are updated dynamically (once per second) from an external source. Purpose is to demonstrate a responsive UI that updates as the data being displayed is updated. Uses the event bus to both receive updates, and to request historical data when the view is mounted.

Charts View:
Uses a 3rd party library (go-chart) to generate graphs showing process cpu usage over time. Data updates are received via EventBus. Historical data is requested via the 
Go-app
[Go-app dev](https://go-app.dev)
https://github.com/maxence-charriere/go-app

[go-chart](https://github.com/wcharczuk/go-chart)

things to explore:
https://github.com/go-echarts/go-echarts