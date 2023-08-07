// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_purchase is a generated GoMock package.
package mock_purchase

import (
	context "context"
	sql "database/sql"
	db "nam_0508/internal/repo/dbmodel"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPurchaseRepository is a mock of PurchaseRepository interface.
type MockPurchaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseRepositoryMockRecorder
}

// MockPurchaseRepositoryMockRecorder is the mock recorder for MockPurchaseRepository.
type MockPurchaseRepositoryMockRecorder struct {
	mock *MockPurchaseRepository
}

// NewMockPurchaseRepository creates a new mock instance.
func NewMockPurchaseRepository(ctrl *gomock.Controller) *MockPurchaseRepository {
	mock := &MockPurchaseRepository{ctrl: ctrl}
	mock.recorder = &MockPurchaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchaseRepository) EXPECT() *MockPurchaseRepositoryMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockPurchaseRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, opts)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockPurchaseRepositoryMockRecorder) BeginTx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockPurchaseRepository)(nil).BeginTx), ctx, opts)
}

// CreatePurchase mocks base method.
func (m *MockPurchaseRepository) CreatePurchase(ctx context.Context, tx *sql.Tx, purchase db.CreatePurchaseParams) (*db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePurchase", ctx, tx, purchase)
	ret0, _ := ret[0].(*db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePurchase indicates an expected call of CreatePurchase.
func (mr *MockPurchaseRepositoryMockRecorder) CreatePurchase(ctx, tx, purchase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePurchase", reflect.TypeOf((*MockPurchaseRepository)(nil).CreatePurchase), ctx, tx, purchase)
}

// MockWagerRepository is a mock of WagerRepository interface.
type MockWagerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWagerRepositoryMockRecorder
}

// MockWagerRepositoryMockRecorder is the mock recorder for MockWagerRepository.
type MockWagerRepositoryMockRecorder struct {
	mock *MockWagerRepository
}

// NewMockWagerRepository creates a new mock instance.
func NewMockWagerRepository(ctrl *gomock.Controller) *MockWagerRepository {
	mock := &MockWagerRepository{ctrl: ctrl}
	mock.recorder = &MockWagerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWagerRepository) EXPECT() *MockWagerRepositoryMockRecorder {
	return m.recorder
}

// GetWagerForUpdate mocks base method.
func (m *MockWagerRepository) GetWagerForUpdate(ctx context.Context, tx *sql.Tx, wagerID int32) (*db.Wager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWagerForUpdate", ctx, tx, wagerID)
	ret0, _ := ret[0].(*db.Wager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWagerForUpdate indicates an expected call of GetWagerForUpdate.
func (mr *MockWagerRepositoryMockRecorder) GetWagerForUpdate(ctx, tx, wagerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWagerForUpdate", reflect.TypeOf((*MockWagerRepository)(nil).GetWagerForUpdate), ctx, tx, wagerID)
}

// UpdatePurchaseWager mocks base method.
func (m *MockWagerRepository) UpdatePurchaseWager(ctx context.Context, tx *sql.Tx, updateWagerParams db.UpdatePurchaseWagerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePurchaseWager", ctx, tx, updateWagerParams)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePurchaseWager indicates an expected call of UpdatePurchaseWager.
func (mr *MockWagerRepositoryMockRecorder) UpdatePurchaseWager(ctx, tx, updateWagerParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePurchaseWager", reflect.TypeOf((*MockWagerRepository)(nil).UpdatePurchaseWager), ctx, tx, updateWagerParams)
}
