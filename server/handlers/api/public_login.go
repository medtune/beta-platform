package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	user "github.com/medtune/beta-platform/pkg/service/users"
	"github.com/medtune/beta-platform/pkg/session"
)

// Login page handler
func Login(c *gin.Context) {

	logins := jsonutil.LoginData{}
	var err error

	// Parse request body
	err = c.ShouldBindJSON(&logins)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Service authentificate user
	ok, isadmin, err := user.Auth(&logins)
	if err != nil || !ok {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	session.SetLoginStatus(c, true, isadmin)
	c.JSON(200, jsonutil.Success())

}
