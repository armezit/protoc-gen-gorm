// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: feature_demo/demo_multi_file_service.proto

package example

import (
	query "github.com/armezit/atlas-app-toolkit/query"
	_ "github.com/armezit/protoc-gen-gorm/options"
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

type ReadAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For a read request, the id field is the only to be specified
	Id     uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields *query.FieldSelection `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *ReadAccountRequest) Reset() {
	*x = ReadAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_demo_demo_multi_file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAccountRequest) ProtoMessage() {}

func (x *ReadAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feature_demo_demo_multi_file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAccountRequest.ProtoReflect.Descriptor instead.
func (*ReadAccountRequest) Descriptor() ([]byte, []int) {
	return file_feature_demo_demo_multi_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReadAccountRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReadAccountRequest) GetFields() *query.FieldSelection {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ReadBlogPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*BlogPost `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ReadBlogPostsResponse) Reset() {
	*x = ReadBlogPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_demo_demo_multi_file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogPostsResponse) ProtoMessage() {}

func (x *ReadBlogPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feature_demo_demo_multi_file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogPostsResponse.ProtoReflect.Descriptor instead.
func (*ReadBlogPostsResponse) Descriptor() ([]byte, []int) {
	return file_feature_demo_demo_multi_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadBlogPostsResponse) GetPosts() []*BlogPost {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_feature_demo_demo_multi_file_service_proto protoreflect.FileDescriptor

var file_feature_demo_demo_multi_file_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61, 0x74, 0x6c, 0x61, 0x73,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0x5e, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f,
	0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feature_demo_demo_multi_file_service_proto_rawDescOnce sync.Once
	file_feature_demo_demo_multi_file_service_proto_rawDescData = file_feature_demo_demo_multi_file_service_proto_rawDesc
)

func file_feature_demo_demo_multi_file_service_proto_rawDescGZIP() []byte {
	file_feature_demo_demo_multi_file_service_proto_rawDescOnce.Do(func() {
		file_feature_demo_demo_multi_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_demo_demo_multi_file_service_proto_rawDescData)
	})
	return file_feature_demo_demo_multi_file_service_proto_rawDescData
}

var file_feature_demo_demo_multi_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feature_demo_demo_multi_file_service_proto_goTypes = []interface{}{
	(*ReadAccountRequest)(nil),    // 0: example.ReadAccountRequest
	(*ReadBlogPostsResponse)(nil), // 1: example.ReadBlogPostsResponse
	(*query.FieldSelection)(nil),  // 2: atlas.query.v1.FieldSelection
	(*BlogPost)(nil),              // 3: example.BlogPost
}
var file_feature_demo_demo_multi_file_service_proto_depIdxs = []int32{
	2, // 0: example.ReadAccountRequest.fields:type_name -> atlas.query.v1.FieldSelection
	3, // 1: example.ReadBlogPostsResponse.posts:type_name -> example.BlogPost
	0, // 2: example.BlogPostService.Read:input_type -> example.ReadAccountRequest
	1, // 3: example.BlogPostService.Read:output_type -> example.ReadBlogPostsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_feature_demo_demo_multi_file_service_proto_init() }
func file_feature_demo_demo_multi_file_service_proto_init() {
	if File_feature_demo_demo_multi_file_service_proto != nil {
		return
	}
	file_feature_demo_demo_multi_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_feature_demo_demo_multi_file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAccountRequest); i {
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
		file_feature_demo_demo_multi_file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogPostsResponse); i {
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
			RawDescriptor: file_feature_demo_demo_multi_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feature_demo_demo_multi_file_service_proto_goTypes,
		DependencyIndexes: file_feature_demo_demo_multi_file_service_proto_depIdxs,
		MessageInfos:      file_feature_demo_demo_multi_file_service_proto_msgTypes,
	}.Build()
	File_feature_demo_demo_multi_file_service_proto = out.File
	file_feature_demo_demo_multi_file_service_proto_rawDesc = nil
	file_feature_demo_demo_multi_file_service_proto_goTypes = nil
	file_feature_demo_demo_multi_file_service_proto_depIdxs = nil
}
