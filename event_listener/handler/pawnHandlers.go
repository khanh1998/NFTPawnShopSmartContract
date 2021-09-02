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

type PawnStatus int

const (
	CREATED PawnStatus = iota
	CANCELLED
	DEAL
	LIQUIDATED
	REPAID
)

func PawnCreated(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	fmt.Println(PawnCreatedName)
	data := UnpackEvent(abi, PawnCreatedName, vlog.Data)
	fmt.Println(data)
	newPawnIdStr := data[1]
	newPawnIdInt := new(big.Int)
	newPawnIdInt, ok := newPawnIdInt.SetString(newPawnIdStr, 10)
	if !ok {
		log.Panic("cannot convert string to bigint")
	} else {
		pawn, err := instance.Pawns(nil, newPawnIdInt)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(pawn)
		client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
		success := client.Pawn.InsertOne(
			newPawnIdStr,
			pawn.Creator.String(),
			pawn.ContractAddress.String(),
			pawn.TokenId.String(),
			pawn.Status,
		)
		log.Println(PawnCreatedName, success)
	}
}

func PawnCancelled(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	log.Println(PawnCancelledName)
	data := UnpackEvent(abi, PawnCancelledName, vlog.Data)
	fmt.Println(data)
	newPawnIdStr := data[1]
	client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
	success := client.Pawn.UpdateOne(newPawnIdStr, int(CANCELLED), "")
	log.Println(PawnCancelledName, success)
}

func PawnRepaid(pawnId string, env *config.Env) {
	log.Println(PawnRepaidName)
	client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
	success := client.Pawn.UpdateOne(pawnId, int(REPAID), "")
	log.Println(PawnRepaidName, success)
}

func PawnLiquidated(pawnId string, env *config.Env) {
	log.Println(PawnLiquidatedName)
	client := client.NewClient(env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH)
	success := client.Pawn.UpdateOne(pawnId, int(LIQUIDATED), "")
	log.Println(PawnLiquidatedName, success)

}
