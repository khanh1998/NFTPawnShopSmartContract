package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"khanh/client"
	pawningShop "khanh/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PawnCreated      = "PawnCreated"
	WhiteListAdded   = "WhiteListAdded"
	WhiteListRemoved = "WhiteListRemoved"
)
const (
	PawnCreatedSignature      = "PawnCreated(address,uint256)"
	WhiteListAddedSignature   = "WhiteListAdded(address)"
	WhiteListRemovedSignature = "WhiteListRemoved(address)"
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:7545")
	// httpClient, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	smartContractAddress := "0x5501064E55e845f8c71fB9C93A6edcdCc23686A6"
	contractAddress := common.HexToAddress(smartContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(pawningShop.ContractsABI)))
	if err != nil {
		log.Panic(err)
	}

	instance, _ := pawningShop.NewContracts(common.HexToAddress(smartContractAddress), client)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			CategorizeEvent(vLog, contractAbi, instance)
		}
	}

}

func UpackEvent(contractAbi abi.ABI, eventName string, data []byte) []string {
	event, err := contractAbi.Unpack(eventName, data)
	if err != nil {
		log.Panic(err)
	}
	strs := make([]string, len(event))
	for i, v := range event {
		strs[i] = fmt.Sprint(v)
	}
	return strs
}

func CategorizeEvent(log types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
	incommingEventHash := log.Topics[0]
	fmt.Println("incomming event hash: ", incommingEventHash)
	switch incommingEventHash {
	case Hash(PawnCreatedSignature):
		PawnCreatedHandler(log, abi, instance)
	case Hash(WhiteListAddedSignature):
		WhiteListAddedHandler(log, abi, instance)
	case Hash(WhiteListRemovedSignature):
		WhiteListRemovedHander(log, abi, instance)
	}
}

func Hash(signature string) common.Hash {
	return crypto.Keccak256Hash([]byte(signature))
}

func PawnCreatedHandler(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
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
		client := client.NewClient("http://localhost:4000")
		success := client.Pawn.Post(
			newPawnIdStr,
			pawn.Creator.String(),
			pawn.ContractAddress.String(),
			pawn.TokenId.String(),
			int(pawn.Status),
		)
		log.Println(PawnCreated, success)
	}
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
