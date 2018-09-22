package demo

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/store"
)

var chexrayPredictionClasses = []string{
	"Atelectasis",
	"Cardiomegaly",
	"Effusion",
	"Infiltration",
	"Mass",
	"Nodule",
	"Pneumonia",
	"Pneumothorax",
	"Consolidation",
	"Edema",
	"Emphysema",
	"Fibrosis",
	"Pleural_Thickening",
	"Hernia",
}

// GetCXPBAForPathology .
func GetCXPBAForPathology(r *jsonutil.GetPathologyAL) (*jsonutil.GetPathologyALResponse, error) {
	ok, err := govalidator.ValidateStruct(r)
	if err != nil || !ok {
		return nil, err
	}

	var known bool
	for _, p := range chexrayPredictionClasses {
		if p == r.Pathology {
			known = true
			break
		}
	}

	if !known {
		return nil, fmt.Errorf("unsupported pathology: %s", r.Pathology)
	}

	result, err := store.Agent.GetPathologyAL(r.Pathology)
	if err != nil {
		return nil, err
	}

	response := &jsonutil.GetPathologyALResponse{
		Pathology: r.Pathology,
		Table:     result,
	}

	return response, nil
}

// GetSpecPoolGrid .
func GetSpecPoolGrid(r *jsonutil.GetSpecPoolGrid) (*jsonutil.GetSpecPoolGridResponse, error) {
	ok, err := govalidator.ValidateStruct(r)
	if err != nil || !ok {
		return nil, err
	}

	res, err := store.Agent.GetSpecs(r.Ignore)
	if err != nil {
		return nil, err
	}

	response := &jsonutil.GetSpecPoolGridResponse{
		Specs: res,
	}

	return response, nil
}
