package capsul

import (
	"context"
	"fmt"
	"time"

	pb "tensorflow_serving/apis"

	"github.com/asaskevich/govalidator"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
)

// GetStatus .
func GetStatus(ctx context.Context, r *jsonutil.GetStatus) (*jsonutil.StatusResult, error) {
	ok, err := govalidator.ValidateStruct(r)
	if err != nil || !ok {
		return nil, err
	}

	var status *pb.GetModelStatusResponse
	var rt time.Duration
	start := time.Now()

	switch r.ModelID {
	case "mnist-lenet":
		resp, err := MnistClient.Status(ctx, pbreq.Status(stdimpl.Mnist))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "imagenet-inception-v3":
		resp, err := InceptionClient.Status(ctx, pbreq.Status(stdimpl.Inception))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "mura-mobilenet-v2":
		resp, err := MuraMNV2Client.Status(ctx, pbreq.Status(stdimpl.MuraMNV2))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "mura-inception-resnet-v2":
		resp, err := MuraIRNV2Client.Status(ctx, pbreq.Status(stdimpl.MuraIRNV2))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "mura-mobilenet-v2-cam":
		resp, err := MuraMNV2CamClient.Status(ctx, nil)
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "chexray-mobilenet-v2":
		resp, err := ChexrayMNV2Client.Status(ctx, pbreq.Status(stdimpl.MuraMNV2))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "chexray-mobilenet-v2-cam":
		resp, err := ChexrayMNV2CamClient.Status(ctx, nil)
		if err != nil {
			return nil, err
		}
		status = resp
		rt = time.Since(start)

	case "chexray-densenet-121":
		resp, err := ChexrayDN121Client.Status(ctx, pbreq.Status(stdimpl.ChexrayDN121))
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	case "chexray-pp-helper":
		resp, err := ChexrayPPHelper.Status(ctx, nil)
		if err != nil {
			return nil, err
		}
		rt = time.Since(start)
		status = resp

	default:
		return nil, fmt.Errorf("unknown model id: %s", r.ModelID)
	}

	return &jsonutil.StatusResult{
		ModelID:   r.ModelID,
		Status:    pb.ModelVersionStatus_State_name[int32(status.GetModelVersionStatus()[0].State)],
		Version:   status.GetModelVersionStatus()[0].Version,
		RoundTrip: rt,
	}, nil
}
