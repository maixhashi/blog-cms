package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IArticleRepository interface {
	GetAllArticles(articles *[]model.Article, userId uint) error
	GetArticleById(article *model.Article, userId uint, articleId uint) error
	CreateArticle(article *model.Article) error
	UpdateArticle(article *model.Article, userId uint, articleId uint) error
	DeleteArticle(userId uint, articleId uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) GetAllArticles(articles *[]model.Article, userId uint) error {
	if err := ar.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Find(articles).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) GetArticleById(article *model.Article, userId uint, articleId uint) error {
	if err := ar.db.Joins("User").Where("user_id=?", userId).First(article, articleId).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) CreateArticle(article *model.Article) error {
	if err := ar.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) UpdateArticle(article *model.Article, userId uint, articleId uint) error {
	result := ar.db.Model(article).Clauses(clause.Returning{}).Where("id=? AND user_id=?", articleId, userId).Updates(map[string]interface{}{
		"title":     article.Title,
		"content":   article.Content,
		"published": article.Published,
		"tags":      article.Tags,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("article does not exist")
	}
	return nil
}

func (ar *articleRepository) DeleteArticle(userId uint, articleId uint) error {
	result := ar.db.Where("id=? AND user_id=?", articleId, userId).Delete(&model.Article{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("article does not exist")
	}
	return nil
}
