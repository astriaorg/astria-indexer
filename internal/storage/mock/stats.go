// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by MockGen. DO NOT EDIT.
// Source: stats.go
//
// Generated by this command:
//
//	mockgen -source=stats.go -destination=mock/stats.go -package=mock -typed
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/celenium-io/astria-indexer/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIStats is a mock of IStats interface.
type MockIStats struct {
	ctrl     *gomock.Controller
	recorder *MockIStatsMockRecorder
}

// MockIStatsMockRecorder is the mock recorder for MockIStats.
type MockIStatsMockRecorder struct {
	mock *MockIStats
}

// NewMockIStats creates a new mock instance.
func NewMockIStats(ctrl *gomock.Controller) *MockIStats {
	mock := &MockIStats{ctrl: ctrl}
	mock.recorder = &MockIStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStats) EXPECT() *MockIStatsMockRecorder {
	return m.recorder
}

// RollupSeries mocks base method.
func (m *MockIStats) RollupSeries(ctx context.Context, rollupId uint64, timeframe storage.Timeframe, name string, req storage.SeriesRequest) ([]storage.SeriesItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollupSeries", ctx, rollupId, timeframe, name, req)
	ret0, _ := ret[0].([]storage.SeriesItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollupSeries indicates an expected call of RollupSeries.
func (mr *MockIStatsMockRecorder) RollupSeries(ctx, rollupId, timeframe, name, req any) *MockIStatsRollupSeriesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollupSeries", reflect.TypeOf((*MockIStats)(nil).RollupSeries), ctx, rollupId, timeframe, name, req)
	return &MockIStatsRollupSeriesCall{Call: call}
}

// MockIStatsRollupSeriesCall wrap *gomock.Call
type MockIStatsRollupSeriesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStatsRollupSeriesCall) Return(arg0 []storage.SeriesItem, arg1 error) *MockIStatsRollupSeriesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStatsRollupSeriesCall) Do(f func(context.Context, uint64, storage.Timeframe, string, storage.SeriesRequest) ([]storage.SeriesItem, error)) *MockIStatsRollupSeriesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStatsRollupSeriesCall) DoAndReturn(f func(context.Context, uint64, storage.Timeframe, string, storage.SeriesRequest) ([]storage.SeriesItem, error)) *MockIStatsRollupSeriesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Series mocks base method.
func (m *MockIStats) Series(ctx context.Context, timeframe storage.Timeframe, name string, req storage.SeriesRequest) ([]storage.SeriesItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Series", ctx, timeframe, name, req)
	ret0, _ := ret[0].([]storage.SeriesItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Series indicates an expected call of Series.
func (mr *MockIStatsMockRecorder) Series(ctx, timeframe, name, req any) *MockIStatsSeriesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Series", reflect.TypeOf((*MockIStats)(nil).Series), ctx, timeframe, name, req)
	return &MockIStatsSeriesCall{Call: call}
}

// MockIStatsSeriesCall wrap *gomock.Call
type MockIStatsSeriesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStatsSeriesCall) Return(arg0 []storage.SeriesItem, arg1 error) *MockIStatsSeriesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStatsSeriesCall) Do(f func(context.Context, storage.Timeframe, string, storage.SeriesRequest) ([]storage.SeriesItem, error)) *MockIStatsSeriesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStatsSeriesCall) DoAndReturn(f func(context.Context, storage.Timeframe, string, storage.SeriesRequest) ([]storage.SeriesItem, error)) *MockIStatsSeriesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Summary mocks base method.
func (m *MockIStats) Summary(ctx context.Context) (storage.NetworkSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Summary", ctx)
	ret0, _ := ret[0].(storage.NetworkSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Summary indicates an expected call of Summary.
func (mr *MockIStatsMockRecorder) Summary(ctx any) *MockIStatsSummaryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Summary", reflect.TypeOf((*MockIStats)(nil).Summary), ctx)
	return &MockIStatsSummaryCall{Call: call}
}

// MockIStatsSummaryCall wrap *gomock.Call
type MockIStatsSummaryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStatsSummaryCall) Return(arg0 storage.NetworkSummary, arg1 error) *MockIStatsSummaryCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStatsSummaryCall) Do(f func(context.Context) (storage.NetworkSummary, error)) *MockIStatsSummaryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStatsSummaryCall) DoAndReturn(f func(context.Context) (storage.NetworkSummary, error)) *MockIStatsSummaryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SummaryTimeframe mocks base method.
func (m *MockIStats) SummaryTimeframe(ctx context.Context, timeframe storage.Timeframe) (storage.NetworkSummaryWithChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SummaryTimeframe", ctx, timeframe)
	ret0, _ := ret[0].(storage.NetworkSummaryWithChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SummaryTimeframe indicates an expected call of SummaryTimeframe.
func (mr *MockIStatsMockRecorder) SummaryTimeframe(ctx, timeframe any) *MockIStatsSummaryTimeframeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SummaryTimeframe", reflect.TypeOf((*MockIStats)(nil).SummaryTimeframe), ctx, timeframe)
	return &MockIStatsSummaryTimeframeCall{Call: call}
}

// MockIStatsSummaryTimeframeCall wrap *gomock.Call
type MockIStatsSummaryTimeframeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIStatsSummaryTimeframeCall) Return(arg0 storage.NetworkSummaryWithChange, arg1 error) *MockIStatsSummaryTimeframeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIStatsSummaryTimeframeCall) Do(f func(context.Context, storage.Timeframe) (storage.NetworkSummaryWithChange, error)) *MockIStatsSummaryTimeframeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIStatsSummaryTimeframeCall) DoAndReturn(f func(context.Context, storage.Timeframe) (storage.NetworkSummaryWithChange, error)) *MockIStatsSummaryTimeframeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
