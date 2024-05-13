// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc2
// source: preference/preference.proto

package preference

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

// PreferenceServiceClient is the client API for PreferenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreferenceServiceClient interface {
	Create(ctx context.Context, in *PreferenceCreateDto, opts ...grpc.CallOption) (*PreferenceResponseDto, error)
	GetByTouristId(ctx context.Context, in *GetPreferenceParams, opts ...grpc.CallOption) (*PreferenceResponseDto, error)
}

type preferenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPreferenceServiceClient(cc grpc.ClientConnInterface) PreferenceServiceClient {
	return &preferenceServiceClient{cc}
}

func (c *preferenceServiceClient) Create(ctx context.Context, in *PreferenceCreateDto, opts ...grpc.CallOption) (*PreferenceResponseDto, error) {
	out := new(PreferenceResponseDto)
	err := c.cc.Invoke(ctx, "/PreferenceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preferenceServiceClient) GetByTouristId(ctx context.Context, in *GetPreferenceParams, opts ...grpc.CallOption) (*PreferenceResponseDto, error) {
	out := new(PreferenceResponseDto)
	err := c.cc.Invoke(ctx, "/PreferenceService/GetByTouristId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreferenceServiceServer is the server API for PreferenceService service.
// All implementations must embed UnimplementedPreferenceServiceServer
// for forward compatibility
type PreferenceServiceServer interface {
	Create(context.Context, *PreferenceCreateDto) (*PreferenceResponseDto, error)
	GetByTouristId(context.Context, *GetPreferenceParams) (*PreferenceResponseDto, error)
	mustEmbedUnimplementedPreferenceServiceServer()
}

// UnimplementedPreferenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPreferenceServiceServer struct {
}

func (UnimplementedPreferenceServiceServer) Create(context.Context, *PreferenceCreateDto) (*PreferenceResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPreferenceServiceServer) GetByTouristId(context.Context, *GetPreferenceParams) (*PreferenceResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByTouristId not implemented")
}
func (UnimplementedPreferenceServiceServer) mustEmbedUnimplementedPreferenceServiceServer() {}

// UnsafePreferenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreferenceServiceServer will
// result in compilation errors.
type UnsafePreferenceServiceServer interface {
	mustEmbedUnimplementedPreferenceServiceServer()
}

func RegisterPreferenceServiceServer(s grpc.ServiceRegistrar, srv PreferenceServiceServer) {
	s.RegisterService(&PreferenceService_ServiceDesc, srv)
}

func _PreferenceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreferenceCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferenceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PreferenceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferenceServiceServer).Create(ctx, req.(*PreferenceCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _PreferenceService_GetByTouristId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreferenceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferenceServiceServer).GetByTouristId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PreferenceService/GetByTouristId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferenceServiceServer).GetByTouristId(ctx, req.(*GetPreferenceParams))
	}
	return interceptor(ctx, in, info, handler)
}

// PreferenceService_ServiceDesc is the grpc.ServiceDesc for PreferenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PreferenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PreferenceService",
	HandlerType: (*PreferenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PreferenceService_Create_Handler,
		},
		{
			MethodName: "GetByTouristId",
			Handler:    _PreferenceService_GetByTouristId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preference/preference.proto",
}
