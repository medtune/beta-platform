package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iallabs/medtune-trials/server"
)

func main() {
	gin.SetMode(gin.DebugMode)
	Server := server.New("./static", ":8007")
	Server.Run()
}
