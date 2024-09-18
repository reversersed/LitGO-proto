// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: genres/genre_service.proto

package genres_pb

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

// Requests
type GetOneOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query to find. can be hex id or translit name
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty" form:"query" validate:"required"`  
}

func (x *GetOneOfRequest) Reset() {
	*x = GetOneOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneOfRequest) ProtoMessage() {}

func (x *GetOneOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneOfRequest.ProtoReflect.Descriptor instead.
func (*GetOneOfRequest) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetOneOfRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Responses
type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*CategoryModel `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetCategories() []*CategoryModel {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *CategoryModel `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryResponse) GetCategory() *CategoryModel {
	if x != nil {
		return x.Category
	}
	return nil
}

type GetCategoryOrGenreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Model:
	//
	//	*GetCategoryOrGenreResponse_Category
	//	*GetCategoryOrGenreResponse_Genre
	Model isGetCategoryOrGenreResponse_Model `protobuf_oneof:"model"`
}

func (x *GetCategoryOrGenreResponse) Reset() {
	*x = GetCategoryOrGenreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryOrGenreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryOrGenreResponse) ProtoMessage() {}

func (x *GetCategoryOrGenreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryOrGenreResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryOrGenreResponse) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{3}
}

func (m *GetCategoryOrGenreResponse) GetModel() isGetCategoryOrGenreResponse_Model {
	if m != nil {
		return m.Model
	}
	return nil
}

func (x *GetCategoryOrGenreResponse) GetCategory() *CategoryModel {
	if x, ok := x.GetModel().(*GetCategoryOrGenreResponse_Category); ok {
		return x.Category
	}
	return nil
}

func (x *GetCategoryOrGenreResponse) GetGenre() *GenreModel {
	if x, ok := x.GetModel().(*GetCategoryOrGenreResponse_Genre); ok {
		return x.Genre
	}
	return nil
}

type isGetCategoryOrGenreResponse_Model interface {
	isGetCategoryOrGenreResponse_Model()
}

type GetCategoryOrGenreResponse_Category struct {
	Category *CategoryModel `protobuf:"bytes,1,opt,name=category,proto3,oneof"`
}

type GetCategoryOrGenreResponse_Genre struct {
	Genre *GenreModel `protobuf:"bytes,2,opt,name=genre,proto3,oneof"`
}

func (*GetCategoryOrGenreResponse_Category) isGetCategoryOrGenreResponse_Model() {}

func (*GetCategoryOrGenreResponse_Genre) isGetCategoryOrGenreResponse_Model() {}

// Models
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_genres_genre_service_proto_rawDescGZIP(), []int{4}
}

type CategoryModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TranslitName string        `protobuf:"bytes,3,opt,name=translitName,proto3" json:"translitName,omitempty"`
	Genres       []*GenreModel `protobuf:"bytes,4,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *CategoryModel) Reset() {
	*x = CategoryModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryModel) ProtoMessage() {}

func (x *CategoryModel) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryModel.ProtoReflect.Descriptor instead.
func (*CategoryModel) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryModel) GetTranslitName() string {
	if x != nil {
		return x.TranslitName
	}
	return ""
}

func (x *CategoryModel) GetGenres() []*GenreModel {
	if x != nil {
		return x.Genres
	}
	return nil
}

type GenreModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TranslitName string `protobuf:"bytes,3,opt,name=translitName,proto3" json:"translitName,omitempty"`
	BookCount    int64  `protobuf:"varint,4,opt,name=bookCount,proto3" json:"bookCount,omitempty"`
}

func (x *GenreModel) Reset() {
	*x = GenreModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genres_genre_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreModel) ProtoMessage() {}

func (x *GenreModel) ProtoReflect() protoreflect.Message {
	mi := &file_genres_genre_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreModel.ProtoReflect.Descriptor instead.
func (*GenreModel) Descriptor() ([]byte, []int) {
	return file_genres_genre_service_proto_rawDescGZIP(), []int{6}
}

func (x *GenreModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenreModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenreModel) GetTranslitName() string {
	if x != nil {
		return x.TranslitName
	}
	return ""
}

func (x *GenreModel) GetBookCount() int64 {
	if x != nil {
		return x.BookCount
	}
	return 0
}

var File_genres_genre_service_proto protoreflect.FileDescriptor

var file_genres_genre_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x47, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x86, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4f, 0x72, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x2a, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x83, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x72, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbf, 0x01, 0x0a, 0x05, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0d, 0x2e,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66,
	0x12, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4f, 0x72,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x64, 0x2f, 0x4c, 0x69, 0x74, 0x47, 0x4f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x3b,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_genres_genre_service_proto_rawDescOnce sync.Once
	file_genres_genre_service_proto_rawDescData = file_genres_genre_service_proto_rawDesc
)

func file_genres_genre_service_proto_rawDescGZIP() []byte {
	file_genres_genre_service_proto_rawDescOnce.Do(func() {
		file_genres_genre_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_genres_genre_service_proto_rawDescData)
	})
	return file_genres_genre_service_proto_rawDescData
}

var file_genres_genre_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_genres_genre_service_proto_goTypes = []any{
	(*GetOneOfRequest)(nil),            // 0: genres.GetOneOfRequest
	(*GetAllResponse)(nil),             // 1: genres.GetAllResponse
	(*CategoryResponse)(nil),           // 2: genres.CategoryResponse
	(*GetCategoryOrGenreResponse)(nil), // 3: genres.GetCategoryOrGenreResponse
	(*Empty)(nil),                      // 4: genres.Empty
	(*CategoryModel)(nil),              // 5: genres.CategoryModel
	(*GenreModel)(nil),                 // 6: genres.GenreModel
}
var file_genres_genre_service_proto_depIdxs = []int32{
	5, // 0: genres.GetAllResponse.categories:type_name -> genres.CategoryModel
	5, // 1: genres.CategoryResponse.category:type_name -> genres.CategoryModel
	5, // 2: genres.GetCategoryOrGenreResponse.category:type_name -> genres.CategoryModel
	6, // 3: genres.GetCategoryOrGenreResponse.genre:type_name -> genres.GenreModel
	6, // 4: genres.CategoryModel.genres:type_name -> genres.GenreModel
	4, // 5: genres.Genre.GetAll:input_type -> genres.Empty
	0, // 6: genres.Genre.GetOneOf:input_type -> genres.GetOneOfRequest
	0, // 7: genres.Genre.GetTree:input_type -> genres.GetOneOfRequest
	1, // 8: genres.Genre.GetAll:output_type -> genres.GetAllResponse
	3, // 9: genres.Genre.GetOneOf:output_type -> genres.GetCategoryOrGenreResponse
	2, // 10: genres.Genre.GetTree:output_type -> genres.CategoryResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_genres_genre_service_proto_init() }
func file_genres_genre_service_proto_init() {
	if File_genres_genre_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genres_genre_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOneOfRequest); i {
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
		file_genres_genre_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllResponse); i {
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
		file_genres_genre_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryResponse); i {
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
		file_genres_genre_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCategoryOrGenreResponse); i {
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
		file_genres_genre_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_genres_genre_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryModel); i {
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
		file_genres_genre_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GenreModel); i {
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
	file_genres_genre_service_proto_msgTypes[3].OneofWrappers = []any{
		(*GetCategoryOrGenreResponse_Category)(nil),
		(*GetCategoryOrGenreResponse_Genre)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_genres_genre_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_genres_genre_service_proto_goTypes,
		DependencyIndexes: file_genres_genre_service_proto_depIdxs,
		MessageInfos:      file_genres_genre_service_proto_msgTypes,
	}.Build()
	File_genres_genre_service_proto = out.File
	file_genres_genre_service_proto_rawDesc = nil
	file_genres_genre_service_proto_goTypes = nil
	file_genres_genre_service_proto_depIdxs = nil
}
