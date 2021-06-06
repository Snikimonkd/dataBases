package usecase

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/Snikimonkd/dataBases/internal/models"
	post_repository "github.com/Snikimonkd/dataBases/internal/post/repository"
)

type PostUseCase struct {
	Repository post_repository.PostRepository
}

func (u *PostUseCase) PostsCreate(slug_or_id string, newPosts []models.Post) (interface{}, int, error) {
	var threads []models.Thread
	id, err := strconv.Atoi(slug_or_id)
	if err == nil {
		threads, err = u.Repository.ThreadGetOneId(id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 1")
		}
	} else {
		threads, err = u.Repository.ThreadGetOneSlug(slug_or_id)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(threads) == 0 {
			return nil, 404, errors.New("cant find thread 2")
		}
	}

	thread := threads[0]

	if len(newPosts) == 0 {
		res := make([]models.Post, 0)
		return res, 201, nil
	}

	users, err := u.Repository.CheckUsers(newPosts[0].Author)
	if err != nil {
		return nil, 500, err
	}
	if len(users) == 0 {
		return nil, 404, errors.New("can`t find user")
	}

	if newPosts[0].Parent != 0 {
		parents, err := u.Repository.CheckParents(newPosts[0].Parent)
		if err != nil {
			log.Println(err)
			return nil, 500, nil
		}
		if len(parents) == 0 {
			return nil, 409, errors.New("parent post was created in another thread 1")
		}
		if parents[0] != thread.Id {
			return nil, 409, errors.New("parent post was created in another thread 2")
		}
	}

	newPosts, err = u.Repository.PostsCreate(newPosts, thread)
	if err != nil {
		log.Println(err)
		return nil, 500, nil

	}

	return newPosts, 201, nil
}

func (u *PostUseCase) PostGetOne(id int, related string) (interface{}, int, error) {
	posts, err := u.Repository.PostGetOne(id)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}
	if len(posts) == 0 {
		return nil, 404, errors.New("cant find post")
	}

	post := posts[0]

	res := models.PostDetails{
		Post: &post,
	}

	if related == "" {
		return res, 200, nil
	}

	relatedArr := strings.Split(related, ",")

	for _, v := range relatedArr {
		switch v {
		case "user":
			{
				users, err := u.Repository.PostGetOneUser(post)
				if err != nil {
					log.Println(err)
					return nil, 500, nil
				}
				if len(users) == 0 {
					return nil, 404, errors.New("cant find user")
				}
				res.Author = &users[0]
			}
		case "forum":
			{
				forums, err := u.Repository.PostGetOneForum(post)
				if err != nil {
					log.Println(err)
					return nil, 500, nil
				}
				if len(forums) == 0 {
					return nil, 404, errors.New("cant find forum")
				}
				res.Forum = &forums[0]
			}
		case "thread":
			{
				threads, err := u.Repository.PostGetOneThread(post)
				if err != nil {
					log.Println(err)
					return nil, 500, nil
				}
				if len(threads) == 0 {
					return nil, 404, errors.New("cant find forum")
				}
				res.Thread = &threads[0]
			}
		}
	}

	return res, 200, nil
}

func (u *PostUseCase) PostUpdate(newPost models.Post) (interface{}, int, error) {
	posts, err := u.Repository.PostGetOne(newPost.Id)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}
	if len(posts) == 0 {
		return nil, 404, errors.New("cant find post")
	}

	post := posts[0]
	if newPost.Message == "" || newPost.Message == post.Message {
		return post, 200, nil
	}

	post.Message = newPost.Message

	err = u.Repository.PostUpdate(newPost)
	if err != nil {
		log.Println(err)
		return nil, 500, nil
	}

	post.IsEdited = true
	return post, 200, nil
}

func (u *PostUseCase) GetStatus() (int, error) {
	return u.Repository.GetStatus()
}
