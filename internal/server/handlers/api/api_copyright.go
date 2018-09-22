package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/jsonutil"
)

// Copyright description handler
func Copyright(c *gin.Context) {
	c.JSON(200, jsonutil.SuccessData(pkg.GetCopyright()))
}
