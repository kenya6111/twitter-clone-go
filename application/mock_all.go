// Package mock_application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"
	domain "twitter-clone-go/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
	isgomock struct{}
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetUserList mocks base method.
func (m *MockUserUsecase) GetUserList() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserList")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserList indicates an expected call of GetUserList.
func (mr *MockUserUsecaseMockRecorder) GetUserList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserList", reflect.TypeOf((*MockUserUsecase)(nil).GetUserList))
}

// SignUp mocks base method.
func (m *MockUserUsecase) SignUp(c context.Context, signUpInfo SignUpInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", c, signUpInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserUsecaseMockRecorder) SignUp(c, signUpInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserUsecase)(nil).SignUp), c, signUpInfo)
}
