package public

const Login = `{{define "content"}}
<div class="center-screen">
	<div class="mdl-card mdl-shadow--6dp login-card">
		<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
			<h2 class="mdl-card__title-text">Login</h2>
		</div>
		<div class="mdl-card__supporting-text">
			<form action="/login" method="POST" id="login-form">
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-username">
					<input class="mdl-textfield__input" type="text" pattern="^[a-zA-Z][a-zA-Z0-9-_\.]{3,20}$" name="username" id="username"/>
					<label class="mdl-textfield__label" for="username">Username</label>
					<span class="mdl-textfield__error" id="username-error">3 - 20 Characters (letters and numbers)</span>
				</div>
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-password">
					<input class="mdl-textfield__input" type="password" name="password" id="password" />
					<label class="mdl-textfield__label" for="userpass">Password</label>
					<span class="mdl-textfield__error" id="password-error"></span>
				</div>
			</form>
		</div>
		<div class="mdl-progress mdl-js-progress mdl-progress__indeterminate load-bar-form" id="login-load"></div>

		<div class="mdl-card__actions mdl-card--border">
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" onclick="Login();">Log in</button>
		</div>
	</div>
</div>
{{end}}`
