package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	pawningShop "khanh/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x4A3f07Cb18ed442EbF2Db20419f8196e52885f10")
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

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			event, err := contractAbi.Unpack("WhiteListAdded", vLog.Data)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(event)
		}
	}

}
