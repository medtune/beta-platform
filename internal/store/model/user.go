package model

// User .
type User struct {
	Id            int64  `yaml:",omitempty"` //`xorm:"pk autoincr"`
	Username      string `xorm:"varchar(16) not null unique" valid:"alphanum,required" yaml:"username"`
	Email         string `xorm:"varchar(64) not null unique" valid:"email,required" yaml:"email"`
	Password      string `xorm:"varchar(128) not null" valid:"required" yaml:"password"`
	AccountType   string `xorm:"varchar(16) not null" valid:"alpha,required" yaml:"account_type"`
	AccountStatus bool   `xorm:"varchar(128) not null" valid:"required" yaml:"account_status"`
	AccountLevel  int8   `xorm:"default 5" yaml:"account_level"`
}
