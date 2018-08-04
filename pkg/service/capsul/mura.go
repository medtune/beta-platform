// +build !cicd

package capsul

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	inception "github.com/medtune/capsules/capsules/inception/v1"
)

func RunMuraInference(infData *jsonutil.RunImageInference) (interface{}, error) {
	if infData.File == "" {
		return nil, fmt.Errorf("File field is empty: Got struct %v", infData)
	}

	// Run inference on image path (mura/images)
	resp, err := inception.RunInferenceOnImagePath(infData.File)
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
