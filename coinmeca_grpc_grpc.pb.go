// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: coinmeca_grpc.proto

package __

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
	CoinmecaGrpcModule_Send_FullMethodName = "/grpcmodule.CoinmecaGrpcModule/Send"
)

// CoinmecaGrpcModuleClient is the client API for CoinmecaGrpcModule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinmecaGrpcModuleClient interface {
	Send(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type coinmecaGrpcModuleClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinmecaGrpcModuleClient(cc grpc.ClientConnInterface) CoinmecaGrpcModuleClient {
	return &coinmecaGrpcModuleClient{cc}
}

func (c *coinmecaGrpcModuleClient) Send(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, CoinmecaGrpcModule_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinmecaGrpcModuleServer is the server API for CoinmecaGrpcModule service.
// All implementations must embed UnimplementedCoinmecaGrpcModuleServer
// for forward compatibility
type CoinmecaGrpcModuleServer interface {
	Send(context.Context, *GeneralRequest) (*GeneralResponse, error)
	mustEmbedUnimplementedCoinmecaGrpcModuleServer()
}

// UnimplementedCoinmecaGrpcModuleServer must be embedded to have forward compatible implementations.
type UnimplementedCoinmecaGrpcModuleServer struct {
}

func (UnimplementedCoinmecaGrpcModuleServer) Send(context.Context, *GeneralRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedCoinmecaGrpcModuleServer) mustEmbedUnimplementedCoinmecaGrpcModuleServer() {}

// UnsafeCoinmecaGrpcModuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinmecaGrpcModuleServer will
// result in compilation errors.
type UnsafeCoinmecaGrpcModuleServer interface {
	mustEmbedUnimplementedCoinmecaGrpcModuleServer()
}

func RegisterCoinmecaGrpcModuleServer(s grpc.ServiceRegistrar, srv CoinmecaGrpcModuleServer) {
	s.RegisterService(&CoinmecaGrpcModule_ServiceDesc, srv)
}

func _CoinmecaGrpcModule_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinmecaGrpcModuleServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoinmecaGrpcModule_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinmecaGrpcModuleServer).Send(ctx, req.(*GeneralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoinmecaGrpcModule_ServiceDesc is the grpc.ServiceDesc for CoinmecaGrpcModule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoinmecaGrpcModule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcmodule.CoinmecaGrpcModule",
	HandlerType: (*CoinmecaGrpcModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _CoinmecaGrpcModule_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinmeca_grpc.proto",
}
