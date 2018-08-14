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
    if (state === true) {
        document.getElementById(id).style.display = "block";
    } else {
        document.getElementById(id).style.display = "none";
    }
}

var diff = function(t1, t2) {
    var dif = t1.getTime() - t2.getTime();
    var Seconds_from_T1_to_T2 = dif / 1000;
    var Seconds_Between_Dates = Math.abs(Seconds_from_T1_to_T2);
    return Seconds_Between_Dates;
}

var getMax = function(l) {
    var maxi = 0;
    var max = l[0];
    for (var i=0; i<10; i++) {
        if (l[i] > max){
            maxi = i;
            max = l[i];
        } 
    }
    console.log(maxi, max);
    var d = {
        'maxi' : maxi,
        'max' : max,
    };
    return d
};