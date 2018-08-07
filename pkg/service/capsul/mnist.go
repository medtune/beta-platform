// +build !cicd

package capsul

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"io/ioutil"

	"github.com/anthonynsimon/bild/transform"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/request/mnist"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
	"github.com/vincent-petithory/dataurl"
	"gocv.io/x/gocv"
)

var MnistClient *tfsclient.Client

func RunMnistInference(ctx context.Context, infData *jsonutil.RunImageInference) (interface{}, error) {
	var data []byte

	// Canvas draw data
	if infData.File == "" {
		dataURL, err := dataurl.DecodeString(infData.Image)
		if err != nil {
			return nil, err
		}

		img, err := png.Decode(bytes.NewReader(dataURL.Data))
		if err != nil {
			return nil, err
		}

		img = transform.Resize(img, 28, 28, transform.Linear)
		buf := new(bytes.Buffer)

		err = png.Encode(buf, img)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()

	} else { // inference on file
		b, err := ioutil.ReadFile("./static/demos/mnist/images/" + infData.File)
		if err != nil {
			return nil, err
		}
		data = b
	}

	imgfloat32, err := bytesToFloat32(data)
	if err != nil {
		return nil, err
	}

	resp, err := MnistClient.Predict(ctx, mnist.Default(imgfloat32))
	if err != nil {
		return nil, err
	}

	return resp.Outputs, nil
}

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
