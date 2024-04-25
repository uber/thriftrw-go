// Code generated by MockGen. DO NOT EDIT.

// Copyright (c) 2024 Uber Technologies, Inc.
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
// Source: go.uber.org/thriftrw/plugin/api (interfaces: Plugin,ServiceGenerator)

// Package plugintest is a generated GoMock package.
package plugintest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "go.uber.org/thriftrw/plugin/api"
)

// MockPlugin is a mock of Plugin interface.
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin.
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance.
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// Goodbye mocks base method.
func (m *MockPlugin) Goodbye() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Goodbye")
	ret0, _ := ret[0].(error)
	return ret0
}

// Goodbye indicates an expected call of Goodbye.
func (mr *MockPluginMockRecorder) Goodbye() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Goodbye", reflect.TypeOf((*MockPlugin)(nil).Goodbye))
}

// Handshake mocks base method.
func (m *MockPlugin) Handshake(arg0 *api.HandshakeRequest) (*api.HandshakeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handshake", arg0)
	ret0, _ := ret[0].(*api.HandshakeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handshake indicates an expected call of Handshake.
func (mr *MockPluginMockRecorder) Handshake(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handshake", reflect.TypeOf((*MockPlugin)(nil).Handshake), arg0)
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
