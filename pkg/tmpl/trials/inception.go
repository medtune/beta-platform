package trials

const Inception = `{{define "base"}}
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="Amine.in website">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0">
    <title> Demo-v0.0.1 </title>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:regular,bold,italic,thin,light,bolditalic,black,medium&amp;lang=en">
    <link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-blue.min.css" />
    <link rel="stylesheet" href="/static/css/styles.css" />
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <script src="/static/js/typed.min.js"></script>
</head>

<body>
    <div class="mdl-layout--fixed-header mdl-layout mdl-js-layout">
        <header class="mdl-layout__header">
            <div class="mdl-layout__header-row">
                <span class="mdl-layout-title">v0.0.2_dev</span>
                <div class="mdl-layout-spacer"></div>

                <button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon">
                    <i class="material-icons">cloud_upload</i>
                </button>
                <button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon">
                    <i class="material-icons">settings_input_svideo</i>
                </button>
                <button class="mdl-button mdl-js-button mdl-button--icon nav-button-icon" id="button_settings">
                    <i class="material-icons">settings</i>
                </button>

                <ul class="mdl-menu mdl-menu--bottom-right mdl-js-menu mdl-js-ripple-effect" for="button_settings">
                    <li class="mdl-menu__item">Settings</li>
                    <li class="mdl-menu__item mdl-menu__item--full-bleed-divider">Endpoints</li>
                    <li disabled class="mdl-menu__item">API</li>
                    <li class="mdl-menu__item">Shutdown</li>
                </ul>
            </div>
    
        </header>
        <div class="mdl-layout__drawer">
                <span class="version">v 0.0.2</span>
            <div style="text-align:center">
                <img src="/static/images/logo.png" class="demo-avatar" onclick="window.location.href='/index'">
            <div>
            <nav class="mdl-navigation">
                <div class="drawer-separator"></div>
                <a class="mdl-navigation__link drawer-nav-link" href="">Tryouts</a>
                <a class="mdl-navigation__link drawer-nav-link" href="">Slides</a>
                <a class="mdl-navigation__link drawer-nav-link" href="">About</a>
                <div class="drawer-separator"></div>
                <a class="mdl-navigation__link drawer-nav-link" href="">Uploads</a>
                <a class="mdl-navigation__link drawer-nav-link" href="">Capsules</a>
                <a class="mdl-navigation__link drawer-nav-link" href="">Settings</a>
                <div class="drawer-separator"></div>
            </nav>
        </div>
    </div>
    </div>
        <main class="mdl-layout__content">
            <div class="page-content"><!-- Your content goes here -->

            </div>
        </main>
    </div>

      <script src="https://code.getmdl.io/1.3.0/material.min.js"></script>

</body>

</html>
{{end}}`
