var method = "POST";

var baseURL = "/";
var baseApiURL = "/api";

var loginFormID = "login-form";
var signupFormID = "signup-form";


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

var loadLogin = function() {
    var un = document.getElementById("username").value;
    var pw = document.getElementById("password").value;
    return {
        "username" : un,
        "password" : pw,
    };
};

var setLoadStatus = function(state, id) {
    if (state == true) {
        document.getElementById(id).style.display = "block";
    } else {
        document.getElementById(id).style.display = "none";
    }
}

var Login = function () {
    var obj = loadLogin();
    setLoadStatus(true, "login-load")
    sendJSON("POST", "/login", obj, (res) => {
        var data = JSON.parse(res);
        //FIXME: Remove
        console.log("Login response:", data);
        if (data.success == true) {
            window.location.href = "/home";
        }
        setLoadStatus(false, "login-load");
    })
};

var loadSignup = function() {
    //
    var un = document.getElementById("username").value;
    var pw = document.getElementById("password").value;
    var pwc = document.getElementById("passwordc").value;
    var secret = document.getElementById("secret").value;
    return {
        "username" : un,
        "password" : pw,
        "passwordc" : pwc,
        "secret" : secret,
    };
};



var Signup = function () { 
    var obj = loadSignup();
    setLoadStatus(true, "signup-load")
    sendJSON("POST", "/signup", obj, (res) => {
        var data = JSON.parse(res);
        console.log("Signup response:", data);
        if (data.success == true) {
            window.location.href = "/signup/success";
        }
        setLoadStatus(false, "signup-load");

    })
};
