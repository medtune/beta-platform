// +build cicd

package capsul

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

func RunMnistInference(infData *jsonutil.RunImageInference) (interface{}, error) {
	return nil, fmt.Errorf("CICD BUILD: +build cicd %v", infData)
}

func RunInceptionInference(infData *jsonutil.RunImageInference) (interface{}, error) {
	return nil, fmt.Errorf("CICD BUILD: +build cicd %v", infData)
}

func RunMuraInference(infData *jsonutil.RunImageInference) (interface{}, error) {
	return nil, fmt.Errorf("CICD BUILD: +build cicd %v", infData)
}

func RunChexrayInference(infData *jsonutil.RunImageInference) (interface{}, error) {
	return nil, fmt.Errorf("CICD BUILD: +build cicd %v", infData)
}
