package api

import (
	"context"
	"time"

	"github.com/medtune/beta-platform/pkg/service/dashboard"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
)

// CapsulGlobalBenchmarks .
func CapsulGlobalBenchmarks(c *gin.Context) {

}

// CapsulGlobalTests .
func CapsulGlobalTests(c *gin.Context) {

}

// CapsulGlobalHealthChecks .
func CapsulGlobalHealthChecks(c *gin.Context) {
	caps, err := dashboard.GetCapsulesList()
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
	}

	lc := make([]*jsonutil.CapsulHealthCheckResponse, 0, 20)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	for _, c := range caps {
		resp, err := capsul.GetStatus(ctx,
			&jsonutil.GetStatus{
				ModelID: c.Name,
			})
		if err != nil {
			lc = append(lc, &jsonutil.CapsulHealthCheckResponse{
				Healthy:   false,
				Status:    err.Error(),
				Version:   0,
				RoundTrip: resp.RoundTrip,
			})
		} else {
			lc = append(lc, &jsonutil.CapsulHealthCheckResponse{
				Healthy:   true,
				Status:    resp.Status,
				Version:   resp.Version,
				RoundTrip: resp.RoundTrip,
			})
		}
	}

	c.JSON(200, jsonutil.SuccessData(&jsonutil.CapsulGlobalHealthCheckResponse{
		RoundTrip: time.Since(start),
		Report:    lc,
	}))
}
