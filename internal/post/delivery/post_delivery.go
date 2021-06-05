package delivery

import (
	"errors"
	"log"
	"net/http"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/Snikimonkd/dataBases/internal/post/usecase"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Usecase usecase.PostUseCase
}

func (a *PostHandler) PostGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *PostHandler) PostUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
