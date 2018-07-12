package model

type User struct {
	Id            int64
	Username      string
	Password      string
	AccountType   string
	AccountStatus string
	OwnData       string
}

type DemoData struct {
	Id      int64
	Name    string
	OwnerId int64
	DemoId  int64
	Type    string

	Bytes string
	Link  string
	//Hash    int64
}

type Demo struct {
	Id          int64
	Name        string
	Status      string
	Version     string
	Description string
	DescImage   string
}
