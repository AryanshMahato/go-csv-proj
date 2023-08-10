// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	model "github.com/AryanshMahato/go-csv-proj/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Saver is an autogenerated mock type for the Saver type
type Saver struct {
	mock.Mock
}

// SaveUsers provides a mock function with given fields: _a0
func (_m *Saver) SaveUsers(_a0 []model.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSaver creates a new instance of Saver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSaver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Saver {
	mock := &Saver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
