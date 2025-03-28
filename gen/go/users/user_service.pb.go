// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.23.4
// source: users/user_service.proto

package users_pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	shared "github.com/reversersed/LitGO-proto/gen/go/shared"
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

// Requests
type LoginRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	 
	Login         string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty" validate:"required" example:"admin"`       // Can be presented as login or email
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required" example:"admin"`  
	RememberMe    bool   `protobuf:"varint,3,opt,name=rememberMe,proto3" json:"rememberMe,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_users_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetRememberMe() bool {
	if x != nil {
		return x.RememberMe
	}
	return false
}

type TokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Refreshtoken  string                 `protobuf:"bytes,1,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty" validate:"required"`  
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	mi := &file_users_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *TokenRequest) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"primitiveid,required_without_all=Login Email" form:"id"`        
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty" validate:"onlyenglish,required_without_all=Id Email" form:"login"`  
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" validate:"required_without_all=Id Login" form:"email"`  
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_users_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegistrationRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Login          string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty" validate:"required,min=4,max=16,onlyenglish"`                                          
	Password       string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validate:"required,min=8,max=32,lowercase,uppercase,digitrequired,specialsymbol"`                                    
	PasswordRepeat string                 `protobuf:"bytes,3,opt,name=password_repeat,json=passwordRepeat,proto3" json:"password_repeat,omitempty" validate:"required,eqfield=Password"`  
	Email          string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" validate:"required,email"`                                          
	RememberMe     bool                   `protobuf:"varint,5,opt,name=rememberMe,proto3" json:"rememberMe,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	mi := &file_users_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *RegistrationRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegistrationRequest) GetPasswordRepeat() string {
	if x != nil {
		return x.PasswordRepeat
	}
	return ""
}

func (x *RegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegistrationRequest) GetRememberMe() bool {
	if x != nil {
		return x.RememberMe
	}
	return false
}

// Responses
type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Token         string                 `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Refreshtoken  string                 `protobuf:"bytes,5,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_users_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

type TokenReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" example:"token"`                
	Refreshtoken  string                 `protobuf:"bytes,2,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty" example:"refresh token"`  
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenReply) Reset() {
	*x = TokenReply{}
	mi := &file_users_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReply) ProtoMessage() {}

func (x *TokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReply.ProtoReflect.Descriptor instead.
func (*TokenReply) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *TokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenReply) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

// Models
type UserModel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Roles         []string               `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	mi := &file_users_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_users_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_users_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *UserModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserModel) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserModel) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_users_user_service_proto protoreflect.FileDescriptor

const file_users_user_service_proto_rawDesc = "" +
	"\n" +
	"\x18users/user_service.proto\x12\x05users\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x13shared/shared.proto\"`\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1e\n" +
	"\n" +
	"rememberMe\x18\x03 \x01(\bR\n" +
	"rememberMe\"2\n" +
	"\fTokenRequest\x12\"\n" +
	"\frefreshtoken\x18\x01 \x01(\tR\frefreshtoken\"I\n" +
	"\vUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"\xa6\x01\n" +
	"\x13RegistrationRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12'\n" +
	"\x0fpassword_repeat\x18\x03 \x01(\tR\x0epasswordRepeat\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x1e\n" +
	"\n" +
	"rememberMe\x18\x05 \x01(\bR\n" +
	"rememberMe\"\x85\x01\n" +
	"\rLoginResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login\x12\x14\n" +
	"\x05roles\x18\x03 \x03(\tR\x05roles\x12\x14\n" +
	"\x05token\x18\x04 \x01(\tR\x05token\x12\"\n" +
	"\frefreshtoken\x18\x05 \x01(\tR\frefreshtoken\"F\n" +
	"\n" +
	"TokenReply\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\"\n" +
	"\frefreshtoken\x18\x02 \x01(\tR\frefreshtoken\"]\n" +
	"\tUserModel\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05roles\x18\x04 \x03(\tR\x05roles2\xa2\a\n" +
	"\x04User\x12\xa5\x01\n" +
	"\x04Auth\x12\r.shared.Empty\x1a\x17.shared.UserCredentials\"u\x92AU\x12\x11Authenticate user\x1a2Returns current user's credentials if there is oneb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/api/v1/users/auth\x12\x8d\x01\n" +
	"\x06Logout\x12\r.shared.Empty\x1a\r.shared.Empty\"e\x92AC\x12\vLogout user\x1a&Removes user cookies if user logged inb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/api/v1/users/logout\x12\xa7\x01\n" +
	"\x05Login\x12\x13.users.LoginRequest\x1a\x14.users.LoginResponse\"s\x92AR\x12\n" +
	"Login user\x1aDLog in user with provided credentials to get auth and refresh tokens\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/api/v1/users/login\x125\n" +
	"\vUpdateToken\x12\x13.users.TokenRequest\x1a\x11.users.TokenReply\x12\xa5\x01\n" +
	"\aGetUser\x12\x12.users.UserRequest\x1a\x10.users.UserModel\"t\x92A\\\x12\vSearch user\x1aMFind user by provided id, login or email. Only one parameter at time allowed.\x82\xd3\xe4\x93\x02\x0f\x12\r/api/v1/users\x12\xd8\x01\n" +
	"\fRegisterUser\x12\x1a.users.RegistrationRequest\x1a\x14.users.LoginResponse\"\x95\x01\x92Aq\x12\rRegister user\x1a`Creates user account with provided credentials and returns auth and refresh tokens to authorize.\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/api/v1/users/registerB\x98\x01\x92A[ZY\n" +
	"W\n" +
	"\x06bearer\x12M\b\x02\x128Authentication token, prefixed by Bearer: Bearer <token>\x1a\rAuthorization \x02Z8github.com/reversersed/LitGO-proto/gen/go/users;users_pbb\x06proto3"

var (
	file_users_user_service_proto_rawDescOnce sync.Once
	file_users_user_service_proto_rawDescData []byte
)

func file_users_user_service_proto_rawDescGZIP() []byte {
	file_users_user_service_proto_rawDescOnce.Do(func() {
		file_users_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_users_user_service_proto_rawDesc), len(file_users_user_service_proto_rawDesc)))
	})
	return file_users_user_service_proto_rawDescData
}

var file_users_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_users_user_service_proto_goTypes = []any{
	(*LoginRequest)(nil),           // 0: users.LoginRequest
	(*TokenRequest)(nil),           // 1: users.TokenRequest
	(*UserRequest)(nil),            // 2: users.UserRequest
	(*RegistrationRequest)(nil),    // 3: users.RegistrationRequest
	(*LoginResponse)(nil),          // 4: users.LoginResponse
	(*TokenReply)(nil),             // 5: users.TokenReply
	(*UserModel)(nil),              // 6: users.UserModel
	(*shared.Empty)(nil),           // 7: shared.Empty
	(*shared.UserCredentials)(nil), // 8: shared.UserCredentials
}
var file_users_user_service_proto_depIdxs = []int32{
	7, // 0: users.User.Auth:input_type -> shared.Empty
	7, // 1: users.User.Logout:input_type -> shared.Empty
	0, // 2: users.User.Login:input_type -> users.LoginRequest
	1, // 3: users.User.UpdateToken:input_type -> users.TokenRequest
	2, // 4: users.User.GetUser:input_type -> users.UserRequest
	3, // 5: users.User.RegisterUser:input_type -> users.RegistrationRequest
	8, // 6: users.User.Auth:output_type -> shared.UserCredentials
	7, // 7: users.User.Logout:output_type -> shared.Empty
	4, // 8: users.User.Login:output_type -> users.LoginResponse
	5, // 9: users.User.UpdateToken:output_type -> users.TokenReply
	6, // 10: users.User.GetUser:output_type -> users.UserModel
	4, // 11: users.User.RegisterUser:output_type -> users.LoginResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_user_service_proto_init() }
func file_users_user_service_proto_init() {
	if File_users_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_users_user_service_proto_rawDesc), len(file_users_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_user_service_proto_goTypes,
		DependencyIndexes: file_users_user_service_proto_depIdxs,
		MessageInfos:      file_users_user_service_proto_msgTypes,
	}.Build()
	File_users_user_service_proto = out.File
	file_users_user_service_proto_goTypes = nil
	file_users_user_service_proto_depIdxs = nil
}
