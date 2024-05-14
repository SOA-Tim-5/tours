// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc2
// source: equipment/equipment.proto

package equipment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EquipmentServiceClient is the client API for EquipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EquipmentServiceClient interface {
	CreateEquipment(ctx context.Context, in *EquipmentCreateDto, opts ...grpc.CallOption) (*EquipmentResponseDto, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EquipmentListResponse, error)
}

type equipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquipmentServiceClient(cc grpc.ClientConnInterface) EquipmentServiceClient {
	return &equipmentServiceClient{cc}
}

func (c *equipmentServiceClient) CreateEquipment(ctx context.Context, in *EquipmentCreateDto, opts ...grpc.CallOption) (*EquipmentResponseDto, error) {
	out := new(EquipmentResponseDto)
	err := c.cc.Invoke(ctx, "/EquipmentService/CreateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EquipmentListResponse, error) {
	out := new(EquipmentListResponse)
	err := c.cc.Invoke(ctx, "/EquipmentService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquipmentServiceServer is the server API for EquipmentService service.
// All implementations must embed UnimplementedEquipmentServiceServer
// for forward compatibility
type EquipmentServiceServer interface {
	CreateEquipment(context.Context, *EquipmentCreateDto) (*EquipmentResponseDto, error)
	GetAll(context.Context, *emptypb.Empty) (*EquipmentListResponse, error)
	mustEmbedUnimplementedEquipmentServiceServer()
}

// UnimplementedEquipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEquipmentServiceServer struct {
}

func (UnimplementedEquipmentServiceServer) CreateEquipment(context.Context, *EquipmentCreateDto) (*EquipmentResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) GetAll(context.Context, *emptypb.Empty) (*EquipmentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEquipmentServiceServer) mustEmbedUnimplementedEquipmentServiceServer() {}

// UnsafeEquipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EquipmentServiceServer will
// result in compilation errors.
type UnsafeEquipmentServiceServer interface {
	mustEmbedUnimplementedEquipmentServiceServer()
}

func RegisterEquipmentServiceServer(s grpc.ServiceRegistrar, srv EquipmentServiceServer) {
	s.RegisterService(&EquipmentService_ServiceDesc, srv)
}

func _EquipmentService_CreateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).CreateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EquipmentService/CreateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).CreateEquipment(ctx, req.(*EquipmentCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EquipmentService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// EquipmentService_ServiceDesc is the grpc.ServiceDesc for EquipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EquipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EquipmentService",
	HandlerType: (*EquipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEquipment",
			Handler:    _EquipmentService_CreateEquipment_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EquipmentService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "equipment/equipment.proto",
}
