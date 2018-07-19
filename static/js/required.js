var scrollTo = function(top) {
    var content = $(".mdl-layout__content");
    var target = top ? 0 : $(".page-content").height();
    content.stop().animate({ scrollTop: target }, "fast");
};
  
var scrollToTop = function() {
    scrollTo(true);
};

var scrollToBottom = function() {
    scrollTo(false);
};

var scrollToWear = function() {
    console.log("hiiii");
    var content = $(".mdl-layout__content");
    var target = top ? 0 : $("#try-beta").height();
    content.stop().animate({ scrollTop: target }, "fast");
}