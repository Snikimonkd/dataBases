package repository

var InsertUserQuery = "INSERT INTO users (nickname, fullname, about, email) VALUES ($1, $2, $3, $4)"

var CheckUserBeforeSignUpQuery = "SELECT * FROM users WHERE nickname = $1 OR email = $2"

var SelectUserWithNicknameQuery = "SELECT * FROM users WHERE nickname = $1"

var CheckUserBeforeUpdateQuery = "SELECT * FROM users WHERE (email = $1 or about = $2 or fullname = $3) AND (nickname <> $4)"

var CheckUserExistQuery = "SELECT * FROM users WHERE nickname = $1"

var UpdateUserQuery = "UPDATE users SET fullname = $1, about = $2, email = $3 WHERE nickname = $4"
