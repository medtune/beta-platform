package store

import (
	"github.com/medtune/go-utils/errors"

	"github.com/medtune/beta-platform/pkg/store/model"
)

// Interface ensuring that store engine contains
// a table synchronisation method.
type syncer interface {
	Sync() error
}

// Sync store models
func (s *Store) Sync() error {
	models := getModels()
	errs := []error{}
	for _, model := range models {
		if err := s.Sync2(model); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.NewList(errs...)
	}
	return nil
}

// Return a map of empty model instances
func getModels() map[string]interface{} {
	m := make(map[string]interface{}, 3)
	m["User"] = &model.User{}
	m["Pathology Analysis Level"] = &model.PathologyAnalysisLevel{}
	m["Analysis Specifications Pool"] = &model.SpecAnalysisPool{}
	return m
}
