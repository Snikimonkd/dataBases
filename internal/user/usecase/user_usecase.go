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

func (u *UserUseCase) CheckFields(oldUser, newUser models.User) models.User {
	if newUser.Fullname != "" {
		oldUser.Fullname = newUser.Fullname
	}

	if newUser.About != "" {
		oldUser.About = newUser.About
	}

	if newUser.Email != "" {
		oldUser.Email = newUser.Email
	}

	return oldUser
}

func (u *UserUseCase) UserUpdate(newUser models.User) (interface{}, int, error) {
	users, err := u.Repository.CheckExist(newUser)
	if err != nil {
		return nil, 500, err
	}
	if len(users) == 0 {
		return nil, 404, errors.New("can't find user")
	}

	newUser = u.CheckFields(users[0], newUser)

	users, err = u.Repository.CheckUpdateData(newUser)
	if err != nil {
		return nil, 500, err
	}
	if len(users) > 0 {
		return nil, 409, errors.New("new data conflicts with existing data")
	}

	err = u.Repository.UserUpdate(newUser)
	if err != nil {
		return nil, 500, err
	}

	return newUser, 200, nil
}

func (u *UserUseCase) GetStatus() (int, error) {
	return u.Repository.GetStatus()
}

func (u *UserUseCase) Clear() error {
	return u.Repository.Clear()
}
