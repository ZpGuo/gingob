$(document).ready(function() {
    /**
     * 代码高亮highlight
     */
    $('pre code').each(function(i, block) {
        hljs.highlightBlock(block);
    });

    /**
     * 科学公式TeX(KaTeX)
     */
    // $("#post-toc").find(".a").each(function(){
    //     var tex = $(this);
    //     katex.render(tex.text(), tex[0]);

    //     tex.find(".katex").css("font-size", "1.6em");
    // });

    /* 回到顶部 */
    var $backToTop = $(".bottom-tools");
    /* 隐藏回顶部按钮 */
    $backToTop.transition('hide');
    $(window).on('scroll', function() {
        if ($(this).scrollTop() > 200) {
            $backToTop.transition('show');
        } else {
            $backToTop.transition('hide');
        }
    });
    $backToTop.on('click', function(e) {
        $("html, body").animate({scrollTop: 0}, 500);
    });
});
