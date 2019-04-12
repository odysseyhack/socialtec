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


  function create_good(uint id, string memory name, string memory ipfs_image, string memory ipfs_details) internal {
    Good memory good;
    good.node = msg.sender;
    good.name = name;
    good.ipfs_image = ipfs_image;
    good.ipfs_details = ipfs_details;
    offers[id].push(good);
  }



  function register(Good[] memory good) public returns (uint id) {
    Node memory node;
    node.addr = msg.sender;
    id = nodes.push(node);
    for(uint cnt=0; cnt<good_count; cnt++){
      create_good(id, good[cnt].name, good[cnt].ipfs_image, good[cnt].ipfs_details);
    }
    //address_to_id[msg.sender] = id;
    //return id;
  }


  function send_mail(address to, address refer, string memory body) public{
    Message memory message;
    message.from = msg.sender;
    message.refer = refer;
    message.body = body;
    Node storage node = nodes[address_to_id[to]];
    node.mail_box.push(message);
  }

  function pop_mail() external returns(address from, address refer, string memory body){
    Node storage node = nodes[address_to_id[msg.sender]];
    require(node.mail_box.length>=1);
    Message storage message = node.mail_box[node.mail_box.length];
    node.mail_box.pop();
    return(message.from, message.refer, message.body);
  }

  function mail_empty() external view returns(bool){
    Node storage node = nodes[address_to_id[msg.sender]];
    return (node.mail_box.length==0);
  }

  //I like your good X please check mine
  //My good X is no more available please remove it from like list.
  //I dont like your good anymore.
  //I will give you good X tell A give good Y to me.
  //Cycle Completed

}
