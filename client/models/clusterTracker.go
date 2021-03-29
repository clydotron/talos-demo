package models

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/clydotron/talos-demo/api/cluster"
	"github.com/clydotron/talos-demo/client/grpc_client"
)

// periodically ping the grpc server for cluster information:
// notify any subscribers over the event bus of any changes
// >work in progress< currently makes a single call to HealthCheck
// should broadcast over the event bus
type ClusterTracker struct {
	cc *grpc_client.ClusterClient
	CI *ClusterInfo
}

// NewClusterTracker ...
func NewClusterTracker(cc *grpc_client.ClusterClient) *ClusterTracker {
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
	ct.CI.ControlPlanes["Control 1"] = ControlPlaneInfo{Name: "Control Plane 1", Status: "Stopped"}
	ct.CI.WorkerNodes["Node 1"] = WorkerNodeInfo{Name: "Worker Node 1", Status: "Stopped"}
	ct.CI.WorkerNodes["Node 2"] = WorkerNodeInfo{Name: "Worker Node 2", Status: "Stopped"}
}

// Start ...
func (ct *ClusterTracker) Start() {
	ct.UpdateStatus()
}

// Stop ...
func (ct *ClusterTracker) Stop() {
	// @todo either add something here, or remove this
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

	req := &cluster.HealthCheckRequest{
		ClusterInfo: &cluster.ClusterInfo{
			ControlPlaneNodes: planes,
			WorkerNodes:       nodes,
		},
		//WaitTimeout: ,
	}

	stream, err := ct.cc.CSC.HealthCheck(context.Background(), req)
	if err != nil {
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

		// check to see if this is a control plane... if so, update
		if cp, ok := ct.CI.ControlPlanes[d.Hostname]; ok {
			cp.Status = x
			//compare old and new...
			ct.CI.ControlPlanes[d.Hostname] = cp
		} else if wn, ok := ct.CI.WorkerNodes[d.Hostname]; ok {
			wn.Status = x
			ct.CI.WorkerNodes[d.Hostname] = wn
		}
	}
}
