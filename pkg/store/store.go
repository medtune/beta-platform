package store

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-xorm/xorm"
	"github.com/medtune/beta-platform/pkg/store/db"
)

var (
	Agent *store
)

type store struct {
	*xorm.Engine
	valid func(interface{}) (bool, error)
}

type Store interface {
	Connect()
	Disconnect()
}

func New(config db.ConnStr) (*store, error) {
	engine, err := db.NewPGEngine(config)
	if err != nil {
		return nil, err
	}
	s := &store{
		Engine: engine,
		valid:  govalidator.ValidateStruct,
	}
	return s, nil
}
