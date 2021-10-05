package main

import (
	"log"
	"math/big"

	"khanh/config"
	pawningShop "khanh/contracts"
	"khanh/handler"
	"khanh/httpClient"
	"khanh/rabbitmq"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	env, err := config.LoadEnv()
	log.Println(env)
	if err != nil {
		log.Panic(err)
	}
	rabbit, err := rabbitmq.NewRabbitMQ(env.RABBIT_MQ_URI)
	if err != nil {
		log.Panic(err)
	}
	err = rabbit.SerializeAndSend("notification", "this is a test message1")
	err = rabbit.SerializeAndSend("notification", "this is a test message2")
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
	pawnRepaidChannelErr, err := instance.WatchPawnRepaid(nil, pawnRepaidChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	pawnLiquidatedChannel := make(chan *pawningShop.ContractsPawnLiquidated)
	pawnLiquidatedChannelErr, err := instance.WatchPawnLiquidated(nil, pawnLiquidatedChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	pawnCreatedChannel := make(chan *pawningShop.ContractsPawnCreated)
	pawnCreatedChannelErr, err := instance.WatchPawnCreated(nil, pawnCreatedChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	pawnCancelledChannel := make(chan *pawningShop.ContractsPawnCancelled)
	pawnCancelledChannelErr, err := instance.WatchPawnCancelled(nil, pawnCancelledChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	bidCreatedChannel := make(chan *pawningShop.ContractsBidCreated)
	bidCreatedChannelErr, err := instance.WatchBidCreated(nil, bidCreatedChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	bidCancelledChannel := make(chan *pawningShop.ContractsBidCancelled)
	bidCancelledChannelErr, err := instance.WatchBidCancelled(nil, bidCancelledChannel, []*big.Int{})
	if err != nil {
		log.Panic(err)
	}

	bidAcceptedChannel := make(chan *pawningShop.ContractsBidAccepted)
	bidAcceptedChannelErr, err := instance.WatchBidAccepted(nil, bidAcceptedChannel, []*big.Int{})
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

	bidHandler := handler.NewBidHandler(
		instance, myClient, rabbit, "notification",
	)
	pawnHandler := handler.NewPawnHandler(
		instance, myClient, rabbit, "notification",
	)
	whiteListHandler := handler.NewWhiteListHandler(
		instance, myClient, rabbit, "notification",
	)

	var data interface{}
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
			data = pawnHandler.PawnRepaid(repay)
			rabbit.SerializeAndSend("notification", data)
		case liquidated := <-pawnLiquidatedChannel:
			data = pawnHandler.PawnLiquidated(liquidated)
			rabbit.SerializeAndSend("notification", data)
		case pawnCreated := <-pawnCreatedChannel:
			data = pawnHandler.PawnCreated(pawnCreated)
			rabbit.SerializeAndSend("notification", data)
		case pawnCancelled := <-pawnCancelledChannel:
			data = pawnHandler.PawnCancelled(pawnCancelled)
			rabbit.SerializeAndSend("notification", data)
		case bidCreated := <-bidCreatedChannel:
			data = bidHandler.BidCreated(bidCreated)
			rabbit.SerializeAndSend("notification", data)
		case bidCancelled := <-bidCancelledChannel:
			data = bidHandler.BidCancelled(bidCancelled)
			rabbit.SerializeAndSend("notification", data)
		case bidAccepted := <-bidAcceptedChannel:
			data = bidHandler.BidAccepted(bidAccepted)
			rabbit.SerializeAndSend("notification", data)
		case whiteListAdded := <-whiteListAddedChannel:
			data = whiteListHandler.WhiteListAdded(whiteListAdded.SmartContract)
			rabbit.SerializeAndSend("notification", data)
		case whiteListRemoved := <-whiteListRemovedChannel:
			data = whiteListHandler.WhiteListRemoved(whiteListRemoved.SmartContract)
			rabbit.SerializeAndSend("notification", data)
		}
	}

}
