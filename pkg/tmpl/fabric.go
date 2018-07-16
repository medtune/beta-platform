package tmpl

import (
	"html/template"
	"log"

	"github.com/medtune/beta-platform/pkg/tmpl/public"
	"github.com/medtune/beta-platform/pkg/tmpl/shared"
	"github.com/medtune/beta-platform/pkg/tmpl/trials"
)

var (
	Index         *template.Template
	Home          *template.Template
	Login         *template.Template
	Signup        *template.Template
	SignupSuccess *template.Template
	Inception     *template.Template
	Error         *template.Template
	ErrorLogged   *template.Template
	Uploads       *template.Template
	Capsules      *template.Template
	Settings      *template.Template
)

func init() {
	index := template.New("base")

	{
		index, err := index.Parse(shared.Base)
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

	home := template.New("base")
	{
		home, err := home.Parse(shared.Base)
		must(err)
		home, err = home.Parse(shared.HeaderLogged)
		must(err)
		home, err = home.Parse(shared.SourceHeaderIndex)
		must(err)
		home, err = home.Parse(shared.Footer)
		must(err)
		home, err = home.Parse(public.Index)
		must(err)
		Home = home
	}

	login := template.New("base")
	{
		login, err := login.Parse(shared.Base)
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

	signup := template.New("base")
	{

		signup, err := signup.Parse(shared.Base)
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

	signupSuccess := template.New("base")
	{

		signupSuccess, err := signupSuccess.Parse(shared.Base)
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

	error_ := template.New("base")
	{
		error_, err := error_.Parse(shared.Base)
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

	error_logged := template.New("base")
	{
		error_logged, err := error_logged.Parse(shared.Base)
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

	inception := template.New("base")
	{
		inception, err := inception.Parse(shared.Base)
		must(err)
		inception, err = inception.Parse(shared.HeaderPublic)
		must(err)
		inception, err = inception.Parse(shared.SourceHeaderEmpty)
		must(err)
		inception, err = inception.Parse(shared.Footer)
		must(err)
		inception, err = inception.Parse(trials.Inception)
		must(err)
		Inception = inception
	}
}

func GetTemplatesMap() map[string]*template.Template {
	m := make(map[string]*template.Template)
	m["index"] = Index
	m["home"] = Home
	m["signup"] = Signup
	m["signup-success"] = SignupSuccess
	m["login"] = Login
	m["error"] = Error
	m["error-logged"] = ErrorLogged
	return m
}

func must(err error) {
	if err != nil {
		log.Panic(err)
	}
}
