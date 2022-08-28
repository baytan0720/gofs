// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: namenode.proto

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

// NameNodeServiceClient is the client API for NameNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeServiceClient interface {
	BlockReport(ctx context.Context, in *BlockReportArgs, opts ...grpc.CallOption) (*BlockReportReply, error)
	Delete(ctx context.Context, in *DeleteArgs, opts ...grpc.CallOption) (*DeleteReply, error)
	GetFile(ctx context.Context, in *GetFileArgs, opts ...grpc.CallOption) (*GetFileReply, error)
	GetLease(ctx context.Context, in *GetLeaseArgs, opts ...grpc.CallOption) (*GetLeaseReply, error)
	GetSystemInfo(ctx context.Context, in *GetSystemInfoArgs, opts ...grpc.CallOption) (*GetSystemInfoReply, error)
	HeartBeat(ctx context.Context, in *HeartBeatArgs, opts ...grpc.CallOption) (*HeartBeatReply, error)
	List(ctx context.Context, in *ListArgs, opts ...grpc.CallOption) (*ListReply, error)
	Mkdir(ctx context.Context, in *MkdirArgs, opts ...grpc.CallOption) (*MkdirReply, error)
	NewBlockReport(ctx context.Context, in *NewBlockReportArgs, opts ...grpc.CallOption) (*NewBlockReportReply, error)
	PutFile(ctx context.Context, in *PutFileArgs, opts ...grpc.CallOption) (*PutFileReply, error)
	PutBlock(ctx context.Context, in *PutBlockArgs, opts ...grpc.CallOption) (*PutBlockReply, error)
	Register(ctx context.Context, in *RegisterArgs, opts ...grpc.CallOption) (*RegisterReply, error)
	Rename(ctx context.Context, in *RenameArgs, opts ...grpc.CallOption) (*RenameReply, error)
	RenewLease(ctx context.Context, in *RenewLeaseArgs, opts ...grpc.CallOption) (*RenewLeaseReply, error)
	ReleaseLease(ctx context.Context, in *ReleaseLeaseArgs, opts ...grpc.CallOption) (*ReleaseLeaseReply, error)
	Stat(ctx context.Context, in *StatArgs, opts ...grpc.CallOption) (*StatReply, error)
}

type nameNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeServiceClient(cc grpc.ClientConnInterface) NameNodeServiceClient {
	return &nameNodeServiceClient{cc}
}

func (c *nameNodeServiceClient) BlockReport(ctx context.Context, in *BlockReportArgs, opts ...grpc.CallOption) (*BlockReportReply, error) {
	out := new(BlockReportReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/BlockReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Delete(ctx context.Context, in *DeleteArgs, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) GetFile(ctx context.Context, in *GetFileArgs, opts ...grpc.CallOption) (*GetFileReply, error) {
	out := new(GetFileReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) GetLease(ctx context.Context, in *GetLeaseArgs, opts ...grpc.CallOption) (*GetLeaseReply, error) {
	out := new(GetLeaseReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/GetLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) GetSystemInfo(ctx context.Context, in *GetSystemInfoArgs, opts ...grpc.CallOption) (*GetSystemInfoReply, error) {
	out := new(GetSystemInfoReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/GetSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) HeartBeat(ctx context.Context, in *HeartBeatArgs, opts ...grpc.CallOption) (*HeartBeatReply, error) {
	out := new(HeartBeatReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) List(ctx context.Context, in *ListArgs, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Mkdir(ctx context.Context, in *MkdirArgs, opts ...grpc.CallOption) (*MkdirReply, error) {
	out := new(MkdirReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) NewBlockReport(ctx context.Context, in *NewBlockReportArgs, opts ...grpc.CallOption) (*NewBlockReportReply, error) {
	out := new(NewBlockReportReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/NewBlockReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) PutFile(ctx context.Context, in *PutFileArgs, opts ...grpc.CallOption) (*PutFileReply, error) {
	out := new(PutFileReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/PutFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) PutBlock(ctx context.Context, in *PutBlockArgs, opts ...grpc.CallOption) (*PutBlockReply, error) {
	out := new(PutBlockReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/PutBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Register(ctx context.Context, in *RegisterArgs, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Rename(ctx context.Context, in *RenameArgs, opts ...grpc.CallOption) (*RenameReply, error) {
	out := new(RenameReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) RenewLease(ctx context.Context, in *RenewLeaseArgs, opts ...grpc.CallOption) (*RenewLeaseReply, error) {
	out := new(RenewLeaseReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/RenewLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) ReleaseLease(ctx context.Context, in *ReleaseLeaseArgs, opts ...grpc.CallOption) (*ReleaseLeaseReply, error) {
	out := new(ReleaseLeaseReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/ReleaseLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeServiceClient) Stat(ctx context.Context, in *StatArgs, opts ...grpc.CallOption) (*StatReply, error) {
	out := new(StatReply)
	err := c.cc.Invoke(ctx, "/service.NameNodeService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServiceServer is the server API for NameNodeService service.
// All implementations must embed UnimplementedNameNodeServiceServer
// for forward compatibility
type NameNodeServiceServer interface {
	BlockReport(context.Context, *BlockReportArgs) (*BlockReportReply, error)
	Delete(context.Context, *DeleteArgs) (*DeleteReply, error)
	GetFile(context.Context, *GetFileArgs) (*GetFileReply, error)
	GetLease(context.Context, *GetLeaseArgs) (*GetLeaseReply, error)
	GetSystemInfo(context.Context, *GetSystemInfoArgs) (*GetSystemInfoReply, error)
	HeartBeat(context.Context, *HeartBeatArgs) (*HeartBeatReply, error)
	List(context.Context, *ListArgs) (*ListReply, error)
	Mkdir(context.Context, *MkdirArgs) (*MkdirReply, error)
	NewBlockReport(context.Context, *NewBlockReportArgs) (*NewBlockReportReply, error)
	PutFile(context.Context, *PutFileArgs) (*PutFileReply, error)
	PutBlock(context.Context, *PutBlockArgs) (*PutBlockReply, error)
	Register(context.Context, *RegisterArgs) (*RegisterReply, error)
	Rename(context.Context, *RenameArgs) (*RenameReply, error)
	RenewLease(context.Context, *RenewLeaseArgs) (*RenewLeaseReply, error)
	ReleaseLease(context.Context, *ReleaseLeaseArgs) (*ReleaseLeaseReply, error)
	Stat(context.Context, *StatArgs) (*StatReply, error)
	mustEmbedUnimplementedNameNodeServiceServer()
}

// UnimplementedNameNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServiceServer struct {
}

func (UnimplementedNameNodeServiceServer) BlockReport(context.Context, *BlockReportArgs) (*BlockReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockReport not implemented")
}
func (UnimplementedNameNodeServiceServer) Delete(context.Context, *DeleteArgs) (*DeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNameNodeServiceServer) GetFile(context.Context, *GetFileArgs) (*GetFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedNameNodeServiceServer) GetLease(context.Context, *GetLeaseArgs) (*GetLeaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLease not implemented")
}
func (UnimplementedNameNodeServiceServer) GetSystemInfo(context.Context, *GetSystemInfoArgs) (*GetSystemInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedNameNodeServiceServer) HeartBeat(context.Context, *HeartBeatArgs) (*HeartBeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedNameNodeServiceServer) List(context.Context, *ListArgs) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNameNodeServiceServer) Mkdir(context.Context, *MkdirArgs) (*MkdirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (UnimplementedNameNodeServiceServer) NewBlockReport(context.Context, *NewBlockReportArgs) (*NewBlockReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBlockReport not implemented")
}
func (UnimplementedNameNodeServiceServer) PutFile(context.Context, *PutFileArgs) (*PutFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutFile not implemented")
}
func (UnimplementedNameNodeServiceServer) PutBlock(context.Context, *PutBlockArgs) (*PutBlockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBlock not implemented")
}
func (UnimplementedNameNodeServiceServer) Register(context.Context, *RegisterArgs) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedNameNodeServiceServer) Rename(context.Context, *RenameArgs) (*RenameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedNameNodeServiceServer) RenewLease(context.Context, *RenewLeaseArgs) (*RenewLeaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewLease not implemented")
}
func (UnimplementedNameNodeServiceServer) ReleaseLease(context.Context, *ReleaseLeaseArgs) (*ReleaseLeaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseLease not implemented")
}
func (UnimplementedNameNodeServiceServer) Stat(context.Context, *StatArgs) (*StatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedNameNodeServiceServer) mustEmbedUnimplementedNameNodeServiceServer() {}

// UnsafeNameNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServiceServer will
// result in compilation errors.
type UnsafeNameNodeServiceServer interface {
	mustEmbedUnimplementedNameNodeServiceServer()
}

func RegisterNameNodeServiceServer(s grpc.ServiceRegistrar, srv NameNodeServiceServer) {
	s.RegisterService(&NameNodeService_ServiceDesc, srv)
}

func _NameNodeService_BlockReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockReportArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).BlockReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/BlockReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).BlockReport(ctx, req.(*BlockReportArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Delete(ctx, req.(*DeleteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).GetFile(ctx, req.(*GetFileArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_GetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaseArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).GetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/GetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).GetLease(ctx, req.(*GetLeaseArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/GetSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).GetSystemInfo(ctx, req.(*GetSystemInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).HeartBeat(ctx, req.(*HeartBeatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).List(ctx, req.(*ListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Mkdir(ctx, req.(*MkdirArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_NewBlockReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBlockReportArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).NewBlockReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/NewBlockReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).NewBlockReport(ctx, req.(*NewBlockReportArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_PutFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutFileArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).PutFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/PutFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).PutFile(ctx, req.(*PutFileArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_PutBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutBlockArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).PutBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/PutBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).PutBlock(ctx, req.(*PutBlockArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Register(ctx, req.(*RegisterArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Rename(ctx, req.(*RenameArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_RenewLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewLeaseArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).RenewLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/RenewLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).RenewLease(ctx, req.(*RenewLeaseArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_ReleaseLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseLeaseArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).ReleaseLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/ReleaseLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).ReleaseLease(ctx, req.(*ReleaseLeaseArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNodeService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NameNodeService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServiceServer).Stat(ctx, req.(*StatArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// NameNodeService_ServiceDesc is the grpc.ServiceDesc for NameNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.NameNodeService",
	HandlerType: (*NameNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockReport",
			Handler:    _NameNodeService_BlockReport_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NameNodeService_Delete_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _NameNodeService_GetFile_Handler,
		},
		{
			MethodName: "GetLease",
			Handler:    _NameNodeService_GetLease_Handler,
		},
		{
			MethodName: "GetSystemInfo",
			Handler:    _NameNodeService_GetSystemInfo_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _NameNodeService_HeartBeat_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NameNodeService_List_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _NameNodeService_Mkdir_Handler,
		},
		{
			MethodName: "NewBlockReport",
			Handler:    _NameNodeService_NewBlockReport_Handler,
		},
		{
			MethodName: "PutFile",
			Handler:    _NameNodeService_PutFile_Handler,
		},
		{
			MethodName: "PutBlock",
			Handler:    _NameNodeService_PutBlock_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _NameNodeService_Register_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _NameNodeService_Rename_Handler,
		},
		{
			MethodName: "RenewLease",
			Handler:    _NameNodeService_RenewLease_Handler,
		},
		{
			MethodName: "ReleaseLease",
			Handler:    _NameNodeService_ReleaseLease_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _NameNodeService_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namenode.proto",
}
