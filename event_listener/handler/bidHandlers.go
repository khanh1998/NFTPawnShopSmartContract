package handler

import (
	"encoding/json"
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"log"
)

type BidStatus int

const (
	BID_CREATED BidStatus = iota
	BID_CANCELLED
	BID_ACCEPTED
)

func BidCreated(bidCreated *pawningShop.ContractsBidCreated, instance *pawningShop.Contracts, client *httpClient.Client) {
	fmt.Println(BidCreatedName)
	bid, err := instance.Bids(nil, bidCreated.BidId)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bid)
	success := client.BidPawn.InsertOne(
		bidCreated.BidId.String(),
		bid.Creator.String(),
		bid.LoanAmount.String(),
		bid.Interest.String(),
		bid.LoanStartTime.String(),
		bid.LoanDuration.String(),
		bid.IsInterestProRated,
		bidCreated.PawnId.String(),
	)
	log.Println("to api", BidCreatedName, success)
	if success {
		bidData, _ := instance.Bids(nil, bidCreated.BidId)
		payload, _ := json.Marshal(bidData)
		client.Notify.SendNotification(httpClient.Notification{
			Message:  "New bid is created",
			Code:     BidCreatedName,
			PawnID:   bidCreated.PawnId.String(),
			BidID:    bidCreated.BidId.String(),
			Lender:   bidCreated.Lender.String(),
			Borrower: bidCreated.Borrower.String(),
			Payload:  string(payload),
		})
	}
	log.Println("to notify", BidCreatedName, success)
}

func BidAccepted(bidAcc *pawningShop.ContractsBidAccepted, instance *pawningShop.Contracts, client *httpClient.Client) {
	log.Println(BidAcceptedName)
	bid, err := instance.Bids(nil, bidAcc.BidId)
	if err != nil {
		log.Panic(err)
	}
	success := client.BidPawn.UpdateOne(bidAcc.BidId.String(), int(BID_ACCEPTED), bid.LoanStartTime.String())
	log.Println("to api", BidAcceptedName, success)
	if success {
		success, _ := client.Notify.SendNotification(httpClient.Notification{
			Code:     BidAcceptedName,
			Message:  "A bid is accepted",
			BidID:    bidAcc.BidId.String(),
			PawnID:   bidAcc.PawnId.String(),
			Lender:   bidAcc.Lender.String(),
			Borrower: bidAcc.Borrower.String(),
		})
		log.Println("to notify", BidAcceptedName, success)
	}
}

func BidCancelled(bidCancel *pawningShop.ContractsBidCancelled, instance *pawningShop.Contracts, client *httpClient.Client) {
	log.Println(BidCancelledName)
	success, resBody := client.Bid.UpdateOne(bidCancel.BidId.String(), int(BID_CANCELLED))
	log.Println("to api", BidCancelledName, success)
	if success {
		success, _ := client.Notify.SendNotification(httpClient.Notification{
			Code:     BidCancelledName,
			Message:  "A bid is cancelled",
			BidID:    bidCancel.BidId.String(),
			PawnID:   bidCancel.PawnId.String(),
			Lender:   bidCancel.Lender.String(),
			Borrower: bidCancel.Borrower.String(),
			Payload:  resBody,
		})
		log.Println("to notify", BidCancelledName, success)
	}
}
