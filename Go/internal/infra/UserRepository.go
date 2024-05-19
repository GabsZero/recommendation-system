package infra

import "github.com/gabszero/recommendation-system-go/internal/infra/models"

func SaveUser(user *models.User) {
	db.Create(user)
}

func CountUsers() int64 {
	var count int64
	db.Model(models.User{}).Count(&count)

	return count
}

func SaveMovieWatched(userId uint, movieId uint) {
	user := FindUser(userId)
	movie := FindMovie(movieId)

	db.Model(&user).Association("Movies").Append(&movie)

}

func FindUser(userId uint) models.User {
	user := models.User{
		ID: userId,
	}
	db.Find(&user)

	return user
}
