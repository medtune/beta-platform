var overlay = $("#overlay"),
        fab = $(".fab"),
     cancel = $("#cancel");
      modal = $("#modal");
    results = $("#results");
      title = $("#title");
        cam = $("#cam");
    preview = $("#preview");
  previewCC = $("#preview-cc");
      camCC = $("#cam-cc");
  camTiming = $("#cam-timing");
       cntt = $("#cntt");
   cnttPrev = $("#cntt-prev");
  cnttPCtrl = $("#cntt-prev-ctrl");
 pageHeader = $("#page-header");
     tabBar = $("#tab-bar");
     
// Set modal title
function _setModalTitle(title) {
    document.getElementById("title").innerHTML = title;
}


cnttPCtrl.hide();
cnttPrev.hide();

// Set Modal results
// Will show grad-cam/preview if a grad cam result exist
// else will only show classification results. if both 
// classification/grad-cam fail, it will log errors and
// not proc
function _setModalResults(image, data, time) {
    //cam.setBackground("url('/static/demos/mura/images" + data.in)
    if (data.cam == null) {
        console.log("hidding cam output");
        previewCC.hide();
        camCC.hide();
        camTiming.hide();
        fab.css('top', '30%');
        cntt.css('padding-bottom', '30px');
    } else {
        console.log("showing preview/cam CC");
        preview.css('background-image', 'url(' + '/static/demos/mura/images/' + image + ')');
        cam.css('background-image', 'url(' + '/static/demos/mura/images/' + data.cam.output + ')');
        fab.css('top', '10%');
        cntt.css('padding-bottom', '0px');
        document.getElementById("grad-cam-time").innerHTML = (data.cam.round_trip/1000000).toFixed(2) + ' ms';
        previewCC.show();
        camCC.show();
        camTiming.show();
    }
    var v1 = (data.inference.scores[0]*100).toFixed(9);
    var v2 = (data.inference.scores[1]*100).toFixed(9);

    document.getElementById("r1").innerHTML = data.inference.keys[0] + "&nbsp;";
    document.getElementById("s1").innerHTML = "&nbsp;" + v1;
    document.getElementById("r2").innerHTML = data.inference.keys[1]; + "&nbsp;"
    document.getElementById("s2").innerHTML = "&nbsp;" + v2;
    document.querySelector('#p1').MaterialProgress.setProgress(v1);
    document.querySelector('#p2').MaterialProgress.setProgress(v2);
    document.getElementById("inference-time").innerHTML = (data.inference.round_trip/1000000).toFixed(2) + ' ms';
    document.getElementById("req-roundtrip").innerHTML = (data.timing/1000000).toFixed(2) + ' ms';

    console.log("time ----", time);
    console.log("inference", data.inference);
}

function _alertFailure(errors) {
    alert(errors);
}

// Erase data at modal view
function _cleanModal() {

}

// Open FAB
function openFAB(event) {
    if (event) event.preventDefault();
    fab.addClass('active');
    overlay.addClass('dark-overlay');
}

// Close FAB
function closeFAB(event) {
    if (event) {
        event.preventDefault();
        event.stopImmediatePropagation();
    };
    fab.removeClass('active');
    overlay.removeClass('dark-overlay');
    cntt.hide();
    cnttPrev.hide();
    cnttPCtrl.hide();
    tabBar.slideDown();
    pageHeader.slideDown();
    _cleanModal();
}

// -------------------------------------

// Open modal pop up to preview image 
function previewCase(image) {
    _setModalTitle("Preview:&nbsp;" + image);
    cntt.hide();
    console.log("preview:", image)
    cnttPrev.html('<canvas id="canvas"></canvas>');
    Caman("#canvas", "/static/demos/mura/images/"+image, function () {
        // operation
        this.render();

    });
    fab.css('top', '10%');
    cnttPrev.show();
    cnttPCtrl.show();
    tabBar.slideUp();
    openFAB();
}

// Query engine type from settings dropdown
function _getEngineType() {
    var m = "";
    _model = document.getElementById('model').value;
    if (_model == "Inception ResNet V2" || _model == "") {
        m = "mura-irn-v2";
    } else if (_model == "MobileNet V2") {
        m = "mura-mn-v2";
    };
    console.log("detected config:", m);
    return m;
}

// Request the server to run cam and inference before
// showing the results in the main modal pop up
function processCase(image) {
    _setModalTitle("Result:&nbsp;" + image);
    var m = _getEngineType();
    data = {
        'target' : image,
        'model_id' : m,
    };
    
    var start = new Date();
    sendJSON(
        'POST',
        '/api/demos/mura/process',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            var r = resp.data;
            var d = r.data;
            if (resp.success == true) {
                console.log('success\n\t', r);
                if (d.errors == null || d.errors.length == 1) {
                    //_alertFailure(d.errors);
                    _setModalResults(image, d, diff(start, end));
                    cntt.show();
                    openFAB();
                } else {
                    console.log("proc failed:", d.errors);
                };
            } else if (resp.success == false) {
                console.log('fail:\n\t', resp);
                _alertFailure(resp.errors);
            };
        });
}

var removeBrightness = $("#remove-brightness");
var addBrightness = $("#add-brightness");
var removeContrast = $("#remove-contrast");
var addContrast = $("#add-contrast");
var removeSaturation = $("#remove-saturation");
var addSaturation = $("#add-saturation");
var removeNoise = $("#remove-noise");
var addNoise = $("#add-noise");
var removeHue = $("#remove-hue");
var addHue = $("#add-hue");
var removeVibrance = $("#remove-vibrance");
var addVibrance = $("#add-vibrance");
var clearBtn = $("#prev-clear");

removeBrightness.click(function() {
    Caman("#canvas", function () {
        // operation
        this.brightness(-5).render();
    });
})

addBrightness.click(function() {
    Caman("#canvas", function () {
        // operation
        this.brightness(5).render();
    });
})

removeContrast.click(function() {
    Caman("#canvas", function () {
        // operation
        this.contrast(-5).render();
    });
})

addContrast.click(function() {
    Caman("#canvas", function () {
        // operation
        this.contrast(5).render();
    });
})

removeSaturation.click(function() {
    Caman("#canvas", function () {
        // operation
        this.saturation(-5).render();
    });
})

addSaturation.click(function() {
    Caman("#canvas", function () {
        // operation
        this.saturation(5).render();
    });
})

removeNoise.click(function() {
    Caman("#canvas", function () {
        // operation
        this.noise(-5).render();
    });
})

addNoise.click(function() {
    Caman("#canvas", function () {
        // operation
        this.noise(5).render();
    });
})

removeHue.click(function() {
    Caman("#canvas", function () {
        // operation
        this.hue(-5).render();
    });
})

addHue.click(function() {
    Caman("#canvas", function () {
        // operation
        this.hue(5).render();
    });
})


removeVibrance.click(function() {
    Caman("#canvas", function () {
        // operation
        this.vibrance(-5).render();
    });
})

addVibrance.click(function() {
    Caman("#canvas", function () {
        // operation
        this.vibrance(5).render();
    });
})

clearBtn.click(function(){
    Caman('#canvas', function() {
        this.revert();
    })
})