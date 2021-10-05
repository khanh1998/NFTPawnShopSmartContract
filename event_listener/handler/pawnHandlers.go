package handler

import (
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"khanh/rabbitmq"
	"log"
)

type PawnStatus int

const (
	CREATED PawnStatus = iota
	CANCELLED
	DEAL
	LIQUIDATED
	REPAID
)

type PawnHandler struct {
	instance    *pawningShop.Contracts
	client      *httpClient.Client
	queue       *rabbitmq.RabbitMQ
	channelName string
}

func NewPawnHandler(instance *pawningShop.Contracts, client *httpClient.Client, rabbit *rabbitmq.RabbitMQ, channelName string) *PawnHandler {
	return &PawnHandler{
		instance: instance,
		client:   client,
		queue:    rabbit,
	}
}

func (p *PawnHandler) PawnCreated(pawnCreated *pawningShop.ContractsPawnCreated) interface{} {
	fmt.Println(PawnCreatedName)
	pawn, err := p.instance.Pawns(nil, pawnCreated.PawnId)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pawn)
	success, resBody := p.client.Pawn.InsertOne(
		pawnCreated.PawnId.String(),
		pawn.Creator.String(),
		pawn.ContractAddress.String(),
		pawn.TokenId.String(),
		pawn.Status,
	)
	log.Println("to api", PawnCreatedName, success)
	if success {
		data := httpClient.Notification{
			Message:  "New pawn is created",
			Code:     PawnCreatedName,
			PawnID:   pawnCreated.PawnId.String(),
			Borrower: pawn.Creator.String(),
			Payload:  string(resBody),
		}
		success, _ := p.client.Notify.SendNotification(data)
		log.Println("to notify", PawnCreatedName, success)
		err := p.queue.SerializeAndSend(p.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}

func (p *PawnHandler) PawnCancelled(pawn *pawningShop.ContractsPawnCancelled) interface{} {
	log.Println(PawnCancelledName)
	success, resBody := p.client.Pawn.UpdateOne(pawn.PawnId.String(), int(CANCELLED), "")
	log.Println("to api", PawnCancelledName, success)
	if success {
		data := httpClient.Notification{
			Message:  "A pawn is cancelled",
			Code:     PawnCancelledName,
			PawnID:   pawn.PawnId.String(),
			Borrower: pawn.Borrower.String(),
			Payload:  string(resBody),
		}
		success, _ := p.client.Notify.SendNotification(data)
		log.Println("to notify", PawnCancelledName, success)
		err := p.queue.SerializeAndSend(p.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}

func (p *PawnHandler) PawnRepaid(pawn *pawningShop.ContractsPawnRepaid) interface{} {
	log.Println(PawnRepaidName)
	success, resBody := p.client.Pawn.UpdateOne(pawn.PawnId.String(), int(REPAID), "")
	log.Println("to api", PawnRepaidName, success)
	if success {
		data := httpClient.Notification{
			Message:  "A pawn is repaid",
			Code:     PawnRepaidName,
			PawnID:   pawn.PawnId.String(),
			BidID:    pawn.BidId.String(),
			Lender:   pawn.Lender.String(),
			Borrower: pawn.Borrower.String(),
			Payload:  resBody,
		}
		success, _ := p.client.Notify.SendNotification(data)
		log.Println("to notify", PawnRepaidName, success)
		err := p.queue.SerializeAndSend(p.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}

func (p *PawnHandler) PawnLiquidated(pawn *pawningShop.ContractsPawnLiquidated) interface{} {
	log.Println(PawnLiquidatedName)
	success, resBody := p.client.Pawn.UpdateOne(pawn.PawnId.String(), int(LIQUIDATED), "")
	log.Println("to api", PawnLiquidatedName, success)
	if success {
		data := httpClient.Notification{
			Message:  "A pawn is liquidated",
			Code:     PawnLiquidatedName,
			PawnID:   pawn.PawnId.String(),
			BidID:    pawn.BidId.String(),
			Lender:   pawn.Lender.String(),
			Borrower: pawn.Borrower.String(),
			Payload:  resBody,
		}
		success, _ := p.client.Notify.SendNotification(data)
		log.Println("to notify", PawnLiquidatedName, success)
		err := p.queue.SerializeAndSend(p.channelName, data)
		if err != nil {
			log.Panic(err)
		} else {
			log.Println("to notify rabbitmq", BidCreatedName, success)
		}
		return data
	}
	return nil
}
