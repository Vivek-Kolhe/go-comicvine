package gocomicvine_test

import (
	"testing"

	gocomicvine "github.com/Vivek-Kolhe/go-comicvine"
	"github.com/Vivek-Kolhe/go-comicvine/models"
	"github.com/Vivek-Kolhe/go-comicvine/testutil"
	"github.com/stretchr/testify/require"
)

func TestCharacterEndpoints(t *testing.T) {
	client := gocomicvine.NewClient("dummy_api_key")

	tests := []testutil.Table{
		{
			Name:        "GetCharacterById",
			Endpoint:    "/character/4005-1253",
			FixtureFile: "character",
			TestFunc: func(client *gocomicvine.Client, t *testing.T) {
				char, err := client.GetCharacterById(1253, make(map[string]string))

				require.NoError(t, err)
				require.NotNil(t, char)
				require.Nil(t, char.Birth)
				require.Equal(t, 1253, *char.ID)
				require.Equal(t, "Garth Ranzz", *char.RealName)
				require.NotEmpty(t, char.VolumeCredits)
				require.IsType(t, &models.GenericPower{}, char.Powers[0])
			},
		},
		{
			Name:        "GetCharacters",
			Endpoint:    "/characters",
			FixtureFile: "characters",
			TestFunc: func(client *gocomicvine.Client, t *testing.T) {
				chars, err := client.GetCharacters(make(map[string]string))

				require.NoError(t, err)
				require.NotEmpty(t, chars)
				if len(chars) > 0 {
					require.IsType(t, &models.CharacterBase{}, chars[0])
					require.NotNil(t, chars[0].ID)
					require.NotNil(t, chars[0].Name)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.Test(t, client)
		})
	}
}
