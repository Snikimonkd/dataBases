/*
 * forum
 *
 * Тестовое задание для реализации проекта \"Форумы\" на курсе по базам данных в Технопарке Mail.ru (https://park.mail.ru).
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package delivery

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB
}

type HandlerInterface interface {
	Clear(w http.ResponseWriter, r *http.Request)
	ForumCreate(w http.ResponseWriter, r *http.Request)
	ForumGetOne(w http.ResponseWriter, r *http.Request)
	ForumGetThreads(w http.ResponseWriter, r *http.Request)
	ForumGetUsers(w http.ResponseWriter, r *http.Request)
	PostGetOne(w http.ResponseWriter, r *http.Request)
	PostUpdate(w http.ResponseWriter, r *http.Request)
	PostsCreate(w http.ResponseWriter, r *http.Request)
	Status(w http.ResponseWriter, r *http.Request)
	ThreadCreate(w http.ResponseWriter, r *http.Request)
	ThreadGetOne(w http.ResponseWriter, r *http.Request)
	ThreadGetPosts(w http.ResponseWriter, r *http.Request)
	ThreadUpdate(w http.ResponseWriter, r *http.Request)
	ThreadVote(w http.ResponseWriter, r *http.Request)
	UserCreate(w http.ResponseWriter, r *http.Request)
	UserGetOne(w http.ResponseWriter, r *http.Request)
	UserUpdate(w http.ResponseWriter, r *http.Request)
}

func (a *App) Clear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ForumCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ForumGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ForumGetThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ForumGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) PostGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) PostUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) PostsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ThreadCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ThreadGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ThreadGetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ThreadUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ThreadVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) UserGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}