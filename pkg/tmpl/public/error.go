package public

const Error = `{{define "content"}}
	<div class="page-error">
		<div class="mdl-grid error-grid">
			<div class="mdl-cell mdl-cell--12-col cell--zone-spacer"></div>
			<div class="mdl-cell mdl-cell--1-col cell--side-spacer"></div>

			<div class="mdl-cell mdl-cell--10-col mdl-card mdl-shadow--4dp">
				
				<div class="mdl-card__actions" style="text-align: center; font-size: 80px;">
				Error {{ .Code }}		</div>

				<div class="mdl-card__actions" style="text-align: center; color: #666699;">
					Message : {{ .Message }}
				</div>

				<div class="mdl-card__actions" style="text-align: center; color: #666699;">
					Details : {{ .Details }}
				</div>

				<div class="mdl-card__actions mdl-card--border" style="text-align:center;">
				<button class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored mdl-js-ripple-effect">
						Report bug
					</button>
					<button class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored mdl-js-ripple-effect">
						Go home
					</button>
				</div>

			</div>
		</div>
	</div>
{{end}}`
