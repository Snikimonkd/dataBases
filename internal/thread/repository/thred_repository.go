package repository

import "github.com/jmoiron/sqlx"

type ThreadRepository struct {
	DB *sqlx.DB
}
