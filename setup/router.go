package setup

import (
	"github.com/imabg/responehq/models"
	"log"
	"net/http"
	"time"
)

func Router(queries *models.Queries) {
	srv := http.Server{
		Addr:         ":8080",
		Handler:      GetRoutes(queries),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
