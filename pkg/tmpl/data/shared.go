package data

import "github.com/medtune/beta-platform/pkg"

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

type genData struct {
	Version   string
	PageTitle string
	Code      int
	Message   string
	Errors    string
	Details   string
	Samples   []Image
}

func Gen() *genData {
	return &genData{
		Version:   pkg.VERSION,
		PageTitle: "GENERATED TEMPLATE",
		Code:      666,
		Message:   "Error GEN TMPL",
		Errors:    "NaN",
		Samples: []Image{
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
			{"/static/images/inception.jpg", "inception", "", ""},
		},
	}
}

func Null() *Main { return &Main{} }
