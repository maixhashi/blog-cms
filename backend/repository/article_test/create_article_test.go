package article_test

import (
    "go-react-app/model"
    "testing"
)

func TestArticleRepository_CreateArticle(t *testing.T) {
    setupArticleTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("新しい記事を作成できる", func(t *testing.T) {
            article := model.Article{
                Title:   "New Article",
                Content: "New Content",
                UserId:  articleTestUser.ID,
            }
            
            err := articleRepo.CreateArticle(&article)
            
            if err != nil {
                t.Errorf("CreateArticle() error = %v", err)
            }
            
            if article.ID == 0 {
                t.Error("CreateArticle() did not set ID")
            }
            
            if article.CreatedAt.IsZero() || article.UpdatedAt.IsZero() {
                t.Error("CreateArticle() did not set timestamps")
            }
        })
    })
}
