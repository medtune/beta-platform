package store

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"
	"github.com/medtune/beta-platform/pkg/store/db"
)

// Compile time check
var _ userStore = &Store{}
var _ bioAnalysisStore = &Store{}
var _ syncer = &Store{}

var (
	// Agent main object used by other packages
	Agent *Store
)

// Store type is the abstraction behind data interactions
// Database io & validation
type Store struct {
	*xorm.Engine
	Valid func(interface{}) (bool, error)
}

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
