package repository

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
	"time"

	"gorm.io/gorm"
)

// 変数名にPrefixをつけて区別する
var articleDB *gorm.DB
var articleRepo IArticleRepository
var articleTestUser model.User
var articleOtherUser model.User
const nonExistentArticleID uint = 9999

func setupArticleTest() {
	// 毎回新しいデータベース接続を作成
	articleDB = testutils.SetupTestDB()
	articleRepo = NewArticleRepository(articleDB)
	
	// テストユーザーを作成
	articleTestUser = testutils.CreateTestUser(articleDB)
	
	// 別のテストユーザーを作成
	articleOtherUser = testutils.CreateOtherUser(articleDB)
}

func TestArticleRepository_GetAllArticles(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	articles := []model.Article{
		{Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
		{Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
		{Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID}, // 別ユーザーの記事
	}
	
	for _, article := range articles {
		articleDB.Create(&article)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDの記事のみを取得する", func(t *testing.T) {
			var result []model.Article
			err := articleRepo.GetAllArticles(&result, articleTestUser.ID)
			
			t.Logf("ユーザーID %d の記事を取得します", articleTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetAllArticles() got %d articles, want 2", len(result))
			}
			
			// 記事タイトルの確認
			titles := make(map[string]bool)
			for _, article := range result {
				titles[article.Title] = true
				t.Logf("取得した記事: ID=%d, Title=%s", article.ID, article.Title)
			}
			
			if !titles["Article 1"] || !titles["Article 2"] {
				t.Errorf("期待した記事が結果に含まれていません: %v", result)
			}
		})
	})
}

func TestArticleRepository_GetArticleById(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	article := model.Article{
		Title:   "Test Article",
		Content: "Test Content",
		UserId:  articleTestUser.ID,
	}
	articleDB.Create(&article)

	t.Run("正常系", func(t *testing.T) {
		t.Run("存在する記事を正しく取得する", func(t *testing.T) {
			t.Logf("記事ID %d を取得します", article.ID)
	
			var result model.Article
			err := articleRepo.GetArticleById(&result, articleTestUser.ID, article.ID)
	
			if err != nil {
				t.Errorf("GetArticleById() error = %v", err)
			}
	
			if result.ID != article.ID || result.Title != article.Title {
				t.Errorf("取得した記事 %+v が、期待値 %+v と一致しません", result, article)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", result.ID, result.Title)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定して記事を取得しようとします", nonExistentArticleID)
	
			var notFound model.Article
			err := articleRepo.GetArticleById(&notFound, articleTestUser.ID, nonExistentArticleID)
	
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:  "Other User's Article",
				Content: "Other Content",
				UserId: articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserArticle.ID, articleTestUser.ID)
			
			var result model.Article
			err := articleRepo.GetArticleById(&result, articleTestUser.ID, otherUserArticle.ID)
			
			if err == nil {
				t.Error("他のユーザーの記事を取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

func TestArticleRepository_CreateArticle(t *testing.T) {
	setupArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しい記事を作成できる", func(t *testing.T) {
			// テスト用記事
			article := model.Article{
				Title:   "New Article",
				Content: "New Content",
				UserId:  articleTestUser.ID,
			}
			
			t.Logf("記事作成: Title=%s, UserId=%d", article.Title, article.UserId)
			
			// テスト実行
			err := articleRepo.CreateArticle(&article)
			
			// 検証
			if err != nil {
				t.Errorf("CreateArticle() error = %v", err)
			}
			
			if article.ID == 0 {
				t.Error("CreateArticle() did not set ID")
			} else {
				t.Logf("生成された記事ID: %d", article.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if article.CreatedAt.IsZero() || article.UpdatedAt.IsZero() {
				t.Error("CreateArticle() did not set timestamps")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", article.CreatedAt, article.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedArticle model.Article
			articleDB.First(&savedArticle, article.ID)
			
			if savedArticle.Title != "New Article" || savedArticle.UserId != articleTestUser.ID {
				t.Errorf("CreateArticle() = %v, want title=%s, userId=%d", savedArticle, "New Article", articleTestUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, UserId=%d", savedArticle.Title, savedArticle.UserId)
			}
		})
	})
}

func TestArticleRepository_UpdateArticle(t *testing.T) {
	setupArticleTest()
	
	// テストデータの作成
	article := model.Article{
		Title:   "Original Title",
		Content: "Original Content",
		UserId:  articleTestUser.ID,
	}
	articleDB.Create(&article)
	t.Logf("元の記事作成: ID=%d, Title=%s", article.ID, article.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("記事のタイトルと内容を更新できる", func(t *testing.T) {
			// 更新する記事
			updatedArticle := model.Article{
				Title:   "Updated Title",
				Content: "Updated Content",
			}
			
			t.Logf("記事更新リクエスト: ID=%d, 新Title=%s", article.ID, updatedArticle.Title)
			
			// テスト実行
			err := articleRepo.UpdateArticle(&updatedArticle, articleTestUser.ID, article.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateArticle() error = %v", err)
			} else {
				t.Log("記事更新成功")
			}
			
			// 返り値の記事が更新されていることを確認
			if updatedArticle.Title != "Updated Title" || updatedArticle.Content != "Updated Content" {
				t.Errorf("UpdateArticle() returned article title = %v, content = %v, want title = %v, content = %v", 
					updatedArticle.Title, updatedArticle.Content, "Updated Title", "Updated Content")
			} else {
				t.Logf("返り値確認: Title=%s, Content=%s", updatedArticle.Title, updatedArticle.Content)
			}
			
			// データベースから直接確認
			var dbArticle model.Article
			articleDB.First(&dbArticle, article.ID)
			
			if dbArticle.Title != "Updated Title" || dbArticle.Content != "Updated Content" {
				t.Errorf("UpdateArticle() database article title = %v, content = %v, want title = %v, content = %v",
					dbArticle.Title, dbArticle.Content, "Updated Title", "Updated Content")
			} else {
				t.Logf("データベース更新確認: Title=%s, Content=%s", dbArticle.Title, dbArticle.Content)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事IDでの更新はエラーになる", func(t *testing.T) {
			invalidArticle := model.Article{Title: "Invalid Update"}
			t.Logf("存在しないID %d で記事更新を試行", nonExistentArticleID)
			
			err := articleRepo.UpdateArticle(&invalidArticle, articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("UpdateArticle() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
				Title:   "Other User's Article",
				Content: "Other Content",
				UserId:  articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
			
			// 他ユーザーの記事を更新しようとする
			updateAttempt := model.Article{Title: "Attempted Update"}
			err := articleRepo.UpdateArticle(&updateAttempt, articleTestUser.ID, otherUserArticle.ID)
			
			if err == nil {
				t.Error("UpdateArticle() should not allow updating other user's article")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}

func TestArticleRepository_DeleteArticle(t *testing.T) {
	setupArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分の記事を削除できる", func(t *testing.T) {
			// テストデータの作成
			article := model.Article{
				Title:   "Article to Delete",
				Content: "Content to Delete",
				UserId:  articleTestUser.ID,
			}
			articleDB.Create(&article)
			t.Logf("削除対象記事作成: ID=%d, Title=%s", article.ID, article.Title)
			
			// テスト実行
			err := articleRepo.DeleteArticle(articleTestUser.ID, article.ID)
			
			// 検証
			if err != nil {
				t.Errorf("DeleteArticle() error = %v", err)
			} else {
				t.Logf("記事削除成功: ID=%d", article.ID)
			}
			
			// データベースから直接確認
			var count int64
			articleDB.Model(&model.Article{}).Where("id = ?", article.ID).Count(&count)
			
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
			
			err := articleRepo.DeleteArticle(articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("DeleteArticle() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーの記事を削除しようとするとエラー", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := model.Article{
					Title:   "Other User's Article",
					Content: "Other Content",
					UserId:  articleOtherUser.ID,
			}
			articleDB.Create(&otherUserArticle)
			t.Logf("他ユーザーの記事作成: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
			
			// 他ユーザーの記事を削除しようとする
			err := articleRepo.DeleteArticle(articleTestUser.ID, otherUserArticle.ID)
			if err == nil {
					t.Error("DeleteArticle() should not allow deleting other user's article")
			} else {
					t.Logf("期待通りエラーが返された: %v", err)
			}
			
			// データベースに残っていることを確認
			var count int64
			articleDB.Model(&model.Article{}).Where("id = ?", otherUserArticle.ID).Count(&count)
			if count == 0 {
					t.Error("他ユーザーの記事が削除されています")
			} else {
					t.Log("他ユーザーの記事は削除されていないことを確認")
			}
	})
})
}
