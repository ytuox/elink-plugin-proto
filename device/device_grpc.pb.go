// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: device/device.proto

package device

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

const (
	RPCDevice_ConnectIotPlatform_FullMethodName      = "/device.RPCDevice/ConnectIotPlatform"
	RPCDevice_DisconnectIotPlatform_FullMethodName   = "/device.RPCDevice/DisconnectIotPlatform"
	RPCDevice_GetDeviceConnectStatus_FullMethodName  = "/device.RPCDevice/GetDeviceConnectStatus"
	RPCDevice_QueryDeviceList_FullMethodName         = "/device.RPCDevice/QueryDeviceList"
	RPCDevice_QueryDeviceById_FullMethodName         = "/device.RPCDevice/QueryDeviceById"
	RPCDevice_QueryDeviceListByUserId_FullMethodName = "/device.RPCDevice/QueryDeviceListByUserId"
	RPCDevice_CreateDevice_FullMethodName            = "/device.RPCDevice/CreateDevice"
	RPCDevice_CreateDeviceAndConnect_FullMethodName  = "/device.RPCDevice/CreateDeviceAndConnect"
	RPCDevice_DeleteDevice_FullMethodName            = "/device.RPCDevice/DeleteDevice"
)

// RPCDeviceClient is the client API for RPCDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCDeviceClient interface {
	// 设备连接云服务 edge s  c
	ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error)
	// 获取驱动所有设备
	QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdResponse, error)
	// 获取用户所有设备
	QueryDeviceListByUserId(ctx context.Context, in *QueryDeviceListByUserIdRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error)
	// 创建设备
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error)
}

type rPCDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCDeviceClient(cc grpc.ClientConnInterface) RPCDeviceClient {
	return &rPCDeviceClient{cc}
}

func (c *rPCDeviceClient) ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error) {
	out := new(ConnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, RPCDevice_ConnectIotPlatform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error) {
	out := new(DisconnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, RPCDevice_DisconnectIotPlatform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error) {
	out := new(GetDeviceConnectStatusResponse)
	err := c.cc.Invoke(ctx, RPCDevice_GetDeviceConnectStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error) {
	out := new(QueryDeviceListResponse)
	err := c.cc.Invoke(ctx, RPCDevice_QueryDeviceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdResponse, error) {
	out := new(QueryDeviceByIdResponse)
	err := c.cc.Invoke(ctx, RPCDevice_QueryDeviceById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) QueryDeviceListByUserId(ctx context.Context, in *QueryDeviceListByUserIdRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error) {
	out := new(QueryDeviceListResponse)
	err := c.cc.Invoke(ctx, RPCDevice_QueryDeviceListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error) {
	out := new(CreateDeviceRequestResponse)
	err := c.cc.Invoke(ctx, RPCDevice_CreateDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error) {
	out := new(CreateDeviceAndConnectRequestResponse)
	err := c.cc.Invoke(ctx, RPCDevice_CreateDeviceAndConnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCDeviceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error) {
	out := new(DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, RPCDevice_DeleteDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCDeviceServer is the server API for RPCDevice service.
// All implementations must embed UnimplementedRPCDeviceServer
// for forward compatibility
type RPCDeviceServer interface {
	// 设备连接云服务 edge s  c
	ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error)
	// 获取驱动所有设备
	QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdResponse, error)
	// 获取用户所有设备
	QueryDeviceListByUserId(context.Context, *QueryDeviceListByUserIdRequest) (*QueryDeviceListResponse, error)
	// 创建设备
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	mustEmbedUnimplementedRPCDeviceServer()
}

// UnimplementedRPCDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedRPCDeviceServer struct {
}

func (UnimplementedRPCDeviceServer) ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectIotPlatform not implemented")
}
func (UnimplementedRPCDeviceServer) DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectIotPlatform not implemented")
}
func (UnimplementedRPCDeviceServer) GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceConnectStatus not implemented")
}
func (UnimplementedRPCDeviceServer) QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceList not implemented")
}
func (UnimplementedRPCDeviceServer) QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceById not implemented")
}
func (UnimplementedRPCDeviceServer) QueryDeviceListByUserId(context.Context, *QueryDeviceListByUserIdRequest) (*QueryDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceListByUserId not implemented")
}
func (UnimplementedRPCDeviceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedRPCDeviceServer) CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceAndConnect not implemented")
}
func (UnimplementedRPCDeviceServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedRPCDeviceServer) mustEmbedUnimplementedRPCDeviceServer() {}

// UnsafeRPCDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCDeviceServer will
// result in compilation errors.
type UnsafeRPCDeviceServer interface {
	mustEmbedUnimplementedRPCDeviceServer()
}

func RegisterRPCDeviceServer(s grpc.ServiceRegistrar, srv RPCDeviceServer) {
	s.RegisterService(&RPCDevice_ServiceDesc, srv)
}

func _RPCDevice_ConnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).ConnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_ConnectIotPlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).ConnectIotPlatform(ctx, req.(*ConnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_DisconnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).DisconnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_DisconnectIotPlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).DisconnectIotPlatform(ctx, req.(*DisconnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_GetDeviceConnectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceConnectStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).GetDeviceConnectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_GetDeviceConnectStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).GetDeviceConnectStatus(ctx, req.(*GetDeviceConnectStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_QueryDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).QueryDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_QueryDeviceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).QueryDeviceList(ctx, req.(*QueryDeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_QueryDeviceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).QueryDeviceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_QueryDeviceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).QueryDeviceById(ctx, req.(*QueryDeviceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_QueryDeviceListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).QueryDeviceListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_QueryDeviceListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).QueryDeviceListByUserId(ctx, req.(*QueryDeviceListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_CreateDeviceAndConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceAndConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).CreateDeviceAndConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_CreateDeviceAndConnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).CreateDeviceAndConnect(ctx, req.(*CreateDeviceAndConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCDevice_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDeviceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCDevice_DeleteDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDeviceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCDevice_ServiceDesc is the grpc.ServiceDesc for RPCDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.RPCDevice",
	HandlerType: (*RPCDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectIotPlatform",
			Handler:    _RPCDevice_ConnectIotPlatform_Handler,
		},
		{
			MethodName: "DisconnectIotPlatform",
			Handler:    _RPCDevice_DisconnectIotPlatform_Handler,
		},
		{
			MethodName: "GetDeviceConnectStatus",
			Handler:    _RPCDevice_GetDeviceConnectStatus_Handler,
		},
		{
			MethodName: "QueryDeviceList",
			Handler:    _RPCDevice_QueryDeviceList_Handler,
		},
		{
			MethodName: "QueryDeviceById",
			Handler:    _RPCDevice_QueryDeviceById_Handler,
		},
		{
			MethodName: "QueryDeviceListByUserId",
			Handler:    _RPCDevice_QueryDeviceListByUserId_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _RPCDevice_CreateDevice_Handler,
		},
		{
			MethodName: "CreateDeviceAndConnect",
			Handler:    _RPCDevice_CreateDeviceAndConnect_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _RPCDevice_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device/device.proto",
}
