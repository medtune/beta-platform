package shared

const (
	SourceHeaderEmpty = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/required.js"></script>
{{end}}`

	SourceHeaderIndex = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/required.js"></script>
		<script src="/static/js/typed.min.js"></script>
	{{end}}`

	SourceHeaderSignup = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/common.js"></script>
		<script src="/static/js/required.js"></script>
		<script src="/static/js/signup.js"></script>
{{end}}`

	SourceHeaderLogin = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/common.js"></script>
		<script src="/static/js/required.js"></script>
		<script src="/static/js/login.js"></script>
{{end}}`

	SourceHeaderTFJS = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

{{end}}`

	SourceHeaderAPIS = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
{{end}}`
)
