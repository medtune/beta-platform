package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/handlers/api"
	"github.com/medtune/beta-platform/server/handlers/hidden"
	"github.com/medtune/beta-platform/server/handlers/public"
	"github.com/medtune/beta-platform/server/middleware"
)

var Default *server

type server struct {
	engine *gin.Engine
	port   string
}

func (s *server) Run() {
	s.engine.Run(s.port)
}

func New(static string, port int) *server {
	engine := gin.New()
	var sport = ":" + strconv.Itoa(port)
	engine.Static(static, static)
	assembleHandlers(engine)
	return &server{
		engine: engine,
		port:   sport,
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
		PUBLIC.GET("/", public.Index)
		PUBLIC.GET("/index", public.Index)
		PUBLIC.GET("/login", public.Login)
		PUBLIC.POST("/login", public.Login)
		PUBLIC.GET("/signup", public.Signup)
		PUBLIC.GET("/signup/success", public.SignupSuccess)
		PUBLIC.POST("/signup", public.Signup)
	}

	// Login protected routes
	PROTECTED := g.Group("/")
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

	DEMOS := g.Group("/demos")
	{
		DEMOS.GET("/image_class", hidden.ImageClassification)
		DEMOS.GET("/linear_regression", hidden.PolynomialRegression)
	}
}
