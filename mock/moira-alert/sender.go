// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Sender)

// Package mock_moira_alert is a generated GoMock package.
package mock_moira_alert

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	moira "github.com/moira-alert/moira"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockSender) Init(arg0 map[string]string, arg1 moira.Logger, arg2 *time.Location, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockSenderMockRecorder) Init(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSender)(nil).Init), arg0, arg1, arg2, arg3)
}

// SendEvents mocks base method
func (m *MockSender) SendEvents(arg0 moira.NotificationEvents, arg1 moira.ContactData, arg2 moira.TriggerData, arg3 [][]byte, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEvents", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvents indicates an expected call of SendEvents
func (mr *MockSenderMockRecorder) SendEvents(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvents", reflect.TypeOf((*MockSender)(nil).SendEvents), arg0, arg1, arg2, arg3, arg4)
}
