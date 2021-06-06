package delivery

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/Snikimonkd/dataBases/internal/thread/usecase"
	"github.com/gorilla/mux"
)

type ThreadHandler struct {
	Usecase usecase.ThreadUseCase
}

func (a *ThreadHandler) ThreadCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newThread, err := models.ReadThread(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	newThread.Forum = slug

	res, status, err := a.Usecase.ThreadCreate(newThread)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ThreadHandler) ThreadGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug_or_id, ok := vars["slug_or_id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.ThreadGetOne(slug_or_id)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ThreadHandler) ThreadGetPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug_or_id, ok := vars["slug_or_id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
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
			models.ResponseError(err.Error(), 404, w)
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

	sort, ok := query["sort"]
	if !ok {
		sort = []string{"flat"}
	}

	res, status, err := a.Usecase.ThreadGetPosts(slug_or_id, limitInt, descBool, since[0], sort[0])
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ThreadHandler) ThreadUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newThread, err := models.ReadThread(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	slug_or_id, ok := vars["slug_or_id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.ThreadUpdate(slug_or_id, newThread)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *ThreadHandler) ThreadVote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newVote, err := models.ReadVote(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	slug_or_id, ok := vars["slug_or_id"]
	if !ok {
		err := errors.New("нет идентификатора")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.ThreadVote(slug_or_id, newVote)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}
