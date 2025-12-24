package entity

import (
	"errors"

	"github.com/flp-fernandes/9-APIS/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNameIsRequired = errors.New("name is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("email is required")
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.ValidateUserInfo(name, email, password)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	return user, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) ValidateUserInfo(name, email, password string) error {
	if u.Name == "" {
		return ErrNameIsRequired
	}
	if u.Email == "" {
		return ErrEmailIsRequired
	}
	if u.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}
