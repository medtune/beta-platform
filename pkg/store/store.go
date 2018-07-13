package store

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"
	"github.com/medtune/beta-platform/pkg/store/db"
)

// Compile static check

var _ userStore = &Store{}
var _ store = &Store{}

var (
	Agent *store
)

func New(config db.ConnStr) (*Store, error) {
	engine, err := db.NewPGEngine(config)
	if err != nil {
		return nil, err
	}
	s := &Store{
		Engine: engine,
		valid:  govalidator.ValidateStruct,
	}
	return s, nil
}+:

// Store
type store interface {
	Sync()
	Connect()
	Abort()
}

type Store struct {
	*xorm.Engine
	valid func(interface{}) (bool, error)
}

func (s *Store) Sync() {

}

func (s *Store) Connect() {

}

func (s *Store) Abort() {

}
