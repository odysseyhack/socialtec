
var Data;
var cntr;

function item_liked(idx) {
  offer= Data[idx];
  $("#card-"+offer.offer.id).remove();
  addInterestButton(offer.offer)
  console.log(offer.offer)
  if(offer.interest != null ){
    initiateCycle(offer)
    return
  }
  sendInterest(offer)
}


function sendInterest(offer) {
  $.ajax({
    url: URL+'/interest/'+offer.offer.id,
    type: 'PUT',
    data: JSON.stringify(offer.offer),
    dataType: 'json',
    contentType: 'application/json',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}


function initiateCycle(offer) {
$.ajax({
    url: URL+'/cycle',
    type: 'POST',
    data: JSON.stringify(offer),
    dataType: 'json',
    contentType: 'application/json',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });

loadOffers();
  

}





function addInterestButton(offer) {
  var intrest = '<li id="intrest-'+offer.id+'" onclick="delete_intrest(' +offer.id+ ')" class="filled" style="background-image: url('+offer['image_url']+');">+</li>';
  $( "#my-interest" ).append(intrest);  
}


function loadInterests() {
$.ajax({
    url: URL+'/me/interests',
    type: 'GET',    
    crossDomain: true,
    success: function(result) {
      if (result != null ){
      result.forEach(addInterestButton);
      }
    }
  });
}


function delete_intrest(id) {
  $('#intrest-'+id).remove();
  $.ajax({
    url: URL+'/interest/'+id,
    type: 'DELETE',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}


function loadOffers() {
  console.log(URL+"/offers")
  $.get(URL+"/offers", function(data, status){
        Data = data;
        console.log(data)
        for( cntr = 0;cntr<data.length; cntr++){
          offer = data[cntr].offer
          var card = '<div class="offer swiper-slide" id="card-'+offer.id+'"><div class="offer__content" style="background-image: url('+offer['image_url']+');"><div class="offer__body"><h3 class="offer__title">'+offer['name']+'-'+offer.id+'</h3><small class="offer__user">by Elit Numquam</small><p>'+offer['details']+'</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul><button class="btn-interested" data-counter="'+ cntr +'" id="'+offer["id"]+'"  onclick="item_liked('+cntr+')"><span>+</span></button></div></div></div>'
          $( ".swiper-wrapper" ).append(card);
        }

      },"json");
}


var prevOffer = null
function slideChange() {
  console.log($(".offer .swiper-slide .swiper-slide-prev"))
}


(function($) {
    // Document Ready
    $(document).ready(function() {
      loadOffers()
    loadInterests()

    });

    var swiper = new Swiper('.swiper-container', {
      slidesPerView: 'auto',
      centeredSlides: true,
      spaceBetween: 0,
      grabCursor: true,
    });

    swiper.on('slideChange', function () {
      slideChange()
  });
})(jQuery);
