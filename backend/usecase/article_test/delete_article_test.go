package article_test

import (
	"testing"
)

func TestArticleUsecase_DeleteArticle(t *testing.T) {
	setupArticleUsecaseTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("自分の記事を削除できる", func(t *testing.T) {
			// テスト用記事の作成
			article := createTestArticle(t, "Article to Delete", "Content to Delete", articleTestUser.ID)
			t.Logf("削除対象記事作成: ID=%d, Title=%s", article.ID, article.Title)
		
			// テスト実行
			err := articleUsecase.DeleteArticle(articleTestUser.ID, article.ID)
		
			// 検証
			if err != nil {
				t.Errorf("DeleteArticle() error = %v", err)
			} else {
				t.Logf("記事削除成功: ID=%d", article.ID)
			}
		
			// データベースから直接確認
			if articleExists(article.ID) {
				t.Error("DeleteArticle() did not delete the article from database")
			} else {
				t.Log("データベースから記事が削除されていることを確認")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事IDでの削除はエラーになる", func(t *testing.T) {
			t.Logf("存在しないID %d で記事削除を試行", nonExistentArticleID)
		
			err := articleUsecase.DeleteArticle(articleTestUser.ID, nonExistentArticleID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーの記事は削除できない", func(t *testing.T) {
			// 他のユーザーの記事を作成
			otherUserArticle := createTestArticle(t, "Other User's Article", "Other User's Content", articleOtherUser.ID)
			t.Logf("他ユーザーの記事作成: ID=%d, Title=%s, UserId=%d", otherUserArticle.ID, otherUserArticle.Title, otherUserArticle.UserId)
		
			// 他ユーザーの記事を削除しようとする
			err := articleUsecase.DeleteArticle(articleTestUser.ID, otherUserArticle.ID)
			if err == nil {
				t.Error("他のユーザーの記事削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに残っていることを確認
			if !articleExists(otherUserArticle.ID) {
				t.Error("他ユーザーの記事が削除されています")
			} else {
				t.Log("他ユーザーの記事は削除されていないことを確認")
			}
		})
	})
}