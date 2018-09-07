package slides

const (
	HelloWorld = `{{define "slide"}}
	<div class="slides">
		<section>
		<div class="mdl-grid">
		<div class="mdl-cell mdl-cell--12-col">
		<h1>Projet MEDTUNE</h1></div>
		<div class="mdl-cell mdl-cell--12-col">
			<h3>Présentation ICL Nancy</h3>
			<h6>12 Septembre 2018</h6>
		</div>
		</div>
		<div class="mdl-grid">
		<div class="mdl-cell mdl-cell--6-col mdl-cell--bottom">
		<p>Agence SII EST
		4 ,rue de Sarrelouis - 67000 Strasbourg</p>
		</div>
		<div class="mdl-cell mdl-cell--6-col mdl-cell--bottom">
		<p>Tél : 03 90 23 62 62 - Fax : 03 88 32 07 66
		www.groupe-sii.com </p>
		</div>
		</div>
		</section>
		<section>
		<h1>Projet MEDTUNE</h1>
		<h3>Présentation ICL Nancy</h3>
		<h6>12 Septembre 2018</h6>
		<h6>Agence SII EST</h6>
		<h6>4 ,rue de Sarrelouis - 67000 Strasbourg</h6>
		<h6>Tél : 03 90 23 62 62 - Fax : 03 88 32 07 66</h6>
		<h6 href=#>www.groupe-sii.com </h6>
		</section>
	</div>
{{end}}`
)
