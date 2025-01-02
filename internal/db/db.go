package db

import (
	"context"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/imabg/responehq/models"
	"github.com/imabg/responehq/pkg/logger"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func SetupDb(connStr string) (*models.Queries, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		logger.DBError(ctx, "while connecting pgsql", err.Error())
		return nil, err
	}
	//defer closeCon(ctx, conn)
	if err = conn.Ping(ctx); err != nil {
		logger.DBError(ctx, "while pinging pgsql", err.Error())
		closeCon(ctx, conn)
	}
	logger.Info(ctx, "database is connected successfully", nil)
	q := models.New(conn)
	return q, nil
}

func closeCon(ctx context.Context, conn *pgx.Conn) {
	err := conn.Close(ctx)
	if err != nil {
		logger.DBError(ctx, "while closing conn", err.Error())
		panic(err)
	}
}

func RunMigration(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.DBError(context.Background(), "while connecting postgres", err.Error())
		return err
	}
	defer runMigrationCleanup(db)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.DBError(context.Background(), "while connecting postgres", err.Error())
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/schemas/", "postgres", driver)
	if err != nil {
		logger.DBError(context.Background(), "while connecting postgres", err.Error())
		return err
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			logger.DBError(context.Background(), "while running migration", err.Error())
			return err
		}
	}
	logger.Info(context.Background(), "database migration successfully", nil)
	return nil
}

func runMigrationCleanup(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.DBError(context.Background(), "while closing db: Migrations", err.Error())
	}
}
