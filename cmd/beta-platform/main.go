package main

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server"
)

func main() {
	gin.SetMode(gin.DebugMode)
	Server := server.New("./static", ":8007")
	Server.Run()

}
