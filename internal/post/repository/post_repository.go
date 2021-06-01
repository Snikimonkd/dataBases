package repository

import "github.com/jmoiron/sqlx"

type PostRepository struct {
	DB *sqlx.DB
}
