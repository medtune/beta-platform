package store

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store/model"
)

type userStore interface {
	CreateUser(string, string) error
	AuthentificateUser(string, string) (bool, error)
	GetUser(string) (*model.User, error)
}

func (s *Store) CreateUser(username, password string) error {
	user := model.User{
		Username: username,
		Password: password,
	}
	if v, err := s.valid(&user); err != nil || !v {
		return err
	}
	if _, err := s.Insert(&user); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUser(username string) (*model.User, error) {
	user := &model.User{}
	has, err := s.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("Record doesnt exist")
	}
	// Hide password
	user.Password = ""
	return user, nil
}

func (s *Store) AuthentificateUser(username, password string) (bool, error) {
	user, err := s.GetUser(username)
	if err != nil {
		return false, err
	}
	if password == user.Password {
		user.Password = ""
		return true, nil
	}
	return false, nil
}
