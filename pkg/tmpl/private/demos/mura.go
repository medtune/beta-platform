package demos

const (
	Mura = `{{define "content"}}
<div class="demo-mura">
	<div class="mdl-grid demo-mura-grid">
		<div class="mdl-cell mdl-cell--12-col demo-menu-card mdl-card mdl-shadow--2dp">
			<div class="mdl-card__title mdl-card__title__image">
				<h2 class="mdl-card__title-text demo-text-dark bold">Image id</h2>
			</div>

			<div class="mdl-card__actions action-left">
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect" onclick="">
					<i class="material-icons">arrow_back</i>
				</button>
			</div>
			<div class="mdl-card__actions action-right">
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect" onclick="">
					<i class="material-icons">arrow_forward</i>
				</button>
			</div>
		</div>
	</div>
</div>
{{end}}`
)
