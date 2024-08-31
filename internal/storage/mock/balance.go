// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: balance.go
//
// Generated by this command:
//
//	mockgen -source=balance.go -destination=mock/balance.go -package=mock -typed
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/astria-indexer/internal/storage"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIBalance is a mock of IBalance interface.
type MockIBalance struct {
	ctrl     *gomock.Controller
	recorder *MockIBalanceMockRecorder
}

// MockIBalanceMockRecorder is the mock recorder for MockIBalance.
type MockIBalanceMockRecorder struct {
	mock *MockIBalance
}

// NewMockIBalance creates a new mock instance.
func NewMockIBalance(ctrl *gomock.Controller) *MockIBalance {
	mock := &MockIBalance{ctrl: ctrl}
	mock.recorder = &MockIBalanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBalance) EXPECT() *MockIBalanceMockRecorder {
	return m.recorder
}

// CursorList mocks base method.
func (m *MockIBalance) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIBalanceMockRecorder) CursorList(ctx, id, limit, order, cmp any) *MockIBalanceCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIBalance)(nil).CursorList), ctx, id, limit, order, cmp)
	return &MockIBalanceCursorListCall{Call: call}
}

// MockIBalanceCursorListCall wrap *gomock.Call
type MockIBalanceCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceCursorListCall) Return(arg0 []*storage.Balance, arg1 error) *MockIBalanceCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Balance, error)) *MockIBalanceCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Balance, error)) *MockIBalanceCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIBalance) GetByID(ctx context.Context, id uint64) (*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIBalanceMockRecorder) GetByID(ctx, id any) *MockIBalanceGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIBalance)(nil).GetByID), ctx, id)
	return &MockIBalanceGetByIDCall{Call: call}
}

// MockIBalanceGetByIDCall wrap *gomock.Call
type MockIBalanceGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceGetByIDCall) Return(arg0 *storage.Balance, arg1 error) *MockIBalanceGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceGetByIDCall) Do(f func(context.Context, uint64) (*storage.Balance, error)) *MockIBalanceGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Balance, error)) *MockIBalanceGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIBalance) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIBalanceMockRecorder) IsNoRows(err any) *MockIBalanceIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIBalance)(nil).IsNoRows), err)
	return &MockIBalanceIsNoRowsCall{Call: call}
}

// MockIBalanceIsNoRowsCall wrap *gomock.Call
type MockIBalanceIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceIsNoRowsCall) Return(arg0 bool) *MockIBalanceIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceIsNoRowsCall) Do(f func(error) bool) *MockIBalanceIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceIsNoRowsCall) DoAndReturn(f func(error) bool) *MockIBalanceIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIBalance) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIBalanceMockRecorder) LastID(ctx any) *MockIBalanceLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIBalance)(nil).LastID), ctx)
	return &MockIBalanceLastIDCall{Call: call}
}

// MockIBalanceLastIDCall wrap *gomock.Call
type MockIBalanceLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceLastIDCall) Return(arg0 uint64, arg1 error) *MockIBalanceLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceLastIDCall) Do(f func(context.Context) (uint64, error)) *MockIBalanceLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *MockIBalanceLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIBalance) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIBalanceMockRecorder) List(ctx, limit, offset, order any) *MockIBalanceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIBalance)(nil).List), ctx, limit, offset, order)
	return &MockIBalanceListCall{Call: call}
}

// MockIBalanceListCall wrap *gomock.Call
type MockIBalanceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceListCall) Return(arg0 []*storage.Balance, arg1 error) *MockIBalanceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Balance, error)) *MockIBalanceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Balance, error)) *MockIBalanceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIBalance) Save(ctx context.Context, m *storage.Balance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIBalanceMockRecorder) Save(ctx, m any) *MockIBalanceSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIBalance)(nil).Save), ctx, m)
	return &MockIBalanceSaveCall{Call: call}
}

// MockIBalanceSaveCall wrap *gomock.Call
type MockIBalanceSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceSaveCall) Return(arg0 error) *MockIBalanceSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceSaveCall) Do(f func(context.Context, *storage.Balance) error) *MockIBalanceSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceSaveCall) DoAndReturn(f func(context.Context, *storage.Balance) error) *MockIBalanceSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIBalance) Update(ctx context.Context, m *storage.Balance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIBalanceMockRecorder) Update(ctx, m any) *MockIBalanceUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIBalance)(nil).Update), ctx, m)
	return &MockIBalanceUpdateCall{Call: call}
}

// MockIBalanceUpdateCall wrap *gomock.Call
type MockIBalanceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIBalanceUpdateCall) Return(arg0 error) *MockIBalanceUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIBalanceUpdateCall) Do(f func(context.Context, *storage.Balance) error) *MockIBalanceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIBalanceUpdateCall) DoAndReturn(f func(context.Context, *storage.Balance) error) *MockIBalanceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
