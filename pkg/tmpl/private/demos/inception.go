package demos

const (
	InceptionImagenet = `{{define "content"}}
<div class="page-inception">
	<div class="mdl-grid grid-inception">
		<div class="mdl-cell mdl-cell--3-col mdl-cell--8-col-tablet mdl-cell--4-col-phone side-cell mdl-grid">
			<div class="settings-model mdl-cell--12-col">
				<div>
					<span class="mdl-layout__title"> Settings </span>
				</div>
				<div>
					<div class="mdl-textfield mdl-js-textfield getmdl-select">
						<input class="mdl-textfield__input" value="" id="model" readonly/>
						<input value="" type="hidden" name="model"/>
						<label class="mdl-textfield__label" for="model">Model type</label>
						<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
						<li class="mdl-menu__item" data-val="BLR">VGG-16</li>
						<li class="mdl-menu__item" data-val="BLR" disabled>VGG-18</li>

						</ul>
					</div>
				</div>
				<div>
					<div class="">
						<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" id="fti-username">
							<input class="mdl-textfield__input" type="text" pattern="^[0-9]{0,2}$" name="numofpreds" id="numofpreds"/>
							<label class="mdl-textfield__label" for="numofpreds">Number of predictions</label>
							<span class="mdl-textfield__error" id="username-error"> 1 or 2 digits number only </span>
						</div>
					</div>
				</div>
				<div>
					<span class="mdl-layout__title title-top-pred"> Top predictions </span>
				</div>
				<div class="pred-time">
					<span class="mdl-list__item-primary-content">
								<i class="material-icons mdl-list__item-icon">av_timer</i>
								<span id="elapsed-time">Elapsed time: 0 ms</span>
					</span>
				</div>
				<div id="predictions-results">
						<div id="result-1">
							<div class="mdl-list__item">
								<span id="r1">0% :</span>
								<span id="s1" class=""> Undefined </span>
							</div>
							<div id="p1" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p1').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(10);
								});
							</script>
						</div>
						<div id="result-2">
							<div class="mdl-list__item">
								<span id="r2">0% :</span>
								<span id="s2" class=""> Undefined </span>
							</div>
							<div id="p2" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p2').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(10);
								});
							</script>
						</div>
						<div id="result-3">
							<div class="mdl-list__item">
								<span id="r3">0% :</span>
								<span id="s3" class=""> Undefined </span>
							</div>
							<div id="p3" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p3').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(10);
								});
							</script>
						</div>
						<div id="result-4">
							<div class="mdl-list__item">
								<span id="r4">0% :</span>
								<span id="s4" class=""> Undefined </span>
							</div>
							<div id="p4" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p4').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(10);
								});
							</script>
						</div>
						<div id="result-5">
							<div class="mdl-list__item">
								<span id="r5">0% :</span>
								<span id="s5" class=""> Undefined </span>
							</div>
							<div id="p5" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p5').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
						<div id="result-6" class="start-hidden">
							<div class="mdl-list__item">
								<span id="r6">0% :</span>
								<span id="s6" class=""> Undefined </span>
							</div>
							<div id="p6" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p6').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
						<div id="result-7" class="start-hidden">
							<div class="mdl-list__item">
								<span id="r7">0% :</span>
								<span id="s7" class=""> Undefined </span>
							</div>
							<div id="p7" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p7').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
						<div id="result-8" class="start-hidden">
							<div class="mdl-list__item">
								<span id="r8">0% :</span>
								<span id="s8" class=""> Undefined </span>
							</div>
							<div id="p8" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p8').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
						<div id="result-9" class="start-hidden">
							<div class="mdl-list__item">
								<span id="r9">0% :</span>
								<span id="s9" class=""> Undefined </span>
							</div>
							<div id="p9" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p9').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
						<div id="result-10" class="start-hidden">
							<div class="mdl-list__item">
								<span id="r10">0% :</span>
								<span id="s10" class=""> Undefined </span>
							</div>
							<div id="p10" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
							<script>
								document.querySelector('#p10').addEventListener('mdl-componentupgraded', function() {
									this.MaterialProgress.setProgress(15);
								});
							</script>
						</div>
				</div>	
			</div>

		</div>
		<div class="mdl-cell mdl-cell--9-col mdl-cell--8-col-tablet mdl-grid">
			<div class="mdl-cell mdl-cell--12-col page-info">
				<span class="title">Inception V3: ImageNet</span>
			</div>
			{{range .Samples}}
			<div class="mdl-cell mdl-cell--4-col mdl-cell--4-col-tablet mdl-card  mdl-card mdl-shadow--4dp" id="test--">
				<div class="mdl-card__media">
					<img class="article-image" src="{{ .URL }}" border="0" alt="">
				</div>
				<div class="mdl-card__title">
					<h2 class="mdl-card__title-text" style="padding-right: 15px;"> {{ .Name }} </h2>
					<div class="mdl-progress mdl-js-progress mdl-progress__indeterminate" style="display:none;position: absolute; top: 309; left: 0;" id="load_{{ .Name }}"></div>
				</div>
				<div class="" style="position: absolute; text-align: right; width: 100%; bottom: 35px; right: 20px;">
					<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button__run" onclick="run('{{ .URL }}');">
						<i class="material-icons">arrow_forward_alt</i>
					</button>
					<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button__delete" onclick="drop('{{ .URL }}');">
						<i class="material-icons">delete</i>
					</button>
				</div>
			</div>
			{{end}}

		</div>
	</div>
</div>
{{end}}`
)
