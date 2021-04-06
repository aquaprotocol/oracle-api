package main

import (
	con "./connector"
	contract "./contract"
	deposition "./deposition"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
)

func main() {
	getEvent()
}

func getEvent() {

	client := con.ClientFactory()

	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	invokeFiatWithdrawEventSig := []byte("InvokeFiatWithdrawEvent(uint256,string,uint256)")
	invokeFiatWithdrawEventSigHash := crypto.Keccak256Hash(invokeFiatWithdrawEventSig)

	invokeFiatDepositEventSig := []byte("InvokeFiatDepositEvent(uint256,string,uint256)")
	invokeFiatDepositEventSigHash := crypto.Keccak256Hash(invokeFiatDepositEventSig)

	addressNew := common.HexToAddress(ORACLE_CALLER_ADDRESS)


	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("In case:")
			fmt.Println(vLog.TxHash.Hex()) // pointer to event log
			contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))

			if err != nil {
				log.Fatal(err)
			}

			switch vLog.Topics[0].Hex() {
			case invokeFiatWithdrawEventSigHash.Hex():
				fmt.Println("InvokeFiatWithdrawEventSigHash")
				var transferEvent deposition.LogInvokeFiatWithdrawEvent

				transactOpts := con.WalletFactory(client)

				err = contractAbi.UnpackIntoInterface(&transferEvent, "InvokeFiatWithdrawEvent", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Bank Account Number: %s\n", transferEvent.BankAccountNumber)
				fmt.Printf("Nonce: %s\n", transactOpts.Nonce.String())
				fmt.Printf("Tokens: %s\n", transferEvent.Id.String())
				fmt.Printf("Amount: %s\n", transferEvent.Amount.String())

		



				res := deposition.ProvideWithdrawDomesticPayment(&transferEvent)
				fmt.Println(&res)


				//////////////////////////////////////////


				tx, err := instance.SetFiatWithdrawResult(transactOpts, "true", addressNew, transferEvent.Id)

		
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("tx sent: %s", tx.Hash().Hex()) 

				fmt.Println("after ChainID")

			case invokeFiatDepositEventSigHash.Hex():

				fmt.Println("invokeFiatDepositEventSigHash")
				var transferEvent deposition.LogInvokeFiatDepositEvent

				transactOpts := con.WalletFactory(client)

				err = contractAbi.UnpackIntoInterface(&transferEvent, "InvokeFiatDepositEvent", vLog.Data)
				if err != nil {
					fmt.Println("BLAD")
					log.Fatal(err)
				}

				fmt.Printf("Bank Account Number: %s\n", transferEvent.BankAccountNumber)
				fmt.Printf("Nonce: %s\n", transactOpts.Nonce.String())
				fmt.Printf("Tokens: %s\n", transferEvent.Id.String())
				fmt.Printf("Amount: %s\n", transferEvent.Amount.String())

			
				/////////////////////////



				res := deposition.ProvideDepositDomesticPayment(&transferEvent)
				fmt.Println(&res)

				tx, err := instance.SetFiatDepositResult(transactOpts, "true", addressNew, transferEvent.Id)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("tx sent: %s", tx.Hash().Hex())

				fmt.Println("after ChainID")

			}

		}
	}

}
