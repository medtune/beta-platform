package shared

var HeaderPublic = `{{define "header"}}
<div class="mdl-layout--fixed-header mdl-layout mdl-js-layout">
	<header class="mdl-layout__header">
		<div class="mdl-layout__header-row">
			<span class="mdl-layout-title">MedTune</span>

			<div class="mdl-layout-spacer"></div>

			<!-- Navigation -->
			<nav class="mdl-navigation">
				<a class="mdl-navigation__link" href="/signup">Sign up</a>
				<a class="mdl-navigation__link" href="/login">Login</a>
			</nav>

		</div>

	</header>
	<div class="mdl-layout__drawer">
		<span class="version">v 0.0.2</span>

		<div style="text-align:center">
			<img src="/static/images/logo.png" class="demo-avatar" onclick="window.location.href='/index'">
		<div>

		<nav class="mdl-navigation">
			<div class="drawer-separator"></div>
			<a class="mdl-navigation__link drawer-nav-link" href="/tryouts">About</a>
			<div class="drawer-separator"></div>
			<a class="mdl-navigation__link drawer-nav-link" href="/uploads">Newsletter</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/uploads">Beta-Test</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/tryouts">Contact</a>
			<div class="drawer-separator"></div>
		</nav>
	</div>
</div>
{{end}}`
