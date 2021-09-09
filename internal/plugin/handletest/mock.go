// Code generated by MockGen. DO NOT EDIT.

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
// Source: go.uber.org/thriftrw/internal/plugin (interfaces: Handle,ServiceGenerator)

// Package handletest is a generated GoMock package.
package handletest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	plugin "go.uber.org/thriftrw/internal/plugin"
	api "go.uber.org/thriftrw/plugin/api"
)

// MockHandle is a mock of Handle interface.
type MockHandle struct {
	ctrl     *gomock.Controller
	recorder *MockHandleMockRecorder
}

// MockHandleMockRecorder is the mock recorder for MockHandle.
type MockHandleMockRecorder struct {
	mock *MockHandle
}

// NewMockHandle creates a new mock instance.
func NewMockHandle(ctrl *gomock.Controller) *MockHandle {
	mock := &MockHandle{ctrl: ctrl}
	mock.recorder = &MockHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandle) EXPECT() *MockHandleMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockHandle) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockHandleMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHandle)(nil).Close))
}

// Name mocks base method.
func (m *MockHandle) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockHandleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHandle)(nil).Name))
}

// ServiceGenerator mocks base method.
func (m *MockHandle) ServiceGenerator() plugin.ServiceGenerator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceGenerator")
	ret0, _ := ret[0].(plugin.ServiceGenerator)
	return ret0
}

// ServiceGenerator indicates an expected call of ServiceGenerator.
func (mr *MockHandleMockRecorder) ServiceGenerator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceGenerator", reflect.TypeOf((*MockHandle)(nil).ServiceGenerator))
}

// MockServiceGenerator is a mock of ServiceGenerator interface.
type MockServiceGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockServiceGeneratorMockRecorder
}

// MockServiceGeneratorMockRecorder is the mock recorder for MockServiceGenerator.
type MockServiceGeneratorMockRecorder struct {
	mock *MockServiceGenerator
}

// NewMockServiceGenerator creates a new mock instance.
func NewMockServiceGenerator(ctrl *gomock.Controller) *MockServiceGenerator {
	mock := &MockServiceGenerator{ctrl: ctrl}
	mock.recorder = &MockServiceGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceGenerator) EXPECT() *MockServiceGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockServiceGenerator) Generate(arg0 *api.GenerateServiceRequest) (*api.GenerateServiceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(*api.GenerateServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockServiceGeneratorMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockServiceGenerator)(nil).Generate), arg0)
}

// Handle mocks base method.
func (m *MockServiceGenerator) Handle() plugin.Handle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(plugin.Handle)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockServiceGeneratorMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockServiceGenerator)(nil).Handle))
}
