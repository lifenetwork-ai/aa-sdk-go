# Go SDK for Account Abstraction

This repository contains a Go SDK implementation for Account Abstraction.

Key features:

- [x] Support entrypoint v0.7.0
- [x] Simple account factory
- [x] Paymaster data encoding and signing
- [x] Handle atomic ops support
- [x] Generic transaction builder for any calldata

# Example

## Setup

```go
config := &aasdk.Config{
		NodeUrl:             "https://testnet-lifeaiv1-c648f.avax-test.network/ext/bc/62fkxYTWbGBfXoHNXcGJbq2dTXba2uoCFySzdHy87iovJj2F4/rpc?token=25e957a027b09bb006da7e9fc981100ce25f333cd998a76eb36a842fcb5ba63a",
		BundlerUrl:          "http://34.126.118.65:8080/rpc",
		WaitReceiptInterval: 2 * time.Second,
		Entrypoint:          common.HexToAddress("0xd308aE59cb31932E8D9305BAda32Fa782d3D5d42"),
		AccountFactory:      common.HexToAddress("0xD421D8470b577f6A64992132D04906EfE51F1dE3"),
		PaymasterAddress:    common.HexToAddress("0xe7db0C105Ac75A493B0413046417e48594360542"),
		VerifyingSigner:     verifyingSigner, // optional, only used for verifying paymaster data
		ExecutorSigner:      executorSigner,  // optional, only used for sending atomic transactions
	}
```

To setup the config, you need to provide the following:

- NodeUrl: The URL of the node to send the transaction to.
- BundlerUrl: The URL of the bundler to send the transaction to.
- WaitReceiptInterval: The interval to wait for the receipt of the transaction.
- Entrypoint: The address of the entrypoint contract.
- AccountFactory: The address of the account factory contract.
- PaymasterAddress: The address of the paymaster contract. <optional>
- VerifyingSigner: The address of the verifying signer. <optional>
- ExecutorSigner: The address of the executor signer. <optional>

## Transfer Example

```go
calldata, err := aasdk.PackTransferData(client.AccountABI(), target, amount)
	if err != nil {
		log.Fatalf("Failed to pack transfer data: %v", err)
	}
	userOp := aasdk.NewUserOpWithDefault(sender, calldata, salt)

	tx, err := client.SendUserOp(context.Background(), userOp, signer)
	if err != nil {
		log.Fatalf("Failed to send user operation: %v", err)
	}
```

For a blockchain call, we will need the calldata of the function we want to call, basically contains of function signature and parameters.

Here, we provide the native support for transfer operation data packing, to get the `calldata` for transfer.

> For any smart contract call, the calldata can be packed by the in the `abigen` itself.
