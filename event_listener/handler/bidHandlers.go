package handler

import (
	"encoding/json"
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"khanh/rabbitmq"
	"log"
)

type BidStatus int

const (
	BID_CREATED BidStatus = iota
	BID_CANCELLED
	BID_ACCEPTED
)

type BidHandler struct {
	instance    *pawningShop.Contracts
	client      *httpClient.Client
	queue       *rabbitmq.RabbitMQ
	channelName string
}

func NewBidHandler(instance *pawningShop.Contracts, client *httpClient.Client, rabbit *rabbitmq.RabbitMQ, channelName string) *BidHandler {
	return &BidHandler{
		instance: instance,
		client:   client,
		queue:    rabbit,
	}
}

func (b *BidHandler) BidCreated(bidCreated *pawningShop.ContractsBidCreated) interface{} {
	fmt.Println(BidCreatedName)
	bid, err := b.instance.Bids(nil, bidCreated.BidId)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bid)
	success := b.client.BidPawn.InsertOne(
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
		bidData, _ := b.instance.Bids(nil, bidCreated.BidId)
		payload, _ := json.Marshal(bidData)
		data := httpClient.Notification{
			Message:  "New bid is created",
			Code:     BidCreatedName,
			PawnID:   bidCreated.PawnId.String(),
			BidID:    bidCreated.BidId.String(),
			Lender:   bidCreated.Lender.String(),
			Borrower: bidCreated.Borrower.String(),
			Payload:  string(payload),
		}
		success, _ := b.client.Notify.SendNotification(data)
		log.Println("to notify", BidCreatedName, success)
		err := b.queue.SerializeAndSend(b.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}

func (b *BidHandler) BidAccepted(bidAcc *pawningShop.ContractsBidAccepted) interface{} {
	log.Println(BidAcceptedName)
	bid, err := b.instance.Bids(nil, bidAcc.BidId)
	if err != nil {
		log.Panic(err)
	}
	success := b.client.BidPawn.UpdateOne(bidAcc.BidId.String(), int(BID_ACCEPTED), bid.LoanStartTime.String())
	log.Println("to api", BidAcceptedName, success)
	if success {
		data := httpClient.Notification{
			Code:     BidAcceptedName,
			Message:  "A bid is accepted",
			BidID:    bidAcc.BidId.String(),
			PawnID:   bidAcc.PawnId.String(),
			Lender:   bidAcc.Lender.String(),
			Borrower: bidAcc.Borrower.String(),
		}
		success, _ := b.client.Notify.SendNotification(data)
		log.Println("to notify", BidAcceptedName, success)
		err := b.queue.SerializeAndSend(b.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}

func (b *BidHandler) BidCancelled(bidCancel *pawningShop.ContractsBidCancelled) interface{} {
	log.Println(BidCancelledName)
	success, resBody := b.client.Bid.UpdateOne(bidCancel.BidId.String(), int(BID_CANCELLED))
	log.Println("to api", BidCancelledName, success)
	if success {
		data := httpClient.Notification{
			Code:     BidCancelledName,
			Message:  "A bid is cancelled",
			BidID:    bidCancel.BidId.String(),
			PawnID:   bidCancel.PawnId.String(),
			Lender:   bidCancel.Lender.String(),
			Borrower: bidCancel.Borrower.String(),
			Payload:  resBody,
		}
		success, _ := b.client.Notify.SendNotification(data)
		log.Println("to notify", BidCancelledName, success)
		err := b.queue.SerializeAndSend(b.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}
