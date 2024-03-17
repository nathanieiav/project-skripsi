// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "project-skbackend/internal/models"

	mock "github.com/stretchr/testify/mock"

	responses "project-skbackend/internal/controllers/responses"

	utpagination "project-skbackend/packages/utils/utpagination"

	uuid "github.com/google/uuid"
)

// ICartRepository is an autogenerated mock type for the ICartRepository type
type ICartRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: m
func (_m *ICartRepository) Create(m models.Cart) (*models.Cart, error) {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Cart) (*models.Cart, error)); ok {
		return rf(m)
	}
	if rf, ok := ret.Get(0).(func(models.Cart) *models.Cart); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Cart) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: m
func (_m *ICartRepository) Delete(m models.Cart) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Cart) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: p
func (_m *ICartRepository) FindAll(p utpagination.Pagination) (*utpagination.Pagination, error) {
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

// FindByCaregiverID provides a mock function with given fields: cgid
func (_m *ICartRepository) FindByCaregiverID(cgid uuid.UUID) ([]*models.Cart, error) {
	ret := _m.Called(cgid)

	if len(ret) == 0 {
		panic("no return value specified for FindByCaregiverID")
	}

	var r0 []*models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*models.Cart, error)); ok {
		return rf(cgid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*models.Cart); ok {
		r0 = rf(cgid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(cgid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *ICartRepository) FindByID(id uuid.UUID) (*models.Cart, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Cart, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Cart); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByMealID provides a mock function with given fields: mid
func (_m *ICartRepository) FindByMealID(mid uuid.UUID) ([]*models.Cart, error) {
	ret := _m.Called(mid)

	if len(ret) == 0 {
		panic("no return value specified for FindByMealID")
	}

	var r0 []*models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*models.Cart, error)); ok {
		return rf(mid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*models.Cart); ok {
		r0 = rf(mid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(mid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByMemberID provides a mock function with given fields: mid
func (_m *ICartRepository) FindByMemberID(mid uuid.UUID) ([]*models.Cart, error) {
	ret := _m.Called(mid)

	if len(ret) == 0 {
		panic("no return value specified for FindByMemberID")
	}

	var r0 []*models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*models.Cart, error)); ok {
		return rf(mid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*models.Cart); ok {
		r0 = rf(mid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(mid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartReferenceObject provides a mock function with given fields: cart
func (_m *ICartRepository) GetCartReferenceObject(cart models.Cart) (*responses.Member, *responses.Caregiver, error) {
	ret := _m.Called(cart)

	if len(ret) == 0 {
		panic("no return value specified for GetCartReferenceObject")
	}

	var r0 *responses.Member
	var r1 *responses.Caregiver
	var r2 error
	if rf, ok := ret.Get(0).(func(models.Cart) (*responses.Member, *responses.Caregiver, error)); ok {
		return rf(cart)
	}
	if rf, ok := ret.Get(0).(func(models.Cart) *responses.Member); ok {
		r0 = rf(cart)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Cart) *responses.Caregiver); ok {
		r1 = rf(cart)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*responses.Caregiver)
		}
	}

	if rf, ok := ret.Get(2).(func(models.Cart) error); ok {
		r2 = rf(cart)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Read provides a mock function with given fields:
func (_m *ICartRepository) Read() ([]*models.Cart, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Cart, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Cart); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: m
func (_m *ICartRepository) Update(m models.Cart) (*models.Cart, error) {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Cart) (*models.Cart, error)); ok {
		return rf(m)
	}
	if rf, ok := ret.Get(0).(func(models.Cart) *models.Cart); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Cart) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICartRepository creates a new instance of ICartRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICartRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICartRepository {
	mock := &ICartRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
