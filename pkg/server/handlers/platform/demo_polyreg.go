package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// PolynomialRegression demo
func PolynomialRegression(c *gin.Context) {
	c.Status(200)
	tmpl.DemoPolynomialRegression.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Polynomial Regression",
	})
}
