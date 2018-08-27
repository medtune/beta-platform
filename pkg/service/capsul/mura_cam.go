package capsul

import (
	"context"

	"github.com/medtune/capsul/pkg/tfs-client"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// MuraCamClient .
var MuraCamClient *tfsclient.RestClient

// RunMuraCAM .
func RunMuraCAM(ctx context.Context, camData *jsonutil.RunImageCam) (*jsonutil.CamResult, error) {
	return nil, nil
}
