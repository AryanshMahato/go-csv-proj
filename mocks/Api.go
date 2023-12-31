// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	model "github.com/AryanshMahato/go-csv-proj/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Api is an autogenerated mock type for the Api type
type Api struct {
	mock.Mock
}

// GetUsers provides a mock function with given fields:
func (_m *Api) GetUsers() ([]model.User, error) {
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

// NewApi creates a new instance of Api. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *Api {
	mock := &Api{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
