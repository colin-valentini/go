// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: internal/proto/example/example.proto

package example

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

type Language int32

const (
	Language_LANGUAGE_UNSPECIFIED Language = 0
	Language_LANGUAGE_ENGLISH     Language = 1
	Language_LANGUAGE_SPANISH     Language = 2
	Language_LANGUAGE_ITALIAN     Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "LANGUAGE_UNSPECIFIED",
		1: "LANGUAGE_ENGLISH",
		2: "LANGUAGE_SPANISH",
		3: "LANGUAGE_ITALIAN",
	}
	Language_value = map[string]int32{
		"LANGUAGE_UNSPECIFIED": 0,
		"LANGUAGE_ENGLISH":     1,
		"LANGUAGE_SPANISH":     2,
		"LANGUAGE_ITALIAN":     3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_example_example_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_internal_proto_example_example_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_example_example_proto_rawDescGZIP(), []int{0}
}

type GreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=example.Language" json:"language,omitempty"`
}

func (x *GreetingRequest) Reset() {
	*x = GreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRequest) ProtoMessage() {}

func (x *GreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRequest.ProtoReflect.Descriptor instead.
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GreetingRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

type GreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_example_example_proto protoreflect.FileDescriptor

var file_internal_proto_example_example_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22,
	0x50, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0x66, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4c,
	0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47,
	0x45, 0x5f, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4c,
	0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x54,
	0x41, 0x4c, 0x49, 0x41, 0x4e, 0x10, 0x03, 0x32, 0x52, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69, 0x6e, 0x2d,
	0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_example_example_proto_rawDescOnce sync.Once
	file_internal_proto_example_example_proto_rawDescData = file_internal_proto_example_example_proto_rawDesc
)

func file_internal_proto_example_example_proto_rawDescGZIP() []byte {
	file_internal_proto_example_example_proto_rawDescOnce.Do(func() {
		file_internal_proto_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_example_example_proto_rawDescData)
	})
	return file_internal_proto_example_example_proto_rawDescData
}

var file_internal_proto_example_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_proto_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_example_example_proto_goTypes = []interface{}{
	(Language)(0),            // 0: example.Language
	(*GreetingRequest)(nil),  // 1: example.GreetingRequest
	(*GreetingResponse)(nil), // 2: example.GreetingResponse
}
var file_internal_proto_example_example_proto_depIdxs = []int32{
	0, // 0: example.GreetingRequest.language:type_name -> example.Language
	1, // 1: example.HelloWorld.GetGreeting:input_type -> example.GreetingRequest
	2, // 2: example.HelloWorld.GetGreeting:output_type -> example.GreetingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_example_example_proto_init() }
func file_internal_proto_example_example_proto_init() {
	if File_internal_proto_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRequest); i {
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
		file_internal_proto_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingResponse); i {
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
			RawDescriptor: file_internal_proto_example_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_example_example_proto_goTypes,
		DependencyIndexes: file_internal_proto_example_example_proto_depIdxs,
		EnumInfos:         file_internal_proto_example_example_proto_enumTypes,
		MessageInfos:      file_internal_proto_example_example_proto_msgTypes,
	}.Build()
	File_internal_proto_example_example_proto = out.File
	file_internal_proto_example_example_proto_rawDesc = nil
	file_internal_proto_example_example_proto_goTypes = nil
	file_internal_proto_example_example_proto_depIdxs = nil
}
