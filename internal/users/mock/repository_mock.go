// Code generated by MockGen. DO NOT EDIT.
// Source: internal/users/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	entities "github.com/DoWithLogic/go-echo-realworld/internal/users/entities"
	repository "github.com/DoWithLogic/go-echo-realworld/internal/users/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Atomic mocks base method.
func (m *MockRepository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(repository.Repository) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Atomic", ctx, opt, repo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Atomic indicates an expected call of Atomic.
func (mr *MockRepositoryMockRecorder) Atomic(ctx, opt, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Atomic", reflect.TypeOf((*MockRepository)(nil).Atomic), ctx, opt, repo)
}

// GetUserByEmail mocks base method.
func (m *MockRepository) GetUserByEmail(arg0 context.Context, arg1 string) (entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockRepositoryMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockRepository)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockRepository) GetUserByID(arg0 context.Context, arg1 int64, arg2 ...entities.LockingOpt) (entities.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByID", varargs...)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockRepositoryMockRecorder) GetUserByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRepository)(nil).GetUserByID), varargs...)
}

// GetUserProfile mocks base method.
func (m *MockRepository) GetUserProfile(ctx context.Context, userName string) (entities.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", ctx, userName)
	ret0, _ := ret[0].(entities.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockRepositoryMockRecorder) GetUserProfile(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockRepository)(nil).GetUserProfile), ctx, userName)
}

// IsUserExist mocks base method.
func (m *MockRepository) IsUserExist(ctx context.Context, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExist", ctx, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserExist indicates an expected call of IsUserExist.
func (mr *MockRepositoryMockRecorder) IsUserExist(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExist", reflect.TypeOf((*MockRepository)(nil).IsUserExist), ctx, email)
}

// IsUserFollowed mocks base method.
func (m *MockRepository) IsUserFollowed(ctx context.Context, user_id, follow_user_id int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserFollowed", ctx, user_id, follow_user_id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserFollowed indicates an expected call of IsUserFollowed.
func (mr *MockRepositoryMockRecorder) IsUserFollowed(ctx, user_id, follow_user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserFollowed", reflect.TypeOf((*MockRepository)(nil).IsUserFollowed), ctx, user_id, follow_user_id)
}

// SaveNewProfile mocks base method.
func (m *MockRepository) SaveNewProfile(ctx context.Context, req entities.Profile) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNewProfile", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveNewProfile indicates an expected call of SaveNewProfile.
func (mr *MockRepositoryMockRecorder) SaveNewProfile(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNewProfile", reflect.TypeOf((*MockRepository)(nil).SaveNewProfile), ctx, req)
}

// SaveNewUser mocks base method.
func (m *MockRepository) SaveNewUser(arg0 context.Context, arg1 entities.Users) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNewUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveNewUser indicates an expected call of SaveNewUser.
func (mr *MockRepositoryMockRecorder) SaveNewUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNewUser", reflect.TypeOf((*MockRepository)(nil).SaveNewUser), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockRepository) UpdateUser(arg0 context.Context, arg1 entities.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), arg0, arg1)
}
