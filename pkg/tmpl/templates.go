package tmpl

import (
	"html/template"
)

var (
	// Public
	Index         *template.Template
	Login         *template.Template
	Signup        *template.Template
	SignupSuccess *template.Template
	Error         *template.Template

	// Logged
	ErrorLogged *template.Template
	Home        *template.Template
	DataHub     *template.Template
	Settings    *template.Template

	// Slides
	SlidesMenu      *template.Template
	SlideHelloWorld *template.Template

	//Demos
	DemosMenu                *template.Template
	DemoPolynomialRegression *template.Template
	DemoMnist                *template.Template
	DemoInceptionImagenet    *template.Template
	DemoMura                 *template.Template
	DemoMuraV2               *template.Template
	DemoChexray              *template.Template
	DemoChexrayV2            *template.Template
	DemoSentimentAnalysis    *template.Template
)

func init() {
	CompileTemplates()
}
