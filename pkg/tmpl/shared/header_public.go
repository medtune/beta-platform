package shared

var HeaderPublic = `{{define "header"}}
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
					<a class="mdl-navigation__link" href="/signup">Sign up</a>
					<a class="mdl-navigation__link" href="/login">Login</a>
				</nav>
			</div>

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
