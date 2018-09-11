package demos

const (
	ChexrayV2 = `{{define "content"}}
<div id="overlay" onclick="closeFAB();"></div>
<div class="demo-chexray">
	<div class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">

		<div class="mdl-tabs__tab-bar" id="tab-bar">
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

				<div class="cntt-prev-ctrl mdl-grid" id="cntt-prev-ctrl">
					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-brightness">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Brightness</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-brightness">
							<i class="material-icons">add</i>
						</button>
					</div>

					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-contrast">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Contrast</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-contrast">
							<i class="material-icons">add</i>
						</button>
					</div>

					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-saturation">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Saturation</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-saturation">
							<i class="material-icons">add</i>
						</button>
					</div>

					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-noise">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Noise</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-noise">
							<i class="material-icons">add</i>
						</button>
					</div>

					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-hue">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Hue</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-hue">
							<i class="material-icons">add</i>
						</button>
					</div>
					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="remove-vibrance">
							<i class="material-icons">remove</i>
						</button>
						<span style="color: black">Vibrance</span>
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab" id="add-vibrance">
							<i class="material-icons">add</i>
						</button>
					</div>
					<div class="mdl-cell mdl-cell-3--col" style="width: auto;">
						<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab prev-clear-btn" id="prev-clear">
							<i class="material-icons">clear</i>
						</button>
					</div>
				</div>
				<div class="cntt-prev mdl-grid" id="cntt-prev">
					<canvas class="mdl-cell mdl-cell-12--col" id="canvas" style="height: 570px;"></canvas>
				</div>
		
				<div class="cntt mdl-grid" id="cntt">


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
							</div>
						</div>
					</div>
					<div class="mdl-cell mdl-grid mdl-cell--6-col metadata">
						<div class="container-ctx">
							<div class="demo-list-action mdl-list stats-list">

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

					<div class="mdl-cell mdl-cell--4-col" style="margin-top: -35px;">
						<div class="results-bars">
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r1">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s1" class=""> {score} </span>
							</div>
							<div id="p1" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r2">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s2" class=""> {score}</span>
							</div>
							<div id="p2" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r3">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s3" class=""> {score}</span>
							</div>
							<div id="p3" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r4">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s4" class=""> {score}</span>
							</div>
							<div id="p4" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r5">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s5" class=""> {score}</span>
							</div>
							<div id="p5" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
						</div>	
					</div>

					<div class="mdl-cell mdl-cell--4-col" style="margin-top: -35px;">
						<div class="results-bars">
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r6">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s6" class=""> {score} </span>
							</div>
							<div id="p6" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r7">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s7" class=""> {score}</span>
							</div>
							<div id="p7" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r8">
									Button
								</button>
								<div class="mdl-layout-spacer"></div>
								<span id="s8" class=""> {score}</span>
							</div>
							<div id="p8" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r9">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s9" class=""> {score}</span>
							</div>
							<div id="p9" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r10">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s10" class=""> {score}</span>
							</div>
							<div id="p10" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
						</div>	
					</div>
					
					<div class="mdl-cell mdl-cell--4-col" style="margin-top: -35px;">
						<div class="results-bars">
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r11">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s11" class=""> {score} </span>
							</div>
							<div id="p11" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r12">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s12" class=""> {score}</span>
							</div>
							<div id="p12" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r13">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s13" class=""> {score}</span>
							</div>
							<div id="p13" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r14">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s14" class=""> {score}</span>
							</div>
							<div id="p14" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
							<div class="mdl-list__item" style="margin-bottom: -15px;margin-left: -15px;">
								<button class="mdl-button mdl-js-button mdl-button--primary" id="r15">
									Button
								  </button>
								  <div class="mdl-layout-spacer"></div>
								<span id="s15" class=""> {score}</span>
							</div>
							<div id="p15" class="mdl-progress prog-bar mdl-js-progress" style="width: 400px; height: 11px;"></div>
						</div>	
					</div>
					
					

					<div class="mdl-cell mdl-cell--12-col demo-card-square mdl-card mdl-shadow--4dp" id="preview-cc">
						<div class="mdl-card__supporting-text">
							Preview
						</div>
						<div class="mdl-card__title mdl-card--expand" id="preview">
						</div>
					</div>

					<div class="mdl-cell mdl-cell--12-col demo-card-square mdl-card mdl-shadow--4dp" id="cxpba-cc" style="margin-top: 50px;">

						<button class="mdl-button mdl-js-button mdl-button--primary" id="hide-btn" onclick="cxpbaCC.slideUp();">
							Hide
						 </button>
						<div class="mdl-card__supporting-text" id="table-title">
							CXPBA
						</div>
						<div class="" id="cxpba-table">
							<table class="mdl-data-table mdl-js-data-table mdl-shadow--2dp table-cxpba">
								<thead>
									<tr>
										<th class="mdl-data-table__cell--non-numeric">Analysis</th>
										<th>Unit</th>
										<th>Min</th>
										<th>Max</th>
										<th>Expected Result</th>
									</tr>
								</thead>
								<tbody id="table-tbody">
									{{range .Cxpbatable}}
									<tr>
										<td class="mdl-data-table__cell--non-numeric">{{ .Name }}</td>
										<td>{{.Unit}}</td>
										<td>{{.Min}}</td>
										<td>{{.Max}}</td>
										<td id="{{ .Name }}">0</td>
									</tr>
									{{end}}
								</tbody>
							</table>
						</div>
					</div>

				</div>
		 	</div>
		</div>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/camanjs/4.1.2/caman.full.min.js"></script>
		<script src="/static/demos/chexray/v2/load.js"></script>

		<!-- ----------- -->
		<!-- CASES PANEL -->
		<!-- ----------- -->

		<div class="mdl-tabs__panel is-active" id="cases-panel">
			<div class="mdl-grid chexray-grid-view">
				{{range .Samples}}
				<div class="mdl-cell mdl-cell--4-col mdl-cell--4-col-tablet demo-card-square mdl-card mdl-shadow--2dp">
					<div class="mdl-card__media">
						<img class="article-image" src="/static/demos/chexray/images/{{ .Filename }}" border="0" alt="">
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

					<div class="mdl-card__supporting-text chexray-settings-ctx">
						<form action="/demos/chexray/upload" method="POST" id="upload-form" enctype="multipart/form-data">
						
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
					<div class="mdl-card__supporting-text chexray-settings-ctx">

						<div class="mdl-textfield mdl-js-textfield getmdl-select model-dropdown">
							<input class="mdl-textfield__input" value="" id="model" readonly/>
							<input value="" type="hidden" name="model-engine"/>
							<label class="mdl-textfield__label" for="model">Inference Model</label>
							<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
								<li class="mdl-menu__item" data-val="densenet-121">DenseNet 121</li>
								<li class="mdl-menu__item" data-val="mobilenet-v2">MobileNet V2</li>
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
								<span class="mdl-switch__label" disabled>Always run GRAD-CAM</span>
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
