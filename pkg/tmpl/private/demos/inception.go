package demos

const (
	InceptionImagenet = `{{define "content"}}
<div class="mdl-grid portfolio-max-width">
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
				<div class="mdl-textfield mdl-js-textfield getmdl-select">
     				 <input class="mdl-textfield__input" value="" id="country" readonly/>
      				<input value="" type="hidden" name="country"/>
      				<label class="mdl-textfield__label" for="country">Number of predictions</label>
					<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="country">
					<li class="mdl-menu__item" data-val="BLR">3</li>
					<li class="mdl-menu__item" data-val="BRA">5</li>
					<li class="mdl-menu__item" data-val="FRA">10</li>
					<li class="mdl-menu__item" data-val="DEU">15</li>
					<li class="mdl-menu__item" data-val="RUS">20</li>
					</ul>
				</div>
			</div>
			<div>
				<span class="mdl-layout__title"> Top predictions </span>
				<div class="mdl-list__item">
				<span id="s">0% :</span>
				<span id="n" class="pred-class"> Undefined </span>
				</div>
				<div id="p" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
				<div class="mdl-list__item">
				<span id="s">0% :</span>
				<span id="n" class="pred-class"> Undefined </span>
				</div>
				<div id="p" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
				<div class="mdl-list__item">
				<span id="s">0% :</span>
				<span id="n" class="pred-class"> Undefined </span>
				</div>
				<div id="p" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
				<div class="mdl-list__item">
				<span id="s">0% :</span>
				<span id="n" class="pred-class"> Undefined </span>
				</div>
				<div id="p" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
				<div class="mdl-list__item">
				<span id="s">0% :</span>
				<span id="n" class="pred-class"> Undefined </span>
				</div>
				<div id="p" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
			</div>
		</div>

    </div>
	<div class="mdl-cell mdl-cell--9-col mdl-cell--8-col-tablet mdl-grid">

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
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button__run" onclick="runDemo('{{ .URL }}');">
					<i class="material-icons">arrow_forward_alt</i>
				</button>
				<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect button__delete" onclick="deleteImage('{{ .URL }}');">
					<i class="material-icons">delete</i>
				</button>
			</div>
		</div>
		{{end}}

	</div>
</div>
{{end}}`
)
