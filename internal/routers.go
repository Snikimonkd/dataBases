package internal

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	user_handler "github.com/Snikimonkd/dataBases/internal/user/delivery"
	user_repository "github.com/Snikimonkd/dataBases/internal/user/repository"
	user_usecase "github.com/Snikimonkd/dataBases/internal/user/usecase"

	forum_handler "github.com/Snikimonkd/dataBases/internal/forum/delivery"
	forum_repository "github.com/Snikimonkd/dataBases/internal/forum/repository"
	forum_usecase "github.com/Snikimonkd/dataBases/internal/forum/usecase"

	post_handler "github.com/Snikimonkd/dataBases/internal/post/delivery"
	post_repository "github.com/Snikimonkd/dataBases/internal/post/repository"
	post_usecase "github.com/Snikimonkd/dataBases/internal/post/usecase"

	thread_handler "github.com/Snikimonkd/dataBases/internal/thread/delivery"
	thread_repository "github.com/Snikimonkd/dataBases/internal/thread/repository"
	thread_usecase "github.com/Snikimonkd/dataBases/internal/thread/usecase"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getPostgres() *sql.DB {
	dsn := "host=localhost user=docker password=docker dbname=docker sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic("cant parse config" + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("can`t ping db" + err.Error())
	}

	db.SetMaxOpenConns(0)

	return db
}

func Init() *sqlx.DB {
	return sqlx.NewDb(getPostgres(), "psx")
}

var db = Init()

var userRepository = user_repository.UserRepository{
	DB: db,
}

var threadRepository = thread_repository.ThreadRepository{
	DB: db,
}

var postRepository = post_repository.PostRepository{
	DB: db,
}

var forumRepository = forum_repository.ForumRepository{
	DB: db,
}

var userUsacese = user_usecase.UserUseCase{
	Repository: userRepository,
}

var threadUsacese = thread_usecase.ThreadUseCase{
	Repository: threadRepository,
}

var postUsacese = post_usecase.PostUseCase{
	Repository: postRepository,
}

var forumUsacese = forum_usecase.ForumUseCase{
	Repository: forumRepository,
}

var user = user_handler.UserHandler{
	Usecase: userUsacese,
}

var thread = thread_handler.ThreadHandler{
	Usecase: threadUsacese,
}

var post = post_handler.PostHandler{
	Usecase: postUsacese,
}

var forum = forum_handler.ForumHandler{
	Usecase: forumUsacese,
}

var a = App{
	User:   user,
	Thread: thread,
	Forum:  forum,
	Post:   post,
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"Clear",
		strings.ToUpper("Post"),
		"/api/service/clear",
		a.Clear,
	},

	Route{
		"ForumCreate",
		strings.ToUpper("Post"),
		"/api/forum/create",
		a.Forum.ForumCreate,
	},

	Route{
		"ForumGetOne",
		strings.ToUpper("Get"),
		"/api/forum/{slug}/details",
		a.Forum.ForumGetOne,
	},

	Route{
		"ForumGetThreads",
		strings.ToUpper("Get"),
		"/api/forum/{slug}/threads",
		a.Forum.ForumGetThreads,
	},

	Route{
		"ForumGetUsers",
		strings.ToUpper("Get"),
		"/api/forum/{slug}/users",
		a.Forum.ForumGetUsers,
	},

	Route{
		"PostGetOne",
		strings.ToUpper("Get"),
		"/api/post/{id}/details",
		a.Post.PostGetOne,
	},

	Route{
		"PostUpdate",
		strings.ToUpper("Post"),
		"/api/post/{id}/details",
		a.Post.PostUpdate,
	},

	Route{
		"PostsCreate",
		strings.ToUpper("Post"),
		"/api/thread/{slug_or_id}/create",
		a.Post.PostsCreate,
	},

	Route{
		"Status",
		strings.ToUpper("Get"),
		"/api/service/status",
		a.Status,
	},

	Route{
		"ThreadCreate",
		strings.ToUpper("Post"),
		"/api/forum/{slug}/create",
		a.Thread.ThreadCreate,
	},

	Route{
		"ThreadGetOne",
		strings.ToUpper("Get"),
		"/api/thread/{slug_or_id}/details",
		a.Thread.ThreadGetOne,
	},

	Route{
		"ThreadGetPosts",
		strings.ToUpper("Get"),
		"/api/thread/{slug_or_id}/posts",
		a.Thread.ThreadGetPosts,
	},

	Route{
		"ThreadUpdate",
		strings.ToUpper("Post"),
		"/api/thread/{slug_or_id}/details",
		a.Thread.ThreadUpdate,
	},

	Route{
		"ThreadVote",
		strings.ToUpper("Post"),
		"/api/thread/{slug_or_id}/vote",
		a.Thread.ThreadVote,
	},

	Route{
		"UserCreate",
		strings.ToUpper("Post"),
		"/api/user/{nickname}/create",
		a.User.UserCreate,
	},

	Route{
		"UserGetOne",
		strings.ToUpper("Get"),
		"/api/user/{nickname}/profile",
		a.User.UserGetOne,
	},

	Route{
		"UserUpdate",
		strings.ToUpper("Post"),
		"/api/user/{nickname}/profile",
		a.User.UserUpdate,
	},
}
