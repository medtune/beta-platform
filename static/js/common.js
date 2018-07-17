var baseURL = "/";
var baseApiURL = "/api";

var sendJSON = function(method, url, obj, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            callback(xhr.responseText)
        }
    };
    var data = JSON.stringify(obj);
    xhr.send(data);
};

var setDisplay = function(state, id) {
    if (state == true) {
        document.getElementById(id).style.display = "block";
    } else {
        document.getElementById(id).style.display = "none";
    }
}

