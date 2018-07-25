package demos

const Mnist = `{{define "content"}}

<div class="demo-mnist">
	<div class="mdl-grid demo-mnist-grid">
		<div class="mdl-cell mdl-cell--12-col page-info">
			<span class="title">MNIST: Handwritten digits</span>
		</div>
		<div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet" id="demo">
			<canvas id="c" width="500" height="500"></canvas>
		</div>
		<div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-grid mdl-shadow--2dp demo-mnist-ctrl">
			<div class=" mdl-cell mdl-cell--12-col">
				<h3 class="demo-mnist-ctrl__title">Control board</h3>
			</div>
			<div class="drawer-separator"></div>

			<div class=" mdl-cell mdl-cell--12-col demo-mnist-ctrl__div">
			<ul class="demo-list-icon mdl-list">
				<li class="mdl-list__item">
					<span class="mdl-list__item-primary-content">
						<i class="material-icons mdl-list__item-icon">av_timer</i>
						<span id="elapsed-time">Elapsed time:</span>
					</span>
				</li>
				<li class="mdl-list__item">
					<span class="mdl-list__item-primary-content">
						<i class="material-icons mdl-list__item-icon">check_circle_outline</i>
						<span id="top-prediction">Top prediction:</span>
					</span>
				</li>
				
			</ul>
				
			</div>

			<div class="mdl-cell mdl-cell--12-col demo-mnist-ctrl__div demo-mnist-ctrl__buttons">
				<div class="drawer-separator"></div>

				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-run-button" onclick="saveAndRun();">
				    <i class="material-icons">play_arrow</i>
				</button>

				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-edit-button" onclick="draw();">
					<i class="material-icons">edit</i>
				</button>

				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-shuffle-button" onclick="shuffle();">
					<i class="material-icons">shuffle</i>
				</button>

				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-clear-button" onclick="clearCanvas();">
					<i class="material-icons">clear</i>
				</button>
			
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect demo-mnist-ctrl__button demo-mnist-info-button">
					<i class="material-icons">info_outline</i>
				</button>
				
			</div>
		</div>
		<div class="mdl-cell mdl-cell--12-col mdl-cell--8-col-tablet mdl-grid mdl-shadow--2dp demo-mnist-logs">
			<div><h3>LOGS [SERVICE LOCKED]</h3></div>
		</div>
	</div>
</div>
<script src="/static/demos/mnist/draw.js"></script>
{{end}}`
