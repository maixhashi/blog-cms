package article_test

import (
	"fmt"
	"go-react-app/controller"
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/testutils"
	"go-react-app/usecase"
	"go-react-app/validator"
	"net/http"
	"net/http/httptest"
	"strings"

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
	articleController controller.IArticleController
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
		articleController = controller.NewArticleController(articleUsecase)
	}
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDB)
	articleOtherUser = testutils.CreateOtherUser(articleDB)
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

// ArticleIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithArticleId(userId uint, articleId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("articleId")
	c.SetParamValues(fmt.Sprintf("%d", articleId))
	return e, c, rec
}
