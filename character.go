package gocomicvine

import (
	"fmt"

	"github.com/Vivek-Kolhe/go-comicvine/models"
)

const CHARACTER_RESOURCE = 4005

// GetCharacterById fetches a list of characters from the ComicVine API.
//
// It makes a GET request to the following endpoint:
//
//	GET /characters
//
// API reference: https://comicvine.gamespot.com/api/documentation#toc-0-3
//
// Parameters:
//   - params (map[string]string): Optional query parameters to include in the request.
//
// Returns:
//   - []*models.CharacterBase: A slice of pointers to models.CharacterBase struct containing details.
//   - error: An error if something goes wrong; otherwise nil.
//
// Example Usage:
//
//	client := NewClient("your_api_key")
//	characters, err := client.GetCharacters(make(map[string]string))
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	for _, character := range characters {
//		fmt.Println(*character.Name)
//	}
func (c *Client) GetCharacters(params map[string]string) ([]*models.CharacterBase, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, "characters")

	response := &models.ApiResponseMultipleResult[models.CharacterBase]{}
	err := c.doGetRequest(url, params, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// GetCharacterById fetches detailed information about a character
// from the ComicVine API using its ID.
//
// It makes a GET request to the following endpoint:
//
//	GET /character/4005-{id}
//
// API reference: https://comicvine.gamespot.com/api/documentation#toc-0-2
//
// Parameters:
//   - id (int): Unique ID of the resource.
//   - params (map[string]string): Optional query parameters to include in the request.
//
// Returns:
//   - *models.Character: A pointer to the *models.Character struct containing all the details.
//   - error: An error if something goes wrong; otherwise nil.
//
// Example Usage:
//
//	client := NewClient("your_api_key")
//	character, err := client.GetCharacterById(1253, make(map[string]string))
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	fmt.Println(*character.Name)
func (c *Client) GetCharacterById(id int, params map[string]string) (*models.Character, error) {
	url := fmt.Sprintf("%s/%s/%d-%d", c.BaseURL, "character", CHARACTER_RESOURCE, id)

	response := &models.ApiResponseSingleResult[models.Character]{}
	err := c.doGetRequest(url, params, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}
