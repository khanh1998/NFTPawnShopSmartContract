package handler

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	PawnCreatedName      = "PawnCreated"
	PawnCancelledName    = "PawnCancelled"
	WhiteListAddedName   = "WhiteListAdded"
	WhiteListRemovedName = "WhiteListRemoved"
	BidCreatedName       = "BidCreated"
	BidCancelledName     = "BidCancelled"
	BidAcceptedName      = "BidAccepted"
)

const (
	PawnCreatedSignature      = "PawnCreated(address,uint256)"
	PawnCancelledSignature    = "PawnCancelled(address,uint256)"
	WhiteListAddedSignature   = "WhiteListAdded(address)"
	WhiteListRemovedSignature = "WhiteListRemoved(address)"
	BidCreatedNameSignature   = "BidCreated(address,uint256,uint256)"
	BidCancelledNameSignature = "BidCancelled(address,uint256,uint256)"
	BidAcceptedNameSignature  = "BidAccepted(uint256,uint256)"
)

func UnpackEvent(contractAbi abi.ABI, eventName string, data []byte) []string {
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
