// Derived from BlockBench's IOHeavy benchmark.

pragma solidity ^0.8.24;

contract LargeMemo {
    string public str;

    constructor() {
        str = "Hello, World";
    }

    function setName(string memory _str) public {
        str = _str;
    }

    function run() public view returns(string memory) {
        return str;
    }
}
