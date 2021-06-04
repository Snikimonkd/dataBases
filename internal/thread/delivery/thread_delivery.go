package delivery

import (
	"errors"
	"log"
	"net/http"

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *ThreadHandler) ThreadUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *ThreadHandler) ThreadVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
