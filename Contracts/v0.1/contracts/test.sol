pragma solidity >=0.5.6;

contract Mortal {
    /* Define variable owner of the type address */
    address owner;

    /* This constructor is executed at initialization and sets the owner of the contract */
    constructor() public { owner = msg.sender; }

    /* Function to recover the funds on the contract */
    function kill() public { if (msg.sender == owner) selfdestruct(msg.sender); }
}

contract Greeter is Mortal {
    /* Define variable greeting of the type string */
    bytes32  greeting;
    event NewGreeting(bytes32 _greeting);
    /* This runs when the contract is executed */
    constructor(bytes32  _greeting) public {
        greeting = _greeting;
    }

    /* Main function */
    function greet(bytes32  _name) public returns (bytes32) {
        emit NewGreeting(greeting);
        return greeting;
    }
}
