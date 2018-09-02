package capsul

import (
	"bytes"
	"context"
	"image/png"
	"io/ioutil"

	"github.com/anthonynsimon/bild/transform"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
	"github.com/vincent-petithory/dataurl"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// MnistClient .
var MnistClient *tfsclient.Client

// RunMnistInference .
func RunMnistInference(ctx context.Context, infData *jsonutil.RunImageInference) (interface{}, error) {
	var data []byte

	// Canvas draw data
	if infData.File == "" {
		// decode base64 image
		dataURL, err := dataurl.DecodeString(infData.Image)
		if err != nil {
			return nil, err
		}

		// decode png
		img, err := png.Decode(bytes.NewReader(dataURL.Data))
		if err != nil {
			return nil, err
		}

		// resize image
		img = transform.Resize(img, 28, 28, transform.Linear)
		buf := new(bytes.Buffer)

		// encode png 'resized image'
		err = png.Encode(buf, img)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()

	} else { // inference on file
		b, err := ioutil.ReadFile(infData.File)
		if err != nil {
			return nil, err
		}
		data = b
	}

	imgfloat32, err := bytesToFloat32(data)
	if err != nil {
		return nil, err
	}

	request := pbreq.PredictF32(stdimpl.Mnist, imgfloat32)

	// Run inference on image path (chexray/images)
	resp, err := MnistClient.Predict(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.Outputs, nil
}
