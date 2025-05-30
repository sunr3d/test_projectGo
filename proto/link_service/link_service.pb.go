// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: link_service/link_service.proto

package link_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type GetLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Link          string                 `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLinkRequest) Reset() {
	*x = GetLinkRequest{}
	mi := &file_link_service_link_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkRequest) ProtoMessage() {}

func (x *GetLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_service_link_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkRequest.ProtoReflect.Descriptor instead.
func (*GetLinkRequest) Descriptor() ([]byte, []int) {
	return file_link_service_link_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Link          string                 `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLinkResponse) Reset() {
	*x = GetLinkResponse{}
	mi := &file_link_service_link_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkResponse) ProtoMessage() {}

func (x *GetLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_service_link_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkResponse.ProtoReflect.Descriptor instead.
func (*GetLinkResponse) Descriptor() ([]byte, []int) {
	return file_link_service_link_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLinkResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type InputLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Link          string                 `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	FakeLink      string                 `protobuf:"bytes,2,opt,name=fakeLink,proto3" json:"fakeLink,omitempty"`
	EraseTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=eraseTime,proto3" json:"eraseTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputLinkRequest) Reset() {
	*x = InputLinkRequest{}
	mi := &file_link_service_link_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputLinkRequest) ProtoMessage() {}

func (x *InputLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_service_link_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputLinkRequest.ProtoReflect.Descriptor instead.
func (*InputLinkRequest) Descriptor() ([]byte, []int) {
	return file_link_service_link_service_proto_rawDescGZIP(), []int{2}
}

func (x *InputLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *InputLinkRequest) GetFakeLink() string {
	if x != nil {
		return x.FakeLink
	}
	return ""
}

func (x *InputLinkRequest) GetEraseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EraseTime
	}
	return nil
}

var File_link_service_link_service_proto protoreflect.FileDescriptor

const file_link_service_link_service_proto_rawDesc = "" +
	"\n" +
	"\x1flink_service/link_service.proto\x12\flink_service\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\"$\n" +
	"\x0eGetLinkRequest\x12\x12\n" +
	"\x04link\x18\x01 \x01(\tR\x04link\"%\n" +
	"\x0fGetLinkResponse\x12\x12\n" +
	"\x04link\x18\x01 \x01(\tR\x04link\"|\n" +
	"\x10InputLinkRequest\x12\x12\n" +
	"\x04link\x18\x01 \x01(\tR\x04link\x12\x1a\n" +
	"\bfakeLink\x18\x02 \x01(\tR\bfakeLink\x128\n" +
	"\teraseTime\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\teraseTime2\xc2\x01\n" +
	"\vLinkService\x12\\\n" +
	"\aGetLink\x12\x1c.link_service.GetLinkRequest\x1a\x1d.link_service.GetLinkResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/link/{link}\x12U\n" +
	"\tInputLink\x12\x1e.link_service.InputLinkRequest\x1a\x16.google.protobuf.Empty\"\x10\x82\xd3\xe4\x93\x02\n" +
	":\x01*\"\x05/linkB\x1cZ\x1a/link_service;link_serviceb\x06proto3"

var (
	file_link_service_link_service_proto_rawDescOnce sync.Once
	file_link_service_link_service_proto_rawDescData []byte
)

func file_link_service_link_service_proto_rawDescGZIP() []byte {
	file_link_service_link_service_proto_rawDescOnce.Do(func() {
		file_link_service_link_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_link_service_link_service_proto_rawDesc), len(file_link_service_link_service_proto_rawDesc)))
	})
	return file_link_service_link_service_proto_rawDescData
}

var file_link_service_link_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_link_service_link_service_proto_goTypes = []any{
	(*GetLinkRequest)(nil),        // 0: link_service.GetLinkRequest
	(*GetLinkResponse)(nil),       // 1: link_service.GetLinkResponse
	(*InputLinkRequest)(nil),      // 2: link_service.InputLinkRequest
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_link_service_link_service_proto_depIdxs = []int32{
	3, // 0: link_service.InputLinkRequest.eraseTime:type_name -> google.protobuf.Timestamp
	0, // 1: link_service.LinkService.GetLink:input_type -> link_service.GetLinkRequest
	2, // 2: link_service.LinkService.InputLink:input_type -> link_service.InputLinkRequest
	1, // 3: link_service.LinkService.GetLink:output_type -> link_service.GetLinkResponse
	4, // 4: link_service.LinkService.InputLink:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_link_service_link_service_proto_init() }
func file_link_service_link_service_proto_init() {
	if File_link_service_link_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_link_service_link_service_proto_rawDesc), len(file_link_service_link_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_service_link_service_proto_goTypes,
		DependencyIndexes: file_link_service_link_service_proto_depIdxs,
		MessageInfos:      file_link_service_link_service_proto_msgTypes,
	}.Build()
	File_link_service_link_service_proto = out.File
	file_link_service_link_service_proto_goTypes = nil
	file_link_service_link_service_proto_depIdxs = nil
}
