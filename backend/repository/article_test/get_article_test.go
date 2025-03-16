package article_test

import (
    "go-react-app/model"
    "testing"
)

func TestArticleRepository_GetAllArticles(t *testing.T) {
    setupArticleTest()
    
    // テスト用記事の作成
    articles := []model.ArticleRequest{
        {Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
        {Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
        {Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID},
    }
    
    for _, articleReq := range articles {
        article := articleReq.ToModel()
        articleDB.Create(&article)
    }
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("正しいユーザーIDの記事のみを取得する", func(t *testing.T) {
            var result []model.Article
            err := articleRepo.GetAllArticles(&result, articleTestUser.ID)
            
            if err != nil {
                t.Errorf("GetAllArticles() error = %v", err)
            }
            
            if len(result) != 2 {
                t.Errorf("GetAllArticles() got %d articles, want 2", len(result))
            }
            
            titles := make(map[string]bool)
            for _, article := range result {
                titles[article.Title] = true
            }
            
            if !titles["Article 1"] || !titles["Article 2"] {
                t.Errorf("期待した記事が結果に含まれていません: %v", result)
            }
        })
    })
}

func TestArticleRepository_GetArticleById(t *testing.T) {
    setupArticleTest()
    
    // テスト用記事の作成
    articleReq := model.ArticleRequest{
        Title:   "Test Get Article",
        Content: "Test Content for Get",
        UserId:  articleTestUser.ID,
    }
    article := articleReq.ToModel()
    articleDB.Create(&article)
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("IDで記事を取得できる", func(t *testing.T) {
            var result model.Article
            err := articleRepo.GetArticleById(&result, articleTestUser.ID, article.ID)
            
            if err != nil {
                t.Errorf("GetArticleById() error = %v", err)
            }
            
            if result.ID != article.ID || result.Title != article.Title || result.Content != article.Content {
                t.Errorf("GetArticleById() = %v, want %v", result, article)
            }
        })
    })
    
    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しない記事IDではエラーになる", func(t *testing.T) {
            var result model.Article
            err := articleRepo.GetArticleById(&result, articleTestUser.ID, nonExistentArticleID)
            
            if err == nil {
                t.Error("GetArticleById() with non-existent ID should return error")
            }
        })
        
        t.Run("他のユーザーの記事は取得できない", func(t *testing.T) {
            // 他のユーザーの記事を作成
            otherArticleReq := model.ArticleRequest{
                Title:   "Other User Article",
                Content: "Other User Content",
                UserId:  articleOtherUser.ID,
            }
            otherArticle := otherArticleReq.ToModel()
            articleDB.Create(&otherArticle)
            
            var result model.Article
            err := articleRepo.GetArticleById(&result, articleTestUser.ID, otherArticle.ID)
            
            if err == nil {
                t.Error("GetArticleById() with other user's article should return error")
            }
        })
    })
}