// Derived from BlockBench's IOHeavy benchmark.
pragma solidity ^0.8.24;

contract IOHeavy {

    bytes constant ALPHABET = "abcdefghijklmnopqrstuvwxy#$%^&*()_+[]{}|;:,./<>?`~abcdefghijklmnopqrstuvwxy#$%^&*()_+[]{}|;:,./<>?`~abcdefghijklmnopqrstuvwxy#$%^&*()_+[]{}|;:,./<>?`~";

    function uintToBytes(uint v) internal pure returns (bytes20 ret) {
        assembly {
            // Create a pointer to the start of the memory location for ret
            let ptr := add(ret, 0x20)

            // Initialize the pointer with 0
            mstore(ptr, 0)

            // Set each byte of ret
            for { let i := 0 } lt(i, 20) { i := add(i, 1) } {
                // Shift and mask the uint value to get the current byte
                let byteVal := and(shr(mul(8, sub(19, i)), v), 0xFF)

                // Store the byte at the correct position
                mstore8(add(ptr, i), byteVal)
            }
        }
    }

    function getKey(uint k) internal pure returns(bytes20) {
        return uintToBytes(k);
    }

    function getVal(uint k) internal pure returns(bytes memory ret) {
        ret = new bytes(100);
        for (uint i = 0; i < 100; i++) {
            ret[i] = ALPHABET[k%50+i];
        }
    }

    event finishWrite(uint size, uint signature);
    event finishScan(uint size, uint signature);

    mapping(bytes20=>bytes) store;

    function get(bytes20 key) public view returns(bytes memory) {
        return store[key];
    }
    function set(bytes20 key, bytes memory value) public {
        store[key] = value;
    }

    function write(uint start_key, uint size, uint signature) public {
        for (uint i = 0; i < size; i++) {
            set(getKey(start_key+i), getVal(start_key+i));
        }
        emit finishWrite(size, signature);
    }

    function scan(uint start_key, uint size, uint signature) public {
        bytes memory ret;
        for (uint i = 0; i < size; i++) {
            ret = get(getKey(start_key+i));
        }
        emit finishScan(size, signature);
    }

    function revert_scan(uint start_key, uint size, uint signature) public {
        bytes memory ret;
        for (uint i = 0; i < size; i++) {
            ret = get(getKey(start_key+size-i-1));
        }
        emit finishScan(size, signature);
    }
}
