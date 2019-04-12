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

  function create_good(uint id, string memory name, string memory ipfs_image, string memory ipfs_details) external returns (bool) {
    return true;
  }
}
