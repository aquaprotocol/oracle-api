# oracle-api

The Oracle provider for the BankFunctionOracle.sol smart contract.

### Event listener:
The most important events listeners is defined for:

#### [InvokeFiatDepositEvent](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L10)
This event is invokes by [invokeFiatDeposit](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L34) function in BankFunctionOracle.sol.

```go
event InvokeFiatDepositEvent(
    uint256 amount,
    string bankAccountNumber,
    uint256 id
  );
  ```
  1. When the event is catched, domestic payment is involved
  2. Request to [open-banking-api](https://github.com/aquaprotocol/open-banking-api) is sending
  3. Response from [open-banking-api](https://github.com/aquaprotocol/open-banking-api) is forwarded to [SetFiatDepositResult](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L56) in BankFunctionOracle.sol contract.
  4. From [SetFiatDepositResult](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L56) function in BankFunctionOracle.sol [callbackDepositeFiat](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L65) is executed.
 

#### [InvokeFiatWithdrawEvent](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L16)
This event is invokes by [invokeFiatWithdraw](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L45) function in BankFunctionOracle.sol.

```go
 event InvokeFiatWithdrawEvent(
    uint256 amount,
    string bankAccountNumber, 
    uint256 id
  );
  ```
  1. When the event is catched, domestic payment is involved
  2. Request to [open-banking-api](https://github.com/aquaprotocol/open-banking-api) is sending
  3. Response from [open-banking-api](https://github.com/aquaprotocol/open-banking-api) is forwarded to [SetFiatWithdrawResult](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L69) in BankFunctionOracle.sol contract.
  4. From [SetFiatDepositResult](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L69) function in BankFunctionOracle.sol [callbackWithdrawFiat](https://github.com/aquaprotocol/smart-contract/blob/a2ab5b59a261337ab4536362d3c5af4fe8ae77a1/BankFunctionOracle.sol#L78) is executed.

#### Run service
`go build main.go`

