
var Data = [];
var cntr;


function item_disliked(idx) {
  offer=Data[idx];
  $("#card-"+offer.offer.id).remove();

  if (offer.interest == null) {
    return
  }

  if(idx >= Data.length) {
    if (idx.interest != null) {
     sendRefer(offer)
  }
  return
}

  next = Data[idx+1]
  if (next.interest != null && next.interest.who != offer.interest.who) {
    sendRefer(offer)
  }
  return
}



function sendRefer(offer) {
  $.ajax({
    url: URL+'/dontLikeAny/'+offer.interest.who,
    type: 'POST',
    dataType: 'json',
    contentType: 'application/json',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}

function Popup(offer,myOfferId) {
  $("#model-offer-name").text(offer.offer.name)
  $("#model-offer-details").text(offer.offer.details)
  $("#match-found").show()
}

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
  Popup(offer)
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
          var card = '<div class="offer swiper-slide" id="card-'+offer.id+'"><div class="offer__content" style="background-image: url('+offer['image_url']+');"><div class="offer__body"><h3 class="offer__title">'+offer['name']+'-'+offer.id+'</h3><small class="offer__user">by Elit Numquam</small><p>'+offer['details']+'</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul>'+
          '<button class="btn-interested" data-counter="'+ cntr +'" id="'+offer["id"]+'"  onclick="item_liked('+cntr+')"><span>+</span></button>'+
          '<button class="btn-not-interested"  data-counter="'+ cntr +'" id="'+offer["id"]+'"  onclick="item_disliked('+cntr+')"><span>&times;</span></button></div></div></div>'
          $( ".swiper-wrapper" ).append(card);
        }

      },"json");
}


var prevOffer = null
function slideChange() {
  console.log($(".offer .swiper-slide .swiper-slide-prev"))
}

function add_card(id, name, details, image_url) {
  card_data = {
    name: name,
    details: details,
    image_url: image_url
  }
  Data[id]=card_data;
  var card = '<div class="offer swiper-slide" id="card-'+id+'"><div class="offer__content" style="background-image: url('+image_url+');"><div class="offer__body"><h3 class="offer__title">'+name+'-'+id+'</h3><small class="offer__user">by Elit Numquam</small><p>'+details+'</p></div><div class="offer__meta"><ul><li><strong>4</strong>Likes</li><li><strong>11</strong>Trades</li><li><strong>680</strong>Respect</li></ul>'+
  '<button class="btn-interested" data-counter="'+ id +'" id="'+id+'"  onclick="item_liked('+id+')"><span>+</span></button>'+
  '<button class="btn-not-interested"  data-counter="'+ id +'" id="'+id+'"  onclick="item_disliked('+id+')"><span>&times;</span></button></div></div></div>'
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
      App.init(on_contract_loaded);
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
