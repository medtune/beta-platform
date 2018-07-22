var el = document.getElementById('c');
var ctx = el.getContext('2d');
var isDrawing;

var getMousePos = function(canvas, evt) {
    var rect = canvas.getBoundingClientRect();
    return {
      x: evt.clientX - rect.left,
      y: evt.clientY - rect.top
    };
}

var clearCanvas = function() {
    ctx.clearRect(0, 0, el.width, el.height);
    var demo = document.getElementById("demo");
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
    el.onmousedown = function(e) {
        isDrawing = true;
        ctx.lineWidth = 15;
        ctx.lineJoin = ctx.lineCap = 'round';
        ctx.shadowBlur = 5;
        ctx.shadowColor = 'rgb(255, 255, 255)';
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
    console.log(el.toDataURL("image/png"));
}

