pragma solidity >=0.4.0 <0.7.0;

pragma experimental ABIEncoderV2;

contract givo {

  struct Good {
    address node;
    string name;
    string ipfs_image;
    string ipfs_details;
  }

  uint good_count = 5;

  struct Node {
    address addr;
    Message[] mail_box;
  }

  struct Message {
    address from;
    address refer;
    string body;
  }

  Node[] nodes;
  mapping (uint => Good[]) public offers;
  mapping (address => uint) address_to_id;

  function create_offer(string memory name, string memory ipfs_image, string memory ipfs_details) public returns (bool) {
    Good[] storage my_offers = offers[address_to_id[msg.sender]];
    Good memory new_good = Good(msg.sender, name, ipfs_image, ipfs_details);
    my_offers.push(new_good);
    return true;
  }

  function get_offers() public view returns (Good[1][1] memory goods){
      Good memory good;
      good.node = msg.sender;
      good.name = "apple";
      good.ipfs_image = "ball";
      good.ipfs_details = "ball";
      goods[0][0]= good;
      return goods;
  }

  function add_intrest(Good memory good) public returns (bool) {
      return true;
  }

  function delete_intrest(Good memory good) public returns (bool) {
      return true;
  }

}
