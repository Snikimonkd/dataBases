package internal

import (
	"net/http"

	forum_handler "github.com/Snikimonkd/dataBases/internal/forum/delivery"
	post_handler "github.com/Snikimonkd/dataBases/internal/post/delivery"
	thread_handler "github.com/Snikimonkd/dataBases/internal/thread/delivery"
	user_handler "github.com/Snikimonkd/dataBases/internal/user/delivery"
)

type App struct {
	User   user_handler.UserHandler
	Thread thread_handler.ThreadHandler
	Post   post_handler.PostHandler
	Forum  forum_handler.ForumHandler
}

func (a *App) Clear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (a *App) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
