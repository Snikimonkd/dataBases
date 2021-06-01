package repository

import (
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

func (u *UserRepository) Check(newUser models.User) ([]models.User, error) {
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
