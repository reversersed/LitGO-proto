// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: books/book_service.proto

package books_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Book_FindBook_FullMethodName        = "/books.Book/FindBook"
	Book_CreateBook_FullMethodName      = "/books.Book/CreateBook"
	Book_GetBook_FullMethodName         = "/books.Book/GetBook"
	Book_GetBookByGenre_FullMethodName  = "/books.Book/GetBookByGenre"
	Book_GetBookList_FullMethodName     = "/books.Book/GetBookList"
	Book_GetBookByAuthor_FullMethodName = "/books.Book/GetBookByAuthor"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
//go:generate mockgen -source=book_service_grpc.pb.go -destination=./mocks/book_service_mock.go
type BookClient interface {
	FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*FindBookResponse, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	GetBookByGenre(ctx context.Context, in *GetBookByGenreRequest, opts ...grpc.CallOption) (*GetBookByGenreResponse, error)
	GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error)
	GetBookByAuthor(ctx context.Context, in *GetBookByAuthorRequest, opts ...grpc.CallOption) (*GetBookByAuthorResponse, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*FindBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindBookResponse)
	err := c.cc.Invoke(ctx, Book_FindBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, Book_CreateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, Book_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookByGenre(ctx context.Context, in *GetBookByGenreRequest, opts ...grpc.CallOption) (*GetBookByGenreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookByGenreResponse)
	err := c.cc.Invoke(ctx, Book_GetBookByGenre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookListResponse)
	err := c.cc.Invoke(ctx, Book_GetBookList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookByAuthor(ctx context.Context, in *GetBookByAuthorRequest, opts ...grpc.CallOption) (*GetBookByAuthorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookByAuthorResponse)
	err := c.cc.Invoke(ctx, Book_GetBookByAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility.
//
//go:generate mockgen -source=book_service_grpc.pb.go -destination=./mocks/book_service_mock.go
type BookServer interface {
	FindBook(context.Context, *FindBookRequest) (*FindBookResponse, error)
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	GetBookByGenre(context.Context, *GetBookByGenreRequest) (*GetBookByGenreResponse, error)
	GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error)
	GetBookByAuthor(context.Context, *GetBookByAuthorRequest) (*GetBookByAuthorResponse, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServer struct{}

func (UnimplementedBookServer) FindBook(context.Context, *FindBookRequest) (*FindBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBook not implemented")
}
func (UnimplementedBookServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServer) GetBookByGenre(context.Context, *GetBookByGenreRequest) (*GetBookByGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByGenre not implemented")
}
func (UnimplementedBookServer) GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (UnimplementedBookServer) GetBookByAuthor(context.Context, *GetBookByAuthorRequest) (*GetBookByAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByAuthor not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}
func (UnimplementedBookServer) testEmbeddedByValue()              {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	// If the following call pancis, it indicates UnimplementedBookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_FindBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).FindBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_FindBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).FindBook(ctx, req.(*FindBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookByGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookByGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookByGenre(ctx, req.(*GetBookByGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookList(ctx, req.(*GetBookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookByAuthor(ctx, req.(*GetBookByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBook",
			Handler:    _Book_FindBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Book_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Book_GetBook_Handler,
		},
		{
			MethodName: "GetBookByGenre",
			Handler:    _Book_GetBookByGenre_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _Book_GetBookList_Handler,
		},
		{
			MethodName: "GetBookByAuthor",
			Handler:    _Book_GetBookByAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books/book_service.proto",
}
