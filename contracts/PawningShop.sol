pragma solidity ^0.8.4;

import "./interfaces/IERC721.sol";

contract NFTPawnShop {
    enum PawnStatus { CREATED, CANCELLED, DEAL, LIQUIDATED, REPAID }
    struct Pawn {
        uint256 id;
        // who borrows money
        address creator;
        address contractAddress;
        uint256 tokenId;
        PawnStatus status;
    }
    struct Bid {
        uint256 id;
        address creator;
        uint256 loanAmount;
        uint8 interestRate;
        uint256 loanStartTime;
        uint256 loanDuration;
        bool isInterestProRated;
    }
    address public owner;
    address[] public _whiteListNFT;
    uint256 public _feeRate;
    uint256 public _totalNumberOfPawn = 0;
    uint256 public _totalNumberOfBid = 0;
    // mapping nft address -> token id -> pawn
    Pawn[] public _pawns;
    // mapping bid id => bid
    Bid[] public _bids;
    // mapping bid id to pawn id
    mapping(uint256 => uint256) public _bidToPawn;
    mapping(uint256 => uint256) public _pawnToBid;

    event PawnCreated(address pawner, uint256 pawnId);
    event PawnCancelled(address pawner, uint256 pawnId);
    event PawnDeal(address pawner, address lender, uint256 id);
    event PawnRepaid(address pawner, address lender, uint256 id);
    event PawnLiquidated(address pawner, address lender, uint256 id);
    event BidCreated(address creator, uint pawnId);
    event BidCancelled(address creator, uint pawnId);

    function createPawn(address tokenAddress, uint256 tokenId) public {
        bool isInWhiteList = false;
        for (uint i = 0; i < _whiteListNFT.length; i++) {
            if (tokenAddress == _whiteListNFT[i]) {
                isInWhiteList = true;
            }
        }
        address sender = msg.sender;
        require(isInWhiteList == true, "PawningShop: smart contract is not in white list");
        bool isApproved = IERC721(tokenAddress).getApproved(tokenId) == address(this);
        bool isApprovedForAll = IERC721(tokenAddress).isApprovedForAll(sender, address(this));
        require(isApproved || isApprovedForAll, "PawningShop: haven't got permission to transfer");
        Pawn memory pawn = Pawn({
            id: _totalNumberOfPawn,
            creator: sender,
            contractAddress: tokenAddress,
            tokenId: tokenId,
            status: PawnStatus.CREATED
        });

        _pawns[_totalNumberOfPawn] = pawn;
        emit PawnCreated(sender, _totalNumberOfPawn);
        _totalNumberOfPawn += 1;
    }
    function cancelPawn(uint256 pawnId) public {
        Pawn storage pawn = _pawns[pawnId];
        address creator = pawn.creator;
        require(pawn.status == PawnStatus.CREATED, "PawningShop: Only can cancel when it has status of CREATED");
        require(msg.sender == creator, "PawningShop: Only owner of the pawn can cancel it");
        require(_pawnToBid[pawnId] == 0, "PawningShop: Only can cancel when no bid is accepted");
        pawn.status = PawnStatus.CANCELLED;

        emit PawnCancelled(owner, pawnId);
    }
    function bid(uint8 rate, uint256 duration, bool isInterestProRated, uint256 loanStartTime, uint256 pawnId) public payable {
        address creator = msg.sender;
        uint256 amount = msg.value;
        Pawn storage pawn = _pawns[pawnId];
        require(pawn.status == PawnStatus.CREATED, "PawningShop: cannot bid this pawn");
        require(amount > 0, "PawningShop: amount of money must be bigger than 0");
        Bid memory newBid = Bid({
            id: _totalNumberOfBid,
            creator: creator,
            loanAmount: amount,
            interestRate: rate,
            loanDuration: duration,
            isInterestProRated: isInterestProRated,
            loanStartTime: loanStartTime
        });
        _bids[_totalNumberOfBid] = newBid;
        _bidToPawn[_totalNumberOfBid] = pawnId;
        emit BidCreated(creator, pawnId);
        _totalNumberOfBid += 1;
    }
    function cancelBid(uint256 bidId) public {
        Bid memory currBid = _bids[bidId];
        address sender = msg.sender;
        require(sender == currBid.creator, "PawningShop: only creator can cancel the bid");
        require(currBid.id == bidId, "PawningShop: the bid is not exited");
        require(_bidToPawn[bidId] == 0, "PawningShop: cannot cancel a accepted bid");
        delete _bidToPawn[bidId];
    }
    function repaid() public {}
    function liquidate() public {}
}
