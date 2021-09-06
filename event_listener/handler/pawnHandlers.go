package handler

import (
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"log"
	"math/big"
)

type PawnStatus int

const (
	CREATED PawnStatus = iota
	CANCELLED
	DEAL
	LIQUIDATED
	REPAID
)

func PawnCreated(pawnId *big.Int, instance *pawningShop.Contracts, client *httpClient.Client) {
	fmt.Println(PawnCreatedName)
	pawn, err := instance.Pawns(nil, pawnId)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pawn)
	success := client.Pawn.InsertOne(
		pawnId.String(),
		pawn.Creator.String(),
		pawn.ContractAddress.String(),
		pawn.TokenId.String(),
		pawn.Status,
	)
	log.Println("to api", PawnCreatedName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Message:  "New pawn is created",
			Code:     PawnCreatedName,
			PawnID:   pawnId.String(),
			Borrower: pawn.Creator.String(),
		})
		log.Println("to notify", PawnCreatedName, success)
	}
}

func PawnCancelled(pawnId *big.Int, instance *pawningShop.Contracts, client *httpClient.Client) {
	log.Println(PawnCancelledName)
	success := client.Pawn.UpdateOne(pawnId.String(), int(CANCELLED), "")
	log.Println("to api", PawnCancelledName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Message: "A pawn is cancelled",
			Code:    PawnCancelledName,
			PawnID:  pawnId.String(),
		})
		log.Println("to notify", PawnCancelledName, success)
	}
}

func PawnRepaid(pawnId *big.Int, client *httpClient.Client) {
	log.Println(PawnRepaidName)
	success := client.Pawn.UpdateOne(pawnId.String(), int(REPAID), "")
	log.Println("to api", PawnRepaidName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Message: "A pawn is repaid",
			Code:    PawnRepaidName,
			PawnID:  pawnId.String(),
		})
		log.Println("to notify", PawnRepaidName, success)
	}
}

func PawnLiquidated(pawnId *big.Int, client *httpClient.Client) {
	log.Println(PawnLiquidatedName)
	success := client.Pawn.UpdateOne(pawnId.String(), int(LIQUIDATED), "")
	log.Println("to api", PawnLiquidatedName, success)
	if success {
		success := client.Notify.SendNotification(httpClient.Notification{
			Message: "A pawn is liquidated",
			Code:    PawnLiquidatedName,
			PawnID:  pawnId.String(),
		})
		log.Println("to notify", PawnLiquidatedName, success)
	}
}
