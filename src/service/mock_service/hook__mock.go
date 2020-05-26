// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/service/hook.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	thrid "github.com/teambition/urbs-console/src/dto/thrid"
	reflect "reflect"
)

// MockHookInterface is a mock of HookInterface interface
type MockHookInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHookInterfaceMockRecorder
}

// MockHookInterfaceMockRecorder is the mock recorder for MockHookInterface
type MockHookInterfaceMockRecorder struct {
	mock *MockHookInterface
}

// NewMockHookInterface creates a new mock instance
func NewMockHookInterface(ctrl *gomock.Controller) *MockHookInterface {
	mock := &MockHookInterface{ctrl: ctrl}
	mock.recorder = &MockHookInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHookInterface) EXPECT() *MockHookInterfaceMockRecorder {
	return m.recorder
}

// SendAsync mocks base method
func (m *MockHookInterface) SendAsync(ctx context.Context, body *thrid.HookSendReq) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAsync", ctx, body)
}

// SendAsync indicates an expected call of SendAsync
func (mr *MockHookInterfaceMockRecorder) SendAsync(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsync", reflect.TypeOf((*MockHookInterface)(nil).SendAsync), ctx, body)
}

// Send mocks base method
func (m *MockHookInterface) Send(ctx context.Context, body *thrid.HookSendReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockHookInterfaceMockRecorder) Send(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockHookInterface)(nil).Send), ctx, body)
}