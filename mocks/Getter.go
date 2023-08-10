// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	model "github.com/AryanshMahato/go-csv-proj/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Getter is an autogenerated mock type for the Getter type
type Getter struct {
	mock.Mock
}

// GetUsers provides a mock function with given fields:
func (_m *Getter) GetUsers() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetter creates a new instance of Getter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Getter {
	mock := &Getter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}