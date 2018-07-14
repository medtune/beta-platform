package shared

const Footer = `{{define "footer"}}
<footer class="mdl-mega-footer">

	<div class="mdl-mega-footer--top-section">
			<div class="mdl-mega-footer--left-section">
			<button class="mdl-mini-footer--social-btn social-btn social-btn__googleplus" onclick="">
					<span class="visuallyhidden">Mattermost</span>
				</button>
				<button class="mdl-mini-footer--social-btn social-btn social-btn__twitter" onclick="window.location.href='https://github.com/medtune-eu'">
					<span class="visuallyhidden">Twitter</span>
				</button>
				<button class="mdl-mini-footer--social-btn social-btn social-btn__github" onclick="window.location.href='https://github.com/medtune'">
				<span class="visuallyhidden">Github</span>
			</button>
            </div>
            <div class="mdl-mega-footer--right-section">

			  	<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab mdl-button--colored" onclick="scroll('', 'top', 1000);">
					<i class="material-icons">expand_less</i>
				</button>
            </div>
  </div>
  <div class="mdl-mega-footer__middle-section">

    <div class="mdl-mega-footer__drop-down-section">
      <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
      <h1 class="mdl-mega-footer__heading">Features</h1>
      <ul class="mdl-mega-footer__link-list">
        <li><a href="#">About</a></li>
        <li><a href="#">Terms</a></li>
        <li><a href="#">Partners</a></li>
      </ul>
    </div>

    <div class="mdl-mega-footer__drop-down-section">
      <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
      <h1 class="mdl-mega-footer__heading">Details</h1>
      <ul class="mdl-mega-footer__link-list">
        <li><a href="#">Specs</a></li>
        <li><a href="#">Tools</a></li>
        <li><a href="#">Resources</a></li>
      </ul>
    </div>

    <div class="mdl-mega-footer__drop-down-section">
      <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
      <h1 class="mdl-mega-footer__heading">Technology</h1>
      <ul class="mdl-mega-footer__link-list">
        <li><a href="#">How it works</a></li>
        <li><a href="#">Patterns</a></li>
        <li><a href="#">Usage</a></li>
        <li><a href="#">Products</a></li>
      </ul>
    </div>

    <div class="mdl-mega-footer__drop-down-section">
      <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
      <h1 class="mdl-mega-footer__heading">FAQ</h1>
      <ul class="mdl-mega-footer__link-list">
        <li><a href="#">Questions</a></li>
        <li><a href="#">Answers</a></li>
        <li><a href="#">Contact us</a></li>
      </ul>
    </div>

  </div>

  <div class="mdl-mega-footer__bottom-section">
    <div class="mdl-logo">Medtune Beta (version iron-0.0.3) </div>
    <ul class="mdl-mega-footer__link-list">
      <li><a href="#">Help</a></li>
      <li><a href="#">Privacy & Terms</a></li>
    </ul>
  </div>

</footer>
{{end}}`
