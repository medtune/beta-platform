package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/server"
)

func main() {

	gin.SetMode(gin.DebugMode)
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		Server := server.Debug("./static", ":8007")
		Server.Run()
	} else {
		Server := server.New("./static", ":8007")
		Server.Run()
	}

}
