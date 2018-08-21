package capsul

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/request/inception"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

// InceptionClient .
var InceptionClient *tfsclient.Client

// RunInceptionInference .
func RunInceptionInference(ctx context.Context, infData *jsonutil.RunImageInference) (interface{}, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("file field is empty: got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	// Send prediction request
	resp, err := InceptionClient.Predict(ctx, inception.Default(b))
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

	return result, nil
}
