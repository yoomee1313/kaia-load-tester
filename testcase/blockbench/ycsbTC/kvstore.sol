// Derived from BlockBench's YCSB benchmark.
pragma solidity ^0.8.24;

contract KVstore {
    mapping(string=>string) store;

    function get(string memory key) public view returns(string memory) {
        return store[key];
    }

    function set(string memory key, string memory value) public {
        store[key] = value;
    }
}
