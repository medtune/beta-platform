package server

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/debug"
	"github.com/medtune/beta-platform/server/public"
)

func Debug(static, port string) *server {
	engine := gin.New()
	engine.Static(static, static)
	assembleHandlersDebug(engine)
	return &server{
		engine: engine,
		port:   port,
	}
}

func assembleHandlersDebug(g *gin.Engine) {
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
