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

type DemoData struct {
	Id         int64
	Name       string `xorm:"varchar(64) not null"`
	Visibility string `xorm:"not null default 'basic'"`
	OwnerId    int64  `xorm:"not null"`
	DemoId     int64  `xorm:"not null"`
	Type       string `xorm:"varchar(16) not null"`

	Bytes string `xorm:"varchar(2097152) not null" valid:""`
	Link  string `xorm:"varchar(256) not null" valid:""`
	//TODO hash data
}

type Demo struct {
	Id          int64
	Name        string `xorm:"varchar(64) not null" valid:"required"`
	Status      bool
	Version     string `xorm:"varchar(25) not null" valid:"required"`
	Description string `xorm:"varchar(1024)"`
	DescImage   string `xorm:"varchar(128)" valid:"required"`
}
