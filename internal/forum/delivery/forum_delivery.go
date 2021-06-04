package delivery

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Snikimonkd/dataBases/internal/forum/usecase"
	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/gorilla/mux"
)

type ForumHandler struct {
	Usecase usecase.ForumUseCase
}

func (a *ForumHandler) ForumCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newForum, err := models.ReadForum(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	res, status, err := a.Usecase.ForumCreate(newForum)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ForumHandler) ForumGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.ForumGetOne(slug)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ForumHandler) ForumGetThreads(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	limitInt := 100
	var err error
	limit, ok := vars["limit"]
	if ok {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			log.Println(err)
			models.ResponseError(err.Error(), 404, w)
		}
	}

	descBool := false
	desc, ok := vars["desc"]
	if ok {
		if desc == "true" {
			descBool = true
		}
	}

	since, ok := vars["since"]
	if !ok {
		since = ""
	}

	res, status, err := a.Usecase.ForumGetThreads(slug, limitInt, descBool, since)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ForumHandler) ForumGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
