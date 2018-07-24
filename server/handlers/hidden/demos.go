package hidden

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/service/global"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func DemosMenu(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemosMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demonstrations",
	})
}

func PolynomialRegression(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoPolynomialRegression.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Polynomial Regression",
	})
}

func Mnist(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoMnist.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Handwritten digits classification",
	})
}

func InceptionImagenet(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	images, err := global.CollectImagesData("inception")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	fmt.Println(images)

	c.Status(200)
	tmpl.DemoInceptionImagenet.Execute(c.Writer, &data.InceptionDemo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: images,
	})
}

func Mura(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoMura.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: MURA Classification",
	})
}

func Chexray(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}
