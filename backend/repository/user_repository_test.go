package repository

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	userRepoDb   *gorm.DB
	userRepo     IUserRepository
	testUserData = model.User{
		Email:    "test@example.com",
		Password: "password123",
	}
)

func setupUserRepositoryTest() {
	// テストごとにデータベースをクリーンアップ
	if userRepoDb != nil {
		testutils.CleanupTestDB(userRepoDb)
	} else {
		// 初回のみデータベース接続を作成
		userRepoDb = testutils.SetupTestDB()
		userRepo = NewUserRepository(userRepoDb)
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	setupUserRepositoryTest()

	t.Run("ユーザーを正常に作成できる", func(t *testing.T) {
		// テスト用ユーザーデータ
		user := testUserData
		
		// テスト実行
		err := userRepo.CreateUser(&user)
		
		// 検証
		assert.NoError(t, err)
		assert.NotZero(t, user.ID, "ユーザーIDが設定されていません")
		assert.NotZero(t, user.CreatedAt, "作成日時が設定されていません")
		assert.NotZero(t, user.UpdatedAt, "更新日時が設定されていません")
	})

	t.Run("同じメールアドレスのユーザーは作成できない", func(t *testing.T) {
		// 最初のユーザーを作成
		firstUser := testUserData
		_ = userRepo.CreateUser(&firstUser)
		
		// 同じメールアドレスで2人目のユーザーを作成
		duplicateUser := model.User{
			Email:    testUserData.Email,
			Password: "different_password",
		}
		
		// テスト実行
		err := userRepo.CreateUser(&duplicateUser)
		
		// 検証
		assert.Error(t, err, "重複するメールアドレスでエラーが発生するはず")
	})
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	setupUserRepositoryTest()

	t.Run("存在するユーザーを取得できる", func(t *testing.T) {
		// テスト用ユーザーを作成
		user := testUserData
		err := userRepo.CreateUser(&user)
		assert.NoError(t, err)
		
		// テスト実行
		foundUser, err := userRepo.GetUserByEmail(user.Email)
		
		// 検証
		assert.NoError(t, err)
		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Email, foundUser.Email)
		assert.Equal(t, user.Password, foundUser.Password)
	})

	t.Run("存在しないユーザーを取得するとエラーになる", func(t *testing.T) {
		// 存在しないメールアドレス
		nonExistentEmail := "nonexistent@example.com"
		
		// テスト実行
		_, err := userRepo.GetUserByEmail(nonExistentEmail)
		
		// 検証
		assert.Error(t, err, "存在しないユーザーを取得するとエラーになるはず")
		assert.Contains(t, err.Error(), "record not found", "レコードが見つからないエラーが発生するはず")
	})
}
