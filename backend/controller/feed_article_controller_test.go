package controller

import (
	"encoding/json"
	"fmt"
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// テスト用の共通変数 - 変数名を変更して衝突を避ける
var (
	articleDB                *gorm.DB
	articleFeedRepo          repository.IFeedRepository
	articleFeedArticleRepo   repository.IFeedArticleRepository
	articleFeedArticleUcase  usecase.IFeedArticleUsecase
	feedArticleCtrl          IFeedArticleController // 変数名を変更
	articleTestUser          model.User
	articleOtherUser         model.User
	articleTestFeed          model.Feed
	articleOtherUserFeed     model.Feed
)

// テストセットアップ関数
func setupArticleControllerTest() {
	// テストごとにデータベースをクリーンアップ
	if articleDB != nil {
		testutils.CleanupTestDB(articleDB)
	} else {
		// 初回のみデータベース接続を作成
		articleDB = testutils.SetupTestDB()
		articleFeedRepo = repository.NewFeedRepository(articleDB)
		articleFeedArticleRepo = repository.NewFeedArticleRepository(articleFeedRepo)
		articleFeedArticleUcase = usecase.NewFeedArticleUsecase(articleFeedArticleRepo)
		feedArticleCtrl = NewFeedArticleController(articleFeedArticleUcase) // 変数名を変更
	}
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDB)
	articleOtherUser = testutils.CreateOtherUser(articleDB)
	
	// テスト用フィードを作成 - 実際のRSSフィードURLを使用
	articleTestFeed = model.Feed{
		Title:       "Test Feed",
		URL:         "https://blog.golang.org/feed.atom", // Golangの公式ブログフィード
		Description: "Test Feed Description",
		UserId:      articleTestUser.ID,
	}
	articleDB.Create(&articleTestFeed)
	
	// 別ユーザーのテスト用フィードを作成
	articleOtherUserFeed = model.Feed{
		Title:       "Other User Feed",
		URL:         "https://blog.golang.org/feed.atom", // 同じURLでも別ユーザー
		Description: "Other User Feed Description",
		UserId:      articleOtherUser.ID,
	}
	articleDB.Create(&articleOtherUserFeed)
}

// JWT認証をセットアップするヘルパー関数
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

// フィードIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupArticleEchoWithFeedId(userId uint, feedId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupArticleEchoWithJWT(userId)
	c.SetParamNames("feedId")
	c.SetParamValues(fmt.Sprintf("%d", feedId))
	return e, c, rec
}

// フィードIDと記事IDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupArticleEchoWithFeedAndArticleId(userId uint, feedId uint, articleId string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupArticleEchoWithFeedId(userId, feedId)
	c.SetParamNames("feedId", "articleId")
	c.SetParamValues(fmt.Sprintf("%d", feedId), articleId)
	return e, c, rec
}

func TestFeedArticleController_GetAllArticles(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("ユーザーの全フィードの記事を取得する", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupArticleEchoWithJWT(articleTestUser.ID)
			err := feedArticleCtrl.GetAllArticles(c) // 変数名を変更
			
			// 検証
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetAllArticles() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response []model.FeedArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			// 通常はGolangのブログフィードから少なくとも1つ以上の記事が取得できるはず
			// 外部依存があるため、記事数の正確な検証は難しいが、基本的な検証は行う
			if len(response) == 0 {
				t.Logf("記事が取得できませんでした。外部フィードに問題がある可能性があります。")
			} else {
				t.Logf("取得した記事数: %d", len(response))
				
				// レスポンスのフォーマットを確認
				for i, article := range response {
					if article.Title == "" || article.URL == "" {
						t.Errorf("記事 %d のフォーマットが不正: %+v", i, article)
					}
				}
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		// 外部依存があるため、異常系のテストは限定的
		// 実際のアプリケーションでは、モックを使用してより詳細なテストを行うことが望ましい
	})
}

func TestFeedArticleController_GetArticlesByFeedID(t *testing.T) {
	setupArticleControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("フィードIDから記事一覧を取得する", func(t *testing.T) {
			// テスト実行
			_, c, rec := setupArticleEchoWithFeedId(articleTestUser.ID, articleTestFeed.ID)
			err := feedArticleCtrl.GetArticlesByFeedID(c) // 変数名を変更
			
			// 検証
			if err != nil {
				t.Errorf("GetArticlesByFeedID() error = %v", err)
			}
			
			if rec.Code != http.StatusOK {
				t.Errorf("GetArticlesByFeedID() status code = %d, want %d", rec.Code, http.StatusOK)
			}
			
			// レスポンスボディをパース
			var response []model.FeedArticleResponse
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}
			
			// 外部依存があるため、記事数の正確な検証は難しいが、基本的な検証は行う
			if len(response) == 0 {
				t.Logf("記事が取得できませんでした。外部フィードに問題がある可能性があります。")
			} else {
				t.Logf("取得した記事数: %d", len(response))
				
				// レスポンスのフォーマットを確認
				for i, article := range response {
					if article.Title == "" || article.URL == "" || article.FeedID != articleTestFeed.ID {
						t.Errorf("記事 %d のフォーマットが不正: %+v", i, article)
					}
				}
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("無効なフィードIDを指定した場合はエラーを返す", func(t *testing.T) {
			// 無効なフィードID
			invalidFeedID := uint(99999)
			
			_, c, rec := setupArticleEchoWithFeedId(articleTestUser.ID, invalidFeedID)
			err := feedArticleCtrl.GetArticlesByFeedID(c) // 変数名を変更
			
			// コントローラーはJSONレスポンスを返すのでエラーオブジェクトを返さない
			if err != nil {
				t.Errorf("GetArticlesByFeedID() unexpected error: %v", err)
			}
			
			// エラーはHTTPステータスコードで判断
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetArticlesByFeedID() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("他のユーザーのフィードにアクセスするとエラーを返す", func(t *testing.T) {
			// 別ユーザーのフィードIDを指定
			_, c, rec := setupArticleEchoWithFeedId(articleTestUser.ID, articleOtherUserFeed.ID)
			err := feedArticleCtrl.GetArticlesByFeedID(c) // 変数名を変更
			
			// コントローラーはJSONレスポンスを返すのでエラーオブジェクトを返さない
			if err != nil {
				t.Errorf("GetArticlesByFeedID() unexpected error: %v", err)
			}
			
			// エラーはHTTPステータスコードで判断
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("GetArticlesByFeedID() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}

func TestFeedArticleController_GetArticleByID(t *testing.T) {
	setupArticleControllerTest()
	
	// このテストは外部依存があるため、先に記事一覧を取得してから特定の記事IDをテストに使用
	var testArticleID string
	
	// 事前に記事を取得してテストに使用する記事IDを決定
	t.Run("事前準備", func(t *testing.T) {
		_, c, rec := setupArticleEchoWithFeedId(articleTestUser.ID, articleTestFeed.ID)
		err := feedArticleCtrl.GetArticlesByFeedID(c) // 変数名を変更
		
		if err != nil || rec.Code != http.StatusOK {
			t.Skipf("事前準備でフィード記事を取得できませんでした: %v", err)
			return
		}
		
		var articles []model.FeedArticleResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &articles); err != nil {
			t.Skipf("事前準備で記事一覧のパースに失敗: %v", err)
			return
		}
		
		if len(articles) == 0 {
			t.Skipf("事前準備で記事が取得できませんでした")
			return
		}
		
		// 最初の記事のIDを使用
		testArticleID = articles[0].ID
		t.Logf("テスト用記事ID: %s", testArticleID)
	})
	
	// 記事IDが取得できた場合のみテストを実行
	if testArticleID != "" {
		t.Run("正常系", func(t *testing.T) {
			t.Run("フィードIDと記事IDから特定の記事を取得する", func(t *testing.T) {
				// テスト実行
				_, c, rec := setupArticleEchoWithFeedAndArticleId(articleTestUser.ID, articleTestFeed.ID, testArticleID)
				err := feedArticleCtrl.GetArticleByID(c)
				
				// 検証
				if err != nil {
					t.Errorf("GetArticleByID() error = %v", err)
				}
				
				if rec.Code != http.StatusOK {
					t.Errorf("GetArticleByID() status code = %d, want %d", rec.Code, http.StatusOK)
				}
				
				// レスポンスボディをパース
				var article model.FeedArticleResponse
				err = json.Unmarshal(rec.Body.Bytes(), &article)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				
				// レスポンスの検証
				if article.ID != testArticleID {
					t.Errorf("GetArticleByID() returned wrong article, got ID=%s, want ID=%s", article.ID, testArticleID)
				}
				
				if article.Title == "" || article.URL == "" || article.FeedID != articleTestFeed.ID {
					t.Errorf("GetArticleByID() returned invalid article: %+v", article)
				} else {
					t.Logf("記事を正常に取得: ID=%s, Title=%s", article.ID, article.Title)
				}
			})
		})
		
		t.Run("異常系", func(t *testing.T) {
			t.Run("存在しない記事IDを指定した場合はエラーを返す", func(t *testing.T) {
				// 存在しない記事ID
				nonExistingArticleID := "non-existing-article-id"
				
				_, c, rec := setupArticleEchoWithFeedAndArticleId(articleTestUser.ID, articleTestFeed.ID, nonExistingArticleID)
				err := feedArticleCtrl.GetArticleByID(c)
				
				// コントローラーはJSONレスポンスを返すのでエラーオブジェクトを返さない
				if err != nil {
					t.Errorf("GetArticleByID() unexpected error: %v", err)
				}
				
				// エラーはHTTPステータスコードで判断
				if rec.Code != http.StatusInternalServerError {
					t.Errorf("GetArticleByID() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
				}
			})
			
			t.Run("他のユーザーのフィードの記事にアクセスするとエラーを返す", func(t *testing.T) {
				// 存在する記事IDを使って別ユーザーのフィードからアクセス
				_, c, rec := setupArticleEchoWithFeedAndArticleId(articleTestUser.ID, articleOtherUserFeed.ID, testArticleID)
				err := feedArticleCtrl.GetArticleByID(c)
				
				// コントローラーはJSONレスポンスを返すのでエラーオブジェクトを返さない
				if err != nil {
					t.Errorf("GetArticleByID() unexpected error: %v", err)
				}
				
				// エラーはHTTPステータスコードで判断
				if rec.Code != http.StatusInternalServerError {
					t.Errorf("GetArticleByID() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
				}
			})
			
			t.Run("無効なフィードIDパラメータを指定した場合はエラーを返す", func(t *testing.T) {
				// echoコンテキストを設定
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/feeds/invalid/articles/"+testArticleID, nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				
				// JWTトークンを設定
				token := jwt.New(jwt.SigningMethodHS256)
				claims := token.Claims.(jwt.MapClaims)
				claims["user_id"] = float64(articleTestUser.ID)
				c.Set("user", token)
				
				// 無効なフィードIDパラメータを設定
				c.SetParamNames("feedId", "articleId")
				c.SetParamValues("invalid", testArticleID)
				
				err := feedArticleCtrl.GetArticleByID(c)
				
				// コントローラーはJSONレスポンスを返すのでエラーオブジェクトを返さない
				if err != nil {
					t.Errorf("GetArticleByID() unexpected error: %v", err)
				}
				
				// エラーはHTTPステータスコードで判断
				if rec.Code != http.StatusBadRequest {
					t.Errorf("GetArticleByID() status code = %d, want %d", rec.Code, http.StatusBadRequest)
				}
			})
		})
	} else {
		t.Log("記事IDが取得できなかったため、GetArticleByIDのテストをスキップします")
	}
}
