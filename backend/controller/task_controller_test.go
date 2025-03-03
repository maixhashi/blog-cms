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

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数
var (
	db             *gorm.DB
	taskRepo       repository.ITaskRepository
	taskValidator  validator.ITaskValidator
	taskUsecase    usecase.ITaskUsecase
	tc             ITaskController  // taskController → tc に変更（衝突を避ける）
	testUser       model.User
	otherUser      model.User
)

const nonExistentTaskID uint = 9999

// テストセットアップ関数
func setupTaskControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if db != nil {
		testutils.CleanupTestDB(db)
	} else {
		// 初回のみデータベース接続を作成
		db = testutils.SetupTestDB()
		taskRepo = repository.NewTaskRepository(db)
		taskValidator = validator.NewTaskValidator()
		taskUsecase = usecase.NewTaskUsecase(taskRepo, taskValidator)
		tc = NewTaskController(taskUsecase)  // taskController → tc に変更
	}
	
	// テストユーザーを作成
	testUser = testutils.CreateTestUser(db)
	otherUser = testutils.CreateOtherUser(db)
}

// JWT認証をモックするヘルパー関数
func setupEchoWithJWT(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	
	// コンテキストにトークンを設定
	c.Set("user", token)
	
	return e, c, rec
}

// リクエストボディ付きのコンテキストを設定するヘルパー関数
func setupEchoWithJWTAndBody(userId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	c.Set("user", token)
	
	return e, c, rec
}

// TaskIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithTaskId(userId uint, method, path string, taskId uint, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("taskId")
	c.SetParamValues(fmt.Sprintf("%d", taskId))
	return e, c, rec
}

func TestTaskController_GetAllTasks(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのタスクを全て取得する", func(t *testing.T) {
			// テスト用タスクの作成
			tasks := []model.Task{
				{Title: "Task 1", UserId: testUser.ID},
				{Title: "Task 2", UserId: testUser.ID},
				{Title: "Task 3", UserId: otherUser.ID}, // 別ユーザーのタスク
			}
			
			for _, task := range tasks {
				db.Create(&task)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWT(testUser.ID)
			err := tc.GetAllTasks(c)  // taskController → tc に変更
			
			// 検証
			if err != nil {
				t.Errorf("GetAllTasks() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllTasks() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response []model.TaskResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if len(response) != 2 {
				t.Errorf("GetAllTasks() returned %d tasks, want 2", len(response))
			}
			
			// タスクタイトルの確認
			titles := make(map[string]bool)
			for _, task := range response {
				titles[task.Title] = true
			}
			
			if !titles["Task 1"] || !titles["Task 2"] {
				t.Errorf("期待したタスクが結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可
	})
}

func TestTaskController_GetTaskById(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するタスクを正しく取得する", func(t *testing.T) {
			// テスト用タスクの作成
			task := model.Task{
				Title:  "Test Task",
				UserId: testUser.ID,
			}
			db.Create(&task)
			
			// テスト実行
			_, c, rec := setupEchoWithTaskId(testUser.ID, http.MethodGet, "/tasks/"+fmt.Sprintf("%d", task.ID), task.ID, "")
			err := tc.GetTaskById(c)  // taskController → tc に変更
			
			// 検証
			if err != nil {
				t.Errorf("GetTaskById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetTaskById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response model.TaskResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.ID != task.ID || response.Title != task.Title {
				t.Errorf("GetTaskById() = %v, want id=%d, title=%s", response, task.ID, task.Title)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーのタスクは取得できない", func(t *testing.T) {
			// 他のユーザーのタスクを作成
			otherUserTask := model.Task{
				Title:  "Other User's Task",
				UserId: otherUser.ID,
			}
			db.Create(&otherUserTask)
			
			// テスト実行 - testUserとして他のユーザーのタスクにアクセス
			_, c, rec := setupEchoWithTaskId(testUser.ID, http.MethodGet, "/tasks/"+fmt.Sprintf("%d", otherUserTask.ID), otherUserTask.ID, "")
			err := tc.GetTaskById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetTaskById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetTaskById() with other user's task status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestTaskController_CreateTask(t *testing.T) {
	setupTaskControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいタスクを作成できる", func(t *testing.T) {
			// テストリクエストの準備
			validTitle := testutils.GenerateValidTitle()
			reqBody := fmt.Sprintf(`{"title":"%s"}`, validTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(testUser.ID, http.MethodPost, "/tasks", reqBody)
			err := tc.CreateTask(c)
			
			// 検証
			if err != nil {
				t.Errorf("CreateTask() error = %v", err)
			}
			
			if rec.Code != http.StatusCreated {
				t.Errorf("CreateTask() status code = %d, want %d", rec.Code, http.StatusCreated)
			}
			
			// レスポンスボディをパース
			var response model.TaskResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.Title != validTitle {
				t.Errorf("CreateTask() = %v, want title=%s", response, validTitle)
			}
			
			// データベースから直接確認
			var dbTask model.Task
			db.First(&dbTask, response.ID)
			if dbTask.Title != validTitle || dbTask.UserId != testUser.ID {
				t.Errorf("CreateTask() did not save task correctly, got title=%s, userId=%d", dbTask.Title, dbTask.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するタスクは作成できない", func(t *testing.T) {
			// 無効なタイトル（最大長超過）
			invalidTitle := testutils.GenerateInvalidTitle()
			reqBody := fmt.Sprintf(`{"title":"%s"}`, invalidTitle)
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(testUser.ID, http.MethodPost, "/tasks", reqBody)
			err := tc.CreateTask(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("CreateTask() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("CreateTask() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBody(testUser.ID, http.MethodPost, "/tasks", invalidJSON)
			err := tc.CreateTask(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("CreateTask() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("CreateTask() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}