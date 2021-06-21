package models

import (
	"io"
	"log"
	"net/http"

	"github.com/mailru/easyjson"
)

// Информация о форуме.
type Forum struct {

	// Название форума.
	Title string `json:"title"`

	// Nickname пользователя, который отвечает за форум.
	User string `json:"user" db:"user_nickname"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL), уникальное поле.
	Slug string `json:"slug"`

	// Общее кол-во сообщений в данном форуме.
	Posts int `json:"posts,omitempty"`

	// Общее кол-во ветвей обсуждения в данном форуме.
	Threads int `json:"threads,omitempty"`
}

func ResponseForum(res Forum, status int, w http.ResponseWriter) {
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

func ReadForum(body io.ReadCloser) (Forum, error) {
	var newForum Forum
	err := easyjson.UnmarshalFromReader(body, &newForum)
	return newForum, err
}
