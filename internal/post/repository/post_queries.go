package repository

var SelectThreadWithIdQuery = "SELECT * FROM threads WHERE id = $1"

var SelectThreadWithSlugQuery = "SELECT * FROM threads WHERE slug = $1"

var SelectParentQuery = "SELECT thread FROM posts WHERE id = $1"

var InsertPostQuery = "INSERT INTO posts (author, created, forum, message, parent, thread) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

var CheckUserExistQuery = "SELECT * FROM users WHERE nickname = $1"

var PostGetOneQuery = "SELECT id, parent, author, message, isEdited, forum, thread, created FROM posts where id = $1"

var PostGetOneUserQuery = "SELECT * FROM users WHERE nickname = $1"

var PostGetOneForumQuery = "SELECT * FROM forums WHERE slug = $1"

var PostGetOneThreadQuery = "SELECT * FROM threads WHERE id = $1"

var PostUpdateQuery = "UPDATE posts SET message = $1, isEdited = TRUE WHERE id = $2 RETURNING id, parent, author, message, isEdited, forum, thread, created"

//-----------

var CheckForumExistQuery = "SELECT * FROM forums WHERE slug = $1"

var CheckThreadExistQuery = "SELECT * FROM threads WHERE slug = $1 AND slug <> ''"

var InsertThreadQuery = "INSERT INTO threads (title, author, forum, message, created, slug) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

var InsertUserQuery = "INSERT INTO users (nickname, fullname, about, email) VALUES ($1, $2, $3, $4)"

var CheckUserBeforeSignUpQuery = "SELECT * FROM users WHERE nickname = $1 OR email = $2"

var SelectUserWithNicknameQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckUserBeforeUpdateQuery = "SELECT * FROM users WHERE (email = $1 or about = $2 or fullname = $3) AND (nickname <> $4)"

var UpdateUserQuery = "UPDATE users SET fullname = $1, about = $2, email = $3 WHERE nickname = $4"
