package models

import (
	"context"
	"fmt"

	"github.com/daniel-z-johnson/journal-of-self/config"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Services struct {
	Uservice UserService
	pool     *pgxpool.Pool
}

func NewServices(dbConfig *config.DatabaseSQL) (*Services, error) {
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
	if err != nil {
		return nil, err
	}
	return &Services{
		pool:     pgxPool,
		Uservice: newUserService(pgxPool),
	}, nil

}

func (s *Services) Close() {
	s.pool.Close()
}
