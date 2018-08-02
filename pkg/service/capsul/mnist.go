// +build !cicd

package capsul

import (
	"bytes"
	"context"
	"image/png"

	"github.com/anthonynsimon/bild/transform"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	mnist "github.com/medtune/capsules/capsules/mnist/v1"
	"github.com/vincent-petithory/dataurl"
)

func RunMnistInference(infData *jsonutil.RunImageInference) (interface{}, error) {
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

		sends3 := buf.Bytes()
		req, err := mnist.PredictRequestFromBytes(sends3)
		if err != nil {
			return nil, err
		}
		resp, err := mnist.RunInference(context.Background(), req)
		if err != nil {
			return nil, err
		}
		return resp.Outputs, nil

	}
	resp, err := mnist.RunInferenceOnImagePath("./static/demos/mnist/images/" + infData.File)
	if err != nil {
		return nil, err
	}

	return resp.Outputs, nil
}
