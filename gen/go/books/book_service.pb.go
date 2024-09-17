// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: book_service.proto

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
type GetSuggestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query to find
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"` // @gotags: form:"query" validate:"required"
	// max objects to find
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"` // @gotags: form:"limit" validate:"required,gte=1,lte=10"
}

func (x *GetSuggestionRequest) Reset() {
	*x = GetSuggestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestionRequest) ProtoMessage() {}

func (x *GetSuggestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestionRequest.ProtoReflect.Descriptor instead.
func (*GetSuggestionRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSuggestionRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetSuggestionRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Responses
type GetBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookModel `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksResponse) Reset() {
	*x = GetBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksResponse) ProtoMessage() {}

func (x *GetBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksResponse.ProtoReflect.Descriptor instead.
func (*GetBooksResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBooksResponse) GetBooks() []*BookModel {
	if x != nil {
		return x.Books
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
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[2]
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
	return file_book_service_proto_rawDescGZIP(), []int{2}
}

type BookModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Translitname string `protobuf:"bytes,3,opt,name=translitname,proto3" json:"translitname,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Picture      string `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Filepath     string `protobuf:"bytes,6,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *BookModel) Reset() {
	*x = BookModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookModel) ProtoMessage() {}

func (x *BookModel) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_book_service_proto_rawDescGZIP(), []int{3}
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

var File_book_service_proto protoreflect.FileDescriptor

var file_book_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0xab, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x32, 0x52, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x4a, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_book_service_proto_rawDescOnce sync.Once
	file_book_service_proto_rawDescData = file_book_service_proto_rawDesc
)

func file_book_service_proto_rawDescGZIP() []byte {
	file_book_service_proto_rawDescOnce.Do(func() {
		file_book_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_service_proto_rawDescData)
	})
	return file_book_service_proto_rawDescData
}

var file_book_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_book_service_proto_goTypes = []any{
	(*GetSuggestionRequest)(nil), // 0: books.GetSuggestionRequest
	(*GetBooksResponse)(nil),     // 1: books.GetBooksResponse
	(*Empty)(nil),                // 2: books.Empty
	(*BookModel)(nil),            // 3: books.BookModel
}
var file_book_service_proto_depIdxs = []int32{
	3, // 0: books.GetBooksResponse.books:type_name -> books.BookModel
	0, // 1: books.Book.GetBookSuggestions:input_type -> books.GetSuggestionRequest
	1, // 2: books.Book.GetBookSuggestions:output_type -> books.GetBooksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_book_service_proto_init() }
func file_book_service_proto_init() {
	if File_book_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetSuggestionRequest); i {
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
		file_book_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetBooksResponse); i {
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
		file_book_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_book_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BookModel); i {
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
			RawDescriptor: file_book_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_service_proto_goTypes,
		DependencyIndexes: file_book_service_proto_depIdxs,
		MessageInfos:      file_book_service_proto_msgTypes,
	}.Build()
	File_book_service_proto = out.File
	file_book_service_proto_rawDesc = nil
	file_book_service_proto_goTypes = nil
	file_book_service_proto_depIdxs = nil
}
