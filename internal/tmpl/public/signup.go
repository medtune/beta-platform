package public

const Signup = `{{define "content"}}
<div class="center-screen">
	<div class="mdl-card mdl-shadow--6dp signup-card">
		<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
			<h2 class="mdl-card__title-text">Sign up</h2>
		</div>
		<div class="mdl-card__supporting-text">
			<form action="/api/signup" method="post" id="signup-form">
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-email">
					<input class="mdl-textfield__input" type="text" name="email" id="email"/>
					<label class="mdl-textfield__label" for="email">Email</label>
					<span class="mdl-textfield__error" id="email-error">Email error</span>

				</div>
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-username">
					<input class="mdl-textfield__input" type="text" pattern="^[a-zA-Z][a-zA-Z0-9-_\.]{3,20}$" name="username" id="username"/>
					<label class="mdl-textfield__label" for="username">Username</label>
					<span class="mdl-textfield__error" id="username-error">3 - 20 characters (Letters and numbers)</span>

				</div>
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-password">
					<input class="mdl-textfield__input" pattern=".{8,}" type="password" name="password" id="password" />
					<label class="mdl-textfield__label" for="userpass">Password</label>
					<span class="mdl-textfield__error" id="password-error">At least 8 characters</span>

				</div>
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-passwordc">
					<input class="mdl-textfield__input" type="password" name="passwordc" id="passwordc" />
					<label class="mdl-textfield__label" for="passwordc">Confirm password</label>
					<span class="mdl-textfield__error" id="passwordc-error">Passwords don't match</span>

				</div>
				
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-secret">
					<input class="mdl-textfield__input" type="text" pattern="^[a-zA-Z][a-zA-Z0-9-_\.]{0,20}$" name="secret" autocomplete="" id="secret"/>
					<label class="mdl-textfield__label" for="secret">Secret Key</label>
					<span class="mdl-textfield__error" id="secret-error">Invalid secret key</span>
				</div>

			</form>
		</div>
		<div class="mdl-progress mdl-js-progress mdl-progress__indeterminate load-bar-form" id="signup-load"></div>
		<div class="mdl-card__actions mdl-card--border">
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" onclick="Signup();">Sign up</button>
		</div>
	</div>
</div>
<script>
$(document).keydown(function(e){
    if (e.keyCode == 13) { 
		Signup();
		return false;
    }
});
</script>
{{end}}`
