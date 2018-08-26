var loadLogin = function() {
    var un = document.getElementById("username").value;
    var pw = document.getElementById("password").value;
    return {
        "username" : un,
        "password" : pw,
    };
};

var setInputState = function(id, state) {
    if (state == 1) {
        document.getElementById(id).classList.add('is-invalid');
    } else if (state == 0) {
        document.getElementById(id).classList.remove('is-invalid');
    };
};


var loginCorrect = function () {
    var un = document.getElementById('fti-username').classList.contains('is-invalid');
    var pw = document.getElementById('fti-password').classList.contains('is-invalid');
    return un && pw 
}

var setErrorMsg = function(id, msg) {
    document.getElementById(id).innerHTML = msg;
};

var setErrors = function(errors) {
    for (var i=0; i<errors.length; i++) {
        var err = errors[i];
        
        if (err.includes('username')) {
            setErrorMsg('username-error', err);
            setInputState('fti-username', 1);
        } else if (err.includes('password')) {
            setErrorMsg('password-error', err);
            setInputState('fti-password', 1)
        } 
    }
};

var Login = function () {
    var obj = loadLogin();
    setDisplay(true, "login-load")
    sendJSON("POST", "/api/login", obj, (res) => {
        var data = JSON.parse(res);

        if (data.success == true) {
            setTimeout(function(){
                setDisplay(false, 'login-load');
                window.location.href = '/home';
            }, 500)
        } else {
            //console.log(data.errors);
            setTimeout(function() {
                setErrors(data.errors);
                setDisplay(false, 'login-load');
            }, 500)
        };
    })
};
