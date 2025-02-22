package usecase

import (
	"go-react-app/model"
	"go-react-app/repository"
	"go-react-app/validator"
)

type IFeedUsecase interface {
	GetAllFeeds(userId uint) ([]model.FeedResponse, error)
	GetFeedById(userId uint, feedId uint) (model.FeedResponse, error)
	CreateFeed(feed model.Feed) (model.FeedResponse, error)
	UpdateFeed(feed model.Feed, userId uint, feedId uint) (model.FeedResponse, error)
	DeleteFeed(userId uint, feedId uint) error
}

type feedUsecase struct {
	fr repository.IFeedRepository
	fv validator.IFeedValidator
}

func NewFeedUsecase(fr repository.IFeedRepository, fv validator.IFeedValidator) IFeedUsecase {
	return &feedUsecase{fr, fv}
}

func (fu *feedUsecase) GetAllFeeds(userId uint) ([]model.FeedResponse, error) {
	feeds := []model.Feed{}
	if err := fu.fr.GetAllFeeds(&feeds, userId); err != nil {
		return nil, err
	}
	resFeeds := []model.FeedResponse{}
	for _, v := range feeds {
		f := model.FeedResponse{
			ID:        v.ID,
			Title:     v.Title,
			URL:       v.URL,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resFeeds = append(resFeeds, f)
	}
	return resFeeds, nil
}

func (fu *feedUsecase) GetFeedById(userId uint, feedId uint) (model.FeedResponse, error) {
	feed := model.Feed{}
	if err := fu.fr.GetFeedById(&feed, userId, feedId); err != nil {
		return model.FeedResponse{}, err
	}
	resFeed := model.FeedResponse{
		ID:        feed.ID,
		Title:     feed.Title,
		URL:       feed.URL,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
	return resFeed, nil
}

func (fu *feedUsecase) CreateFeed(feed model.Feed) (model.FeedResponse, error) {
	if err := fu.fv.FeedValidate(feed); err != nil {
		return model.FeedResponse{}, err
	}
	if err := fu.fr.CreateFeed(&feed); err != nil {
		return model.FeedResponse{}, err
	}
	resFeed := model.FeedResponse{
		ID:        feed.ID,
		Title:     feed.Title,
		URL:       feed.URL,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
	return resFeed, nil
}

func (fu *feedUsecase) UpdateFeed(feed model.Feed, userId uint, feedId uint) (model.FeedResponse, error) {
	if err := fu.fv.FeedValidate(feed); err != nil {
		return model.FeedResponse{}, err
	}
	if err := fu.fr.UpdateFeed(&feed, userId, feedId); err != nil {
		return model.FeedResponse{}, err
	}
	resFeed := model.FeedResponse{
		ID:        feed.ID,
		Title:     feed.Title,
		URL:       feed.URL,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
	return resFeed, nil
}

func (fu *feedUsecase) DeleteFeed(userId uint, feedId uint) error {
	if err := fu.fr.DeleteFeed(userId, feedId); err != nil {
		return err
	}
	return nil
}
