package handler

import (
	"fmt"
	"khanh/client"
	"khanh/config"
	pawningShop "khanh/contracts"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

type BidStatus int

const (
	BID_CREATED BidStatus = iota
	BID_CANCELLED
	BID_ACCEPTED
)

func BidCreated(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	fmt.Println(BidCreatedName)
	data := UnpackEvent(abi, BidCreatedName, vlog.Data)
	fmt.Println(data)
	pawnIdStr := data[1]
	newBidIdStr := data[2]
	newBidIdInt := new(big.Int)
	newBidIdInt, ok := newBidIdInt.SetString(newBidIdStr, 10)
	if !ok {
		log.Panic("cannot convert string to bigint")
	} else {
		bid, err := instance.Bids(nil, newBidIdInt)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(bid)
		client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
		success := client.BidPawn.InsertOne(
			newBidIdStr,
			bid.Creator.String(),
			bid.LoanAmount.String(),
			bid.Interest.String(),
			bid.LoanStartTime.String(),
			bid.LoanDuration.String(),
			bid.IsInterestProRated,
			pawnIdStr,
		)
		log.Println(BidCreatedName, success)
	}
}

func BidAccepted(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	log.Println(BidAcceptedName)
	data := UnpackEvent(abi, BidAcceptedName, vlog.Data)
	bidIdStr := data[1]
	client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
	bidIdInt, ok := new(big.Int), false
	if bidIdInt, ok = bidIdInt.SetString(bidIdStr, 10); !ok {
		log.Panic("cannot convert string to big int")
	}
	bid, err := instance.Bids(nil, bidIdInt)
	if err != nil {
		log.Panic(err)
	}
	success := client.BidPawn.UpdateOne(bidIdStr, int(BID_ACCEPTED), bid.LoanStartTime.String())
	log.Println(BidAcceptedName, success)
}

func BidCancelled(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	log.Println(BidCancelledName)
	data := UnpackEvent(abi, BidCancelledName, vlog.Data)
	bidIdStr := data[2]
	client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
	success := client.Bid.UpdateOne(bidIdStr, int(BID_CANCELLED))
	log.Println(BidCancelledName, success)
}
