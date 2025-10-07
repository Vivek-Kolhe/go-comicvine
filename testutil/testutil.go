package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	gocomicvine "github.com/Vivek-Kolhe/go-comicvine"
)

type Table struct {
	Name        string
	FixtureFile string
	Endpoint    string
	TestFunc    func(client *gocomicvine.Client, t *testing.T)
}

func (tb *Table) Test(t *testing.T, client *gocomicvine.Client) {
	fixture := LoadFixture(t, tb.FixtureFile)
	srv := NewMockServer(t, fixture, tb.Endpoint)
	defer srv.Close()

	client.BaseURL = srv.URL
	tb.TestFunc(client, t)
}

func LoadFixture(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", filename))
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	return data
}

func NewMockServer(t *testing.T, fixture []byte, expectedPath string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || expectedPath != r.URL.Path {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(fixture)
	}))
}
