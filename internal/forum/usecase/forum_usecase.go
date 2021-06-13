package usecase

import (
	"errors"

	forum_repository "github.com/Snikimonkd/dataBases/internal/forum/repository"
	"github.com/Snikimonkd/dataBases/internal/models"
)

type ForumUseCase struct {
	Repository forum_repository.ForumRepository
}

func (u *ForumUseCase) ForumCreate(newForum models.Forum) (interface{}, int, error) {
	users, err := u.Repository.CheckUsers(newForum)
	if err != nil {
		return nil, 500, err
	}
	if len(users) == 0 {
		return nil, 404, errors.New("can`t find user")
	}

	newForum.User = users[0].Nickname
	forums, err := u.Repository.CheckForums(newForum)
	if err != nil {
		return nil, 500, err
	}
	if len(forums) != 0 {
		return forums[0], 409, nil
	}

	err = u.Repository.CreateForum(newForum)
	if err != nil {
		return nil, 500, err
	}

	return newForum, 201, nil
}

func (u *ForumUseCase) ForumGetOne(slug string) (interface{}, int, error) {
	forums, err := u.Repository.ForumGetOne(slug)
	if err != nil {
		return nil, 500, err
	}
	if len(forums) == 0 {
		return nil, 404, errors.New("can`t find forum")
	}

	return forums[0], 200, nil
}

func (u *ForumUseCase) ForumGetThreads(slug string, limitInt int, descBool bool, since string) (interface{}, int, error) {
	var newForum models.Forum
	newForum.Slug = slug
	forums, err := u.Repository.CheckForums(newForum)
	if err != nil {
		return nil, 500, err
	}
	if len(forums) == 0 {
		return nil, 404, errors.New("can`t find forum")
	}

	threads, err := u.Repository.ForumGetThreads(slug, limitInt, descBool, since)
	if err != nil {
		return nil, 500, err
	}

	return threads, 200, err
}

func (u *ForumUseCase) ForumGetUsers(slug string, limitInt int, descBool bool, since string) (interface{}, int, error) {
	var newForum models.Forum
	newForum.Slug = slug
	forums, err := u.Repository.CheckForums(newForum)
	if err != nil {
		return nil, 500, err
	}
	if len(forums) == 0 {
		return nil, 404, errors.New("can`t find forum")
	}

	users, err := u.Repository.ForumGetUsers(slug, limitInt, descBool, since)
	if err != nil {
		return nil, 500, err
	}

	return users, 200, err
}

func (u *ForumUseCase) GetStatus() (int, error) {
	return u.Repository.GetStatus()
}
