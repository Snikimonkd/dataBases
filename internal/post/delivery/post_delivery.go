package delivery

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/Snikimonkd/dataBases/internal/post/usecase"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Usecase usecase.PostUseCase
}

func (a *PostHandler) PostGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("пользователя с таким id нет")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	query := r.URL.Query()
	related, ok := query["related"]
	if !ok {
		related = []string{""}
	}

	res, status, err := a.Usecase.PostGetOne(idInt, related[0])
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *PostHandler) PostUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newPost, err := models.ReadPost(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("пользователя с таким id нет")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	newPost.Id = idInt

	res, status, err := a.Usecase.PostUpdate(newPost)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *PostHandler) PostsCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newPosts, err := models.ReadPosts(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	slug, ok := vars["slug_or_id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.PostsCreate(slug, newPosts)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}
