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

// CreateProject mocks base method.
func (m *MockAPI) CreateProject(workspaceID string, project model.Project) (*client.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", workspaceID, project)
	ret0, _ := ret[0].(*client.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockAPIMockRecorder) CreateProject(workspaceID, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockAPI)(nil).CreateProject), workspaceID, project)
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

// CreateWorkspace mocks base method.
func (m *MockAPI) CreateWorkspace(workspace model.Workspace) (*client.CreateWorkspacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkspace", workspace)
	ret0, _ := ret[0].(*client.CreateWorkspacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkspace indicates an expected call of CreateWorkspace.
func (mr *MockAPIMockRecorder) CreateWorkspace(workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkspace", reflect.TypeOf((*MockAPI)(nil).CreateWorkspace), workspace)
}

// DuplicateStudy mocks base method.
func (m *MockAPI) DuplicateStudy(ID string) (*model.Study, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DuplicateStudy", ID)
	ret0, _ := ret[0].(*model.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DuplicateStudy indicates an expected call of DuplicateStudy.
func (mr *MockAPIMockRecorder) DuplicateStudy(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateStudy", reflect.TypeOf((*MockAPI)(nil).DuplicateStudy), ID)
}

// GetCampaigns mocks base method.
func (m *MockAPI) GetCampaigns(workspaceID string, limit, offset int) (*client.ListCampaignsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCampaigns", workspaceID, limit, offset)
	ret0, _ := ret[0].(*client.ListCampaignsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCampaigns indicates an expected call of GetCampaigns.
func (mr *MockAPIMockRecorder) GetCampaigns(workspaceID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCampaigns", reflect.TypeOf((*MockAPI)(nil).GetCampaigns), workspaceID, limit, offset)
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

// GetEvents mocks base method.
func (m *MockAPI) GetEvents(subscriptionID string, limit, offset int) (*client.ListHookEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", subscriptionID, limit, offset)
	ret0, _ := ret[0].(*client.ListHookEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockAPIMockRecorder) GetEvents(subscriptionID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockAPI)(nil).GetEvents), subscriptionID, limit, offset)
}

// GetFilterSet mocks base method.
func (m *MockAPI) GetFilterSet(ID string) (*model.FilterSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilterSet", ID)
	ret0, _ := ret[0].(*model.FilterSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilterSet indicates an expected call of GetFilterSet.
func (mr *MockAPIMockRecorder) GetFilterSet(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilterSet", reflect.TypeOf((*MockAPI)(nil).GetFilterSet), ID)
}

// GetFilterSets mocks base method.
func (m *MockAPI) GetFilterSets(workspaceID string, limit, offset int) (*client.ListFilterSetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilterSets", workspaceID, limit, offset)
	ret0, _ := ret[0].(*client.ListFilterSetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilterSets indicates an expected call of GetFilterSets.
func (mr *MockAPIMockRecorder) GetFilterSets(workspaceID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilterSets", reflect.TypeOf((*MockAPI)(nil).GetFilterSets), workspaceID, limit, offset)
}

// GetFilters mocks base method.
func (m *MockAPI) GetFilters() (*client.ListFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilters")
	ret0, _ := ret[0].(*client.ListFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilters indicates an expected call of GetFilters.
func (mr *MockAPIMockRecorder) GetFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilters", reflect.TypeOf((*MockAPI)(nil).GetFilters))
}

// GetHookEventTypes mocks base method.
func (m *MockAPI) GetHookEventTypes() (*client.ListHookEventTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHookEventTypes")
	ret0, _ := ret[0].(*client.ListHookEventTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHookEventTypes indicates an expected call of GetHookEventTypes.
func (mr *MockAPIMockRecorder) GetHookEventTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHookEventTypes", reflect.TypeOf((*MockAPI)(nil).GetHookEventTypes))
}

// GetHookSecrets mocks base method.
func (m *MockAPI) GetHookSecrets(workspaceID string) (*client.ListSecretsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHookSecrets", workspaceID)
	ret0, _ := ret[0].(*client.ListSecretsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHookSecrets indicates an expected call of GetHookSecrets.
func (mr *MockAPIMockRecorder) GetHookSecrets(workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHookSecrets", reflect.TypeOf((*MockAPI)(nil).GetHookSecrets), workspaceID)
}

// GetHooks mocks base method.
func (m *MockAPI) GetHooks(workspaceID string, enabled bool, limit, offset int) (*client.ListHooksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHooks", workspaceID, enabled, limit, offset)
	ret0, _ := ret[0].(*client.ListHooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHooks indicates an expected call of GetHooks.
func (mr *MockAPIMockRecorder) GetHooks(workspaceID, enabled, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHooks", reflect.TypeOf((*MockAPI)(nil).GetHooks), workspaceID, enabled, limit, offset)
}

// GetMe mocks base method.
func (m *MockAPI) GetMe() (*client.MeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe")
	ret0, _ := ret[0].(*client.MeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockAPIMockRecorder) GetMe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockAPI)(nil).GetMe))
}

// GetMessages mocks base method.
func (m *MockAPI) GetMessages(userID, createdAfter *string) (*client.ListMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", userID, createdAfter)
	ret0, _ := ret[0].(*client.ListMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockAPIMockRecorder) GetMessages(userID, createdAfter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockAPI)(nil).GetMessages), userID, createdAfter)
}

// GetParticipantGroup mocks base method.
func (m *MockAPI) GetParticipantGroup(groupID string) (*client.ViewParticipantGroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantGroup", groupID)
	ret0, _ := ret[0].(*client.ViewParticipantGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantGroup indicates an expected call of GetParticipantGroup.
func (mr *MockAPIMockRecorder) GetParticipantGroup(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantGroup", reflect.TypeOf((*MockAPI)(nil).GetParticipantGroup), groupID)
}

// GetParticipantGroups mocks base method.
func (m *MockAPI) GetParticipantGroups(projectID string, limit, offset int) (*client.ListParticipantGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantGroups", projectID, limit, offset)
	ret0, _ := ret[0].(*client.ListParticipantGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantGroups indicates an expected call of GetParticipantGroups.
func (mr *MockAPIMockRecorder) GetParticipantGroups(projectID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantGroups", reflect.TypeOf((*MockAPI)(nil).GetParticipantGroups), projectID, limit, offset)
}

// GetProject mocks base method.
func (m *MockAPI) GetProject(ID string) (*model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", ID)
	ret0, _ := ret[0].(*model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockAPIMockRecorder) GetProject(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockAPI)(nil).GetProject), ID)
}

// GetProjects mocks base method.
func (m *MockAPI) GetProjects(workspaceID string, limit, offset int) (*client.ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", workspaceID, limit, offset)
	ret0, _ := ret[0].(*client.ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockAPIMockRecorder) GetProjects(workspaceID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockAPI)(nil).GetProjects), workspaceID, limit, offset)
}

// GetStudies mocks base method.
func (m *MockAPI) GetStudies(status, projectID string) (*client.ListStudiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudies", status, projectID)
	ret0, _ := ret[0].(*client.ListStudiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudies indicates an expected call of GetStudies.
func (mr *MockAPIMockRecorder) GetStudies(status, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudies", reflect.TypeOf((*MockAPI)(nil).GetStudies), status, projectID)
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
func (m *MockAPI) GetSubmissions(ID string, limit, offset int) (*client.ListSubmissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissions", ID, limit, offset)
	ret0, _ := ret[0].(*client.ListSubmissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissions indicates an expected call of GetSubmissions.
func (mr *MockAPIMockRecorder) GetSubmissions(ID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissions", reflect.TypeOf((*MockAPI)(nil).GetSubmissions), ID, limit, offset)
}

// GetUnreadMessages mocks base method.
func (m *MockAPI) GetUnreadMessages() (*client.ListUnreadMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreadMessages")
	ret0, _ := ret[0].(*client.ListUnreadMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreadMessages indicates an expected call of GetUnreadMessages.
func (mr *MockAPIMockRecorder) GetUnreadMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreadMessages", reflect.TypeOf((*MockAPI)(nil).GetUnreadMessages))
}

// GetWorkspaces mocks base method.
func (m *MockAPI) GetWorkspaces(limit, offset int) (*client.ListWorkspacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaces", limit, offset)
	ret0, _ := ret[0].(*client.ListWorkspacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaces indicates an expected call of GetWorkspaces.
func (mr *MockAPIMockRecorder) GetWorkspaces(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaces", reflect.TypeOf((*MockAPI)(nil).GetWorkspaces), limit, offset)
}

// SendMessage mocks base method.
func (m *MockAPI) SendMessage(body, recipientID, studyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", body, recipientID, studyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockAPIMockRecorder) SendMessage(body, recipientID, studyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockAPI)(nil).SendMessage), body, recipientID, studyID)
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

// UpdateStudy mocks base method.
func (m *MockAPI) UpdateStudy(ID string, study model.UpdateStudy) (*model.Study, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudy", ID, study)
	ret0, _ := ret[0].(*model.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudy indicates an expected call of UpdateStudy.
func (mr *MockAPIMockRecorder) UpdateStudy(ID, study interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudy", reflect.TypeOf((*MockAPI)(nil).UpdateStudy), ID, study)
}
