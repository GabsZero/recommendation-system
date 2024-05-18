package models

import "time"

type Movie struct {
	Adult                 bool
	Belongs_to_collection string
	Budget                float32
	Genres                string
	Homepage              string
	Id                    uint
	Imdb_id               uint
	Original_language     string
	Original_title        string
	Overview              string
	Popularity            float32
	Poster_path           string
	Production_companies  string
	Production_countries  string
	Release_date          time.Time
	Revenue               float32
	Runtime               float32
	Spoken_languages      string
	Status                string
	Tagline               string
	Title                 string
	Video                 bool
	Vote_average          float32
	Vote_count            uint
}
