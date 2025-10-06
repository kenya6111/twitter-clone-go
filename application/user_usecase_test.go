package application

import (
	"context"
	"errors"
	"testing"
	"time"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"

	"go.uber.org/mock/gomock"
)

type UserUsecaseTester struct {
	userRepo          *domain.MockUserRepository
	emailVerifyRepo   *domain.MockEmailVerifyTokenRepository
	Transaction       *domain.MockTransaction
	userDomainService *domain.MockUserDomainService
	emailService      *domain.MockEmailService
	passwordHasher    *domain.MockPasswordHasher

	Usecase UserUsecase
}

func newUserUsecaseTester(ctrl *gomock.Controller) *UserUsecaseTester {
	var (
		userRepo          = domain.NewMockUserRepository(ctrl)
		emailVerifyRepo   = domain.NewMockEmailVerifyTokenRepository(ctrl)
		Transaction       = domain.NewMockTransaction(ctrl)
		userDomainService = domain.NewMockUserDomainService(ctrl)
		emailService      = domain.NewMockEmailService(ctrl)
		passwordHasher    = domain.NewMockPasswordHasher(ctrl)
		sessionStore      = domain.NewMockSessionStore(ctrl)
	)

	return &UserUsecaseTester{
		userRepo:          userRepo,
		emailVerifyRepo:   emailVerifyRepo,
		Transaction:       Transaction,
		userDomainService: userDomainService,
		emailService:      emailService,
		passwordHasher:    passwordHasher,
		Usecase:           NewUserUsecase(userRepo, emailVerifyRepo, Transaction, userDomainService, emailService, passwordHasher, sessionStore),
	}
}
func TestUserUsecaseImpl_SignUp(t *testing.T) {
	ctx := context.Background()
	password, err := domain.NewPassword("hashed_PW1!")
	if err != nil {
		panic(1)
	}
	tests := []struct {
		name       string
		input      *SignUpInfo
		mockExpect func(tester *UserUsecaseTester)
		wantErr    bool
	}{
		{
			name: "should signup user successfully",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", nil),
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error { // gomock で EXPECT().Do(ctx, gomock.Any()) だけだと、f は実行されない
							return f(ctx)
						}),
					tester.userRepo.EXPECT().CreateUser(ctx, "user1", "user1@example.com", "hashed_password").Return(&domain.User{
						ID:       "1",
						Name:     "user1",
						Email:    "user1@example.com",
						Password: password,
						IsActive: false,
					}, nil),
					tester.passwordHasher.EXPECT().GenerateSecureToken(32).Return("secure_token!!!", nil),
					tester.emailVerifyRepo.EXPECT().Save(ctx, "1", "secure_token!!!").Return(nil, nil),
					tester.emailService.EXPECT().SendInvitationEmail("user1@example.com", "secure_token!!!").Return(nil),
				)
			},
			wantErr: false,
		},
		{
			name: "should not signup when create user object fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!!!!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {},
			wantErr:    true,
		},
		{
			name: "should not signup when email is duplicated",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(errors.New(string(apperrors.DuplicateData))),
				)
			},
			wantErr: true,
		},
		{
			name: "should not signup when create hashed password fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", errors.New(string(apperrors.GenerateHashFailed))),
				)
			},
			wantErr: true,
		},
		{
			name: "should not signup when create user record fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", nil),
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.userRepo.EXPECT().CreateUser(ctx, "user1", "user1@example.com", "hashed_password").Return(&domain.User{
						ID:       "1",
						Name:     "user1",
						Email:    "user1@example.com",
						Password: password,
						IsActive: false,
					}, errors.New(string(apperrors.InsertDataFailed))),
				)
			},
			wantErr: true,
		},
		{
			name: "should not signup when create secure token fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", nil),
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.userRepo.EXPECT().CreateUser(ctx, "user1", "user1@example.com", "hashed_password").Return(&domain.User{
						ID:       "1",
						Name:     "user1",
						Email:    "user1@example.com",
						Password: password,
						IsActive: false,
					}, nil),
					tester.passwordHasher.EXPECT().GenerateSecureToken(32).Return("secure_token!!!", errors.New(string(apperrors.GenerateTokenFailed))),
				)
			},
			wantErr: true,
		},
		{
			name: "should not signup when create emailverifytoken fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", nil),
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.userRepo.EXPECT().CreateUser(ctx, "user1", "user1@example.com", "hashed_password").Return(&domain.User{
						ID:       "1",
						Name:     "user1",
						Email:    "user1@example.com",
						Password: password,
						IsActive: false,
					}, nil),
					tester.passwordHasher.EXPECT().GenerateSecureToken(32).Return("secure_token!!!", nil),
					tester.emailVerifyRepo.EXPECT().Save(ctx, "1", "secure_token!!!").Return(nil, errors.New(string(apperrors.InsertDataFailed))),
				)
			},
			wantErr: true,
		},
		{
			name: "should not signup when send invitation email fail",
			input: &SignUpInfo{
				Name:            "user1",
				Email:           "user1@example.com",
				Password:        "Password1234!",
				ConfirmPassword: "Password1234!",
			},
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.userDomainService.EXPECT().IsDuplicatedEmail(ctx, "user1@example.com").Return(nil),
					tester.passwordHasher.EXPECT().HashPassword("Password1234!").Return("hashed_password", nil),
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.userRepo.EXPECT().CreateUser(ctx, "user1", "user1@example.com", "hashed_password").Return(&domain.User{
						ID:       "1",
						Name:     "user1",
						Email:    "user1@example.com",
						Password: password,
						IsActive: false,
					}, nil),
					tester.passwordHasher.EXPECT().GenerateSecureToken(32).Return("secure_token!!!", nil),
					tester.emailVerifyRepo.EXPECT().Save(ctx, "1", "secure_token!!!").Return(nil, nil),
					tester.emailService.EXPECT().SendInvitationEmail("user1@example.com", "secure_token!!!").Return(errors.New(string(apperrors.SendEmailFailed))),
				)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tester := newUserUsecaseTester(ctrl)
			if me := tt.mockExpect; me != nil {
				me(tester)
			}
			err := tester.Usecase.SignUp(ctx, *tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestUserUsecaseImpl_Activate(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		token      string
		mockExpect func(tester *UserUsecaseTester)
		wantErr    bool
	}{
		{
			name:  "should activate user successfully",
			token: "email_verify_token",
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.emailVerifyRepo.EXPECT().FindByToken(ctx, "email_verify_token").Return(&domain.EmailVerifyToken{
						ID:        "1",
						UserID:    "1",
						Token:     "email_verify_token",
						ExpiresAt: time.Now(),
						CreatedAt: time.Now(),
					}, nil),
					tester.userRepo.EXPECT().ActivateUser(ctx, "1").Return(nil, nil),
					tester.emailVerifyRepo.EXPECT().DeleteByToken(ctx, "email_verify_token").Return(nil),
				)
			},
			wantErr: false,
		},
		{
			name:  "should not activate when FindByToken fail",
			token: "email_verify_token",
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.emailVerifyRepo.EXPECT().FindByToken(ctx, "email_verify_token").Return(nil, errors.New(string(apperrors.NoTargetData))),
				)
			},
			wantErr: true,
		},
		{
			name:  "should not activate when FindByToken result nil ",
			token: "email_verify_token",
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.emailVerifyRepo.EXPECT().FindByToken(ctx, "email_verify_token").Return(nil, nil),
				)
			},
			wantErr: true,
		},
		{
			name:  "should not activate when activate user fail",
			token: "email_verify_token",
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.emailVerifyRepo.EXPECT().FindByToken(ctx, "email_verify_token").Return(&domain.EmailVerifyToken{
						ID:        "1",
						UserID:    "1",
						Token:     "email_verify_token",
						ExpiresAt: time.Now(),
						CreatedAt: time.Now(),
					}, nil),

					tester.userRepo.EXPECT().ActivateUser(ctx, "1").Return(nil, errors.New(string(apperrors.UpdateDataFailed))),
				)
			},
			wantErr: true,
		},
		{
			name:  "should not activate when delete email_verify_token fail",
			token: "email_verify_token",
			mockExpect: func(tester *UserUsecaseTester) {
				gomock.InOrder(
					tester.Transaction.EXPECT().
						Do(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).
						DoAndReturn(func(_ context.Context, f func(context.Context) error) error {
							return f(ctx)
						}),
					tester.emailVerifyRepo.EXPECT().FindByToken(ctx, "email_verify_token").Return(&domain.EmailVerifyToken{
						ID:        "1",
						UserID:    "1",
						Token:     "email_verify_token",
						ExpiresAt: time.Now(),
						CreatedAt: time.Now(),
					}, nil),
					tester.userRepo.EXPECT().ActivateUser(ctx, "1").Return(nil, nil),
					tester.emailVerifyRepo.EXPECT().DeleteByToken(ctx, "email_verify_token").Return(errors.New(string(apperrors.DeleteDataFailed))),
				)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tester := newUserUsecaseTester(ctrl)
			if me := tt.mockExpect; me != nil {
				me(tester)
			}
			err := tester.Usecase.Activate(ctx, tt.token)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
