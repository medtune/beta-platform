package server

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/api"
	"github.com/medtune/beta-platform/server/hidden"
	"github.com/medtune/beta-platform/server/middleware"
	"github.com/medtune/beta-platform/server/public"
)

var Default *server

type server struct {
	engine *gin.Engine
	port   string
}

func (s *server) Run() {
	s.engine.Run(s.port)
}

func New(static string, port string) *server {
	engine := gin.New()
	engine.Static(static, static)
	assembleHandlers(engine)
	return &server{
		engine: engine,
		port:   port,
	}
}

func assembleHandlers(g *gin.Engine) {
	// Gin recovery and logger
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	// No route set
	g.NoRoute(public.NoRouteProxy)
	g.Use(middleware.Session())

	// Public routes
	PUBLIC := g.Group("/")
	{
		PUBLIC.GET("/index", public.Index)
		PUBLIC.GET("/login", public.Login)
		PUBLIC.POST("/login", public.Login)
		PUBLIC.GET("/signup", public.Signup)
		PUBLIC.POST("/signup", public.Signup)
	}

	PROTECTED := g.Group("/")
	PROTECTED.Use(middleware.LoggedOnly())
	{
		PROTECTED.GET("/logout", hidden.Logout)
		PROTECTED.GET("/home", hidden.Home)
	}

	// Api routes
	API := g.Group("/api")
	{
		API.GET("/capsule", api.Capsule)
	}

	ERRORS := g.Group("/error")
	{
		ERRORS.GET("/:code", public.Error)
	}
}
