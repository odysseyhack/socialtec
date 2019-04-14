
var Data;
function item_liked(x) {
  var id = "#card-"+x;
  console.log(id);
  $(id).remove();
  for(var cntr = 0;cntr<Data.length; cntr++){
    console.log(Data[cntr]["id"]);
    if(Data[cntr]["id"]==x){
      var intrest = '<li id="intrest-'+x+'" onclick="delete_intrest(this.id)" class="filled" style="background-image: url('+Data[cntr]['image_url']+');">+</li>';
      $( "#my-interest" ).append(intrest);
      break;
    }
  }
  $.ajax({
    url: URL+'/interest/'+x,
    type: 'PUT',
    data: Data[cntr]
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}

(function($) {

    // Document Ready
    $(document).ready(function() {

      $.get( URL+"/offers", function(data, status){
        Data = data;
        for(var cntr = 0;cntr<data.length; cntr++){
          var card = '<div class="offer swiper-slide" id="card-'+data[cntr]["id"]+'"><div class="offer__content" style="background-image: url('+data[cntr]['image_url']+');"><div class="offer__body"><h3 class="offer__title">'+data[cntr]['name']+'</h3><small class="offer__user">by Elit Numquam</small><p>'+data[cntr]['details']+'</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul><button class="btn-interested" id="'+data[cntr]["id"]+'"  onclick="item_liked(this.id)"><span>+</span></button></div></div></div>'
          $( ".swiper-wrapper" ).append(card);
        }

      },"json");


    });

    var swiper = new Swiper('.swiper-container', {
      slidesPerView: 'auto',
      centeredSlides: true,
      spaceBetween: 0,
      grabCursor: true,
    });
})(jQuery);
