// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: address.proto

package address

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
	Address_GetProvice_FullMethodName         = "/address.Address/GetProvice"
	Address_GetDistrict_FullMethodName        = "/address.Address/GetDistrict"
	Address_GetWard_FullMethodName            = "/address.Address/GetWard"
	Address_GetListAddressUser_FullMethodName = "/address.Address/GetListAddressUser"
	Address_CreateAddressUser_FullMethodName  = "/address.Address/CreateAddressUser"
	Address_UpdateAddressUser_FullMethodName  = "/address.Address/UpdateAddressUser"
)

// AddressClient is the client API for Address service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressClient interface {
	GetProvice(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseProvince, error)
	GetDistrict(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseDistrict, error)
	GetWard(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseWard, error)
	GetListAddressUser(ctx context.Context, in *RequestListAddress, opts ...grpc.CallOption) (*ResponseListAddress, error)
	CreateAddressUser(ctx context.Context, in *AddressCreateRequest, opts ...grpc.CallOption) (*AddressUser, error)
	UpdateAddressUser(ctx context.Context, in *AddressUser, opts ...grpc.CallOption) (*AddressUser, error)
}

type addressClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressClient(cc grpc.ClientConnInterface) AddressClient {
	return &addressClient{cc}
}

func (c *addressClient) GetProvice(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseProvince, error) {
	out := new(ResponseProvince)
	err := c.cc.Invoke(ctx, Address_GetProvice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) GetDistrict(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseDistrict, error) {
	out := new(ResponseDistrict)
	err := c.cc.Invoke(ctx, Address_GetDistrict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) GetWard(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ResponseWard, error) {
	out := new(ResponseWard)
	err := c.cc.Invoke(ctx, Address_GetWard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) GetListAddressUser(ctx context.Context, in *RequestListAddress, opts ...grpc.CallOption) (*ResponseListAddress, error) {
	out := new(ResponseListAddress)
	err := c.cc.Invoke(ctx, Address_GetListAddressUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) CreateAddressUser(ctx context.Context, in *AddressCreateRequest, opts ...grpc.CallOption) (*AddressUser, error) {
	out := new(AddressUser)
	err := c.cc.Invoke(ctx, Address_CreateAddressUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressClient) UpdateAddressUser(ctx context.Context, in *AddressUser, opts ...grpc.CallOption) (*AddressUser, error) {
	out := new(AddressUser)
	err := c.cc.Invoke(ctx, Address_UpdateAddressUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServer is the server API for Address service.
// All implementations must embed UnimplementedAddressServer
// for forward compatibility
type AddressServer interface {
	GetProvice(context.Context, *Request) (*ResponseProvince, error)
	GetDistrict(context.Context, *Request) (*ResponseDistrict, error)
	GetWard(context.Context, *Request) (*ResponseWard, error)
	GetListAddressUser(context.Context, *RequestListAddress) (*ResponseListAddress, error)
	CreateAddressUser(context.Context, *AddressCreateRequest) (*AddressUser, error)
	UpdateAddressUser(context.Context, *AddressUser) (*AddressUser, error)
	mustEmbedUnimplementedAddressServer()
}

// UnimplementedAddressServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServer struct {
}

func (UnimplementedAddressServer) GetProvice(context.Context, *Request) (*ResponseProvince, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvice not implemented")
}
func (UnimplementedAddressServer) GetDistrict(context.Context, *Request) (*ResponseDistrict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistrict not implemented")
}
func (UnimplementedAddressServer) GetWard(context.Context, *Request) (*ResponseWard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWard not implemented")
}
func (UnimplementedAddressServer) GetListAddressUser(context.Context, *RequestListAddress) (*ResponseListAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListAddressUser not implemented")
}
func (UnimplementedAddressServer) CreateAddressUser(context.Context, *AddressCreateRequest) (*AddressUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddressUser not implemented")
}
func (UnimplementedAddressServer) UpdateAddressUser(context.Context, *AddressUser) (*AddressUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressUser not implemented")
}
func (UnimplementedAddressServer) mustEmbedUnimplementedAddressServer() {}

// UnsafeAddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServer will
// result in compilation errors.
type UnsafeAddressServer interface {
	mustEmbedUnimplementedAddressServer()
}

func RegisterAddressServer(s grpc.ServiceRegistrar, srv AddressServer) {
	s.RegisterService(&Address_ServiceDesc, srv)
}

func _Address_GetProvice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetProvice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_GetProvice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetProvice(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_GetDistrict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetDistrict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_GetDistrict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetDistrict(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_GetWard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetWard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_GetWard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetWard(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_GetListAddressUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).GetListAddressUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_GetListAddressUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).GetListAddressUser(ctx, req.(*RequestListAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_CreateAddressUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).CreateAddressUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_CreateAddressUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).CreateAddressUser(ctx, req.(*AddressCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Address_UpdateAddressUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServer).UpdateAddressUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Address_UpdateAddressUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServer).UpdateAddressUser(ctx, req.(*AddressUser))
	}
	return interceptor(ctx, in, info, handler)
}

// Address_ServiceDesc is the grpc.ServiceDesc for Address service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Address_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "address.Address",
	HandlerType: (*AddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProvice",
			Handler:    _Address_GetProvice_Handler,
		},
		{
			MethodName: "GetDistrict",
			Handler:    _Address_GetDistrict_Handler,
		},
		{
			MethodName: "GetWard",
			Handler:    _Address_GetWard_Handler,
		},
		{
			MethodName: "GetListAddressUser",
			Handler:    _Address_GetListAddressUser_Handler,
		},
		{
			MethodName: "CreateAddressUser",
			Handler:    _Address_CreateAddressUser_Handler,
		},
		{
			MethodName: "UpdateAddressUser",
			Handler:    _Address_UpdateAddressUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
