package tmpl

import (
	"html/template"
)

func init() {
	CompileTemplates()
}

var (
	// Public templates
	Index         *template.Template
	Login         *template.Template
	Signup        *template.Template
	SignupSuccess *template.Template
	Error         *template.Template

	// Private templates
	ErrorLogged *template.Template
	Home        *template.Template
	DataHub     *template.Template
	Dashboard   *template.Template
	Settings    *template.Template

	// Slides templates
	SlidesMenu      *template.Template
	SlideHelloWorld *template.Template
	SlideTest *template.Template

	// Demos templates
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
