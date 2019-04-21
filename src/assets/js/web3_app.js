/*jshint esversion: 8 */

App = {
  account: '0x08427F4196800c59c925cb1D0c1BcfcB759f13C6',
  address: '0x1fc3a9b1ec7fddd718ecddec6bde4dcaf1472bff',
  contract: '',
  options: {
    from: this.account
  },

  init: function(run) {
    if (typeof web3 !== 'undefined') {
      // If a web3 instance is already provided by Meta Mask.
      web3Provider = web3.currentProvider;
      web3 = new Web3(web3.currentProvider);
    } else {
      // Specify default instance if no web3 instance provided
      console.log("Metamask Unavailable");
      web3Provider = new Web3.providers.HttpProvider('http://localhost:7545');
      web3 = new Web3(web3Provider);
    }
    return App.init_contract(run);
  },

  get_balance: function() {
    web3.eth.getBalance(App.account, (err, wei) => {
      var eth = web3._extend.utils.fromWei(wei, 'ether');
      console.log(eth);
    });
  },

  init_contract: function(run) {
    $.getJSON( "../build/contracts/givo.json", function( jsonInterface ) {
      //console.log(jsonInterface);
      App.contract = web3.eth.contract(jsonInterface.abi).at(App.address);
      //console.log(App.contract);
      ethereum.enable();
      run();
    });
  },

 get_id: async function() {
    if(Storage.my_id() == 0) {
      await promisify(cb => App.contract.assign_id(App.options, cb));
      id = await promisify(cb => App.contract.address_to_id(App.account, App.options, cb));
      try {
        console.log("Id Assigned:"+id);
        Storage.set_id(id);
        return id;
      } catch(error) {
        console.log("Error Assigning Id");
      }
    }else {
      return Storage.my_id();
    }
  },

  get_offers: async function(id) {
    offers = promisify(cb => App.contract.get_offers(id, App.option, cb));
    console.log(offers);
    return await offers;
  }

};
