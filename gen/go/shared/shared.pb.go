// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.23.4
// source: shared/shared.proto

package shared_pb

import (
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

// @Description Error detail contains information about error
type ErrorDetail struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	 
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty" example:"Token"` // Field that error occured on
	 
	Struct string `protobuf:"bytes,2,opt,name=struct,proto3" json:"struct,omitempty" example:"users_pb.TokenRequest"` // Structure that contains field
	 
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty" example:"jwt"` // Failed validation tag
	 
	TagValue string `protobuf:"bytes,4,opt,name=tagValue,proto3" json:"tagValue,omitempty" example:"5"` // Valitation tag value
	 
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" example:"Field must be a jwt token"` // Error description. Only development purposes, do not show users
	 
	Actualvalue   string `protobuf:"bytes,6,opt,name=actualvalue,proto3" json:"actualvalue,omitempty" example:"token"` // Actual value of field that causes the error. Note: 'password' field will be hidden
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_shared_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_shared_shared_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_shared_shared_proto_rawDescGZIP(), []int{2}
}

var File_shared_shared_proto protoreflect.FileDescriptor

const file_shared_shared_proto_rawDesc = "" +
	"\n" +
	"\x13shared/shared.proto\x12\x06shared\"\xad\x01\n" +
	"\vErrorDetail\x12\x14\n" +
	"\x05field\x18\x01 \x01(\tR\x05field\x12\x16\n" +
	"\x06struct\x18\x02 \x01(\tR\x06struct\x12\x10\n" +
	"\x03tag\x18\x03 \x01(\tR\x03tag\x12\x1a\n" +
	"\btagValue\x18\x04 \x01(\tR\btagValue\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12 \n" +
	"\vactualvalue\x18\x06 \x01(\tR\vactualvalue\"M\n" +
	"\x0fUserCredentials\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login\x12\x14\n" +
	"\x05roles\x18\x03 \x03(\tR\x05roles\"\a\n" +
	"\x05EmptyB<Z:github.com/reversersed/LitGO-proto/gen/go/shared;shared_pbb\x06proto3"

var (
	file_shared_shared_proto_rawDescOnce sync.Once
	file_shared_shared_proto_rawDescData []byte
)

func file_shared_shared_proto_rawDescGZIP() []byte {
	file_shared_shared_proto_rawDescOnce.Do(func() {
		file_shared_shared_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shared_shared_proto_rawDesc), len(file_shared_shared_proto_rawDesc)))
	})
	return file_shared_shared_proto_rawDescData
}

var file_shared_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shared_shared_proto_goTypes = []any{
	(*ErrorDetail)(nil),     // 0: shared.ErrorDetail
	(*UserCredentials)(nil), // 1: shared.UserCredentials
	(*Empty)(nil),           // 2: shared.Empty
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shared_shared_proto_rawDesc), len(file_shared_shared_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_shared_proto_goTypes,
		DependencyIndexes: file_shared_shared_proto_depIdxs,
		MessageInfos:      file_shared_shared_proto_msgTypes,
	}.Build()
	File_shared_shared_proto = out.File
	file_shared_shared_proto_goTypes = nil
	file_shared_shared_proto_depIdxs = nil
}
