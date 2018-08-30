package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/session"
)

// MuraRunInference .
func MuraRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

	// Parse body data
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()

	// Run inference
	result, err := capsul.RunMuraInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(result))
}

// MuraRunCam .
func MuraRunCam(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

	// Parse data from body
	camData := jsonutil.RunImageCam{}
	err := c.ShouldBindJSON(&camData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()
	// Run cam
	resp, err := capsul.RunMuraCAM(ctx, &camData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// TODO cam file creation
	// should be used as static

	// Success
	c.JSON(200, jsonutil.SuccessData(resp))
}

// MuraProcess .
func MuraProcess(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

	// Parse body data
	procData := jsonutil.ProcessImage{}
	if err := c.ShouldBindJSON(&procData); err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Option data
	infData := jsonutil.RunImageInference{}
	camData := jsonutil.RunImageCam{}

	// result
	var infResult *jsonutil.InferenceResult
	var camResult *jsonutil.CamResult

	// errors
	var infError error
	var camError error

	// timings
	var infTiming time.Duration
	var camTiming time.Duration

	{
		// WaitGroup
		// Will block until wg.Count = 0
		var wg sync.WaitGroup
		// Init with count = 2 (2 goroutines)
		wg.Add(2)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		start := time.Now()

		go func() {
			// run inference
			infresult, inferr := capsul.RunMuraInference(ctx, &infData)
			// copy result
			infTiming = time.Since(start)
			infResult = infresult
			infError = inferr
			wg.Done()
		}()

		go func() {
			// run cam
			camresult, camerr := capsul.RunMuraCAM(ctx, &camData)
			// copy result
			camTiming = time.Since(start)
			camResult = camresult
			camError = camerr
			wg.Done()
		}()

		// wait for goroutines to finish
		wg.Wait()
		cancel()
	}

	// response objects
	response := &jsonutil.DefaultResponse{}
	processResult := &jsonutil.ProcessResult{}
	// init success
	response.Success = true

	// inf check
	if infError != nil {
		response.Success = false
		response.Errors = append(response.Errors, infError.Error())
	} else {
		processResult.Inference = infResult
	}

	// cam check
	if camError != nil {
		response.Success = false
		response.Errors = append(response.Errors, camError.Error())
	} else {
		processResult.Cam = camResult
	}

	response.Data = processResult

	// Success
	c.JSON(200, jsonutil.SuccessData(response))
}
