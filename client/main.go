// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/clydotron/talos-demo/client/api"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	serverAddr := flag.String("grpc", "localhost:50051", "gRPC server address")
	port := flag.String("port", ":8000", "Port to listen on (default 8000)")

	flag.Parse()

	// ClusterAPI is responsible for establishing a gRPC connect to get cluster information
	// also implements a handleFunc to return this info as JSON to the caller (from inside wasm-land)
	cx := &api.ClusterAPI{}
	cx.Init(*serverAddr)
	defer cx.Close()

	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "Experimental",
		Styles: []string{
			"/web/tailwind.css", "/web/test.css",
		},
	})

	// handle any API reqests (since wasm doesnt support gRPC)
	http.HandleFunc("/api/v1/cluster", cx.HandleClusterReq)

	fmt.Println("up and running...")

	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
