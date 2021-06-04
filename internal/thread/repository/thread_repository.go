package repository

import (
	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/jmoiron/sqlx"
)

type ThreadRepository struct {
	DB *sqlx.DB
}

func (f *ThreadRepository) CheckForum(newThread models.Thread) ([]models.Forum, error) {
	var forums []models.Forum
	err := f.DB.Select(&forums, CheckForumExistQuery,
		newThread.Forum,
	)

	return forums, err
}

func (f *ThreadRepository) CheckUser(newThread models.Thread) ([]models.User, error) {
	var users []models.User
	err := f.DB.Select(&users, CheckUserExistQuery,
		newThread.Author,
	)

	return users, err
}

func (f *ThreadRepository) CheckThread(newThread models.Thread) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, CheckThreadExistQuery,
		newThread.Slug,
	)

	return threads, err
}

func (f *ThreadRepository) CreateThread(newThread models.Thread) (int, error) {
	var id int
	err := f.DB.QueryRow(InsertThreadQuery,
		newThread.Title, newThread.Author, newThread.Forum, newThread.Message, newThread.Created, newThread.Slug,
	).Scan(&id)

	return id, err
}

func (f *ThreadRepository) ThreadGetOneId(id int) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, SelectThreadWithIdQuery,
		id,
	)

	return threads, err
}

func (f *ThreadRepository) ThreadGetOneSlug(slug string) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, SelectThreadWithSlugQuery,
		slug,
	)

	return threads, err
}
