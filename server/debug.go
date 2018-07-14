package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/handlers/debug"
	"github.com/medtune/beta-platform/server/handlers/public"
)

var ()

func Debug(static string, port int) *server {
	engine := gin.New()
	engine.Static(static, static)
	debugHandlers(engine)
	var sport = ":" + strconv.Itoa(port)
	return &server{
		engine: engine,
		port:   sport,
	}
}

func debugHandlers(g *gin.Engine) {
	// Gin recovery and logger
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	// No route set
	g.NoRoute(public.NoRouteProxy)

	DEBUG := g.Group("/")
	{
		DEBUG.GET("/index", debug.Index)
		DEBUG.GET("/home", debug.Home)
		DEBUG.GET("/login", debug.Login)
		DEBUG.GET("/signup", debug.Signup)
		DEBUG.GET("/error/:code", debug.Error)
	}
}
