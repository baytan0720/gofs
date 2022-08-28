// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: datanode.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WriteBlockArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId     string  `protobuf:"bytes,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
	Md5         string  `protobuf:"bytes,2,opt,name=Md5,proto3" json:"Md5,omitempty"`
	DataBuf     []byte  `protobuf:"bytes,3,opt,name=DataBuf,proto3" json:"DataBuf,omitempty"`
	Size        int64   `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	EntryId     int64   `protobuf:"varint,5,opt,name=EntryId,proto3" json:"EntryId,omitempty"`
	Index       int32   `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	DatanodeIds []int32 `protobuf:"varint,7,rep,packed,name=DatanodeIds,proto3" json:"DatanodeIds,omitempty"`
}

func (x *WriteBlockArgs) Reset() {
	*x = WriteBlockArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBlockArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBlockArgs) ProtoMessage() {}

func (x *WriteBlockArgs) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBlockArgs.ProtoReflect.Descriptor instead.
func (*WriteBlockArgs) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{0}
}

func (x *WriteBlockArgs) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *WriteBlockArgs) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *WriteBlockArgs) GetDataBuf() []byte {
	if x != nil {
		return x.DataBuf
	}
	return nil
}

func (x *WriteBlockArgs) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WriteBlockArgs) GetEntryId() int64 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *WriteBlockArgs) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *WriteBlockArgs) GetDatanodeIds() []int32 {
	if x != nil {
		return x.DatanodeIds
	}
	return nil
}

type WriteBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusCode `protobuf:"varint,1,opt,name=Status,proto3,enum=service.StatusCode" json:"Status,omitempty"`
}

func (x *WriteBlockReply) Reset() {
	*x = WriteBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBlockReply) ProtoMessage() {}

func (x *WriteBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBlockReply.ProtoReflect.Descriptor instead.
func (*WriteBlockReply) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{1}
}

func (x *WriteBlockReply) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

type ReadBlockArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *ReadBlockArgs) Reset() {
	*x = ReadBlockArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockArgs) ProtoMessage() {}

func (x *ReadBlockArgs) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockArgs.ProtoReflect.Descriptor instead.
func (*ReadBlockArgs) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{2}
}

func (x *ReadBlockArgs) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type ReadBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  StatusCode `protobuf:"varint,1,opt,name=Status,proto3,enum=service.StatusCode" json:"Status,omitempty"`
	Md5     string     `protobuf:"bytes,2,opt,name=Md5,proto3" json:"Md5,omitempty"`
	DataBuf []byte     `protobuf:"bytes,3,opt,name=DataBuf,proto3" json:"DataBuf,omitempty"`
}

func (x *ReadBlockReply) Reset() {
	*x = ReadBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlockReply) ProtoMessage() {}

func (x *ReadBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlockReply.ProtoReflect.Descriptor instead.
func (*ReadBlockReply) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{3}
}

func (x *ReadBlockReply) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

func (x *ReadBlockReply) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *ReadBlockReply) GetDataBuf() []byte {
	if x != nil {
		return x.DataBuf
	}
	return nil
}

type GetBlockInfoArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId int64 `protobuf:"varint,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *GetBlockInfoArgs) Reset() {
	*x = GetBlockInfoArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockInfoArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockInfoArgs) ProtoMessage() {}

func (x *GetBlockInfoArgs) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockInfoArgs.ProtoReflect.Descriptor instead.
func (*GetBlockInfoArgs) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlockInfoArgs) GetBlockId() int64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

type GetBlockInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusCode `protobuf:"varint,1,opt,name=Status,proto3,enum=service.StatusCode" json:"Status,omitempty"`
	Info   *BlockInfo `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetBlockInfoReply) Reset() {
	*x = GetBlockInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockInfoReply) ProtoMessage() {}

func (x *GetBlockInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockInfoReply.ProtoReflect.Descriptor instead.
func (*GetBlockInfoReply) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlockInfoReply) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

func (x *GetBlockInfoReply) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreatePipelineArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     int32              `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	DataNodes []*DataNodeNetInfo `protobuf:"bytes,2,rep,name=DataNodes,proto3" json:"DataNodes,omitempty"`
}

func (x *CreatePipelineArgs) Reset() {
	*x = CreatePipelineArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineArgs) ProtoMessage() {}

func (x *CreatePipelineArgs) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineArgs.ProtoReflect.Descriptor instead.
func (*CreatePipelineArgs) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePipelineArgs) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CreatePipelineArgs) GetDataNodes() []*DataNodeNetInfo {
	if x != nil {
		return x.DataNodes
	}
	return nil
}

type CreatePipelineReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusCode `protobuf:"varint,1,opt,name=Status,proto3,enum=service.StatusCode" json:"Status,omitempty"`
}

func (x *CreatePipelineReply) Reset() {
	*x = CreatePipelineReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineReply) ProtoMessage() {}

func (x *CreatePipelineReply) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineReply.ProtoReflect.Descriptor instead.
func (*CreatePipelineReply) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePipelineReply) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_OK
}

var File_datanode_proto protoreflect.FileDescriptor

var file_datanode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x64, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x75, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x75, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x64, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x42, 0x75, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x44, 0x61, 0x74, 0x61, 0x42, 0x75, 0x66, 0x22, 0x2c,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x62, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x36, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa4,
	0x02, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_datanode_proto_rawDescOnce sync.Once
	file_datanode_proto_rawDescData = file_datanode_proto_rawDesc
)

func file_datanode_proto_rawDescGZIP() []byte {
	file_datanode_proto_rawDescOnce.Do(func() {
		file_datanode_proto_rawDescData = protoimpl.X.CompressGZIP(file_datanode_proto_rawDescData)
	})
	return file_datanode_proto_rawDescData
}

var file_datanode_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_datanode_proto_goTypes = []interface{}{
	(*WriteBlockArgs)(nil),      // 0: service.WriteBlockArgs
	(*WriteBlockReply)(nil),     // 1: service.WriteBlockReply
	(*ReadBlockArgs)(nil),       // 2: service.ReadBlockArgs
	(*ReadBlockReply)(nil),      // 3: service.ReadBlockReply
	(*GetBlockInfoArgs)(nil),    // 4: service.GetBlockInfoArgs
	(*GetBlockInfoReply)(nil),   // 5: service.GetBlockInfoReply
	(*CreatePipelineArgs)(nil),  // 6: service.CreatePipelineArgs
	(*CreatePipelineReply)(nil), // 7: service.CreatePipelineReply
	(StatusCode)(0),             // 8: service.StatusCode
	(*BlockInfo)(nil),           // 9: service.BlockInfo
	(*DataNodeNetInfo)(nil),     // 10: service.DataNodeNetInfo
}
var file_datanode_proto_depIdxs = []int32{
	8,  // 0: service.WriteBlockReply.Status:type_name -> service.StatusCode
	8,  // 1: service.ReadBlockReply.Status:type_name -> service.StatusCode
	8,  // 2: service.GetBlockInfoReply.Status:type_name -> service.StatusCode
	9,  // 3: service.GetBlockInfoReply.Info:type_name -> service.BlockInfo
	10, // 4: service.CreatePipelineArgs.DataNodes:type_name -> service.DataNodeNetInfo
	8,  // 5: service.CreatePipelineReply.Status:type_name -> service.StatusCode
	0,  // 6: service.DataNodeService.WriteBlock:input_type -> service.WriteBlockArgs
	2,  // 7: service.DataNodeService.ReadBlock:input_type -> service.ReadBlockArgs
	4,  // 8: service.DataNodeService.GetBlockInfo:input_type -> service.GetBlockInfoArgs
	6,  // 9: service.DataNodeService.CreatePipeline:input_type -> service.CreatePipelineArgs
	1,  // 10: service.DataNodeService.WriteBlock:output_type -> service.WriteBlockReply
	3,  // 11: service.DataNodeService.ReadBlock:output_type -> service.ReadBlockReply
	5,  // 12: service.DataNodeService.GetBlockInfo:output_type -> service.GetBlockInfoReply
	7,  // 13: service.DataNodeService.CreatePipeline:output_type -> service.CreatePipelineReply
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_datanode_proto_init() }
func file_datanode_proto_init() {
	if File_datanode_proto != nil {
		return
	}
	file_statuscode_proto_init()
	file_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datanode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBlockArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBlockReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlockReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockInfoArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datanode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datanode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datanode_proto_goTypes,
		DependencyIndexes: file_datanode_proto_depIdxs,
		MessageInfos:      file_datanode_proto_msgTypes,
	}.Build()
	File_datanode_proto = out.File
	file_datanode_proto_rawDesc = nil
	file_datanode_proto_goTypes = nil
	file_datanode_proto_depIdxs = nil
}
