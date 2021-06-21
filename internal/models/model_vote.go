package models

import (
	"io"

	easyjson "github.com/mailru/easyjson"
)

// Информация о голосовании пользователя.
type Vote struct {

	// Идентификатор пользователя.
	Nickname string `json:"nickname"`

	// Отданный голос.
	Voice int `json:"voice" db:"vote"`
}

func ReadVote(body io.ReadCloser) (Vote, error) {
	var newVote Vote
	err := easyjson.UnmarshalFromReader(body, &newVote)
	return newVote, err
}
