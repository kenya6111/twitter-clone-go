package domain

import (
	"context"
	"reflect"

	"go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CountByEmail mocks base method.
func (m *MockUserRepository) CountByEmail(c context.Context, email string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByEmail", c, email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByEmail indicates an expected call of CountByEmail.
func (mr *MockUserRepositoryMockRecorder) CountByEmail(c, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByEmail", reflect.TypeOf((*MockUserRepository)(nil).CountByEmail), c, email)
}

// CreateEmailVerifyToken mocks base method.
func (m *MockUserRepository) CreateEmailVerifyToken(ctx context.Context, userId, token string) (*EmailVerifyToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmailVerifyToken", ctx, userId, token)
	ret0, _ := ret[0].(*EmailVerifyToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmailVerifyToken indicates an expected call of CreateEmailVerifyToken.
func (mr *MockUserRepositoryMockRecorder) CreateEmailVerifyToken(ctx, userId, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmailVerifyToken", reflect.TypeOf((*MockUserRepository)(nil).CreateEmailVerifyToken), ctx, userId, token)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(c context.Context, email string, hash string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", c, email, hash)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(c, email, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), c, email, hash)
}

// FindAll mocks base method.
func (m *MockUserRepository) FindAll() ([]User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepository)(nil).FindAll))
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(c context.Context, email string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", c, email)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(c, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), c, email)
}

// MockUserDomainService is a mock of UserDomainService interface.

type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
	isgomock struct{}
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// IsDuplicatedEmail mocks base method.
func (m *MockUserDomainService) IsDuplicatedEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDuplicatedEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsDuplicatedEmail indicates an expected call of IsDuplicatedEmail.
func (mr *MockUserDomainServiceMockRecorder) IsDuplicatedEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDuplicatedEmail", reflect.TypeOf((*MockUserDomainService)(nil).IsDuplicatedEmail), ctx, email)
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
	isgomock struct{}
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockTransaction) Do(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockTransactionMockRecorder) Do(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockTransaction)(nil).Do), ctx, fn)
}

// MockEmailService is a mock of EmailService interface.
type MockEmailService struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceMockRecorder
	isgomock struct{}
}

// MockEmailServiceMockRecorder is the mock recorder for MockEmailService.
type MockEmailServiceMockRecorder struct {
	mock *MockEmailService
}

// NewMockEmailService creates a new mock instance.
func NewMockEmailService(ctrl *gomock.Controller) *MockEmailService {
	mock := &MockEmailService{ctrl: ctrl}
	mock.recorder = &MockEmailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailService) EXPECT() *MockEmailServiceMockRecorder {
	return m.recorder
}

// SendInvitationEmail mocks base method.
func (m *MockEmailService) SendInvitationEmail(email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendInvitationEmail", email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendInvitationEmail indicates an expected call of SendInvitationEmail.
func (mr *MockEmailServiceMockRecorder) SendInvitationEmail(email, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendInvitationEmail", reflect.TypeOf((*MockEmailService)(nil).SendInvitationEmail), email, token)
}

type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
	isgomock struct{}
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// CompareHashAndPassword mocks base method.
func (m *MockPasswordHasher) CompareHashAndPassword(hashedPassword, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashAndPassword", hashedPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareHashAndPassword indicates an expected call of CompareHashAndPassword.
func (mr *MockPasswordHasherMockRecorder) CompareHashAndPassword(hashedPassword, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashAndPassword", reflect.TypeOf((*MockPasswordHasher)(nil).CompareHashAndPassword), hashedPassword, password)
}

// HashPassword mocks base method.
func (m *MockPasswordHasher) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockPasswordHasherMockRecorder) HashPassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockPasswordHasher)(nil).HashPassword), password)
}

// GenerateSecureToken mocks base method.
func (m *MockPasswordHasher) GenerateSecureToken(n int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSecureToken", n)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSecureToken indicates an expected call of GenerateSecureToken.
func (mr *MockPasswordHasherMockRecorder) GenerateSecureToken(n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSecureToken", reflect.TypeOf((*MockPasswordHasher)(nil).GenerateSecureToken), n)
}
