package demos

const PolynomialRegression = `{{define "content"}}
<div class="demo-polyreg">
	<div class="demo-goback">
			<button class="mdl-button mdl-js-button mdl-js-ripple-effect">
				Go back
			</button>
	</div>
	<div class="mdl-grid demo-polyreg-grid">
		<div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet" id="demo">
		</div>
		<div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet mdl-grid mdl-shadow--2dp demo-polyreg-ctrl">
			<div class=" mdl-cell mdl-cell--12-col">
				<h3 class="demo-polyreg-ctrl__title">Control board</h3>
			</div>
			<div class=" mdl-cell mdl-cell--12-col demo-polyreg-ctrl__div">
				<div id="outputs"></div>
			</div>
			<div class="mdl-cell mdl-cell--12-col demo-polyreg-ctrl__div" id="inputs">
				<label>Order of polynomial</label>
				<input class="mdl-slider mdl-js-slider" type="range" id="orderPolySlider" min="1" max="10" step="1" value="1">
			</div>
			<div class="mdl-cell mdl-cell--12-col demo-polyreg-ctrl__div">
				<label>Learning Rate</label>
				<input class="mdl-slider mdl-js-slider" id="learningRateSlider" type="range" min="0.001" max="0.3" step="0.001" value="0.001">
			</div>
			<div class="mdl-cell mdl-cell--12-col demo-polyreg-ctrl__div demo-polyreg-ctrl__buttons">
				<button id="clearPointsButton" class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-polyreg-ctrl__button demo-polyreg-clear-button">
					<i class="material-icons">clear</i>
				</button>
				<button id="clearPointsButton" class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-polyreg-ctrl__button demo-polyreg-info-button">
					<i class="material-icons">info</i>
				</button>
			</div>
		</div>
	</div>
</div>
{{end}}`
