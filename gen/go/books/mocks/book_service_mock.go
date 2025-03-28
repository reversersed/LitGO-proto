// Code generated by MockGen. DO NOT EDIT.
// Source: book_service_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=book_service_grpc.pb.go -destination=./mocks/book_service_mock.go
//

// Package mock_books_pb is a generated GoMock package.
package mock_books_pb

import (
	context "context"
	reflect "reflect"

	books "github.com/reversersed/LitGO-proto/gen/go/books"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBookClient is a mock of BookClient interface.
type MockBookClient struct {
	ctrl     *gomock.Controller
	recorder *MockBookClientMockRecorder
	isgomock struct{}
}

// MockBookClientMockRecorder is the mock recorder for MockBookClient.
type MockBookClientMockRecorder struct {
	mock *MockBookClient
}

// NewMockBookClient creates a new mock instance.
func NewMockBookClient(ctrl *gomock.Controller) *MockBookClient {
	mock := &MockBookClient{ctrl: ctrl}
	mock.recorder = &MockBookClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookClient) EXPECT() *MockBookClientMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookClient) CreateBook(ctx context.Context, in *books.CreateBookRequest, opts ...grpc.CallOption) (*books.CreateBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBook", varargs...)
	ret0, _ := ret[0].(*books.CreateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookClientMockRecorder) CreateBook(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookClient)(nil).CreateBook), varargs...)
}

// FindBook mocks base method.
func (m *MockBookClient) FindBook(ctx context.Context, in *books.FindBookRequest, opts ...grpc.CallOption) (*books.FindBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBook", varargs...)
	ret0, _ := ret[0].(*books.FindBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBook indicates an expected call of FindBook.
func (mr *MockBookClientMockRecorder) FindBook(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBook", reflect.TypeOf((*MockBookClient)(nil).FindBook), varargs...)
}

// GetBook mocks base method.
func (m *MockBookClient) GetBook(ctx context.Context, in *books.GetBookRequest, opts ...grpc.CallOption) (*books.GetBookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBook", varargs...)
	ret0, _ := ret[0].(*books.GetBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookClientMockRecorder) GetBook(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookClient)(nil).GetBook), varargs...)
}

// GetBookByGenre mocks base method.
func (m *MockBookClient) GetBookByGenre(ctx context.Context, in *books.GetBookByGenreRequest, opts ...grpc.CallOption) (*books.GetBookByGenreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBookByGenre", varargs...)
	ret0, _ := ret[0].(*books.GetBookByGenreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByGenre indicates an expected call of GetBookByGenre.
func (mr *MockBookClientMockRecorder) GetBookByGenre(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByGenre", reflect.TypeOf((*MockBookClient)(nil).GetBookByGenre), varargs...)
}

// GetBookList mocks base method.
func (m *MockBookClient) GetBookList(ctx context.Context, in *books.GetBookListRequest, opts ...grpc.CallOption) (*books.GetBookListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBookList", varargs...)
	ret0, _ := ret[0].(*books.GetBookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookList indicates an expected call of GetBookList.
func (mr *MockBookClientMockRecorder) GetBookList(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookList", reflect.TypeOf((*MockBookClient)(nil).GetBookList), varargs...)
}

// MockBookServer is a mock of BookServer interface.
type MockBookServer struct {
	ctrl     *gomock.Controller
	recorder *MockBookServerMockRecorder
	isgomock struct{}
}

// MockBookServerMockRecorder is the mock recorder for MockBookServer.
type MockBookServerMockRecorder struct {
	mock *MockBookServer
}

// NewMockBookServer creates a new mock instance.
func NewMockBookServer(ctrl *gomock.Controller) *MockBookServer {
	mock := &MockBookServer{ctrl: ctrl}
	mock.recorder = &MockBookServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServer) EXPECT() *MockBookServerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServer) CreateBook(arg0 context.Context, arg1 *books.CreateBookRequest) (*books.CreateBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(*books.CreateBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServerMockRecorder) CreateBook(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServer)(nil).CreateBook), arg0, arg1)
}

// FindBook mocks base method.
func (m *MockBookServer) FindBook(arg0 context.Context, arg1 *books.FindBookRequest) (*books.FindBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBook", arg0, arg1)
	ret0, _ := ret[0].(*books.FindBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBook indicates an expected call of FindBook.
func (mr *MockBookServerMockRecorder) FindBook(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBook", reflect.TypeOf((*MockBookServer)(nil).FindBook), arg0, arg1)
}

// GetBook mocks base method.
func (m *MockBookServer) GetBook(arg0 context.Context, arg1 *books.GetBookRequest) (*books.GetBookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0, arg1)
	ret0, _ := ret[0].(*books.GetBookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServerMockRecorder) GetBook(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookServer)(nil).GetBook), arg0, arg1)
}

// GetBookByGenre mocks base method.
func (m *MockBookServer) GetBookByGenre(arg0 context.Context, arg1 *books.GetBookByGenreRequest) (*books.GetBookByGenreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByGenre", arg0, arg1)
	ret0, _ := ret[0].(*books.GetBookByGenreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByGenre indicates an expected call of GetBookByGenre.
func (mr *MockBookServerMockRecorder) GetBookByGenre(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByGenre", reflect.TypeOf((*MockBookServer)(nil).GetBookByGenre), arg0, arg1)
}

// GetBookList mocks base method.
func (m *MockBookServer) GetBookList(arg0 context.Context, arg1 *books.GetBookListRequest) (*books.GetBookListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookList", arg0, arg1)
	ret0, _ := ret[0].(*books.GetBookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookList indicates an expected call of GetBookList.
func (mr *MockBookServerMockRecorder) GetBookList(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookList", reflect.TypeOf((*MockBookServer)(nil).GetBookList), arg0, arg1)
}

// mustEmbedUnimplementedBookServer mocks base method.
func (m *MockBookServer) mustEmbedUnimplementedBookServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServer")
}

// mustEmbedUnimplementedBookServer indicates an expected call of mustEmbedUnimplementedBookServer.
func (mr *MockBookServerMockRecorder) mustEmbedUnimplementedBookServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServer", reflect.TypeOf((*MockBookServer)(nil).mustEmbedUnimplementedBookServer))
}

// MockUnsafeBookServer is a mock of UnsafeBookServer interface.
type MockUnsafeBookServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBookServerMockRecorder
	isgomock struct{}
}

// MockUnsafeBookServerMockRecorder is the mock recorder for MockUnsafeBookServer.
type MockUnsafeBookServerMockRecorder struct {
	mock *MockUnsafeBookServer
}

// NewMockUnsafeBookServer creates a new mock instance.
func NewMockUnsafeBookServer(ctrl *gomock.Controller) *MockUnsafeBookServer {
	mock := &MockUnsafeBookServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBookServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBookServer) EXPECT() *MockUnsafeBookServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBookServer mocks base method.
func (m *MockUnsafeBookServer) mustEmbedUnimplementedBookServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServer")
}

// mustEmbedUnimplementedBookServer indicates an expected call of mustEmbedUnimplementedBookServer.
func (mr *MockUnsafeBookServerMockRecorder) mustEmbedUnimplementedBookServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServer", reflect.TypeOf((*MockUnsafeBookServer)(nil).mustEmbedUnimplementedBookServer))
}
