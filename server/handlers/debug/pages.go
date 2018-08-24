package debug

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Home(c *gin.Context) {
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Home",
	}
	c.Status(200)
	tmpl.Home.Execute(c.Writer, &inject)
}

func Index(c *gin.Context) {
	c.Status(200)
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Index",
	}
	tmpl.Index.Execute(c.Writer, &inject)
}

func Login(c *gin.Context) {
	c.Status(200)
	tmpl.Login.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "login",
	})
}

func Signup(c *gin.Context) {
	c.Status(200)
	tmpl.Signup.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Signup",
	})
}

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

func DemosMenu(c *gin.Context) {
	c.Status(200)
	tmpl.DemosMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demos menu",
	})
}

func PolynomialRegression(c *gin.Context) {
	c.Status(200)
	tmpl.DemoPolynomialRegression.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Polynomial Regression",
	})
}

func Mnist(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMnist.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Image classification",
	})
}

func InceptionImagenet(c *gin.Context) {
	c.Status(200)
	tmpl.DemoInceptionImagenet.Execute(c.Writer, &data.InceptionDemo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: []data.Image{
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
		},
	},
	)
}

func Mura(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMura.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: MURA Classification",
	})
}

func MuraV2(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMuraV2.Execute(c.Writer, &data.MuraV2Demo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: []data.Image{
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
			{"/static/demos/mura/images/image_2.png", "image_2", "", ""},
		},
	},
	)
}

func Chexray(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}

func ChexrayV2(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexrayV2.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo V2: Chest X-Ray Classification",
	})
}

func SentimentAnalysis(c *gin.Context) {
	c.Status(200)
	tmpl.DemoSentimentAnalysis.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Sentiment Analysis",
	})
}

func Datahub(c *gin.Context) {
	c.Status(200)
	tmpl.DataHub.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Datahub",
	})
}

func SlidesMenu(c *gin.Context) {
	c.Status(200)
	tmpl.SlidesMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Slides menu",
	})
}
