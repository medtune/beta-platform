package tmpl

import "html/template"

// GetTemplatesMap Return map of templates. Used by cmd/gen-views
func GetTemplatesMap() map[string]*template.Template {
	m := make(map[string]*template.Template, 30)
	// Public
	m["index"] = Index
	m["home"] = Home
	m["signup"] = Signup
	m["signup-success"] = SignupSuccess
	m["login"] = Login
	m["error"] = Error
	m["error-logged"] = ErrorLogged
	m["demos-menu"] = DemosMenu
	m["demo-polynomial-regression"] = DemoPolynomialRegression
	m["demo-inception-imagenet"] = DemoInceptionImagenet
	m["demo-mnist"] = DemoMnist
	m["demo-mura"] = DemoMura
	m["demo-mura-v2"] = DemoMuraV2
	m["demo-chexray"] = DemoChexray
	m["demo-chexray-v2"] = DemoChexrayV2
	m["demo-sentiment-analysis"] = DemoSentimentAnalysis
	m["slides-menu"] = SlidesMenu
	m["slide-hello-world"] = SlideHelloWorld
	m["datahub"] = DataHub
	m["settings"] = Settings
	return m
}
