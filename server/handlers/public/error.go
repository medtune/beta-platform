package public

import (
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Error(c *gin.Context) {
	errStatus := c.Param("code")
	code, err := strconv.Atoi(errStatus)
	var response *template.Template
	if session.GetLoginStatus(c) {
		response = tmpl.ErrorLogged
	} else {
		response = tmpl.Error
	}
	if err != nil {
		c.Status(418)
		response.Execute(c.Writer, data.ErrorImAteaPot)
	} else if code == 404 {
		c.Status(code)
		response.Execute(c.Writer, data.Error404)
	} else if code == 401 {
		c.Status(401)
		response.Execute(c.Writer, &data.Error401)
	} else if code == 500 {
		c.Status(500)
		response.Execute(c.Writer, data.Error500)
	} else {
		c.Status(418)
		response.Execute(c.Writer, data.ErrorFinalBoss)
	}
}
