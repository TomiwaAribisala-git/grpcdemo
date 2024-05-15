// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/greet.proto

package proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Sayhello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	Sayhelloclientstreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayhelloclientstreamingClient, error)
	Sayhelloserverstreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayhelloserverstreamingClient, error)
	Sayhellobidirectionalstreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayhellobidirectionalstreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Sayhello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreetService/sayhello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) Sayhelloclientstreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayhelloclientstreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet_service.GreetService/sayhelloclientstreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayhelloclientstreamingClient{stream}
	return x, nil
}

type GreetService_SayhelloclientstreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceSayhelloclientstreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayhelloclientstreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayhelloclientstreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) Sayhelloserverstreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_SayhelloserverstreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet_service.GreetService/sayhelloserverstreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayhelloserverstreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayhelloserverstreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayhelloserverstreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayhelloserverstreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) Sayhellobidirectionalstreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayhellobidirectionalstreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet_service.GreetService/sayhellobidirectionalstreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayhellobidirectionalstreamingClient{stream}
	return x, nil
}

type GreetService_SayhellobidirectionalstreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayhellobidirectionalstreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayhellobidirectionalstreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayhellobidirectionalstreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Sayhello(context.Context, *NoParam) (*HelloResponse, error)
	Sayhelloclientstreaming(GreetService_SayhelloclientstreamingServer) error
	Sayhelloserverstreaming(*NamesList, GreetService_SayhelloserverstreamingServer) error
	Sayhellobidirectionalstreaming(GreetService_SayhellobidirectionalstreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Sayhello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayhello not implemented")
}
func (UnimplementedGreetServiceServer) Sayhelloclientstreaming(GreetService_SayhelloclientstreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method Sayhelloclientstreaming not implemented")
}
func (UnimplementedGreetServiceServer) Sayhelloserverstreaming(*NamesList, GreetService_SayhelloserverstreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method Sayhelloserverstreaming not implemented")
}
func (UnimplementedGreetServiceServer) Sayhellobidirectionalstreaming(GreetService_SayhellobidirectionalstreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method Sayhellobidirectionalstreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Sayhello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Sayhello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreetService/sayhello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Sayhello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_Sayhelloclientstreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).Sayhelloclientstreaming(&greetServiceSayhelloclientstreamingServer{stream})
}

type GreetService_SayhelloclientstreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayhelloclientstreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayhelloclientstreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayhelloclientstreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_Sayhelloserverstreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Sayhelloserverstreaming(m, &greetServiceSayhelloserverstreamingServer{stream})
}

type GreetService_SayhelloserverstreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceSayhelloserverstreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayhelloserverstreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_Sayhellobidirectionalstreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).Sayhellobidirectionalstreaming(&greetServiceSayhellobidirectionalstreamingServer{stream})
}

type GreetService_SayhellobidirectionalstreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayhellobidirectionalstreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayhellobidirectionalstreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayhellobidirectionalstreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayhello",
			Handler:    _GreetService_Sayhello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sayhelloclientstreaming",
			Handler:       _GreetService_Sayhelloclientstreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "sayhelloserverstreaming",
			Handler:       _GreetService_Sayhelloserverstreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "sayhellobidirectionalstreaming",
			Handler:       _GreetService_Sayhellobidirectionalstreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}
