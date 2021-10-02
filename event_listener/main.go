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
	err = rabbit.Send("test", []byte("this is a test message"))
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
			pawnHandler.PawnRepaid(repay)
		case liquidated := <-pawnLiquidatedChannel:
			pawnHandler.PawnLiquidated(liquidated)
		case pawnCreated := <-pawnCreatedChannel:
			pawnHandler.PawnCreated(pawnCreated)
		case pawnCancelled := <-pawnCancelledChannel:
			pawnHandler.PawnCancelled(pawnCancelled)
		case bidCreated := <-bidCreatedChannel:
			bidHandler.BidCreated(bidCreated)
		case bidCancelled := <-bidCancelledChannel:
			bidHandler.BidCancelled(bidCancelled)
		case bidAccepted := <-bidAcceptedChannel:
			bidHandler.BidAccepted(bidAccepted)
		case whiteListAdded := <-whiteListAddedChannel:
			whiteListHandler.WhiteListAdded(whiteListAdded.SmartContract)
		case whiteListRemoved := <-whiteListRemovedChannel:
			whiteListHandler.WhiteListRemoved(whiteListRemoved.SmartContract)
		}
	}

}
