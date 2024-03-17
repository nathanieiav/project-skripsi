// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "project-skbackend/internal/models"

	mock "github.com/stretchr/testify/mock"

	utpagination "project-skbackend/packages/utils/utpagination"

	uuid "github.com/google/uuid"
)

// IOrganizationRepository is an autogenerated mock type for the IOrganizationRepository type
type IOrganizationRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: o
func (_m *IOrganizationRepository) Create(o models.Organization) (*models.Organization, error) {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Organization) (*models.Organization, error)); ok {
		return rf(o)
	}
	if rf, ok := ret.Get(0).(func(models.Organization) *models.Organization); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Organization) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: o
func (_m *IOrganizationRepository) Delete(o models.Organization) error {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Organization) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: p
func (_m *IOrganizationRepository) FindAll(p utpagination.Pagination) (*utpagination.Pagination, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 *utpagination.Pagination
	var r1 error
	if rf, ok := ret.Get(0).(func(utpagination.Pagination) (*utpagination.Pagination, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(utpagination.Pagination) *utpagination.Pagination); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utpagination.Pagination)
		}
	}

	if rf, ok := ret.Get(1).(func(utpagination.Pagination) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: email
func (_m *IOrganizationRepository) FindByEmail(email string) (*models.Organization, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Organization, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Organization); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *IOrganizationRepository) FindByID(id uuid.UUID) (*models.Organization, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Organization, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Organization); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: uid
func (_m *IOrganizationRepository) FindByUserID(uid uuid.UUID) (*models.Organization, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for FindByUserID")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Organization, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Organization); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields:
func (_m *IOrganizationRepository) Read() ([]*models.Organization, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Organization, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Organization); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: o
func (_m *IOrganizationRepository) Update(o models.Organization) (*models.Organization, error) {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Organization) (*models.Organization, error)); ok {
		return rf(o)
	}
	if rf, ok := ret.Get(0).(func(models.Organization) *models.Organization); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Organization) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIOrganizationRepository creates a new instance of IOrganizationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrganizationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrganizationRepository {
	mock := &IOrganizationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
