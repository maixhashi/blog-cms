package usecase

import (
	"errors"
	"go-react-app/model"
	"reflect"
	"testing"
	"time"
)

// モックリポジトリの定義
type mockFeedArticleRepository struct {
	articles         map[uint][]model.FeedArticle // フィードIDごとの記事マップ
	allArticles      []model.FeedArticle          // すべての記事
	shouldReturnErr  bool
	getUserArticleErr bool
}

func (m *mockFeedArticleRepository) GetArticlesByFeedID(userId uint, feedID uint) ([]model.FeedArticle, error) {
	if m.shouldReturnErr {
		return nil, errors.New("フィードの取得に失敗しました")
	}
	
	if articles, ok := m.articles[feedID]; ok {
		return articles, nil
	}
	return []model.FeedArticle{}, nil
}

func (m *mockFeedArticleRepository) GetArticleByID(userId uint, feedID uint, articleID string) (model.FeedArticle, error) {
	if m.shouldReturnErr {
		return model.FeedArticle{}, errors.New("記事の取得に失敗しました")
	}
	
	articles, ok := m.articles[feedID]
	if !ok {
		return model.FeedArticle{}, errors.New("フィードが見つかりません")
	}
	
	for _, article := range articles {
		if article.ID == articleID {
			return article, nil
		}
	}
	
	return model.FeedArticle{}, errors.New("記事が見つかりません")
}

func (m *mockFeedArticleRepository) GetAllArticles(userId uint) ([]model.FeedArticle, error) {
	if m.getUserArticleErr {
		return nil, errors.New("記事の取得に失敗しました")
	}
	return m.allArticles, nil
}

// テスト用変数
var (
	mockRepo         *mockFeedArticleRepository
	feedArticleUc    IFeedArticleUsecase
	testUserId       uint = 1
)

// テスト用記事データ
var testArticles = []model.FeedArticle{
	{
		ID:          "article1",
		FeedID:      1,
		Title:       "Test Article 1",
		URL:         "https://example.com/article1",
		Summary:     "Summary of article 1",
		Categories:  []string{"tech", "news"},
		PublishedAt: time.Now(),
		Author:      "Test Author",
	},
	{
		ID:          "article2",
		FeedID:      1,
		Title:       "Test Article 2",
		URL:         "https://example.com/article2",
		Summary:     "Summary of article 2",
		Categories:  []string{"science"},
		PublishedAt: time.Now(),
		Author:      "Another Author",
	},
	{
		ID:          "article3",
		FeedID:      2,
		Title:       "Test Article 3",
		URL:         "https://example.com/article3",
		Summary:     "Summary of article 3",
		Categories:  []string{"politics"},
		PublishedAt: time.Now(),
		Author:      "Third Author",
	},
}

func setupFeedArticleTest() {
	mockRepo = &mockFeedArticleRepository{
		articles: map[uint][]model.FeedArticle{
			1: {testArticles[0], testArticles[1]},
			2: {testArticles[2]},
		},
		allArticles: testArticles,
		shouldReturnErr: false,
		getUserArticleErr: false,
	}
	
	feedArticleUc = NewFeedArticleUsecase(mockRepo)
}

func TestFeedArticleUsecase_GetArticlesByFeedID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("フィードの記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			responses, err := feedArticleUc.GetArticlesByFeedID(testUserId, 1)
			
			// 検証
			if err != nil {
				t.Errorf("GetArticlesByFeedID() error = %v", err)
			}
			
			if len(responses) != 2 {
				t.Errorf("GetArticlesByFeedID() got %d articles, want 2", len(responses))
			}
			
			// レスポンスの確認
			expectedFirst := model.FeedArticleResponse{
				ID:          testArticles[0].ID,
				FeedID:      testArticles[0].FeedID,
				Title:       testArticles[0].Title,
				URL:         testArticles[0].URL,
				Summary:     testArticles[0].Summary,
				Categories:  testArticles[0].Categories,
				PublishedAt: testArticles[0].PublishedAt,
				Author:      testArticles[0].Author,
			}
			
			if !reflect.DeepEqual(responses[0], expectedFirst) {
				t.Errorf("GetArticlesByFeedID() response mismatch: got %+v, want %+v", responses[0], expectedFirst)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("フィードの取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.shouldReturnErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetArticlesByFeedID(testUserId, 1)
			
			// 検証
			if err == nil {
				t.Error("GetArticlesByFeedID() should return error")
			}
			
			// モックをリセット
			mockRepo.shouldReturnErr = false
		})
	})
}

func TestFeedArticleUsecase_GetArticleByID(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("指定した記事IDの記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			response, err := feedArticleUc.GetArticleByID(testUserId, 1, "article1")
			
			// 検証
			if err != nil {
				t.Errorf("GetArticleByID() error = %v", err)
			}
			
			expectedResponse := model.FeedArticleResponse{
				ID:          testArticles[0].ID,
				FeedID:      testArticles[0].FeedID,
				Title:       testArticles[0].Title,
				URL:         testArticles[0].URL,
				Summary:     testArticles[0].Summary,
				Categories:  testArticles[0].Categories,
				PublishedAt: testArticles[0].PublishedAt,
				Author:      testArticles[0].Author,
			}
			
			if !reflect.DeepEqual(response, expectedResponse) {
				t.Errorf("GetArticleByID() response mismatch: got %+v, want %+v", response, expectedResponse)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しない記事IDを指定した場合はエラーを返す", func(t *testing.T) {
			// テスト実行
			_, err := feedArticleUc.GetArticleByID(testUserId, 1, "non-existent-id")
			
			// 検証
			if err == nil {
				t.Error("GetArticleByID() should return error for non-existent article ID")
			}
		})
		
		t.Run("フィードの取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.shouldReturnErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetArticleByID(testUserId, 1, "article1")
			
			// 検証
			if err == nil {
				t.Error("GetArticleByID() should return error when feed fetch fails")
			}
			
			// モックをリセット
			mockRepo.shouldReturnErr = false
		})
	})
}

func TestFeedArticleUsecase_GetAllArticles(t *testing.T) {
	setupFeedArticleTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("すべての記事を正しく取得できる", func(t *testing.T) {
			// テスト実行
			responses, err := feedArticleUc.GetAllArticles(testUserId)
			
			// 検証
			if err != nil {
				t.Errorf("GetAllArticles() error = %v", err)
			}
			
			if len(responses) != 3 {
				t.Errorf("GetAllArticles() got %d articles, want 3", len(responses))
			}
			
			// すべての記事のIDを確認
			articleIDs := make(map[string]bool)
			for _, article := range responses {
				articleIDs[article.ID] = true
			}
			
			expectedIDs := []string{"article1", "article2", "article3"}
			for _, id := range expectedIDs {
				if !articleIDs[id] {
					t.Errorf("GetAllArticles() missing article with ID: %s", id)
				}
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("記事の取得に失敗した場合はエラーを返す", func(t *testing.T) {
			// モックリポジトリにエラーを返すよう設定
			mockRepo.getUserArticleErr = true
			
			// テスト実行
			_, err := feedArticleUc.GetAllArticles(testUserId)
			
			// 検証
			if err == nil {
				t.Error("GetAllArticles() should return error")
			}
			
			// モックをリセット
			mockRepo.getUserArticleErr = false
		})
	})
}
