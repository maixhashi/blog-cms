package article_test

import (
    "go-react-app/model"
    "testing"
)

func TestArticleRepository_CreateArticle(t *testing.T) {
    setupArticleTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("新しい記事を作成できる", func(t *testing.T) {
            // ArticleRequestを使用
            articleReq := model.ArticleRequest{
                Title:   "New Article",
                Content: "New Content",
                UserId:  articleTestUser.ID,
            }
            
            article := articleReq.ToModel()
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
            
            // データベースから記事を取得して確認
            var dbArticle model.Article
            articleDB.First(&dbArticle, article.ID)
            
            if dbArticle.Title != articleReq.Title || dbArticle.Content != articleReq.Content {
                t.Errorf("CreateArticle() database record doesn't match: got=%v, want title=%s, content=%s",
                    dbArticle, articleReq.Title, articleReq.Content)
            }
        })
    })
}