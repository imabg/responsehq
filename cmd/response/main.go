package main

import (
	"github.com/imabg/responehq/config"
	"github.com/imabg/responehq/internal/db"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/imabg/responehq/setup"
)

func main() {
	logger.New()
	c := config.NewConfig()
	q, err := db.SetupDb(c.PostgresURL)
	if err != nil {
		panic(err)
	}
	err = db.RunMigration(c.PostgresURL)
	if err != nil {
		panic(err)
	}
	setup.Router(q)
	select {}
}
