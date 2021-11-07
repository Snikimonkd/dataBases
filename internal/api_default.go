package internal

import (
	"log"
	"net/http"

	forum_handler "github.com/Snikimonkd/dataBases/internal/forum/delivery"
	"github.com/Snikimonkd/dataBases/internal/models"
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
	err := a.User.Usecase.Clear()
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 500, w, r)
		return
	}

	models.ResponseJson(nil, 200, w, r)
}

func (a *App) Status(w http.ResponseWriter, r *http.Request) {
	var st models.Status

	users, err := a.User.Usecase.GetStatus()
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 500, w, r)
		return
	}
	st.User = users

	forums, err := a.Forum.Usecase.GetStatus()
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 500, w, r)
		return
	}
	st.Forum = forums

	threads, err := a.Thread.Usecase.GetStatus()
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 500, w, r)
		return
	}
	st.Thread = threads

	posts, err := a.Post.Usecase.GetStatus()
	if err != nil {
		log.Println(err)
		models.ResponseError(err.Error(), 500, w, r)
		return
	}
	st.Post = posts

	models.ResponseJson(st, 200, w, r)
}
