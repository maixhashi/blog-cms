package repository

import (
	"go-react-app/model"
	"go-react-app/testutils"
	"testing"
	"time"

	"gorm.io/gorm"
)

var (
	feedDB           *gorm.DB
	feedRepo         IFeedRepository
	feedTestUser     model.User
	feedOtherUser    model.User
)
const nonExistentFeedID uint = 9999
func setupFeedTest() {
	// テストごとにデータベースをクリーンアップ
	if feedDB != nil {
		testutils.CleanupTestDB(feedDB)
	} else {
		// 初回のみデータベース接続を作成
		feedDB = testutils.SetupTestDB()
		feedRepo = NewFeedRepository(feedDB)
	}
	
	// テストユーザーを作成
	feedTestUser = testutils.CreateTestUser(feedDB)
	
	// 別のテストユーザーを作成
	feedOtherUser = testutils.CreateOtherUser(feedDB)
}

func TestFeedRepository_GetAllFeeds(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feeds := []model.Feed{
		{Title: "Feed 1", URL: "https://example.com/feed1", UserId: feedTestUser.ID},
		{Title: "Feed 2", URL: "https://example.com/feed2", UserId: feedTestUser.ID},
		{Title: "Feed 3", URL: "https://example.com/feed3", UserId: feedOtherUser.ID}, // 別ユーザーのフィード
	}
	
	for _, feed := range feeds {
		feedDB.Create(&feed)
	}
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("正しいユーザーIDのフィードのみを取得する", func(t *testing.T) {
			var result []model.Feed
			err := feedRepo.GetAllFeeds(&result, feedTestUser.ID)
			
			t.Logf("ユーザーID %d のフィードを取得します", feedTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllFeeds() error = %v", err)
			}
			
			if len(result) != 2 {
				t.Errorf("GetAllFeeds() got %d feeds, want 2", len(result))
			}
			
			// フィードタイトルの確認
			titles := make(map[string]bool)
			for _, feed := range result {
				titles[feed.Title] = true
				t.Logf("取得したフィード: ID=%d, Title=%s", feed.ID, feed.Title)
			}
			
			if !titles["Feed 1"] || !titles["Feed 2"] {
				t.Errorf("期待したフィードが結果に含まれていません: %v", result)
			}
		})
	})
}

func TestFeedRepository_GetFeedById(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feed := model.Feed{
		Title:  "Test Feed",
		URL:    "https://example.com/test",
		UserId: feedTestUser.ID,
	}
	feedDB.Create(&feed)

	t.Run("正常系", func(t *testing.T) {
		t.Run("存在するフィードを正しく取得する", func(t *testing.T) {
			t.Logf("フィードID %d を取得します", feed.ID)
	
			var result model.Feed
			err := feedRepo.GetFeedById(&result, feedTestUser.ID, feed.ID)
	
			if err != nil {
				t.Errorf("GetFeedById() error = %v", err)
			}
	
			if result.ID != feed.ID || result.Title != feed.Title {
				t.Errorf("取得したフィード %+v が、期待値 %+v と一致しません", result, feed)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", result.ID, result.Title)
			}
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してフィードを取得しようとします", nonExistentFeedID)
	
			var notFound model.Feed
			err := feedRepo.GetFeedById(&notFound, feedTestUser.ID, nonExistentFeedID)
	
			if err == nil {
				t.Error("存在しないIDを指定したときにエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
		
		t.Run("他のユーザーのフィードは取得できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			t.Logf("他ユーザーのフィード(ID=%d)を別ユーザー(ID=%d)として取得しようとします", otherUserFeed.ID, feedTestUser.ID)
			
			var result model.Feed
			err := feedRepo.GetFeedById(&result, feedTestUser.ID, otherUserFeed.ID)
			
			if err == nil {
				t.Error("他のユーザーのフィードを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

func TestFeedRepository_CreateFeed(t *testing.T) {
	setupFeedTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("新しいフィードを作成できる", func(t *testing.T) {
			// テスト用フィード
			feed := model.Feed{
				Title:  "New Feed",
				URL:    "https://example.com/new",
				UserId: feedTestUser.ID,
			}
			
			t.Logf("フィード作成: Title=%s, URL=%s, UserId=%d", feed.Title, feed.URL, feed.UserId)
			
			// テスト実行
			err := feedRepo.CreateFeed(&feed)
			
			// 検証
			if err != nil {
				t.Errorf("CreateFeed() error = %v", err)
			}
			
			if feed.ID == 0 {
				t.Error("CreateFeed() did not set ID")
			} else {
				t.Logf("生成されたフィードID: %d", feed.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if feed.CreatedAt.IsZero() || feed.UpdatedAt.IsZero() {
				t.Error("CreateFeed() did not set timestamps")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", feed.CreatedAt, feed.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedFeed model.Feed
			feedDB.First(&savedFeed, feed.ID)
			
			if savedFeed.Title != "New Feed" || savedFeed.URL != "https://example.com/new" || savedFeed.UserId != feedTestUser.ID {
				t.Errorf("CreateFeed() = %v, want title=%s, url=%s, userId=%d", savedFeed, "New Feed", "https://example.com/new", feedTestUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, URL=%s, UserId=%d", savedFeed.Title, savedFeed.URL, savedFeed.UserId)
			}
		})
	})
}

func TestFeedRepository_UpdateFeed(t *testing.T) {
	setupFeedTest()
	
	// テストデータの作成
	feed := model.Feed{
		Title:  "Original Feed",
		URL:    "https://example.com/original",
		UserId: feedTestUser.ID,
	}
	feedDB.Create(&feed)
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
			err := feedRepo.UpdateFeed(&updatedFeed, feedTestUser.ID, feed.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateFeed() error = %v", err)
			} else {
				t.Log("フィード更新成功")
			}
			
			// 返り値のフィードが更新されていることを確認
			if updatedFeed.Title != "Updated Feed" || updatedFeed.URL != "https://example.com/updated" {
				t.Errorf("UpdateFeed() returned feed title = %v, url = %v, want title = %v, url = %v", 
					updatedFeed.Title, updatedFeed.URL, "Updated Feed", "https://example.com/updated")
			} else {
				t.Logf("返り値確認: Title=%s, URL=%s", updatedFeed.Title, updatedFeed.URL)
			}
			
			// データベースから直接確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, feed.ID)
			
			if dbFeed.Title != "Updated Feed" || dbFeed.URL != "https://example.com/updated" {
				t.Errorf("UpdateFeed() database feed title = %v, url = %v, want title = %v, url = %v", 
					dbFeed.Title, dbFeed.URL, "Updated Feed", "https://example.com/updated")
			} else {
				t.Logf("データベース更新確認: Title=%s, URL=%s", dbFeed.Title, dbFeed.URL)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないフィードIDでの更新はエラーになる", func(t *testing.T) {
			invalidFeed := model.Feed{Title: "Invalid Update", URL: "https://example.com/invalid"}
			t.Logf("存在しないID %d でフィード更新を試行", nonExistentFeedID)
			
			err := feedRepo.UpdateFeed(&invalidFeed, feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("UpdateFeed() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーのフィードは更新できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			t.Logf("他ユーザーのフィード: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
			
			// 他ユーザーのフィードを更新しようとする
			updateAttempt := model.Feed{Title: "Attempted Update", URL: "https://example.com/attempt"}
			err := feedRepo.UpdateFeed(&updateAttempt, feedTestUser.ID, otherUserFeed.ID)
			
			if err == nil {
				t.Error("UpdateFeed() should not allow updating other user's feed")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}

func TestFeedRepository_DeleteFeed(t *testing.T) {
	setupFeedTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のフィードを削除できる", func(t *testing.T) {
			// テストデータの作成
			feed := model.Feed{
				Title:  "Feed to Delete",
				URL:    "https://example.com/delete",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			t.Logf("削除対象フィード作成: ID=%d, Title=%s", feed.ID, feed.Title)
			
			// テスト実行
			err := feedRepo.DeleteFeed(feedTestUser.ID, feed.ID)
			
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
			
			err := feedRepo.DeleteFeed(feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("DeleteFeed() with non-existent ID should return error")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
		
		t.Run("他のユーザーのフィードを削除しようとするとエラー", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other Users Feed",
				URL:    "https://example.com/other",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			t.Logf("他ユーザーのフィード作成: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
			
			// 他ユーザーのフィードを削除しようとする
			err := feedRepo.DeleteFeed(feedTestUser.ID, otherUserFeed.ID)
			if err == nil {
				t.Error("DeleteFeed() should not allow deleting other user's feed")
			} else {
				t.Logf("期待通りエラーが返された: %v", err)
			}
		})
	})
}
