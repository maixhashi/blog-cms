package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/validator"
	"testing"
	"time"
	
	"gorm.io/gorm"
)

// テスト用の共通変数とセットアップ関数
var (
	db             *gorm.DB
	taskRepo       repository.ITaskRepository
	taskValidator  validator.ITaskValidator
	tu             ITaskUsecase
	testUser       model.User
	otherUser      model.User
)

const nonExistentTaskID uint = 9999

func setupTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		taskRepo = repository.NewTaskRepository(db)
		taskValidator = validator.NewTaskValidator()
		tu = NewTaskUsecase(taskRepo, taskValidator)
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	
	// 別のテストユーザーを作成
	otherUser = testutils.CreateOtherUser(db)
}

func TestTaskUsecase_GetAllTasks(t *testing.T) {
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
			t.Logf("ユーザーID %d のタスクを取得します", testUser.ID)
			
			taskResponses, err := tu.GetAllTasks(testUser.ID)
			
			if err != nil {
				t.Errorf("GetAllTasks() error = %v", err)
			}
			
			if len(taskResponses) != 2 {
				t.Errorf("GetAllTasks() got %d tasks, want 2", len(taskResponses))
			}
			
			// タスクタイトルの確認
			titles := make(map[string]bool)
			for _, task := range taskResponses {
				titles[task.Title] = true
				t.Logf("取得したタスク: ID=%d, Title=%s", task.ID, task.Title)
			}
			
			if !titles["Task 1"] || !titles["Task 2"] {
				t.Errorf("期待したタスクが結果に含まれていません: %v", taskResponses)
			}
			
			// レスポンス形式の検証
			for _, task := range taskResponses {
				if task.ID == 0 || task.Title == "" || task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
					t.Errorf("GetAllTasks() returned invalid task: %+v", task)
				}
			}
		})
	})
}

func TestTaskUsecase_GetTaskById(t *testing.T) {
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
			
			response, err := tu.GetTaskById(testUser.ID, task.ID)
			
			if err != nil {
				t.Errorf("GetTaskById() error = %v", err)
			}
			
			if response.ID != task.ID || response.Title != task.Title {
				t.Errorf("取得したタスク %+v が、期待値 %+v と一致しません", response, task)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", response.ID, response.Title)
			}
			
			// レスポンス形式の検証
			if response.ID == 0 || response.Title == "" || response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Errorf("GetTaskById() returned invalid task: %+v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してタスクを取得しようとします", nonExistentTaskID)
			
			_, err := tu.GetTaskById(testUser.ID, nonExistentTaskID)
			
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
			
			_, err := tu.GetTaskById(testUser.ID, otherUserTask.ID)
			
			if err == nil {
				t.Error("他のユーザーのタスクを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

func TestTaskUsecase_CreateTask(t *testing.T) {
	setupTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいタスクを作成できる", func(t *testing.T) {
			// テスト用タスク - 有効なタイトルを生成関数から取得
			validTitle := testutils.GenerateValidTitle()
			validRequest := model.TaskRequest{
				Title:  validTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("タスク作成: Title=%s, UserId=%d", validRequest.Title, validRequest.UserId)
			
			// テスト実行
			response, err := tu.CreateTask(validRequest)
			
			// 検証
			if err != nil {
				t.Errorf("CreateTask() error = %v", err)
			}
			
			if response.ID == 0 || response.Title != validRequest.Title {
				t.Errorf("CreateTask() returned unexpected response: %+v", response)
			} else {
				t.Logf("生成されたタスクID: %d", response.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Error("CreateTask() did not set timestamps in response")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", response.CreatedAt, response.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedTask model.Task
			db.First(&savedTask, response.ID)
			
			if savedTask.Title != validRequest.Title || savedTask.UserId != testUser.ID {
				t.Errorf("CreateTask() = %v, want title=%s, userId=%d", savedTask, validRequest.Title, testUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, UserId=%d", savedTask.Title, savedTask.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは作成できない", func(t *testing.T) {
			// 無効なタスク - 長すぎるタイトルをヘルパー関数で生成
			invalidTitle := testutils.GenerateInvalidTitle()
			invalidRequest := model.TaskRequest{
				Title:  invalidTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("無効なタスク作成を試行: Title=%s (長さ: %d)", 
				invalidRequest.Title, len(invalidRequest.Title))
			
			_, err := tu.CreateTask(invalidRequest)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なタスクでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに保存されていないことを確認
			var count int64
			db.Model(&model.Task{}).Where("title = ?", invalidRequest.Title).Count(&count)
			if count > 0 {
				t.Error("バリデーションエラーのタスクがデータベースに保存されています")
			} else {
				t.Log("バリデーションエラーのタスクは保存されていないことを確認")
			}
		})
	})
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	setupTest()
	
	// テストデータの作成
	task := model.Task{
		Title:  testutils.GenerateValidTitle(), // 有効なタイトルを使用
		UserId: testUser.ID,
	}
	db.Create(&task)
	t.Logf("元のタスク作成: ID=%d, Title=%s", task.ID, task.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("タスクのタイトルを更新できる", func(t *testing.T) {
			// 更新するタスク - 別の有効なタイトルを使用
			updatedTitle := testutils.GenerateValidTitle() + "2" // 末尾に2を追加して別のタイトルに
			updateRequest := model.TaskRequest{
				Title:  updatedTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("タスク更新リクエスト: ID=%d, 新Title=%s", task.ID, updateRequest.Title)
			
			// テスト実行
			response, err := tu.UpdateTask(updateRequest, testUser.ID, task.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateTask() error = %v", err)
			} else {
				t.Log("タスク更新成功")
			}
			
			// 返り値のタスクが更新されていることを確認
			if response.ID != task.ID || response.Title != updateRequest.Title {
				t.Errorf("UpdateTask() = %+v, want id=%d, title=%s", response, task.ID, updateRequest.Title)
			} else {
				t.Logf("返り値確認: ID=%d, Title=%s", response.ID, response.Title)
			}
			
			// データベースから直接確認
			var dbTask model.Task
			db.First(&dbTask, task.ID)
			
			if dbTask.Title != updateRequest.Title {
				t.Errorf("UpdateTask() database task title = %v, want %v", dbTask.Title, updateRequest.Title)
			} else {
				t.Logf("データベース更新確認: Title=%s", dbTask.Title)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは更新できない", func(t *testing.T) {
			// 無効な更新 - 長すぎるタイトル
			invalidTitle := testutils.GenerateInvalidTitle()
			invalidRequest := model.TaskRequest{
				Title:  invalidTitle,
				UserId: testUser.ID,
			}
			
			t.Logf("無効なタイトルでの更新を試行: %s (長さ: %d)", 
				invalidRequest.Title, len(invalidRequest.Title))
			
			_, err := tu.UpdateTask(invalidRequest, testUser.ID, task.ID)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なタイトルでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに反映されていないことを確認
			var dbTask model.Task
			db.First(&dbTask, task.ID)
			if dbTask.Title == invalidRequest.Title {
				t.Error("バリデーションエラーの更新がデータベースに反映されています")
			} else {
				t.Logf("データベース確認: Title=%s (変更されていない)", dbTask.Title)
			}
		})
		
		t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
			updateAttempt := model.TaskRequest{Title: "Valid"}
			t.Logf("存在しないID %d でタスク更新を試行", nonExistentTaskID)
		
			_, err := tu.UpdateTask(updateAttempt, testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
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
			updateAttempt := model.TaskRequest{Title: "Attempted Update"}
			_, err := tu.UpdateTask(updateAttempt, testUser.ID, otherUserTask.ID)
		
			if err == nil {
				t.Error("他のユーザーのタスク更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに反映されていないことを確認
			var dbTask model.Task
			db.First(&dbTask, otherUserTask.ID)
			if dbTask.Title != otherUserTask.Title {
				t.Errorf("他ユーザーのタスクが変更されています: %s → %s", otherUserTask.Title, dbTask.Title)
			} else {
				t.Log("他ユーザーのタスクは変更されていないことを確認")
			}
		})
	})
}

func TestTaskUsecase_DeleteTask(t *testing.T) {
	setupTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のタスクを削除できる", func(t *testing.T) {
			// テスト用タスクの作成
			task := model.Task{
				Title:  "Task to Delete",
				UserId: testUser.ID,
			}
			db.Create(&task)
			t.Logf("削除対象タスク作成: ID=%d, Title=%s", task.ID, task.Title)
		
			// テスト実行
			err := tu.DeleteTask(testUser.ID, task.ID)
		
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
		
			err := tu.DeleteTask(testUser.ID, nonExistentTaskID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのタスクは削除できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := model.Task{
				Title:  "Other User's Task",
				UserId: otherUser.ID,
			}
			db.Create(&otherUserTask)
			t.Logf("他ユーザーのタスク作成: ID=%d, Title=%s, UserId=%d", otherUserTask.ID, otherUserTask.Title, otherUserTask.UserId)
		
			// 他ユーザーのタスクを削除しようとする
			err := tu.DeleteTask(testUser.ID, otherUserTask.ID)
			if err == nil {
				t.Error("他のユーザーのタスク削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに残っていることを確認
			var count int64
			db.Model(&model.Task{}).Where("id = ?", otherUserTask.ID).Count(&count)
			if count == 0 {
				t.Error("他ユーザーのタスクが削除されています")
			} else {
				t.Log("他ユーザーのタスクは削除されていないことを確認")
			}
		})
	})
}
