package article_test

import (
    "go-react-app/model"
    "testing"
    "time"
)

func TestArticleRepository_UpdateArticle(t *testing.T) {
    setupArticleTest()
    
    article := model.Article{
        Title:   "Original Title",
        Content: "Original Content",
        UserId:  articleTestUser.ID,
    }
    articleDB.Create(&article)
    
    time.Sleep(10 * time.Millisecond)
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("記事のタイトルと内容を更新できる", func(t *testing.T) {
            updatedArticle := model.Article{
                Title:   "Updated Title",
                Content: "Updated Content",
            }
            
            err := articleRepo.UpdateArticle(&updatedArticle, articleTestUser.ID, article.ID)
            
            if err != nil {
                t.Errorf("UpdateArticle() error = %v", err)
            }
            
            if updatedArticle.Title != "Updated Title" || updatedArticle.Content != "Updated Content" {
                t.Errorf("UpdateArticle() returned unexpected content")
            }
            
            var dbArticle model.Article
            articleDB.First(&dbArticle, article.ID)
            
            if dbArticle.Title != "Updated Title" || dbArticle.Content != "Updated Content" {
                t.Errorf("UpdateArticle() database update failed")
            }
        })
    })

    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
            invalidArticle := model.Article{Title: "Invalid Update"}
            err := articleRepo.UpdateArticle(&invalidArticle, articleTestUser.ID, nonExistentArticleID)
            
            if err == nil {
                t.Error("UpdateArticle() should return error for non-existent ID")
            }
        })

        t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
            otherUserArticle := model.Article{
                Title:   "Other User's Article",
                Content: "Other Content",
                UserId:  articleOtherUser.ID,
            }
            articleDB.Create(&otherUserArticle)
            
            updateAttempt := model.Article{Title: "Attempted Update"}
            err := articleRepo.UpdateArticle(&updateAttempt, articleTestUser.ID, otherUserArticle.ID)
            
            if err == nil {
                t.Error("UpdateArticle() should not allow updating other user's article")
            }
        })
    })
}
