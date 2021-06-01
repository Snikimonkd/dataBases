package models

import (
	"encoding/json"
	"io"
)

// Информация о пользователе.
type User struct {
	Id int `json:"-"`

	// Имя пользователя (уникальное поле). Данное поле допускает только латиницу, цифры и знак подчеркивания. Сравнение имени регистронезависимо.
	Nickname string `json:"nickname,omitempty"`

	// Полное имя пользователя.
	Fullname string `json:"fullname"`

	// Описание пользователя.
	About string `json:"about,omitempty"`

	// Почтовый адрес пользователя (уникальное поле).
	Email string `json:"email"`
}

type UserUpdate struct {

	// Полное имя пользователя.
	Fullname string `json:"fullname,omitempty"`

	// Описание пользователя.
	About string `json:"about,omitempty"`

	// Почтовый адрес пользователя (уникальное поле).
	Email string `json:"email,omitempty"`
}

func ReadUser(body io.ReadCloser) (User, error) {
	var newUser User
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&newUser)
	return newUser, err
}
