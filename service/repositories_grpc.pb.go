// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: repositories.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Repo_GetRepos_FullMethodName = "/Repo/GetRepos"
)

// RepoClient is the client API for Repo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoClient interface {
	GetRepos(ctx context.Context, in *RepoGetRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RepoGetReply], error)
}

type repoClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoClient(cc grpc.ClientConnInterface) RepoClient {
	return &repoClient{cc}
}

func (c *repoClient) GetRepos(ctx context.Context, in *RepoGetRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RepoGetReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Repo_ServiceDesc.Streams[0], Repo_GetRepos_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RepoGetRequest, RepoGetReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Repo_GetReposClient = grpc.ServerStreamingClient[RepoGetReply]

// RepoServer is the server API for Repo service.
// All implementations must embed UnimplementedRepoServer
// for forward compatibility.
type RepoServer interface {
	GetRepos(*RepoGetRequest, grpc.ServerStreamingServer[RepoGetReply]) error
	mustEmbedUnimplementedRepoServer()
}

// UnimplementedRepoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRepoServer struct{}

func (UnimplementedRepoServer) GetRepos(*RepoGetRequest, grpc.ServerStreamingServer[RepoGetReply]) error {
	return status.Errorf(codes.Unimplemented, "method GetRepos not implemented")
}
func (UnimplementedRepoServer) mustEmbedUnimplementedRepoServer() {}
func (UnimplementedRepoServer) testEmbeddedByValue()              {}

// UnsafeRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoServer will
// result in compilation errors.
type UnsafeRepoServer interface {
	mustEmbedUnimplementedRepoServer()
}

func RegisterRepoServer(s grpc.ServiceRegistrar, srv RepoServer) {
	// If the following call pancis, it indicates UnimplementedRepoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Repo_ServiceDesc, srv)
}

func _Repo_GetRepos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RepoGetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepoServer).GetRepos(m, &grpc.GenericServerStream[RepoGetRequest, RepoGetReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Repo_GetReposServer = grpc.ServerStreamingServer[RepoGetReply]

// Repo_ServiceDesc is the grpc.ServiceDesc for Repo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Repo",
	HandlerType: (*RepoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRepos",
			Handler:       _Repo_GetRepos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "repositories.proto",
}