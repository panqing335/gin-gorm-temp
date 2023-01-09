// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: protos/grpcDemo.proto

package grpcService

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

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpcDemo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpcDemo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_protos_grpcDemo_proto_rawDescGZIP(), []int{0}
}

func (x *DemoRequest) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type DemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DemoReply) Reset() {
	*x = DemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_grpcDemo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoReply) ProtoMessage() {}

func (x *DemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_grpcDemo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoReply.ProtoReflect.Descriptor instead.
func (*DemoReply) Descriptor() ([]byte, []int) {
	return file_protos_grpcDemo_proto_rawDescGZIP(), []int{1}
}

func (x *DemoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_grpcDemo_proto protoreflect.FileDescriptor

var file_protos_grpcDemo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x44, 0x65, 0x6d,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x45,
	0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x3d, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x3b, 0x20, 0x67, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_grpcDemo_proto_rawDescOnce sync.Once
	file_protos_grpcDemo_proto_rawDescData = file_protos_grpcDemo_proto_rawDesc
)

func file_protos_grpcDemo_proto_rawDescGZIP() []byte {
	file_protos_grpcDemo_proto_rawDescOnce.Do(func() {
		file_protos_grpcDemo_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_grpcDemo_proto_rawDescData)
	})
	return file_protos_grpcDemo_proto_rawDescData
}

var file_protos_grpcDemo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_grpcDemo_proto_goTypes = []interface{}{
	(*DemoRequest)(nil), // 0: grpcService.DemoRequest
	(*DemoReply)(nil),   // 1: grpcService.DemoReply
}
var file_protos_grpcDemo_proto_depIdxs = []int32{
	0, // 0: grpcService.Demo.UnaryCall:input_type -> grpcService.DemoRequest
	1, // 1: grpcService.Demo.UnaryCall:output_type -> grpcService.DemoReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_grpcDemo_proto_init() }
func file_protos_grpcDemo_proto_init() {
	if File_protos_grpcDemo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_grpcDemo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoRequest); i {
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
		file_protos_grpcDemo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoReply); i {
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
			RawDescriptor: file_protos_grpcDemo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_grpcDemo_proto_goTypes,
		DependencyIndexes: file_protos_grpcDemo_proto_depIdxs,
		MessageInfos:      file_protos_grpcDemo_proto_msgTypes,
	}.Build()
	File_protos_grpcDemo_proto = out.File
	file_protos_grpcDemo_proto_rawDesc = nil
	file_protos_grpcDemo_proto_goTypes = nil
	file_protos_grpcDemo_proto_depIdxs = nil
}