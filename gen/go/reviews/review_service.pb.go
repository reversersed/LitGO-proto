// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.23.4
// source: reviews/review_service.proto

package reviews_pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// Models
type UserActionEnum int32

const (
	UserActionEnum_noAction UserActionEnum = 0
	UserActionEnum_like     UserActionEnum = 1
	UserActionEnum_dislike  UserActionEnum = 2
)

// Enum value maps for UserActionEnum.
var (
	UserActionEnum_name = map[int32]string{
		0: "noAction",
		1: "like",
		2: "dislike",
	}
	UserActionEnum_value = map[string]int32{
		"noAction": 0,
		"like":     1,
		"dislike":  2,
	}
)

func (x UserActionEnum) Enum() *UserActionEnum {
	p := new(UserActionEnum)
	*p = x
	return p
}

func (x UserActionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserActionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_reviews_review_service_proto_enumTypes[0].Descriptor()
}

func (UserActionEnum) Type() protoreflect.EnumType {
	return &file_reviews_review_service_proto_enumTypes[0]
}

func (x UserActionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserActionEnum.Descriptor instead.
func (UserActionEnum) EnumDescriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{0}
}

// Requests
type GetBookReviewsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Book ID to search reviews of
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" validate:"required,primitiveid"`  
	// Page number to search
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page" validate:"gte=0"`  
	// Reviews count per page
	PageSize      int32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty" form:"pagesize" validate:"gte=1"`  
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookReviewsRequest) Reset() {
	*x = GetBookReviewsRequest{}
	mi := &file_reviews_review_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReviewsRequest) ProtoMessage() {}

func (x *GetBookReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetBookReviewsRequest) Descriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBookReviewsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBookReviewsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBookReviewsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CreateBookReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty" form:"Id" validate:"required"`            
	Rating        float32                `protobuf:"fixed32,2,opt,name=rating,proto3" json:"rating,omitempty" form:"Rating" validate:"required,gte=0,lte=5"`      
	CreatorId     string                 `protobuf:"bytes,3,opt,name=creatorId,proto3" json:"creatorId,omitempty" validate:"required,primitiveid" swaggerignore:"true"`  
	ModelId       string                 `protobuf:"bytes,4,opt,name=modelId,proto3" json:"modelId,omitempty" form:"ModelId" validate:"required,primitiveid"`      
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookReviewRequest) Reset() {
	*x = CreateBookReviewRequest{}
	mi := &file_reviews_review_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookReviewRequest) ProtoMessage() {}

func (x *CreateBookReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateBookReviewRequest) Descriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookReviewRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateBookReviewRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateBookReviewRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *CreateBookReviewRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

// Responses
type GetBookReviewsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reviews       []*ReviewModel         `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookReviewsResponse) Reset() {
	*x = GetBookReviewsResponse{}
	mi := &file_reviews_review_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReviewsResponse) ProtoMessage() {}

func (x *GetBookReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetBookReviewsResponse) Descriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookReviewsResponse) GetReviews() []*ReviewModel {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type CreateBookReviewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *ReviewModel           `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookReviewResponse) Reset() {
	*x = CreateBookReviewResponse{}
	mi := &file_reviews_review_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookReviewResponse) ProtoMessage() {}

func (x *CreateBookReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookReviewResponse.ProtoReflect.Descriptor instead.
func (*CreateBookReviewResponse) Descriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookReviewResponse) GetReview() *ReviewModel {
	if x != nil {
		return x.Review
	}
	return nil
}

type ReviewModel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Created       int64                  `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Creator       *UserModel             `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Rating        float32                `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
	UserAction    UserActionEnum         `protobuf:"varint,6,opt,name=userAction,proto3,enum=reviews.UserActionEnum" json:"userAction,omitempty"`
	Upvotes       int32                  `protobuf:"varint,7,opt,name=upvotes,proto3" json:"upvotes,omitempty"`
	Downvotes     int32                  `protobuf:"varint,8,opt,name=downvotes,proto3" json:"downvotes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewModel) Reset() {
	*x = ReviewModel{}
	mi := &file_reviews_review_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewModel) ProtoMessage() {}

func (x *ReviewModel) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewModel.ProtoReflect.Descriptor instead.
func (*ReviewModel) Descriptor() ([]byte, []int) {
	return file_reviews_review_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReviewModel) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ReviewModel) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *ReviewModel) GetCreator() *UserModel {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *ReviewModel) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ReviewModel) GetUserAction() UserActionEnum {
	if x != nil {
		return x.UserAction
	}
	return UserActionEnum_noAction
}

func (x *ReviewModel) GetUpvotes() int32 {
	if x != nil {
		return x.Upvotes
	}
	return 0
}

func (x *ReviewModel) GetDownvotes() int32 {
	if x != nil {
		return x.Downvotes
	}
	return 0
}

type UserModel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	mi := &file_reviews_review_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_review_service_proto_msgTypes[5]
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
	return file_reviews_review_service_proto_rawDescGZIP(), []int{5}
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

var File_reviews_review_service_proto protoreflect.FileDescriptor

const file_reviews_review_service_proto_rawDesc = "" +
	"\n" +
	"\x1creviews/review_service.proto\x12\areviews\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1cgoogle/api/annotations.proto\"W\n" +
	"\x15GetBookReviewsRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x05R\x04page\x12\x1a\n" +
	"\bpageSize\x18\x03 \x01(\x05R\bpageSize\"}\n" +
	"\x17CreateBookReviewRequest\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\x12\x16\n" +
	"\x06rating\x18\x02 \x01(\x02R\x06rating\x12\x1c\n" +
	"\tcreatorId\x18\x03 \x01(\tR\tcreatorId\x12\x18\n" +
	"\amodelId\x18\x04 \x01(\tR\amodelId\"H\n" +
	"\x16GetBookReviewsResponse\x12.\n" +
	"\areviews\x18\x01 \x03(\v2\x14.reviews.ReviewModelR\areviews\"H\n" +
	"\x18CreateBookReviewResponse\x12,\n" +
	"\x06review\x18\x01 \x01(\v2\x14.reviews.ReviewModelR\x06review\"\x82\x02\n" +
	"\vReviewModel\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x12\x18\n" +
	"\acreated\x18\x03 \x01(\x03R\acreated\x12,\n" +
	"\acreator\x18\x04 \x01(\v2\x12.reviews.UserModelR\acreator\x12\x16\n" +
	"\x06rating\x18\x05 \x01(\x02R\x06rating\x127\n" +
	"\n" +
	"userAction\x18\x06 \x01(\x0e2\x17.reviews.UserActionEnumR\n" +
	"userAction\x12\x18\n" +
	"\aupvotes\x18\a \x01(\x05R\aupvotes\x12\x1c\n" +
	"\tdownvotes\x18\b \x01(\x05R\tdownvotes\"1\n" +
	"\tUserModel\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login*5\n" +
	"\x0eUserActionEnum\x12\f\n" +
	"\bnoAction\x10\x00\x12\b\n" +
	"\x04like\x10\x01\x12\v\n" +
	"\adislike\x10\x022\xba\x02\n" +
	"\x06Review\x12\x93\x01\n" +
	"\x10CreateBookReview\x12 .reviews.CreateBookReviewRequest\x1a!.reviews.CreateBookReviewResponse\":\x92A\x19\x12\x17Adds new review to book\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/api/v1/review/book\x12\x99\x01\n" +
	"\x0eGetBookReviews\x12\x1e.reviews.GetBookReviewsRequest\x1a\x1f.reviews.GetBookReviewsResponse\"F\x92A(\x12&Get reviews of book with provided {id}\x82\xd3\xe4\x93\x02\x15\x12\x13/api/v1/review/{id}B>Z<github.com/reversersed/LitGO-proto/gen/go/reviews;reviews_pbb\x06proto3"

var (
	file_reviews_review_service_proto_rawDescOnce sync.Once
	file_reviews_review_service_proto_rawDescData []byte
)

func file_reviews_review_service_proto_rawDescGZIP() []byte {
	file_reviews_review_service_proto_rawDescOnce.Do(func() {
		file_reviews_review_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_reviews_review_service_proto_rawDesc), len(file_reviews_review_service_proto_rawDesc)))
	})
	return file_reviews_review_service_proto_rawDescData
}

var file_reviews_review_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reviews_review_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reviews_review_service_proto_goTypes = []any{
	(UserActionEnum)(0),              // 0: reviews.UserActionEnum
	(*GetBookReviewsRequest)(nil),    // 1: reviews.GetBookReviewsRequest
	(*CreateBookReviewRequest)(nil),  // 2: reviews.CreateBookReviewRequest
	(*GetBookReviewsResponse)(nil),   // 3: reviews.GetBookReviewsResponse
	(*CreateBookReviewResponse)(nil), // 4: reviews.CreateBookReviewResponse
	(*ReviewModel)(nil),              // 5: reviews.ReviewModel
	(*UserModel)(nil),                // 6: reviews.UserModel
}
var file_reviews_review_service_proto_depIdxs = []int32{
	5, // 0: reviews.GetBookReviewsResponse.reviews:type_name -> reviews.ReviewModel
	5, // 1: reviews.CreateBookReviewResponse.review:type_name -> reviews.ReviewModel
	6, // 2: reviews.ReviewModel.creator:type_name -> reviews.UserModel
	0, // 3: reviews.ReviewModel.userAction:type_name -> reviews.UserActionEnum
	2, // 4: reviews.Review.CreateBookReview:input_type -> reviews.CreateBookReviewRequest
	1, // 5: reviews.Review.GetBookReviews:input_type -> reviews.GetBookReviewsRequest
	4, // 6: reviews.Review.CreateBookReview:output_type -> reviews.CreateBookReviewResponse
	3, // 7: reviews.Review.GetBookReviews:output_type -> reviews.GetBookReviewsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_reviews_review_service_proto_init() }
func file_reviews_review_service_proto_init() {
	if File_reviews_review_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_reviews_review_service_proto_rawDesc), len(file_reviews_review_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reviews_review_service_proto_goTypes,
		DependencyIndexes: file_reviews_review_service_proto_depIdxs,
		EnumInfos:         file_reviews_review_service_proto_enumTypes,
		MessageInfos:      file_reviews_review_service_proto_msgTypes,
	}.Build()
	File_reviews_review_service_proto = out.File
	file_reviews_review_service_proto_goTypes = nil
	file_reviews_review_service_proto_depIdxs = nil
}
