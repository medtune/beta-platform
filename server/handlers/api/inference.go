package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/service/demos"
	"github.com/medtune/beta-platform/pkg/session"
)

// MnistRunInference .
func MnistRunInference(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()
	// Run inference
	resp, err := capsul.RunMnistInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(resp))
}

// InceptionImagenetRunInference .
func InceptionImagenetRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()
	// Run inference
	result, err := capsul.RunInceptionInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(result))
}

// InceptionImagenetDropImage .
func InceptionImagenetDropImage(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// drop image
	err = demos.DropImage(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.Success())
}

// MuraRunInference .
func MuraRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
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

// ChexrayRunInference .
func ChexrayRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()

	// Run inference
	result, err := capsul.RunChexrayInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(result))

}
