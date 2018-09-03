package api

import (
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
)

// MuraRunInference .
func MuraRunInference(c *gin.Context) {
	// Parse body data
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()

	infData.File = staticPath("mura", infData.File)

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
	// Parse body data
	procData := jsonutil.ProcessImage{}
	if err := c.ShouldBindJSON(&procData); err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Option data
	infData := jsonutil.RunImageInference{}
	camData := jsonutil.RunImageCam{}

	// Set inference data
	infData.ModelID = procData.ModelID
	infData.File = staticPath("mura", procData.Target)
	infData.NumPreds = procData.NumPreds

	// set cam data
	camData.ModelID = procData.ModelID
	camData.Target = procData.Target
	camData.Output = procData.Output
	camData.Force = true

	// result
	var infResult *jsonutil.InferenceResult
	var camResult *jsonutil.CamResult

	// errors
	var infError error
	var camError error

	// timings
	var Timing time.Duration

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
			infResult = infresult
			infError = inferr
			wg.Done()
		}()

		go func() {
			// run cam
			camresult, camerr := capsul.RunMuraCAM(ctx, &camData)
			// copy result
			camResult = camresult
			camError = camerr
			wg.Done()
		}()

		// wait for goroutines to finish
		wg.Wait()
		cancel()
		Timing = time.Since(start)

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

	processResult.Timing = Timing
	response.Data = processResult

	// Success
	c.JSON(200, jsonutil.SuccessData(response))
}
