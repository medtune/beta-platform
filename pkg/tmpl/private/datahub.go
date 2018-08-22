package private

const (
	Datahub = `{{define "content"}}
<div class="page-datahub">
	<div class="mdl-grid grid-datahub">
		<div class="mdl-cell mdl-cell--12-col mdl-card mdl-shadow--4dp">
			<div class="mdl-card__title">
				<h2 class="mdl-card__title-text">Add image</h2>
			</div>
			<div class="mdl-card__supporting-text">
				<form action="/datahub_upload" class="frm" method="post" enctype="multipart/form-data">

					<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label getmdl-select" style="width: 100%;">
						<input class="mdl-textfield__input" value="" id="model" readonly/>
						<input value="" type="hidden" name="demo"/>
						<label class="mdl-textfield__label" for="model">Demo</label>
						<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
							<li class="mdl-menu__item" data-val="inception">Inception</li>
							<li class="mdl-menu__item" data-val="mnist">MNIST</li>
							<li class="mdl-menu__item" data-val="mura">MURA</li>
							<li class="mdl-menu__item" data-val="chexray">CheXray</li>
						</ul>
					</div>

					<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" style="width: 100%;">
						<input class="mdl-textfield__input" type="text" pattern="[A-Z,a-z,0-9, ]*" name="name"/>
						<label class="mdl-textfield__label">Name</label>
						<span class="mdl-textfield__error">Letters and spaces only</span>
					</div>

					<div class="mdl-textfield mdl-js-textfield mdl-textfield--file" style="width: 100%;">
						<input class="mdl-textfield__input" placeholder="File" type="text" id="TEXT" readonly/>
						<div class="mdl-button mdl-button--primary mdl-button--icon mdl-button--file">
							<i class="material-icons">attach_file</i>
							<input type="file" id="ID" onchange="document.getElementById('TEXT').value=this.files[0].name;" name="file">
						</div>
					</div>

					<div class="datahub-uplaod-format">
						<label class="mdl-radio mdl-js-radio mdl-js-ripple-effect datahub-radio-jpeg" for="option-1">
							<input type="radio" id="option-1" class="mdl-radio__button" name="format" value="jpeg" checked>
							<span class="mdl-radio__label">JPEG</span>
						</label>
						<label class="mdl-radio mdl-js-radio mdl-js-ripple-effect datahub-radio-png" for="option-2">
							<input type="radio" id="option-2" class="mdl-radio__button" name="format" value="png">
							<span class="mdl-radio__label">PNG</span>
						</label>
					</div>

					<button class="mdl-button mdl-js-button standard__button mdl-js-ripple-effect mdl-button--raised mdl-button--colored" type="submit" value="Upload" />
						Upload
					</button>
				</form>
			</div>
		</div>
	</div> 
</div>
{{end}}`
)
