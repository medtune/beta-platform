package capsul

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// ChexrayClient .
var ChexrayClient *tfsclient.Client

// RunChexrayInference .
func RunChexrayInference(ctx context.Context, infData *jsonutil.RunImageInference) (*jsonutil.InferenceResult, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	request := pbreq.Predict(stdimpl.Chexray, b)

	// Run inference on image path (chexray/images)
	resp, err := ChexrayClient.Predict(ctx, request)
	if err != nil {
		return nil, err
	}

	// construct response
	result := jsonutil.InferenceResult{}
	result.Scores = resp.Outputs["scores"].FloatVal

	var s []string
	for _, e := range resp.Outputs["classes"].StringVal {
		s = append(s, string(e))
	}
	result.Keys = s

	return &result, nil
}
