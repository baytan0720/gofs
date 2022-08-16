// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: PipelineToDN.proto

package service

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

// PipelineToDNServiceClient is the client API for PipelineToDNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineToDNServiceClient interface {
	PipelineToDN(ctx context.Context, in *PipelineToDNArgs, opts ...grpc.CallOption) (*PipelineToDNReply, error)
}

type pipelineToDNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineToDNServiceClient(cc grpc.ClientConnInterface) PipelineToDNServiceClient {
	return &pipelineToDNServiceClient{cc}
}

func (c *pipelineToDNServiceClient) PipelineToDN(ctx context.Context, in *PipelineToDNArgs, opts ...grpc.CallOption) (*PipelineToDNReply, error) {
	out := new(PipelineToDNReply)
	err := c.cc.Invoke(ctx, "/pb.PipelineToDNService/PipelineToDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineToDNServiceServer is the server API for PipelineToDNService service.
// All implementations must embed UnimplementedPipelineToDNServiceServer
// for forward compatibility
type PipelineToDNServiceServer interface {
	PipelineToDN(context.Context, *PipelineToDNArgs) (*PipelineToDNReply, error)
	mustEmbedUnimplementedPipelineToDNServiceServer()
}

// UnimplementedPipelineToDNServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineToDNServiceServer struct {
}

func (UnimplementedPipelineToDNServiceServer) PipelineToDN(context.Context, *PipelineToDNArgs) (*PipelineToDNReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineToDN not implemented")
}
func (UnimplementedPipelineToDNServiceServer) mustEmbedUnimplementedPipelineToDNServiceServer() {}

// UnsafePipelineToDNServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineToDNServiceServer will
// result in compilation errors.
type UnsafePipelineToDNServiceServer interface {
	mustEmbedUnimplementedPipelineToDNServiceServer()
}

func RegisterPipelineToDNServiceServer(s grpc.ServiceRegistrar, srv PipelineToDNServiceServer) {
	s.RegisterService(&PipelineToDNService_ServiceDesc, srv)
}

func _PipelineToDNService_PipelineToDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineToDNArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineToDNServiceServer).PipelineToDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PipelineToDNService/PipelineToDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineToDNServiceServer).PipelineToDN(ctx, req.(*PipelineToDNArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineToDNService_ServiceDesc is the grpc.ServiceDesc for PipelineToDNService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineToDNService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PipelineToDNService",
	HandlerType: (*PipelineToDNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PipelineToDN",
			Handler:    _PipelineToDNService_PipelineToDN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "PipelineToDN.proto",
}