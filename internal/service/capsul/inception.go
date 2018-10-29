package capsul

import (
	"context"
	"fmt"
	"io/ioutil"

	tfsclient "github.com/medtune/capsul/pkg/tfs-client"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
)

// InceptionClient .
var InceptionClient *tfsclient.Client

// RunInceptionInference .
func RunInceptionInference(ctx context.Context, infData *jsonutil.RunImageInference) (*jsonutil.InferenceResult, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	request := pbreq.Predict(stdimpl.Inception, b)

	// Run inference on image path (chexray/images)
	resp, err := InceptionClient.Predict(ctx, request)
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
