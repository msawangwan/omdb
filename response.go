package omdb

// SearchResponse represents a response payload as a result of
// querying the omdb api by search.
type SearchResponse struct {
	Search []struct {
		Title  string `json:"Title,omitempty"`
		Year   string `json:"Year,omitempty"`
		IMDBID string `json:"imdbID,omitempty"`
		Type   string `json:"Type,omitempty"`
		Poster string `json:"Poster,omitempty"`
	} `json:"Search,omitempty"`
}

// QueryResponse represents a response payload as a result of
// querying the omdb api by title or ID.
type QueryResponse struct {
	Title    string `json:"Title,omitempty"`
	Year     string `json:"Year,omitempty"`
	Rated    string `json:"Rated,omitempty"`
	Released string `json:"Released,omitempty"`
	Runtime  string `json:"Runtime,omitempty"`
	Genre    string `json:"Genre,omitempty"`
	Director string `json:"Director,omitempty"`
	Writer   string `json:"Writer,omitempty"`
	Actors   string `json:"Actors,omitempty"`
	Plot     string `json:"Plot,omitempty"`
	Language string `json:"Language,omitempty"`
	Country  string `json:"Country,omitempty"`
	Awards   string `json:"Awards,omitempty"`
	Poster   string `json:"Poster,omitempty"`
	Ratings  []struct {
		Source string `json:"Source,omitempty"`
		Value  string `json:"Value,omitempty"`
	} `json:"Ratings,omitempty"`
	Metascore  string `json:"Metascore,omitempty"`
	IMDBRating string `json:"imdbRating,omitempty"`
	IMDBVotes  string `json:"imdbVotes,omitempty"`
	IMDBID     string `json:"imdbid,omitempty"`
	Type       string `json:"Type,omitempty"`
	DVD        string `json:"DVD,omitempty"`
	BoxOffice  string `json:"BoxOffice,omitempty"`
	Production string `json:"Production,omitempty"`
	Website    string `json:"Website,omitempty"`
}
