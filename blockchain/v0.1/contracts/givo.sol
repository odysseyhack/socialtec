pragma solidity >=0.4.0 <0.7.0;



contract givo {

  struct Good {
    address node;
    string name;
    string ipfs_image;
    string ipfs_details;
  }

  uint constant good_count = 5;
  uint constant page_size = 5;
  uint public node_cntr = 0;
  mapping (uint => Good[]) public offers;
  mapping (address => uint) address_to_id;

  event intrested(uint to, uint good_id, uint from);
  event not_intrested(uint to, uint good_id, uint from);
  event refer(uint to, uint good_id, uint from, uint refer_id);
  event deleted(uint owner, uint good_id);
  event chained(uint to, uint get_owner, uint get_good_id, uint give_owner, uint give_good_id);

  function create_offer(string memory name, string memory ipfs_image, string memory ipfs_details) public returns (bool) {
    Good[] storage my_offers = offers[address_to_id[msg.sender]];
    if(my_offers.length==0){
        address_to_id[msg.sender] = node_cntr;
        node_cntr++;
    }
    Good memory new_good = Good(msg.sender, name, ipfs_image, ipfs_details);
    my_offers.push(new_good);
    return true;
  }

  function delete_offer(uint good_id) public{
    Good[] storage my_offers = offers[address_to_id[msg.sender]];
    require(my_offers.length>=good_id);
    my_offers[good_id] = my_offers[my_offers.length-1];
    my_offers.pop();
    emit deleted(address_to_id[msg.sender], good_id);
  }

/*  function get_offers(uint page) public view returns (Good[page_size][good_count] memory goods){
      uint last_visit = (page-1)*page_size;
      require(last_visit<node_cntr);
      uint max_nodes = page_size;
      if(page*page_size>node_cntr){
          max_nodes = node_cntr - last_visit;
      }

      for(uint node_cntr = 0; node_cntr<max_nodes; node_cntr++){
          for(uint good_cntr = 0; good_cntr<good_count; good_cntr++){
              goods[node_cntr][good_cntr] = new Good(msg.sender,"a","b","c");
             //goods[node_cntr][good_cntr] = offers[last_visit+node_cntr][good_cntr];
          }
      }
      return goods;
  }*/

  function add_intrest(uint node_id, uint good_id) public {
      emit intrested(node_id, good_id, address_to_id[msg.sender]);
  }

  function delete_intrest(uint node_id, uint good_id) public {
      emit intrested(node_id, good_id, address_to_id[msg.sender]);
  }

  function refer_intrest(uint node_id, uint good_id, uint refer_id) public {
      emit refer(node_id, good_id, address_to_id[msg.sender], refer_id);
  }

  function cycle_formed(uint to, uint get_good_id, uint give_owner, uint give_good_id) public{
      emit chained(to, address_to_id[msg.sender], get_good_id, give_owner, give_good_id);
  }



}
