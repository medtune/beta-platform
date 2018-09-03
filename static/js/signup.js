var loadSignup = function() {
    //
    var email = document.getElementById('email').value;
    var un = document.getElementById('username').value;
    var pw = document.getElementById('password').value;
    var pwc = document.getElementById('passwordc').value;
    var secret = document.getElementById('secret').value;
    return {
        'email' : email,
        'username' : un,
        'password' : pw,
        'passwordc' : pwc,
        'secret' : secret,
    };
};

var setInputState = function(id, state) {
    if (state == 1) {
        document.getElementById(id).classList.add('is-invalid');
    } else if (state == 0) {
        document.getElementById(id).classList.remove('is-invalid');
    };
};

var setErrorMsg = function(id, msg) {
    document.getElementById(id).innerHTML = msg;
};

var signupCorrect = function () {
    var email = document.getElementById('fti-error').classList.contains('is-invalid');
    var un = document.getElementById('fti-username').classList.contains('is-invalid');
    var pw = document.getElementById('fti-password').classList.contains('is-invalid');
    var pwc = document.getElementById('fti-passwordc').classList.contains('is-invalid');
    var secret = document.getElementById('fti-email').classList.contains('is-invalid');
    return email && un && pw && pwc && secret;
}

var setErrors = function(errors) {
    for (var i=0; i<errors.length; i++) {
        var err = errors[i];

        if (err.includes('secret')) {
            setErrorMsg('secret-error', err);
            setInputState('fti-secret', 1);
        } else if (err.includes('username')) {
            setErrorMsg('username-error', err);
            setInputState('fti-username', 1);
        } else if (err.includes('password c')) {
            setErrorMsg('passwordc-error', err);
            setInputState('fti-passwordc', 1);
        } else if (err.includes('password')) {
            setErrorMsg('password-error', err);
            setInputState('fti-password', 1)
        } else if (err.includes('email', err)) {
            setErrorMsg('email-error');
            setInputState('fti-email', 1);
        };
    }
};

var Signup = function () { 
    var obj = loadSignup();
    setDisplay(true, 'signup-load');
    sendJSON('POST', '/api/signup', obj, (res) => {
        var data = JSON.parse(res);
        //console.log('Signup response:', data);
        if (data.success == true) {
            setTimeout(function(){
                setDisplay(false, 'signup-load');
                window.location.href = '/signup/success';
            }, 500);
        } else {
            //console.log(data.errors);
            setTimeout(function() {
                setErrors(data.errors);
                setDisplay(false, 'signup-load');
            }, 500);
        };
    });
};
