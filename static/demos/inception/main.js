var incr = 0;
var defaultPreds = 5;

var _loadingStatus = function(status, id) {
    if (status == true) {
        document.getElementById(id).style.display = "block";
    } else {
        document.getElementById(id).style.display = "none";
    }
}

var _getSettings = function() {
    var numpreds = document.getElementById("numofpreds").value;
    if (numpreds == 0) {
        numpreds = 5;
    };
    return {
        'numpreds': numpreds,
        'model': 'VGG-16',
    };
};

var _setPredsUI = function(num) {
    console.log(num);
    for (var i = 1; i <=10 ; i++) {
        console.log('s' + i)
        if (i <= num ) {
            document.getElementById('result-' + i).style.display = 'block';
        } else {
            document.getElementById('result-' + i).style.display = 'none';
        }
    }
    console.log("hehehe");
}

var _createResultStatOne = function(tags, perc, id) {
    document.getElementById('s' + id).innerHTML = tags;
    document.getElementById('r' + id).innerHTML = perc + '%:';
    document.querySelector('#p' + id).MaterialProgress.setProgress(perc);
};

var _alertFailure = function() {
    console.log("failure protocol");
};

var _showTime = function(t) {
    document.getElementById('elapsed-time').innerHTML = 'Elapsed time: ' + t * 1000 + 'ms';
};

var _requestInference = function(image, numpreds) {
    data = {
        'file' : '.' + image,
        'numpreds' : numpreds,
    };
    var start = new Date();
    sendJSON(
        'POST',
        '/api/inception_imagenet/run_inference',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            if (resp.success == true) {
                console.log('success //!!', resp.data);
                _showResults(resp.data);
                _showTime(diff(start, end));
            } else if (resp.success == false) {
                console.log('fail //!!', resp);
                _alertFailure();
            }
        });
};

var _deleteOldResults = function() {
    incr = 0;
    resultDiv = document.getElementById('predictions-results');
    resultDiv.innerHTML = ``
};

var _showResults = function(res) {
    var i = 0;
    console.log("***called***");
    for (var key in res.keys) {
        console.log("going on key: ", res.keys[i], res.scores[i], i+1)
        _createResultStatOne(res.keys[i], res.scores[i], i+1)
        i++;
    };
};

var run = function(image) {
    settings = _getSettings();
    _setPredsUI(settings.numpreds);
    _requestInference(image, settings.numpreds);
};

var drop = function(image) {
    data = {
        'file' : image,
    };
    sendJSON(
        'POST',
        '/api/inception_imagenet/drop_image',
        data,
        (res) => {
            resp = JSON.parse(res);
            console.log(resp);
            if (resp.success == true) {
                console.log("droping ok");
                //alert("-----");
                window.location.href = '/demos/inception_imagenet';
            } else {
                console.log("failed to drop file:" + image)
            };
        },
    );
};


