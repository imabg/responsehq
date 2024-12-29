package setup

import (
	"github.com/gorilla/mux"
	"github.com/imabg/responehq/internal/services"
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
	compCtx := services.NewCompany(queries)
	userCtx := services.NewUser(queries, compCtx)
	subCtx := services.NewSubscription(queries)
	m.HandleFunc("/sub/create", subCtx.CreateSub).Methods("POST")
	m.HandleFunc("/c/create", compCtx.CreateCompany).Methods("POST")
	m.HandleFunc("/user/create", userCtx.CreateUser).Methods("POST")
	return m
}
