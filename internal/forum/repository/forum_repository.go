package repository

import (
	"log"

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
	query := "SELECT * FROM threads WHERE forum = $1"
	queryLimit := " ORDER BY created"

	if since != "" {
		query += " AND created"
		if descBool {
			query += " <= '" + since + "'"
		} else {
			query += " >= '" + since + "'"
		}
	}
	if descBool {
		query += queryLimit + " DESC"
	} else {
		query += queryLimit
	}
	query += " LIMIT $2"

	var threads []models.Thread

	err := f.DB.Select(&threads, query,
		slug, limitInt,
	)

	if len(threads) == 0 {
		threads = make([]models.Thread, 0)
	}

	return threads, err
}

func (f *ForumRepository) ForumGetUsers(slug string, limitInt int, descBool bool, since string, forumSlug string) ([]models.User, error) {
	query := "SELECT DISTINCT u.nickname, u.fullname, u.about, u.email FROM users AS u JOIN forum_participants as f ON f.user_nickname = u.nickname WHERE f.forum = $1"
	queryLimit := " ORDER BY u.nickname"

	if since != "" {
		query += " AND f.user_nickname"
		if descBool {
			query += " < '" + since + "'"
		} else {
			query += " > '" + since + "'"
		}
	}

	if descBool {
		query += queryLimit + " DESC"
	} else {
		query += queryLimit
	}

	query += " LIMIT $2"

	var users []models.User

	log.Println(query)

	err := f.DB.Select(&users, query,
		slug, limitInt,
	)

	if len(users) == 0 {
		users = make([]models.User, 0)
	}

	return users, err
}
