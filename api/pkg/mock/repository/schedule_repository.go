// Code generated by MockGen. DO NOT EDIT.
// Source: schedule_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	infrastructure "api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScheduleRepository is a mock of ScheduleRepository interface.
type MockScheduleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleRepositoryMockRecorder
}

// MockScheduleRepositoryMockRecorder is the mock recorder for MockScheduleRepository.
type MockScheduleRepositoryMockRecorder struct {
	mock *MockScheduleRepository
}

// NewMockScheduleRepository creates a new mock instance.
func NewMockScheduleRepository(ctrl *gomock.Controller) *MockScheduleRepository {
	mock := &MockScheduleRepository{ctrl: ctrl}
	mock.recorder = &MockScheduleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleRepository) EXPECT() *MockScheduleRepositoryMockRecorder {
	return m.recorder
}

// FIndNewest mocks base method.
func (m *MockScheduleRepository) FIndNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Schedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FIndNewest", db, petId)
	ret0, _ := ret[0].(*dbmodel.Schedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FIndNewest indicates an expected call of FIndNewest.
func (mr *MockScheduleRepositoryMockRecorder) FIndNewest(db, petId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FIndNewest", reflect.TypeOf((*MockScheduleRepository)(nil).FIndNewest), db, petId)
}

// Finds mocks base method.
func (m *MockScheduleRepository) Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finds", db, petId)
	ret0, _ := ret[0].([]*dbmodel.Schedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Finds indicates an expected call of Finds.
func (mr *MockScheduleRepositoryMockRecorder) Finds(db, petId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finds", reflect.TypeOf((*MockScheduleRepository)(nil).Finds), db, petId)
}

// Post mocks base method.
func (m *MockScheduleRepository) Post(db *infrastructure.RDB, petId uint64, schedule *dbmodel.Schedules) (*dbmodel.Schedules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", db, petId, schedule)
	ret0, _ := ret[0].(*dbmodel.Schedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockScheduleRepositoryMockRecorder) Post(db *infrastructure.RDB, petId uint64, schedule *dbmodel.Schedules) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockScheduleRepository)(nil).Post), db, petId, schedule)
}
