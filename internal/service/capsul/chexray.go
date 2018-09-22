package capsul

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	pb "tensorflow_serving/apis"

	"github.com/asaskevich/govalidator"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
)

var (
	// ChexrayDN121Client .
	ChexrayDN121Client *tfsclient.Client
	// ChexrayDN121CamClient .
	ChexrayDN121CamClient *tfsclient.RestClient
	// ChexrayPPHelper .
	ChexrayPPHelper *tfsclient.RestClient
	// ChexrayMNV2Client .
	ChexrayMNV2Client *tfsclient.Client
	// ChexrayMNV2CamClient .
	ChexrayMNV2CamClient *tfsclient.RestClient
	// ChexrayPredictionClasses .
	chexrayPredictionClasses = []string{
		"Atelectasis",
		"Cardiomegaly",
		"Effusion",
		"Infiltration",
		"Mass",
		"Nodule",
		"Pneumonia",
		"Pneumothorax",
		"Consolidation",
		"Edema",
		"Emphysema",
		"Fibrosis",
		"Pleural_Thickening",
		"Hernia",
	}
)

// RunChexrayInference .
func RunChexrayInference(ctx context.Context, infData *jsonutil.RunImageInference) (*jsonutil.InferenceResult, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	if infData.ModelID == "" {
		return nil, fmt.Errorf("model id field is empty: got struct %v", infData)
	}

	var resp *pb.PredictResponse
	var roundTrip time.Duration

	if infData.ModelID == "chexray-mn-v2" {
		// Read file
		b, err := ioutil.ReadFile(infData.File)
		if err != nil {
			return nil, err
		}

		request := pbreq.Predict(stdimpl.ChexrayMNV2, b)
		start := time.Now()

		response, err := ChexrayMNV2Client.Predict(ctx, request)
		if err != nil {
			return nil, err
		}

		roundTrip = time.Since(start)
		resp = response

	} else if infData.ModelID == "chexray-dn-121" {

		// Preprocess Image request
		pprequest, err := ChexrayPPHelper.Process(ctx, &tfsclient.ProcessRequest{
			Target: filepath.Base(infData.File),
		})
		if err != nil {
			return nil, err
		}

		// validate success
		if !pprequest.Success {
			return nil, fmt.Errorf("chexray preprocessing request failed: %v:%v", pprequest, err)
		}

		// convert string to [][][][][]float32
		var data [][][][][]float32
		err = json.Unmarshal([]byte(pprequest.Out), &data)
		if err != nil {
			return nil, fmt.Errorf("preprocessing output unmarshing failed: %v", err)
		}

		// init looping over data
		var dataList []float32
		d := data[0][0]

		// loop over data 224*224*3 =~ 150000
		for _, i := range d {
			for _, j := range i {
				dataList = append(dataList, j[0], j[1], j[2])
			}
		}

		// prepare request
		request := pbreq.PredictFTest(stdimpl.ChexrayDN121, dataList)
		start := time.Now()

		// send prediction request
		response, err := ChexrayDN121Client.Predict(ctx, request)
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
	result.Scores = resp.GetOutputs()["scores"].GetFloatVal()
	result.Keys = chexrayPredictionClasses
	result.RoundTrip = roundTrip

	return &result, nil
}

// RunChexrayCAM .
func RunChexrayCAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	ok, err := govalidator.ValidateStruct(camData)
	if err != nil || !ok {
		return nil, err
	}

	if camData.ModelID == "chexray-mn-v2-cam" {
		start := time.Now()
		r, err := ChexrayMNV2CamClient.Cam(ctx, &tfsclient.CamRequest{
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

	} else if camData.ModelID == "chexray-dn-121-cam" {
		return nil, fmt.Errorf("unavailable custom capsule model: %s", camData.ModelID)
	}
	return nil, fmt.Errorf("unknown model: %s", camData.ModelID)
}
