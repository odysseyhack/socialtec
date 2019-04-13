(function($) {

    // Document Ready
    $(document).ready(function() {

      $.get( "http://6ec69753.ngrok.io/offers", function(data, status){

        for(var cntr = 0;cntr<data.length; cntr++){
          var card = '<div class="offer swiper-slide"><div class="offer__content" style="background-image: url('+data[cntr]['image_url']+');"><div class="offer__body"><h3 class="offer__title">'+data[cntr]['name']+'</h3><small class="offer__user">by Elit Numquam</small><p>'+data[cntr]['details']+'</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul></div></div></div>'
          $( ".swiper-wrapper" ).append(card);
        }

      },"json");

      var swiper = new Swiper('.swiper-container', {
        slidesPerView: 'auto',
        centeredSlides: true,
        spaceBetween: 0,
        grabCursor: true,
      });

    });

})(jQuery);
