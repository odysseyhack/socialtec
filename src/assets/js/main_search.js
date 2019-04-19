var URL = "http://localhost:3333"

function item_liked(x) {
  var id = "#card-" + x;
  console.log(id);
  $(id).remove();
  $.ajax({
    url: URL + '/interest/' + x,
    type: 'PUT',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}

function add_card(id, name, details, image_url) {
  var card = '<div class="offer swiper-slide" id="card-' + id + '"><div class="offer__content" style="background-image: url(' + image_url + ');"><div class="offer__body"><h3 class="offer__title">' + name + '</h3><small class="offer__user">by Elit Numquam</small><p>' + details + '</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul><button class="btn-interested" id="' + id + '"  onclick="item_liked(this.id)"><span>+</span></button>	<button class="btn-not-interested"><span>&times;</span></button></div></div></div>'
  $(".swiper-wrapper").append(card);
}

function get_offers(id) {
  options = {
    from: App.address
  }
  App.contract.all_goods(id,
    options,
    function(err, result) {
    if (!err) {
      add_card(id, result[3], result[5], result[4]);
    } else {
      console.log(err);
    }
  });
}

function populate_offers() {
  options = {
    from: App.address
  }
  App.contract.max_all_good(
    options,
    function(err, result) {
    if (!err) {
      //console.log(result)
      for(var id = 0; id<=result; id++){
        get_offers(id);
      }
    } else {
      console.log(err);
    }
  });
}

function on_contract_loaded() {
  populate_offers();
}

(function($) {

  // Document Ready
  $(document).ready(function() {

  });

  var swiper = new Swiper('.swiper-container', {
    slidesPerView: 'auto',
    centeredSlides: true,
    spaceBetween: 0,
    grabCursor: true,
  });

})(jQuery);

window.addEventListener('load', function() {
  App.init(on_contract_loaded);
  //App.get_balance();
})
