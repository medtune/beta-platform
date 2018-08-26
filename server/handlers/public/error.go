package public

import (
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Error page for public and logged views
func Error(c *gin.Context) {
	errStatus := c.Param("code")
	code, err := strconv.Atoi(errStatus)
	var respTmpl *template.Template
	if session.GetLoginStatus(c) {
		respTmpl = tmpl.ErrorLogged
	} else {
		respTmpl = tmpl.Error
	}
	if err != nil {
		c.Status(418)
		respTmpl.Execute(c.Writer, data.ErrorImAteaPot)
	} else if code == 404 {
		c.Status(code)
		respTmpl.Execute(c.Writer, data.Error404)
	} else if code == 401 {
		c.Status(401)
		respTmpl.Execute(c.Writer, &data.Error401)
	} else if code == 500 {
		c.Status(500)
		respTmpl.Execute(c.Writer, data.Error500)
	} else {
		c.Status(418)
		respTmpl.Execute(c.Writer, data.ErrorFinalBoss)
	}
}
