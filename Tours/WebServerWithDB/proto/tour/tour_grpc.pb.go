// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: tour/tour.proto

package tour

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

// TourServiceClient is the client API for TourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TourServiceClient interface {
	Create(ctx context.Context, in *TourCreateDto, opts ...grpc.CallOption) (*TourResponseDto, error)
	GetAuthorsTours(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*TourListResponse, error)
	CreateKeyPoint(ctx context.Context, in *KeyPointCreateDto, opts ...grpc.CallOption) (*KeyPointResponseDto, error)
	GetById(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*TourResponseDto, error)
	GetKeyPoints(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*KeyPointList, error)
}

type tourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTourServiceClient(cc grpc.ClientConnInterface) TourServiceClient {
	return &tourServiceClient{cc}
}

func (c *tourServiceClient) Create(ctx context.Context, in *TourCreateDto, opts ...grpc.CallOption) (*TourResponseDto, error) {
	out := new(TourResponseDto)
	err := c.cc.Invoke(ctx, "/TourService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) GetAuthorsTours(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*TourListResponse, error) {
	out := new(TourListResponse)
	err := c.cc.Invoke(ctx, "/TourService/GetAuthorsTours", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) CreateKeyPoint(ctx context.Context, in *KeyPointCreateDto, opts ...grpc.CallOption) (*KeyPointResponseDto, error) {
	out := new(KeyPointResponseDto)
	err := c.cc.Invoke(ctx, "/TourService/CreateKeyPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) GetById(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*TourResponseDto, error) {
	out := new(TourResponseDto)
	err := c.cc.Invoke(ctx, "/TourService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) GetKeyPoints(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*KeyPointList, error) {
	out := new(KeyPointList)
	err := c.cc.Invoke(ctx, "/TourService/GetKeyPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TourServiceServer is the server API for TourService service.
// All implementations must embed UnimplementedTourServiceServer
// for forward compatibility
type TourServiceServer interface {
	Create(context.Context, *TourCreateDto) (*TourResponseDto, error)
	GetAuthorsTours(context.Context, *GetParams) (*TourListResponse, error)
	CreateKeyPoint(context.Context, *KeyPointCreateDto) (*KeyPointResponseDto, error)
	GetById(context.Context, *GetParams) (*TourResponseDto, error)
	GetKeyPoints(context.Context, *GetParams) (*KeyPointList, error)
	mustEmbedUnimplementedTourServiceServer()
}

// UnimplementedTourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTourServiceServer struct {
}

func (UnimplementedTourServiceServer) Create(context.Context, *TourCreateDto) (*TourResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTourServiceServer) GetAuthorsTours(context.Context, *GetParams) (*TourListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorsTours not implemented")
}
func (UnimplementedTourServiceServer) CreateKeyPoint(context.Context, *KeyPointCreateDto) (*KeyPointResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeyPoint not implemented")
}
func (UnimplementedTourServiceServer) GetById(context.Context, *GetParams) (*TourResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTourServiceServer) GetKeyPoints(context.Context, *GetParams) (*KeyPointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyPoints not implemented")
}
func (UnimplementedTourServiceServer) mustEmbedUnimplementedTourServiceServer() {}

// UnsafeTourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TourServiceServer will
// result in compilation errors.
type UnsafeTourServiceServer interface {
	mustEmbedUnimplementedTourServiceServer()
}

func RegisterTourServiceServer(s grpc.ServiceRegistrar, srv TourServiceServer) {
	s.RegisterService(&TourService_ServiceDesc, srv)
}

func _TourService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TourCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TourService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).Create(ctx, req.(*TourCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_GetAuthorsTours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).GetAuthorsTours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TourService/GetAuthorsTours",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).GetAuthorsTours(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_CreateKeyPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPointCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).CreateKeyPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TourService/CreateKeyPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).CreateKeyPoint(ctx, req.(*KeyPointCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TourService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).GetById(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_GetKeyPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).GetKeyPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TourService/GetKeyPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).GetKeyPoints(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

// TourService_ServiceDesc is the grpc.ServiceDesc for TourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TourService",
	HandlerType: (*TourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TourService_Create_Handler,
		},
		{
			MethodName: "GetAuthorsTours",
			Handler:    _TourService_GetAuthorsTours_Handler,
		},
		{
			MethodName: "CreateKeyPoint",
			Handler:    _TourService_CreateKeyPoint_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _TourService_GetById_Handler,
		},
		{
			MethodName: "GetKeyPoints",
			Handler:    _TourService_GetKeyPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tour/tour.proto",
}
