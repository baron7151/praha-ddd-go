// Code generated by MockGen. DO NOT EDIT.
// Source: team_repository.go

// Package mock_domainteam is a generated GoMock package.
package domainteam

import (
	reflect "reflect"

	common "github.com/baron7151/praha-ddd-go/src/domain/common"
	gomock "github.com/golang/mock/gomock"
)

// MockITeamRepository is a mock of ITeamRepository interface.
type MockITeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITeamRepositoryMockRecorder
}

// MockITeamRepositoryMockRecorder is the mock recorder for MockITeamRepository.
type MockITeamRepositoryMockRecorder struct {
	mock *MockITeamRepository
}

// NewMockITeamRepository creates a new mock instance.
func NewMockITeamRepository(ctrl *gomock.Controller) *MockITeamRepository {
	mock := &MockITeamRepository{ctrl: ctrl}
	mock.recorder = &MockITeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITeamRepository) EXPECT() *MockITeamRepositoryMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockITeamRepository) Exists(teamName TeamName) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", teamName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockITeamRepositoryMockRecorder) Exists(teamName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockITeamRepository)(nil).Exists), teamName)
}

// FindAllTeams mocks base method.
func (m *MockITeamRepository) FindAllTeams() ([]TeamEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTeams")
	ret0, _ := ret[0].([]TeamEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTeams indicates an expected call of FindAllTeams.
func (mr *MockITeamRepositoryMockRecorder) FindAllTeams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTeams", reflect.TypeOf((*MockITeamRepository)(nil).FindAllTeams))
}

// FindByTeamId mocks base method.
func (m *MockITeamRepository) FindByTeamId(teamId common.BaseUUID) (TeamEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTeamId", teamId)
	ret0, _ := ret[0].(TeamEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTeamId indicates an expected call of FindByTeamId.
func (mr *MockITeamRepositoryMockRecorder) FindByTeamId(teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTeamId", reflect.TypeOf((*MockITeamRepository)(nil).FindByTeamId), teamId)
}

// Save mocks base method.
func (m *MockITeamRepository) Save(team TeamEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", team)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockITeamRepositoryMockRecorder) Save(team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockITeamRepository)(nil).Save), team)
}