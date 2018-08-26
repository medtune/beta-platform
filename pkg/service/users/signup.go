package user

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/secret"
	"github.com/medtune/beta-platform/pkg/store"
)

// Signup will 'try' to create a new user from json req
func Signup(signupData *jsonutil.SignupData) (bool, error) {
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
