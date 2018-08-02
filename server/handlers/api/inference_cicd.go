// +build cicd

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
)

func MnistRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:cicd:")))
		return
	}
	c.JSON(200, struct{ Mock bool }{true})
}

func InceptionImagenetRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:cicd")))
		return
	}
	c.JSON(200, struct{ Mock bool }{true})
}

func InceptionImagenetDropImage(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:cicd")))
		return
	}
	c.JSON(200, struct{ Mock bool }{true})
}

func MuraRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:cicd")))
		return
	}
	c.JSON(200, struct{ Mock bool }{true})
}

func ChexrayRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:cicd")))
		return
	}
	c.JSON(200, struct{ Mock bool }{true})
}
