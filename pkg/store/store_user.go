package store

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store/model"
	"github.com/medtune/go-utils/crypto"
)

const (
	ACCOUNT_ADMIN      = "admin"
	ACCOUNT_BETATESTER = "betatester"
	ACCOUNT_SII        = "sii"
)

type userStore interface {
	CreateUser(string, string, string) error
	AuthentificateUser(string, string) (bool, error)
	GetUser(string) (*model.User, error)
}

func (s *Store) CreateUser(email, username, password string) error {
	user := model.User{
		Email:         email,
		Username:      username,
		Password:      crypto.Sha256(password),
		AccountStatus: true,
		AccountType:   ACCOUNT_BETATESTER,
		AccountLevel:  5,
	}
	v, err := s.Valid(user)
	if err != nil || !v {
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
		return nil, fmt.Errorf("record doesnt exist")
	}
	return user, nil
}

func (s *Store) getUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	has, err := s.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("record doesnt exist")
	}
	return user, nil
}

func (s *Store) AuthentificateUser(username, password string) (bool, error) {
	user, err := s.GetUser(username)
	if err != nil {
		// Database server error or record not found
		return false, fmt.Errorf("username or password incorrect")
	}

	if crypto.Sha256(password) == user.Password {
		return true, nil
	}

	// Password is incorrect
	return false, fmt.Errorf("username or password incorrect")
}
