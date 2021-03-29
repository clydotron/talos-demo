package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/clydotron/talos-demo/api/cluster"
	"google.golang.org/grpc"
)

type ClusterClient struct {
	cconn *grpc.ClientConn
	CSC   cluster.ClusterServiceClient
}

func NewClusterClient() *ClusterClient {

	cc := &ClusterClient{}

	return cc
}

func (cc *ClusterClient) Connect(server string) error {

	fmt.Println("ClusterClient Connecting...")
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial:", err) //pick one or the other (log fatal or return err)
		return err
	}
	//defer cc.Close()

	cc.cconn = conn
	cc.CSC = cluster.NewClusterServiceClient(conn)

	fmt.Println(">>> ClusterClient Connected >>>")
	//doHealthCheck(cc.CSC)

	return nil
}

func (cc *ClusterClient) Close() {
	cc.cconn.Close()
}

func doHealthCheck(csc cluster.ClusterServiceClient) {

	fmt.Println("testing connection...")

	controlPlanes := []string{"control plane 1", "control plane 2"}
	workerNodes := []string{"node 1", "node 2", "node 3", "node 4", "node 5"}

	req := &cluster.HealthCheckRequest{
		ClusterInfo: &cluster.ClusterInfo{
			ControlPlaneNodes: controlPlanes,
			WorkerNodes:       workerNodes,
		},
		//WaitTimeout: ,
	}

	stream, err := csc.HealthCheck(context.Background(), req)
	if err != nil {
		fmt.Println("### HealthCheck RPC error:", err)
		return
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// reached end of stream
			break
		}

		if err != nil {
			log.Fatalln("stream Recv error:", err)
		}
		// handle the actual result:
		d := msg.GetMetadata()
		x := msg.GetMessage()
		fmt.Println("msg:", d, x)
	}
}
