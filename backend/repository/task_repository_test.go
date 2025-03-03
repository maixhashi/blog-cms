package repository

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB
var taskRepo ITaskRepository
var testUser model.User
var otherUser model.User
const nonExistentTaskID uint = 9999

func setupTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		taskRepo = NewTaskRepository(db)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(db)
}

func TestTaskRepository_GetAllTasks(t *testing.T) {
	setupTest()
	
	// テストデータの作成
	tasks := []model.Task{
		{Title: "Task 1", UserId: testUser.ID},
		{Title: "Task 2", UserId: testUser.ID},
		{Title: "Task 3", UserId: otherUser.ID}, // 別ユーザーのタスク
	}
	
	for _, task := range tasks {
		db.Create(&task)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのタスクのみを取得する", func(t *testing.T) {
			var result []model.Task
			err := taskRepo.GetAllTasks(&result, testUser.ID)
			
			t.Logf("ユーザーID %d のタスクを取得します", testUser.ID)
			
			if err != nil {
				t.Errorf("GetAllTasks() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetAllTasks() got %d tasks, want 2", len(result))
			}
			
			// タスクタイトルの確認
			titles := make(map[string]bool)
			for _, task := range result {
				titles[task.Title] = true
				t.Logf("取得したタスク: ID=%d, Title=%s", task.ID, task.Title)
			}
			
			if !titles["Task 1"] || !titles["Task 2"] {
				t.Errorf("期待したタスクが結果に含まれていません: %v", result)
			}
		})
	})
	
	// 異常系のテストケースを追加する場合はここに
}

func TestTaskRepository_GetTaskById(t *testing.T) {
	setupTest()
	
	// テストデータの作成
	task := model.Task{
		Title:  "Test Task",
		UserId: testUser.ID,
	}
	db.Create(&task)

	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するタスクを正しく取得する", func(t *testing.T) {
			t.Logf("タスクID %d を取得します", task.ID)
	
			var result model.Task
			err := taskRepo.GetTaskById(&result, testUser.ID, task.ID)
	
			if err != nil {
				t.Errorf("GetTaskById() error = %v", err)
			}
	
			if result.ID != task.ID || result.Title != task.Title {
				t.Errorf("取得したタスク %+v が、期待値 %+v と一致しません", result, task)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", result.ID, result.Title)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してタスクを取得しようとします", nonExistentTaskID)
	
			var notFound model.Task
			err := taskRepo.GetTaskById(&notFound, testUser.ID, nonExistentTaskID)
	
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーのタスクは取得できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := model.Task{
				Title:  "Other User's Task",
				UserId: otherUser.ID,
			}
			db.Create(&otherUserTask)
			t.Logf("他ユーザーのタスク(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserTask.ID, testUser.ID)
			
			var result model.Task
			err := taskRepo.GetTaskById(&result, testUser.ID, otherUserTask.ID)
			
			if err == nil {
				t.Error("他のユーザーのタスクを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

func TestTaskRepository_CreateTask(t *testing.T) {
	setupTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいタスクを作成できる", func(t *testing.T) {
			// テスト用タスク
			task := model.Task{
				Title:  "New Task",
				UserId: testUser.ID,
			}
			
			t.Logf("タスク作成: Title=%s, UserId=%d", task.Title, task.UserId)
			
			// テスト実行
			err := taskRepo.CreateTask(&task)
			
			// 検証
			if err != nil {
				t.Errorf("CreateTask() error = %v", err)
			}
			
			if task.ID == 0 {
				t.Error("CreateTask() did not set ID")
			} else {
				t.Logf("生成されたタスクID: %d", task.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
				t.Error("CreateTask() did not set timestamps")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", task.CreatedAt, task.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedTask model.Task
			db.First(&savedTask, task.ID)
			
			if savedTask.Title != "New Task" || savedTask.UserId != testUser.ID {
				t.Errorf("CreateTask() = %v, want title=%s, userId=%d", savedTask, "New Task", testUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, UserId=%d", savedTask.Title, savedTask.UserId)
			}
		})
	})
	
	// CreateTaskの異常系があれば追加
}

func TestTaskRepository_UpdateTask(t *testing.T) {
	setupTest()
	
	// テストデータの作成
	task := model.Task{
		Title:  "Original Title",
		UserId: testUser.ID,
	}
	db.Create(&task)
	t.Logf("元のタスク作成: ID=%d, Title=%s", task.ID, task.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("タスクのタイトルを更新できる", func(t *testing.T) {
			// 更新するタスク
			updatedTask := model.Task{
				Title: "Updated Title",
			}
			
			t.Logf("タスク更新リクエスト: ID=%d, 新Title=%s", task.ID, updatedTask.Title)
			
			// テスト実行
			err := taskRepo.UpdateTask(&updatedTask, testUser.ID, task.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateTask() error = %v", err)
			} else {
				t.Log("タスク更新成功")
			}
			
			// 返り値のタスクが更新されていることを確認
			if updatedTask.Title != "Updated Title" {
				t.Errorf("UpdateTask() returned task title = %v, want %v", updatedTask.Title, "Updated Title")
			} else {
				t.Logf("返り値確認: Title=%s", updatedTask.Title)
			}
			
			// データベースから直接確認
			var dbTask model.Task
			db.First(&dbTask, task.ID)
			
			if dbTask.Title != "Updated Title" {
				t.Errorf("UpdateTask() database task title = %v, want %v", dbTask.Title, "Updated Title")
			} else {
				t.Logf("データベース更新確認: Title=%s", dbTask.Title)
			}
			
			// 更新日時が変更されていることを確認
			if !dbTask.UpdatedAt.After(task.UpdatedAt) {
				t.Error("UpdateTask() did not update the updated_at timestamp")
			} else {
				t.Logf("タイムスタンプ更新確認: 元=%v, 更新後=%v", task.UpdatedAt, dbTask.UpdatedAt)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
			invalidTask := model.Task{Title: "Invalid Update"}
			t.Logf("存在しないID %d でタスク更新を試行", nonExistentTaskID)
			
			err := taskRepo.UpdateTask(&invalidTask, testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("UpdateTask() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーのタスクは更新できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := model.Task{
				Title:  "Other User's Task",
				UserId: otherUser.ID,
			}
			db.Create(&otherUserTask)
			t.Logf("他ユーザーのタスク: ID=%d, Title=%s, UserId=%d", otherUserTask.ID, otherUserTask.Title, otherUserTask.UserId)
			
			// 他ユーザーのタスクを更新しようとする
			updateAttempt := model.Task{Title: "Attempted Update"}
			err := taskRepo.UpdateTask(&updateAttempt, testUser.ID, otherUserTask.ID)
			
			if err == nil {
				t.Error("UpdateTask() should not allow updating other user's task")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}

func TestTaskRepository_DeleteTask(t *testing.T) {
	setupTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のタスクを削除できる", func(t *testing.T) {
			// テストデータの作成
			task := model.Task{
				Title:  "Task to Delete",
				UserId: testUser.ID,
			}
			db.Create(&task)
			t.Logf("削除対象タスク作成: ID=%d, Title=%s", task.ID, task.Title)
			
			// テスト実行
			err := taskRepo.DeleteTask(testUser.ID, task.ID)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteTask() error = %v", err)
			} else {
				t.Logf("タスク削除成功: ID=%d", task.ID)
			}
			
			// データベースから直接確認
			var count int64
			db.Model(&model.Task{}).Where("id = ?", task.ID).Count(&count)
			
			if count != 0 {
				t.Error("DeleteTask() did not delete the task from database")
			} else {
				t.Log("データベースからタスクが削除されていることを確認")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないタスクIDでの削除はエラーになる", func(t *testing.T) {
			t.Logf("存在しないID %d でタスク削除を試行", nonExistentTaskID)
			
			err := taskRepo.DeleteTask(testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("DeleteTask() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーのタスクを削除しようとするとエラー", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := model.Task{
				Title:  "Other User's Task",
				UserId: otherUser.ID,
			}
			db.Create(&otherUserTask)
			t.Logf("他ユーザーのタスク作成: ID=%d, Title=%s, UserId=%d", otherUserTask.ID, otherUserTask.Title, otherUserTask.UserId)
			
			// 他ユーザーのタスクを削除しようとする
			err := taskRepo.DeleteTask(testUser.ID, otherUserTask.ID)
			if err == nil {
				t.Error("DeleteTask() should not allow deleting other user's task")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}