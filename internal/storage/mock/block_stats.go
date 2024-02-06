// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT


// Code generated by MockGen. DO NOT EDIT.
// Source: block_stats.go
//
// Generated by this command:
//
//	mockgen -source=block_stats.go -destination=mock/block_stats.go -package=mock -typed
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/celestia-indexer/internal/storage"
	types "github.com/celenium-io/celestia-indexer/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockIBlockStats is a mock of IBlockStats interface.
type MockIBlockStats struct {
	ctrl     *gomock.Controller
	recorder *MockIBlockStatsMockRecorder
}

// MockIBlockStatsMockRecorder is the mock recorder for MockIBlockStats.
type MockIBlockStatsMockRecorder struct {
	mock *MockIBlockStats
}

// NewMockIBlockStats creates a new mock instance.
func NewMockIBlockStats(ctrl *gomock.Controller) *MockIBlockStats {
	mock := &MockIBlockStats{ctrl: ctrl}
	mock.recorder = &MockIBlockStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBlockStats) EXPECT() *MockIBlockStatsMockRecorder {
	return m.recorder
}

// ByHeight mocks base method.
func (m *MockIBlockStats) ByHeight(ctx context.Context, height types.Level) (storage.BlockStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByHeight", ctx, height)
	ret0, _ := ret[0].(storage.BlockStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByHeight indicates an expected call of ByHeight.
func (mr *MockIBlockStatsMockRecorder) ByHeight(ctx, height any) *IBlockStatsByHeightCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByHeight", reflect.TypeOf((*MockIBlockStats)(nil).ByHeight), ctx, height)
	return &IBlockStatsByHeightCall{Call: call}
}

// IBlockStatsByHeightCall wrap *gomock.Call
type IBlockStatsByHeightCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBlockStatsByHeightCall) Return(arg0 storage.BlockStats, arg1 error) *IBlockStatsByHeightCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBlockStatsByHeightCall) Do(f func(context.Context, types.Level) (storage.BlockStats, error)) *IBlockStatsByHeightCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBlockStatsByHeightCall) DoAndReturn(f func(context.Context, types.Level) (storage.BlockStats, error)) *IBlockStatsByHeightCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastFrom mocks base method.
func (m *MockIBlockStats) LastFrom(ctx context.Context, head types.Level, limit int) ([]storage.BlockStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastFrom", ctx, head, limit)
	ret0, _ := ret[0].([]storage.BlockStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastFrom indicates an expected call of LastFrom.
func (mr *MockIBlockStatsMockRecorder) LastFrom(ctx, head, limit any) *IBlockStatsLastFromCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastFrom", reflect.TypeOf((*MockIBlockStats)(nil).LastFrom), ctx, head, limit)
	return &IBlockStatsLastFromCall{Call: call}
}

// IBlockStatsLastFromCall wrap *gomock.Call
type IBlockStatsLastFromCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IBlockStatsLastFromCall) Return(arg0 []storage.BlockStats, arg1 error) *IBlockStatsLastFromCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IBlockStatsLastFromCall) Do(f func(context.Context, types.Level, int) ([]storage.BlockStats, error)) *IBlockStatsLastFromCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IBlockStatsLastFromCall) DoAndReturn(f func(context.Context, types.Level, int) ([]storage.BlockStats, error)) *IBlockStatsLastFromCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
