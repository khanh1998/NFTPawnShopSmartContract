pragma solidity ^0.8.4;

import './interfaces/IERC165.sol';

contract ERC165 is IERC165 {
    mapping(bytes4 => bool) _supportedInterfaces;

    function supportsInterface(bytes4 interfaceId) external view virtual override returns (bool) {
        return _supportedInterfaces[interfaceId];
    }

    function _registerInterface(bytes4 interfaceId) internal virtual {
        require(interfaceId != 0xffffffff, 'ERC165: Interface ID is not correct');
    }
}