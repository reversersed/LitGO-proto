// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.23.4
// source: shared/shared.proto

package shared_pb

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

// @Description Error detail contains information about error
type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	 
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty" example:"Token"` // Field that error occured on
	 
	Struct string `protobuf:"bytes,2,opt,name=struct,proto3" json:"struct,omitempty" example:"users_pb.TokenRequest"` // Structure that contains field
	 
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty" example:"jwt"` // Failed validation tag
	 
	TagValue string `protobuf:"bytes,4,opt,name=tagValue,proto3" json:"tagValue,omitempty" example:"5"` // Valitation tag value
	 
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" example:"Field must be a jwt token"` // Error description. Only development purposes, do not show users
	 
	Actualvalue string `protobuf:"bytes,6,opt,name=actualvalue,proto3" json:"actualvalue,omitempty" example:"token"` // Actual value of field that causes the error. Note: 'password' field will be hidden
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	mi := &file_shared_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_shared_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_shared_shared_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetail) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ErrorDetail) GetStruct() string {
	if x != nil {
		return x.Struct
	}
	return ""
}

func (x *ErrorDetail) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ErrorDetail) GetTagValue() string {
	if x != nil {
		return x.TagValue
	}
	return ""
}

func (x *ErrorDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ErrorDetail) GetActualvalue() string {
	if x != nil {
		return x.Actualvalue
	}
	return ""
}

type UserCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Roles []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserCredentials) Reset() {
	*x = UserCredentials{}
	mi := &file_shared_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredentials) ProtoMessage() {}

func (x *UserCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_shared_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredentials.ProtoReflect.Descriptor instead.
func (*UserCredentials) Descriptor() ([]byte, []int) {
	return file_shared_shared_proto_rawDescGZIP(), []int{1}
}

func (x *UserCredentials) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserCredentials) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserCredentials) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_shared_shared_proto protoreflect.FileDescriptor

var file_shared_shared_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x22, 0xad, 0x01,
	0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x75, 0x61, 0x6c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4d, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x64, 0x2f, 0x4c, 0x69, 0x74, 0x47, 0x4f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_shared_shared_proto_rawDescOnce sync.Once
	file_shared_shared_proto_rawDescData = file_shared_shared_proto_rawDesc
)

func file_shared_shared_proto_rawDescGZIP() []byte {
	file_shared_shared_proto_rawDescOnce.Do(func() {
		file_shared_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_shared_proto_rawDescData)
	})
	return file_shared_shared_proto_rawDescData
}

var file_shared_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_shared_proto_goTypes = []any{
	(*ErrorDetail)(nil),     // 0: shared.ErrorDetail
	(*UserCredentials)(nil), // 1: shared.UserCredentials
}
var file_shared_shared_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shared_shared_proto_init() }
func file_shared_shared_proto_init() {
	if File_shared_shared_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_shared_proto_goTypes,
		DependencyIndexes: file_shared_shared_proto_depIdxs,
		MessageInfos:      file_shared_shared_proto_msgTypes,
	}.Build()
	File_shared_shared_proto = out.File
	file_shared_shared_proto_rawDesc = nil
	file_shared_shared_proto_goTypes = nil
	file_shared_shared_proto_depIdxs = nil
}
