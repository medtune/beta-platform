function doForModels(callback) {
    var start = new Date();
    sendJSON(
        'GET',
        '/api/capsul/list',
        {},
        (res) => {
            resp = JSON.parse(res);
            callback(resp);
        }
    )
}

function updateModelStatus(modelid) {
    $("#referesh-" + modelid).hide();
    $("#load-" + modelid).show();

    data = {
        'model_id' : modelid,
    };
    
    var start = new Date();
    sendJSON(
        'POST',
        '/api/capsul/status',
        data,
        (res) => {
            resp = JSON.parse(res);
            var end = new Date();
            if (resp.success == true) {
                console.log('success /!/', resp.data);
                if (resp.data.status == "AVAILABLE") {
                    setTimeout(function(){
                        $("#status-"+modelid).html("done_outline");
                        $("#status-"+modelid).css("color", "green");
                        $("#referesh-" + modelid).show();
                        $("#load-" + modelid).hide();
                    }, 500);
                } else {
                    setTimeout(function(){
                        $("#status-"+modelid).html("cancel");
                        $("#status-"+modelid).css("color", "red");
                        $("#referesh-" + modelid).show();
                        $("#load-" + modelid).hide();
                    }, 1000);
                };
            } else if (resp.success == false) {
                console.log('fail /!/', resp);
                setTimeout(function(){
                    $("#status-"+modelid).html("cancel");
                    $("#status-"+modelid).css("color", "red");
                    $("#referesh-" + modelid).show();
                    $("#load-" + modelid).hide();
                }, 1500);
            };
        });
}

function runGlobalHealthCheck(req) {
    doForModels(function(resp) {
        if (resp.success) {
            for (var i in resp.data) {
                updateModelStatus(resp.data[i].Name);
            }
        } else {
            console.log("error", resp)
        }
    });
}

function runGlobalBenchmarks() {

}

function runGlobalTests() {
    
}
