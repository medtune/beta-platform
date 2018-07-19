package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/handlers/api"
	"github.com/medtune/beta-platform/server/handlers/hidden"
	"github.com/medtune/beta-platform/server/handlers/public"
	"github.com/medtune/beta-platform/server/middleware"
)

type server struct {
	engine *gin.Engine
	port   string
}

func (s *server) Run() {
	s.engine.Run(s.port)
}

// New .
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
	// Set gin middlewares
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.Session())

	// Handler for non set routes
	g.NoRoute(public.NoRouteProxy)

	// Public routes handlers
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

		// Demonstrations routes
		DEMOS := PROTECTED.Group("/demos")
		{
			DEMOS.GET("/", hidden.DemosMenu)
			DEMOS.GET("/image_class", hidden.ImageClassification)
			DEMOS.GET("/polynomial_regression", hidden.PolynomialRegression)
		}

		// Api routes
		API := PROTECTED.Group("/api")
		{
			API.GET("/capsule", api.Capsule)
		}
	}

	// Errors handler
	ERRORS := g.Group("/error")
	{
		ERRORS.GET("/:code", public.Error)
	}

}
