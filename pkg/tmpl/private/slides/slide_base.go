package slides

const Base = `{{define "base"}}
<!doctype html>
<html lang="en">

	<head>
		<meta charset="utf-8">

		<title> {{ .Title }} </title>
		
		<meta name="description" content="">
		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:regular,bold,italic,thin,light,bolditalic,black,medium&amp;lang=en">
    	<link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-blue.min.css" />
    	<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
		<link rel="stylesheet" href="/static/reveal.js-3.7.0/css/reveal.css">
		<link rel="stylesheet" href="/static/reveal.js-3.7.0/css/theme/white.css" id="theme">
		<link rel="stylesheet" href="/static/reveal.js-3.7.0/lib/css/zenburn.css">

		<script>
			var link = document.createElement( 'link' );
			link.rel = 'stylesheet';
			link.type = 'text/css';
			link.href = window.location.search.match( /print-pdf/gi ) ? '/static/reveal.js-3.7.0/css/print/pdf.css' : '/static/reveal.js-3.7.0/css/print/paper.css';
			document.getElementsByTagName( 'head' )[0].appendChild( link );
		</script>
	</head>

	<body>
		<div class="reveal">
			{{template "slide" .}}
		</div>

		<script src="/static/reveal.js-3.7.0/lib/js/head.min.js"></script>
		<script src="/static/reveal.js-3.7.0/js/reveal.js"></script>

		<script>

			Reveal.initialize({
				controls: true,
				progress: true,
				history: true,
				center: true,
				transition: 'slide',
				dependencies: [
					{ src: '/static/reveal.js-3.7.0/lib/js/classList.js', condition: function() { return !document.body.classList; } },
					{ src: '/static/reveal.js-3.7.0/plugin/markdown/marked.js', condition: function() { return !!document.querySelector( '[data-markdown]' ); } },
					{ src: '/static/reveal.js-3.7.0/plugin/markdown/markdown.js', condition: function() { return !!document.querySelector( '[data-markdown]' ); } },
					{ src: '/static/reveal.js-3.7.0/plugin/highlight/highlight.js', async: true, callback: function() { hljs.initHighlightingOnLoad(); } },
					{ src: '/static/reveal.js-3.7.0/plugin/search/search.js', async: true },
					{ src: '/static/reveal.js-3.7.0/plugin/zoom-js/zoom.js', async: true },
					{ src: '/static/reveal.js-3.7.0/plugin/notes/notes.js', async: true }
				]
			});

		</script>

	</body>
</html>
{{end}}`
