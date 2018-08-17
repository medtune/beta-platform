package private

const (
	Datahub = `{{define "content"}}
<div class="page-datahub">
<div class="mdl-grid grid-datahub">
    <div class="mdl-cell mdl-cell--12-col mdl-card mdl-shadow--4dp">
        <div class="mdl-card__title">
            <h2 class="mdl-card__title-text">Upload Image</h2>
        </div>
        <div class="mdl-card__supporting-text">
            <form action="/api/datahub_upload" class="frm" method="post" enctype="multipart/form-data">
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label" style="width: 100%;">
					<input class="mdl-textfield__input" type="text" pattern="[A-Z,a-z,0-9, ]*" name="name"/>
					<label class="mdl-textfield__label">Name</label>
					<span class="mdl-textfield__error">Letters and spaces only</span>
				</div>
				<div>
					<div class="mdl-textfield mdl-js-textfield getmdl-select" style="width: 100%;">
						<input class="mdl-textfield__input" value="" id="model" readonly/>
						<input value="" type="hidden" name="model"/>
						<label class="mdl-textfield__label" for="model">Demo</label>
						<ul class="mdl-menu mdl-menu--bottom-left mdl-js-menu" for="model">
							<li class="mdl-menu__item" data-val="MNS" disabled>MNIST</li>
							<li class="mdl-menu__item" data-val="INC">Inception</li>
						</ul>
					</div>
				</div>
				<div class="mdl-textfield mdl-js-textfield mdl-textfield--file" style="width: 100%;">
					<input class="mdl-textfield__input" placeholder="File" type="text" id="TEXT" readonly/>
					<div class="mdl-button mdl-button--primary mdl-button--icon mdl-button--file">
						<i class="material-icons">attach_file</i>
						<input type="file" id="ID" onchange="document.getElementById('TEXT').value=this.files[0].name;" name="file">
					</div>
				</div>
				<button class="mdl-button mdl-js-button standard__button mdl-js-ripple-effect" type="submit" value="Upload" />
                    Send
            	</button>
			</form>
        </div>
    </div>
</div> 
</div>
{{end}}`
)
