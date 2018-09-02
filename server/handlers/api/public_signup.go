package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	user "github.com/medtune/beta-platform/pkg/service/users"
	"github.com/medtune/beta-platform/pkg/session"
)

// Signup page handler
func Signup(c *gin.Context) {

	signupData := &jsonutil.SignupData{}
	var err error

	// Parse request body (json)
	err = c.ShouldBindJSON(&signupData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ok, err := user.Signup(signupData)
	if err != nil || !ok {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	session.SetLoginStatus(c, true, false)
	c.JSON(200, jsonutil.Success())
}
