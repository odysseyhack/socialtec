var URL = "http://localhost:3333"

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
    $(document).ready(function() {});


    // Window Load
    $(window).load(function() {
    });

})(jQuery);
