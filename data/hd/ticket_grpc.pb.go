// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: ticket.proto

package hd_ticket

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

// HDTabTicketServiceClient is the client API for HDTabTicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HDTabTicketServiceClient interface {
	CreateNewTicketTable(ctx context.Context, in *CreateNewTicketTableRequest, opts ...grpc.CallOption) (*CreateNewTicketTableResponse, error)
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
}

type hDTabTicketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHDTabTicketServiceClient(cc grpc.ClientConnInterface) HDTabTicketServiceClient {
	return &hDTabTicketServiceClient{cc}
}

func (c *hDTabTicketServiceClient) CreateNewTicketTable(ctx context.Context, in *CreateNewTicketTableRequest, opts ...grpc.CallOption) (*CreateNewTicketTableResponse, error) {
	out := new(CreateNewTicketTableResponse)
	err := c.cc.Invoke(ctx, "/HDTabTicketService/CreateNewTicketTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hDTabTicketServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, "/HDTabTicketService/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hDTabTicketServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	out := new(GetTicketResponse)
	err := c.cc.Invoke(ctx, "/HDTabTicketService/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HDTabTicketServiceServer is the server API for HDTabTicketService service.
// All implementations must embed UnimplementedHDTabTicketServiceServer
// for forward compatibility
type HDTabTicketServiceServer interface {
	CreateNewTicketTable(context.Context, *CreateNewTicketTableRequest) (*CreateNewTicketTableResponse, error)
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	mustEmbedUnimplementedHDTabTicketServiceServer()
}

// UnimplementedHDTabTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHDTabTicketServiceServer struct {
}

func (UnimplementedHDTabTicketServiceServer) CreateNewTicketTable(context.Context, *CreateNewTicketTableRequest) (*CreateNewTicketTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTicketTable not implemented")
}
func (UnimplementedHDTabTicketServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedHDTabTicketServiceServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedHDTabTicketServiceServer) mustEmbedUnimplementedHDTabTicketServiceServer() {}

// UnsafeHDTabTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HDTabTicketServiceServer will
// result in compilation errors.
type UnsafeHDTabTicketServiceServer interface {
	mustEmbedUnimplementedHDTabTicketServiceServer()
}

func RegisterHDTabTicketServiceServer(s grpc.ServiceRegistrar, srv HDTabTicketServiceServer) {
	s.RegisterService(&HDTabTicketService_ServiceDesc, srv)
}

func _HDTabTicketService_CreateNewTicketTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewTicketTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HDTabTicketServiceServer).CreateNewTicketTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HDTabTicketService/CreateNewTicketTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HDTabTicketServiceServer).CreateNewTicketTable(ctx, req.(*CreateNewTicketTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HDTabTicketService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HDTabTicketServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HDTabTicketService/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HDTabTicketServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HDTabTicketService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HDTabTicketServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HDTabTicketService/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HDTabTicketServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HDTabTicketService_ServiceDesc is the grpc.ServiceDesc for HDTabTicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HDTabTicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HDTabTicketService",
	HandlerType: (*HDTabTicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewTicketTable",
			Handler:    _HDTabTicketService_CreateNewTicketTable_Handler,
		},
		{
			MethodName: "CreateTicket",
			Handler:    _HDTabTicketService_CreateTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _HDTabTicketService_GetTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}