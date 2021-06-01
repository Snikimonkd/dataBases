package usecase

import (
	forum_repository "github.com/Snikimonkd/dataBases/internal/forum/repository"
	"github.com/Snikimonkd/dataBases/internal/models"
)

type ForumUseCase struct {
	Repository forum_repository.ForumRepository
}

func (u *ForumUseCase) ForumCreate(newUser models.Forum) (interface{}, int, error) {
	users, err := u.Repository.CheckSignUp(newUser)
	if err != nil {
		return nil, 500, err
	}
	if len(users) != 0 {
		return users, 409, nil
	}

	err = u.Repository.UserCreate(newUser)
	if err != nil {
		return nil, 500, err
	}

	return newUser, 201, nil
}
