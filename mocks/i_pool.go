// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	wpool "github.com/egnd/go-wpool"
	mock "github.com/stretchr/testify/mock"
)

// IPool is an autogenerated mock type for the IPool type
type IPool struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *IPool) Add(_a0 wpool.ITask) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(wpool.ITask) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *IPool) Close() {
	_m.Called()
}
