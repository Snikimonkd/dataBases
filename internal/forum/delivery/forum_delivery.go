package delivery

import (
	"log"
	"net/http"

	"github.com/Snikimonkd/dataBases/internal/forum/usecase"
	"github.com/Snikimonkd/dataBases/internal/models"
)

type ForumHandler struct {
	Usecase usecase.ForumUseCase
}

func (a *ForumHandler) ForumCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newForum, err := models.ReadUser(r.Body)
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *ForumHandler) ForumGetThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *ForumHandler) ForumGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
