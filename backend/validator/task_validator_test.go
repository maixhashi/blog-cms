package validator

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
)

func TestTaskValidate(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// テストユーザーの作成
	user := testutils.CreateTestUser(db)
	
	validator := NewTaskValidator()

	testCases := []struct {
		name     string
		task     model.Task
		hasError bool
	}{
		{
			name: "Valid task with valid title",
			task: model.Task{
				Title:  testutils.GenerateValidTitle(),
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Empty title",
			task: model.Task{
				Title:  "",
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Title too long",
			task: model.Task{
				Title:  testutils.GenerateInvalidTitle(),
				UserId: user.ID,
			},
			hasError: true,
		},
		{
			name: "Valid title with exact max length",
			task: model.Task{
				Title:  generateExactMaxLengthTitle(),
				UserId: user.ID,
			},
			hasError: false,
		},
		{
			name: "Zero user ID",
			task: model.Task{
				Title:  "Valid Title",
				UserId: 0,
			},
			hasError: false, // UserIDはバリデーションしていないので、エラーにならないはず
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.TaskValidate(tc.task)
			if (err != nil) != tc.hasError {
				t.Errorf("TaskValidate() error = %v, want error: %v", err, tc.hasError)
			}
		})
	}
}

// 追加の別テストケース - テスト用DBを使ったバリデーション
func TestTaskValidateWithDB(t *testing.T) {
	// テスト用DBの設定
	db := testutils.SetupTestDB()
	defer testutils.CleanupTestDB(db)
	
	// 異なるテストユーザーを作成
	user1 := testutils.CreateTestUser(db)
	user2 := testutils.CreateOtherUser(db)
	
	validator := NewTaskValidator()

	// ユーザー1のタスクを作成
	task1 := model.Task{
		Title:  "User 1 Task",
		UserId: user1.ID,
	}
	
	// ユーザー2のタスクを作成
	task2 := model.Task{
		Title:  "User 2 Task",
		UserId: user2.ID,
	}
	
	// バリデーションのテスト
	err1 := validator.TaskValidate(task1)
	if err1 != nil {
		t.Errorf("TaskValidate() for user1 should not return error, got: %v", err1)
	}
	
	err2 := validator.TaskValidate(task2)
	if err2 != nil {
		t.Errorf("TaskValidate() for user2 should not return error, got: %v", err2)
	}
}

// 最大長さちょうどのタイトルを生成する関数
func generateExactMaxLengthTitle() string {
	title := ""
	for i := 0; i < model.TaskTitleMaxLength; i++ {
		title += "x"
	}
	return title
}
