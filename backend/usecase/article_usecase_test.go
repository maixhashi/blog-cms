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
	articleDb         *gorm.DB
	articleRepo       repository.IArticleRepository
	articleValidator  validator.IArticleValidator
	au                IArticleUsecase
	articleTestUser   model.User
	articleOtherUser  model.User
)

const nonExistentArticleID uint = 9999

func setupArticleTest() {
	// テストごとにデータベースをクリーンアップ
	if articleDb != nil {
		testutils.CleanupTestDB(articleDb)
	} else {
		// 初回のみデータベース接続を作成
		articleDb = testutils.SetupTestDB()
		articleRepo = repository.NewArticleRepository(articleDb)
		articleValidator = validator.NewArticleValidator()
		au = NewArticleUsecase(articleRepo, articleValidator)
	}
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDb)
	
	// 別のテストユーザーを作成
	articleOtherUser = testutils.CreateOtherUser(articleDb)
}

func TestArticleUsecase_GetAllArticles(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	articles := []model.Article{
		{Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
		{Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
		{Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID}, // 別ユーザーの記事
	}
	
	for _, article := range articles {
		articleDb.Create(&article)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDの記事のみを取得する", func(t *testing.T) {
			t.Logf("ユーザーID %d の記事を取得します", articleTestUser.ID)
			
			articleResponses, err := au.GetAllArticles(articleTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if len(articleResponses) != 2 {
				t.Errorf("GetAllArticles() got %d articles, want 2", len(articleResponses))
			}
			
			// 記事タイトルの確認
			titles := make(map[string]bool)
			for _, article := range articleResponses {
				titles[article.Title] = true
				t.Logf("取得した記事: ID=%d, Title=%s", article.ID, article.Title)
			}
			
			if !titles["Article 1"] || !titles["Article 2"] {
				t.Errorf("期待した記事が結果に含まれていません: %v", articleResponses)
			}
			
			// レスポンス形式の検証
			for _, article := range articleResponses {
				if article.ID == 0 || article.Title == "" || article.CreatedAt.IsZero() || article.UpdatedAt.IsZero() {
					t.Errorf("GetAllArticles() returned invalid article: %+v", article)
				}
			}
		})
	})
}

func TestArticleUsecase_GetArticleById(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	article := model.Article{
		Title:   "Test Article",
		Content: "Test Content",
		UserId:  articleTestUser.ID,
	}
	articleDb.Create(&article)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する記事を正しく取得する", func(t *testing.T) {
			t.Logf("記事ID %d を取得します", article.ID)
			
			response, err := au.GetArticleById(articleTestUser.ID, article.ID)
			
			if err != nil {
				t.Errorf("GetArticleById() error = %v", err)
			}
			
			if response.ID != article.ID || response.Title != article.Title || response.Content != article.Content {
				t.Errorf("取得した記事 %+v が、期待値 %+v と一致しません", response, article)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", response.ID, response.Title)
			}
			
			// レスポンス形式の検証
			if response.ID == 0 || response.Title == "" || response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Errorf("GetArticleById() returned invalid article: %+v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定して記事を取得しようとします", nonExistentArticleID)
			
			_, err := au.GetArticleById(articleTestUser.ID, nonExistentArticleID)
			
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDb.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserArticle.ID, articleTestUser.ID)
			
			_, err := au.GetArticleById(articleTestUser.ID, otherUserArticle.ID)
			
			if err == nil {
				t.Error("他のユーザーの記事を取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

func TestArticleUsecase_CreateArticle(t *testing.T) {
	setupArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しい記事を作成できる", func(t *testing.T) {
			// テスト用記事
			validArticle := model.Article{
				Title:   "New Test Article",
				Content: "This is a test article content",
				UserId:  articleTestUser.ID,
			}
			
			t.Logf("記事作成: Title=%s, UserId=%d", validArticle.Title, validArticle.UserId)
			
			// テスト実行
			response, err := au.CreateArticle(validArticle)
			
			// 検証
			if err != nil {
				t.Errorf("CreateArticle() error = %v", err)
			}
			
			if response.ID == 0 || response.Title != validArticle.Title || response.Content != validArticle.Content {
				t.Errorf("CreateArticle() returned unexpected response: %+v", response)
			} else {
				t.Logf("生成された記事ID: %d", response.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Error("CreateArticle() did not set timestamps in response")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", response.CreatedAt, response.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedArticle model.Article
			articleDb.First(&savedArticle, response.ID)
			
			if savedArticle.Title != validArticle.Title || savedArticle.Content != validArticle.Content || savedArticle.UserId != articleTestUser.ID {
				t.Errorf("CreateArticle() = %v, want title=%s, content=%s, userId=%d", savedArticle, validArticle.Title, validArticle.Content, articleTestUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, Content=%s, UserId=%d", savedArticle.Title, savedArticle.Content, savedArticle.UserId)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する記事は作成できない", func(t *testing.T) {
			// 無効な記事（タイトルなし）
			invalidArticle := model.Article{
				Title:   "", // 空のタイトル
				Content: "Invalid article content",
				UserId:  articleTestUser.ID,
			}
			
			t.Logf("無効な記事作成を試行: Title=%s (空)", invalidArticle.Title)
			
			_, err := au.CreateArticle(invalidArticle)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効な記事でエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに保存されていないことを確認
			var count int64
			articleDb.Model(&model.Article{}).Where("content = ? AND title = ?", invalidArticle.Content, invalidArticle.Title).Count(&count)
			if count > 0 {
				t.Error("バリデーションエラーの記事がデータベースに保存されています")
			} else {
				t.Log("バリデーションエラーの記事は保存されていないことを確認")
			}
		})
	})
}

func TestArticleUsecase_UpdateArticle(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	article := model.Article{
		Title:   "Original Article",
		Content: "Original Content",
		UserId:  articleTestUser.ID,
	}
	articleDb.Create(&article)
	t.Logf("元の記事作成: ID=%d, Title=%s", article.ID, article.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("記事のタイトルと内容を更新できる", func(t *testing.T) {
			// 更新する記事
			updatedArticle := model.Article{
				Title:   "Updated Article Title",
				Content: "Updated Article Content",
			}
			
			t.Logf("記事更新リクエスト: ID=%d, 新Title=%s", article.ID, updatedArticle.Title)
			
			// テスト実行
			response, err := au.UpdateArticle(updatedArticle, articleTestUser.ID, article.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateArticle() error = %v", err)
			} else {
				t.Log("記事更新成功")
			}
			
			// 返り値の記事が更新されていることを確認
			if response.ID != article.ID || response.Title != updatedArticle.Title || response.Content != updatedArticle.Content {
				t.Errorf("UpdateArticle() = %+v, want id=%d, title=%s, content=%s", response, article.ID, updatedArticle.Title, updatedArticle.Content)
			} else {
				t.Logf("返り値確認: ID=%d, Title=%s, Content=%s", response.ID, response.Title, response.Content)
			}
			
			// データベースから直接確認
			var dbArticle model.Article
			articleDb.First(&dbArticle, article.ID)
			
			if dbArticle.Title != updatedArticle.Title || dbArticle.Content != updatedArticle.Content {
				t.Errorf("UpdateArticle() database article title = %v, content = %v, want title = %v, content = %v", 
					dbArticle.Title, dbArticle.Content, updatedArticle.Title, updatedArticle.Content)
			} else {
				t.Logf("データベース更新確認: Title=%s, Content=%s", dbArticle.Title, dbArticle.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生する記事は更新できない", func(t *testing.T) {
			// 無効な更新（空のタイトル）
			invalidUpdate := model.Article{
				Title:   "", // 空のタイトル
				Content: "Valid Content",
			}
			
			t.Logf("無効なタイトルでの更新を試行: タイトルが空")
			
			_, err := au.UpdateArticle(invalidUpdate, articleTestUser.ID, article.ID)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効な記事でエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに反映されていないことを確認
			var dbArticle model.Article
			articleDb.First(&dbArticle, article.ID)
			if dbArticle.Title == invalidUpdate.Title {
				t.Error("バリデーションエラーの更新がデータベースに反映されています")
			} else {
				t.Logf("データベース確認: Title=%s (変更されていない)", dbArticle.Title)
			}
		})
		
		t.Run("存在しないタスクIDでの更新はエラーになる", func(t *testing.T) {
			updateAttempt := model.Article{Title: "Valid Title", Content: "Valid Content"}
			t.Logf("存在しないID %d で記事更新を試行", nonExistentArticleID)
		
			_, err := au.UpdateArticle(updateAttempt, articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDb.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
		
			// 他ユーザーの記事を更新しようとする
			updateAttempt := model.Article{Title: "Attempted Update", Content: "Attempted Content"}
			_, err := au.UpdateArticle(updateAttempt, articleTestUser.ID, otherUserArticle.ID)
		
			if err == nil {
				t.Error("他のユーザーの記事更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに反映されていないことを確認
			var dbArticle model.Article
			articleDb.First(&dbArticle, otherUserArticle.ID)
			if dbArticle.Title != otherUserArticle.Title {
				t.Errorf("他ユーザーの記事が変更されています: %s → %s", otherUserArticle.Title, dbArticle.Title)
			} else {
				t.Log("他ユーザーの記事は変更されていないことを確認")
			}
		})
	})
}

func TestArticleUsecase_DeleteArticle(t *testing.T) {
	setupArticleTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("自分の記事を削除できる", func(t *testing.T) {
			// テスト用記事の作成
			article := model.Article{
				Title:   "Article to Delete",
				Content: "Content to Delete",
				UserId:  articleTestUser.ID,
			}
			articleDb.Create(&article)
			t.Logf("削除対象記事作成: ID=%d, Title=%s", article.ID, article.Title)
		
			// テスト実行
			err := au.DeleteArticle(articleTestUser.ID, article.ID)
		
			// 検証
			if err != nil {
				t.Errorf("DeleteArticle() error = %v", err)
			} else {
				t.Logf("記事削除成功: ID=%d", article.ID)
			}
		
			// データベースから直接確認
			var count int64
			articleDb.Model(&model.Article{}).Where("id = ?", article.ID).Count(&count)
			if count != 0 {
				t.Error("DeleteArticle() did not delete the article from database")
			} else {
				t.Log("データベースから記事が削除されていることを確認")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事IDでの削除はエラーになる", func(t *testing.T) {
			t.Logf("存在しないID %d で記事削除を試行", nonExistentArticleID)
		
			err := au.DeleteArticle(articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーの記事は削除できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other User's Content",
				UserId:  articleOtherUser.ID,
			}
			articleDb.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事作成: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
		
			// 他ユーザーの記事を削除しようとする
			err := au.DeleteArticle(articleTestUser.ID, otherUserArticle.ID)
			if err == nil {
				t.Error("他のユーザーの記事削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに残っていることを確認
			var count int64
			articleDb.Model(&model.Article{}).Where("id = ?", otherUserArticle.ID).Count(&count)
			if count == 0 {
				t.Error("他ユーザーの記事が削除されています")
			} else {
				t.Log("他ユーザーの記事は削除されていないことを確認")
			}
		})
	})
}
