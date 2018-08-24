var overlay = $("#overlay"),
        fab = $(".fab"),
     cancel = $("#cancel");

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
}

function _setModalTitle(title) {

}

function _setModalContent(content) {

}

function _setModalPreview(image) {

}

function _setModalResult(inference, image, cam) {

}

// Erase data at modal view
function _cleanModal() {

}


// Run inference only
function _runInference(image) {
    console.log("not implemented");
}

// Run cam only
function _runCam(image) {
    console.log("not implemented");
}

// Run inference and cam using custom exection api ep
function _runInferenceAndCam(image) {

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

}