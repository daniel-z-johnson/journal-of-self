package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Services struct {
	Uservice UserService
}

func NewServices(connConfig string) (*Services, error) {
	pgxPool, err := pgxpool.Connect(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}
	return &Services{
		Uservice: newUserService(pgxPool),
	}, nil
}
