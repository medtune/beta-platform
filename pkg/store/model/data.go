package model

// GlobalData .
type GlobalData struct {
	Id   int64
	Name int64
}

// DemoData .
type DemoData struct {
	Id   int64
	Name string `xorm:"varchar(64) not null"`

	Visibility string `xorm:"not null default 'basic'"`
	OwnerId    int64  `xorm:"not null"`
	DemoId     string `xorm:"not null"`
	Type       string `xorm:"varchar(16) not null"`
	Kind       string `xorm:"varchar(16) not null"`

	Bytes string `xorm:"varchar(2097152) not null" valid:""`
	Link  string `xorm:"varchar(256) not null" valid:""`
	//TODO hash data
}

// Demo .
type Demo struct {
	Id          int64
	Name        string `xorm:"varchar(64) not null" valid:"required"`
	Status      bool
	Version     string `xorm:"varchar(25) not null" valid:"required"`
	Description string `xorm:"varchar(1024)"`
	DescImage   string `xorm:"varchar(128)" valid:"required"`
}
