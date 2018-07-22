package demos

const Mnist = `{{define "content"}}

<div class="demo-mnist">
	<div class="demo-goback">
			<button class="mdl-button mdl-js-button mdl-js-ripple-effect" onclick="window.location.href='/demos'">
				Go back
			</button>
	</div>
	<div class="draw-space">

	</div>
	<div class="mdl-grid demo-mnist-grid">

		<div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet" id="demo">
			<canvas id="c" width="500" height="500"></canvas>
			<script src="/static/demos/mnist/draw.js"></script>
		</div>
		<div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet mdl-grid mdl-shadow--2dp demo-mnist-ctrl">
			<div class=" mdl-cell mdl-cell--12-col">
				<h3 class="demo-mnist-ctrl__title">Control board</h3>
			</div>
			<div class="drawer-separator"></div>

			<div class=" mdl-cell mdl-cell--12-col demo-mnist-ctrl__div">
				<div id="outputs"></div>
			</div>

			<div class="mdl-cell mdl-cell--12-col demo-mnist-ctrl__div demo-mnist-ctrl__buttons">
			<div class="drawer-separator"></div>

				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-run-button" onclick="saveAndRun();">
					<i class="material-icons">play_arrow</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-clear-button" onclick="clearCanvas();">
					<i class="material-icons">clear</i>
				</button>
			
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-info-button">
					<i class="material-icons">info_outline</i>
				</button>
			</div>
		</div>
	</div>
</div>
{{end}}`
