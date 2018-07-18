package data

type Main struct {
	Version   string
	PageTitle string
}

func Default(args ...string) *Main {
	return &Main{}
}

func Null() *Main { return &Main{} }
