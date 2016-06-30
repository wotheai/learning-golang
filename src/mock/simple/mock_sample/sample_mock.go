// Automatically generated by MockGen. DO NOT EDIT!
// Source: sample.go

package mock_sample

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Sample interface
type MockSample struct {
	ctrl     *gomock.Controller
	recorder *_MockSampleRecorder
}

// Recorder for MockSample (not exported)
type _MockSampleRecorder struct {
	mock *MockSample
}

func NewMockSample(ctrl *gomock.Controller) *MockSample {
	mock := &MockSample{ctrl: ctrl}
	mock.recorder = &_MockSampleRecorder{mock}
	return mock
}

func (_m *MockSample) EXPECT() *_MockSampleRecorder {
	return _m.recorder
}

func (_m *MockSample) Method(str string) int {
	ret := _m.ctrl.Call(_m, "Method", str)
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockSampleRecorder) Method(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Method", arg0)
}