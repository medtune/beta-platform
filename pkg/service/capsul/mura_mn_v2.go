package capsul

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

var MuraMNV2Client *tfsclient.Client

// MuraMNV2CamClient .
var MuraMNV2CamClient *tfsclient.RestClient

// RunMuraMNV2Inference .
func RunMuraMNV2Inference(ctx context.Context, infData *jsonutil.RunImageInference, mtype *string) (*jsonutil.InferenceResult, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	request := pbreq.Predict(stdimpl.MuraMNV2, b)

	resp, err := MuraMNV2Client.Predict(ctx, request)
	if err != nil {
		return nil, err
	}

	// construct responses
	result := jsonutil.InferenceResult{}
	result.Scores = resp.Outputs["scores"].FloatVal

	var s []string
	for _, e := range resp.Outputs["classes"].StringVal {
		s = append(s, string(e))
	}
	result.Keys = s

	return &result, nil
}

// RunMuraMNV2CAM .
func RunMuraMNV2CAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	return nil, nil
}
