package main

import (
	"log"

	"khanh/config"
	pawningShop "khanh/contracts"
	"khanh/handler"
	"khanh/httpClient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	instance, _ := pawningShop.NewContracts(common.HexToAddress(env.CONTRACT_ADDRESS), client)
	if err != nil {
		log.Panic(err)
	}

	myClient := httpClient.NewClient(
		env.API_HOST, env.PAWN_PATH, env.BID_PATH, env.BID_PAWN_PATH, env.NOTIFY_HOST, env.NOTIFICATION_PATH,
	)

	pawnRepaidChannel := make(chan *pawningShop.ContractsPawnRepaid)
	pawnRepaidChannelErr, err := instance.WatchPawnRepaid(nil, pawnRepaidChannel)
	if err != nil {
		log.Panic(err)
	}

	pawnLiquidatedChannel := make(chan *pawningShop.ContractsPawnLiquidated)
	pawnLiquidatedChannelErr, err := instance.WatchPawnLiquidated(nil, pawnLiquidatedChannel)
	if err != nil {
		log.Panic(err)
	}

	pawnCreatedChannel := make(chan *pawningShop.ContractsPawnCreated)
	pawnCreatedChannelErr, err := instance.WatchPawnCreated(nil, pawnCreatedChannel)
	if err != nil {
		log.Panic(err)
	}

	pawnCancelledChannel := make(chan *pawningShop.ContractsPawnCancelled)
	pawnCancelledChannelErr, err := instance.WatchPawnCancelled(nil, pawnCancelledChannel)
	if err != nil {
		log.Panic(err)
	}

	bidCreatedChannel := make(chan *pawningShop.ContractsBidCreated)
	bidCreatedChannelErr, err := instance.WatchBidCreated(nil, bidCreatedChannel)
	if err != nil {
		log.Panic(err)
	}

	bidCancelledChannel := make(chan *pawningShop.ContractsBidCancelled)
	bidCancelledChannelErr, err := instance.WatchBidCancelled(nil, bidCancelledChannel)
	if err != nil {
		log.Panic(err)
	}

	bidAcceptedChannel := make(chan *pawningShop.ContractsBidAccepted)
	bidAcceptedChannelErr, err := instance.WatchBidAccepted(nil, bidAcceptedChannel)
	if err != nil {
		log.Panic(err)
	}

	whiteListAddedChannel := make(chan *pawningShop.ContractsWhiteListAdded)
	whiteListAddedChannelErr, err := instance.WatchWhiteListAdded(nil, whiteListAddedChannel)
	if err != nil {
		log.Panic(err)
	}

	whiteListRemovedChannel := make(chan *pawningShop.ContractsWhiteListAdded)
	whiteListRemovedChannelErr, err := instance.WatchWhiteListAdded(nil, whiteListRemovedChannel)
	if err != nil {
		log.Panic(err)
	}

	log.Println("started to listen to ", env.CONTRACT_ADDRESS)

	for {
		select {
		case err := <-pawnRepaidChannelErr.Err():
			log.Panic(err)
		case err := <-pawnLiquidatedChannelErr.Err():
			log.Panic(err)
		case err := <-pawnCreatedChannelErr.Err():
			log.Panic(err)
		case err := <-pawnCancelledChannelErr.Err():
			log.Panic(err)
		case err := <-bidCreatedChannelErr.Err():
			log.Panic(err)
		case err := <-bidCancelledChannelErr.Err():
			log.Panic(err)
		case err := <-bidAcceptedChannelErr.Err():
			log.Panic(err)
		case err := <-whiteListAddedChannelErr.Err():
			log.Panic(err)
		case err := <-whiteListRemovedChannelErr.Err():
			log.Panic(err)

		case repay := <-pawnRepaidChannel:
			handler.PawnRepaid(repay.PawnId, myClient)
		case liquidated := <-pawnLiquidatedChannel:
			handler.PawnLiquidated(liquidated.PawnId, myClient)
		case pawnCreated := <-pawnCreatedChannel:
			handler.PawnCreated(pawnCreated.PawnId, instance, myClient)
		case pawnCancelled := <-pawnCancelledChannel:
			handler.PawnCancelled(pawnCancelled.PawnId, instance, myClient)
		case bidCreated := <-bidCreatedChannel:
			handler.BidCreated(bidCreated.BidId, bidCreated.PawnId, instance, myClient)
		case bidCancelled := <-bidCancelledChannel:
			handler.BidCancelled(bidCancelled.BidId, bidCancelled.Creator, instance, myClient)
		case bidAccepted := <-bidAcceptedChannel:
			handler.BidAccepted(bidAccepted.BidId, instance, myClient)
		case whiteListAdded := <-whiteListAddedChannel:
			handler.WhiteListAdded(whiteListAdded.SmartContract, myClient)
		case whiteListRemoved := <-whiteListRemovedChannel:
			handler.WhiteListRemoved(whiteListRemoved.SmartContract, myClient)
		}
	}

}
