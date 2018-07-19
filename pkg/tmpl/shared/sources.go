package shared

const (
	SourceHeaderEmpty = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/required.js"></script>
{{end}}`

	SourceHeaderIndex = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="/static/js/required.js"></script>
		<script src="/static/js/third_party/typed.min.js"></script>
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

	SourceHeaderPolyReg = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.6.0/p5.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.6.0/addons/p5.dom.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/@tensorflow/tfjs@0.11.1"> </script>
		<script src="/static/js/common.js"></script>
		<script src="/static/js/required.js"></script>
		<link rel="stylesheet" href="/static/demos/polyreg/styles.css" />
		<script src="/static/demos/polyreg/main.js"></script>
	
{{end}}`

	SourceHeaderImgClass = `{{define "srcheader"}}
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
{{end}}`
)
