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
	BulkGetFields(ctx context.Context, in *BulkFields, opts ...grpc.CallOption) (CommonService_BulkGetFieldsClient, error)
	//Query endpoint
	GetField(ctx context.Context, in *GetFieldParams, opts ...grpc.CallOption) (*FieldObject, error)
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

func (c *commonServiceClient) BulkGetFields(ctx context.Context, in *BulkFields, opts ...grpc.CallOption) (CommonService_BulkGetFieldsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommonService_serviceDesc.Streams[0], "/CommonService/BulkGetFields", opts...)
	if err != nil {
		return nil, err
	}
	x := &commonServiceBulkGetFieldsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommonService_BulkGetFieldsClient interface {
	Recv() (*FieldObject, error)
	grpc.ClientStream
}

type commonServiceBulkGetFieldsClient struct {
	grpc.ClientStream
}

func (x *commonServiceBulkGetFieldsClient) Recv() (*FieldObject, error) {
	m := new(FieldObject)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commonServiceClient) GetField(ctx context.Context, in *GetFieldParams, opts ...grpc.CallOption) (*FieldObject, error) {
	out := new(FieldObject)
	err := c.cc.Invoke(ctx, "/CommonService/GetField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
// All implementations must embed UnimplementedCommonServiceServer
// for forward compatibility
type CommonServiceServer interface {
	//Command endpoint
	AddField(context.Context, *AddFieldParams) (*FieldObject, error)
	//Query endpoint
	BulkGetFields(*BulkFields, CommonService_BulkGetFieldsServer) error
	//Query endpoint
	GetField(context.Context, *GetFieldParams) (*FieldObject, error)
	mustEmbedUnimplementedCommonServiceServer()
}

// UnimplementedCommonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (UnimplementedCommonServiceServer) AddField(context.Context, *AddFieldParams) (*FieldObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddField not implemented")
}
func (UnimplementedCommonServiceServer) BulkGetFields(*BulkFields, CommonService_BulkGetFieldsServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkGetFields not implemented")
}
func (UnimplementedCommonServiceServer) GetField(context.Context, *GetFieldParams) (*FieldObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetField not implemented")
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

func _CommonService_BulkGetFields_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BulkFields)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommonServiceServer).BulkGetFields(m, &commonServiceBulkGetFieldsServer{stream})
}

type CommonService_BulkGetFieldsServer interface {
	Send(*FieldObject) error
	grpc.ServerStream
}

type commonServiceBulkGetFieldsServer struct {
	grpc.ServerStream
}

func (x *commonServiceBulkGetFieldsServer) Send(m *FieldObject) error {
	return x.ServerStream.SendMsg(m)
}

func _CommonService_GetField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFieldParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).GetField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/GetField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).GetField(ctx, req.(*GetFieldParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddField",
			Handler:    _CommonService_AddField_Handler,
		},
		{
			MethodName: "GetField",
			Handler:    _CommonService_GetField_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkGetFields",
			Handler:       _CommonService_BulkGetFields_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common.proto",
}
