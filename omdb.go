package omdb

import (
	"io"
	"strings"

	"fmt"
	"time"

	"io/ioutil"
	"net/http"
	"net/url"

	"encoding/json"
)

// QueryRequest encapsulates a query request's query parameters.
type QueryRequest struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Year  string `json:"year,omitempty"`
	Plot  string `json:"plot,omitempty"`
}

// SearchRequest encapsulates a search request's query parameters.
type SearchRequest struct {
	Search string `json:"search,omitempty"`
	Type   string `json:"type,omitempty"`
	Year   string `json:"year,omitempty"`
	Page   string `json:"page,omitempty"`
}

// APIClientConfig exposes fields that are mapped to a JSON configuration file.
type APIClientConfig struct {
	API struct {
		Key      string `json:"key,omitempty"`
		Endpoint struct {
			Data  string `json:"data,omitempty"`
			Image string `json:"image,omitempty"`
		} `json:"endpoint,omitempty"`
	} `json:"api,omitempty"`
}

// APIClient wraps the standard library http.Client and maintains any state required
// for interacting with the OMDb API.
type APIClient struct {
	http.Client `json:"-"`

	DataEndpoint  string `json:"data_endpoint,omitempty"`
	ImageEndpoint string `json:"image_endpoint,omitempty"`

	*APIClientConfig `json:"conf,omitempty"`
}

// New creates a new client, initialized with the standard library http.Client.
func New(config io.Reader, timeoutSeconds int) (*APIClient, error) {
	raw, err := ioutil.ReadAll(config)
	if err != nil {
		return nil, err
	}

	var (
		ac *APIClientConfig = &APIClientConfig{}
	)

	if err := json.Unmarshal(raw, ac); err != nil {
		return nil, err
	}

	dataendpoint := strings.Trim(ac.API.Endpoint.Data, "/")
	imgendpoint := strings.Trim(ac.API.Endpoint.Image, "/")

	return &APIClient{
		APIClientConfig: ac,
		DataEndpoint:    fmt.Sprintf("%s/?apikey=%s", dataendpoint, ac.API.Key),
		ImageEndpoint:   fmt.Sprintf("%s/?apikey=%s", imgendpoint, ac.API.Key),
		Client: http.Client{
			Timeout: time.Second * time.Duration(timeoutSeconds),
		},
	}, nil
}

// Query queries an external API for a movie by title.
func (api *APIClient) Query(q QueryRequest) (*QueryResponse, error) {
	var (
		o *QueryResponse = &QueryResponse{}
	)

	resource := ""

	if q.Title != "" {
		resource = fmt.Sprintf("%s&t=%s", api.DataEndpoint, url.QueryEscape(q.Title))
	} else if q.ID != "" {
		resource = fmt.Sprintf("%s&i=%s", api.DataEndpoint, q.ID)
	} else {
		return nil, fmt.Errorf("missing required query parameter: need title or id")
	}

	if q.Year != "" {
		resource = resource + "&y=" + q.Year
	}

	if q.Plot != "" {
		resource = resource + "&plot=full"
	}

	if q.Type == "movie" || q.Type == "series" || q.Type == "episode" {
		resource = resource + "&type=" + q.Type
	}

	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return nil, err
	}

	res, err := api.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	if err := json.Unmarshal(data, o); err != nil {
		return nil, err
	}

	return o, nil
}

// Search submits a search for movies match a search string.
func (api *APIClient) Search(q SearchRequest) (*SearchResponse, error) {
	var (
		o *SearchResponse = &SearchResponse{}
	)

	resource := fmt.Sprintf("%s&s=%s", api.DataEndpoint, url.QueryEscape(q.Search))

	if q.Page != "" {
		resource = resource + "&page=full"
	}

	if q.Year != "" {
		resource = resource + "&y=" + q.Year
	}

	if q.Type == "movie" || q.Type == "series" || q.Type == "episode" {
		resource = resource + "&type=" + q.Type
	}

	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return nil, err
	}

	res, err := api.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	if err := json.Unmarshal(data, o); err != nil {
		return nil, err
	}

	return o, nil
}
