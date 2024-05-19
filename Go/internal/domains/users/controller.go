package users

import (
	"net/http"
	"strconv"

	"github.com/gabszero/recommendation-system-go/internal/infra"
)

func WatchedMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userId, err := strconv.ParseUint(r.Form["userId"][0], 10, 32)
	if err != nil {
		panic(err)
	}

	movieId, err := strconv.ParseUint(r.Form["movieId"][0], 10, 32)
	if err != nil {
		panic(err)
	}

	infra.SaveMovieWatched(uint(userId), uint(movieId))
}
