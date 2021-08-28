package main

import (
	"fmt"
	"khanh/client"
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

func PawnCreatedHandler(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, apiHost string) {
	fmt.Println(PawnCreated)
	data := UpackEvent(abi, PawnCreated, vlog.Data)
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
		client := client.NewClient(apiHost)
		success := client.Pawn.Post(
			newPawnIdStr,
			pawn.Creator.String(),
			pawn.ContractAddress.String(),
			pawn.TokenId.String(),
			pawn.Status,
		)
		log.Println(PawnCreated, success)
	}
}

func PawnCancelledHandler(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts, apiHost string) {
	fmt.Println(PawnCancelled)
	data := UpackEvent(abi, PawnCancelled, vlog.Data)
	fmt.Println(data)
	newPawnIdStr := data[1]
	client := client.NewClient(apiHost)
	const CANCELLED = 1
	success := client.Pawn.Patch(newPawnIdStr, CANCELLED, "")
	log.Println(PawnCancelled, success)
}

func WhiteListAddedHandler(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
	fmt.Println(WhiteListAdded)
	data := UpackEvent(abi, WhiteListAdded, vlog.Data)
	fmt.Println(data)
	address, err := instance.WhiteListNFT(nil, big.NewInt(0))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(address)
}
func WhiteListRemovedHander(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
	fmt.Println(WhiteListRemoved)
	data := UpackEvent(abi, WhiteListRemoved, vlog.Data)
	fmt.Println(data)
	instance.WhiteListNFT(nil, big.NewInt(0))
	address, err := instance.WhiteListNFT(nil, big.NewInt(0))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(address)
}
