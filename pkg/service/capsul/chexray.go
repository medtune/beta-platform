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
	// ChexrayMNV2Client .
	ChexrayMNV2Client *tfsclient.Client
	// ChexrayMNV2CamClient .
	ChexrayMNV2CamClient *tfsclient.RestClient
)

// RunChexrayInference .
func RunChexrayInference(ctx context.Context, infData *jsonutil.RunImageInference) (*jsonutil.InferenceResult, error) {
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

	if infData.ModelID == "chexray-mn-v2" {
		request := pbreq.Predict(stdimpl.ChexrayMNV2, b)
		start := time.Now()

		response, err := ChexrayMNV2Client.Predict(ctx, request)
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
	result.Keys = []string{"positive", "negative"}
	result.RoundTrip = roundTrip

	return &result, nil
}

// RunChexrayCAM .
func RunChexrayCAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	ok, err := govalidator.ValidateStruct(camData)
	if err != nil || !ok {
		return nil, err
	}

	if camData.ModelID == "chexray-mn-v2" {
		start := time.Now()
		r, err := ChexrayMNV2CamClient.Cam(context.Background(), &tfsclient.CamRequest{
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

	} else if camData.ModelID == "chexray-irn-v2" {
		return nil, fmt.Errorf("unavailable model: %s", camData.ModelID)
	}
	return nil, fmt.Errorf("unknown model: %s", camData.ModelID)
}
