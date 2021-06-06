package repository

import (
	"errors"

	"github.com/Snikimonkd/dataBases/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (u *UserRepository) UserCreate(newUser models.User) error {
	_, err := u.DB.Exec(InsertUserQuery,
		newUser.Nickname, newUser.Fullname, newUser.About, newUser.Email,
	)

	return err
}

func (u *UserRepository) CheckSignUp(newUser models.User) ([]models.User, error) {
	var users []models.User
	err := u.DB.Select(&users, CheckUserBeforeSignUpQuery,
		newUser.Nickname, newUser.Email,
	)

	return users, err
}

func (u *UserRepository) UserGetOne(nickname string) ([]models.User, error) {
	var users []models.User
	err := u.DB.Select(&users, SelectUserWithNicknameQuery,
		nickname,
	)

	return users, err
}

func (u *UserRepository) UserUpdate(newUser models.User) error {
	_, err := u.DB.Exec(UpdateUserQuery,
		newUser.Fullname, newUser.About, newUser.Email, newUser.Nickname,
	)

	return err
}

func (u *UserRepository) CheckExist(newUser models.User) ([]models.User, error) {
	var users []models.User
	err := u.DB.Select(&users, CheckUserExistQuery,
		newUser.Nickname,
	)

	return users, err
}

func (u *UserRepository) CheckUpdateData(newUser models.User) ([]models.User, error) {
	var users []models.User
	err := u.DB.Select(&users, CheckUserBeforeUpdateQuery,
		newUser.Email, newUser.Nickname,
	)

	return users, err
}

func (u *UserRepository) GetStatus() (int, error) {
	var ret []int
	err := u.DB.Select(&ret, GetStatusQuery)
	if err != nil {
		return -1, err
	}
	if len(ret) == 0 {
		return -1, errors.New("nothing in return")
	}

	return ret[0], nil
}

func (u *UserRepository) Clear() error {
	_, err := u.DB.Exec(ClearUsersQuery)

	return err
}
