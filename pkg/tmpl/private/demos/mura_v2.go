package demos

const (
	MuraV2 = `{{define "content"}}
<div id="overlay" onclick="closeFAB();"></div>
<div class="demo-mura">
	<div class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">

		<div class="mdl-tabs__tab-bar">
			<a href="#cases-panel" class="mdl-tabs__tab is-active">cases</a>
			<a href="#upload-panel" class="mdl-tabs__tab">Upload</a>
			<a href="#settings-panel" class="mdl-tabs__tab">Settings</a>
			<a href="#informations-panel" class="mdl-tabs__tab">Informations</a>
		</div>

		<!-- ----------- -->
		<!-- RESULT POP UP -->
		<!-- ----------- -->

		<div class="fab" id="fabm">
		
		    <div class="cntt-wrapper" id="cntt-wrapper">
				<div id="fab-hdr">
			  		<h3 id="title">Result</h3>
			  		<div class="btn-wrapper">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab btn-cancel" id="cancel" onclick="closeFAB();">
							<i class="material-icons">clear</i>
						</button>
					</div>
				</div>
		
				<div class="cntt mdl-grid" id="cntt">
					<div class="mdl-cell mdl-cell--6-col">
						<div class="results-bars">
						<div class="mdl-list__item">
							<span id="r1">{class} :</span>
							<span id="s1" class=""> {score} </span>
						</div>
						<div id="p1" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
						<div class="mdl-list__item">
							<span id="r2">{class} :</span>
							<span id="s2" class=""> {score}</span>
						</div>
						<div id="p2" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 20px"></div>
					</div>	
					</div>
					<div class="mdl-cell mdl-grid mdl-cell--6-col metadata">
						<div class="container-ctx">
						<div class="demo-list-action mdl-list stats-list">
						<div class="mdl-list__item">

						  <span class="mdl-list__item-primary-content">
							<i class="material-icons mdl-list__item-icon">timer</i>
							<span class="stats-sub-title">Inference time</span>
						  </span>
						  <a class="mdl-list__item-secondary-action stats-value" href="#" id="inference-time">392</a>
						</div>
						<div class="mdl-list__item" id="cam-timing">
						  <span class="mdl-list__item-primary-content">
							<i class="material-icons mdl-list__item-icon">timer</i>
							<span class="stats-sub-title">Grad cam time</span>
						  </span>
						  <a class="mdl-list__item-secondary-action stats-value" href="#" id="grad-cam-time">11</a>
						</div>
						<div class="mdl-list__item">
						  <span class="mdl-list__item-primary-content">
							<i class="material-icons mdl-list__item-icon">timer</i>
							<span class="stats-sub-title">Request roundtrip</span>
						  </span>
						  <a class="mdl-list__item-secondary-action  stats-value" href="#" id="req-roundtrip">1232</a>
						</div>
					  </div>

					  </div>
						
					</div>
					

					<div class="mdl-cell mdl-cell--6-col demo-card-square mdl-card mdl-shadow--4dp" id="preview-cc">
						<div class="mdl-card__supporting-text">
							Preview
						</div>
						<div class="mdl-card__title mdl-card--expand" id="preview">
						</div>
					</div>

					<div class="mdl-cell mdl-cell--6-col demo-card-square mdl-card mdl-shadow--4dp" id="cam-cc">
						<div class="mdl-card__supporting-text">
							Grad CAM
						</div>
						<div class="mdl-card__title mdl-card--expand" id="cam">
						</div>
					</div>

				</div>
		 	</div>
		</div>

		<script src="/static/demos/mura/v2/load.js"></script>

		<!-- ----------- -->
		<!-- CASES PANEL -->
		<!-- ----------- -->

		<div class="mdl-tabs__panel is-active" id="cases-panel">
			<div class="mdl-grid mura-grid-view">
				{{range .Samples}}
				<div class="mdl-cell mdl-cell--4-col mdl-cell--4-col-tablet demo-card-square mdl-card mdl-shadow--2dp">
					<div class="mdl-card__media">
						<img class="article-image" src="/static/demos/mura/images/{{ .Filename }}" border="0" alt="">
					</div>

					<div class="mdl-card__supporting-text">
						{{ .Name }}
					</div>

					<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect ctrl-btn run-btn" onclick="processCase('{{ .Filename }}');">
						<i class="material-icons">arrow_forward_ios</i>
					</button>

					<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect ctrl-btn preview-btn" onclick="previewCase('{{ .Filename }}');">
						<i class="material-icons">remove_red_eye</i>
					</button>
	
				</div>
				{{end}}
			</div>
		</div>

		<!-- ------------ -->
		<!-- UPLOAD PANEL -->
		<!-- ------------ -->

		<div class="mdl-tabs__panel" id="upload-panel">
			<div class="settings-container">
				<div class="mdl-card mdl-shadow--4dp settings-card">

					<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
						<h2 class="mdl-card__title-text">Upload</h2>
					</div>

					<div class="mdl-card__supporting-text mura-settings-ctx">
						<form action="/demos/mura/upload" method="POST" id="upload-form" enctype="multipart/form-data">
						
							<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label name-label" style="">
								<input class="mdl-textfield__input" type="text" pattern="[A-Z,a-z, ]*" name="name"/>
								<label class="mdl-textfield__label">Name</label>
								<span class="mdl-textfield__error">Letters and spaces only</span>
							</div>

							<div class="mdl-textfield mdl-js-textfield mdl-textfield--file file-label">
								<input class="mdl-textfield__input" placeholder="File" type="text" id="FF" readonly/>
								<div class="mdl-button mdl-button--primary mdl-button--icon mdl-button--file">
									<i class="material-icons">attach_file</i>
									<input type="file" id="ID" onchange="document.getElementById('FF').value=this.files[0].name;" name="file">
								</div>
							</div>
							<div class="mdl-card__actions mdl-card--border" style="margin-top:50px;">
								<button class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" type="submit">Send</button>
							</div>
						</form>
						
					</div>
				</div>
			</div>
		</div>

		<!-- -------------- -->
		<!-- SETTINGS PANEL -->
		<!-- -------------- -->

		<div class="mdl-tabs__panel" id="settings-panel">
			<div class="settings-container">
				<div class="mdl-card mdl-shadow--4dp settings-card">
					<div class="mdl-card__title mdl-color--primary mdl-color-text--white">
						<h2 class="mdl-card__title-text">Settings</h2>
					</div>
					<div class="mdl-card__supporting-text mura-settings-ctx">

						<div class="mdl-textfield mdl-js-textfield getmdl-select model-dropdown">
							<input class="mdl-textfield__input" value="" id="model" readonly/>
							<input value="" type="hidden" name="model-engine"/>
							<label class="mdl-textfield__label" for="model">Inference Model</label>
							<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
								<li class="mdl-menu__item" data-val="mobilenet-v2">MobileNet V2</li>
								<li class="mdl-menu__item" data-val="inception-resnet-v2">Inception ResNet V2</li>
							</ul>
						</div>

						<div>
							<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-1">
								<input type="checkbox" id="switch-1" class="mdl-switch__input">
								<span class="mdl-switch__label">Normalize</span>
							</label>
							<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-2">
								<input type="checkbox" id="switch-2" class="mdl-switch__input">
								<span class="mdl-switch__label">Auto save results</span>
							</label>
							<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-3">
								<input type="checkbox" id="switch-3" class="mdl-switch__input">
								<span class="mdl-switch__label">Show logs</span>
							</label>
							<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-4">
								<input type="checkbox" id="switch-4" class="mdl-switch__input">
								<span class="mdl-switch__label">Calculate round trip</span>
							</label>
							<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect" for="switch-5">
								<input type="checkbox" id="switch-5" class="mdl-switch__input">
								<span class="mdl-switch__label">Always run GRAD-CAM</span>
							</label>
						</div>


					</div>

				</div>

			</div>
		</div>

		<div class="mdl-tabs__panel" id="informations-panel">
		test
		</div>

	</div>
</div>
{{end}}`
)
