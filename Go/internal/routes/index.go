package routes

import (
	"net/http"

	"github.com/gabszero/recommendation-system-go/internal/domains/users"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/movies/watched", users.WatchedMovie).Methods("Post")

	http.ListenAndServe(":8000", r)
}
