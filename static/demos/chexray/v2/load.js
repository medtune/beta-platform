var overlay = $("#overlay"),
        fab = $(".fab"),
     cancel = $("#cancel");
      modal = $("#modal");
    results = $("#results");
      title = $("#title");
        cam = $("#cam");
    preview = $("#preview");


function _setModalTitle(title) {
    document.getElementById("title").innerHTML = title;
}

function _setModalResults(image, data, time) {
    //cam.setBackground("url('/static/demos/chexray/images" + data.in)
    preview.css('background-image', 'url(' + '/static/demos/chexray/images/' + image + ')');
    cam.css('background-image', 'url(' + '/static/demos/chexray/images/' + data.cam.output + ')');
    var v1 = (data.inference.scores[0]*100).toFixed(9);
    var v2 = (data.inference.scores[1]*100).toFixed(9);
    document.getElementById("r1").innerHTML = data.inference.keys[0] + "&nbsp;";
    document.getElementById("s1").innerHTML = "&nbsp;" + v1;
    document.getElementById("r2").innerHTML = data.inference.keys[1]; + "&nbsp;"
    document.getElementById("s2").innerHTML = "&nbsp;" + v2;
    document.querySelector('#p1').MaterialProgress.setProgress(v1);
    document.querySelector('#p2').MaterialProgress.setProgress(v2);
    document.getElementById("inference-time").innerHTML = (data.inference.round_trip/1000000).toFixed(2) + ' ms';
    document.getElementById("grad-cam-time").innerHTML = (data.cam.round_trip/1000000).toFixed(2) + ' ms';
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
    }
    fab.removeClass('active');
    overlay.removeClass('dark-overlay');
    _cleanModal();
}

// -------------------------------------

// Open modal pop up to preview image 
function previewCase(image) {
    _setModalTitle("Preview");
    _setModalPreview(image);
}

// Request the server to run cam and inference before
// showing the results in the main modal pop up
function processCase(image) {
    _setModalTitle("Result 1");
    data = {
        'target' : image,
        'model_id' : 'chexray-mn-v2',
    };
    var start = new Date();
    sendJSON(
        'POST',
        '/api/demos/chexray/process',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            if (resp.success == true) {
                console.log('success\n\t', resp.data);
                if (resp.data.success == true) {
                    _setModalResults(image, resp.data.data, diff(start, end));
                    openFAB();
                } else {
                    console.log("not enough data to proc");
                };
            } else if (resp.success == false) {
                console.log('fail:\n\t', resp);
                _alertFailure(resp.errors);
            }
        });
}