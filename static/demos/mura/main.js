var actualSampleId = 0;

var numSamples = 4;

var samples = [
    {
        'id': 0,
        'pbody' : 'elbow',
        'label' : 'positive',
    },
];

var getSample = function(id) {
    return s
};

var _showActualSample = function() {
    var s = getSample(actualSampleId);
};

var runInference = function() {
    
};

var back = function() {
    actualSampleId--;
    actualSampleId = actualSampleId % numSamples;
    _showActualSample();
};

var next = function() {
    actualSampleId++;
    actualSampleId = actualSampleId % numSamples;
    _showActualSample();
};

var information = function() {
    console.log("informations xd lol")
};