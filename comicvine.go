package gocomicvine

import (
	"encoding/json"
	"io"
	"net/http"
)

// Default ComicVine API url.
const BaseURL = "https://comicvine.gamespot.com/api"

// Client represents a ComicVine API client.
//
// It holds the API key and the base URL used to make requests
// to the ComicVine API.
type Client struct {
	// ApiKey is your personal ComicVine API key used for authentication.
	ApiKey string

	// BaseURL is the default base url for ComicVine API endpoints.
	BaseURL string
}

// NewClient creates and returns a new ComicVine API client.
//
// It takes an API key as param and initializes a Client struct
// with default API url and provided API key.
//
// Parameters:
// 	- apiKey: ComicVine API key (Can be grabbed from https://comicvine.gamespot.com/api/)
//
// Returns:
// 	- A pointer to a Client struct, which can be used to make
//	  make requests to the ComicVine API.

// Example Usage:
//
//	client := NewClient("your_api_key")
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
func (c *Client) doGetRequest(url string, params map[string]string, target any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	queries := req.URL.Query()
	for key, val := range params {
		queries.Add(key, val)
	}
	queries.Add("api_key", c.ApiKey)
	queries.Add("format", "json")

	req.URL.RawQuery = queries.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
