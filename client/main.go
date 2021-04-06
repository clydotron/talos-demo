// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/clydotron/talos-demo/client/api"
	"github.com/clydotron/talos-demo/client/ws"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	//"github.com/wcharczuk/go-chart" //exposes "chart"
)

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {

	serverAddr := flag.String("grpc", "localhost:50051", "gRPC server address")
	port := flag.String("port", ":8000", "Port to listen on (default 8000)")

	flag.Parse()

	// ClusterAPI is responsible for establishing a gRPC connect to get cluster information
	// also implements a handleFunc to return this info as JSON to the caller (from inside wasm-land)
	cx := &api.ClusterAPI{}
	cx.Init(*serverAddr)
	defer cx.Close()

	// Test websocket server
	tsx := ws.NewTestServerX()
	http.HandleFunc("/echo", tsx.ServeHttp)
	defer tsx.Close()

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

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// create a new server, have it listen and serve in its own go routine
	srv := &http.Server{Addr: *port}

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	fmt.Println("up and running...")

	// listen for an interrupt from the OS.
	// when we get one, shutdown the server. When we leave this function, the deferred
	// close/cleanup methods will be triggered for the gRPC and Websockets
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-interrupt:
		fmt.Println("### interrupted ###")
		if err := srv.Shutdown(context.TODO()); err != nil {
			panic(err) // failure/timeout shutting down the server gracefully
		}
	}

	// wait for goroutine to stop
	wg.Wait()

	fmt.Println("All done.")
	return nil
}
