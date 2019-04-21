/*jshint esversion: 8 */
Events = {
  listen_refer: async function() {
    filter = {to: "2"};
    additional_filter = {fromBlock: Storage.get_last_block("refer"), toBlock: 'latest'};
    refer_event = App.contract.refer(filter, additional_filter);
    refer_event.watch(function(error, result){
     if (!error){
       console.log("Refer Event:");
       console.log(result);
       Events.react_refer(result);
     }
    });
  },

  react_refer: async function(result) {
    reference_offers = await App.get_offers(result.args.reference_id);
    common_offer = in_my_interest(reference_offers);
    if(common_offer.length==0){
      //the goods from reference_id is already in user's reference list
      reference_offers.forEach(function(offer_id){
        add_offer(offer_id, true, result.args.referer, result.args.referer_interest, result.blockNumber);
      });
    }else {
      //the goods from reference_id is already in user's reference list
      //initiate_cycle with first common good
      // TODO: ask user for which good would he like more?
      initiate_cycle(common_offer[0]);
    }
  },

  init_events: function() {
      Events.listen_refer();
  }

};
