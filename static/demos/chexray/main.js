var actualSampleId = 0;
var numSamples = 2;

var samples = [
    {
        'id': 0,
        'label' : 'Positive',
    },
    {
        'id': 1,
        'label' : 'Positive',
    },
];

var getSample = function(id) {
    var s = {};
    for (var x in samples) {
        if (samples[x].id == id) {
            s = samples[x];
            break;
        }
    };
    return s
};

var _setImageInformations = function(id, zone, label) {
    document.getElementById('sample-id').innerHTML = id;
    document.getElementById('sample-label').innerHTML = label;
}

var _setClassPrediction = function(name, time, isCorrect) {
    document.getElementById('pred-zone').innerHTML = name + ' (Timing: ' + time + ' ms)';
    var color = '';
    if (isCorrect === true) {
        color = 'green';
    } else {
        color = 'red';
    }
    document.getElementById('pred-zone').style.backgroundColor = color;
}

var _initClassPrediction = function() {
    document.getElementById('pred-zone').innerHTML = 'Undefined (Timing: NaN ms)';
    document.getElementById('pred-zone').style.backgroundColor = 'grey';

};

var _setMainImage = function(id) {
    document.getElementById('main-image').style.backgroundImage = "url('/static/demos/chexray/images/image" + id + ".png')"
}

var _showActualSample = function() {
    var s = getSample(actualSampleId);
    _setMainImage(s.id);
    _setImageInformations(s.id, s.pbody, s.label);
};

var _alertFailure = function() {

};

var runInference = function() {
    data = {
        'id' : actualSampleId,
    };
    var start = new Date();
    sendJSON(
        'POST',
        '/api/chexray/run_inference',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            if (resp.success == true) {
                console.log('success /!/', resp.data);
                _setClassPrediction(resp.data.prediction, diff(start, end), resp.data.correct);
            } else if (resp.success == false) {
                console.log('fail /!/', resp);
                _alertFailure();
            }
        });
};

var back = function() {
    actualSampleId--;
    actualSampleId = Math.pow(actualSampleId % numSamples, 2);
    console.log(actualSampleId);
    _showActualSample();
    _initClassPrediction();
};

var next = function() {
    actualSampleId++;
    actualSampleId = actualSampleId % numSamples;
    console.log(actualSampleId);
    _showActualSample();
    _initClassPrediction();
};

var information = function() {
    console.log("informations xd lol")
};


