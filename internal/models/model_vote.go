package models

import (
	"encoding/json"
	"io"
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
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&newVote)
	return newVote, err
}
