// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greetpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greetpb.GreetService/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Greet(*GreetRequest, GreetService_GreetServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet(*GreetRequest, GreetService_GreetServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Greet(m, &greetServiceGreetServer{stream})
}

type GreetService_GreetServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greetServiceGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greetpb.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _GreetService_Greet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/gRPC/greet.proto",
}
