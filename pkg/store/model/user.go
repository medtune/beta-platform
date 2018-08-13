package model

type User struct {
	Id            int64  //`xorm:"pk autoincr"`
	Username      string `xorm:"varchar(16) not null unique" valid:"alphanum,required"`
	Email         string `xorm:"varchar(64) not null unique" valid:"email,required"`
	Password      string `xorm:"varchar(128) not null" valid:"required"`
	AccountType   string `xorm:"varchar(16) not null" valid:"alpha,required"`
	AccountStatus bool   `xorm:"varchar(128) not null" valid:"required"`
	AccountLevel  int8   `xorm:"default 5"`
}
