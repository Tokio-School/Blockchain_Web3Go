// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Tokio-School-Test
 * @dev Hola mundo
 */
contract HolaMundo {

    string message;

    /**
     * @dev Store message in variable
     * @param mes value to store
     */
    function storeMessage(string memory mes) public {
        message = mes;
    }

    /**
     * @dev Return value 
     * @return value of 'message'
     */
    function retrieveMessage() public view returns (string memory){
        return message;
    }
}


