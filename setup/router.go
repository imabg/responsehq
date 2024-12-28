package setup

import (
	"github.com/gorilla/mux"
	"github.com/imabg/responehq/internal/services/user"
	"github.com/imabg/responehq/models"
	"log"
	"net/http"
	"time"
)

func Router(queries *models.Queries) {
	srv := http.Server{
		Addr:         ":8080",
		Handler:      getRoutes(queries),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func getRoutes(queries *models.Queries) *mux.Router {
	m := mux.NewRouter().PathPrefix("/api").Subrouter()
	userCtx := user.NewUser(queries)
	m.HandleFunc("/user/create", userCtx.Create).Methods("POST")
	return m
}
