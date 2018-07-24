package demos

const (
	Chexray = `{{define "content"}}
	<div class="demo-chexray">
	<div class="mdl-grid demo-chexray-grid">
		<div class="mdl-cell mdl-cell--12-col page-info">
			<span class="title">Chexray</span>
		</div>
		<div class="mdl-cell mdl-cell--12-col demo-chexray-card mdl-card mdl-shadow--2dp">
			<div class="mdl-card__title mdl-card__title__image">
				<h2 class="mdl-card__title-text demo-text-dark bold">Image id</h2>
			</div>

			<div class="mdl-card__actions action">
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button-play" onclick="runInference();">
					<i class="material-icons">play_arrow</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button-play" onclick="runInference();">
					<i class="material-icons>graphic_eq</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button-change" onclick="back();">
					<i class="material-icons">arrow_back</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button-change" onclick="next();">
					<i class="material-icons">arrow_forward</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button-info" onclick="information();">
					<i class="material-icons">info_outline</i>
				</button>
			</div>

		</div>
		<div class="mdl-cell mdl-cell--12-col mdl-shadow--2dp mdl-grid">
			<div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone">
				<div class="real-info">
					<h5>Image Informations</h5>
				</div>
				<div class="keykey">
					<strong>Image ID:</strong> <span id="sample-id"> 1 </span>
				</div>
				<div class="keykey">
					<strong>Approved label:</strong> <span id="sample-label"> Negative </span>
				</div>
			</div>
			<div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone">
				<div class="class-pred">
					<h5>Class Prediction</h5>
				</div>
				
			</div>
		</div>
	</div>
</div>
{{end}}`
)
