(function($) {

    // Document Ready
    $(document).ready(function() {

        $(".offers").on("touchstart", function(event){
                var xClick = event.originalEvent.touches[0].pageX;
            $(this).one("touchmove", function(event){
                var xMove = event.originalEvent.touches[0].pageX;
                if( Math.floor(xClick - xMove) > 5 ){
                    $(this).carousel('next');
                }
                else if( Math.floor(xClick - xMove) < -5 ){
                    $(this).carousel('prev');
                }
            });
            $(".offers").on("touchend", function(){
                    $(this).off("touchmove");
            });
        });

    });


    // Window Load
    $(window).load(function() {
    });

})(jQuery); 