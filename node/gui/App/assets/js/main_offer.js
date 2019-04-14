var URL = "http://localhost:3333"

function load_my_offers(){
  var offer_element = '<div class="my-offer"><div class="my-offer__content"><div class="my-offer__image" style="background-image: url(https://source.unsplash.com/700x700/?product);"></div><div class="my-offer__body"><h3 class="my-offer__title">Something</h3><p>Provident, nam? Lorem ipsum dolor sit amet adipisicing elit.</p><a class="btn-remove-my-offer"><span>X</span></a></div></div></div>'
    $( "#my-offers" ).append(offer_element);
}

function add_offer(){
  console.log("offer added");
  var offer_data = {
    "name": $('#offer-title').val(),
    "details": $('#offer-details').val(),
    "image_url": "https://via.placeholder.com/150",
  }
  console.log(offer_data);
  $.post(URL+"/offers",
  offer_data,
  function(data, status){
    console.log("Data: " + data + "\nStatus: " + status);
  });

}

(function($) {

    // Document Ready
    $(document).ready(function() {
      load_my_offers();
    });

})(jQuery);
