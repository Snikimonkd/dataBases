package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/go-openapi/strfmt"
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

func (f *PostRepository) PostsCreate(posts []models.Post, thread models.Thread) ([]models.Post, error) {
	query := "INSERT INTO posts (author, created, forum, message, parent, thread) VALUES"
	res := " RETURNING id"
	var args []interface{}

	created := strfmt.DateTime(time.Now())

	for i, post := range posts {
		query += fmt.Sprintf(
			" ($%d, $%d, $%d, $%d, $%d, $%d),",
			i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6,
		)

		args = append(args, post.Author, created, thread.Forum, post.Message, post.Parent, thread.Id)
	}

	query = query[:len(query)-1]
	query += res

	rows, err := f.DB.Queryx(query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		err := rows.Scan(&posts[i].Id)
		if err != nil {
			return nil, err
		}

		posts[i].Forum = thread.Forum
		posts[i].Thread = thread.Id
		posts[i].Created = created

		i++
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (f *PostRepository) CheckUsers(nickname string) ([]models.User, error) {
	var users []models.User
	err := f.DB.Select(&users, CheckUserExistQuery,
		nickname,
	)

	return users, err
}

func (f *PostRepository) PostGetOne(id int) ([]models.Post, error) {
	var posts []models.Post
	err := f.DB.Select(&posts, PostGetOneQuery,
		id,
	)

	return posts, err
}

func (f *PostRepository) PostGetOneUser(post models.Post) ([]models.User, error) {
	var users []models.User
	err := f.DB.Select(&users, PostGetOneUserQuery,
		post.Author,
	)

	return users, err
}

func (f *PostRepository) PostGetOneForum(post models.Post) ([]models.Forum, error) {
	var forums []models.Forum
	err := f.DB.Select(&forums, PostGetOneForumQuery,
		post.Forum,
	)

	return forums, err
}

func (f *PostRepository) PostGetOneThread(post models.Post) ([]models.Thread, error) {
	var threads []models.Thread
	err := f.DB.Select(&threads, PostGetOneThreadQuery,
		post.Thread,
	)

	return threads, err
}

func (f *PostRepository) PostUpdate(post models.Post) error {
	_, err := f.DB.Exec(PostUpdateQuery, post.Message, post.Id)

	return err
}

func (f *PostRepository) GetStatus() (int, error) {
	var ret []int
	err := f.DB.Select(&ret, GetStatusQuery)
	if err != nil {
		return -1, err
	}
	if len(ret) == 0 {
		return -1, errors.New("nothing in return")
	}

	return ret[0], nil
}
