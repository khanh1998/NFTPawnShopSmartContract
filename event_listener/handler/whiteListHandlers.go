package handler

import (
	"fmt"
	pawningShop "khanh/contracts"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func WhiteListAdded(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
	fmt.Println(WhiteListAddedName)
	data := UnpackEvent(abi, WhiteListAddedName, vlog.Data)
	fmt.Println(data)
	address, err := instance.WhiteListNFT(nil, big.NewInt(0))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(address)
}

func WhiteListRemoved(vlog types.Log, abi abi.ABI, instance *pawningShop.Contracts) {
	fmt.Println(WhiteListRemovedName)
	data := UnpackEvent(abi, WhiteListRemovedName, vlog.Data)
	fmt.Println(data)
	instance.WhiteListNFT(nil, big.NewInt(0))
	address, err := instance.WhiteListNFT(nil, big.NewInt(0))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(address)
}
