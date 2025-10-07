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
//   - apiKey (string): ComicVine API key (Can be grabbed from https://comicvine.gamespot.com/api/).
//
// Returns:
//   - *Client: A pointer to a Client struct, which can be used to make
//     make requests to the ComicVine API.
func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey:  apiKey,
		BaseURL: BaseURL,
	}
}

// GetApiKey returns the api_key associated with the client
//
// Returns:
//   - string: The API key of the client.
func (c *Client) GetApiKey() string {
	return c.ApiKey
}

// SetApiKey sets the api_key for the client
//
// Parameters:
//   - apiKey (string): ComicVine API key (Can be grabbed from https://comicvine.gamespot.com/api/).
func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}

// doGetRequest performs an HTTP GET request to the specified ComicVine API endpoint.
//
// It constructs and sends a GET request using the provided URL and query parameters,
// automatically adding required parameters such as api_key and response format.
// The JSON response is then decoded into the specified type.
//
// Parameters:
//   - url (string): Complete URL of the API endpoint.
//   - params (map[string]string): Optional query paramters to include in the request.
//   - target (any): A pointer to the variable where the JSON response will be unmarshalled into.
//
// Returns:
//   - error: An error if something goes wrong; otherwise nil.
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
