package tmpl

import (
	"html/template"
	"log"

	"github.com/medtune/beta-platform/internal/tmpl/private"
	"github.com/medtune/beta-platform/internal/tmpl/private/demos"
	"github.com/medtune/beta-platform/internal/tmpl/private/slides"
	"github.com/medtune/beta-platform/internal/tmpl/public"
	"github.com/medtune/beta-platform/internal/tmpl/shared"
)

// CompileTemplates will construct known templates.
// by default it will panic when error is occured
func CompileTemplates() {
	// Just helping garbage collector
	var err error

	// Index page
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

	// login page
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

	// Signup page
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

	// Signup success page
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

	// Error public page
	{
		errorPublic := template.New("base")
		errorPublic, err = errorPublic.Parse(shared.Base)
		must(err)
		errorPublic, err = errorPublic.Parse(shared.HeaderPublic)
		must(err)
		errorPublic, err = errorPublic.Parse(shared.SourceHeaderEmpty)
		must(err)
		errorPublic, err = errorPublic.Parse(shared.Footer)
		must(err)
		errorPublic, err = errorPublic.Parse(public.Error)
		must(err)
		Error = errorPublic
	}

	// Error private page
	{
		errorLogged := template.New("base")
		errorLogged, err = errorLogged.Parse(shared.Base)
		must(err)
		errorLogged, err = errorLogged.Parse(shared.HeaderLogged)
		must(err)
		errorLogged, err = errorLogged.Parse(shared.SourceHeaderEmpty)
		must(err)
		errorLogged, err = errorLogged.Parse(shared.Footer)
		must(err)
		errorLogged, err = errorLogged.Parse(public.Error)
		must(err)
		ErrorLogged = errorLogged
	}

	// Home page
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

	// Demos menu page
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

	// Polynomial regression page
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

	// Mnist page
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

	// Inception imagenet page
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

	// Mura imagenet
	{
		mura := template.New("base")
		mura, err = mura.Parse(shared.Base)
		must(err)
		mura, err = mura.Parse(shared.HeaderLogged)
		must(err)
		mura, err = mura.Parse(shared.SourceHeaderMura)
		must(err)
		mura, err = mura.Parse(shared.Footer)
		must(err)
		mura, err = mura.Parse(demos.Mura)
		must(err)
		DemoMura = mura
	}

	// Mura imagenet
	{
		muraV2 := template.New("base")
		muraV2, err = muraV2.Parse(shared.Base)
		must(err)
		muraV2, err = muraV2.Parse(shared.HeaderLogged)
		must(err)
		muraV2, err = muraV2.Parse(shared.SourceHeaderMuraV2)
		must(err)
		muraV2, err = muraV2.Parse(shared.Footer)
		must(err)
		muraV2, err = muraV2.Parse(demos.MuraV2)
		must(err)
		DemoMuraV2 = muraV2
	}

	// Chexray page
	{
		chexray := template.New("base")
		chexray, err = chexray.Parse(shared.Base)
		must(err)
		chexray, err = chexray.Parse(shared.HeaderLogged)
		must(err)
		chexray, err = chexray.Parse(shared.SourceHeaderChexray)
		must(err)
		chexray, err = chexray.Parse(shared.Footer)
		must(err)
		chexray, err = chexray.Parse(demos.Chexray)
		must(err)
		DemoChexray = chexray
	}

	// Chexray V2 page
	{
		chexrayV2 := template.New("base")
		chexrayV2, err = chexrayV2.Parse(shared.Base)
		must(err)
		chexrayV2, err = chexrayV2.Parse(shared.HeaderLogged)
		must(err)
		chexrayV2, err = chexrayV2.Parse(shared.SourceHeaderChexrayV2)
		must(err)
		chexrayV2, err = chexrayV2.Parse(shared.Footer)
		must(err)
		chexrayV2, err = chexrayV2.Parse(demos.ChexrayV2)
		must(err)
		DemoChexrayV2 = chexrayV2
	}

	// Sentiment analysis page
	{
		sentimentAnalysis := template.New("base")
		sentimentAnalysis, err = sentimentAnalysis.Parse(shared.Base)
		must(err)
		sentimentAnalysis, err = sentimentAnalysis.Parse(shared.HeaderLogged)
		must(err)
		sentimentAnalysis, err = sentimentAnalysis.Parse(shared.SourceHeaderSentimentAnalysis)
		must(err)
		sentimentAnalysis, err = sentimentAnalysis.Parse(shared.Footer)
		must(err)
		sentimentAnalysis, err = sentimentAnalysis.Parse(demos.SentimentAnalysis)
		must(err)
		DemoSentimentAnalysis = sentimentAnalysis
	}

	// Slides menu page
	{
		slidesMenu := template.New("base")
		slidesMenu, err = slidesMenu.Parse(shared.Base)
		must(err)
		slidesMenu, err = slidesMenu.Parse(shared.HeaderLogged)
		must(err)
		slidesMenu, err = slidesMenu.Parse(shared.SourceHeaderIndex)
		must(err)
		slidesMenu, err = slidesMenu.Parse(shared.Footer)
		must(err)
		slidesMenu, err = slidesMenu.Parse(private.Slides)
		must(err)
		SlidesMenu = slidesMenu
	}

	// slide Hello world
	{
		slideMedtunePresentation := template.New("base")
		slideMedtunePresentation, err = slideMedtunePresentation.Parse(slides.Base)
		must(err)
		slideMedtunePresentation, err = slideMedtunePresentation.Parse(slides.MedtunePresentation)
		must(err)
		SlideMedtunePresentation = slideMedtunePresentation
	}

	// Datahub page
	{
		datahub := template.New("base")
		datahub, err = datahub.Parse(shared.Base)
		must(err)
		datahub, err = datahub.Parse(shared.HeaderLogged)
		must(err)
		datahub, err = datahub.Parse(shared.SourceHeaderDatahub)
		must(err)
		datahub, err = datahub.Parse(shared.Footer)
		must(err)
		datahub, err = datahub.Parse(private.Datahub)
		DataHub = datahub
	}

	// Settings page
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

	// Dashboard menu page
	{
		dashboard := template.New("base")
		dashboard, err = dashboard.Parse(shared.BaseDashboard)
		must(err)
		dashboard, err = dashboard.Parse(private.Dashboard)
		must(err)
		Dashboard = dashboard
	}
}

var must = func(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// SetErrorHandler sets the 'must' function
func SetErrorHandler(f func(err error)) {
	must = f
}
