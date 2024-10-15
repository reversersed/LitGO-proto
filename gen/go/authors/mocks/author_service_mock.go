// Code generated by MockGen. DO NOT EDIT.
// Source: author_service_grpc.pb.go

// Package mock_authors_pb is a generated GoMock package.
package mock_authors_pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authors_pb "github.com/reversersed/LitGO-proto/gen/go/authors"
	grpc "google.golang.org/grpc"
)

// MockAuthorClient is a mock of AuthorClient interface.
type MockAuthorClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorClientMockRecorder
}

// MockAuthorClientMockRecorder is the mock recorder for MockAuthorClient.
type MockAuthorClientMockRecorder struct {
	mock *MockAuthorClient
}

// NewMockAuthorClient creates a new mock instance.
func NewMockAuthorClient(ctrl *gomock.Controller) *MockAuthorClient {
	mock := &MockAuthorClient{ctrl: ctrl}
	mock.recorder = &MockAuthorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorClient) EXPECT() *MockAuthorClientMockRecorder {
	return m.recorder
}

// GetAuthorSuggestion mocks base method.
func (m *MockAuthorClient) GetAuthorSuggestion(ctx context.Context, in *authors_pb.GetSuggestionRequest, opts ...grpc.CallOption) (*authors_pb.GetAuthorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorSuggestion", varargs...)
	ret0, _ := ret[0].(*authors_pb.GetAuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorSuggestion indicates an expected call of GetAuthorSuggestion.
func (mr *MockAuthorClientMockRecorder) GetAuthorSuggestion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorSuggestion", reflect.TypeOf((*MockAuthorClient)(nil).GetAuthorSuggestion), varargs...)
}

// GetAuthors mocks base method.
func (m *MockAuthorClient) GetAuthors(ctx context.Context, in *authors_pb.GetAuthorsRequest, opts ...grpc.CallOption) (*authors_pb.GetAuthorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthors", varargs...)
	ret0, _ := ret[0].(*authors_pb.GetAuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthors indicates an expected call of GetAuthors.
func (mr *MockAuthorClientMockRecorder) GetAuthors(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockAuthorClient)(nil).GetAuthors), varargs...)
}

// MockAuthorServer is a mock of AuthorServer interface.
type MockAuthorServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorServerMockRecorder
}

// MockAuthorServerMockRecorder is the mock recorder for MockAuthorServer.
type MockAuthorServerMockRecorder struct {
	mock *MockAuthorServer
}

// NewMockAuthorServer creates a new mock instance.
func NewMockAuthorServer(ctrl *gomock.Controller) *MockAuthorServer {
	mock := &MockAuthorServer{ctrl: ctrl}
	mock.recorder = &MockAuthorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorServer) EXPECT() *MockAuthorServerMockRecorder {
	return m.recorder
}

// GetAuthorSuggestion mocks base method.
func (m *MockAuthorServer) GetAuthorSuggestion(arg0 context.Context, arg1 *authors_pb.GetSuggestionRequest) (*authors_pb.GetAuthorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorSuggestion", arg0, arg1)
	ret0, _ := ret[0].(*authors_pb.GetAuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorSuggestion indicates an expected call of GetAuthorSuggestion.
func (mr *MockAuthorServerMockRecorder) GetAuthorSuggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorSuggestion", reflect.TypeOf((*MockAuthorServer)(nil).GetAuthorSuggestion), arg0, arg1)
}

// GetAuthors mocks base method.
func (m *MockAuthorServer) GetAuthors(arg0 context.Context, arg1 *authors_pb.GetAuthorsRequest) (*authors_pb.GetAuthorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthors", arg0, arg1)
	ret0, _ := ret[0].(*authors_pb.GetAuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthors indicates an expected call of GetAuthors.
func (mr *MockAuthorServerMockRecorder) GetAuthors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockAuthorServer)(nil).GetAuthors), arg0, arg1)
}

// mustEmbedUnimplementedAuthorServer mocks base method.
func (m *MockAuthorServer) mustEmbedUnimplementedAuthorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthorServer")
}

// mustEmbedUnimplementedAuthorServer indicates an expected call of mustEmbedUnimplementedAuthorServer.
func (mr *MockAuthorServerMockRecorder) mustEmbedUnimplementedAuthorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthorServer", reflect.TypeOf((*MockAuthorServer)(nil).mustEmbedUnimplementedAuthorServer))
}

// MockUnsafeAuthorServer is a mock of UnsafeAuthorServer interface.
type MockUnsafeAuthorServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuthorServerMockRecorder
}

// MockUnsafeAuthorServerMockRecorder is the mock recorder for MockUnsafeAuthorServer.
type MockUnsafeAuthorServerMockRecorder struct {
	mock *MockUnsafeAuthorServer
}

// NewMockUnsafeAuthorServer creates a new mock instance.
func NewMockUnsafeAuthorServer(ctrl *gomock.Controller) *MockUnsafeAuthorServer {
	mock := &MockUnsafeAuthorServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuthorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthorServer) EXPECT() *MockUnsafeAuthorServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuthorServer mocks base method.
func (m *MockUnsafeAuthorServer) mustEmbedUnimplementedAuthorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthorServer")
}

// mustEmbedUnimplementedAuthorServer indicates an expected call of mustEmbedUnimplementedAuthorServer.
func (mr *MockUnsafeAuthorServerMockRecorder) mustEmbedUnimplementedAuthorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthorServer", reflect.TypeOf((*MockUnsafeAuthorServer)(nil).mustEmbedUnimplementedAuthorServer))
}
