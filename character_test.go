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
			Name:        "GetCharacter",
			Endpoint:    "/character/4005-1253",
			FixtureFile: "character",
			TestFunc: func(client *gocomicvine.Client, t *testing.T) {
				char, err := client.GetCharacterById(1253, make(map[string]string))
				if err != nil {
					t.Fatalf("%s", err.Error())
				}

				require.NotNil(t, char)
				require.Nil(t, char.Birth)
				require.Equal(t, 1253, *char.ID)
				require.Equal(t, "Garth Ranzz", *char.RealName)
				require.NotEmpty(t, char.VolumeCredits)
				require.IsType(t, &models.GenericPower{}, char.Powers[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			fixture := testutil.LoadFixture(t, tt.FixtureFile)
			srv := testutil.NewMockServer(t, fixture, tt.Endpoint)
			defer srv.Close()

			client.BaseURL = srv.URL

			tt.TestFunc(client, t)
		})
	}
}
