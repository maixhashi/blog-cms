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

// テスト用の共通変数
var (
	feedDB          *gorm.DB
	feedRepo        repository.IFeedRepository
	feedValidator   validator.IFeedValidator
	fu              IFeedUsecase // feedUsecase -> fu に変更
	feedTestUser    model.User
	feedOtherUser   model.User
)

const nonExistentFeedID uint = 9999

func setupFeedTest() {
	// テストごとにデータベースをクリーンアップ
	if feedDB != nil {
		testutils.CleanupTestDB(feedDB)
	} else {
		// 初回のみデータベース接続を作成
		feedDB = testutils.SetupTestDB()
		feedRepo = repository.NewFeedRepository(feedDB)
		feedValidator = validator.NewFeedValidator()
		fu = NewFeedUsecase(feedRepo, feedValidator)
	}
	
	// テストユーザーを作成
	feedTestUser = testutils.CreateTestUser(feedDB)
	
	// 別のテストユーザーを作成
	feedOtherUser = testutils.CreateOtherUser(feedDB)
}

func TestFeedUsecase_GetAllFeeds(t *testing.T) {
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
			t.Logf("ユーザーID %d のフィードを取得します", feedTestUser.ID)
			
			feedResponses, err := fu.GetAllFeeds(feedTestUser.ID)
			
			if err != nil {
				t.Errorf("GetAllFeeds() error = %v", err)
			}
			
			if len(feedResponses) != 2 {
				t.Errorf("GetAllFeeds() got %d feeds, want 2", len(feedResponses))
			}
			
			// フィードタイトルの確認
			titles := make(map[string]bool)
			for _, feed := range feedResponses {
				titles[feed.Title] = true
				t.Logf("取得したフィード: ID=%d, Title=%s", feed.ID, feed.Title)
			}
			
			if !titles["Feed 1"] || !titles["Feed 2"] {
				t.Errorf("期待したフィードが結果に含まれていません: %v", feedResponses)
			}
			
			// レスポンス形式の検証
			for _, feed := range feedResponses {
				if feed.ID == 0 || feed.Title == "" || feed.CreatedAt.IsZero() || feed.UpdatedAt.IsZero() {
					t.Errorf("GetAllFeeds() returned invalid feed: %+v", feed)
				}
			}
		})
	})
}

func TestFeedUsecase_GetFeedById(t *testing.T) {
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
			
			response, err := fu.GetFeedById(feedTestUser.ID, feed.ID)
			
			if err != nil {
				t.Errorf("GetFeedById() error = %v", err)
			}
			
			if response.ID != feed.ID || response.Title != feed.Title || response.URL != feed.URL {
				t.Errorf("取得したフィード %+v が、期待値 %+v と一致しません", response, feed)
			} else {
				t.Logf("正常に取得: ID=%d, Title=%s", response.ID, response.Title)
			}
			
			// レスポンス形式の検証
			if response.ID == 0 || response.Title == "" || response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Errorf("GetFeedById() returned invalid feed: %+v", response)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
			t.Logf("存在しないID %d を指定してフィードを取得しようとします", nonExistentFeedID)
			
			_, err := fu.GetFeedById(feedTestUser.ID, nonExistentFeedID)
			
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
			
			_, err := fu.GetFeedById(feedTestUser.ID, otherUserFeed.ID)
			
			if err == nil {
				t.Error("他のユーザーのフィードを取得できてしまいました")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	})
}

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
			response, err := fu.CreateFeed(validFeed)
			
			// 検証
			if err != nil {
				t.Errorf("CreateFeed() error = %v", err)
			}
			
			if response.ID == 0 || response.Title != validFeed.Title || response.URL != validFeed.URL {
				t.Errorf("CreateFeed() returned unexpected response: %+v", response)
			} else {
				t.Logf("生成されたフィードID: %d", response.ID)
			}
			
			// タイムスタンプが設定されていることを確認
			if response.CreatedAt.IsZero() || response.UpdatedAt.IsZero() {
				t.Error("CreateFeed() did not set timestamps in response")
			} else {
				t.Logf("タイムスタンプ設定済み: CreatedAt=%v, UpdatedAt=%v", response.CreatedAt, response.UpdatedAt)
			}
			
			// データベースから直接確認
			var savedFeed model.Feed
			feedDB.First(&savedFeed, response.ID)
			
			if savedFeed.Title != validFeed.Title || savedFeed.URL != validFeed.URL || savedFeed.UserId != feedTestUser.ID {
				t.Errorf("CreateFeed() = %v, want title=%s, url=%s, userId=%d", savedFeed, validFeed.Title, validFeed.URL, feedTestUser.ID)
			} else {
				t.Logf("データベース保存確認: Title=%s, URL=%s, UserId=%d", savedFeed.Title, savedFeed.URL, savedFeed.UserId)
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
			
			_, err := fu.CreateFeed(invalidFeed)
			
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

func TestFeedUsecase_UpdateFeed(t *testing.T) {
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
			response, err := fu.UpdateFeed(updatedFeed, feedTestUser.ID, feed.ID)
			
			// 検証
			if err != nil {
				t.Errorf("UpdateFeed() error = %v", err)
			} else {
				t.Log("フィード更新成功")
			}
			
			// 返り値のフィードが更新されていることを確認
			if response.ID != feed.ID || response.Title != updatedFeed.Title || response.URL != updatedFeed.URL {
				t.Errorf("UpdateFeed() = %+v, want id=%d, title=%s, url=%s", response, feed.ID, updatedFeed.Title, updatedFeed.URL)
			} else {
				t.Logf("返り値確認: ID=%d, Title=%s, URL=%s", response.ID, response.Title, response.URL)
			}
			
			// データベースから直接確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, feed.ID)
			
			if dbFeed.Title != updatedFeed.Title || dbFeed.URL != updatedFeed.URL {
				t.Errorf("UpdateFeed() database feed = %v, want title=%v, url=%v", dbFeed, updatedFeed.Title, updatedFeed.URL)
			} else {
				t.Logf("データベース更新確認: Title=%s, URL=%s", dbFeed.Title, dbFeed.URL)
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
			t.Logf("無効なタイトルでの更新を試行: %s", invalidUpdate.Title)
			
			_, err := fu.UpdateFeed(invalidUpdate, feedTestUser.ID, feed.ID)
			
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
		
			_, err := fu.UpdateFeed(updateAttempt, feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("存在しないIDでの更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのフィードは更新できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other-feed",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			t.Logf("他ユーザーのフィード: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
		
			// 他ユーザーのフィードを更新しようとする
			updateAttempt := model.Feed{Title: "Attempted Update"}
			_, err := fu.UpdateFeed(updateAttempt, feedTestUser.ID, otherUserFeed.ID)
		
			if err == nil {
				t.Error("他のユーザーのフィード更新でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		
			// データベースに反映されていないことを確認
			var dbFeed model.Feed
			feedDB.First(&dbFeed, otherUserFeed.ID)
			if dbFeed.Title != otherUserFeed.Title {
				t.Errorf("他ユーザーのフィードが変更されています: %s → %s", otherUserFeed.Title, dbFeed.Title)
			} else {
				t.Log("他ユーザーのフィードは変更されていないことを確認")
			}
		})
	})
}

func TestFeedUsecase_DeleteFeed(t *testing.T) {
	setupFeedTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("自分のフィードを削除できる", func(t *testing.T) {
			// テスト用フィードの作成
			feed := model.Feed{
				Title:  "Feed to Delete",
				URL:    "https://example.com/delete",
				UserId: feedTestUser.ID,
			}
			feedDB.Create(&feed)
			t.Logf("削除対象フィード作成: ID=%d, Title=%s", feed.ID, feed.Title)
		
			// テスト実行
			err := fu.DeleteFeed(feedTestUser.ID, feed.ID)
		
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
		
			err := fu.DeleteFeed(feedTestUser.ID, nonExistentFeedID)
			if err == nil {
				t.Error("存在しないIDでの削除でエラーが返されませんでした")
			} else {
				t.Logf("期待通りエラーが返されました: %v", err)
			}
		})
	
		t.Run("他のユーザーのフィードは削除できない", func(t *testing.T) {
			// 他のユーザーのフィードを作成
			otherUserFeed := model.Feed{
				Title:  "Other User's Feed",
				URL:    "https://example.com/other-feed-delete",
				UserId: feedOtherUser.ID,
			}
			feedDB.Create(&otherUserFeed)
			t.Logf("他ユーザーのフィード作成: ID=%d, Title=%s, UserId=%d", otherUserFeed.ID, otherUserFeed.Title, otherUserFeed.UserId)
		
			// 他ユーザーのフィードを削除しようとする
			err := fu.DeleteFeed(feedTestUser.ID, otherUserFeed.ID)
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
