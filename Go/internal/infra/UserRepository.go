package infra

import "github.com/gabszero/recommendation-system-go/internal/infra/models"

func SaveUser(user *models.User) {
	db.Create(user)
}
