package setup

import (
	"github.com/gorilla/mux"
	"github.com/imabg/responehq/internal/services"
	"github.com/imabg/responehq/models"
)

func GetRoutes(queries *models.Queries) *mux.Router {
	m := mux.NewRouter().PathPrefix("/api").Subrouter()
	compCtx := services.NewCompany(queries)
	userCtx := services.NewUser(queries, compCtx)
	subCtx := services.NewSubscription(queries)
	m.HandleFunc("/sub/create", subCtx.CreateSub).Methods("POST")
	m.HandleFunc("/c/create", compCtx.CreateCompany).Methods("POST")
	m.HandleFunc("/user/create", userCtx.CreateUser).Methods("POST")
	m.HandleFunc("/user/login", userCtx.Login).Methods("POST")
	return m
}
