package api

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/custom"
)

// CustomExecution .
func CustomExecution(c *gin.Context) {
	start := time.Now()

	// Parse data from body
	cx := jsonutil.CustomExecutionRequest{}
	err := c.ShouldBindJSON(&cx)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// response object
	resp := &jsonutil.CustomExecutionResponse{}

	ok, err := govalidator.ValidateStruct(cx)
	if err != nil || !ok {
		resp.Warnings = append(resp.Warnings, fmt.Sprintf("non valid request: %v ,error: %v", ok, err))
	}

	// time execution
	startX := time.Now()
	jobs, err := custom.Execute(&cx)

	resp.ExecutionRoundTrip = time.Since(startX)

	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	resp.Jobs = jobs
	resp.RoundTrip = time.Since(start)
	c.JSON(200, jsonutil.SuccessData(resp))
}
