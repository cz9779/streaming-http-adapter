// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import metadata "google.golang.org/grpc/metadata"
import mock "github.com/stretchr/testify/mock"
import rpc "github.com/projectriff/streaming-http-adapter/pkg/rpc"

// Riff_InvokeClient is an autogenerated mock type for the Riff_InvokeClient type
type Riff_InvokeClient struct {
	mock.Mock
}

// CloseSend provides a mock function with given fields:
func (_m *Riff_InvokeClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *Riff_InvokeClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *Riff_InvokeClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Recv provides a mock function with given fields:
func (_m *Riff_InvokeClient) Recv() (*rpc.OutputSignal, error) {
	ret := _m.Called()

	var r0 *rpc.OutputSignal
	if rf, ok := ret.Get(0).(func() *rpc.OutputSignal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.OutputSignal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvMsg provides a mock function with given fields: m
func (_m *Riff_InvokeClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *Riff_InvokeClient) Send(_a0 *rpc.InputSignal) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rpc.InputSignal) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *Riff_InvokeClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trailer provides a mock function with given fields:
func (_m *Riff_InvokeClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}
