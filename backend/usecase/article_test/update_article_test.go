package article_test

import (
	"go-react-app/model"
	"testing"
	"time"
)

func TestArticleUsecase_UpdateArticle(t *testing.T) {
	setupArticleUsecaseTest()
	
	// テストデータの作成
	article := createTestArticle(t, "Original Article", "Original Content", articleTestUser.ID)
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
			response, err := articleUsecase.UpdateArticle(updatedArticle, articleTestUser.ID, article.ID)
			
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
			verifyDatabaseArticle(t, article.ID, updatedArticle.Title, updatedArticle.Content, articleTestUser.ID)
			t.Logf("データベース更新確認: Title=%s, Content=%s", updatedArticle.Title, updatedArticle.Content)
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
			
			_, err := articleUsecase.UpdateArticle(invalidUpdate, articleTestUser.ID, article.ID)
			
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
		
			_, err := articleUsecase.UpdateArticle(updateAttempt, articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle(t, "Other User's Article", "Other User's Content", articleOtherUser.ID)
			t.Logf("他ユーザーの記事: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
		
			// 他ユーザーの記事を更新しようとする
			updateAttempt := model.Article{Title: "Attempted Update", Content: "Attempted Content"}
			_, err := articleUsecase.UpdateArticle(updateAttempt, articleTestUser.ID, otherUserArticle.ID)
		
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