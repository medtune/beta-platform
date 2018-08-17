package demos

const (
	MuraV2 = `{{define "content"}}
<div class="demo-mura">
	<div class="mdl-grid">
		<div class="mdl-cell mdl-cell--3-col mdl-card"></div>
		<div class="mdl-cell mdl-cell--3-col mdl-card mdl-shadow--6dp">

			<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
				<h2 class="mdl-card__title-text">Control board</h2>
			</div>

			<div class="mdl-card__supporting-text">

				<div class="mdl-textfield mdl-js-textfield getmdl-select" style="width: 100%;">
					<input class="mdl-textfield__input" value="" id="model" readonly/>
					<input value="" type="hidden" name="model"/>
					<label class="mdl-textfield__label" for="model">Model base</label>
					<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
						<li class="mdl-menu__item" data-val="MBL">MobileNet</li>
						<li class="mdl-menu__item" data-val="DNS" disabled>DenseNet</li>
					</ul>
				</div>

				<div class="mdl-textfield mdl-js-textfield mdl-textfield--file" style="width: 100%;">
					<input class="mdl-textfield__input" placeholder="File" type="text" id="TEXT" readonly/>
					<div class="mdl-button mdl-button--primary mdl-button--icon mdl-button--file">
						<i class="material-icons">attach_file</i>
						<input type="file" id="ID" onchange="document.getElementById('TEXT').value=this.files[0].name;" name="file">
					</div>
				</div>
				<div>
					<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-2">
						<input type="checkbox" id="switch-2" class="mdl-switch__input" checked>
						<span class="mdl-switch__label">Normalize</span>
					</label>
				</div>
			</div>

			<div class="mdl-card__actions mdl-card--border">
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Preview</button>
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Predict</button>
				<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">GRAD-CAM</button>

			</div>

		</div>
	</div>
</div>
{{end}}`
)
