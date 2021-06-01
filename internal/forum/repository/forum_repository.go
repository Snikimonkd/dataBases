package repository

import "github.com/jmoiron/sqlx"

type ForumRepository struct {
	DB *sqlx.DB
}
