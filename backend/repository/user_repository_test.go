package repository

import (
	"fmt"
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
	"time"

	"gorm.io/gorm"
)

var (
	userDB   *gorm.DB
	userRepo IUserRepository
)

// 一意なメールアドレスを生成するヘルパー関数
func generateUniqueEmail() string {
	return fmt.Sprintf("test%d@example.com", time.Now().UnixNano())
}

func setupUserRepoTest() {
	// テストごとにデータベースをクリーンアップ
	if userDB != nil {
		testutils.CleanupTestDB(userDB)
	} else {
		// 初回のみデータベース接続を作成
		userDB = testutils.SetupTestDB()
		userRepo = NewUserRepository(userDB)
	}
	
	// データベースが本当にクリーンか確認するために追加チェック
	var count int64
	userDB.Model(&model.User{}).Count(&count)
	if count > 0 {
		// もし残っているレコードがあれば、強制的に全削除
		userDB.Exec("DELETE FROM users")
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	setupUserRepoTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいユーザーを作成できる", func(t *testing.T) {
			// テスト用ユーザー（一意のメールアドレスを使用）
			testEmail := generateUniqueEmail()
			user := model.User{
				Email:    testEmail,
				Password: "password123",
			}

			t.Logf("ユーザー作成: Email=%s", user.Email)

			// テスト実行
			err := userRepo.CreateUser(&user)

			// 検証
			if err != nil {
				t.Errorf("CreateUser() error = %v", err)
			}

			if user.ID == 0 {
				t.Error("CreateUser() did not set ID")
			} else {
				t.Logf("生成されたユーザーID: %d", user.ID)
			}

			// タイムスタンプが設定されていることを確認
			if user.CreatedAt.IsZero() || user.UpdatedAt.IsZero() {
				t.Error("CreateUser() did not set timestamps")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", user.CreatedAt, user.UpdatedAt)
			}

			// データベースから直接確認
			var savedUser model.User
			userDB.First(&savedUser, user.ID)

			if savedUser.Email != testEmail {
				t.Errorf("CreateUser() = %v, want email=%s", savedUser, testEmail)
			} else {
				t.Logf("データベース保存確認: Email=%s", savedUser.Email)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("同じメールアドレスで重複ユーザーを作成しようとするとエラー", func(t *testing.T) {
			// まず一意のメールを持つユーザーを作成
			testEmail := generateUniqueEmail()
			originalUser := model.User{
				Email:    testEmail,
				Password: "password123",
			}
			userRepo.CreateUser(&originalUser)
			
			// 同じメールで別ユーザーを作成
			duplicateUser := model.User{
				Email:    testEmail,
				Password: "anotherpassword",
			}

			t.Logf("重複ユーザー作成試行: Email=%s", duplicateUser.Email)

			// テスト実行
			err := userRepo.CreateUser(&duplicateUser)

			// 検証
			if err == nil {
				t.Error("CreateUser() with duplicate email should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	setupUserRepoTest()

	// テストデータの作成（一意のメールアドレスを使用）
	testEmail := generateUniqueEmail()
	user := model.User{
		Email:    testEmail,
		Password: "password123",
	}
	userDB.Create(&user)

	t.Run("正常系", func(t *testing.T) {
		t.Run("メールアドレスでユーザーを取得できる", func(t *testing.T) {
			t.Logf("メールアドレス %s でユーザーを取得します", testEmail)

			var result model.User
			err := userRepo.GetUserByEmail(&result, testEmail)

			if err != nil {
				t.Errorf("GetUserByEmail() error = %v", err)
			}

			if result.ID != user.ID || result.Email != user.Email {
				t.Errorf("GetUserByEmail() = %v, want user with id=%d, email=%s", result, user.ID, user.Email)
			} else {
				t.Logf("正常に取得: ID=%d, Email=%s", result.ID, result.Email)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないメールアドレスを指定した場合はエラーを返す", func(t *testing.T) {
			nonExistentEmail := "nonexistent@example.com"
			t.Logf("存在しないメールアドレス %s を指定してユーザーを取得しようとします", nonExistentEmail)

			var result model.User
			err := userRepo.GetUserByEmail(&result, nonExistentEmail)

			if err == nil {
				t.Error("GetUserByEmail() with non-existent email should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}
