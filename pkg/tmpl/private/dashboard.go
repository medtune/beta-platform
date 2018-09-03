package private

const Dashboard = `{{define "dashboard"}}
<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">
	<header class="mdl-layout__header mdl-layout__header--scroll mdl-color--primary">
		<div class="mdl-layout__tab-bar mdl-js-ripple-effect mdl-color--primary-dark">
		<a href="#overview" class="mdl-layout__tab is-active">Overview</a>
			<a href="#features" class="mdl-layout__tab">Stats</a>
			<a href="#features" class="mdl-layout__tab">Capsules</a>
			<a href="#features" class="mdl-layout__tab">Storage</a>
			<a href="#config" class="mdl-layout__tab">Config</a>
			<a href="#features" class="mdl-layout__tab"></a>
			<button class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--colored mdl-shadow--4dp mdl-color--accent" id="add">
			<i class="material-icons" role="presentation">add</i>
			<span class="visuallyhidden">Add</span>
			</button>
		</div>
	</header>
		<main class="mdl-layout__content">
			<div class="mdl-layout__tab-panel is-active" id="overview">
		
            	<section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
					<div class="mdl-card mdl-cell mdl-cell--12-col">
					<div class="mdl-card__supporting-text">
						<h4>Test</h4>
						<pre>{{ .Version }}</pre>
					</div>
						<div class="mdl-card__actions">
							<a href="#" class="mdl-button">Test</a>
						</div>
					</div>
					<button class="mdl-button mdl-js-button mdl-js-ripple-effect mdl-button--icon" id="btn1">
						<i class="material-icons">more_vert</i>
					</button>
					<ul class="mdl-menu mdl-js-menu mdl-menu--bottom-right" for="btn1">
						<li class="mdl-menu__item">Lorem</li>
						<li class="mdl-menu__item" disabled>Ipsum</li>
						<li class="mdl-menu__item">Dolor</li>
					</ul>
				  </section>
				  

          <section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
            <div class="mdl-card mdl-cell mdl-cell--12-col">
              <div class="mdl-card__supporting-text mdl-grid mdl-grid--no-spacing">
                <h4 class="mdl-cell mdl-cell--12-col">Details</h4>
                <div class="section__circle-container mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
                  <div class="section__circle-container__circle mdl-color--primary"></div>
                </div>
                <div class="section__text mdl-cell mdl-cell--10-col-desktop mdl-cell--6-col-tablet mdl-cell--3-col-phone">
                  <h5>Lorem ipsum dolor sit amet</h5>
                  Dolore ex deserunt aute fugiat aute nulla ea sunt aliqua nisi cupidatat eu. Duis nulla tempor do aute et eiusmod velit exercitation nostrud quis <a href="#">proident minim</a>.
                </div>
                <div class="section__circle-container mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
                  <div class="section__circle-container__circle mdl-color--primary"></div>
                </div>
                <div class="section__text mdl-cell mdl-cell--10-col-desktop mdl-cell--6-col-tablet mdl-cell--3-col-phone">
                  <h5>Lorem ipsum dolor sit amet</h5>
                  Dolore ex deserunt aute fugiat aute nulla ea sunt aliqua nisi cupidatat eu. Duis nulla tempor do aute et eiusmod velit exercitation nostrud quis <a href="#">proident minim</a>.
                </div>
                <div class="section__circle-container mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
                  <div class="section__circle-container__circle mdl-color--primary"></div>
                </div>
                <div class="section__text mdl-cell mdl-cell--10-col-desktop mdl-cell--6-col-tablet mdl-cell--3-col-phone">
                  <h5>Lorem ipsum dolor sit amet</h5>
                  Dolore ex deserunt aute fugiat aute nulla ea sunt aliqua nisi cupidatat eu. Duis nulla tempor do aute et eiusmod velit exercitation nostrud quis <a href="#">proident minim</a>.
                </div>
              </div>
              <div class="mdl-card__actions">
                <a href="#" class="mdl-button">Test</a>
              </div>
            </div>
            <button class="mdl-button mdl-js-button mdl-js-ripple-effect mdl-button--icon" id="btn2">
              <i class="material-icons">more_vert</i>
            </button>
            <ul class="mdl-menu mdl-js-menu mdl-menu--bottom-right" for="btn2">
              <li class="mdl-menu__item">Lorem</li>
              <li class="mdl-menu__item" disabled>Ipsum</li>
              <li class="mdl-menu__item">Dolor</li>
            </ul>
		  </section>
		

		</div>
		

		<div class="mdl-layout__tab-panel" id="features">
		

          <section class="section--center mdl-grid mdl-grid--no-spacing">
            <div class="mdl-cell mdl-cell--12-col">
              <h4>Test</h4>
              Minim duis incididunt est cillum est ex occaecat consectetur. Qui sint ut et qui nisi cupidatat. Reprehenderit nostrud proident officia exercitation anim et pariatur ex.
              <ul class="toc">
                <h4>Contents</h4>
                <a href="#lorem1">Lorem ipsum</a>
                <a href="#lorem2">Lorem ipsum</a>
                <a href="#lorem3">Lorem ipsum</a>
                <a href="#lorem4">Lorem ipsum</a>
                <a href="#lorem5">Lorem ipsum</a>
              </ul>

              <h5 id="lorem1">Lorem ipsum dolor sit amet</h5>
              Excepteur et pariatur officia veniam anim culpa cupidatat consequat ad velit culpa est non.
              <ul>
                <li>Nisi qui nisi duis commodo duis reprehenderit consequat velit aliquip.</li>
                <li>Dolor consectetur incididunt in ipsum laborum non et irure pariatur excepteur anim occaecat officia sint.</li>
                <li>Lorem labore proident officia excepteur do.</li>
              </ul>


            </div>
          </section>
		</div>
		<div class="mdl-layout__tab-panel is-active" id="config">
			<section class="section--center mdl-grid mdl-grid--no-spacing mdl-shadow--2dp">
			<div class="mdl-card mdl-cell mdl-cell--12-col">
			<div class="mdl-card__supporting-text">
				<h4>Test</h4>
				<pre>{{ .Config }}</pre>
			</div>
				<div class="mdl-card__actions">
					<a href="#" class="mdl-button">Test</a>
				</div>
			</div>
			<button class="mdl-button mdl-js-button mdl-js-ripple-effect mdl-button--icon" id="btn1">
				<i class="material-icons">more_vert</i>
			</button>
			<ul class="mdl-menu mdl-js-menu mdl-menu--bottom-right" for="btn1">
				<li class="mdl-menu__item">Lorem</li>
				<li class="mdl-menu__item" disabled>Ipsum</li>
				<li class="mdl-menu__item">Dolor</li>
			</ul>
		</section>
		</div>
        <footer class="mdl-mega-footer">
          <div class="mdl-mega-footer--bottom-section">
            <div class="mdl-logo">
              More Information
            </div>
            <ul class="mdl-mega-footer--link-list">
              <li><a href="#">Source</a></li>
              <li><a href="#">App</a></li>
            </ul>
          </div>
        </footer>
      </main>
    </div>
    <a href="https://github.com/google/material-design-lite/blob/mdl-1.x/templates/text-only/" target="_blank" id="go-to-app" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-color--accent mdl-color-text--accent-contrast">Go to app</a>
    <script src="https://code.getmdl.io/1.3.0/material.min.js"></script>

{{end}}`
