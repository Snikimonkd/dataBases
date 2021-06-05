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

func (f *ThreadRepository) ThreadVote(vote models.Vote, thread models.Thread) error {
	_, err := f.DB.Exec(ThreadVoteQuery,
		vote.Voice, thread.Id,
	)

	return err
}

func (f *ThreadRepository) GetVote(vote models.Vote, threadId int) ([]models.Vote, error) {
	var votes []models.Vote
	err := f.DB.Select(&votes, SelectVoteQuery,
		vote.Nickname, threadId,
	)

	return votes, err
}

func (f *ThreadRepository) InsertNewVote(vote models.Vote, threadId int) error {
	_, err := f.DB.Exec(InsertVoteQuery,
		vote.Nickname, threadId, vote.Voice,
	)

	return err
}

func (f *ThreadRepository) UpdateVote(vote models.Vote, threadId int) error {
	_, err := f.DB.Exec(UpdateVoteQuery,
		vote.Voice, threadId, vote.Nickname,
	)

	return err
}

func (f *ThreadRepository) ThreadGetPostsFlat(limitInt int, descBool bool, since string, thread models.Thread) ([]models.Post, error) {
	query := "SELECT id, author, created, forum, isEdited, message, parent, thread FROM posts WHERE thread = $1"
	queryLimit := " ORDER BY created"

	if since != "" {
		if descBool {
			query += " AND id < " + since
		} else {
			query += " AND id > " + since
		}
	}

	if descBool {
		query += queryLimit + " DESC, id DESC"
	} else {
		query += queryLimit + " ASC, id ASC"
	}
	query += " LIMIT $2"

	var posts []models.Post

	err := f.DB.Select(&posts, query,
		thread.Id, limitInt,
	)

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	return posts, err
}

func (f *ThreadRepository) ThreadGetPostsTree(limitInt int, descBool bool, since string, thread models.Thread) ([]models.Post, error) {
	query := "SELECT id, author, created, forum, message, parent, thread FROM posts WHERE thread = $1"
	queryLimit := " ORDER BY tree"

	if since != "" {
		if descBool {
			query += " AND tree < (SELECT tree FROM posts WHERE id = " + since + ")"
		} else {
			query += " AND tree > (SELECT tree FROM posts WHERE id = " + since + ")"
		}
	}

	if descBool {
		query += queryLimit + " DESC, id DESC"
	} else {
		query += queryLimit + " ASC, id ASC"
	}

	query += " LIMIT $2"

	var posts []models.Post

	err := f.DB.Select(&posts, query,
		thread.Id, limitInt,
	)

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	return posts, err
}

func (f *ThreadRepository) ThreadGetPostsParentTree(limitInt int, descBool bool, since string, thread models.Thread) ([]models.Post, error) {
	query := "(SELECT id FROM posts WHERE thread = $1 AND parent = 0"
	queryLimit := " ORDER BY id"

	if since != "" {
		query += " AND tree[1]"
		if descBool {
			query += " < " + "(SELECT tree[1] FROM posts where id = " + since + ")" + queryLimit + " DESC"
		} else {
			query += " > " + "(SELECT tree[1] FROM posts where id = " + since + ")" + queryLimit + " ASC"
		}
	} else {
		if descBool {
			query += queryLimit + " DESC"
		} else {
			query += queryLimit + " ASC"
		}
	}

	query += " LIMIT $2)"

	query = "SELECT id, author, created, forum, isEdited, message, parent, thread FROM posts where tree[1] IN " + query

	if descBool {
		query += " ORDER BY tree[1] DESC, tree, id"
	} else {
		query += " ORDER BY tree, id"
	}

	var posts []models.Post

	err := f.DB.Select(&posts, query,
		thread.Id, limitInt,
	)

	if len(posts) == 0 {
		posts = make([]models.Post, 0)
	}

	return posts, err
}
