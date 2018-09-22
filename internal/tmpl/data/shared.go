package data

import "github.com/medtune/beta-platform/internal"

var (
	// Default .
	Default = &defaultHolder{}
)

// Main .
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

// Gen erate data
// Used for inception in debug mode
func Gen() *genData {
	return &genData{
		Version:   pkg.VERSION,
		PageTitle: "GENERATED TEMPLATE",
		Code:      666,
		Message:   "Error GEN TMPL",
		Errors:    "NaN",
		Samples: []Image{
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
			{"debug.png", "debug-name", "", ""},
		},
	}
}
