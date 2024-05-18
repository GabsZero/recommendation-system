package importmovies

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gabszero/recommendation-system-go/internal/infra"
	"github.com/gabszero/recommendation-system-go/internal/infra/models"
)

func ImportMovie() {
	moviesCount := infra.CountMovies()

	fmt.Println(moviesCount)
	if moviesCount > 0 {
		return
	}
	file, err := os.Open("movies_metadata.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for ok := true; ok; ok = ok != false {

		row, err := reader.Read()
		movie := models.Movie{}
		movie.Adult, _ = strconv.ParseBool(row[0])
		movie.Belongs_to_collection = row[1]
		movie.Budget, _ = strconv.ParseFloat(row[2], 64)
		movie.Genres = row[3]
		movie.Homepage = row[4]
		movie.Id, _ = strconv.ParseUint(row[5], 10, 64)
		movie.Imdb_id, _ = strconv.ParseUint(row[6], 10, 64)
		movie.Original_language = row[7]
		movie.Original_title = row[8]
		movie.Overview = row[9]
		movie.Popularity, _ = strconv.ParseFloat(row[10], 64)
		movie.Poster_path = row[11]
		movie.Production_companies = row[12]
		movie.Production_countries = row[13]
		movie.Release_date, _ = time.Parse("2006-01-02", row[14])
		movie.Revenue, _ = strconv.ParseFloat(row[14], 64)
		movie.Runtime, _ = strconv.ParseFloat(row[15], 64)
		movie.Spoken_languages = row[16]
		movie.Status = row[17]
		movie.Tagline = row[18]
		movie.Title = row[19]
		movie.Video, _ = strconv.ParseBool(row[20])
		movie.Vote_average, _ = strconv.ParseFloat(row[21], 64)
		movie.Vote_count, _ = strconv.ParseUint(row[22], 10, 64)

		if err != nil {
			println(err)
			ok = false
		}

		infra.SaveMovie(&movie)
	}

}
