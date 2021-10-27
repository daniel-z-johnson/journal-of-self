package models

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/daniel-z-johnson/journal-of-self/config"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migration
var migrations embed.FS

type Services struct {
	UserService UserService
	pool        *pgxpool.Pool
}

func NewServices(dbConfig *config.DatabaseSQL, migrate bool) (*Services, error) {
	connConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username,
		dbConfig.Password, dbConfig.Database)
	conn, err := pgxpool.ParseConfig(connConfig)
	if err != nil {
		return nil, err
	}
	conn.MaxConns = 20
	conn.BeforeAcquire = func(context context.Context, conn *pgx.Conn) bool {
		fmt.Println("Making a connection")
		return true
	}
	pgxPool, err := pgxpool.ConnectConfig(context.Background(), conn)
	if migrate {
		fmt.Println("Doing migration")
		if err := migration(connConfig); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return &Services{
		pool:        pgxPool,
		UserService: newUserService(pgxPool),
	}, nil

}

func migration(postgresConfig string) error {
	goose.SetBaseFS(migrations)

	db, err := sql.Open("pgx", postgresConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.Up(db, "migration"); err != nil {
		return err
	}
	return nil
}

func (s *Services) Close() {
	s.pool.Close()
}
