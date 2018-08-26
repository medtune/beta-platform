package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/store"
)

// Auth authentificate user from json data
func Auth(loginData *jsonutil.LoginData) (bool, error) {
	// Validate data
	ok, err := govalidator.ValidateStruct(loginData)
	if err != nil || !ok {
		return false, err
	}

	// Authentificate user
	ok, err = store.Agent.AuthentificateUser(loginData.Username, loginData.Password)
	if err != nil || !ok {
		return false, err
	}

	return true, nil
}
