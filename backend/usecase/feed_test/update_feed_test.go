package feed_test

import (
	"go-react-app/model"
	"testing"
	"time"
)

func TestFeedUsecase_UpdateFeed(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feed := createTestFeed("Original Feed", "https://example.com/original", feedTestUser.ID)
	t.Logf("元のフィード作成: ID=%d, Title=%s", feed.ID, feed.Title)
	
	// 少し待って更新時間に差をつける
	time.Sleep(10 * time.Millisecond)
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("フィードのタイトルとURLを更新できる", func(t *testing.T) {
			// 更新するフィード
			updatedFeed := model.Feed{
				Title: "Updated Feed",
				URL:   "https://example.com/updated",
			}
			
			t.Logf("フィード更新リクエスト: ID=%d, 新Title=%s, 新URL=%s", feed.ID, updatedFeed.Title, updatedFeed.URL)
			
			// テスト実行
			response, err := feedUsecase.UpdateFeed(updatedFeed, feedTestUser.ID, feed.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateFeed() error = %v", err)
			} else {
				t.Log("フィード更新成功")
			}
			
			// 返り値のフィードが更新されていることを確認
			if validateFeedResponse(t, response, updatedFeed.Title, updatedFeed.URL) {
				t.Logf("返り値確認: ID=%d, Title=%s, URL=%s", response.ID, response.Title, response.URL)
			}
			
			// データベースから直接確認
			if validateDatabaseFeed(t, feed.ID, updatedFeed.Title, updatedFeed.URL, feedTestUser.ID) {
				t.Logf("データベース更新確認: Title=%s, URL=%s", updatedFeed.Title, updatedFeed.URL)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("バリデーションエラーが発生するフィードは更新できない", func(t *testing.T) {
			// 無効な更新 - タイトルなし
			invalidUpdate := model.Feed{
				Title: "",
				URL:   "https://example.com/invalid-update",
			}
			
			t.Logf("無効なタイトルでの更新を試行: %s", invalidUpdate.Title)
			
			_, err := feedUsecase.UpdateFeed(invalidUpdate, feedTestUser.ID, feed.ID)
			
			// バリデーションエラーが発生するはず
			if err == nil {
				t.Error("無効なタイトルでエラーが返されませんでした")
			} else {
				t.Logf("期待通りバリデーションエラーが返されました: %v", err)
			}
			
			// データベースに反映されていないことを確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, feed.ID)
			if dbFeed.Title == invalidUpdate.Title {
				t.Error("バリデーションエラーの更新がデータベースに反映されています")
			} else {
				t.Logf("データベース確認: Title=%s (変更されていない)", dbFeed.Title)
			}
		})
		
		t.Run("存在しないフィードIDでの更新はエラーになる", func(t *testing.T) {
			updateAttempt := model.Feed{Title: "Valid Title"}
			t.Logf("存在しないID %d でフィード更新を試行", nonExistentFeedID)
		
			_, err := feedUsecase.UpdateFeed(updateAttempt, feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのフィードは更新できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := createTestFeed("Other User's Feed", "https://example.com/other-feed", feedOtherUser.ID)
			t.Logf("他ユーザーのフィード: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
		
			// 他ユーザーのフィードを更新しようとする
			updateAttempt := model.Feed{Title: "Attempted Update"}
			_, err := feedUsecase.UpdateFeed(updateAttempt, feedTestUser.ID, otherUserFeed.ID)
		
			if err == nil {
				t.Error("他のユーザーのフィード更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに反映されていないことを確認
			if validateDatabaseFeed(t, otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.URL, feedOtherUser.ID) {
				t.Log("他ユーザーのフィードは変更されていないことを確認")
			}
		})
	})
}