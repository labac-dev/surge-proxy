// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: protobuf/request/common.proto

package common

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_protobuf_request_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_request_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_protobuf_request_common_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdSignalsInfo *AdSignalsInfo `protobuf:"bytes,9,opt,name=adSignalsInfo,proto3" json:"adSignalsInfo,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	mi := &file_protobuf_request_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_request_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_protobuf_request_common_proto_rawDescGZIP(), []int{1}
}

func (x *Context) GetAdSignalsInfo() *AdSignalsInfo {
	if x != nil {
		return x.AdSignalsInfo
	}
	return nil
}

type AdSignalsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []*Params `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *AdSignalsInfo) Reset() {
	*x = AdSignalsInfo{}
	mi := &file_protobuf_request_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdSignalsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdSignalsInfo) ProtoMessage() {}

func (x *AdSignalsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_request_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdSignalsInfo.ProtoReflect.Descriptor instead.
func (*AdSignalsInfo) Descriptor() ([]byte, []int) {
	return file_protobuf_request_common_proto_rawDescGZIP(), []int{2}
}

func (x *AdSignalsInfo) GetParams() []*Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	mi := &file_protobuf_request_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_request_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_protobuf_request_common_proto_rawDescGZIP(), []int{3}
}

func (x *Params) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Params) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_protobuf_request_common_proto protoreflect.FileDescriptor

var file_protobuf_request_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x56, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x61, 0x64, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x61, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x41, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x30,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x61, 0x62, 0x61, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x75, 0x72, 0x67, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_request_common_proto_rawDescOnce sync.Once
	file_protobuf_request_common_proto_rawDescData = file_protobuf_request_common_proto_rawDesc
)

func file_protobuf_request_common_proto_rawDescGZIP() []byte {
	file_protobuf_request_common_proto_rawDescOnce.Do(func() {
		file_protobuf_request_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_request_common_proto_rawDescData)
	})
	return file_protobuf_request_common_proto_rawDescData
}

var file_protobuf_request_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_request_common_proto_goTypes = []any{
	(*Request)(nil),       // 0: youtube.request.common.Request
	(*Context)(nil),       // 1: youtube.request.common.Context
	(*AdSignalsInfo)(nil), // 2: youtube.request.common.AdSignalsInfo
	(*Params)(nil),        // 3: youtube.request.common.Params
}
var file_protobuf_request_common_proto_depIdxs = []int32{
	1, // 0: youtube.request.common.Request.context:type_name -> youtube.request.common.Context
	2, // 1: youtube.request.common.Context.adSignalsInfo:type_name -> youtube.request.common.AdSignalsInfo
	3, // 2: youtube.request.common.AdSignalsInfo.params:type_name -> youtube.request.common.Params
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_request_common_proto_init() }
func file_protobuf_request_common_proto_init() {
	if File_protobuf_request_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_request_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_request_common_proto_goTypes,
		DependencyIndexes: file_protobuf_request_common_proto_depIdxs,
		MessageInfos:      file_protobuf_request_common_proto_msgTypes,
	}.Build()
	File_protobuf_request_common_proto = out.File
	file_protobuf_request_common_proto_rawDesc = nil
	file_protobuf_request_common_proto_goTypes = nil
	file_protobuf_request_common_proto_depIdxs = nil
}
