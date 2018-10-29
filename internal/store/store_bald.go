package store

import (
	"fmt"

	"github.com/medtune/beta-platform/internal/store/model"
)

type bioAnalysisStore interface {
	// Pathology spec analysis level
	CreatePathologyAL(*model.PathologyAnalysisLevel) error
	GetPathologyAL(string) (*model.PathologyAnalysisLevel, error)
	GetPathologiesAL() (*[]model.PathologyAnalysisLevel, error)

	// Specs
	CreateSpec(string, string, string, string) error
	GetSpec(string) (*model.SpecAnalysisPool, error)
	GetSpecs([]string) (*[]model.SpecAnalysisPool, error)
}

// CreatePathologyAL .
func (s *Store) CreatePathologyAL(m *model.PathologyAnalysisLevel) error {
	// Validate
	if k, err := s.Valid(m); err != nil || !k {
		return err
	}
	// insert value
	if _, err := s.Insert(m); err != nil {
		return err
	}
	return nil
}

// GetPathologyAL .
func (s *Store) GetPathologyAL(name string) (*model.PathologyAnalysisLevel, error) {
	pal := &model.PathologyAnalysisLevel{}
	has, err := s.Where("name = ?", name).Get(pal)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("record doesnt exist: %v", err)
	}

	return pal, nil
}

// GetPathologiesAL .
func (s *Store) GetPathologiesAL() (*[]model.PathologyAnalysisLevel, error) {
	var pal []model.PathologyAnalysisLevel
	err := s.Find(&pal)
	if err != nil {
		return nil, err
	}

	return &pal, nil
}

// CreateSpec .
func (s *Store) CreateSpec(name string, unit string, min string, max string) error {
	m := &model.SpecAnalysisPool{
		Name: name,
		Unit: unit,
		Min:  min,
		Max:  max,
	}
	if k, err := s.Valid(m); err != nil || !k {
		return err
	}
	if _, err := s.Insert(m); err != nil {
		return err
	}
	return nil
}

// GetSpec .
func (s *Store) GetSpec(name string) (*model.SpecAnalysisPool, error) {
	spec := &model.SpecAnalysisPool{}
	has, err := s.Where("name = ?", name).Get(spec)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("record doesnt exist: %v", err)
	}
	return spec, nil
}

// GetSpecs .
func (s *Store) GetSpecs(ignore []string) (*[]model.SpecAnalysisPool, error) {
	var specs []model.SpecAnalysisPool
	// TODO Ignore specs

	err := s.Find(&specs)
	if err != nil {
		return nil, fmt.Errorf("find specs failed: %v", err)
	}

	return &specs, nil
}
