package deposition

import "math/big"

//event GetLatestEthPriceEvent(address callerAddress, uint id);
type LogInvokeFiatWithdrawEvent struct {
	Amount             *big.Int
	BankAccountNumber string
	Id                 *big.Int
}

//event InvokeFiatDepositEvent(uint256 amount, string bankAccountNumber, uint id);
type LogInvokeFiatDepositEvent struct {
	Amount             *big.Int
	BankAccountNumber string
	Id                 *big.Int
}