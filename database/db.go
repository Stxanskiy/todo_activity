package database

import "github.com/jackc/pgx/v5/pgxpool"

type pg struct {
	db *pgxpool.Pool
}

func NewPg(db *pgxpool.Pool) i.Repo {
	return &pg{
		db: db,
	}
}
