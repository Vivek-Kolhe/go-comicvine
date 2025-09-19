package gocomicvine

import (
	"fmt"

	"github.com/Vivek-Kolhe/go-comicvine/models"
)

// GetCharacters (Endpoint: /characters) returns a slice of *models.Character,
// takes in all query params as a map[string]string.
// More info on query params: https://comicvine.gamespot.com/api/documentation#toc-0-3
func (c *Client) GetCharacters(params map[string]string) ([]*models.Character, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, "characters")

	response, err := doGetRequest[models.Character](c, url, params)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}
