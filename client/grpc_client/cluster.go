package grpc_client

import (
	"fmt"

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

	fmt.Println("ClusterClient Connecting to:", server)
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial:", err) //pick one or the other (log fatal or return err)
		return err
	}

	cc.cconn = conn
	cc.CSC = cluster.NewClusterServiceClient(conn)

	return nil
}

func (cc *ClusterClient) Close() {
	cc.cconn.Close()
}
