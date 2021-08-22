pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/presets/ERC721PresetMinterPauserAutoId.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract TestToken is ERC721PresetMinterPauserAutoId {
    constructor() ERC721PresetMinterPauserAutoId("TestToken", "TT", "example.com") {}
}
