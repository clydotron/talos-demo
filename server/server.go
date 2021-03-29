package main

import (
	"fmt"
	"log"
	"net"

	"github.com/clydotron/talos-demo/api/cluster"
	"google.golang.org/grpc"
)

type clusterServer struct {
	cluster.UnimplementedClusterServiceServer
	counter int
}

func (cs *clusterServer) HealthCheck(req *cluster.HealthCheckRequest, stream cluster.ClusterService_HealthCheckServer) error {

	fmt.Println("HealthCheck:", req, "counter", cs.counter)

	//@todo this includes a timeout/duration, implement a timer and pass

	status := "idle"
	errorMsg := "none"
	if cs.counter < 2 {
		status = "starting"
	} else if cs.counter < 3 {
		status = "running"
	} else if cs.counter < 4 {
		status = "stopped"
		errorMsg = "crash!"
	}
	cs.counter++

	// send info for each control plane
	cplanes := req.ClusterInfo.GetControlPlaneNodes()
	for _, cplane := range cplanes {

		resp := &cluster.HealthCheckProgress{
			Metadata: &cluster.Metadata{
				Hostname: cplane,
				Error:    errorMsg,
			},
			Message: status,
		}
		stream.Send(resp)
	}

	// send info for each worker node
	//wnodes := req.ClusterInfo.GetWorkerNodes()
	for _, node := range req.ClusterInfo.GetWorkerNodes() {
		resp := &cluster.HealthCheckProgress{
			Metadata: &cluster.Metadata{
				Hostname: node,
				Error:    "none",
			},
			Message: "active",
		}
		stream.Send(resp)
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	cluster.RegisterClusterServiceServer(s, &clusterServer{})

	fmt.Println("gRPC server ACTIVE")

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
