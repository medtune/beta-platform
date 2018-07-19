package shared

var HeaderPublic = `{{define "header"}}
<div class="mdl-layout--fixed-header mdl-layout mdl-js-layout">
	<header class="mdl-layout__header page-header">
		<div class="mdl-layout__header-row">
			<span class="mdl-layout-title">MedTune <span class="v-beta">beta</span></span>

			<div class="mdl-layout-spacer"></div>

			<!-- Navigation -->
			<div class="navigation-container">
				<nav class="nav-cc mdl-navigation">
					<a class="mdl-navigation__link" href="/signup">Sign up</a>
					<a class="mdl-navigation__link" href="/login">Login</a>
				</nav>
			</div>

		</div>

	</header>
	<div class="mdl-layout__drawer">
		<span class="version">{{ .Version }}</span>

		<div style="text-align:center">
			<img src="/static/images/logo.png" class="demo-avatar" onclick="window.location.href='/index'">
		<div>

		<nav class="mdl-navigation">

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">The project</span>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">About</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">Terms</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">Partners</a>

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">Account</span>
			<a class="mdl-navigation__link drawer-nav-link" href="/login">Login</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/signup">Signup</a>

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">Contact</span>
			<a class="mdl-navigation__link drawer-nav-link" href="/beta-test">Beta Test</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/contact">Contact us</a>

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">Contribute</span>
			<a class="mdl-navigation__link drawer-nav-link" href="/contact">Report a bug</a>

			<div class="drawer-separator"></div>

			
		</nav>
	</div>
</div>
{{end}}`
