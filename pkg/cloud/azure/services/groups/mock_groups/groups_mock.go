/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources/resourcesapi (interfaces: GroupsClientAPI)

// Package mock_groups is a generated GoMock package.
package mock_groups

import (
	context "context"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGroupsClientAPI is a mock of GroupsClientAPI interface
type MockGroupsClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockGroupsClientAPIMockRecorder
}

// MockGroupsClientAPIMockRecorder is the mock recorder for MockGroupsClientAPI
type MockGroupsClientAPIMockRecorder struct {
	mock *MockGroupsClientAPI
}

// NewMockGroupsClientAPI creates a new mock instance
func NewMockGroupsClientAPI(ctrl *gomock.Controller) *MockGroupsClientAPI {
	mock := &MockGroupsClientAPI{ctrl: ctrl}
	mock.recorder = &MockGroupsClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupsClientAPI) EXPECT() *MockGroupsClientAPIMockRecorder {
	return m.recorder
}

// CheckExistence mocks base method
func (m *MockGroupsClientAPI) CheckExistence(arg0 context.Context, arg1 string) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistence", arg0, arg1)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistence indicates an expected call of CheckExistence
func (mr *MockGroupsClientAPIMockRecorder) CheckExistence(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistence", reflect.TypeOf((*MockGroupsClientAPI)(nil).CheckExistence), arg0, arg1)
}

// CreateOrUpdate mocks base method
func (m *MockGroupsClientAPI) CreateOrUpdate(arg0 context.Context, arg1 string, arg2 resources.Group) (resources.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockGroupsClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockGroupsClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockGroupsClientAPI) Delete(arg0 context.Context, arg1 string) (resources.GroupsDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(resources.GroupsDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockGroupsClientAPIMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupsClientAPI)(nil).Delete), arg0, arg1)
}

// ExportTemplate mocks base method
func (m *MockGroupsClientAPI) ExportTemplate(arg0 context.Context, arg1 string, arg2 resources.ExportTemplateRequest) (resources.GroupExportResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.GroupExportResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportTemplate indicates an expected call of ExportTemplate
func (mr *MockGroupsClientAPIMockRecorder) ExportTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportTemplate", reflect.TypeOf((*MockGroupsClientAPI)(nil).ExportTemplate), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockGroupsClientAPI) Get(arg0 context.Context, arg1 string) (resources.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(resources.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGroupsClientAPIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGroupsClientAPI)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockGroupsClientAPI) List(arg0 context.Context, arg1 string, arg2 *int32) (resources.GroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.GroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGroupsClientAPIMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGroupsClientAPI)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockGroupsClientAPI) Update(arg0 context.Context, arg1 string, arg2 resources.GroupPatchable) (resources.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGroupsClientAPIMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupsClientAPI)(nil).Update), arg0, arg1, arg2)
}
