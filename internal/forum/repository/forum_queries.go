package repository

var CheckUserExistQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckForumExistQuery = "SELECT * FROM forums WHERE slug = $1"

var InsertForumQuery = "INSERT INTO forums (title, user_nickname, slug) VALUES ($1, $2, $3)"

var SelectForumQuery = "SELECT * FROM forums WHERE slug = $1"

var SelectThreadsQuery = "SELECT * FROM threads WHERE forum = $1 ORDER BY created LIMIT $2"

var SelectThreadsQueryDesc = "SELECT * FROM threads WHERE forum = $1 ORDER BY created DESC LIMIT $2"

var GetStatusQuery = "SELECT COUNT(*) FROM forums"

//--------------------

var InsertUserQuery = "INSERT INTO users (nickname, fullname, about, email) VALUES ($1, $2, $3, $4)"

var CheckUserBeforeSignUpQuery = "SELECT * FROM users WHERE nickname = $1 OR email = $2"

var SelectUserWithNicknameQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckUserBeforeUpdateQuery = "SELECT * FROM users WHERE (email = $1 or about = $2 or fullname = $3) AND (nickname <> $4)"

var UpdateUserQuery = "UPDATE users SET fullname = $1, about = $2, email = $3 WHERE nickname = $4"
