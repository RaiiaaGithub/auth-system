package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
}

type authUser struct {
	email        string
	passwordHash string
}

var authUserDb = map[string]authUser{}

var DefaultUserService userLocalService

type userLocalService struct {
	authUserDb map[string]authUser
}

func (userLocalService) CreateUser(newUser User) error {
	_, ok := authUserDb[newUser.Email]
	if ok {
		return errors.New("User already exists")
	}
	passwordHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return err
	}
	newAuthUser := authUser{
		email:        newUser.Email,
		passwordHash: passwordHash,
	}
	authUserDb[newAuthUser.email] = newAuthUser
	return nil
}

func (userLocalService) VerifyUser(user User) error {
	authUser, ok := authUserDb[user.Email]
	if !ok {
		return errors.New("User don't exist")
	}

	err := decryptPasswordHash(authUser, user)
	if err != nil {
		return errors.New("Password don't match")
	}

	return nil
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}

func decryptPasswordHash(authUser authUser, user User) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password))
}
