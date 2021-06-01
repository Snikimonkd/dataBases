package delivery

import (
	"errors"
	"log"
	"net/http"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/Snikimonkd/dataBases/internal/user/usecase"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Usecase usecase.UserUseCase
}

func (a *UserHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newUser, err := models.ReadUser(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		err = errors.New("нет никнэйма")
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}
	newUser.Nickname = nickname

	res, status, err := a.Usecase.UserCreate(newUser)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *UserHandler) UserGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		err := errors.New("нет никнэйма")
		log.Println(err)
		models.ResponseError(err.Error(), 404, w)
		return
	}

	res, status, err := a.Usecase.UserGetOne(nickname)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}

func (a *UserHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newUser, err := models.ReadUser(r.Body)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}

	vars := mux.Vars(r)
	nickname, ok := vars["nickname"]
	if !ok {
		err = errors.New("нет никнэйма")
		log.Println(err)
		models.ResponseError(err.Error(), 400, w)
		return
	}
	newUser.Nickname = nickname

	res, status, err := a.Usecase.UserUpdate(newUser)
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), status, w)
		return
	}

	models.ResponseJson(res, status, w)
}
