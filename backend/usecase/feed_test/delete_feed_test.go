package feed_test

import (
	"go-react-app/model"
	"testing"
)

func TestFeedUsecase_DeleteFeed(t *testing.T) {
	setupFeedTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のフィードを削除できる", func(t *testing.T) {
			// テスト用フィードの作成
			feed := createTestFeed("Feed to Delete", "https://example.com/delete", feedTestUser.ID)
			t.Logf("削除対象フィード作成: ID=%d, Title=%s", feed.ID, feed.Title)
		
			// テスト実行
			err := feedUsecase.DeleteFeed(feedTestUser.ID, feed.ID)
		
			// 検証
			if err != nil {
				t.Errorf("DeleteFeed() error = %v", err)
			} else {
				t.Logf("フィード削除成功: ID=%d", feed.ID)
			}
		
			// データベースから直接確認
			var count int64
			feedDB.Model(&model.Feed{}).Where("id = ?", feed.ID).Count(&count)
			if count != 0 {
				t.Error("DeleteFeed() did not delete the feed from database")
			} else {
				t.Log("データベースからフィードが削除されていることを確認")
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないフィードIDでの削除はエラーになる", func(t *testing.T) {
			t.Logf("存在しないID %d でフィード削除を試行", nonExistentFeedID)
		
			err := feedUsecase.DeleteFeed(feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのフィードは削除できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := createTestFeed("Other User's Feed", "https://example.com/other-feed-delete", feedOtherUser.ID)
			t.Logf("他ユーザーのフィード作成: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
		
			// 他ユーザーのフィードを削除しようとする
			err := feedUsecase.DeleteFeed(feedTestUser.ID, otherUserFeed.ID)
			if err == nil {
				t.Error("他のユーザーのフィード削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに残っていることを確認
			var count int64
			feedDB.Model(&model.Feed{}).Where("id = ?", otherUserFeed.ID).Count(&count)
			if count == 0 {
				t.Error("他ユーザーのフィードが削除されています")
			} else {
				t.Log("他ユーザーのフィードは削除されていないことを確認")
			}
		})
	})
}
