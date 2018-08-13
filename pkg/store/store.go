package store

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/medtune/beta-platform/pkg/store/model"
)

// Compile static check
var _ userStore = &Store{}
var _ bioAnalysisStore = &Store{}

var (
	// Agent main object used by other packages
	Agent *Store
)

// New return a database engine
func New(config db.ConnStr) (*Store, error) {
	engine, err := db.NewPGEngine(config)
	if err != nil {
		return nil, err
	}
	s := &Store{
		Engine: engine,
		Valid:  govalidator.ValidateStruct,
	}
	return s, nil
}

// Type store is the abstraction behind data interactions
// Database io & validation
type Store struct {
	*xorm.Engine
	Valid func(interface{}) (bool, error)
}

// Sync store models
func (s *Store) Sync() error {
	if err := s.Sync2(&model.User{}); err != nil {
		return err
	} else if err := s.Sync2(&model.PathologyAnalysisLevel{}); err != nil {
		return err
	} else if err := s.Sync2(&model.SpecAnalysisPool{}); err != nil {
		return err
	} else {
		fmt.Printf("migrated: %s\n", "model.{User|PathologyAnalysisLevel|SpecAnalysisPool}")
	}
	return nil
}
