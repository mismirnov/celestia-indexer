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

// ByAddress mocks base method.
func (m *MockIRedelegation) ByAddress(ctx context.Context, addressId uint64, limit, offset int) ([]storage.Redelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByAddress", ctx, addressId, limit, offset)
	ret0, _ := ret[0].([]storage.Redelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByAddress indicates an expected call of ByAddress.
func (mr *MockIRedelegationMockRecorder) ByAddress(ctx, addressId, limit, offset any) *IRedelegationByAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByAddress", reflect.TypeOf((*MockIRedelegation)(nil).ByAddress), ctx, addressId, limit, offset)
	return &IRedelegationByAddressCall{Call: call}
}

// IRedelegationByAddressCall wrap *gomock.Call
type IRedelegationByAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationByAddressCall) Return(arg0 []storage.Redelegation, arg1 error) *IRedelegationByAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationByAddressCall) Do(f func(context.Context, uint64, int, int) ([]storage.Redelegation, error)) *IRedelegationByAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationByAddressCall) DoAndReturn(f func(context.Context, uint64, int, int) ([]storage.Redelegation, error)) *IRedelegationByAddressCall {
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
func (mr *MockIRedelegationMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IRedelegationCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIRedelegation)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IRedelegationCursorListCall{Call: call}
}

// IRedelegationCursorListCall wrap *gomock.Call
type IRedelegationCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationCursorListCall) Return(arg0 []*storage.Redelegation, arg1 error) *IRedelegationCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Redelegation, error)) *IRedelegationCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Redelegation, error)) *IRedelegationCursorListCall {
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
func (mr *MockIRedelegationMockRecorder) GetByID(ctx, id any) *IRedelegationGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIRedelegation)(nil).GetByID), ctx, id)
	return &IRedelegationGetByIDCall{Call: call}
}

// IRedelegationGetByIDCall wrap *gomock.Call
type IRedelegationGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationGetByIDCall) Return(arg0 *storage.Redelegation, arg1 error) *IRedelegationGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationGetByIDCall) Do(f func(context.Context, uint64) (*storage.Redelegation, error)) *IRedelegationGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Redelegation, error)) *IRedelegationGetByIDCall {
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
func (mr *MockIRedelegationMockRecorder) IsNoRows(err any) *IRedelegationIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIRedelegation)(nil).IsNoRows), err)
	return &IRedelegationIsNoRowsCall{Call: call}
}

// IRedelegationIsNoRowsCall wrap *gomock.Call
type IRedelegationIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationIsNoRowsCall) Return(arg0 bool) *IRedelegationIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationIsNoRowsCall) Do(f func(error) bool) *IRedelegationIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationIsNoRowsCall) DoAndReturn(f func(error) bool) *IRedelegationIsNoRowsCall {
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
func (mr *MockIRedelegationMockRecorder) LastID(ctx any) *IRedelegationLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIRedelegation)(nil).LastID), ctx)
	return &IRedelegationLastIDCall{Call: call}
}

// IRedelegationLastIDCall wrap *gomock.Call
type IRedelegationLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationLastIDCall) Return(arg0 uint64, arg1 error) *IRedelegationLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationLastIDCall) Do(f func(context.Context) (uint64, error)) *IRedelegationLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IRedelegationLastIDCall {
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
func (mr *MockIRedelegationMockRecorder) List(ctx, limit, offset, order any) *IRedelegationListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRedelegation)(nil).List), ctx, limit, offset, order)
	return &IRedelegationListCall{Call: call}
}

// IRedelegationListCall wrap *gomock.Call
type IRedelegationListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationListCall) Return(arg0 []*storage.Redelegation, arg1 error) *IRedelegationListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Redelegation, error)) *IRedelegationListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Redelegation, error)) *IRedelegationListCall {
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
func (mr *MockIRedelegationMockRecorder) Save(ctx, m any) *IRedelegationSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRedelegation)(nil).Save), ctx, m)
	return &IRedelegationSaveCall{Call: call}
}

// IRedelegationSaveCall wrap *gomock.Call
type IRedelegationSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationSaveCall) Return(arg0 error) *IRedelegationSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationSaveCall) Do(f func(context.Context, *storage.Redelegation) error) *IRedelegationSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationSaveCall) DoAndReturn(f func(context.Context, *storage.Redelegation) error) *IRedelegationSaveCall {
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
func (mr *MockIRedelegationMockRecorder) Update(ctx, m any) *IRedelegationUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRedelegation)(nil).Update), ctx, m)
	return &IRedelegationUpdateCall{Call: call}
}

// IRedelegationUpdateCall wrap *gomock.Call
type IRedelegationUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRedelegationUpdateCall) Return(arg0 error) *IRedelegationUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRedelegationUpdateCall) Do(f func(context.Context, *storage.Redelegation) error) *IRedelegationUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRedelegationUpdateCall) DoAndReturn(f func(context.Context, *storage.Redelegation) error) *IRedelegationUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
