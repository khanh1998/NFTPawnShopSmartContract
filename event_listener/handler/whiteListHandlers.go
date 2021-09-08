package handler

import (
	"fmt"
	"khanh/httpClient"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func WhiteListAdded(SmartContract common.Address, client *httpClient.Client) {
	fmt.Println(WhiteListAddedName)
	success, _ := client.Notify.SendNotification(httpClient.Notification{
		Code:    WhiteListAddedName,
		Message: "A new smart contract is add to white list",
		Payload: SmartContract.String(),
	})
	log.Println("to notify", WhiteListAddedName, success)
}

func WhiteListRemoved(smartContract common.Address, client *httpClient.Client) {
	fmt.Println(WhiteListRemovedName)
	success, _ := client.Notify.SendNotification(httpClient.Notification{
		Code:    WhiteListRemovedName,
		Message: "A new smart contract is removed from white list",
		Payload: smartContract.String(),
	})
	log.Println("to notify", WhiteListRemovedName, success)
}
