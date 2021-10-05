package handler

import (
	"fmt"
	pawningShop "khanh/contracts"
	"khanh/httpClient"
	"khanh/rabbitmq"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

type WhiteListHandler struct {
	instance    *pawningShop.Contracts
	client      *httpClient.Client
	queue       *rabbitmq.RabbitMQ
	Q           *rabbitmq.RabbitMQ
	channelName string
}

func NewWhiteListHandler(instance *pawningShop.Contracts, client *httpClient.Client, rabbit *rabbitmq.RabbitMQ, channelName string) *WhiteListHandler {
	return &WhiteListHandler{
		instance: instance,
		client:   client,
		queue:    rabbit,
		Q:        rabbit,
	}
}

func (w *WhiteListHandler) WhiteListAdded(SmartContract common.Address) interface{} {
	fmt.Println(WhiteListAddedName)
	w.queue.SerializeAndSend(w.channelName, WhiteListAddedName)
	data := httpClient.Notification{
		Code:    WhiteListAddedName,
		Message: "A new smart contract is add to white list",
		Payload: SmartContract.String(),
	}
	success, _ := w.client.Notify.SendNotification(data)
	log.Println("to notify", WhiteListAddedName, success)
	// err := w.queue.SerializeAndSend(w.channelName, data)
	// if err != nil {
	// 	log.Panic(err)
	// } else {
	// 	log.Println("to notify rabbitmq", BidCreatedName, success)
	// }
	return data
}

func (w *WhiteListHandler) WhiteListRemoved(smartContract common.Address) interface{} {
	fmt.Println(WhiteListRemovedName)
	w.queue.SerializeAndSend(w.channelName, WhiteListRemovedName)
	data := httpClient.Notification{
		Code:    WhiteListRemovedName,
		Message: "A new smart contract is removed from white list",
		Payload: smartContract.String(),
	}
	success, _ := w.client.Notify.SendNotification(data)
	log.Println("to notify", WhiteListRemovedName, success)
	// err := w.queue.SerializeAndSend(w.channelName, data)
	// if err != nil {
	// 	log.Panic(err)
	// } else {
	// 	log.Println("to notify rabbitmq", BidCreatedName, success)
	// }
	return data
}
