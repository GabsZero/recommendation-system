package imports

import (
	"github.com/gabszero/recommendation-system-go/internal/infra"
	"github.com/gabszero/recommendation-system-go/internal/infra/models"
)

func ImportUsers() {
	usersCount := infra.CountUsers()
	if usersCount > 0 {
		return
	}

	usersToImport := []models.User{
		{
			Name:  "Test 1",
			Email: "test1@test.com",
		},
		{
			Name:  "Test 2",
			Email: "test2@test.com",
		},
		{
			Name:  "Test 3",
			Email: "test3@test.com",
		},
	}

	for _, user := range usersToImport {
		infra.SaveUser(&user)
	}
}
