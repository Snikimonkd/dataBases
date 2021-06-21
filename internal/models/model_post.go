package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-openapi/strfmt"
	easyjson "github.com/mailru/easyjson"
)

// Сообщение внутри ветки обсуждения на форуме.
type Post struct {

	// Идентификатор данного сообщения.
	Id int `json:"id,omitempty"`

	// Идентификатор родительского сообщения (0 - корневое сообщение обсуждения).
	Parent int `json:"parent,omitempty"`

	// Автор, написавший данное сообщение.
	Author string `json:"author"`

	// Собственно сообщение форума.
	Message string `json:"message"`

	// Истина, если данное сообщение было изменено.
	IsEdited bool `json:"isEdited,omitempty"`

	// Идентификатор форума (slug) данного сообещния.
	Forum string `json:"forum,omitempty"`

	// Идентификатор ветви (id) обсуждения данного сообещния.
	Thread int `json:"thread,omitempty"`

	// Дата создания сообщения на форуме.
	Created strfmt.DateTime `json:"created,omitempty"`
}

type PostDetails struct {
	Post   *Post   `json:"post"`
	Author *User   `json:"author,omitempty"`
	Thread *Thread `json:"thread,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
}

//easyjson:json
type Posts []Post

func ReadPost(body io.ReadCloser) (Post, error) {
	var newPost Post
	err := easyjson.UnmarshalFromReader(body, &newPost)
	return newPost, err
}

func ReadPosts(body io.ReadCloser) ([]Post, error) {
	var newPosts []Post
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&newPosts)
	return newPosts, err
}

func ResponsePost(res Post, status int, w http.ResponseWriter) {
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

func ResponsePosts(res Posts, status int, w http.ResponseWriter) {
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

func ResponsePostDetails(res PostDetails, status int, w http.ResponseWriter) {
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
