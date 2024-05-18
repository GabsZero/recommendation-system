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
