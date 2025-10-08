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

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(c context.Context, name string, email string, hash string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", c, name, email, hash)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(c, name, email, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), c, name, email, hash)
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

// ActivateUser mocks base method.
func (m *MockUserRepository) ActivateUser(ctx context.Context, userId string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateUser", ctx, userId)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateUser indicates an expected call of ActivateUser.
func (mr *MockUserRepositoryMockRecorder) ActivateUser(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateUser", reflect.TypeOf((*MockUserRepository)(nil).ActivateUser), ctx, userId)
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

// MockEmailVerifyTokenRepository is a mock of EmailVerifyTokenRepository interface.
type MockEmailVerifyTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmailVerifyTokenRepositoryMockRecorder
	isgomock struct{}
}

// MockEmailVerifyTokenRepositoryMockRecorder is the mock recorder for MockEmailVerifyTokenRepository.
type MockEmailVerifyTokenRepositoryMockRecorder struct {
	mock *MockEmailVerifyTokenRepository
}

// NewMockEmailVerifyTokenRepository creates a new mock instance.
func NewMockEmailVerifyTokenRepository(ctrl *gomock.Controller) *MockEmailVerifyTokenRepository {
	mock := &MockEmailVerifyTokenRepository{ctrl: ctrl}
	mock.recorder = &MockEmailVerifyTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailVerifyTokenRepository) EXPECT() *MockEmailVerifyTokenRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockEmailVerifyTokenRepository) Save(ctx context.Context, userId, token string) (*EmailVerifyToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, userId, token)
	ret0, _ := ret[0].(*EmailVerifyToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockEmailVerifyTokenRepositoryMockRecorder) Save(ctx, userId, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockEmailVerifyTokenRepository)(nil).Save), ctx, userId, token)
}

// DeleteByToken mocks base method.
func (m *MockEmailVerifyTokenRepository) DeleteByToken(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByToken indicates an expected call of DeleteByToken.
func (mr *MockEmailVerifyTokenRepositoryMockRecorder) DeleteByToken(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByToken", reflect.TypeOf((*MockEmailVerifyTokenRepository)(nil).DeleteByToken), ctx, token)
}

func (m *MockEmailVerifyTokenRepository) FindByToken(ctx context.Context, token string) (*EmailVerifyToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", ctx, token)
	ret0, _ := ret[0].(*EmailVerifyToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockEmailVerifyTokenRepositoryMockRecorder) FindByToken(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockEmailVerifyTokenRepository)(nil).FindByToken), ctx, token)
}

// MockSessionStore is a mock of SessionStore interface.
type MockSessionStore struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStoreMockRecorder
	isgomock struct{}
}

// MockSessionStoreMockRecorder is the mock recorder for MockSessionStore.
type MockSessionStoreMockRecorder struct {
	mock *MockSessionStore
}

// NewMockSessionStore creates a new mock instance.
func NewMockSessionStore(ctrl *gomock.Controller) *MockSessionStore {
	mock := &MockSessionStore{ctrl: ctrl}
	mock.recorder = &MockSessionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionStore) EXPECT() *MockSessionStoreMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockSessionStore) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockSessionStoreMockRecorder) Clear(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockSessionStore)(nil).Clear), ctx)
}

// Delete mocks base method.
func (m *MockSessionStore) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSessionStoreMockRecorder) Delete(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionStore)(nil).Delete), ctx)
}

// Get mocks base method.
func (m *MockSessionStore) Get(ctx context.Context, key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSessionStoreMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionStore)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockSessionStore) Set(ctx context.Context, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSessionStoreMockRecorder) Set(ctx, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSessionStore)(nil).Set), ctx, value)
}

// MockSessionMiddleware is a mock of SessionMiddleware interface.
type MockSessionMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMiddlewareMockRecorder
	isgomock struct{}
}

// MockSessionMiddlewareMockRecorder is the mock recorder for MockSessionMiddleware.
type MockSessionMiddlewareMockRecorder struct {
	mock *MockSessionMiddleware
}

// NewMockSessionMiddleware creates a new mock instance.
func NewMockSessionMiddleware(ctrl *gomock.Controller) *MockSessionMiddleware {
	mock := &MockSessionMiddleware{ctrl: ctrl}
	mock.recorder = &MockSessionMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionMiddleware) EXPECT() *MockSessionMiddlewareMockRecorder {
	return m.recorder
}

// GetMiddleware mocks base method.
func (m *MockSessionMiddleware) GetMiddleware(sessionName string) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMiddleware", sessionName)
	ret0, _ := ret[0].(any)
	return ret0
}

// GetMiddleware indicates an expected call of GetMiddleware.
func (mr *MockSessionMiddlewareMockRecorder) GetMiddleware(sessionName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMiddleware", reflect.TypeOf((*MockSessionMiddleware)(nil).GetMiddleware), sessionName)
}
