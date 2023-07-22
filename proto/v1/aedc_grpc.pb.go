// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/aedc.proto

package v1

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

// AEDCServiceClient is the client API for AEDCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AEDCServiceClient interface {
	VerifyInfo(ctx context.Context, in *VerifyInfoReq, opts ...grpc.CallOption) (*VerifyInfoResp, error)
}

type aEDCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAEDCServiceClient(cc grpc.ClientConnInterface) AEDCServiceClient {
	return &aEDCServiceClient{cc}
}

func (c *aEDCServiceClient) VerifyInfo(ctx context.Context, in *VerifyInfoReq, opts ...grpc.CallOption) (*VerifyInfoResp, error) {
	out := new(VerifyInfoResp)
	err := c.cc.Invoke(ctx, "/proto.v1.AEDCService/VerifyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AEDCServiceServer is the server API for AEDCService service.
// All implementations must embed UnimplementedAEDCServiceServer
// for forward compatibility
type AEDCServiceServer interface {
	VerifyInfo(context.Context, *VerifyInfoReq) (*VerifyInfoResp, error)
	mustEmbedUnimplementedAEDCServiceServer()
}

// UnimplementedAEDCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAEDCServiceServer struct {
}

func (UnimplementedAEDCServiceServer) VerifyInfo(context.Context, *VerifyInfoReq) (*VerifyInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyInfo not implemented")
}
func (UnimplementedAEDCServiceServer) mustEmbedUnimplementedAEDCServiceServer() {}

// UnsafeAEDCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AEDCServiceServer will
// result in compilation errors.
type UnsafeAEDCServiceServer interface {
	mustEmbedUnimplementedAEDCServiceServer()
}

func RegisterAEDCServiceServer(s grpc.ServiceRegistrar, srv AEDCServiceServer) {
	s.RegisterService(&AEDCService_ServiceDesc, srv)
}

func _AEDCService_VerifyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AEDCServiceServer).VerifyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.AEDCService/VerifyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AEDCServiceServer).VerifyInfo(ctx, req.(*VerifyInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AEDCService_ServiceDesc is the grpc.ServiceDesc for AEDCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AEDCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.AEDCService",
	HandlerType: (*AEDCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyInfo",
			Handler:    _AEDCService_VerifyInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/aedc.proto",
}