pragma solidity >=0.4.21 <0.7.0;

contract node_maker{
  struct Node {
    address addr;
    Message[] mail_box;
  }

  mapping (address => uint) address_to_id;
  Node[] nodes;

}
