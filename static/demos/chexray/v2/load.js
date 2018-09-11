var overlay = $("#overlay");
        fab = $(".fab");
     cancel = $("#cancel");
      modal = $("#modal");
    results = $("#results");
      title = $("#title");
        cam = $("#cam");
    preview = $("#preview");
  previewCC = $("#preview-cc");
      camCC = $("#cam-cc");
  camTiming = $("#cam-timing");
    cxpbaCC = $("#cxpba-cc");
 cxpbaTable = $("#cxpba-table");
 tableTbody = $("#table-tbody");
       cntt = $("#cntt");
   cnttPrev = $("#cntt-prev");
  cnttPCtrl = $("#cntt-prev-ctrl");
 pageHeader = $("#page-header");
    tabBar = $("#tab-bar");
         r1 = $("#r1");
         r2 = $("#r2");
         r3 = $("#r3");
         r4 = $("#r4");
         r5 = $("#r5");
         r6 = $("#r6");
         r7 = $("#r7");
         r8 = $("#r8");
         r9 = $("#r9");
        r10 = $("#r10");
        r11 = $("#r11");
        r12 = $("#r12");
        r13 = $("#r13");
        r14 = $("#r14");
        r15 = $("#r15");

cxpbaCC.hide();

function _setCXPBAdata(pathology) {
    console.log("pathology:", pathology);
    data = {
        "pathology" : pathology,
    };
    var start = new Date();
    sendJSON(
        'POST',
        '/api/demos/chexray/cxpba/pathology_al',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            var r = resp.data;
            if (resp.success == true) {
                console.log('success\n\t', r);
                $("#table-title").html(r.pathology);
                for (var k in r.table) {
                    var tag = "";
                    if (r.table[k] == "Elevée") {
                        $("#" + k).html("High");
                        $("#" + k).css("color", "#e53935");
                    } else if (r.table[k] == "Faible") {
                        $("#" + k).html("Low");
                        $("#" + k).css("color", "rgb(63,81,181)");
                    } else {
                        $("#" + k).html(r.table[k]);
                        $("#" + k).css("color", "");
                    }
                };
                
                console.log("setup done.");
                cxpbaCC.slideDown();
            } else if (resp.success == false) {
                console.log('fail:\n\t', resp);
            };
        },
    );
};

r1.click(function() {_setCXPBAdata(r1.html())});
r2.click(function() {_setCXPBAdata(r2.html())});
r3.click(function() {_setCXPBAdata(r3.html())});
r4.click(function() {_setCXPBAdata(r4.html())});
r5.click(function() {_setCXPBAdata(r5.html())});
r6.click(function() {_setCXPBAdata(r6.html())});
r7.click(function() {_setCXPBAdata(r7.html())});
r8.click(function() {_setCXPBAdata(r8.html())});
r9.click(function() {_setCXPBAdata(r9.html())});
r10.click(function() {_setCXPBAdata(r10.html())});
r11.click(function() {_setCXPBAdata(r11.html())});
r12.click(function() {_setCXPBAdata(r12.html())});
r13.click(function() {_setCXPBAdata(r13.html())});
r14.click(function() {_setCXPBAdata(r14.html())});
r15.click(function() {_setCXPBAdata(r15.html())});


console.log("initilized rcct");

// Set modal title
function _setModalTitle(title) {
    document.getElementById("title").innerHTML = title;
}

// Set Modal results
// Will show grad-cam/preview if a grad cam result exist
// else will only show classification results. if both 
// classification/grad-cam fail, it will log errors and
// not proc
function _setModalResults(image, data, time) {
    //cam.setBackground("url('/static/demos/chexray/images" + data.in)

    cnttPCtrl.hide();
    cnttPrev.hide();
    if (data.cam == null) {
        console.log("hidding cam output");
        previewCC.hide();
        camCC.hide();
        camTiming.hide();
        fab.css('top', '10%');
        cntt.css('padding-bottom', '10px');
    } else {
        console.log("showing preview/cam CC");
        preview.css('background-image', 'url(' + '/static/demos/chexray/images/' + image + ')');
        cam.css('background-image', 'url(' + '/static/demos/chexray/images/' + data.cam.output + ')');
        fab.css('top', '10%');
        cntt.css('padding-bottom', '0px');
        document.getElementById("grad-cam-time").innerHTML = (data.cam.round_trip/1000000).toFixed(2) + ' ms';
        previewCC.show();
        camCC.show();
        camTiming.show();
    }
    
    var v1 = (data.inference.scores[0]*100).toFixed(4);
    var v2 = (data.inference.scores[1]*100).toFixed(4);
    var v3 = (data.inference.scores[2]*100).toFixed(4);
    var v4 = (data.inference.scores[3]*100).toFixed(4);
    var v5 = (data.inference.scores[4]*100).toFixed(4);
    var v6 = (data.inference.scores[5]*100).toFixed(4);
    var v7 = (data.inference.scores[6]*100).toFixed(4);
    var v8 = (data.inference.scores[7]*100).toFixed(4);
    var v9 = (data.inference.scores[8]*100).toFixed(4);
    var v10 = (data.inference.scores[9]*100).toFixed(4);
    var v11 = (data.inference.scores[10]*100).toFixed(4);
    var v12 = (data.inference.scores[11]*100).toFixed(4);
    var v13 = (data.inference.scores[12]*100).toFixed(4);
    var v14 = (data.inference.scores[13]*100).toFixed(4);
    var v15 = (data.inference.scores[14]*100).toFixed(4);

    document.getElementById("r1").innerHTML = data.inference.keys[0];
    document.getElementById("s1").innerHTML = "&nbsp;" + v1;
    document.getElementById("r2").innerHTML = data.inference.keys[1];
    document.getElementById("s2").innerHTML = "&nbsp;" + v2;
    document.getElementById("r3").innerHTML = data.inference.keys[2];
    document.getElementById("s3").innerHTML = "&nbsp;" + v3;
    document.getElementById("r4").innerHTML = data.inference.keys[3];
    document.getElementById("s4").innerHTML = "&nbsp;" + v4;
    document.getElementById("r5").innerHTML = data.inference.keys[4];
    document.getElementById("s5").innerHTML = "&nbsp;" + v5;
    document.getElementById("r6").innerHTML = data.inference.keys[5];
    document.getElementById("s6").innerHTML = "&nbsp;" + v6;
    document.getElementById("r7").innerHTML = data.inference.keys[6];
    document.getElementById("s7").innerHTML = "&nbsp;" + v7;
    document.getElementById("r8").innerHTML = data.inference.keys[7];
    document.getElementById("s8").innerHTML = "&nbsp;" + v8;
    document.getElementById("r9").innerHTML = data.inference.keys[8];
    document.getElementById("s9").innerHTML = "&nbsp;" + v9;
    document.getElementById("r10").innerHTML = data.inference.keys[9];
    document.getElementById("s10").innerHTML = "&nbsp;" + v10;
    document.getElementById("r11").innerHTML = data.inference.keys[10];
    document.getElementById("s11").innerHTML = "&nbsp;" + v11;
    document.getElementById("r12").innerHTML = data.inference.keys[11];
    document.getElementById("s12").innerHTML = "&nbsp;" + v12;
    document.getElementById("r13").innerHTML = data.inference.keys[12];
    document.getElementById("s13").innerHTML = "&nbsp;" + v13;
    document.getElementById("r14").innerHTML = data.inference.keys[13];
    document.getElementById("s14").innerHTML = "&nbsp;" + v14;
    document.getElementById("r15").innerHTML = data.inference.keys[14];
    document.getElementById("s15").innerHTML = "&nbsp;" + v15;

    document.querySelector('#p1').MaterialProgress.setProgress(v1);
    document.querySelector('#p2').MaterialProgress.setProgress(v2);
    document.querySelector('#p3').MaterialProgress.setProgress(v3);
    document.querySelector('#p4').MaterialProgress.setProgress(v4);
    document.querySelector('#p5').MaterialProgress.setProgress(v5);
    document.querySelector('#p6').MaterialProgress.setProgress(v6);
    document.querySelector('#p7').MaterialProgress.setProgress(v7);
    document.querySelector('#p8').MaterialProgress.setProgress(v8);
    document.querySelector('#p9').MaterialProgress.setProgress(v9);
    document.querySelector('#p10').MaterialProgress.setProgress(v10);
    document.querySelector('#p11').MaterialProgress.setProgress(v11);
    document.querySelector('#p12').MaterialProgress.setProgress(v12);
    document.querySelector('#p13').MaterialProgress.setProgress(v13);
    document.querySelector('#p14').MaterialProgress.setProgress(v14);
    document.querySelector('#p15').MaterialProgress.setProgress(v15);

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
    cxpbaCC.hide();
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
    _setModalTitle("Preview:&nbsp;"+image);
    cntt.hide();
    console.log("preview:", image)
    cnttPrev.html('<canvas id="canvas"></canvas>');
    Caman("#canvas", "/static/demos/chexray/images/"+image, function () {
        // operation
        this.render();

    });
    fab.css('top', '5%');
    cnttPrev.show();
    cnttPCtrl.show();
    tabBar.slideUp();
    pageHeader.slideUp();
    openFAB();
}

// Query engine type from settings dropdown
function _getEngineType() {
    var m = "";
    _model = document.getElementById('model').value;
    if (_model == "DenseNet 121" || _model == "") {
        m = "chexray-dn-121";
    } else if (_model == "MobileNet V2") {
        m = "chexray-mn-v2";
    };
    console.log("detected config:", m);
    return m;
}

// Request the server to run cam and inference before
// showing the results in the main modal pop up
function processCase(image) {
    _setModalTitle('Result:&nbsp;' + image);
    var m = _getEngineType();
    data = {
        'target' : image,
        'model_id' : m,
    };
    
    var start = new Date();
    sendJSON(
        'POST',
        '/api/demos/chexray/process',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            var r = resp.data;
            console.log("r", r);
            if (resp.success == true) {
                console.log('success\n\t', r);
                if (r.errors == null || r.errors.length == 1) {
                    //_alertFailure(d.errors);
                    _setModalResults(image, r.data, diff(start, end));
                    cntt.show();
                    openFAB();
                } else {
                    console.log("proc failed:", r.errors);
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