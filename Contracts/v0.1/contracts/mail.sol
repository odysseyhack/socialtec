pragma solidity >=0.4.21 <0.7.0;

contract mail{

  struct Message {
    address from;
    address refer;
    string body;
  }

  function send_mail(address to, address refer, string body) {
    Message message;
    message.from = msg.sender;
    message.refer = refer;
    message.body = body;
    Node storage node = nodes[address_to_id[to]];
    node.mail_box.push(message);
  }

  function pop_mail() external return(address from, address refer, string body){
    Node storage node = nodes[address_to_id[msg.sender]];
    require(nodes.mail_box.length>=1);
    Message message = nodes.mail_box.pop();
    return(message.from, message.refer, message.body);
  }

  function mail_empty() external view return(bool){
    Node storage node = nodes[address_to_id[msg.sender]];
    return (nodes.mail_box.length==0)
  }

  //I like your good X please check mine
  //My good X is no more available please remove it from like list.
  //I will give you good X tell A give good Y to me.
}
