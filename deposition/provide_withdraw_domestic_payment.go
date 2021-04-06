package deposition

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func ProvideWithdrawDomesticPayment(transferEvent *LogInvokeFiatWithdrawEvent) *AddPaymentResponse {

	pDR := PaymentDomesticRequest{
		Sender:
		&SenderPisDomestic{AccountNumber:
		&AccountNumber{Value: "OPERATOR_BANK_ACCOUNT_NUMBER"}},
		Recipient:
		&RecipientPis{AccountNumber:
		&AccountNumber{Value: transferEvent.BankAccountNumber}},
		TransferData:
		&TransferDataCurrencyRequired{
			Currency: "PLN", Description: "Deposition", Amount: transferEvent.Amount.String(), ExecutionDate: time.Now().Format("2021-01-16"),
		},
	}
	jsonData, err := json.Marshal(pDR)

	if err != nil {
		log.Fatal(err)
	}

	//SendNotification("Send Fiat between bank account")

     resp := ProvideOpenBankingApi(jsonData)
	var res AddPaymentResponse

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &res)

	return &res
}