// +build !cicd

package capsul

import (
	"fmt"

	"gocv.io/x/gocv"
)

// This file is compiled only with build tags doesnt contain cicd #see utils_cicd.go

// Convert images bytes to float32 array
// Need gocv
func bytesToFloat32(ib []byte) ([]float32, error) {
	matb, err := gocv.IMDecode(ib, -1)
	if err != nil {
		return nil, fmt.Errorf("not an image %v", err)
	}

	mat := gocv.NewMat()
	matb.ConvertTo(&mat, gocv.MatTypeCV32F)

	imgfloat := make([]float32, 0, 0)

	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			imgfloat = append(imgfloat, mat.GetFloatAt(i, j))
		}
	}
	return imgfloat, nil
}
