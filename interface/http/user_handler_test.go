package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"twitter-clone-go/application"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

type UserHandlerTester struct {
	userUseCase *application.MockUserUsecase
	Handler     *UserHandler
}

func newUserHandlerTester(ctrl *gomock.Controller) *UserHandlerTester {
	userUseCase := application.NewMockUserUsecase(ctrl)
	return &UserHandlerTester{
		userUseCase: userUseCase,
		Handler:     NewUserHandler(userUseCase),
	}
}

func TestSignUpHandler_Signup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*UserHandlerTester)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "successfully",
			requestBody: map[string]interface{}{
				"name":            "user1",
				"email":           "user1@example.com",
				"password":        "Password1234!",
				"confirmPassword": "Password1234!",
			},
			mockSetup: func(tester *UserHandlerTester) {
				tester.userUseCase.EXPECT().
					SignUp(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, input application.SignUpInfo) error {
						// 入力データの検証
						if input.Name != "user1" {
							t.Errorf("Expected name 'user1', got %s", input.Name)
						}
						if input.Email != "user1@example.com" {
							t.Errorf("Expected email 'user1@example.com', got %s", input.Email)
						}
						if input.Password != "Password1234!" {
							t.Errorf("Expected role 'Password1234!', got %s", input.Password)
						}
						if input.ConfirmPassword != "Password1234!" {
							t.Errorf("Expected role 'Password1234!', got %s", input.ConfirmPassword)
						}
						return nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"code":    "0",
				"message": "success",
			},
		},
		{
			name: "missing name",
			requestBody: map[string]interface{}{
				"email":           "user1@example.com",
				"password":        "Password1234!",
				"confirmPassword": "Password1234!",
			},
			mockSetup:      func(tester *UserHandlerTester) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"ErrCode": "R001",
				"Message": "bad request body",
			},
		},

		{
			name: "failed to signUp",
			requestBody: map[string]interface{}{
				"name":            "user1",
				"email":           "user1@example.com",
				"password":        "Password1234!",
				"confirmPassword": "Password1234!",
			},
			mockSetup: func(tester *UserHandlerTester) {
				tester.userUseCase.EXPECT().
					SignUp(gomock.Any(), gomock.Any()).
					Return(errors.New("internal error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"ErrCode": "U000",
				"Message": "internal process failed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tester := newUserHandlerTester(ctrl)
			tt.mockSetup(tester)

			// リクエストボディの準備
			jsonBody, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// HTTPリクエストの作成
			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			// レスポンスレコーダーの作成
			w := httptest.NewRecorder()

			// Ginコンテキストの作成
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// ハンドラーの実行
			tester.Handler.SignUp(c)

			// ステータスコードの確認
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// レスポンスボディの確認
			var responseBody map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			// レスポンスボディの比較（簡易的な比較）
			if tt.expectedBody != nil {
				expectedJSON, _ := json.Marshal(tt.expectedBody)
				actualJSON, _ := json.Marshal(responseBody)
				if string(expectedJSON) != string(actualJSON) {
					t.Errorf("Expected response body %s, got %s", string(expectedJSON), string(actualJSON))
				}
			}
		})
	}
}

func TestUserHandler_SignUp_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tester := newUserHandlerTester(ctrl)
	// 無効なJSONなので、UseCaseは呼ばれない

	// 無効なJSONリクエスト
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// ハンドラーの実行
	tester.Handler.SignUp(c)

	// ステータスコードの確認
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	// レスポンスボディの確認
	var responseBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedMessage := "bad request body"
	if responseBody["Message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, responseBody["Message"])
	}
}
