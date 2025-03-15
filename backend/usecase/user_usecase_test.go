package usecase

import (
	"fmt"
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/validator"
	"os"
	"testing"
	
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// テスト用の共通変数とセットアップ関数
var (
	userDb          *gorm.DB
	userRepo        repository.IUserRepository
	userValidator   validator.IUserValidator
	uu              IUserUsecase
	wrongUserEmail  string = "wrong@example.com"
	wrongUserPwd    string = "wrongpassword"
)

// UUID を使用する方法（非常に低い確率で衝突）
func generateUniqueEmail() string {
	// UUIDを生成（ほぼ確実に一意）
	uuid := uuid.New().String()
	// UUIDの最初の8文字だけを使用
	shortUUID := uuid[:8]
	
	return fmt.Sprintf("test-%s@example.com", shortUUID)
	// 例: "test550e8400@example.com" (長さは22文字)
}

func setupUserTest() {
	// テストごとにデータベースをクリーンアップ
	if userDb != nil {
		testutils.CleanupTestDB(userDb)
	} else {
		// 初回のみデータベース接続を作成
		userDb = testutils.SetupTestDB()
		userRepo = repository.NewUserRepository(userDb)
		userValidator = validator.NewUserValidator()
		uu = NewUserUsecase(userRepo, userValidator)
		
		// JWT用のSECRET環境変数を設定
		os.Setenv("SECRET", "test-secret-key")
	}

	// 既存のテストユーザーを明示的に削除（念のため）
	userDb.Exec("DELETE FROM users WHERE email LIKE 'test%@example.com'")
}
	func TestUserUsecase_SignUp(t *testing.T) {
		setupUserTest()
	
		t.Run("正常系", func(t *testing.T) {
			t.Run("新規ユーザーを登録できる", func(t *testing.T) {
				// 一意のメールアドレスを生成
				testUserEmail := generateUniqueEmail()
				testUserPwd := "password123"

				// テスト用ユーザー
				signupReq := model.UserSignupRequest{
					Email:    testUserEmail,
					Password: testUserPwd,
				}
			
				t.Logf("ユーザー登録: Email=%s", signupReq.Email)
			
				// テスト実行
				userRes, err := uu.SignUp(signupReq)
			
				// 検証
				if err != nil {
					t.Errorf("SignUp() error = %v", err)
				}
			
				if userRes.ID == 0 || userRes.Email != signupReq.Email {
					t.Errorf("SignUp() = %v, want email=%s and ID > 0", userRes, signupReq.Email)
				} else {
					t.Logf("生成されたユーザーID: %d", userRes.ID)
				}
			
				// データベースから直接確認
				var savedUser model.User
				userDb.First(&savedUser, userRes.ID)
			
				if savedUser.Email != signupReq.Email {
					t.Errorf("SignUp() saved email = %v, want %v", savedUser.Email, signupReq.Email)
				}
			
				// パスワードがハッシュ化されていることを確認
				err = bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(signupReq.Password))
				if err != nil {
					t.Errorf("パスワードが正しくハッシュ化されていません: %v", err)
				} else {
					t.Log("パスワードが正しくハッシュ化されています")
				}
			})
		})
	
		t.Run("異常系", func(t *testing.T) {
			t.Run("バリデーションエラーが発生する場合はユーザー登録に失敗する", func(t *testing.T) {
				// 無効なメールアドレス
				invalidUser := model.UserSignupRequest{
					Email:    "invalid-email",
					Password: "password123",
				}
			
				t.Logf("無効なユーザー登録を試行: Email=%s", invalidUser.Email)
			
				_, err := uu.SignUp(invalidUser)
			
				// バリデーションエラーが発生するはず
				if err == nil {
					t.Error("無効なメールアドレスでエラーが返されませんでした")
				} else {
					t.Logf("期待通りバリデーションエラーが返されました: %v", err)
				}
			})
		
			t.Run("すでに登録されているメールアドレスで登録に失敗する", func(t *testing.T) {
				// 一意のメールアドレスを生成
				duplicateEmail := generateUniqueEmail()
			
				// 最初に1人目のユーザーを登録
				firstUser := model.UserSignupRequest{
					Email:    duplicateEmail,
					Password: "password123",
				}
				_, err := uu.SignUp(firstUser)
				if err != nil {
					t.Fatalf("最初のユーザー登録に失敗しました: %v", err)
				}
			
				// 同じメールアドレスで2人目のユーザーを登録
				duplicateUser := model.UserSignupRequest{
					Email:    duplicateEmail,
					Password: "different_password",
				}
			
				t.Logf("重複するメールアドレスでユーザー登録を試行: Email=%s", duplicateUser.Email)
			
				_, err = uu.SignUp(duplicateUser)
			
				// エラーが発生するはず
				if err == nil {
					t.Error("重複するメールアドレスでエラーが返されませんでした")
				} else {
					t.Logf("期待通りエラーが返されました: %v", err)
				}
			})
		})
	}

	func TestUserUsecase_Login(t *testing.T) {
		setupUserTest()
	
		// 一意のメールアドレスを生成
		testUserEmail := generateUniqueEmail()
		testUserPwd := "password123"
	
		// テスト用ユーザーを登録
		testUser := model.UserSignupRequest{
			Email:    testUserEmail,
			Password: testUserPwd,
		}
		_, err := uu.SignUp(testUser)
		if err != nil {
			t.Fatalf("テストユーザーの登録に失敗しました: %v", err)
		}
	
		t.Run("正常系", func(t *testing.T) {
			t.Run("有効な認証情報でログインに成功する", func(t *testing.T) {
				loginReq := model.UserLoginRequest{
					Email:    testUserEmail,
					Password: testUserPwd,
				}
			
				t.Logf("ログイン試行: Email=%s", loginReq.Email)
			
				// テスト実行
				tokenString, err := uu.Login(loginReq)
			
				// 検証
				if err != nil {
					t.Errorf("Login() error = %v", err)
				}
			
				if tokenString == "" {
					t.Error("Login() returned empty token")
				} else {
					t.Log("JWTトークンが正常に生成されました")
				}
			
				// トークンが有効なJWTかどうか検証
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("SECRET")), nil
				})
			
				if err != nil || !token.Valid {
					t.Errorf("生成されたトークンが無効です: %v", err)
				} else {
					// クレームの検証
					claims, ok := token.Claims.(jwt.MapClaims)
					if !ok {
						t.Error("トークンクレームの取得に失敗しました")
					} else {
						userId, exists := claims["user_id"]
						if !exists || userId == 0 {
							t.Errorf("トークンにuser_idが含まれていないか、無効な値です: %v", userId)
						} else {
							t.Logf("トークン内のuser_id: %v", userId)
						}
					}
				}
			})
		})
	
		t.Run("異常系", func(t *testing.T) {
			t.Run("バリデーションエラーが発生する場合はログインに失敗する", func(t *testing.T) {
				// 無効なメールアドレス
				invalidUser := model.UserLoginRequest{
					Email:    "invalid-email",
					Password: testUserPwd,
				}
			
				t.Logf("無効なユーザーでログインを試行: Email=%s", invalidUser.Email)
			
				_, err := uu.Login(invalidUser)
			
				// バリデーションエラーが発生するはず
				if err == nil {
					t.Error("無効なメールアドレスでエラーが返されませんでした")
				} else {
					t.Logf("期待通りバリデーションエラーが返されました: %v", err)
				}
			})
		
			t.Run("存在しないユーザーでログインに失敗する", func(t *testing.T) {
				nonExistUser := model.UserLoginRequest{
					Email:    wrongUserEmail,
					Password: testUserPwd,
				}
			
				t.Logf("存在しないユーザーでログインを試行: Email=%s", nonExistUser.Email)
			
				_, err := uu.Login(nonExistUser)
			
				if err == nil {
					t.Error("存在しないユーザーでエラーが返されませんでした")
				} else {
					t.Logf("期待通りエラーが返されました: %v", err)
				}
			})
		
			t.Run("パスワードが間違っている場合はログインに失敗する", func(t *testing.T) {
				wrongPwdUser := model.UserLoginRequest{
					Email:    testUserEmail,
					Password: wrongUserPwd,
				}
			
				t.Logf("間違ったパスワードでログインを試行: Email=%s", wrongPwdUser.Email)
			
				_, err := uu.Login(wrongPwdUser)
			
				if err == nil {
					t.Error("間違ったパスワードでエラーが返されませんでした")
				} else {
					t.Logf("期待通りエラーが返されました: %v", err)
				}
			})
		})
	}
