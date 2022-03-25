package omdb_test

import (
	"fmt"
	"os"
	"testing"

	"encoding/json"
	"strings"

	"github.com/msawangwan/omdb"
)

func pretty(t *testing.T, o interface{}) {
	raw, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(raw))
}

func TestCreateClient(t *testing.T) {
	confdata := strings.NewReader(`{
		"api": {
			"key": "SECRET_KEY",
			"endpoint": {
				"data": "http://www.omdbapi.com/",
				"image": "http://img.omdbapi.com/"
			}
		}
	}`)

	client, err := omdb.New(confdata, 1)
	if err != nil {
		t.Fatal(err)
	}

	pretty(t, client)
}

func TestQuery(t *testing.T) {
	var (
		secret string
	)

	if v, isset := os.LookupEnv("OMDB_API_KEY"); isset == false {
		t.Skip("no valid api key found, set one with 'OMDB_API_KEY'")
	} else {
		secret = v
	}

	confdata := strings.NewReader(fmt.Sprintf(`{
		"api": {
			"key": "%s",
			"endpoint": {
				"data": "http://www.omdbapi.com/",
				"image": "http://img.omdbapi.com/"
			}
		}
	}`, secret))

	client, err := omdb.New(confdata, 1)
	if err != nil {
		t.Fatal(err)
	}

	var testcases = []struct {
		label string
		q     omdb.QueryRequest
	}{
		{"title without space", omdb.QueryRequest{Title: "shrek"}},
		{"title with space", omdb.QueryRequest{Title: "star wars"}},
		{"id", omdb.QueryRequest{ID: "tt0395789"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			res, err := client.Query(testcase.q)
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			pretty(t, res)
		})
	}
}

func TestSearch(t *testing.T) {
	var (
		secret string
	)

	if v, isset := os.LookupEnv("OMDB_API_KEY"); isset == false {
		t.Skip("no valid api key found, set one with 'OMDB_API_KEY'")
	} else {
		secret = v
	}

	confdata := strings.NewReader(fmt.Sprintf(`{
		"api": {
			"key": "%s",
			"endpoint": {
				"data": "http://www.omdbapi.com/",
				"image": "http://img.omdbapi.com/"
			}
		}
	}`, secret))

	client, err := omdb.New(confdata, 1)
	if err != nil {
		t.Fatal(err)
	}

	var testcases = []struct {
		label string
		q     omdb.SearchRequest
	}{
		{"title without space", omdb.SearchRequest{Search: "shrek"}},
		{"title with space", omdb.SearchRequest{Search: "toy story"}},
		{"title with special character", omdb.SearchRequest{Search: "star wars - a new hope"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			res, err := client.Search(testcase.q)
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			pretty(t, res)
		})
	}
}
