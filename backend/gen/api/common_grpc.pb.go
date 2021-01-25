// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package myapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonServiceClient interface {
	//Command endpoint
	AddField(ctx context.Context, in *AddFieldParams, opts ...grpc.CallOption) (*FieldObject, error)
	//Query endpoint
	GetField(ctx context.Context, in *GetFieldParams, opts ...grpc.CallOption) (CommonService_GetFieldClient, error)
}

type commonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonServiceClient(cc grpc.ClientConnInterface) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) AddField(ctx context.Context, in *AddFieldParams, opts ...grpc.CallOption) (*FieldObject, error) {
	out := new(FieldObject)
	err := c.cc.Invoke(ctx, "/CommonService/AddField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) GetField(ctx context.Context, in *GetFieldParams, opts ...grpc.CallOption) (CommonService_GetFieldClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommonService_serviceDesc.Streams[0], "/CommonService/GetField", opts...)
	if err != nil {
		return nil, err
	}
	x := &commonServiceGetFieldClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommonService_GetFieldClient interface {
	Recv() (*FieldObject, error)
	grpc.ClientStream
}

type commonServiceGetFieldClient struct {
	grpc.ClientStream
}

func (x *commonServiceGetFieldClient) Recv() (*FieldObject, error) {
	m := new(FieldObject)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommonServiceServer is the server API for CommonService service.
// All implementations must embed UnimplementedCommonServiceServer
// for forward compatibility
type CommonServiceServer interface {
	//Command endpoint
	AddField(context.Context, *AddFieldParams) (*FieldObject, error)
	//Query endpoint
	GetField(*GetFieldParams, CommonService_GetFieldServer) error
	mustEmbedUnimplementedCommonServiceServer()
}

// UnimplementedCommonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (UnimplementedCommonServiceServer) AddField(context.Context, *AddFieldParams) (*FieldObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddField not implemented")
}
func (UnimplementedCommonServiceServer) GetField(*GetFieldParams, CommonService_GetFieldServer) error {
	return status.Errorf(codes.Unimplemented, "method GetField not implemented")
}
func (UnimplementedCommonServiceServer) mustEmbedUnimplementedCommonServiceServer() {}

// UnsafeCommonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServiceServer will
// result in compilation errors.
type UnsafeCommonServiceServer interface {
	mustEmbedUnimplementedCommonServiceServer()
}

func RegisterCommonServiceServer(s grpc.ServiceRegistrar, srv CommonServiceServer) {
	s.RegisterService(&_CommonService_serviceDesc, srv)
}

func _CommonService_AddField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFieldParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).AddField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/AddField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).AddField(ctx, req.(*AddFieldParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_GetField_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFieldParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommonServiceServer).GetField(m, &commonServiceGetFieldServer{stream})
}

type CommonService_GetFieldServer interface {
	Send(*FieldObject) error
	grpc.ServerStream
}

type commonServiceGetFieldServer struct {
	grpc.ServerStream
}

func (x *commonServiceGetFieldServer) Send(m *FieldObject) error {
	return x.ServerStream.SendMsg(m)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddField",
			Handler:    _CommonService_AddField_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetField",
			Handler:       _CommonService_GetField_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common.proto",
}
