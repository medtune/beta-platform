package shared

var HeaderLogged = `{{define "header"}}
<div class="mdl-layout--fixed-header mdl-layout mdl-js-layout">
	<header class="mdl-layout__header page-header">
	<div class="mdl-layout__header-row">
			
		<span class="mdl-layout-title">MedTune <span class="v-beta">beta</span></span>
		
		<div class="mdl-layout-spacer"></div>

		<!-- Navigation -->
		<div class="navigation-container">
			<nav class="mdl-navigation">
				<a class="mdl-navigation__link" href="/home">Explore</a>
				<a class="mdl-navigation__link" href="">Account</a>
				<a class="mdl-navigation__link" href="">The Project</a>
				<a class="mdl-navigation__link" href="">Contribute</a>
			</nav>
		</div>
		<button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon">
			<i class="material-icons">cloud_upload</i>
		</button>
		<button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon">
			<i class="material-icons">settings_input_svideo</i>
		</button>
		<button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon" id="button_settings">
			<i class="material-icons">settings</i>
		</button>

		<ul class="mdl-menu mdl-menu--bottom-right mdl-js-menu mdl-js-ripple-effect" for="button_settings">
			<li class="mdl-menu__item">Settings</li>
			<li class="mdl-menu__item mdl-menu__item--full-bleed-divider">Endpoints</li>
			<li disabled class="mdl-menu__item">API</li>
			<li class="mdl-menu__item" onclick="location.href='/logout'">Logout</li>
		</ul>

	</div>

	</header>
	
	<div class="mdl-layout__drawer">
		<span class="version">v0.0.3</span>

		<div style="text-align:center">
			<img src="/static/images/logo.png" class="demo-avatar" onclick="window.location.href='/home'">
		<div>

		<nav class="mdl-navigation">

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">Explore</span>
			<a class="mdl-navigation__link drawer-nav-link" href="/tryouts">Demos</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/slides">Case studies</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/about">Slides</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/about">Data hub</a>
		
			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">Account</span>
			
			<a class="mdl-navigation__link drawer-nav-link" href="/uploads">My data</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/uploads">Informations</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/settings">Settings</a>

			<div class="drawer-separator"></div>
			<span class="mdl-navigation__link drawer-nav-title" href="">The project</span>

			<a class="mdl-navigation__link drawer-nav-link" href="/partners">About</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">Terms</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">RGPD</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/partners">Partners</a>

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
