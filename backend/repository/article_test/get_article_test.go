package article_test

import (
    "go-react-app/model"
    "testing"
)

func TestArticleRepository_GetAllArticles(t *testing.T) {
    setupArticleTest()
    
    articles := []model.Article{
        {Title: "Article 1", Content: "Content 1", UserId: articleTestUser.ID},
        {Title: "Article 2", Content: "Content 2", UserId: articleTestUser.ID},
        {Title: "Article 3", Content: "Content 3", UserId: articleOtherUser.ID},
    }
    
    for _, article := range articles {
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
