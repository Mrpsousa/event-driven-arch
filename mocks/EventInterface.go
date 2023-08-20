// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EventInterface is an autogenerated mock type for the EventInterface type
type EventInterface struct {
	mock.Mock
}

// GetName provides a mock function with given fields:
func (_m *EventInterface) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPayload provides a mock function with given fields:
func (_m *EventInterface) GetPayload() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// SetPayload provides a mock function with given fields: payload
func (_m *EventInterface) SetPayload(payload interface{}) {
	_m.Called(payload)
}
