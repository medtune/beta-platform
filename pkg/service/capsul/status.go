package capsul

import (
	"context"
	"fmt"
	"time"

	pb "tensorflow_serving/apis"

	"github.com/asaskevich/govalidator"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
)

// GetStatus .
func GetStatus(r *jsonutil.GetStatus) (*jsonutil.StatusResult, error) {
	ok, err := govalidator.ValidateStruct(r)
	if err != nil || !ok {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var status *pb.GetModelStatusResponse
	var rt time.Duration

	start := time.Now()
	switch r.ModelID {
	case "mnist":
		resp, err := MnistClient.Status(ctx, pbreq.Status(stdimpl.Mnist))
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	case "inception":
		resp, err := InceptionClient.Status(ctx, pbreq.Status(stdimpl.Inception))
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	case "mura-mn-v2":
		resp, err := MuraMNV2Client.Status(ctx, pbreq.Status(stdimpl.MuraMNV2))
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	case "mura-irn-v2":
		resp, err := MuraIRNV2Client.Status(ctx, pbreq.Status(stdimpl.MuraIRNV2))
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	case "mura-mn-v2-cam":
		resp, err := MuraMNV2CamClient.Status(ctx, nil)
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	default:
		return nil, fmt.Errorf("unkown model id: %s", r.ModelID)
	}

	return &jsonutil.StatusResult{
		Status:    pb.ModelVersionStatus_State_name[int32(status.GetModelVersionStatus()[0].State)],
		Version:   status.GetModelVersionStatus()[0].Version,
		RoundTrip: rt,
	}, nil
}

/*

// RunMuraCAM .
func RunMuraCAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	ok, err := govalidator.ValidateStruct(camData)
	if err != nil || !ok {
		return nil, err
	}

	var resp *pb.GetModelStatusResponse
	var roundTrip time.Duration

	if camData.ModelID == "mura-mn-v2" {
		// make inception restnet v2 inference
		request := pbreq.Status(stdimpl.MuraIRNV2)
		start := time.Now()

		response, err := MuraMNV2Client.Status(ctx, request)
		if err != nil {
			return nil, err
		}

		roundTrip = time.Since(start)
		resp = response

	} else if camData.ModelID == "mura-irn-v2" {

	}

	response := jsonutil.StatusResult{}
	response.ModelID = camData.ModelID
	response.RoundTrip = roundTrip
	info := resp.GetModelVersionStatus()[0]
	response.Version = info.Version
	response.Status = info.State

	return nil, nil
}
*/
