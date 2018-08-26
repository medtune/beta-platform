package debug

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Home debug view
func Home(c *gin.Context) {
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Home",
	}
	c.Status(200)
	tmpl.Home.Execute(c.Writer, &inject)
}

// Index debug view
func Index(c *gin.Context) {
	c.Status(200)
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Index",
	}
	tmpl.Index.Execute(c.Writer, &inject)
}

// Login debug view
func Login(c *gin.Context) {
	c.Status(200)
	tmpl.Login.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "login",
	})
}

// Signup debug view
func Signup(c *gin.Context) {
	c.Status(200)
	tmpl.Signup.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Signup",
	})
}

// Error debug view
func Error(c *gin.Context) {
	errStatus := c.Param("code")
	code, err := strconv.Atoi(errStatus)
	if err != nil {
		c.Status(418)
		tmpl.Error.Execute(c.Writer, &data.ErrorImAteaPot)
	} else if code == 404 {
		c.Status(code)
		tmpl.Error.Execute(c.Writer, &data.Error404)
	} else if code == 401 {
		c.Status(401)
		tmpl.Error.Execute(c.Writer, &data.Error401)
	} else if code == 500 {
		c.Status(500)
		tmpl.Error.Execute(c.Writer, &data.Error500)
	} else {
		c.Status(418)
		tmpl.Error.Execute(c.Writer, &data.ErrorFinalBoss)
	}
}

// ErrorLogged debug view
func ErrorLogged(c *gin.Context) {
	errStatus := c.Param("code")
	code, err := strconv.Atoi(errStatus)
	if err != nil {
		c.Status(418)
		tmpl.ErrorLogged.Execute(c.Writer, &data.ErrorImAteaPot)
	} else if code == 404 {
		c.Status(code)
		tmpl.ErrorLogged.Execute(c.Writer, &data.Error404)
	} else if code == 401 {
		c.Status(401)
		tmpl.ErrorLogged.Execute(c.Writer, &data.Error401)
	} else if code == 500 {
		c.Status(500)
		tmpl.ErrorLogged.Execute(c.Writer, &data.Error500)
	} else {
		c.Status(418)
		tmpl.ErrorLogged.Execute(c.Writer, &data.ErrorFinalBoss)
	}
}

// DemosMenu debug view
func DemosMenu(c *gin.Context) {
	c.Status(200)
	tmpl.DemosMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demos menu",
	})
}

// PolynomialRegression debug view
func PolynomialRegression(c *gin.Context) {
	c.Status(200)
	tmpl.DemoPolynomialRegression.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Polynomial Regression",
	})
}

// Mnist debug view
func Mnist(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMnist.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Image classification",
	})
}

// InceptionImagenet debug view
func InceptionImagenet(c *gin.Context) {
	c.Status(200)
	tmpl.DemoInceptionImagenet.Execute(c.Writer, &data.InceptionDemo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: data.Gen().Samples,
	},
	)
}

// Mura debug view
func Mura(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMura.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: MURA Classification",
	})
}

// MuraV2 debug view
func MuraV2(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMuraV2.Execute(c.Writer, &data.MuraV2Demo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: data.Gen().Samples,
	},
	)
}

// Chexray debug view
func Chexray(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}

// ChexrayV2 debug view
func ChexrayV2(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexrayV2.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo V2: Chest X-Ray Classification",
	})
}

// SentimentAnalysis debug view
func SentimentAnalysis(c *gin.Context) {
	c.Status(200)
	tmpl.DemoSentimentAnalysis.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Sentiment Analysis",
	})
}

// Datahub debug view
func Datahub(c *gin.Context) {
	c.Status(200)
	tmpl.DataHub.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Datahub",
	})
}

// SlidesMenu debug view
func SlidesMenu(c *gin.Context) {
	c.Status(200)
	tmpl.SlidesMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Slides menu",
	})
}
