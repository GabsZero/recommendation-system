package models

import "time"

type Movie struct {
	Adult                 bool
	Belongs_to_collection string
	Budget                float64
	Genres                string
	Homepage              string
	Id                    uint64
	Imdb_id               uint64
	Original_language     string
	Original_title        string
	Overview              string
	Popularity            float64
	Poster_path           string
	Production_companies  string
	Production_countries  string
	Release_date          time.Time
	Revenue               float64
	Runtime               float64
	Spoken_languages      string
	Status                string
	Tagline               string
	Title                 string
	Video                 bool
	Vote_average          float64
	Vote_count            uint64
}
