package feed_test

import (
	"go-react-app/model"
	"testing"
)

func TestFeedUsecase_CreateFeed(t *testing.T) {
	setupFeedTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいフィードを作成できる", func(t *testing.T) {
			// テスト用フィード
			validFeed := model.Feed{
				Title:  "New Feed",
				URL:    "https://example.com/new",
				UserId: feedTestUser.ID,
			}
			
			t.Logf("フィード作成: Title=%s, URL=%s, UserId=%d", validFeed.Title, validFeed.URL, validFeed.UserId)
			
			// テスト実行
			response, err := feedUsecase.CreateFeed(validFeed)
			
			// 検証
			if err != nil {
				t.Errorf("CreateFeed() error = %v", err)
			}
			
			if validateFeedResponse(t, response, validFeed.Title, validFeed.URL) {
				t.Logf("生成されたフィードID: %d", response.ID)
			}
			
			// データベースから直接確認
			if validateDatabaseFeed(t, response.ID, validFeed.Title, validFeed.URL, feedTestUser.ID) {
				t.Logf("データベース保存確認: Title=%s, URL=%s, UserId=%d", validFeed.Title, validFeed.URL, feedTestUser.ID)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するフィードは作成できない", func(t *testing.T) {
			// 無効なフィード - タイトルなし
			invalidFeed := model.Feed{
				Title:  "",
				URL:    "https://example.com/invalid",
				UserId: feedTestUser.ID,
			}
			
			t.Logf("無効なフィード作成を試行: Title=%s", invalidFeed.Title)
			
			_, err := feedUsecase.CreateFeed(invalidFeed)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なフィードでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに保存されていないことを確認
			var count int64
			feedDB.Model(&model.Feed{}).Where("url = ?", invalidFeed.URL).Count(&count)
			if count > 0 {
				t.Error("バリデーションエラーのフィードがデータベースに保存されています")
			} else {
				t.Log("バリデーションエラーのフィードは保存されていないことを確認")
			}
		})
	})
}
