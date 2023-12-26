// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: rollup_provider.go
//
// Generated by this command:
//
//	mockgen -source=rollup_provider.go -destination=mock/rollup_provider.go -package=mock -typed
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

// MockIRollupProvider is a mock of IRollupProvider interface.
type MockIRollupProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIRollupProviderMockRecorder
}

// MockIRollupProviderMockRecorder is the mock recorder for MockIRollupProvider.
type MockIRollupProviderMockRecorder struct {
	mock *MockIRollupProvider
}

// NewMockIRollupProvider creates a new mock instance.
func NewMockIRollupProvider(ctrl *gomock.Controller) *MockIRollupProvider {
	mock := &MockIRollupProvider{ctrl: ctrl}
	mock.recorder = &MockIRollupProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRollupProvider) EXPECT() *MockIRollupProviderMockRecorder {
	return m.recorder
}

// CursorList mocks base method.
func (m *MockIRollupProvider) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.RollupProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.RollupProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIRollupProviderMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IRollupProviderCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIRollupProvider)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IRollupProviderCursorListCall{Call: call}
}

// IRollupProviderCursorListCall wrap *gomock.Call
type IRollupProviderCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderCursorListCall) Return(arg0 []*storage.RollupProvider, arg1 error) *IRollupProviderCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.RollupProvider, error)) *IRollupProviderCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.RollupProvider, error)) *IRollupProviderCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIRollupProvider) GetByID(ctx context.Context, id uint64) (*storage.RollupProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.RollupProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIRollupProviderMockRecorder) GetByID(ctx, id any) *IRollupProviderGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIRollupProvider)(nil).GetByID), ctx, id)
	return &IRollupProviderGetByIDCall{Call: call}
}

// IRollupProviderGetByIDCall wrap *gomock.Call
type IRollupProviderGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderGetByIDCall) Return(arg0 *storage.RollupProvider, arg1 error) *IRollupProviderGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderGetByIDCall) Do(f func(context.Context, uint64) (*storage.RollupProvider, error)) *IRollupProviderGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.RollupProvider, error)) *IRollupProviderGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIRollupProvider) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIRollupProviderMockRecorder) IsNoRows(err any) *IRollupProviderIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIRollupProvider)(nil).IsNoRows), err)
	return &IRollupProviderIsNoRowsCall{Call: call}
}

// IRollupProviderIsNoRowsCall wrap *gomock.Call
type IRollupProviderIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderIsNoRowsCall) Return(arg0 bool) *IRollupProviderIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderIsNoRowsCall) Do(f func(error) bool) *IRollupProviderIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderIsNoRowsCall) DoAndReturn(f func(error) bool) *IRollupProviderIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIRollupProvider) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIRollupProviderMockRecorder) LastID(ctx any) *IRollupProviderLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIRollupProvider)(nil).LastID), ctx)
	return &IRollupProviderLastIDCall{Call: call}
}

// IRollupProviderLastIDCall wrap *gomock.Call
type IRollupProviderLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderLastIDCall) Return(arg0 uint64, arg1 error) *IRollupProviderLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderLastIDCall) Do(f func(context.Context) (uint64, error)) *IRollupProviderLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IRollupProviderLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIRollupProvider) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.RollupProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.RollupProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIRollupProviderMockRecorder) List(ctx, limit, offset, order any) *IRollupProviderListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRollupProvider)(nil).List), ctx, limit, offset, order)
	return &IRollupProviderListCall{Call: call}
}

// IRollupProviderListCall wrap *gomock.Call
type IRollupProviderListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderListCall) Return(arg0 []*storage.RollupProvider, arg1 error) *IRollupProviderListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.RollupProvider, error)) *IRollupProviderListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.RollupProvider, error)) *IRollupProviderListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIRollupProvider) Save(ctx context.Context, m *storage.RollupProvider) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRollupProviderMockRecorder) Save(ctx, m any) *IRollupProviderSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRollupProvider)(nil).Save), ctx, m)
	return &IRollupProviderSaveCall{Call: call}
}

// IRollupProviderSaveCall wrap *gomock.Call
type IRollupProviderSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderSaveCall) Return(arg0 error) *IRollupProviderSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderSaveCall) Do(f func(context.Context, *storage.RollupProvider) error) *IRollupProviderSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderSaveCall) DoAndReturn(f func(context.Context, *storage.RollupProvider) error) *IRollupProviderSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIRollupProvider) Update(ctx context.Context, m *storage.RollupProvider) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRollupProviderMockRecorder) Update(ctx, m any) *IRollupProviderUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRollupProvider)(nil).Update), ctx, m)
	return &IRollupProviderUpdateCall{Call: call}
}

// IRollupProviderUpdateCall wrap *gomock.Call
type IRollupProviderUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupProviderUpdateCall) Return(arg0 error) *IRollupProviderUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupProviderUpdateCall) Do(f func(context.Context, *storage.RollupProvider) error) *IRollupProviderUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupProviderUpdateCall) DoAndReturn(f func(context.Context, *storage.RollupProvider) error) *IRollupProviderUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
