// Code generated by MockGen. DO NOT EDIT.
// Source: /media/formiga/code/go_ws/src/github.com/caioformiga/go_mongodb_crud_cryptovote/bo/interfaceBO.go

// mockgen -source=/media/formiga/code/go_ws/src/github.com/caioformiga/go_mongodb_crud_cryptovote/bo/interfaceBO.go -destination=/media/formiga/code/go_ws/src/github.com/caioformiga/go_mongodb_crud_cryptovote/mock/mock_interfaceBO.go -package=mock

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	gomock "github.com/golang/mock/gomock"
)

// MockInterfaceCryptoVoteBO is a mock of InterfaceCryptoVoteBO interface.
type MockInterfaceCryptoVoteBO struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceCryptoVoteBOMockRecorder
}

// MockInterfaceCryptoVoteBOMockRecorder is the mock recorder for MockInterfaceCryptoVoteBO.
type MockInterfaceCryptoVoteBOMockRecorder struct {
	mock *MockInterfaceCryptoVoteBO
}

// NewMockInterfaceCryptoVoteBO creates a new mock instance.
func NewMockInterfaceCryptoVoteBO(ctrl *gomock.Controller) *MockInterfaceCryptoVoteBO {
	mock := &MockInterfaceCryptoVoteBO{ctrl: ctrl}
	mock.recorder = &MockInterfaceCryptoVoteBOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceCryptoVoteBO) EXPECT() *MockInterfaceCryptoVoteBOMockRecorder {
	return m.recorder
}

// AddDownVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) AddDownVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDownVote", filterCryptoVote)
	ret0, _ := ret[0].(model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDownVote indicates an expected call of AddDownVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) AddDownVote(filterCryptoVote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDownVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).AddDownVote), filterCryptoVote)
}

// AddUpVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) AddUpVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUpVote", filterCryptoVote)
	ret0, _ := ret[0].(model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUpVote indicates an expected call of AddUpVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) AddUpVote(filterCryptoVote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).AddUpVote), filterCryptoVote)
}

// CreateCryptoVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) CreateCryptoVote(cryptoVote model.CryptoVote) (model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCryptoVote", cryptoVote)
	ret0, _ := ret[0].(model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCryptoVote indicates an expected call of CreateCryptoVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) CreateCryptoVote(cryptoVote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCryptoVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).CreateCryptoVote), cryptoVote)
}

// DeleteAllCryptoVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) DeleteAllCryptoVote() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllCryptoVote")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllCryptoVote indicates an expected call of DeleteAllCryptoVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) DeleteAllCryptoVote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllCryptoVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).DeleteAllCryptoVote))
}

// DeleteAllCryptoVoteByFilter mocks base method.
func (m *MockInterfaceCryptoVoteBO) DeleteAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllCryptoVoteByFilter", filterCryptoVote)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllCryptoVoteByFilter indicates an expected call of DeleteAllCryptoVoteByFilter.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) DeleteAllCryptoVoteByFilter(filterCryptoVote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllCryptoVoteByFilter", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).DeleteAllCryptoVoteByFilter), filterCryptoVote)
}

// RetrieveAllCryptoVoteByFilter mocks base method.
func (m *MockInterfaceCryptoVoteBO) RetrieveAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) ([]model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAllCryptoVoteByFilter", filterCryptoVote)
	ret0, _ := ret[0].([]model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAllCryptoVoteByFilter indicates an expected call of RetrieveAllCryptoVoteByFilter.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) RetrieveAllCryptoVoteByFilter(filterCryptoVote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAllCryptoVoteByFilter", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).RetrieveAllCryptoVoteByFilter), filterCryptoVote)
}

// RetrieveOneCryptoVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) RetrieveOneCryptoVote(name, symbol string) (model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveOneCryptoVote", name, symbol)
	ret0, _ := ret[0].(model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveOneCryptoVote indicates an expected call of RetrieveOneCryptoVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) RetrieveOneCryptoVote(name, symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveOneCryptoVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).RetrieveOneCryptoVote), name, symbol)
}

// SumaryAllCryptoVote mocks base method.
func (m *MockInterfaceCryptoVoteBO) SumaryAllCryptoVote(pageSize int64) ([]model.SumaryCryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumaryAllCryptoVote", pageSize)
	ret0, _ := ret[0].([]model.SumaryCryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumaryAllCryptoVote indicates an expected call of SumaryAllCryptoVote.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) SumaryAllCryptoVote(pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumaryAllCryptoVote", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).SumaryAllCryptoVote), pageSize)
}

// UpdateOneCryptoVoteByFilter mocks base method.
func (m *MockInterfaceCryptoVoteBO) UpdateOneCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneCryptoVoteByFilter", filterCryptoVote, cryptoNewData)
	ret0, _ := ret[0].(model.CryptoVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOneCryptoVoteByFilter indicates an expected call of UpdateOneCryptoVoteByFilter.
func (mr *MockInterfaceCryptoVoteBOMockRecorder) UpdateOneCryptoVoteByFilter(filterCryptoVote, cryptoNewData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneCryptoVoteByFilter", reflect.TypeOf((*MockInterfaceCryptoVoteBO)(nil).UpdateOneCryptoVoteByFilter), filterCryptoVote, cryptoNewData)
}
