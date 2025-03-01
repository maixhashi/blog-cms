package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go-react-app/model"
)

type IQiitaRepository interface {
	GetQiitaArticles() ([]model.QiitaArticle, error)
	GetQiitaArticleByID(id string) (model.QiitaArticle, error)
}

type qiitaRepository struct {
	baseURL string
	token   string
}

func NewQiitaRepository() IQiitaRepository {
	return &qiitaRepository{
		baseURL: "https://qiita.com/api/v2",
		token:   os.Getenv("QIITA_ACCESS_TOKEN"),
	}
}

func (qr *qiitaRepository) GetQiitaArticles() ([]model.QiitaArticle, error) {
	url := fmt.Sprintf("%s/items", qr.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+qr.token)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var articles []model.QiitaArticle
	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func (qr *qiitaRepository) GetQiitaArticleByID(id string) (model.QiitaArticle, error) {
	url := fmt.Sprintf("%s/items/%s", qr.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.QiitaArticle{}, err
	}

	req.Header.Set("Authorization", "Bearer "+qr.token)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.QiitaArticle{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.QiitaArticle{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var article model.QiitaArticle
	if err := json.NewDecoder(resp.Body).Decode(&article); err != nil {
		return model.QiitaArticle{}, err
	}

	return article, nil
}
