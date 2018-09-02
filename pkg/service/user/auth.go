package user

import (
	"github.com/asaskevich/govalidator"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/store"
)

// Auth authentificate user from json data
func Auth(loginData *jsonutil.LoginData) (bool, bool, error) {
	// Validate data
	ok, err := govalidator.ValidateStruct(loginData)
	if err != nil || !ok {
		return false, false, err
	}

	// Authentificate user
	ok, isadmin, err := store.Agent.AuthentificateUser(loginData.Username, loginData.Password)
	if err != nil || !ok {
		return false, false, err
	}

	return true, isadmin, nil
}
