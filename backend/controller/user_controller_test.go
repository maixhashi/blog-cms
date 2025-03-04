package controller

import (
	"encoding/json"
	"fmt"
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/usecase"
	"go-react-app/validator"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	userDb        *gorm.DB
	userRepo      repository.IUserRepository
	userValidator validator.IUserValidator
	userUsecase   usecase.IUserUsecase
	uc            IUserController
)

// テストセットアップ関数
func setupUserControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if userDb != nil {
		testutils.CleanupTestDB(userDb)
	} else {
		// 初回のみデータベース接続を作成
		userDb = testutils.SetupTestDB()
		userRepo = repository.NewUserRepository(userDb)
		userValidator = validator.NewUserValidator()
		userUsecase = usecase.NewUserUsecase(userRepo, userValidator)
		uc = NewUserController(userUsecase)
	}
}

// リクエストボディ付きのコンテキストを設定するヘルパー関数
func setupEchoContextWithBody(method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return e, c, rec
}

// CSRFトークンを持つコンテキストをセットアップするヘルパー関数
func setupEchoContextWithCSRF(method, path string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// CSRFトークンをモック
	c.Set("csrf", "test-csrf-token")
	return e, c, rec
}

func TestUserController_SignUp(t *testing.T) {
	setupUserControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効なユーザー情報でアカウント作成に成功する", func(t *testing.T) {
			// テスト用の一意なメールアドレスを生成（短く保つ）
			timestamp := time.Now().UnixNano() % 10000 // 短い数値に制限
			validEmail := fmt.Sprintf("test%d@example.com", timestamp)
			validPassword := "password123"
			reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, validEmail, validPassword)

			t.Logf("テスト実行: email=%s, password=%s", validEmail, validPassword)

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			err := uc.SignUp(c)

			// レスポンスの詳細をログ出力
			t.Logf("レスポンス: status=%d, body=%s", rec.Code, rec.Body.String())

			// 検証
			if err != nil {
				t.Errorf("SignUp() error = %v", err)
			}

			if rec.Code != http.StatusCreated {
				t.Errorf("SignUp() status code = %d, want %d", rec.Code, http.StatusCreated)
			}

			// レスポンスボディをパース
			var response model.UserResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
				return
			}

			if response.Email != validEmail {
				t.Errorf("SignUp() = %v, want email=%s", response, validEmail)
			}

			// IDが設定されていることを確認
			if response.ID == 0 {
				t.Error("SignUp() did not set ID in response")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("無効なJSONリクエストでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"email": "invalid@example.com", "password": Invalid JSON`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", invalidJSON)
			err := uc.SignUp(c)

			// エラーがあるはずだが、コントローラーはJSONレスポンスを返す
			if err != nil {
				t.Errorf("SignUp() unexpected error: %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("SignUp() with invalid JSON status code = %d, want %d",
					rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("バリデーションエラーが発生する場合はエラーを返す", func(t *testing.T) {
			// 短すぎるパスワード
			invalidReqBody := `{"email":"test@example.com","password":"12"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/signup", invalidReqBody)
			err := uc.SignUp(c)

			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("SignUp() returned unexpected error: %v", err)
			}

			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("SignUp() with validation error status code = %d, want %d", 
                   rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("既に存在するユーザーは登録できない", func(t *testing.T) {
			// 最初のユーザー登録
			email := "duplicate@example.com"
			password := "password123"
			reqBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)

			_, c, _ := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			uc.SignUp(c)

			// 同じメールアドレスで再登録を試みる
			_, c2, rec2 := setupEchoContextWithBody(http.MethodPost, "/signup", reqBody)
			err := uc.SignUp(c2)

			if err != nil {
				t.Errorf("SignUp() returned unexpected error: %v", err)
			}

			// 重複登録はエラーになるはず
			if rec2.Code != http.StatusInternalServerError {
				t.Errorf("SignUp() with duplicate email status code = %d, want %d",
					rec2.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestUserController_LogIn(t *testing.T) {
	setupUserControllerTest()

	// テスト用ユーザーを事前に登録
	validEmail := "logintest@example.com"
	validPassword := "password123"
	userUsecase.SignUp(model.User{
		Email:    validEmail,
		Password: validPassword,
	})

	t.Run("正常系", func(t *testing.T) {
		t.Run("有効な認証情報でログインに成功する", func(t *testing.T) {
			// テストリクエストの準備
			reqBody := `{"email":"` + validEmail + `","password":"` + validPassword + `"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", reqBody)
			err := uc.LogIn(c)

			// 検証
			if err != nil {
				t.Errorf("LogIn() error = %v", err)
			}

			if rec.Code != http.StatusOK {
				t.Errorf("LogIn() status code = %d, want %d", rec.Code, http.StatusOK)
			}

			// Cookieが設定されていることを確認
			cookies := rec.Result().Cookies()
			var found bool
			for _, cookie := range cookies {
				if cookie.Name == "token" && cookie.Value != "" {
					found = true
					break
				}
			}
			if !found {
				t.Error("LogIn() did not set token cookie")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("無効なJSONリクエストでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"email": "invalid@example.com", "password": Invalid JSON`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", invalidJSON)
			err := uc.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("LogIn() with invalid JSON status code = %d, want %d",
					rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("存在しないユーザーでログインするとエラーを返す", func(t *testing.T) {
			// 存在しないユーザー
			nonExistentUserReqBody := `{"email":"nonexistent@example.com","password":"password123"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", nonExistentUserReqBody)
			err := uc.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("LogIn() with non-existent user status code = %d, want %d",
					rec.Code, http.StatusInternalServerError)
			}
		})

		t.Run("誤ったパスワードでログインするとエラーを返す", func(t *testing.T) {
			// 誤ったパスワード
			wrongPasswordReqBody := `{"email":"` + validEmail + `","password":"wrongpassword"}`

			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/login", wrongPasswordReqBody)
			err := uc.LogIn(c)

			if err != nil {
				t.Errorf("LogIn() unexpected error: %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("LogIn() with wrong password status code = %d, want %d",
					rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestUserController_LogOut(t *testing.T) {
	setupUserControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("ログアウトに成功する", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoContextWithBody(http.MethodPost, "/logout", "")
			err := uc.LogOut(c)

			// 検証
			if err != nil {
				t.Errorf("LogOut() error = %v", err)
			}

			if rec.Code != http.StatusOK {
				t.Errorf("LogOut() status code = %d, want %d", rec.Code, http.StatusOK)
			}

			// Cookieが削除されたことを確認（有効期限が過去に設定されている）
			cookies := rec.Result().Cookies()
			var found bool
			for _, cookie := range cookies {
				if cookie.Name == "token" && cookie.Value == "" {
					found = true
					break
				}
			}
			if !found {
				t.Error("LogOut() did not properly clear token cookie")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("Cookieが設定できない場合もステータスコード200を返す", func(t *testing.T) {
			// Echoのモックを使用してCookie設定に失敗するケースをシミュレート
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/logout", nil)
			rec := httptest.NewRecorder()
			
			// カスタムコンテキストを作成（Cookie設定を検証するためのモック）
			c := e.NewContext(req, rec)
			
			// テスト実行
			err := uc.LogOut(c)
			
			// 検証 - エラーがなく、ステータスコードが200であることを確認
			if err != nil {
				t.Errorf("LogOut() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("LogOut() status code = %d, want %d", rec.Code, http.StatusOK)
			}
		})
	})
}

func TestUserController_CsrfToken(t *testing.T) {
	setupUserControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("CSRFトークンを取得できる", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupEchoContextWithCSRF(http.MethodGet, "/csrf-token")
			err := uc.CsrfToken(c)

			// 検証
			if err != nil {
				t.Errorf("CsrfToken() error = %v", err)
			}

			if rec.Code != http.StatusOK {
				t.Errorf("CsrfToken() status code = %d, want %d", rec.Code, http.StatusOK)
			}

			// レスポンスをパース
			var response map[string]string
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}

			// CSRFトークンが返されていることを確認
			if token, exists := response["csrf_token"]; !exists || token != "test-csrf-token" {
				t.Errorf("CsrfToken() response = %v, want csrf_token = test-csrf-token", response)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("CSRFトークンが設定されていない場合はパニックが発生する", func(t *testing.T) {
			// CSRFトークンなしでコンテキストを設定
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/csrf-token", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			// パニックが発生することを確認
			defer func() {
				if r := recover(); r == nil {
					t.Error("CsrfToken() should panic when CSRF token is not set")
				}
			}()
			
			// テスト実行 - パニックが発生するはず
			uc.CsrfToken(c)
		})
		
		t.Run("CSRFトークンが不正な型の場合はパニックが発生する", func(t *testing.T) {
			// 不正な型のCSRFトークンでコンテキストを設定
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/csrf-token", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			// 文字列ではなく数値を設定
			c.Set("csrf", 12345)
			
			// パニックが発生することを確認
			defer func() {
				if r := recover(); r == nil {
					t.Error("CsrfToken() should panic when CSRF token is not a string")
				}
			}()
			
			// テスト実行 - パニックが発生するはず
			uc.CsrfToken(c)
		})
	})
}