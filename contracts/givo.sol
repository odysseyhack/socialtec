pragma solidity >=0.4.0 <0.7.0;



contract givo {
  struct Good {
    uint node;
    uint index;
    bool available;
    string name;
    string ipfs_image;
    string ipfs_details;
  }

  uint constant good_count = 5;
  uint public node_cntr = 1;
  uint public max_all_good =0;
  mapping (uint => uint[]) public offers;
  mapping (address => uint) public address_to_id;
  Good[] public all_goods;

  event not_intrested(uint indexed to, uint good_id, uint from);
  event refer(uint indexed to, uint referer_interest, uint referer, uint reference_id);
  event deleted(uint owner, uint indexed good_id);
  event chained(uint if_owner, uint if_good, uint if_receiver, uint then_good, uint indexed then_receiver);

  function create_offer(string memory name, string memory ipfs_image, string memory ipfs_details) public returns (uint good_id) {
    uint id = address_to_id[msg.sender];
    if (id == 0) {
        address_to_id[msg.sender] = node_cntr;
        node_cntr=node_cntr+1;
    }
    uint[] storage my_offers = offers[address_to_id[msg.sender]];
    Good memory new_good = Good(address_to_id[msg.sender], 0, true, name, ipfs_image, ipfs_details);
    good_id = all_goods.push(new_good)-1;
    uint index = my_offers.push(good_id) -1;
    new_good.index = index;
    max_all_good = good_id;
    return good_id;
  }

  function assign_id() external {
    uint id = address_to_id[msg.sender];
    if (id == 0) {
        address_to_id[msg.sender] = node_cntr;
        id = node_cntr;
        node_cntr=node_cntr+1;
    }
  }

  function delete_offer(uint good_id) public{
    uint[] storage my_offers = offers[address_to_id[msg.sender]];
    Good storage my_good = all_goods[good_id];
    require(max_all_good>good_id);
    require(my_good.node==address_to_id[msg.sender]);
    my_good.available =false;
    my_offers[my_good.index] = my_offers[my_offers.length-1];
    Good storage last_good = all_goods[my_offers[my_good.index]];
    last_good.index = my_good.index;
    my_offers.pop();
    emit deleted(address_to_id[msg.sender], good_id);
  }

  function get_offers(uint node_id) public view returns (uint[] memory my_offers){
      my_offers = offers[node_id];
      return my_offers;
  }

  function add_interest(uint interested_good_id) public {
      uint owner_interested_good = all_goods[interested_good_id].node;
      emit refer(owner_interested_good, interested_good_id, address_to_id[msg.sender], address_to_id[msg.sender]);
  }

  function delete_interest(uint good_id) public {
      uint node_id = all_goods[good_id].node;
      emit not_intrested(node_id, good_id, address_to_id[msg.sender]);
  }

  function refer_interest(uint intrested_good_id, uint reference_id) public {
        uint owner_interested_good = all_goods[intrested_good_id].node;
        emit refer(owner_interested_good, intrested_good_id, address_to_id[msg.sender], reference_id);

  }

  function cycle_formed(uint if_owner, uint if_good,  uint then_good, uint then_receiver) public{
      emit chained(if_owner, if_good, address_to_id[msg.sender], then_good,then_receiver);
  }

  function cycle_propagate(uint if_owner, uint if_good,uint if_receiver,  uint then_good, uint then_receiver) public{
      emit chained(if_owner, if_good, if_receiver, then_good,then_receiver);
  }

}
