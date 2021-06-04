package repository

import (
	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/jmoiron/sqlx"
)

type ForumRepository struct {
	DB *sqlx.DB
}

func (f *ForumRepository) CheckUsers(newForum models.Forum) ([]models.User, error) {
	var users []models.User
	err := f.DB.Select(&users, CheckUserExistQuery,
		newForum.User,
	)

	return users, err
}

func (f *ForumRepository) CheckForums(newForum models.Forum) ([]models.Forum, error) {
	var forums []models.Forum
	err := f.DB.Select(&forums, CheckForumExistQuery,
		newForum.Slug,
	)

	return forums, err
}

func (f *ForumRepository) CreateForum(newForum models.Forum) error {
	_, err := f.DB.Exec(InsertForumQuery,
		newForum.Title, newForum.User, newForum.Slug,
	)

	return err
}

func (f *ForumRepository) ForumGetOne(slug string) ([]models.Forum, error) {
	var forums []models.Forum
	err := f.DB.Select(&forums, SelectForumQuery,
		slug,
	)

	return forums, err
}

func (f *ForumRepository) ForumGetThreads(slug string, limitInt int, descBool bool, since string) ([]models.Thread, error) {
	var threads []models.Thread
	if !descBool {
		err := f.DB.Select(&threads, SelectThreadsQueryDesc,
			slug, limitInt,
		)
		return threads, err
	}

	err := f.DB.Select(&threads, SelectThreadsQuery,
		slug, limitInt,
	)

	return threads, err
}
