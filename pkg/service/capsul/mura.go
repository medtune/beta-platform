package capsul

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	pb "tensorflow_serving/apis"

	"github.com/asaskevich/govalidator"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
)

var (
	// MuraIRNV2Client .
	MuraIRNV2Client *tfsclient.Client
	// MuraIRNV2CamClient .
	MuraIRNV2CamClient *tfsclient.RestClient
	// MuraMNV2Client .
	MuraMNV2Client *tfsclient.Client
	// MuraMNV2CamClient .
	MuraMNV2CamClient *tfsclient.RestClient
	// MuraPredictionClasses .
	muraPredictionClasses = []string{"positive", "negative"}
)

// RunMuraInference .
func RunMuraInference(ctx context.Context, infData *jsonutil.RunImageInference) (*jsonutil.InferenceResult, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	if infData.ModelID == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	var resp *pb.PredictResponse
	var roundTrip time.Duration

	if infData.ModelID == "mura-mn-v2" {
		request := pbreq.Predict(stdimpl.MuraMNV2, b)
		start := time.Now()

		response, err := MuraMNV2Client.Predict(ctx, request)
		if err != nil {
			return nil, err
		}

		roundTrip = time.Since(start)
		resp = response

	} else if infData.ModelID == "mura-irn-v2" {
		// make inception restnet v2 inference
		request := pbreq.Predict(stdimpl.MuraIRNV2, b)
		start := time.Now()

		response, err := MuraIRNV2Client.Predict(ctx, request)
		if err != nil {
			return nil, err
		}

		roundTrip = time.Since(start)
		resp = response
	} else {
		return nil, fmt.Errorf("unkown model id: %s", infData.ModelID)
	}

	// construct responses
	result := jsonutil.InferenceResult{}
	result.Scores = resp.Outputs["scores"].FloatVal
	result.Keys = muraPredictionClasses
	result.RoundTrip = roundTrip

	return &result, nil
}

// RunMuraCAM .
func RunMuraCAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	ok, err := govalidator.ValidateStruct(camData)
	if err != nil || !ok {
		return nil, err
	}

	if camData.ModelID == "mura-mn-v2" {
		// make inception restnet v2 cam
		start := time.Now()
		r, err := MuraMNV2CamClient.Cam(context.Background(), &tfsclient.CamRequest{
			Target: camData.Target,
			Dest:   camData.Output,
			Force:  camData.Force,
		})

		if err != nil {
			return nil, err
		}

		roundTrip := time.Since(start)
		if !r.Success {
			return nil, fmt.Errorf("err: target:%s dest:%s errors: %v", r.Target, r.Dest, r.Errors)
		}

		return &jsonutil.CamResult{
			ModelID:   camData.ModelID,
			Output:    r.Dest,
			RoundTrip: roundTrip,
		}, nil

	} else if camData.ModelID == "mura-irn-v2" {
		return nil, fmt.Errorf("unavailable cam model: %s", camData.ModelID)
	}
	return nil, fmt.Errorf("unknown model: %s", camData.ModelID)
}
