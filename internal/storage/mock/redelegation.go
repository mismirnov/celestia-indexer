// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: redelegation.go
//
// Generated by this command:
//
//	mockgen -source=redelegation.go -destination=mock/redelegation.go -package=mock -typed
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/celestia-indexer/internal/storage"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIRedelegation is a mock of IRedelegation interface.
type MockIRedelegation struct {
	ctrl     *gomock.Controller
	recorder *MockIRedelegationMockRecorder
}

// MockIRedelegationMockRecorder is the mock recorder for MockIRedelegation.
type MockIRedelegationMockRecorder struct {
	mock *MockIRedelegation
}

// NewMockIRedelegation creates a new mock instance.
func NewMockIRedelegation(ctrl *gomock.Controller) *MockIRedelegation {
	mock := &MockIRedelegation{ctrl: ctrl}
	mock.recorder = &MockIRedelegationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedelegation) EXPECT() *MockIRedelegationMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockIRedelegation) ISGOMOCK() struct{} {
	return struct{}{}
}

// ByAddress mocks base method.
func (m *MockIRedelegation) ByAddress(ctx context.Context, addressId uint64, limit, offset int) ([]storage.Redelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByAddress", ctx, addressId, limit, offset)
	ret0, _ := ret[0].([]storage.Redelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByAddress indicates an expected call of ByAddress.
func (mr *MockIRedelegationMockRecorder) ByAddress(ctx, addressId, limit, offset any) *MockIRedelegationByAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByAddress", reflect.TypeOf((*MockIRedelegation)(nil).ByAddress), ctx, addressId, limit, offset)
	return &MockIRedelegationByAddressCall{Call: call}
}

// MockIRedelegationByAddressCall wrap *gomock.Call
type MockIRedelegationByAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationByAddressCall) Return(arg0 []storage.Redelegation, arg1 error) *MockIRedelegationByAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationByAddressCall) Do(f func(context.Context, uint64, int, int) ([]storage.Redelegation, error)) *MockIRedelegationByAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationByAddressCall) DoAndReturn(f func(context.Context, uint64, int, int) ([]storage.Redelegation, error)) *MockIRedelegationByAddressCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CursorList mocks base method.
func (m *MockIRedelegation) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Redelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Redelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIRedelegationMockRecorder) CursorList(ctx, id, limit, order, cmp any) *MockIRedelegationCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIRedelegation)(nil).CursorList), ctx, id, limit, order, cmp)
	return &MockIRedelegationCursorListCall{Call: call}
}

// MockIRedelegationCursorListCall wrap *gomock.Call
type MockIRedelegationCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationCursorListCall) Return(arg0 []*storage.Redelegation, arg1 error) *MockIRedelegationCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Redelegation, error)) *MockIRedelegationCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Redelegation, error)) *MockIRedelegationCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIRedelegation) GetByID(ctx context.Context, id uint64) (*storage.Redelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Redelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIRedelegationMockRecorder) GetByID(ctx, id any) *MockIRedelegationGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIRedelegation)(nil).GetByID), ctx, id)
	return &MockIRedelegationGetByIDCall{Call: call}
}

// MockIRedelegationGetByIDCall wrap *gomock.Call
type MockIRedelegationGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationGetByIDCall) Return(arg0 *storage.Redelegation, arg1 error) *MockIRedelegationGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationGetByIDCall) Do(f func(context.Context, uint64) (*storage.Redelegation, error)) *MockIRedelegationGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Redelegation, error)) *MockIRedelegationGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIRedelegation) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIRedelegationMockRecorder) IsNoRows(err any) *MockIRedelegationIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIRedelegation)(nil).IsNoRows), err)
	return &MockIRedelegationIsNoRowsCall{Call: call}
}

// MockIRedelegationIsNoRowsCall wrap *gomock.Call
type MockIRedelegationIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationIsNoRowsCall) Return(arg0 bool) *MockIRedelegationIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationIsNoRowsCall) Do(f func(error) bool) *MockIRedelegationIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationIsNoRowsCall) DoAndReturn(f func(error) bool) *MockIRedelegationIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIRedelegation) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIRedelegationMockRecorder) LastID(ctx any) *MockIRedelegationLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIRedelegation)(nil).LastID), ctx)
	return &MockIRedelegationLastIDCall{Call: call}
}

// MockIRedelegationLastIDCall wrap *gomock.Call
type MockIRedelegationLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationLastIDCall) Return(arg0 uint64, arg1 error) *MockIRedelegationLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationLastIDCall) Do(f func(context.Context) (uint64, error)) *MockIRedelegationLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *MockIRedelegationLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIRedelegation) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Redelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Redelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIRedelegationMockRecorder) List(ctx, limit, offset, order any) *MockIRedelegationListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRedelegation)(nil).List), ctx, limit, offset, order)
	return &MockIRedelegationListCall{Call: call}
}

// MockIRedelegationListCall wrap *gomock.Call
type MockIRedelegationListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationListCall) Return(arg0 []*storage.Redelegation, arg1 error) *MockIRedelegationListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Redelegation, error)) *MockIRedelegationListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Redelegation, error)) *MockIRedelegationListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIRedelegation) Save(ctx context.Context, m *storage.Redelegation) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRedelegationMockRecorder) Save(ctx, m any) *MockIRedelegationSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRedelegation)(nil).Save), ctx, m)
	return &MockIRedelegationSaveCall{Call: call}
}

// MockIRedelegationSaveCall wrap *gomock.Call
type MockIRedelegationSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationSaveCall) Return(arg0 error) *MockIRedelegationSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationSaveCall) Do(f func(context.Context, *storage.Redelegation) error) *MockIRedelegationSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationSaveCall) DoAndReturn(f func(context.Context, *storage.Redelegation) error) *MockIRedelegationSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIRedelegation) Update(ctx context.Context, m *storage.Redelegation) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRedelegationMockRecorder) Update(ctx, m any) *MockIRedelegationUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRedelegation)(nil).Update), ctx, m)
	return &MockIRedelegationUpdateCall{Call: call}
}

// MockIRedelegationUpdateCall wrap *gomock.Call
type MockIRedelegationUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIRedelegationUpdateCall) Return(arg0 error) *MockIRedelegationUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIRedelegationUpdateCall) Do(f func(context.Context, *storage.Redelegation) error) *MockIRedelegationUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIRedelegationUpdateCall) DoAndReturn(f func(context.Context, *storage.Redelegation) error) *MockIRedelegationUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
