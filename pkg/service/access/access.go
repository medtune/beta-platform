package access

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/secret"
	"github.com/medtune/beta-platform/pkg/store"
)

// AuthUser authentificate user from json data
func AuthUser(loginData *jsonutil.LoginData) (bool, error) {
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

// CreateUser will create a user from json data
func CreateUser(signupData *jsonutil.SignupData) (bool, error) {
	// Validate data
	ok, err := govalidator.ValidateStruct(signupData)
	if err != nil || !ok {
		return false, err
	}

	// Password confirmation
	if signupData.Password != signupData.PasswordConfirm {
		return false, fmt.Errorf("password confirmation error")
	}

	// Signup secret unknown
	ok, err = secret.Check("signup", signupData.Secret)
	if err != nil || !ok {
		return ok, err
	}

	// Try to create user in database
	err = store.Agent.CreateUser(signupData.Email, signupData.Username, signupData.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}
