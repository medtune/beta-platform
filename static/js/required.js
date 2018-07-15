//TODO Scroll

/*
 * Important things to note-
 *  the usual 'html, body' is replaced with mdl-layout__content
 *  the height of the document is now the page content's height
 */
var scrollTo = function(top) {
    var content = $(".mdl-layout__content");
    var target = top ? 0 : $(".page-content").height();
    content.stop().animate({ scrollTop: target }, "slow");
};
  
var scrollToTop = function() {
    scrollTo(true);
};

var scrollToBottom = function() {
    scrollTo(false);
};
