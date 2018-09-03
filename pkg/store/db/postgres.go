package db

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
	// Load postgres lib
	_ "github.com/lib/pq"
)

const (
	// postgres constant
	postgres      = "postgres"
	sslverifyFull = "verify-full"
	ssldisable    = "disable"
)

// ConnStr contain necessary information to establish
// secure and non-secure connections with a postgres db
type ConnStr struct {
	Host         string
	User         string
	Password     string
	Database     string
	Port         int
	SslMode      int8
	MaxIdleConns int
	MaxOpenConns int
}

// MakeConnectionString return a string that express
// postgres connection string. Something on form of
// dbname=test host=localhost user=admin password=acunt
func MakeConnectionString(c ConnStr) string {
	var ssl string
	if c.SslMode == 2 {
		ssl = sslverifyFull
	} else if c.SslMode == 0 {
		ssl = ssldisable
	}
	str := fmt.Sprintf(
		"dbname=%v sslmode=%v host=%v user=%v password=%v",
		c.Database,
		ssl,
		c.Host,
		c.User,
		c.Password,
	)
	if c.Port != 0 {
		str += " port=" + strconv.Itoa(c.Port)
	}
	return str
}

// NewPGEngine return a new postgres engine from
// Configuration string.
func NewPGEngine(c ConnStr) (*xorm.Engine, error) {
	cstr := MakeConnectionString(c)
	engine, err := xorm.NewEngine(postgres, cstr)
	engine.SetMaxOpenConns(c.MaxOpenConns)
	engine.SetMaxIdleConns(c.MaxIdleConns)
	return engine, err
}
