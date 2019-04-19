App = {
  account: '0x08427F4196800c59c925cb1D0c1BcfcB759f13C6',
  address: '0x5db3b3b518e5ce17821da3f363bf432791d102d0',
  contract: '',

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
    })
  },

  init_contract: function(run) {
    $.getJSON( "../build/contracts/givo.json", function( jsonInterface ) {
      console.log(jsonInterface);
      App.contract = web3.eth.contract(jsonInterface.abi).at(App.address);
      console.log(App.contract);
      run();
    });
  }
}
