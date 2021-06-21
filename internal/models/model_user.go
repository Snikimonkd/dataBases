package models

import (
	"io"
	"log"
	"net/http"

	easyjson "github.com/mailru/easyjson"
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

//easyjson:json
type Users []User

func ReadUser(body io.ReadCloser) (User, error) {
	var newUser User
	err := easyjson.UnmarshalFromReader(body, &newUser)
	return newUser, err
}

func ResponseUsers(res Users, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := res.MarshalJSON()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

func ResponseUser(res User, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := res.MarshalJSON()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
