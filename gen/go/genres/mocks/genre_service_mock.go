// Code generated by MockGen. DO NOT EDIT.
// Source: genre_service_grpc.pb.go

// Package mock_genres_pb is a generated GoMock package.
package mock_genres_pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	genres_pb "github.com/reversersed/LitGO-proto/gen/go/genres"
	grpc "google.golang.org/grpc"
)

// MockGenreClient is a mock of GenreClient interface.
type MockGenreClient struct {
	ctrl     *gomock.Controller
	recorder *MockGenreClientMockRecorder
}

// MockGenreClientMockRecorder is the mock recorder for MockGenreClient.
type MockGenreClientMockRecorder struct {
	mock *MockGenreClient
}

// NewMockGenreClient creates a new mock instance.
func NewMockGenreClient(ctrl *gomock.Controller) *MockGenreClient {
	mock := &MockGenreClient{ctrl: ctrl}
	mock.recorder = &MockGenreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenreClient) EXPECT() *MockGenreClientMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockGenreClient) GetAll(ctx context.Context, in *genres_pb.Empty, opts ...grpc.CallOption) (*genres_pb.GetAllResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAll", varargs...)
	ret0, _ := ret[0].(*genres_pb.GetAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockGenreClientMockRecorder) GetAll(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGenreClient)(nil).GetAll), varargs...)
}

// GetOneOf mocks base method.
func (m *MockGenreClient) GetOneOf(ctx context.Context, in *genres_pb.GetOneOfRequest, opts ...grpc.CallOption) (*genres_pb.GetCategoryOrGenreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOneOf", varargs...)
	ret0, _ := ret[0].(*genres_pb.GetCategoryOrGenreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneOf indicates an expected call of GetOneOf.
func (mr *MockGenreClientMockRecorder) GetOneOf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneOf", reflect.TypeOf((*MockGenreClient)(nil).GetOneOf), varargs...)
}

// GetTree mocks base method.
func (m *MockGenreClient) GetTree(ctx context.Context, in *genres_pb.GetOneOfRequest, opts ...grpc.CallOption) (*genres_pb.CategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTree", varargs...)
	ret0, _ := ret[0].(*genres_pb.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockGenreClientMockRecorder) GetTree(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockGenreClient)(nil).GetTree), varargs...)
}

// MockGenreServer is a mock of GenreServer interface.
type MockGenreServer struct {
	ctrl     *gomock.Controller
	recorder *MockGenreServerMockRecorder
}

// MockGenreServerMockRecorder is the mock recorder for MockGenreServer.
type MockGenreServerMockRecorder struct {
	mock *MockGenreServer
}

// NewMockGenreServer creates a new mock instance.
func NewMockGenreServer(ctrl *gomock.Controller) *MockGenreServer {
	mock := &MockGenreServer{ctrl: ctrl}
	mock.recorder = &MockGenreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenreServer) EXPECT() *MockGenreServerMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockGenreServer) GetAll(arg0 context.Context, arg1 *genres_pb.Empty) (*genres_pb.GetAllResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].(*genres_pb.GetAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockGenreServerMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGenreServer)(nil).GetAll), arg0, arg1)
}

// GetOneOf mocks base method.
func (m *MockGenreServer) GetOneOf(arg0 context.Context, arg1 *genres_pb.GetOneOfRequest) (*genres_pb.GetCategoryOrGenreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneOf", arg0, arg1)
	ret0, _ := ret[0].(*genres_pb.GetCategoryOrGenreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneOf indicates an expected call of GetOneOf.
func (mr *MockGenreServerMockRecorder) GetOneOf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneOf", reflect.TypeOf((*MockGenreServer)(nil).GetOneOf), arg0, arg1)
}

// GetTree mocks base method.
func (m *MockGenreServer) GetTree(arg0 context.Context, arg1 *genres_pb.GetOneOfRequest) (*genres_pb.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", arg0, arg1)
	ret0, _ := ret[0].(*genres_pb.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockGenreServerMockRecorder) GetTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockGenreServer)(nil).GetTree), arg0, arg1)
}

// mustEmbedUnimplementedGenreServer mocks base method.
func (m *MockGenreServer) mustEmbedUnimplementedGenreServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGenreServer")
}

// mustEmbedUnimplementedGenreServer indicates an expected call of mustEmbedUnimplementedGenreServer.
func (mr *MockGenreServerMockRecorder) mustEmbedUnimplementedGenreServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGenreServer", reflect.TypeOf((*MockGenreServer)(nil).mustEmbedUnimplementedGenreServer))
}

// MockUnsafeGenreServer is a mock of UnsafeGenreServer interface.
type MockUnsafeGenreServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeGenreServerMockRecorder
}

// MockUnsafeGenreServerMockRecorder is the mock recorder for MockUnsafeGenreServer.
type MockUnsafeGenreServerMockRecorder struct {
	mock *MockUnsafeGenreServer
}

// NewMockUnsafeGenreServer creates a new mock instance.
func NewMockUnsafeGenreServer(ctrl *gomock.Controller) *MockUnsafeGenreServer {
	mock := &MockUnsafeGenreServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeGenreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeGenreServer) EXPECT() *MockUnsafeGenreServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedGenreServer mocks base method.
func (m *MockUnsafeGenreServer) mustEmbedUnimplementedGenreServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGenreServer")
}

// mustEmbedUnimplementedGenreServer indicates an expected call of mustEmbedUnimplementedGenreServer.
func (mr *MockUnsafeGenreServerMockRecorder) mustEmbedUnimplementedGenreServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGenreServer", reflect.TypeOf((*MockUnsafeGenreServer)(nil).mustEmbedUnimplementedGenreServer))
}