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
		models.ResponseError(err.Error(), 400, w, r)
		return
	}

	res, status, err := a.Usecase.ForumCreate(newForum)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w, r)
		return
	}

	models.ResponseJson(res, status, w, r)
}

func (a *ForumHandler) ForumGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w, r)
		return
	}

	res, status, err := a.Usecase.ForumGetOne(slug)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w, r)
		return
	}

	models.ResponseForum(res, status, w)
}

func (a *ForumHandler) ForumGetThreads(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w, r)
		return
	}

	query := r.URL.Query()

	limitInt := 100
	var err error
	limit, ok := query["limit"]
	if ok {
		limitInt, err = strconv.Atoi(limit[0])
		if err != nil {
			log.Println(err)
			models.ResponseError(err.Error(), 404, w, r)
		}
	}

	descBool := false
	desc, ok := query["desc"]
	if ok {
		if desc[0] == "true" {
			descBool = true
		}
	}

	since, ok := query["since"]
	if !ok {
		since = []string{""}
	}

	res, status, err := a.Usecase.ForumGetThreads(slug, limitInt, descBool, since[0])
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w, r)
		return
	}

	models.ResponseThreads(res, status, w)
}

func (a *ForumHandler) ForumGetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w, r)
		return
	}

	query := r.URL.Query()

	limitInt := 100
	var err error
	limit, ok := query["limit"]
	if ok {
		limitInt, err = strconv.Atoi(limit[0])
		if err != nil {
			log.Println(err)
			models.ResponseError(err.Error(), 404, w, r)
		}
	}

	descBool := false
	desc, ok := query["desc"]
	if ok {
		if desc[0] == "true" {
			descBool = true
		}
	}

	since, ok := query["since"]
	if !ok {
		since = []string{""}
	}

	res, status, err := a.Usecase.ForumGetUsers(slug, limitInt, descBool, since[0])
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w, r)
		return
	}

	models.ResponseUsers(res, status, w)
}
