package shared

var HeaderLogged = `{{define "header"}}
<div class="mdl-layout--fixed-header mdl-layout mdl-js-layout">
	<header class="mdl-layout__header page-header">
	<div class="mdl-layout__header-row">
			<div class="search-box mdl-textfield mdl-js-textfield mdl-textfield--expandable mdl-textfield--floating-label mdl-textfield--align-right mdl-textfield--full-width">
			<label class="mdl-button mdl-js-button mdl-button--icon" for="search-field">
			<i class="material-icons">search</i>
			</label>
			<div class="mdl-textfield__expandable-holder">
			<input class="mdl-textfield__input" type="text" id="search-field">
			</div>
		</div>
		<span class="mdl-layout-title">MedTune</span>

		<div class="mdl-layout-spacer"></div>

		<!-- Navigation -->
		<div class="navigation-container">
			<nav class="mdl-navigation">
				<a class="mdl-navigation__link" href="/home">Home</a>
				<a class="mdl-navigation__link" href="">Demos</a>
				<a class="mdl-navigation__link" href="">Slides</a>
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
		<span class="version">v 0.0.2</span>

		<div style="text-align:center">
			<img src="/static/images/logo.png" class="demo-avatar" onclick="window.location.href='/home'">
		<div>

		<nav class="mdl-navigation">
			<div class="drawer-separator"></div>
			<a class="mdl-navigation__link drawer-nav-link" href="/tryouts">Tryouts</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/slides">Slides</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/about">About</a>
			<div class="drawer-separator"></div>
			<a class="mdl-navigation__link drawer-nav-link" href="/uploads">Uploads</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/capsules">Capsules</a>
			<a class="mdl-navigation__link drawer-nav-link" href="/settings">Settings</a>
			<div class="drawer-separator"></div>
		</nav>
	</div>
</div>
{{end}}`
