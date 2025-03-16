package article_test

import (
    "go-react-app/model"
    "testing"
)

func TestArticleRepository_DeleteArticle(t *testing.T) {
    setupArticleTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("自分の記事を削除できる", func(t *testing.T) {
            // ArticleRequestを使用して記事を作成
            articleReq := model.ArticleRequest{
                Title:   "Article to Delete",
                Content: "Content to Delete",
                UserId:  articleTestUser.ID,
            }
            
            article := articleReq.ToModel()
            articleDB.Create(&article)
            
            // 削除を実行
            err := articleRepo.DeleteArticle(articleTestUser.ID, article.ID)
            
            if err != nil {
                t.Errorf("DeleteArticle() error = %v", err)
            }
            
            // データベースから削除されていることを確認
            var count int64
            articleDB.Model(&model.Article{}).Where("id = ?", article.ID).Count(&count)
            
            if count != 0 {
                t.Error("DeleteArticle() did not delete the article from database")
            }
        })
    })
    
    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しない記事IDでの削除はエラーになる", func(t *testing.T) {
            err := articleRepo.DeleteArticle(articleTestUser.ID, nonExistentArticleID)
            
            if err == nil {
                t.Error("DeleteArticle() with non-existent ID should return error")
            }
        })
        
        t.Run("他のユーザーの記事を削除しようとするとエラー", func(t *testing.T) {
            // 他のユーザーの記事を作成
            otherUserReq := model.ArticleRequest{
                Title:   "Other User's Article",
                Content: "Other Content",
                UserId:  articleOtherUser.ID,
            }
            
            otherUserArticle := otherUserReq.ToModel()
            articleDB.Create(&otherUserArticle)
            
            // 削除を試みる
            err := articleRepo.DeleteArticle(articleTestUser.ID, otherUserArticle.ID)
            
            if err == nil {
                t.Error("DeleteArticle() should not allow deleting other user's article")
            }
            
            // データベースから削除されていないことを確認
            var count int64
            articleDB.Model(&model.Article{}).Where("id = ?", otherUserArticle.ID).Count(&count)
            
            if count == 0 {
                t.Error("他ユーザーの記事が削除されています")
            }
        })
    })
}