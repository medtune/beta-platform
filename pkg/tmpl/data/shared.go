package data

var (
	Default = &defaultHolder{}
)

type Main struct {
	Version   string
	PageTitle string
}

type defaultHolder struct {
	Main
}

func Null() *Main { return &Main{} }
