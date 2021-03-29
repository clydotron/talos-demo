// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cluster

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (ClusterService_HealthCheckClient, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (ClusterService_HealthCheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClusterService_ServiceDesc.Streams[0], "/cluster.ClusterService/HealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterServiceHealthCheckClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClusterService_HealthCheckClient interface {
	Recv() (*HealthCheckProgress, error)
	grpc.ClientStream
}

type clusterServiceHealthCheckClient struct {
	grpc.ClientStream
}

func (x *clusterServiceHealthCheckClient) Recv() (*HealthCheckProgress, error) {
	m := new(HealthCheckProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations must embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	HealthCheck(*HealthCheckRequest, ClusterService_HealthCheckServer) error
	mustEmbedUnimplementedClusterServiceServer()
}

// UnimplementedClusterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) HealthCheck(*HealthCheckRequest, ClusterService_HealthCheckServer) error {
	return status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedClusterServiceServer) mustEmbedUnimplementedClusterServiceServer() {}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_HealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClusterServiceServer).HealthCheck(m, &clusterServiceHealthCheckServer{stream})
}

type ClusterService_HealthCheckServer interface {
	Send(*HealthCheckProgress) error
	grpc.ServerStream
}

type clusterServiceHealthCheckServer struct {
	grpc.ServerStream
}

func (x *clusterServiceHealthCheckServer) Send(m *HealthCheckProgress) error {
	return x.ServerStream.SendMsg(m)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HealthCheck",
			Handler:       _ClusterService_HealthCheck_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cluster/cluster.proto",
}
