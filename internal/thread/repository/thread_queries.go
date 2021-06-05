package repository

var CheckUserExistQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckForumExistQuery = "SELECT * FROM forums WHERE slug = $1"

var CheckThreadExistQuery = "SELECT * FROM threads WHERE slug = $1 AND slug <> ''"

var InsertThreadQuery = "INSERT INTO threads (title, author, forum, message, created, slug) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

var SelectThreadWithIdQuery = "SELECT * FROM threads WHERE id = $1"

var SelectThreadWithSlugQuery = "SELECT * FROM threads WHERE slug = $1"

var ThreadVoteQuery = "UPDATE threads SET votes = vote + $1 WHERE id = $2"

var SelectVoteQuery = "SELECT nickname, vote FROM votes WHERE nickname = $1 AND thread_id = $2"

var InsertVoteQuery = "INSERT INTO votes (nickname, thread_id, vote) VALUES ($1, $2, $3)"

var UpdateVoteQuery = "UPDATE votes SET vote = $1 WHERE thread_id = $2 AND nickname = $3"

//-----------

var InsertUserQuery = "INSERT INTO users (nickname, fullname, about, email) VALUES ($1, $2, $3, $4)"

var CheckUserBeforeSignUpQuery = "SELECT * FROM users WHERE nickname = $1 OR email = $2"

var SelectUserWithNicknameQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckUserBeforeUpdateQuery = "SELECT * FROM users WHERE (email = $1 or about = $2 or fullname = $3) AND (nickname <> $4)"

var UpdateUserQuery = "UPDATE users SET fullname = $1, about = $2, email = $3 WHERE nickname = $4"
