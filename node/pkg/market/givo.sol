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
  uint public node_cntr = 0;
  uint public max_all_good =0;
  mapping (uint => uint[]) public offers;
  mapping (address => uint) address_to_id;
  Good[] public all_goods;

  event intrested(uint to, uint good_id, uint from);
  event not_intrested(uint to, uint good_id, uint from);
  event refer(uint to, uint good_id, uint from, uint refer_id);
  event deleted(uint owner, uint good_id);
  event chained(uint to, uint get_owner, uint get_good_id, uint give_owner, uint give_good_id);

  function create_offer(string memory name, string memory ipfs_image, string memory ipfs_details) public returns (uint good_id) {
    uint[] storage my_offers = offers[address_to_id[msg.sender]];
    if(my_offers.length==0){
        address_to_id[msg.sender] = node_cntr;
        node_cntr++;
    }
    Good memory new_good = Good(address_to_id[msg.sender], 0, true, name, ipfs_image, ipfs_details);
    good_id = all_goods.push(new_good)-1;
    uint index = my_offers.push(good_id) -1;
    new_good.index = index;
    max_all_good = good_id;
    return good_id;
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

  function add_intrest(uint good_id) public {
      uint node_id = all_goods[good_id].node;
      emit intrested(node_id, good_id, address_to_id[msg.sender]);
  }

  function delete_intrest(uint good_id) public {
      uint node_id = all_goods[good_id].node;
      emit intrested(node_id, good_id, address_to_id[msg.sender]);
  }

  function refer_intrest(uint good_id, uint refer_id) public {
      uint node_id = all_goods[good_id].node;
      emit refer(node_id, good_id, address_to_id[msg.sender], refer_id);
  }

  function cycle_formed(uint to, uint get_good_id, uint give_owner, uint give_good_id) public{
      emit chained(to, address_to_id[msg.sender], get_good_id, give_owner, give_good_id);
  }

  //I give (good) (to) you if (owner) gives (con_good) to (getter).
  function cycle_propagate(uint to, uint good, uint owner, uint con_good, uint getter) public{
      emit chained(to, getter, good, owner, con_good);
  }



}
