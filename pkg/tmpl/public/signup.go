package public

const Signup = `{{define "content"}}
<div class="center-screen">
	<div class="mdl-card mdl-shadow--6dp signup-card">
		<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
			<h2 class="mdl-card__title-text">Sign up</h2>
		</div>
		<div class="mdl-card__supporting-text">
			<form action="/signup" method="post" id="signup-form">
				<div class="mdl-textfield mdl-js-textfield">
					<input class="mdl-textfield__input" type="text" name="username" id="username"/>
					<label class="mdl-textfield__label" for="username">Username</label>
				</div>
				<div class="mdl-textfield mdl-js-textfield">
					<input class="mdl-textfield__input" type="password" name="password" id="password" />
					<label class="mdl-textfield__label" for="userpass">Password</label>
				</div>
				<div class="mdl-textfield mdl-js-textfield">
					<input class="mdl-textfield__input" type="password" name="passwordc" id="passwordc" />
					<label class="mdl-textfield__label" for="passwordc">Confirm password</label>
				</div>
				<div class="mdl-textfield mdl-js-textfield">
					<input class="mdl-textfield__input" type="text" name="secret" id="secret"/>
					<label class="mdl-textfield__label" for="secret">Secret Key</label>
				</div>
			</form>
		</div>
		<div class="mdl-progress mdl-js-progress mdl-progress__indeterminate load-bar-form" id="signup-load"></div>
		<div class="mdl-card__actions mdl-card--border">
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" onclick="Signup();">Sign up</button>
		</div>
	</div>
</div>
{{end}}`
