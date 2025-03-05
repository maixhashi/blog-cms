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
	feedDB          *gorm.DB
	feedRepo        repository.IFeedRepository
	feedValidator   validator.IFeedValidator
	feedUsecase     usecase.IFeedUsecase
	fc              IFeedController
	feedTestUser    model.User
	feedOtherUser   model.User
)

const nonExistentFeedID uint = 9999

// テストセットアップ関数
func setupFeedControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if feedDB != nil {
		testutils.CleanupTestDB(feedDB)
	} else {
		// 初回のみデータベース接続を作成
		feedDB = testutils.SetupTestDB()
		feedRepo = repository.NewFeedRepository(feedDB)
		feedValidator = validator.NewFeedValidator()
		feedUsecase = usecase.NewFeedUsecase(feedRepo, feedValidator)
		fc = NewFeedController(feedUsecase)
	}
	
	// テストユーザーを作成
	feedTestUser = testutils.CreateTestUser(feedDB)
	feedOtherUser = testutils.CreateOtherUser(feedDB)
}

// JWT認証をモックするヘルパー関数
func setupEchoWithJWTForFeed(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
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
func setupEchoWithJWTAndBodyForFeed(userId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
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

// FeedIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithFeedId(userId uint, method, path string, feedId uint, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBodyForFeed(userId, method, path, body)
	c.SetParamNames("feedId")
	c.SetParamValues(fmt.Sprintf("%d", feedId))
	return e, c, rec
}

func TestFeedController_GetAllFeeds(t *testing.T) {
	setupFeedControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーのフィードを全て取得する", func(t *testing.T) {
			// テスト用フィードの作成
			feeds := []model.Feed{
				{Title: "Feed 1", URL: "https://example.com/feed1", UserId: feedTestUser.ID},
				{Title: "Feed 2", URL: "https://example.com/feed2", UserId: feedTestUser.ID},
				{Title: "Feed 3", URL: "https://example.com/feed3", UserId: feedOtherUser.ID}, // 別ユーザーのフィード
			}
			
			for _, feed := range feeds {
				feedDB.Create(&feed)
			}
			
			// テスト実行
			_, c, rec := setupEchoWithJWTForFeed(feedTestUser.ID)
			err := fc.GetAllFeeds(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllFeeds() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllFeeds() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response []model.FeedResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if len(response) != 2 {
				t.Errorf("GetAllFeeds() returned %d feeds, want 2", len(response))
			}
			
			// フィードタイトルの確認
			titles := make(map[string]bool)
			for _, feed := range response {
				titles[feed.Title] = true
			}
			
			if !titles["Feed 1"] || !titles["Feed 2"] {
				t.Errorf("期待したフィードが結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可
	})
}

func TestFeedController_GetFeedById(t *testing.T) {
	setupFeedControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するフィードを正しく取得する", func(t *testing.T) {
			// テスト用フィードの作成
			feed := model.Feed{
				Title:  "Test Feed",
				URL:    "https://example.com/test",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			
			// テスト実行
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodGet, "/feeds/"+fmt.Sprintf("%d", feed.ID), feed.ID, "")
			err := fc.GetFeedById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetFeedById() error = %v", err)
			}
			// 検証
			if err != nil {
				t.Errorf("GetFeedById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetFeedById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response model.FeedResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.ID != feed.ID || response.Title != feed.Title || response.URL != feed.URL {
				t.Errorf("GetFeedById() = %v, want id=%d, title=%s, url=%s", response, feed.ID, feed.Title, feed.URL)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーのフィードは取得できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			
			// テスト実行 - feedTestUserとして他のユーザーのフィードにアクセス
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodGet, "/feeds/"+fmt.Sprintf("%d", otherUserFeed.ID), otherUserFeed.ID, "")
			err := fc.GetFeedById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetFeedById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetFeedById() with other user's feed status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestFeedController_CreateFeed(t *testing.T) {
	setupFeedControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいフィードを作成できる", func(t *testing.T) {
			// テストリクエストの準備
			validTitle := "Test Feed Title"
			validURL := "https://example.com/valid-feed"
			reqBody := fmt.Sprintf(`{"title":"%s", "url":"%s"}`, validTitle, validURL)
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBodyForFeed(feedTestUser.ID, http.MethodPost, "/feeds", reqBody)
			err := fc.CreateFeed(c)
			
			// 検証
			if err != nil {
				t.Errorf("CreateFeed() error = %v", err)
			}
			
			if rec.Code != http.StatusCreated {
				t.Errorf("CreateFeed() status code = %d, want %d", rec.Code, http.StatusCreated)
			}
			
			// レスポンスボディをパース
			var response model.FeedResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.Title != validTitle || response.URL != validURL {
				t.Errorf("CreateFeed() = %v, want title=%s, url=%s", response, validTitle, validURL)
			}
			
			// データベースから直接確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, response.ID)
			if dbFeed.Title != validTitle || dbFeed.URL != validURL || dbFeed.UserId != feedTestUser.ID {
				t.Errorf("CreateFeed() did not save feed correctly, got title=%s, url=%s, userId=%d", dbFeed.Title, dbFeed.URL, dbFeed.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するフィードは作成できない", func(t *testing.T) {
			// 無効なフィード（タイトルなし）
			invalidReqBody := `{"title":"", "url":"https://example.com/invalid"}`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBodyForFeed(feedTestUser.ID, http.MethodPost, "/feeds", invalidReqBody)
			err := fc.CreateFeed(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("CreateFeed() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("CreateFeed() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupEchoWithJWTAndBodyForFeed(feedTestUser.ID, http.MethodPost, "/feeds", invalidJSON)
			err := fc.CreateFeed(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("CreateFeed() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("CreateFeed() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}

func TestFeedController_UpdateFeed(t *testing.T) {
	setupFeedControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("フィードを正常に更新できる", func(t *testing.T) {
			// テスト用フィードの作成
			feed := model.Feed{
				Title:  "Original Feed",
				URL:    "https://example.com/original",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			
			// 更新用のリクエストボディ
			updatedTitle := "Updated Feed"
			updatedURL := "https://example.com/updated"
			reqBody := fmt.Sprintf(`{"title":"%s", "url":"%s"}`, updatedTitle, updatedURL)
			
			// テスト実行
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodPut, "/feeds/"+fmt.Sprintf("%d", feed.ID), feed.ID, reqBody)
			err := fc.UpdateFeed(c)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateFeed() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("UpdateFeed() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response model.FeedResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.ID != feed.ID || response.Title != updatedTitle || response.URL != updatedURL {
				t.Errorf("UpdateFeed() = %v, want id=%d, title=%s, url=%s", response, feed.ID, updatedTitle, updatedURL)
			}
			
			// データベースから直接確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, feed.ID)
			if dbFeed.Title != updatedTitle || dbFeed.URL != updatedURL {
				t.Errorf("UpdateFeed() did not update feed correctly, got title=%s, url=%s", dbFeed.Title, dbFeed.URL)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーのフィードは更新できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			
			// 更新用のリクエストボディ
			reqBody := `{"title":"Attempted Update", "url":"https://example.com/attempt"}`
			
			// テスト実行 - feedTestUserとして他のユーザーのフィードを更新しようとする
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodPut, "/feeds/"+fmt.Sprintf("%d", otherUserFeed.ID), otherUserFeed.ID, reqBody)
			err := fc.UpdateFeed(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateFeed() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateFeed() with other user's feed status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから直接確認して、フィードが更新されていないことを確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, otherUserFeed.ID)
			if dbFeed.Title != otherUserFeed.Title || dbFeed.URL != otherUserFeed.URL {
				t.Errorf("UpdateFeed() incorrectly updated other user's feed, got title=%s, url=%s", dbFeed.Title, dbFeed.URL)
			}
		})
		
		t.Run("バリデーションエラーが発生するフィードは更新できない", func(t *testing.T) {
			// テスト用フィードの作成
			feed := model.Feed{
				Title:  "Test Feed for Validation",
				URL:    "https://example.com/test-validation",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			
			// 無効なリクエストボディ（タイトルなし）
			invalidReqBody := `{"title":"", "url":"https://example.com/invalid-update"}`
			
			// テスト実行
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodPut, "/feeds/"+fmt.Sprintf("%d", feed.ID), feed.ID, invalidReqBody)
			err := fc.UpdateFeed(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateFeed() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateFeed() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestFeedController_DeleteFeed(t *testing.T) {
	setupFeedControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のフィードを削除できる", func(t *testing.T) {
			// テスト用フィードの作成
			feed := model.Feed{
				Title:  "Feed to Delete",
				URL:    "https://example.com/delete",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			
			// テスト実行
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodDelete, "/feeds/"+fmt.Sprintf("%d", feed.ID), feed.ID, "")
			err := fc.DeleteFeed(c)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteFeed() error = %v", err)
			}
			
			if rec.Code != http.StatusNoContent {
				t.Errorf("DeleteFeed() status code = %d, want %d", rec.Code, http.StatusNoContent)
			}
			
			// データベースから直接確認
			var count int64
			feedDB.Model(&model.Feed{}).Where("id = ?", feed.ID).Count(&count)
			if count != 0 {
				t.Error("DeleteFeed() did not delete the feed from database")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーのフィードは削除できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed to Not Delete",
				URL:    "https://example.com/not-delete",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			
			// テスト実行 - feedTestUserとして他のユーザーのフィードを削除しようとする
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodDelete, "/feeds/"+fmt.Sprintf("%d", otherUserFeed.ID), otherUserFeed.ID, "")
			err := fc.DeleteFeed(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteFeed() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteFeed() with other user's feed status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから直接確認
			var count int64
			feedDB.Model(&model.Feed{}).Where("id = ?", otherUserFeed.ID).Count(&count)
			if count != 1 {
				t.Error("DeleteFeed() incorrectly deleted other user's feed")
			}
		})
		
		t.Run("存在しないフィードIDでの削除はエラーになる", func(t *testing.T) {
			// テスト実行 - 存在しないフィードIDを指定
			_, c, rec := setupEchoWithFeedId(feedTestUser.ID, http.MethodDelete, "/feeds/"+fmt.Sprintf("%d", nonExistentFeedID), nonExistentFeedID, "")
			err := fc.DeleteFeed(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteFeed() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteFeed() with non-existent ID status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}
