package repository

import (
	"fmt"
	"go-react-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IFeedRepository interface {
	GetAllFeeds(feeds *[]model.Feed, userId uint) error
	GetFeedById(feed *model.Feed, userId uint, feedId uint) error
	CreateFeed(feed *model.Feed) error
	UpdateFeed(feed *model.Feed, userId uint, feedId uint) error
	DeleteFeed(userId uint, feedId uint) error
}

type feedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) IFeedRepository {
	return &feedRepository{db}
}

func (fr *feedRepository) GetAllFeeds(feeds *[]model.Feed, userId uint) error {
	if err := fr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(feeds).Error; err != nil {
		return err
	}
	return nil
}

func (fr *feedRepository) GetFeedById(feed *model.Feed, userId uint, feedId uint) error {
	if err := fr.db.Joins("User").Where("user_id=?", userId).First(feed, feedId).Error; err != nil {
		return err
	}
	return nil
}

func (fr *feedRepository) CreateFeed(feed *model.Feed) error {
	if err := fr.db.Create(feed).Error; err != nil {
		return err
	}
	return nil
}

func (fr *feedRepository) UpdateFeed(feed *model.Feed, userId uint, feedId uint) error {
	result := fr.db.Model(feed).Clauses(clause.Returning{}).Where("id=? AND user_id=?", feedId, userId).Updates(map[string]interface{}{
		"title":       feed.Title,
		"url":         feed.URL,
		"site_url":         feed.SiteURL,
		"description": feed.Description,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (fr *feedRepository) DeleteFeed(userId uint, feedId uint) error {
	result := fr.db.Where("id=? AND user_id=?", feedId, userId).Delete(&model.Feed{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
