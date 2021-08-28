package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"khanh/config"
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
	PawnCancelled    = "PawnCancelled"
	WhiteListAdded   = "WhiteListAdded"
	WhiteListRemoved = "WhiteListRemoved"
)
const (
	PawnCreatedSignature      = "PawnCreated(address,uint256)"
	PawnCancelledSignature    = "PawnCancelled(address,uint256)"
	WhiteListAddedSignature   = "WhiteListAdded(address)"
	WhiteListRemovedSignature = "WhiteListRemoved(address)"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Panic(err)
	}
	client, err := ethclient.Dial(env.NETWORK_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(env.CONTRACT_ADDRESS)
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

	instance, _ := pawningShop.NewContracts(common.HexToAddress(env.CONTRACT_ADDRESS), client)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			CategorizeEvent(vLog, contractAbi, instance, env)
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

func CategorizeEvent(log types.Log, abi abi.ABI, instance *pawningShop.Contracts, env *config.Env) {
	incommingEventHash := log.Topics[0]
	fmt.Println("incomming event hash: ", incommingEventHash)
	switch incommingEventHash {
	case Hash(PawnCreatedSignature):
		PawnCreatedHandler(log, abi, instance, env.API_HOST)
	case Hash(PawnCancelledSignature):
		PawnCancelledHandler(log, abi, instance, env.API_HOST)
	case Hash(WhiteListAddedSignature):
		WhiteListAddedHandler(log, abi, instance)
	case Hash(WhiteListRemovedSignature):
		WhiteListRemovedHander(log, abi, instance)
	}
}

func Hash(signature string) common.Hash {
	return crypto.Keccak256Hash([]byte(signature))
}
