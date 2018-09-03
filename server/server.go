package server

import "github.com/gin-gonic/gin"

// Wrapping gin server implementation
type engine struct {
	engine *gin.Engine
	port   string
}

// Engine interface is described by the Run method
type Engine interface {
	Run() error
}

// Run server
func (s *engine) Run() error {
	return s.engine.Run(s.port)
}
