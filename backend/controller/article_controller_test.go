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
	articleDB         *gorm.DB
	articleRepo       repository.IArticleRepository
	articleValidator  validator.IArticleValidator
	articleUsecase    usecase.IArticleUsecase
	ac                IArticleController
	articleTestUser   model.User
	articleOtherUser  model.User
)

const nonExistentArticleID uint = 9999

// テストセットアップ関数
func setupArticleControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if articleDB != nil {
		testutils.CleanupTestDB(articleDB)
	} else {
		// 初回のみデータベース接続を作成
		articleDB = testutils.SetupTestDB()
		articleRepo = repository.NewArticleRepository(articleDB)
		articleValidator = validator.NewArticleValidator()
		articleUsecase = usecase.NewArticleUsecase(articleRepo, articleValidator)
		ac = NewArticleController(articleUsecase)
	}
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDB)
	articleOtherUser = testutils.CreateOtherUser(articleDB)
}

// JWT認証をモックするヘルパー関数
func setupArticleEchoWithJWT(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
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
func setupArticleEchoWithJWTAndBody(userId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
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

// ArticleIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupArticleEchoWithArticleId(userId uint, method, path string, articleId uint, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupArticleEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("articleId")
	c.SetParamValues(fmt.Sprintf("%d", articleId))
	return e, c, rec
}

func TestArticleController_GetAllArticles(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーの記事を全て取得する", func(t *testing.T) {
			// テスト用記事の作成
			articles := []model.Article{
				{Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
				{Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
				{Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID}, // 別ユーザーの記事
			}
			
			for _, article := range articles {
				articleDB.Create(&article)
			}
			
			// テスト実行
			_, c, rec := setupArticleEchoWithJWT(articleTestUser.ID)
			err := ac.GetAllArticles(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllArticles() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response []model.ArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if len(response) != 2 {
				t.Errorf("GetAllArticles() returned %d articles, want 2", len(response))
			}
			
			// 記事タイトルの確認
			titles := make(map[string]bool)
			for _, article := range response {
				titles[article.Title] = true
			}
			
			if !titles["Article 1"] || !titles["Article 2"] {
				t.Errorf("期待した記事が結果に含まれていません: %v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// データベース接続エラーなどのケースをモックして追加可
	})
}

func TestArticleController_GetArticleById(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する記事を正しく取得する", func(t *testing.T) {
			// テスト用記事の作成
			article := model.Article{
				Title:   "Test Article",
				Content: "Test Content",
				UserId:  articleTestUser.ID,
			}
			articleDB.Create(&article)
			
			// テスト実行
			_, c, rec := setupArticleEchoWithArticleId(articleTestUser.ID, http.MethodGet, "/articles/"+fmt.Sprintf("%d", article.ID), article.ID, "")
			err := ac.GetArticleById(c)
			
			// 検証
			if err != nil {
				t.Errorf("GetArticleById() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetArticleById() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response model.ArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.ID != article.ID || response.Title != article.Title || response.Content != article.Content {
				t.Errorf("GetArticleById() = %v, want id=%d, title=%s, content=%s", response, article.ID, article.Title, article.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			
			// テスト実行 - testUserとして他のユーザーの記事にアクセス
			_, c, rec := setupArticleEchoWithArticleId(articleTestUser.ID, http.MethodGet, "/articles/"+fmt.Sprintf("%d", otherUserArticle.ID), otherUserArticle.ID, "")
			err := ac.GetArticleById(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("GetArticleById() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetArticleById() with other user's article status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestArticleController_CreateArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しい記事を作成できる", func(t *testing.T) {
			// テストリクエストの準備
			reqBody := `{"title":"New Test Article","content":"This is a test article content"}`
			
			// テスト実行
			_, c, rec := setupArticleEchoWithJWTAndBody(articleTestUser.ID, http.MethodPost, "/articles", reqBody)
			err := ac.CreateArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("CreateArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusCreated {
				t.Errorf("CreateArticle() status code = %d, want %d", rec.Code, http.StatusCreated)
			}
			
			// レスポンスボディをパース
			var response model.ArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.Title != "New Test Article" || response.Content != "This is a test article content" {
				t.Errorf("CreateArticle() = %v, want title=%s, content=%s", response, "New Test Article", "This is a test article content")
			}
			
			// データベースから直接確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, response.ID)
			if dbArticle.Title != "New Test Article" || dbArticle.UserId != articleTestUser.ID {
				t.Errorf("CreateArticle() did not save article correctly, got title=%s, userId=%d", dbArticle.Title, dbArticle.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する記事は作成できない", func(t *testing.T) {
			// タイトルが空のリクエスト
			reqBody := `{"title":"","content":"Invalid article with empty title"}`
			
			// テスト実行
			_, c, rec := setupArticleEchoWithJWTAndBody(articleTestUser.ID, http.MethodPost, "/articles", reqBody)
			err := ac.CreateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("CreateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("CreateArticle() with invalid title status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("JSONデコードエラーでバッドリクエストを返す", func(t *testing.T) {
			// 無効なJSON
			invalidJSON := `{"title": Invalid JSON`
			
			// テスト実行
			_, c, rec := setupArticleEchoWithJWTAndBody(articleTestUser.ID, http.MethodPost, "/articles", invalidJSON)
			err := ac.CreateArticle(c)
			
			// この場合はコントローラーがJSONレスポンスを返すので、
			// エラーオブジェクトではなくレスポンスのステータスコードを確認
			if err != nil {
				t.Errorf("CreateArticle() unexpected error: %v", err)
			}
			
			if rec.Code != http.StatusBadRequest {
				t.Errorf("CreateArticle() with invalid JSON status code = %d, want %d", 
					rec.Code, http.StatusBadRequest)
			}
		})
	})
}

func TestArticleController_UpdateArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("既存の記事を更新できる", func(t *testing.T) {
			// テスト用記事の作成
			article := model.Article{
				Title:   "Original Article",
				Content: "Original Content",
				UserId:  articleTestUser.ID,
			}
			articleDB.Create(&article)
			
			// 更新リクエストの準備
			updateReqBody := `{"title":"Updated Article","content":"Updated Content"}`
			
			// テスト実行
			_, c, rec := setupArticleEchoWithArticleId(articleTestUser.ID, http.MethodPut, "/articles/"+fmt.Sprintf("%d", article.ID), article.ID, updateReqBody)
			err := ac.UpdateArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("UpdateArticle() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response model.ArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			if response.ID != article.ID || response.Title != "Updated Article" || response.Content != "Updated Content" {
				t.Errorf("UpdateArticle() = %v, want id=%d, title=%s, content=%s", 
					response, article.ID, "Updated Article", "Updated Content")
			}
			
			// データベースから直接確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, article.ID)
			if dbArticle.Title != "Updated Article" || dbArticle.Content != "Updated Content" {
				t.Errorf("UpdateArticle() did not update article correctly, got title=%s, content=%s", 
					dbArticle.Title, dbArticle.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			
			// 更新リクエストの準備
			updateReqBody := `{"title":"Attempted Update","content":"Attempted Content Update"}`
			
			// テスト実行 - testUserとして他のユーザーの記事を更新しようとする
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodPut, 
				"/articles/"+fmt.Sprintf("%d", otherUserArticle.ID), 
				otherUserArticle.ID, 
				updateReqBody)
			err := ac.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with other user's article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースに変更が反映されていないことを確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, otherUserArticle.ID)
			if dbArticle.Title != "Other User's Article" {
				t.Errorf("UpdateArticle() should not update other user's article, but got title=%s", dbArticle.Title)
			}
		})
		
		t.Run("存在しないIDの記事は更新できない", func(t *testing.T) {
			// 更新リクエストの準備
			updateReqBody := `{"title":"Update Non-existent","content":"Update Content"}`
			
			// テスト実行 - 存在しないIDの記事を更新しようとする
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodPut, 
				"/articles/"+fmt.Sprintf("%d", nonExistentArticleID), 
				nonExistentArticleID, 
				updateReqBody)
			err := ac.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with non-existent article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("バリデーションエラーが発生する記事は更新できない", func(t *testing.T) {
			// テスト用記事の作成
			article := model.Article{
				Title:   "Article for Validation Test",
				Content: "Original Content",
				UserId:  articleTestUser.ID,
			}
			articleDB.Create(&article)
			
			// タイトルが空の無効なリクエスト
			invalidReqBody := `{"title":"","content":"Content with empty title"}`
			
			// テスト実行
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodPut, 
				"/articles/"+fmt.Sprintf("%d", article.ID), 
				article.ID, 
				invalidReqBody)
			err := ac.UpdateArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("UpdateArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UpdateArticle() with invalid article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースに変更が反映されていないことを確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, article.ID)
			if dbArticle.Title != "Article for Validation Test" {
				t.Errorf("UpdateArticle() should not update article with validation error, but got title=%s", dbArticle.Title)
			}
		})
	})
}

func TestArticleController_DeleteArticle(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分の記事を削除できる", func(t *testing.T) {
			// テスト用記事の作成
			article := model.Article{
				Title:   "Article to Delete",
				Content: "Content to Delete",
				UserId:  articleTestUser.ID,
			}
			articleDB.Create(&article)
			
			// テスト実行
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodDelete, 
				"/articles/"+fmt.Sprintf("%d", article.ID), 
				article.ID, 
				"")
			err := ac.DeleteArticle(c)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteArticle() error = %v", err)
			}
			
			if rec.Code != http.StatusNoContent {
				t.Errorf("DeleteArticle() status code = %d, want %d", rec.Code, http.StatusNoContent)
			}
			
			// データベースから削除されていることを確認
			var count int64
			articleDB.Model(&model.Article{}).Where("id = ?", article.ID).Count(&count)
			if count != 0 {
				t.Error("DeleteArticle() did not delete the article from database")
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("他のユーザーの記事は削除できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article to Not Delete",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			
			// テスト実行 - testUserとして他のユーザーの記事を削除しようとする
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodDelete, 
				"/articles/"+fmt.Sprintf("%d", otherUserArticle.ID), 
				otherUserArticle.ID, 
				"")
			err := ac.DeleteArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteArticle() with other user's article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
			
			// データベースから削除されていないことを確認
			var count int64
			articleDB.Model(&model.Article{}).Where("id = ?", otherUserArticle.ID).Count(&count)
			if count == 0 {
				t.Error("DeleteArticle() deleted other user's article from database")
			}
		})
		
		t.Run("存在しない記事の削除はエラーになる", func(t *testing.T) {
			// テスト実行 - 存在しないIDの記事を削除しようとする
			_, c, rec := setupArticleEchoWithArticleId(
				articleTestUser.ID, 
				http.MethodDelete, 
				"/articles/"+fmt.Sprintf("%d", nonExistentArticleID), 
				nonExistentArticleID, 
				"")
			err := ac.DeleteArticle(c)
			
			// コントローラーはエラーを返さずにJSONレスポンスとしてエラーメッセージを返す
			if err != nil {
				t.Errorf("DeleteArticle() returned unexpected error: %v", err)
			}
			
			// ステータスコードの確認
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("DeleteArticle() with non-existent article status code = %d, want %d", 
					rec.Code, http.StatusInternalServerError)
			}
		})
	})
}
