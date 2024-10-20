// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.23.4
// source: books/book_service.proto

package books_pb

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
type FindBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query to find
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty" form:"query" validate:"required"`  
	// max objects to find
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty" form:"limit" validate:"required,gte=1,lte=50"`  
	// page to find
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty" form:"page" validate:"gte=0"`  
	// minimum rating to find
	Rating float32 `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty" form:"rating" validate:"gte=0,lte=5"`  
}

func (x *FindBookRequest) Reset() {
	*x = FindBookRequest{}
	mi := &file_books_book_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookRequest) ProtoMessage() {}

func (x *FindBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookRequest.ProtoReflect.Descriptor instead.
func (*FindBookRequest) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{0}
}

func (x *FindBookRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *FindBookRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindBookRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindBookRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book name, required, 4 <= length <= 64
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" validate:"required,min=4,max=64" form:"Name"`  
	// description, required, 16 <= length <= 1024
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" validate:"required,min=16,max=1024" form:"Description"`  
	// picture path, required
	Picture string `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty" validate:"required" swaggerignore:"true"`  
	// file path, required
	Filepath string `protobuf:"bytes,4,opt,name=filepath,proto3" json:"filepath,omitempty" validate:"required" swaggerignore:"true"`  
	// genre primitive id, required
	Genre string `protobuf:"bytes,5,opt,name=genre,proto3" json:"genre,omitempty" validate:"required,primitiveid" form:"Genre"`  
	// authors primitive id, at least one required
	Authors []string `protobuf:"bytes,6,rep,name=authors,proto3" json:"authors,omitempty" validate:"required,primitiveid" form:"Authors"`  
	Price   int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty" validate:"required,min=0" form:"Price"`     
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	mi := &file_books_book_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBookRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateBookRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *CreateBookRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *CreateBookRequest) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *CreateBookRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book's translit name or id to find
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty" form:"query" validate:"required"`  
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	mi := &file_books_book_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Responses
type FindBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookModel `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *FindBookResponse) Reset() {
	*x = FindBookResponse{}
	mi := &file_books_book_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookResponse) ProtoMessage() {}

func (x *FindBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookResponse.ProtoReflect.Descriptor instead.
func (*FindBookResponse) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{3}
}

func (x *FindBookResponse) GetBooks() []*BookModel {
	if x != nil {
		return x.Books
	}
	return nil
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *BookModel `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	mi := &file_books_book_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBookResponse) GetBook() *BookModel {
	if x != nil {
		return x.Book
	}
	return nil
}

type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *BookModel `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	mi := &file_books_book_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookResponse) GetBook() *BookModel {
	if x != nil {
		return x.Book
	}
	return nil
}

// Models
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_books_book_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[6]
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
	return file_books_book_service_proto_rawDescGZIP(), []int{6}
}

type BookModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Translitname string         `protobuf:"bytes,3,opt,name=translitname,proto3" json:"translitname,omitempty"`
	Description  string         `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Picture      string         `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Filepath     string         `protobuf:"bytes,6,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Authors      []*AuthorModel `protobuf:"bytes,7,rep,name=authors,proto3" json:"authors,omitempty"`
	Category     *CategoryModel `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Genre        *GenreModel    `protobuf:"bytes,9,opt,name=genre,proto3" json:"genre,omitempty"`
	Rating       float32        `protobuf:"fixed32,10,opt,name=rating,proto3" json:"rating,omitempty"`
	Reviews      int32          `protobuf:"varint,11,opt,name=reviews,proto3" json:"reviews,omitempty"`
	Price        int32          `protobuf:"varint,12,opt,name=price,proto3" json:"price,omitempty"`
	Published    int64          `protobuf:"varint,13,opt,name=published,proto3" json:"published,omitempty"`
}

func (x *BookModel) Reset() {
	*x = BookModel{}
	mi := &file_books_book_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookModel) ProtoMessage() {}

func (x *BookModel) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookModel.ProtoReflect.Descriptor instead.
func (*BookModel) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{7}
}

func (x *BookModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BookModel) GetTranslitname() string {
	if x != nil {
		return x.Translitname
	}
	return ""
}

func (x *BookModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookModel) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *BookModel) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *BookModel) GetAuthors() []*AuthorModel {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *BookModel) GetCategory() *CategoryModel {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *BookModel) GetGenre() *GenreModel {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *BookModel) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *BookModel) GetReviews() int32 {
	if x != nil {
		return x.Reviews
	}
	return 0
}

func (x *BookModel) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookModel) GetPublished() int64 {
	if x != nil {
		return x.Published
	}
	return 0
}

type AuthorModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Translitname   string `protobuf:"bytes,3,opt,name=translitname,proto3" json:"translitname,omitempty"`
	Profilepicture string `protobuf:"bytes,4,opt,name=profilepicture,proto3" json:"profilepicture,omitempty"`
}

func (x *AuthorModel) Reset() {
	*x = AuthorModel{}
	mi := &file_books_book_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorModel) ProtoMessage() {}

func (x *AuthorModel) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorModel.ProtoReflect.Descriptor instead.
func (*AuthorModel) Descriptor() ([]byte, []int) {
	return file_books_book_service_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorModel) GetTranslitname() string {
	if x != nil {
		return x.Translitname
	}
	return ""
}

func (x *AuthorModel) GetProfilepicture() string {
	if x != nil {
		return x.Profilepicture
	}
	return ""
}

type CategoryModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Translitname string `protobuf:"bytes,3,opt,name=translitname,proto3" json:"translitname,omitempty"`
}

func (x *CategoryModel) Reset() {
	*x = CategoryModel{}
	mi := &file_books_book_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryModel) ProtoMessage() {}

func (x *CategoryModel) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[9]
	if x != nil {
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
	return file_books_book_service_proto_rawDescGZIP(), []int{9}
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

func (x *CategoryModel) GetTranslitname() string {
	if x != nil {
		return x.Translitname
	}
	return ""
}

type GenreModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Translitname string `protobuf:"bytes,3,opt,name=translitname,proto3" json:"translitname,omitempty"`
}

func (x *GenreModel) Reset() {
	*x = GenreModel{}
	mi := &file_books_book_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenreModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreModel) ProtoMessage() {}

func (x *GenreModel) ProtoReflect() protoreflect.Message {
	mi := &file_books_book_service_proto_msgTypes[10]
	if x != nil {
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
	return file_books_book_service_proto_rawDescGZIP(), []int{10}
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

func (x *GenreModel) GetTranslitname() string {
	if x != nil {
		return x.Translitname
	}
	return ""
}

var File_books_book_service_proto protoreflect.FileDescriptor

var file_books_book_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x22, 0x69, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xc5, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x3a, 0x0a, 0x10,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x9a, 0x03, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x22, 0x7d, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0a, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xc0, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x72, 0x73, 0x65, 0x64, 0x2f, 0x4c,
	0x69, 0x74, 0x47, 0x4f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_book_service_proto_rawDescOnce sync.Once
	file_books_book_service_proto_rawDescData = file_books_book_service_proto_rawDesc
)

func file_books_book_service_proto_rawDescGZIP() []byte {
	file_books_book_service_proto_rawDescOnce.Do(func() {
		file_books_book_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_book_service_proto_rawDescData)
	})
	return file_books_book_service_proto_rawDescData
}

var file_books_book_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_books_book_service_proto_goTypes = []any{
	(*FindBookRequest)(nil),    // 0: books.FindBookRequest
	(*CreateBookRequest)(nil),  // 1: books.CreateBookRequest
	(*GetBookRequest)(nil),     // 2: books.GetBookRequest
	(*FindBookResponse)(nil),   // 3: books.FindBookResponse
	(*CreateBookResponse)(nil), // 4: books.CreateBookResponse
	(*GetBookResponse)(nil),    // 5: books.GetBookResponse
	(*Empty)(nil),              // 6: books.Empty
	(*BookModel)(nil),          // 7: books.BookModel
	(*AuthorModel)(nil),        // 8: books.AuthorModel
	(*CategoryModel)(nil),      // 9: books.CategoryModel
	(*GenreModel)(nil),         // 10: books.GenreModel
}
var file_books_book_service_proto_depIdxs = []int32{
	7,  // 0: books.FindBookResponse.books:type_name -> books.BookModel
	7,  // 1: books.CreateBookResponse.book:type_name -> books.BookModel
	7,  // 2: books.GetBookResponse.book:type_name -> books.BookModel
	8,  // 3: books.BookModel.authors:type_name -> books.AuthorModel
	9,  // 4: books.BookModel.category:type_name -> books.CategoryModel
	10, // 5: books.BookModel.genre:type_name -> books.GenreModel
	0,  // 6: books.Book.FindBook:input_type -> books.FindBookRequest
	1,  // 7: books.Book.CreateBook:input_type -> books.CreateBookRequest
	2,  // 8: books.Book.GetBook:input_type -> books.GetBookRequest
	3,  // 9: books.Book.FindBook:output_type -> books.FindBookResponse
	4,  // 10: books.Book.CreateBook:output_type -> books.CreateBookResponse
	5,  // 11: books.Book.GetBook:output_type -> books.GetBookResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_books_book_service_proto_init() }
func file_books_book_service_proto_init() {
	if File_books_book_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_book_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_book_service_proto_goTypes,
		DependencyIndexes: file_books_book_service_proto_depIdxs,
		MessageInfos:      file_books_book_service_proto_msgTypes,
	}.Build()
	File_books_book_service_proto = out.File
	file_books_book_service_proto_rawDesc = nil
	file_books_book_service_proto_goTypes = nil
	file_books_book_service_proto_depIdxs = nil
}
