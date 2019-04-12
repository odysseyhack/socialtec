pragma solidity >=0.4.0 <0.7.0;


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
  mapping (uint => Good[]) offers;
  mapping (address => uint) address_to_id;

  function get_offers() public view returns (
    //address[1] memory node,
    //string[1] memory name,
    //string[1] memory ipfs_image,
    string memory ipfs_details){
       // node[0] = msg.sender;
        //name[0] = "apple";
        //ipfs_image[0] = "ball";
        ipfs_details = "cat";
    return(ipfs_details);
  }

  function create_offer(string memory name, string memory ipfs_image, string memory ipfs_details) public returns (bool) {
    return true;
  }


  function add_intrest(uint node_id, uint good_id) public returns (bool) {
      return true;
  }

  function delete_intrest(uint node_id, uint good_id) public returns (bool) {
      return true;
  }

}
