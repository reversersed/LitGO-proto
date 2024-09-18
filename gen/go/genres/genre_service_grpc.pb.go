// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: genres/genre_service.proto

package genres_pb

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
	Genre_GetAll_FullMethodName   = "/genres.Genre/GetAll"
	Genre_GetOneOf_FullMethodName = "/genres.Genre/GetOneOf"
	Genre_GetTree_FullMethodName  = "/genres.Genre/GetTree"
)

// GenreClient is the client API for Genre service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
//go:generate mockgen -source=genre_service_grpc.pb.go -destination=./mock/genre_service_mock.go
type GenreClient interface {
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetOneOf(ctx context.Context, in *GetOneOfRequest, opts ...grpc.CallOption) (*GetCategoryOrGenreResponse, error)
	GetTree(ctx context.Context, in *GetOneOfRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
}

type genreClient struct {
	cc grpc.ClientConnInterface
}

func NewGenreClient(cc grpc.ClientConnInterface) GenreClient {
	return &genreClient{cc}
}

func (c *genreClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, Genre_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreClient) GetOneOf(ctx context.Context, in *GetOneOfRequest, opts ...grpc.CallOption) (*GetCategoryOrGenreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryOrGenreResponse)
	err := c.cc.Invoke(ctx, Genre_GetOneOf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreClient) GetTree(ctx context.Context, in *GetOneOfRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, Genre_GetTree_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenreServer is the server API for Genre service.
// All implementations must embed UnimplementedGenreServer
// for forward compatibility.
//
//go:generate mockgen -source=genre_service_grpc.pb.go -destination=./mock/genre_service_mock.go
type GenreServer interface {
	GetAll(context.Context, *Empty) (*GetAllResponse, error)
	GetOneOf(context.Context, *GetOneOfRequest) (*GetCategoryOrGenreResponse, error)
	GetTree(context.Context, *GetOneOfRequest) (*CategoryResponse, error)
	mustEmbedUnimplementedGenreServer()
}

// UnimplementedGenreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGenreServer struct{}

func (UnimplementedGenreServer) GetAll(context.Context, *Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGenreServer) GetOneOf(context.Context, *GetOneOfRequest) (*GetCategoryOrGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneOf not implemented")
}
func (UnimplementedGenreServer) GetTree(context.Context, *GetOneOfRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTree not implemented")
}
func (UnimplementedGenreServer) mustEmbedUnimplementedGenreServer() {}
func (UnimplementedGenreServer) testEmbeddedByValue()               {}

// UnsafeGenreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenreServer will
// result in compilation errors.
type UnsafeGenreServer interface {
	mustEmbedUnimplementedGenreServer()
}

func RegisterGenreServer(s grpc.ServiceRegistrar, srv GenreServer) {
	// If the following call pancis, it indicates UnimplementedGenreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Genre_ServiceDesc, srv)
}

func _Genre_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Genre_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Genre_GetOneOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServer).GetOneOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Genre_GetOneOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServer).GetOneOf(ctx, req.(*GetOneOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Genre_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Genre_GetTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServer).GetTree(ctx, req.(*GetOneOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Genre_ServiceDesc is the grpc.ServiceDesc for Genre service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Genre_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genres.Genre",
	HandlerType: (*GenreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Genre_GetAll_Handler,
		},
		{
			MethodName: "GetOneOf",
			Handler:    _Genre_GetOneOf_Handler,
		},
		{
			MethodName: "GetTree",
			Handler:    _Genre_GetTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "genres/genre_service.proto",
}
