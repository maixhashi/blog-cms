package testutils

import (
	"go-react-app/model"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// テスト用データベースを設定する
func SetupTestDB() *gorm.DB {
	// テスト用のSQLiteインメモリデータベースを使用
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Silent, // テスト時はログを抑制
			},
		),
	})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// テスト用のテーブルを作成
	db.AutoMigrate(&model.User{}, &model.Task{})

	return db
}

// テストユーザーの作成
func CreateTestUser(db *gorm.DB) model.User {
	user := model.User{
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(&user)
	return user
}

// テストユーザーの作成
func CreateOtherUser(db *gorm.DB) model.User {
	user := model.User{
		Email:    "other@example.com",
		Password: "otherpassword",
	}
	db.Create(&user)
	return user
}

// テストデータのクリーンアップ
func CleanupTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM users")
}

