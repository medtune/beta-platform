package public

const Error = `{{define "content"}}
	<div class="center-screen">
		error {{ .Code }}
		{{ .Message }}
		{{ .Details }}
	</div>
{{end}}`
