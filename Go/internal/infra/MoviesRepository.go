package infra

import "github.com/gabszero/recommendation-system-go/internal/infra/models"

func SaveMovie(movie *models.Movie) {
	db.Create(movie)
}

func CountMovies() int64 {
	var count int64
	db.Model(models.Movie{}).Count(&count)

	return count

}
