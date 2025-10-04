package gocomicvine

import (
	"fmt"

	"github.com/Vivek-Kolhe/go-comicvine/models"
)

const CHARACTER_RESOURCE = 4005

// GetCharacters (Endpoint: /characters) returns a slice of *models.Character,
// takes in all query params as a map[string]string.
// More info on query params: https://comicvine.gamespot.com/api/documentation#toc-0-3
func (c *Client) GetCharacters(params map[string]string) ([]*models.CharacterBase, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, "characters")

	response := &models.ApiResponseMultipleResult[models.CharacterBase]{}
	err := c.doGetRequest(url, params, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

func (c *Client) GetCharacterById(id int, params map[string]string) (*models.Character, error) {
	url := fmt.Sprintf("%s/%s/%d-%d", c.BaseURL, "character", CHARACTER_RESOURCE, id)

	response := &models.ApiResponseSingleResult[models.Character]{}
	err := c.doGetRequest(url, params, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}
