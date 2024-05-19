package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gabszero/recommendation-system-go/internal/domains/users"
	"github.com/gorilla/mux"
)

func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/movies/watched", users.WatchedMovie).Methods("Post")

	port := os.Getenv("RECOMMENDATION_SYSTEM_HOST_PORT")
	fmt.Println("Listening request on port " + port)

	http.ListenAndServe(":"+port, r)
}
