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

func BidCreated(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	fmt.Println(BidCreatedName)
	data := UpackEvent(abi, BidCreatedName, vlog.Data)
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
		client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH)
		success := client.Bid.Post(
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
