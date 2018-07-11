package db

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

const (
	Postgres      = "postgres"
	sslverifyFull = "verify-full"
	ssldisable    = "disable"
)

type ConnStr struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
	sslMode  int8
}

func MakeConnectionString(c ConnStr) string {
	var ssl string
	if c.sslMode == 2 {
		ssl = sslverifyFull
	} else if c.sslMode == 0 {
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

func NewPGEngine(c ConnStr) (*xorm.Engine, error) {
	cstr := MakeConnectionString(c)
	engine, err := xorm.NewEngine(Postgres, cstr)
	return engine, err
}
