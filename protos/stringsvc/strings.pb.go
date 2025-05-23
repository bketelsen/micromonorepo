// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: stringsvc/strings.proto

package stringsvc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpperRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpperRequest) Reset() {
	*x = UpperRequest{}
	mi := &file_stringsvc_strings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpperRequest) ProtoMessage() {}

func (x *UpperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_strings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpperRequest.ProtoReflect.Descriptor instead.
func (*UpperRequest) Descriptor() ([]byte, []int) {
	return file_stringsvc_strings_proto_rawDescGZIP(), []int{0}
}

func (x *UpperRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpperResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpperResponse) Reset() {
	*x = UpperResponse{}
	mi := &file_stringsvc_strings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpperResponse) ProtoMessage() {}

func (x *UpperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stringsvc_strings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpperResponse.ProtoReflect.Descriptor instead.
func (*UpperResponse) Descriptor() ([]byte, []int) {
	return file_stringsvc_strings_proto_rawDescGZIP(), []int{1}
}

func (x *UpperResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_stringsvc_strings_proto protoreflect.FileDescriptor

const file_stringsvc_strings_proto_rawDesc = "" +
	"\n" +
	"\x17stringsvc/strings.proto\x1a\x1cgoogle/api/annotations.proto\"$\n" +
	"\fUpperRequest\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\"%\n" +
	"\rUpperResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value2B\n" +
	"\x05Tools\x129\n" +
	"\x05Upper\x12\r.UpperRequest\x1a\x0e.UpperResponse\"\x11\x82\xd3\xe4\x93\x02\v:\x01*\"\x06/upperB\x0eZ\f../stringsvcb\x06proto3"

var (
	file_stringsvc_strings_proto_rawDescOnce sync.Once
	file_stringsvc_strings_proto_rawDescData []byte
)

func file_stringsvc_strings_proto_rawDescGZIP() []byte {
	file_stringsvc_strings_proto_rawDescOnce.Do(func() {
		file_stringsvc_strings_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stringsvc_strings_proto_rawDesc), len(file_stringsvc_strings_proto_rawDesc)))
	})
	return file_stringsvc_strings_proto_rawDescData
}

var file_stringsvc_strings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stringsvc_strings_proto_goTypes = []any{
	(*UpperRequest)(nil),  // 0: UpperRequest
	(*UpperResponse)(nil), // 1: UpperResponse
}
var file_stringsvc_strings_proto_depIdxs = []int32{
	0, // 0: Tools.Upper:input_type -> UpperRequest
	1, // 1: Tools.Upper:output_type -> UpperResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stringsvc_strings_proto_init() }
func file_stringsvc_strings_proto_init() {
	if File_stringsvc_strings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stringsvc_strings_proto_rawDesc), len(file_stringsvc_strings_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stringsvc_strings_proto_goTypes,
		DependencyIndexes: file_stringsvc_strings_proto_depIdxs,
		MessageInfos:      file_stringsvc_strings_proto_msgTypes,
	}.Build()
	File_stringsvc_strings_proto = out.File
	file_stringsvc_strings_proto_goTypes = nil
	file_stringsvc_strings_proto_depIdxs = nil
}
