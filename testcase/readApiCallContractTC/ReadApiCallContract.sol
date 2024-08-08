pragma solidity ^0.8.24;  // (required) version pragma

contract ReadApiCallContract {
        uint num = 4;  // state variable
        uint val;
        function get() public view returns (uint) {
                return num; //4
        }
        function set() public {
                val = 8;
        }
}