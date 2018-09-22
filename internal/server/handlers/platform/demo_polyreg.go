package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// PolynomialRegression demo
func PolynomialRegression(c *gin.Context) {
	c.Status(200)
	tmpl.DemoPolynomialRegression.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Polynomial Regression",
	})
}
