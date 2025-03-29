package repository

import (
	"encoding/json"
	"fmt"
	"go-react-app/model"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type IGoogleBookRepository interface {
	SearchBooks(query string, maxResults int) (model.GoogleBookSearchResponse, error)
	GetBookByID(id string) (model.GoogleBook, error)
}

type googleBookRepository struct{}

func NewGoogleBookRepository() IGoogleBookRepository {
	return &googleBookRepository{}
}

func (gbr *googleBookRepository) SearchBooks(query string, maxResults int) (model.GoogleBookSearchResponse, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		return model.GoogleBookSearchResponse{}, fmt.Errorf("GOOGLE_BOOKS_API_KEY is not set")
	}

	if maxResults <= 0 {
		maxResults = 10 // デフォルト値
	}

	baseURL := "https://www.googleapis.com/books/v1/volumes"
	params := url.Values{}
	params.Add("q", query)
	params.Add("maxResults", fmt.Sprintf("%d", maxResults))
	params.Add("key", apiKey)

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return model.GoogleBookSearchResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.GoogleBookSearchResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var apiResp struct {
		TotalItems int `json:"totalItems"`
		Items      []struct {
			ID         string `json:"id"`
			VolumeInfo struct {
				Title         string   `json:"title"`
				Authors       []string `json:"authors"`
				Description   string   `json:"description"`
				PublishedDate string   `json:"publishedDate"`
				ImageLinks    struct {
					Thumbnail string `json:"thumbnail"`
				} `json:"imageLinks"`
				IndustryIdentifiers []struct {
					Type       string `json:"type"`
					Identifier string `json:"identifier"`
				} `json:"industryIdentifiers"`
			} `json:"volumeInfo"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return model.GoogleBookSearchResponse{}, err
	}

	result := model.GoogleBookSearchResponse{
		TotalItems: apiResp.TotalItems,
		Items:      make([]model.GoogleBook, 0, len(apiResp.Items)),
	}

	for _, item := range apiResp.Items {
		book := model.GoogleBook{
			ID:           item.ID,
			Title:        item.VolumeInfo.Title,
			Authors:      item.VolumeInfo.Authors,
			Description:  item.VolumeInfo.Description,
			ImageURL:     item.VolumeInfo.ImageLinks.Thumbnail,
			PublishedDate: item.VolumeInfo.PublishedDate,
		}

		// ISBNを取得
		for _, identifier := range item.VolumeInfo.IndustryIdentifiers {
			if strings.Contains(identifier.Type, "ISBN") {
				book.ISBN = identifier.Identifier
				break
			}
		}

		result.Items = append(result.Items, book)
	}

	return result, nil
}

func (gbr *googleBookRepository) GetBookByID(id string) (model.GoogleBook, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		return model.GoogleBook{}, fmt.Errorf("GOOGLE_BOOKS_API_KEY is not set")
	}

	baseURL := "https://www.googleapis.com/books/v1/volumes"
	url := fmt.Sprintf("%s/%s?key=%s", baseURL, id, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return model.GoogleBook{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.GoogleBook{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var apiResp struct {
		ID         string `json:"id"`
		VolumeInfo struct {
			Title         string   `json:"title"`
			Authors       []string `json:"authors"`
			Description   string   `json:"description"`
			PublishedDate string   `json:"publishedDate"`
			ImageLinks    struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
		} `json:"volumeInfo"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return model.GoogleBook{}, err
	}

	book := model.GoogleBook{
		ID:           apiResp.ID,
		Title:        apiResp.VolumeInfo.Title,
		Authors:      apiResp.VolumeInfo.Authors,
		Description:  apiResp.VolumeInfo.Description,
		ImageURL:     apiResp.VolumeInfo.ImageLinks.Thumbnail,
		PublishedDate: apiResp.VolumeInfo.PublishedDate,
	}

	// ISBNを取得
	for _, identifier := range apiResp.VolumeInfo.IndustryIdentifiers {
		if strings.Contains(identifier.Type, "ISBN") {
			book.ISBN = identifier.Identifier
			break
		}
	}

	return book, nil
}
