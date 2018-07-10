package shared

const (
	SourceHeaderEmpty = `{{define "srcheader"}}{{end}}`

	SourceHeaderIndex = `{{define "srcheader"}}
		<script src="/static/js/typed.min.js"></script>
	{{end}}`

	SourceHeaderApiHelpers = `{{define "srcheader"}}
		<script src="/static/js/api.js"></script>
	{{end}}`
)
