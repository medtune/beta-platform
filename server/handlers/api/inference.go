// +build !cicd

package api

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/service/demos"
	"github.com/medtune/beta-platform/pkg/session"
)

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

	// Run inference
	resp, err := capsul.RunMnistInference(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(resp))
}

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

	// Run inference
	result, err := capsul.RunInceptionInference(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(result))
}

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

	err = demos.DropImage(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.Success())
}

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

	var response = struct {
		Prediction string `json:"prediction"`
		Correct    bool   `json:"correct"`
	}{}

	switch infData.Id {
	case 0:
		response.Prediction = "Positive"
		response.Correct = true
	case 1:
		response.Prediction = "Negative"
		response.Correct = true
	case 2:
		response.Prediction = "Positive"
		response.Correct = false
	case 3:
		response.Prediction = "Negative"
		response.Correct = false

	default:
		response.Prediction = "Undefined"
		response.Correct = false
	}

	time.Sleep(time.Duration(rand.Intn(184)+1073) * time.Millisecond)

	c.JSON(200, jsonutil.SuccessData(response))
}

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

	var response = struct {
		Prediction string `json:"prediction"`
		Correct    bool   `json:"correct"`
	}{}

	switch infData.Id {
	case 0:
		response.Prediction = "Positive"
		response.Correct = true
	case 1:
		response.Prediction = "Negative"
		response.Correct = false
	case 2:
		response.Prediction = "CLASS4"
		response.Correct = false
	case 3:
		response.Prediction = "CLASS4"
		response.Correct = false
	default:
		response.Prediction = "Class----"
		response.Correct = false
	}

	time.Sleep(time.Duration(rand.Intn(300)+950) * time.Millisecond)

	c.JSON(200, jsonutil.SuccessData(response))
}
