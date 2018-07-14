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

var (
	Agent *Store
)

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

type Store struct {
	*xorm.Engine
	Valid func(interface{}) (bool, error)
}

func (s *Store) Sync() error {
	if err := s.Sync2(&model.User{}); err != nil {
		return err
	} else {
		fmt.Printf("migrated: %s\n", "model.User")
	}
	return nil
}
