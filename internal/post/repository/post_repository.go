package repository

import (
	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB *sqlx.DB
}

func (f *PostRepository) ThreadGetOneId(id int) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, SelectThreadWithIdQuery,
		id,
	)

	return threads, err
}

func (f *PostRepository) ThreadGetOneSlug(slug string) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, SelectThreadWithSlugQuery,
		slug,
	)

	return threads, err
}

func (f *PostRepository) CheckParents(parent int) ([]int, error) {
	var threads []int
	err := f.DB.Select(&threads, SelectParentQuery,
		parent,
	)

	return threads, err
}

func (f *PostRepository) PostsCreate(post models.Post) (int, error) {
	var id int
	err := f.DB.QueryRow(InsertPostQuery,
		post.Author, post.Created, post.Forum, post.Message, post.Parent, post.Thread,
	).Scan(&id)

	return id, err
}

func (f *PostRepository) CheckUsers(nickname string) ([]models.User, error) {
	var users []models.User
	err := f.DB.Select(&users, CheckUserExistQuery,
		nickname,
	)

	return users, err
}
