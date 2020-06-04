package main

import (
	"fmt"
	"os"

	"flag"
	"strings"

	"encoding/json"

	"github.com/msawangwan/omdb"
)

var (
	flagCmd        string
	flagAPIKey     string
	flagOutputFile bool

	flagByID           string
	flagByTitle        string
	flagSearchKeywords string
)

func display(o interface{}) {
	raw, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(raw))

	if flagOutputFile {
		titleNoWS := strings.Replace(flagByTitle, " ", "", -1)
		searchNoWS := strings.Replace(flagSearchKeywords, " ", "", -1)

		f, err := os.Create(strings.TrimSpace(flagByID) + strings.TrimSpace(titleNoWS) + strings.TrimSpace(searchNoWS) + ".json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, err := f.Write(raw); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func init() {
	flag.StringVar(&flagCmd, "cmd", "", "query or search")
	flag.StringVar(&flagAPIKey, "key", "", "specify an OMDb API key")
	flag.StringVar(&flagByID, "id", "", "query by IMDB ID")
	flag.StringVar(&flagByTitle, "title", "", "query by movie title")
	flag.BoolVar(&flagOutputFile, "out", false, "write output to a file")
	flag.StringVar(&flagSearchKeywords, "keywords", "", "search keywords")

	flag.Parse()
}

func main() {
	secret := flagAPIKey

	if len(strings.TrimSpace(secret)) == 0 {
		secret = strings.TrimSpace(os.Getenv("OMDB_API_KEY"))
	}

	if len(strings.TrimSpace(secret)) == 0 {
		fmt.Println("please specify an API key (either export 'OMDB_API_KEY' or specify '--key' as command-line argument)")
		os.Exit(1)
	}

	cfg := strings.NewReader(fmt.Sprintf(`{
		"api": {
			"key": "%s",
			"endpoint": {
				"data": "http://www.omdbapi.com/",
				"image": "http://img.omdbapi.com/"
			}
		}
	}`, secret))

	client, err := omdb.New(cfg, 30)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("===")
	fmt.Println("submitting", flagCmd)
	fmt.Println("  title:", flagByTitle)
	fmt.Println("  id:", flagByID)
	fmt.Println("  keywords:", flagSearchKeywords)
	fmt.Println()

	switch flagCmd {
	case "query":
		res, err := client.Query(omdb.QueryRequest{ID: flagByID, Title: flagByTitle})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		display(res)
	case "search":
		res, err := client.Search(omdb.SearchRequest{Search: flagSearchKeywords})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		display(res)
	default:
		fmt.Println("unrecogized command:", flagCmd)
		os.Exit(1)
	}

	os.Exit(0)
}
