package tmpl

import (
	"html/template"
	"log"

	"github.com/medtune/beta-platform/pkg/tmpl/private"
	"github.com/medtune/beta-platform/pkg/tmpl/private/demos"
	"github.com/medtune/beta-platform/pkg/tmpl/public"
	"github.com/medtune/beta-platform/pkg/tmpl/shared"
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

	//Demos
	DemosMenu                *template.Template
	DemoPolynomialRegression *template.Template
	DemoMnist                *template.Template
	DemoInceptionImagenet    *template.Template
)

func init() {

	var err error

	{
		index := template.New("base")
		index, err = index.Parse(shared.Base)
		must(err)
		index, err = index.Parse(shared.HeaderPublic)
		must(err)
		index, err = index.Parse(shared.SourceHeaderIndex)
		must(err)
		index, err = index.Parse(shared.Footer)
		must(err)
		index, err = index.Parse(public.Index)
		must(err)
		Index = index
	}

	{
		login := template.New("base")
		login, err = login.Parse(shared.Base)
		must(err)
		login, err = login.Parse(shared.HeaderPublic)
		must(err)
		login, err = login.Parse(shared.SourceHeaderLogin)
		must(err)
		login, err = login.Parse(shared.Footer)
		must(err)
		login, err = login.Parse(public.Login)
		Login = login
	}

	{
		signup := template.New("base")
		signup, err = signup.Parse(shared.Base)
		must(err)
		signup, err = signup.Parse(shared.HeaderPublic)
		must(err)
		signup, err = signup.Parse(shared.SourceHeaderSignup)
		must(err)
		signup, err = signup.Parse(shared.Footer)
		must(err)
		signup, err = signup.Parse(public.Signup)
		must(err)
		Signup = signup
	}

	{
		signupSuccess := template.New("base")
		signupSuccess, err = signupSuccess.Parse(shared.Base)
		must(err)
		signupSuccess, err = signupSuccess.Parse(shared.HeaderLogged)
		must(err)
		signupSuccess, err = signupSuccess.Parse(shared.SourceHeaderEmpty)
		must(err)
		signupSuccess, err = signupSuccess.Parse(shared.Footer)
		must(err)
		signupSuccess, err = signupSuccess.Parse(public.SignupSucces)
		must(err)
		SignupSuccess = signupSuccess
	}

	{
		error_ := template.New("base")
		error_, err = error_.Parse(shared.Base)
		must(err)
		error_, err = error_.Parse(shared.HeaderPublic)
		must(err)
		error_, err = error_.Parse(shared.SourceHeaderEmpty)
		must(err)
		error_, err = error_.Parse(shared.Footer)
		must(err)
		error_, err = error_.Parse(public.Error)
		must(err)
		Error = error_
	}

	{
		error_logged := template.New("base")
		error_logged, err = error_logged.Parse(shared.Base)
		must(err)
		error_logged, err = error_logged.Parse(shared.HeaderLogged)
		must(err)
		error_logged, err = error_logged.Parse(shared.SourceHeaderEmpty)
		must(err)
		error_logged, err = error_logged.Parse(shared.Footer)
		must(err)
		error_logged, err = error_logged.Parse(public.Error)
		must(err)
		ErrorLogged = error_logged
	}

	{
		home := template.New("base")
		home, err = home.Parse(shared.Base)
		must(err)
		home, err = home.Parse(shared.HeaderLogged)
		must(err)
		home, err = home.Parse(shared.SourceHeaderIndex)
		must(err)
		home, err = home.Parse(shared.Footer)
		must(err)
		home, err = home.Parse(private.Home)
		must(err)
		Home = home
	}

	{
		demosMenu := template.New("base")
		demosMenu, err = demosMenu.Parse(shared.Base)
		must(err)
		demosMenu, err = demosMenu.Parse(shared.HeaderLogged)
		must(err)
		demosMenu, err = demosMenu.Parse(shared.SourceHeaderIndex)
		must(err)
		demosMenu, err = demosMenu.Parse(shared.Footer)
		must(err)
		demosMenu, err = demosMenu.Parse(private.Demos)
		must(err)
		DemosMenu = demosMenu
	}

	{
		polynomialRegression := template.New("base")
		polynomialRegression, err = polynomialRegression.Parse(shared.Base)
		must(err)
		polynomialRegression, err = polynomialRegression.Parse(shared.HeaderLogged)
		must(err)
		polynomialRegression, err = polynomialRegression.Parse(shared.SourceHeaderPolyReg)
		must(err)
		polynomialRegression, err = polynomialRegression.Parse(shared.Footer)
		must(err)
		polynomialRegression, err = polynomialRegression.Parse(demos.PolynomialRegression)
		must(err)
		DemoPolynomialRegression = polynomialRegression
	}

	{
		mnist := template.New("base")
		mnist, err = mnist.Parse(shared.Base)
		must(err)
		mnist, err = mnist.Parse(shared.HeaderLogged)
		must(err)
		mnist, err = mnist.Parse(shared.SourceHeaderMnist)
		must(err)
		mnist, err = mnist.Parse(shared.Footer)
		must(err)
		mnist, err = mnist.Parse(demos.Mnist)
		must(err)
		DemoMnist = mnist
	}

	{
		inceptionImagenet := template.New("base")
		inceptionImagenet, err = inceptionImagenet.Parse(shared.Base)
		must(err)
		inceptionImagenet, err = inceptionImagenet.Parse(shared.HeaderLogged)
		must(err)
		inceptionImagenet, err = inceptionImagenet.Parse(shared.SourceHeaderInceptionImagenet)
		must(err)
		inceptionImagenet, err = inceptionImagenet.Parse(shared.Footer)
		must(err)
		inceptionImagenet, err = inceptionImagenet.Parse(demos.InceptionImagenet)
		must(err)
		DemoInceptionImagenet = inceptionImagenet
	}

	{
		datahub := template.New("base")
		datahub, err = datahub.Parse(shared.Base)
		must(err)
		datahub, err = datahub.Parse(shared.HeaderLogged)
		must(err)
		datahub, err = datahub.Parse(shared.SourceHeaderEmpty)
		must(err)
		datahub, err = datahub.Parse(shared.Footer)
		must(err)
		datahub, err = datahub.Parse(private.Datahub)
		DataHub = datahub
	}

	{
		settings := template.New("base")
		settings, err = settings.Parse(shared.Base)
		must(err)
		settings, err = settings.Parse(shared.HeaderLogged)
		must(err)
		settings, err = settings.Parse(shared.SourceHeaderEmpty)
		must(err)
		settings, err = settings.Parse(shared.Footer)
		must(err)
		settings, err = settings.Parse(private.Datahub)
		Settings = settings
	}
}

// Return map of templates. Used by cmd/gen-views
func GetTemplatesMap() map[string]*template.Template {
	m := make(map[string]*template.Template)
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
	m["datahub"] = DataHub
	m["settings"] = Settings
	return m
}

func must(err error) {
	if err != nil {
		log.Panic(err)
	}
}
