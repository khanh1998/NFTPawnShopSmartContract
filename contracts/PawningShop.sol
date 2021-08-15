pragma solidity ^0.8.4;

import "./interfaces/IERC721.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract NFTPawnShop {

    using SafeMath for uint256;
    using SafeMath for uint8;

    enum PawnStatus {
        CREATED,
        CANCELLED,
        DEAL,
        LIQUIDATED,
        REPAID
    }
    struct Pawn {
        // who borrows money
        address creator;
        address contractAddress;
        uint256 tokenId;
        PawnStatus status;
    }
    struct Bid {
        address creator;
        uint256 loanAmount;
        // the amount of wei borrower have to pay more
        uint256 interest;
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
    mapping(uint256 => Pawn) _pawns;
    // mapping bid id => bid
    mapping(uint256 => Bid) _bids;
    // mapping bid id to pawn id
    mapping(uint256 => uint256) public _bidToPawn;
    mapping(uint256 => uint256) public _pawnToBid;

    constructor() {
        owner = msg.sender;
        Pawn memory newPawn = Pawn({
            creator: address(0),
            contractAddress: address(0),
            tokenId: 0,
            status: PawnStatus.CANCELLED
        });
        _pawns[0] = newPawn;
        Bid memory newBid = Bid({
            creator: address(0),
            loanAmount: 0,
            interest: 0,
            loanStartTime: 0,
            loanDuration: 0,
            isInterestProRated: false
        });
        _bids[0] = newBid;
    }

    event PawnCreated(address pawner, uint256 pawnId);
    event PawnCancelled(address pawner, uint256 pawnId);
    event PawnDeal(address pawner, address lender, uint256 id);
    event PawnRepaid(address pawner, address lender, uint256 id);
    event PawnLiquidated(address pawner, address lender, uint256 id);
    event BidCreated(address creator, uint256 pawnId);
    event BidCancelled(address creator, uint256 pawnId);

    function createPawn(address tokenAddress, uint256 tokenId) public {
        _totalNumberOfPawn += 1;
        bool isInWhiteList = false;
        for (uint256 i = 0; i < _whiteListNFT.length; i++) {
            if (tokenAddress == _whiteListNFT[i]) {
                isInWhiteList = true;
            }
        }
        address sender = msg.sender;
        require(
            isInWhiteList == true,
            "PawningShop: smart contract is not in white list"
        );
        bool isApproved = IERC721(tokenAddress).getApproved(tokenId) ==
            address(this);
        bool isApprovedForAll = IERC721(tokenAddress).isApprovedForAll(
            sender,
            address(this)
        );
        require(
            isApproved || isApprovedForAll,
            "PawningShop: haven't got permission to transfer"
        );
        Pawn memory pawn = Pawn({
            creator: sender,
            contractAddress: tokenAddress,
            tokenId: tokenId,
            status: PawnStatus.CREATED
        });

        _pawns[_totalNumberOfPawn] = pawn;
        emit PawnCreated(sender, _totalNumberOfPawn);
    }

    function cancelPawn(uint256 pawnId) public {
        Pawn storage pawn = _pawns[pawnId];
        address creator = pawn.creator;
        require(
            pawn.status == PawnStatus.CREATED,
            "PawningShop: Only can cancel when it has status of CREATED"
        );
        require(
            msg.sender == creator,
            "PawningShop: Only owner of the pawn can cancel it"
        );
        require(
            _pawnToBid[pawnId] == 0,
            "PawningShop: Only can cancel when no bid is accepted"
        );
        pawn.status = PawnStatus.CANCELLED;

        emit PawnCancelled(owner, pawnId);
    }

    function bid(
        uint8 rate,
        uint256 duration,
        bool isInterestProRated,
        uint256 loanStartTime,
        uint256 pawnId
    ) public payable {
        _totalNumberOfBid += 1;
        address creator = msg.sender;
        uint256 amount = msg.value;
        Pawn storage pawn = _pawns[pawnId];
        require(pawnId > 0, "PawningShop: pawn id is not valid");
        require(
            pawn.status == PawnStatus.CREATED,
            "PawningShop: cannot bid this pawn"
        );
        require(
            amount > 0,
            "PawningShop: amount of money must be bigger than 0"
        );
        Bid memory newBid = Bid({
            creator: creator,
            loanAmount: amount,
            interest: rate,
            loanDuration: duration,
            isInterestProRated: isInterestProRated,
            loanStartTime: loanStartTime
        });
        _bids[_totalNumberOfBid] = newBid;
        _bidToPawn[_totalNumberOfBid] = pawnId;
        emit BidCreated(creator, pawnId);
    }

    function cancelBid(uint256 bidId) public {
        Bid memory currBid = _bids[bidId];
        address sender = msg.sender;
        require(
            sender == currBid.creator,
            "PawningShop: only creator can cancel the bid"
        );
        uint256 pawnId = _bidToPawn[bidId];
        require(_pawnToBid[pawnId] != bidId, "PawningShop: your bid is accepted, cannot cancel");
        address payable lender = payable(currBid.creator);
        lender.transfer(currBid.loanAmount);
        delete _bids[bidId];
        delete _bidToPawn[bidId];
    }

    function acceptBid(uint256 bidId) public {
        Bid storage currBid = _bids[bidId];
        uint256 pawnId = _bidToPawn[bidId];
        require(pawnId > 0, "PawningShop: The pawn is not existed");
        Pawn storage pawn = _pawns[pawnId];
        require(pawn.creator == msg.sender, "PawningShop: only creator of pawn can accept bid");
        IERC721(pawn.contractAddress).transferFrom(pawn.creator, address(this), pawn.tokenId);
        address payable borrower = payable(pawn.creator);
        borrower.transfer(currBid.loanAmount);
        pawn.status = PawnStatus.DEAL;
        _pawnToBid[pawnId] = bidId;
        pawn.status = PawnStatus.DEAL;
        currBid.loanStartTime = block.timestamp;
    }

    function repaid(uint256 pawnId) public payable {
        Pawn storage currPawn = _pawns[pawnId];
        require(
            currPawn.creator == msg.sender,
            "PawningShop: Only creator of pawn can repay"
        );
        uint256 bidId = _pawnToBid[pawnId];
        require(
            bidId != 0,
            "PawningShop: This pawn doen't have any accepted bid"
        );
        Bid storage currBid = _bids[bidId];
        require(block.timestamp <= currBid.loanStartTime + currBid.loanDuration, "PawningShop: to late to repaid");
        uint256 value = msg.value;
        uint256 elapsedDuration = block.timestamp - currBid.loanStartTime;
        uint256 repaidAmount = _calculateRepaidAmount(currBid.loanAmount, currBid.interest, currBid.loanDuration, elapsedDuration, currBid.isInterestProRated);
        require(value == repaidAmount, "PawningShop: pay exactly repaid amount");
        // transfer token back to borrower
        IERC721(currPawn.contractAddress).transferFrom(address(this), currPawn.creator, currPawn.tokenId);
        // transfer money to lender
        address payable lender = payable(currBid.creator);
        lender.transfer(value);
        delete _pawnToBid[pawnId];
        delete _bidToPawn[bidId];
    }

    function _calculateRepaidAmount(uint256 original, uint256 interest, uint256 totalDuration, uint256 elapsedDuration, bool isInterestProRated) internal pure returns (uint256) {
        uint256 interestDue = interest;
        if (isInterestProRated) {
            interestDue = (interest.div(totalDuration)).mul(elapsedDuration);
        }
        return original.add(interest);
    }

    function liquidate(uint256 bidId) public {
        Bid storage currBid = _bids[bidId];
        require(currBid.creator == msg.sender, "PawningShop: only creator of bid can liquidate token");
        require(block.timestamp > currBid.loanStartTime + currBid.loanAmount, "PawningShop: Not valid time to liquidate");
        uint256 pawnId = _bidToPawn[bidId];
        require(_pawnToBid[pawnId] == bidId, "PawningShop: this bid is not accepted by borrower");
        Pawn storage currPawn = _pawns[pawnId];
        IERC721(currPawn.contractAddress).transferFrom(address(this), currBid.creator, currPawn.tokenId);
    }
}
