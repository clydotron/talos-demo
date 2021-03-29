// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/clydotron/talos-demo/client/models"
	"github.com/clydotron/talos-demo/grpc_client/client"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	//"github.com/wcharczuk/go-chart" //exposes "chart"
)

func lager(format string, v ...interface{}) {
	fmt.Println(format, v)

}

type ClientX struct {
	cc *client.ClusterClient
	ct *models.ClusterTracker
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// apiCluster -- this is called on its own go routine, so we can block
func (cx *ClientX) apiCluster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api request: cluster")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if (*r).Method != "GET" {
		fmt.Println("Incorrect method:", (*r).Method)
		return
	}
	//make sure this is a get
	//get the latest cluster info from the store and json it...

	cx.ct.UpdateStatus()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cx.ct.CI)
}

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	// could move all of this into init
	client := client.NewClusterClient()
	err := client.Connect("localhost:50051")
	if err != nil {
		log.Fatalln("### Client failed to connect:", err)
	}
	defer client.Close()

	clusterTracker := models.NewClusterTracker(client)
	clusterTracker.InitWithFakeData()
	clusterTracker.Start()
	defer clusterTracker.Stop()

	cx := &ClientX{
		cc: client,
		ct: clusterTracker,
	}
	// sequence of events:
	// create client (gRPC connection to server)
	// create the cluster tracker - responsible for making the HealthCheck gRPC call and [optionally] maintaining a snaphot of the cluster (so can diff)
	// ClientX(Api) - implements the http.HandleFunc call to respond to api requests

	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "Experimental",
		Styles: []string{
			"/web/tailwind.css", "/web/test.css", // Include .css file.
		},
	})

	http.HandleFunc("/api/v1/cluster", cx.apiCluster)

	fmt.Println("up and running...")

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
