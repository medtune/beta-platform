package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server/handlers/api"
	"github.com/medtune/beta-platform/server/handlers/platform"
	"github.com/medtune/beta-platform/server/handlers/public"
	"github.com/medtune/beta-platform/server/middleware"
)

// New .
func New(static string, port int) Engine {
	server := gin.New()
	server.Static(static, static)

	set404Handler(server, func(c *gin.Context) {
		c.Redirect(302, "/error/404")
	})
	setMiddlewares(server)
	assembleHandlers(server)

	return &engine{
		server,
		fmt.Sprintf(":%d", port),
	}
}

func setMiddlewares(server *gin.Engine) {
	// Set gin middlewares
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	server.Use(middleware.Session())
}

func set404Handler(server *gin.Engine, nrh gin.HandlerFunc) {
	server.NoRoute(nrh)
}

func setStaticFiles(server *gin.Engine, staticDir string, root string) {
	server.Static(staticDir, root)
}

// Make production engine server
func assembleHandlers(server *gin.Engine) {

	// Public routes handlers
	PUBLIC := server.Group("/")
	PUBLIC.GET("/", public.Index)
	PUBLIC.GET("/index", public.Index)
	PUBLIC.GET("/login", public.Login)
	PUBLIC.POST("/login", public.Login)
	PUBLIC.GET("/signup", public.Signup)
	PUBLIC.POST("/signup", public.Signup)
	PUBLIC.GET("/signup/success", public.SignupSuccess)

	// Errors
	ERRORS := PUBLIC.Group("/error")
	ERRORS.GET("/:code", public.Error)

	// Login platform routes
	PROTECTED := server.Group("/")
	PROTECTED.GET("/logout", platform.Logout)
	PROTECTED.GET("/home", platform.Home)
	PROTECTED.GET("/demos", platform.DemosMenu)
	PROTECTED.GET("/slides", platform.SlidesMenu)
	PROTECTED.GET("/datahub", platform.Datahub)
	PROTECTED.POST("/datahub_upload", platform.DatahubUpload)

	// Demonstrations routes
	DEMOS := PROTECTED.Group("/demos")
	DEMOS.GET("/polynomial_regression", platform.PolynomialRegression)
	DEMOS.GET("/mnist", platform.Mnist)
	DEMOS.GET("/inception_imagenet", platform.InceptionImagenet)
	DEMOS.GET("/mura", platform.Mura)
	DEMOS.GET("/mura.v2", platform.MuraV2)
	DEMOS.GET("/chexray", platform.Chexray)
	DEMOS.GET("/chexray.v2", platform.ChexrayV2)
	DEMOS.GET("/sentiment_analysis", platform.SentimentAnalysis)

	// Api routes
	// public API
	PUBLICAPI := PUBLIC.Group("/api")
	PUBLICAPI.GET("/version", api.Version)
	PUBLICAPI.POST("/test", api.Test)
	PUBLICAPI.POST("/login", api.Login)
	PUBLICAPI.POST("/signup", api.Signup)

	API := PROTECTED.Group("/api")
	API.POST("/custom/exec", api.CustomExecution)

	APIDEMOS := API.Group("/demos")
	APIDEMOS.POST("/mnist/run_inference", api.MnistRunInference)
	APIDEMOS.POST("/inception_imagenet/run_inference", api.InceptionImagenetRunInference)
	APIDEMOS.POST("/mura/process", api.MuraProcess)
	APIDEMOS.POST("/mura/run_cam", api.MuraRunCam)
	APIDEMOS.POST("/mura/run_inference", api.MuraRunInference)
	APIDEMOS.POST("/chexray/run_inference", api.ChexrayRunInference)

	DATAHUB := API.Group("/datahub")
	DATAHUB.POST("/upload/:demo", api.DemoDataUpload)
	DATAHUB.POST("/drop/:demo", api.DemoDataDrop)
}
