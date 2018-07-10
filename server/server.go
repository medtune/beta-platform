package server

import (
	"github.com/gin-gonic/gin"
	"github.com/iallabs/medtune-trials/server/api"
	"github.com/iallabs/medtune-trials/server/hidden"
	"github.com/iallabs/medtune-trials/server/middleware"
	"github.com/iallabs/medtune-trials/server/public"
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
	AssembleHandlers(engine)
	return &server{
		engine: engine,
		port:   port,
	}
}

func AssembleHandlers(g *gin.Engine) {
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
