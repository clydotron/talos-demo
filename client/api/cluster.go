package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clydotron/talos-demo/client/grpc_client"
	"github.com/clydotron/talos-demo/client/models"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type ClusterAPI struct {
	cc *grpc_client.ClusterClient
	ct *models.ClusterTracker
}

//"localhost:50051"
func (c *ClusterAPI) Init(addr string) {

	fmt.Println("Init...")

	client := grpc_client.NewClusterClient()
	err := client.Connect(addr)
	if err != nil {
		fmt.Println("### failed to connect to:", addr)
		//log.Fatalln("### Client failed to connect:", err)
		return
	}
	//defer client.Close()

	clusterTracker := models.NewClusterTracker(client)
	clusterTracker.InitWithFakeData()
	clusterTracker.Start()
	//defer clusterTracker.Stop()

	c.cc = client
	c.ct = clusterTracker
}

func (cx *ClusterAPI) Close() {
	if cx.cc == nil {
		return
	}

	cx.cc.Close()
	cx.ct.Stop()
}

func (cx *ClusterAPI) HandleClusterReq(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("api request: cluster")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	// make sure this is a get
	if (*r).Method != "GET" {
		fmt.Println("Incorrect method:", (*r).Method)
		return
	}

	if cx.ct != nil {
		//get the latest cluster info from the store and json it...
		cx.ct.UpdateStatus()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cx.ct.CI)
}
