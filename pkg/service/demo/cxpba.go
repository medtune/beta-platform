package demo

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/model"
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
func GetCXPBAForPathology(pathology string) (*model.PathologyAnalysisLevel, error) {
	var ok bool
	for _, p := range chexrayPredictionClasses {
		if p == pathology {
			ok = true
			break
		}
	}
	if !ok {
		return nil, fmt.Errorf("unsupported pathology: %s", pathology)
	}
	result, err := store.Agent.GetPathologyAL(pathology)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSpecPoolGrid .
func GetSpecPoolGrid() (*[]model.SpecAnalysisPool, error) {
	return store.Agent.GetSpecs()
}
