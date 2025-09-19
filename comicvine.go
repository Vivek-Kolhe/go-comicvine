package gocomicvine

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Vivek-Kolhe/go-comicvine/models"
)

const BaseURL = "https://comicvine.gamespot.com/api/"

type Client struct {
	ApiKey  string
	BaseURL string
}

// NewClient creates a new instance of ComicVine API client with the given api_key
func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey:  apiKey,
		BaseURL: BaseURL,
	}
}

// GetApiKey returns the api_key of the client
func (c *Client) GetApiKey() string {
	return c.ApiKey
}

// SetApiKey sets the api_key of the client
func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}

// Standalone function for making get requests to API endpoints
func doGetRequest[T any](client *Client, url string, params map[string]string) (*models.ApiResponse[T], error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	queries := req.URL.Query()
	for key, val := range params {
		queries.Add(key, val)
	}
	queries.Add("api_key", client.ApiKey)
	queries.Add("format", "json")

	req.URL.RawQuery = queries.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.ApiResponse[T]
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
