// +build !cicd

package capsul

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/request/mura"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

var MuraClient *tfsclient.Client

func RunMuraInference(ctx context.Context, infData *jsonutil.RunImageInference) (interface{}, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("File field is empty: Got struct %v", infData)
	}

	// Read file
	b, err := ioutil.ReadFile(infData.File)
	if err != nil {
		panic(err)
	}

	resp, err := MuraClient.Predict(ctx, mura.Default(b))
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
