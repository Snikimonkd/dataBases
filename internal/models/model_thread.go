package models

import (
	"io"
	"log"
	"net/http"
	"time"

	easyjson "github.com/mailru/easyjson"
)

// Ветка обсуждения на форуме.
type Thread struct {

	// Идентификатор ветки обсуждения.
	Id int `json:"id,omitempty"`

	// Заголовок ветки обсуждения.
	Title string `json:"title"`

	// Пользователь, создавший данную тему.
	Author string `json:"author"`

	// Форум, в котором расположена данная ветка обсуждения.
	Forum string `json:"forum,omitempty"`

	// Описание ветки обсуждения.
	Message string `json:"message"`

	// Кол-во голосов непосредственно за данное сообщение форума.
	Votes int `json:"votes,omitempty"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL). В данной структуре slug опционален и не может быть числом.
	Slug string `json:"slug,omitempty"`

	// Дата создания ветки на форуме.
	Created time.Time `json:"created,omitempty"`
}

//easyjson:json
type Threads []Thread

func ReadThread(body io.ReadCloser) (Thread, error) {
	var newThread Thread
	err := easyjson.UnmarshalFromReader(body, &newThread)
	return newThread, err
}

func ResponseThread(res Thread, status int, w http.ResponseWriter) {
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

func ResponseThreads(res Threads, status int, w http.ResponseWriter) {
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
