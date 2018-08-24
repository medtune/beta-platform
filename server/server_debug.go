package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/handlers/debug"
)

// Debug return a ui oriented debug version of the plaform server
// it will serve static content only
// pkg session / db / cache are inexistent in the debug mode
func Debug(static string, port int) Engine {
	server := gin.New()
	server.Static(static, static)
	debugHandlers(server)
	var sport = ":" + strconv.Itoa(port)
	return &engine{
		engine: server,
		port:   sport,
	}
}

// Assemble static ui
func debugHandlers(g *gin.Engine) {
	// Gin recovery and logger
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	DEBUG := g.Group("/")
	{
		DEBUG.GET("/", debug.Index)
		DEBUG.GET("/index", debug.Index)
		DEBUG.GET("/home", debug.Home)
		DEBUG.GET("/login", debug.Login)
		DEBUG.GET("/signup", debug.Signup)
		DEBUG.GET("/error/:code", debug.Error)
		DEBUG.GET("/demos", debug.DemosMenu)
		DEBUG.GET("/demos/mnist", debug.Mnist)
		DEBUG.GET("/demos/inception_imagenet", debug.InceptionImagenet)
		DEBUG.GET("/demos/polynomial_regression", debug.PolynomialRegression)
		DEBUG.GET("/demos/mura", debug.Mura)
		DEBUG.GET("/demos/mura.v2", debug.MuraV2)
		DEBUG.GET("/demos/chexray", debug.Chexray)
		DEBUG.GET("/demos/chexray.v2", debug.ChexrayV2)
		DEBUG.GET("/demos/sentiment_analysis", debug.SentimentAnalysis)
		DEBUG.GET("/datahub", debug.Datahub)
		DEBUG.GET("/slides", debug.SlidesMenu)
	}
}
