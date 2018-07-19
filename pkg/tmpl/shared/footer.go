package shared

const Footer = `{{define "footer"}}
<div class="social-section">
    <div class="">
        <div class="madeby">
            <span class="madeby-text"> Made with <i class="material-icons madeby-love">favorite</i> by <img class="madeby-img" height="20" src="/static/images/siiresearch.png"></img> </span>
        </div>
        <div>
            <button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab scroll-top-button" onclick="scrollToTop();">
                <i class="material-icons">expand_less</i>
            </button>
        </div>
    </div>
</div>

<footer class="mdl-mega-footer footing">

	
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

    <div class="mdl-mega-footer__drop-down-section footer-faq">
      <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
      <h1 class="mdl-mega-footer__heading">FAQ</h1>
      <ul class="mdl-mega-footer__link-list">
        <li><a href="#">Questions</a></li>
        <li><a href="#">Answers</a></li>
        <li><a href="#">Contact us</a></li>
        <li><a href="#">Help</a></li>
      </ul>
    </div>

  </div>

  <div class="mdl-mega-footer__bottom-section">
  <div class="social-links">
            <button class="social-btn social-btn__googleplus" onclick="">
                <span class="visuallyhidden">Mattermost</span>
            </button>
            <button class="social-btn social-btn__twitter" onclick="window.location.href='https://github.com/medtune-eu'">
                <span class="visuallyhidden">Twitter</span>
            </button>
            <button class="social-btn social-btn__github" onclick="window.location.href='https://github.com/medtune'">
                <span class="visuallyhidden">Github</span>
            </button>
      </div>
  </div>
</footer>
{{end}}`
