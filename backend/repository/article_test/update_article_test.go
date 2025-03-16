package article_test

import (
    "go-react-app/model"
    "testing"
    "time"
)

func TestArticleRepository_UpdateArticle(t *testing.T) {
    setupArticleTest()
    
    // ArticleRequestを使用して記事を作成
    originalReq := model.ArticleRequest{
        Title:   "Original Title",
        Content: "Original Content",
        UserId:  articleTestUser.ID,
    }
    
    article := originalReq.ToModel()
    articleDB.Create(&article)
    
    time.Sleep(10 * time.Millisecond) // タイムスタンプの違いを確認するため
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("記事のタイトルと内容を更新できる", func(t *testing.T) {
            // 更新用のArticleRequestを作成
            updateReq := model.ArticleRequest{
                Title:   "Updated Title",
                Content: "Updated Content",
                UserId:  articleTestUser.ID,
            }
            
            updatedArticle := updateReq.ToModel()
            err := articleRepo.UpdateArticle(&updatedArticle, articleTestUser.ID, article.ID)
            
            if err != nil {
                t.Errorf("UpdateArticle() error = %v", err)
            }
            
            if updatedArticle.Title != updateReq.Title || updatedArticle.Content != updateReq.Content {
                t.Errorf("UpdateArticle() returned unexpected content")
            }
            
            // データベースから記事を取得して確認
            var dbArticle model.Article
            articleDB.First(&dbArticle, article.ID)
            
            if dbArticle.Title != updateReq.Title || dbArticle.Content != updateReq.Content {
                t.Errorf("UpdateArticle() database update failed")
            }
            
            // 更新日時が変更されていることを確認
            if dbArticle.UpdatedAt.Equal(article.UpdatedAt) {
                t.Error("UpdateArticle() did not update the UpdatedAt timestamp")
            }
        })
    })

    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないIDを指定した場合はエラーを返す", func(t *testing.T) {
            invalidReq := model.ArticleRequest{
                Title:   "Invalid Update",
                Content: "Invalid Content",
                UserId:  articleTestUser.ID,
            }
            
            invalidArticle := invalidReq.ToModel()
            err := articleRepo.UpdateArticle(&invalidArticle, articleTestUser.ID, nonExistentArticleID)
            
            if err == nil {
                t.Error("UpdateArticle() should return error for non-existent ID")
            }
        })

        t.Run("他のユーザーの記事は更新できない", func(t *testing.T) {
            // 他のユーザーの記事を作成
            otherUserReq := model.ArticleRequest{
                Title:   "Other User's Article",
                Content: "Other Content",
                UserId:  articleOtherUser.ID,
            }
            
            otherUserArticle := otherUserReq.ToModel()
            articleDB.Create(&otherUserArticle)
            
            // 更新を試みる
            updateAttemptReq := model.ArticleRequest{
                Title:   "Attempted Update",
                Content: "Attempted Content",
                UserId:  articleTestUser.ID,
            }
            
            updateAttempt := updateAttemptReq.ToModel()
            err := articleRepo.UpdateArticle(&updateAttempt, articleTestUser.ID, otherUserArticle.ID)
            
            if err == nil {
                t.Error("UpdateArticle() should not allow updating other user's article")
            }
            
            // データベースから記事を取得して変更されていないことを確認
            var dbArticle model.Article
            articleDB.First(&dbArticle, otherUserArticle.ID)
            
            if dbArticle.Title != otherUserReq.Title {
                t.Errorf("UpdateArticle() should not update other user's article, but got title=%s", dbArticle.Title)
            }
        })
    })
}