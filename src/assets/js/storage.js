Storage = {
  set_interest: function(id){
    key = "my_interest";
    my_interest = [];
    if(localStorage.getItem(key) != null){
      my_interest = JSON.parse(localStorage.getItem(key));
    }
    my_interest.push(id);
    localStorage.setItem(key, JSON.stringify(my_interest));
  },

  get_my_interests: function(){
    key = "my_interest";
    my_interests = JSON.parse(localStorage.getItem(key));
    if(my_interests == null){
      return [];
    }
    return my_interests;
  },

  set_id: function(id) {
    key = "my_id";
    localStorage.setItem(key, id);
  },

  my_id: function() {
    key = "my_id";
    id = my_interest = localStorage.getItem(key);
    if(id == null){
      return 0;
    }
    return id;
  },

  get_last_block: function(event) {
    block_number = my_interest = localStorage.getItem(event);
    if(block_number == null || block_number == "NaN"){
      return 0;
    }
    return block_number;
  },

  set_last_block: function(event, block_number) {
    localStorage.setItem(event, block_number);
  },

  set_offer: function(offer){
    key = "my_offers";
    my_offers = [];
    if(localStorage.getItem(key) != null){
      my_offers = JSON.parse(localStorage.getItem(key));
    }
    my_offers.push(offer);
    localStorage.setItem(key, JSON.stringify(my_offers));
  },

  get_my_offers: function(){
    key = "my_offers";
    my_offers = JSON.parse(localStorage.getItem(key));
    if(my_offers == null){
      return [];
    }
    return my_offers;
  },

  reset_memory: function() {
    localStorage.removeItem("my_interest");
    localStorage.removeItem("my_id");
    lacalStorage.removeItem("refer");
  }

};
