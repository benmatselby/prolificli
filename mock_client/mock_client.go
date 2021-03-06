// Code generated by MockGen. DO NOT EDIT.
// Source: client/client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	client "github.com/benmatselby/prolificli/client"
	model "github.com/benmatselby/prolificli/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// CreateStudy mocks base method.
func (m *MockAPI) CreateStudy(arg0 model.CreateStudy) (*model.Study, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudy", arg0)
	ret0, _ := ret[0].(*model.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudy indicates an expected call of CreateStudy.
func (mr *MockAPIMockRecorder) CreateStudy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudy", reflect.TypeOf((*MockAPI)(nil).CreateStudy), arg0)
}

// GetEligibilityRequirements mocks base method.
func (m *MockAPI) GetEligibilityRequirements() (*client.ListRequirementsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEligibilityRequirements")
	ret0, _ := ret[0].(*client.ListRequirementsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEligibilityRequirements indicates an expected call of GetEligibilityRequirements.
func (mr *MockAPIMockRecorder) GetEligibilityRequirements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEligibilityRequirements", reflect.TypeOf((*MockAPI)(nil).GetEligibilityRequirements))
}

// GetMe mocks base method.
func (m *MockAPI) GetMe() (*client.Me, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe")
	ret0, _ := ret[0].(*client.Me)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockAPIMockRecorder) GetMe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockAPI)(nil).GetMe))
}

// GetStudies mocks base method.
func (m *MockAPI) GetStudies(status string) (*client.ListStudiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudies", status)
	ret0, _ := ret[0].(*client.ListStudiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudies indicates an expected call of GetStudies.
func (mr *MockAPIMockRecorder) GetStudies(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudies", reflect.TypeOf((*MockAPI)(nil).GetStudies), status)
}

// GetStudy mocks base method.
func (m *MockAPI) GetStudy(ID string) (*model.Study, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudy", ID)
	ret0, _ := ret[0].(*model.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudy indicates an expected call of GetStudy.
func (mr *MockAPIMockRecorder) GetStudy(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudy", reflect.TypeOf((*MockAPI)(nil).GetStudy), ID)
}

// GetSubmissions mocks base method.
func (m *MockAPI) GetSubmissions(ID string) (*client.ListSubmissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissions", ID)
	ret0, _ := ret[0].(*client.ListSubmissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissions indicates an expected call of GetSubmissions.
func (mr *MockAPIMockRecorder) GetSubmissions(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissions", reflect.TypeOf((*MockAPI)(nil).GetSubmissions), ID)
}

// TransitionStudy mocks base method.
func (m *MockAPI) TransitionStudy(ID, action string) (*client.TransitionStudyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransitionStudy", ID, action)
	ret0, _ := ret[0].(*client.TransitionStudyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransitionStudy indicates an expected call of TransitionStudy.
func (mr *MockAPIMockRecorder) TransitionStudy(ID, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransitionStudy", reflect.TypeOf((*MockAPI)(nil).TransitionStudy), ID, action)
}
