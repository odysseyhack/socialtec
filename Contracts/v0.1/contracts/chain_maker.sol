pragma solidity >=0.4.21 <0.7.0;

require("./mail.sol");

contract chain_maker is mail{

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

  Node[] nodes;
  mapping (uint => Good[]) offers;
  mapping (address => uint) address_to_id;

  function create_good(uint id, string name, string ipfs_image, string ipfs_details) internal {
    Good good;
    good.node = msg.sender;
    good.name = name;
    good.ipfs_image = ipfs_image;
    good.ipfs_details = ipfs_details;
    offers[id].push(good);
  }

  function register(Good[] memory good) public returns (uint) {
    Node node;
    node.addr = msg.sender;
    uint id = nodes.push(node);
    for(uint cnt=0; cnt<good_count; cnt++){
      create_good(id, name, ipfs_image, ipfs_details);
    }
    address_to_id[msg.sender] = id;
    return id;
  }

  function get_offers(uint last_id=0) external view returns(Goods[]) {
    Goods[] result;
    for(int itr_id=last_id; ctr<nodes.length; ctr++){
      for(int itr_good=0; itr_good<good_count; itr_good++){
        results.push(offers[itr_id][itr_good]);
      }
    }
    return results;
  }

}
