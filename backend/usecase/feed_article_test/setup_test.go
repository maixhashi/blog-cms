package feed_article_test

import (
	"errors"
	"go-react-app/model"
	"go-react-app/usecase"
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
	feedArticleUc    usecase.IFeedArticleUsecase
)

// テスト環境のセットアップ
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
	
	feedArticleUc = usecase.NewFeedArticleUsecase(mockRepo)
}
