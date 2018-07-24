var el = document.getElementById('c');
var ctx = el.getContext('2d');
var isDrawing;
var imageName = "";
var basesrc = '/static/demos/mnist/images/';

var getMousePos = function(canvas, evt) {
    var rect = canvas.getBoundingClientRect();
    return {
      x: evt.clientX - rect.left,
      y: evt.clientY - rect.top
    };
}

var clearCanvas = function() {
    imageName = "";
    ctx.clearRect(0, 0, el.width, el.height);
    var demo = document.getElementById('demo');
    demo.innerHTML = '';
    var canvas = document.createElement('canvas');
    canvas.id = "c";
    canvas.width = 500;
    canvas.height = 500;
    canvas.style.border = "1px solid";
    demo.appendChild(canvas);
    el = canvas;
    ctx = el.getContext('2d');
    setupCanvas(el);
}

var setupCanvas = function(el) {
    ctx.fillStyle = "black";
    ctx.fillRect(0, 0, el.width, el.height);
    el.onmousedown = function(e) {
        isDrawing = true;
        ctx.lineWidth = 25;
        ctx.lineJoin = ctx.lineCap = 'round';
        ctx.strokeStyle = 'white';
        pos = getMousePos(el, e);
        ctx.moveTo(pos.x, pos.y);
      };
      el.onmousemove = function(e) {
        if (isDrawing) {
          ctx.strokeStyle = 'white';
          pos = getMousePos(el, e);
          ctx.lineTo(pos.x, pos.y);
          ctx.stroke();
        }
      };
      el.onmouseup = function() {
        isDrawing = false;
      };
}

setupCanvas(el);

var saveAndRun = function() {
    dataURL = el.toDataURL('image/png');
    //console.log(dataURL);
    data = {'image': null, 'file': null};
    if (imageName != "") {
        data['file'] = imageName;
    } else {
        data['image'] = dataURL;
    }
    console.log(data);
    var start = new Date();
    sendJSON(
        'POST',
        '/api/mnist/run_inference',
        data,
        (res) => {
            resp = JSON.parse(res);
            console.log(resp);
            var end = new Date();
            if (resp.success) {
                var x = getMax(resp.data.scores.float_val);
                document.getElementById('elapsed-time').innerHTML = "Elapsed time: " + diff(start, end) + "s";
                document.getElementById('top-prediction').innerHTML = "Top prediction: " + x.maxi + '(scalar : ' + x.max + ')';
            } else if (resp.success == false) {
                document.getElementById('elapsed-time').innerHTML = "Elapsed time:" + diff(start, end) + "s";
                document.getElementById("top-prediction").innerHTML = "Top prediction: ERROR";

            }
        });
}

var drawImage = function(i) {
    imageName = i + '.png';
    var img = new Image();
    img.src = basesrc + i + '.png';
    console.log(img.src, imageName);
    ctx.drawImage(img, 0, 0, 500, 500);
};

var shuffle = function() {
    clearCanvas();
    drawImage(Math.floor((Math.random() * 10)));
}

var draw = function() {
    clearCanvas();
}