// Automatically generated by MockGen. DO NOT EDIT!
// Source: client.go

package envelope

import (
	gomock "github.com/golang/mock/gomock"
	wire "github.com/thriftrw/thriftrw-go/wire"
)

// Mock of Transport interface
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *_MockTransportRecorder
}

// Recorder for MockTransport (not exported)
type _MockTransportRecorder struct {
	mock *MockTransport
}

func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &_MockTransportRecorder{mock}
	return mock
}

func (_m *MockTransport) EXPECT() *_MockTransportRecorder {
	return _m.recorder
}

func (_m *MockTransport) Send(_param0 []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTransportRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0)
}

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Send(name string, body wire.Value) (wire.Value, error) {
	ret := _m.ctrl.Call(_m, "Send", name, body)
	ret0, _ := ret[0].(wire.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0, arg1)
}
