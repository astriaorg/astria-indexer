// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: rollup.go
//
// Generated by this command:
//
//	mockgen -source=rollup.go -destination=mock/rollup.go -package=mock -typed
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/astria-indexer/internal/storage"
	types "github.com/celenium-io/astria-indexer/pkg/types"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIRollup is a mock of IRollup interface.
type MockIRollup struct {
	ctrl     *gomock.Controller
	recorder *MockIRollupMockRecorder
}

// MockIRollupMockRecorder is the mock recorder for MockIRollup.
type MockIRollupMockRecorder struct {
	mock *MockIRollup
}

// NewMockIRollup creates a new mock instance.
func NewMockIRollup(ctrl *gomock.Controller) *MockIRollup {
	mock := &MockIRollup{ctrl: ctrl}
	mock.recorder = &MockIRollupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRollup) EXPECT() *MockIRollupMockRecorder {
	return m.recorder
}

// ActionsByHeight mocks base method.
func (m *MockIRollup) ActionsByHeight(ctx context.Context, height types.Level, limit, offset int) ([]storage.RollupAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionsByHeight", ctx, height, limit, offset)
	ret0, _ := ret[0].([]storage.RollupAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionsByHeight indicates an expected call of ActionsByHeight.
func (mr *MockIRollupMockRecorder) ActionsByHeight(ctx, height, limit, offset any) *IRollupActionsByHeightCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionsByHeight", reflect.TypeOf((*MockIRollup)(nil).ActionsByHeight), ctx, height, limit, offset)
	return &IRollupActionsByHeightCall{Call: call}
}

// IRollupActionsByHeightCall wrap *gomock.Call
type IRollupActionsByHeightCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupActionsByHeightCall) Return(arg0 []storage.RollupAction, arg1 error) *IRollupActionsByHeightCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupActionsByHeightCall) Do(f func(context.Context, types.Level, int, int) ([]storage.RollupAction, error)) *IRollupActionsByHeightCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupActionsByHeightCall) DoAndReturn(f func(context.Context, types.Level, int, int) ([]storage.RollupAction, error)) *IRollupActionsByHeightCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ActionsByTxId mocks base method.
func (m *MockIRollup) ActionsByTxId(ctx context.Context, txId uint64, limit, offset int) ([]storage.RollupAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionsByTxId", ctx, txId, limit, offset)
	ret0, _ := ret[0].([]storage.RollupAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionsByTxId indicates an expected call of ActionsByTxId.
func (mr *MockIRollupMockRecorder) ActionsByTxId(ctx, txId, limit, offset any) *IRollupActionsByTxIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionsByTxId", reflect.TypeOf((*MockIRollup)(nil).ActionsByTxId), ctx, txId, limit, offset)
	return &IRollupActionsByTxIdCall{Call: call}
}

// IRollupActionsByTxIdCall wrap *gomock.Call
type IRollupActionsByTxIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupActionsByTxIdCall) Return(arg0 []storage.RollupAction, arg1 error) *IRollupActionsByTxIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupActionsByTxIdCall) Do(f func(context.Context, uint64, int, int) ([]storage.RollupAction, error)) *IRollupActionsByTxIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupActionsByTxIdCall) DoAndReturn(f func(context.Context, uint64, int, int) ([]storage.RollupAction, error)) *IRollupActionsByTxIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Addresses mocks base method.
func (m *MockIRollup) Addresses(ctx context.Context, rollupId uint64, limit, offset int, sort storage0.SortOrder) ([]storage.RollupAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses", ctx, rollupId, limit, offset, sort)
	ret0, _ := ret[0].([]storage.RollupAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses.
func (mr *MockIRollupMockRecorder) Addresses(ctx, rollupId, limit, offset, sort any) *IRollupAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockIRollup)(nil).Addresses), ctx, rollupId, limit, offset, sort)
	return &IRollupAddressesCall{Call: call}
}

// IRollupAddressesCall wrap *gomock.Call
type IRollupAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupAddressesCall) Return(arg0 []storage.RollupAddress, arg1 error) *IRollupAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupAddressesCall) Do(f func(context.Context, uint64, int, int, storage0.SortOrder) ([]storage.RollupAddress, error)) *IRollupAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupAddressesCall) DoAndReturn(f func(context.Context, uint64, int, int, storage0.SortOrder) ([]storage.RollupAddress, error)) *IRollupAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ByHash mocks base method.
func (m *MockIRollup) ByHash(ctx context.Context, hash []byte) (storage.Rollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByHash", ctx, hash)
	ret0, _ := ret[0].(storage.Rollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByHash indicates an expected call of ByHash.
func (mr *MockIRollupMockRecorder) ByHash(ctx, hash any) *IRollupByHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByHash", reflect.TypeOf((*MockIRollup)(nil).ByHash), ctx, hash)
	return &IRollupByHashCall{Call: call}
}

// IRollupByHashCall wrap *gomock.Call
type IRollupByHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupByHashCall) Return(arg0 storage.Rollup, arg1 error) *IRollupByHashCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupByHashCall) Do(f func(context.Context, []byte) (storage.Rollup, error)) *IRollupByHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupByHashCall) DoAndReturn(f func(context.Context, []byte) (storage.Rollup, error)) *IRollupByHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CountActionsByHeight mocks base method.
func (m *MockIRollup) CountActionsByHeight(ctx context.Context, height types.Level) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountActionsByHeight", ctx, height)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountActionsByHeight indicates an expected call of CountActionsByHeight.
func (mr *MockIRollupMockRecorder) CountActionsByHeight(ctx, height any) *IRollupCountActionsByHeightCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountActionsByHeight", reflect.TypeOf((*MockIRollup)(nil).CountActionsByHeight), ctx, height)
	return &IRollupCountActionsByHeightCall{Call: call}
}

// IRollupCountActionsByHeightCall wrap *gomock.Call
type IRollupCountActionsByHeightCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupCountActionsByHeightCall) Return(arg0 int64, arg1 error) *IRollupCountActionsByHeightCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupCountActionsByHeightCall) Do(f func(context.Context, types.Level) (int64, error)) *IRollupCountActionsByHeightCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupCountActionsByHeightCall) DoAndReturn(f func(context.Context, types.Level) (int64, error)) *IRollupCountActionsByHeightCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CountActionsByTxId mocks base method.
func (m *MockIRollup) CountActionsByTxId(ctx context.Context, txId uint64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountActionsByTxId", ctx, txId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountActionsByTxId indicates an expected call of CountActionsByTxId.
func (mr *MockIRollupMockRecorder) CountActionsByTxId(ctx, txId any) *IRollupCountActionsByTxIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountActionsByTxId", reflect.TypeOf((*MockIRollup)(nil).CountActionsByTxId), ctx, txId)
	return &IRollupCountActionsByTxIdCall{Call: call}
}

// IRollupCountActionsByTxIdCall wrap *gomock.Call
type IRollupCountActionsByTxIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupCountActionsByTxIdCall) Return(arg0 int64, arg1 error) *IRollupCountActionsByTxIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupCountActionsByTxIdCall) Do(f func(context.Context, uint64) (int64, error)) *IRollupCountActionsByTxIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupCountActionsByTxIdCall) DoAndReturn(f func(context.Context, uint64) (int64, error)) *IRollupCountActionsByTxIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CursorList mocks base method.
func (m *MockIRollup) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Rollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Rollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIRollupMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IRollupCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIRollup)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IRollupCursorListCall{Call: call}
}

// IRollupCursorListCall wrap *gomock.Call
type IRollupCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupCursorListCall) Return(arg0 []*storage.Rollup, arg1 error) *IRollupCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Rollup, error)) *IRollupCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Rollup, error)) *IRollupCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIRollup) GetByID(ctx context.Context, id uint64) (*storage.Rollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Rollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIRollupMockRecorder) GetByID(ctx, id any) *IRollupGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIRollup)(nil).GetByID), ctx, id)
	return &IRollupGetByIDCall{Call: call}
}

// IRollupGetByIDCall wrap *gomock.Call
type IRollupGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupGetByIDCall) Return(arg0 *storage.Rollup, arg1 error) *IRollupGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupGetByIDCall) Do(f func(context.Context, uint64) (*storage.Rollup, error)) *IRollupGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Rollup, error)) *IRollupGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIRollup) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIRollupMockRecorder) IsNoRows(err any) *IRollupIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIRollup)(nil).IsNoRows), err)
	return &IRollupIsNoRowsCall{Call: call}
}

// IRollupIsNoRowsCall wrap *gomock.Call
type IRollupIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupIsNoRowsCall) Return(arg0 bool) *IRollupIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupIsNoRowsCall) Do(f func(error) bool) *IRollupIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupIsNoRowsCall) DoAndReturn(f func(error) bool) *IRollupIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIRollup) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIRollupMockRecorder) LastID(ctx any) *IRollupLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIRollup)(nil).LastID), ctx)
	return &IRollupLastIDCall{Call: call}
}

// IRollupLastIDCall wrap *gomock.Call
type IRollupLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupLastIDCall) Return(arg0 uint64, arg1 error) *IRollupLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupLastIDCall) Do(f func(context.Context) (uint64, error)) *IRollupLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IRollupLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIRollup) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Rollup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Rollup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIRollupMockRecorder) List(ctx, limit, offset, order any) *IRollupListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRollup)(nil).List), ctx, limit, offset, order)
	return &IRollupListCall{Call: call}
}

// IRollupListCall wrap *gomock.Call
type IRollupListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupListCall) Return(arg0 []*storage.Rollup, arg1 error) *IRollupListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Rollup, error)) *IRollupListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Rollup, error)) *IRollupListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListRollupsByAddress mocks base method.
func (m *MockIRollup) ListRollupsByAddress(ctx context.Context, addressId uint64, limit, offset int, sort storage0.SortOrder) ([]storage.RollupAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRollupsByAddress", ctx, addressId, limit, offset, sort)
	ret0, _ := ret[0].([]storage.RollupAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRollupsByAddress indicates an expected call of ListRollupsByAddress.
func (mr *MockIRollupMockRecorder) ListRollupsByAddress(ctx, addressId, limit, offset, sort any) *IRollupListRollupsByAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRollupsByAddress", reflect.TypeOf((*MockIRollup)(nil).ListRollupsByAddress), ctx, addressId, limit, offset, sort)
	return &IRollupListRollupsByAddressCall{Call: call}
}

// IRollupListRollupsByAddressCall wrap *gomock.Call
type IRollupListRollupsByAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupListRollupsByAddressCall) Return(arg0 []storage.RollupAddress, arg1 error) *IRollupListRollupsByAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupListRollupsByAddressCall) Do(f func(context.Context, uint64, int, int, storage0.SortOrder) ([]storage.RollupAddress, error)) *IRollupListRollupsByAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupListRollupsByAddressCall) DoAndReturn(f func(context.Context, uint64, int, int, storage0.SortOrder) ([]storage.RollupAddress, error)) *IRollupListRollupsByAddressCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIRollup) Save(ctx context.Context, m *storage.Rollup) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRollupMockRecorder) Save(ctx, m any) *IRollupSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRollup)(nil).Save), ctx, m)
	return &IRollupSaveCall{Call: call}
}

// IRollupSaveCall wrap *gomock.Call
type IRollupSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupSaveCall) Return(arg0 error) *IRollupSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupSaveCall) Do(f func(context.Context, *storage.Rollup) error) *IRollupSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupSaveCall) DoAndReturn(f func(context.Context, *storage.Rollup) error) *IRollupSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIRollup) Update(ctx context.Context, m *storage.Rollup) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRollupMockRecorder) Update(ctx, m any) *IRollupUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRollup)(nil).Update), ctx, m)
	return &IRollupUpdateCall{Call: call}
}

// IRollupUpdateCall wrap *gomock.Call
type IRollupUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IRollupUpdateCall) Return(arg0 error) *IRollupUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IRollupUpdateCall) Do(f func(context.Context, *storage.Rollup) error) *IRollupUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IRollupUpdateCall) DoAndReturn(f func(context.Context, *storage.Rollup) error) *IRollupUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
