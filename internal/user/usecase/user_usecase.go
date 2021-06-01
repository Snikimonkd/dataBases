package usecase

import (
	"errors"

	"github.com/Snikimonkd/dataBases/internal/models"
	user_repository "github.com/Snikimonkd/dataBases/internal/user/repository"
)

type UserUseCase struct {
	Repository user_repository.UserRepository
}

func (u *UserUseCase) UserCreate(newUser models.User) (interface{}, int, error) {
	users, err := u.Repository.Check(newUser)
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

func (u *UserUseCase) UserGetOne(nickname string) (interface{}, int, error) {
	users, err := u.Repository.UserGetOne(nickname)
	if err != nil {
		return nil, 500, err
	}

	if len(users) == 0 {
		return nil, 404, errors.New("can't find user")
	}

	return users[0], 200, nil
}
