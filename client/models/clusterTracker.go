package models

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/clydotron/talos-demo/api/cluster"
	"github.com/clydotron/talos-demo/grpc_client/client"
)

// periodically ping the grpc server for cluster information:
// notify any subscribers over the event bus of any changes

type ClusterTracker struct {
	//eb     *utils.EventBus
	cc     *client.ClusterClient
	ticker *time.Ticker
	doneCh chan bool
	CI     *ClusterInfo
}

// NewClusterTracker ...
func NewClusterTracker(cc *client.ClusterClient) *ClusterTracker {
	ct := &ClusterTracker{
		//eb: eb,
		cc: cc,
		CI: &ClusterInfo{
			ControlPlanes: make(map[string]ControlPlaneInfo),
			WorkerNodes:   make(map[string]WorkerNodeInfo),
		},
	}
	return ct
}

// InitWithFakeData ...
func (ct *ClusterTracker) InitWithFakeData() {
	ct.CI.ControlPlanes["Control 1"] = ControlPlaneInfo{Name: "Control 1", Status: "Stopped"}
	ct.CI.WorkerNodes["Node 1"] = WorkerNodeInfo{Name: "Node 1", Status: "Stopped"}
	ct.CI.WorkerNodes["Node 2"] = WorkerNodeInfo{Name: "Node 2", Status: "Stopped"}
}

// Start ...
func (ct *ClusterTracker) Start() {

	fmt.Println("ClusterTracker - Start")

	ct.UpdateStatus()

	// ct.ticker = time.NewTicker(2000 * time.Millisecond)
	// ct.doneCh = make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ct.doneCh:
	// 			return
	// 		case <-ct.ticker.C:
	// 			ct.UpdateStatus()
	// 		}
	// 	}
	// }()
}

// Stop ...
func (ct *ClusterTracker) Stop() {
	fmt.Println("ClusterTracker - Stop")

	ct.ticker.Stop()
	ct.doneCh <- true
}

// UpdateStatus ...
func (ct *ClusterTracker) UpdateStatus() {

	//
	planes := []string{}
	nodes := []string{}

	for k := range ct.CI.ControlPlanes {
		planes = append(planes, k)
	}

	for k := range ct.CI.WorkerNodes {
		nodes = append(nodes, k)
	}

	//@todo why dont i just make a copy? (or pass the cluter info?)

	req := &cluster.HealthCheckRequest{
		ClusterInfo: &cluster.ClusterInfo{
			ControlPlaneNodes: planes,
			WorkerNodes:       nodes,
		},
		//WaitTimeout: ,
	}

	stream, err := ct.cc.CSC.HealthCheck(context.Background(), req)
	if err != nil {
		//og.Fatalln("HealthCheck RPC error:", err)
		fmt.Println("#### HealthCheck RPC error:", err)
		ct.Stop()
		return
	}
	for {
		msg, err := stream.Recv() // _ -> msg
		if err == io.EOF {
			// reached end of stream
			break
		}

		if err != nil {
			log.Fatalln("stream Recv error:", err)
		}

		// @todo do a diff to calculate what changed
		//
		// handle the actual result:
		d := msg.GetMetadata()
		x := msg.GetMessage()
		//fmt.Println("msg: meta", d, "msg:", x)

		// check to see if this is a control plane... if so, update
		if cp, ok := ct.CI.ControlPlanes[d.Hostname]; ok {
			cp.Status = x
			//compare old and new...
			ct.CI.ControlPlanes[d.Hostname] = cp
		} else if wn, ok := ct.CI.WorkerNodes[d.Hostname]; ok {
			wn.Status = x
			ct.CI.WorkerNodes[d.Hostname] = wn
		}
		//@todo store this locally... check if values have changed, set a flag
	}

}
