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
		logger.Error(ctx, "while connecting pgsql", err)
		return nil, err
	}
	//defer closeCon(ctx, conn)
	if err = conn.Ping(ctx); err != nil {
		logger.Error(ctx, "while pinging pgsql", err)
		closeCon(ctx, conn)
	}
	logger.Info(ctx, "database is connected successfully", nil)
	q := models.New(conn)
	return q, nil
}

func closeCon(ctx context.Context, conn *pgx.Conn) {
	err := conn.Close(ctx)
	if err != nil {
		logger.Error(ctx, "while closing conn", err)
		panic(err)
	}
}

func RunMigration(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error(context.Background(), "while connecting postgres", err)
		return err
	}
	defer runMigrationCleanup(db)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error(context.Background(), "while connecting postgres", err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/schemas/", "postgres", driver)
	if err != nil {
		logger.Error(context.Background(), "while connecting postgres", err)
		return err
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			logger.Error(context.Background(), "while running migration", err)
			return err
		}
	}
	logger.Info(context.Background(), "database migration successfully", nil)
	return nil
}

func runMigrationCleanup(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Error(context.Background(), "while closing db: Migrations", err)
	}
}
