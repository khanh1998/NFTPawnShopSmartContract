package handler

import (
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BidStatus int

const (
	BID_CREATED BidStatus = iota
	BID_CANCELLED
	BID_ACCEPTED
)

func BidCreated(bidId *big.Int, pawnId *big.Int, instance *pawningShop.Contracts, client *httpClient.Client) {
	fmt.Println(BidCreatedName)
	bid, err := instance.Bids(nil, bidId)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bid)
	success := client.BidPawn.InsertOne(
		bidId.String(),
		bid.Creator.String(),
		bid.LoanAmount.String(),
		bid.Interest.String(),
		bid.LoanStartTime.String(),
		bid.LoanDuration.String(),
		bid.IsInterestProRated,
		pawnId.String(),
	)
	log.Println("to api", BidCreatedName, success)
	if success {
		client.Notify.SendNotification(httpClient.Notification{
			Message: "New bid is created",
			Code:    BidCreatedName,
			PawnID:  pawnId.String(),
			BidID:   bidId.String(),
			Lender:  bid.Creator.String(),
		})
	}
	log.Println("to notify", BidCreatedName, success)
}

func BidAccepted(bidId *big.Int, instance *pawningShop.Contracts, client *httpClient.Client) {
	log.Println(BidAcceptedName)
	bid, err := instance.Bids(nil, bidId)
	if err != nil {
		log.Panic(err)
	}
	success := client.BidPawn.UpdateOne(bidId.String(), int(BID_ACCEPTED), bid.LoanStartTime.String())
	log.Println("to api", BidAcceptedName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Code:    BidAcceptedName,
			Message: "A bid is accepted",
			BidID:   bidId.String(),
		})
		log.Println("to notify", BidAcceptedName, success)
	}
}

func BidCancelled(bidId *big.Int, creator common.Address, instance *pawningShop.Contracts, client *httpClient.Client) {
	log.Println(BidCancelledName)
	success := client.Bid.UpdateOne(bidId.String(), int(BID_CANCELLED))
	log.Println("to api", BidCancelledName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Code:    BidCancelledName,
			Message: "A bid is cancelled",
			BidID:   bidId.String(),
			Lender:  creator.String(),
		})
		log.Println("to notify", BidCancelledName, success)
	}
}
